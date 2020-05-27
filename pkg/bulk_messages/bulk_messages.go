package bulk_messages

import (
	"errors"

	"github.com/influenzanet/messaging-service/pkg/types"
	umAPI "github.com/influenzanet/user-management-service/pkg/api"
)

func AsynSendToAllUsers(apiClients *types.APIClients) {
	// define async methods to fetch users and trigger email sending here
	// don't send to unconfirmed emails
	// generate tempLogin, and unsubscribe tokens

}

func AsyncSendToStudyParticipants(apiClients *types.APIClients) {
	// define async methods to fetch users, check study states and trigger email sending here
	// don't send to unconfirmed emails
	// generate tempLogin, and unsubscribe tokens
}

func getTemploginToken(
	userClient umAPI.UserManagementApiClient,
	instanceID string,
	user *umAPI.User,
) (token string, err error) {

	return "", errors.New("unimplemented")
}

func getUnsubscribeToken(
	userClient umAPI.UserManagementApiClient,
	instanceID string,
	user *umAPI.User,
) (token string, err error) {
	return "", errors.New("unimplemented")
}
