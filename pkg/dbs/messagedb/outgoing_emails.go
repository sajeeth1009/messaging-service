package messagedb

import (
	"github.com/influenzanet/messaging-service/pkg/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (dbService *MessageDBService) AddToOutgoingEmails(instanceID string, email types.OutgoingEmail) (types.OutgoingEmail, error) {
	ctx, cancel := dbService.getContext()
	defer cancel()

	res, err := dbService.collectionRefOutgoingEmails(instanceID).InsertOne(ctx, email)
	if err != nil {
		return email, err
	}
	email.ID = res.InsertedID.(primitive.ObjectID)
	return email, nil
}

func (dbService *MessageDBService) FetchOutgoingEmails(instanceID string, amount int) (emails []types.OutgoingEmail, err error) {
	ctx, cancel := dbService.getContext()
	defer cancel()

	counter := 0
	for counter < amount {
		var newEmail types.OutgoingEmail
		filter := bson.M{}
		if err := dbService.collectionRefOutgoingEmails(instanceID).FindOneAndDelete(ctx, filter).Decode(&newEmail); err != nil {
			break
		}
		emails = append(emails, newEmail)
		counter += 1
	}
	return emails, nil
}
