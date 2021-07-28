.PHONY: test-up test-down test

start-test: ## Build and start docker containers (test)
	docker-compose -f docker-compose.test.yml up -d $(service)_test

exec-test: ## Execute test suite
	docker-compose -f docker-compose.test.yml exec -T $(service)_test go test -v -vet=off -coverprofile=/tmp/coverage.out ./...

stop-test: ## Stop docker containers (test)
	docker-compose -f docker-compose.test.yml stop $(service)_test

clean-test: stop-test ## Stop docker containers, clean data and workspace (test)
	docker-compose -f docker-compose.test.yml rm -v

final-clean-test: ## Stop docker containers, clean data and workspace (test)
	docker-compose -f docker-compose.test.yml down -v --remove-orphans --rmi all

test: start-test exec-test final-clean-test ## Run test suite
