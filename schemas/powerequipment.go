//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/PowerEquipment.v1_2_3.json
// 2021.3 - #PowerEquipment.v1_2_3.PowerEquipment

package schemas

import (
	"encoding/json"
)

// PowerEquipment shall represent the set of power equipment for a Redfish
// implementation.
type PowerEquipment struct {
	Entity
	// ElectricalBuses shall contain a link to a resource collection of type
	// 'PowerDistributionCollection' that contains a set of electrical bus units.
	//
	// Version added: v1.2.0
	electricalBuses string
	// FloorPDUs shall contain a link to a resource collection of type
	// 'PowerDistributionCollection' that contains a set of floor power
	// distribution units.
	floorPDUs string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerShelves shall contain a link to a resource collection of type
	// 'PowerDistributionCollection' that contains a set of power shelves.
	//
	// Version added: v1.1.0
	powerShelves string
	// RackPDUs shall contain a link to a resource collection of type
	// 'PowerDistributionCollection' that contains a set of rack-level power
	// distribution units.
	rackPDUs string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// Switchgear shall contain a link to a resource collection of type
	// 'PowerDistributionCollection' that contains a set of switchgear.
	switchgear string
	// TransferSwitches shall contain a link to a resource collection of type
	// 'PowerDistributionCollection' that contains a set of transfer switches.
	transferSwitches string
	// managedBy are the URIs for ManagedBy.
	managedBy []string
}

// UnmarshalJSON unmarshals a PowerEquipment object from the raw JSON.
func (p *PowerEquipment) UnmarshalJSON(b []byte) error {
	type temp PowerEquipment
	type pLinks struct {
		ManagedBy Links `json:"ManagedBy"`
	}
	var tmp struct {
		temp
		Links            pLinks
		ElectricalBuses  Link `json:"ElectricalBuses"`
		FloorPDUs        Link `json:"FloorPDUs"`
		PowerShelves     Link `json:"PowerShelves"`
		RackPDUs         Link `json:"RackPDUs"`
		Switchgear       Link `json:"Switchgear"`
		TransferSwitches Link `json:"TransferSwitches"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PowerEquipment(tmp.temp)

	// Extract the links to other entities for later
	p.managedBy = tmp.Links.ManagedBy.ToStrings()
	p.electricalBuses = tmp.ElectricalBuses.String()
	p.floorPDUs = tmp.FloorPDUs.String()
	p.powerShelves = tmp.PowerShelves.String()
	p.rackPDUs = tmp.RackPDUs.String()
	p.switchgear = tmp.Switchgear.String()
	p.transferSwitches = tmp.TransferSwitches.String()

	return nil
}

// GetPowerEquipment will get a PowerEquipment instance from the service.
func GetPowerEquipment(c Client, uri string) (*PowerEquipment, error) {
	return GetObject[PowerEquipment](c, uri)
}

// ListReferencedPowerEquipments gets the collection of PowerEquipment from
// a provided reference.
func ListReferencedPowerEquipments(c Client, link string) ([]*PowerEquipment, error) {
	return GetCollectionObjects[PowerEquipment](c, link)
}

// ManagedBy gets the ManagedBy linked resources.
func (p *PowerEquipment) ManagedBy() ([]*Manager, error) {
	return GetObjects[Manager](p.client, p.managedBy)
}

// ElectricalBuses gets the ElectricalBuses collection.
func (p *PowerEquipment) ElectricalBuses() ([]*PowerDistribution, error) {
	if p.electricalBuses == "" {
		return nil, nil
	}
	return GetCollectionObjects[PowerDistribution](p.client, p.electricalBuses)
}

// FloorPDUs gets the FloorPDUs collection.
func (p *PowerEquipment) FloorPDUs() ([]*PowerDistribution, error) {
	if p.floorPDUs == "" {
		return nil, nil
	}
	return GetCollectionObjects[PowerDistribution](p.client, p.floorPDUs)
}

// PowerShelves gets the PowerShelves collection.
func (p *PowerEquipment) PowerShelves() ([]*PowerDistribution, error) {
	if p.powerShelves == "" {
		return nil, nil
	}
	return GetCollectionObjects[PowerDistribution](p.client, p.powerShelves)
}

// RackPDUs gets the RackPDUs collection.
func (p *PowerEquipment) RackPDUs() ([]*PowerDistribution, error) {
	if p.rackPDUs == "" {
		return nil, nil
	}
	return GetCollectionObjects[PowerDistribution](p.client, p.rackPDUs)
}

// Switchgear gets the Switchgear collection.
func (p *PowerEquipment) Switchgear() ([]*PowerDistribution, error) {
	if p.switchgear == "" {
		return nil, nil
	}
	return GetCollectionObjects[PowerDistribution](p.client, p.switchgear)
}

// TransferSwitches gets the TransferSwitches collection.
func (p *PowerEquipment) TransferSwitches() ([]*PowerDistribution, error) {
	if p.transferSwitches == "" {
		return nil, nil
	}
	return GetCollectionObjects[PowerDistribution](p.client, p.transferSwitches)
}
