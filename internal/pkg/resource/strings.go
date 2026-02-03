// SPDX-License-Identifier: Apache-2.0

package resource

// Error message constants used to create sentinel errors.
//
// These messages provide human-readable descriptions of error conditions
// encountered during percentage calculations.
const (
	// OutOfRangeErrorMessage is the message for values outside valid bounds.
	OutOfRangeErrorMessage = "pkg percent: out of the range"

	// DivideByZeroErrorMessage is the message for division by zero attempts.
	DivideByZeroErrorMessage = "pkg percent: division by zero"

	// PartGreaterThanTotalErrorMessage is the message for invalid part/total relationships.
	PartGreaterThanTotalErrorMessage = "pkg percent: part cannot be greater than total"
)
