package email_client_service

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	api "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *emailClientServer) Status(ctx context.Context, _ *empty.Empty) (*api.ServiceStatus, error) {
	return &api.ServiceStatus{
		Status:  api.ServiceStatus_NORMAL,
		Msg:     "service running",
		Version: apiVersion,
	}, nil
}

func (s *emailClientServer) SendEmail(ctx context.Context, req *api.SendEmailReq) (*api.ServiceStatus, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
