---
name: cicd
description: CI/CD workflows and automation using GitHub Actions for the Go percent package
allowed-tools: bash, view, edit, grep
version: 1.0.0
---

# CI/CD Skill

This skill provides instructions and capabilities for understanding and working with the CI/CD workflows and automation in the percent project.

## Purpose

Enable developers and AI agents to:
- Understand CI/CD workflows
- Interact with GitHub Actions
- Troubleshoot workflow failures
- Manage automated processes

## GitHub Actions Workflows

All workflows are located in `.github/workflows/` directory.

### Testing Workflow

**File**: `go-tests.yml`

**Triggers**:
- Push to any branch
- Pull requests

**Actions**:
1. Checkout code
2. Setup Go environment
3. Run unit tests with race detection
4. Generate coverage reports
5. Upload test results as artifacts

**Integration**:
```yaml
uses: sentenz/actions/go-tests@latest
```

**Artifacts**:
- JUnit test reports
- Coverage reports (HTML and XML)

### Linting Workflow

**File**: `golangci.yml`

**Triggers**:
- Push to any branch
- Pull requests

**Actions**:
1. Checkout code
2. Setup Go environment
3. Run golangci-lint
4. Report linting issues

**Configuration**: `.golangci.yml`

### Security Workflow

**File**: `semgrep.yml`

**Triggers**:
- Push to any branch
- Pull requests

**Actions**:
1. Checkout code
2. Run Semgrep security analysis
3. Identify security vulnerabilities
4. Report findings

**Purpose**: Static analysis security testing (SAST)

### Policy Workflow

**File**: `conftest.yml`

**Triggers**:
- Push to any branch
- Pull requests

**Actions**:
1. Checkout code
2. Run Conftest policy validation
3. Validate against Rego policies
4. Report policy violations

**Configuration**: `conftest.toml`

**Policies**: `tests/policy/`

### Regal Workflow

**File**: `regal.yml`

**Triggers**:
- Push to any branch
- Pull requests

**Actions**:
1. Checkout code
2. Lint Rego policies
3. Validate policy syntax
4. Report issues

**Integration**:
```yaml
uses: sentenz/actions/regal@latest
```

### Release Workflow

**File**: `semantic-release.yml`

**Triggers**:
- Push to main branch

**Actions**:
1. Analyze commit messages
2. Determine next version
3. Generate changelog
4. Create GitHub release
5. Publish release notes

**Integration**:
```yaml
uses: sentenz/actions/semantic-release@latest
```

**Configuration**: `.releaserc.json`

### Dependency Update Workflow

**File**: `renovate.yml`

**Triggers**:
- Scheduled (defined in workflow)

**Actions**:
1. Check for dependency updates
2. Create pull requests for updates
3. Include security patches
4. Version upgrade proposals

**Integration**:
```yaml
uses: sentenz/actions/renovate@latest
```

**Configuration**: `renovate.json`

## Workflow Actions

### Reusable Actions

The project uses standardized reusable actions from `sentenz/actions`:

- `go-tests@latest` - Go testing with coverage
- `semantic-release@latest` - Automated releases
- `renovate@latest` - Dependency updates
- `trivy@latest` - Security scanning
- `conftest@latest` - Policy validation
- `regal@latest` - Rego policy linting

### Benefits

- Consistent CI/CD patterns across projects
- Centralized action maintenance
- Standardized workflows
- Easy updates

## Quality Gates

All pull requests must pass these checks:

1. ✅ **Unit Tests** - All tests passing
2. ✅ **Race Detection** - No race conditions
3. ✅ **Code Coverage** - Coverage requirements met
4. ✅ **Linting** - No linting errors
5. ✅ **Security Scan** - No security issues
6. ✅ **Policy Validation** - Policies enforced

## Branch Protection

Main branch is protected with:

- Required status checks
- Required code review
- No force push
- Linear history enforcement

## Workflow Monitoring

### View Workflow Runs

Check workflow status:
1. Navigate to repository on GitHub
2. Click "Actions" tab
3. View recent workflow runs

### Troubleshooting Failed Workflows

1. **Identify the failure**
   - Check workflow run logs
   - Identify failing step

2. **Common Issues**
   - Test failures
   - Linting errors
   - Security vulnerabilities
   - Policy violations

3. **Resolution**
   - Fix identified issue locally
   - Run checks locally before pushing
   - Push fix to trigger re-run

### Local Validation

Before pushing changes, run locally:

```bash
# Run tests
make go-test-unit

# Run linting
make go-check

# Run security scan
make go-vuln
```

## Automation Features

### Automated Processes

1. **Testing** - Automatic on every push/PR
2. **Linting** - Automatic code quality checks
3. **Security Scanning** - Continuous vulnerability monitoring
4. **Policy Enforcement** - Automated compliance checking
5. **Releases** - Automated version management
6. **Dependency Updates** - Automated update proposals

### Notifications

Workflow results are visible:
- GitHub Actions tab
- Pull request checks
- Commit status checks
- Email notifications (configurable)

## References

- [README.md](../../../README.md) - Project documentation
- [.github/workflows/](../../workflows/) - Workflow definitions
- [Makefile](../../../Makefile) - Local command equivalents
- [sentenz/actions](https://github.com/sentenz/actions) - Reusable actions repository
