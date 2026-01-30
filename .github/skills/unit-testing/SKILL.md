---
name: unit-testing
description: Automates unit test creation for Go projects using consistent software testing patterns including In-Got-Want, Table-Driven Testing, and AAA patterns. Use when creating, modifying, or reviewing unit tests, or when the user mentions unit tests, test coverage, or Go testing.
metadata:
  version: "1.0"
  activation:
    implicit: true
    priority: 1
    triggers:
      - "unit test"
      - "go test"
      - "testing"
      - "create test"
      - "add test"
      - "write test"
      - "test coverage"
    match:
      languages: ["go", "golang"]
      paths: ["pkg/**/*_test.go", "internal/**/*_test.go"]
      prompt_regex: "(?i)(unit test|go test|testing|create test|add test|write test|test coverage)"
  usage:
    load_on_prompt: true
    autodispatch: true
---

# Unit Testing

Instructions for AI coding agents on automating unit test creation using consistent software testing patterns in this Go project.

- [1. Features and Benefits](#1-features-and-benefits)
- [2. Unit Testing Patterns](#2-unit-testing-patterns)
- [3. Unit Test Workflow](#3-unit-test-workflow)
- [4. Unit Test Commands](#4-unit-test-commands)
- [5. Unit Test Style](#5-unit-test-style)
- [6. Unit Test Template](#6-unit-test-template)

## 1. Features and Benefits

- Readability
  > Ensures high code quality and reliability. Tests are self-documenting, reducing cognitive load for reviewers and maintainers.

- Consistency
  > Uniform structure across tests ensures predictable, familiar code that team members can navigate efficiently.

- Scalability
  > Table-driven and data-driven approaches minimize boilerplate code when adding new test cases, making it simple to expand coverage.

- Debuggability
  > Scoped traces and detailed assertion messages pinpoint failures quickly during continuous integration and local testing.

## 2. Unit Testing Patterns

- In-Got-Want
  > In-Got-Want is a software testing pattern that structures test cases into three distinct sections of In (input), Got (actual output), and Want (expected output).

- Table-Driven Testing
  > Table-Driven Testing is a software testing technique in which test cases are organized in a tabular format.

- Data-Driven Testing (DDT)
  > Data-Driven Testing (DDT) is a software testing methodology that separates test data from test logic, allowing the same test logic to be executed with multiple sets of input data (e.g., JSON, CSV).

- Arrange, Act, Assert (AAA)
  > Arrange, Act, Assert (AAA) is a software testing pattern that structures test cases into three distinct phases of Arrange (setup), Act (execution), and Assert (verification).

- Test Fixtures
  > Test Fixtures are a software testing pattern that provides a consistent and reusable setup and teardown mechanism for test cases.

## 3. Unit Test Workflow

1. Identify

    Identify new functions in `pkg/` or `internal/` (e.g., `pkg/<package>/<file>.go`).

2. Add/Create

    Create new tests in the same package (e.g., `pkg/<package>/<file>_test.go`).

3. Test Coverage Requirements

    Include comprehensive edge cases:
    - Coverage-guided cases
    - Boundary values (min/max limits, edge thresholds)
    - Empty/null inputs
    - Null pointers and invalid references
    - Overflow/underflow scenarios
    - Special cases (negative numbers, zero, special states)

4. Apply Templates

    Structure all tests using the [template](#6-unit-test-template) pattern.

## 4. Unit Test Commands

- Run Unit Tests
  > Execute tests with race detection and generate JUnit XML report.

  ```bash
  make go-test-unit
  ```

- Run Code Coverage
  > Generate coverage reports (HTML and XML).

  ```bash
  make go-test-coverage
  ```

- Run Benchmarks
  > Execute benchmarks.

  ```bash
  make go-test-bench
  ```

- Run Fuzz Tests
  > Execute fuzz tests.

  ```bash
  make go-test-fuzz
  ```

## 5. Unit Test Style

- Test Framework
  > Use the standard Go `testing` package.

- Include Imports
  > Include `testing` and `github.com/google/go-cmp/cmp` for comparisons.

- Parallelism
  > Use `t.Parallel()` to run tests in parallel.

- Test Organization
  > Consolidate test cases for a single function into **one `TestXxx(t *testing.T)` function** using table-driven testing. This approach:
  > - Eliminates redundant test function definitions
  > - Simplifies maintenance by grouping related scenarios together
  > - Reduces code duplication in setup and teardown phases
  > - Makes it easier to add or modify test cases

- Assertions
  > Use `cmp.Equal` for value comparisons and `errors.Is` for error checking.

## 6. Unit Test Template

Use this template (In-Got-Want + Table-Driven + AAA) for new test functions. Replace placeholders with actual values and adjust as needed for the use case.

```go
func Test<FunctionName>(t *testing.T) {
	t.Parallel()

	type in struct {
		/* input fields */
	}

	type want struct {
		/* expected output fields */
		err error
	}

	tests := []struct {
		name string
		in   in
		want want
	}{
		{
			name: "case description 1",
			in: in{
				/* input values */
			},
			want: want{
				/* expected output */
				err: nil,
			},
		},
		{
			name: "case description 2",
			in: in{
				/* input values */
			},
			want: want{
				/* expected output */
				err: nil, // or specific error
			},
		},
		// add more cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			// additional setup as needed

			// Act
			got, err := <Function>(tt.in.<input>)

			// Assert
			if !errors.Is(err, tt.want.err) {
				t.Errorf("<Function>() error = %v, want err %v", err, tt.want.err)
			}
			if !cmp.Equal(got, tt.want.<value>) {
				t.Errorf("<Function>(%+v) = %v, want %v", tt.in, got, tt.want.<value>)
			}
		})
	}
}
```
