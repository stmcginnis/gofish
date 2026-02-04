//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Facility.v1_4_2.json
// 2023.1 - #Facility.v1_4_2.Facility

package schemas

import (
	"encoding/json"
)

type FacilityType string

const (
	// RoomFacilityType is a room inside of a building or floor.
	RoomFacilityType FacilityType = "Room"
	// FloorFacilityType is a floor inside of a building.
	FloorFacilityType FacilityType = "Floor"
	// BuildingFacilityType is a structure with a roof and walls.
	BuildingFacilityType FacilityType = "Building"
	// SiteFacilityType is a small area consisting of several buildings.
	SiteFacilityType FacilityType = "Site"
)

// Facility shall be used to represent a location containing equipment, such as
// a room, building, or campus, for a Redfish implementation.
type Facility struct {
	Entity
	// AmbientMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that specifies the outdoor environment metrics for this
	// facility.
	//
	// Version added: v1.1.0
	ambientMetrics string
	// EnvironmentMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that specifies the environment metrics for this
	// facility.
	//
	// Version added: v1.1.0
	environmentMetrics string
	// FacilityType shall contain the type of location this resource represents.
	FacilityType FacilityType
	// Location shall contain the location information of the associated facility.
	Location Location
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerDomains shall contain a link to a resource collection of type
	// 'PowerDomainCollection' that contains the power domains associated with this
	// facility.
	powerDomains string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// cDUs are the URIs for CDUs.
	cDUs []string
	// containedByFacility is the URI for ContainedByFacility.
	containedByFacility string
	// containsChassis are the URIs for ContainsChassis.
	containsChassis []string
	// containsFacilities are the URIs for ContainsFacilities.
	containsFacilities []string
	// coolingLoops are the URIs for CoolingLoops.
	coolingLoops []string
	// electricalBuses are the URIs for ElectricalBuses.
	electricalBuses []string
	// floorPDUs are the URIs for FloorPDUs.
	floorPDUs []string
	// immersionUnits are the URIs for ImmersionUnits.
	immersionUnits []string
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

// UnmarshalJSON unmarshals a Facility object from the raw JSON.
func (f *Facility) UnmarshalJSON(b []byte) error {
	type temp Facility
	type fLinks struct {
		CDUs                Links `json:"CDUs"`
		ContainedByFacility Link  `json:"ContainedByFacility"`
		ContainsChassis     Links `json:"ContainsChassis"`
		ContainsFacilities  Links `json:"ContainsFacilities"`
		CoolingLoops        Links `json:"CoolingLoops"`
		ElectricalBuses     Links `json:"ElectricalBuses"`
		FloorPDUs           Links `json:"FloorPDUs"`
		ImmersionUnits      Links `json:"ImmersionUnits"`
		ManagedBy           Links `json:"ManagedBy"`
		PowerShelves        Links `json:"PowerShelves"`
		RackPDUs            Links `json:"RackPDUs"`
		Switchgear          Links `json:"Switchgear"`
		TransferSwitches    Links `json:"TransferSwitches"`
	}
	var tmp struct {
		temp
		Links              fLinks
		AmbientMetrics     Link `json:"AmbientMetrics"`
		EnvironmentMetrics Link `json:"EnvironmentMetrics"`
		PowerDomains       Link `json:"PowerDomains"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = Facility(tmp.temp)

	// Extract the links to other entities for later
	f.cDUs = tmp.Links.CDUs.ToStrings()
	f.containedByFacility = tmp.Links.ContainedByFacility.String()
	f.containsChassis = tmp.Links.ContainsChassis.ToStrings()
	f.containsFacilities = tmp.Links.ContainsFacilities.ToStrings()
	f.coolingLoops = tmp.Links.CoolingLoops.ToStrings()
	f.electricalBuses = tmp.Links.ElectricalBuses.ToStrings()
	f.floorPDUs = tmp.Links.FloorPDUs.ToStrings()
	f.immersionUnits = tmp.Links.ImmersionUnits.ToStrings()
	f.managedBy = tmp.Links.ManagedBy.ToStrings()
	f.powerShelves = tmp.Links.PowerShelves.ToStrings()
	f.rackPDUs = tmp.Links.RackPDUs.ToStrings()
	f.switchgear = tmp.Links.Switchgear.ToStrings()
	f.transferSwitches = tmp.Links.TransferSwitches.ToStrings()
	f.ambientMetrics = tmp.AmbientMetrics.String()
	f.environmentMetrics = tmp.EnvironmentMetrics.String()
	f.powerDomains = tmp.PowerDomains.String()

	return nil
}

// GetFacility will get a Facility instance from the service.
func GetFacility(c Client, uri string) (*Facility, error) {
	return GetObject[Facility](c, uri)
}

// ListReferencedFacilitys gets the collection of Facility from
// a provided reference.
func ListReferencedFacilitys(c Client, link string) ([]*Facility, error) {
	return GetCollectionObjects[Facility](c, link)
}

// CDUs gets the CDUs linked resources.
func (f *Facility) CDUs() ([]*CoolingUnit, error) {
	return GetObjects[CoolingUnit](f.client, f.cDUs)
}

// ContainedByFacility gets the ContainedByFacility linked resource.
func (f *Facility) ContainedByFacility() (*Facility, error) {
	if f.containedByFacility == "" {
		return nil, nil
	}
	return GetObject[Facility](f.client, f.containedByFacility)
}

// ContainsChassis gets the ContainsChassis linked resources.
func (f *Facility) ContainsChassis() ([]*Chassis, error) {
	return GetObjects[Chassis](f.client, f.containsChassis)
}

// ContainsFacilities gets the ContainsFacilities linked resources.
func (f *Facility) ContainsFacilities() ([]*Facility, error) {
	return GetObjects[Facility](f.client, f.containsFacilities)
}

// CoolingLoops gets the CoolingLoops linked resources.
func (f *Facility) CoolingLoops() ([]*CoolingLoop, error) {
	return GetObjects[CoolingLoop](f.client, f.coolingLoops)
}

// ElectricalBuses gets the ElectricalBuses linked resources.
func (f *Facility) ElectricalBuses() ([]*PowerDistribution, error) {
	return GetObjects[PowerDistribution](f.client, f.electricalBuses)
}

// FloorPDUs gets the FloorPDUs linked resources.
func (f *Facility) FloorPDUs() ([]*PowerDistribution, error) {
	return GetObjects[PowerDistribution](f.client, f.floorPDUs)
}

// ImmersionUnits gets the ImmersionUnits linked resources.
func (f *Facility) ImmersionUnits() ([]*CoolingUnit, error) {
	return GetObjects[CoolingUnit](f.client, f.immersionUnits)
}

// ManagedBy gets the ManagedBy linked resources.
func (f *Facility) ManagedBy() ([]*Manager, error) {
	return GetObjects[Manager](f.client, f.managedBy)
}

// PowerShelves gets the PowerShelves linked resources.
func (f *Facility) PowerShelves() ([]*PowerDistribution, error) {
	return GetObjects[PowerDistribution](f.client, f.powerShelves)
}

// RackPDUs gets the RackPDUs linked resources.
func (f *Facility) RackPDUs() ([]*PowerDistribution, error) {
	return GetObjects[PowerDistribution](f.client, f.rackPDUs)
}

// Switchgear gets the Switchgear linked resources.
func (f *Facility) Switchgear() ([]*PowerDistribution, error) {
	return GetObjects[PowerDistribution](f.client, f.switchgear)
}

// TransferSwitches gets the TransferSwitches linked resources.
func (f *Facility) TransferSwitches() ([]*PowerDistribution, error) {
	return GetObjects[PowerDistribution](f.client, f.transferSwitches)
}

// AmbientMetrics gets the AmbientMetrics linked resource.
func (f *Facility) AmbientMetrics() (*EnvironmentMetrics, error) {
	if f.ambientMetrics == "" {
		return nil, nil
	}
	return GetObject[EnvironmentMetrics](f.client, f.ambientMetrics)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (f *Facility) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if f.environmentMetrics == "" {
		return nil, nil
	}
	return GetObject[EnvironmentMetrics](f.client, f.environmentMetrics)
}

// PowerDomains gets the PowerDomains collection.
func (f *Facility) PowerDomains() ([]*PowerDomain, error) {
	if f.powerDomains == "" {
		return nil, nil
	}
	return GetCollectionObjects[PowerDomain](f.client, f.powerDomains)
}
