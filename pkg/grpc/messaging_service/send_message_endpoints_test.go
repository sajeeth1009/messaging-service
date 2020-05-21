package messaging_service

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"github.com/influenzanet/messaging-service/pkg/types"
	emailMock "github.com/influenzanet/messaging-service/test/mocks/email-client-service"
)

func TestSendInstantEmailEndpoint(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockEmailClient := emailMock.NewMockEmailClientServiceApiClient(mockCtrl)

	s := messagingServer{
		messageDBservice: testMessageDBService,
		clients: &types.APIClients{
			EmailClientService: mockEmailClient,
		},
	}

	t.Run("without payload", func(t *testing.T) {
		_, err := s.SendInstantEmail(context.Background(), nil)
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with empty payload", func(t *testing.T) {
		_, err := s.SendInstantEmail(context.Background(), &api.SendEmailReq{})
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with not existing template", func(t *testing.T) {
		_, err := s.SendInstantEmail(context.Background(), &api.SendEmailReq{
			InstanceId:  testInstanceID,
			To:          []string{"test@test.test"},
			MessageType: "wrong",
		})
		ok, msg := shouldHaveGrpcErrorStatus(err, "template not found")
		if !ok {
			t.Error(msg)
		}
	})

	t.Error("add test template")
	t.Error("test for sending failed") // sending failed (should save to outgoing)
	t.Error("test unimplemented")      // sending succeeded
}
