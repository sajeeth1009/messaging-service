package email_client_service

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	api "github.com/influenzanet/messaging-service/pkg/api/email_client_service"
	"google.golang.org/grpc"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
)

type emailClientServer struct {
}

// NewEmailClientServiceServer creates a new service instance
func NewEmailClientServiceServer() api.EmailClientServiceApiServer {
	return &emailClientServer{}
}

// RunServer runs gRPC service to publish ToDo service
func RunServer(
	ctx context.Context, port string,
) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// register service
	server := grpc.NewServer()
	api.RegisterEmailClientServiceApiServer(server, NewEmailClientServiceServer())

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")
			server.GracefulStop()
			<-ctx.Done()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")
	log.Println("wait connections on port " + port)
	return server.Serve(lis)
}
