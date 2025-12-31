# Skills Directory

This directory contains modular skill definitions for AI agents working with the percent project. Each skill is organized in its own subdirectory with a `SKILL.md` file that includes YAML frontmatter for configuration and markdown content for instructions.

## Structure

```
.github/skills/
├── development/          # Development environment and tools
│   └── SKILL.md
├── testing/             # Testing capabilities and frameworks
│   └── SKILL.md
├── build-deployment/    # Build and release management
│   └── SKILL.md
├── security/            # Security scanning and policies
│   └── SKILL.md
└── cicd/                # CI/CD workflows and automation
    └── SKILL.md
```

## Skill Format

Each `SKILL.md` file follows a two-part structure:

### 1. YAML Frontmatter (Configuration)

The frontmatter configures **HOW** the skill runs:

```yaml
---
name: skill-name
description: Brief overview of the skill
allowed-tools: "bash, view, edit, grep, glob"
version: 1.0.0
---
```

### 2. Markdown Content (Instructions)

The markdown content tells the AI agent **WHAT** to do:

- Purpose and objectives
- Detailed instructions
- Commands and usage
- Examples and guidelines
- Step-by-step procedures
- References to related documentation

## Available Skills

### Development (`development/`)

**Purpose**: Development environment setup, dependency management, and code quality tools

**Key Capabilities**:
- Bootstrap and setup workspace
- Manage Go modules and vendoring
- Code formatting, vetting, and linting
- Security vulnerability scanning

**Commands**: `make bootstrap`, `make setup`, `make go-check`, `make go-vuln`

### Testing (`testing/`)

**Purpose**: Comprehensive testing capabilities including unit, fuzz, and benchmark tests

**Key Capabilities**:
- Unit testing with race detection
- Coverage-guided fuzz testing
- Performance benchmarking
- Coverage analysis

**Commands**: `make go-test-unit`, `make go-test-fuzz`, `make go-test-bench`, `make go-test-coverage`

### Build and Deployment (`build-deployment/`)

**Purpose**: Build management and automated release processes

**Key Capabilities**:
- Module management and vendoring
- Semantic versioning
- Automated releases
- Dependency updates

**Commands**: `make go-mod-tidy`, `make go-mod-vendor`

**Automation**: Semantic-Release, Renovate, Dependabot

### Security (`security/`)

**Purpose**: Security scanning, vulnerability management, and policy enforcement

**Key Capabilities**:
- Vulnerability scanning with govulncheck and Trivy
- SBOM generation
- Policy-as-Code enforcement
- License analysis

**Commands**: `make go-vuln`, `make sast-trivy-fs`, `make policy-analysis-conftest`

### CI/CD (`cicd/`)

**Purpose**: Continuous integration and deployment workflows

**Key Capabilities**:
- GitHub Actions workflows
- Automated testing and linting
- Security scanning
- Automated releases

**Workflows**: Testing, linting, security, policy validation, releases

## Usage

AI agents can reference specific skills by navigating to the appropriate subdirectory and reading the `SKILL.md` file. Each skill is self-contained with all necessary information to perform its designated tasks.

## Relationship to Other Documentation

- **[AGENTS.md](../../AGENTS.md)**: Provides detailed testing patterns and templates (HOW to test)
- **[README.md](../../README.md)**: User-facing documentation
- **[.github/copilot-instructions.md](../copilot-instructions.md)**: GitHub Copilot specific instructions
- **[Makefile](../../Makefile)**: Complete list of available commands

## Version

All skills are currently at version 1.0.0. Versions will be incremented as skills evolve and capabilities are added.
