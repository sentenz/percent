# SKILLS.md

Skills index for the percent project. This document provides an overview of available skills and links to detailed skill definitions.

## Purpose

This index helps AI agents and developers discover and navigate the modular skill system. Each skill is self-contained with YAML frontmatter configuration and detailed markdown instructions.

## Available Skills

### 🛠️ [Development](./development/SKILL.md)

Development environment setup, dependency management, and code quality tools.

**Key Capabilities**:
- Bootstrap and setup workspace (`make bootstrap`, `make setup`)
- Manage Go modules (`make go-mod-tidy`, `make go-mod-vendor`)
- Code quality checks (`make go-fmt`, `make go-vet`, `make go-check`)
- Security scanning (`make go-vuln`)

**Version**: 1.0.0

---

### 🧪 [Testing](./testing/SKILL.md)

Comprehensive testing capabilities including unit tests, fuzz tests, benchmarks, and coverage analysis.

**Key Capabilities**:
- Unit testing with race detection (`make go-test-unit`)
- Coverage-guided fuzz testing (`make go-test-fuzz`)
- Performance benchmarking (`make go-test-bench`)
- Coverage analysis (`make go-test-coverage`)

**Test Patterns**: In-Got-Want, Table-Driven, Arrange-Act-Assert

**Version**: 1.0.0

---

### 🏗️ [Build and Deployment](./build-deployment/SKILL.md)

Build management, semantic versioning, and automated release processes.

**Key Capabilities**:
- Module management and vendoring
- Semantic versioning with Conventional Commits
- Automated releases with Semantic-Release
- Dependency updates (Renovate, Dependabot)

**Versioning**: SemVer 2.0.0

**Version**: 1.0.0

---

### 🔒 [Security](./security/SKILL.md)

Security scanning, vulnerability management, SBOM generation, and policy enforcement.

**Key Capabilities**:
- Vulnerability scanning (`make go-vuln`, `make sast-trivy-fs`)
- SBOM generation (CycloneDX format)
- Policy-as-Code enforcement (`make policy-analysis-conftest`)
- License analysis

**Tools**: govulncheck, Trivy, Conftest, Regal

**Version**: 1.0.0

---

### 🔄 [CI/CD](./cicd/SKILL.md)

Continuous integration and deployment workflows using GitHub Actions.

**Key Capabilities**:
- Automated testing and linting
- Security scanning workflows
- Policy validation
- Automated releases
- Dependency update automation

**Workflows**: go-tests, golangci, semgrep, conftest, regal, semantic-release, renovate

**Version**: 1.0.0

---

## Skill Structure

Each skill follows this format:

```
.github/skills/<skill-name>/
└── SKILL.md
    ├── YAML Frontmatter (Configuration)
    │   ├── name
    │   ├── description
    │   ├── allowed-tools
    │   └── version
    └── Markdown Content (Instructions)
        ├── Purpose
        ├── Detailed instructions
        ├── Commands and usage
        ├── Examples
        └── References
```

## Quick Access

- **Setup Development**: [development/SKILL.md](./development/SKILL.md)
- **Run Tests**: [testing/SKILL.md](./testing/SKILL.md)
- **Build & Release**: [build-deployment/SKILL.md](./build-deployment/SKILL.md)
- **Security Scanning**: [security/SKILL.md](./security/SKILL.md)
- **CI/CD Workflows**: [cicd/SKILL.md](./cicd/SKILL.md)

## Related Documentation

- **[AGENTS.md](../../AGENTS.md)**: AI agent instructions for testing patterns (HOW to test)
- **[README.md](../../README.md)**: User-facing documentation
- **[.github/copilot-instructions.md](../copilot-instructions.md)**: GitHub Copilot instructions
- **[Makefile](../../Makefile)**: Complete list of available commands
- **[README.md](./README.md)**: Detailed skills directory documentation

## Usage for AI Agents

AI agents should:
1. Navigate to the appropriate skill directory
2. Read the `SKILL.md` file
3. Follow the YAML configuration for tool permissions
4. Execute the markdown instructions

## Project Overview

**Project**: Percent - Go package for percentage calculations

**Language**: Go 1.22.0+

**Architecture**: Standard Go Project Layout
- `pkg/percent/` - Public API
- `internal/pkg/resource/` - Internal resources
- Tests alongside source code

**Quality**: 100% code coverage with comprehensive testing

## Version History

- **1.0.0** (2025-12-31): Initial skill modularization
  - Created individual skill directories
  - Added YAML frontmatter configuration
  - Separated concerns by capability
