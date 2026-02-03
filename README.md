# Percent

[![Go Version](https://img.shields.io/github/go-mod/go-version/sentenz/percent?logo=go)](https://go.dev/doc/install)
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
  - [2.1. AI Agents](#21-ai-agents)
  - [2.2. Task Runner](#22-task-runner)
    - [2.2.1. Make](#221-make)
  - [2.3. Bootstrap](#23-bootstrap)
    - [2.3.1. Scripts](#231-scripts)
  - [2.4. Dev Containers](#24-dev-containers)
  - [2.5. Dependency Manager](#25-dependency-manager)
    - [2.5.1. Go Modules](#251-go-modules)
  - [2.6. Software Testing](#26-software-testing)
    - [2.6.1. Unit Testing](#261-unit-testing)
    - [2.6.2. Fuzz Testing](#262-fuzz-testing)
    - [2.6.3. Benchmarks](#263-benchmarks)
    - [2.6.4. Code Coverage](#264-code-coverage)
  - [2.7. Release Manager](#27-release-manager)
    - [2.7.1. Semantic-Release](#271-semantic-release)
  - [2.8. Update Manager](#28-update-manager)
    - [2.8.1. Renovate](#281-renovate)
    - [2.8.2. Dependabot](#282-dependabot)
  - [2.9. Policy Manager](#29-policy-manager)
    - [2.9.1. Conftest](#291-conftest)
  - [2.10. Supply Chain Manager](#210-supply-chain-manager)
    - [2.10.1. Trivy](#2101-trivy)
- [3. References](#3-references)

## 1. Module

### 1.1. Prerequisites

- [Go](https://golang.org/)
  > Go programming language environment for building and running Go applications.

### 1.2. Installation

- Install
  > Install the package via `go get`.

  ```bash
  go get github.com/sentenz/percent/v3
  ```

- Import
  > Import the package in a called Go source file.

  ```go
  import "github.com/sentenz/percent/v3/pkg/percent"
  ```

### 1.3. Usage

- Examples
  > Examples of how to use the Percent package.

  ```go
  package main

  import (
      "fmt"
      "log"
      
      "github.com/sentenz/percent/v3/pkg/percent"
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

### 2.1. AI Agents

AI Agents are automated tools that assist in various development tasks such as code generation, testing, and documentation.

1. Insights and Details

    - [AGENTS.md](AGENTS.md)
      > Instructions for AI coding agents working with the project.

    - [SKILL.md](.github/skills/README.md)
      > Instructions for AI agent skills used in the project.

2. Usage and Instructions

    - Implicit Invocation
      > AI Agents can be implicitly invoked based on file paths, programming languages, or specific keywords in user prompts.

      ```plaintext
      .github/skills/<skill-name>/SKILL.md
      ```

    - Explicit Invocation
      > AI Agents can be explicitly invoked by specifying the skill name in user prompts.

      ```plaintext
      @agent <skill-name> <task-description>
      ```

### 2.2. Task Runner

#### 2.2.1. Make

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

### 2.3. Bootstrap

#### 2.3.1. Scripts

1. Insights and Details

    - [scripts/](scripts/README.md)
      > Provides scripts to bootstrap, setup, and teardown a software development workspace with requisites.

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

### 2.4. Dev Containers

1. Insights and Details

    - [.devcontainer/](.devcontainer/README.md)
      > Provides Dev Containers as a consistent development environment using Docker containers.

2. Usage and Instructions

    - Tasks

      ```bash
      # TODO
      # make devcontainer-go
      ```

### 2.5. Dependency Manager

#### 2.5.1. Go Modules

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

### 2.6. Software Testing

#### 2.6.1. Unit Testing

[Go testing](https://pkg.go.dev/testing) is the standard library package for unit testing in Go.

1. Insights and Details

    - `testing.T`
      > Unit tests use the standard Go testing package with `testing.T`.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/go-tests@latest
      ```

    - Tasks

      ```bash
      make go-test-unit
      ```

    - AI Agents
      > Instruct Agent Skills capabilities to to perform [Unit Testing](.github/skills/unit-testing/SKILL.md) tasks.

#### 2.6.2. Fuzz Testing

[Go fuzzing](https://go.dev/security/fuzz/) is a testing technique that uses randomized inputs to find bugs and security vulnerabilities.

1. Insights and Details

    - `testing.F`
      > Fuzz tests use the standard Go testing package with `testing.F`.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/go-tests@latest
      ```

    - Tasks

      ```bash
      make go-test-fuzz
      ```

    - AI Agents
      > Instruct Agent Skills capabilities to to perform [Fuzz Testing](.github/skills/fuzz-testing/SKILL.md) tasks.

#### 2.6.3. Benchmarks

[Go benchmarks](https://pkg.go.dev/testing#hdr-Benchmarks) measure the performance of code and track performance regressions.

1. Insights and Details

    - `testing.B`
      > Benchmark tests use the standard Go testing package with `testing.B`.

2. Usage and Instructions

    - Tasks

      ```bash
      make go-test-bench
      ```

    - AI Agents
      > Instruct Agent Skills capabilities to to perform [Benchmark Testing](.github/skills/benchmark-testing/SKILL.md) tasks.

#### 2.6.4. Code Coverage

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

### 2.7. Release Manager

#### 2.7.1. Semantic-Release

[Semantic-Release](https://github.com/semantic-release/semantic-release) automates the release process by analyzing commit messages to determine the next version number, generating changelog and release notes, and publishing the release.

1. Insights and Details

    - [.releaserc.json](.releaserc.json)
      > Configuration file for Semantic-Release specifying release rules and plugins.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/semantic-release@latest
      ```

### 2.8. Update Manager

#### 2.8.1. Renovate

[Renovate](https://github.com/renovatebot/renovate) automates dependency updates by creating merge requests for outdated dependencies, libraries and packages.

1. Insights and Details

    - [renovate.json](renovate.json)
      > Configuration file for Renovate specifying update rules and schedules.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/renovate@latest
      ```

#### 2.8.2. Dependabot

[Dependabot](https://github.com/dependabot/dependabot-core) automates dependency updates by creating pull requests for outdated dependencies, libraries and packages.

1. Insights and Details

    - [.github/dependabot.yml](.github/dependabot.yml)
      > Configuration file for Dependabot specifying update rules and schedules.

### 2.9. Policy Manager

#### 2.9.1. Conftest

[Conftest](https://www.conftest.dev/) is a **Policy as Code (PaC)** tool to streamline policy management for improved development, security and audit capability.

1. Insights and Details

    - [conftest.toml](conftest.toml)
      > Configuration file for Conftest specifying policy paths and output formats.

    - [tests/policy/](tests/policy/)
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
      make policy-regal-lint <filepath>
      ```

      ```bash
      make policy-conftest-test <filepath>
      ```

### 2.10. Supply Chain Manager

#### 2.10.1. Trivy

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
      make sast-trivy-sbom-scan <sbom_path>
      ```

      ```bash
      make sast-trivy-sbom-license <sbom_path>
      ```

## 3. References

- Sentenz [Template DX](https://github.com/sentenz/template-dx) repository.
- Sentenz [Actions](https://github.com/sentenz/actions) repository.
- Sentenz [Manager Tools](https://github.com/sentenz/convention/issues/392) article.
