# AGENTS.md

- [1. Unit Testing](#1-unit-testing)
  - [1.1. Testing Patterns](#11-testing-patterns)
  - [1.2. Test Workflow](#12-test-workflow)
  - [1.3. Test Commands](#13-test-commands)
  - [1.4. Test Style](#14-test-style)
  - [1.5. Test Template](#15-test-template)
- [2. Fuzz Testing](#2-fuzz-testing)
  - [2.1. Fuzz Testing Patterns](#21-fuzz-testing-patterns)
  - [2.2. Fuzz Test Workflow](#22-fuzz-test-workflow)
  - [2.3. Fuzz Test Commands](#23-fuzz-test-commands)
  - [2.4. Fuzz Test Style](#24-fuzz-test-style)
  - [2.5. Fuzz Test Template](#25-fuzz-test-template)

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

- Run Fuzz Tests
  > Execute fuzz tests.

  ```bash
  make go-test-fuzz
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

## 2. Fuzz Testing

Instructions for AI coding agents on automating fuzz test creation using consistent software testing patterns in this Go project.

1. Features and Benefits

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

### 2.1. Fuzz Testing Patterns

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

### 2.2. Fuzz Test Workflow

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

    Structure all fuzz tests using this [template](#25-fuzz-test-template) pattern.

5. Seed Corpus

    Optionally provide seed inputs in `testdata/fuzz/<FuzzTestName>/` directory to guide fuzzing toward interesting inputs.

### 2.3. Fuzz Test Commands

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

### 2.4. Fuzz Test Style

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

### 2.5. Fuzz Test Template

Use this template for new fuzz test functions. Replace placeholders with actual values and adjust as needed for the use case.

- Multi-Parameter Functions
  > For functions with multiple parameters, use a struct array to define test cases.

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

- Single-Parameter Functions
  > For functions with a single parameter, use a slice array to define test cases.

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
