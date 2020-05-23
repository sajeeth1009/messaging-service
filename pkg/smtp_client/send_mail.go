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
		selectedServer.Host,
	)

	from := sc.servers.FromAddr
	if len(fromAddressOverride) > 0 {
		from = fromAddressOverride
	}

	fromName := sc.servers.FromName
	if len(fromNameOverride) > 0 {
		fromName = fromNameOverride
	}

	message := "To: " + strings.Join(to, ",") + "\r\n"
	message += "From: \"" + fromName + "\" <" + from + ">\r\n"
	message += "Subject: " + subject + "\r\n"
	message += "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	message += content
	log.Println(message)

	return smtp.SendMail(selectedServer.Address(), auth, from, to, []byte(message))
}
