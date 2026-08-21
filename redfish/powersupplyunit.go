//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/coreweave/gofish/common"
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
	LocationIndicatorActive *bool
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

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
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
	powerSupplyUnit.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (powerSupplyUnit *PowerSupplyUnit) Update() error {
	return powerSupplyUnit.UpdateWithContext(common.ContextOf(powerSupplyUnit.GetClient()))
}

// UpdateWithContext commits updates to this object's properties to the running system.
func (powerSupplyUnit *PowerSupplyUnit) UpdateWithContext(ctx context.Context) error {
	readWriteFields := []string{
		"ElectricalSourceManagerURIs",
		"ElectricalSourceNames",
		"LocationIndicatorActive",
	}

	return powerSupplyUnit.UpdateFromRawDataWithContext(ctx, powerSupplyUnit, powerSupplyUnit.RawData, readWriteFields)
}

// GetPowerSupplyUnit will get a PowerSupplyUnit instance from the Redfish service.
func GetPowerSupplyUnit(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*PowerSupplyUnit, error) {
	return GetPowerSupplyUnitWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetPowerSupplyUnitWithContext will get a PowerSupplyUnit instance from the Redfish service.
func GetPowerSupplyUnitWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*PowerSupplyUnit, error) {
	return common.GetObjectWithContext[PowerSupplyUnit](ctx, c, uri, queryOpts...)
}

// ListReferencedPowerSupplyUnits gets the collection of PowerSupplies from
// a provided reference.
func ListReferencedPowerSupplyUnits(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*PowerSupplyUnit, error) {
	return ListReferencedPowerSupplyUnitsWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedPowerSupplyUnitsWithContext gets the collection of PowerSupplies from
// a provided reference.
func ListReferencedPowerSupplyUnitsWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*PowerSupplyUnit, error) {
	return common.GetCollectionObjectsWithContext[PowerSupplyUnit](ctx, c, link, queryOpts...)
}

// This action shall reset a power supply. A GracefulRestart ResetType shall reset the power supply
// but shall not affect the power output. A ForceRestart ResetType can affect the power supply output.
func (powerSupplyUnit *PowerSupplyUnit) Reset(resetType ResetType) error {
	return powerSupplyUnit.ResetWithContext(common.ContextOf(powerSupplyUnit.GetClient()), resetType)
}

// This action shall reset a power supply. A GracefulRestart ResetType shall reset the power supply
// but shall not affect the power output. A ForceRestart ResetType can affect the power supply output.
func (powerSupplyUnit *PowerSupplyUnit) ResetWithContext(ctx context.Context, resetType ResetType) error {
	if powerSupplyUnit.resetTarget == "" {
		return errors.New("Reset is not supported")
	}

	t := struct {
		ResetType ResetType
	}{ResetType: resetType}

	return powerSupplyUnit.PostWithContext(ctx, powerSupplyUnit.resetTarget, t)
}

// Assembly gets the containing assembly for this power supply.
func (powerSupplyUnit *PowerSupplyUnit) Assembly(queryOpts ...common.QueryGroupOption) (*Assembly, error) {
	return powerSupplyUnit.AssemblyWithContext(common.ContextOf(powerSupplyUnit.GetClient()), queryOpts...)
}

// AssemblyWithContext gets the containing assembly for this power supply.
func (powerSupplyUnit *PowerSupplyUnit) AssemblyWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*Assembly, error) {
	if powerSupplyUnit.assembly == "" {
		return nil, nil
	}
	return GetAssemblyWithContext(ctx, powerSupplyUnit.GetClient(), powerSupplyUnit.assembly, queryOpts...)
}

// Metrics gets the metrics associated with this power supply.
func (powerSupplyUnit *PowerSupplyUnit) Metrics(queryOpts ...common.QueryGroupOption) (*PowerSupplyUnitMetrics, error) {
	return powerSupplyUnit.MetricsWithContext(common.ContextOf(powerSupplyUnit.GetClient()), queryOpts...)
}

// MetricsWithContext gets the metrics associated with this power supply.
func (powerSupplyUnit *PowerSupplyUnit) MetricsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*PowerSupplyUnitMetrics, error) {
	if powerSupplyUnit.metrics == "" {
		return nil, nil
	}
	return GetPowerSupplyUnitMetricsWithContext(ctx, powerSupplyUnit.GetClient(), powerSupplyUnit.metrics, queryOpts...)
}

// Outlet get the outlet connected to this power supply.
// Deprecated (v1.4)
func (powerSupplyUnit *PowerSupplyUnit) Outlet(queryOpts ...common.QueryGroupOption) (*Outlet, error) {
	return powerSupplyUnit.OutletWithContext(common.ContextOf(powerSupplyUnit.GetClient()), queryOpts...)
}

// OutletWithContext get the outlet connected to this power supply.
// Deprecated (v1.4)
func (powerSupplyUnit *PowerSupplyUnit) OutletWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*Outlet, error) {
	if powerSupplyUnit.metrics == "" {
		return nil, nil
	}
	return GetOutletWithContext(ctx, powerSupplyUnit.GetClient(), powerSupplyUnit.outlet, queryOpts...)
}

// PowerOutlets gets the outlets that supply power to this power supply.
func (powerSupplyUnit *PowerSupplyUnit) PowerOutlets(queryOpts ...common.QueryGroupOption) ([]*Outlet, error) {
	return powerSupplyUnit.PowerOutletsWithContext(common.ContextOf(powerSupplyUnit.GetClient()), queryOpts...)
}

// PowerOutletsWithContext gets the outlets that supply power to this power supply.
func (powerSupplyUnit *PowerSupplyUnit) PowerOutletsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Outlet, error) {
	return common.GetObjectsWithContext[Outlet](ctx, powerSupplyUnit.GetClient(), powerSupplyUnit.powerOutlets, queryOpts...)
}

// PoweringChassis gets the collection of the chassis directly powered by this power supply.
func (powerSupplyUnit *PowerSupplyUnit) PoweringChassis(queryOpts ...common.QueryGroupOption) ([]*Chassis, error) {
	return powerSupplyUnit.PoweringChassisWithContext(common.ContextOf(powerSupplyUnit.GetClient()), queryOpts...)
}

// PoweringChassisWithContext gets the collection of the chassis directly powered by this power supply.
func (powerSupplyUnit *PowerSupplyUnit) PoweringChassisWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Chassis, error) {
	return common.GetObjectsWithContext[Chassis](ctx, powerSupplyUnit.GetClient(), powerSupplyUnit.poweringChassis, queryOpts...)
}
