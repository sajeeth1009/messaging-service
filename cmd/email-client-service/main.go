package main

import (
	"context"
	"log"
	"os"

	"github.com/influenzanet/messaging-service/pkg/grpc/email_client_service"
)

// Config is the structure that holds all global configuration data
type Config struct {
	Port string
}

func initConfig() Config {
	conf := Config{}
	conf.Port = os.Getenv("EMAIL_CLIENT_SERVICE_LISTEN_PORT")
	return conf
}

func main() {
	conf := initConfig()

	ctx := context.Background()
	if err := email_client_service.RunServer(
		ctx,
		conf.Port,
	); err != nil {
		log.Fatal(err)
	}
}
