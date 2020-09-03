package messaging_service

import (
	"context"
	"encoding/base64"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/influenzanet/go-utils/pkg/token_checks"
	emailAPI "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"github.com/influenzanet/messaging-service/pkg/bulk_messages"
	"github.com/influenzanet/messaging-service/pkg/templates"
	"github.com/influenzanet/messaging-service/pkg/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *messagingServer) Status(ctx context.Context, _ *empty.Empty) (*api.ServiceStatus, error) {
	return &api.ServiceStatus{
		Status:  api.ServiceStatus_NORMAL,
		Msg:     "service running",
		Version: apiVersion,
	}, nil
}

func (s *messagingServer) SendMessageToAllUsers(ctx context.Context, req *api.SendMessageToAllUsersReq) (*api.ServiceStatus, error) {
	if req == nil || token_checks.IsTokenEmpty(req.Token) || req.Template == nil {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}

	if !token_checks.CheckIfAnyRolesInToken(req.Token, []string{"RESEARCHER", "ADMIN"}) {
		return nil, status.Error(codes.PermissionDenied, "no permission to send messages")
	}

	// use go method (don't wait for result since it can take long)
	go bulk_messages.AsyncSendToAllUsers(
		s.clients,
		s.messageDBservice,
		req.Token.InstanceId,
		types.EmailTemplateFromAPI(req.Template),
	)
	return &api.ServiceStatus{
		Msg:     "message sending triggered",
		Status:  api.ServiceStatus_NORMAL,
		Version: apiVersion,
	}, nil
}

func (s *messagingServer) SendMessageToStudyParticipants(ctx context.Context, req *api.SendMessageToStudyParticipantsReq) (*api.ServiceStatus, error) {
	if req == nil || token_checks.IsTokenEmpty(req.Token) || req.StudyKey == "" || req.Template == nil {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}
	if !token_checks.CheckIfAnyRolesInToken(req.Token, []string{"RESEARCHER", "ADMIN"}) {
		return nil, status.Error(codes.PermissionDenied, "no permission to send messages")
	}
	req.Template.StudyKey = req.StudyKey

	// use go method (don't wait for result since it can take long)
	go bulk_messages.AsyncSendToStudyParticipants(
		s.clients,
		s.messageDBservice,
		req.Token.InstanceId,
		types.EmailTemplateFromAPI(req.Template),
		req.Condition,
	)
	return &api.ServiceStatus{
		Msg:     "message sending triggered",
		Status:  api.ServiceStatus_NORMAL,
		Version: apiVersion,
	}, nil
}

func (s *messagingServer) SendInstantEmail(ctx context.Context, req *api.SendEmailReq) (*api.ServiceStatus, error) {
	if req == nil || req.InstanceId == "" || len(req.To) < 1 || req.MessageType == "" {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}

	templateDef, err := s.messageDBservice.FindEmailTemplateByType(req.InstanceId, req.MessageType, req.StudyKey)
	if err != nil {
		return nil, status.Error(codes.Internal, "template not found")
	}

	translation := templates.GetTemplateTranslation(templateDef, req.PreferredLanguage)

	decodedTemplate, err := base64.StdEncoding.DecodeString(translation.TemplateDef)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// execute template
	content, err := templates.ResolveTemplate(
		req.InstanceId+req.MessageType+req.PreferredLanguage,
		string(decodedTemplate),
		req.ContentInfos,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "content could not be generated")
	}

	outgoingEmail := types.OutgoingEmail{
		MessageType:     req.MessageType,
		To:              req.To,
		HeaderOverrides: templateDef.HeaderOverrides,
		Subject:         translation.Subject,
		Content:         content,
	}

	_, err = s.clients.EmailClientService.SendEmail(ctx, &emailAPI.SendEmailReq{
		To:              outgoingEmail.To,
		HeaderOverrides: outgoingEmail.HeaderOverrides.ToEmailClientAPI(),
		Subject:         outgoingEmail.Subject,
		Content:         content,
		HighPrio:        true,
	})
	if err != nil {
		_, errS := s.messageDBservice.AddToOutgoingEmails(req.InstanceId, outgoingEmail)
		if errS != nil {
			log.Printf("Error while saving to outgoing: %v", errS)
		}
		return &api.ServiceStatus{
			Version: apiVersion,
			Msg:     "failed sending message, added to outgoing",
			Status:  api.ServiceStatus_PROBLEM,
		}, nil
	}

	_, err = s.messageDBservice.AddToSentEmails(req.InstanceId, outgoingEmail)
	if err != nil {
		log.Printf("Saving to sent: %v", err)
	}

	return &api.ServiceStatus{
		Version: apiVersion,
		Msg:     "message sent",
		Status:  api.ServiceStatus_NORMAL,
	}, nil
}
