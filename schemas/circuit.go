//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #Circuit.v1_9_0.Circuit

package schemas

import (
	"encoding/json"
)

type BreakerStates string

const (
	// NormalBreakerStates The breaker is powered on.
	NormalBreakerStates BreakerStates = "Normal"
	// TrippedBreakerStates The breaker has been tripped.
	TrippedBreakerStates BreakerStates = "Tripped"
	// OffBreakerStates The breaker is off.
	OffBreakerStates BreakerStates = "Off"
)

type CircuitType string

const (
	// MainsCircuitType is a mains input or utility circuit.
	MainsCircuitType CircuitType = "Mains"
	// BranchCircuitType is a branch (output) circuit.
	BranchCircuitType CircuitType = "Branch"
	// SubfeedCircuitType is a subfeed (output) circuit.
	SubfeedCircuitType CircuitType = "Subfeed"
	// FeederCircuitType is a feeder (output) circuit.
	FeederCircuitType CircuitType = "Feeder"
	// BusCircuitType is an electrical bus circuit.
	BusCircuitType CircuitType = "Bus"
)

type NominalVoltageType string

const (
	// AC100To127VNominalVoltageType shall indicate the device supports a nominal
	// voltage in the complete range of 100-127VAC. Range values are generally used
	// to describe support on device inputs or inlets.
	AC100To127VNominalVoltageType NominalVoltageType = "AC100To127V"
	// AC100To240VNominalVoltageType shall indicate the device supports a nominal
	// voltage in the complete range of 100-240VAC. Range values are generally used
	// to describe support on device inputs or inlets.
	AC100To240VNominalVoltageType NominalVoltageType = "AC100To240V"
	// AC100To277VNominalVoltageType shall indicate the device supports a nominal
	// voltage in the complete range of 100-277VAC. Range values are generally used
	// to describe support on device inputs or inlets.
	AC100To277VNominalVoltageType NominalVoltageType = "AC100To277V"
	// AC120VNominalVoltageType shall indicate the device supports a nominal
	// voltage of 120VAC. Specific values are generally used to describe support on
	// device outputs or outlets.
	AC120VNominalVoltageType NominalVoltageType = "AC120V"
	// AC200To240VNominalVoltageType shall indicate the device supports a nominal
	// voltage in the complete range of 200-240VAC. Range values are generally used
	// to describe support on device inputs or inlets.
	AC200To240VNominalVoltageType NominalVoltageType = "AC200To240V"
	// AC200To277VNominalVoltageType shall indicate the device supports a nominal
	// voltage in the complete range of 200-277VAC. Range values are generally used
	// to describe support on device inputs or inlets.
	AC200To277VNominalVoltageType NominalVoltageType = "AC200To277V"
	// AC208VNominalVoltageType shall indicate the device supports a nominal
	// voltage of 208VAC. Specific AC voltage values are generally used to describe
	// support on device outputs or outlets.
	AC208VNominalVoltageType NominalVoltageType = "AC208V"
	// AC230VNominalVoltageType shall indicate the device supports a nominal
	// voltage of 230AC. Specific AC voltage values are generally used to describe
	// support on device outputs or outlets.
	AC230VNominalVoltageType NominalVoltageType = "AC230V"
	// AC240VNominalVoltageType shall indicate the device supports a nominal
	// voltage of 240VAC. Specific AC voltage values are generally used to describe
	// support on device outputs or outlets.
	AC240VNominalVoltageType NominalVoltageType = "AC240V"
	// AC240AndDC380VNominalVoltageType shall indicate the device supports a
	// nominal voltage in the complete range of 200-240VAC or a value of 380VDC.
	// Range values are generally used to describe support on device inputs or
	// inlets.
	AC240AndDC380VNominalVoltageType NominalVoltageType = "AC240AndDC380V"
	// AC277VNominalVoltageType shall indicate the device supports a nominal
	// voltage of 277VAC. Specific AC high-voltage values may be used to describe
	// support on device inputs or outputs.
	AC277VNominalVoltageType NominalVoltageType = "AC277V"
	// AC277AndDC380VNominalVoltageType shall indicate the device supports a
	// nominal voltage in the complete range of 200-277VAC or a value of 380VDC.
	// Range values are generally used to describe support on device inputs or
	// inlets.
	AC277AndDC380VNominalVoltageType NominalVoltageType = "AC277AndDC380V"
	// AC400VNominalVoltageType shall indicate the device supports a nominal
	// voltage of 400VAC or 415VAC. Specific AC high-voltage values may be used to
	// describe support on device inputs or outputs.
	AC400VNominalVoltageType NominalVoltageType = "AC400V"
	// AC480VNominalVoltageType shall indicate the device supports a nominal
	// voltage of 480VAC. Specific AC high-voltage values may be used to describe
	// support on device inputs or outputs.
	AC480VNominalVoltageType NominalVoltageType = "AC480V"
	// DC48VNominalVoltageType shall indicate the device supports a nominal voltage
	// of 48VDC.
	DC48VNominalVoltageType NominalVoltageType = "DC48V"
	// DC240VNominalVoltageType shall indicate the device supports a nominal
	// voltage of 240VDC.
	DC240VNominalVoltageType NominalVoltageType = "DC240V"
	// DC380VNominalVoltageType shall indicate the device supports a nominal
	// voltage of 380VDC.
	DC380VNominalVoltageType NominalVoltageType = "DC380V"
	// DC400VNominalVoltageType shall indicate the device supports a nominal
	// voltage of 400VDC.
	DC400VNominalVoltageType NominalVoltageType = "DC400V"
	// DC800VNominalVoltageType shall indicate the device supports a nominal
	// voltage of 800VDC.
	DC800VNominalVoltageType NominalVoltageType = "DC800V"
	// DCNeg48VNominalVoltageType shall indicate the device supports a nominal
	// voltage of -48VDC.
	DCNeg48VNominalVoltageType NominalVoltageType = "DCNeg48V"
	// DC16VNominalVoltageType shall indicate the device supports a nominal voltage
	// of 16VDC.
	DC16VNominalVoltageType NominalVoltageType = "DC16V"
	// DC12VNominalVoltageType shall indicate the device supports a nominal voltage
	// of 12VDC.
	DC12VNominalVoltageType NominalVoltageType = "DC12V"
	// DC9VNominalVoltageType shall indicate the device supports a nominal voltage
	// of 9VDC.
	DC9VNominalVoltageType NominalVoltageType = "DC9V"
	// DC5VNominalVoltageType shall indicate the device supports a nominal voltage
	// of 5VDC.
	DC5VNominalVoltageType NominalVoltageType = "DC5V"
	// DC33VNominalVoltageType shall indicate the device supports a nominal voltage
	// of 3.3VDC.
	DC33VNominalVoltageType NominalVoltageType = "DC3_3V"
	// DC18VNominalVoltageType shall indicate the device supports a nominal voltage
	// of 1.8VDC.
	DC18VNominalVoltageType NominalVoltageType = "DC1_8V"
)

type PhaseWiringType string

const (
	// OnePhase3WirePhaseWiringType shall represent a single-phase / 3-wire (Line1,
	// Neutral, Protective Earth) wiring.
	OnePhase3WirePhaseWiringType PhaseWiringType = "OnePhase3Wire"
	// TwoPhase3WirePhaseWiringType shall represent a two-phase / 3-wire (Line1,
	// Line2, Protective Earth) wiring.
	TwoPhase3WirePhaseWiringType PhaseWiringType = "TwoPhase3Wire"
	// OneOrTwoPhase3WirePhaseWiringType shall represent a single or two-phase /
	// 3-wire (Line1, Line2 or Neutral, Protective Earth) wiring. This value shall
	// be used when both phase configurations are supported. This is most common
	// where detachable cordsets are used. If poly-phase properties such as
	// 'PolyPhaseVoltage' are supported, the service should populate the
	// measurements as if the circuit is wired as Line1, Neutral, and Protective
	// Earth.
	OneOrTwoPhase3WirePhaseWiringType PhaseWiringType = "OneOrTwoPhase3Wire"
	// TwoPhase4WirePhaseWiringType shall represent a two-phase / 4-wire (Line1,
	// Line2, Neutral, Protective Earth) wiring.
	TwoPhase4WirePhaseWiringType PhaseWiringType = "TwoPhase4Wire"
	// ThreePhase4WirePhaseWiringType shall represent a three-phase / 4-wire
	// (Line1, Line2, Line3, Protective Earth) wiring.
	ThreePhase4WirePhaseWiringType PhaseWiringType = "ThreePhase4Wire"
	// ThreePhase5WirePhaseWiringType shall represent a three-phase / 5-wire
	// (Line1, Line2, Line3, Neutral, Protective Earth) wiring.
	ThreePhase5WirePhaseWiringType PhaseWiringType = "ThreePhase5Wire"
)

type PlugType string

const (
	// NEMA515PPlugType shall represent a plug that matches the NEMA specified 5-15
	// straight (non-locking) plug (Single-phase 125V; 15A; 1P3W).
	NEMA515PPlugType PlugType = "NEMA_5_15P"
	// NEMAL515PPlugType shall represent a plug that matches the NEMA specified
	// locking L5-15 plug (Single-phase 125V; 15A; 1P3W).
	NEMAL515PPlugType PlugType = "NEMA_L5_15P"
	// NEMA520PPlugType shall represent a plug that matches the NEMA specified 5-20
	// straight (non-locking) plug that exhibits a T-slot (Single-phase 125V; 20A;
	// 1P3W).
	NEMA520PPlugType PlugType = "NEMA_5_20P"
	// NEMAL520PPlugType shall represent a plug that matches the NEMA specified
	// locking L5-20 plug (Single-phase 125V; 20A; 1P3W).
	NEMAL520PPlugType PlugType = "NEMA_L5_20P"
	// NEMAL530PPlugType shall represent a plug that matches the NEMA specified
	// locking L5-30 plug (Single-phase 125V; 30A; 1P3W).
	NEMAL530PPlugType PlugType = "NEMA_L5_30P"
	// NEMA615PPlugType shall represent a plug that matches the NEMA specified 6-15
	// straight (non-locking) plug (Single-phase 250V; 15A; 2P3W).
	NEMA615PPlugType PlugType = "NEMA_6_15P"
	// NEMAL615PPlugType shall represent a plug that matches the NEMA specified
	// locking L6-15 plug (Single-phase 250V; 15A; 2P3W).
	NEMAL615PPlugType PlugType = "NEMA_L6_15P"
	// NEMA620PPlugType shall represent a plug that matches the NEMA specified 6-20
	// straight (non-locking) plug (Single-phase 250V; 20A; 2P3W).
	NEMA620PPlugType PlugType = "NEMA_6_20P"
	// NEMAL620PPlugType shall represent a plug that matches the NEMA specified
	// locking L6-20 plug (Single-phase 250V; 20A; 2P3W).
	NEMAL620PPlugType PlugType = "NEMA_L6_20P"
	// NEMAL630PPlugType shall represent a plug that matches the NEMA specified
	// locking L6-30 plug (Single-phase 250V; 30A; 2P3W).
	NEMAL630PPlugType PlugType = "NEMA_L6_30P"
	// NEMAL1420PPlugType shall represent a plug that matches the NEMA specified
	// locking L14-20 plug (Split-phase 125/250V; 20A; 2P4W).
	NEMAL1420PPlugType PlugType = "NEMA_L14_20P"
	// NEMAL1430PPlugType shall represent a plug that matches the NEMA specified
	// locking L14-30 plug (Split-phase 125/250V; 30A; 2P4W).
	NEMAL1430PPlugType PlugType = "NEMA_L14_30P"
	// NEMAL1520PPlugType shall represent a plug that matches the NEMA specified
	// locking L15-20 plug (Three-phase 250V; 20A; 3P4W).
	NEMAL1520PPlugType PlugType = "NEMA_L15_20P"
	// NEMAL1530PPlugType shall represent a plug that matches the NEMA specified
	// locking L15-30 plug (Three-phase 250V; 30A; 3P4W).
	NEMAL1530PPlugType PlugType = "NEMA_L15_30P"
	// NEMAL2120PPlugType shall represent a plug that matches the NEMA specified
	// locking L21-20 plug (Three-phase 120/208V; 20A; 3P5W).
	NEMAL2120PPlugType PlugType = "NEMA_L21_20P"
	// NEMAL2130PPlugType shall represent a plug that matches the NEMA specified
	// locking L21-30 plug (Three-phase 120/208V; 30A; 3P5W).
	NEMAL2130PPlugType PlugType = "NEMA_L21_30P"
	// NEMAL2220PPlugType shall represent a plug that matches the NEMA specified
	// locking L22-20 plug (Three-phase 277/480V; 20A; 3P5W).
	NEMAL2220PPlugType PlugType = "NEMA_L22_20P"
	// NEMAL2230PPlugType shall represent a plug that matches the NEMA specified
	// locking L22-30 plug (Three-phase 277/480V; 30A; 3P5W).
	NEMAL2230PPlugType PlugType = "NEMA_L22_30P"
	// CaliforniaCS8265PlugType shall represent a plug that matches the 'California
	// Standard' CS8265 style plug (Single-phase 250V; 50A; 2P3W).
	CaliforniaCS8265PlugType PlugType = "California_CS8265"
	// CaliforniaCS8365PlugType shall represent a plug that matches the 'California
	// Standard' CS8365 style plug (Three-phase 250V; 50A; 3P4W).
	CaliforniaCS8365PlugType PlugType = "California_CS8365"
	// IEC60320C14PlugType shall represent a plug that matches the IEC 60320
	// specified C14 input (Single-phase 250V; 10A; 1P3W).
	IEC60320C14PlugType PlugType = "IEC_60320_C14"
	// IEC60320C20PlugType shall represent a plug that matches the IEC 60320
	// specified C20 input (Single-phase 250V; 16A; 1P3W).
	IEC60320C20PlugType PlugType = "IEC_60320_C20"
	// IEC60309316P6PlugType shall represent a plug that matches the IEC 60309
	// 316P6 plug (Single-phase 200-250V; 16A; 1P3W; Blue, 6-hour).
	IEC60309316P6PlugType PlugType = "IEC_60309_316P6"
	// IEC60309332P6PlugType shall represent a plug that matches the IEC 60309
	// 332P6 plug (Single-phase 200-250V; 32A; 1P3W; Blue, 6-hour).
	IEC60309332P6PlugType PlugType = "IEC_60309_332P6"
	// IEC60309363P6PlugType shall represent a plug that matches the IEC 60309
	// 363P6 plug (Single-phase 200-250V; 63A; 1P3W; Blue, 6-hour).
	IEC60309363P6PlugType PlugType = "IEC_60309_363P6"
	// IEC60309516P6PlugType shall represent a plug that matches the IEC 60309
	// 516P6 plug (Three-phase 200-240/346-415V; 16A; 3P5W; Red; 6-hour).
	IEC60309516P6PlugType PlugType = "IEC_60309_516P6"
	// IEC60309532P6PlugType shall represent a plug that matches the IEC 60309 plug
	// 532P6 (Three-phase 200-240/346-415V; 32A; 3P5W; Red; 6-hour).
	IEC60309532P6PlugType PlugType = "IEC_60309_532P6"
	// IEC60309563P6PlugType shall represent a plug that matches the IEC 60309
	// 563P6 plug (Three-phase 200-240/346-415V; 63A; 3P5W; Red; 6-hour).
	IEC60309563P6PlugType PlugType = "IEC_60309_563P6"
	// IEC60309460P9PlugType shall represent a plug that matches the IEC 60309
	// 460P9 plug (Three-phase 200-250V; 60A; 3P4W; Blue; 9-hour).
	IEC60309460P9PlugType PlugType = "IEC_60309_460P9"
	// IEC60309560P9PlugType shall represent a plug that matches the IEC 60309 plug
	// 560P9 (Three-phase 120-144/208-250V; 60A; 3P5W; Blue; 9-hour).
	IEC60309560P9PlugType PlugType = "IEC_60309_560P9"
	// Field208V3P4W60APlugType shall represent field-wired input that is
	// three-phase 200-250V; 60A; 3P4W.
	Field208V3P4W60APlugType PlugType = "Field_208V_3P4W_60A"
	// Field400V3P5W32APlugType shall represent field-wired input that is
	// three-phase 200-240/346-415V; 32A; 3P5W.
	Field400V3P5W32APlugType PlugType = "Field_400V_3P5W_32A"
)

// Circuit shall be used to represent an electrical circuit for a Redfish
// implementation.
type Circuit struct {
	Entity
	// BreakerState shall contain the state of the overcurrent protection device.
	BreakerState BreakerStates
	// CircuitType shall contain the type of circuit.
	CircuitType CircuitType
	// ConfigurationLocked shall indicate whether modification requests to this
	// resource are not permitted. If 'true', services shall reject modification
	// requests to other properties in this resource.
	//
	// Version added: v1.5.0
	ConfigurationLocked bool
	// CriticalCircuit shall indicate whether the circuit is designated as a
	// critical circuit, and therefore is excluded from autonomous logic that could
	// affect the state of the circuit. The value shall be 'true' if the circuit is
	// deemed critical, and 'false' if the circuit is not critical.
	CriticalCircuit bool
	// CurrentAmps shall contain the current, in ampere units, for this circuit.
	// The value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Current'. This property shall not be present if 'PhaseWiringType'
	// contains a value that indicates a 4-wire or greater configuration, such as
	// 'TwoPhase4Wire'.
	CurrentAmps SensorCurrentExcerpt
	// ElectricalConsumerNames shall contain an array of user-assigned identifying
	// strings that describe downstream devices that are powered by this circuit.
	//
	// Version added: v1.4.0
	ElectricalConsumerNames []string
	// ElectricalContext shall contain the combination of current-carrying
	// conductors that distribute power.
	ElectricalContext ElectricalContext
	// ElectricalSourceManagerURI shall contain a URI to the management application
	// or device that provides monitoring or control of the upstream electrical
	// source that provides power to this circuit. If a value has not been assigned
	// by a user, the value of this property shall be an empty string.
	//
	// Version added: v1.4.0
	ElectricalSourceManagerURI string
	// ElectricalSourceName shall contain a string that identifies the upstream
	// electrical source, such as the name of a circuit or outlet, that provides
	// power to this circuit. If a value has not been assigned by a user, the value
	// of this property shall be an empty string.
	//
	// Version added: v1.4.0
	ElectricalSourceName string
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for this
	// circuit that represents the 'Total' 'ElectricalContext' sensor when multiple
	// energy sensors exist for this circuit. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// FrequencyHz shall contain the frequency, in hertz units, for this circuit.
	// The value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Frequency'.
	FrequencyHz SensorExcerpt
	// IndicatorLED shall contain the indicator light state for the indicator light
	// associated with this circuit.
	//
	// Deprecated: v1.1.0
	// This property has been deprecated in favor of the 'LocationIndicatorActive'
	// property.
	IndicatorLED IndicatorLED
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource. A write to this property shall
	// update the value of 'IndicatorLED' in this resource, if supported, to
	// reflect the implementation of the locating function.
	//
	// Version added: v1.1.0
	LocationIndicatorActive bool
	// NominalFrequencyHz shall contain the nominal frequency for this circuit, in
	// hertz units.
	//
	// Version added: v1.8.0
	NominalFrequencyHz *float64 `json:",omitempty"`
	// NominalVoltage shall contain the nominal voltage for this circuit, in volt
	// units.
	NominalVoltage NominalVoltageType
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhaseWiringType shall contain the number of ungrounded current-carrying
	// conductors (phases) and the total number of conductors (wires).
	PhaseWiringType PhaseWiringType
	// PlugType shall contain the type of physical plug used for this circuit, as
	// defined by IEC, NEMA, or regional standards.
	PlugType PlugType
	// PolyPhaseCurrentAmps shall contain the current sensors for this circuit. For
	// 3-wire circuits, this property shall contain a duplicate copy of the current
	// sensor referenced in the 'CurrentAmps' property, if present. For other
	// circuits, this property should contain multiple current sensor readings used
	// to fully describe the circuit.
	PolyPhaseCurrentAmps CurrentSensors
	// PolyPhaseEnergykWh shall contain the energy sensors for this circuit. For
	// 3-wire circuits, this property shall contain a duplicate copy of the energy
	// sensor referenced in the 'EnergykWh' property, if present. For other
	// circuits, this property should contain multiple energy sensor readings used
	// to fully describe the circuit.
	PolyPhaseEnergykWh EnergySensors
	// PolyPhasePowerWatts shall contain the power sensors for this circuit. For
	// 3-wire circuits, this property shall contain a duplicate copy of the power
	// sensor referenced in the 'PowerWatts' property, if present. For other
	// circuits, this property should contain multiple power sensor readings used
	// to fully describe the circuit.
	PolyPhasePowerWatts PowerSensors
	// PolyPhaseVoltage shall contain the voltage sensors for this circuit. For
	// 3-wire circuits, this property shall contain a duplicate copy of the voltage
	// sensor referenced in the 'Voltage' property, if present. For other circuits,
	// this property should contain multiple voltage sensor readings used to fully
	// describe the circuit.
	PolyPhaseVoltage VoltageSensors
	// PowerControlLocked shall indicate whether requests to the 'PowerControl'
	// action are locked. If 'true', services shall reject requests to the
	// 'PowerControl' action.
	//
	// Version added: v1.5.0
	PowerControlLocked bool
	// PowerCycleDelaySeconds shall contain the number of seconds to delay power on
	// after a 'PowerControl' action to cycle power. The value '0' shall indicate
	// no delay to power on.
	PowerCycleDelaySeconds *float64 `json:",omitempty"`
	// PowerEnabled shall indicate the power enable state of the circuit. The value
	// 'true' shall indicate that the circuit can be powered on, and 'false' shall
	// indicate that the circuit cannot be powered.
	PowerEnabled bool
	// PowerLoadPercent shall contain the power load, in percent units, for this
	// circuit that represents the 'Total' 'ElectricalContext' for this circuit.
	//
	// Version added: v1.3.0
	PowerLoadPercent SensorExcerpt
	// PowerOffDelaySeconds shall contain the number of seconds to delay power off
	// after a 'PowerControl' action. The value '0' shall indicate no delay to
	// power off.
	PowerOffDelaySeconds *float64 `json:",omitempty"`
	// PowerOnDelaySeconds shall contain the number of seconds to delay power up
	// after a power cycle or a 'PowerControl' action. The value '0' shall indicate
	// no delay to power up.
	PowerOnDelaySeconds *float64 `json:",omitempty"`
	// PowerRestoreDelaySeconds shall contain the number of seconds to delay power
	// on after a power fault. The value '0' shall indicate no delay to power on.
	PowerRestoreDelaySeconds *float64 `json:",omitempty"`
	// PowerRestorePolicy shall contain the desired 'PowerState' of the circuit
	// when power is applied. The value 'LastState' shall return the circuit to the
	// 'PowerState' it was in when power was lost.
	PowerRestorePolicy PowerRestorePolicyTypes
	// PowerState shall contain the power state of the circuit.
	PowerState PowerState
	// PowerStateInTransition shall indicate whether the 'PowerState' property will
	// undergo a transition between on and off states due to a configured delay.
	// The transition may be due to the configuration of the power on, off, or
	// restore delay properties. If 'true', the 'PowerState' property will
	// transition at the conclusion of a configured delay.
	//
	// Version added: v1.5.0
	PowerStateInTransition bool
	// PowerWatts shall contain the total power, in watt units, for this circuit
	// that represents the 'Total' 'ElectricalContext' sensor when multiple power
	// sensors exist for this circuit. The value of the 'DataSourceUri' property,
	// if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'Power'.
	PowerWatts SensorPowerExcerpt
	// RatedCurrentAmps shall contain the rated maximum current for this circuit,
	// in ampere units, after any required de-rating, due to safety agency or other
	// regulatory requirements, has been applied.
	RatedCurrentAmps *float64 `json:",omitempty"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UnbalancedCurrentPercent shall contain the current imbalance, in percent
	// units, between phases in a poly-phase circuit. The value of the
	// 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value 'Percent'.
	//
	// Version added: v1.5.0
	UnbalancedCurrentPercent SensorExcerpt
	// UnbalancedVoltagePercent shall contain the voltage imbalance, in percent
	// units, between phases in a poly-phase circuit. The value of the
	// 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value 'Percent'.
	//
	// Version added: v1.5.0
	UnbalancedVoltagePercent SensorExcerpt
	// UserLabel shall contain a user-assigned label used to identify this
	// resource. If a value has not been assigned by a user, the value of this
	// property shall be an empty string.
	//
	// Version added: v1.4.0
	UserLabel string
	// Voltage shall contain the voltage, in volt units, for this circuit. The
	// value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Voltage'. This property shall not be present if 'PhaseWiringType'
	// contains a value that indicates a 4-wire or greater configuration, such as
	// 'TwoPhase4Wire'.
	Voltage SensorVoltageExcerpt
	// VoltageType shall contain the type of voltage applied to the circuit.
	VoltageType VoltageType
	// breakerControlTarget is the URL to send BreakerControl requests.
	breakerControlTarget string
	// powerControlTarget is the URL to send PowerControl requests.
	powerControlTarget string
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
	// branchCircuit is the URI for BranchCircuit.
	branchCircuit string
	// distributionCircuits are the URIs for DistributionCircuits.
	distributionCircuits []string
	// outlets are the URIs for Outlets.
	outlets []string
	// powerOutlet is the URI for PowerOutlet.
	powerOutlet string
	// sourceCircuit is the URI for SourceCircuit.
	sourceCircuit string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Circuit object from the raw JSON.
func (c *Circuit) UnmarshalJSON(b []byte) error {
	type temp Circuit
	type cActions struct {
		BreakerControl ActionTarget `json:"#Circuit.BreakerControl"`
		PowerControl   ActionTarget `json:"#Circuit.PowerControl"`
		ResetMetrics   ActionTarget `json:"#Circuit.ResetMetrics"`
	}
	type cLinks struct {
		BranchCircuit        Link  `json:"BranchCircuit"`
		DistributionCircuits Links `json:"DistributionCircuits"`
		Outlets              Links `json:"Outlets"`
		PowerOutlet          Link  `json:"PowerOutlet"`
		SourceCircuit        Link  `json:"SourceCircuit"`
	}
	var tmp struct {
		temp
		Actions cActions
		Links   cLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = Circuit(tmp.temp)

	// Extract the links to other entities for later
	c.breakerControlTarget = tmp.Actions.BreakerControl.Target
	c.powerControlTarget = tmp.Actions.PowerControl.Target
	c.resetMetricsTarget = tmp.Actions.ResetMetrics.Target
	c.branchCircuit = tmp.Links.BranchCircuit.String()
	c.distributionCircuits = tmp.Links.DistributionCircuits.ToStrings()
	c.outlets = tmp.Links.Outlets.ToStrings()
	c.powerOutlet = tmp.Links.PowerOutlet.String()
	c.sourceCircuit = tmp.Links.SourceCircuit.String()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *Circuit) Update() error {
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

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetCircuit will get a Circuit instance from the service.
func GetCircuit(c Client, uri string) (*Circuit, error) {
	return GetObject[Circuit](c, uri)
}

// ListReferencedCircuits gets the collection of Circuit from
// a provided reference.
func ListReferencedCircuits(c Client, link string) ([]*Circuit, error) {
	return GetCollectionObjects[Circuit](c, link)
}

// This action shall control the state of the circuit breaker or over-current
// protection device.
// powerState - This parameter shall contain the desired power state of the
// circuit.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *Circuit) BreakerControl(powerState PowerState) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["PowerState"] = powerState
	resp, taskInfo, err := PostWithTask(c.client,
		c.breakerControlTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall control the power state of the circuit.
// powerState - This parameter shall contain the desired power state of the
// circuit.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *Circuit) PowerControl(powerState PowerState) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["PowerState"] = powerState
	resp, taskInfo, err := PostWithTask(c.client,
		c.powerControlTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset any time intervals or counted values for this
// circuit.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *Circuit) ResetMetrics() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(c.client,
		c.resetMetricsTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// BranchCircuit gets the BranchCircuit linked resource.
func (c *Circuit) BranchCircuit() (*Circuit, error) {
	if c.branchCircuit == "" {
		return nil, nil
	}
	return GetObject[Circuit](c.client, c.branchCircuit)
}

// DistributionCircuits gets the DistributionCircuits linked resources.
func (c *Circuit) DistributionCircuits() ([]*Circuit, error) {
	return GetObjects[Circuit](c.client, c.distributionCircuits)
}

// Outlets gets the Outlets linked resources.
func (c *Circuit) Outlets() ([]*Outlet, error) {
	return GetObjects[Outlet](c.client, c.outlets)
}

// PowerOutlet gets the PowerOutlet linked resource.
func (c *Circuit) PowerOutlet() (*Outlet, error) {
	if c.powerOutlet == "" {
		return nil, nil
	}
	return GetObject[Outlet](c.client, c.powerOutlet)
}

// SourceCircuit gets the SourceCircuit linked resource.
func (c *Circuit) SourceCircuit() (*Circuit, error) {
	if c.sourceCircuit == "" {
		return nil, nil
	}
	return GetObject[Circuit](c.client, c.sourceCircuit)
}

// CurrentSensors shall contain properties that describe current sensor readings
// for a circuit.
type CurrentSensors struct {
	// Line1 shall contain the line current, in ampere units, for L1. The value of
	// the 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value 'Current'.
	// This property shall not be present if the equipment does not include an L1
	// measurement.
	Line1 SensorCurrentExcerpt
	// Line2 shall contain the line current, in ampere units, for L2. The value of
	// the 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value 'Current'.
	// This property shall not be present if the equipment does not include an L2
	// measurement.
	Line2 SensorCurrentExcerpt
	// Line3 shall contain the line current, in ampere units, for L3. The value of
	// the 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value 'Current'.
	// This property shall not be present if the equipment does not include an L3
	// measurement.
	Line3 SensorCurrentExcerpt
	// Neutral shall contain the line current, in ampere units, for the Neutral
	// line. The value of the 'DataSourceUri' property, if present, shall reference
	// a resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Current'. This property shall not be present if the equipment does
	// not include a Neutral line measurement.
	Neutral SensorCurrentExcerpt
}

// EnergySensors shall contain properties that describe energy sensor readings
// for a circuit.
type EnergySensors struct {
	// Line1ToLine2 shall contain the energy, in kilowatt-hour units, between L1
	// and L2. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'EnergykWh'. This property shall not be present if the
	// equipment does not include an L1-L2 measurement.
	Line1ToLine2 SensorEnergykWhExcerpt
	// Line1ToNeutral shall contain the energy, in kilowatt-hour units, between L1
	// and Neutral. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'EnergykWh'. This property shall not be present if the
	// equipment does not include an L1-Neutral measurement.
	Line1ToNeutral SensorEnergykWhExcerpt
	// Line2ToLine3 shall contain the energy, in kilowatt-hour units, between L2
	// and L3. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'EnergykWh'. This property shall not be present if the
	// equipment does not include an L2-L3 measurement.
	Line2ToLine3 SensorEnergykWhExcerpt
	// Line2ToNeutral shall contain the energy, in kilowatt-hour units, between L2
	// and Neutral. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'EnergykWh'. This property shall not be present if the
	// equipment does not include an L2-Neutral measurement.
	Line2ToNeutral SensorEnergykWhExcerpt
	// Line3ToLine1 shall contain the energy, in kilowatt-hour units, between L3
	// and L1. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'EnergykWh'. This property shall not be present if the
	// equipment does not include an L3-L1 measurement.
	Line3ToLine1 SensorEnergykWhExcerpt
	// Line3ToNeutral shall contain the energy, in kilowatt-hour units, between L3
	// and Neutral. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'EnergykWh'. This property shall not be present if the
	// equipment does not include an L3-Neutral measurement.
	Line3ToNeutral SensorEnergykWhExcerpt
}

// PowerSensors shall contain properties that describe power sensor readings for
// a circuit.
type PowerSensors struct {
	// Line1ToLine2 shall contain the power, in watt units, between L1 and L2. The
	// value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Power'. This property shall not be present if the equipment does not
	// include an L1-L2 measurement.
	Line1ToLine2 SensorPowerExcerpt
	// Line1ToNeutral shall contain the power, in watt units, between L1 and
	// Neutral. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Power'. This property shall not be present if the
	// equipment does not include an L1-Neutral measurement.
	Line1ToNeutral SensorPowerExcerpt
	// Line2ToLine3 shall contain the power, in watt units, between L2 and L3. The
	// value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Power'. This property shall not be present if the equipment does not
	// include an L2-L3 measurement.
	Line2ToLine3 SensorPowerExcerpt
	// Line2ToNeutral shall contain the power, in watt units, between L2 and
	// Neutral. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Power'. This property shall not be present if the
	// equipment does not include an L2-Neutral measurement.
	Line2ToNeutral SensorPowerExcerpt
	// Line3ToLine1 shall contain the power, in watt units, between L3 and L1. The
	// value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Power'. This property shall not be present if the equipment does not
	// include an L3-L1 measurement.
	Line3ToLine1 SensorPowerExcerpt
	// Line3ToNeutral shall contain the power, in watt units, between L3 and
	// Neutral. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Power'. This property shall not be present if the
	// equipment does not include an L3-Neutral measurement.
	Line3ToNeutral SensorPowerExcerpt
}

// VoltageSensors shall contain properties that describe voltage sensor readings
// for a circuit.
type VoltageSensors struct {
	// Line1ToLine2 shall contain the line-to-line voltage, in volt units, between
	// L1 and L2. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Voltage'. This property shall not be present if the
	// equipment does not include an L1-L2 measurement.
	Line1ToLine2 SensorVoltageExcerpt
	// Line1ToNeutral shall contain the line-to-line voltage, in volt units,
	// between L1 and Neutral. The value of the 'DataSourceUri' property, if
	// present, shall reference a resource of type 'Sensor' with the 'ReadingType'
	// property containing the value 'Voltage'. This property shall not be present
	// if the equipment does not include an L1-Neutral measurement.
	Line1ToNeutral SensorVoltageExcerpt
	// Line2ToLine3 shall contain the line-to-line voltage, in volt units, between
	// L2 and L3. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Voltage'. This property shall not be present if the
	// equipment does not include an L2-L3 measurement.
	Line2ToLine3 SensorVoltageExcerpt
	// Line2ToNeutral shall contain the line-to-line voltage, in volt units,
	// between L2 and Neutral. The value of the 'DataSourceUri' property, if
	// present, shall reference a resource of type 'Sensor' with the 'ReadingType'
	// property containing the value 'Voltage'. This property shall not be present
	// if the equipment does not include an L2-Neutral measurement.
	Line2ToNeutral SensorVoltageExcerpt
	// Line3ToLine1 shall contain the line-to-line voltage, in volt units, between
	// L3 and L1. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Voltage'. This property shall not be present if the
	// equipment does not include an L3-L1 measurement.
	Line3ToLine1 SensorVoltageExcerpt
	// Line3ToNeutral shall contain the line-to-line voltage, in volt units,
	// between L3 and Neutral. The value of the 'DataSourceUri' property, if
	// present, shall reference a resource of type 'Sensor' with the 'ReadingType'
	// property containing the value 'Voltage'. This property shall not be present
	// if the equipment does not include an L3-Neutral measurement.
	Line3ToNeutral SensorVoltageExcerpt
}
