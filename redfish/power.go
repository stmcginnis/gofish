//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.4 - #Power.v1_7_3.Power

package redfish

import (
	"encoding/json"
	"math"
	"strconv"

	"github.com/stmcginnis/gofish/common"
)

type InputType string

const (
	// ACInputType Alternating Current (AC) input range.
	ACInputType InputType = "AC"
	// DCInputType Direct Current (DC) input range.
	DCInputType InputType = "DC"
)

type LineInputVoltageType string

const (
	// UnknownLineInputVoltageType The power supply line input voltage type cannot
	// be determined.
	UnknownLineInputVoltageType LineInputVoltageType = "Unknown"
	// ACLowLineLineInputVoltageType 100-127V AC input.
	ACLowLineLineInputVoltageType LineInputVoltageType = "ACLowLine"
	// ACMidLineLineInputVoltageType 200-240V AC input.
	ACMidLineLineInputVoltageType LineInputVoltageType = "ACMidLine"
	// ACHighLineLineInputVoltageType 277V AC input.
	ACHighLineLineInputVoltageType LineInputVoltageType = "ACHighLine"
	// DCNeg48VLineInputVoltageType -48V DC input.
	DCNeg48VLineInputVoltageType LineInputVoltageType = "DCNeg48V"
	// DC380VLineInputVoltageType High-voltage DC input (380V).
	DC380VLineInputVoltageType LineInputVoltageType = "DC380V"
	// AC120VLineInputVoltageType AC 120V nominal input.
	AC120VLineInputVoltageType LineInputVoltageType = "AC120V"
	// AC240VLineInputVoltageType AC 240V nominal input.
	AC240VLineInputVoltageType LineInputVoltageType = "AC240V"
	// AC277VLineInputVoltageType AC 277V nominal input.
	AC277VLineInputVoltageType LineInputVoltageType = "AC277V"
	// ACandDCWideRangeLineInputVoltageType Wide range AC or DC input.
	ACandDCWideRangeLineInputVoltageType LineInputVoltageType = "ACandDCWideRange"
	// ACWideRangeLineInputVoltageType Wide range AC input.
	ACWideRangeLineInputVoltageType LineInputVoltageType = "ACWideRange"
	// DC240VLineInputVoltageType DC 240V nominal input.
	DC240VLineInputVoltageType LineInputVoltageType = "DC240V"
)

type PowerLimitException string

const (
	// NoActionPowerLimitException Take no action when the limit is exceeded.
	NoActionPowerLimitException PowerLimitException = "NoAction"
	// HardPowerOffPowerLimitException Turn the power off immediately when the
	// limit is exceeded.
	HardPowerOffPowerLimitException PowerLimitException = "HardPowerOff"
	// LogEventOnlyPowerLimitException Log an event when the limit is exceeded, but
	// take no further action.
	LogEventOnlyPowerLimitException PowerLimitException = "LogEventOnly"
	// OemPowerLimitException Take an OEM-defined action.
	OemPowerLimitException PowerLimitException = "Oem"
)

// Power shall contain the power metrics for a Redfish implementation.
type Power struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerControl shall contain the set of power control readings and settings.
	PowerControl []PowerControl
	// PowerControl@odata.count
	PowerControlCount int `json:"PowerControl@odata.count"`
	// PowerSupplies shall contain the set of power supplies associated with this
	// system or device.
	PowerSupplies []PowerPowerSupply
	// PowerSupplies@odata.count
	PowerSuppliesCount int `json:"PowerSupplies@odata.count"`
	// Redundancy shall contain redundancy information for the set of power
	// supplies in this system or device.
	Redundancy []common.Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Voltages shall contain the set of voltage sensors for this chassis.
	Voltages []Voltage
	// Voltages@odata.count
	VoltagesCount int `json:"Voltages@odata.count"`
	// powerSupplyResetTarget is the URL to send PowerSupplyReset requests.
	powerSupplyResetTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Power object from the raw JSON.
func (p *Power) UnmarshalJSON(b []byte) error {
	type temp Power
	type pActions struct {
		PowerSupplyReset common.ActionTarget `json:"#Power.PowerSupplyReset"`
	}
	var tmp struct {
		temp
		Actions      pActions
		PowerControl json.RawMessage
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = Power(tmp.temp)

	powerControls := []PowerControl{}
	err = json.Unmarshal(tmp.PowerControl, &powerControls)
	if err != nil {
		// Some Cisco implementations return a singular object instead of the
		// expected array.
		powerControl := PowerControl{}
		err2 := json.Unmarshal(tmp.PowerControl, &powerControl)
		if err2 != nil {
			// Return the original error
			return err
		}

		powerControls = append(powerControls, powerControl)
	}

	p.PowerControl = powerControls

	// Extract the links to other entities for later
	p.powerSupplyResetTarget = tmp.Actions.PowerSupplyReset.Target

	// This is a read/write object, so we need to save the raw object data for later
	p.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *Power) Update() error {
	readWriteFields := []string{
		"PowerControl",
		"PowerControl@odata.count",
		"PowerSupplies",
		"PowerSupplies@odata.count",
		"Redundancy",
		"Redundancy@odata.count",
		"Voltages",
		"Voltages@odata.count",
	}

	return p.UpdateFromRawData(p, p.rawData, readWriteFields)
}

// GetPower will get a Power instance from the service.
func GetPower(c common.Client, uri string) (*Power, error) {
	return common.GetObject[Power](c, uri)
}

// ListReferencedPowers gets the collection of Power from
// a provided reference.
func ListReferencedPowers(c common.Client, link string) ([]*Power, error) {
	return common.GetCollectionObjects[Power](c, link)
}

// PowerSupplyReset shall reset a power supply specified by the 'MemberId' from the
// 'PowerSupplies' array. A 'GracefulRestart' 'ResetType' shall reset the power
// supply but shall not affect the power output. A 'ForceRestart' 'ResetType'
// can affect the power supply output.
// memberID - This parameter shall contain the identifier of the member within
// the 'PowerSupplies' array on which to perform the reset.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and shall perform a
// 'GracefulRestart'.
func (p *Power) PowerSupplyReset(memberID string, resetType common.ResetType) error {
	payload := make(map[string]any)
	payload["MemberId"] = memberID
	payload["ResetType"] = resetType
	return p.Post(p.powerSupplyResetTarget, payload)
}

// PowerInputRange shall describe an input range that the associated power supply can
// utilize.
type PowerInputRange struct {
	// InputType shall contain the input type (AC or DC) of the associated range.
	//
	// Version added: v1.1.0
	InputType InputType
	// MaximumFrequencyHz shall contain the value, in hertz units, of the maximum
	// line input frequency that the power supply is capable of consuming for this
	// range.
	//
	// Version added: v1.1.0
	MaximumFrequencyHz *float64 `json:",omitempty"`
	// MaximumVoltage shall contain the value, in volt units, of the maximum line
	// input voltage that the power supply is capable of consuming for this range.
	//
	// Version added: v1.1.0
	MaximumVoltage *float64 `json:",omitempty"`
	// MinimumFrequencyHz shall contain the value, in hertz units, of the minimum
	// line input frequency that the power supply is capable of consuming for this
	// range.
	//
	// Version added: v1.1.0
	MinimumFrequencyHz *float64 `json:",omitempty"`
	// MinimumVoltage shall contain the value, in volt units, of the minimum line
	// input voltage that the power supply is capable of consuming for this range.
	//
	// Version added: v1.1.0
	MinimumVoltage *float64 `json:",omitempty"`
	// Oem shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.1.0
	OEM json.RawMessage `json:"Oem"`
	// OutputWattage shall contain the maximum amount of power, in watt units, that
	// the associated power supply is rated to deliver while operating in this
	// input range.
	//
	// Version added: v1.1.0
	OutputWattage *float64 `json:",omitempty"`
}

// PowerControl represents the PowerControl type.
type PowerControl struct {
	common.Entity
	// MemberId shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// Name is the name of the resource or array element.
	Name string
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhysicalContext shall contain a description of the affected device(s) or
	// region within the chassis to which this power control applies.
	//
	// Version added: v1.4.0
	PhysicalContext PhysicalContext
	// PowerAllocatedWatts shall represent the total power currently allocated or
	// budgeted to the chassis.
	PowerAllocatedWatts *float32 `json:",omitempty"`
	// PowerAvailableWatts shall represent the amount of reserve power capacity, in
	// watt units, that remains. This value is the PowerCapacityWatts value minus
	// the 'PowerAllocatedWatts' value.
	PowerAvailableWatts *float32 `json:",omitempty"`
	// PowerCapacityWatts shall represent the total power capacity that can be
	// allocated to the chassis.
	PowerCapacityWatts *float32 `json:",omitempty"`
	// PowerConsumedWatts shall represent the actual power that the chassis
	// consumes, in watt units.
	PowerConsumedWatts *float32 `json:",omitempty"`
	// PowerLimit shall contain power limit status and configuration information
	// for this chassis.
	PowerLimit PowerLimit
	// PowerMetrics shall contain power metrics for power readings, such as
	// interval, minimum, maximum, and average power consumption, for the chassis.
	PowerMetrics PowerMetric
	// PowerRequestedWatts shall represent the amount of power, in watt units, that
	// the chassis currently requests to be budgeted for future use.
	PowerRequestedWatts *float32 `json:",omitempty"`
	// RelatedItem shall contain an array of links to resources or objects
	// associated with this power limit.
	relatedItem []string
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// powerSupplyResetTarget is the URL to send PowerSupplyReset requests.
	powerSupplyResetTarget string
}

// UnmarshalJSON unmarshals a PowerControl object from the raw JSON.
func (p *PowerControl) UnmarshalJSON(b []byte) error {
	type temp PowerControl
	type pActions struct {
		PowerSupplyReset common.ActionTarget `json:"#Power.PowerSupplyReset"`
	}
	var tmp struct {
		temp
		Actions     pActions
		RelatedItem common.Links `json:"RelatedItem"`

		// Need to work around some non-standard data types in Dell and Cisco
		// systems.
		MemberID            any `json:"MemberId"`
		PowerAllocatedWatts any
		PowerAvailableWatts any
		PowerCapacityWatts  any
		PowerConsumedWatts  any
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PowerControl(tmp.temp)

	// Standardize the property types
	p.relatedItem = tmp.RelatedItem.ToStrings()
	p.MemberID = parseMemberID(tmp.MemberID)
	p.PowerAllocatedWatts = toFloat32(tmp.PowerAllocatedWatts)
	p.PowerAvailableWatts = toFloat32(tmp.PowerAvailableWatts)
	p.PowerCapacityWatts = toFloat32(tmp.PowerCapacityWatts)
	p.PowerConsumedWatts = toFloat32(tmp.PowerConsumedWatts)

	// Extract the links to other entities for later
	p.powerSupplyResetTarget = tmp.Actions.PowerSupplyReset.Target

	return nil
}

// RelatedItem gets the RelatedItem linked resources.
func (p *PowerControl) RelatedItem(client common.Client) ([]*common.Entity, error) {
	return common.GetObjects[common.Entity](client, p.relatedItem)
}

// PowerSupplyReset shall reset a power supply specified by the 'MemberId' from the
// 'PowerSupplies' array. A 'GracefulRestart' 'ResetType' shall reset the power
// supply but shall not affect the power output. A 'ForceRestart' 'ResetType'
// can affect the power supply output.
// memberID - This parameter shall contain the identifier of the member within
// the 'PowerSupplies' array on which to perform the reset.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and shall perform a
// 'GracefulRestart'.
func (p *PowerControl) PowerSupplyReset(memberID string, resetType common.ResetType) error {
	payload := make(map[string]any)
	payload["MemberId"] = memberID
	payload["ResetType"] = resetType
	return p.Post(p.powerSupplyResetTarget, payload)
}

// PowerLimit shall contain power limit status and configuration information for
// this chassis.
type PowerLimit struct {
	// CorrectionInMs shall represent the time interval in ms required for the
	// limiting process to react and reduce the power consumption below the limit.
	CorrectionInMs *int `json:",omitempty"`
	// LimitException shall represent the action to be taken if the resource power
	// consumption cannot be limited below the specified limit after several
	// correction time periods.
	LimitException PowerLimitException
	// LimitInWatts shall represent the power capping limit, in watt units, for the
	// resource. If 'null', power capping shall be disabled.
	LimitInWatts *float64 `json:",omitempty"`
}

// PowerMetric shall contain power metrics for power readings, such as interval,
// minimum, maximum, and average power consumption, for a resource.
type PowerMetric struct {
	// AverageConsumedWatts shall represent the average power level that occurred
	// over the last 'IntervalInMin' minutes.
	AverageConsumedWatts *float32 `json:",omitempty"`
	// IntervalInMin shall represent the time interval or window, in minutes, over
	// which the power metrics are measured.
	IntervalInMin *int `json:",omitempty"`
	// MaxConsumedWatts shall represent the maximum power level, in watt units,
	// that occurred within the last 'IntervalInMin' minutes.
	MaxConsumedWatts *float32 `json:",omitempty"`
	// MinConsumedWatts shall represent the minimum power level, in watt units,
	// that occurred within the last 'IntervalInMin' minutes.
	MinConsumedWatts *float32 `json:",omitempty"`
}

func (p *PowerMetric) UnmarshalJSON(b []byte) error {
	type temp PowerMetric
	type t1 struct {
		temp
		IntervalInMin    any
		MaxConsumedWatts any
		MinConsumedWatts any
	}
	var t t1

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*p = PowerMetric(t.temp)

	if t.IntervalInMin != nil {
		val := int(math.Round(float64(*toFloat32(t.IntervalInMin))))
		p.IntervalInMin = &val
	}
	p.MaxConsumedWatts = toFloat32(t.MaxConsumedWatts)
	p.MinConsumedWatts = toFloat32(t.MinConsumedWatts)

	return nil
}

// InputRange describes an input range that a power supply can utilize.
type InputRange struct {
	// InputType indicates the input type (AC or DC) of this range.
	InputType InputType `json:"InputType,omitempty"`
	// MaximumFrequencyHz indicates the maximum line input frequency in Hertz that the power supply can use for this range.
	MaximumFrequencyHz *float32 `json:"MaximumFrequencyHz,omitempty"`
	// MaximumVoltage indicates the maximum line input voltage in Volts that the power supply can use for this range.
	MaximumVoltage *float32 `json:"MaximumVoltage,omitempty"`
	// MinimumFrequencyHz indicates the minimum line input frequency in Hertz that the power supply can use for this range.
	MinimumFrequencyHz *float32 `json:"MinimumFrequencyHz,omitempty"`
	// MinimumVoltage indicates the minimum line input voltage in Volts that the power supply can use for this range.
	MinimumVoltage *float32 `json:"MinimumVoltage,omitempty"`
	// OutputWattage indicates the maximum power in Watts that the power supply can deliver while operating in this range.
	OutputWattage *float32 `json:"OutputWattage,omitempty"`
	// Oem contains OEM-specific extensions.
	OEM json.RawMessage `json:"Oem,omitempty"`
}

// PowerPowerSupply Details of a power supplies associated with this system or
// device.
type PowerPowerSupply struct {
	common.Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	//
	// Version added: v1.5.0
	assembly string
	// EfficiencyPercent shall contain the measured power efficiency, as a
	// percentage, of the associated power supply.
	//
	// Version added: v1.5.0
	EfficiencyPercent *float64 `json:",omitempty"`
	// FirmwareVersion shall contain the firmware version as defined by the
	// manufacturer for the associated power supply.
	FirmwareVersion string
	// HotPluggable shall indicate whether the device can be inserted or removed
	// while the underlying equipment otherwise remains in its current operational
	// state. Devices indicated as hot-pluggable shall allow the device to become
	// operable without altering the operational state of the underlying equipment.
	// Devices that cannot be inserted or removed from equipment in operation, or
	// devices that cannot become operable without affecting the operational state
	// of that equipment, shall be indicated as not hot-pluggable.
	//
	// Version added: v1.5.0
	HotPluggable bool
	// IndicatorLED shall contain the indicator light state for the indicator light
	// associated with this power supply.
	//
	// Version added: v1.2.0
	IndicatorLED common.IndicatorLED
	// InputRanges shall contain a collection of ranges usable by the power supply
	// unit.
	//
	// Version added: v1.1.0
	InputRanges []InputRange
	// LastPowerOutputWatts shall contain the average power output, measured in
	// watt units, of the associated power supply.
	LastPowerOutputWatts *float32 `json:",omitempty"`
	// LineInputVoltage shall contain the value in volt units of the line input
	// voltage (measured or configured for) that the power supply has been
	// configured to operate with or is currently receiving.
	LineInputVoltage *float32 `json:",omitempty"`
	// LineInputVoltageType shall contain the type of input line voltage supported
	// by the associated power supply.
	LineInputVoltageType LineInputVoltageType
	// Location shall contain the location information of the associated power
	// supply.
	//
	// Version added: v1.5.0
	Location common.Location
	// Manufacturer shall contain the name of the organization responsible for
	// producing the power supply. This organization may be the entity from whom
	// the power supply is purchased, but this is not necessarily true.
	//
	// Version added: v1.1.0
	Manufacturer string
	// MemberId shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// Model shall contain the model information as defined by the manufacturer for
	// the associated power supply.
	Model string
	// Name is the name of the resource or array element.
	Name string
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for
	// the associated power supply.
	PartNumber string
	// PowerCapacityWatts shall contain the maximum amount of power, in watt units,
	// that the associated power supply is rated to deliver.
	PowerCapacityWatts *float32 `json:",omitempty"`
	// PowerInputWatts shall contain the measured input power, in watt units, of
	// the associated power supply.
	//
	// Version added: v1.5.0
	PowerInputWatts *float32 `json:",omitempty"`
	// PowerOutputWatts shall contain the measured output power, in watt units, of
	// the associated power supply.
	//
	// Version added: v1.5.0
	PowerOutputWatts *float32 `json:",omitempty"`
	// PowerSupplyType shall contain the input power type (AC or DC) of the
	// associated power supply.
	PowerSupplyType PowerSupplyType
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// RelatedItem shall contain an array of links to resources or objects
	// associated with this power supply.
	relatedItem []string
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SerialNumber shall contain the serial number as defined by the manufacturer
	// for the associated power supply.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as
	// defined by the manufacturer for the associated power supply.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status

	rawData         []byte
	metrics         string
	redundancyLinks []string
	// powerSupplyResetTarget is the URL to send PowerSupplyReset requests.
	powerSupplyResetTarget string
}

// UnmarshalJSON unmarshals a PowerPowerSupply object from the raw JSON.
func (p *PowerPowerSupply) UnmarshalJSON(b []byte) error {
	type temp PowerPowerSupply
	type pActions struct {
		PowerSupplyReset common.ActionTarget `json:"#Power.PowerSupplyReset"`
	}
	var tmp struct {
		temp
		Actions              pActions
		Assembly             common.Link `json:"assembly"`
		Metrics              common.Link
		Redundancy           common.Links
		RelatedItem          common.Links
		MemberID             any `json:"MemberId"`
		LineInputVoltage     any
		LastPowerOutputWatts any
		PowerInputWatts      any
		PowerOutputWatts     any
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PowerPowerSupply(tmp.temp)

	// Extract the links to other entities for later
	p.powerSupplyResetTarget = tmp.Actions.PowerSupplyReset.Target
	p.assembly = tmp.Assembly.String()
	p.metrics = tmp.Metrics.String()
	p.redundancyLinks = tmp.Redundancy.ToStrings()
	p.relatedItem = tmp.RelatedItem.ToStrings()

	p.MemberID = parseMemberID(tmp.MemberID)
	p.LineInputVoltage = toFloat32(tmp.LineInputVoltage)
	p.LastPowerOutputWatts = toFloat32(tmp.LastPowerOutputWatts)
	p.PowerInputWatts = toFloat32(tmp.PowerInputWatts)
	p.PowerOutputWatts = toFloat32(tmp.PowerOutputWatts)

	// This is a read/write object, so we need to save the raw object data for later
	p.rawData = b

	return nil
}

// RelatedItem gets the RelatedItem linked resources.
func (p *PowerPowerSupply) RelatedItem(client common.Client) ([]*common.Entity, error) {
	return common.GetObjects[common.Entity](client, p.relatedItem)
}

// PowerSupplyReset shall reset a power supply specified by the 'MemberId' from the
// 'PowerSupplies' array. A 'GracefulRestart' 'ResetType' shall reset the power
// supply but shall not affect the power output. A 'ForceRestart' 'ResetType'
// can affect the power supply output.
// memberID - This parameter shall contain the identifier of the member within
// the 'PowerSupplies' array on which to perform the reset.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and shall perform a
// 'GracefulRestart'.
func (p *PowerPowerSupply) Reset(memberID string, resetType common.ResetType) error {
	payload := make(map[string]any)
	payload["MemberId"] = memberID
	payload["ResetType"] = resetType
	return p.Post(p.powerSupplyResetTarget, payload)
}

// Assembly gets the Assembly linked resource.
func (p *PowerPowerSupply) Assembly(client common.Client) (*Assembly, error) {
	if p.assembly == "" {
		return nil, nil
	}
	return common.GetObject[Assembly](client, p.assembly)
}

// Metrics gets the metrics associated with this power supply.
func (powersupply *PowerPowerSupply) Metrics() (*PowerSupplyMetrics, error) {
	if powersupply.metrics == "" {
		return nil, nil
	}
	return GetPowerSupplyMetrics(powersupply.GetClient(), powersupply.metrics)
}

// Redundancy gets the endpoints at the other end of the link.
func (powersupply *PowerPowerSupply) Redundancy() ([]*common.Redundancy, error) {
	return common.GetObjects[common.Redundancy](powersupply.GetClient(), powersupply.redundancyLinks)
}

// GetPowerPowerSupply retrieves a PowerPowerSupply instance from the service.
func GetPowerPowerSupply(c common.Client, uri string) (*PowerPowerSupply, error) {
	var powerSupply PowerPowerSupply
	return &powerSupply, powerSupply.Get(c, uri, &powerSupply)
}

// ListReferencedPowerPowerSupplies retrieves a collection of PowerSupplies from a reference.
func ListReferencedPowerPowerSupplies(c common.Client, link string) ([]*PowerPowerSupply, error) {
	return common.GetCollectionObjects[PowerPowerSupply](c, link)
}

// Update commits updates to this object's properties to the running system.
func (powersupply *PowerPowerSupply) Update() error {
	readWriteFields := []string{"IndicatorLED"}

	return powersupply.UpdateFromRawData(powersupply, powersupply.rawData, readWriteFields)
}

// Voltage represents the Voltage type.
type Voltage struct {
	common.Entity
	// LowerThresholdCritical shall contain the value at which the 'ReadingVolts'
	// property is below the normal range but is not yet fatal. The value of the
	// property shall use the same units as the 'ReadingVolts' property.
	LowerThresholdCritical *float32 `json:",omitempty"`
	// LowerThresholdFatal shall contain the value at which the 'ReadingVolts'
	// property is below the normal range and is fatal. The value of the property
	// shall use the same units as the 'ReadingVolts' property.
	LowerThresholdFatal *float32 `json:",omitempty"`
	// LowerThresholdNonCritical shall contain the value at which the
	// 'ReadingVolts' property is below normal range. The value of the property
	// shall use the same units as the 'ReadingVolts' property.
	LowerThresholdNonCritical *float32 `json:",omitempty"`
	// MaxReadingRange shall indicate the highest possible value for the
	// 'ReadingVolts' property. The value of the property shall use the same units
	// as the 'ReadingVolts' property.
	MaxReadingRange *float32 `json:",omitempty"`
	// MemberId shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// MinReadingRange shall indicate the lowest possible value for the
	// 'ReadingVolts' property. The value of the property shall use the same units
	// as the 'ReadingVolts' property.
	MinReadingRange *float32 `json:",omitempty"`
	// Name is the name of the resource or array element.
	Name string
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhysicalContext shall contain a description of the affected device or region
	// within the chassis to which this voltage measurement applies.
	PhysicalContext PhysicalContext
	// ReadingVolts shall contain the voltage sensor's reading.
	ReadingVolts *float32 `json:",omitempty"`
	// RelatedItem shall contain an array of links to resources or objects to which
	// this voltage measurement applies.
	relatedItem []string
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SensorNumber shall contain a numerical identifier for this voltage sensor
	// that is unique within this resource.
	SensorNumber *int `json:",omitempty"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UpperThresholdCritical shall contain the value at which the 'ReadingVolts'
	// property is above the normal range but is not yet fatal. The value of the
	// property shall use the same units as the 'ReadingVolts' property.
	UpperThresholdCritical *float32 `json:",omitempty"`
	// UpperThresholdFatal shall contain the value at which the 'ReadingVolts'
	// property is above the normal range and is fatal. The value of the property
	// shall use the same units as the 'ReadingVolts' property.
	UpperThresholdFatal *float32 `json:",omitempty"`
	// UpperThresholdNonCritical shall contain the value at which the
	// 'ReadingVolts' property is above the normal range. The value of the property
	// shall use the same units as the 'ReadingVolts' property.
	UpperThresholdNonCritical *float32 `json:",omitempty"`
	// powerSupplyResetTarget is the URL to send PowerSupplyReset requests.
	powerSupplyResetTarget string
}

// UnmarshalJSON unmarshals a Voltage object from the raw JSON.
func (v *Voltage) UnmarshalJSON(b []byte) error {
	type temp Voltage
	type vActions struct {
		PowerSupplyReset common.ActionTarget `json:"#Power.PowerSupplyReset"`
	}
	var tmp struct {
		temp
		Actions     vActions
		RelatedItem common.Links

		// Need to work around some non-standard data types in Dell and Cisco
		// systems.
		MemberID                  any `json:"MemberId"`
		UpperThresholdCritical    any
		UpperThresholdFatal       any
		UpperThresholdNonCritical any
		LowerThresholdCritical    any
		LowerThresholdFatal       any
		LowerThresholdNonCritical any
		ReadingVolts              any
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*v = Voltage(tmp.temp)

	// Extract the links to other entities for later
	v.powerSupplyResetTarget = tmp.Actions.PowerSupplyReset.Target
	v.relatedItem = tmp.RelatedItem.ToStrings()

	// Standardize the property types
	v.MemberID = parseMemberID(tmp.MemberID)
	v.UpperThresholdCritical = toFloat32(tmp.UpperThresholdCritical)
	v.UpperThresholdFatal = toFloat32(tmp.UpperThresholdFatal)
	v.UpperThresholdNonCritical = toFloat32(tmp.UpperThresholdNonCritical)
	v.LowerThresholdCritical = toFloat32(tmp.LowerThresholdCritical)
	v.LowerThresholdFatal = toFloat32(tmp.LowerThresholdFatal)
	v.LowerThresholdNonCritical = toFloat32(tmp.LowerThresholdNonCritical)
	v.ReadingVolts = toFloat32(tmp.ReadingVolts)

	return nil
}

// RelatedItem gets the RelatedItem linked resources.
func (v *Voltage) RelatedItem(client common.Client) ([]*common.Entity, error) {
	return common.GetObjects[common.Entity](client, v.relatedItem)
}

// PowerSupplyReset shall reset a power supply specified by the 'MemberId' from the
// 'PowerSupplies' array. A 'GracefulRestart' 'ResetType' shall reset the power
// supply but shall not affect the power output. A 'ForceRestart' 'ResetType'
// can affect the power supply output.
// memberID - This parameter shall contain the identifier of the member within
// the 'PowerSupplies' array on which to perform the reset.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and shall perform a
// 'GracefulRestart'.
func (v *Voltage) PowerSupplyReset(memberID string, resetType common.ResetType) error {
	payload := make(map[string]any)
	payload["MemberId"] = memberID
	payload["ResetType"] = resetType
	return v.Post(v.powerSupplyResetTarget, payload)
}

func parseMemberID(val any) string {
	switch id := val.(type) {
	case string:
		return id
	case json.Number:
		return id.String()
	case int:
		return strconv.Itoa(id)
	case float32:
		return strconv.Itoa(int(id))
	case float64:
		return strconv.Itoa(int(id))
	}

	return ""
}

func toFloat32(val any) *float32 {
	if val == nil {
		return nil
	}

	var ret float32 = 0.0
	switch valu := val.(type) {
	case string:
		fl, err := strconv.ParseFloat(valu, 32)
		if err == nil {
			ret = float32(fl)
		}
	case int:
		ret = float32(valu)
	case float32:
		ret = float32(valu)
	case float64:
		conv := float32(valu)
		if math.IsInf(float64(conv), 1) {
			// Too big, return float32 max as a fallback
			ret = math.MaxFloat32
		} else if math.IsInf(float64(conv), 0) {
			// Too large negative
			ret = -math.MaxFloat32
		} else {
			ret = conv
		}
	}

	return &ret
}

func toInt(val any) *int {
	if val == nil {
		return nil
	}

	ret := int(*toFloat32(val))
	return &ret
}
