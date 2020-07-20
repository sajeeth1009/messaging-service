package messaging_service

import (
	"context"

	"github.com/influenzanet/go-utils/pkg/token_checks"
	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"github.com/influenzanet/messaging-service/pkg/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *messagingServer) GetAutoMessages(ctx context.Context, req *api.GetAutoMessagesReq) (*api.AutoMessages, error) {
	if req == nil || token_checks.IsTokenEmpty(req.Token) {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}

	if !token_checks.CheckIfAnyRolesInToken(req.Token, []string{"RESEARCHER", "ADMIN"}) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	autoMessages, err := s.messageDBservice.FindAutoMessages(req.Token.InstanceId, false)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	resp := &api.AutoMessages{
		AutoMessages: make([]*api.AutoMessage, len(autoMessages)),
	}
	for i, v := range autoMessages {
		resp.AutoMessages[i] = v.ToAPI()
	}
	return resp, nil
}

func (s *messagingServer) SaveAutoMessage(ctx context.Context, req *api.SaveAutoMessageReq) (*api.AutoMessage, error) {
	if req == nil || token_checks.IsTokenEmpty(req.Token) || req.AutoMessage == nil {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}
	if !token_checks.CheckIfAnyRolesInToken(req.Token, []string{"RESEARCHER", "ADMIN"}) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}

	reqMsg := types.AutoMessageFromAPI(req.AutoMessage)
	autoMsg, err := s.messageDBservice.SaveAutoMessage(req.Token.InstanceId, *reqMsg)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return autoMsg.ToAPI(), nil
}

func (s *messagingServer) DeleteAutoMessage(ctx context.Context, req *api.DeleteAutoMessageReq) (*api.ServiceStatus, error) {
	if req == nil || token_checks.IsTokenEmpty(req.Token) || req.AutoMessageId == "" {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}
	if !token_checks.CheckIfAnyRolesInToken(req.Token, []string{"RESEARCHER", "ADMIN"}) {
		return nil, status.Error(codes.PermissionDenied, "permission denied")
	}
	err := s.messageDBservice.DeleteAutoMessage(req.Token.InstanceId, req.AutoMessageId)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &api.ServiceStatus{
		Status: api.ServiceStatus_NORMAL,
		Msg:    "auto message deleted",
	}, nil
}
