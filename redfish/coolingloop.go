//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
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

// Coolant shall describe the coolant used with a device.
type Coolant struct {
	// AdditiveName shall contain the name of the additive contained in the coolant.
	AdditiveName string
	// AdditivePercent shall contain the percent of additives, '0' to '100', by volume, contained in the coolant
	// mixture.
	AdditivePercent float64
	// CoolantType shall contain the type of coolant used by this resource.
	CoolantType CoolantType
	// DensityKgPerCubicMeter shall contain the density of the coolant, in kilograms per cubic meter units, as measured
	// at room temperature (20-25 degrees C) and atmospheric pressure.
	DensityKgPerCubicMeter float64
	// RatedServiceHours shall contain the number of hours of service that the coolant is rated to provide before
	// servicing or replacement is necessary.
	RatedServiceHours float64
	// ServiceHours shall contain the number of hours of service that the coolant has provided.
	ServiceHours float64
	// ServicedDate shall contain the date the coolant was last serviced or tested for quality.
	ServicedDate string
	// SpecificHeatkJoulesPerKgK shall contain the specific heat capacity of the coolant, in kilojoules per kilogram
	// per degree kelvin units, as measured at room temperature (20-25 degrees C) and atmospheric pressure.
	SpecificHeatkJoulesPerKgK float64
}

// CoolingLoop shall represent a cooling loop for a Redfish implementation.
type CoolingLoop struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ConsumingEquipmentNames shall contain an array of user-assigned identifying strings that describe downstream
	// devices that receive coolant from this cooling loop.
	ConsumingEquipmentNames []string
	// Coolant shall contain the details about the coolant contained in this cooling loop.
	Coolant Coolant
	// CoolantLevelPercent shall contain the amount of coolant capacity, in percent units, filled in this cooling loop.
	// The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'Percent'. Services that support this property shall also return the
	// CoolantLevelStatus property.
	CoolantLevelPercent SensorExcerpt
	// CoolantLevelStatus shall indicate the status of the coolant level in this cooling loop.
	CoolantLevelStatus common.Health
	// CoolantQuality shall indicate the quality of the coolant contained in this cooling loop.
	CoolantQuality common.Health
	// CoolingManagerURI shall contain a URI to the application or device that provides administration or management of
	// the cooling loop associated with this interface.
	CoolingManagerURI string
	// Description provides a description of this resource.
	Description string
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// PrimaryCoolantConnectors shall contain a link to a resource collection of type CoolantConnectorCollection that
	// contains the primary coolant connectors for this equipment.
	PrimaryCoolantConnectors string
	// RatedFlowLitersPerMinute shall contain the rated liquid flow, in liters per minute units, for this cooling loop.
	RatedFlowLitersPerMinute float64
	// RatedPressurekPa shall contain the rated maximum pressure, in kilopascal units, for this cooling loop.
	RatedPressurekPa float64
	// SecondaryCoolantConnectors shall contain a link to a resource collection of type CoolantConnectorCollection that
	// contains the secondary coolant connectors for this equipment.
	secondaryCoolantConnectors string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupplyEquipmentNames shall contain an array of user-assigned identifying strings that describe upstream devices
	// that supply coolant to this cooling loop.
	SupplyEquipmentNames []string
	// UserLabel shall contain a user-assigned label used to identify this resource. If a value has not been assigned
	// by a user, the value of this property shall be an empty string.
	UserLabel string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData   []byte
	chassis   string
	facility  string
	managedBy []string
	// ManagedByCount is the number of managers that manage this equipment.
	ManagedByCount int
}

// UnmarshalJSON unmarshals a CoolingLoop object from the raw JSON.
func (coolingloop *CoolingLoop) UnmarshalJSON(b []byte) error {
	type temp CoolingLoop
	type Links struct {
		// Chassis shall contain a link to resources of type Chassis that represent the physical container that contains
		// this resource.
		Chassis common.Link
		// Facility shall contain a link to a resource of type Facility that represents the facility that contains this
		// equipment.
		Facility common.Link
		// ManagedBy shall contain an array of links to resources of type Manager that represent the managers that manage
		// this equipment.
		ManagedBy common.Links
		// ManagedBy@odata.count
		ManagedByCount int `json:"ManagedBy@odata.count"`
	}
	var t struct {
		temp
		SecondaryCoolantConnectors common.Link
		Links                      Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*coolingloop = CoolingLoop(t.temp)

	// Extract the links to other entities for later
	coolingloop.secondaryCoolantConnectors = t.SecondaryCoolantConnectors.String()
	coolingloop.chassis = t.Links.Chassis.String()
	coolingloop.facility = t.Links.Facility.String()
	coolingloop.managedBy = t.Links.ManagedBy.ToStrings()
	coolingloop.ManagedByCount = t.Links.ManagedByCount

	// This is a read/write object, so we need to save the raw object data for later
	coolingloop.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (coolingloop *CoolingLoop) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(CoolingLoop)
	original.UnmarshalJSON(coolingloop.rawData)

	readWriteFields := []string{
		"ConsumingEquipmentNames",
		"CoolingManagerURI",
		"LocationIndicatorActive",
		"SupplyEquipmentNames",
		"UserLabel",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(coolingloop).Elem()

	return coolingloop.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetCoolingLoop will get a CoolingLoop instance from the service.
func GetCoolingLoop(c common.Client, uri string) (*CoolingLoop, error) {
	return common.GetObject[CoolingLoop](c, uri)
}

// ListReferencedCoolingLoops gets the collection of CoolingLoop from
// a provided reference.
func ListReferencedCoolingLoops(c common.Client, link string) ([]*CoolingLoop, error) {
	return common.GetCollectionObjects[CoolingLoop](c, link)
}

// SecondaryCoolantConnectors gets the secondary coolant connectors for this equipment.
func (coolingloop *CoolingLoop) SecondaryCoolantConnectors() ([]*CoolantConnector, error) {
	return ListReferencedCoolantConnectors(coolingloop.GetClient(), coolingloop.secondaryCoolantConnectors)
}

// Chassis gets the physical container that contains this resource.
func (coolingloop *CoolingLoop) Chassis() (*Chassis, error) {
	return GetChassis(coolingloop.GetClient(), coolingloop.chassis)
}

// Facility gets the physical container that contains this resource.
func (coolingloop *CoolingLoop) Facility() (*Facility, error) {
	return GetFacility(coolingloop.GetClient(), coolingloop.chassis)
}

// ManagedBy gets the collection of managers of this equipment.
func (coolingloop *CoolingLoop) ManagedBy() ([]*Manager, error) {
	return common.GetObjects[Manager](coolingloop.GetClient(), coolingloop.managedBy)
}
