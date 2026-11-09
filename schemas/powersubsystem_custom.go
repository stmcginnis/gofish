//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// PowerSupplies gets the PowerSupplies collection.
//
// The generator derives a collection's member type from the collection name
// (PowerSupplyCollection -> PowerSupply), which resolves to the legacy
// Power.PowerSupply type that has no Metrics link. PowerSupplyCollection
// members are standalone PowerSupply resources, so this returns PowerSupplyUnit,
// which carries the Metrics link.
func (p *PowerSubsystem) PowerSupplies() ([]*PowerSupplyUnit, error) {
	if p.powerSupplies == "" {
		return nil, nil
	}
	return GetCollectionObjects[PowerSupplyUnit](p.client, p.powerSupplies)
}

// powerAllocation converts Allocation, whose watt values some NVIDIA services
// send with non-numeric types. Values that are not numbers become 0.
func powerAllocation(v any) PowerAllocation {
	var a PowerAllocation
	data, err := json.Marshal(v)
	if err == nil && json.Unmarshal(data, &a) == nil {
		return a
	}

	m, ok := v.(map[string]any)
	if !ok {
		return a
	}
	watts := func(v any) *float64 {
		if v == nil {
			return nil
		}
		f, _ := v.(float64)
		return &f
	}
	if w, ok := m["AllocatedWatts"]; ok {
		a.AllocatedWatts = watts(w)
	}
	if w, ok := m["RequestedWatts"]; ok {
		a.RequestedWatts = watts(w)
	}
	return a
}
