# SKILLS.md

- [1. Project Overview](#1-project-overview)
  - [1.1. Purpose](#11-purpose)
  - [1.2. Architecture](#12-architecture)
- [2. Project Structure](#2-project-structure)
  - [2.1. Directory Layout](#21-directory-layout)
  - [2.2. Package Organization](#22-package-organization)
- [3. Development Skills](#3-development-skills)
  - [3.1. Environment Setup](#31-environment-setup)
  - [3.2. Dependency Management](#32-dependency-management)
  - [3.3. Code Quality](#33-code-quality)
- [4. Testing Skills](#4-testing-skills)
  - [4.1. Unit Testing](#41-unit-testing)
  - [4.2. Fuzz Testing](#42-fuzz-testing)
  - [4.3. Benchmark Testing](#43-benchmark-testing)
  - [4.4. Coverage Analysis](#44-coverage-analysis)
- [5. Build and Deployment Skills](#5-build-and-deployment-skills)
  - [5.1. Build Commands](#51-build-commands)
  - [5.2. Release Management](#52-release-management)
- [6. Security Skills](#6-security-skills)
  - [6.1. Vulnerability Scanning](#61-vulnerability-scanning)
  - [6.2. Policy Management](#62-policy-management)
- [7. CI/CD Skills](#7-cicd-skills)
  - [7.1. Workflows](#71-workflows)
  - [7.2. Automation](#72-automation)

## 1. Project Overview

### 1.1. Purpose

Percent is a Go package that provides utility functions for calculating percentages and performing related operations. The package is designed to be simple, reliable, and well-tested with 100% code coverage.

### 1.2. Architecture

The project follows the **Standard Go Project Layout** with clear separation between:
- **Public API** (`pkg/percent/`): Exported functions for percentage calculations
- **Internal Resources** (`internal/pkg/resource/`): Private constants, errors, and strings
- **Testing**: Comprehensive unit, fuzz, and benchmark tests
- **Scripts**: Bootstrap, setup, and teardown automation
- **Policies**: Rego-based policy enforcement

## 2. Project Structure

### 2.1. Directory Layout

```
.
├── pkg/
│   └── percent/              # Public API for percentage calculations
│       ├── percent.go        # Core percentage functions
│       └── percent_test.go   # Comprehensive test suite
├── internal/
│   └── pkg/
│       └── resource/         # Internal resources
│           ├── errors.go     # Error definitions
│           ├── strings.go    # String constants
│           └── constants.go  # Shared constants
├── scripts/                  # Development automation
│   ├── bootstrap.sh          # Initialize workspace
│   ├── setup.sh              # Install dependencies
│   └── teardown.sh           # Clean workspace
├── tests/
│   └── policy/               # Rego policy definitions
├── .github/
│   └── workflows/            # CI/CD pipelines
├── .devcontainer/            # Dev container configs
├── vendor/                   # Vendored dependencies
├── logs/                     # Test and coverage reports
├── Makefile                  # Task automation
├── AGENTS.md                 # AI agent instructions
├── .github/
│   └── skills/
│       └── SKILLS.md         # Capability documentation
└── README.md                 # User documentation
```

### 2.2. Package Organization

1. **Public Packages** (`pkg/`)
   - Exported API that external users can import
   - Follows semantic versioning
   - Fully documented with godoc comments
   - Example: `github.com/sentenz/percent/pkg/percent`

2. **Internal Packages** (`internal/`)
   - Private implementation details
   - Cannot be imported by external packages
   - Shared resources and utilities
   - Example: `internal/pkg/resource/errors.go`

3. **Test Organization**
   - Tests live alongside source files (`*_test.go`)
   - Test fixtures in `testdata/` directory
   - Fuzz corpus in `testdata/fuzz/` directory

## 3. Development Skills

### 3.1. Environment Setup

1. **Prerequisites**
   - Go 1.22.0 or later
   - Make (GNU Make)
   - Git

2. **Bootstrap**
   ```bash
   make bootstrap
   ```
   - Initializes development workspace
   - Installs required tools
   - Sets up pre-commit hooks

3. **Setup**
   ```bash
   make setup
   ```
   - Installs dependencies
   - Configures development environment
   - Prepares vendored packages

4. **Teardown**
   ```bash
   make teardown
   ```
   - Removes development artifacts
   - Cleans build outputs
   - Restores to pre-setup state

### 3.2. Dependency Management

1. **Go Modules**
   ```bash
   make go-mod-tidy
   ```
   - Adds missing module dependencies
   - Removes unused dependencies
   - Updates `go.mod` and `go.sum`

2. **Vendoring**
   ```bash
   make go-mod-vendor
   ```
   - Downloads dependencies to `vendor/` directory
   - Creates local copy of all dependencies
   - Ensures reproducible builds

### 3.3. Code Quality

1. **Formatting**
   ```bash
   make go-fmt
   ```
   - Formats Go source files
   - Uses `gofmt` with standard settings
   - Ensures consistent code style

2. **Vetting**
   ```bash
   make go-vet
   ```
   - Examines Go source code
   - Reports suspicious constructs
   - Catches common mistakes

3. **Linting**
   ```bash
   make go-check
   ```
   - Runs golangci-lint with configuration
   - Enforces code quality standards
   - Checks for potential issues

4. **Security**
   ```bash
   make go-vuln
   ```
   - Scans for known vulnerabilities
   - Uses govulncheck tool
   - Reports security issues

## 4. Testing Skills

### 4.1. Unit Testing

1. **Run Tests**
   ```bash
   make go-test-unit
   ```
   - Executes all unit tests
   - Enables race detection (`-race`)
   - Generates JUnit XML report
   - Output: `logs/unit/junit-report.xml`

2. **Test Organization**
   - Uses standard Go `testing` package
   - Table-driven test approach
   - In-Got-Want pattern
   - Arrange-Act-Assert structure

3. **Test Framework**
   - `testing.T` for unit tests
   - `github.com/google/go-cmp/cmp` for comparisons
   - `errors.Is()` for error checking

### 4.2. Fuzz Testing

1. **Run Fuzz Tests**
   ```bash
   make go-test-fuzz
   ```
   - Executes coverage-guided fuzzing
   - Tests: `FuzzPercent`, `FuzzOf`, `FuzzChange`, `FuzzRemain`, `FuzzFromRatio`, `FuzzToRatio`
   - Duration: 10 seconds per function
   - Automatically explores code paths

2. **Fuzz Framework**
   - `testing.F` for fuzz tests
   - Coverage-guided input generation
   - Corpus-driven mutations
   - Automatic crash detection

3. **Corpus Management**
   - Seed corpus in `testdata/fuzz/<FuzzName>/`
   - Generated inputs saved on crash
   - Regression prevention

### 4.3. Benchmark Testing

1. **Run Benchmarks**
   ```bash
   make go-test-bench
   ```
   - Measures execution time
   - Tracks memory allocations
   - Reports operations per second

2. **Compare Benchmarks**
   ```bash
   make go-test-bench-compare
   ```
   - Runs benchmarks 10 times
   - Compares before/after results
   - Uses `benchstat` for analysis
   - Output: `logs/bench/`

3. **Benchmark Framework**
   - `testing.B` for benchmarks
   - `b.Loop()` for iterations
   - Memory profiling with `-benchmem`
   - Sub-benchmarks with `b.Run()`

### 4.4. Coverage Analysis

1. **Generate Coverage**
   ```bash
   make go-test-coverage
   ```
   - Runs tests with coverage tracking
   - Generates HTML report
   - Creates Cobertura XML
   - Output: `logs/coverage/`

2. **Coverage Reports**
   - `coverage.out`: Raw coverage data
   - `coverage.html`: Interactive HTML view
   - `coverage.xml`: Cobertura format for CI/CD

## 5. Build and Deployment Skills

### 5.1. Build Commands

1. **Module Management**
   - `go mod tidy`: Clean up dependencies
   - `go mod vendor`: Vendor dependencies
   - `go mod verify`: Verify checksums

2. **Test Execution**
   - `go test ./...`: Run all tests
   - `go test -race ./...`: Enable race detector
   - `go test -cover ./...`: Measure coverage

3. **Benchmark Execution**
   - `go test -bench=.`: Run all benchmarks
   - `go test -bench=BenchmarkName`: Run specific benchmark
   - `go test -benchmem`: Include memory stats

### 5.2. Release Management

1. **Semantic Release**
   - Automated version management
   - Changelog generation
   - GitHub release creation
   - Configuration: `.releaserc.json`

2. **Versioning**
   - Follows semantic versioning (SemVer)
   - Based on conventional commits
   - Automated by CI/CD

## 6. Security Skills

### 6.1. Vulnerability Scanning

1. **Govulncheck**
   ```bash
   make go-vuln
   ```
   - Scans Go dependencies
   - Reports known vulnerabilities
   - Uses Go vulnerability database

2. **Trivy**
   ```bash
   make sast-trivy-fs <path>
   ```
   - Comprehensive security scanner
   - Checks dependencies, configs, secrets
   - Generates SBOM (CycloneDX format)

3. **SBOM Generation**
   ```bash
   make sast-trivy-sbom-cyclonedx-fs <path>
   ```
   - Creates Software Bill of Materials
   - CycloneDX format
   - License information

### 6.2. Policy Management

1. **Conftest**
   ```bash
   make policy-analysis-conftest <filepath>
   ```
   - Policy-as-Code enforcement
   - Rego policy language
   - Configuration validation

2. **Regal**
   ```bash
   make policy-lint-regal <filepath>
   ```
   - Lints Rego policies
   - Ensures policy quality
   - Validates policy syntax

## 7. CI/CD Skills

### 7.1. Workflows

1. **Testing Workflow** (`go-tests.yml`)
   - Runs on push and pull requests
   - Executes unit tests with coverage
   - Generates test reports
   - Uploads coverage artifacts

2. **Linting Workflow** (`golangci.yml`)
   - Runs golangci-lint
   - Enforces code quality
   - Catches potential issues

3. **Security Workflow** (`semgrep.yml`)
   - Static analysis security testing
   - Identifies security vulnerabilities
   - Reviews code patterns

4. **Policy Workflow** (`conftest.yml`)
   - Validates policies
   - Enforces compliance
   - Runs Rego policies

5. **Release Workflow** (`semantic-release.yml`)
   - Automated releases
   - Version bumping
   - Changelog generation

6. **Dependency Updates** (`renovate.yml`)
   - Automated dependency updates
   - Security patches
   - Version upgrades

### 7.2. Automation

1. **Actions Integration**
   - Uses `sentenz/actions` reusable workflows
   - Standardized CI/CD patterns
   - Consistent across projects

2. **Quality Gates**
   - All tests must pass
   - Coverage requirements met
   - Linting passes
   - Security scans clean

3. **Branch Protection**
   - Required status checks
   - Code review required
   - No force push
   - Linear history

## References

- [AGENTS.md](../../AGENTS.md) - AI agent instructions for testing patterns
- [README.md](../../README.md) - User-facing documentation
- [Makefile](../../Makefile) - Complete list of available commands
- [.github/copilot-instructions.md](../copilot-instructions.md) - GitHub Copilot instructions
