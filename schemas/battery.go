//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #Battery.v1_5_0.Battery

package schemas

import (
	"encoding/json"
)

type BatteryChemistryType string

const (
	// LeadAcidBatteryChemistryType shall indicate that the battery stores chemical
	// energy using lead and acid electrochemical reactions.
	LeadAcidBatteryChemistryType BatteryChemistryType = "LeadAcid"
	// LithiumIonBatteryChemistryType shall indicate that the battery stores
	// chemical energy using lithium ion intercalation.
	LithiumIonBatteryChemistryType BatteryChemistryType = "LithiumIon"
	// NickelCadmiumBatteryChemistryType shall indicate that the battery stores
	// chemical energy using nickel and cadmium electrochemical reactions.
	NickelCadmiumBatteryChemistryType BatteryChemistryType = "NickelCadmium"
)

type ChargeState string

const (
	// IdleChargeState shall indicate the battery is idle and energy is not
	// entering or leaving the battery. Small amounts of energy may enter or leave
	// the battery while in this state if the battery is regulating itself.
	IdleChargeState ChargeState = "Idle"
	// ChargingChargeState shall indicate the battery is charging and energy is
	// entering the battery.
	ChargingChargeState ChargeState = "Charging"
	// DischargingChargeState shall indicate the battery is discharging and energy
	// is leaving the battery.
	DischargingChargeState ChargeState = "Discharging"
)

type EnergyStorageType string

const (
	// BatteryEnergyStorageType shall indicate that the battery stores energy using
	// one or more electrochemical cells.
	BatteryEnergyStorageType EnergyStorageType = "Battery"
	// SupercapacitorEnergyStorageType shall indicate that the battery stores
	// energy using electrostatic double-layer capacitors or electrodes with
	// electrochemical pseudocapacitance.
	SupercapacitorEnergyStorageType EnergyStorageType = "Supercapacitor"
)

// Battery shall represent an energy storage device for a Redfish
// implementation. It may also represent a location, such as a slot, socket, or
// bay, where a unit may be installed if the 'State' property within the
// 'Status' property contains 'Absent'.
type Battery struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	assembly string
	// BatteryChemistryType shall contain the chemistry of the battery. This
	// property shall only be present if the 'EnergyStorageType' property contains
	// 'Battery'.
	//
	// Version added: v1.4.0
	BatteryChemistryType BatteryChemistryType
	// CapacityActualAmpHours shall contain the actual maximum capacity of this
	// battery in amp-hour units.
	CapacityActualAmpHours *float64 `json:",omitempty"`
	// CapacityActualWattHours shall contain the actual maximum capacity of this
	// battery in watt-hour units.
	CapacityActualWattHours *float64 `json:",omitempty"`
	// CapacityRatedAmpHours shall contain the rated maximum capacity of this
	// battery in amp-hour units.
	CapacityRatedAmpHours *float64 `json:",omitempty"`
	// CapacityRatedWattHours shall contain the rated maximum capacity of this
	// battery in watt-hour units.
	CapacityRatedWattHours *float64 `json:",omitempty"`
	// ChargeState shall contain the charge state of this battery.
	ChargeState ChargeState
	// EnergyStorageType shall contain the energy storage technology used in the
	// battery.
	//
	// Version added: v1.4.0
	EnergyStorageType EnergyStorageType
	// FirmwareVersion shall contain the firmware version as defined by the
	// manufacturer for this battery.
	FirmwareVersion string
	// HotPluggable shall indicate whether the device can be inserted or removed
	// while the underlying equipment otherwise remains in its current operational
	// state. Devices indicated as hot-pluggable shall allow the device to become
	// operable without altering the operational state of the underlying equipment.
	// Devices that cannot be inserted or removed from equipment in operation, or
	// devices that cannot become operable without affecting the operational state
	// of that equipment, shall be indicated as not hot-pluggable.
	HotPluggable bool
	// Location shall contain the location information of this battery.
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for
	// producing the battery. This organization may be the entity from whom the
	// battery is purchased, but this is not necessarily true.
	Manufacturer string
	// MaxChargeRateAmps shall contain the maximum charge rate at the input of this
	// battery in amp units.
	MaxChargeRateAmps *float64 `json:",omitempty"`
	// MaxChargeVoltage shall contain the maximum charge voltage across the cell
	// pack of this battery when it is fully charged. This property should not be
	// present if the battery contains an internal charger that regulates the
	// voltage applied to the cell pack from the input of the battery.
	MaxChargeVoltage *float64 `json:",omitempty"`
	// MaxDischargeRateAmps shall contain the maximum discharge rate at the output
	// of this battery in amp units.
	MaxDischargeRateAmps *float64 `json:",omitempty"`
	// Metrics shall contain a link to a resource of type 'BatteryMetrics'.
	metrics string
	// Model shall contain the model information as defined by the manufacturer for
	// this battery.
	Model string
	// NominalOutputVoltage shall contain the nominal output voltage of this
	// battery.
	//
	// Version added: v1.3.0
	NominalOutputVoltage *float64 `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for
	// this battery.
	PartNumber string
	// ProductionDate shall contain the date of production or manufacture for this
	// battery.
	ProductionDate string
	// Replaceable shall indicate whether this component can be independently
	// replaced as allowed by the vendor's replacement policy. A value of 'false'
	// indicates the component needs to be replaced by policy as part of another
	// component. If the 'LocationType' property of this component contains
	// 'Embedded', this property shall contain 'false'.
	//
	// Version added: v1.2.0
	Replaceable bool
	// SerialNumber shall contain the serial number as defined by the manufacturer
	// for this battery.
	SerialNumber string
	// ServicedDate shall contain the date the battery was put into active service.
	//
	// Version added: v1.5.0
	ServicedDate string
	// SparePartNumber shall contain the spare or replacement part number as
	// defined by the manufacturer for this battery.
	SparePartNumber string
	// StateOfHealthPercent shall contain the state of health, in percent units, of
	// this battery. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Percent'.
	StateOfHealthPercent SensorExcerpt
	// Status shall contain any status or health properties of the resource.
	Status Status
	// Version shall contain the hardware version of this battery as determined by
	// the vendor or supplier.
	Version string
	// calibrateTarget is the URL to send Calibrate requests.
	calibrateTarget string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// selfTestTarget is the URL to send SelfTest requests.
	selfTestTarget string
	// memory are the URIs for Memory.
	memory []string
	// storageControllers are the URIs for StorageControllers.
	storageControllers []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Battery object from the raw JSON.
func (ba *Battery) UnmarshalJSON(b []byte) error {
	type temp Battery
	type baActions struct {
		Calibrate ActionTarget `json:"#Battery.Calibrate"`
		Reset     ActionTarget `json:"#Battery.Reset"`
		SelfTest  ActionTarget `json:"#Battery.SelfTest"`
	}
	type baLinks struct {
		Memory             Links `json:"Memory"`
		StorageControllers Links `json:"StorageControllers"`
	}
	var tmp struct {
		temp
		Actions  baActions
		Links    baLinks
		Assembly Link `json:"Assembly"`
		Metrics  Link `json:"Metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*ba = Battery(tmp.temp)

	// Extract the links to other entities for later
	ba.calibrateTarget = tmp.Actions.Calibrate.Target
	ba.resetTarget = tmp.Actions.Reset.Target
	ba.selfTestTarget = tmp.Actions.SelfTest.Target
	ba.memory = tmp.Links.Memory.ToStrings()
	ba.storageControllers = tmp.Links.StorageControllers.ToStrings()
	ba.assembly = tmp.Assembly.String()
	ba.metrics = tmp.Metrics.String()

	// This is a read/write object, so we need to save the raw object data for later
	ba.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (ba *Battery) Update() error {
	readWriteFields := []string{
		"LocationIndicatorActive",
		"ServicedDate",
	}

	return ba.UpdateFromRawData(ba, ba.RawData, readWriteFields)
}

// GetBattery will get a Battery instance from the service.
func GetBattery(c Client, uri string) (*Battery, error) {
	return GetObject[Battery](c, uri)
}

// ListReferencedBatterys gets the collection of Battery from
// a provided reference.
func ListReferencedBatterys(c Client, link string) ([]*Battery, error) {
	return GetCollectionObjects[Battery](c, link)
}

// This action shall perform a self-calibration, or learn cycle, of the
// battery.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (ba *Battery) Calibrate() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(ba.client,
		ba.calibrateTarget, payload, ba.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset the battery.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and shall perform a
// 'GracefulRestart'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (ba *Battery) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(ba.client,
		ba.resetTarget, payload, ba.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall perform a self-test of the battery.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (ba *Battery) SelfTest() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(ba.client,
		ba.selfTestTarget, payload, ba.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Memory gets the Memory linked resources.
func (ba *Battery) Memory() ([]*Memory, error) {
	return GetObjects[Memory](ba.client, ba.memory)
}

// StorageControllers gets the StorageControllers linked resources.
func (ba *Battery) StorageControllers() ([]*StorageController, error) {
	return GetObjects[StorageController](ba.client, ba.storageControllers)
}

// Assembly gets the Assembly linked resource.
func (ba *Battery) Assembly() (*Assembly, error) {
	if ba.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](ba.client, ba.assembly)
}

// Metrics gets the Metrics linked resource.
func (ba *Battery) Metrics() (*BatteryMetrics, error) {
	if ba.metrics == "" {
		return nil, nil
	}
	return GetObject[BatteryMetrics](ba.client, ba.metrics)
}
