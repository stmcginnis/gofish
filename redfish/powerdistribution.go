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

// The type of equipment this resource represents.
type PowerEquipmentType string

const (
	// RackPDUPowerEquipmentType A power distribution unit providing outlets for a rack or similar quantity of devices.
	RackPDUPowerEquipmentType PowerEquipmentType = "RackPDU"
	// FloorPDUPowerEquipmentType A power distribution unit providing feeder circuits for further power distribution.
	FloorPDUPowerEquipmentType PowerEquipmentType = "FloorPDU"
	// ManualTransferSwitchPowerEquipmentType A manual power transfer switch.
	ManualTransferSwitchPowerEquipmentType PowerEquipmentType = "ManualTransferSwitch"
	// AutomaticTransferSwitchPowerEquipmentType An automatic power transfer switch.
	AutomaticTransferSwitchPowerEquipmentType PowerEquipmentType = "AutomaticTransferSwitch"
	// SwitchgearPowerEquipmentType Electrical switchgear.
	SwitchgearPowerEquipmentType PowerEquipmentType = "Switchgear"
	// PowerShelfPowerEquipmentType A power shelf.
	PowerShelfPowerEquipmentType PowerEquipmentType = "PowerShelf"
	// BusPowerEquipmentType An electrical bus.
	BusPowerEquipmentType PowerEquipmentType = "Bus"
	// BatteryShelfPowerEquipmentType A battery shelf or battery-backed unit (BBU).
	BatteryShelfPowerEquipmentType PowerEquipmentType = "BatteryShelf"
)

// The sensitivity to voltage waveform quality to satisfy the criterion for initiating a transfer.
type TransferSensitivity string

const (
	// High sensitivity for initiating a transfer.
	HighTransferSensitivity TransferSensitivity = "High"
	// Low sensitivity for initiating a transfer.
	LowTransferSensitivity TransferSensitivity = "Low"
	// Medium sensitivity for initiating a transfer.
	MediumTransferSensitivity TransferSensitivity = "Medium"
)

// The configuration settings for an automatic transfer switch.
type TransferConfiguration struct {
	// The mains circuit that is switched on and qualified
	// to supply power to the output circuit.
	ActiveMainsID string `json:"ActiveMainsId"`
	// Indicates if the qualified alternate mains circuit is automatically switched on
	// when the preferred mains circuit becomes unqualified and is automatically switched off.
	AutoTransferEnabled bool
	// Indicates if a make-before-break switching sequence of the mains circuits is permitted
	// when they are both qualified and in synchronization.
	ClosedTransitionAllowed bool
	// The time in seconds to wait for a closed transition to occur.
	ClosedTransitionTimeoutSeconds int64
	// The preferred source for the mains circuit to this equipment.
	PreferredMainsID string `json:"PreferredMainsId"`
	// The time in seconds to delay the automatic transfer
	// from the alternate mains circuit back to the preferred mains circuit.
	RetransferDelaySeconds int64
	// Indicates if the automatic transfer is permitted from the alternate mains circuit
	// back to the preferred mains circuit after the preferred mains circuit is qualified again
	// and the Retransfer Delay time has expired.
	RetransferEnabled bool
	// The time in seconds to delay the automatic transfer from the preferred mains circuit
	// to the alternate mains circuit when the preferred mains circuit is disqualified.
	TransferDelaySeconds int64
	// Indicates if any transfer is inhibited.
	TransferInhibit bool
}

// The criteria used to initiate a transfer for an automatic transfer switch.
type TransferCriteria struct {
	// The frequency in hertz over the nominal value
	// that satisfies a criterion for transfer.
	OverNominalFrequencyHz float32
	// The positive percentage of voltage RMS over the nominal value
	// that satisfies a criterion for transfer.
	OverVoltageRMSPercentage int
	// The sensitivity to voltage waveform quality
	// to satisfy the criterion for initiating a transfer.
	TransferSensitivity TransferSensitivity
	// The frequency in hertz under the nominal value
	// that satisfies a criterion for transfer.
	UnderNominalFrequencyHz float32
	// The negative percentage of voltage RMS under the nominal value
	// that satisfies a criterion for transfer.
	UnderVoltageRMSPercentage int
}

// PowerDistribution shall be used to represent
// a power distribution component or unit for a Redfish implementation.
type PowerDistribution struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string

	// The user-assigned asset tag for this equipment.
	AssetTag string
	// A link to the branch circuits for this equipment.
	branches string
	// The type of equipment this resource represents.
	EquipmentType PowerEquipmentType
	// A link to the feeder circuits for this equipment.
	feeders string
	// The firmware version of this equipment.
	FirmwareVersion string
	// The location of the equipment.
	Location common.Location
	// A link to the power input circuits for this equipment.
	mains string
	// The redundancy information for the mains (input) circuits for this equipment.
	MainsRedundancy RedundantGroup
	// The manufacturer of this equipment.
	Manufacturer string
	// A link to the summary metrics for this equipment.
	metrics string
	// The product model number of this equipment.
	Model string
	// A link to the outlet groups for this equipment.
	outletGroups string
	// A link to the outlets for this equipment.
	outlets string
	// The part number for this equipment.
	PartNumber string
	// Deprecated: (v1.3) The link to the collection of power supplies for this equipment.
	// This property has been deprecated in favor of the PowerSupplies link in the Chassis resource.
	powerSupplies string
	// Deprecated: (v1.3) The redundancy information for the devices in a redundancy group.
	// This property has been deprecated in favor of the PowerSupplyRedundancy property in the Chassis resource.
	PowerSupplyRedundancy []RedundantGroup
	// The production or manufacturing date of this equipment.
	ProductionDate string
	// Deprecated: (v1.3) A link to the collection of sensors located in the equipment and sub-components.
	// This property has been deprecated in favor of the Sensors link in the Chassis resource.
	sensors string
	// The serial number for this equipment.
	SerialNumber string
	// The status and health of the resource and its subordinate or dependent resources.
	Status common.Status
	// A link to the subfeed circuits for this equipment.
	subfeeds string
	// The configuration settings for an automatic transfer switch.
	TransferConfiguration TransferConfiguration
	// The criteria used to initiate a transfer for an automatic transfer switch.
	TransferCriteria TransferCriteria
	// A user-assigned label.
	UserLabel string
	// The UUID for this equipment.
	UUID string
	// The hardware version of this equipment.
	Version string

	// Links section
	// An array of links to the chassis that contain this equipment.
	chassis      []string
	ChassisCount int
	// A link to the facility that contains this equipment.
	facility string
	// An array of links to the managers responsible for managing this equipment.
	managedBy      []string
	ManagedByCount int
	// OemLinks are all OEM data under link section
	OemLinks json.RawMessage

	// Actions section
	// This action transfers control to the alternative input circuit.
	transferControlTarget string
	// OemActions contains all the vendor specific actions.
	// It is vendor responsibility to parse this field accordingly
	OemActions json.RawMessage

	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PowerDistribution object from the raw JSON.
func (powerDistribution *PowerDistribution) UnmarshalJSON(b []byte) error {
	type temp PowerDistribution
	type linkReference struct {
		Chassis        common.Links
		ChassisCount   int `json:"Chassis@odata.count"`
		Facility       common.Link
		ManagedBy      common.Links
		ManagedByCount int `json:"ManagedBy@odata.count"`
		Oem            json.RawMessage
	}
	type actions struct {
		TransferControl common.ActionTarget `json:"#PowerDistribution.TransferControl"`
		Oem             json.RawMessage     // OEM actions will be stored here
	}
	var t struct {
		temp

		Branches      common.Link
		Feeders       common.Link
		Mains         common.Link
		Metrics       common.Link
		OutletGroups  common.Link
		Outlets       common.Link
		PowerSupplies common.Link
		Sensors       common.Link
		Subfeeds      common.Link

		Links   linkReference
		Actions actions
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	// Extract the links to other entities for later
	*powerDistribution = PowerDistribution(t.temp)
	powerDistribution.branches = t.Branches.String()
	powerDistribution.feeders = t.Feeders.String()
	powerDistribution.mains = t.Mains.String()
	powerDistribution.metrics = t.Metrics.String()
	powerDistribution.outletGroups = t.OutletGroups.String()
	powerDistribution.outlets = t.Outlets.String()
	powerDistribution.powerSupplies = t.PowerSupplies.String()
	powerDistribution.sensors = t.Sensors.String()
	powerDistribution.subfeeds = t.Subfeeds.String()

	powerDistribution.chassis = t.Links.Chassis.ToStrings()
	powerDistribution.ChassisCount = t.Links.ChassisCount
	powerDistribution.facility = t.Links.Facility.String()
	powerDistribution.managedBy = t.Links.ManagedBy.ToStrings()
	powerDistribution.ManagedByCount = t.Links.ManagedByCount
	powerDistribution.OemLinks = t.Links.Oem

	powerDistribution.transferControlTarget = t.Actions.TransferControl.Target
	powerDistribution.OemActions = t.Actions.Oem

	// This is a read/write object, so we need to save the raw object data for later
	powerDistribution.rawData = b

	return nil
}

// GetPowerDistribution will get a PowerDistribution instance from the Redfish service.
func GetPowerDistribution(c common.Client, uri string) (*PowerDistribution, error) {
	return common.GetObject[PowerDistribution](c, uri)
}

// Update commits updates to this object's properties to the running system.
func (powerDistribution *PowerDistribution) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	pd := new(PowerDistribution)
	err := pd.UnmarshalJSON(powerDistribution.rawData)
	if err != nil {
		return err
	}

	// Note: current definition (2023.3) only includes AssetTag and UserLabel.
	// May have errors trying to set other values, but keeping in here for backwards
	// compatibility.
	readWriteFields := []string{
		"AssetTag",
		"UserLabel",
		"ActiveMainsId",
		"AutoTransferEnabled",
		"ClosedTransitionAllowed",
		"ClosedTransitionTimeoutSeconds",
		"PreferredMainsId",
		"RetransferDelaySeconds",
		"RetransferEnabled",
		"TransferDelaySeconds",
		"TransferInhibit",
		"OverNominalFrequencyHz",
		"OverVoltageRMSPercentage",
		"TransferSensitivity",
		"UnderNominalFrequencyHz",
		"UnderVoltageRMSPercentage",
	}

	originalElement := reflect.ValueOf(pd).Elem()
	currentElement := reflect.ValueOf(powerDistribution).Elem()

	return powerDistribution.Entity.Update(originalElement, currentElement, readWriteFields)
}

// This action shall transfer power input from the existing mains circuit to the alternative mains circuit.
func (powerDistribution *PowerDistribution) TransferControl() error {
	if powerDistribution.transferControlTarget == "" {
		return errors.New("TransferControl is not supported") //nolint:golint
	}

	return powerDistribution.Post(powerDistribution.transferControlTarget, nil)
}

// ListReferencedPowerDistribution gets the collection of PowerDistribution from
// a provided reference.
func ListReferencedPowerDistributionUnits(c common.Client, link string) ([]*PowerDistribution, error) {
	return common.GetCollectionObjects[PowerDistribution](c, link)
}

// Deprecated: (v1.3) in favor of the Sensors link in the Chassis resource.
func (powerDistribution *PowerDistribution) Sensors() ([]*Sensor, error) {
	return ListReferencedSensors(powerDistribution.GetClient(), powerDistribution.sensors)
}

// Deprecated: (v1.3) in favor of the PowerSupplies link in the Chassis resource.
func (powerDistribution *PowerDistribution) PowerSupplies() ([]*PowerSupplyUnit, error) {
	return ListReferencedPowerSupplyUnits(powerDistribution.GetClient(), powerDistribution.powerSupplies)
}

// ManagedBy gets the collection of managers for this equipment.
func (powerDistribution *PowerDistribution) ManagedBy() ([]*Manager, error) {
	return common.GetObjects[Manager](powerDistribution.GetClient(), powerDistribution.managedBy)
}

// Chassis gets the collection of chassis for this equipment.
func (powerDistribution *PowerDistribution) Chassis() ([]*Chassis, error) {
	return common.GetObjects[Chassis](powerDistribution.GetClient(), powerDistribution.chassis)
}

// Branches gets the collection that contains the branch circuits for this equipment.
func (powerDistribution *PowerDistribution) Branches() ([]*Circuit, error) {
	return ListReferencedCircuits(powerDistribution.GetClient(), powerDistribution.branches)
}

// Feeders gets the collection that contains the feeder circuits for this equipment.
func (powerDistribution *PowerDistribution) Feeders() ([]*Circuit, error) {
	return ListReferencedCircuits(powerDistribution.GetClient(), powerDistribution.feeders)
}

// Mains gets the collection that contains the power input circuits for this equipment.
func (powerDistribution *PowerDistribution) Mains() ([]*Circuit, error) {
	return ListReferencedCircuits(powerDistribution.GetClient(), powerDistribution.mains)
}

// Subfeeds gets the collection that contains the subfeed circuits for this equipment.
func (powerDistribution *PowerDistribution) Subfeeds() ([]*Circuit, error) {
	return ListReferencedCircuits(powerDistribution.GetClient(), powerDistribution.subfeeds)
}

// Facility gets a resource that represents the facility that contains this equipment.
func (powerDistribution *PowerDistribution) Facility() (*Facility, error) {
	return GetFacility(powerDistribution.GetClient(), powerDistribution.facility)
}

// Metrics gets the metrics of a power distribution component or unit.
func (powerDistribution *PowerDistribution) Metrics() (metrics *PowerDistributionMetrics, err error) {
	if powerDistribution.metrics == "" {
		return
	}
	return GetPowerDistributionMetrics(powerDistribution.GetClient(), powerDistribution.metrics)
}

// OutletGroups gets the collection that contains the outlet groups for this equipment.
func (powerDistribution *PowerDistribution) OutletGroups() ([]*OutletGroup, error) {
	return ListReferencedOutletGroups(powerDistribution.GetClient(), powerDistribution.outletGroups)
}

// Outlets gets the collection that contains the outlets for this equipment.
func (powerDistribution *PowerDistribution) Outlets() ([]*Outlet, error) {
	return ListReferencedOutlets(powerDistribution.GetClient(), powerDistribution.outlets)
}
