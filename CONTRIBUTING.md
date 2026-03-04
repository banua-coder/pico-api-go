# Contributing — Git Flow

This project follows **Git Flow** branching strategy.

## Branch Structure

```
main        ← production (stable releases only)
develop     ← integration branch (all features merged here)
feature/*   ← new features
bugfix/*    ← bug fixes
release/*   ← release preparation
hotfix/*    ← critical production fixes
```

## Workflow

### Feature / Bugfix
```bash
git checkout develop
git checkout -b feature/your-feature-name
# or
git checkout -b bugfix/your-bug-fix

# ... make changes ...

git push -u origin feature/your-feature-name
# Open PR → target: develop
```

### Release
```bash
git checkout develop
git checkout -b release/v2.5.0

# Bump version, update CHANGELOG, final fixes only
# Open PR → target: main
# After merge, back-merge main → develop
```

### Hotfix
```bash
git checkout main
git checkout -b hotfix/critical-fix

# Fix the issue
# Open PR → target: main
# After merge, back-merge main → develop
```

## CI/CD
- Push to `main` → triggers Docker build + push to GHCR → Watchtower auto-deploys
- Push to `develop` → triggers CI checks only (no deploy)
- All other branches → CI checks only
