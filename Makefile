# Simple Makefile for a Go project

# Build the application
all: build

build:
	@echo "Building..."
	@go build -o main cmd/api/main.go

# Run the application
run:
	@go run cmd/api/main.go

# Create DB container
dc-up:
	@echo "Starting your application..."
	@docker compose up -d

# Shutdown DB container
dc-down:
	@echo "Closing..."
	@docker compose down

# Test the application
test:
	@echo "Testing..."
	@go test ./tests/... -v

# Clean the binary
clean:
	@echo "Cleaning..."
	@go mod tidy

migration-generate:
	@echo "Generate migration in ./migrations"
	@touch migrations/`date '+%Y-%m-%d_%s'`_migration.sql

secret-key:
	@node -e "console.log(require('crypto').randomBytes(256).toString('base64'));"

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/cosmtrek/air@latest; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run test clean
