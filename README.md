# Percent

[![Go Version](https://img.shields.io/github/go-mod/go-version/sentenz/percent?logo=go)](https://go.dev/doc/install)
[![GoDoc](https://pkg.go.dev/badge/github.com/sentenz/percent/v3)](https://pkg.go.dev/github.com/sentenz/percent/v3/pkg/percent)
[![Go Report Card](https://goreportcard.com/badge/github.com/sentenz/percent/v3)](https://goreportcard.com/report/github.com/sentenz/percent/v3)
[![Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)](https://github.com/sentenz/percent/actions/workflows/go-tests.yml)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

Percent is a Go package that provides utility functions for calculating percentages and performing related operations.

- [1. Module](#1-module)
  - [1.1. Prerequisites](#11-prerequisites)
  - [1.2. Installation](#12-installation)
  - [1.3. Usage](#13-usage)
  - [1.4. Versioning](#14-versioning)
- [2. Contribute](#2-contribute)
- [3. References](#3-references)

## 1. Module

### 1.1. Prerequisites

- [Go](https://golang.org/)
  > Go programming language environment for building and running Go applications.

  ```bash
  # For linux/amd64
  wget https://go.dev/dl/go1.22.0.linux-amd64.tar.gz
  tar -C /usr/local -xzf go1.22.0.linux-amd64.tar.gz
  export PATH=$PATH:/usr/local/go/bin
  ```

### 1.2. Installation

- Install
  > Install the package via `go get` command.

  ```bash
  go get github.com/sentenz/percent/v3@latest
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

### 1.4. Versioning

This module follows [Go's semantic import versioning](https://go.dev/ref/mod#major-version-suffixes) for major versions v2 and above.

- Module Path
  > For major versions v2+, the module path includes a version suffix.

  | Version        | Module Path                     |
  | -------------- | ------------------------------- |
  | v0.x.x, v1.x.x | `github.com/sentenz/percent`    |
  | v2.x.x         | `github.com/sentenz/percent/v2` |
  | v3.x.x         | `github.com/sentenz/percent/v3` |

- Go Proxy
  > Verify module availability on the [Go Module Proxy](https://proxy.golang.org/).

  ```bash
  # For v3.x.x versions
  curl https://proxy.golang.org/github.com/sentenz/percent/v3/@v/v3.0.3.info
  ```

- Documentation
  > Access documentation on [pkg.go.dev](https://pkg.go.dev/) using the versioned path.

  | Version | Documentation URL                                |
  | ------- | ------------------------------------------------ |
  | v1.x.x  | https://pkg.go.dev/github.com/sentenz/percent    |
  | v3.x.x  | https://pkg.go.dev/github.com/sentenz/percent/v3 |

- Migration
  > When upgrading to a new major version, update imports in all Go files.

  ```go
  // Before (v1.x.x)
  import "github.com/sentenz/percent/pkg/percent"

  // After (v3.x.x)
  import "github.com/sentenz/percent/v3/pkg/percent"
  ```

## 2. Contribute

[CONTRIBUTING.md](CONTRIBUTING.md) provides guidens and instructions for contributing to the project.

- [AI Agents](CONTRIBUTING.md#1-ai-agents)
  > Automated tools that assist in various development tasks such as code generation, testing, and documentation.

- [Skills Manager](CONTRIBUTING.md#2-skills-manager)
  > CLI tool for managing AI agent skills in development projects.

- [Task Runner](CONTRIBUTING.md#3-task-runner)
  > Make automation tool that defines and manages tasks to streamline development workflows.

- [Bootstrap](CONTRIBUTING.md#4-bootstrap)
  > Scripts to bootstrap, setup, and teardown a software development workspace with requisites.

- [Dev Containers](CONTRIBUTING.md#5-dev-containers)
  > Consistent development environments using Docker containers.

- [Dependency Manager](CONTRIBUTING.md#6-dependency-manager)
  > Go Modules for managing dependencies and vendored modules.

- [Software Testing](CONTRIBUTING.md#7-software-testing)
  > Unit tests, fuzz tests, benchmarks, and coverage for ensuring code quality and reliability.

- [Release Manager](CONTRIBUTING.md#8-release-manager)
  > Semantic-Release automates the release process by analyzing commit messages.

- [Update Manager](CONTRIBUTING.md#9-update-manager)
  > Renovate and Dependabot automate dependency updates by creating pull requests.

- [Policy Manager](CONTRIBUTING.md#12-policy-manager)
  > Conftest for policy-as-code enforcement.

- [Supply Chain Manager](CONTRIBUTING.md#13-supply-chain-manager)
  > Trivy for security scanning of vulnerabilities, misconfigurations, and compliance issues.

## 3. References

- Sentenz [Template DX](https://github.com/sentenz/template-dx) repository.
- Sentenz [Actions](https://github.com/sentenz/actions) repository.
- Sentenz [Manager Tools](https://github.com/sentenz/convention/issues/392) article.
- Sentenz [Skills](https://github.com/sentenz/skills) repository.
