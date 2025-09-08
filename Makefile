.PHONY: build run test test-unit test-integration clean help

# Build the application
build:
	go build -o bin/pico-api-go cmd/main.go

# Run the application
run:
	go run cmd/main.go

# Run all tests
test:
	go test -v ./...

# Run unit tests only
test-unit:
	go test -v ./internal/...

# Run integration tests only
test-integration:
	go test -v ./test/integration/...

# Run tests with coverage
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Run tests with race detection
test-race:
	go test -v -race ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Install dependencies
deps:
	go mod tidy
	go mod download

# Format code
fmt:
	go fmt ./...

# Run linter (requires golangci-lint)
lint:
	golangci-lint run

# Setup development environment
setup: deps
	cp .env.example .env
	@echo "Don't forget to update .env with your database credentials"

# Development server with hot reload (requires air)
dev:
	air

# Run benchmarks
bench:
	go test -bench=. -benchmem ./...

# Check for vulnerabilities
security:
	govulncheck ./...

help:
	@echo "Available commands:"
	@echo "  build            - Build the application"
	@echo "  run              - Run the application"
	@echo "  test             - Run all tests"
	@echo "  test-unit        - Run unit tests only"
	@echo "  test-integration - Run integration tests only"
	@echo "  test-coverage    - Run tests with coverage report"
	@echo "  test-race        - Run tests with race detection"
	@echo "  clean            - Clean build artifacts"
	@echo "  deps             - Install dependencies"
	@echo "  fmt              - Format code"
	@echo "  lint             - Run linter"
	@echo "  setup            - Setup development environment"
	@echo "  dev              - Run development server with hot reload"
	@echo "  bench            - Run benchmarks"
	@echo "  security         - Check for vulnerabilities"
	@echo "  help             - Show this help message"
	