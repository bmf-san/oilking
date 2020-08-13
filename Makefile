.PHONY: help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help

build: ## Build containers by docker-compose.
	docker-compose build

up: ## Run containers by docker-compose.
	docker-compose up

up-d: ## Run container in the background by docker-compose.
	docker-compose up -d

lint: ## Run golint.
	docker exec -it oilking golint ./...

test: ## Run tests.
	docker exec -it oilking go test -v ./...
