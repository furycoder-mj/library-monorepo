.PHONY: status logs start-dev stop-dev clean-dev 

status: ## Get status of containers
	docker-compose ps

logs: ## Get logs of containers
	docker-compose logs --tail=0 --follow
	
start-dev: ## Build and start docker container
	docker-compose -f docker-compose.dev.yml up -d $(service)

stop-dev: ## Stop docker container
	docker-compose -f docker-compose.dev.yml stop $(service)

clean-dev: stop-dev ## Stop docker containers, clean data and workspace (test)
	docker-compose -f docker-compose.dev.yml rm -f -s -v $(service)


