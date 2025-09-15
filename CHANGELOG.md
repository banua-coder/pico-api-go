# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]





## [v2.4.0] - 2025-09-15

### Added

- Configure deploy workflow for minimal production build (6.1mb) ([1d362189](https://github.com/banua-coder/pico-api-go/commit/1d362189e56fe7e7cea6b616593944761d1fdbbd))
- Optimize binary size with conditional swagger compilation ([64a56304](https://github.com/banua-coder/pico-api-go/commit/64a56304257db60a1e6db0561431701b9ecd57c9))
- Enhance ci with intelligent testing and coverage thresholds ([3058e376](https://github.com/banua-coder/pico-api-go/commit/3058e37690ea612eaa50aa4fa401cb2b366db931))
- Enhance release workflow with swagger regeneration and script organization ([cf94c807](https://github.com/banua-coder/pico-api-go/commit/cf94c807fe32b6970da58e527dbb68285628a3df))
- Simplify changelog generator and remove unnecessary complexity ([fc609bb7](https://github.com/banua-coder/pico-api-go/commit/fc609bb75f535ca6b8962886146eda655c5d894a))

### Fixed

- Exclude test files from golangci-lint to resolve mock interface issues ([d77e2c0c](https://github.com/banua-coder/pico-api-go/commit/d77e2c0cbe5004d33c35fec60f1e16a24983689c))
- Add golangci-lint configuration to resolve test file issues ([73b1e516](https://github.com/banua-coder/pico-api-go/commit/73b1e516277da92c5b65411f4a1b15b2ad163f81))
- Explicitly reference embedded db methods to resolve linter issues ([f86a70b1](https://github.com/banua-coder/pico-api-go/commit/f86a70b1736cc654841f4d8c3639296f0bea1b21))
- Resolve golangci-lint version compatibility issue in ci ([86872eec](https://github.com/banua-coder/pico-api-go/commit/86872eeca96a5e13920181a6dbfaec71ea7dc397))
- Resolve ci failures - integration tests and code formatting ([1f26b10f](https://github.com/banua-coder/pico-api-go/commit/1f26b10f03fa297c734c27b119db04fbf55e4d51))
- Remove redundant province data from latest_case in province list api ([00d63ebc](https://github.com/banua-coder/pico-api-go/commit/00d63ebc908a3cfcd2484a6d64aaa1fd4f402a2e))
- Implement config-based version management system ([3a68d854](https://github.com/banua-coder/pico-api-go/commit/3a68d85496d8bf3fbfd3ea44fae5dfc515f1b21c))
- Resolve workflow duplicates and conflicts ([2b756609](https://github.com/banua-coder/pico-api-go/commit/2b756609e3cb5701496ca699af64a1d22f083f36))
- Simplify workflows and restore working deploy.yml ([547f4556](https://github.com/banua-coder/pico-api-go/commit/547f45566a4f89d91ca4b15a595c784d0dbafe83))
- Fix generate changelog script (script) ([9ab39f0f](https://github.com/banua-coder/pico-api-go/commit/9ab39f0f40b73bad51fbc34b52645c97f2a25839))

### Documentation

- Update readme with latest project structure and ci features ([82cc7c9a](https://github.com/banua-coder/pico-api-go/commit/82cc7c9aa86c4518b15489d15c8a41689835fa25))

## [v2.3.0] - 2025-09-08

### Documentation

- Remove rate_limiting.md documentation file (6c43bb3)

## [v2.2.0] - 2025-09-08

### Fixed

- Handle error return values in rate limit tests for linter (9597171)

### Documentation

- Update .env.example with rate limiting configuration (73a9d63)

### CI/CD

- Add cleanup job (workflow) (cdf5dc7)

### Maintenance

- Update for next version (version) (80d5451)

### Style

- Add missing newline at end of workflow file (5266ab2)
- Fix formatting and add missing newlines (f2f373a)

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
