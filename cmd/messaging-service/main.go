package main

import (
	"context"
	"log"
	"os"

	"github.com/influenzanet/messaging-service/internal/config"
	"github.com/influenzanet/messaging-service/pkg/dbs/messagedb"
	"github.com/influenzanet/messaging-service/pkg/types"

	gc "github.com/influenzanet/messaging-service/pkg/grpc/clients"
	"github.com/influenzanet/messaging-service/pkg/grpc/messaging_service"
)

// Config is the structure that holds all global configuration data
type Config struct {
	Port            string
	MessageDBConfig types.DBConfig
	ServiceURLs     struct {
		UserManagementService string
		EmailClientService string
	}
}

func initConfig() Config {
	conf := Config{}
	conf.Port = os.Getenv("MESSAGING_SERVICE_LISTEN_PORT")
	conf.ServiceURLs.UserManagementService = os.Getenv("ADDR_USER_MANAGEMENT_SERVICE")
	conf.ServiceURLs.EmailClientService = os.Getenv("ADDR_EMAIL_CLIENT_SERVICE")
	conf.MessageDBConfig = config.GetMessageDBConfig()
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
	clients. = emailClient
	// <---

	messageDBService := messagedb.NewMessageDBService(conf.MessageDBConfig)

	ctx := context.Background()

	if err := messaging_service.RunServer(
		ctx,
		conf.Port,
		clients,
		messageDBService,
	); err != nil {
		log.Fatal(err)
	}
}
