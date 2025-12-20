// SPDX-License-Identifier: Apache-2.0

package percent

import (
	"math"

	"github.com/sentenz/percent/internal/pkg/resource"
	"golang.org/x/exp/constraints"
)

// Percent returns the percentage of value.
func Percent[T constraints.Integer | constraints.Float](percent, value T) (float64, error) {
	if float64(percent) < 0 || float64(percent) > 100 {
		return 0, resource.ErrOutOfRange
	}

	return float64(value) * (float64(percent) / resource.PercentMax), nil
}

// Of calculates the percentage of the part relative to the total.
func Of[T constraints.Integer | constraints.Float](part, total T) (float64, error) {
	if float64(total) == 0 {
		return 0, resource.ErrDivideByZero
	}

	if float64(part) > float64(total) {
		return 0, resource.ErrPartGreaterThanTotal
	}

	return float64(part) / float64(total) * resource.PercentMax, nil
}

// Change calculates the percentage change between two values.
func Change[T constraints.Integer | constraints.Float](oldValue, newValue T) (float64, error) {
	if float64(oldValue) == 0 {
		return 0, resource.ErrDivideByZero
	}

	return (float64(newValue) - float64(oldValue)) / math.Abs(float64(oldValue)) * resource.PercentMax, nil
}

// Remain returns the percentage of value that remains after subtracting the percentage.
func Remain[T constraints.Integer | constraints.Float](percent, value T) (float64, error) {
	if float64(percent) < 0 || float64(percent) > 100 {
		return 0, resource.ErrOutOfRange
	}

	return float64(value) * ((resource.PercentMax - float64(percent)) / resource.PercentMax), nil
}

// FromRatio returns the percent of ratio.
func FromRatio[T constraints.Integer | constraints.Float](ratio T) (float64, error) {
	if float64(ratio) < 0 || float64(ratio) > 1 {
		return 0, resource.ErrOutOfRange
	}

	return float64(ratio) * resource.PercentMax, nil
}

// ToRatio returns the ratio of percent.
func ToRatio[T constraints.Integer | constraints.Float](percent T) (float64, error) {
	if float64(percent) < 0 || float64(percent) > 100 {
		return 0, resource.ErrOutOfRange
	}

	return float64(percent) / resource.PercentMax, nil
}
