# Makefile for a Go project

# Variables
GO := go
GOLANGCI := golangci-lint

.PHONY: all fmt tidy lint

# Default target
all: fmt tidy lint

# Format code
fmt:
	@echo "Formatting Go code..."
	@gofmt -s -w .
	@goimports -w .

# Clean up module files
tidy:
	@echo "Tidying Go modules..."
	@$(GO) mod tidy

# Run linter
lint:
	@echo "Running golangci-lint..."
	@$(GOLANGCI) run --fix

