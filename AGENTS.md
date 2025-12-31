# AGENTS.md

AI Coding Agent Instructions for Software Testing in Go Projects

- [1. Purpose and Scope](#1-purpose-and-scope)
  - [1.1. Purpose](#11-purpose)
  - [1.2. Scope](#12-scope)
  - [1.3. Complementary Documentation](#13-complementary-documentation)
- [2. Software Testing](#2-software-testing)
  - [2.1. Unit Testing](#21-unit-testing)
    - [2.1.1. Unit Testing Patterns](#211-unit-testing-patterns)
    - [2.1.2. Unit Test Workflow](#212-unit-test-workflow)
    - [2.1.3. Unit Test Commands](#213-unit-test-commands)
    - [2.1.4. Unit Test Style](#214-unit-test-style)
    - [2.1.5. Unit Test Template](#215-unit-test-template)
    - [2.1.6. Unit Test Examples](#216-unit-test-examples)
  - [2.2. Fuzz Testing](#22-fuzz-testing)
    - [2.2.1. Fuzz Testing Patterns](#221-fuzz-testing-patterns)
    - [2.2.2. Fuzz Test Workflow](#222-fuzz-test-workflow)
    - [2.2.3. Fuzz Test Commands](#223-fuzz-test-commands)
    - [2.2.4. Fuzz Test Style](#224-fuzz-test-style)
    - [2.2.5. Fuzz Test Template](#225-fuzz-test-template)
    - [2.2.6. Fuzz Test Examples](#226-fuzz-test-examples)
  - [2.3. Benchmark Testing](#23-benchmark-testing)
    - [2.3.1. Benchmark Testing Patterns](#231-benchmark-testing-patterns)
    - [2.3.2. Benchmark Test Workflow](#232-benchmark-test-workflow)
    - [2.3.3. Benchmark Test Commands](#233-benchmark-test-commands)
    - [2.3.4. Benchmark Test Style](#234-benchmark-test-style)
    - [2.3.5. Benchmark Test Template](#235-benchmark-test-template)
    - [2.3.6. Benchmark Test Examples](#236-benchmark-test-examples)

## 1. Purpose and Scope

### 1.1. Purpose

This document provides **prescriptive instructions** for AI coding agents on **how to automate test creation** using consistent software testing patterns in this Go project. It defines the testing methodology, patterns, workflows, and templates that agents must follow when generating or modifying tests.

**Key Objectives:**

1. **Consistency**: Ensure all AI-generated tests follow uniform patterns and structure
2. **Quality**: Maintain high code quality through comprehensive test coverage
3. **Maintainability**: Create self-documenting, readable tests that are easy to maintain
4. **Automation**: Enable AI agents to generate tests autonomously with minimal human intervention

### 1.2. Scope

**This document covers:**

- ✅ Testing patterns and methodologies (In-Got-Want, Table-Driven, AAA)
- ✅ Detailed workflows for creating unit, fuzz, and benchmark tests
- ✅ Code templates and style guidelines for AI agents
- ✅ Examples demonstrating proper test implementation
- ✅ Coverage requirements and edge case identification

**This document does NOT cover:**

- ❌ Project structure and organization (see [SKILLS.md](./SKILLS.md))
- ❌ Build and deployment processes (see [README.md](./README.md))
- ❌ Development environment setup (see [.github/copilot-instructions.md](./.github/copilot-instructions.md))
- ❌ CI/CD workflow details (see [SKILLS.md](./SKILLS.md))

### 1.3. Complementary Documentation

- **[SKILLS.md](./SKILLS.md)**: Describes WHAT capabilities, commands, and tools are available in the project
- **[README.md](./README.md)**: User-facing documentation for package usage and contribution
- **[.github/copilot-instructions.md](./.github/copilot-instructions.md)**: GitHub Copilot-specific instructions
- **[Makefile](./Makefile)**: Complete list of available commands and tasks

## 2. Software Testing

### 2.1. Unit Testing

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

#### 2.1.1. Unit Testing Patterns

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

#### 2.1.2. Unit Test Workflow

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

    Structure all tests using this [template](#215-unit-test-template) pattern.

#### 2.1.3. Unit Test Commands

See [SKILLS.md - Unit Testing](./SKILLS.md#41-unit-testing) for detailed command documentation.

**Quick Reference:**

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

#### 2.1.4. Unit Test Style

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

#### 2.1.5. Unit Test Template

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

#### 2.1.6. Unit Test Examples

**Example 1: Testing a simple percentage calculation function**

```go
// Function to test: Percent calculates what percentage of a value is
func Percent[T Number](percent T, value T) (T, error)

// Test implementation
func TestPercent(t *testing.T) {
 t.Parallel()

 type in struct {
  percent float64
  value   float64
 }

 type want struct {
  result float64
  err    error
 }

 tests := []struct {
  name string
  in   in
  want want
 }{
  {
   name: "25 percent of 200",
   in: in{
    percent: 25.0,
    value:   200.0,
   },
   want: want{
    result: 50.0,
    err:    nil,
   },
  },
  {
   name: "zero percent",
   in: in{
    percent: 0.0,
    value:   100.0,
   },
   want: want{
    result: 0.0,
    err:    nil,
   },
  },
  {
   name: "negative percent error",
   in: in{
    percent: -10.0,
    value:   100.0,
   },
   want: want{
    result: 0.0,
    err:    resource.ErrInvalidPercent,
   },
  },
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   // Arrange
   // (no additional setup needed)

   // Act
   got, err := Percent(tt.in.percent, tt.in.value)

   // Assert
   if !errors.Is(err, tt.want.err) {
    t.Errorf("Percent() error = %v, want err %v", err, tt.want.err)
   }
   if !cmp.Equal(got, tt.want.result) {
    t.Errorf("Percent(%+v) = %v, want %v", tt.in, got, tt.want.result)
   }
  })
 }
}
```

**Example 2: Testing with complex error scenarios**

```go
func TestOf(t *testing.T) {
 t.Parallel()

 type in struct {
  value float64
  total float64
 }

 type want struct {
  percent float64
  err     error
 }

 tests := []struct {
  name string
  in   in
  want want
 }{
  {
   name: "50 of 200 is 25 percent",
   in: in{
    value: 50.0,
    total: 200.0,
   },
   want: want{
    percent: 25.0,
    err:     nil,
   },
  },
  {
   name: "division by zero error",
   in: in{
    value: 50.0,
    total: 0.0,
   },
   want: want{
    percent: 0.0,
    err:     resource.ErrDivisionByZero,
   },
  },
  {
   name: "negative value error",
   in: in{
    value: -50.0,
    total: 200.0,
   },
   want: want{
    percent: 0.0,
    err:     resource.ErrNegativeValue,
   },
  },
 }

 for _, tt := range tests {
  t.Run(tt.name, func(t *testing.T) {
   // Arrange
   // (no additional setup needed)

   // Act
   got, err := Of(tt.in.value, tt.in.total)

   // Assert
   if !errors.Is(err, tt.want.err) {
    t.Errorf("Of() error = %v, want err %v", err, tt.want.err)
   }
   if !cmp.Equal(got, tt.want.percent) {
    t.Errorf("Of(%+v) = %v, want %v", tt.in, got, tt.want.percent)
   }
  })
 }
}
```

### 2.2. Fuzz Testing

Instructions for AI coding agents on **how to automate fuzz test creation** using consistent software testing patterns in this Go project.

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

#### 2.2.1. Fuzz Testing Patterns

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

#### 2.2.2. Fuzz Test Workflow

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

    Structure all fuzz tests using this [template](#225-fuzz-test-template) pattern.

5. Seed Corpus

    Optionally provide seed inputs in `testdata/fuzz/<FuzzTestName>/` directory to guide fuzzing toward interesting inputs.

#### 2.2.3. Fuzz Test Commands

See [SKILLS.md - Fuzz Testing](./SKILLS.md#42-fuzz-testing) for detailed command documentation.

**Quick Reference:**

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

#### 2.2.4. Fuzz Test Style

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

#### 2.2.5. Fuzz Test Template

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

#### 2.2.6. Fuzz Test Examples

**Example 1: Fuzzing a percentage calculation function**

```go
func FuzzPercent(f *testing.F) {
 // Seed corpus with edge cases
 testcases := []struct {
  percent float64
  value   float64
 }{
  {0.0, 100.0},      // Zero percent
  {100.0, 100.0},    // Full percent
  {50.0, 200.0},     // Normal case
  {-10.0, 100.0},    // Negative percent (should error)
  {110.0, 100.0},    // Over 100 percent
  {0.5, 1000000.0},  // Small percent, large value
 }
 for _, tc := range testcases {
  f.Add(tc.percent, tc.value)
 }

 f.Fuzz(func(t *testing.T, percent float64, value float64) {
  // Arrange
  // Skip NaN and Inf values
  if math.IsNaN(percent) || math.IsInf(percent, 0) ||
   math.IsNaN(value) || math.IsInf(value, 0) {
   t.Skip("skipping NaN or Inf values")
  }

  // Act
  got, err := Percent(percent, value)

  // Assert
  // Property 1: Function should never panic
  // Property 2: If error, result should be zero
  if err != nil {
   if got != 0.0 {
    t.Errorf("expected zero result on error, got %v", got)
   }
  } else {
   // Property 3: Valid result should be calculable
   expected := (percent / 100.0) * value
   if !almostEqual(got, expected, 0.0001) {
    t.Errorf("Percent(%v, %v) = %v, expected %v", percent, value, got, expected)
   }
  }
 })
}
```

**Example 2: Fuzzing with single parameter**

```go
func FuzzFromRatio(f *testing.F) {
 // Seed corpus with boundary values
 testcases := []float64{
  0.0,   // Minimum
  0.5,   // Half
  1.0,   // Maximum
  -0.1,  // Below minimum
  1.5,   // Above maximum
  0.001, // Very small
  0.999, // Almost 1
 }
 for _, tc := range testcases {
  f.Add(tc)
 }

 f.Fuzz(func(t *testing.T, ratio float64) {
  // Arrange
  // Skip NaN and Inf
  if math.IsNaN(ratio) || math.IsInf(ratio, 0) {
   t.Skip("skipping NaN or Inf")
  }

  // Act
  got, err := FromRatio(ratio)

  // Assert
  if err != nil {
   // Property 1: Errors occur for out-of-range ratios
   if ratio >= 0 && ratio <= 1 {
    t.Errorf("unexpected error for valid ratio %v: %v", ratio, err)
   }
  } else {
   // Property 2: Result should be ratio * 100
   expected := ratio * 100.0
   if !almostEqual(got, expected, 0.0001) {
    t.Errorf("FromRatio(%v) = %v, expected %v", ratio, got, expected)
   }
   // Property 3: Result should be in range [0, 100]
   if got < 0 || got > 100 {
    t.Errorf("FromRatio(%v) = %v, out of range [0, 100]", ratio, got)
   }
  }
 })
}
```

### 2.3. Benchmark Testing

Instructions for AI coding agents on **how to automate benchmark test creation** using consistent software testing patterns in this Go project.

1. Features and Benefits

    - Performance Measurement
      > Benchmark tests measure the execution time and memory allocation of functions, providing quantifiable metrics for performance analysis.

    - Regression Detection
      > Continuous benchmarking helps identify performance regressions early in the development cycle before they reach production.

    - Optimization Guidance
      > Benchmark results guide optimization efforts by identifying bottlenecks and quantifying the impact of performance improvements.

    - Comparative Analysis
      > Benchmarks enable comparison of different implementations or algorithms to make informed decisions about performance trade-offs.

    - Resource Profiling
      > Memory allocation tracking helps identify unnecessary allocations and optimize memory usage patterns.

#### 2.3.1. Benchmark Testing Patterns

- Microbenchmarking
  > Microbenchmarking is a software testing technique that measures the performance of small, isolated code units to identify performance characteristics and bottlenecks.

- Comparative Benchmarking
  > Comparative Benchmarking is a testing approach that compares the performance of different implementations or algorithms side-by-side using consistent workloads.

- Memory Profiling
  > Memory Profiling is the process of measuring memory allocations and usage patterns during benchmark execution using `-benchmem` flag.

- Statistical Benchmarking
  > Statistical Benchmarking uses multiple iterations to calculate statistical measures (mean, variance) to ensure reliable and reproducible results.

- Sub-benchmarks
  > Sub-benchmarks organize related benchmark cases using `b.Run()` to group variations of the same function with different input scenarios.

- Table-Driven Testing
  > Table-Driven Testing is a software testing technique in which benchmark cases are organized in a tabular format to systematically cover different input scenarios.

#### 2.3.2. Benchmark Test Workflow

1. Identify

    Identify performance-critical functions in `pkg/` or `internal/` that benefit from performance tracking (e.g., `pkg/<package>/<file>.go`).

2. Add/Create

    Create benchmark tests in the same package (e.g., `pkg/<package>/<file>_test.go`).

3. Benchmark Test Coverage Requirements

    Focus on functions that:
    - Are called frequently in hot paths
    - Perform mathematical operations or calculations
    - Process data structures or collections
    - Have multiple implementation approaches to compare
    - Are candidates for optimization

4. Apply Templates

    Structure all benchmark tests using this [template](#235-benchmark-test-template) pattern.

5. Baseline Measurements

    Establish performance baselines by running benchmarks on stable code before making changes.

#### 2.3.3. Benchmark Test Commands

See [SKILLS.md - Benchmark Testing](./SKILLS.md#43-benchmark-testing) for detailed command documentation.

**Quick Reference:**

- Run Benchmarks
  > Execute all benchmarks with memory statistics.

  ```bash
  make go-test-bench
  ```

- Run Specific Benchmark
  > Execute a specific benchmark function.

  ```bash
  go test -bench=BenchmarkPercent -benchmem ./pkg/percent
  ```

- Run with CPU Profiling
  > Generate CPU profile for performance analysis.

  ```bash
  go test -bench=. -benchmem -cpuprofile=cpu.prof ./pkg/percent
  ```

- Run with Memory Profiling
  > Generate memory profile for allocation analysis.

  ```bash
  go test -bench=. -benchmem -memprofile=mem.prof ./pkg/percent
  ```

- Benchmark Time Control
  > Run benchmarks for a specific duration.

  ```bash
  go test -bench=. -benchtime=10s ./pkg/percent
  ```

- Compare Benchmark Results
  > Use `benchstat` to compare benchmark results before and after changes.

  ```bash
  go test -bench=. -benchmem -count=10 ./pkg/percent > old.txt
  # Make changes
  go test -bench=. -benchmem -count=10 ./pkg/percent > new.txt
  benchstat old.txt new.txt
  ```

#### 2.3.4. Benchmark Test Style

- Test Framework
  > Use the standard Go `testing` package with `testing.B` for benchmark tests.

- Include Imports
  > Include `testing` and any packages needed for the function under test.

- Benchmark Function Naming
  > Name benchmark functions with the `Benchmark` prefix followed by the function name (e.g., `BenchmarkPercent` for testing `Percent()`).

- Benchmark Loop
  > Use `b.Loop()` to control the number of iterations. The testing framework automatically adjusts the loop iterations to get reliable timing measurements. `b.Loop()` is preferred over `b.N` as it provides better integration with the testing framework and more accurate measurements. Unlike `b.N`-style benchmarks, `b.Loop()` integrates timer management, it automatically handles `b.ResetTimer()` at the loop's start and `b.StopTimer()` at its end, eliminating the need to manually manage the benchmark timer for setup and cleanup code.

- Timer Control
  > When using `b.Loop()`, timer management is automatic and no manual `b.ResetTimer()`, `b.StopTimer()`, or `b.StartTimer()` calls are needed for typical benchmarks. For advanced scenarios not using `b.Loop()`, use `b.ResetTimer()` to exclude setup time from measurements and `b.StopTimer()`/`b.StartTimer()` to exclude specific operations.

- Sub-benchmarks
  > Use `b.Run()` to organize related benchmark cases with different input scenarios. Each sub-benchmark runs independently with its own `b.N` iterations.

- Memory Reporting
  > Use `b.ReportAllocs()` to report memory allocations per operation when not using `-benchmem` flag.

- Result Validation
  > Optionally validate results in benchmarks to prevent compiler optimizations from eliminating dead code.

#### 2.3.5. Benchmark Test Template

Use this template for new benchmark test functions. Replace placeholders with actual values and adjust as needed for the use case.

- Multi-Scenario Benchmarks
  > For benchmarking multiple scenarios or input variations, use sub-benchmarks with table-driven approach.

  ```go
  func Benchmark<FunctionName>(b *testing.B) {
   // Define benchmark cases with different scenarios
   benchmarks := []struct {
    name   string
    param1 <type>
    param2 <type>
    // Add more parameters as needed
   }{
    {
     name:   "scenario description 1",
     param1: <value1>,
     param2: <value2>,
    },
    {
     name:   "scenario description 2",
     param1: <value1>,
     param2: <value2>,
    },
    // Add more benchmark cases
   }

   for _, bm := range benchmarks {
    b.Run(bm.name, func(b *testing.B) {
     // Arrange
     // Setup code here (automatically excluded from timing by b.Loop)

     // Act
     for b.Loop() {
      _, _ = <Function>(bm.param1, bm.param2)
     }
    })
   }
  }
  ```

- Simple Benchmarks
  > For benchmarking a single scenario, use a simple loop without sub-benchmarks.

  ```go
  func Benchmark<FunctionName>(b *testing.B) {
   // Arrange
   // Setup code here (automatically excluded from timing by b.Loop)
   param1 := <value1>
   param2 := <value2>

   // Act
   for b.Loop() {
    _, _ = <Function>(param1, param2)
   }
  }
  ```

- Benchmarks with Validation
  > For benchmarks that need to prevent compiler optimizations, store results in package-level variables.

  ```go
  var (
   benchResult <type>
   benchError  error
  )

  func Benchmark<FunctionName>(b *testing.B) {
   // Arrange
   // Setup code here (automatically excluded from timing by b.Loop)
   param1 := <value1>
   param2 := <value2>

   // Act
   for b.Loop() {
    benchResult, benchError = <Function>(param1, param2)
   }
  }
  ```

#### 2.3.6. Benchmark Test Examples

**Example 1: Simple benchmark for percentage calculation**

```go
func BenchmarkPercent(b *testing.B) {
 // Arrange
 percent := 25.0
 value := 200.0

 // Act
 for b.Loop() {
  _, _ = Percent(percent, value)
 }
}
```

**Example 2: Multi-scenario benchmark with table-driven approach**

```go
func BenchmarkOf(b *testing.B) {
 // Define benchmark cases
 benchmarks := []struct {
  name  string
  value float64
  total float64
 }{
  {
   name:  "small values",
   value: 50.0,
   total: 200.0,
  },
  {
   name:  "large values",
   value: 5000000.0,
   total: 20000000.0,
  },
  {
   name:  "fractional values",
   value: 0.5,
   total: 2.0,
  },
 }

 for _, bm := range benchmarks {
  b.Run(bm.name, func(b *testing.B) {
   // Arrange
   // (setup is automatic excluded by b.Loop)

   // Act
   for b.Loop() {
    _, _ = Of(bm.value, bm.total)
   }
  })
 }
}
```

**Example 3: Benchmark with result validation**

```go
var (
 benchResult float64
 benchError  error
)

func BenchmarkChange(b *testing.B) {
 // Arrange
 oldValue := 100.0
 newValue := 150.0

 // Act
 for b.Loop() {
  benchResult, benchError = Change(oldValue, newValue)
 }
}
```

**Example 4: Comparative benchmark for different implementations**

```go
func BenchmarkToRatio(b *testing.B) {
 benchmarks := []struct {
  name    string
  percent float64
 }{
  {
   name:    "zero percent",
   percent: 0.0,
  },
  {
   name:    "fifty percent",
   percent: 50.0,
  },
  {
   name:    "hundred percent",
   percent: 100.0,
  },
 }

 for _, bm := range benchmarks {
  b.Run(bm.name, func(b *testing.B) {
   // Arrange
   // (automatic with b.Loop)

   // Act
   for b.Loop() {
    _, _ = ToRatio(bm.percent)
   }
  })
 }
}
```

## References

- **[SKILLS.md](./SKILLS.md)**: Project capabilities, commands, and tools documentation
- **[README.md](./README.md)**: User-facing documentation and contribution guidelines
- **[.github/copilot-instructions.md](./.github/copilot-instructions.md)**: GitHub Copilot instructions
- **[Makefile](./Makefile)**: Complete list of available commands
- **Go Testing**: [https://pkg.go.dev/testing](https://pkg.go.dev/testing)
- **Go Fuzzing**: [https://go.dev/security/fuzz/](https://go.dev/security/fuzz/)
- **Table-Driven Tests**: [https://go.dev/wiki/TableDrivenTests](https://go.dev/wiki/TableDrivenTests)
