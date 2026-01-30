# AGENTS.md

- [1. Tech Stack](#1-tech-stack)
  - [1.1. Programming Language](#11-programming-language)
  - [1.2. Dependency Manager](#12-dependency-manager)
  - [1.3. Test Framework](#13-test-framework)
  - [1.4. Build System](#14-build-system)
- [2. Project Layout](#2-project-layout)
  - [2.1. Directory Structure](#21-directory-structure)
  - [2.2. Key Directories](#22-key-directories)
- [3. Task Runners](#3-task-runners)
  - [3.1. Make Commands](#31-make-commands)
    - [3.1.1. Setup & Teardown](#311-setup--teardown)
    - [3.1.2. Go Module Management](#312-go-module-management)
    - [3.1.3. Go Testing](#313-go-testing)
    - [3.1.4. Go Code Quality](#314-go-code-quality)
    - [3.1.5. Policy & Security](#315-policy--security)

## 1. Tech Stack

The technology stack for this Go project includes the tools and frameworks used for development, testing, and deployment.

### 1.1. Programming Language

- Go (Golang)
  > A statically typed, compiled programming language designed for simplicity, efficiency, and concurrency.
  >
  > Version: Go 1.22.0 or higher

### 1.2. Dependency Manager

- [Go Modules](https://go.dev/ref/mod)
  > Go's official dependency management system that handles versioning and package dependencies through `go.mod` and `go.sum` files.
  >
  > Commands:
  > - `go mod tidy` - Add missing and remove unused modules
  > - `go mod vendor` - Make vendored copy of dependencies
  > - `go mod download` - Download modules to local cache

### 1.3. Test Framework

- [Testing](https://pkg.go.dev/testing)
  > Go's built-in testing framework for unit tests, benchmarks, and fuzz tests.
  >
  > Supports:
  > - Unit tests with table-driven patterns
  > - Benchmark tests for performance measurement
  > - Fuzz tests with coverage-guided fuzzing
  > - Code coverage analysis

### 1.4. Build System

- [Make](https://www.gnu.org/software/make/)
  > Task runner and build automation tool using Makefile to define project tasks and commands.
  >
  > Used for: build automation, testing, linting, and deployment tasks

## 2. Project Layout

Standard Go project layout following community best practices for organizing code and resources.

### 2.1. Directory Structure

```
.
├── .github/            # GitHub workflows, actions, and configuration
│   ├── skills/         # Agent skills for AI coding assistants
│   └── workflows/      # GitHub Actions CI/CD workflows
├── internal/           # Private application and library code
│   └── pkg/            # Internal packages
│       └── resource/   # Internal resources (errors, constants, strings)
├── pkg/                # Public library code (importable by external projects)
│   └── percent/        # Main package with percentage utilities
├── scripts/            # Scripts for build, setup, and deployment
├── tests/              # Additional test files and test data
│   └── policy/         # Policy tests with Conftest/Rego
├── vendor/             # Vendored dependencies (managed by Go modules)
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── Makefile            # Task runner definitions
└── AGENTS.md           # AI agent instructions and guidelines
```

### 2.2. Key Directories

- `pkg/`
  > Contains packages that can be imported by external projects. This is the public API of the module.

- `internal/`
  > Contains private code that cannot be imported by external projects. Go enforces this restriction.

- `internal/pkg/resource/`
  > Shared internal resources like error definitions, constants, and string templates used across packages.

- `.github/skills/`
  > Agent skills documentation following the Agent Skills specification for AI coding assistants.

- `tests/policy/`
  > Policy-as-code tests using Conftest and Rego for validating configurations.

- `scripts/`
  > Automation scripts for bootstrapping, setup, and teardown of the development environment.

- `vendor/`
  > Local copy of dependencies for reproducible builds and offline development.

## 3. Task Runners

Make is used as the task runner to automate common development tasks.

### 3.1. Make Commands

Run `make help` to see all available commands with descriptions.

#### 3.1.1. Setup & Teardown

- `make bootstrap`
  > Initialize a software development workspace with requisites.
  >
  > Runs: `scripts/bootstrap.sh`

- `make setup`
  > Install and configure all dependencies essential for development.
  >
  > Runs: `scripts/setup.sh`

- `make teardown`
  > Remove development artifacts and restore the host to its pre-setup state.
  >
  > Runs: `scripts/teardown.sh`

#### 3.1.2. Go Module Management

- `make go-mod-tidy`
  > Tidy Go modules by adding missing and removing unused modules.
  >
  > Runs: `go mod tidy`

- `make go-mod-vendor`
  > Vendor Go modules to create a local copy of dependencies.
  >
  > Runs: `go mod vendor`

#### 3.1.3. Go Testing

- `make go-test-unit`
  > Run Go unit tests with race detection and generate JUnit XML report.
  >
  > Flags: `-race`, `-coverprofile`, `-v`

- `make go-test-coverage`
  > Run Go tests with coverage report (HTML and XML formats).
  >
  > Generates: `coverage.html`, `coverage.xml`

- `make go-test-bench`
  > Run Go benchmarks to measure performance and memory usage.
  >
  > Flags: `-bench=.`, `-benchmem`

- `make go-test-fuzz`
  > Run fuzz tests with Go's native fuzzing engine.
  >
  > Duration: Configurable via environment variable

#### 3.1.4. Go Code Quality

- `make go-fmt`
  > Format Go code according to Go standards using `gofmt`.
  >
  > Runs: `gofmt -s -w`

- `make go-fix`
  > Update Go source files to use new APIs with `go fix`.

- `make go-vet`
  > Check Go code for common mistakes using `go vet`.

- `make go-vuln`
  > Check Go code for known vulnerabilities using `govulncheck`.

- `make go-check`
  > Run all Go code quality checks (fmt, vet, vuln).
  >
  > Runs: `go-fmt`, `go-vet`, `go-vuln`

#### 3.1.5. Policy & Security

- `make conftest-test`
  > Analyze configuration files using Conftest for policy violations.

- `make regal-lint`
  > Lint Rego policies using Regal.

- `make trivy-scan`
  > Scan Infrastructure-as-Code files for misconfigurations using Trivy.

