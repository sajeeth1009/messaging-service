package smtp_client

import (
	"errors"
	"net/textproto"
	"time"

	"github.com/jordan-wright/email"
)

func (sc *SmtpClients) SendMail(
	to []string,
	fromAddressOverride string,
	fromNameOverride string,
	subject string,
	htmlContent string,
) error {
	sc.counter += 1
	if len(sc.connectionPool) < 1 {
		return errors.New("no servers defined")
	}

	index := sc.counter % len(sc.connectionPool)
	selectedServer := sc.connectionPool[index]

	e := &email.Email{
		To:      to,
		From:    formatFrom(sc.servers.FromAddr, sc.servers.FromName, fromAddressOverride, fromNameOverride),
		Subject: subject,
		HTML:    []byte(htmlContent),
		Headers: textproto.MIMEHeader{},
	}
	return selectedServer.Send(e, time.Second*15)
}

func formatFrom(defaultAddr string, defaultName string, overrideAddr string, overrideName string) string {
	fromAddr := defaultAddr
	if len(overrideAddr) > 0 {
		fromAddr = overrideAddr
	}

	fromName := defaultName
	if len(overrideName) > 0 {
		fromName = overrideName
	}
	from := fromAddr
	if len(fromName) > 0 {
		from = fromName + " <" + fromAddr + ">"
	}
	return from
}
