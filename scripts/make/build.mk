.PHONY: build build-prod

build: ## Build docker image
	docker-compose -f docker-compose.dev.yml build -q $(service)

push: ## Push docker image
	docker-compose -f docker-compose.dev.yml push $(service)

logout: ## docker logout
	docker logout

build-push: build push logout## Build docker image (production)

