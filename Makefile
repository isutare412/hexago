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

##@ Applications

.PHONY: image-gateway
image-gateway: ## Build docker image of hexago gateway.
	$(MAKE) -C gateway image

.PHONY: image-payment
image-payment: ## Build docker image of hexago payment.
	$(MAKE) -C payment image

.PHONY: run-gateway
run-gateway: ## Run docker image of hexago gateway.
	$(MAKE) -C gateway run-docker

.PHONY: run-payment
run-payment: ## Run docker image of hexago payment.
	$(MAKE) -C payment run-docker

.PHONY: stop-gateway
stop-gateway: ## Stop docker image of hexago gateway.
	$(MAKE) -C gateway stop-docker

.PHONY: stop-payment
stop-payment: ## Stop docker image of hexago payment.
	$(MAKE) -C payment stop-docker
