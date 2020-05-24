package types

import (
	api "github.com/influenzanet/messaging-service/pkg/api/messaging_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EmailTemplate struct {
	ID              primitive.ObjectID  `bson:"_id,omitempty"`
	MessageType     string              `bson:"messageType"`
	StudyKey        string              `bson:"studyKey,omitempty"`
	DefaultLanguage string              `bson:"defaultLanguage"`
	FromName        string              `bson:"fromName"`
	FromAddress     string              `bson:"fromAddress"`
	Translations    []LocalizedTemplate `bson:"translations"`
}

type LocalizedTemplate struct {
	Lang        string `bson:"languageCode"`
	Subject     string `bson:"subject"`
	TemplateDef string `bson:"templateDef"`
}

func EmailTemplateFromAPI(obj *api.EmailTemplate) EmailTemplate {
	if obj == nil {
		return EmailTemplate{}
	}
	_id, _ := primitive.ObjectIDFromHex(obj.Id)
	translations := make([]LocalizedTemplate, len(obj.Translations))
	for i, t := range obj.Translations {
		translations[i] = LocalizedTemplateFromAPI(t)
	}
	return EmailTemplate{
		ID:              _id,
		MessageType:     obj.MessageType,
		StudyKey:        obj.StudyKey,
		DefaultLanguage: obj.DefaultLanguage,
		FromName:        obj.FromName,
		FromAddress:     obj.FromAddress,
		Translations:    translations,
	}
}

// ToAPI converts a email template object from DB format into the API format
func (obj EmailTemplate) ToAPI() *api.EmailTemplate {
	translations := make([]*api.LocalizedTemplate, len(obj.Translations))
	for i, t := range obj.Translations {
		translations[i] = t.ToAPI()
	}
	return &api.EmailTemplate{
		Id:              obj.ID.Hex(),
		MessageType:     obj.MessageType,
		StudyKey:        obj.StudyKey,
		DefaultLanguage: obj.DefaultLanguage,
		FromAddress:     obj.FromAddress,
		FromName:        obj.FromName,
		Translations:    translations,
	}
}

func LocalizedTemplateFromAPI(obj *api.LocalizedTemplate) LocalizedTemplate {
	if obj == nil {
		return LocalizedTemplate{}
	}
	return LocalizedTemplate{
		Lang:        obj.Lang,
		Subject:     obj.Subject,
		TemplateDef: obj.TemplateDef,
	}
}

// ToAPI converts a localized template from DB format into the API format
func (obj LocalizedTemplate) ToAPI() *api.LocalizedTemplate {
	return &api.LocalizedTemplate{
		Lang:        obj.Lang,
		Subject:     obj.Subject,
		TemplateDef: obj.TemplateDef,
	}
}
