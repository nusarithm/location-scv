#!/bin/bash

# Development setup script for Location Service

set -e

echo "🚀 Location Service Development Setup"
echo "======================================"

# Function to check if command exists
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Check prerequisites
echo "📋 Checking prerequisites..."

if ! command_exists docker; then
    echo "❌ Docker not found. Please install Docker first."
    exit 1
fi

if ! command_exists docker-compose; then
    echo "❌ Docker Compose not found. Please install Docker Compose first."
    exit 1
fi

if ! command_exists go; then
    echo "❌ Go not found. Please install Go first."
    exit 1
fi

echo "✅ All prerequisites found!"

# Install Go dependencies
echo "📦 Installing Go dependencies..."
go mod tidy

# Install swag if not exists
if ! command_exists swag; then
    echo "📦 Installing swag for Swagger documentation..."
    go install github.com/swaggo/swag/cmd/swag@latest
fi

# Generate Swagger docs
echo "📚 Generating Swagger documentation..."
swag init -g cmd/main.go

# Build the application to check for errors
echo "🔨 Building application..."
go build -o bin/location-svc cmd/main.go

echo ""
echo "✅ Setup complete!"
echo ""
echo "🐳 To start with Docker (recommended):"
echo "   make docker-up"
echo ""
echo "🏃 To run locally (need PostgreSQL running):"
echo "   make run"
echo ""
echo "📖 Access Swagger documentation at:"
echo "   http://localhost:8080/swagger/index.html"
echo ""
echo "🏥 Health check:"
echo "   http://localhost:8080/health"