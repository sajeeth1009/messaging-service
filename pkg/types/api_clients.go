package types

import (
	loggingAPI "github.com/influenzanet/logging-service/pkg/api"
	emailAPI "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	studyAPI "github.com/influenzanet/study-service/pkg/api"
	umAPI "github.com/influenzanet/user-management-service/pkg/api"
)

type APIClients struct {
	UserManagementService umAPI.UserManagementApiClient
	StudyService          studyAPI.StudyServiceApiClient
	EmailClientService    emailAPI.EmailClientServiceApiClient
	LoggingService        loggingAPI.LoggingServiceApiClient
}
