//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/PowerSupply.v1_6_0.json
// 2024.1 - #PowerSupply.v1_6_0.PowerSupply

package schemas

import (
	"encoding/json"
)

type LineStatus string

const (
	// NormalLineStatus Line input is within normal operating range.
	NormalLineStatus LineStatus = "Normal"
	// LossOfInputLineStatus No power detected at line input.
	LossOfInputLineStatus LineStatus = "LossOfInput"
	// OutOfRangeLineStatus Line input voltage or current is outside of normal
	// operating range.
	OutOfRangeLineStatus LineStatus = "OutOfRange"
)

type PowerSupplyUnitType string

const (
	// ACPowerSupplyUnitType Alternating Current (AC) power supply.
	ACPowerSupplyUnitType PowerSupplyUnitType = "AC"
	// DCPowerSupplyUnitType Direct Current (DC) power supply.
	DCPowerSupplyUnitType PowerSupplyUnitType = "DC"
	// ACorDCPowerSupplyUnitType The power supply supports both DC and AC.
	ACorDCPowerSupplyUnitType PowerSupplyUnitType = "ACorDC"
	// DCRegulatorPowerSupplyUnitType Direct Current (DC) voltage regulator.
	DCRegulatorPowerSupplyUnitType PowerSupplyUnitType = "DCRegulator"
)

// PowerSupplyUnit shall represent a power supply unit for a Redfish implementation.
// It may also represent a location, such as a slot, socket, or bay, where a
// unit may be installed, but the 'State' property within the 'Status' property
// contains 'Absent'.
type PowerSupplyUnit struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	assembly string
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains certificates for device identity and
	// attestation.
	//
	// Version added: v1.6.0
	certificates string
	// EfficiencyRatings shall contain an array of efficiency ratings for this
	// power supply.
	EfficiencyRatings []EfficiencyRating
	// ElectricalSourceManagerURIs shall contain an array of URIs to the management
	// applications or devices that provide monitoring or control of the upstream
	// electrical sources that provide power to this power supply.
	//
	// Version added: v1.2.0
	ElectricalSourceManagerURIs []string
	// ElectricalSourceNames shall contain an array of strings that identify the
	// upstream electrical sources, such as the names of circuits or outlets, that
	// provide power to this power supply.
	//
	// Version added: v1.2.0
	ElectricalSourceNames []string
	// FirmwareVersion shall contain the firmware version as defined by the
	// manufacturer for this power supply.
	FirmwareVersion string
	// HotPluggable shall indicate whether the device can be inserted or removed
	// while the underlying equipment otherwise remains in its current operational
	// state. Devices indicated as hot-pluggable shall allow the device to become
	// operable without altering the operational state of the underlying equipment.
	// Devices that cannot be inserted or removed from equipment in operation, or
	// devices that cannot become operable without affecting the operational state
	// of that equipment, shall be indicated as not hot-pluggable.
	HotPluggable bool
	// InputNominalVoltageType shall contain the nominal voltage type that is
	// detected on the line input to this power supply. This value shall be one of
	// the values shown in the 'NominalVoltageType' property in the 'PowerSupplyUnitInputRanges'
	// array, if present. If the line input voltage is unknown, out of range, or
	// there is no input provided to the power supply, the value shall be 'null'.
	InputNominalVoltageType NominalVoltageType
	// InputRanges shall contain a collection of ranges usable by this power
	// supply.
	InputRanges []PowerSupplyUnitInputRange
	// LineInputStatus shall contain the status of the power line input for this
	// power supply.
	//
	// Version added: v1.3.0
	LineInputStatus LineStatus
	// Location shall contain the location information of the power supply. For a
	// resource in the 'Absent' state, this property describes the empty location,
	// such as a slot, socket, or bay, to represent the available capacity.
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for
	// producing the power supply. This organization may be the entity from whom
	// the power supply is purchased, but this is not necessarily true.
	Manufacturer string
	// Metrics shall contain a link to a resource of type 'PowerSupplyMetrics'.
	metrics string
	// Model shall contain the model information as defined by the manufacturer for
	// this power supply.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OutputNominalVoltageType shall contain the nominal voltage type of the
	// single output line of this power supply. This property is intended to
	// describe power supply types that connect to additional power infrastructure
	// components, such as a rectifier component in a modular power system. This
	// property shall not be present for power supplies not intended to connect to
	// additional power infrastructure components.
	//
	// Version added: v1.5.0
	OutputNominalVoltageType NominalVoltageType
	// OutputRails shall contain an array of output power rails provided by this
	// power supply. The elements shall be ordered in ascending nominal voltage
	// order. This ordering is necessary for consistency with 'Sensor' properties
	// in an associated 'PowerSupplyMetrics' resource.
	OutputRails []OutputRail
	// PartNumber shall contain the part number as defined by the manufacturer for
	// this power supply.
	PartNumber string
	// PhaseWiringType shall contain the number of ungrounded current-carrying
	// conductors (phases) and the total number of conductors (wires) included in
	// the input connector for the power supply.
	PhaseWiringType PhaseWiringType
	// PlugType shall contain the type of physical plug used for the input to this
	// power supply, as defined by IEC, NEMA, or regional standards.
	PlugType PlugType
	// PowerCapacityWatts shall contain the maximum amount of power, in watt units,
	// that this power supply is rated to deliver.
	PowerCapacityWatts *float64 `json:",omitempty"`
	// PowerSupplyType shall contain the input power type (AC or DC) of this power
	// supply.
	PowerSupplyType PowerSupplyUnitType
	// ProductionDate shall contain the date of production or manufacture for this
	// power supply.
	//
	// Version added: v1.1.0
	ProductionDate string
	// Replaceable shall indicate whether this component can be independently
	// replaced as allowed by the vendor's replacement policy. A value of 'false'
	// indicates the component needs to be replaced by policy as part of another
	// component. If the 'LocationType' property of this component contains
	// 'Embedded', this property shall contain 'false'.
	//
	// Version added: v1.5.0
	Replaceable bool
	// SerialNumber shall contain the serial number as defined by the manufacturer
	// for this power supply.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as
	// defined by the manufacturer for this power supply.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// Version shall contain the hardware version of this power supply as
	// determined by the vendor or supplier.
	//
	// Version added: v1.1.0
	Version string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// outlet is the URI for Outlet.
	outlet string
	// powerOutlets are the URIs for PowerOutlets.
	powerOutlets []string
	// poweringChassis are the URIs for PoweringChassis.
	poweringChassis []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a PowerSupplyUnit object from the raw JSON.
func (p *PowerSupplyUnit) UnmarshalJSON(b []byte) error {
	type temp PowerSupplyUnit
	type pActions struct {
		Reset ActionTarget `json:"#PowerSupply.Reset"`
	}
	type pLinks struct {
		Outlet          Link  `json:"Outlet"`
		PowerOutlets    Links `json:"PowerOutlets"`
		PoweringChassis Links `json:"PoweringChassis"`
	}
	var tmp struct {
		temp
		Actions      pActions
		Links        pLinks
		Assembly     Link `json:"Assembly"`
		Certificates Link `json:"Certificates"`
		Metrics      Link `json:"Metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PowerSupplyUnit(tmp.temp)

	// Extract the links to other entities for later
	p.resetTarget = tmp.Actions.Reset.Target
	p.outlet = tmp.Links.Outlet.String()
	p.powerOutlets = tmp.Links.PowerOutlets.ToStrings()
	p.poweringChassis = tmp.Links.PoweringChassis.ToStrings()
	p.assembly = tmp.Assembly.String()
	p.certificates = tmp.Certificates.String()
	p.metrics = tmp.Metrics.String()

	// This is a read/write object, so we need to save the raw object data for later
	p.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *PowerSupplyUnit) Update() error {
	readWriteFields := []string{
		"ElectricalSourceManagerURIs",
		"ElectricalSourceNames",
		"LocationIndicatorActive",
	}

	return p.UpdateFromRawData(p, p.RawData, readWriteFields)
}

// GetPowerSupplyUnit will get a PowerSupplyUnit instance from the service.
func GetPowerSupplyUnit(c Client, uri string) (*PowerSupplyUnit, error) {
	return GetObject[PowerSupplyUnit](c, uri)
}

// ListReferencedPowerSupplyUnits gets the collection of PowerSupplyUnit from
// a provided reference.
func ListReferencedPowerSupplyUnits(c Client, link string) ([]*PowerSupplyUnit, error) {
	return GetCollectionObjects[PowerSupplyUnit](c, link)
}

// This action shall reset a power supply. A 'GracefulRestart' 'ResetType'
// shall reset the power supply but shall not affect the power output. A
// 'ForceRestart' 'ResetType' can affect the power supply output.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and shall perform a
// 'GracefulRestart'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *PowerSupplyUnit) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(p.client,
		p.resetTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Outlet gets the Outlet linked resource.
func (p *PowerSupplyUnit) Outlet() (*Outlet, error) {
	if p.outlet == "" {
		return nil, nil
	}
	return GetObject[Outlet](p.client, p.outlet)
}

// PowerOutlets gets the PowerOutlets linked resources.
func (p *PowerSupplyUnit) PowerOutlets() ([]*Outlet, error) {
	return GetObjects[Outlet](p.client, p.powerOutlets)
}

// PoweringChassis gets the PoweringChassis linked resources.
func (p *PowerSupplyUnit) PoweringChassis() ([]*Chassis, error) {
	return GetObjects[Chassis](p.client, p.poweringChassis)
}

// Assembly gets the Assembly linked resource.
func (p *PowerSupplyUnit) Assembly() (*Assembly, error) {
	if p.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](p.client, p.assembly)
}

// Certificates gets the Certificates collection.
func (p *PowerSupplyUnit) Certificates() ([]*Certificate, error) {
	if p.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](p.client, p.certificates)
}

// Metrics gets the Metrics linked resource.
func (p *PowerSupplyUnit) Metrics() (*PowerSupplyMetrics, error) {
	if p.metrics == "" {
		return nil, nil
	}
	return GetObject[PowerSupplyMetrics](p.client, p.metrics)
}

// EfficiencyRating shall describe an efficiency rating for a power supply.
type EfficiencyRating struct {
	// EfficiencyPercent shall contain the rated efficiency, as a percentage, '0'
	// to '100', of this power supply at the specified load.
	EfficiencyPercent *float64 `json:",omitempty"`
	// LoadPercent shall contain the load, as a percentage, '0' to '100', of this
	// power supply at which this efficiency rating is valid.
	LoadPercent *float64 `json:",omitempty"`
}

// PowerSupplyUnitInputRange shall describe an input range that the associated power supply can
// utilize.
type PowerSupplyUnitInputRange struct {
	// CapacityWatts shall contain the maximum amount of power, in watt units, that
	// the associated power supply is rated to deliver while operating in this
	// input range.
	CapacityWatts *float64 `json:",omitempty"`
	// NominalVoltageType shall contain the input voltage type of the associated
	// range.
	NominalVoltageType NominalVoltageType
}

// OutputRail shall describe an output power rail provided by a power supply.
type OutputRail struct {
	// NominalVoltage shall contain the nominal voltage of the associated output
	// power rail.
	NominalVoltage *float64 `json:",omitempty"`
	// PhysicalContext shall contain a description of the device or region within
	// the chassis to which this power rail applies.
	PhysicalContext PhysicalContext
}
