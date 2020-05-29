package messaging_service

import (
	"context"

	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"github.com/influenzanet/messaging-service/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *messagingServer) GetAutoMessages(ctx context.Context, req *api.GetAutoMessagesReq) (*api.AutoMessages, error) {
	if req == nil || utils.IsTokenEmpty(req.Token) {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}

	if !utils.CheckIfAnyRolesInToken(req.Token, []string{"RESEARCHER", "ADMIN"}) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	return nil, status.Error(codes.Unimplemented, "unimplemented")
	/*
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
		return resp, nil*/
}

func (s *messagingServer) SaveAutoMessage(ctx context.Context, req *api.SaveAutoMessageReq) (*api.AutoMessage, error) {
	if req == nil || utils.IsTokenEmpty(req.Token) || req.AutoMessage == nil {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}
	if !utils.CheckIfAnyRolesInToken(req.Token, []string{"RESEARCHER", "ADMIN"}) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	return nil, status.Error(codes.Unimplemented, "unimplemented")
	/*templ := types.EmailTemplateFromAPI(req.Template)
	templ, err := s.messageDBservice.SaveEmailTemplate(req.Token.InstanceId, templ)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return templ.ToAPI(), nil*/
}

func (s *messagingServer) DeleteAutoMessage(ctx context.Context, req *api.DeleteAutoMessageReq) (*api.ServiceStatus, error) {
	if req == nil || utils.IsTokenEmpty(req.Token) || req.AutoMessageId == "" {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}
	if !utils.CheckIfAnyRolesInToken(req.Token, []string{"RESEARCHER", "ADMIN"}) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	return nil, status.Error(codes.Unimplemented, "unimplemented")
	/*err := s.messageDBservice.DeleteEmailTemplate(req.Token.InstanceId, req.MessageType, req.StudyKey)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.ServiceStatus{
		Status: api.ServiceStatus_NORMAL,
		Msg:    "template deleted",
	}, nil*/
}
