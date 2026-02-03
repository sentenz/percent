// SPDX-License-Identifier: Apache-2.0

// Package percent provides utility functions for calculating percentages and
// performing related mathematical operations.
//
// The package supports generic numeric types through Go generics, accepting
// both integer and floating-point values. All functions return float64 results
// for precision and consistency.
//
// Example usage:
//
//	// Calculate 25% of 200
//	result, err := percent.Percent(25, 200)
//	// result: 50.0
//
//	// Calculate what percentage 30 is of 120
//	result, err := percent.Of(30, 120)
//	// result: 25.0
//
// All functions validate their inputs and return appropriate errors for
// invalid operations such as division by zero or out-of-range values.
package percent

import (
	"math"

	"github.com/sentenz/percent/internal/pkg/resource"
	"golang.org/x/exp/constraints"
)

// Percent calculates the percentage of a given value.
//
// The calculation follows the formula:
//
//	result = value × (percent / 100)
//
// Where:
//   - value is the base amount
//   - percent is the percentage rate in range [0, 100]
//   - result is the calculated percentage of the value
//
// For example, 25% of 200:
//
//	result = 200 × (25 / 100) = 200 × 0.25 = 50
//
// The function returns the calculated result as float64.
// Returns 0 and resource.ErrOutOfRange if percent is outside the range [0, 100].
func Percent[T constraints.Integer | constraints.Float](percent, value T) (float64, error) {
	if float64(percent) < 0 || float64(percent) > 100 {
		return 0, resource.ErrOutOfRange
	}

	return float64(value) * (float64(percent) / resource.PercentMax), nil
}

// Of calculates the percentage that part represents of total.
//
// The calculation follows the formula:
//
//	             part
//	result = ─────────── × 100
//	            total
//
// Where:
//   - part is the portion being measured (part ≤ total)
//   - total is the whole amount (total ≠ 0)
//   - result is the percentage that part represents of total
//
// For example, what percentage is 30 of 120:
//
//	result = (30 / 120) × 100 = 0.25 × 100 = 25
//
// The function returns the calculated percentage as float64.
// Returns 0 and resource.ErrDivideByZero if total is zero.
// Returns 0 and resource.ErrPartGreaterThanTotal if part exceeds total.
func Of[T constraints.Integer | constraints.Float](part, total T) (float64, error) {
	if float64(total) == 0 {
		return 0, resource.ErrDivideByZero
	}

	if float64(part) > float64(total) {
		return 0, resource.ErrPartGreaterThanTotal
	}

	return float64(part) / float64(total) * resource.PercentMax, nil
}

// Change calculates the percentage change between an old value and a new value.
//
// The calculation follows the formula:
//
//	            (newValue - oldValue)
//	change = ─────────────────────────── × 100
//	                 |oldValue|
//
// Where:
//   - oldValue is the original value (oldValue ≠ 0)
//   - newValue is the new value
//   - change is the percentage change (positive for increase, negative for decrease)
//
// For example, change from 50 to 75:
//
//	change = (75 - 50) / |50| × 100 = 25 / 50 × 100 = 50%
//
// The function returns the percentage change as float64.
// Returns 0 and resource.ErrDivideByZero if oldValue is zero.
func Change[T constraints.Integer | constraints.Float](oldValue, newValue T) (float64, error) {
	if float64(oldValue) == 0 {
		return 0, resource.ErrDivideByZero
	}

	return (float64(newValue) - float64(oldValue)) / math.Abs(float64(oldValue)) * resource.PercentMax, nil
}

// Remain calculates the remaining value after subtracting a percentage.
//
// The calculation follows the formula:
//
//	                  (100 - percent)
//	remain = value × ─────────────────
//	                       100
//
// Where:
//   - value is the base amount
//   - percent is the percentage to subtract in range [0, 100]
//   - remain is the value remaining after percentage reduction
//
// For example, what remains after removing 25% from 200:
//
//	remain = 200 × ((100 - 25) / 100) = 200 × 0.75 = 150
//
// The function returns the remaining value as float64.
// Returns 0 and resource.ErrOutOfRange if percent is outside the range [0, 100].
func Remain[T constraints.Integer | constraints.Float](percent, value T) (float64, error) {
	if float64(percent) < 0 || float64(percent) > 100 {
		return 0, resource.ErrOutOfRange
	}

	return float64(value) * ((resource.PercentMax - float64(percent)) / resource.PercentMax), nil
}

// FromRatio converts a decimal ratio to a percentage.
//
// The calculation follows the formula:
//
//	percent = ratio × 100
//
// Where:
//   - ratio is the decimal value in range [0, 1]
//   - percent is the equivalent percentage value
//
// For example, convert ratio 0.25 to percentage:
//
//	percent = 0.25 × 100 = 25
//
// The function returns the percentage as float64.
// Returns 0 and resource.ErrOutOfRange if ratio is outside the range [0, 1].
func FromRatio[T constraints.Integer | constraints.Float](ratio T) (float64, error) {
	if float64(ratio) < 0 || float64(ratio) > 1 {
		return 0, resource.ErrOutOfRange
	}

	return float64(ratio) * resource.PercentMax, nil
}

// ToRatio converts a percentage to a decimal ratio.
//
// The calculation follows the formula:
//
//	          percent
//	ratio = ───────────
//	            100
//
// Where:
//   - percent is the percentage value in range [0, 100]
//   - ratio is the equivalent decimal value in range [0, 1]
//
// For example, convert 25% to ratio:
//
//	ratio = 25 / 100 = 0.25
//
// The function returns the ratio as float64.
// Returns 0 and resource.ErrOutOfRange if percent is outside the range [0, 100].
func ToRatio[T constraints.Integer | constraints.Float](percent T) (float64, error) {
	if float64(percent) < 0 || float64(percent) > 100 {
		return 0, resource.ErrOutOfRange
	}

	return float64(percent) / resource.PercentMax, nil
}
