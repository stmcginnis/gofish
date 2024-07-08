//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// PowerEquipment shall be used to represent
// the set of power equipment for a Redfish implementation.
type PowerEquipment struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string

	// A link to a collection of electrical buses.
	electricalBuses string
	// A link to a collection of floor power distribution units.
	floorPDUs string
	// A link to a collection of power shelves.
	powerShelves string
	// A link to a collection of rack-level power distribution units.
	rackPDUs string
	// The status and health of the resource and its subordinate or dependent resources.
	Status common.Status
	// A link to a collection of switchgear.
	switchgear string
	// A link to a collection of transfer switches.
	transferSwitches string

	// Links section
	// An array of links to the managers responsible for managing this power equipment.
	managedBy      []string
	ManagedByCount int
	// OemLinks are all OEM data under link section
	OemLinks json.RawMessage

	// Actions section
	// OemActions contains all the vendor specific actions.
	// It is vendor responsibility to parse this field accordingly
	OemActions json.RawMessage
}

// UnmarshalJSON unmarshals a PowerEquipment object from the raw JSON.
func (powerEquipment *PowerEquipment) UnmarshalJSON(b []byte) error {
	type temp PowerEquipment
	type linkReference struct {
		ManagedBy      common.Links
		ManagedByCount int `json:"ManagedBy@odata.count"`
		Oem            json.RawMessage
	}
	type actions struct {
		Oem json.RawMessage // OEM actions will be stored here
	}
	var t struct {
		temp
		ElectricalBuses  common.Link
		FloorPDUs        common.Link
		PowerShelves     common.Link
		RackPDUs         common.Link
		Switchgear       common.Link
		TransferSwitches common.Link
		Links            linkReference
		Actions          actions
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	// Extract the links to other entities for later
	*powerEquipment = PowerEquipment(t.temp)
	powerEquipment.electricalBuses = t.ElectricalBuses.String()
	powerEquipment.floorPDUs = t.FloorPDUs.String()
	powerEquipment.powerShelves = t.PowerShelves.String()
	powerEquipment.rackPDUs = t.RackPDUs.String()
	powerEquipment.switchgear = t.Switchgear.String()
	powerEquipment.transferSwitches = t.TransferSwitches.String()

	powerEquipment.managedBy = t.Links.ManagedBy.ToStrings()
	powerEquipment.ManagedByCount = t.Links.ManagedByCount
	powerEquipment.OemLinks = t.Links.Oem

	powerEquipment.OemActions = t.Actions.Oem

	return nil
}

// GetPowerEquipment will get a PowerEquipment instance from the Redfish service.
func GetPowerEquipment(c common.Client, uri string) (*PowerEquipment, error) {
	return common.GetObject[PowerEquipment](c, uri)
}

// ManagedBy gets the collection of managers of this PowerEquipment
func (powerEquipment *PowerEquipment) ManagedBy() ([]*Manager, error) {
	return common.GetObjects[Manager](powerEquipment.GetClient(), powerEquipment.managedBy)
}

// ElectricalBuses gets the collection that contains a set of electrical bus units.
func (powerEquipment *PowerEquipment) ElectricalBuses() ([]*PowerDistribution, error) {
	return ListReferencedPowerDistributionUnits(powerEquipment.GetClient(), powerEquipment.electricalBuses)
}

// FloorPDUs gets the collection that contains a set of floor power distribution units.
func (powerEquipment *PowerEquipment) FloorPDUs() ([]*PowerDistribution, error) {
	return ListReferencedPowerDistributionUnits(powerEquipment.GetClient(), powerEquipment.floorPDUs)
}

// PowerShelves gets the collection that contains a set of power shelves.
func (powerEquipment *PowerEquipment) PowerShelves() ([]*PowerDistribution, error) {
	return ListReferencedPowerDistributionUnits(powerEquipment.GetClient(), powerEquipment.powerShelves)
}

// RackPDUs gets the collection that contains a set of rack-level power distribution units.
func (powerEquipment *PowerEquipment) RackPDUs() ([]*PowerDistribution, error) {
	return ListReferencedPowerDistributionUnits(powerEquipment.GetClient(), powerEquipment.rackPDUs)
}

// Switchgear gets the collection that contains a set of switchgear.
func (powerEquipment *PowerEquipment) Switchgear() ([]*PowerDistribution, error) {
	return ListReferencedPowerDistributionUnits(powerEquipment.GetClient(), powerEquipment.switchgear)
}

// TransferSwitches gets the collection that contains a set of transfer switches.
func (powerEquipment *PowerEquipment) TransferSwitches() ([]*PowerDistribution, error) {
	return ListReferencedPowerDistributionUnits(powerEquipment.GetClient(), powerEquipment.transferSwitches)
}
