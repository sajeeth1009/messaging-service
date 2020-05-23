package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type OutgoingEmail struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	MessageType string             `bson:"messageType"`
	To          []string           `bson:"to"`
	FromAddress string             `bson:"fromAddress"`
	FromName    string             `bson:"fromName"`
	Subject     string             `bson:"subject"`
	Content     string             `bson:"content"`
	AddedAt     int64              `bson:"addedAt"`
}
