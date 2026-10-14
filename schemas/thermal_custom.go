//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// Update commits updates to this object's properties to the running system.
func (f *ThermalFan) Update() error {
	readWriteFields := []string{
		"IndicatorLED",
	}

	return f.UpdateFromRawData(f, f.RawData, readWriteFields)
}

// Assembly gets the Assembly linked resource.
func (f *ThermalFan) Assembly() (*Assembly, error) {
	return ResolveOrGet[Assembly](f.client, f.assembly)
}

// Redundancy gets the Redundancy linked resource.
func (f *ThermalFan) Redundancy() (*Redundancy, error) {
	if f.redundancy == "" {
		return nil, nil
	}
	return GetObject[Redundancy](f.client, f.redundancy)
}

// RelatedItem gets the RelatedItem linked resources.
func (f *ThermalFan) RelatedItem() ([]*Entity, error) {
	return GetObjects[Entity](f.client, f.relatedItem)
}

// postUnmarshal records the Redundancy link.
func (f *ThermalFan) postUnmarshal(b []byte) error {
	var t struct {
		Redundancy Link
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	f.redundancy = t.Redundancy.String()

	return nil
}
