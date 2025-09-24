# Location Service Makefile

.PHONY: help build run test clean docs docker-up docker-down

# Default target
help:
	@echo "Available commands:"
	@echo "  build      - Build the application"
	@echo "  run        - Run the application locally"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean build artifacts"
	@echo "  docs       - Generate Swagger documentation"
	@echo "  docker-up  - Start services with Docker Compose"
	@echo "  docker-down- Stop Docker Compose services"
	@echo "  docker-logs- View Docker logs"

# Build the application
build:
	@echo "Building Location Service..."
	go build -o bin/location-svc cmd/main.go

# Run the application locally
run:
	@echo "Starting Location Service..."
	go run cmd/main.go

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...

# Clean build artifacts
clean:
	@echo "Cleaning..."
	rm -rf bin/
	rm -f main

# Generate Swagger documentation
docs:
	@echo "Generating Swagger docs..."
	swag init -g cmd/main.go

# Docker commands
docker-up:
	@echo "Starting services with Docker Compose..."
	docker-compose up --build

docker-down:
	@echo "Stopping Docker Compose services..."
	docker-compose down

docker-logs:
	@echo "Viewing Docker logs..."
	docker-compose logs -f app

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod tidy

# Format code
fmt:
	@echo "Formatting code..."
	go fmt ./...

# Lint code (requires golangci-lint)
lint:
	@echo "Linting code..."
	golangci-lint run