package messaging_service

import (
	"context"

	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *messagingServer) GetEmailTemplates(ctx context.Context, req *api.GetEmailTemplatesReq) (*api.EmailTemplates, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *messagingServer) SaveEmailTemplate(ctx context.Context, req *api.SaveEmailTemplateReq) (*api.EmailTemplate, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (s *messagingServer) DeleteEmailTemplate(ctx context.Context, req *api.DeleteEmailTemplateReq) (*api.ServiceStatus, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
