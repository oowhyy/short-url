include .env
MIGRATION_FOLDER=migrations
POSTGRES_SETUP = "user=$(POSTGRES_USER) password=$(POSTGRES_PASSWORD) dbname=$(POSTGRES_DB) host=localhost port=$(POSTGRES_HOST_PORT) sslmode=disable"

testcoverage:
	@go test -coverprofile=c.out ./...
	@go tool cover -func=c.out
	@go tool cover -html=c.out

.PHONY: proto
proto:
	@protoc  --go_out=. --go-grpc_out=. api/v1/shorturl/urls.proto

.PHONY: mig-make
mig-make:
	@mkdir -p $(MIGRATION_FOLDER)
	@goose -dir "$(MIGRATION_FOLDER)" create "$(name)" sql


mig-up:
	@goose -dir "$(MIGRATION_FOLDER)" postgres $(POSTGRES_SETUP) up

mig-down:
	@goose -dir "$(MIGRATION_FOLDER)" postgres $(POSTGRES_SETUP) down