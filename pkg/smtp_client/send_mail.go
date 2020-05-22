package smtp_client

import (
	"errors"
	"log"
	"net/smtp"
	"strings"
)

func (sc *SmtpClients) SendMail(
	to []string,
	fromAddressOverride string,
	fromNameOverride string,
	subject string,
	content string,
) error {
	sc.counter += 1
	if len(sc.servers.Servers) < 1 {
		return errors.New("no servers defined")
	}
	selectedServerIndex := sc.counter % len(sc.servers.Servers)
	selectedServer := sc.servers.Servers[selectedServerIndex]

	auth := smtp.PlainAuth(
		"",
		selectedServer.AuthData.Username,
		selectedServer.AuthData.Password,
		selectedServer.Address(),
	)

	from := sc.servers.FromAddr
	if len(fromAddressOverride) > 0 {
		from = fromAddressOverride
	}

	fromName := sc.servers.FromName
	if len(fromNameOverride) > 0 {
		fromName = fromAddressOverride
	}

	message := "To: " + strings.Join(to, ",") + "\r\n"
	message += "From: \"" + fromName + "\" " + from + "\r\n"
	message += "Subject: " + subject + "\r\n"
	message += "\r\n"
	message += content

	log.Println(auth)
	log.Println(message)
	/*
		if err := smtp.SendMail(selectedServer.Address(), auth, from, []string{"hevesi.peti@gmail.com"}, []byte(message)); err != nil {
			fmt.Println("Error SendMail: ", err)
		}*/
	return nil
}
