//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// RedundancySet gets the RedundancySet linked resources.
func (r *Redundancy) RedundancySet(client Client) ([]*Resource, error) {
	return GetObjects[Resource](client, r.redundancySet)
}

// postUnmarshal records the RedundancySet links.
func (r *Redundancy) postUnmarshal(b []byte) error {
	var t struct {
		RedundancySet Links
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	r.redundancySet = t.RedundancySet.ToStrings()

	return nil
}
