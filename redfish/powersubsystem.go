//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.2 - #PowerSubsystem.v1_1_3.PowerSubsystem

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// PowerSubsystem shall represent a power subsystem for a Redfish
// implementation.
type PowerSubsystem struct {
	common.Entity
	// Allocation shall contain the set of properties describing the allocation of
	// power for this subsystem as part of the power infrastructure for the chassis
	// or an upstream chassis. This property should not be present in resources
	// that are not part of a shared power infrastructure.
	Allocation PowerAllocation
	// Batteries shall contain a link to a resource collection of type
	// 'BatteryCollection'.
	//
	// Version added: v1.1.0
	batteries string
	// CapacityWatts shall represent the total power capacity that can be allocated
	// to this subsystem.
	CapacityWatts *float64 `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerSupplies shall contain a link to a resource collection of type
	// 'PowerSupplyCollection'.
	powerSupplies string
	// PowerSupplyRedundancy shall contain redundancy information for the set of
	// power supplies in this subsystem. The values of the 'RedundancyGroup' array
	// shall reference resources of type 'PowerSupply'.
	PowerSupplyRedundancy []common.RedundantGroup
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PowerSubsystem object from the raw JSON.
func (p *PowerSubsystem) UnmarshalJSON(b []byte) error {
	type temp PowerSubsystem
	var tmp struct {
		temp
		Batteries     common.Link `json:"batteries"`
		PowerSupplies common.Link `json:"powerSupplies"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		// Work around a bug in NVIDIA's implementation
		var u struct {
			temp
			Batteries     common.Link
			PowerSupplies common.Link
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
	p.batteries = tmp.Batteries.String()
	p.powerSupplies = tmp.PowerSupplies.String()

	// This is a read/write object, so we need to save the raw object data for later
	p.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *PowerSubsystem) Update() error {
	readWriteFields := []string{
		"Allocation",
		"PowerSupplyRedundancy",
		"Status",
	}

	return p.UpdateFromRawData(p, p.rawData, readWriteFields)
}

// GetPowerSubsystem will get a PowerSubsystem instance from the service.
func GetPowerSubsystem(c common.Client, uri string) (*PowerSubsystem, error) {
	return common.GetObject[PowerSubsystem](c, uri)
}

// ListReferencedPowerSubsystems gets the collection of PowerSubsystem from
// a provided reference.
func ListReferencedPowerSubsystems(c common.Client, link string) ([]*PowerSubsystem, error) {
	return common.GetCollectionObjects[PowerSubsystem](c, link)
}

// Batteries gets the Batteries collection.
func (p *PowerSubsystem) Batteries(client common.Client) ([]*Battery, error) {
	if p.batteries == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Battery](client, p.batteries)
}

// PowerSupplies gets the PowerSupplies collection.
func (p *PowerSubsystem) PowerSupplies(client common.Client) ([]*PowerSupply, error) {
	if p.powerSupplies == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[PowerSupply](client, p.powerSupplies)
}

// PowerAllocation shall contain the set of properties describing the allocation
// of power for a subsystem.
type PowerAllocation struct {
	// AllocatedWatts shall represent the total power currently allocated or
	// budgeted to this subsystem.
	AllocatedWatts *float64 `json:",omitempty"`
	// RequestedWatts shall represent the amount of power, in watt units, that the
	// subsystem currently requests to be budgeted for future use.
	RequestedWatts *float64 `json:",omitempty"`
}
