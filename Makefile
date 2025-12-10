.PHONY: help run build test fmt lint clean info

# Default target
.DEFAULT_GOAL := help

# Variables
PORT = 8888
BINARY_NAME = testing-studio

##@ General

help: ## Display this help message
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Running

run: ## Run the application
	@echo "🚀 Starting Testing Studio..."
	@go run cmd/server/main.go

start: run ## Alias for 'run'

##@ Build Operations

build: ## Build the application binary
	@echo "🔨 Building Testing Studio..."
	@go build -o $(BINARY_NAME) cmd/server/main.go
	@echo "✅ Build complete: ./$(BINARY_NAME)"

build-prod: ## Build optimized production binary
	@echo "🔨 Building optimized Testing Studio..."
	@go build -ldflags="-s -w" -o $(BINARY_NAME) cmd/server/main.go
	@echo "✅ Production build complete: ./$(BINARY_NAME)"

##@ Development

test: ## Run tests
	@echo "🧪 Running tests..."
	@go test ./... -v

fmt: ## Format code
	@echo "✨ Formatting code..."
	@go fmt ./...
	@echo "✅ Code formatted"

lint: ## Run linter
	@echo "🔍 Running linter..."
	@golangci-lint run || echo "⚠️  Install golangci-lint: https://golangci-lint.run/usage/install/"

deps: ## Download dependencies
	@echo "📦 Downloading dependencies..."
	@go mod download
	@echo "✅ Dependencies downloaded"

tidy: ## Tidy go.mod
	@echo "🧹 Tidying go.mod..."
	@go mod tidy
	@echo "✅ go.mod tidied"

##@ Cleanup

clean: ## Remove binary and configs
	@echo "🧹 Cleaning up..."
	@rm -f $(BINARY_NAME)
	@echo "✅ Cleanup complete"

clean-all: clean ## Remove binary, configs, and generated files
	@echo "🧹 Deep cleaning..."
	@rm -f configs.json
	@echo "✅ Deep cleanup complete"

##@ Quick Actions

open: ## Open Testing Studio in browser
	@echo "🌐 Opening Testing Studio..."
	@open http://localhost:$(PORT) || xdg-open http://localhost:$(PORT) || echo "Please open http://localhost:$(PORT) in your browser"

health: ## Check application health
	@echo "🏥 Checking health..."
	@curl -sf http://localhost:$(PORT)/api/configs > /dev/null && echo "✅ Healthy" || echo "❌ Unhealthy"

##@ Information

info: ## Show project information
	@echo "📋 Testing Studio Information"
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo "Name:          Testing Studio"
	@echo "Version:       3.0.0"
	@echo "Port:          $(PORT)"
	@echo "Binary:        $(BINARY_NAME)"
	@echo "URL:           http://localhost:$(PORT)"
	@echo "━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━"
	@echo ""
	@echo "🚀 Quick Start:"
	@echo "   make run      # Run the application"
	@echo "   make build    # Build binary"
	@echo "   make test     # Run tests"
