package messaging_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	emailAPI "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
)

func (s *messagingServer) Status(ctx context.Context, _ *empty.Empty) (*api.ServiceStatus, error) {
	return &api.ServiceStatus{
		Status:  api.ServiceStatus_NORMAL,
		Msg:     "service running",
		Version: apiVersion,
	}, nil
}

func (s *messagingServer) SendInstantEmail(ctx context.Context, req *api.SendEmailReq) (*api.ServiceStatus, error) {
	// TODO: check inputs
	// TODO: find and execute template
	_, err := s.clients.EmailClientService.SendEmail(ctx, &emailAPI.SendEmailReq{
		To:      []string{"test"},
		Content: "testcontent",
	})

	return nil, err
}
