.PHONY: run test build clean help

BINARY_NAME=app

.DEFAULT_GOAL := help

## Run the application
run:
	go run cmd/server/main.go

## Test
test:
	go test ./...

## Build the application
build:
	go build -o bin/$(BINARY_NAME) cmd/server/main.go

## Clean build files
clean:
	rm -rf bin/$(BINARY_NAME)

## Show available make commands
help:
	@echo "Available commands:"
	@echo "  make run    - Run the application"
	@echo "  make test   - Run tests"
	@echo "  make build  - Build the application"
	@echo "  make clean  - Clean build files"
	@echo "  make help   - Show this help message"