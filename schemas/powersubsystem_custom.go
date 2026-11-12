//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

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
