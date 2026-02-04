//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/PowerDomain.v1_2_2.json
// 2021.3 - #PowerDomain.v1_2_2.PowerDomain

package schemas

import (
	"encoding/json"
)

// PowerDomain shall represent a DCIM power domain for a Redfish implementation.
type PowerDomain struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status Status
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
}

// UnmarshalJSON unmarshals a PowerDomain object from the raw JSON.
func (p *PowerDomain) UnmarshalJSON(b []byte) error {
	type temp PowerDomain
	type pLinks struct {
		ElectricalBuses  Links `json:"ElectricalBuses"`
		FloorPDUs        Links `json:"FloorPDUs"`
		ManagedBy        Links `json:"ManagedBy"`
		PowerShelves     Links `json:"PowerShelves"`
		RackPDUs         Links `json:"RackPDUs"`
		Switchgear       Links `json:"Switchgear"`
		TransferSwitches Links `json:"TransferSwitches"`
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

	return nil
}

// GetPowerDomain will get a PowerDomain instance from the service.
func GetPowerDomain(c Client, uri string) (*PowerDomain, error) {
	return GetObject[PowerDomain](c, uri)
}

// ListReferencedPowerDomains gets the collection of PowerDomain from
// a provided reference.
func ListReferencedPowerDomains(c Client, link string) ([]*PowerDomain, error) {
	return GetCollectionObjects[PowerDomain](c, link)
}

// ElectricalBuses gets the ElectricalBuses linked resources.
func (p *PowerDomain) ElectricalBuses() ([]*PowerDistribution, error) {
	return GetObjects[PowerDistribution](p.client, p.electricalBuses)
}

// FloorPDUs gets the FloorPDUs linked resources.
func (p *PowerDomain) FloorPDUs() ([]*PowerDistribution, error) {
	return GetObjects[PowerDistribution](p.client, p.floorPDUs)
}

// ManagedBy gets the ManagedBy linked resources.
func (p *PowerDomain) ManagedBy() ([]*Manager, error) {
	return GetObjects[Manager](p.client, p.managedBy)
}

// PowerShelves gets the PowerShelves linked resources.
func (p *PowerDomain) PowerShelves() ([]*PowerDistribution, error) {
	return GetObjects[PowerDistribution](p.client, p.powerShelves)
}

// RackPDUs gets the RackPDUs linked resources.
func (p *PowerDomain) RackPDUs() ([]*PowerDistribution, error) {
	return GetObjects[PowerDistribution](p.client, p.rackPDUs)
}

// Switchgear gets the Switchgear linked resources.
func (p *PowerDomain) Switchgear() ([]*PowerDistribution, error) {
	return GetObjects[PowerDistribution](p.client, p.switchgear)
}

// TransferSwitches gets the TransferSwitches linked resources.
func (p *PowerDomain) TransferSwitches() ([]*PowerDistribution, error) {
	return GetObjects[PowerDistribution](p.client, p.transferSwitches)
}
