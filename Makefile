.PHONY: test api mock docker-email-client docker-message-scheduler docker-messaging-service

PROTO_BUILD_DIR = ../../..
PROTO_TARGET = ./pkg/api
DOCKER_OPTS ?= --rm

#TEST_ARGS = -v
VERSION := $(shell git describe --tags --abbrev=0)

help:
	@echo "Service building targets"
	@echo "  test  : run test suites"
	@echo "  api: compile protobuf files for go"
	@echo "  mock: generate mockup Services for testing"
	@echo "  docker-email-client: build docker container for email_client_service"
	@echo "  docker-message-scheduler: build docker container for message scheduler"
	@echo "  docker-messaging-service: build docker container for messaging-service"
	@echo "Env:"
	@echo "  DOCKER_OPTS : default docker build options (default : $(DOCKER_OPTS))"
	@echo "  TEST_ARGS : Arguments to pass to go test call"

api:
	find "$(PROTO_TARGET)" -type f -delete
	find ./api/*.proto -maxdepth 1 -type f -exec protoc {} --go_out=plugins=grpc:$(PROTO_BUILD_DIR) \;

mock:
	mockgen -source=./pkg/api/email_client_service/email-client-service.pb.go EmailClientServiceApiClient > test/mocks/email-client-service/email_client_service.go
	mockgen github.com/influenzanet/user-management-service/pkg/api UserManagementApiClient,UserManagementApi_StreamUsersClient > test/mocks/user-management-service/user_management_service.go
	mockgen github.com/influenzanet/study-service/pkg/api StudyServiceApiClient > test/mocks/study-service/study_service.go

test:
	./test/test.sh $(TEST_ARGS)

docker-email-client:
	docker build -t  github.com/influenzanet/email-client-service:$(VERSION)  -f build/docker/email-client-service/Dockerfile $(DOCKER_OPTS) .

docker-message-scheduler:
	docker build -t  github.com/influenzanet/message-scheduler:$(VERSION)  -f build/docker/message-scheduler/Dockerfile $(DOCKER_OPTS) .

docker-messaging-service:
	docker build -t  github.com/influenzanet/messaging-service:$(VERSION)  -f build/docker/messaging-service/Dockerfile $(DOCKER_OPTS) .
