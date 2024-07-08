//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type VoltageType string

const (
	// ACVoltageType Alternating Current (AC) outlet.
	ACVoltageType VoltageType = "AC"
	// DCVoltageType Direct Current (DC) outlet.
	DCVoltageType VoltageType = "DC"
)

type ReceptacleType string

const (
	// BS1363TypeGReceptacleType is BS 1363 Type G (250V; 13A).
	BS1363TypeGReceptacleType ReceptacleType = "BS_1363_Type_G"
	// BusConnectionReceptacleType is an Electrical bus connection.
	BusConnectionReceptacleType ReceptacleType = "BusConnection"
	// CEE7TypeEReceptacleType is EE 7/7 Type E (250V; 16A).
	CEE7TypeEReceptacleType ReceptacleType = "CEE_7_Type_E"
	// CEE7TypeFReceptacleType is CEE 7/7 Type F (250V; 16A).
	CEE7TypeFReceptacleType ReceptacleType = "CEE_7_Type_F"
	// IEC60320C13ReceptacleType is IEC C13 (250V; 10A or 15A).
	IEC60320C13ReceptacleType ReceptacleType = "IEC_60320_C13"
	// IEC60320C19ReceptacleType is IEC C19 (250V; 16A or 20A).
	IEC60320C19ReceptacleType ReceptacleType = "IEC_60320_C19"
	// NEMA515RReceptacleType is NEMA 5-15R (120V; 15A).
	NEMA515RReceptacleType ReceptacleType = "NEMA_5_15R"
	// NEMA520RReceptacleType is NEMA 5-20R (120V; 20A).
	NEMA520RReceptacleType ReceptacleType = "NEMA_5_20R"
	// NEMAL520RReceptacleType is NEMA L5-20R (120V; 20A).
	NEMAL520RReceptacleType ReceptacleType = "NEMA_L5_20R"
	// NEMAL530RReceptacleType is NEMA L5-30R (120V; 30A).
	NEMAL530RReceptacleType ReceptacleType = "NEMA_L5_30R"
	// NEMAL620RReceptacleType is NEMA L6-20R (250V; 20A).
	NEMAL620RReceptacleType ReceptacleType = "NEMA_L6_20R"
	// NEMAL630RReceptacleType is NEMA L6-30R (250V; 30A).
	NEMAL630RReceptacleType ReceptacleType = "NEMA_L6_30R"
	// SEV1011TYPE12ReceptacleType is SEV 1011 Type 12 (250V; 10A).
	SEV1011TYPE12ReceptacleType ReceptacleType = "SEV_1011_TYPE_12"
	// SEV1011TYPE23ReceptacleTypeSEV 1011 Type 23 (250V; 16A).
	SEV1011TYPE23ReceptacleType ReceptacleType = "SEV_1011_TYPE_23"
)

// CurrentSensors shall contain properties that describe current sensor readings for an outlet.
type CurrentSensors struct {
	// Line1 shall contain the line current, in ampere units, for L1. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value 'Current'.
	// This property shall not be present if the equipment does not include an L1 measurement.
	Line1 SensorCurrentExcerpt
	// Line2 shall contain the line current, in ampere units, for L2. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value 'Current'.
	// This property shall not be present if the equipment does not include an L2 measurement.
	Line2 SensorCurrentExcerpt
	// Line3 shall contain the line current, in ampere units, for L3. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value 'Current'.
	// This property shall not be present if the equipment does not include an L3 measurement.
	Line3 SensorCurrentExcerpt
	// Neutral shall contain the line current, in ampere units, for the Neutral line. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Current'. This property shall not be present if the equipment does not include a Neutral line
	// measurement.
	Neutral SensorCurrentExcerpt
}

// Outlet shall be used to represent an electrical outlet for a Redfish implementation.
type Outlet struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ConfigurationLocked shall indicate whether modification requests to this resource are not permitted. If 'true',
	// services shall reject modification requests to other properties in this resource.
	ConfigurationLocked bool
	// CurrentAmps shall contain the current, in ampere units, for this single-phase outlet. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Current'. This property shall not appear in resource instances representing poly-phase
	// outlets.
	CurrentAmps SensorCurrentExcerpt
	// Description provides a description of this resource.
	Description string
	// ElectricalConsumerNames shall contain an array of user-assigned identifying strings that describe downstream
	// devices that are powered by this outlet.
	ElectricalConsumerNames []string
	// ElectricalContext shall contain the combination of current-carrying conductors that distribute power.
	ElectricalContext common.ElectricalContext
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for this outlet that represents the 'Total'
	// ElectricalContext sensor when multiple energy sensors exist for this outlet. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// FrequencyHz shall contain the frequency, in hertz units, for this outlet. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Frequency'.
	FrequencyHz SensorExcerpt
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// NominalVoltage shall contain the nominal voltage for this outlet, in volt units.
	NominalVoltage NominalVoltage
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutletType shall contain the type of physical receptacle used for this outlet, as defined by IEC, NEMA, or
	// regional standards.
	OutletType ReceptacleType
	// PhaseWiringType shall contain the number of ungrounded current-carrying conductors (phases) and the total number
	// of conductors (wires).
	PhaseWiringType PhaseWiringType
	// PolyPhaseCurrentAmps shall contain the current readings for this outlet. For single-phase outlets, this property
	// shall contain a duplicate copy of the current sensor referenced in the CurrentAmps property, if present. For
	// poly-phase outlets, this property should contain multiple current sensor readings used to fully describe the
	// outlet.
	PolyPhaseCurrentAmps CurrentSensors
	// PolyPhaseVoltage shall contain the voltage readings for this outlet. For single-phase outlets, this property
	// shall contain a duplicate copy of the voltage sensor referenced in the Voltage property, if present. For poly-
	// phase outlets, this property should contain multiple voltage sensor readings used to fully describe the outlet.
	PolyPhaseVoltage VoltageSensors
	// PowerControlLocked shall indicate whether requests to the PowerControl action are locked. If 'true', services
	// shall reject requests to the PowerControl action.
	PowerControlLocked bool
	// PowerCycleDelaySeconds shall contain the number of seconds to delay power on after a PowerControl action to
	// cycle power. The value '0' shall indicate no delay to power on.
	PowerCycleDelaySeconds float64
	// PowerEnabled shall indicate the power enable state of the outlet. The value 'true' shall indicate that the
	// outlet can be powered on, and 'false' shall indicate that the outlet cannot be powered.
	PowerEnabled bool
	// PowerLoadPercent shall contain the power load, in percent units, for this outlet that represents the 'Total'
	// ElectricalContext for this outlet. The value of the DataSourceUri property, if present, shall reference a
	// resource of type Sensor with the ReadingType property containing the value 'Percent'.
	PowerLoadPercent SensorExcerpt
	// PowerOffDelaySeconds shall contain the number of seconds to delay power off after a PowerControl action. The
	// value '0' shall indicate no delay to power off.
	PowerOffDelaySeconds float64
	// PowerOnDelaySeconds shall contain the number of seconds to delay power up after a power cycle or a PowerControl
	// action. The value '0' shall indicate no delay to power up.
	PowerOnDelaySeconds float64
	// PowerRestoreDelaySeconds shall contain the number of seconds to delay power on after a power fault. The value
	// '0' shall indicate no delay to power on.
	PowerRestoreDelaySeconds float64
	// PowerRestorePolicy shall contain the desired PowerState of the outlet when power is applied. The value
	// 'LastState' shall return the outlet to the PowerState it was in when power was lost.
	PowerRestorePolicy PowerRestorePolicyTypes
	// PowerState shall contain the power state of the outlet.
	PowerState PowerState
	// PowerStateInTransition shall indicate whether the PowerState property will undergo a transition between on and
	// off states due to a configured delay. The transition may be due to the configuration of the power on, off, or
	// restore delay properties. If 'true', the PowerState property will transition at the conclusion of a configured
	// delay.
	PowerStateInTransition bool
	// PowerWatts shall contain the total power, in watt units, for this outlet that represents the 'Total'
	// ElectricalContext sensor when multiple power sensors exist for this outlet. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Power'.
	PowerWatts SensorPowerExcerpt
	// RatedCurrentAmps shall contain the rated maximum current for this outlet, in ampere units, after any required
	// de-rating, due to safety agency or other regulatory requirements, has been applied.
	RatedCurrentAmps float64
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UserLabel shall contain a user-assigned label used to identify this resource. If a value has not been assigned
	// by a user, the value of this property shall be an empty string.
	UserLabel string
	// Voltage shall contain the voltage, in volt units, for this single-phase outlet. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Voltage'. This property shall not appear in resource instances representing poly-phase outlets.
	Voltage SensorVoltageExcerpt
	// VoltageType shall contain the type of voltage applied to the outlet.
	VoltageType VoltageType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	powerControlTarget string
	resetMetricsTarget string

	branchCircuit string
	chassis       []string
	// ChassisCount is the number of chassis connected to this outlet.
	ChassisCount         int
	distributionCircuits []string
	// DistributionCircuitsCount is the number of mains or input circuits powered by this outlet.
	DistributionCircuitsCount int
	powerSupplies             []string
	// PowerSuppliesCount is the number of power supplies connected to this outlet.
	PowerSuppliesCount int
}

// UnmarshalJSON unmarshals a Outlet object from the raw JSON.
func (outlet *Outlet) UnmarshalJSON(b []byte) error {
	type temp Outlet
	type Actions struct {
		PowerControl common.ActionTarget `json:"#Outlet.PowerControl"`
		ResetMetrics common.ActionTarget `json:"#Outlet.ResetMetrics"`
	}
	type Links struct {
		BranchCircuit             common.Link
		Chassis                   common.Links
		ChassisCount              int `json:"Chassis@odata.count"`
		DistributionCircuits      common.Links
		DistributionCircuitsCount int `json:"DistributionCircuits@odata.count"`
		PowerSupplies             common.Links
		PowerSuppliesCount        int `json:"PowerSupplies@odata.count"`
	}
	var t struct {
		temp
		Actions Actions
		Links   Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*outlet = Outlet(t.temp)

	// Extract the links to other entities for later
	outlet.powerControlTarget = t.Actions.PowerControl.Target
	outlet.resetMetricsTarget = t.Actions.ResetMetrics.Target

	outlet.branchCircuit = t.Links.BranchCircuit.String()
	outlet.chassis = t.Links.Chassis.ToStrings()
	outlet.ChassisCount = t.Links.ChassisCount
	outlet.distributionCircuits = t.Links.DistributionCircuits.ToStrings()
	outlet.DistributionCircuitsCount = t.Links.DistributionCircuitsCount
	outlet.powerSupplies = t.Links.PowerSupplies.ToStrings()
	outlet.PowerSuppliesCount = t.Links.PowerSuppliesCount

	// This is a read/write object, so we need to save the raw object data for later
	outlet.rawData = b

	return nil
}

// PowerControl controls the power state of the outlet.
func (outlet *Outlet) PowerControl(powerState ActionPowerState) error {
	params := struct {
		PowerState ActionPowerState
	}{
		PowerState: powerState,
	}
	return outlet.Post(outlet.powerControlTarget, params)
}

// ResetMetrics resets metrics related to this outlet.
func (outlet *Outlet) ResetMetrics() error {
	return outlet.Post(outlet.resetMetricsTarget, nil)
}

// BranchCircuit gets the branch circuit associated with this outlet.
func (outlet *Outlet) BranchCircuit() (*Circuit, error) {
	if outlet.branchCircuit == "" {
		return nil, nil
	}
	return GetCircuit(outlet.GetClient(), outlet.branchCircuit)
}

// Chassis gets the chassis connected to this outlet.
func (outlet *Outlet) Chassis() ([]*Chassis, error) {
	return common.GetObjects[Chassis](outlet.GetClient(), outlet.chassis)
}

// DistributionCircuits gets the circuits powered by this outlet.
func (outlet *Outlet) DistributionCircuits() ([]*Circuit, error) {
	return common.GetObjects[Circuit](outlet.GetClient(), outlet.distributionCircuits)
}

// PowerSupplies gets the power supplies connected to this outlet.
func (outlet *Outlet) PowerSupplies() ([]*PowerSupply, error) {
	return common.GetObjects[PowerSupply](outlet.GetClient(), outlet.powerSupplies)
}

// Update commits updates to this object's properties to the running system.
func (outlet *Outlet) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Outlet)
	original.UnmarshalJSON(outlet.rawData)

	readWriteFields := []string{
		"ConfigurationLocked",
		"ElectricalConsumerNames",
		"LocationIndicatorActive",
		"PowerControlLocked",
		"PowerCycleDelaySeconds",
		"PowerOffDelaySeconds",
		"PowerOnDelaySeconds",
		"PowerRestoreDelaySeconds",
		"PowerRestorePolicy",
		"UserLabel",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(outlet).Elem()

	return outlet.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetOutlet will get a Outlet instance from the service.
func GetOutlet(c common.Client, uri string) (*Outlet, error) {
	return common.GetObject[Outlet](c, uri)
}

// ListReferencedOutlets gets the collection of Outlet from
// a provided reference.
func ListReferencedOutlets(c common.Client, link string) ([]*Outlet, error) {
	return common.GetCollectionObjects[Outlet](c, link)
}

// VoltageSensors shall contain properties that describe voltage sensor readings for an outlet.
type VoltageSensors struct {
	// Line1ToLine2 shall contain the line-to-line voltage, in volt units, between L1 and L2. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. This property shall not be present if the equipment does not include an L1-L2
	// measurement.
	Line1ToLine2 SensorVoltageExcerpt
	// Line1ToNeutral shall contain the line-to-line voltage, in volt units, between L1 and Neutral. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. This property shall not be present if the equipment does not include an
	// L1-Neutral measurement.
	Line1ToNeutral SensorVoltageExcerpt
	// Line2ToLine3 shall contain the line-to-line voltage, in volt units, between L2 and L3. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. This property shall not be present if the equipment does not include an L2-L3
	// measurement.
	Line2ToLine3 SensorVoltageExcerpt
	// Line2ToNeutral shall contain the line-to-line voltage, in volt units, between L2 and Neutral. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. This property shall not be present if the equipment does not include an
	// L2-Neutral measurement.
	Line2ToNeutral SensorVoltageExcerpt
	// Line3ToLine1 shall contain the line-to-line voltage, in volt units, between L3 and L1. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. This property shall not be present if the equipment does not include an L3-L1
	// measurement.
	Line3ToLine1 SensorVoltageExcerpt
	// Line3ToNeutral shall contain the line-to-line voltage, in volt units, between L3 and Neutral. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Voltage'. This property shall not be present if the equipment does not include an
	// L3-Neutral measurement.
	Line3ToNeutral SensorVoltageExcerpt
}
