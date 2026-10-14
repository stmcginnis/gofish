//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strconv"
)

func (p *PartLocation) UnmarshalJSON(b []byte) error {
	type temp PartLocation
	var t struct {
		temp
		// Some implementations incorrectly have this as a string, e.g. "1" instead of 1.
		LocationOrdinalValue any
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	i := 0
	if t.LocationOrdinalValue != nil {
		i, err = strconv.Atoi(parseString(t.LocationOrdinalValue))
		if err != nil {
			return err
		}
	}

	*p = PartLocation(t.temp)
	p.LocationOrdinalValue = &i

	return nil
}
