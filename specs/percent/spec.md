# Feature Specification: Percent Package

**Feature Branch**: `pkg/percent`

**Created**: 2025-01-01

**Status**: Stable

## Overview

The `percent` package provides utility functions for calculating percentages and performing related mathematical operations.
It supports generic numeric types through Go generics and returns `float64` results for precision.

---

## User Stories & Testing

### User Story 1 — Calculate a percentage of a value (Priority: P1)

A developer wants to know how much a certain percentage of a base amount is.

**Why this priority**: This is the most fundamental percentage operation and the primary entry point for users of the library.

**Independent Test**: Call `percent.Percent(25, 200.0)` and verify the result is `50.0`.

**Acceptance Scenarios**:

1. **Given** a valid percent `p` in `[0, 100]` and a value `v`, **When** `Percent(p, v)` is called, **Then** it returns `v × (p / 100)` and `nil` error.
2. **Given** `percent = 0`, **When** `Percent(0, v)` is called, **Then** it returns `0` and `nil` error.
3. **Given** `percent = 100`, **When** `Percent(100, v)` is called, **Then** it returns `v` and `nil` error.
4. **Given** `percent < 0` or `percent > 100`, **When** `Percent(p, v)` is called, **Then** it returns `0` and `ErrOutOfRange`.

---

### User Story 2 — Calculate what percentage a part is of a total (Priority: P1)

A developer wants to express a part as a percentage of a whole.

**Why this priority**: Equally fundamental to US1; both are required to make the library minimally useful.

**Independent Test**: Call `percent.Of(50.0, 200.0)` and verify the result is `25.0`.

**Acceptance Scenarios**:

1. **Given** a `part` ≤ `total` and `total ≠ 0`, **When** `Of(part, total)` is called, **Then** it returns `(part / total) × 100` and `nil` error.
2. **Given** `part = 0`, **When** `Of(0, total)` is called, **Then** it returns `0` and `nil` error.
3. **Given** `part = total`, **When** `Of(total, total)` is called, **Then** it returns `100` and `nil` error.
4. **Given** `total = 0`, **When** `Of(part, 0)` is called, **Then** it returns `0` and `ErrDivideByZero`.
5. **Given** `part > total`, **When** `Of(part, total)` is called, **Then** it returns `0` and `ErrPartGreaterThanTotal`.

---

### User Story 3 — Calculate percentage change between two values (Priority: P2)

A developer wants to quantify how much a value has grown or shrunk relative to its original amount.

**Why this priority**: Derived metric building on US1/US2; important but not required for an MVP.

**Independent Test**: Call `percent.Change(50.0, 75.0)` and verify the result is `50.0`.

**Acceptance Scenarios**:

1. **Given** `oldValue ≠ 0`, **When** `Change(oldValue, newValue)` is called, **Then** it returns `((newValue - oldValue) / |oldValue|) × 100` and `nil` error.
2. **Given** `oldValue = newValue`, **When** `Change(v, v)` is called, **Then** it returns `0` and `nil` error.
3. **Given** `oldValue = 0`, **When** `Change(0, newValue)` is called, **Then** it returns `0` and `ErrDivideByZero`.
4. **Given** `oldValue < 0` and `newValue < 0` (decrease), **When** `Change(oldValue, newValue)` is called, **Then** it returns a negative percentage and `nil` error.

---

### User Story 4 — Calculate the remaining value after a percentage reduction (Priority: P2)

A developer wants to know what remains after removing a given percentage from a base value.

**Why this priority**: Useful complement to US1 for discount/tax scenarios; not required for MVP.

**Independent Test**: Call `percent.Remain(25, 200.0)` and verify the result is `150.0`.

**Acceptance Scenarios**:

1. **Given** a valid `percent` in `[0, 100]` and `value`, **When** `Remain(percent, value)` is called, **Then** it returns `value × ((100 - percent) / 100)` and `nil` error.
2. **Given** `percent = 0`, **When** `Remain(0, value)` is called, **Then** it returns `value` unchanged and `nil` error.
3. **Given** `percent = 100`, **When** `Remain(100, value)` is called, **Then** it returns `0` and `nil` error.
4. **Given** `percent < 0` or `percent > 100`, **When** `Remain(percent, value)` is called, **Then** it returns `0` and `ErrOutOfRange`.

---

### User Story 5 — Convert a ratio to a percentage (Priority: P3)

A developer wants to convert a decimal ratio (e.g., `0.75`) to a percentage (`75`).

**Independent Test**: Call `percent.FromRatio(0.75)` and verify the result is `75.0`.

**Acceptance Scenarios**:

1. **Given** a `ratio` in `[0, 1]`, **When** `FromRatio(ratio)` is called, **Then** it returns `ratio × 100` and `nil` error.
2. **Given** `ratio = 0`, **When** `FromRatio(0)` is called, **Then** it returns `0` and `nil` error.
3. **Given** `ratio = 1`, **When** `FromRatio(1)` is called, **Then** it returns `100` and `nil` error.
4. **Given** `ratio < 0` or `ratio > 1`, **When** `FromRatio(ratio)` is called, **Then** it returns `0` and `ErrOutOfRange`.

---

### User Story 6 — Convert a percentage to a ratio (Priority: P3)

A developer wants to convert a percentage (e.g., `75`) to a decimal ratio (`0.75`).

**Independent Test**: Call `percent.ToRatio(75.0)` and verify the result is `0.75`.

**Acceptance Scenarios**:

1. **Given** a `percent` in `[0, 100]`, **When** `ToRatio(percent)` is called, **Then** it returns `percent / 100` and `nil` error.
2. **Given** `percent = 0`, **When** `ToRatio(0)` is called, **Then** it returns `0` and `nil` error.
3. **Given** `percent = 100`, **When** `ToRatio(100)` is called, **Then** it returns `1` and `nil` error.
4. **Given** `percent < 0` or `percent > 100`, **When** `ToRatio(percent)` is called, **Then** it returns `0` and `ErrOutOfRange`.

---

### Edge Cases

- What happens when both `part` and `total` are negative in `Of()`?
  > `Of(-200, -50)` returns `400.0`, `nil` — the ratio of two negative values is a valid positive percentage.
- What happens with floating-point inputs that are very close to boundary values (e.g., `100.0000000001`)?
  > The function returns `ErrOutOfRange`; callers must ensure values are within `[0, 100]`.
- What happens when `value = 0` in `Percent()` or `Remain()`?
  > Returns `0` and `nil` error — zero value is valid and the result is always zero.
- What happens with integer types as input (e.g., `int`, `int64`)?
  > All functions accept any `constraints.Integer | constraints.Float` type via generics; integer inputs are promoted to `float64` internally.

---

## Golden Dataset

The following input/output pairs are the authoritative test cases for this package.
Any implementation must produce these exact outputs for these inputs.

### `Percent(percent, value)`

| percent | value  | result  | error          |
|---------|--------|---------|----------------|
| 25      | 100.0  | 25.0    | nil            |
| 50      | -200.0 | -100.0  | nil            |
| 0       | 100.0  | 0.0     | nil            |
| 100     | 50.0   | 50.0    | nil            |
| -10     | 100.0  | 0.0     | ErrOutOfRange  |
| 150     | 100.0  | 0.0     | ErrOutOfRange  |

### `Of(part, total)`

| part    | total  | result  | error                    |
|---------|--------|---------|--------------------------|
| 25.0    | 100.0  | 25.0    | nil                      |
| -200.0  | -50.0  | 400.0   | nil                      |
| 100.0   | 100.0  | 100.0   | nil                      |
| 0.0     | 100.0  | 0.0     | nil                      |
| 150.0   | 100.0  | 0.0     | ErrPartGreaterThanTotal  |
| 150.0   | 0.0    | 0.0     | ErrDivideByZero          |

### `Change(oldValue, newValue)`

| oldValue | newValue | result  | error           |
|----------|----------|---------|-----------------|
| 25.0     | 100.0    | 300.0   | nil             |
| -50.0    | -200.0   | -300.0  | nil             |
| 100.0    | 100.0    | 0.0     | nil             |
| 0.0      | 100.0    | 0.0     | ErrDivideByZero |

### `Remain(percent, value)`

| percent | value  | result  | error          |
|---------|--------|---------|----------------|
| 25.0    | 100.0  | 75.0    | nil            |
| 50.0    | -200.0 | -100.0  | nil            |
| 0.0     | 100.0  | 100.0   | nil            |
| 100.0   | 50.0   | 0.0     | nil            |
| -10.0   | 100.0  | 0.0     | ErrOutOfRange  |
| 150.0   | 100.0  | 0.0     | ErrOutOfRange  |

### `FromRatio(ratio)`

| ratio  | result | error          |
|--------|--------|----------------|
| 0.25   | 25.0   | nil            |
| 0.0    | 0.0    | nil            |
| 1.0    | 100.0  | nil            |
| -0.1   | 0.0    | ErrOutOfRange  |
| 2.0    | 0.0    | ErrOutOfRange  |

### `ToRatio(percent)`

| percent | result | error          |
|---------|--------|----------------|
| 50.0    | 0.5    | nil            |
| 0.0     | 0.0    | nil            |
| 100.0   | 1.0    | nil            |
| -10.0   | 0.0    | ErrOutOfRange  |
| 150.0   | 0.0    | ErrOutOfRange  |

---

## Requirements

### Functional Requirements

- **FR-001**: The library MUST provide a `Percent` function to calculate `value × (percent / 100)`.
- **FR-002**: The library MUST provide an `Of` function to calculate `(part / total) × 100`.
- **FR-003**: The library MUST provide a `Change` function to calculate `((newValue - oldValue) / |oldValue|) × 100`.
- **FR-004**: The library MUST provide a `Remain` function to calculate `value × ((100 - percent) / 100)`.
- **FR-005**: The library MUST provide a `FromRatio` function to convert a `[0, 1]` ratio to a percentage.
- **FR-006**: The library MUST provide a `ToRatio` function to convert a `[0, 100]` percentage to a ratio.
- **FR-007**: All functions MUST accept any `constraints.Integer | constraints.Float` type via Go generics.
- **FR-008**: All functions MUST return `(float64, error)`.
- **FR-009**: `Percent`, `Remain`, and `ToRatio` MUST return `ErrOutOfRange` when `percent` is outside `[0, 100]`.
- **FR-010**: `FromRatio` MUST return `ErrOutOfRange` when `ratio` is outside `[0, 1]`.
- **FR-011**: `Of` MUST return `ErrDivideByZero` when `total = 0`.
- **FR-012**: `Of` MUST return `ErrPartGreaterThanTotal` when `part > total`.
- **FR-013**: `Change` MUST return `ErrDivideByZero` when `oldValue = 0`.

### Key Entities

- **`percent` (input)**: A numeric value representing a percentage in the range `[0, 100]`.
- **`ratio` (input)**: A decimal value representing a proportion in the range `[0, 1]`.
- **`part` / `total`**: Numerator and denominator for proportional calculation; `part ≤ total` when both are positive.
- **`oldValue` / `newValue`**: Before and after values for change calculation; `oldValue ≠ 0`.
- **Sentinel errors**: `ErrOutOfRange`, `ErrDivideByZero`, `ErrPartGreaterThanTotal` — defined in `internal/pkg/resource/`.

---

## Success Criteria

- **SC-001**: All six functions exist and are exported from `pkg/percent`.
- **SC-002**: Every function in the golden dataset produces the exact listed result and error.
- **SC-003**: Statement coverage is 100 % (CI badge remains green).
- **SC-004**: The package compiles and all tests pass with `-race` flag enabled.
- **SC-005**: All exported symbols are documented with godoc-compatible comments.
- **SC-006**: `go vet` and `golangci-lint` report zero issues.

---

## Assumptions

- Callers provide numeric inputs without pre-validation; the library owns all range checks.
- Negative `value` inputs for `Percent` and `Remain` are valid and return negative results.
- Negative `part` and `total` in `Of` are valid when their ratio produces a meaningful percentage.
- No floating-point rounding or precision normalization beyond standard `float64` arithmetic is required.
- Thread safety is provided by Go's immutable function semantics; no shared mutable state exists.
