package types

type EmailTemplate struct {
	MessageType  string `bson:"messageType"`
	StudyKey     string `bson:"studyKey,omitempty"`
	Translations string `bson:"translations"`
}

type LocalizedTemplate struct {
	Lang        string `bson:"languageCode"`
	TemplateDef string `bson:"templateDef"`
}
