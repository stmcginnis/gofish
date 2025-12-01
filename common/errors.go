//
// SPDX-License-Identifier: BSD-3-Clause
//

package common

import "errors"

var (
	ErrIsEmpty    = errors.New("value is empty")
	ErrIsNegative = errors.New("value is negative")
)
