package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/influenzanet/messaging-service/internal/config"
	emailAPI "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	"github.com/influenzanet/messaging-service/pkg/bulk_messages"
	"github.com/influenzanet/messaging-service/pkg/dbs/globaldb"
	"github.com/influenzanet/messaging-service/pkg/dbs/messagedb"
	gc "github.com/influenzanet/messaging-service/pkg/grpc/clients"
	"github.com/influenzanet/messaging-service/pkg/types"
)

// Config is the structure that holds all global configuration data
type Config struct {
	Frequencies struct {
		HighPrio    int
		LowPrio     int
		AutoMessage int
	}
	MessageDBConfig types.DBConfig
	GlobalDBConfig  types.DBConfig
	ServiceURLs     struct {
		UserManagementService string
		EmailClientService    string
		StudyService          string
	}
}

func initConfig() Config {
	conf := Config{}
	hp, err := strconv.Atoi(os.Getenv("MESSAGE_SCHEDULER_INTERVAL_HIGH_PRIO"))
	if err != nil {
		log.Fatal(err)
	}
	lp, err := strconv.Atoi(os.Getenv("MESSAGE_SCHEDULER_INTERVAL_LOW_PRIO"))
	if err != nil {
		log.Fatal(err)
	}
	am, err := strconv.Atoi(os.Getenv("MESSAGE_SCHEDULER_INTERVAL_AUTO_MESSAGE"))
	if err != nil {
		log.Fatal(err)
	}

	conf.Frequencies = struct {
		HighPrio    int
		LowPrio     int
		AutoMessage int
	}{
		HighPrio:    hp,
		LowPrio:     lp,
		AutoMessage: am,
	}
	conf.ServiceURLs.UserManagementService = os.Getenv("ADDR_USER_MANAGEMENT_SERVICE")
	conf.ServiceURLs.StudyService = os.Getenv("ADDR_STUDY_SERVICE")
	conf.ServiceURLs.EmailClientService = os.Getenv("ADDR_EMAIL_CLIENT_SERVICE")
	conf.MessageDBConfig = config.GetMessageDBConfig()
	conf.GlobalDBConfig = config.GetGlobalDBConfig()
	return conf
}

func main() {
	conf := initConfig()

	// ---> client connections
	clients := &types.APIClients{}
	umClient, close := gc.ConnectToUserManagementService(conf.ServiceURLs.UserManagementService)
	defer close()
	clients.UserManagementService = umClient

	emailClient, close := gc.ConnectToEmailClientService(conf.ServiceURLs.EmailClientService)
	defer close()
	clients.EmailClientService = emailClient

	studyClient, close := gc.ConnectToStudyService(conf.ServiceURLs.StudyService)
	defer close()
	clients.StudyService = studyClient
	// <---

	messageDBService := messagedb.NewMessageDBService(conf.MessageDBConfig)
	globalDBService := globaldb.NewGlobalDBService(conf.GlobalDBConfig)

	go runnerForLowPrioOutgoingEmails(messageDBService, globalDBService, clients, conf.Frequencies.LowPrio)
	go runnerForAutoMessages(messageDBService, globalDBService, clients, conf.Frequencies.AutoMessage)
	runnerForHighPrioOutgoingEmails(messageDBService, globalDBService, clients, conf.Frequencies.HighPrio)
}

func runnerForHighPrioOutgoingEmails(mdb *messagedb.MessageDBService, gdb *globaldb.GlobalDBService, clients *types.APIClients, freq int) {
	olderThan := int64(float64(freq) * 0.8)
	for {
		log.Println("Fetch and send high prio outgoing emails.")
		go handleOutgoingEmails(mdb, gdb, clients, olderThan, true)
		time.Sleep(time.Duration(freq) * time.Second)
	}
}

func runnerForLowPrioOutgoingEmails(mdb *messagedb.MessageDBService, gdb *globaldb.GlobalDBService, clients *types.APIClients, freq int) {
	olderThan := int64(float64(freq) * 0.8)
	for {
		log.Println("Fetch and send low prio outgoing emails.")
		go handleOutgoingEmails(mdb, gdb, clients, olderThan, false)
		time.Sleep(time.Duration(freq) * time.Second)
	}
}

func runnerForAutoMessages(mdb *messagedb.MessageDBService, gdb *globaldb.GlobalDBService, clients *types.APIClients, freq int) {
	for {
		log.Println("Fetch and send scheduled bulk messages.")
		go handleAutoMessages(mdb, gdb, clients)
		time.Sleep(time.Duration(freq) * time.Second)
	}
}

func handleOutgoingEmails(mdb *messagedb.MessageDBService, gdb *globaldb.GlobalDBService, clients *types.APIClients, olderThan int64, onlyHighPrio bool) {
	instances, err := gdb.GetAllInstances()
	if err != nil {
		log.Printf("handleOutgoingEmails.GetAllInstances: %v", err)
	}
	for _, instance := range instances {
		emails, err := mdb.FetchOutgoingEmails(instance.InstanceID, 500, olderThan, onlyHighPrio)
		if err != nil {
			log.Printf("handleOutgoingEmails.FetchOutgoingEmails for %s: %v", instance.InstanceID, err)
			continue
		}
		if len(emails) < 1 {
			continue
		}
		log.Printf("%d outgoing emails found in instance %s", len(emails), instance.InstanceID)
		for _, email := range emails {
			_, err := clients.EmailClientService.SendEmail(context.Background(), &emailAPI.SendEmailReq{
				To:              email.To,
				HeaderOverrides: email.HeaderOverrides.ToEmailClientAPI(),
				Subject:         email.Subject,
				Content:         email.Content,
				HighPrio:        email.HighPrio,
			})
			if err != nil {
				log.Printf("Could not send email in instance %s: %v", instance.InstanceID, err)
				continue
			}

			_, err = mdb.AddToSentEmails(instance.InstanceID, email)
			if err != nil {
				log.Printf("Error while saving to sent: %v", err)
				continue
			}
			err = mdb.DeleteOutgoingEmail(instance.InstanceID, email.ID.Hex())
			if err != nil {
				log.Printf("Error while deleting outgoing: %v", err)
			}
		}
	}
}

func handleAutoMessages(mdb *messagedb.MessageDBService, gdb *globaldb.GlobalDBService, clients *types.APIClients) {
	instances, err := gdb.GetAllInstances()
	if err != nil {
		log.Printf("handleAutoMessages.GetAllInstances: %v", err)
	}
	for _, instance := range instances {
		activeMessages, err := mdb.FindAutoMessages(instance.InstanceID, true)
		if err != nil {
			log.Printf("handleAutoMessages.FindAutoMessages for %s: %v", instance.InstanceID, err)
			continue
		}
		if len(activeMessages) < 1 {
			continue
		}

		for _, messageDef := range activeMessages {
			switch messageDef.Type {
			case "all-users":
				go bulk_messages.SendToAllUsers(
					clients,
					mdb,
					instance.InstanceID,
					messageDef.Template,
				)
			case "study-participants":
				messageDef.Template.StudyKey = messageDef.StudyKey
				go bulk_messages.SendToStudyParticipants(
					clients,
					mdb,
					instance.InstanceID,
					messageDef.Template,
					messageDef.Condition.ToAPI(),
				)
			default:
				log.Printf("handleAutoMessages: message type unknown: %s", messageDef.Type)
			}

			messageDef.NextTime += messageDef.Period
			_, err := mdb.SaveAutoMessage(instance.InstanceID, messageDef)
			if err != nil {
				log.Printf("handleAutoMessages.SaveAutoMessage for %s: %v", instance.InstanceID, err)
				continue
			}
		}
	}

}
