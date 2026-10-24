# Implementation Plan: Percent Package

**Branch**: `pkg/percent` | **Date**: 2025-01-01 | **Spec**: [spec.md](spec.md)

## Summary

The `percent` package is a pure Go library that provides six percentage-calculation functions over generic numeric types.
All functions validate inputs against well-defined ranges and return typed sentinel errors for invalid states.
The implementation resides entirely in `pkg/percent/percent.go` with corresponding tests in `pkg/percent/percent_test.go`.

## Technical Context

**Language/Version**: Go 1.22+

**Primary Dependencies**:
- `golang.org/x/exp/constraints` — generic type constraints (`Integer | Float`)
- `github.com/google/go-cmp` — test assertion library (dev dependency)

**Storage**: N/A — pure calculation library, no persistence.

**Testing**: Go standard `testing` package; table-driven tests; `make go-test-unit`, `make go-test-coverage`.

**Target Platform**: Any platform supported by the Go toolchain (linux/amd64, darwin/arm64, windows/amd64, …).

**Project Type**: Library (importable Go module).

**Performance Goals**: All functions are O(1) arithmetic; no performance budgets required.

**Constraints**: Zero external runtime dependencies beyond `golang.org/x/exp`; no CGO.

**Scale/Scope**: Single package, six exported functions, ≤200 LOC.

## Constitution Check

| Principle | Status | Notes |
|-----------|--------|-------|
| I. Library-First | ✅ | All code lives under `pkg/percent/` |
| II. Generic Numeric API | ✅ | `[T constraints.Integer \| constraints.Float]` on all functions |
| III. Explicit Error Handling | ✅ | Returns `(float64, error)`; sentinel errors in `resource` |
| IV. Test-First | ✅ | 100% coverage gate enforced in CI |
| V. Input Validation at Boundary | ✅ | Each function validates its own inputs using `resource` constants |
| VI. Immutable Internal Resources | ✅ | Constants and errors defined in `internal/pkg/resource/` |
| VII. Semantic Versioning | ✅ | Module path is `github.com/sentenz/percent/v3` |
| VIII. Simplicity | ✅ | Flat package, no abstraction layers |

## Project Structure

```text
github.com/sentenz/percent/v3
├── internal/
│   └── pkg/
│       └── resource/
│           ├── constant.go      # PercentMin, PercentMax
│           ├── error.go         # ErrOutOfRange, ErrDivideByZero, ErrPartGreaterThanTotal
│           └── message.go       # Error message string constants
├── pkg/
│   └── percent/
│       ├── percent.go           # Exported functions: Percent, Of, Change, Remain, FromRatio, ToRatio
│       └── percent_test.go      # Table-driven unit tests (100% coverage)
├── specs/
│   ├── constitution.md          # Project governing principles
│   └── percent/
│       ├── spec.md              # Feature specification and golden dataset (this file's companion)
│       └── plan.md              # This file
├── go.mod
├── go.sum
├── vendor/                      # Vendored dependencies
└── Makefile                     # Task runner
```

## Function Designs

### `Percent[T](percent, value T) (float64, error)`

```
formula:  result = float64(value) × (float64(percent) / 100)
validate: percent ∈ [0, 100] → else ErrOutOfRange
```

### `Of[T](part, total T) (float64, error)`

```
formula:  result = float64(part) / float64(total) × 100
validate: total ≠ 0 → else ErrDivideByZero
          float64(part) ≤ float64(total) → else ErrPartGreaterThanTotal
```

### `Change[T](oldValue, newValue T) (float64, error)`

```
formula:  result = (float64(newValue) - float64(oldValue)) / math.Abs(float64(oldValue)) × 100
validate: oldValue ≠ 0 → else ErrDivideByZero
```

### `Remain[T](percent, value T) (float64, error)`

```
formula:  result = float64(value) × ((100 - float64(percent)) / 100)
validate: percent ∈ [0, 100] → else ErrOutOfRange
```

### `FromRatio[T](ratio T) (float64, error)`

```
formula:  result = float64(ratio) × 100
validate: ratio ∈ [0, 1] → else ErrOutOfRange
```

### `ToRatio[T](percent T) (float64, error)`

```
formula:  result = float64(percent) / 100
validate: percent ∈ [0, 100] → else ErrOutOfRange
```

## Internal Resources

### `internal/pkg/resource/constant.go`

```go
const (
    PercentMin = 0.0
    PercentMax = 100.0
)
```

### `internal/pkg/resource/message.go`

```go
const (
    OutOfRangeErrorMessage           = "pkg percent: out of the range"
    DivideByZeroErrorMessage         = "pkg percent: division by zero"
    PartGreaterThanTotalErrorMessage  = "pkg percent: part cannot be greater than total"
)
```

### `internal/pkg/resource/error.go`

```go
var (
    ErrOutOfRange            = errors.New(OutOfRangeErrorMessage)
    ErrDivideByZero          = errors.New(DivideByZeroErrorMessage)
    ErrPartGreaterThanTotal  = errors.New(PartGreaterThanTotalErrorMessage)
)
```

## Testing Strategy

All tests are table-driven and co-located with the implementation (`percent_test.go`).
Each test function covers one exported function and uses the In-Got-Want pattern.
The golden dataset from `spec.md` is the minimum set of required test cases.
Additional boundary and edge cases are included to maintain 100% coverage.

### Test Commands

| Command                 | Purpose                                        |
| ----------------------- | ---------------------------------------------- |
| `make go-test-unit`     | Run tests with race detection; generate JUnit  |
| `make go-test-coverage` | Generate HTML and XML coverage reports         |
| `make go-test-bench`    | Run benchmarks to validate O(1) performance    |

## Release Checklist

- [ ] All golden dataset cases pass (`make go-test-unit`)
- [ ] Coverage is 100% (`make go-test-coverage`)
- [ ] `go vet` reports zero issues (`make go-vet`)
- [ ] `golangci-lint` reports zero issues (`make go-check`)
- [ ] All exported symbols have godoc comments
- [ ] `go mod tidy` and `go mod vendor` are up to date
- [ ] CHANGELOG.md entry is added
- [ ] Version tag follows semantic versioning
