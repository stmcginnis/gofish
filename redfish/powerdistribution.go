//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #PowerDistribution.v1_5_0.PowerDistribution

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type PowerEquipmentType string

const (
	// RackPDUPowerEquipmentType is a power distribution unit providing outlets for
	// a rack or similar quantity of devices.
	RackPDUPowerEquipmentType PowerEquipmentType = "RackPDU"
	// FloorPDUPowerEquipmentType is a power distribution unit providing feeder
	// circuits for further power distribution.
	FloorPDUPowerEquipmentType PowerEquipmentType = "FloorPDU"
	// ManualTransferSwitchPowerEquipmentType is a manual power transfer switch.
	ManualTransferSwitchPowerEquipmentType PowerEquipmentType = "ManualTransferSwitch"
	// AutomaticTransferSwitchPowerEquipmentType is an automatic power transfer
	// switch.
	AutomaticTransferSwitchPowerEquipmentType PowerEquipmentType = "AutomaticTransferSwitch"
	// SwitchgearPowerEquipmentType Electrical switchgear.
	SwitchgearPowerEquipmentType PowerEquipmentType = "Switchgear"
	// PowerShelfPowerEquipmentType is a power shelf.
	PowerShelfPowerEquipmentType PowerEquipmentType = "PowerShelf"
	// BusPowerEquipmentType is an electrical bus.
	BusPowerEquipmentType PowerEquipmentType = "Bus"
	// BatteryShelfPowerEquipmentType is a battery shelf or battery-backed unit
	// (BBU).
	BatteryShelfPowerEquipmentType PowerEquipmentType = "BatteryShelf"
)

type TransferSensitivityType string

const (
	// HighTransferSensitivityType High sensitivity for initiating a transfer.
	HighTransferSensitivityType TransferSensitivityType = "High"
	// MediumTransferSensitivityType Medium sensitivity for initiating a transfer.
	MediumTransferSensitivityType TransferSensitivityType = "Medium"
	// LowTransferSensitivityType Low sensitivity for initiating a transfer.
	LowTransferSensitivityType TransferSensitivityType = "Low"
)

// PowerDistribution shall represent a power distribution component or unit for
// a Redfish implementation.
type PowerDistribution struct {
	common.Entity
	// AssetTag shall contain the user-assigned asset tag, which is an identifying
	// string that tracks the equipment for inventory purposes. Modifying this
	// property may modify the 'AssetTag' in the containing 'Chassis' resource.
	AssetTag string
	// Branches shall contain a link to a resource collection of type
	// 'CircuitCollection' that contains the branch circuits for this equipment.
	branches string
	// EquipmentType shall contain the type of equipment this resource represents.
	EquipmentType PowerEquipmentType
	// Feeders shall contain a link to a resource collection of type
	// 'CircuitCollection' that contains the feeder circuits for this equipment.
	feeders string
	// FirmwareVersion shall contain a string describing the firmware version of
	// this equipment as provided by the manufacturer.
	FirmwareVersion string
	// Location shall contain the location information of the associated equipment.
	Location common.Location
	// Mains shall contain a link to a resource collection of type
	// 'CircuitCollection' that contains the power input circuits for this
	// equipment.
	mains string
	// MainsRedundancy shall contain redundancy information for the mains (input)
	// circuits for this equipment. The values of the 'RedundancyGroup' array shall
	// reference resources of type 'Circuit'.
	//
	// Version added: v1.1.0
	MainsRedundancy common.RedundantGroup
	// Manufacturer shall contain the name of the organization responsible for
	// producing the equipment. This organization may be the entity from which the
	// equipment is purchased, but this is not necessarily true.
	Manufacturer string
	// Metrics shall contain a link to a resource of type
	// 'PowerDistributionMetrics'.
	metrics string
	// Model shall contain the manufacturer-provided model information of this
	// equipment.
	Model string
	// MultipartImportConfigurationPushURI shall contain a URI used to perform a
	// multipart HTTP or HTTPS 'POST' of a vendor-specific configuration file for
	// the purpose of importing the configuration contained within the file as
	// defined by the 'Import configuration data' clause of the Redfish
	// Specification. The value of this property should not contain a URI of a
	// Redfish resource. See the 'Redfish-defined URIs and relative reference
	// rules' clause in the Redfish Specification.
	//
	// Version added: v1.5.0
	MultipartImportConfigurationPushURI string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutletGroups shall contain a link to a resource collection of type
	// 'OutletCollection' that contains the outlet groups for this equipment.
	outletGroups string
	// Outlets shall contain a link to a resource collection of type
	// 'OutletCollection' that contains the outlets for this equipment.
	outlets string
	// PartNumber shall contain the manufacturer-provided part number for the
	// equipment.
	PartNumber string
	// PowerCapacityVA shall contain the maximum power capacity, rated as apparent
	// power, of this equipment, in volt-ampere units.
	//
	// Version added: v1.4.0
	PowerCapacityVA *uint `json:",omitempty"`
	// PowerSupplies shall contain a link to a resource collection of type
	// 'PowerSupplyCollection'.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of the 'PowerSupplies' link in
	// the 'Chassis' resource.
	powerSupplies string
	// PowerSupplyRedundancy shall contain redundancy information for the set of
	// power supplies for this equipment. The values of the 'RedundancyGroup' array
	// shall reference resources of type 'PowerSupply'.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of the 'PowerSupplyRedundancy'
	// property in the 'Chassis' resource.
	PowerSupplyRedundancy []common.RedundantGroup
	// ProductionDate shall contain the date of production or manufacture for this
	// equipment.
	ProductionDate string
	// Sensors shall be a link to a resource collection of type 'SensorCollection'
	// that contains the sensors located in the equipment and sub-components.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated in favor of the 'Sensors' link in the
	// 'Chassis' resource.
	sensors string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the equipment.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Subfeeds shall contain a link to a resource collection of type
	// 'CircuitCollection' that contains the subfeed circuits for this equipment.
	subfeeds string
	// TransferConfiguration shall contain the configuration information regarding
	// an automatic transfer switch function for this resource.
	TransferConfiguration TransferConfiguration
	// TransferCriteria shall contain the criteria for initiating a transfer within
	// an automatic transfer switch function for this resource.
	TransferCriteria TransferCriteria
	// UUID shall contain the UUID for the equipment.
	UUID string
	// UserLabel shall contain a user-assigned label used to identify this
	// resource. If a value has not been assigned by a user, the value of this
	// property shall be an empty string.
	//
	// Version added: v1.3.0
	UserLabel string
	// Version shall contain the hardware version of this equipment as determined
	// by the vendor or supplier.
	Version string
	// exportConfigurationTarget is the URL to send ExportConfiguration requests.
	exportConfigurationTarget string
	// transferControlTarget is the URL to send TransferControl requests.
	transferControlTarget string
	// chassis are the URIs for Chassis.
	chassis []string
	// facility is the URI for Facility.
	facility string
	// managedBy are the URIs for ManagedBy.
	managedBy []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PowerDistribution object from the raw JSON.
func (p *PowerDistribution) UnmarshalJSON(b []byte) error {
	type temp PowerDistribution
	type pActions struct {
		ExportConfiguration common.ActionTarget `json:"#PowerDistribution.ExportConfiguration"`
		TransferControl     common.ActionTarget `json:"#PowerDistribution.TransferControl"`
	}
	type pLinks struct {
		Chassis   common.Links `json:"Chassis"`
		Facility  common.Link  `json:"Facility"`
		ManagedBy common.Links `json:"ManagedBy"`
	}
	var tmp struct {
		temp
		Actions       pActions
		Links         pLinks
		Branches      common.Link `json:"branches"`
		Feeders       common.Link `json:"feeders"`
		Mains         common.Link `json:"mains"`
		Metrics       common.Link `json:"metrics"`
		OutletGroups  common.Link `json:"outletGroups"`
		Outlets       common.Link `json:"outlets"`
		PowerSupplies common.Link `json:"powerSupplies"`
		Sensors       common.Link `json:"sensors"`
		Subfeeds      common.Link `json:"subfeeds"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PowerDistribution(tmp.temp)

	// Extract the links to other entities for later
	p.exportConfigurationTarget = tmp.Actions.ExportConfiguration.Target
	p.transferControlTarget = tmp.Actions.TransferControl.Target
	p.chassis = tmp.Links.Chassis.ToStrings()
	p.facility = tmp.Links.Facility.String()
	p.managedBy = tmp.Links.ManagedBy.ToStrings()
	p.branches = tmp.Branches.String()
	p.feeders = tmp.Feeders.String()
	p.mains = tmp.Mains.String()
	p.metrics = tmp.Metrics.String()
	p.outletGroups = tmp.OutletGroups.String()
	p.outlets = tmp.Outlets.String()
	p.powerSupplies = tmp.PowerSupplies.String()
	p.sensors = tmp.Sensors.String()
	p.subfeeds = tmp.Subfeeds.String()

	// This is a read/write object, so we need to save the raw object data for later
	p.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *PowerDistribution) Update() error {
	readWriteFields := []string{
		"AssetTag",
		"Location",
		"MainsRedundancy",
		"PowerSupplyRedundancy",
		"Status",
		"TransferConfiguration",
		"TransferCriteria",
		"UserLabel",
	}

	return p.UpdateFromRawData(p, p.rawData, readWriteFields)
}

// GetPowerDistribution will get a PowerDistribution instance from the service.
func GetPowerDistribution(c common.Client, uri string) (*PowerDistribution, error) {
	return common.GetObject[PowerDistribution](c, uri)
}

// ListReferencedPowerDistributions gets the collection of PowerDistribution from
// a provided reference.
func ListReferencedPowerDistributions(c common.Client, link string) ([]*PowerDistribution, error) {
	return common.GetCollectionObjects[PowerDistribution](c, link)
}

// ExportConfiguration shall export the specified configuration of the equipment in a
// vendor-specific format. Upon successful completion of the action and any
// asynchronous processing, the 'Location' header in the response shall contain
// a URI to a file that contains the configuration data.
// components - This parameter shall contain an array of components of the
// equipment for which to export configuration data.
// encryptionPassphrase - This parameter shall contain the encryption
// passphrase for the exported file. If this parameter is specified and has a
// non-zero length, the service shall encrypt the exported file with the
// passphrase. Otherwise, the service shall not encrypt the exported file.
// exportType - This parameter shall contain the type of export to perform.
// oEMComponents - This parameter shall contain an array of OEM-specific
// components of the equipment for which to export configuration data.
// security - This parameter shall contain the policy to apply when exporting
// secure information.
func (p *PowerDistribution) ExportConfiguration(components []string, encryptionPassphrase string, exportType ExportType, oEMComponents string, security ExportSecurity) error {
	payload := make(map[string]any)
	payload["Components"] = components
	payload["EncryptionPassphrase"] = encryptionPassphrase
	payload["ExportType"] = exportType
	payload["OEMComponents"] = oEMComponents
	payload["Security"] = security
	return p.Post(p.exportConfigurationTarget, payload)
}

// TransferControl shall transfer power input from the existing mains circuit to
// the alternative mains circuit.
func (p *PowerDistribution) TransferControl() error {
	payload := make(map[string]any)
	return p.Post(p.transferControlTarget, payload)
}

// Chassis gets the Chassis linked resources.
func (p *PowerDistribution) Chassis(client common.Client) ([]*Chassis, error) {
	return common.GetObjects[Chassis](client, p.chassis)
}

// Facility gets the Facility linked resource.
func (p *PowerDistribution) Facility(client common.Client) (*Facility, error) {
	if p.facility == "" {
		return nil, nil
	}
	return common.GetObject[Facility](client, p.facility)
}

// ManagedBy gets the ManagedBy linked resources.
func (p *PowerDistribution) ManagedBy(client common.Client) ([]*Manager, error) {
	return common.GetObjects[Manager](client, p.managedBy)
}

// Branches gets the Branches collection.
func (p *PowerDistribution) Branches(client common.Client) ([]*Circuit, error) {
	if p.branches == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Circuit](client, p.branches)
}

// Feeders gets the Feeders collection.
func (p *PowerDistribution) Feeders(client common.Client) ([]*Circuit, error) {
	if p.feeders == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Circuit](client, p.feeders)
}

// Mains gets the Mains collection.
func (p *PowerDistribution) Mains(client common.Client) ([]*Circuit, error) {
	if p.mains == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Circuit](client, p.mains)
}

// Metrics gets the Metrics linked resource.
func (p *PowerDistribution) Metrics(client common.Client) (*PowerDistributionMetrics, error) {
	if p.metrics == "" {
		return nil, nil
	}
	return common.GetObject[PowerDistributionMetrics](client, p.metrics)
}

// OutletGroups gets the OutletGroups collection.
func (p *PowerDistribution) OutletGroups(client common.Client) ([]*OutletGroup, error) {
	if p.outletGroups == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[OutletGroup](client, p.outletGroups)
}

// Outlets gets the Outlets collection.
func (p *PowerDistribution) Outlets(client common.Client) ([]*Outlet, error) {
	if p.outlets == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Outlet](client, p.outlets)
}

// PowerSupplies gets the PowerSupplies collection.
func (p *PowerDistribution) PowerSupplies(client common.Client) ([]*PowerSupply, error) {
	if p.powerSupplies == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[PowerSupply](client, p.powerSupplies)
}

// Sensors gets the Sensors collection.
func (p *PowerDistribution) Sensors(client common.Client) ([]*Sensor, error) {
	if p.sensors == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Sensor](client, p.sensors)
}

// Subfeeds gets the Subfeeds collection.
func (p *PowerDistribution) Subfeeds(client common.Client) ([]*Circuit, error) {
	if p.subfeeds == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Circuit](client, p.subfeeds)
}

// TransferConfiguration shall contain the configuration information regarding
// an automatic transfer switch function for this resource.
type TransferConfiguration struct {
	// ActiveMainsId shall contain the mains circuit that is switched on and
	// qualified to supply power to the output circuit. The value shall be a string
	// that matches the 'Id' property value of a circuit contained in the
	// collection referenced by the 'Mains' property.
	ActiveMainsID string `json:"ActiveMainsId"`
	// AutoTransferEnabled shall indicate if the qualified alternate mains circuit
	// is automatically switched on when the preferred mains circuit becomes
	// unqualified and is automatically switched off.
	AutoTransferEnabled bool
	// ClosedTransitionAllowed shall indicate if a make-before-break switching
	// sequence of the mains circuits is permitted when they are both qualified and
	// in synchronization.
	ClosedTransitionAllowed bool
	// ClosedTransitionTimeoutSeconds shall contain the time in seconds to wait for
	// a closed transition to occur.
	ClosedTransitionTimeoutSeconds *int `json:",omitempty"`
	// PreferredMainsId shall contain the preferred source for mains circuit to
	// this equipment. The value shall be a string that matches the 'Id' property
	// value of a circuit contained in the collection referenced by the 'Mains'
	// property.
	PreferredMainsID string `json:"PreferredMainsId"`
	// RetransferDelaySeconds shall contain the time in seconds to delay the
	// automatic transfer from the alternate mains circuit back to the preferred
	// mains circuit.
	RetransferDelaySeconds *int `json:",omitempty"`
	// RetransferEnabled shall indicate if the automatic transfer is permitted from
	// the alternate mains circuit back to the preferred mains circuit after the
	// preferred mains circuit is qualified again and the 'RetransferDelaySeconds'
	// time has expired.
	RetransferEnabled bool
	// TransferDelaySeconds shall contain the time in seconds to delay the
	// automatic transfer from the preferred mains circuit to the alternate mains
	// circuit when the preferred mains circuit is disqualified. A value of zero
	// shall mean it transfers as fast as possible.
	TransferDelaySeconds *int `json:",omitempty"`
	// TransferInhibit shall indicate if any transfer is inhibited.
	TransferInhibit bool
}

// TransferCriteria shall contain the criteria for initiating a transfer within
// an automatic transfer switch function for this resource.
type TransferCriteria struct {
	// OverNominalFrequencyHz shall contain the frequency in hertz units over the
	// nominal value that satisfies a criterion for transfer.
	OverNominalFrequencyHz *float64 `json:",omitempty"`
	// OverVoltageRMSPercentage shall contain the positive percentage of voltage
	// RMS over the nominal value that satisfies a criterion for transfer.
	OverVoltageRMSPercentage *float64 `json:",omitempty"`
	// TransferSensitivity shall contain the setting that adjusts the analytical
	// sensitivity of the detection of the quality of voltage waveform that
	// satisfies a criterion for transfer.
	TransferSensitivity TransferSensitivityType
	// UnderNominalFrequencyHz shall contain the frequency in hertz units under the
	// nominal value that satisfies a criterion for transfer.
	UnderNominalFrequencyHz *float64 `json:",omitempty"`
	// UnderVoltageRMSPercentage shall contain the negative percentage of voltage
	// RMS under the nominal value that satisfies a criterion for transfer.
	UnderVoltageRMSPercentage *float64 `json:",omitempty"`
}
