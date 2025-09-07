# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]


## [v2.1.0] - 2025-09-08

### Maintenance

- Bump to 2.1.0 (version)
- Update version to 2.0.2 in health endpoint

## [v2.0.1] - 2025-09-07

### Hotfixes

- Ensure reproduction rate always appears in json response even when null - fix rt values to always be included in api responses for consistency - update reproductionrate struct to use pointer types for proper null handling - modify transformation logic to always include rt structure - update tests to handle new pointer-based rt values - bump version to 2.0.1 for hotfix release this ensures api consumers always receive the reproduction_rate object structure, with null values when data is not available.

### Fixed

- Correct database column typo and ensure rt fields always present in json response - fix database column name typo from 'cumulative_finished_persoon_under_observation' to 'cumulative_finished_person_under_observation' - remove 'omitempty' from rt fields (rt, rt_upper, rt_lower) to ensure they always appear in json responses even when null - update all sql queries and tests to use correct column name - critical production hotfix for database errors and missing rt data
- Correct [Unreleased] section format in CHANGELOG.md

### Added

- Add hotfix branch support to changelog generator

### Maintenance

- Bump version to 2.0.1 for hotfix release (version)

## [v2.0.0] - 2025-09-07

### Maintenance

- Bump version to 2.0.0 (versino)

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
