package email_client_service

import (
	"context"
	"testing"

	"github.com/influenzanet/go-utils/pkg/testutils"
	api "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
)

func TestSendEmailEndpoint(t *testing.T) {
	s := emailClientServer{
		StmpClients: nil,
	}

	t.Run("with missing payload", func(t *testing.T) {
		_, err := s.SendEmail(context.Background(), nil)
		ok, msg := testutils.ShouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with empty payload", func(t *testing.T) {
		_, err := s.SendEmail(context.Background(), &api.SendEmailReq{})
		ok, msg := testutils.ShouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with empty to list", func(t *testing.T) {
		_, err := s.SendEmail(context.Background(), &api.SendEmailReq{
			Subject: "test",
			Content: "hello",
			HeaderOverrides: &api.HeaderOverrides{
				From: `"Test" <test@test.de>`,
			},
		})
		ok, msg := testutils.ShouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})
}
