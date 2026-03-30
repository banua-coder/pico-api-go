# Changelog
## [2.8.0] - 2026-03-30

### ✨ Features

-  restructure gender stats response for better readability (#67) (ab2968f)
-  restructure vaccination response payload for better readability (6eb9b35)
-  add vaccination coverage percentage to response (af6133d)
- feat(cache): add in-memory caching layer with TTL and admin invalidation (dd0dba8)
- feat(cache): add Redis-backed dual-layer cache support (d79cd9b)

### 🐛 Bug Fixes

-  use WORKFLOW_TOKEN for gh CLI calls in release-branch-creation workflow (f8d618e)
-  use WORKFLOW_PAT org secret instead of WORKFLOW_TOKEN (d6d932c)
-  remove /api/v1 prefix from @Router annotations (d6ba522)
-  remove /api/v1 prefix from @Router annotations (fdac8c4)
-  remove /api/v1 prefix from @Router annotations (9e1d85e)
-  remove /api/v1 prefix from @Router annotations (d56aeba)
-  remove /api/v1 prefix from @Router annotations (a2e3f46)
-  remove /api/v1 prefix from @Router annotations (208a4c7)
-  accept both v-prefixed and plain semver tags in deploy trigger (684e843)
-  use WORKFLOW_TOKEN for gh CLI calls in release-branch-creation workflow (662fd4c)
-  rewrite vaccination DTO tests to match actual struct fields (e67b84b)
-  remove duplicate bump-develop-version job (handled by reusable workflow) (a50d92f)
-  remove duplicate release-branch-creation workflow (fully handled by reusable workflow) (a4ba4d7)
-  resolve errcheck linter errors in cached service tests (92aebd4)

### 📚 Documentation

-  update changelog for v2.6.0 (ebde3af)
-  regenerate swagger with fixed paths (no double /api/v1) (a659834)
-  regenerate swagger with fixed paths (no double /api/v1) (9b4bcee)
-  regenerate swagger with fixed paths (no double /api/v1) (30d9fd5)
-  update changelog for 2.6.1 (cf610ca)
-  update changelog for 2.7.0 (67f765f)

### 🧪 Tests

-  add unit tests for cached services and admin handler (39f1559)

### 🔧 Chores

-  bump version to v2.7.0 for next development cycle (6356336)
-  bump version to 2.6.1 (520cf6d)
-  bump version to v2.7.0 for next development cycle (be91838)
-  prepare v2.7.0 release (5fc39d9)
-  bump version to 2.8.0 for next development cycle (3927c36)

### 📝 Other Changes

- Merge pull request #62 from banua-coder/fix/release-workflow-token (b70fa3c)
- Merge pull request #60 from banua-coder/chore/bump-next-version (8ac5cd0)
- Merge branch 'develop' into chore/back-merge-v2.6.0 (c8c5136)
- Merge pull request #63 from banua-coder/chore/back-merge-v2.6.0 (7c96678)
- Merge pull request #64 from banua-coder/hotfix/2.6.1 (e4adcd0)
- Merge pull request #65 from banua-coder/chore/back-merge-2.6.1 (4e29a65)
- Merge pull request #66 from banua-coder/feature/vaccination-response-restructure (934ac0c)
- Merge pull request #69 from banua-coder/feature/vaccination-coverage-v2 (d5bf33d)
- Merge pull request #72 from banua-coder/chore/prepare-release-v2.7.0 (e3c764d)
- Merge pull request #74 from banua-coder/chore/bump-version-to-v2.8.0-dev (9d67911)
- Merge pull request #70 from banua-coder/release/2.7.0 (802e888)
- Merge branch 'develop' into chore/back-merge-2.7.0 (213c3ad)
- Merge pull request #77 from banua-coder/chore/back-merge-2.7.0 (3509785)
- Merge branch 'develop' into feat/in-memory-cache (a805d9c)
- Merge pull request #78 from banua-coder/feat/in-memory-cache (8b547c2)
- release: 2.8.0 (448e788)

## [2.7.0] - 2026-03-25

### ✨ Features

-  restructure gender stats response for better readability (#67) (ab2968f)
-  restructure vaccination response payload for better readability (6eb9b35)
-  add vaccination coverage percentage to response (af6133d)

### 🐛 Bug Fixes

-  use WORKFLOW_TOKEN for gh CLI calls in release-branch-creation workflow (f8d618e)
-  use WORKFLOW_PAT org secret instead of WORKFLOW_TOKEN (d6d932c)
-  remove /api/v1 prefix from @Router annotations (d6ba522)
-  remove /api/v1 prefix from @Router annotations (fdac8c4)
-  remove /api/v1 prefix from @Router annotations (9e1d85e)
-  remove /api/v1 prefix from @Router annotations (d56aeba)
-  remove /api/v1 prefix from @Router annotations (a2e3f46)
-  remove /api/v1 prefix from @Router annotations (208a4c7)
-  accept both v-prefixed and plain semver tags in deploy trigger (684e843)
-  use WORKFLOW_TOKEN for gh CLI calls in release-branch-creation workflow (662fd4c)
-  rewrite vaccination DTO tests to match actual struct fields (e67b84b)
-  remove duplicate bump-develop-version job (handled by reusable workflow) (a50d92f)
-  remove duplicate release-branch-creation workflow (fully handled by reusable workflow) (a4ba4d7)

### 📚 Documentation

-  update changelog for v2.6.0 (ebde3af)
-  regenerate swagger with fixed paths (no double /api/v1) (a659834)
-  regenerate swagger with fixed paths (no double /api/v1) (9b4bcee)
-  regenerate swagger with fixed paths (no double /api/v1) (30d9fd5)
-  update changelog for 2.6.1 (cf610ca)

### 🔧 Chores

-  bump version to v2.7.0 for next development cycle (6356336)
-  bump version to 2.6.1 (520cf6d)
-  bump version to v2.7.0 for next development cycle (be91838)
-  prepare v2.7.0 release (5fc39d9)

### 📝 Other Changes

- Merge pull request #62 from banua-coder/fix/release-workflow-token (b70fa3c)
- Merge pull request #60 from banua-coder/chore/bump-next-version (8ac5cd0)
- Merge branch 'develop' into chore/back-merge-v2.6.0 (c8c5136)
- Merge pull request #63 from banua-coder/chore/back-merge-v2.6.0 (7c96678)
- Merge pull request #64 from banua-coder/hotfix/2.6.1 (e4adcd0)
- Merge pull request #65 from banua-coder/chore/back-merge-2.6.1 (4e29a65)
- Merge pull request #66 from banua-coder/feature/vaccination-response-restructure (934ac0c)
- Merge pull request #69 from banua-coder/feature/vaccination-coverage-v2 (d5bf33d)
- Merge pull request #72 from banua-coder/chore/prepare-release-v2.7.0 (e3c764d)
- Merge pull request #70 from banua-coder/release/2.7.0 (802e888)

## [2.6.1] - 2026-03-25

### 🐛 Bug Fixes

-  use WORKFLOW_PAT org secret instead of WORKFLOW_TOKEN (d6d932c)
-  remove /api/v1 prefix from @Router annotations (d6ba522)
-  remove /api/v1 prefix from @Router annotations (fdac8c4)
-  remove /api/v1 prefix from @Router annotations (9e1d85e)
-  remove /api/v1 prefix from @Router annotations (d56aeba)
-  remove /api/v1 prefix from @Router annotations (a2e3f46)
-  remove /api/v1 prefix from @Router annotations (208a4c7)

### 📚 Documentation

-  update changelog for v2.6.0 (ebde3af)
-  regenerate swagger with fixed paths (no double /api/v1) (a659834)
-  regenerate swagger with fixed paths (no double /api/v1) (9b4bcee)
-  regenerate swagger with fixed paths (no double /api/v1) (30d9fd5)

### 🔧 Chores

-  bump version to 2.6.1 (520cf6d)

### 📝 Other Changes

- Merge pull request #64 from banua-coder/hotfix/2.6.1 (e4adcd0)

## [v2.6.0] - 2026-03-25

### ✨ Features

-  add pagination to hospitals, vaccination, regencies, and task-force endpoints (8dd0903)

### 🐛 Bug Fixes

-  use update-version.sh script for develop version bump (750a7e7)
-  handle errcheck lint issues in GetPaginatedByProvinceID (057a22e)

### 📚 Documentation

-  update changelog for v2.5.1 (5999231)

### 🧪 Tests

-  improve code coverage from 47% to 61.6% (63552ac)
-  push code coverage from 57% to 80.5% (adc36b5)
-  push coverage to 82.3% (above 80% CI threshold) (988ccc5)

### 👷 CI

-  exclude cmd, docs, pkg/database from coverage calculation (aba5f17)

### 🔧 Chores

-  bump version to v2.6.0 for next development cycle (85e547b)
-  bump version to 4.6.0 and fix lint errors (f6a1bd1)
-  correct version to 2.6.0 (not 4.6.0) (6fa7209)
-  bump version to v2.6.0 (59fcbca)
-  prepare v2.6.0 release (2421171)

### 📝 Other Changes

- Merge pull request #49 from banua-coder/chore/bump-version-to-v2.6.0-dev (7f42cdc)
- Merge branch 'develop' into chore/back-merge-v2.5.0 (46efc5f)
- Merge pull request #53 from banua-coder/chore/back-merge-v2.5.0 (9b873c5)
- Merge branch 'develop' into chore/back-merge-v2.5.1 (fb8a336)
- Merge pull request #57 from banua-coder/chore/back-merge-v2.5.1 (bbffd95)
- Merge pull request #58 from banua-coder/feature/pagination-and-restructure-api (d967794)
- Merge pull request #61 from banua-coder/chore/prepare-release-v2.6.0 (999d8a1)
- Merge pull request #59 from banua-coder/release/v2.6.0 (c675ca6)

## [v2.5.1] - 2026-03-15

### 🐛 Bug Fixes

-  add all new endpoints to /api/v1 index response (0bb07af)

### 📚 Documentation

-  update changelog for v2.5.0 (ed28b3f)

### 🔧 Chores

-  bump version to v2.5.1 (31e554d)

### 📝 Other Changes

- Merge pull request #56 from banua-coder/hotfix/v2.5.1 (2ed01ec)

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








## [v2.7.0] - 2026-03-25

### Added

- Add vaccination coverage percentage to response ([af6133dc](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/af6133dcfeeb3577b774513149f93b4035612731))
- Restructure vaccination response payload for better readability ([6eb9b350](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/6eb9b3503c55c53c4301a904db3df15d2668cf9a))
- Restructure gender stats response for better readability (#67) ([ab2968f6](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/ab2968f644bdf589ca51c1d8c63f9f616741e0dd))

### Fixed

- Rewrite vaccination dto tests to match actual struct fields ([e67b84b3](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/e67b84b3be2c284a7cfe8762c4404cf0bad669df))
- Use workflow_token for gh cli calls in release-branch-creation workflow ([662fd4c5](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/662fd4c51d6637c1e1a3db50e7fb5d19f4921aa0))
- Accept both v-prefixed and plain semver tags in deploy trigger ([684e843d](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/684e843d0965378905948ba44ba674818f742599))
- Remove /api/v1 prefix from @router annotations ([208a4c77](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/208a4c778ef2d8dd4d8dbf330286c8f860301985))
- Remove /api/v1 prefix from @router annotations ([a2e3f46d](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/a2e3f46d89240a8aa1117b98024b2da9555eed0d))
- Remove /api/v1 prefix from @router annotations ([d56aeba8](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/d56aeba895c8d624563c0d166fc952540a92196f))
- Remove /api/v1 prefix from @router annotations ([9e1d85e8](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/9e1d85e863c9de9ca8da654f0b5b37a8477449de))
- Remove /api/v1 prefix from @router annotations ([fdac8c4a](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/fdac8c4a8576d974f2f1e75723698f641fc33856))
- Remove /api/v1 prefix from @router annotations ([d6ba522d](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/d6ba522dca902246bfee5f4276a79f9de6a410ef))
- Use workflow_pat org secret instead of workflow_token ([d6d932c6](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/d6d932c6a295afd69a03b554a38cb5a18236bde4))
- Use workflow_token for gh cli calls in release-branch-creation workflow ([f8d618eb](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/f8d618eb7617b32b11f333fbd92e759570417977))

### Documentation

- Update changelog for 2.6.1 ([cf610cad](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/cf610cada7b0caad0ed5b0d726586c256cf35cc3))
- Regenerate swagger with fixed paths (no double /api/v1) ([30d9fd5f](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/30d9fd5f4d0915733f88c31e09944c04d58a174d))
- Regenerate swagger with fixed paths (no double /api/v1) ([9b4bcee4](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/9b4bcee4c7426ce654ba3e7dba0d888b6103862d))
- Regenerate swagger with fixed paths (no double /api/v1) ([a6598347](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/a6598347595d16eaf622913298f458fc6df3368f))
- Update changelog for v2.6.0 ([ebde3afb](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/ebde3afb5b2efc83aab6400d559bc710e39fa99a))

### Maintenance

- Bump version to v2.7.0 for next development cycle ([be91838f](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/be91838fd5faf918e67f86923b142a650338d4e5))
- Bump version to 2.6.1 ([520cf6d7](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/520cf6d7c099df460bf3da25e4c59294645fddf9))
- Bump version to v2.7.0 for next development cycle ([63563362](https://x-access-token:ghs_xElwuof11GZrjZv1U555gtZHuHLXen0hJCDx@github.com/banua-coder/pico-api-go/commit/63563362f07beac63049bb2759044a323073a809))

## [v2.6.0] - 2026-03-16

### Added

- Add pagination to hospitals, vaccination, regencies, and task-force endpoints ([8dd09036](https://x-access-token:ghs_20RzSAyPWuWvofxgjetv5CUGy6X8Vg3HsG1Q@github.com/banua-coder/pico-api-go/commit/8dd09036743bbc22a65040f979cb248b3da56c9c))

### Fixed

- Handle errcheck lint issues in getpaginatedbyprovinceid ([057a22ee](https://x-access-token:ghs_20RzSAyPWuWvofxgjetv5CUGy6X8Vg3HsG1Q@github.com/banua-coder/pico-api-go/commit/057a22ee182d30c4f530fcbe14f14c5610d620a3))
- Use update-version.sh script for develop version bump ([750a7e7f](https://x-access-token:ghs_20RzSAyPWuWvofxgjetv5CUGy6X8Vg3HsG1Q@github.com/banua-coder/pico-api-go/commit/750a7e7fdf4e1c2d40752738ce21c58888bd8e40))

### Documentation

- Update changelog for v2.5.1 ([59992313](https://x-access-token:ghs_20RzSAyPWuWvofxgjetv5CUGy6X8Vg3HsG1Q@github.com/banua-coder/pico-api-go/commit/59992313650247aab20b9f1a4d385f48f6905681))

### Tests

- Push coverage to 82.3% (above 80% ci threshold) ([988ccc5a](https://x-access-token:ghs_20RzSAyPWuWvofxgjetv5CUGy6X8Vg3HsG1Q@github.com/banua-coder/pico-api-go/commit/988ccc5a9523867d91eae3b96b41c2ad560441dc))
- Push code coverage from 57% to 80.5% ([adc36b55](https://x-access-token:ghs_20RzSAyPWuWvofxgjetv5CUGy6X8Vg3HsG1Q@github.com/banua-coder/pico-api-go/commit/adc36b55e3a59d930d552c072590a411ba8d79fb))
- Improve code coverage from 47% to 61.6% ([63552acc](https://x-access-token:ghs_20RzSAyPWuWvofxgjetv5CUGy6X8Vg3HsG1Q@github.com/banua-coder/pico-api-go/commit/63552acc69b053b79f8469506880ae47d6a50703))

### CI/CD

- Exclude cmd, docs, pkg/database from coverage calculation ([aba5f170](https://x-access-token:ghs_20RzSAyPWuWvofxgjetv5CUGy6X8Vg3HsG1Q@github.com/banua-coder/pico-api-go/commit/aba5f17038916e4ca1de175aabf04a71d6576dec))

### Maintenance

- Correct version to 2.6.0 (not 4.6.0) ([6fa7209d](https://x-access-token:ghs_20RzSAyPWuWvofxgjetv5CUGy6X8Vg3HsG1Q@github.com/banua-coder/pico-api-go/commit/6fa7209def16d7a32ffd2f15caca86bba1e69af6))
- Bump version to 4.6.0 and fix lint errors ([f6a1bd1a](https://x-access-token:ghs_20RzSAyPWuWvofxgjetv5CUGy6X8Vg3HsG1Q@github.com/banua-coder/pico-api-go/commit/f6a1bd1a50c5714b91d8942edba74c8522856666))
- Bump version to v2.6.0 for next development cycle ([85e547b2](https://x-access-token:ghs_20RzSAyPWuWvofxgjetv5CUGy6X8Vg3HsG1Q@github.com/banua-coder/pico-api-go/commit/85e547b23618c91616471f302328ec2a0e425d6a))

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
