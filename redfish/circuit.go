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

// The state of the over current protection device.
type BreakerState string

const (
	// The breaker is powered on.
	NormalBreakerState BreakerState = "Normal"
	// The breaker is off.
	OffBreakerState BreakerState = "Off"
	// The breaker has been tripped.
	TrippedBreakerState BreakerState = "Tripped"
)

// In Actions: BreakerControl, PowerControl:
// The desired power state of the circuit if the breaker is reset successfully.
type ActionPowerState string

const (
	// Power off.
	OffActionPowerState ActionPowerState = "Off"
	// Power on.
	OnActionPowerState ActionPowerState = "On"
	// Power cycle.
	PowerCycleActionPowerState ActionPowerState = "PowerCycle"
)

// The type of circuit.
type CircuitType string

const (
	// A branch (output) circuit.
	BranchCircuitType CircuitType = "Branch"
	// (v1.3+)	An electrical bus circuit.
	BusCircuitType CircuitType = "Bus"
	// A feeder (output) circuit.
	FeederCircuitType CircuitType = "Feeder"
	// A mains input or utility circuit.
	MainsCircuitType CircuitType = "Mains"
	// A subfeed (output) circuit.
	SubfeedCircuitType CircuitType = "Subfeed"
)

// The nominal voltage for an outlet.
type NominalVoltage string

const (
	// AC 100-127V nominal.
	AC100To127VNominal NominalVoltage = "AC100To127V"
	// AC 100-240V nominal.
	AC100To240VNominal NominalVoltage = "AC100To240V"
	// AC 100-277V nominal.
	AC100To277VNominal NominalVoltage = "AC100To277V"
	// AC 120V nominal.
	AC120VNominal NominalVoltage = "AC120V"
	// AC 200-240V nominal.
	AC200To240VNominal NominalVoltage = "AC200To240V"
	// AC 200-277V nominal.
	AC200To277VNominal NominalVoltage = "AC200To277V"
	// AC 208V nominal.
	AC208VNominal NominalVoltage = "AC208V"
	// AC 230V nominal.
	AC230VNominal NominalVoltage = "AC230V"
	// AC 200-240V and DC 380V.
	AC240AndDC380VNominal NominalVoltage = "AC240AndDC380V"
	// AC 240V nominal.
	AC240VNominal NominalVoltage = "AC240V"
	// AC 200-277V and DC 380V.
	AC277AndDC380VNominal NominalVoltage = "AC277AndDC380V"
	// AC 277V nominal.
	AC277VNominal NominalVoltage = "AC277V"
	// AC 400V or 415V nominal.
	AC400VNominal NominalVoltage = "AC400V"
	// AC 480V nominal.
	AC480VNominal NominalVoltage = "AC480V"
	// DC 12V nominal.
	DC12VNominal NominalVoltage = "DC12V"
	// DC 16V nominal.
	DC16VNominal NominalVoltage = "DC16V"
	// DC 1.8V nominal.
	DC1_8VNominal NominalVoltage = "DC1_8V"
	// DC 240V nominal.
	DC240VNominal NominalVoltage = "DC240V"
	// High Voltage DC (380V).
	DC380VNominal NominalVoltage = "DC380V"
	// DC 3.3V nominal.
	DC3_3VNominal NominalVoltage = "DC3_3V"
	// DC 48V nominal.
	DC48VNominal NominalVoltage = "DC48V"
	// DC 5V nominal.
	DC5VNominal NominalVoltage = "DC5V"
	// DC 9V nominal.
	DC9VNominal NominalVoltage = "DC9V"
	// -48V DC.
	DCNeg48VNominal NominalVoltage = "DCNeg48V"
)

// The number of ungrounded current-carrying conductors (phases)
// and the total number of conductors (wires).
type PhaseWiringType string

const (
	// Single or Two-Phase / 3-Wire (Line1, Line2 or Neutral, Protective Earth).
	OneOrTwoPhase3WirePhaseWiringType PhaseWiringType = "OneOrTwoPhase3Wire"
	// Single-phase / 3-Wire (Line1, Neutral, Protective Earth).
	OnePhase3WirePhaseWiringType PhaseWiringType = "OnePhase3Wire"
	// Three-phase / 4-Wire (Line1, Line2, Line3, Protective Earth).
	ThreePhase4WirePhaseWiringType PhaseWiringType = "ThreePhase4Wire"
	// Three-phase / 5-Wire (Line1, Line2, Line3, Neutral, Protective Earth).
	ThreePhase5WirePhaseWiringType PhaseWiringType = "ThreePhase5Wire"
	// Two-phase / 3-Wire (Line1, Line2, Protective Earth).
	TwoPhase3WirePhaseWiringType PhaseWiringType = "TwoPhase3Wire"
	// Two-phase / 4-Wire (Line1, Line2, Neutral, Protective Earth).
	TwoPhase4WirePhaseWiringType PhaseWiringType = "TwoPhase4Wire"
)

// The type of plug according to NEMA, IEC, or regional standards.
type PlugType string

const (
	// California Standard CS8265 (Single-phase 250V; 50A; 2P3W).
	CaliforniaCS8265PlugType PlugType = "California_CS8265"
	// California Standard CS8365 (Three-phase 250V; 50A; 3P4W).
	CaliforniaCS8365PlugType PlugType = "California_CS8365"
	// Field-wired; Three-phase 200-250V; 60A; 3P4W.
	Field208V3P4W60APlugType PlugType = "Field_208V_3P4W_60A"
	// Field-wired; Three-phase 200-240/346-415V; 32A; 3P5W.
	Field400V3P5W32APlugType PlugType = "Field_400V_3P5W_32A"
	// IEC 60309 316P6 (Single-phase 200-250V; 16A; 1P3W; Blue, 6-hour).
	IEC60309316P6PlugType PlugType = "IEC_60309_316P6"
	// IEC 60309 332P6 (Single-phase 200-250V; 32A; 1P3W; Blue, 6-hour).
	IEC60309332P6PlugType PlugType = "IEC_60309_332P6"
	// IEC 60309 363P6 (Single-phase 200-250V; 63A; 1P3W; Blue, 6-hour).
	IEC60309363P6PlugType PlugType = "IEC_60309_363P6"
	// IEC 60309 460P9 (Three-phase 200-250V; 60A; 3P4W; Blue; 9-hour).
	IEC60309460P9PlugType PlugType = "IEC_60309_460P9"
	// IEC 60309 516P6 (Three-phase 200-240/346-415V; 16A; 3P5W; Red; 6-hour).
	IEC60309516P6PlugType PlugType = "IEC_60309_516P6"
	// IEC 60309 532P6 (Three-phase 200-240/346-415V; 32A; 3P5W; Red; 6-hour).
	IEC60309532P6PlugType PlugType = "IEC_60309_532P6"
	// IEC 60309 560P9 (Three-phase 120-144/208-250V; 60A; 3P5W; Blue; 9-hour).
	IEC60309560P9PlugType PlugType = "IEC_60309_560P9"
	// IEC 60309 563P6 (Three-phase 200-240/346-415V; 63A; 3P5W; Red; 6-hour).
	IEC60309563P6PlugType PlugType = "IEC_60309_563P6"
	// IEC C14 (Single-phase 250V; 10A; 1P3W).
	IEC60320C14PlugType PlugType = "IEC_60320_C14"
	// IEC C20 (Single-phase 250V; 16A; 1P3W).
	IEC60320C20PlugType PlugType = "IEC_60320_C20"
	// NEMA 5-15P (Single-phase 125V; 15A; 1P3W).
	NEMA515PPlugType PlugType = "NEMA_5_15P"
	// NEMA 5-20P (Single-phase 125V; 20A; 1P3W).
	NEMA520PPlugType PlugType = "NEMA_5_20P"
	// NEMA 6-15P (Single-phase 250V; 15A; 2P3W).
	NEMA615PPlugType PlugType = "NEMA_6_15P"
	// NEMA 6-20P (Single-phase 250V; 20A; 2P3W).
	NEMA620PPlugType PlugType = "NEMA_6_20P"
	// NEMA L14-20P (Split-phase 125/250V; 20A; 2P4W).
	NEMAL1420PPlugType PlugType = "NEMA_L14_20P"
	// NEMA L14-30P (Split-phase 125/250V; 30A; 2P4W).
	NEMAL1430PPlugType PlugType = "NEMA_L14_30P"
	// NEMA L15-20P (Three-phase 250V; 20A; 3P4W).
	NEMAL1520PPlugType PlugType = "NEMA_L15_20P"
	// NEMA L15-30P (Three-phase 250V; 30A; 3P4W).
	NEMAL1530PPlugType PlugType = "NEMA_L15_30P"
	// NEMA L21-20P (Three-phase 120/208V; 20A; 3P5W).
	NEMAL2120PPlugType PlugType = "NEMA_L21_20P"
	// NEMA L21-30P (Three-phase 120/208V; 30A; 3P5W).
	NEMAL2130PPlugType PlugType = "NEMA_L21_30P"
	// NEMA L22-20P (Three-phase 277/480V; 20A; 3P5W).
	NEMAL2220PPlugType PlugType = "NEMA_L22_20P"
	// NEMA L22-30P (Three-phase 277/480V; 30A; 3P5W).
	NEMAL2230PPlugType PlugType = "NEMA_L22_30P"
	// NEMA L5-15P (Single-phase 125V; 15A; 1P3W).
	NEMAL515PPlugType PlugType = "NEMA_L5_15P"
	// NEMA L5-20P (Single-phase 125V; 20A; 1P3W).
	NEMAL520PPlugType PlugType = "NEMA_L5_20P"
	// NEMA L5-30P (Single-phase 125V; 30A; 1P3W).
	NEMAL530PPlugType PlugType = "NEMA_L5_30P"
	// NEMA L6-15P (Single-phase 250V; 15A; 2P3W).
	NEMAL615PPlugType PlugType = "NEMA_L6_15P"
	// NEMA L6-20P (Single-phase 250V; 20A; 2P3W).
	NEMAL620PPlugType PlugType = "NEMA_L6_20P"
	// NEMA L6-30P (Single-phase 250V; 30A; 2P3W).
	NEMAL630PPlugType PlugType = "NEMA_L6_30P"
)

// The current readings for this circuit.
type PolyPhaseCurrentAmps struct {
	// 	Line 1 current (A).
	Line1 SensorVoltageExcerpt
	// 	Line 2 current (A).
	Line2 SensorVoltageExcerpt
	// 	Line 3 current (A).
	Line3 SensorVoltageExcerpt
	// Neutral line current (A).
	Neutral SensorVoltageExcerpt
}

// The energy readings for this circuit.
type PolyPhaseEnergykWh struct {
	// The Line 1 to Line 2 energy (kWh) for this circuit.
	Line1ToLine2 SensorEnergykWhExcerpt
	// The Line 1 to Neutral energy (kWh) for this circuit.
	Line1ToNeutral SensorEnergykWhExcerpt
	// The Line 2 to Line 3 energy (kWh) for this circuit.
	Line2ToLine3 SensorEnergykWhExcerpt
	// The Line 2 to Neutral energy (kWh) for this circuit.
	Line2ToNeutral SensorEnergykWhExcerpt
	// The Line 3 to Line 1 energy (kWh) for this circuit.
	Line3ToLine1 SensorEnergykWhExcerpt
	// The Line 3 to Neutral energy (kWh) for this circuit.
	Line3ToNeutral SensorEnergykWhExcerpt
}

// The power readings for this circuit.
type PolyPhasePowerWatts struct {
	// The Line 1 to Line 2 power (W) for this circuit.
	Line1ToLine2 SensorPowerExcerpt
	// The Line 1 to Neutral power (W) for this circuit.
	Line1ToNeutral SensorPowerExcerpt
	// The Line 2 to Line 3 power (W) for this circuit.
	Line2ToLine3 SensorPowerExcerpt
	// The Line 2 to Neutral power (W) for this circuit.
	Line2ToNeutral SensorPowerExcerpt
	// The Line 3 to Line 1 power (W) for this circuit.
	Line3ToLine1 SensorPowerExcerpt
	// The Line 3 to Neutral power (W) for this circuit.
	Line3ToNeutral SensorPowerExcerpt
}

// The voltage readings for this circuit.
type PolyPhaseVoltage struct {
	// The Line 1 to Line 2 voltage (V) for this circuit.
	Line1ToLine2 SensorVoltageExcerpt
	// The Line 1 to Neutral voltage (V) for this circuit.
	Line1ToNeutral SensorVoltageExcerpt
	// The Line 2 to Line 3 voltage (V) for this circuit.
	Line2ToLine3 SensorVoltageExcerpt
	// The Line 2 to Neutral voltage (V) for this circuit.
	Line2ToNeutral SensorVoltageExcerpt
	// The Line 3 to Line 1 voltage (V) for this circuit.
	Line3ToLine1 SensorVoltageExcerpt
	// The Line 3 to Neutral voltage (V) for this circuit.
	Line3ToNeutral SensorVoltageExcerpt
}

// Circuit shall be used to represent
// an electrical circuit for a Redfish implementation.
type Circuit struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string

	// The state of the over current protection device.
	BreakerState BreakerState
	// The type of circuit.
	CircuitType CircuitType
	// Indicates whether the configuration is locked.
	ConfigurationLocked bool
	// Designates if this is a critical circuit.
	CriticalCircuit bool
	// The current (A) for this single phase circuit.
	CurrentAmps SensorVoltageExcerpt
	// An array of names of downstream devices that are powered by this circuit.
	ElectricalConsumerNames []string
	// The combination of current-carrying conductors.
	ElectricalContext common.ElectricalContext
	// The URI of the management interface
	// for the upstream electrical source connection for this circuit.
	ElectricalSourceManagerURI string
	// The name of the upstream electrical source,
	// such as a circuit or outlet, connected to this circuit.
	ElectricalSourceName string
	// The energy (kWh) for this circuit.
	EnergykWh SensorEnergykWhExcerpt
	// The frequency (Hz) for this circuit.
	FrequencyHz SensorExcerpt
	// Deprecated: (v1.1)
	// The state of the indicator LED, which identifies the circuit.
	IndicatorLED common.IndicatorLED
	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool
	// The nominal voltage for this circuit.
	NominalVoltage NominalVoltage
	// The number of ungrounded current-carrying conductors (phases)
	// and the total number of conductors (wires).
	PhaseWiringType PhaseWiringType
	// The type of plug according to NEMA, IEC, or regional standards.
	PlugType PlugType
	// The current readings for this circuit.
	PolyPhaseCurrentAmps PolyPhaseCurrentAmps
	// The energy readings for this circuit.
	PolyPhaseEnergykWh PolyPhaseEnergykWh
	// The power readings for this circuit.
	PolyPhasePowerWatts PolyPhasePowerWatts
	// The voltage readings for this circuit.
	PolyPhaseVoltage PolyPhaseVoltage
	// Indicates whether power control requests are locked.
	PowerControlLocked bool
	// The number of seconds to delay power on
	// after a PowerControl action to cycle power.
	// Zero seconds indicates no delay.
	PowerCycleDelaySeconds float32
	// Indicates if the circuit can be powered.
	PowerEnabled bool
	// The power load (percent) for this circuit.
	PowerLoadPercent SensorExcerpt
	// The number of seconds to delay power off after a PowerControl action.
	// Zero seconds indicates no delay to power off.
	PowerOffDelaySeconds float32
	// The number of seconds to delay power up after a power cycle or
	// a PowerControl action. Zero seconds indicates no delay to power up.
	PowerOnDelaySeconds float32
	// The number of seconds to delay power on after power has been restored.
	// Zero seconds indicates no delay.
	PowerRestoreDelaySeconds float32
	// The desired power state of the circuit
	// when power is restored after a power loss.
	PowerRestorePolicy PowerRestorePolicyTypes
	// The power state of the circuit.
	PowerState PowerState
	// Indicates whether the power state is undergoing a delayed transition.
	PowerStateInTransition bool
	// The power (W) for this circuit.
	PowerWatts SensorPowerExcerpt
	// The rated maximum current allowed for this circuit.
	RatedCurrentAmps float32
	// The status and health of the resource and its subordinate or dependent resources.
	Status common.Status
	// The current imbalance (percent) between phases.
	UnbalancedCurrentPercent SensorExcerpt
	// The voltage imbalance (percent) between phases.
	UnbalancedVoltagePercent SensorExcerpt
	// A user-assigned label.
	UserLabel string
	// The voltage (V) for this single phase circuit.
	Voltage SensorVoltageExcerpt
	// The type of voltage applied to the circuit.
	VoltageType InputType

	// Links section
	// A reference to the branch circuit related to this circuit.
	branchCircuit string
	// An array of links to the circuits powered by this circuit.
	distributionCircuits      []string
	DistributionCircuitsCount int
	// OemLinks are all OEM data under link section
	OemLinks json.RawMessage
	// An array of references to the outlets contained by this circuit.
	outlets      []string
	OutletsCount int
	// A link to the power outlet that provides power to this circuit.
	powerOutlet string
	// A link to the circuit that provides power to this circuit.
	sourceCircuit string

	// Actions section
	// This action attempts to reset the circuit breaker.
	breakerControlTarget string
	// This action turns the circuit on or off.
	powerControlTarget string
	// This action resets metrics related to this circuit.
	resetMetricsTarget string
	// OemActions contains all the vendor specific actions.
	// It is vendor responsibility to parse this field accordingly
	OemActions json.RawMessage

	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Circuit object from the raw JSON.
func (circuit *Circuit) UnmarshalJSON(b []byte) error {
	type temp Circuit
	type linkReference struct {
		BranchCircuit             common.Link
		DistributionCircuits      common.Links
		DistributionCircuitsCount int `json:"DistributionCircuits@odata.count"`
		Outlets                   common.Links
		OutletsCount              int `json:"Outlets@odata.count"`
		Oem                       json.RawMessage
		PowerOutlet               common.Link
		SourceCircuit             common.Link
	}
	type actions struct {
		BreakerControl common.ActionTarget `json:"#Circuit.BreakerControl"`
		PowerControl   common.ActionTarget `json:"#Circuit.PowerControl"`
		ResetMetrics   common.ActionTarget `json:"#Circuit.ResetMetrics"`
		Oem            json.RawMessage     // OEM actions will be stored here
	}
	var t struct {
		temp

		Links   linkReference
		Actions actions
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	// Extract the links to other entities for later
	*circuit = Circuit(t.temp)

	circuit.branchCircuit = t.Links.BranchCircuit.String()
	circuit.distributionCircuits = t.Links.DistributionCircuits.ToStrings()
	circuit.DistributionCircuitsCount = t.Links.DistributionCircuitsCount
	circuit.outlets = t.Links.Outlets.ToStrings()
	circuit.OutletsCount = t.Links.OutletsCount
	circuit.OemLinks = t.Links.Oem
	circuit.powerOutlet = t.Links.PowerOutlet.String()
	circuit.sourceCircuit = t.Links.SourceCircuit.String()

	circuit.breakerControlTarget = t.Actions.BreakerControl.Target
	circuit.powerControlTarget = t.Actions.PowerControl.Target
	circuit.resetMetricsTarget = t.Actions.ResetMetrics.Target
	circuit.OemActions = t.Actions.Oem

	// This is a read/write object, so we need to save the raw object data for later
	circuit.rawData = b

	return nil
}

// GetCircuit will get a Circuit instance from the Redfish service.
func GetCircuit(c common.Client, uri string) (*Circuit, error) {
	return common.GetObject[Circuit](c, uri)
}

// Update commits updates to this object's properties to the running system.
func (circuit *Circuit) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	ct := new(Circuit)
	err := ct.UnmarshalJSON(circuit.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"ConfigurationLocked",
		"CriticalCircuit",
		"ElectricalConsumerNames",
		"ElectricalSourceManagerURI",
		"ElectricalSourceName",
		"IndicatorLED",
		"LocationIndicatorActive",
		"PowerControlLocked",
		"PowerCycleDelaySeconds",
		"PowerOffDelaySeconds",
		"PowerOnDelaySeconds",
		"PowerRestoreDelaySeconds",
		"PowerRestorePolicy",
		"UserLabel",
	}

	originalElement := reflect.ValueOf(ct).Elem()
	currentElement := reflect.ValueOf(circuit).Elem()

	return circuit.Entity.Update(originalElement, currentElement, readWriteFields)
}

// This action shall control the state of the circuit breaker or over-current protection device.
func (circuit *Circuit) BreakerControl(powerState ActionPowerState) error {
	if circuit.breakerControlTarget == "" {
		return errors.New("BreakerControl is not supported")
	}

	t := struct {
		PowerState ActionPowerState
	}{PowerState: powerState}

	return circuit.Post(circuit.breakerControlTarget, t)
}

// This action shall control the power state of the circuit.
func (circuit *Circuit) PowerControl(powerState ActionPowerState) error {
	if circuit.powerControlTarget == "" {
		return errors.New("PowerControl is not supported")
	}

	t := struct {
		PowerState ActionPowerState
	}{PowerState: powerState}

	return circuit.Post(circuit.powerControlTarget, t)
}

// This action shall reset any time intervals or counted values for this circuit.
func (circuit *Circuit) ResetMetrics() error {
	if circuit.resetMetricsTarget == "" {
		return errors.New("ResetMetrics is not supported")
	}

	return circuit.Post(circuit.resetMetricsTarget, nil)
}

// ListReferencedCircuits gets the collection of Circuits from
// a provided reference.
func ListReferencedCircuits(c common.Client, link string) ([]*Circuit, error) {
	return common.GetCollectionObjects[Circuit](c, link)
}

// BranchCircuit gets a resource that represents the branch circuit associated with this circuit.
func (circuit *Circuit) BranchCircuit() (*Circuit, error) {
	return GetCircuit(circuit.GetClient(), circuit.branchCircuit)
}

// SourceCircuit gets a resource that represents the circuit that provides power to this circuit.
func (circuit *Circuit) SourceCircuit() (*Circuit, error) {
	return GetCircuit(circuit.GetClient(), circuit.sourceCircuit)
}

// DistributionCircuits gets the collection that contains the circuits powered by this circuit.
func (circuit *Circuit) DistributionCircuits() ([]*Circuit, error) {
	return common.GetObjects[Circuit](circuit.GetClient(), circuit.distributionCircuits)
}

// TODO: outlets, power outlet
