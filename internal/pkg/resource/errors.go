// SPDX-License-Identifier: Apache-2.0

// Package resource provides internal constants, error definitions, and shared
// resources used across the percent package.
package resource

import (
	"errors"
)

// Sentinel errors returned by percentage calculation functions.
//
// These errors indicate specific failure conditions and can be checked using
// errors.Is() for proper error handling.
var (
	// ErrOutOfRange indicates that a value is outside the valid range.
	// For percentages, the valid range is 0 to 100 inclusive.
	// For ratios, the valid range is 0 to 1 inclusive.
	ErrOutOfRange = errors.New(OutOfRangeErrorMessage)

	// ErrDivideByZero indicates an attempt to divide by zero.
	// This occurs when a denominator value (such as total or oldValue) is zero.
	ErrDivideByZero = errors.New(DivideByZeroErrorMessage)

	// ErrPartGreaterThanTotal indicates that the part value exceeds the total.
	// This is an invalid condition for calculating proportional percentages.
	ErrPartGreaterThanTotal = errors.New(PartGreaterThanTotalErrorMessage)
)
