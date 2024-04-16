//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// PowerAllocation shall contain the set of properties describing the allocation of power for a subsystem.
type PowerAllocation struct {
	// AllocatedWatts shall represent the total power currently allocated or budgeted to this subsystem.
	AllocatedWatts float64
	// RequestedWatts shall represent the amount of power, in watt units, that the subsystem currently requests to be
	// budgeted for future use.
	RequestedWatts float64
}

// PowerSubsystem shall represent a power subsystem for a Redfish implementation.
type PowerSubsystem struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Allocation shall contain the set of properties describing the allocation of power for this subsystem.
	Allocation PowerAllocation
	// Batteries shall contain a link to a resource collection of type BatteryCollection.
	batteries []string
	// CapacityWatts shall represent the total power capacity that can be allocated to this subsystem.
	CapacityWatts float64
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerSupplies shall contain a link to a resource collection of type PowerSupplyCollection.
	powerSupplies []string
	// PowerSupplyRedundancy shall contain redundancy information for the set of power supplies in this subsystem. The
	// values of the RedundancyGroup array shall reference resources of type PowerSupply.
	PowerSupplyRedundancy []RedundantGroup
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a PowerSubsystem object from the raw JSON.
func (powersubsystem *PowerSubsystem) UnmarshalJSON(b []byte) error {
	type temp PowerSubsystem
	var t struct {
		temp
		Batteries     common.LinksCollection
		PowerSupplies common.LinksCollection
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powersubsystem = PowerSubsystem(t.temp)

	// Extract the links to other entities for later
	powersubsystem.batteries = t.Batteries.ToStrings()
	powersubsystem.powerSupplies = t.PowerSupplies.ToStrings()

	return nil
}

// Batteries gets the batteries in this power subsystem.
func (powersubsystem *PowerSubsystem) Batteries() ([]*Battery, error) {
	var result []*Battery

	collectionError := common.NewCollectionError()
	for _, ethLink := range powersubsystem.batteries {
		eth, err := GetBattery(powersubsystem.GetClient(), ethLink)
		if err != nil {
			collectionError.Failures[ethLink] = err
		} else {
			result = append(result, eth)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// PowerSupplies gets the power supplies in this power subsystem.
func (powersubsystem *PowerSubsystem) PowerSupplies() ([]*PowerSupply, error) {
	var result []*PowerSupply

	collectionError := common.NewCollectionError()
	for _, ethLink := range powersubsystem.powerSupplies {
		eth, err := GetPowerSupply(powersubsystem.GetClient(), ethLink)
		if err != nil {
			collectionError.Failures[ethLink] = err
		} else {
			result = append(result, eth)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// GetPowerSubsystem will get a PowerSubsystem instance from the service.
func GetPowerSubsystem(c common.Client, uri string) (*PowerSubsystem, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var powersubsystem PowerSubsystem
	err = json.NewDecoder(resp.Body).Decode(&powersubsystem)
	if err != nil {
		return nil, err
	}

	powersubsystem.SetClient(c)
	return &powersubsystem, nil
}

// ListReferencedPowerSubsystems gets the collection of PowerSubsystem from
// a provided reference.
func ListReferencedPowerSubsystems(c common.Client, link string) ([]*PowerSubsystem, error) {
	var result []*PowerSubsystem
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *PowerSubsystem
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		powersubsystem, err := GetPowerSubsystem(c, link)
		ch <- GetResult{Item: powersubsystem, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
