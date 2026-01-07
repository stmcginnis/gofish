//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.3 - #PowerEquipment.v1_2_3.PowerEquipment

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// PowerEquipment shall represent the set of power equipment for a Redfish
// implementation.
type PowerEquipment struct {
	common.Entity
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
	// Oem shall contain the OEM extensions. All values for properties that this
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
	Status common.Status
	// Switchgear shall contain a link to a resource collection of type
	// 'PowerDistributionCollection' that contains a set of switchgear.
	switchgear string
	// TransferSwitches shall contain a link to a resource collection of type
	// 'PowerDistributionCollection' that contains a set of transfer switches.
	transferSwitches string
	// managedBy are the URIs for ManagedBy.
	managedBy []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PowerEquipment object from the raw JSON.
func (p *PowerEquipment) UnmarshalJSON(b []byte) error {
	type temp PowerEquipment
	type pLinks struct {
		ManagedBy common.Links `json:"ManagedBy"`
	}
	var tmp struct {
		temp
		Links            pLinks
		ElectricalBuses  common.Link `json:"electricalBuses"`
		FloorPDUs        common.Link `json:"floorPDUs"`
		PowerShelves     common.Link `json:"powerShelves"`
		RackPDUs         common.Link `json:"rackPDUs"`
		Switchgear       common.Link `json:"switchgear"`
		TransferSwitches common.Link `json:"transferSwitches"`
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

	// This is a read/write object, so we need to save the raw object data for later
	p.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *PowerEquipment) Update() error {
	readWriteFields := []string{
		"Status",
	}

	return p.UpdateFromRawData(p, p.rawData, readWriteFields)
}

// GetPowerEquipment will get a PowerEquipment instance from the service.
func GetPowerEquipment(c common.Client, uri string) (*PowerEquipment, error) {
	return common.GetObject[PowerEquipment](c, uri)
}

// ListReferencedPowerEquipments gets the collection of PowerEquipment from
// a provided reference.
func ListReferencedPowerEquipments(c common.Client, link string) ([]*PowerEquipment, error) {
	return common.GetCollectionObjects[PowerEquipment](c, link)
}

// ManagedBy gets the ManagedBy linked resources.
func (p *PowerEquipment) ManagedBy(client common.Client) ([]*Manager, error) {
	return common.GetObjects[Manager](client, p.managedBy)
}

// ElectricalBuses gets the ElectricalBuses collection.
func (p *PowerEquipment) ElectricalBuses(client common.Client) ([]*PowerDistribution, error) {
	if p.electricalBuses == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[PowerDistribution](client, p.electricalBuses)
}

// FloorPDUs gets the FloorPDUs collection.
func (p *PowerEquipment) FloorPDUs(client common.Client) ([]*PowerDistribution, error) {
	if p.floorPDUs == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[PowerDistribution](client, p.floorPDUs)
}

// PowerShelves gets the PowerShelves collection.
func (p *PowerEquipment) PowerShelves(client common.Client) ([]*PowerDistribution, error) {
	if p.powerShelves == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[PowerDistribution](client, p.powerShelves)
}

// RackPDUs gets the RackPDUs collection.
func (p *PowerEquipment) RackPDUs(client common.Client) ([]*PowerDistribution, error) {
	if p.rackPDUs == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[PowerDistribution](client, p.rackPDUs)
}

// Switchgear gets the Switchgear collection.
func (p *PowerEquipment) Switchgear(client common.Client) ([]*PowerDistribution, error) {
	if p.switchgear == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[PowerDistribution](client, p.switchgear)
}

// TransferSwitches gets the TransferSwitches collection.
func (p *PowerEquipment) TransferSwitches(client common.Client) ([]*PowerDistribution, error) {
	if p.transferSwitches == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[PowerDistribution](client, p.transferSwitches)
}
