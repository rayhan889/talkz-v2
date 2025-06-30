include .env
MIGRATIONS_PATH = ./migrate/migrations

.PHONY: build
build:
	@go build -o bin/talkz-v2 cmd/main.go

.PHONY: run
run: build
	@./bin/talkz-v2

.PHONY: test
test:
	@go test -v ./...

.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))

.PHONY: migrate-up
migrate-up:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) up

.PHONY: migrate-down
migrate-down:
	@migrate -path=$(MIGRATIONS_PATH) -database=$(DB_ADDR) down $(filter-out $@,$(MAKECMDGOALS))

.PHONY: seed
seed: 
	@go run ./migrate/seed/main.go

.PHONY: wire
wire:
	@wire gen github.com/rayhan889/talkz-v2/app