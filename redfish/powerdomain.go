//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// PowerDomain shall be used to represent a DCIM power domain for a Redfish implementation.
type PowerDomain struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	Actions string
	// Description provides a description of this resource.
	Description string
	// Links shall contain links to resources that are related to but are not contained by, or subordinate to, this
	// resource.
	Links string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status

	electricalBuses []string
	// ElectricalBusesCount is the number of electrical buses in this power domain.
	ElectricalBusesCount int
	floorPDUs            []string
	// FloorPDUsCount is the number of floor power distribution units in this power domain.
	FloorPDUsCount int
	managedBy      []string
	// ManagedByCount is the number of managers for this power domain.
	ManagedByCount int
	powerShelves   []string
	// PowerShelvesCount is the number of power shelves in this power domain.
	PowerShelvesCount int
	rackPDUs          []string
	// RackPDUsCount is the number of rack-level power distribution units in this power domain.
	RackPDUsCount int
	switchgear    []string
	// SwitchGearCount is the number of switchgear in this power domain.
	SwitchgearCount  int
	transferSwitches []string
	// TransferSwitchesCount is the number of transfer switches in this power domain.
	TransferSwitchesCount int
}

// UnmarshalJSON unmarshals a PowerDomain object from the raw JSON.
func (powerdomain *PowerDomain) UnmarshalJSON(b []byte) error {
	type temp PowerDomain
	type Links struct {
		ElectricalBuses       common.Links
		ElectricalBusesCount  int `json:"ElectricalBuses@odata.count"`
		FloorPDUs             common.Links
		FloorPDUsCount        int `json:"FloorPDUs@odata.count"`
		ManagedBy             common.Links
		ManagedByCount        int `json:"ManagedBy@odata.count"`
		PowerShelves          common.Links
		PowerShelvesCount     int `json:"PowerShelves@odata.count"`
		RackPDUs              common.Links
		RackPDUsCount         int `json:"RackPDUs@odata.count"`
		Switchgear            common.Links
		SwitchgearCount       int `json:"Switchgear@odata.count"`
		TransferSwitches      common.Links
		TransferSwitchesCount int `json:"TransferSwitches@odata.count"`
	}
	var t struct {
		temp
		Links Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powerdomain = PowerDomain(t.temp)

	// Extract the links to other entities for later
	powerdomain.electricalBuses = t.Links.ElectricalBuses.ToStrings()
	powerdomain.ElectricalBusesCount = t.Links.ElectricalBusesCount
	powerdomain.floorPDUs = t.Links.FloorPDUs.ToStrings()
	powerdomain.FloorPDUsCount = t.Links.FloorPDUsCount
	powerdomain.managedBy = t.Links.ManagedBy.ToStrings()
	powerdomain.ManagedByCount = t.Links.ManagedByCount
	powerdomain.powerShelves = t.Links.PowerShelves.ToStrings()
	powerdomain.PowerShelvesCount = t.Links.PowerShelvesCount
	powerdomain.rackPDUs = t.Links.RackPDUs.ToStrings()
	powerdomain.RackPDUsCount = t.Links.RackPDUsCount
	powerdomain.switchgear = t.Links.Switchgear.ToStrings()
	powerdomain.SwitchgearCount = t.Links.SwitchgearCount
	powerdomain.transferSwitches = t.Links.TransferSwitches.ToStrings()
	powerdomain.TransferSwitchesCount = t.Links.TransferSwitchesCount

	return nil
}

// ElectricalBuses gets the electrical buses in this power domain.
func (powerdomain *PowerDomain) ElectricalBuses() ([]*PowerDistribution, error) {
	var result []*PowerDistribution

	collectionError := common.NewCollectionError()
	for _, ethLink := range powerdomain.electricalBuses {
		eth, err := GetPowerDistribution(powerdomain.GetClient(), ethLink)
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

// FloorPDUs gets the floor power distribution units in this power domain.
func (powerdomain *PowerDomain) FloorPDUs() ([]*PowerDistribution, error) {
	var result []*PowerDistribution

	collectionError := common.NewCollectionError()
	for _, ethLink := range powerdomain.floorPDUs {
		eth, err := GetPowerDistribution(powerdomain.GetClient(), ethLink)
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

// ManagedBy gets the managers that manage this power domain.
func (powerdomain *PowerDomain) ManagedBy() ([]*Manager, error) {
	var result []*Manager

	collectionError := common.NewCollectionError()
	for _, ethLink := range powerdomain.managedBy {
		eth, err := GetManager(powerdomain.GetClient(), ethLink)
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

// PowerShelves gets the power shelves in this power domain.
func (powerdomain *PowerDomain) PowerShelves() ([]*PowerDistribution, error) {
	var result []*PowerDistribution

	collectionError := common.NewCollectionError()
	for _, ethLink := range powerdomain.powerShelves {
		eth, err := GetPowerDistribution(powerdomain.GetClient(), ethLink)
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

// RackPDUs gets the rack-level power distribution units in this power domain.
func (powerdomain *PowerDomain) RackPDUs() ([]*PowerDistribution, error) {
	var result []*PowerDistribution

	collectionError := common.NewCollectionError()
	for _, ethLink := range powerdomain.rackPDUs {
		eth, err := GetPowerDistribution(powerdomain.GetClient(), ethLink)
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

// Switchgear gets the switchgear in this power domain.
func (powerdomain *PowerDomain) Switchgear() ([]*PowerDistribution, error) {
	var result []*PowerDistribution

	collectionError := common.NewCollectionError()
	for _, ethLink := range powerdomain.switchgear {
		eth, err := GetPowerDistribution(powerdomain.GetClient(), ethLink)
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

// TransferSwitches gets the transfer switches in this power domain.
func (powerdomain *PowerDomain) TransferSwitches() ([]*PowerDistribution, error) {
	var result []*PowerDistribution

	collectionError := common.NewCollectionError()
	for _, ethLink := range powerdomain.transferSwitches {
		eth, err := GetPowerDistribution(powerdomain.GetClient(), ethLink)
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

// GetPowerDomain will get a PowerDomain instance from the service.
func GetPowerDomain(c common.Client, uri string) (*PowerDomain, error) {
	return common.GetObject[PowerDomain](c, uri)
}

// ListReferencedPowerDomains gets the collection of PowerDomain from
// a provided reference.
func ListReferencedPowerDomains(c common.Client, link string) ([]*PowerDomain, error) {
	return common.GetCollectionObjects[PowerDomain](c, link)
}
