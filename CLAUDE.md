# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Development Commands

### Running the Application
```bash
# Development mode
go run cmd/main.go

# Build for production
go build -o pico-api-go cmd/main.go

# Install dependencies
go mod tidy
```

### Environment Setup
```bash
# Copy environment template
cp .env.example .env
# Edit .env with actual database credentials
```

### Testing Commands
```bash
# Run all tests
make test
# or
go test -v ./...

# Run unit tests only
make test-unit
# or
go test -v ./internal/...

# Run integration tests only  
make test-integration
# or
go test -v ./test/integration/...

# Run tests with coverage
make test-coverage

# Run tests with race detection
make test-race
```

### Git Flow Workflow
```bash
# Start new feature
git flow feature start feature-name

# Finish feature
git flow feature finish feature-name
```

## Architecture Overview

### Clean Architecture Layers
The application follows a clean architecture pattern with clear separation of concerns:

1. **Handler Layer** (`internal/handler/`) - HTTP request/response handling, route definitions
2. **Service Layer** (`internal/service/`) - Business logic, date parsing, data transformation
3. **Repository Layer** (`internal/repository/`) - Database access, SQL queries
4. **Models Layer** (`internal/models/`) - Data structures and domain entities

### Key Architecture Patterns

#### Dependency Injection Flow
- `cmd/main.go` orchestrates the entire dependency chain
- Database connection → Repositories → Service → Handlers → Router
- Each layer only depends on interfaces from the layer below

#### Database Relationships
- `national_cases.id` serves as the primary key and is referenced by `province_cases.day`
- Province cases link to dates through the national_cases table
- Province codes use Indonesian administration codes (e.g., Aceh: 11, Sulawesi Tengah: 72)

#### Response Structure
All API responses follow a consistent JSON structure defined in `internal/handler/response.go`:
```json
{
  "status": "success|error",
  "data": {},
  "error": "error message"
}
```

#### Configuration Management
- Environment-based configuration using `.env` files
- Production-ready for shared hosting deployment
- Database credentials and server settings centralized in `internal/config/`

### COVID-19 Data Context
- **National Cases**: Daily and cumulative COVID-19 statistics for Indonesia
- **Province Cases**: Provincial data including ODP (Orang Dalam Pemantauan) and PDP (Pasien Dalam Pengawasan) tracking
- **R-rate**: Reproductive rate data when available (nullable fields)
- **Date Filtering**: All endpoints support `start_date` and `end_date` query parameters (YYYY-MM-DD format)

### Database Schema Understanding
- There's a typo in the database schema: `cumulative_finished_persoon_under_observation` (should be `person`)
- The repository layer handles this mapping correctly
- Province cases are linked to national cases via the `day` field, which references `national_cases.id`

### Middleware Stack
Applied in order: Recovery → Logging → CORS
- Recovery: Panic handling with structured error responses  
- Logging: Request logging with method, path, status, size, duration
- CORS: Configured for web frontend integration

### Testing Architecture
- **Unit Tests**: Located alongside source code (e.g., `*_test.go` files)
- **Repository Tests**: Use `go-sqlmock` for database mocking
- **Service Tests**: Use `testify/mock` for repository mocking  
- **Handler Tests**: Use `httptest` for HTTP request/response testing
- **Integration Tests**: Located in `test/integration/` for end-to-end API testing

### Development Workflow
When adding new features, follow the established patterns:
1. Create models first with corresponding unit tests
2. Implement repository interfaces and implementations with mock-based tests
3. Create service layer methods with comprehensive unit tests
4. Build HTTP handlers with HTTP testing
5. Add integration tests for complete API workflows
6. Use `make test` to run the full test suite before committing
