package email_client_service

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	api "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const maxRetry = 3

func (s *emailClientServer) Status(ctx context.Context, _ *empty.Empty) (*api.ServiceStatus, error) {
	return &api.ServiceStatus{
		Status:  api.ServiceStatus_NORMAL,
		Msg:     "service running",
		Version: apiVersion,
	}, nil
}

func (s *emailClientServer) SendEmail(ctx context.Context, req *api.SendEmailReq) (*api.ServiceStatus, error) {
	if req == nil || len(req.To) < 1 {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}

	retryCounter := 0
	for {
		if err := s.StmpClients.SendMail(
			req.To,
			req.FromAddress,
			req.FromName,
			req.Subject,
			req.Content,
		); err != nil {
			retryCounter += 1
			if retryCounter >= maxRetry {
				return nil, status.Error(codes.Internal, err.Error())
			}
			log.Printf("SendEmail: %v", err)
		} else {
			break
		}
	}

	return &api.ServiceStatus{
		Version: apiVersion,
		Status:  api.ServiceStatus_NORMAL,
		Msg:     "email sent",
	}, nil
}
