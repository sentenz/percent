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
	@for fuzz in Percent Of Change Remain FromRatio ToRatio; do \
		echo "Fuzzing: Fuzz$${fuzz}"; \
		go test -fuzz=Fuzz"$${fuzz}" -fuzztime=10s ./pkg/percent || exit 1; \
	done
.PHONY: go-test-fuzz

## Generate fuzz tests for Go functions using AI agents-based with AGENTS.md
copilot-agents-test-fuzz:
	# TODO: Implement AI agents-based fuzz test generation
	@echo "AI agents-based fuzz test generation is not yet implemented."
.PHONY: copilot-agents-test-fuzz

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
