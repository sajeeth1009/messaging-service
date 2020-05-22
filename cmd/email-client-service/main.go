package main

import (
	"context"
	"log"
	"os"

	"github.com/influenzanet/messaging-service/pkg/grpc/email_client_service"
	sc "github.com/influenzanet/messaging-service/pkg/smtp_client"
)

// Config is the structure that holds all global configuration data
type Config struct {
	Port             string
	ServerConfigPath string
}

func initConfig() Config {
	conf := Config{}
	conf.Port = os.Getenv("EMAIL_CLIENT_SERVICE_LISTEN_PORT")
	conf.ServerConfigPath = os.Getenv("SMTP_SERVER_CONFIG_PATH")
	return conf
}

func main() {
	conf := initConfig()

	smtpClients, err := sc.NewSmtpClients(conf.ServerConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	if err := email_client_service.RunServer(
		ctx,
		conf.Port,
		smtpClients,
	); err != nil {
		log.Fatal(err)
	}
}
