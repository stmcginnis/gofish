//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a Fan object from the raw JSON.
func (f *ThermalFan) UnmarshalJSON(b []byte) error {
	type temp ThermalFan
	var tmp struct {
		temp
		Assembly    ResolvedLink[Assembly] `json:"Assembly"`
		Redundancy  Link                   `json:"Redundancy"`
		RelatedItem Links                  `json:"RelatedItem"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = ThermalFan(tmp.temp)

	// Extract the links to other entities for later
	f.assembly = tmp.Assembly

	f.redundancy = tmp.Redundancy.String()
	f.relatedItem = tmp.RelatedItem.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	f.RawData = b

	return nil
}

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
