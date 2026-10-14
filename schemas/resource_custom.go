//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"strconv"
)

// locationOrdinal converts LocationOrdinalValue, which some services send as a
// string. A missing value is 0, as in earlier releases.
func locationOrdinal(v any) *int {
	i := 0
	if v != nil {
		n, err := strconv.Atoi(parseString(v))
		if err != nil {
			return nil
		}
		i = n
	}
	return &i
}
