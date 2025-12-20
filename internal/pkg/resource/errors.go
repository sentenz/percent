// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"errors"
)

var (
	ErrOutOfRange           = errors.New(OutOfRangeErrorMessage)
	ErrDivideByZero         = errors.New(DivideByZeroErrorMessage)
	ErrPartGreaterThanTotal = errors.New(PartGreaterThanTotalErrorMessage)
)
