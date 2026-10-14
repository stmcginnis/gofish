//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a Redundancy object from the raw JSON.
func (r *Redundancy) UnmarshalJSON(b []byte) error {
	type temp Redundancy
	var tmp struct {
		temp
		ActiveRedundancySet Links `json:"ActiveRedundancySet"`
		RedundancySet       Links
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = Redundancy(tmp.temp)

	// Extract the links to other entities for later
	r.activeRedundancySet = tmp.ActiveRedundancySet.ToStrings()
	r.redundancySet = tmp.RedundancySet.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	r.RawData = b

	return nil
}

// RedundancySet gets the RedundancySet linked resources.
func (r *Redundancy) RedundancySet(client Client) ([]*Resource, error) {
	return GetObjects[Resource](client, r.redundancySet)
}
