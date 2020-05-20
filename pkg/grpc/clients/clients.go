package clients

import (
	"log"

	emailAPI "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	umAPI "github.com/influenzanet/user-management-service/pkg/api"
	"google.golang.org/grpc"
)

func connectToGRPCServer(addr string) *grpc.ClientConn {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to connect to %s: %v", addr, err)
	}
	return conn
}

func ConnectToUserManagementService(addr string) (client umAPI.UserManagementApiClient, close func() error) {
	serverConn := connectToGRPCServer(addr)
	return umAPI.NewUserManagementApiClient(serverConn), serverConn.Close
}

func ConnectToEmailClientService(addr string) (client emailAPI.EmailClientServiceApiClient, close func() error) {
	serverConn := connectToGRPCServer(addr)
	return emailAPI.NewEmailClientServiceApiClient(serverConn), serverConn.Close
}
