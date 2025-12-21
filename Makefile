# SPDX-License-Identifier: Apache-2.0

ifneq (,$(wildcard .env))
	include .env
	export
endif

# Define Variables

PIP_VENV := .venv/bin

SHELL := bash
.SHELLFLAGS := -euo pipefail -c
.ONESHELL:

# Define Targets

default: help

# NOTE Targets MUST have a leading comment line starting with `##` to be included in the list. See the targets below for examples.
help:
	@awk 'BEGIN {printf "Tasks\n\tA collection of tasks used in the current project.\n\n"}'
	@awk 'BEGIN {printf "Usage\n\tmake $(shell tput -Txterm setaf 6)<task>$(shell tput -Txterm sgr0)\n\n"}' $(MAKEFILE_LIST)
	@awk '/^##/{c=substr($$0,3);next}c&&/^[[:alpha:]][[:alnum:]_-]+:/{print "$(shell tput -Txterm setaf 6)\t" substr($$1,1,index($$1,":")) "$(shell tput -Txterm sgr0)",c}1{c=0}' $(MAKEFILE_LIST) | column -s: -t
.PHONY: help

# ── Setup & Teardown ─────────────────────────────────────────────────────────────────────────────

## Initialize a software development workspace with requisites
bootstrap:
	@cd ./scripts/ && bash ./bootstrap.sh
.PHONY: bootstrap

## Install and configure all dependencies essential for development
setup:
	@cd ./scripts/ && bash ./setup.sh
.PHONY: setup

## Remove development artifacts and restore the host to its pre-setup state
teardown:
	@cd ./scripts/ && bash ./teardown.sh
.PHONY: teardown

# ── Go Tools ─────────────────────────────────────────────────────────────────────────────────────

## Tidy Go modules
go-mod-tidy:
	go mod tidy
.PHONY: go-mod-tidy

## Vendor Go modules
go-mod-vendor:
	go mod vendor
.PHONY: go-mod-vendor

## Run Go unit tests with race detection and JUnit XML report
go-test-unit:
	@mkdir -p logs/unit
	go test -race -json ./... 2>&1 | tee logs/unit/test-output.json | go run -mod=vendor github.com/jstemmer/go-junit-report/v2 -set-exit-code -iocopy -out logs/unit/junit-report.xml
.PHONY: go-test-unit

## Run Go tests with coverage report
go-test-coverage:
	@mkdir -p logs/coverage

	go test -coverprofile=logs/coverage/coverage.out ./...
	go tool cover -html=logs/coverage/coverage.out -o logs/coverage/coverage.html
	go run -mod=vendor github.com/boumenot/gocover-cobertura < logs/coverage/coverage.out > logs/coverage/coverage.xml
.PHONY: go-test-coverage

## Run Go benchmarks
go-test-bench:
	go test -bench=. -benchmem ./...
.PHONY: go-test-bench

## Run fuzz tests
go-test-fuzz:
	go test -fuzz=Fuzz -fuzztime=10s ./pkg/percent
.PHONY: go-test-fuzz

## Format Go code according to Go standards
go-fmt:
	go fmt ./...
.PHONY: go-fmt

## Update Go source files to use new APIs
go-fix:
	go fix ./...
.PHONY: go-fix

## Check Go code for common mistakes
go-vet:
	go vet ./...
.PHONY: go-vet

## Check Go code for known vulnerabilities
go-vuln:
	go run -mod=vendor golang.org/x/vuln/cmd/govulncheck ./...
.PHONY: go-vuln

## Run all Go code quality checks
go-check:
	$(MAKE) go-fmt
	$(MAKE) go-vet
	$(MAKE) go-vuln
.PHONY: go-check

# ── Secrets Manager ──────────────────────────────────────────────────────────────────────────────

SECRETS_SOPS_UID ?= sops-percent

# Usage: make secrets-gpg-generate SECRETS_SOPS_UID=<uid>
#
## Generate a new GPG key pair for SOPS
secrets-gpg-generate:
	@gpg --batch --quiet --passphrase '' --quick-generate-key "$(SECRETS_SOPS_UID)" ed25519 cert,sign 0
	@NEW_FPR="$$(gpg --list-keys --with-colons "$(SECRETS_SOPS_UID)" | awk -F: '/^fpr:/ {print $$10; exit}')"
	@gpg --batch --quiet --passphrase '' --quick-add-key "$${NEW_FPR}" cv25519 encrypt 0
.PHONY: secrets-gpg-generate

# Usage: make secrets-gpg-show SECRETS_SOPS_UID=<uid>
#
## Print the GPG key fingerprint for SOPS (.sops.yaml)
secrets-gpg-show:
	@FPR="$$(gpg --list-keys --with-colons "$(SECRETS_SOPS_UID)" | awk -F: '/^fpr:/ {print $$10; exit}')"; \
	if [ -z "$${FPR}" ]; then \
		echo "error: no fingerprint found for UID '$(SECRETS_SOPS_UID)'" >&2; \
		exit 1; \
	fi; \
	echo -e "UID: $(SECRETS_SOPS_UID)\nFingerprint: $${FPR}"
.PHONY: secrets-gpg-show

# Usage: make secrets-gpg-remove SECRETS_SOPS_UID=<uid>
#
## Remove an existing GPG key for SOPS (interactive)
secrets-gpg-remove:
	if ! gpg --list-keys "$(SECRETS_SOPS_UID)" >/dev/null 2>&1; then
		echo "warning: no key found for '$(SECRETS_SOPS_UID)'" >&2
		exit 0
	fi
	echo "info: deleting key for '$(SECRETS_SOPS_UID)'"
	# Delete private key first, then public key
	gpg --yes --delete-secret-keys "$(SECRETS_SOPS_UID)"
	gpg --yes --delete-keys "$(SECRETS_SOPS_UID)"
.PHONY: secrets-gpg-remove

# Usage: make secrets-sops-encrypt <files>
#
## Encrypt file using SOPS
secrets-sops-encrypt:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "usage: make secrets-sops-encrypt <files>"; \
		exit 1; \
	fi

	export PATH="$$PATH:$(shell go env GOPATH 2>/dev/null)/bin"
	@for file in $(filter-out $@,$(MAKECMDGOALS)); do \
		if [ -f "$$file" ]; then \
			sops --encrypt --in-place "$$file"; \
		else \
			echo "Skipping (not found): $$file" >&2; \
		fi; \
	done
.PHONY: secrets-sops-encrypt

# Usage: make secrets-sops-decrypt <files>
#
## Decrypt file using SOPS
secrets-sops-decrypt:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "usage: make secrets-sops-encrypt <files>"; \
		exit 1; \
	fi

	export PATH="$$PATH:$(shell go env GOPATH 2>/dev/null)/bin"
	@for file in $(filter-out $@,$(MAKECMDGOALS)); do \
		if [ -f "$$file" ]; then \
			sops --decrypt --in-place "$$file"; \
		else \
			echo "Skipping (not found): $$file" >&2; \
		fi; \
	done
.PHONY: secrets-sops-decrypt

# Usage: make secrets-sops-view <file>
#
## View a file encrypted with SOPS
secrets-sops-view:
	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "usage: make secrets-sops-view <file>"; \
		exit 1; \
	fi

	export PATH="$$PATH:$(shell go env GOPATH 2>/dev/null)/bin"
	sops --decrypt "$(filter-out $@,$(MAKECMDGOALS))"
.PHONY: secrets-sops-view

# ── Policy Manager ───────────────────────────────────────────────────────────────────────────────

POLICY_IMAGE_CONFTEST ?= openpolicyagent/conftest:v0.65.0@sha256:afa510df6d4562ebe24fb3e457da6f6d6924124140a13b51b950cc6cb1d25525

# Usage: make policy-analysis-conftest <filepath>
#
## Analyze configuration files using Conftest for policy violations and generate a report
policy-analysis-conftest:
	@mkdir -p logs/policy

	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "usage: make policy-analysis-conftest <filepath>"; \
		exit 1; \
	fi

	docker run --rm -v "${PWD}:/workspace" -w /workspace "$(POLICY_IMAGE_CONFTEST)" test "$(filter-out $@,$(MAKECMDGOALS))" > logs/policy/conftest.json 2>&1
.PHONY: policy-analysis-conftest

POLICY_IMAGE_REGAL ?= ghcr.io/openpolicyagent/regal:0.37.0@sha256:a09884658f3c8c9cc30de136b664b3afdb7927712927184ba891a155a9676050

# Usage: make policy-lint-regal <filepath>
#
## Lint Rego policies using Regal and generate a report
policy-lint-regal:
	@mkdir -p logs/analysis

	@if [ -z "$(filter-out $@,$(MAKECMDGOALS))" ]; then \
		echo "usage: make policy-lint-regal"; \
		exit 1; \
	fi

	docker pull "$(POLICY_IMAGE_REGAL)"
	docker run --rm -v "${PWD}:/workspace" -w /workspace "$(POLICY_IMAGE_REGAL)" regal lint "$(filter-out $@,$(MAKECMDGOALS))" --format json > logs/analysis/regal.json 2>&1
.PHONY: policy-lint-regal
