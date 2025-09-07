
# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Calculated `active` field in both `daily` and `cumulative` objects for national endpoint
- New `statistics.percentages` object containing:
  - `active` - Percentage of active cases
  - `recovered` - Percentage of recovered cases
  - `deceased` - Percentage of deceased cases
- Comprehensive API documentation for all endpoints
- Health check endpoint at `/health`
- Graceful shutdown handling with 30-second timeout
- Connection pooling configuration for database
- Request logging middleware
- CORS support for API endpoints

### Changed

- **BREAKING**: Complete restructure of national endpoint JSON response format
  - All response keys changed from Indonesian to English
  - Data reorganized into logical nested groups:
    - `daily` object for new daily cases
    - `cumulative` object for total cumulative counts
    - `statistics` object for calculated metrics and reproduction rate
- **BREAKING**: Reproduction rate (Rt) fields renamed and nested:
  - `rt` → `statistics.reproduction_rate.value`
  - `rt_upper` → `statistics.reproduction_rate.upper_bound`
  - `rt_lower` → `statistics.reproduction_rate.lower_bound`
- Improved error handling with structured error responses
- Database queries optimized with proper indexing
- Response structure now uses nested objects for better organization

### Removed

- **BREAKING**: `id` field no longer exposed in national endpoint responses
- Redundant database columns that duplicated calculated values

### Fixed

- N+1 query issues in data fetching
- Memory leaks in database connection handling
- Proper null handling for optional fields (reproduction rate)

## [1.0.0] - 2024-03-15

### Added

- Initial release of PICO COVID-19 API for Indonesia
- National endpoint for COVID-19 statistics
- Provincial endpoint for regional data
- Date range filtering support
- Latest data endpoint
- Basic authentication for admin endpoints
- PostgreSQL database integration
- Docker support for containerized deployment
- Environment-based configuration
- Automatic database migrations

### Security

- API rate limiting to prevent abuse
- Input validation for all query parameters
- SQL injection prevention through parameterized queries
- HTTPS enforcement in production

[Unreleased]: https://github.com/ryanaidilp/pico-api-go/compare/v1.0.0...HEAD
[1.0.0]: https://github.com/ryanaidilp/pico-api-go/releases/tag/v1.0.0
