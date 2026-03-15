# Changelog
## [v2.5.0] - 2026-03-15

### ✨ Features

-  enable Swagger UI in production build (fe19501)
-  migrate Lumen endpoints — regencies, hospitals, task forces (29d7863)
-  add vaccination endpoints — national, province, locations (32ce6f4)
-  add province stats, national by day, single province endpoints (16ad143)

### 🐛 Bug Fixes

-  create-release job now runs after successful deployment (73b1fbb)
-  update host to pico-api-go.banuacoder.com (f97f6c9)
-  use ./cmd/ package path in Dockerfile (main_production.go not found) (8359a50)
-  update swagger host to pico-api-go.banuacoder.com (983fa25)
-  handle NULL values in province case ODP/PDP fields (8238837)
-  upgrade golangci-lint to v2.10.1 for Go 1.26 compatibility (aa3bcf9)
-  remove redundant embedded DB field selectors (QF1008) (cfbeecc)
-  remove all redundant embedded DB field selectors (QF1008) (d85e9d6)
-  resolve golangci-lint errcheck and govet issues (5428b62)
-  resolve git auth error in release branch creation workflow (6de04ee)
-  add nolint:errcheck to remaining defer Close() calls (a4d975f)
-  properly handle rows.Close() and db.Close() errors (29041b6)
-  properly handle rows/db Close() errors (07fb831)
-  pin release workflow to @v1 tag (08a1f1e)
-  restore corrupt assert.Equal args and fix .version-config.yml patterns (a016065)

### 📚 Documentation

-  add CONTRIBUTING.md with git flow workflow (401732c)

### 🧪 Tests

-  add unit tests for all new handlers + refactor to service interfaces (6ff000d)
-  add service layer unit tests to improve coverage (9066ddd)
-  add repository tests to meet 85% coverage threshold (0d01572)

### 👷 CI

-  add Docker build and deploy workflow to GHCR (3928c39)
-  add setup-buildx-action to fix GHA cache export (56a74d0)
-  separate deploy (main only) and CI workflows (develop/feature/PR) (3331bac)
-  deploy triggers on main + hotfix/**, CI on develop + feature/bugfix/release + PRs (820f648)
-  restore original ci.yml; deploy triggers on version tags (Docker/GHCR) (75c1982)
-  add SSH deploy step — pull & restart container after image push (9936208)
-  retrigger CI after golangci-lint-action upgrade to v7 (5ce2c25)
-  re-trigger release workflow with new WORKFLOW_TOKEN (556305d)

### 🔧 Chores

-  back-merge v2.4.0 from main to develop (832fc4f)
-  update version to 2.5.0 and fix remaining issues (367fac9)
-  sync develop with main (resolve conflicts, keep main version) (b0d0bd6)
-  upgrade Go to 1.26 and migrate to reusable workflows (31439dd)
-  integrate branch cleanup from reusable workflow (4d473d6)
-  prepare v2.5.0 release (0aca0fc)
-  bump version to v2.5.0 (b47db4a)

### 📝 Other Changes

- Merge pull request #38 from banua-coder/chore/back-merge-v2.4.0-to-develop (38837aa)
- Merge pull request #40 from banua-coder/fix/nullable-province-case-fields (8c77691)
- hotfix: fix deploy workflow — trigger on tags, add SSH deploy step (776d957)
- Merge pull request #43 from banua-coder/hotfix/fix-deploy-workflow (461c699)
- Merge pull request #45 from banua-coder/chore/upgrade-go-and-integrate-reusable-workflows (787b598)
- Merge pull request #44 from banua-coder/feature/migrate-lumen-endpoints (cb0a27f)
- Merge pull request #46 from banua-coder/chore/integrate-branch-cleanup-housekeeping (81ec193)
- Merge pull request #47 from banua-coder/chore/prepare-release-v2.5.0 (70d1dac)
- Merge pull request #48 from banua-coder/release/v2.5.0 (fad460f)


All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]






## [v2.5.0] - 2026-03-15

### Hotfixes

- Fix deploy workflow — trigger on tags, add ssh deploy step ([776d9578](https://github.com/banua-coder/pico-api-go/commit/776d9578d3efc9ac73d8a8eaab1354b22ceaba1d))

### Added

- Add province stats, national by day, single province endpoints ([16ad143b](https://github.com/banua-coder/pico-api-go/commit/16ad143b4e07f699e20bde8b7cc7719d34bdd851))
- Add vaccination endpoints — national, province, locations ([32ce6f4a](https://github.com/banua-coder/pico-api-go/commit/32ce6f4a8dbdb9bf7e3d8ae08d160febe250b4ce))
- Migrate lumen endpoints — regencies, hospitals, task forces ([29d7863f](https://github.com/banua-coder/pico-api-go/commit/29d7863fe8fd2118e5e90d5bf4aadf7030d05fc7))
- Enable swagger ui in production build ([fe19501b](https://github.com/banua-coder/pico-api-go/commit/fe19501ba81f45538cc2a0afba1f16801f2f5abd))

### Fixed

- Remove all redundant embedded db field selectors (qf1008) ([d85e9d61](https://github.com/banua-coder/pico-api-go/commit/d85e9d61bd4f679baf4223f25b1aae035a0cf9fb))
- Remove redundant embedded db field selectors (qf1008) ([cfbeecc3](https://github.com/banua-coder/pico-api-go/commit/cfbeecc3acdf76bb8185ab251bfa642ad88f6b09))
- Upgrade golangci-lint to v2.10.1 for go 1.26 compatibility ([aa3bcf9e](https://github.com/banua-coder/pico-api-go/commit/aa3bcf9e7058a84b2bc1010a89ec2fa13dba10f3))
- Handle null values in province case odp/pdp fields ([82388377](https://github.com/banua-coder/pico-api-go/commit/82388377d6d7f4449ab3bc50b9e958b8777d0012))
- Update swagger host to pico-api-go.banuacoder.com ([983fa25e](https://github.com/banua-coder/pico-api-go/commit/983fa25e636735ac1d3c92a1abee885f9736bfa0))
- Use ./cmd/ package path in dockerfile (main_production.go not found) ([8359a50c](https://github.com/banua-coder/pico-api-go/commit/8359a50c57024cbf63419d8b5356e12357cd795a))
- Update host to pico-api-go.banuacoder.com ([f97f6c9e](https://github.com/banua-coder/pico-api-go/commit/f97f6c9e64267460b40079809607d1938b865ba7))
- Create-release job now runs after successful deployment ([73b1fbbc](https://github.com/banua-coder/pico-api-go/commit/73b1fbbc0361b7340839ffcf7c82a212d04eb39e))

### Documentation

- Add contributing.md with git flow workflow ([401732c9](https://github.com/banua-coder/pico-api-go/commit/401732c955c1d42924fbbc9938b4920439fa0673))

### Tests

- Add repository tests to meet 85% coverage threshold ([0d01572b](https://github.com/banua-coder/pico-api-go/commit/0d01572b7a90b375ca446ccdfcf079a36eaeed48))
- Add service layer unit tests to improve coverage ([9066ddd6](https://github.com/banua-coder/pico-api-go/commit/9066ddd6dbca1fec8fece3bafc2fca1033f0ac12))
- Add unit tests for all new handlers + refactor to service interfaces ([6ff000db](https://github.com/banua-coder/pico-api-go/commit/6ff000db026e7d1b8cd9506a196acefa91f90180))

### CI/CD

- Retrigger ci after golangci-lint-action upgrade to v7 ([5ce2c25e](https://github.com/banua-coder/pico-api-go/commit/5ce2c25e686a0aa529cf0d20f35bb9fcd9d71cc0))
- Add ssh deploy step — pull & restart container after image push ([9936208a](https://github.com/banua-coder/pico-api-go/commit/9936208a4c1927ec04607a9c1cd19f3320bf0647))
- Restore original ci.yml; deploy triggers on version tags (docker/ghcr) ([75c19820](https://github.com/banua-coder/pico-api-go/commit/75c19820821009fec7bbadaa942c90f8182d6cff))
- Deploy triggers on main + hotfix/**, ci on develop + feature/bugfix/release + prs ([820f648d](https://github.com/banua-coder/pico-api-go/commit/820f648d0c455b4f8fbac747d5c97880cb6b2f76))
- Separate deploy (main only) and ci workflows (develop/feature/pr) ([3331bace](https://github.com/banua-coder/pico-api-go/commit/3331bace0d9b18b6dad84729ea72df85ce714c80))
- Add setup-buildx-action to fix gha cache export ([56a74d00](https://github.com/banua-coder/pico-api-go/commit/56a74d0036838b949eb354a542b9c169f84de5ca))
- Add docker build and deploy workflow to ghcr ([3928c39c](https://github.com/banua-coder/pico-api-go/commit/3928c39c60a81b9c0b2d10dd7a0dde282d08e677))

### Maintenance

- Upgrade go to 1.26 and migrate to reusable workflows ([31439dd0](https://github.com/banua-coder/pico-api-go/commit/31439dd0ad3b77944fabf14eab5eaeeedc9d23bc))
- Sync develop with main (resolve conflicts, keep main version) ([b0d0bd6d](https://github.com/banua-coder/pico-api-go/commit/b0d0bd6db6b5412a6baf8d87c34e31805edbfd5f))
- Update version to 2.5.0 and fix remaining issues ([367fac96](https://github.com/banua-coder/pico-api-go/commit/367fac96a7f71f7ab04c299c6ae9e9ab1d44c86f))

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
