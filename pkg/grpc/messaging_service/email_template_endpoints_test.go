package messaging_service

import (
	"context"
	"testing"

	"github.com/influenzanet/go-utils/pkg/api_types"
	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"github.com/influenzanet/messaging-service/pkg/types"
)

func TestGetEmailTemplatesEndpoint(t *testing.T) {
	s := messagingServer{
		messageDBservice: testMessageDBService,
	}

	_, err := s.messageDBservice.SaveEmailTemplate(testInstanceID, types.EmailTemplate{MessageType: "B"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	_, err = s.messageDBservice.SaveEmailTemplate(testInstanceID, types.EmailTemplate{MessageType: "A"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
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
		resp, err := s.GetEmailTemplates(context.Background(), &api.GetEmailTemplatesReq{
			Token: &api_types.TokenInfos{
				Id:         "uid",
				InstanceId: testInstanceID,
				Payload: map[string]string{
					"roles":    "PARTICIPANT,RESEARCHER",
					"username": "testuser",
				},
			},
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		if len(resp.Templates) != 2 {
			t.Errorf("unexpected number of templates: %d", len(resp.Templates))
			return
		}
	})
}

func TestSaveEmailTemplateEndpoint(t *testing.T) {
	s := messagingServer{
		messageDBservice: testMessageDBService,
	}

	userToken := &api_types.TokenInfos{
		Id:         "uid",
		InstanceId: testInstanceID,
		Payload: map[string]string{
			"roles":    "PARTICIPANT,RESEARCHER",
			"username": "testuser",
		},
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
		_, err := s.SaveEmailTemplate(context.Background(), &api.SaveEmailTemplateReq{
			Token: userToken,
			Template: &api.EmailTemplate{
				MessageType:     "test",
				DefaultLanguage: "de",
			},
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	})

	t.Run("with new template with study", func(t *testing.T) {
		_, err := s.SaveEmailTemplate(context.Background(), &api.SaveEmailTemplateReq{
			Token: userToken,
			Template: &api.EmailTemplate{
				MessageType:     "test",
				StudyKey:        "testStudy",
				DefaultLanguage: "en",
			},
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	})

	t.Run("with existing template without study", func(t *testing.T) {
		_, err := s.SaveEmailTemplate(context.Background(), &api.SaveEmailTemplateReq{
			Token: userToken,
			Template: &api.EmailTemplate{
				MessageType:     "test",
				DefaultLanguage: "de",
			},
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	})

	t.Run("with existing template with study", func(t *testing.T) {
		_, err := s.SaveEmailTemplate(context.Background(), &api.SaveEmailTemplateReq{
			Token: userToken,
			Template: &api.EmailTemplate{
				MessageType:     "test",
				StudyKey:        "testStudy",
				DefaultLanguage: "de",
			},
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
	})
}

func TestDeleteEmailTemplateEndpoint(t *testing.T) {
	s := messagingServer{
		messageDBservice: testMessageDBService,
	}
	userToken := &api_types.TokenInfos{
		Id:         "uid",
		InstanceId: testInstanceID,
		Payload: map[string]string{
			"roles":    "PARTICIPANT,RESEARCHER",
			"username": "testuser",
		},
	}

	_, err := s.messageDBservice.SaveEmailTemplate(testInstanceID, types.EmailTemplate{MessageType: "B"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	_, err = s.messageDBservice.SaveEmailTemplate(testInstanceID, types.EmailTemplate{MessageType: "A", StudyKey: "al"})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	t.Run("without payload", func(t *testing.T) {
		_, err := s.DeleteEmailTemplate(context.Background(), nil)
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with empty payload", func(t *testing.T) {
		_, err := s.DeleteEmailTemplate(context.Background(), &api.DeleteEmailTemplateReq{})
		ok, msg := shouldHaveGrpcErrorStatus(err, "missing argument")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with not existing template", func(t *testing.T) {
		_, err := s.DeleteEmailTemplate(context.Background(), &api.DeleteEmailTemplateReq{})
		ok, msg := shouldHaveGrpcErrorStatus(err, "")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with existing template but wrong study", func(t *testing.T) {
		_, err := s.DeleteEmailTemplate(context.Background(), &api.DeleteEmailTemplateReq{})
		ok, msg := shouldHaveGrpcErrorStatus(err, "")
		if !ok {
			t.Error(msg)
		}
	})

	t.Run("with existing template without study", func(t *testing.T) {
		_, err := s.DeleteEmailTemplate(context.Background(), &api.DeleteEmailTemplateReq{
			Token:       userToken,
			MessageType: "test",
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}
		_, err = s.messageDBservice.FindEmailTemplateByType(testInstanceID, "test", "")
		if err == nil {
			t.Error("should return error")
			return
		}
	})

	t.Run("with existing template with study", func(t *testing.T) {
		_, err := s.DeleteEmailTemplate(context.Background(), &api.DeleteEmailTemplateReq{
			Token:       userToken,
			MessageType: "test",
			StudyKey:    "testStudy",
		})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		_, err = s.messageDBservice.FindEmailTemplateByType(testInstanceID, "test", "testStudy")
		if err == nil {
			t.Error("should return error")
			return
		}
	})
}
