---
name: build-deployment
description: Build management and release automation for the Go percent package using semantic versioning
allowed-tools: bash, view, edit, grep
version: 1.0.0
---

# Build and Deployment Skill

This skill provides instructions and capabilities for building, versioning, and releasing the percent package.

## Purpose

Enable developers and AI agents to:
- Build and verify Go modules
- Execute automated releases
- Manage semantic versioning
- Generate changelogs

## Module Management

### Build Commands

#### Tidy Modules

Clean up dependencies and update module files:

```bash
go mod tidy
```

Or use the Makefile target:

```bash
make go-mod-tidy
```

This command:
- Adds missing module dependencies
- Removes unused dependencies
- Updates `go.mod` and `go.sum`

#### Vendor Dependencies

Create a local copy of all dependencies:

```bash
go mod vendor
```

Or use the Makefile target:

```bash
make go-mod-vendor
```

This command:
- Downloads dependencies to `vendor/` directory
- Ensures reproducible builds
- Creates offline-capable builds

#### Verify Modules

Verify dependency checksums:

```bash
go mod verify
```

This ensures the integrity of downloaded dependencies.

### Test Execution

Execute all tests:

```bash
go test ./...
```

With race detection:

```bash
go test -race ./...
```

With coverage:

```bash
go test -cover ./...
```

## Release Management

### Semantic Release

The project uses [Semantic-Release](https://github.com/semantic-release/semantic-release) for automated version management.

#### Features

- Automated version bumping based on commit messages
- Changelog generation
- GitHub release creation
- Tag creation

#### Configuration

Release configuration is defined in `.releaserc.json`.

#### Versioning

The project follows [Semantic Versioning (SemVer)](https://semver.org/):
- **MAJOR**: Incompatible API changes
- **MINOR**: New functionality (backward compatible)
- **PATCH**: Bug fixes (backward compatible)

#### Commit Convention

Uses [Conventional Commits](https://www.conventionalcommits.org/):

- `feat:` - New feature (triggers MINOR version bump)
- `fix:` - Bug fix (triggers PATCH version bump)
- `BREAKING CHANGE:` - Breaking change (triggers MAJOR version bump)
- `chore:` - Maintenance tasks (no version bump)
- `docs:` - Documentation changes (no version bump)
- `test:` - Test changes (no version bump)

#### CI/CD Integration

Releases are automated through GitHub Actions workflow:

```yaml
uses: sentenz/actions/semantic-release@latest
```

The workflow:
1. Analyzes commit messages
2. Determines next version
3. Generates changelog
4. Creates GitHub release
5. Publishes release notes

## Dependency Updates

### Renovate

Automated dependency updates using [Renovate](https://github.com/renovatebot/renovate).

Configuration: `renovate.json`

Features:
- Automatic pull requests for outdated dependencies
- Security vulnerability updates
- Version upgrade proposals

CI/CD workflow:
```yaml
uses: sentenz/actions/renovate@latest
```

### Dependabot

Alternative dependency management with [Dependabot](https://github.com/dependabot/dependabot-core).

Configuration: `.github/dependabot.yml`

Features:
- Automated security updates
- Dependency version updates
- Pull request creation

## Build Artifacts

The project generates the following artifacts:

### Test Reports
- `logs/unit/junit-report.xml` - JUnit test results
- `logs/coverage/coverage.xml` - Cobertura coverage report
- `logs/coverage/coverage.html` - HTML coverage report

### Benchmark Results
- `logs/bench/` - Benchmark comparison results

## Quality Gates

Before release, all checks must pass:
- ✅ All tests passing
- ✅ Coverage requirements met
- ✅ Linting passing
- ✅ Security scans clean
- ✅ No vulnerable dependencies

## References

- [README.md](../../../README.md) - Project documentation
- [Makefile](../../../Makefile) - Complete list of available commands
- [.releaserc.json](../../../.releaserc.json) - Semantic release configuration
- [renovate.json](../../../renovate.json) - Renovate configuration
