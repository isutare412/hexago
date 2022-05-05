.DEFAULT_TARGET := help

ROOTDIR := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))

IMG ?= hexago/payment
TAG ?= v0.1.0

##@ General

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Deployment

.PHONY: image
image: ## Build server image.
	@echo "Build image: ${IMG}:${TAG}"
	@docker build \
		-f deployments/arm64.Dockerfile \
		-t ${IMG}:${TAG} \
		.

##@ Development

.PHONY: test
test: ## Run go test.
	@go test -v -count=1 ./...

.PHONY: run-local
run-local: ## Run server locally.
	@go run ./cmd/main.go ./cmd/wire.go

.PHONY: run-docker
run-docker: ## Run server in docker.
	@docker run --rm -d \
		--name hexago-payment \
		--network hexago_app-tier \
		${IMG}:${TAG}

.PHONY: stop-docker
stop-docker: ## Stop server in docker.
	@docker stop hexago-payment
