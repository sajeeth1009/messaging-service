package bulk_messages

import (
	"context"
	"encoding/base64"
	"io"
	"log"
	"time"

	emailAPI "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	"github.com/influenzanet/messaging-service/pkg/dbs/messagedb"
	"github.com/influenzanet/messaging-service/pkg/templates"
	"github.com/influenzanet/messaging-service/pkg/types"
	umAPI "github.com/influenzanet/user-management-service/pkg/api"
)

func AsyncSendToAllUsers(
	apiClients *types.APIClients,
	messageDBService *messagedb.MessageDBService,
	instanceID string,
	messageTemplate types.EmailTemplate,
) {

	stream, err := apiClients.UserManagementService.StreamUsers(context.Background(), &umAPI.StreamUsersMsg{InstanceId: instanceID})
	if err != nil {
		log.Printf("AsyncSendToAllUsers: %v", err)
		return
	}
	for {
		user, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("%v.AsyncSendToAllUsers(_) = _, %v", apiClients.UserManagementService, err)
			break
		}

		outgoingEmail := types.OutgoingEmail{
			MessageType:     messageTemplate.MessageType,
			HeaderOverrides: messageTemplate.HeaderOverrides,
		}
		contentInfos := map[string]string{}

		if user.Account.Type == "email" {
			outgoingEmail.To = []string{user.Account.AccountId}
		}

		if user.Account.AccountConfirmedAt < 1 {
			log.Println("message is not sent, if account not confirmed")
			continue
		}
		if messageTemplate.MessageType == "newsletter" {
			if !user.ContactPreferences.SubscribedToNewsletter {
				// user does not want to get newsletter
				continue
			}

			outgoingEmail.To = getEmailsByIds(user.ContactInfos, user.ContactPreferences.SendNewsletterTo)

			token, err := getUnsubscribeToken(apiClients.UserManagementService, instanceID, user)
			if err != nil {
				log.Printf("AsyncSendToAllUsers: %v", err)
				continue
			}
			contentInfos["unsubscribeToken"] = token
		} else if messageTemplate.MessageType == "study-reminder" {
			token, err := getTemploginToken(apiClients.UserManagementService, instanceID, user, messageTemplate.StudyKey, 604800)
			if err != nil {
				log.Printf("AsyncSendToAllUsers: %v", err)
				continue
			}
			contentInfos["loginToken"] = token
		}

		subject, content, err := prepareContent(messageTemplate, user.Account.PreferredLanguage, contentInfos)
		if err != nil {
			log.Printf("AsyncSendToAllUsers: %v", err)
			continue
		}

		outgoingEmail.Subject = subject
		outgoingEmail.Content = content

		go sendMail(
			apiClients.EmailClientService,
			instanceID,
			messageDBService,
			outgoingEmail,
		)
	}
}

func AsyncSendToStudyParticipants(apiClients *types.APIClients) {
	// define async methods to fetch users, check study states and trigger email sending here
	// don't send to unconfirmed emails
	// generate tempLogin, and unsubscribe tokens
}

func getEmailsByIds(contacts []*umAPI.ContactInfo, ids []string) []string {
	emails := []string{}
	for _, c := range contacts {
		if c.Type == "email" {
			for _, id := range ids {
				if c.Id == id {
					emails = append(emails, c.GetEmail())
				}
			}
		}
	}
	return emails
}

func prepareContent(temp types.EmailTemplate, prefLang string, contentInfos map[string]string) (subject string, content string, err error) {
	translation := templates.GetTemplateTranslation(temp, prefLang)
	subject = translation.Subject
	decodedTemplate, err := base64.StdEncoding.DecodeString(translation.TemplateDef)
	if err != nil {
		return "", "", err
	}

	// execute template
	content, err = templates.ResolveTemplate(
		temp.MessageType+prefLang,
		string(decodedTemplate),
		contentInfos,
	)
	return
}

func sendMail(
	emailClient emailAPI.EmailClientServiceApiClient,
	instanceID string,
	messageDBService *messagedb.MessageDBService,
	mail types.OutgoingEmail,
) {
	_, err := emailClient.SendEmail(context.Background(), &emailAPI.SendEmailReq{
		To:              mail.To,
		HeaderOverrides: mail.HeaderOverrides.ToEmailClientAPI(),
		Subject:         mail.Subject,
		Content:         mail.Content,
	})
	if err != nil {
		log.Printf("Error when sending: %v", err)
		_, errS := messageDBService.AddToOutgoingEmails(instanceID, mail)
		if errS != nil {
			log.Printf("Error when saving to outgoing: %v", err)
		}
		return
	}

	_, err = messageDBService.AddToSentEmails(instanceID, mail)
	if err != nil {
		log.Printf("Error when saving to sent: %v", err)
	}
}

func getTemploginToken(
	userClient umAPI.UserManagementApiClient,
	instanceID string,
	user *umAPI.User,
	studyKey string,
	expiresIn int64,
) (token string, err error) {
	resp, err := userClient.GenerateTempToken(context.Background(), &umAPI.TempTokenInfo{
		UserId:     user.Id,
		Purpose:    "survey-login",
		InstanceId: instanceID,
		Expiration: time.Now().Unix() + expiresIn,
		Info: map[string]string{
			"studyKey": studyKey,
		},
	})
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}

func getUnsubscribeToken(
	userClient umAPI.UserManagementApiClient,
	instanceID string,
	user *umAPI.User,
) (token string, err error) {
	resp, err := userClient.GenerateTempToken(context.Background(), &umAPI.TempTokenInfo{
		UserId:     user.Id,
		Purpose:    "unsubscribe-newsletter",
		InstanceId: instanceID,
		Expiration: time.Now().Unix() + 157680000,
	})
	if err != nil {
		return "", err
	}
	return resp.Token, nil
}
