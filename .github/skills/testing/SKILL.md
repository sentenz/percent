---
name: testing
description: Comprehensive testing capabilities including unit tests, fuzz tests, benchmarks, and coverage analysis for the Go percent package
allowed-tools: bash, view, edit, grep, glob
version: 1.0.0
---

# Testing Skill

This skill provides instructions and capabilities for executing various types of tests in the percent project, including unit tests, fuzz tests, benchmarks, and coverage analysis.

## Purpose

Enable developers and AI agents to:
- Execute unit tests with race detection
- Run fuzz tests for discovering edge cases
- Measure performance with benchmarks
- Analyze code coverage

## Unit Testing

### Run Unit Tests

Execute all unit tests with race detection and generate JUnit XML report:

```bash
make go-test-unit
```

Output: `logs/unit/junit-report.xml`

Features:
- Uses standard Go `testing` package
- Enables race detection (`-race`)
- Generates JUnit XML report
- Parallel test execution with `t.Parallel()`

### Test Framework

- **Framework**: Standard Go `testing` package with `testing.T`
- **Comparison**: `github.com/google/go-cmp/cmp` for value comparisons
- **Error Checking**: `errors.Is()` for error validation

### Test Patterns

Tests follow these patterns:
- **In-Got-Want**: Explicit variable names for inputs, actual, and expected values
- **Table-Driven**: Test cases organized in slices of structs
- **Arrange-Act-Assert (AAA)**: Three distinct test phases

For detailed testing patterns and templates, see [AGENTS.md](../../../AGENTS.md).

## Fuzz Testing

### Run Fuzz Tests

Execute coverage-guided fuzz tests:

```bash
make go-test-fuzz
```

Tests executed:
- `FuzzPercent` - Test percentage calculation
- `FuzzOf` - Test percentage of value
- `FuzzChange` - Test percentage change
- `FuzzRemain` - Test percentage remaining
- `FuzzFromRatio` - Test ratio to percentage conversion
- `FuzzToRatio` - Test percentage to ratio conversion

Duration: 10 seconds per function

Features:
- Coverage-guided input generation
- Automatic crash detection
- Corpus-driven mutations
- Seed corpus in `testdata/fuzz/<FuzzName>/`

### Fuzz Framework

- **Framework**: Standard Go `testing` package with `testing.F`
- **Coverage**: Automatic coverage-guided fuzzing
- **Corpus**: Seeds stored in `testdata/fuzz/` directories

## Benchmark Testing

### Run Benchmarks

Measure execution time and memory allocations:

```bash
make go-test-bench
```

Features:
- Measures operations per second
- Tracks memory allocations
- Reports with `-benchmem` flag

### Compare Benchmarks

Compare benchmark results before and after changes:

```bash
make go-test-bench-compare
```

This command:
- Runs benchmarks 10 times
- Uses `benchstat` for statistical analysis
- Saves results to `logs/bench/`

### Benchmark Framework

- **Framework**: Standard Go `testing` package with `testing.B`
- **Loop Control**: Uses `b.Loop()` for iterations
- **Timer Management**: Automatic with `b.Loop()`
- **Sub-benchmarks**: Organized with `b.Run()`

## Coverage Analysis

### Generate Coverage Reports

Run tests with coverage tracking and generate reports:

```bash
make go-test-coverage
```

Output files:
- `logs/coverage/coverage.out` - Raw coverage data
- `logs/coverage/coverage.html` - Interactive HTML view
- `logs/coverage/coverage.xml` - Cobertura format for CI/CD

Features:
- Tracks line coverage
- Generates multiple report formats
- Integrates with CI/CD pipelines

### Coverage Requirements

The project maintains 100% code coverage through:
- Comprehensive test cases
- Edge case validation
- Error condition testing
- Boundary value testing

## Test Organization

Tests are located alongside the source code:
- `pkg/percent/percent_test.go` - Unit, fuzz, and benchmark tests
- `testdata/fuzz/` - Fuzz test corpus

## References

- [AGENTS.md](../../../AGENTS.md) - Detailed testing patterns and templates
- [README.md](../../../README.md) - Project documentation
- [Makefile](../../../Makefile) - Complete list of available commands
