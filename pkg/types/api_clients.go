package types

import (
	emailAPI "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	umAPI "github.com/influenzanet/user-management-service/pkg/api"
)

type APIClients struct {
	UserManagementService umAPI.UserManagementApiClient
	EmailClientService    emailAPI.EmailClientServiceApiClient
}
