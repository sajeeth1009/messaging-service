package smtp_client

import (
	"errors"
	"net/textproto"
	"time"

	"github.com/influenzanet/messaging-service/pkg/types"
	"github.com/jordan-wright/email"
)

func (sc *SmtpClients) SendMail(
	to []string,
	subject string,
	htmlContent string,
	overrides *types.HeaderOverrides,
) error {
	sc.counter += 1
	if len(sc.connectionPool) < 1 {
		return errors.New("no servers defined")
	}

	index := sc.counter % len(sc.connectionPool)
	selectedServer := sc.connectionPool[index]

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

	e := &email.Email{
		To:      to,
		From:    From,
		Sender:  Sender,
		ReplyTo: ReplyTo,
		Subject: subject,
		HTML:    []byte(htmlContent),
		Headers: textproto.MIMEHeader{},
	}
	return selectedServer.Send(e, time.Second*15)
}
