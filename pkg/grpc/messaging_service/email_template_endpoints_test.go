package messaging_service

import (
	"context"
	"testing"

	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
)

func TestGetEmailTemplatesEndpoint(t *testing.T) {
	s := messagingServer{
		messageDBservice: testMessageDBService,
	}
	t.Run("without payload", func(t *testing.T) {
		_, err := s.GetEmailTemplates(context.Background(), nil)
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with empty payload", func(t *testing.T) {
		_, err := s.GetEmailTemplates(context.Background(), &api.GetEmailTemplatesReq{})
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with valid arguments", func(t *testing.T) {
		t.Error("test unimplemented")
	})
}

func TestSaveEmailTemplateEndpoint(t *testing.T) {
	s := messagingServer{
		messageDBservice: testMessageDBService,
	}
	t.Run("without payload", func(t *testing.T) {
		_, err := s.SaveEmailTemplate(context.Background(), nil)
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with empty payload", func(t *testing.T) {
		_, err := s.SaveEmailTemplate(context.Background(), &api.SaveEmailTemplateReq{})
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with new template without study", func(t *testing.T) {
		t.Error("test unimplemented")
	})

	t.Run("with new template with study", func(t *testing.T) {
		t.Error("test unimplemented")
	})

	t.Run("with existing template without study", func(t *testing.T) {
		t.Error("test unimplemented")
	})

	t.Run("with existing template with study", func(t *testing.T) {
		t.Error("test unimplemented")
	})
}

func TestDeleteEmailTemplateEndpoint(t *testing.T) {
	s := messagingServer{
		messageDBservice: testMessageDBService,
	}
	t.Run("without payload", func(t *testing.T) {
		_, err := s.GetEmailTemplates(context.Background(), nil)
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with empty payload", func(t *testing.T) {
		_, err := s.GetEmailTemplates(context.Background(), &api.GetEmailTemplatesReq{})
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with not existing template", func(t *testing.T) {
		t.Error("test unimplemented")
	})

	t.Run("with existing template but wrong study", func(t *testing.T) {
		t.Error("test unimplemented")
	})

	t.Run("with existing template without study", func(t *testing.T) {
		t.Error("test unimplemented")
	})

	t.Run("with existing template with study", func(t *testing.T) {
		t.Error("test unimplemented")
	})
}
