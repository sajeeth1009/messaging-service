package smtp_client

import (
	"log"
	"net/smtp"
	"strings"
	"time"

	"github.com/influenzanet/messaging-service/pkg/types"
)

func (sc *SmtpClients) SendMail(
	to []string,
	subject string,
	htmlContent string,
	overrides *types.HeaderOverrides,
) error {
	sc.counter++

	index := sc.counter % len(sc.servers.Servers)

	From := sc.servers.From
	Sender := sc.servers.Sender
	ReplyTo := sc.servers.ReplyTo

	if overrides != nil {
		if overrides.From != "" {
			From = overrides.From
		}
		if overrides.Sender != "" {
			Sender = overrides.Sender
		}

		if overrides.NoReplyTo {
			ReplyTo = []string{}
		} else if len(overrides.ReplyTo) > 0 {
			ReplyTo = overrides.ReplyTo
		}
	}

	msg := []byte("To: " + strings.Join(to, " ") + " \r\n" +
		"From: " + From + " \r\n" +
		"Sender: " + Sender + " \r\n" +
		"ReplyTo: " + strings.Join(ReplyTo, " ") + " \r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		htmlContent + " \r\n")
	server := sc.servers.Servers[index]

	auth := smtp.PlainAuth(
		"",
		server.AuthData.Username,
		server.AuthData.Password,
		server.Host,
	)
	if server.AuthData.Username == "" && server.AuthData.Password == "" {
		auth = nil
	}

	start := time.Now().UnixNano()

	err := smtp.SendMail(server.Address(), auth, From, to, msg)
	log.Printf("Time taken to send message : %v ms", (time.Now().UnixNano()-start)/int64(time.Millisecond))

	if err != nil {
		log.Printf("error when trying to send email: %v", err)
	}
	return err
}
