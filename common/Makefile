.DEFAULT_TARGET := help

ROOTDIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Protobuf

.PHONY: protoc
protoc: ## Install protobuf compiler.
	@brew install -q protobuf
	@echo "Installed protoc"

	@go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0
	@echo "Installed protoc-gen-go"

	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0
	@echo "Installed protoc-gen-grpc"

.PHONY: proto
proto: ## Compile proto files into go language implementation.
	@protoc \
		--go_out=. \
		--go_opt=paths=import \
		--go-grpc_out=. \
		--go-grpc_opt=paths=import \
		protos/*.proto
