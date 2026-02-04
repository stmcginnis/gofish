//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/PowerDistribution.v1_6_0.json
// 2025.4 - #PowerDistribution.v1_6_0.PowerDistribution

package schemas

import (
	"encoding/json"
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
	Entity
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
	Location Location
	// Mains shall contain a link to a resource collection of type
	// 'CircuitCollection' that contains the power input circuits for this
	// equipment.
	mains string
	// MainsRedundancy shall contain redundancy information for the mains (input)
	// circuits for this equipment. The values of the 'RedundancyGroup' array shall
	// reference resources of type 'Circuit'.
	//
	// Version added: v1.1.0
	MainsRedundancy RedundantGroup
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
	// OEM shall contain the OEM extensions. All values for properties that this
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
	// PowerDistributionRedundancy shall contain redundancy information for the
	// groups of power distribution units.
	//
	// Version added: v1.6.0
	PowerDistributionRedundancy []RedundantGroup
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
	PowerSupplyRedundancy []RedundantGroup
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
	Status Status
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
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a PowerDistribution object from the raw JSON.
func (p *PowerDistribution) UnmarshalJSON(b []byte) error {
	type temp PowerDistribution
	type pActions struct {
		ExportConfiguration ActionTarget `json:"#PowerDistribution.ExportConfiguration"`
		TransferControl     ActionTarget `json:"#PowerDistribution.TransferControl"`
	}
	type pLinks struct {
		Chassis   Links `json:"Chassis"`
		Facility  Link  `json:"Facility"`
		ManagedBy Links `json:"ManagedBy"`
	}
	var tmp struct {
		temp
		Actions       pActions
		Links         pLinks
		Branches      Link `json:"Branches"`
		Feeders       Link `json:"Feeders"`
		Mains         Link `json:"Mains"`
		Metrics       Link `json:"Metrics"`
		OutletGroups  Link `json:"OutletGroups"`
		Outlets       Link `json:"Outlets"`
		PowerSupplies Link `json:"PowerSupplies"`
		Sensors       Link `json:"Sensors"`
		Subfeeds      Link `json:"Subfeeds"`
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
	p.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *PowerDistribution) Update() error {
	readWriteFields := []string{
		"AssetTag",
		"UserLabel",
	}

	return p.UpdateFromRawData(p, p.RawData, readWriteFields)
}

// GetPowerDistribution will get a PowerDistribution instance from the service.
func GetPowerDistribution(c Client, uri string) (*PowerDistribution, error) {
	return GetObject[PowerDistribution](c, uri)
}

// ListReferencedPowerDistributions gets the collection of PowerDistribution from
// a provided reference.
func ListReferencedPowerDistributions(c Client, link string) ([]*PowerDistribution, error) {
	return GetCollectionObjects[PowerDistribution](c, link)
}

// PowerDistributionExportConfigurationParameters holds the parameters for the ExportConfiguration action.
type PowerDistributionExportConfigurationParameters struct {
	// Components shall contain an array of components of the equipment for which
	// to export configuration data.
	Components []Component `json:"Components,omitempty"`
	// EncryptionPassphrase shall contain the encryption passphrase for the
	// exported file. If this parameter is specified and has a non-zero length, the
	// service shall encrypt the exported file with the passphrase. Otherwise, the
	// service shall not encrypt the exported file.
	EncryptionPassphrase string `json:"EncryptionPassphrase,omitempty"`
	// ExportType shall contain the type of export to perform.
	ExportType ExportType `json:"ExportType,omitempty"`
	// OEMComponents shall contain an array of OEM-specific components of the
	// equipment for which to export configuration data.
	OEMComponents []string `json:"OEMComponents,omitempty"`
	// Security shall contain the policy to apply when exporting secure
	// information.
	Security ExportSecurity `json:"Security,omitempty"`
}

// This action shall export the specified configuration of the equipment in a
// vendor-specific format. Upon successful completion of the action and any
// asynchronous processing, the 'Location' header in the response shall contain
// a URI to a file that contains the configuration data.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *PowerDistribution) ExportConfiguration(params *PowerDistributionExportConfigurationParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(p.client,
		p.exportConfigurationTarget, params, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall transfer power input from the existing mains circuit to
// the alternative mains circuit.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *PowerDistribution) TransferControl() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(p.client,
		p.transferControlTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Chassis gets the Chassis linked resources.
func (p *PowerDistribution) Chassis() ([]*Chassis, error) {
	return GetObjects[Chassis](p.client, p.chassis)
}

// Facility gets the Facility linked resource.
func (p *PowerDistribution) Facility() (*Facility, error) {
	if p.facility == "" {
		return nil, nil
	}
	return GetObject[Facility](p.client, p.facility)
}

// ManagedBy gets the ManagedBy linked resources.
func (p *PowerDistribution) ManagedBy() ([]*Manager, error) {
	return GetObjects[Manager](p.client, p.managedBy)
}

// Branches gets the Branches collection.
func (p *PowerDistribution) Branches() ([]*Circuit, error) {
	if p.branches == "" {
		return nil, nil
	}
	return GetCollectionObjects[Circuit](p.client, p.branches)
}

// Feeders gets the Feeders collection.
func (p *PowerDistribution) Feeders() ([]*Circuit, error) {
	if p.feeders == "" {
		return nil, nil
	}
	return GetCollectionObjects[Circuit](p.client, p.feeders)
}

// Mains gets the Mains collection.
func (p *PowerDistribution) Mains() ([]*Circuit, error) {
	if p.mains == "" {
		return nil, nil
	}
	return GetCollectionObjects[Circuit](p.client, p.mains)
}

// Metrics gets the Metrics linked resource.
func (p *PowerDistribution) Metrics() (*PowerDistributionMetrics, error) {
	if p.metrics == "" {
		return nil, nil
	}
	return GetObject[PowerDistributionMetrics](p.client, p.metrics)
}

// OutletGroups gets the OutletGroups collection.
func (p *PowerDistribution) OutletGroups() ([]*OutletGroup, error) {
	if p.outletGroups == "" {
		return nil, nil
	}
	return GetCollectionObjects[OutletGroup](p.client, p.outletGroups)
}

// Outlets gets the Outlets collection.
func (p *PowerDistribution) Outlets() ([]*Outlet, error) {
	if p.outlets == "" {
		return nil, nil
	}
	return GetCollectionObjects[Outlet](p.client, p.outlets)
}

// PowerSupplies gets the PowerSupplies collection.
func (p *PowerDistribution) PowerSupplies() ([]*PowerSupply, error) {
	if p.powerSupplies == "" {
		return nil, nil
	}
	return GetCollectionObjects[PowerSupply](p.client, p.powerSupplies)
}

// Sensors gets the Sensors collection.
func (p *PowerDistribution) Sensors() ([]*Sensor, error) {
	if p.sensors == "" {
		return nil, nil
	}
	return GetCollectionObjects[Sensor](p.client, p.sensors)
}

// Subfeeds gets the Subfeeds collection.
func (p *PowerDistribution) Subfeeds() ([]*Circuit, error) {
	if p.subfeeds == "" {
		return nil, nil
	}
	return GetCollectionObjects[Circuit](p.client, p.subfeeds)
}

// TransferConfiguration shall contain the configuration information regarding
// an automatic transfer switch function for this resource.
type TransferConfiguration struct {
	// ActiveMainsID shall contain the mains circuit that is switched on and
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
	// PreferredMainsID shall contain the preferred source for mains circuit to
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
