//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"math"
)

// UnmarshalJSON unmarshals a Power object from the raw JSON.
func (p *Power) UnmarshalJSON(b []byte) error {
	type temp Power
	type pActions struct {
		PowerSupplyReset ActionTarget `json:"#Power.PowerSupplyReset"`
	}
	var tmp struct {
		temp
		Actions      pActions
		PowerControl json.RawMessage
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = Power(tmp.temp)

	powerControls := []PowerControl{}
	err = json.Unmarshal(tmp.PowerControl, &powerControls)
	if err != nil {
		// Some Cisco implementations return a singular object instead of the
		// expected array.
		powerControl := PowerControl{}
		err2 := json.Unmarshal(tmp.PowerControl, &powerControl)
		if err2 != nil {
			// Return the original error
			return err
		}

		powerControls = append(powerControls, powerControl)
	}

	p.PowerControl = powerControls

	// Extract the links to other entities for later
	p.powerSupplyResetTarget = tmp.Actions.PowerSupplyReset.Target

	return nil
}

// UnmarshalJSON unmarshals a PowerControl object from the raw JSON.
func (p *PowerControl) UnmarshalJSON(b []byte) error {
	type temp PowerControl
	type pActions struct {
		PowerSupplyReset ActionTarget `json:"#Power.PowerSupplyReset"`
	}
	var tmp struct {
		temp
		Actions     pActions
		RelatedItem Links `json:"RelatedItem"`

		// Need to work around some non-standard data types in Dell and Cisco
		// systems.
		MemberID            any `json:"MemberId"`
		PowerAllocatedWatts any
		PowerAvailableWatts any
		PowerCapacityWatts  any
		PowerConsumedWatts  any
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PowerControl(tmp.temp)

	// Standardize the property types
	p.relatedItem = tmp.RelatedItem.ToStrings()
	p.MemberID = parseString(tmp.MemberID)
	p.PowerAllocatedWatts = toFloat32(tmp.PowerAllocatedWatts)
	p.PowerAvailableWatts = toFloat32(tmp.PowerAvailableWatts)
	p.PowerCapacityWatts = toFloat32(tmp.PowerCapacityWatts)
	p.PowerConsumedWatts = toFloat32(tmp.PowerConsumedWatts)

	// Extract the links to other entities for later
	p.powerSupplyResetTarget = tmp.Actions.PowerSupplyReset.Target
	p.relatedItem = tmp.RelatedItem.ToStrings()

	return nil
}

func (p *PowerMetric) UnmarshalJSON(b []byte) error {
	type temp PowerMetric
	type t1 struct {
		temp
		IntervalInMin    any
		MaxConsumedWatts any
		MinConsumedWatts any
	}
	var t t1

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*p = PowerMetric(t.temp)

	if t.IntervalInMin != nil {
		val := uint(math.Round(float64(*toFloat32(t.IntervalInMin))))
		p.IntervalInMin = &val
	}
	p.MaxConsumedWatts = toFloat32(t.MaxConsumedWatts)
	p.MinConsumedWatts = toFloat32(t.MinConsumedWatts)

	return nil
}

// UnmarshalJSON unmarshals a PowerSupply object from the raw JSON.
func (p *PowerSupply) UnmarshalJSON(b []byte) error {
	type temp PowerSupply
	type pActions struct {
		PowerSupplyReset ActionTarget `json:"#Power.PowerSupplyReset"`
	}
	var tmp struct {
		temp
		Actions              pActions
		Assembly             ResolvedLink[Assembly] `json:"Assembly"`
		Redundancy           Links                  `json:"Redundancy"`
		RelatedItem          Links                  `json:"RelatedItem"`
		MemberID             any                    `json:"MemberId"`
		LineInputVoltage     any
		LastPowerOutputWatts any
		PowerInputWatts      any
		PowerOutputWatts     any
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PowerSupply(tmp.temp)

	// Extract the links to other entities for later
	p.powerSupplyResetTarget = tmp.Actions.PowerSupplyReset.Target
	p.assembly = tmp.Assembly

	p.redundancy = tmp.Redundancy.ToStrings()
	p.relatedItem = tmp.RelatedItem.ToStrings()

	p.MemberID = parseString(tmp.MemberID)
	p.LineInputVoltage = toFloat32(tmp.LineInputVoltage)
	p.LastPowerOutputWatts = toFloat32(tmp.LastPowerOutputWatts)
	p.PowerInputWatts = toFloat32(tmp.PowerInputWatts)
	p.PowerOutputWatts = toFloat32(tmp.PowerOutputWatts)

	// This is a read/write object, so we need to save the raw object data for later
	p.RawData = b

	return nil
}

// Redundancy gets the Redundancy linked resources.
func (p *PowerSupply) Redundancy() ([]*Redundancy, error) {
	return GetObjects[Redundancy](p.client, p.redundancy)
}

// UnmarshalJSON unmarshals a Voltage object from the raw JSON.
func (v *Voltage) UnmarshalJSON(b []byte) error {
	type temp Voltage
	type vActions struct {
		PowerSupplyReset ActionTarget `json:"#Power.PowerSupplyReset"`
	}
	var tmp struct {
		temp
		Actions     vActions
		RelatedItem Links `json:"RelatedItem"`

		// Need to work around some non-standard data types in Dell and Cisco
		// systems.
		MemberID                  any `json:"MemberId"`
		UpperThresholdCritical    any
		UpperThresholdFatal       any
		UpperThresholdNonCritical any
		LowerThresholdCritical    any
		LowerThresholdFatal       any
		LowerThresholdNonCritical any
		ReadingVolts              any
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*v = Voltage(tmp.temp)

	// Extract the links to other entities for later
	v.powerSupplyResetTarget = tmp.Actions.PowerSupplyReset.Target
	v.relatedItem = tmp.RelatedItem.ToStrings()

	v.MemberID = parseString(tmp.MemberID)
	v.UpperThresholdCritical = toFloat32(tmp.UpperThresholdCritical)
	v.UpperThresholdFatal = toFloat32(tmp.UpperThresholdFatal)
	v.UpperThresholdNonCritical = toFloat32(tmp.UpperThresholdNonCritical)
	v.LowerThresholdCritical = toFloat32(tmp.LowerThresholdCritical)
	v.LowerThresholdFatal = toFloat32(tmp.LowerThresholdFatal)
	v.LowerThresholdNonCritical = toFloat32(tmp.LowerThresholdNonCritical)
	v.ReadingVolts = toFloat32(tmp.ReadingVolts)

	return nil
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
