//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #CoolingUnit.v1_5_0.CoolingUnit

package schemas

import (
	"encoding/json"
)

type CoolingEquipmentType string

const (
	// CDUCoolingEquipmentType is a coolant distribution unit (CDU).
	CDUCoolingEquipmentType CoolingEquipmentType = "CDU"
	// HeatExchangerCoolingEquipmentType is a heat exchanger.
	HeatExchangerCoolingEquipmentType CoolingEquipmentType = "HeatExchanger"
	// ImmersionUnitCoolingEquipmentType is an immersion cooling unit.
	ImmersionUnitCoolingEquipmentType CoolingEquipmentType = "ImmersionUnit"
	// RPUCoolingEquipmentType is a reservoir and pumping unit (RPU).
	RPUCoolingEquipmentType CoolingEquipmentType = "RPU"
)

type CoolingUnitMode string

const (
	// EnabledCoolingUnitMode shall indicate a request to enable the cooling unit.
	// Upon successful completion, the 'State' property within 'Status', shall
	// contain the value 'Enabled'.
	EnabledCoolingUnitMode CoolingUnitMode = "Enabled"
	// DisabledCoolingUnitMode shall indicate a request to disable the cooling
	// unit. When disabled, primary functions of the cooling unit, such as pump
	// activity, are also disabled. When disabled, the cooling unit may perform
	// administrative functions, such as monitoring sensors, controlling valves,
	// and accepting new firmware. Upon successful completion, the 'State' property
	// within 'Status', shall contain the value 'Disabled'.
	DisabledCoolingUnitMode CoolingUnitMode = "Disabled"
)

// CoolingUnit shall represent a cooling system component or unit for a Redfish
// implementation.
type CoolingUnit struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	assembly string
	// AssetTag shall contain the user-assigned asset tag, which is an identifying
	// string that tracks the equipment for inventory purposes.
	AssetTag string
	// Coolant shall contain details regarding the coolant contained or used by
	// this unit.
	Coolant Coolant
	// CoolantConnectorRedundancy shall contain redundancy information for the set
	// of coolant connectors attached to this equipment. The values of the
	// 'RedundancyGroup' array shall reference resources of type
	// 'CoolantConnector'.
	//
	// Version added: v1.1.0
	CoolantConnectorRedundancy []RedundantGroup
	// CoolingCapacityWatts shall contain the manufacturer-provided cooling
	// capacity, in watt units, of this equipment.
	CoolingCapacityWatts *int `json:",omitempty"`
	// CoolingUnitRedundancy shall contain redundancy information for the groups of
	// cooling units.
	//
	// Version added: v1.5.0
	CoolingUnitRedundancy []RedundantGroup
	// EnvironmentMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that specifies the environment metrics for this
	// equipment.
	environmentMetrics string
	// EquipmentType shall contain the type of equipment this resource represents.
	EquipmentType CoolingEquipmentType
	// FilterRedundancy shall contain redundancy information for the groups of
	// filters in this unit.
	FilterRedundancy []RedundantGroup
	// Filters shall contain a link to a resource collection of type
	// 'FilterCollection' that contains the filter information for this equipment.
	filters string
	// FirmwareVersion shall contain a string describing the firmware version of
	// this equipment as provided by the manufacturer.
	FirmwareVersion string
	// LeakDetection shall contain a link to a resource of type 'LeakDetection'
	// that contains the leak detection component information for this equipment.
	// This link should be used when the leak detection capabilities are tied to a
	// particular cooling unit or system which may span multiple 'Chassis'
	// resources. For equipment represented with a single 'Chassis' resource or
	// detection inside a particular 'Chassis' resource, populating the
	// 'LeakDetection' resource under 'ThermalSubsystem' for the relevant 'Chassis'
	// is the preferred approach.
	leakDetection string
	// Location shall contain the location information of the associated equipment.
	Location Location
	// Manufacturer shall contain the name of the organization responsible for
	// producing the equipment. This organization may be the entity from which the
	// equipment is purchased, but this is not necessarily true.
	Manufacturer string
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
	// Version added: v1.4.0
	MultipartImportConfigurationPushURI string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the manufacturer-provided part number for the
	// equipment.
	PartNumber string
	// PrimaryCoolantConnectors shall contain a link to a resource collection of
	// type 'CoolantConnectorCollection' that contains the primary coolant
	// connectors for this equipment.
	primaryCoolantConnectors string
	// ProductionDate shall contain the date of production or manufacture for this
	// equipment.
	ProductionDate string
	// PumpRedundancy shall contain redundancy information for the groups of pumps
	// in this unit.
	PumpRedundancy []RedundantGroup
	// Pumps shall contain a link to a resource collection of type 'PumpCollection'
	// that contains the pump information for this equipment.
	pumps string
	// RatedThermalLossToAirWatts shall contain the rated maximum amount of heat,
	// in watt units, lost to the surrounding environment during normal operation.
	//
	// Version added: v1.4.0
	RatedThermalLossToAirWatts *int `json:",omitempty"`
	// Reservoirs shall contain a link to a resource collection of type
	// 'ReservoirCollection' that contains the reservoir information for this
	// equipment.
	reservoirs string
	// SecondaryCoolantConnectors shall contain a link to a resource collection of
	// type 'CoolantConnectorCollection' that contains the secondary coolant
	// connectors for this equipment.
	secondaryCoolantConnectors string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the equipment.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UserLabel shall contain a user-assigned label used to identify this
	// resource. If a value has not been assigned by a user, the value of this
	// property shall be an empty string.
	UserLabel string
	// Version shall contain the hardware version of this equipment as determined
	// by the vendor or supplier.
	Version string
	// exportConfigurationTarget is the URL to send ExportConfiguration requests.
	exportConfigurationTarget string
	// setModeTarget is the URL to send SetMode requests.
	setModeTarget string
	// chassis are the URIs for Chassis.
	chassis []string
	// facility is the URI for Facility.
	facility string
	// managedBy are the URIs for ManagedBy.
	managedBy []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a CoolingUnit object from the raw JSON.
func (c *CoolingUnit) UnmarshalJSON(b []byte) error {
	type temp CoolingUnit
	type cActions struct {
		ExportConfiguration ActionTarget `json:"#CoolingUnit.ExportConfiguration"`
		SetMode             ActionTarget `json:"#CoolingUnit.SetMode"`
	}
	type cLinks struct {
		Chassis   Links `json:"Chassis"`
		Facility  Link  `json:"Facility"`
		ManagedBy Links `json:"ManagedBy"`
	}
	var tmp struct {
		temp
		Actions                    cActions
		Links                      cLinks
		Assembly                   Link `json:"Assembly"`
		EnvironmentMetrics         Link `json:"EnvironmentMetrics"`
		Filters                    Link `json:"Filters"`
		LeakDetection              Link `json:"LeakDetection"`
		PrimaryCoolantConnectors   Link `json:"PrimaryCoolantConnectors"`
		Pumps                      Link `json:"Pumps"`
		Reservoirs                 Link `json:"Reservoirs"`
		SecondaryCoolantConnectors Link `json:"SecondaryCoolantConnectors"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CoolingUnit(tmp.temp)

	// Extract the links to other entities for later
	c.exportConfigurationTarget = tmp.Actions.ExportConfiguration.Target
	c.setModeTarget = tmp.Actions.SetMode.Target
	c.chassis = tmp.Links.Chassis.ToStrings()
	c.facility = tmp.Links.Facility.String()
	c.managedBy = tmp.Links.ManagedBy.ToStrings()
	c.assembly = tmp.Assembly.String()
	c.environmentMetrics = tmp.EnvironmentMetrics.String()
	c.filters = tmp.Filters.String()
	c.leakDetection = tmp.LeakDetection.String()
	c.primaryCoolantConnectors = tmp.PrimaryCoolantConnectors.String()
	c.pumps = tmp.Pumps.String()
	c.reservoirs = tmp.Reservoirs.String()
	c.secondaryCoolantConnectors = tmp.SecondaryCoolantConnectors.String()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *CoolingUnit) Update() error {
	readWriteFields := []string{
		"AssetTag",
		"UserLabel",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetCoolingUnit will get a CoolingUnit instance from the service.
func GetCoolingUnit(c Client, uri string) (*CoolingUnit, error) {
	return GetObject[CoolingUnit](c, uri)
}

// ListReferencedCoolingUnits gets the collection of CoolingUnit from
// a provided reference.
func ListReferencedCoolingUnits(c Client, link string) ([]*CoolingUnit, error) {
	return GetCollectionObjects[CoolingUnit](c, link)
}

// CoolingUnitExportConfigurationParameters holds the parameters for the ExportConfiguration action.
type CoolingUnitExportConfigurationParameters struct {
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
func (c *CoolingUnit) ExportConfiguration(params *CoolingUnitExportConfigurationParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(c.client,
		c.exportConfigurationTarget, params, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall set the operating mode of the cooling unit.
// mode - This parameter shall contain the desired operating mode of the
// cooling unit.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *CoolingUnit) SetMode(mode CoolingUnitMode) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Mode"] = mode
	resp, taskInfo, err := PostWithTask(c.client,
		c.setModeTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Chassis gets the Chassis linked resources.
func (c *CoolingUnit) Chassis() ([]*Chassis, error) {
	return GetObjects[Chassis](c.client, c.chassis)
}

// Facility gets the Facility linked resource.
func (c *CoolingUnit) Facility() (*Facility, error) {
	if c.facility == "" {
		return nil, nil
	}
	return GetObject[Facility](c.client, c.facility)
}

// ManagedBy gets the ManagedBy linked resources.
func (c *CoolingUnit) ManagedBy() ([]*Manager, error) {
	return GetObjects[Manager](c.client, c.managedBy)
}

// Assembly gets the Assembly linked resource.
func (c *CoolingUnit) Assembly() (*Assembly, error) {
	if c.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](c.client, c.assembly)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (c *CoolingUnit) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if c.environmentMetrics == "" {
		return nil, nil
	}
	return GetObject[EnvironmentMetrics](c.client, c.environmentMetrics)
}

// Filters gets the Filters collection.
func (c *CoolingUnit) Filters() ([]*Filter, error) {
	if c.filters == "" {
		return nil, nil
	}
	return GetCollectionObjects[Filter](c.client, c.filters)
}

// LeakDetection gets the LeakDetection linked resource.
func (c *CoolingUnit) LeakDetection() (*LeakDetection, error) {
	if c.leakDetection == "" {
		return nil, nil
	}
	return GetObject[LeakDetection](c.client, c.leakDetection)
}

// PrimaryCoolantConnectors gets the PrimaryCoolantConnectors collection.
func (c *CoolingUnit) PrimaryCoolantConnectors() ([]*CoolantConnector, error) {
	if c.primaryCoolantConnectors == "" {
		return nil, nil
	}
	return GetCollectionObjects[CoolantConnector](c.client, c.primaryCoolantConnectors)
}

// Pumps gets the Pumps collection.
func (c *CoolingUnit) Pumps() ([]*Pump, error) {
	if c.pumps == "" {
		return nil, nil
	}
	return GetCollectionObjects[Pump](c.client, c.pumps)
}

// Reservoirs gets the Reservoirs collection.
func (c *CoolingUnit) Reservoirs() ([]*Reservoir, error) {
	if c.reservoirs == "" {
		return nil, nil
	}
	return GetCollectionObjects[Reservoir](c.client, c.reservoirs)
}

// SecondaryCoolantConnectors gets the SecondaryCoolantConnectors collection.
func (c *CoolingUnit) SecondaryCoolantConnectors() ([]*CoolantConnector, error) {
	if c.secondaryCoolantConnectors == "" {
		return nil, nil
	}
	return GetCollectionObjects[CoolantConnector](c.client, c.secondaryCoolantConnectors)
}
