package types

import (
	umAPI "github.com/influenzanet/user-management-service/pkg/api"
)

type APIClients struct {
	UserManagementService umAPI.UserManagementApiClient
}
