//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"math"
)

// Redundancy gets the Redundancy linked resources.
func (p *PowerSupply) Redundancy() ([]*Redundancy, error) {
	return GetObjects[Redundancy](p.client, p.redundancy)
}

// RedundancySet returns the power supplies in the specified redundancy group.
// memberID identifies the redundancy group in the Redundancy array.
func (p *Power) RedundancySet(memberID int) []PowerSupply {
	var powerSupplies []PowerSupply
	if len(p.Redundancy) >= memberID+1 {
		for _, psLink := range p.Redundancy[memberID].redundancySet {
			for i := range p.PowerSupplies {
				if p.PowerSupplies[i].ODataID == psLink {
					powerSupplies = append(powerSupplies, p.PowerSupplies[i])
				}
			}
		}
	}
	return powerSupplies
}

// powerControls converts PowerControl, which some services send as a single
// object instead of an array.
func powerControls(v any) []PowerControl {
	data, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	var controls []PowerControl
	if json.Unmarshal(data, &controls) == nil {
		return controls
	}
	var control PowerControl
	if json.Unmarshal(data, &control) == nil {
		return []PowerControl{control}
	}
	return nil
}

// roundedUint converts IntervalInMin, which some services send as a string or
// a fractional number, rounding to the nearest minute.
func roundedUint(v any) *uint {
	f := toFloat32(v)
	if f == nil {
		return nil
	}
	u := uint(math.Round(float64(*f)))
	return &u
}

// postUnmarshal records the Redundancy links.
func (p *PowerSupply) postUnmarshal(b []byte) error {
	var t struct {
		Redundancy Links
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	p.redundancy = t.Redundancy.ToStrings()

	return nil
}
