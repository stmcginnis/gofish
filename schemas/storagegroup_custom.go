//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a MappedVolume object from the raw JSON.
func (m *MappedVolume) UnmarshalJSON(b []byte) error {
	type temp MappedVolume
	var tmp struct {
		temp
		LogicalUnitNumber any
		Volume            Link `json:"Volume"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MappedVolume(tmp.temp)

	// Extract the links to other entities for later
	m.volume = tmp.Volume.String()
	m.LogicalUnitNumber = parseString(tmp.LogicalUnitNumber)

	return nil
}

// Volume gets the Volume linked resource.
func (m *MappedVolume) Volume(client Client) (*Volume, error) {
	if m.volume == "" {
		return nil, nil
	}
	return GetObject[Volume](client, m.volume)
}
