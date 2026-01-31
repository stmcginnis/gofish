//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.4 - #Outlet.v1_4_4.Outlet

package schemas

import (
	"encoding/json"
)

type OutletPowerState string

const (
	// OnOutletPowerState Power on.
	OnOutletPowerState OutletPowerState = "On"
	// OffOutletPowerState Power off.
	OffOutletPowerState OutletPowerState = "Off"
	// PowerCycleOutletPowerState shall indicate the resource will transition to a power
	// off state, then transition to a power on state. Upon successful completion,
	// the 'OutletPowerState' property, if supported, shall contain the value 'On'.
	PowerCycleOutletPowerState OutletPowerState = "PowerCycle"
)

type ReceptacleType string

const (
	// NEMA515RReceptacleType shall represent a receptacle that matches the NEMA
	// specified 5-15 receptacle (120V; 15A). The current is commonly de-rated to
	// 12A if it is protected by a 15A breaker.
	NEMA515RReceptacleType ReceptacleType = "NEMA_5_15R"
	// NEMA520RReceptacleType shall represent a receptacle that matches the NEMA
	// specified 5-20 receptacle that exhibits a T-slot (120V; 20A). The current is
	// commonly de-rated to 16A if it is protected by a 20A breaker.
	NEMA520RReceptacleType ReceptacleType = "NEMA_5_20R"
	// NEMAL520RReceptacleType shall represent a receptacle that matches the NEMA
	// specified locking L5-20 receptacle (120V; 20A). The current is commonly
	// de-rated to 16A if it is protected by a 20A breaker.
	NEMAL520RReceptacleType ReceptacleType = "NEMA_L5_20R"
	// NEMAL530RReceptacleType shall represent a receptacle that matches the NEMA
	// specified locking L5-30 receptacle (120V; 30A). The current is commonly
	// de-rated to 24A if it is protected by a 30A breaker.
	NEMAL530RReceptacleType ReceptacleType = "NEMA_L5_30R"
	// NEMAL620RReceptacleType shall represent a receptacle that matches the NEMA
	// specified locking L6-20 receptacle (250V; 20A). The current is commonly
	// de-rated to 16A if it is protected by a 20A breaker.
	NEMAL620RReceptacleType ReceptacleType = "NEMA_L6_20R"
	// NEMAL630RReceptacleType shall represent a receptacle that matches the NEMA
	// specified locking L6-30 receptacle (250V; 30A). The current is commonly
	// de-rated to 24A if it is protected by a 30A breaker.
	NEMAL630RReceptacleType ReceptacleType = "NEMA_L6_30R"
	// IEC60320C13ReceptacleType shall represent a receptacle that matches the IEC
	// 60320 Sheet F C13 specified receptacle (250V; 10A per IEC, 15A per UL).
	IEC60320C13ReceptacleType ReceptacleType = "IEC_60320_C13"
	// IEC60320C19ReceptacleType shall represent a receptacle that matches the IEC
	// 60320 Sheet J C19 specified receptacle (250V; 16A per IEC, 20A per UL).
	IEC60320C19ReceptacleType ReceptacleType = "IEC_60320_C19"
	// CEE7TypeEReceptacleType shall represent a receptacle that matches the French
	// specified CEE 7/7 Type E receptacle (250V; 16A).
	CEE7TypeEReceptacleType ReceptacleType = "CEE_7_Type_E"
	// CEE7TypeFReceptacleType shall represent a receptacle that matches the Schuko
	// specified CEE 7/7 Type F receptacle (250V; 16A).
	CEE7TypeFReceptacleType ReceptacleType = "CEE_7_Type_F"
	// SEV1011TYPE12ReceptacleType shall represent a receptacle that matches the
	// SEV 1011 specified Type 12 receptacle (250V; 10A).
	SEV1011TYPE12ReceptacleType ReceptacleType = "SEV_1011_TYPE_12"
	// SEV1011TYPE23ReceptacleType shall represent a receptacle that matches the
	// SEV 1011 specified Type 23 receptacle (250V; 16A).
	SEV1011TYPE23ReceptacleType ReceptacleType = "SEV_1011_TYPE_23"
	// BS1363TypeGReceptacleType shall represent a receptacle that matches the
	// British BS 1363 Type G receptacle (250V; 13A).
	BS1363TypeGReceptacleType ReceptacleType = "BS_1363_Type_G"
	// BusConnectionReceptacleType shall represent a direct connection to an
	// electrical bus.
	BusConnectionReceptacleType ReceptacleType = "BusConnection"
)

type VoltageType string

const (
	// ACVoltageType Alternating Current (AC) outlet.
	ACVoltageType VoltageType = "AC"
	// DCVoltageType Direct Current (DC) outlet.
	DCVoltageType VoltageType = "DC"
)

// Outlet shall be used to represent an electrical outlet for a Redfish
// implementation.
type Outlet struct {
	Entity
	// ConfigurationLocked shall indicate whether modification requests to this
	// resource are not permitted. If 'true', services shall reject modification
	// requests to other properties in this resource.
	//
	// Version added: v1.4.0
	ConfigurationLocked bool
	// CurrentAmps shall contain the current, in ampere units, for this outlet. The
	// value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Current'. This property shall not be present if 'PhaseWiringType'
	// contains a value that indicates a 4-wire or greater configuration, such as
	// 'TwoPhase4Wire'.
	CurrentAmps SensorCurrentExcerpt
	// ElectricalConsumerNames shall contain an array of user-assigned identifying
	// strings that describe downstream devices that are powered by this outlet.
	//
	// Version added: v1.3.0
	ElectricalConsumerNames []string
	// ElectricalContext shall contain the combination of current-carrying
	// conductors that distribute power.
	ElectricalContext ElectricalContext
	// EnergykWh shall contain the total energy, in kilowatt-hour units, for this
	// outlet that represents the 'Total' 'ElectricalContext' sensor when multiple
	// energy sensors exist for this outlet. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'EnergykWh'.
	EnergykWh SensorEnergykWhExcerpt
	// FrequencyHz shall contain the frequency, in hertz units, for this outlet.
	// The value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Frequency'.
	FrequencyHz SensorExcerpt
	// IndicatorLED shall contain the indicator light state for the indicator light
	// associated with this outlet.
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
	// NominalVoltage shall contain the nominal voltage for this outlet, in volt
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
	// OutletType shall contain the type of physical receptacle used for this
	// outlet, as defined by IEC, NEMA, or regional standards.
	OutletType ReceptacleType
	// PhaseWiringType shall contain the number of ungrounded current-carrying
	// conductors (phases) and the total number of conductors (wires).
	PhaseWiringType PhaseWiringType
	// PolyPhaseCurrentAmps shall contain the current readings for this outlet. For
	// 3-wire outlets, this property shall contain a duplicate copy of the current
	// sensor referenced in the 'CurrentAmps' property, if present. For other
	// outlets, this property should contain multiple current sensor readings used
	// to fully describe the outlet.
	PolyPhaseCurrentAmps CurrentSensors
	// PolyPhaseVoltage shall contain the voltage readings for this outlet. For
	// 3-wire outlets, this property shall contain a duplicate copy of the voltage
	// sensor referenced in the 'Voltage' property, if present. For other outlets,
	// this property should contain multiple voltage sensor readings used to fully
	// describe the outlet.
	PolyPhaseVoltage VoltageSensors
	// PowerControlLocked shall indicate whether requests to the 'PowerControl'
	// action are locked. If 'true', services shall reject requests to the
	// 'PowerControl' action.
	//
	// Version added: v1.4.0
	PowerControlLocked bool
	// PowerCycleDelaySeconds shall contain the number of seconds to delay power on
	// after a 'PowerControl' action to cycle power. The value '0' shall indicate
	// no delay to power on.
	PowerCycleDelaySeconds *float64 `json:",omitempty"`
	// PowerEnabled shall indicate the power enable state of the outlet. The value
	// 'true' shall indicate that the outlet can be powered on, and 'false' shall
	// indicate that the outlet cannot be powered.
	PowerEnabled bool
	// PowerLoadPercent shall contain the power load, in percent units, for this
	// outlet that represents the 'Total' 'ElectricalContext' for this outlet. The
	// value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Percent'.
	//
	// Version added: v1.2.0
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
	// PowerRestorePolicy shall contain the desired 'PowerState' of the outlet when
	// power is applied. The value 'LastState' shall return the outlet to the
	// 'PowerState' it was in when power was lost.
	PowerRestorePolicy PowerRestorePolicyTypes
	// PowerState shall contain the power state of the outlet.
	PowerState OutletPowerState
	// PowerStateInTransition shall indicate whether the 'PowerState' property will
	// undergo a transition between on and off states due to a configured delay.
	// The transition may be due to the configuration of the power on, off, or
	// restore delay properties. If 'true', the 'PowerState' property will
	// transition at the conclusion of a configured delay.
	//
	// Version added: v1.4.0
	PowerStateInTransition bool
	// PowerWatts shall contain the total power, in watt units, for this outlet
	// that represents the 'Total' 'ElectricalContext' sensor when multiple power
	// sensors exist for this outlet. The value of the 'DataSourceUri' property, if
	// present, shall reference a resource of type 'Sensor' with the 'ReadingType'
	// property containing the value 'Power'.
	PowerWatts SensorPowerExcerpt
	// RatedCurrentAmps shall contain the rated maximum current for this outlet, in
	// ampere units, after any required de-rating, due to safety agency or other
	// regulatory requirements, has been applied.
	RatedCurrentAmps *float64 `json:",omitempty"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UserLabel shall contain a user-assigned label used to identify this
	// resource. If a value has not been assigned by a user, the value of this
	// property shall be an empty string.
	//
	// Version added: v1.3.0
	UserLabel string
	// Voltage shall contain the voltage, in volt units, for this outlet. The value
	// of the 'DataSourceUri' property, if present, shall reference a resource of
	// type 'Sensor' with the 'ReadingType' property containing the value
	// 'Voltage'. This property shall not be present if 'PhaseWiringType' contains
	// a value that indicates a 4-wire or greater configuration, such as
	// 'TwoPhase4Wire'.
	Voltage SensorVoltageExcerpt
	// VoltageType shall contain the type of voltage applied to the outlet.
	VoltageType VoltageType
	// powerControlTarget is the URL to send PowerControl requests.
	powerControlTarget string
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
	// branchCircuit is the URI for BranchCircuit.
	branchCircuit string
	// chassis are the URIs for Chassis.
	chassis []string
	// distributionCircuits are the URIs for DistributionCircuits.
	distributionCircuits []string
	// powerSupplies are the URIs for PowerSupplies.
	powerSupplies []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Outlet object from the raw JSON.
func (o *Outlet) UnmarshalJSON(b []byte) error {
	type temp Outlet
	type oActions struct {
		PowerControl ActionTarget `json:"#Outlet.PowerControl"`
		ResetMetrics ActionTarget `json:"#Outlet.ResetMetrics"`
	}
	type oLinks struct {
		BranchCircuit        Link  `json:"BranchCircuit"`
		Chassis              Links `json:"Chassis"`
		DistributionCircuits Links `json:"DistributionCircuits"`
		PowerSupplies        Links `json:"PowerSupplies"`
	}
	var tmp struct {
		temp
		Actions oActions
		Links   oLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*o = Outlet(tmp.temp)

	// Extract the links to other entities for later
	o.powerControlTarget = tmp.Actions.PowerControl.Target
	o.resetMetricsTarget = tmp.Actions.ResetMetrics.Target
	o.branchCircuit = tmp.Links.BranchCircuit.String()
	o.chassis = tmp.Links.Chassis.ToStrings()
	o.distributionCircuits = tmp.Links.DistributionCircuits.ToStrings()
	o.powerSupplies = tmp.Links.PowerSupplies.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	o.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (o *Outlet) Update() error {
	readWriteFields := []string{
		"ConfigurationLocked",
		"ElectricalConsumerNames",
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

	return o.UpdateFromRawData(o, o.RawData, readWriteFields)
}

// GetOutlet will get a Outlet instance from the service.
func GetOutlet(c Client, uri string) (*Outlet, error) {
	return GetObject[Outlet](c, uri)
}

// ListReferencedOutlets gets the collection of Outlet from
// a provided reference.
func ListReferencedOutlets(c Client, link string) ([]*Outlet, error) {
	return GetCollectionObjects[Outlet](c, link)
}

// This action shall control the power state of the outlet.
// powerState - This parameter shall contain the desired power state of the
// outlet.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (o *Outlet) PowerControl(powerState OutletPowerState) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["OutletPowerState"] = powerState
	resp, taskInfo, err := PostWithTask(o.client,
		o.powerControlTarget, payload, o.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset any time intervals or counted values for this
// outlet.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (o *Outlet) ResetMetrics() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(o.client,
		o.resetMetricsTarget, payload, o.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// BranchCircuit gets the BranchCircuit linked resource.
func (o *Outlet) BranchCircuit() (*Circuit, error) {
	if o.branchCircuit == "" {
		return nil, nil
	}
	return GetObject[Circuit](o.client, o.branchCircuit)
}

// Chassis gets the Chassis linked resources.
func (o *Outlet) Chassis() ([]*Chassis, error) {
	return GetObjects[Chassis](o.client, o.chassis)
}

// DistributionCircuits gets the DistributionCircuits linked resources.
func (o *Outlet) DistributionCircuits() ([]*Circuit, error) {
	return GetObjects[Circuit](o.client, o.distributionCircuits)
}

// PowerSupplies gets the PowerSupplies linked resources.
func (o *Outlet) PowerSupplies() ([]*PowerSupply, error) {
	return GetObjects[PowerSupply](o.client, o.powerSupplies)
}
