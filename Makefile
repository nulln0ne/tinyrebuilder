# Makefile for tinyrebuilder

.PHONY: all test fmt lint clean

# Default target
all: fmt test lint

# Run all tests
test:
	@echo "==> Running tests..."
	@go test -race ./...

# Format the source code
fmt:
	@echo "==> Formatting code..."
	@go fmt ./...

# Run the linter. Requires golangci-lint.
# Install with: go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
lint:
	@echo "==> Running linter..."
	@golangci-lint run

# Tidy up the module dependencies
tidy:
	@echo "==> Tidying module dependencies..."
	@go mod tidy

# Clean up build artifacts
clean:
	@echo "==> Cleaning up..."
	@go clean
