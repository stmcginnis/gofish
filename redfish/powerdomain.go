//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.3 - #PowerDomain.v1_2_2.PowerDomain

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// PowerDomain shall represent a DCIM power domain for a Redfish implementation.
type PowerDomain struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// electricalBuses are the URIs for ElectricalBuses.
	electricalBuses []string
	// floorPDUs are the URIs for FloorPDUs.
	floorPDUs []string
	// managedBy are the URIs for ManagedBy.
	managedBy []string
	// powerShelves are the URIs for PowerShelves.
	powerShelves []string
	// rackPDUs are the URIs for RackPDUs.
	rackPDUs []string
	// switchgear are the URIs for Switchgear.
	switchgear []string
	// transferSwitches are the URIs for TransferSwitches.
	transferSwitches []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PowerDomain object from the raw JSON.
func (p *PowerDomain) UnmarshalJSON(b []byte) error {
	type temp PowerDomain
	type pLinks struct {
		ElectricalBuses  common.Links `json:"ElectricalBuses"`
		FloorPDUs        common.Links `json:"FloorPDUs"`
		ManagedBy        common.Links `json:"ManagedBy"`
		PowerShelves     common.Links `json:"PowerShelves"`
		RackPDUs         common.Links `json:"RackPDUs"`
		Switchgear       common.Links `json:"Switchgear"`
		TransferSwitches common.Links `json:"TransferSwitches"`
	}
	var tmp struct {
		temp
		Links pLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PowerDomain(tmp.temp)

	// Extract the links to other entities for later
	p.electricalBuses = tmp.Links.ElectricalBuses.ToStrings()
	p.floorPDUs = tmp.Links.FloorPDUs.ToStrings()
	p.managedBy = tmp.Links.ManagedBy.ToStrings()
	p.powerShelves = tmp.Links.PowerShelves.ToStrings()
	p.rackPDUs = tmp.Links.RackPDUs.ToStrings()
	p.switchgear = tmp.Links.Switchgear.ToStrings()
	p.transferSwitches = tmp.Links.TransferSwitches.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	p.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *PowerDomain) Update() error {
	readWriteFields := []string{
		"Status",
	}

	return p.UpdateFromRawData(p, p.rawData, readWriteFields)
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

// ElectricalBuses gets the ElectricalBuses linked resources.
func (p *PowerDomain) ElectricalBuses(client common.Client) ([]*PowerDistribution, error) {
	return common.GetObjects[PowerDistribution](client, p.electricalBuses)
}

// FloorPDUs gets the FloorPDUs linked resources.
func (p *PowerDomain) FloorPDUs(client common.Client) ([]*PowerDistribution, error) {
	return common.GetObjects[PowerDistribution](client, p.floorPDUs)
}

// ManagedBy gets the ManagedBy linked resources.
func (p *PowerDomain) ManagedBy(client common.Client) ([]*Manager, error) {
	return common.GetObjects[Manager](client, p.managedBy)
}

// PowerShelves gets the PowerShelves linked resources.
func (p *PowerDomain) PowerShelves(client common.Client) ([]*PowerDistribution, error) {
	return common.GetObjects[PowerDistribution](client, p.powerShelves)
}

// RackPDUs gets the RackPDUs linked resources.
func (p *PowerDomain) RackPDUs(client common.Client) ([]*PowerDistribution, error) {
	return common.GetObjects[PowerDistribution](client, p.rackPDUs)
}

// Switchgear gets the Switchgear linked resources.
func (p *PowerDomain) Switchgear(client common.Client) ([]*PowerDistribution, error) {
	return common.GetObjects[PowerDistribution](client, p.switchgear)
}

// TransferSwitches gets the TransferSwitches linked resources.
func (p *PowerDomain) TransferSwitches(client common.Client) ([]*PowerDistribution, error) {
	return common.GetObjects[PowerDistribution](client, p.transferSwitches)
}
