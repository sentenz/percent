# Contributing

Contribution guidelines and project management tools.

- [1. AI Agents](#1-ai-agents)
- [2. Skills Manager](#2-skills-manager)
  - [2.1. Skills CLI](#21-skills-cli)
- [3. Task Runner](#3-task-runner)
  - [3.1. Make](#31-make)
- [4. Bootstrap](#4-bootstrap)
  - [4.1. Scripts](#41-scripts)
- [5. Git Hooks Manager](#5-git-hooks-manager)
  - [5.1. Lefthook](#51-lefthook)
- [6. Dev Containers](#6-dev-containers)
- [7. Dependency Manager](#7-dependency-manager)
  - [7.1. Go Modules](#71-go-modules)
- [8. Software Testing](#8-software-testing)
  - [8.1. Unit Testing](#81-unit-testing)
  - [8.2. Fuzz Testing](#82-fuzz-testing)
  - [8.3. Benchmarks](#83-benchmarks)
  - [8.4. Code Coverage](#84-code-coverage)
- [9. Release Manager](#9-release-manager)
  - [9.1. Semantic-Release](#91-semantic-release)
- [10. Update Manager](#10-update-manager)
  - [10.1. Renovate](#101-renovate)
  - [10.2. Dependabot](#102-dependabot)
- [11. Secrets Manager](#11-secrets-manager)
  - [11.1. SOPS](#111-sops)
- [12. Container Manager](#12-container-manager)
  - [12.1. Docker](#121-docker)
- [13. Policy Manager](#13-policy-manager)
  - [13.1. Conftest](#131-conftest)
- [14. SAST Manager](#14-sast-manager)
  - [14.1. Gitleaks](#141-gitleaks)
  - [14.2. TruffleHog](#142-trufflehog)
  - [14.3. Semgrep](#143-semgrep)
- [15. Supply Chain Manager](#15-supply-chain-manager)
  - [15.1. Trivy](#151-trivy)
- [16. Documentation Generators](#16-documentation-generators)
  - [16.1. Doxygen](#161-doxygen)
  - [16.2. MkDocs](#162-mkdocs)

## 1. AI Agents

AI Agents are automated tools that assist in various development tasks such as code generation, testing, and documentation.

1. Insights and Details

    - [AGENTS.md](AGENTS.md)
      > Instructions for AI coding agents working with the project.

    - [.agents/skills/](.agents/skills/)
      > Directory containing AI agent skill definitions and configurations.

2. Usage and Instructions

    - Implicit Invocation
      > AI Agents can be implicitly invoked based on file paths, programming languages, or specific keywords in user prompts.

      ```plaintext
      .agents/skills/<skill-name>/SKILL.md
      ```

    - Explicit Invocation
      > AI Agents can be explicitly invoked by specifying the skill name in user prompts.

      ```plaintext
      @agent <skill-name> <task-description>
      ```

## 2. Skills Manager

### 2.1. Skills CLI

[Skills CLI](https://skills.sh/) is a command-line tool for managing AI agent skills in development projects.

1. Insights and Details

    - [Sentenz Skills](https://github.com/sentenz/skills)
      > Reusable AI agent skills for various development tasks.

    - [skills-lock.json](skills-lock.json)
      > Lock file for managing skill dependencies and versions.

2. Usage and Instructions

    - Tasks

      ```bash
      make skills-add
      ```

      ```bash
      make skills-update
      ```

## 3. Task Runner

### 3.1. Make

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

## 4. Bootstrap

### 4.1. Scripts

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

## 5. Git Hooks Manager

### 5.1. Lefthook

[Lefthook](https://lefthook.dev/) is a fast, language-agnostic Git hooks manager that uses a single `lefthook.yml` configuration file to define hooks for automating tasks during the Git workflow.

1. Insights and Details

    - [lefthook.yml](lefthook.yml)
      > Configuration file for Lefthook specifying Git hooks and associated commands.

2. Usage and Instructions

    - Tasks

      ```bash
      make githooks-lefthook-initialize
      ```

      ```bash
      make githooks-lefthook-deinitialize
      ```

## 6. Dev Containers

1. Insights and Details

    - [.devcontainer/](.devcontainer/README.md)
      > Provides Dev Containers as a consistent development environment using Docker containers.

2. Usage and Instructions

    - Tasks

      ```bash
      make devcontainer-go
      make devcontainer-cpp
      make devcontainer-python
      ```

## 7. Dependency Manager

### 7.1. Go Modules

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

## 8. Software Testing

### 8.1. Unit Testing

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

### 8.2. Fuzz Testing

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

### 8.3. Benchmarks

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

### 8.4. Code Coverage

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

## 9. Release Manager

### 9.1. Semantic-Release

[Semantic-Release](https://github.com/semantic-release/semantic-release) automates the release process by analyzing commit messages to determine the next version number, generating changelog and release notes, and publishing the release.

1. Insights and Details

    - [.releaserc.json](.releaserc.json)
      > Configuration file for Semantic-Release specifying release rules and plugins.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/semantic-release@latest
      ```

## 10. Update Manager

### 10.1. Renovate

[Renovate](https://github.com/renovatebot/renovate) automates dependency updates by creating merge requests for outdated dependencies, libraries and packages.

1. Insights and Details

    - [renovate.json](renovate.json)
      > Configuration file for Renovate specifying update rules and schedules.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/renovate@latest
      ```

### 10.2. Dependabot

[Dependabot](https://github.com/dependabot/dependabot-core) automates dependency updates by creating pull requests for outdated dependencies, libraries and packages.

1. Insights and Details

    - [.github/dependabot.yml](.github/dependabot.yml)
      > Configuration file for Dependabot specifying update rules and schedules.

## 11. Secrets Manager

### 11.1. SOPS

[SOPS (Secrets OPerationS)](https://github.com/getsops/sops) is a tool for managing and encrypting sensitive data such as passwords, API keys, and other secrets.

1. Insights and Details

    - [.sops.yaml](.sops.yaml)
      > Configuration file for SOPS specifying encryption rules and key management.

2. Usage and Instructions

    - GPG Key Pair Generation

      - Tasks
        > Generate a new key pair to be used with SOPS.

        > [!NOTE]
        > Customize the UID by providing the `SECRETS_SOPS_UID` variable. Default UID is `sops-<repo>`.

        ```bash
        make secrets-gpg-generate SECRETS_SOPS_UID=<uid>
        ```

    - GPG Public Key Fingerprint

      - Tasks
        > Print the  GPG Public Key fingerprint associated with a given UID.

        ```bash
        make secrets-gpg-show SECRETS_SOPS_UID=<uid>
        ```

      - [.sops.yaml](.sops.yaml)
        > The GPG UID is required for populating in `.sops.yaml`.

        ```yaml
        creation_rules:
          - pgp: "<fingerprint>" # <uid>
        ```

    - SOPS Encrypt/Decrypt

      - Tasks
        > Encrypt/decrypt one or more files in place using SOPS.

        ```bash
        make secrets-sops-encrypt <files>
        ```

        ```bash
        make secrets-sops-decrypt <files>
        ```

## 12. Container Manager

### 12.1. Docker

[Docker](https://github.com/docker) containerization tool to run applications in isolated container environments and execute container-based tasks.

1. Insights and Details

    - [Dockerfile](Dockerfile)
      > Dockerfile defining the container image for the project.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/container@latest
      ```

    - Tasks

      ```bash
      make container-docker-build
      ```

      ```bash
      make container-docker-run
      ```

## 13. Policy Manager

### 13.1. Conftest

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

## 14. SAST Manager

SAST (Static Application Security Testing) tools for identifying security vulnerabilities and issues in source code, container images, and artifacts.

### 14.1. Gitleaks

[Gitleaks](https://github.com/gitleaks/gitleaks) is a SAST tool for detecting hardcoded secrets such as passwords, API keys, and tokens in git repositories and staged changes.

1. Insights and Details

    - [lefthook.yml](lefthook.yml)
      > Pre-commit hook runs `sast-gitleaks-protect` to scan staged changes before every commit.

2. Usage and Instructions

    - CI/CD

      ```yaml
      - component: sentenz/actions/gitleaks@latest
      ```

    - Tasks

      ```bash
      make sast-gitleaks-detect
      ```

      ```bash
      make sast-gitleaks-protect
      ```

### 14.2. TruffleHog

[TruffleHog](https://github.com/trufflesecurity/trufflehog) is a secret-scanning tool for detecting verified, unverified, and unknown credentials in filesystems and git repositories.

1. Insights and Details

    - [.github/workflows/trufflehog.yml](.github/workflows/trufflehog.yml)
      > Workflow definition for TruffleHog-based secret scanning in CI.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: trufflesecurity/trufflehog@latest
      ```

    - Tasks

      ```bash
      make sast-trufflehog-fs
      ```

      ```bash
      make sast-trufflehog-git
      ```

### 14.3. Semgrep

[Semgrep](https://github.com/semgrep/semgrep) is a static analysis tool for detecting code security issues and enforcing secure coding patterns across source files.

1. Insights and Details

    - [.github/workflows/semgrep.yml](.github/workflows/semgrep.yml)
      > Workflow definition for Semgrep-based static analysis in CI.

    - [lefthook.yml](lefthook.yml)
      > Pre-commit hook runs `sast-semgrep-scan` against staged files before every commit when Docker is available.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/semgrep@latest
      ```

    - Tasks

      ```bash
      make sast-semgrep-scan
      ```

## 15. Supply Chain Manager

Software Supply Chain Security for identifying vulnerabilities in dependencies by scanning SBOMs, container images, and filesystems.

### 15.1. Trivy

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

## 16. Documentation Generators

### 16.1. Doxygen

[Doxygen](https://www.doxygen.nl/) is an **API Documentation Generator** for C++, C programming languages, used to create software reference documentation from annotated source code.

1. Insights and Details

    - [Doxyfile](Doxyfile)
      > Configuration file for Doxygen specifying documentation generation settings.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/doxygen@latest
      ```

    - Tasks

      ```bash
      make pages-doxygen-build
      ```

      ```bash
      make pages-doxygen-serve
      ```

### 16.2. MkDocs

[MkDocs](https://www.mkdocs.org/) is a Static Site Generator (SSG) designed for building project documentation using Markdown files.

1. Insights and Details

    - [mkdocs.yml](mkdocs.yml)
      > Configuration file for MkDocs specifying site settings, theme, plugins, and markdown extensions.

2. Usage and Instructions

    - CI/CD

      ```yaml
      uses: sentenz/actions/mkdocs@latest
      ```

    - Tasks

      ```bash
      make pages-mkdocs-setup
      ```

      ```bash
      make pages-mkdocs-build
      ```

      ```bash
      make pages-mkdocs-serve
      ```
