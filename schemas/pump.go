//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2024.4 - #Pump.v1_2_0.Pump

package schemas

import (
	"encoding/json"
)

type PumpMode string

const (
	// EnabledPumpMode shall indicate a request to enable the pump. Upon successful
	// completion, the 'State' property within 'Status', shall contain the value
	// 'Enabled'.
	EnabledPumpMode PumpMode = "Enabled"
	// DisabledPumpMode shall indicate a request to disable the pump. Upon
	// successful completion, the 'State' property within 'Status', shall contain
	// the value 'Disabled'.
	DisabledPumpMode PumpMode = "Disabled"
)

type PumpType string

const (
	// LiquidPumpType is a water or liquid pump.
	LiquidPumpType PumpType = "Liquid"
	// CompressorPumpType is a compressor.
	CompressorPumpType PumpType = "Compressor"
)

// Pump shall represent the management properties for monitoring and management
// of pumps for a Redfish implementation.
type Pump struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	assembly string
	// AssetTag shall contain the user-assigned asset tag, which is an identifying
	// string that tracks the equipment for inventory purposes.
	AssetTag string
	// Filters shall contain a link to a resource collection of type
	// 'FilterCollection' that contains a set of filters.
	filters string
	// FirmwareVersion shall contain a string describing the firmware version of
	// this equipment as provided by the manufacturer.
	FirmwareVersion string
	// InletPressurekPa shall contain the pressure, in kilopascal units, for the
	// inlet to this pump. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'PressurekPa'.
	//
	// Version added: v1.2.0
	InletPressurekPa SensorExcerpt
	// Location shall contain the location information of this pump.
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for
	// producing the pump. This organization may be the entity from whom the pump
	// is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the model information as defined by the manufacturer for
	// this pump.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for
	// this pump.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region
	// within the chassis with which this pump is associated.
	PhysicalContext PhysicalContext
	// ProductionDate shall contain the date of production or manufacture for this
	// equipment.
	ProductionDate string
	// PumpSpeedPercent shall contain the current speed, in percent units, of this
	// pump. The value of the 'DataSourceUri' property, if present, shall reference
	// a resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Percent'.
	PumpSpeedPercent SensorPumpExcerpt
	// PumpType shall contain the type of pump represented by this resource.
	PumpType PumpType
	// SerialNumber shall contain the serial number as defined by the manufacturer
	// for this pump.
	SerialNumber string
	// ServiceHours shall contain the number of hours of service that the pump has
	// been in operation.
	ServiceHours *float64 `json:",omitempty"`
	// SparePartNumber shall contain the spare or replacement part number as
	// defined by the manufacturer for this pump.
	SparePartNumber string
	// SpeedControlPercent shall contain the desired speed, in percent units, of
	// this pump. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Control' with the 'ControlType' property
	// containing the value 'Percent'.
	//
	// Version added: v1.1.0
	SpeedControlPercent ControlSingleLoopExcerpt
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UserLabel shall contain a user-assigned label used to identify this
	// resource. If a value has not been assigned by a user, the value of this
	// property shall be an empty string.
	UserLabel string
	// Version shall contain the hardware version of this equipment as determined
	// by the vendor or supplier.
	Version string
	// setModeTarget is the URL to send SetMode requests.
	setModeTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Pump object from the raw JSON.
func (p *Pump) UnmarshalJSON(b []byte) error {
	type temp Pump
	type pActions struct {
		SetMode ActionTarget `json:"#Pump.SetMode"`
	}
	var tmp struct {
		temp
		Actions  pActions
		Assembly Link `json:"Assembly"`
		Filters  Link `json:"Filters"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = Pump(tmp.temp)

	// Extract the links to other entities for later
	p.setModeTarget = tmp.Actions.SetMode.Target
	p.assembly = tmp.Assembly.String()
	p.filters = tmp.Filters.String()

	// This is a read/write object, so we need to save the raw object data for later
	p.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *Pump) Update() error {
	readWriteFields := []string{
		"AssetTag",
		"LocationIndicatorActive",
		"ServiceHours",
		"UserLabel",
	}

	return p.UpdateFromRawData(p, p.RawData, readWriteFields)
}

// GetPump will get a Pump instance from the service.
func GetPump(c Client, uri string) (*Pump, error) {
	return GetObject[Pump](c, uri)
}

// ListReferencedPumps gets the collection of Pump from
// a provided reference.
func ListReferencedPumps(c Client, link string) ([]*Pump, error) {
	return GetCollectionObjects[Pump](c, link)
}

// This action shall set the operating mode of the pump.
// mode - This parameter shall contain the desired operating mode of the pump.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *Pump) SetMode(mode PumpMode) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Mode"] = mode
	resp, taskInfo, err := PostWithTask(p.client,
		p.setModeTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Assembly gets the Assembly linked resource.
func (p *Pump) Assembly() (*Assembly, error) {
	if p.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](p.client, p.assembly)
}

// Filters gets the Filters collection.
func (p *Pump) Filters() ([]*Filter, error) {
	if p.filters == "" {
		return nil, nil
	}
	return GetCollectionObjects[Filter](p.client, p.filters)
}
