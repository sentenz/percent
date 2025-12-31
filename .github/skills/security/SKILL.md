---
name: security
description: Security scanning and vulnerability management including govulncheck, Trivy, and policy enforcement for the Go percent package
allowed-tools: bash, view, edit, grep
version: 1.0.0
---

# Security Skill

This skill provides instructions and capabilities for security scanning, vulnerability management, and policy enforcement in the percent project.

## Purpose

Enable developers and AI agents to:
- Scan for known vulnerabilities in dependencies
- Analyze code for security issues
- Generate Software Bill of Materials (SBOM)
- Enforce security policies
- Validate compliance standards

## Vulnerability Scanning

### Govulncheck

Scan Go dependencies for known vulnerabilities using the official Go vulnerability database.

```bash
make go-vuln
```

Features:
- Scans all Go dependencies
- Reports known vulnerabilities
- Uses official Go vulnerability database
- Integration with development workflow

Output: Security vulnerabilities found in dependencies

### Trivy

Comprehensive security scanner for vulnerabilities, misconfigurations, and compliance issues.

#### Filesystem Scan

Scan the project filesystem for security issues:

```bash
make sast-trivy-fs <path>
```

Features:
- Scans dependencies
- Checks configurations
- Detects secrets
- Identifies misconfigurations

Configuration: `trivy.yaml`

#### SBOM Generation

Generate Software Bill of Materials in CycloneDX format:

```bash
make sast-trivy-sbom-cyclonedx-fs <path>
```

Features:
- Creates comprehensive SBOM
- CycloneDX format
- Includes all dependencies
- License information

#### SBOM Analysis

Analyze existing SBOM for vulnerabilities:

```bash
make sast-trivy-sbom <sbom_path>
```

Scans the SBOM for known vulnerabilities.

#### License Analysis

Analyze SBOM for license information:

```bash
make sast-trivy-sbom-license <sbom_path>
```

Features:
- License identification
- Compliance checking
- Risk assessment

### Configuration

#### Trivy Configuration

Location: `trivy.yaml`

Specifies:
- Scan settings
- Severity thresholds
- Ignored vulnerabilities

#### Trivy Ignore

Location: `.trivyignore`

Lists vulnerabilities to ignore during scans (with justification).

## Policy Management

### Conftest

Policy-as-Code (PaC) tool for enforcing best practices and compliance standards using Rego policies.

#### Run Policy Tests

Validate configurations against defined policies:

```bash
make policy-analysis-conftest <filepath>
```

Features:
- Validates against Rego policies
- Enforces compliance
- Checks best practices

Configuration: `conftest.toml`

Policies location: `tests/policy/`

### Regal

Lint Rego policies to ensure policy quality.

```bash
make policy-lint-regal <filepath>
```

Features:
- Validates Rego syntax
- Ensures policy quality
- Checks for common mistakes

## CI/CD Integration

### GitHub Actions Workflows

Security scanning is integrated into CI/CD:

#### Trivy Workflow
```yaml
uses: sentenz/actions/trivy@latest
```

#### Conftest Workflow
```yaml
uses: sentenz/actions/conftest@latest
```

#### Regal Workflow
```yaml
uses: sentenz/actions/regal@latest
```

#### Semgrep Workflow
Security analysis workflow: `.github/workflows/semgrep.yml`

Static analysis security testing to identify vulnerabilities and review code patterns.

## Security Best Practices

### Dependency Management

- Keep dependencies up to date
- Review security advisories
- Use dependabot/renovate for automated updates
- Scan regularly with govulncheck

### Secret Management

- Never commit secrets to source code
- Use environment variables
- Leverage GitHub Secrets
- Scan for exposed secrets with Trivy

### Vulnerability Response

1. Identify vulnerability (automated scanning)
2. Assess impact and severity
3. Update affected dependency
4. Verify fix with security scans
5. Document in `.trivyignore` if accepting risk

## Security Reports

The project generates various security reports:

- Vulnerability scan results
- SBOM (Software Bill of Materials)
- License compliance reports
- Policy validation results

## References

- [README.md](../../../README.md) - Project documentation
- [Makefile](../../../Makefile) - Complete list of available commands
- [trivy.yaml](../../../trivy.yaml) - Trivy configuration
- [.trivyignore](../../../.trivyignore) - Ignored vulnerabilities
- [conftest.toml](../../../conftest.toml) - Conftest configuration
- [tests/policy/](../../../tests/policy/) - Rego policies
