//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strconv"
)

// UnmarshalJSON unmarshals a ConnectedEntity object from the raw JSON.
func (p *PciID) UnmarshalJSON(b []byte) error {
	type temp PciID
	var tmp struct {
		temp
		FunctionNumber any
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PciID(tmp.temp)

	// Extract the links to other entities for later
	if tmp.FunctionNumber != nil {
		switch val := tmp.FunctionNumber.(type) {
		case string:
			fn, err := strconv.Atoi(val)
			if err == nil {
				p.FunctionNumber = &fn
			}
		case int:
			fn := val
			p.FunctionNumber = &fn
		}
	}

	return nil
}
