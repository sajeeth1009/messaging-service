package smtp_client

import (
	"log"
	"net/smtp"

	"github.com/jordan-wright/email"
)

type SmtpClients struct {
	servers        SmtpServerList
	connectionPool []email.Pool
	counter        int
}

func NewSmtpClients(configFile string) (*SmtpClients, error) {
	serverList := SmtpServerList{}
	if err := serverList.ReadFromFile(configFile); err != nil {
		return nil, err
	}

	sc := &SmtpClients{
		servers:        serverList,
		counter:        0,
		connectionPool: initConnectionPool(serverList),
	}
	return sc, nil
}

func initConnectionPool(serverList SmtpServerList) []email.Pool {
	connectionPools := []email.Pool{}
	for _, server := range serverList.Servers {
		//Set number of concurrent connections here
		auth := smtp.PlainAuth(
			"",
			server.AuthData.Username,
			server.AuthData.Password,
			server.Host,
		)
		var pool, err = email.NewPool(server.Address(), server.Connections, auth)
		if err != nil {
			log.Print("Error setting up connection pool for: " + server.Address())
			continue
		} else {
			connectionPools = append(connectionPools, *pool)
		}
	}
	if len(connectionPools) < 1 {
		log.Fatal("no smtp server connection in the pool")
	}
	return connectionPools
}
