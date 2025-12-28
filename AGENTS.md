# AGENTS.md

- [1. Unit Testing](#1-unit-testing)
  - [1.1. Testing Patterns](#11-testing-patterns)
  - [1.2. Test Workflow](#12-test-workflow)
  - [1.3. Test Commands](#13-test-commands)
  - [1.4. Test Style](#14-test-style)
  - [1.5. Test Template](#15-test-template)

## 1. Unit Testing

Instructions for AI coding agents on automating unit test creation using consistent software testing patterns in this Go project.

1. Features and Benefits

    - Readability
      > Ensures high code quality and reliability. Tests are self-documenting, reducing cognitive load for reviewers and maintainers.

    - Consistency
      > Uniform structure across tests ensures predictable, familiar code that team members can navigate efficiently.

    - Scalability
      > Table-driven and data-driven approaches minimize boilerplate code when adding new test cases, making it simple to expand coverage.

    - Debuggability
      > Scoped traces and detailed assertion messages pinpoint failures quickly during continuous integration and local testing.

### 1.1. Testing Patterns

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

### 1.2. Test Workflow

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

    Structure all tests using this [template](#15-test-template) pattern.

### 1.3. Test Commands

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

### 1.4. Test Style

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

### 1.5. Test Template

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
