package messaging_service

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	emailAPI "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
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

func (s *messagingServer) SendInstantEmail(ctx context.Context, req *api.SendEmailReq) (*api.ServiceStatus, error) {
	if req == nil || req.InstanceId == "" || len(req.To) < 1 || req.MessageType == "" {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}

	templateDef, err := s.messageDBservice.FindEmailTemplateByType(req.InstanceId, req.MessageType, req.StudyKey)
	if err != nil {
		return nil, status.Error(codes.Internal, "template not found")
	}

	translation := templates.GetTemplateTranslation(templateDef, req.PreferredLanguage)

	// execute template
	content, err := templates.ResolveTemplate(
		req.InstanceId+req.MessageType+req.PreferredLanguage,
		translation.TemplateDef,
		req.ContentInfos,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, "content could not be generated")
	}

	outgoingEmail := types.OutgoingEmail{
		MessageType: req.MessageType,
		To:          req.To,
		FromName:    templateDef.FromName,
		FromAddress: templateDef.FromAddress,
		Subject:     translation.Subject,
		Content:     content,
	}

	_, err = s.clients.EmailClientService.SendEmail(ctx, &emailAPI.SendEmailReq{
		To:          outgoingEmail.To,
		FromName:    outgoingEmail.FromName,
		FromAddress: outgoingEmail.FromAddress,
		Subject:     outgoingEmail.Subject,
		Content:     content,
	})
	if err != nil {
		_, errS := s.messageDBservice.AddToOutgoingEmails(req.InstanceId, outgoingEmail)
		log.Printf("Saving to outgoing: %v", errS)
		return &api.ServiceStatus{
			Version: apiVersion,
			Msg:     "failed sending message, added to outgoind",
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
