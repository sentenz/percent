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
// The function computes the result as: value * (percent / 100).
//
// The percent parameter must be between 0 and 100 inclusive. Values outside
// this range return an error.
//
// Returns the calculated percentage as a float64. If percent is out of range,
// it returns 0 and resource.ErrOutOfRange.
//
// Example:
//
//	// Calculate 25% of 200
//	result, err := percent.Percent(25, 200)
//	// result: 50.0, err: nil
func Percent[T constraints.Integer | constraints.Float](percent, value T) (float64, error) {
	if float64(percent) < 0 || float64(percent) > 100 {
		return 0, resource.ErrOutOfRange
	}

	return float64(value) * (float64(percent) / resource.PercentMax), nil
}

// Of calculates the percentage that part represents of total.
//
// The function computes the result as: (part / total) * 100.
//
// The total parameter must be non-zero. The part parameter must not exceed
// total. Violations of these constraints return an error.
//
// Returns the calculated percentage as a float64. If total is zero, it returns
// 0 and resource.ErrDivideByZero. If part exceeds total, it returns 0 and
// resource.ErrPartGreaterThanTotal.
//
// Example:
//
//	// Calculate what percentage 30 is of 120
//	result, err := percent.Of(30, 120)
//	// result: 25.0, err: nil
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
// The function computes the result as: ((newValue - oldValue) / |oldValue|) * 100.
//
// The absolute value of oldValue is used in the denominator to correctly
// handle negative base values. The oldValue parameter must be non-zero.
//
// A positive result indicates an increase, while a negative result indicates
// a decrease.
//
// Returns the percentage change as a float64. If oldValue is zero, it returns
// 0 and resource.ErrDivideByZero.
//
// Example:
//
//	// Calculate percentage change from 50 to 75
//	result, err := percent.Change(50, 75)
//	// result: 50.0, err: nil (50% increase)
func Change[T constraints.Integer | constraints.Float](oldValue, newValue T) (float64, error) {
	if float64(oldValue) == 0 {
		return 0, resource.ErrDivideByZero
	}

	return (float64(newValue) - float64(oldValue)) / math.Abs(float64(oldValue)) * resource.PercentMax, nil
}

// Remain calculates the remaining value after subtracting a percentage.
//
// The function computes the result as: value * ((100 - percent) / 100).
//
// The percent parameter must be between 0 and 100 inclusive. Values outside
// this range return an error.
//
// Returns the remaining value as a float64. If percent is out of range,
// it returns 0 and resource.ErrOutOfRange.
//
// Example:
//
//	// Calculate what remains after removing 25% from 200
//	result, err := percent.Remain(25, 200)
//	// result: 150.0, err: nil
func Remain[T constraints.Integer | constraints.Float](percent, value T) (float64, error) {
	if float64(percent) < 0 || float64(percent) > 100 {
		return 0, resource.ErrOutOfRange
	}

	return float64(value) * ((resource.PercentMax - float64(percent)) / resource.PercentMax), nil
}

// FromRatio converts a decimal ratio to a percentage.
//
// The function computes the result as: ratio * 100.
//
// The ratio parameter must be between 0 and 1 inclusive. Values outside
// this range return an error.
//
// Returns the percentage as a float64. If ratio is out of range,
// it returns 0 and resource.ErrOutOfRange.
//
// Example:
//
//	// Convert ratio 0.25 to percentage
//	result, err := percent.FromRatio(0.25)
//	// result: 25.0, err: nil
func FromRatio[T constraints.Integer | constraints.Float](ratio T) (float64, error) {
	if float64(ratio) < 0 || float64(ratio) > 1 {
		return 0, resource.ErrOutOfRange
	}

	return float64(ratio) * resource.PercentMax, nil
}

// ToRatio converts a percentage to a decimal ratio.
//
// The function computes the result as: percent / 100.
//
// The percent parameter must be between 0 and 100 inclusive. Values outside
// this range return an error.
//
// Returns the ratio as a float64. If percent is out of range,
// it returns 0 and resource.ErrOutOfRange.
//
// Example:
//
//	// Convert 25% to ratio
//	result, err := percent.ToRatio(25)
//	// result: 0.25, err: nil
func ToRatio[T constraints.Integer | constraints.Float](percent T) (float64, error) {
	if float64(percent) < 0 || float64(percent) > 100 {
		return 0, resource.ErrOutOfRange
	}

	return float64(percent) / resource.PercentMax, nil
}
