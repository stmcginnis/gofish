//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

func (ml *MemoryLocation) UnmarshalJSON(b []byte) error {
	type temp MemoryLocation
	var tmp struct {
		temp

		// Workaround for Inspur and Kaytus bug
		Channel any
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*ml = MemoryLocation(tmp.temp)
	ml.Channel = toInt(tmp.Channel)

	return nil
}
