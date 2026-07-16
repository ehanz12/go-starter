# ==============================================================
# {{PROJECT_NAME}} – Makefile
# ==============================================================

.PHONY: run build test lint tidy docker-up docker-down migrate help

BINARY=bin/{{PROJECT_NAME}}
ENTRY=./cmd/api/main.go

## run: Start the development server with hot-reload (requires air)
run:
	@which air > /dev/null 2>&1 || (echo "Installing air..." && go install github.com/air-verse/air@latest)
	air

## build: Compile the production binary
build:
	@mkdir -p bin
	go build -ldflags="-s -w" -o $(BINARY) $(ENTRY)

## test: Run all tests with race detector
test:
	go test -race -count=1 ./...

## lint: Run golangci-lint
lint:
	@which golangci-lint > /dev/null 2>&1 || (echo "golangci-lint not found, install from https://golangci-lint.run/")
	golangci-lint run ./...

## tidy: Tidy and vendor go modules
tidy:
	go mod tidy
	go mod verify

## docker-up: Start services defined in docker-compose.yml
docker-up:
	docker compose up -d

## docker-down: Stop all docker compose services
docker-down:
	docker compose down

## migrate: Run database migrations (requires golang-migrate)
migrate:
	@which migrate > /dev/null 2>&1 || (echo "Install golang-migrate: https://github.com/golang-migrate/migrate")
	migrate -path database/migrations -database "$${DATABASE_URL}" up

## help: Display this help message
help:
	@echo "Available targets:"
	@sed -n 's/^## //p' Makefile | column -t -s ':' | sed -e 's/^/ /'
