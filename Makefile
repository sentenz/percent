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

setup-devops: ## Setup dependencies and tools for the devops service
	cd $(DEVOPS_PATH)/scripts && chmod +x setup.sh && ./setup.sh
.PHONY: setup-devops

setup-devcontainer: ## Setup dependencies and tools for the vscode devcontainer
	$(MAKE) update-submodule
	$(MAKE) setup
.PHONY: setup-devcontainer

setup-security: ## Setup dependencies and tools for the security service
	cd $(@D)/scripts/pipeline && chmod +x setup_security.sh && ./setup_security.sh
.PHONY: setup-security

setup-integration: ## Setup dependencies and tools for the integration service
	cd $(@D)/scripts/pipeline && chmod +x setup_integration.sh && ./setup_integration.sh
.PHONY: setup-integration

setup-build: ## Setup dependencies and tools for the build service
	cd $(@D)/scripts/pipeline && chmod +x setup_build.sh && ./setup_build.sh
.PHONY: setup-build

setup-testing: ## Setup dependencies and tools for the testing service
	cd $(@D)/scripts/pipeline && chmod +x setup_testing.sh && ./setup_testing.sh
.PHONY: setup-testing

setup-release: ## Setup dependencies and tools for the release service
	cd $(@D)/scripts/pipeline && chmod +x setup_release.sh && ./setup_release.sh
.PHONY: setup-release

setup-continuous-security: ## Setup dependencies and tools for the continuous security pipeline
	$(MAKE) update-submodule
	$(MAKE) setup-security
.PHONY: setup-continuous-security

setup-continuous-integration: ## Setup dependencies and tools for the continuous integration pipeline
	$(MAKE) update-submodule
	$(MAKE) setup-integration
.PHONY: setup-continuous-integration

setup-continuous-build: ## Setup dependencies and tools for the continuous build pipeline
	$(MAKE) setup-build
.PHONY: setup-continuous-build

setup-continuous-testing: ## Setup dependencies and tools for the continuous testing pipeline
	$(MAKE) setup-testing
.PHONY: setup-continuous-testing

setup-continuous-release: ## Setup dependencies and tools for the continuous release pipeline
	$(MAKE) setup-release
.PHONY: setup-continuous-release

teardown-devops: ## Teardown dependencies and tools for the devops service
	cd $(DEVOPS_PATH)/scripts && chmod +x teardown.sh && ./teardown.sh
.PHONY: teardown-devops

update-devops: ## Update dependencies and tools for the devops service
	$(MAKE) teardown-devops
	$(MAKE) update-submodule
	$(MAKE) setup-devops
.PHONY: update-devops

run-linter-staged: ## Perform analysis of local staged files
	cd $(DEVOPS_PATH)/cmd/app && chmod +x sast.sh && ./sast.sh -l staged
.PHONY: run-linter-staged

run-linter-diff: ## Perform analysis of local modified files
	cd $(DEVOPS_PATH)/cmd/app && chmod +x sast.sh && ./sast.sh -l diff
.PHONY: run-linter-diff

run-linter-ci: ## Perform analysis of modified files in continuous integration pipeline
	cd $(DEVOPS_PATH)/cmd/app && chmod +x sast.sh && ./sast.sh -l ci
.PHONY: run-linter-ci

run-linter-commit: ## Perform analysis of the commit message
	commitlint --edit .git/COMMIT_EDITMSG
.PHONY: run-linter-commit

run-sanitizer-app: ## Perform analysis of the application binary file
	cd $(DEVOPS_PATH)/cmd/app && chmod +x dast.sh && ./dast.sh -b $(PATH_BINARY_APP)
.PHONY: run-sanitizer-app

run-sanitizer-test: ## Perform analysis of the test binary file
	cd $(DEVOPS_PATH)/cmd/app && chmod +x dast.sh && ./dast.sh -b $(PATH_BINARY_TEST)
.PHONY: run-sanitizer-test

run-security-scan: ## Perform security analysis of local project
	cd $(DEVOPS_PATH)/cmd/app && chmod +x sca.sh && ./sca.sh -p $(@D)
.PHONY: run-security-scan

run-release: ## Perform release service task
	npx semantic-release
.PHONY: run-release

run-continuous-security: ## Perform task in continuous security pipeline
	$(MAKE) run-security-scan
.PHONY: run-continuous-security

run-continuous-integration: ## Perform task in continuous integration pipeline
	$(MAKE) run-linter-ci
.PHONY: run-continuous-integration

run-continuous-build: ## Perform task in continuous build pipeline
	$(MAKE) app-build
	$(MAKE) app-run
.PHONY: run-continuous-build

run-continuous-testing: ## Perform task in continuous testing pipeline
	$(MAKE) test-unit
	$(MAKE) test-cover
.PHONY: run-continuous-testing

run-continuous-release: ## Perform task in continuous release pipeline
	$(MAKE) run-release
.PHONY: run-continuous-release

setup-submodule-devops: ## Setup git submodule devops service
	git submodule add -f --name $(DEVOPS_NAME) $(DEVOPS_URL) $(DEVOPS_PATH)
.PHONY: setup-submodule-devops

teardown-submodule-devops: ## Teardown git submodule devops service
	git submodule deinit -f $(DEVOPS_PATH)
	rm -rf .git/modules/$(DEVOPS_NAME)
	git rm -rf $(DEVOPS_PATH)
.PHONY: teardown-submodule-devops

setup-submodule: ## Setup git submodules
	$(MAKE) setup-submodule-devops
.PHONY: setup-submodule

update-submodule: ## Update git submodules
	git submodule update --remote --recursive --merge --init
.PHONY: update-submodule

teardown-submodule: ## Remove git submodules
	$(MAKE) teardown-submodule-devops
.PHONY: teardown-submodule

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
	 docker run --rm -v "${PWD}:/workspace" -w /workspace golangci/golangci-lint:v2.7.2-alpine@sha256:1e1851102b736971267400e08b3e4b2e7799c73976a998820f6f6b6b86b48343 golangci-lint run ./...
.PHONY: analysis-golangci
