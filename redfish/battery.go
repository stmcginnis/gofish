//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ChargeState is the current state of the batter charge.
type ChargeState string

const (
	// IdleChargeState shall indicate the battery is idle and energy is not entering or leaving the battery. Small
	// amounts of energy may enter or leave the battery while in this state if the battery is regulating itself.
	IdleChargeState ChargeState = "Idle"
	// ChargingChargeState shall indicate the battery is charging and energy is entering the battery.
	ChargingChargeState ChargeState = "Charging"
	// DischargingChargeState shall indicate the battery is discharging and energy is leaving the battery.
	DischargingChargeState ChargeState = "Discharging"
)

// Battery shall represent a battery for a Redfish implementation. It may also represent a location, such as a
// slot, socket, or bay, where a unit may be installed if the State property within the Status property contains
// 'Absent'.
type Battery struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Assembly shall contain a link to a resource of type Assembly.
	assembly string
	// CapacityActualAmpHours shall contain the actual maximum capacity of this battery in amp-hour units.
	CapacityActualAmpHours float64
	// CapacityActualWattHours shall contain the actual maximum capacity of this battery in watt-hour units.
	CapacityActualWattHours float64
	// CapacityRatedAmpHours shall contain the rated maximum capacity of this battery in amp-hour units.
	CapacityRatedAmpHours float64
	// CapacityRatedWattHours shall contain the rated maximum capacity of this battery in watt-hour units.
	CapacityRatedWattHours float64
	// ChargeState shall contain the charge state of this battery.
	ChargeState ChargeState
	// Description provides a description of this resource.
	Description string
	// FirmwareVersion shall contain the firmware version as defined by the manufacturer for this battery.
	FirmwareVersion string
	// HotPluggable shall indicate whether the device can be inserted or removed while the underlying equipment
	// otherwise remains in its current operational state. Devices indicated as hot-pluggable shall allow the device to
	// become operable without altering the operational state of the underlying equipment. Devices that cannot be
	// inserted or removed from equipment in operation, or devices that cannot become operable without affecting the
	// operational state of that equipment, shall be indicated as not hot-pluggable.
	HotPluggable bool
	// Location shall contain the location information of this battery.
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for producing the battery. This organization
	// may be the entity from whom the battery is purchased, but this is not necessarily true.
	Manufacturer string
	// MaxChargeRateAmps shall contain the maximum charge rate of this battery in amp units.
	MaxChargeRateAmps float64
	// MaxChargeVoltage shall contain the maximum charge voltage of this battery.
	MaxChargeVoltage float64
	// MaxDischargeRateAmps shall contain the maximum discharge rate of this battery in amp units.
	MaxDischargeRateAmps float64
	// Metrics shall contain a link to a resource of type BatteryMetrics.
	metrics string
	// Model shall contain the model information as defined by the manufacturer for this battery.
	Model string
	// PartNumber shall contain the part number as defined by the manufacturer for this battery.
	PartNumber string
	// ProductionDate shall contain the date of production or manufacture for this battery.
	ProductionDate string
	// Replaceable shall indicate whether this component can be independently replaced as allowed by the vendor's
	// replacement policy. A value of 'false' indicates the component needs to be replaced by policy as part of another
	// component. If the 'LocationType' property of this component contains 'Embedded', this property shall contain
	// 'false'.
	Replaceable bool
	// SerialNumber shall contain the serial number as defined by the manufacturer for this battery.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as defined by the manufacturer for this
	// battery.
	SparePartNumber string
	// StateOfHealthPercent shall contain the state of health, in percent units, of this battery. The value of the
	// DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType property
	// containing the value 'Percent'.
	StateOfHealthPercent SensorExcerpt
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Version shall contain the hardware version of this battery as determined by the vendor or supplier.
	Version string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
	// memory is an array of links to Memory devices.
	memory []string
	// MemoryCount is the number of associated Memory devices.
	MemoryCount int
	// storageControllers is an array of links to associated StorageControllers.
	storageControllers []string
	// StorageControllersCount is the number of associated StorageControllers.
	StorageControllersCount int
	calibrateTarget         string
	resetTarget             string
	selfTestTarget          string
}

// UnmarshalJSON unmarshals a Battery object from the raw JSON.
func (battery *Battery) UnmarshalJSON(b []byte) error {
	type temp Battery
	type Links struct {
		// Memory shall contain an array of links to resources of type Memory that represent the memory devices to which
		// this battery provides power during a power-loss event, such as battery-backed NVDIMMs. This property shall not
		// be present if the battery powers the containing chassis as a whole rather than individual components in a
		// chassis.
		Memory []string
		// Memory@odata.count
		MemoryCount int `json:"Memory@odata.count"`
		// StorageControllers shall contain an array of links to resources of type StorageController that represent the
		// storage controllers to which this battery provides power during a power-loss event, such as battery-backed RAID
		// controllers. This property shall not be present if the battery powers the containing chassis as a whole rather
		// than individual components in a chassis.
		StorageControllers []string
		// StorageControllers@odata.count
		StorageControllersCount int `json:"StorageControllers@odata.count"`
	}
	type Actions struct {
		BatteryCalibrate common.ActionTarget `json:"#Battery.Calibrate"`
		BatteryReset     common.ActionTarget `json:"#Battery.Reset"`
		BatterySelfTest  common.ActionTarget `json:"#Battery.SelfTest"`
	}
	var t struct {
		temp
		Assembly common.Link
		Metrics  common.Link
		Links    Links
		Actions  Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*battery = Battery(t.temp)

	// Extract the links to other entities for later
	battery.assembly = t.Assembly.String()
	battery.metrics = t.Metrics.String()

	battery.memory = t.Links.Memory
	battery.MemoryCount = t.Links.MemoryCount
	battery.storageControllers = t.Links.StorageControllers
	battery.StorageControllersCount = t.Links.StorageControllersCount
	battery.calibrateTarget = t.Actions.BatteryCalibrate.Target
	battery.resetTarget = t.Actions.BatteryReset.Target
	battery.selfTestTarget = t.Actions.BatterySelfTest.Target

	// This is a read/write object, so we need to save the raw object data for later
	battery.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (battery *Battery) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Battery)
	_ = original.UnmarshalJSON(battery.rawData)

	readWriteFields := []string{
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(battery).Elem()

	return battery.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetBattery will get a Battery instance from the service.
func GetBattery(c common.Client, uri string) (*Battery, error) {
	return common.GetObject[Battery](c, uri)
}

// ListReferencedBatterys gets the collection of Battery from
// a provided reference.
func ListReferencedBatterys(c common.Client, link string) ([]*Battery, error) {
	return common.GetCollectionObjects[Battery](c, link)
}

// Assembly get the containing assembly of this battery.
func (battery *Battery) Assembly() (*Assembly, error) {
	if battery.assembly == "" {
		return nil, nil
	}
	return GetAssembly(battery.GetClient(), battery.assembly)
}

// BatteryMetrics get the metrics for this battery.
func (battery *Battery) BatteryMetrics() (*BatteryMetrics, error) {
	if battery.metrics == "" {
		return nil, nil
	}
	return GetBatteryMetrics(battery.GetClient(), battery.metrics)
}

// Memory returns a collection of Memory devices associated with this Battery.
func (battery *Battery) Memory() ([]*Memory, error) {
	return common.GetObjects[Memory](battery.GetClient(), battery.memory)
}

// StorageControllers returns a collection of StorageControllers associated with this Battery.
func (battery *Battery) StorageControllers() ([]*StorageController, error) {
	return common.GetObjects[StorageController](battery.GetClient(), battery.storageControllers)
}

// Calibrate performs a self-calibration, or learn cycle, of the battery.
func (battery *Battery) Calibrate() error {
	payload := struct{}{}
	return battery.Post(battery.calibrateTarget, payload)
}

// Reset resets the battery.
func (battery *Battery) Reset(resetType ResetType) error {
	t := struct {
		ResetType ResetType
	}{ResetType: resetType}
	return battery.Post(battery.resetTarget, t)
}

// SelfTest performs a self-test of the battery.
func (battery *Battery) SelfTest() error {
	payload := struct{}{}
	return battery.Post(battery.selfTestTarget, payload)
}
