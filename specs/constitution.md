# Percent Constitution

## I. Library-First

Every feature is delivered as a standalone, importable Go library under `pkg/`.
Libraries must be self-contained, independently testable, and fully documented using godoc conventions.
A clear purpose is required — no organizational-only packages.

## II. Generic Numeric API

All calculation functions use Go generics (`constraints.Integer | constraints.Float`) to accept any numeric type.
Return types are always `float64` for precision and consistency across integer and floating-point callers.
The public API signature is: `func Name[T constraints.Integer | constraints.Float](inputs T) (float64, error)`.

## III. Explicit Error Handling (NON-NEGOTIABLE)

Every function that can fail returns an `(float64, error)` pair — never panics, never silently swallows errors.
Sentinel errors are defined in `internal/pkg/resource/` and checked with `errors.Is()`.
No new error type may be introduced without a corresponding constant message in `resource`.

## IV. Test-First (NON-NEGOTIABLE)

TDD is mandatory: tests are written and reviewed before implementation.
The Red-Green-Refactor cycle is strictly enforced.
100 % statement coverage is a gate for merging; the CI badge in README must stay green.

## V. Input Validation at the Boundary

All input validation occurs at the entry point of each exported function.
Valid ranges are enforced by constants from `internal/pkg/resource/` (`PercentMin`, `PercentMax`).
No caller should need to pre-validate inputs; the library owns this responsibility.

## VI. Immutable Internal Resources

Constants and sentinel errors in `internal/pkg/resource/` are the single source of truth.
Changing a constant value constitutes a breaking change and requires a major-version bump.
String-based error messages must remain stable across patch and minor releases.

## VII. Semantic Versioning & Module Path

This module follows Go's semantic import versioning (`github.com/sentenz/percent/vN`).
Breaking API changes require a new major version with an updated module path.
Non-breaking additions or fixes increment MINOR or PATCH versions respectively.

## VIII. Simplicity over Cleverness

Avoid abstraction layers that add indirection without clear benefit.
Prefer readable, flat code over clever one-liners.
YAGNI: do not add functionality that is not demanded by a concrete use case.

## Governance

This constitution supersedes all other practices documented in the repository.
Amendments require a pull request, peer review, a documented rationale, and an update to the `Last Amended` date.
All pull requests and code reviews must verify compliance with these principles.
Complexity violations must be explicitly justified in the PR description.

**Version**: 1.0.0 | **Ratified**: 2025-01-01 | **Last Amended**: 2025-01-01
