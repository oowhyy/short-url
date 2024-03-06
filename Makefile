include .env
MIGRATION_FOLDER=migrations
POSTGRES_SETUP = "user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) host=localhost port=$(POSTGRES_HOST_PORT) sslmode=disable"

test-coverage:
	@go test -coverprofile=c.out ./...
	@go tool cover -func=c.out
	@go tool cover -html=c.out

test-up:
	@docker compose up -d postgres-test

test-down:
	@docker compose down postgres-test

proto:
	@protoc  --go_out=. --go-grpc_out=. api/v1/shorturl/urls.proto

mig-make:
	@mkdir -p $(MIGRATION_FOLDER)
	@goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql


mig-up:
	@goose -dir "$(MIGRATION_FOLDER)" postgres $(POSTGRES_SETUP) up

mig-down:
	@goose -dir "$(MIGRATION_FOLDER)" postgres $(POSTGRES_SETUP) down

up:
	@docker compose up -d shorturl postgres

down:
	@docker compose down shorturl postgres
