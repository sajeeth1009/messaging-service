.PHONY: test api

PROTO_BUILD_DIR = ../../..
PROTO_TARGET = ./pkg/api
DOCKER_OPTS ?= --rm

#TEST_ARGS = -v

# VERSION := $(shell git describe --tags --abbrev=0)

help:
	@echo "Service building targets"
	@echo "  test  : run test suites"
	@echo "  api: compile protobuf files for go"
	@echo "Env:"
	@echo "  DOCKER_OPTS : default docker build options (default : $(DOCKER_OPTS))"
	@echo "  TEST_ARGS : Arguments to pass to go test call"

api:
	find "$(PROTO_TARGET)" -type f -delete
	find ./api/*.proto -maxdepth 1 -type f -exec protoc {} --go_out=plugins=grpc:$(PROTO_BUILD_DIR) \;

build:
	go build .

test:
	./test/test.sh $(TEST_ARGS)
