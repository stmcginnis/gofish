//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type FacilityType string

const (
	// RoomFacilityType A room inside of a building or floor.
	RoomFacilityType FacilityType = "Room"
	// FloorFacilityType A floor inside of a building.
	FloorFacilityType FacilityType = "Floor"
	// BuildingFacilityType A structure with a roof and walls.
	BuildingFacilityType FacilityType = "Building"
	// SiteFacilityType A small area consisting of several buildings.
	SiteFacilityType FacilityType = "Site"
)

// Facility shall be used to represent a location containing equipment, such as a room, building, or campus, for a
// Redfish implementation.
type Facility struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AmbientMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the outdoor
	// environment metrics for this facility.
	AmbientMetrics string
	// Description provides a description of this resource.
	Description string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this facility.
	environmentMetrics string
	// FacilityType shall contain the type of location this resource represents.
	FacilityType FacilityType
	// Location shall contain the location information of the associated facility.
	Location common.Location
	// PowerDomains shall contain a link to a resource collection of type PowerDomainCollection that contains the power
	// domains associated with this facility.
	powerDomains string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// CDUs shall contain an array of links to resources of type CoolingUnit that represent the coolant distribution
	// units in this facility.
	cdus []string
	// CDUCount is the number of coolant distribution units in this facility.
	CDUsCount int
	// ContainedByFacility shall contain a link to a resource of type Facility that represents the facility that
	// contains this facility.
	containedByFacility string
	// ContainsChassis shall be an array of links to resources of type Chassis that represent the outermost chassis
	// that this facility contains. This array shall only contain chassis instances that do not include a ContainedBy
	// property within the Links property. That is, only chassis instances that are not contained by another chassis.
	containsChassis []string
	// ContainsChassisCount is the number of Chassis that this facility contains.
	ContainsChassisCount int
	// ContainsFacilities shall be an array of links to resources of type Facility that represent the facilities that
	// this facility contains.
	containsFacilities []string
	// ContainsFacilitiesCount is the number of facilities that this facility contains.
	ContainsFacilitiesCount int
	// CoolingLoops shall contain an array of links to resources of type CoolingLoop that represent the cooling loops
	// in this facility.
	coolingLoops []string
	// CoolingLoopsCount is the number of cooling loops in this facility.
	CoolingLoopsCount int
	// ElectricalBuses shall contain an array of links to resources of type PowerDistribution that represent the
	// electrical buses in this facility.
	electricalBuses []string
	// ElectricalBusesCount is the number of electrical buses in this facility.
	ElectricalBusesCount int
	// FloorPDUs shall be an array of links to resources of type PowerDistribution that represent the floor power
	// distribution units in this facility.
	floorPDUs []string
	// FloorPDUs@odataCount is the number of floor power distribution units in this facility.
	FloorPDUsCount int
	// ImmersionUnits shall contain an array of links to resources of type CoolingUnit that represent the immersion
	// cooling units in this facility.
	immersionUnits []string
	// ImmersionUnitsCount is the number of immersion cooling units in this facility.
	ImmersionUnitsCount int
	// ManagedBy shall be an array of links to resources of type Manager that represent the managers that manager this
	// facility.
	managedBy []string
	// ManagedByCount is the number of Managers that manage this facility.
	ManagedByCount int
	// PowerShelves shall be an array of links to resources of type PowerDistribution that represent the power shelves
	// in this facility.
	powerShelves []string
	// PowerShelvesCount is the number of power shelves in this facility.
	PowerShelvesCount int
	// RackPDUs shall be an array of links to resources of type PowerDistribution that represent the rack-level power
	// distribution units in this facility.
	rackPDUs []string
	// RackPDUsCount is the number of rack-level power distribution units in this facility.
	RackPDUsCount int
	// Switchgear shall be an array of links to resources of type PowerDistribution that represent the switchgear in
	// this facility.
	switchgear []string
	// SwitchgearCount is the number of switch gear in this facility.
	SwitchgearCount int
	// TransferSwitches shall be an array of links to resources of type PowerDistribution that represent the transfer
	// switches in this facility.
	transferSwitches []string
	// TransferSwitchesCount is the number of transfer switches in this community.
	TransferSwitchesCount int
}

type facilityLinks struct {
	// CDUs shall contain an array of links to resources of type CoolingUnit that represent the coolant distribution
	// units in this facility.
	CDUs common.Links
	// CDUs@odata.count
	CDUsCount int `json:"CDUs@odata.count"`
	// ContainedByFacility shall contain a link to a resource of type Facility that represents the facility that
	// contains this facility.
	ContainedByFacility common.Link
	// ContainsChassis shall be an array of links to resources of type Chassis that represent the outermost chassis
	// that this facility contains. This array shall only contain chassis instances that do not include a ContainedBy
	// property within the Links property. That is, only chassis instances that are not contained by another chassis.
	ContainsChassis common.Links
	// ContainsChassis@odata.count
	ContainsChassisCount int `json:"ContainsChassis@odata.count"`
	// ContainsFacilities shall be an array of links to resources of type Facility that represent the facilities that
	// this facility contains.
	ContainsFacilities common.Links
	// ContainsFacilities@odata.count
	ContainsFacilitiesCount int `json:"ContainsFacilities@odata.count"`
	// CoolingLoops shall contain an array of links to resources of type CoolingLoop that represent the cooling loops
	// in this facility.
	CoolingLoops common.Links
	// CoolingLoops@odata.count
	CoolingLoopsCount int `json:"CoolingLoops@odata.count"`
	// ElectricalBuses shall contain an array of links to resources of type PowerDistribution that represent the
	// electrical buses in this facility.
	ElectricalBuses common.Links
	// ElectricalBuses@odata.count
	ElectricalBusesCount int `json:"ElectricalBuses@odata.count"`
	// FloorPDUs shall be an array of links to resources of type PowerDistribution that represent the floor power
	// distribution units in this facility.
	FloorPDUs common.Links
	// FloorPDUs@odata.count
	FloorPDUsCount int `json:"FloorPDUs@odata.count"`
	// ImmersionUnits shall contain an array of links to resources of type CoolingUnit that represent the immersion
	// cooling units in this facility.
	ImmersionUnits common.Links
	// ImmersionUnits@odata.count
	ImmersionUnitsCount int `json:"ImmersionUnits@odata.count"`
	// ManagedBy shall be an array of links to resources of type Manager that represent the managers that manager this
	// facility.
	ManagedBy common.Links
	// ManagedBy@odata.count
	ManagedByCount int `json:"ManagedBy@odata.count"`
	// PowerShelves shall be an array of links to resources of type PowerDistribution that represent the power shelves
	// in this facility.
	PowerShelves common.Links
	// PowerShelves@odata.count
	PowerShelvesCount int `json:"PowerShelves@odata.count"`
	// RackPDUs shall be an array of links to resources of type PowerDistribution that represent the rack-level power
	// distribution units in this facility.
	RackPDUs common.Links
	// RackPDUs@odata.count
	RackPDUsCount int `json:"RackPDUs@odata.count"`
	// Switchgear shall be an array of links to resources of type PowerDistribution that represent the switchgear in
	// this facility.
	Switchgear common.Links
	// Switchgear@odata.count
	SwitchgearCount int `json:"Switchgear@odata.count"`
	// TransferSwitches shall be an array of links to resources of type PowerDistribution that represent the transfer
	// switches in this facility.
	TransferSwitches common.Links
	// TransferSwitches@odata.count
	TransferSwitchesCount int `json:"TransferSwitches@odata.count"`
}

// UnmarshalJSON unmarshals a Facility object from the raw JSON.
func (facility *Facility) UnmarshalJSON(b []byte) error {
	type temp Facility
	var t struct {
		temp
		EnvironmentMetrics common.Link
		PowerDomains       common.Link
		Links              facilityLinks
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*facility = Facility(t.temp)

	// Extract the links to other entities for later
	facility.environmentMetrics = t.EnvironmentMetrics.String()
	facility.powerDomains = t.PowerDomains.String()

	facility.cdus = t.Links.CDUs.ToStrings()
	facility.CDUsCount = t.Links.CDUsCount
	facility.containedByFacility = t.Links.ContainedByFacility.String()
	facility.containsChassis = t.Links.ContainsChassis.ToStrings()
	facility.ContainsChassisCount = t.Links.ContainsChassisCount
	facility.containsFacilities = t.Links.ContainsFacilities.ToStrings()
	facility.ContainsFacilitiesCount = t.Links.ContainsFacilitiesCount
	facility.coolingLoops = t.Links.CoolingLoops.ToStrings()
	facility.CoolingLoopsCount = t.Links.CoolingLoopsCount
	facility.electricalBuses = t.Links.ElectricalBuses.ToStrings()
	facility.ElectricalBusesCount = t.Links.ElectricalBusesCount
	facility.floorPDUs = t.Links.FloorPDUs.ToStrings()
	facility.FloorPDUsCount = t.Links.FloorPDUsCount
	facility.immersionUnits = t.Links.ImmersionUnits.ToStrings()
	facility.ImmersionUnitsCount = t.Links.ImmersionUnitsCount
	facility.managedBy = t.Links.ManagedBy.ToStrings()
	facility.ManagedByCount = t.Links.ManagedByCount
	facility.powerShelves = t.Links.PowerShelves.ToStrings()
	facility.PowerShelvesCount = t.Links.PowerShelvesCount
	facility.rackPDUs = t.Links.RackPDUs.ToStrings()
	facility.RackPDUsCount = t.Links.RackPDUsCount
	facility.switchgear = t.Links.Switchgear.ToStrings()
	facility.SwitchgearCount = t.Links.SwitchgearCount
	facility.transferSwitches = t.Links.TransferSwitches.ToStrings()
	facility.TransferSwitchesCount = t.Links.TransferSwitchesCount

	return nil
}

// CDUs get the cooling distribution units associated with this facility.
func (facility *Facility) CDUs() ([]*CoolingUnit, error) {
	return common.GetObjects[CoolingUnit](facility.GetClient(), facility.cdus)
}

// ContainedByFacility get facility that contains this facility.
func (facility *Facility) ContainedByFacility() (*Facility, error) {
	if facility.containedByFacility == "" {
		return nil, nil
	}
	return GetFacility(facility.GetClient(), facility.containedByFacility)
}

// ContainsChassis get the chassis within this facility.
func (facility *Facility) ContainsChassis() ([]*Chassis, error) {
	return common.GetObjects[Chassis](facility.GetClient(), facility.containsChassis)
}

// ContainsFacilities get facilities within this facility.
func (facility *Facility) ContainsFacilities() ([]*Facility, error) {
	return common.GetObjects[Facility](facility.GetClient(), facility.containsFacilities)
}

// CoolingLoops get cooling loops within this facility.
func (facility *Facility) CoolingLoops() ([]*CoolingLoop, error) {
	return common.GetObjects[CoolingLoop](facility.GetClient(), facility.coolingLoops)
}

// ElectricalBuses get electrical buses within this facility.
func (facility *Facility) ElectricalBuses() ([]*PowerDistribution, error) {
	return common.GetObjects[PowerDistribution](facility.GetClient(), facility.electricalBuses)
}

// FloorPDUs get floor power distribution units within this facility.
func (facility *Facility) FloorPDUs() ([]*PowerDistribution, error) {
	return common.GetObjects[PowerDistribution](facility.GetClient(), facility.floorPDUs)
}

// ImmersionUnits get immersion cooling units within this facility.
func (facility *Facility) ImmersionUnits() ([]*CoolingUnit, error) {
	return common.GetObjects[CoolingUnit](facility.GetClient(), facility.immersionUnits)
}

// ManagedBy gets the managers of this facility.
func (facility *Facility) ManagedBy() ([]*Manager, error) {
	return common.GetObjects[Manager](facility.GetClient(), facility.managedBy)
}

// PowerShelves get power shelves within this facility.
func (facility *Facility) PowerShelves() ([]*PowerDistribution, error) {
	return common.GetObjects[PowerDistribution](facility.GetClient(), facility.powerShelves)
}

// RackPDUs get rack power distribution units within this facility.
func (facility *Facility) RackPDUs() ([]*PowerDistribution, error) {
	return common.GetObjects[PowerDistribution](facility.GetClient(), facility.rackPDUs)
}

// Switchgear get switchgear power distribution units within this facility.
func (facility *Facility) Switchgear() ([]*PowerDistribution, error) {
	return common.GetObjects[PowerDistribution](facility.GetClient(), facility.switchgear)
}

// TransferSwitches get transfer switches within this facility.
func (facility *Facility) TransferSwitches() ([]*PowerDistribution, error) {
	return common.GetObjects[PowerDistribution](facility.GetClient(), facility.transferSwitches)
}

// GetFacility will get a Facility instance from the service.
func GetFacility(c common.Client, uri string) (*Facility, error) {
	return common.GetObject[Facility](c, uri)
}

// ListReferencedFacilities gets the collection of Facility from
// a provided reference.
func ListReferencedFacilities(c common.Client, link string) ([]*Facility, error) {
	return common.GetCollectionObjects[Facility](c, link)
}
