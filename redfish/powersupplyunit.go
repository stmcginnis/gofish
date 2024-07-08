//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type LineInputStatus string

const (
	// No power detected at line input.
	LossOfInputLineInputStatus LineInputStatus = "LossOfInput"
	// Line input is within normal operating range.
	NormalLineInputStatus LineInputStatus = "Normal"
	// Line input voltage or current is outside of normal operating range.
	OutOfRangeLineInputStatus LineInputStatus = "OutOfRange"
)

type PowerSupplyUnitType string

const (
	// Alternating Current (AC) power supply.
	ACPowerSupplyUnitType PowerSupplyUnitType = "AC"
	// The power supply supports both DC or AC.
	ACorDCPowerSupplyUnitType PowerSupplyUnitType = "ACorDC"
	// Direct Current (DC) power supply.
	DCPowerSupplyUnitType PowerSupplyUnitType = "DC"
	// (v1.5+)	Direct Current (DC) voltage regulator.
	DCRegulatorPowerSupplyUnitType PowerSupplyUnitType = "DCRegulator"
)

// The efficiency ratings of this power supply.
type EfficiencyRating struct {
	// The rated efficiency of this power supply at the specified load.
	EfficiencyPercent float32
	// The electrical load for this rating.
	LoadPercent float32
}

// The input ranges that the power supply can use.
type PowerSupplyInputRange struct {
	// The maximum capacity of this power supply when operating in this input range.
	CapacityWatts float32
	// The input voltage range.
	NominalVoltageType NominalVoltage
}

// The input ranges that the power supply can use.
type OutputRail struct {
	// The nominal voltage of this output power rail.
	NominalVoltage float32
	// The area or device to which this power rail applies.
	PhysicalContext common.PhysicalContext
}

// PowerSupplyUnit shall represent a power supply unit for a Redfish implementation.
// It may also represent a location, such as a slot, socket, or bay, where a unit may be installed,
// but the State property within the Status property contains Absent.
type PowerSupplyUnit struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string

	// The link to the assembly associated with this power supply.
	assembly string
	// The efficiency ratings of this power supply.
	EfficiencyRatings []EfficiencyRating
	// The URIs of the management interfaces
	// for the upstream electrical source connections for this power supply.
	ElectricalSourceManagerURIs []string
	// The names of the upstream electrical sources,
	// such as circuits or outlets, connected to this power supply.
	ElectricalSourceNames []string
	// The firmware version for this power supply.
	FirmwareVersion string
	// An indication of whether this device can be inserted
	// or removed while the equipment is in operation.
	HotPluggable bool
	// The nominal voltage type of the line input to this power supply.
	InputNominalVoltageType NominalVoltage
	// The input ranges that the power supply can use.
	InputRanges []PowerSupplyInputRange
	// The status of the line input.
	LineInputStatus LineInputStatus
	// The location of the power supply.
	Location common.Location
	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool
	// The manufacturer of this power supply.
	Manufacturer string
	// The link to the power supply metrics resource associated with this power supply.
	metrics string
	// The model number for this power supply.
	Model string
	// The nominal output voltage type of this power supply.
	OutputNominalVoltageType NominalVoltage
	// The output power rails provided by this power supply.
	OutputRails []OutputRail
	// The part number for this power supply.
	PartNumber string
	// The number of ungrounded current-carrying conductors (phases)
	// and the total number of conductors (wires)
	// provided for the power supply input connector.
	PhaseWiringType PhaseWiringType
	// The type of plug according to NEMA, IEC, or regional standards.
	PlugType PlugType
	// The maximum capacity of this power supply.
	PowerCapacityWatts float32
	// The power supply type (AC or DC).
	PowerSupplyType PowerSupplyUnitType
	// The production or manufacturing date of this power supply.
	ProductionDate string
	// An indication of whether this component can be independently replaced
	// as allowed by the vendor's replacement policy.
	Replaceable bool
	// The serial number for this power supply.
	SerialNumber string
	// The spare part number for this power supply.
	SparePartNumber string
	// The status and health of the resource and its subordinate or dependent resources.
	Status common.Status
	// The hardware version of this power supply.
	Version string

	// Links section
	// Deprecated (v1.4): A link to the outlet connected to this power supply.
	outlet          string
	poweringChassis []string
	// PoweringChassisCount is the number of chassis that are directly powered by this power supply.
	PoweringChassisCount int
	powerOutlets         []string
	// PowerOutletsCount is the number of outlets that provide power to this power supply.
	PowerOutletsCount int
	// OemLinks are all OEM data under link section
	OemLinks json.RawMessage

	// Actions section
	// This action resets the power supply.
	resetTarget string
	// OemActions contains all the vendor specific actions.
	// It is vendor responsibility to parse this field accordingly
	OemActions json.RawMessage

	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PowerSupplyUnit object from the raw JSON.
func (powerSupplyUnit *PowerSupplyUnit) UnmarshalJSON(b []byte) error {
	type temp PowerSupplyUnit
	type linkReference struct {
		Outlet               common.Link
		PoweringChassis      common.Links
		PoweringChassisCount int `json:"PoweringChassis@odata.count"`
		PowerOutlets         common.Links
		PowerOutletsCount    int `json:"PowerOutlets@odata.count"`
		Oem                  json.RawMessage
	}
	type actions struct {
		Reset common.ActionTarget `json:"#PowerSupply.Reset"`
		Oem   json.RawMessage     // OEM actions will be stored here
	}
	var t struct {
		temp

		Assembly common.Link
		Metrics  common.Link

		Links   linkReference
		Actions actions
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	// Extract the links to other entities for later
	*powerSupplyUnit = PowerSupplyUnit(t.temp)
	powerSupplyUnit.assembly = t.Assembly.String()
	powerSupplyUnit.metrics = t.Metrics.String()

	powerSupplyUnit.outlet = t.Links.Outlet.String()
	powerSupplyUnit.poweringChassis = t.Links.PoweringChassis.ToStrings()
	powerSupplyUnit.PoweringChassisCount = t.Links.PoweringChassisCount
	powerSupplyUnit.powerOutlets = t.Links.PowerOutlets.ToStrings()
	powerSupplyUnit.PowerOutletsCount = t.Links.PowerOutletsCount
	powerSupplyUnit.OemLinks = t.Links.Oem

	powerSupplyUnit.resetTarget = t.Actions.Reset.Target
	powerSupplyUnit.OemActions = t.Actions.Oem

	// This is a read/write object, so we need to save the raw object data for later
	powerSupplyUnit.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (powerSupplyUnit *PowerSupplyUnit) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	psu := new(PowerSupplyUnit)
	err := psu.UnmarshalJSON(powerSupplyUnit.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"ElectricalSourceManagerURIs",
		"ElectricalSourceNames",
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(psu).Elem()
	currentElement := reflect.ValueOf(powerSupplyUnit).Elem()

	return powerSupplyUnit.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetPowerSupplyUnit will get a PowerSupplyUnit instance from the Redfish service.
func GetPowerSupplyUnit(c common.Client, uri string) (*PowerSupplyUnit, error) {
	return common.GetObject[PowerSupplyUnit](c, uri)
}

// ListReferencedPowerSupplyUnits gets the collection of PowerSupplies from
// a provided reference.
func ListReferencedPowerSupplyUnits(c common.Client, link string) ([]*PowerSupplyUnit, error) {
	return common.GetCollectionObjects[PowerSupplyUnit](c, link)
}

// This action shall reset a power supply. A GracefulRestart ResetType shall reset the power supply
// but shall not affect the power output. A ForceRestart ResetType can affect the power supply output.
func (powerSupplyUnit *PowerSupplyUnit) Reset(resetType ResetType) error {
	if powerSupplyUnit.resetTarget == "" {
		return errors.New("Reset is not supported") //nolint:golint
	}

	t := struct {
		ResetType ResetType
	}{ResetType: resetType}

	return powerSupplyUnit.Post(powerSupplyUnit.resetTarget, t)
}

// Assembly gets the containing assembly for this power supply.
func (powerSupplyUnit *PowerSupplyUnit) Assembly() (*Assembly, error) {
	if powerSupplyUnit.assembly == "" {
		return nil, nil
	}
	return GetAssembly(powerSupplyUnit.GetClient(), powerSupplyUnit.assembly)
}

// Metrics gets the metrics associated with this power supply.
func (powerSupplyUnit *PowerSupplyUnit) Metrics() (*PowerSupplyUnitMetrics, error) {
	if powerSupplyUnit.metrics == "" {
		return nil, nil
	}
	return GetPowerSupplyUnitMetrics(powerSupplyUnit.GetClient(), powerSupplyUnit.metrics)
}

// Outlet get the outlet connected to this power supply.
// Deprecated (v1.4)
func (powerSupplyUnit *PowerSupplyUnit) Outlet() (*Outlet, error) {
	if powerSupplyUnit.metrics == "" {
		return nil, nil
	}
	return GetOutlet(powerSupplyUnit.GetClient(), powerSupplyUnit.outlet)
}

// PowerOutlets gets the outlets that supply power to this power supply.
func (powerSupplyUnit *PowerSupplyUnit) PowerOutlets() ([]*Outlet, error) {
	return common.GetObjects[Outlet](powerSupplyUnit.GetClient(), powerSupplyUnit.powerOutlets)
}

// PoweringChassis gets the collection of the chassis directly powered by this power supply.
func (powerSupplyUnit *PowerSupplyUnit) PoweringChassis() ([]*Chassis, error) {
	return common.GetObjects[Chassis](powerSupplyUnit.GetClient(), powerSupplyUnit.poweringChassis)
}
