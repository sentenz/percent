# Percent

[![GoDoc](https://godoc.org/github.com/sentenz/percent?status.svg)](https://godoc.org/github.com/sentenz/percent/pkg/percent)
[![Go Report Card](https://goreportcard.com/badge/github.com/sentenz/percent)](https://goreportcard.com/report/github.com/sentenz/percent)
[![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)](https://github.com/sentenz/percent/actions/workflows/go-tests.yml)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Percent is a Go package that provides utility functions for calculating percentages and performing related operations.

- [1. Module](#1-module)
  - [1.1. Prerequisites](#11-prerequisites)
  - [1.2. Installation](#12-installation)
  - [1.3. Usage](#13-usage)
- [2. Contribute](#2-contribute)
  - [2.1. Task Runner](#21-task-runner)
    - [2.1.1. Make](#211-make)
  - [2.2. Bootstrap](#22-bootstrap)
    - [2.2.1. Scripts](#221-scripts)
  - [2.3. Dev Containers](#23-dev-containers)
  - [2.4. Dependency Manager](#24-dependency-manager)
    - [2.4.1. Go Modules](#241-go-modules)
  - [2.5. Software Testing](#25-software-testing)
    - [2.5.1. Unit Testing](#251-unit-testing)
    - [2.5.2. Fuzz Testing](#252-fuzz-testing)
    - [2.5.3. Benchmarks](#253-benchmarks)
    - [2.5.4. Code Coverage](#254-code-coverage)
  - [2.6. Release Manager](#26-release-manager)
    - [2.6.1. Semantic-Release](#261-semantic-release)
  - [2.7. Update Manager](#27-update-manager)
    - [2.7.1. Renovate](#271-renovate)
    - [2.7.2. Dependabot](#272-dependabot)
  - [2.8. Policy Manager](#28-policy-manager)
    - [2.8.1. Conftest](#281-conftest)
  - [2.9. Supply Chain Manager](#29-supply-chain-manager)
    - [2.9.1. Trivy](#291-trivy)
- [3. References](#3-references)

## 1. Module

### 1.1. Prerequisites

- [Go](https://golang.org/)
  > Go programming language environment for building and running Go applications.

### 1.2. Installation

- Install
  > Install the package via `go get`.

  ```bash
  go get github.com/sentenz/percent
  ```

- Import
  > Import the package in a called Go source file.

  ```go
  import "github.com/sentenz/percent/pkg/percent"
  ```

### 1.3. Usage

- Examples
  > Examples of how to use the Percent package.

  ```go
  package main

  import (
      "fmt"
      "log"
      
      "github.com/sentenz/percent/pkg/percent"
  )

  func main() {
      // Example 1: Calculate what percentage of a value is
      // What is 25% of 200?
      result, err := percent.Percent(25, 200.0)
      if err != nil {
          log.Fatalf("Error calculating percent: %v", err)
      }
      fmt.Printf("25%% of 200 = %.2f\n", result) // Output: 25% of 200 = 50.00

      // Example 2: Calculate what percentage a value is of a total
      // What percentage is 50 of 200?
      pct, err := percent.Of(50.0, 200.0)
      if err != nil {
          log.Fatalf("Error calculating percentage: %v", err)
      }
      fmt.Printf("50 is %.2f%% of 200\n", pct) // Output: 50 is 25.00% of 200
  }
  ```

## 2. Contribute

Contribution guidelines and project management tools.

### 2.1. Task Runner

#### 2.1.1. Make

[Make](https://www.gnu.org/software/make/) is a automation tool that defines and manages tasks to streamline development workflows.

1. Insights and Details

    - [Makefile](Makefile)
      > Makefile defining tasks for building, testing, and managing the project.

2. Usage and Instructions

    - Tasks

      ```bash
      make help
      ```

      > [!NOTE]
      > - Each task description must begin with `##` to be included in the task list.

      ```plaintext
      $ make help

      Tasks
              A collection of tasks used in the current project.

      Usage
              make <task>

              bootstrap         Initialize a software development workspace with requisites
              setup             Install and configure all dependencies essential for development
              teardown          Remove development artifacts and restore the host to its pre-setup state
      ```

### 2.2. Bootstrap

#### 2.2.1. Scripts

[scripts/](scripts/README.md) provides scripts to bootstrap, setup, and teardown a software development workspace with requisites.

1. Insights and Details

    - [bootstrap.sh](scripts/bootstrap.sh)
      > Initializes a software development workspace with requisites.

    - [setup.sh](scripts/setup.sh)
      > Installs and configures all dependencies essential for development.

    - [teardown.sh](scripts/teardown.sh)
      > Removes development artifacts and restores the host to its pre-setup state.

2. Usage and Instructions

    - Tasks

      ```bash
      make bootstrap
      ```

      ```bash
      make setup
      ```

      ```bash
      make teardown
      ```

### 2.3. Dev Containers

[.devcontainer/](.devcontainer/README.md) provides Dev Containers as a consistent development environment using Docker containers.

1. Insights and Details

    - [go/](.devcontainer/go/)
      > Dev Container configuration for Go development environment.

      ```json
      // ...
      "postCreateCommand": "sudo make bootstrap && sudo make setup",
      // ...
      ```

      > [!NOTE]
      > The `devcontainer.json` runs the `bootstrap` and `setup` tasks to initialize and configure the development environment.

2. Usage and Instructions

    - Tasks

      ```bash
      # TODO
      # make devcontainer-go
      ```

### 2.4. Dependency Manager

#### 2.4.1. Go Modules

[Go Modules](https://go.dev/ref/mod) is the dependency management system for Go that simplifies the process of managing dependencies and libraries.

1. Insights and Details

    - [go.mod](go.mod)
      > Go module file defining the module path and dependencies.

    - [go.sum](go.sum)
      > Go checksum file containing expected cryptographic checksums of module dependencies.

    - [vendor/](vendor/)
      > Directory containing vendored dependencies.

2. Usage and Instructions

    - Tasks

      ```bash
      make go-mod-tidy
      ```

      ```bash
      make go-mod-vendor
      ```

### 2.5. Software Testing

#### 2.5.1. Unit Testing

[Go testing](https://pkg.go.dev/testing) is the standard library package for unit testing in Go.

1. Insights and Details

    - `testing.T`
      > Unit tests use the standard Go testing package with `testing.T`.

    - [AGENTS.md](./AGENTS.md)
      > Automate unit test generation using Large Language Models (LLMs) Agents.

    - [SKILLS.md](./SKILLS.md)
      > Project capabilities, commands, and tools documentation.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/go-tests@latest
      ```

    - Tasks

      ```bash
      make go-test-unit
      ```

#### 2.5.2. Fuzz Testing

[Go fuzzing](https://go.dev/security/fuzz/) is a testing technique that uses randomized inputs to find bugs and security vulnerabilities.

1. Insights and Details

    - `testing.F`
      > Fuzz tests use the standard Go testing package with `testing.F`.

    - [AGENTS.md](./AGENTS.md)
      > Automate fuzz test generation using Large Language Models (LLMs) Agents.

1. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/go-tests@latest
      ```

    - Tasks

      ```bash
      make go-test-fuzz
      ```

#### 2.5.3. Benchmarks

[Go benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks) measure the performance of code and track performance regressions.

1. Insights and Details

    - `testing.B`
      > Benchmark tests use the standard Go testing package with `testing.B`.

    - [AGENTS.md](./AGENTS.md)
      > Automate benchmark test generation using Large Language Models (LLMs) Agents.

1. Usage and Instructions

    - Tasks

      ```bash
      make go-test-bench
      ```

#### 2.5.4. Code Coverage

[go tool cover](https://pkg.go.dev/cmd/cover) provides code coverage analysis for Go tests.

1. Insights and Details

    - Code coverage reports are generated in HTML and XML formats.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/go-tests@latest
      ```

    - Tasks

      ```bash
      make go-test-coverage
      ```

### 2.6. Release Manager

#### 2.6.1. Semantic-Release

[Semantic-Release](https://github.com/semantic-release/semantic-release) automates the release process by analyzing commit messages to determine the next version number, generating changelog and release notes, and publishing the release.

1. Insights and Details

    - [.releaserc.json](.releaserc.json)
      > Configuration file for Semantic-Release specifying release rules and plugins.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/semantic-release@latest
      ```

### 2.7. Update Manager

#### 2.7.1. Renovate

[Renovate](https://github.com/renovatebot/renovate) automates dependency updates by creating merge requests for outdated dependencies, libraries and packages.

1. Insights and Details

    - [renovate.json](renovate.json)
      > Configuration file for Renovate specifying update rules and schedules.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/renovate@latest
      ```

#### 2.7.2. Dependabot

[Dependabot](https://github.com/dependabot/dependabot-core) automates dependency updates by creating pull requests for outdated dependencies, libraries and packages.

1. Insights and Details

    - [.github/dependabot.yml](.github/dependabot.yml)
      > Configuration file for Dependabot specifying update rules and schedules.

### 2.8. Policy Manager

#### 2.8.1. Conftest

[Conftest](https://www.conftest.dev/) is a **Policy as Code (PaC)** tool to streamline policy management for improved development, security and audit capability.

1. Insights and Details

    - [conftest.toml](conftest.toml)
      > Configuration file for Conftest specifying policy paths and output formats.

    - [tests/policy](tests/policy/)
      > Directory contains Rego policies for Conftest to enforce best practices and compliance standards.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/regal@latest
      ```

      ```yaml
      uses: sentenz/actions/conftest@latest
      ```

    - Tasks

      ```bash
      make policy-lint-regal <filepath>
      ```

      ```bash
      make policy-analysis-conftest <filepath>
      ```

### 2.9. Supply Chain Manager

#### 2.9.1. Trivy

[Trivy](https://github.com/aquasecurity/trivy) is a comprehensive security scanner for vulnerabilities, misconfigurations, and compliance issues in container images, filesystems, and source code.

1. Insights and Details

    - [trivy.yaml](trivy.yaml)
      > Configuration file for Trivy specifying scan settings and options.

    - [.trivyignore](.trivyignore)
      > File specifying vulnerabilities to ignore during Trivy scans.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/trivy@latest
      ```

    - Tasks

      ```bash
      make sast-trivy-fs <path>
      ```

      ```bash
      make sast-trivy-sbom-cyclonedx-fs <path>
      ```

      ```bash
      make sast-trivy-sbom <sbom_path>
      ```

      ```bash
      make sast-trivy-sbom-license <sbom_path>
      ```

## 3. References

- GitHub [Template DX](https://github.com/sentenz/template-dx) repository.
- Sentenz [Actions](https://github.com/sentenz/actions) repository.
- Sentenz [Manager Tools](https://github.com/sentenz/convention/issues/392) article.
