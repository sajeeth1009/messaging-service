package smtp_client

import (
	"errors"
	"log"
)

func (sc *SmtpClients) SendMail(to []string, content string) error {
	if len(sc.servers.Servers) < 1 {
		return errors.New("no servers defined")
	}
	selectedServerIndex := sc.counter % len(sc.servers.Servers)

	selectedServer := sc.servers.Servers[selectedServerIndex]
	log.Println(selectedServer)
	sc.counter += 1
	return errors.New("not implmented")
}
