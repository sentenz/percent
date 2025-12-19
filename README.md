# Percent

[![GoDoc](https://godoc.org/github.com/sentenz/percent?status.svg)](https://godoc.org/github.com/sentenz/percent/pkg/percent)
[![Go Report Card](https://goreportcard.com/badge/github.com/sentenz/percent)](https://goreportcard.com/report/github.com/sentenz/percent)
[![License](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Percent is a Go package that provides utility functions for calculating percentages and performing related operations.

- [1. Details](#1-details)
  - [1.1. Prerequisites](#11-prerequisites)
  - [1.2. Installation](#12-installation)
- [2. Usage](#2-usage)
  - [2.1. Bootstrap](#21-bootstrap)
  - [2.2. Dev Containers](#22-dev-containers)
  - [2.3. Task Runner](#23-task-runner)
    - [2.3.1. Make](#231-make)
  - [2.4. Release Manager](#24-release-manager)
    - [2.4.1. Semantic-Release](#241-semantic-release)
  - [2.5. Update Manager](#25-update-manager)
    - [2.5.1. Renovate](#251-renovate)
  - [2.6. Secrets Manager](#26-secrets-manager)
    - [2.6.1. SOPS](#261-sops)
- [3. References](#3-references)

## 1. Details

### 1.1. Prerequisites

- [Go](https://golang.org/)
  > Go programming language environment for building and running Go applications.

- [Make](https://www.gnu.org/software/make/)
  > Task automation tool to manage build processes and workflows.

- [Docker](https://www.docker.com/)
  > Containerization tool to run applications in isolated container environments and execute container-based tasks.

### 1.2. Installation

Install the package via `go get`.

```bash
go get github.com/sentenz/percent
```

Import the package in a called Go source file.

```go
import "github.com/sentenz/percent/pkg/percent"
```

## 2. Usage

### 2.1. Bootstrap

A bootstrap script initializes an environment, application, or system by installing dependencies, configuring settings, or preparing the system for further operations.

- [scripts/](scripts/README.md)
  > Collection of bootstrap, setup, or teardown scripts across multiple environments.

### 2.2. Dev Containers

A Development Container (Dev Container) utilizes a container as a full-featured development environment.

- [.devcontainer/](.devcontainer/README.md)
  > Collection of Dev Container resources published under `mcr.microsoft.com/devcontainers`.

### 2.3. Task Runner

#### 2.3.1. Make

- [Makefile](Makefile)
  > A Makefile is used to define and manage tasks for building, testing, and maintaining the project.

  > [!NOTE]
  > - Run the `make help` command in the terminal to list the tasks used for the project.
  > - Targets **must** have a leading comment line starting with `##` to be included in the task list.

  ```plaintext
  $ make help

  Tasks
          A collection of tasks used in the current project.

  Usage
          make <task>

          bootstrap                   Initialize a software development workspace with requisites
          setup                       Install and configure all dependencies essential for development
          teardown                    Remove development artifacts and restore the host to its pre-setup state
          go-test-unit                Run Go unit tests with race detection
          go-test-coverage            Run Go tests with coverage report
          go-test-bench               Run Go benchmarks
          go-test-fuzz                Run fuzz tests
          go-check                    Run all Go code quality checks
          secrets-sops-encrypt        Encrypt file using SOPS
          secrets-sops-decrypt        Decrypt file using SOPS
          analysis-policy-conftest    Analyze configuration files using Conftest for policy violations and generate a report
          analysis-lint-regal         Lint Rego policies using Regal and generate a report
  ```

### 2.4. Release Manager

#### 2.4.1. Semantic-Release

The [Semantic-Release CI/CD Component](https://gitlab.samscm.net/explore/catalog/development-environment/ci-cd/manager) tool automates release workflows, including determining the next version number, generating a `CHANGELOG.md` file, and publishing the release notes with artifacts.

- [.gitlab-ci.yml](.gitlab-ci.yml)

  ```yaml
  include:
    - component: $CI_SERVER_FQDN/development-environment/ci-cd/manager/semantic-release@<version>
  ```

### 2.5. Update Manager

#### 2.5.1. Renovate

The [Renovate CI/CD Component](https://gitlab.samscm.net/explore/catalog/development-environment/ci-cd/manager) automates dependency updates by creating Pull Requests (PRs) or Merge Requests (MRs) to update versions.

- [.gitlab-ci.yml](.gitlab-ci.yml)

  ```yaml
  include:
    - component: $CI_SERVER_FQDN/development-environment/ci-cd/manager/renovate@<version>
  ```

### 2.6. Secrets Manager

#### 2.6.1. SOPS

1. GPG Key Pair Generation

    - Task Runner
      > Generate a new key pair to be used with SOPS.

      > [!NOTE]
      > The UID can be customized via the `SECRETS_SOPS_UID` variable (defaults to `sops-dx`).

      ```sh
      make secrets-gpg-generate SECRETS_SOPS_UID=<uid>
      ```

2. GPG Public Key Fingerprint

    - Task Runner
      > Print the  GPG Public Key fingerprint associated with a given UID.

      ```sh
      make secrets-gpg-show SECRETS_SOPS_UID=<uid>
      ```

    - [.sops.yaml](.sops.yaml)
      > The GPG UID is required for populating in `.sops.yaml`.

      ```yaml
      creation_rules:
        - pgp: "<fingerprint>" # <uid>
      ```

3. SOPS Encrypt/Decrypt

    - Task Runner
      > Encrypt/decrypt one or more files in place using SOPS.

      ```sh
      make secrets-sops-encrypt <files>
      make secrets-sops-decrypt <files>
      ```

## 3. References

- GitHub [Template DX](https://github.com/sentenz/template-dx) repository.
- Sentenz [Manager Tools](https://github.com/sentenz/convention/issues/392) article.
