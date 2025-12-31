# GitHub Copilot Instructions

This document provides instructions for GitHub Copilot when working with this repository.

## Project Overview

This is a Go package that provides utility functions for calculating percentages and performing related operations. The package is designed to be simple, reliable, and well-tested.

## Repository Structure

```
.
├── pkg/percent/          # Main package with percentage calculation functions
├── internal/pkg/resource/ # Internal resources (errors, constants, strings)
├── scripts/              # Bootstrap, setup, and teardown scripts
├── tests/policy/         # Rego policies for Conftest
├── .github/              # GitHub workflows and configuration
├── AGENTS.md             # AI agent instructions for testing patterns (HOW)
├── SKILLS.md             # Project capabilities and commands (WHAT)
└── Makefile              # Task automation
```

## Language and Tooling

- **Language**: Go 1.22.0+
- **Dependency Management**: Go modules with vendoring
- **Testing**: Standard Go testing package with `github.com/google/go-cmp/cmp` for comparisons
- **Linting**: golangci-lint (see `.golangci.yml`)
- **Security**: govulncheck for vulnerability scanning

## Build and Test Commands

Use the Makefile for all common tasks. Run `make help` to see all available commands.

### Essential Commands

```bash
# Development setup
make bootstrap          # Initialize development workspace
make setup              # Install dependencies

# Code quality
make go-fmt             # Format Go code
make go-vet             # Check for common mistakes
make go-vuln            # Check for vulnerabilities
make go-check           # Run all quality checks

# Testing
make go-test-unit       # Run unit tests with race detection
make go-test-coverage   # Generate coverage reports
make go-test-bench      # Run benchmarks
make go-test-fuzz       # Run fuzz tests

# Module management
make go-mod-tidy        # Tidy Go modules
make go-mod-vendor      # Vendor dependencies
```

## Coding Standards

### General Conventions

- **License Header**: All source files must include `// SPDX-License-Identifier: Apache-2.0` at the top
- **Line Length**: Max 100 characters (except for Go files where it's not enforced)
- **Line Endings**: LF (Unix-style)
- **Indentation**: Tabs for Go code, spaces for other files (see `.editorconfig`)
- **Imports**: Group standard library, external packages, and internal packages separately

### Go-Specific Standards

- Follow standard Go conventions and idioms
- Use meaningful variable names
- Export only what needs to be public
- Document all exported functions, types, and constants
- Use generics with `golang.org/x/exp/constraints` for numeric types
- Return errors, don't panic (except in truly exceptional circumstances)

### Error Handling

- Define errors in `internal/pkg/resource/errors.go`
- Use `errors.New()` for sentinel errors
- Check errors using `errors.Is()` in tests
- Return descriptive error messages defined in `internal/pkg/resource/strings.go`

### Package Organization

- `pkg/percent/`: Public API functions for percentage calculations
- `internal/pkg/resource/`: Internal constants, errors, and shared resources
- Tests live alongside the code they test (e.g., `percent_test.go` next to `percent.go`)

## Testing Guidelines

**IMPORTANT**: This repository follows specific testing patterns documented in `AGENTS.md`. When creating or modifying tests, follow the guidelines in that file.

### Testing Patterns (Quick Reference)

See `AGENTS.md` for comprehensive testing guidelines. Key patterns:

1. **In-Got-Want Pattern**: Use explicit variable names (`in` for inputs, `got` for actual, `want` for expected)
2. **Table-Driven Testing**: Organize test cases in slices of structs
3. **Arrange-Act-Assert (AAA)**: Structure tests in three distinct phases

### Test Template

Consolidate test cases for a single function into ONE `TestXxx(t *testing.T)` function:

```go
func TestFunctionName(t *testing.T) {
 t.Parallel()

 type in struct {
  // input fields
 }

 type want struct {
  // expected output fields
  err error
 }

 tests := []struct {
  name string
  in   in
  want want
 }{
  {
   name: "case description",
   in: in{
    // input values
   },
   want: want{
    // expected output
    err: nil,
   },
  },
  // add more cases as needed
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   // Arrange
   // additional setup as needed

   // Act
   got, err := FunctionName(tt.in.field)

   // Assert
   if !errors.Is(err, tt.want.err) {
    t.Errorf("FunctionName() error = %v, want err %v", err, tt.want.err)
   }
   if !cmp.Equal(got, tt.want.value) {
    t.Errorf("FunctionName(%+v) = %v, want %v", tt.in, got, tt.want.value)
   }
  })
 }
}
```

### Test Coverage Requirements

Include comprehensive edge cases:
- Boundary values (min/max limits, edge thresholds)
- Empty/null inputs
- Null pointers and invalid references
- Overflow/underflow scenarios
- Special cases (negative numbers, zero, special states)
- Error conditions

### Test Assertions

- Use `github.com/google/go-cmp/cmp` for value comparisons
- Use `errors.Is()` for error checking
- Include descriptive error messages in assertions

## Documentation

- Update `README.md` for user-facing changes
- Update `AGENTS.md` when changing testing patterns or conventions
- Document all exported functions with godoc-style comments
- Include examples in documentation where helpful

## Workflow and CI/CD

The repository uses GitHub Actions for CI/CD:
- `go-tests.yml`: Runs unit tests and coverage
- `golangci.yml`: Lints code with golangci-lint
- `semgrep.yml`: Security analysis with Semgrep
- `semantic-release.yml`: Automated releases
- `renovate.yml`: Dependency updates
- `conftest.yml`: Policy validation with Conftest
- `regal.yml`: Rego policy linting

## Security

- Run `make go-vuln` to check for known vulnerabilities
- Never commit secrets or sensitive data
- Security scanning is performed in CI/CD

## Making Changes

1. Make minimal, focused changes
2. Follow existing code patterns and conventions
3. Add/update tests for all changes (following patterns in `AGENTS.md`)
4. Run `make go-check` before committing
5. Run `make go-test-unit` to verify tests pass
6. Ensure all CI checks pass

## References

- **AI Agent Instructions**: See `AGENTS.md` for testing patterns and workflows (HOW to test)
- **Project Capabilities**: See `SKILLS.md` for commands, tools, and features (WHAT is available)
- **Task Automation**: See `Makefile` for all available commands
- **Development Setup**: See `scripts/README.md` for bootstrap instructions
- **Editor Config**: See `.editorconfig` for formatting rules
