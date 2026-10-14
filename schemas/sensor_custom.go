//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strconv"
)

// UnmarshalJSON unmarshals SensorArrayExcerpt from JSON, handling Reading as either a number or string.
func (s *SensorArrayExcerpt) UnmarshalJSON(b []byte) error {
	type temp SensorArrayExcerpt
	var t temp

	if err := json.Unmarshal(b, &t); err != nil {
		// Try with Reading as string
		var stringT struct {
			temp
			Reading string `json:",omitempty"`
		}
		if err := json.Unmarshal(b, &stringT); err != nil {
			return err
		}
		*s = SensorArrayExcerpt(stringT.temp)
		if stringT.Reading != "" {
			f, err := strconv.ParseFloat(stringT.Reading, 64)
			if err != nil {
				return err
			}
			s.Reading = &f
		}
		return nil
	}

	*s = SensorArrayExcerpt(t)
	return nil
}
