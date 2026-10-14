//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// Volume gets the Volume linked resource.
func (m *MappedVolume) Volume(client Client) (*Volume, error) {
	if m.volume == "" {
		return nil, nil
	}
	return GetObject[Volume](client, m.volume)
}

// postUnmarshal records the Volume link.
func (m *MappedVolume) postUnmarshal(b []byte) error {
	var t struct {
		Volume Link
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	m.volume = t.Volume.String()

	return nil
}
