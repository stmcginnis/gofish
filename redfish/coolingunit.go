//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

type CoolingEquipmentType string

const (
	// CDUCoolingEquipmentType A coolant distribution unit (CDU).
	CDUCoolingEquipmentType CoolingEquipmentType = "CDU"
	// HeatExchangerCoolingEquipmentType A heat exchanger.
	HeatExchangerCoolingEquipmentType CoolingEquipmentType = "HeatExchanger"
	// ImmersionUnitCoolingEquipmentType An immersion cooling unit.
	ImmersionUnitCoolingEquipmentType CoolingEquipmentType = "ImmersionUnit"
)

// CoolingUnit shall represent a cooling system component or unit for a Redfish implementation.
type CoolingUnit struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AllowedCoolingUnitModes shall contain the allowed values for setting the mode of this cooling unit.
	AllowedCoolingUnitModes []CoolingUnitMode
	// Assembly shall contain a link to a resource of type Assembly.
	assembly string
	// AssetTag shall contain the user-assigned asset tag, which is an identifying string that tracks the equipment for
	// inventory purposes.
	AssetTag string
	// Coolant shall contain details regarding the coolant contained or used by this unit.
	Coolant Coolant
	// CoolantConnectorRedundancy shall contain redundancy information for the set of coolant connectors attached to
	// this equipment. The values of the RedundancyGroup array shall reference resources of type CoolantConnector.
	coolantConnectorRedundancy []string
	// CoolingCapacityWatts shall contain the manufacturer-provided cooling capacity, in watt units, of this equipment.
	CoolingCapacityWatts int
	// Description provides a description of this resource.
	Description string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this equipment.
	environmentMetrics string
	// EquipmentType shall contain the type of equipment this resource represents.
	EquipmentType CoolingEquipmentType
	// FilterRedundancy shall contain redundancy information for the groups of filters in this unit.
	FilterRedundancy []RedundantGroup
	// Filters shall contain a link to a resource collection of type FilterCollection that contains the filter
	// information for this equipment.
	filters string
	// FirmwareVersion shall contain a string describing the firmware version of this equipment as provided by the
	// manufacturer.
	FirmwareVersion string
	// LeakDetection shall contain a link to a resource of type LeakDetection that contains the leak detection
	// component information for this equipment.
	leakDetection string
	// Location shall contain the location information of the associated equipment.
	Location common.Location
	// Manufacturer shall contain the name of the organization responsible for producing the equipment. This
	// organization may be the entity from which the equipment is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the manufacturer-provided model information of this equipment.
	Model string
	// PartNumber shall contain the manufacturer-provided part number for the equipment.
	PartNumber string
	// PrimaryCoolantConnectors shall contain a link to a resource collection of type CoolantConnectorCollection that
	// contains the primary coolant connectors for this equipment.
	primaryCoolantConnectors string
	// ProductionDate shall contain the date of production or manufacture for this equipment.
	ProductionDate string
	// PumpRedundancy shall contain redundancy information for the groups of pumps in this unit.
	PumpRedundancy []RedundantGroup
	// Pumps shall contain a link to a resource collection of type PumpCollection that contains the pump information
	// for this equipment.
	pumps string
	// Reservoirs shall contain a link to a resource collection of type ReservoirCollection that contains the reservoir
	// information for this equipment.
	reservoirs string
	// SecondaryCoolantConnectors shall contain a link to a resource collection of type CoolantConnectorCollection that
	// contains the secondary coolant connectors for this equipment.
	secondaryCoolantConnectors string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the equipment.
	SerialNumber string
	// setMode shall contain the action target for setting the mode of this cooling unit.
	setMode string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UserLabel shall contain a user-assigned label used to identify this resource. If a value has not been assigned
	// by a user, the value of this property shall be an empty string.
	UserLabel string
	// Version shall contain the hardware version of this equipment as determined by the vendor or supplier.
	Version string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
	chassis []string
	// ChassisCount is the number of physical containers that contain this equipment.
	ChassisCount int
	facility     string
	managedBy    []string
	// ManagedByCount is the number of managers that manage this equipment.
	ManagedByCount int
}

type CoolingUnitMode string

const (
	EnabledCoolingUnitMode  CoolingUnitMode = "Enabled"
	DisabledCoolingUnitMode CoolingUnitMode = "Disabled"
)

// UnmarshalJSON unmarshals a CoolingUnit object from the raw JSON.
func (coolingunit *CoolingUnit) UnmarshalJSON(b []byte) error {
	type temp CoolingUnit
	type Links struct {
		// Chassis shall contain an array of links to resources of type Chassis that represent the physical containers that
		// contain this equipment.
		Chassis common.Links
		// Chassis@odata.count
		ChassisCount int `json:"Chassis@odata.count"`
		// Facility shall contain a link to a resource of type Facility that represents the facility that contains this
		// equipment.
		Facility common.Link
		// ManagedBy shall contain an array of links to resources of type Manager that represent the managers that manage
		// this equipment.
		ManagedBy common.Links
		// ManagedBy@odata.count
		ManagedByCount int `json:"ManagedBy@odata.count"`
	}
	type CoolingUnitActions struct {
		SetMode struct {
			AllowedCoolingUnitModes []CoolingUnitMode `json:"Mode@Redfish.AllowableValues"`
			Target                  string
		} `json:"#CoolingUnit.SetMode"`
	}

	var t struct {
		temp
		Actions                    CoolingUnitActions
		Assembly                   common.Link
		CoolantConnectorRedundancy common.Links
		EnvironmentMetrics         common.Link
		Filters                    common.Link
		LeakDetection              common.Link
		PrimaryCoolantConnectors   common.Link
		Pumps                      common.Link
		Reservoirs                 common.Link
		SecondaryCoolantConnectors common.Link
		Links                      Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*coolingunit = CoolingUnit(t.temp)

	// Extract the links to other entities for later
	coolingunit.assembly = t.Assembly.String()
	coolingunit.coolantConnectorRedundancy = t.CoolantConnectorRedundancy.ToStrings()
	coolingunit.environmentMetrics = t.EnvironmentMetrics.String()
	coolingunit.filters = t.Filters.String()
	coolingunit.leakDetection = t.LeakDetection.String()
	coolingunit.primaryCoolantConnectors = t.PrimaryCoolantConnectors.String()
	coolingunit.pumps = t.Pumps.String()
	coolingunit.reservoirs = t.Reservoirs.String()
	coolingunit.secondaryCoolantConnectors = t.SecondaryCoolantConnectors.String()
	coolingunit.chassis = t.Links.Chassis.ToStrings()
	coolingunit.ChassisCount = t.Links.ChassisCount
	coolingunit.facility = t.Links.Facility.String()
	coolingunit.managedBy = t.Links.ManagedBy.ToStrings()
	coolingunit.ManagedByCount = t.Links.ManagedByCount
	coolingunit.AllowedCoolingUnitModes = t.Actions.SetMode.AllowedCoolingUnitModes
	coolingunit.setMode = t.Actions.SetMode.Target

	// This is a read/write object, so we need to save the raw object data for later
	coolingunit.rawData = b

	return nil
}

func (coolingunit *CoolingUnit) SetMode(mode CoolingUnitMode) error {
	return coolingunit.SetModeWithContext(common.ContextOf(coolingunit.GetClient()), mode)
}

func (coolingunit *CoolingUnit) SetModeWithContext(ctx context.Context, mode CoolingUnitMode) error {
	// TODO: check if mode is in Allowable values
	properties := map[string]interface{}{
		"Mode": mode,
	}

	return coolingunit.PostWithContext(ctx, coolingunit.setMode, properties)
}

// Update commits updates to this object's properties to the running system.
func (coolingunit *CoolingUnit) Update() error {
	return coolingunit.UpdateWithContext(common.ContextOf(coolingunit.GetClient()))
}

// UpdateWithContext commits updates to this object's properties to the running system.
func (coolingunit *CoolingUnit) UpdateWithContext(ctx context.Context) error {
	readWriteFields := []string{
		"AssetTag",
		"UserLabel",
	}

	return coolingunit.UpdateFromRawDataWithContext(ctx, coolingunit, coolingunit.rawData, readWriteFields)
}

// GetCoolingUnit will get a CoolingUnit instance from the service.
func GetCoolingUnit(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*CoolingUnit, error) {
	return GetCoolingUnitWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetCoolingUnitWithContext will get a CoolingUnit instance from the service.
func GetCoolingUnitWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*CoolingUnit, error) {
	return common.GetObjectWithContext[CoolingUnit](ctx, c, uri, queryOpts...)
}

// ListReferencedCoolingUnits gets the collection of CoolingUnit from
// a provided reference.
func ListReferencedCoolingUnits(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*CoolingUnit, error) {
	return ListReferencedCoolingUnitsWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedCoolingUnitsWithContext gets the collection of CoolingUnit from
// a provided reference.
func ListReferencedCoolingUnitsWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*CoolingUnit, error) {
	return common.GetCollectionObjectsWithContext[CoolingUnit](ctx, c, link, queryOpts...)
}

// Assembly gets a collection of assemblies.
func (coolingunit *CoolingUnit) Assembly(queryOpts ...common.QueryGroupOption) ([]*Assembly, error) {
	return coolingunit.AssemblyWithContext(common.ContextOf(coolingunit.GetClient()), queryOpts...)
}

// AssemblyWithContext gets a collection of assemblies.
func (coolingunit *CoolingUnit) AssemblyWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Assembly, error) {
	return ListReferencedAssemblysWithContext(ctx, coolingunit.GetClient(), coolingunit.assembly, queryOpts...)
}

// EnvironmentMetrics gets the environment metrics for this cooling unit.
func (coolingunit *CoolingUnit) EnvironmentMetrics(queryOpts ...common.QueryGroupOption) (*EnvironmentMetrics, error) {
	return coolingunit.EnvironmentMetricsWithContext(common.ContextOf(coolingunit.GetClient()), queryOpts...)
}

// EnvironmentMetricsWithContext gets the environment metrics for this cooling unit.
func (coolingunit *CoolingUnit) EnvironmentMetricsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*EnvironmentMetrics, error) {
	if coolingunit.environmentMetrics == "" {
		return nil, nil
	}
	return GetEnvironmentMetricsWithContext(ctx, coolingunit.GetClient(), coolingunit.environmentMetrics, queryOpts...)
}

// Filters gets a collection of filters.
func (coolingunit *CoolingUnit) Filters(queryOpts ...common.QueryGroupOption) ([]*Filter, error) {
	return coolingunit.FiltersWithContext(common.ContextOf(coolingunit.GetClient()), queryOpts...)
}

// FiltersWithContext gets a collection of filters.
func (coolingunit *CoolingUnit) FiltersWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Filter, error) {
	return ListReferencedFiltersWithContext(ctx, coolingunit.GetClient(), coolingunit.filters, queryOpts...)
}

// LeakDetection gets the of leak detection of this cooling unit.
func (coolingunit *CoolingUnit) LeakDetection(queryOpts ...common.QueryGroupOption) (*LeakDetection, error) {
	return coolingunit.LeakDetectionWithContext(common.ContextOf(coolingunit.GetClient()), queryOpts...)
}

// LeakDetectionWithContext gets the of leak detection of this cooling unit.
func (coolingunit *CoolingUnit) LeakDetectionWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*LeakDetection, error) {
	if coolingunit.leakDetection == "" {
		return nil, nil
	}
	return GetLeakDetectionWithContext(ctx, coolingunit.GetClient(), coolingunit.leakDetection, queryOpts...)
}

// PrimaryCoolantConnectors gets a collection of primary coolant connectors.
func (coolingunit *CoolingUnit) PrimaryCoolantConnectors(queryOpts ...common.QueryGroupOption) ([]*CoolantConnector, error) {
	return coolingunit.PrimaryCoolantConnectorsWithContext(common.ContextOf(coolingunit.GetClient()), queryOpts...)
}

// PrimaryCoolantConnectorsWithContext gets a collection of primary coolant connectors.
func (coolingunit *CoolingUnit) PrimaryCoolantConnectorsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*CoolantConnector, error) {
	return ListReferencedCoolantConnectorsWithContext(ctx, coolingunit.GetClient(), coolingunit.primaryCoolantConnectors, queryOpts...)
}

// Pumps gets a collection of pumps.
func (coolingunit *CoolingUnit) Pumps(queryOpts ...common.QueryGroupOption) ([]*Pump, error) {
	return coolingunit.PumpsWithContext(common.ContextOf(coolingunit.GetClient()), queryOpts...)
}

// PumpsWithContext gets a collection of pumps.
func (coolingunit *CoolingUnit) PumpsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Pump, error) {
	return ListReferencedPumpsWithContext(ctx, coolingunit.GetClient(), coolingunit.pumps, queryOpts...)
}

// Reservoirs gets a collection of reservoirs.
func (coolingunit *CoolingUnit) Reservoirs(queryOpts ...common.QueryGroupOption) ([]*Reservoir, error) {
	return coolingunit.ReservoirsWithContext(common.ContextOf(coolingunit.GetClient()), queryOpts...)
}

// ReservoirsWithContext gets a collection of reservoirs.
func (coolingunit *CoolingUnit) ReservoirsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Reservoir, error) {
	return ListReferencedReservoirsWithContext(ctx, coolingunit.GetClient(), coolingunit.reservoirs, queryOpts...)
}

// SecondaryCoolantConnectors gets a collection of secondary coolant connectors.
func (coolingunit *CoolingUnit) SecondaryCoolantConnectors(queryOpts ...common.QueryGroupOption) ([]*CoolantConnector, error) {
	return coolingunit.SecondaryCoolantConnectorsWithContext(common.ContextOf(coolingunit.GetClient()), queryOpts...)
}

// SecondaryCoolantConnectorsWithContext gets a collection of secondary coolant connectors.
func (coolingunit *CoolingUnit) SecondaryCoolantConnectorsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*CoolantConnector, error) {
	return ListReferencedCoolantConnectorsWithContext(ctx, coolingunit.GetClient(), coolingunit.secondaryCoolantConnectors, queryOpts...)
}
