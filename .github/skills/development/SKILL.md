---
name: development
description: Development environment setup, dependency management, and code quality tools for the Go percent package
allowed-tools: bash, view, edit, grep, glob
version: 1.0.0
---

# Development Skill

This skill provides instructions and capabilities for setting up the development environment, managing dependencies, and maintaining code quality in the percent project.

## Purpose

Enable developers and AI agents to:
- Bootstrap and setup the development workspace
- Manage Go modules and dependencies
- Ensure code quality through formatting, vetting, and linting
- Scan for security vulnerabilities

## Environment Setup

### Bootstrap

Initialize a software development workspace with requisites:

```bash
make bootstrap
```

This command:
- Installs required tools
- Sets up pre-commit hooks
- Prepares the development environment

### Setup

Install and configure all dependencies essential for development:

```bash
make setup
```

This command:
- Installs dependencies
- Configures development environment
- Prepares vendored packages

### Teardown

Remove development artifacts and restore the host to its pre-setup state:

```bash
make teardown
```

This command:
- Removes development artifacts
- Cleans build outputs
- Restores to pre-setup state

## Dependency Management

### Go Modules

Tidy Go modules to add missing and remove unused dependencies:

```bash
make go-mod-tidy
```

This updates `go.mod` and `go.sum` files.

### Vendoring

Vendor Go modules to create a local copy of all dependencies:

```bash
make go-mod-vendor
```

This downloads dependencies to the `vendor/` directory and ensures reproducible builds.

## Code Quality

### Formatting

Format Go source files using standard Go formatting:

```bash
make go-fmt
```

Uses `gofmt` with standard settings to ensure consistent code style.

### Vetting

Examine Go source code and report suspicious constructs:

```bash
make go-vet
```

Catches common mistakes and potential issues.

### Linting

Run golangci-lint with project configuration:

```bash
make go-check
```

Enforces code quality standards and checks for potential issues based on `.golangci.yml` configuration.

### Security Scanning

Scan for known vulnerabilities using govulncheck:

```bash
make go-vuln
```

Reports security issues from the Go vulnerability database.

## Prerequisites

- Go 1.22.0 or later
- Make (GNU Make)
- Git

## References

- [README.md](../../../README.md) - Project documentation
- [Makefile](../../../Makefile) - Complete list of available commands
- [scripts/](../../../scripts/) - Bootstrap, setup, and teardown scripts
