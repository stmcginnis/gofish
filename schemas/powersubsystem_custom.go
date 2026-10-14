//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a PowerSubsystem object from the raw JSON.
func (p *PowerSubsystem) UnmarshalJSON(b []byte) error {
	type temp PowerSubsystem
	type pLinks struct {
		PowerBrakeLoadPercentSensor Link `json:"PowerBrakeLoadPercentSensor"`
		PowerBrakeTemperatureSensor Link `json:"PowerBrakeTemperatureSensor"`
	}
	var tmp struct {
		temp
		Links         pLinks
		Batteries     Link `json:"Batteries"`
		PowerSupplies Link `json:"PowerSupplies"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		// Work around a bug in NVIDIA's implementation
		var u struct {
			temp
			Batteries     Link
			PowerSupplies Link
			Allocation    any
		}
		err2 := json.Unmarshal(b, &u)
		if err2 != nil {
			// Still didn't work, return original error
			return err
		}

		tmp.temp = u.temp
		tmp.Batteries = u.Batteries
		tmp.PowerSupplies = u.PowerSupplies

		if u.Allocation != nil {
			convert := func(v any) *float64 {
				if v == nil {
					return nil
				}
				switch val := v.(type) {
				case float64:
					return &val
				case int:
					f := float64(val)
					return &f
				}

				var unknown float64
				return &unknown
			}

			switch val := u.Allocation.(type) {
			case PowerAllocation:
				p.Allocation = val
			case map[string]any:
				p.Allocation = PowerAllocation{}
				if v, ok := val["AllocatedWatts"]; ok {
					p.Allocation.AllocatedWatts = convert(v)
				}
				if v, ok := val["RequestedWatts"]; ok {
					p.Allocation.RequestedWatts = convert(v)
				}
			default:
				p.Allocation = PowerAllocation{}
			}
		}
	}

	*p = PowerSubsystem(tmp.temp)

	// Extract the links to other entities for later
	p.powerBrakeLoadPercentSensor = tmp.Links.PowerBrakeLoadPercentSensor.String()
	p.powerBrakeTemperatureSensor = tmp.Links.PowerBrakeTemperatureSensor.String()
	p.batteries = tmp.Batteries.String()
	p.powerSupplies = tmp.PowerSupplies.String()

	return nil
}
