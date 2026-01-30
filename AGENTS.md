# AGENTS.md

- [1. Tech Stack](#1-tech-stack)
  - [1.1. Programming Language](#11-programming-language)
  - [1.2. Dependency Manager](#12-dependency-manager)
  - [1.3. Test Framework](#13-test-framework)
  - [1.4. Build System](#14-build-system)
- [2. Project Layout](#2-project-layout)
  - [2.1. Directory Structure](#21-directory-structure)
  - [2.2. Key Directories](#22-key-directories)
- [3. Task Runners](#3-task-runners)
  - [3.1. Make Commands](#31-make-commands)
    - [3.1.1. Setup & Teardown](#311-setup--teardown)
    - [3.1.2. Go Module Management](#312-go-module-management)
    - [3.1.3. Go Testing](#313-go-testing)
    - [3.1.4. Go Code Quality](#314-go-code-quality)
    - [3.1.5. Policy & Security](#315-policy--security)
- [4. Software Testing](#4-software-testing)
  - [4.1. Unit Testing](#41-unit-testing)
    - [4.1.1. Unit Testing Patterns](#411-unit-testing-patterns)
    - [4.1.2. Unit Test Workflow](#412-unit-test-workflow)
    - [4.1.3. Unit Test Commands](#413-unit-test-commands)
    - [4.1.4. Unit Test Style](#414-unit-test-style)
    - [4.1.5. Unit Test Template](#415-unit-test-template)
  - [4.2. Fuzz Testing](#42-fuzz-testing)
    - [4.2.1. Fuzz Testing Patterns](#421-fuzz-testing-patterns)
    - [4.2.2. Fuzz Test Workflow](#422-fuzz-test-workflow)
    - [4.2.3. Fuzz Test Commands](#423-fuzz-test-commands)
    - [4.2.4. Fuzz Test Style](#424-fuzz-test-style)
    - [4.2.5. Fuzz Test Template](#425-fuzz-test-template)
  - [4.3. Benchmark Testing](#43-benchmark-testing)
    - [4.3.1. Benchmark Testing Patterns](#431-benchmark-testing-patterns)
    - [4.3.2. Benchmark Test Workflow](#432-benchmark-test-workflow)
    - [4.3.3. Benchmark Test Commands](#433-benchmark-test-commands)
    - [4.3.4. Benchmark Test Style](#434-benchmark-test-style)
    - [4.3.5. Benchmark Test Template](#435-benchmark-test-template)

## 1. Tech Stack

The technology stack for this Go project includes the tools and frameworks used for development, testing, and deployment.

### 1.1. Programming Language

- Go (Golang)
  > A statically typed, compiled programming language designed for simplicity, efficiency, and concurrency.
  >
  > Version: Go 1.22.0 or higher

### 1.2. Dependency Manager

- [Go Modules](https://go.dev/ref/mod)
  > Go's official dependency management system that handles versioning and package dependencies through `go.mod` and `go.sum` files.
  >
  > Commands:
  > - `go mod tidy` - Add missing and remove unused modules
  > - `go mod vendor` - Make vendored copy of dependencies
  > - `go mod download` - Download modules to local cache

### 1.3. Test Framework

- [Testing](https://pkg.go.dev/testing)
  > Go's built-in testing framework for unit tests, benchmarks, and fuzz tests.
  >
  > Supports:
  > - Unit tests with table-driven patterns
  > - Benchmark tests for performance measurement
  > - Fuzz tests with coverage-guided fuzzing
  > - Code coverage analysis

### 1.4. Build System

- [Make](https://www.gnu.org/software/make/)
  > Task runner and build automation tool using Makefile to define project tasks and commands.
  >
  > Used for: build automation, testing, linting, and deployment tasks

## 2. Project Layout

Standard Go project layout following community best practices for organizing code and resources.

### 2.1. Directory Structure

```
.
├── .github/            # GitHub workflows, actions, and configuration
│   ├── skills/         # Agent skills for AI coding assistants
│   └── workflows/      # GitHub Actions CI/CD workflows
├── internal/           # Private application and library code
│   └── pkg/            # Internal packages
│       └── resource/   # Internal resources (errors, constants, strings)
├── pkg/                # Public library code (importable by external projects)
│   └── percent/        # Main package with percentage utilities
├── scripts/            # Scripts for build, setup, and deployment
├── tests/              # Additional test files and test data
│   └── policy/         # Policy tests with Conftest/Rego
├── vendor/             # Vendored dependencies (managed by Go modules)
├── go.mod              # Go module definition
├── go.sum              # Go module checksums
├── Makefile            # Task runner definitions
└── AGENTS.md           # AI agent instructions and guidelines
```

### 2.2. Key Directories

- `pkg/`
  > Contains packages that can be imported by external projects. This is the public API of the module.

- `internal/`
  > Contains private code that cannot be imported by external projects. Go enforces this restriction.

- `internal/pkg/resource/`
  > Shared internal resources like error definitions, constants, and string templates used across packages.

- `.github/skills/`
  > Agent skills documentation following the Agent Skills specification for AI coding assistants.

- `tests/policy/`
  > Policy-as-code tests using Conftest and Rego for validating configurations.

- `scripts/`
  > Automation scripts for bootstrapping, setup, and teardown of the development environment.

- `vendor/`
  > Local copy of dependencies for reproducible builds and offline development.

## 3. Task Runners

Make is used as the task runner to automate common development tasks.

### 3.1. Make Commands

Run `make help` to see all available commands with descriptions.

#### 3.1.1. Setup & Teardown

- `make bootstrap`
  > Initialize a software development workspace with requisites.
  >
  > Runs: `scripts/bootstrap.sh`

- `make setup`
  > Install and configure all dependencies essential for development.
  >
  > Runs: `scripts/setup.sh`

- `make teardown`
  > Remove development artifacts and restore the host to its pre-setup state.
  >
  > Runs: `scripts/teardown.sh`

#### 3.1.2. Go Module Management

- `make go-mod-tidy`
  > Tidy Go modules by adding missing and removing unused modules.
  >
  > Runs: `go mod tidy`

- `make go-mod-vendor`
  > Vendor Go modules to create a local copy of dependencies.
  >
  > Runs: `go mod vendor`

#### 3.1.3. Go Testing

- `make go-test-unit`
  > Run Go unit tests with race detection and generate JUnit XML report.
  >
  > Flags: `-race`, `-coverprofile`, `-v`

- `make go-test-coverage`
  > Run Go tests with coverage report (HTML and XML formats).
  >
  > Generates: `coverage.html`, `coverage.xml`

- `make go-test-bench`
  > Run Go benchmarks to measure performance and memory usage.
  >
  > Flags: `-bench=.`, `-benchmem`

- `make go-test-fuzz`
  > Run fuzz tests with Go's native fuzzing engine.
  >
  > Duration: Configurable via environment variable

#### 3.1.4. Go Code Quality

- `make go-fmt`
  > Format Go code according to Go standards using `gofmt`.
  >
  > Runs: `gofmt -s -w`

- `make go-fix`
  > Update Go source files to use new APIs with `go fix`.

- `make go-vet`
  > Check Go code for common mistakes using `go vet`.

- `make go-vuln`
  > Check Go code for known vulnerabilities using `govulncheck`.

- `make go-check`
  > Run all Go code quality checks (fmt, vet, vuln).
  >
  > Runs: `go-fmt`, `go-vet`, `go-vuln`

#### 3.1.5. Policy & Security

- `make conftest-test`
  > Analyze configuration files using Conftest for policy violations.

- `make regal-lint`
  > Lint Rego policies using Regal.

- `make trivy-scan`
  > Scan Infrastructure-as-Code files for misconfigurations using Trivy.

## 4. Software Testing

### 4.1. Unit Testing

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

#### 1.1.1. Unit Testing Patterns

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

#### 1.1.2. Unit Test Workflow

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

#### 1.1.3. Unit Test Commands

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

#### 1.1.4. Unit Test Style

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

#### 1.1.5. Unit Test Template

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

### 1.2. Fuzz Testing

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

#### 4.2.1. Fuzz Testing Patterns

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

#### 4.2.2. Fuzz Test Workflow

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

    Structure all fuzz tests using this [template](#425-fuzz-test-template) pattern.

5. Seed Corpus

    Optionally provide seed inputs in `testdata/fuzz/<FuzzTestName>/` directory to guide fuzzing toward interesting inputs.

#### 4.2.3. Fuzz Test Commands

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

#### 4.2.4. Fuzz Test Style

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

#### 4.2.5. Fuzz Test Template

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

### 4.3. Benchmark Testing

Instructions for AI coding agents on automating benchmark test creation using consistent software testing patterns in this Go project.

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

#### 4.3.1. Benchmark Testing Patterns

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

#### 4.3.2. Benchmark Test Workflow

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

    Structure all benchmark tests using this [template](#435-benchmark-test-template) pattern.

5. Baseline Measurements

    Establish performance baselines by running benchmarks on stable code before making changes.

#### 4.3.3. Benchmark Test Commands

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

#### 4.3.4. Benchmark Test Style

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

#### 4.3.5. Benchmark Test Template

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
