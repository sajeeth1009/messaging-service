package messaging_service

import (
	"context"

	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"github.com/influenzanet/messaging-service/pkg/types"
	"github.com/influenzanet/messaging-service/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *messagingServer) GetEmailTemplates(ctx context.Context, req *api.GetEmailTemplatesReq) (*api.EmailTemplates, error) {
	if req == nil || utils.IsTokenEmpty(req.Token) {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}

	if !utils.CheckIfAnyRolesInToken(req.Token, []string{"RESEARCHER", "ADMIN"}) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	templates, err := s.messageDBservice.FindAllEmailTempates(req.Token.InstanceId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &api.EmailTemplates{
		Templates: make([]*api.EmailTemplate, len(templates)),
	}
	for i, v := range templates {
		resp.Templates[i] = v.ToAPI()
	}
	return resp, nil
}

func (s *messagingServer) SaveEmailTemplate(ctx context.Context, req *api.SaveEmailTemplateReq) (*api.EmailTemplate, error) {
	if req == nil || utils.IsTokenEmpty(req.Token) || req.Template.MessageType == "" {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}
	if !utils.CheckIfAnyRolesInToken(req.Token, []string{"RESEARCHER", "ADMIN"}) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	templ := types.EmailTemplateFromAPI(req.Template)
	templ, err := s.messageDBservice.SaveEmailTemplate(req.Token.InstanceId, templ)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return templ.ToAPI(), nil
}

func (s *messagingServer) DeleteEmailTemplate(ctx context.Context, req *api.DeleteEmailTemplateReq) (*api.ServiceStatus, error) {
	if req == nil || utils.IsTokenEmpty(req.Token) || req.MessageType == "" {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}
	if !utils.CheckIfAnyRolesInToken(req.Token, []string{"RESEARCHER", "ADMIN"}) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	err := s.messageDBservice.DeleteEmailTemplate(req.Token.InstanceId, req.MessageType, req.StudyKey)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.ServiceStatus{
		Status: api.ServiceStatus_NORMAL,
		Msg:    "template deleted",
	}, nil
}
