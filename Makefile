ifneq (,$(wildcard ./.env))
    include .env
    export
endif

help: ## This help dialog.
	@echo "Available commands:"
	@grep -F -h "##" $(MAKEFILE_LIST) | grep -F -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

run-local: ## Run the app locally
	go run app.go

run-air : ## Run the app with air - live reload
	air

requirements: ## Generate go.mod & go.sum files
	go mod tidy

clean-packages: ## Clean packages
	go clean -modcache

DATABASE_URL="mysql://$(DB_USER):$(DB_PASS)@tcp($(DB_HOST):$(DB_PORT))/$(DB_NAME)"

migrate-up: ## Migrates the database to the latest version
	migrate -path database/migrations -database $(DATABASE_URL) -lock-timeout=20 -verbose up

migrate-down: ## Rollback the database to the first version
	migrate -path database/migrations -database $(DATABASE_URL) -lock-timeout=20 -verbose down

migrate-force: ## Force the database to a specific version with the VERSION parameter
	migrate -path database/migrations -database $(DATABASE_URL) -lock-timeout=20 -verbose force $(VERSION)

test-url: ## Test the database connection
	@echo $(DATABASE_URL)
