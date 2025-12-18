ifneq (,$(wildcard .env))
	include .env
	export
endif

# See https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Display help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-30s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
.PHONY: help

setup: ## Setup dependencies and tools
	cd $(@D)/scripts && chmod +x setup.sh && ./setup.sh
.PHONY: setup

app-build: ## Perform application build
	mkdir -p $(@D)/cmd/bin
	go build -buildvcs=false -o $(@D)/cmd/bin -race -v $(@D)/cmd/app
.PHONY: app-build

app-run: ## Perform application run
	go run $(@D)/cmd/app/main.go
.PHONY: app-run

app-audit: ## Perform application audit
	mkdir -p $(@D)/logs/audit
	go mod tidy
	go mod vendor
	go list -json -m | tee $(@D)/logs/audit/audit.log
.PHONY: app-audit

test-unit: ## Perform unit test
	mkdir -p $(@D)/logs/test
	go test -race ./... | tee $(@D)/logs/test/test.log
.PHONY: app-test

test-fuzz: ## ## Perform fuzz test
	go test --fuzz=Fuzz -fuzztime=10s
.PHONY: test-fuzz

test-bench: ## Perform benchmark test
	mkdir -p $(@D)/logs/test
	go test -timeout 30s -race -count 3 -bench=. -benchmem ./... | tee $(@D)/logs/test/benchmark.log
.PHONY: test-bench

test-cover: ## Perform code coverage
	mkdir -p $(@D)/logs/test
	go test -race -coverprofile=logs/test/coverage.log -covermode=atomic ./...
	go tool cover -func=$(@D)/logs/test/coverage.log
	go tool cover -html=$(@D)/logs/test/coverage.log
.PHONY: test-cover

analysis-golangci: ## Perform static code analysis
	@mkdir -p $(@D)/logs/analysis
	docker run --rm -v "${PWD}:/workspace" -w /workspace golangci/golangci-lint:v2.7.2-alpine@sha256:1e1851102b736971267400e08b3e4b2e7799c73976a998820f6f6b6b86b48343 golangci-lint run ./...
.PHONY: analysis-golangci
