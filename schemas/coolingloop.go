//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #CoolingLoop.v1_1_0.CoolingLoop

package schemas

import (
	"encoding/json"
)

type CoolantType string

const (
	// WaterCoolantType Water or glycol mixture, including additives.
	WaterCoolantType CoolantType = "Water"
	// HydrocarbonCoolantType Hydrocarbon-based.
	HydrocarbonCoolantType CoolantType = "Hydrocarbon"
	// FluorocarbonCoolantType Fluorocarbon-based.
	FluorocarbonCoolantType CoolantType = "Fluorocarbon"
	// DielectricCoolantType Dielectric fluid.
	DielectricCoolantType CoolantType = "Dielectric"
)

type CoolingLoopType string

const (
	// FWSCoolingLoopType Facility Water System (FWS).
	FWSCoolingLoopType CoolingLoopType = "FWS"
	// TCSCoolingLoopType Technology Cooling System (TCS).
	TCSCoolingLoopType CoolingLoopType = "TCS"
	// RowTCSCoolingLoopType is a loop connecting to one or more racks or similar
	// scope. May connect to multiple TCS loops.
	RowTCSCoolingLoopType CoolingLoopType = "RowTCS"
)

// CoolingLoop shall represent a cooling loop for a Redfish implementation.
type CoolingLoop struct {
	Entity
	// ConsumingEquipmentNames shall contain an array of user-assigned identifying
	// strings that describe downstream devices that receive coolant from this
	// cooling loop.
	ConsumingEquipmentNames []string
	// Coolant shall contain the details about the coolant contained in this
	// cooling loop.
	Coolant Coolant
	// CoolantLevelPercent shall contain the amount of coolant capacity, in percent
	// units, filled in this cooling loop. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'Percent'. Services that support
	// this property shall also return the 'CoolantLevelStatus' property.
	CoolantLevelPercent SensorExcerpt
	// CoolantLevelStatus shall indicate the status of the coolant level in this
	// cooling loop.
	CoolantLevelStatus Health
	// CoolantQuality shall indicate the quality of the coolant contained in this
	// cooling loop.
	CoolantQuality Health
	// CoolingLoopType shall contain the type of cooling loop represented by this
	// resource.
	//
	// Version added: v1.1.0
	CoolingLoopType CoolingLoopType
	// CoolingManagerURI shall contain a URI to the application or device that
	// provides administration or management of the cooling loop associated with
	// this interface.
	CoolingManagerURI string
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource.
	LocationIndicatorActive bool
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PrimaryCoolantConnectors shall contain a link to a resource collection of
	// type 'CoolantConnectorCollection' that contains the primary coolant
	// connectors for this equipment.
	primaryCoolantConnectors string
	// RatedFlowLitersPerMinute shall contain the rated liquid flow, in liters per
	// minute units, for this cooling loop.
	RatedFlowLitersPerMinute *float64 `json:",omitempty"`
	// RatedPressurekPa shall contain the rated maximum pressure, in kilopascal
	// units, for this cooling loop.
	RatedPressurekPa *float64 `json:",omitempty"`
	// SecondaryCoolantConnectors shall contain a link to a resource collection of
	// type 'CoolantConnectorCollection' that contains the secondary coolant
	// connectors for this equipment.
	secondaryCoolantConnectors string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SupplyEquipmentNames shall contain an array of user-assigned identifying
	// strings that describe upstream devices that supply coolant to this cooling
	// loop.
	SupplyEquipmentNames []string
	// UserLabel shall contain a user-assigned label used to identify this
	// resource. If a value has not been assigned by a user, the value of this
	// property shall be an empty string.
	UserLabel string
	// chassis is the URI for Chassis.
	chassis string
	// facility is the URI for Facility.
	facility string
	// managedBy are the URIs for ManagedBy.
	managedBy []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a CoolingLoop object from the raw JSON.
func (c *CoolingLoop) UnmarshalJSON(b []byte) error {
	type temp CoolingLoop
	type cLinks struct {
		Chassis   Link  `json:"Chassis"`
		Facility  Link  `json:"Facility"`
		ManagedBy Links `json:"ManagedBy"`
	}
	var tmp struct {
		temp
		Links                      cLinks
		PrimaryCoolantConnectors   Link `json:"PrimaryCoolantConnectors"`
		SecondaryCoolantConnectors Link `json:"SecondaryCoolantConnectors"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CoolingLoop(tmp.temp)

	// Extract the links to other entities for later
	c.chassis = tmp.Links.Chassis.String()
	c.facility = tmp.Links.Facility.String()
	c.managedBy = tmp.Links.ManagedBy.ToStrings()
	c.primaryCoolantConnectors = tmp.PrimaryCoolantConnectors.String()
	c.secondaryCoolantConnectors = tmp.SecondaryCoolantConnectors.String()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *CoolingLoop) Update() error {
	readWriteFields := []string{
		"ConsumingEquipmentNames",
		"CoolingLoopType",
		"CoolingManagerURI",
		"LocationIndicatorActive",
		"SupplyEquipmentNames",
		"UserLabel",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetCoolingLoop will get a CoolingLoop instance from the service.
func GetCoolingLoop(c Client, uri string) (*CoolingLoop, error) {
	return GetObject[CoolingLoop](c, uri)
}

// ListReferencedCoolingLoops gets the collection of CoolingLoop from
// a provided reference.
func ListReferencedCoolingLoops(c Client, link string) ([]*CoolingLoop, error) {
	return GetCollectionObjects[CoolingLoop](c, link)
}

// Chassis gets the Chassis linked resource.
func (c *CoolingLoop) Chassis() (*Chassis, error) {
	if c.chassis == "" {
		return nil, nil
	}
	return GetObject[Chassis](c.client, c.chassis)
}

// Facility gets the Facility linked resource.
func (c *CoolingLoop) Facility() (*Facility, error) {
	if c.facility == "" {
		return nil, nil
	}
	return GetObject[Facility](c.client, c.facility)
}

// ManagedBy gets the ManagedBy linked resources.
func (c *CoolingLoop) ManagedBy() ([]*Manager, error) {
	return GetObjects[Manager](c.client, c.managedBy)
}

// PrimaryCoolantConnectors gets the PrimaryCoolantConnectors collection.
func (c *CoolingLoop) PrimaryCoolantConnectors() ([]*CoolantConnector, error) {
	if c.primaryCoolantConnectors == "" {
		return nil, nil
	}
	return GetCollectionObjects[CoolantConnector](c.client, c.primaryCoolantConnectors)
}

// SecondaryCoolantConnectors gets the SecondaryCoolantConnectors collection.
func (c *CoolingLoop) SecondaryCoolantConnectors() ([]*CoolantConnector, error) {
	if c.secondaryCoolantConnectors == "" {
		return nil, nil
	}
	return GetCollectionObjects[CoolantConnector](c.client, c.secondaryCoolantConnectors)
}

// Coolant shall describe the coolant used with a device.
type Coolant struct {
	// AdditiveName shall contain the name of the additive contained in the
	// coolant.
	AdditiveName string
	// AdditivePercent shall contain the percent of additives, '0' to '100', by
	// volume, contained in the coolant mixture.
	AdditivePercent *float64 `json:",omitempty"`
	// CoolantType shall contain the type of coolant used by this resource.
	CoolantType CoolantType
	// DensityKgPerCubicMeter shall contain the density of the coolant, in
	// kilograms per cubic meter units, as measured at room temperature (20-25
	// degrees C) and atmospheric pressure.
	DensityKgPerCubicMeter *float64 `json:",omitempty"`
	// RatedServiceHours shall contain the number of hours of service that the
	// coolant is rated to provide before servicing or replacement is necessary.
	RatedServiceHours *float64 `json:",omitempty"`
	// ServiceHours shall contain the number of hours of service that the coolant
	// has provided.
	ServiceHours *float64 `json:",omitempty"`
	// ServicedDate shall contain the date the coolant was last serviced or tested
	// for quality.
	ServicedDate string
	// SpecificHeatkJoulesPerKgK shall contain the specific heat capacity of the
	// coolant, in kilojoules per kilogram per degree kelvin units, as measured at
	// room temperature (20-25 degrees C) and atmospheric pressure.
	SpecificHeatkJoulesPerKgK *float64 `json:",omitempty"`
}
