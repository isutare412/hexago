.DEFAULT_TARGET := help

ROOTDIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Infrastructure

.PHONY: infra
infra: ## Docker compose entrypoint of infra.
	@./scripts/compose_infra.sh

##@ Development

test: ## Run go test.
	@go test -v -count=1 ./...

run: ## Run hexago locally.
	@go run ./cmd/main.go ./cmd/wire.go
