---
name: fuzz-testing
description: Automates fuzz test creation for Go projects using native fuzzing engine with coverage-guided testing. Use when creating, modifying, or reviewing fuzz tests, or when the user mentions fuzz testing, fuzzing, or security testing.
metadata:
  version: "1.0"
  activation:
    implicit: true
    priority: 1
    triggers:
      - "fuzz test"
      - "fuzzing"
      - "fuzz testing"
      - "create fuzz test"
      - "add fuzz test"
      - "write fuzz test"
      - "security test"
    match:
      languages: ["go", "golang"]
      paths: ["pkg/**/*_test.go", "internal/**/*_test.go", "testdata/fuzz/**"]
      prompt_regex: "(?i)(fuzz test|fuzzing|fuzz testing|create fuzz|add fuzz|write fuzz|security test)"
  usage:
    load_on_prompt: true
    autodispatch: true
---

# Fuzz Testing

Instructions for AI coding agents on automating fuzz test creation using consistent software testing patterns in this Go project.

- [1. Features and Benefits](#1-features-and-benefits)
- [2. Fuzz Testing Patterns](#2-fuzz-testing-patterns)
- [3. Fuzz Test Workflow](#3-fuzz-test-workflow)
- [4. Fuzz Test Commands](#4-fuzz-test-commands)
- [5. Fuzz Test Style](#5-fuzz-test-style)
- [6. Fuzz Test Templates](#6-fuzz-test-templates)

## 1. Features and Benefits

- Coverage-Guided Exploration
  > Go's native fuzzing engine uses code coverage instrumentation to automatically guide input generation toward unexplored code paths, maximizing bug discovery without manual intervention.

- Automated Input Generation
  > Fuzz testing automatically generates random inputs to discover edge cases and unexpected behaviors that manual testing might miss.

- Security Vulnerability Discovery
  > Helps identify security vulnerabilities, crashes, and undefined behaviors by testing with malformed, unexpected, or extreme inputs.

- Continuous Testing
  > Fuzz tests can run continuously to explore the input space over time, discovering new edge cases as the corpus grows.

- Regression Prevention
  > Once a crash or bug is found, the input is saved in the corpus to prevent regression in future test runs.

## 2. Fuzz Testing Patterns

- Coverage-Guided Fuzzing
  > Coverage-Guided Fuzzing is the primary fuzzing technique used by Go's native fuzzing engine. It automatically instruments code to track coverage and guides input generation toward unexplored code paths, maximizing code exploration and bug discovery.

- Corpus-Driven Fuzzing
  > Corpus-Driven Fuzzing is a software testing technique that uses a collection of seed inputs (corpus) as the starting point for generating new test inputs through mutation.

- Property-Based Testing
  > Property-Based Testing is a testing approach that verifies invariants and properties that should hold true for all inputs, rather than testing specific input-output pairs.

- Boundary Value Fuzzing
  > Boundary Value Fuzzing focuses on testing edge cases and boundary conditions with randomly generated inputs around critical thresholds.

- Crash Detection
  > Crash Detection is the process of identifying inputs that cause panics, runtime errors, or undefined behavior in the code under test.

## 3. Fuzz Test Workflow

1. Identify

    Identify functions in `pkg/` or `internal/` that accept external inputs, perform calculations, or have edge cases worth fuzzing (e.g., `pkg/<package>/<file>.go`).

2. Add/Create

    Create fuzz tests in the same package (e.g., `pkg/<package>/<file>_test.go`).

3. Fuzz Test Coverage Requirements

    Focus on functions that:
    - Accept numeric inputs (integers, floats)
    - Perform mathematical operations (division, multiplication)
    - Have boundary conditions (min/max values, zero checks)
    - Return errors for invalid inputs
    - Use generics or type constraints

4. Apply Templates

    Structure all fuzz tests using the [template](#6-fuzz-test-templates) pattern.

5. Seed Corpus

    Optionally provide seed inputs in `testdata/fuzz/<FuzzTestName>/` directory to guide fuzzing toward interesting inputs.

## 4. Fuzz Test Commands

- Run Fuzz Tests
  > Execute fuzz tests for a specified duration.

  ```bash
  make go-test-fuzz
  ```

- View Fuzz Corpus
  > Inspect the seed corpus and generated inputs.

  ```bash
  ls -la testdata/fuzz/<FuzzTestName>/
  ```

## 5. Fuzz Test Style

- Test Framework
  > Use the standard Go `testing` package with `testing.F` for fuzz tests. Go's fuzzing engine automatically uses coverage-guided fuzzing to explore code paths.

- Coverage-Guided Behavior
  > The fuzzing engine tracks code coverage during execution and prioritizes inputs that explore new code paths. No manual configuration is required - coverage guidance is automatic.

- Include Imports
  > Include `testing` and any packages needed for the function under test.

- Fuzz Function Naming
  > Name fuzz functions with the `Fuzz` prefix followed by the function name (e.g., `FuzzPercent` for testing `Percent()`).

- Seed Corpus
  > Use `f.Add()` to provide seed inputs that cover important edge cases and known valid/invalid inputs. The fuzzer will mutate these seeds while maximizing coverage.

- Fuzz Target
  > The fuzz target function receives `*testing.T` and randomly generated inputs. It should:
  > - Validate inputs before calling the function under test (skip invalid inputs with `t.Skip()` if necessary)
  > - Call the function with fuzzed inputs
  > - Assert invariants and properties that must always hold true
  > - Not crash or panic for any input

- Error Handling
  > Fuzz tests should verify that functions handle errors gracefully without panicking.

- Assertions
  > Use explicit checks with `t.Errorf()` or `t.Fatalf()` to report violations of expected properties.

## 6. Fuzz Test Templates

Use these templates for new fuzz test functions. Replace placeholders with actual values and adjust as needed for the use case.

### 6.1. Multi-Parameter Functions

For functions with multiple parameters, use a struct array to define test cases.

```go
func Fuzz<FunctionName>(f *testing.F) {
	// Seed corpus with edge cases using testcases array
	testcases := []struct {
		param1 <type>
		param2 <type>
		// Add more parameters as needed
	}{
		{<value1>, <value2>}, // description of test case
		{<value1>, <value2>}, // description of test case
		// Add more test cases
	}
	for _, tc := range testcases {
		f.Add(tc.param1, tc.param2) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, param1 <type>, param2 <type>) {
		// Arrange
		// Optional: skip invalid inputs or prepare test conditions
		// Example: if input < 0 { t.Skip("negative inputs not interesting") }

		// Act
		got, err := <Function>(param1, param2)

		// Assert
		// Verify properties that should always hold true
		// Example 1: Function should never panic
		// Example 2: If no error, result should meet certain properties
		// Example 3: If error, result should be in expected error state

		if err != nil {
			// Verify error cases
			// Example: if got != 0 { t.Errorf("expected zero result on error, got %v", got) }
		} else {
			// Verify success cases and invariants
			// Example: if got < 0 { t.Errorf("result should be non-negative, got %v", got) }
		}
	})
}
```

### 6.2. Single-Parameter Functions

For functions with a single parameter, use a slice array to define test cases.

```go
func Fuzz<FunctionName>(f *testing.F) {
	// Seed corpus with edge cases using testcases array
	testcases := []<type>{
		<value1>, // description
		<value2>, // description
		// Add more test cases
	}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}

	f.Fuzz(func(t *testing.T, param <type>) {
		// Arrange
		// Optional: skip invalid inputs or prepare test conditions

		// Act
		got, err := <Function>(param)

		// Assert
		// Verify properties and invariants
		if err != nil {
			// Verify error cases
		} else {
			// Verify success cases
		}
	})
}
```
