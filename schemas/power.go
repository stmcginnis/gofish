//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.4 - #Power.v1_7_3.Power

package schemas

import (
	"encoding/json"
	"math"
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

type PowerSupplyType string

const (
	// UnknownPowerSupplyType The power supply type cannot be determined.
	UnknownPowerSupplyType PowerSupplyType = "Unknown"
	// ACPowerSupplyType Alternating Current (AC) power supply.
	ACPowerSupplyType PowerSupplyType = "AC"
	// DCPowerSupplyType Direct Current (DC) power supply.
	DCPowerSupplyType PowerSupplyType = "DC"
	// ACorDCPowerSupplyType The power supply supports both DC and AC.
	ACorDCPowerSupplyType PowerSupplyType = "ACorDC"
)

// Power shall contain the power metrics for a Redfish implementation.
type Power struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PowerControl shall contain the set of power control readings and settings.
	PowerControl []PowerControl
	// PowerControlCount
	PowerControlCount int `json:"PowerControl@odata.count"`
	// PowerSupplies shall contain the set of power supplies associated with this
	// system or device.
	PowerSupplies []PowerSupply
	// PowerSuppliesCount
	PowerSuppliesCount int `json:"PowerSupplies@odata.count"`
	// Redundancy shall contain redundancy information for the set of power
	// supplies in this system or device.
	Redundancy []Redundancy
	// RedundancyCount
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Voltages shall contain the set of voltage sensors for this chassis.
	Voltages []Voltage
	// VoltagesCount
	VoltagesCount int `json:"Voltages@odata.count"`
	// powerSupplyResetTarget is the URL to send PowerSupplyReset requests.
	powerSupplyResetTarget string
}

// UnmarshalJSON unmarshals a Power object from the raw JSON.
func (p *Power) UnmarshalJSON(b []byte) error {
	type temp Power
	type pActions struct {
		PowerSupplyReset ActionTarget `json:"#Power.PowerSupplyReset"`
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

	return nil
}

// GetPower will get a Power instance from the service.
func GetPower(c Client, uri string) (*Power, error) {
	return GetObject[Power](c, uri)
}

// ListReferencedPowers gets the collection of Power from
// a provided reference.
func ListReferencedPowers(c Client, link string) ([]*Power, error) {
	return GetCollectionObjects[Power](c, link)
}

// This action shall reset a power supply specified by the 'MemberId' from the
// 'PowerSupplies' array. A 'GracefulRestart' 'ResetType' shall reset the power
// supply but shall not affect the power output. A 'ForceRestart' 'ResetType'
// can affect the power supply output.
// memberID - This parameter shall contain the identifier of the member within
// the 'PowerSupplies' array on which to perform the reset.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and shall perform a
// 'GracefulRestart'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *Power) PowerSupplyReset(memberID string, resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["MemberId"] = memberID
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(p.client,
		p.powerSupplyResetTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// InputRange shall describe an input range that the associated power supply can
// utilize.
type InputRange struct {
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
	// OEM shall contain the OEM extensions. All values for properties contained in
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
	Entity
	// MemberID shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// OEM shall contain the OEM extensions. All values for properties that this
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
	PowerRequestedWatts *float64 `json:",omitempty"`
	// RelatedItem shall contain an array of links to resources or objects
	// associated with this power limit.
	relatedItem []string
	// RelatedItemCount
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// powerSupplyResetTarget is the URL to send PowerSupplyReset requests.
	powerSupplyResetTarget string
}

// UnmarshalJSON unmarshals a PowerControl object from the raw JSON.
func (p *PowerControl) UnmarshalJSON(b []byte) error {
	type temp PowerControl
	type pActions struct {
		PowerSupplyReset ActionTarget `json:"#Power.PowerSupplyReset"`
	}
	var tmp struct {
		temp
		Actions     pActions
		RelatedItem Links `json:"RelatedItem"`

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
	p.MemberID = parseString(tmp.MemberID)
	p.PowerAllocatedWatts = toFloat32(tmp.PowerAllocatedWatts)
	p.PowerAvailableWatts = toFloat32(tmp.PowerAvailableWatts)
	p.PowerCapacityWatts = toFloat32(tmp.PowerCapacityWatts)
	p.PowerConsumedWatts = toFloat32(tmp.PowerConsumedWatts)

	// Extract the links to other entities for later
	p.powerSupplyResetTarget = tmp.Actions.PowerSupplyReset.Target
	p.relatedItem = tmp.RelatedItem.ToStrings()

	return nil
}

// GetPowerControl will get a PowerControl instance from the service.
func GetPowerControl(c Client, uri string) (*PowerControl, error) {
	return GetObject[PowerControl](c, uri)
}

// ListReferencedPowerControls gets the collection of PowerControl from
// a provided reference.
func ListReferencedPowerControls(c Client, link string) ([]*PowerControl, error) {
	return GetCollectionObjects[PowerControl](c, link)
}

// This action shall reset a power supply specified by the 'MemberId' from the
// 'PowerSupplies' array. A 'GracefulRestart' 'ResetType' shall reset the power
// supply but shall not affect the power output. A 'ForceRestart' 'ResetType'
// can affect the power supply output.
// memberID - This parameter shall contain the identifier of the member within
// the 'PowerSupplies' array on which to perform the reset.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and shall perform a
// 'GracefulRestart'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *PowerControl) PowerSupplyReset(memberID string, resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["MemberId"] = memberID
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(p.client,
		p.powerSupplyResetTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// RelatedItem gets the RelatedItem linked resources.
func (p *PowerControl) RelatedItem() ([]*Entity, error) {
	return GetObjects[Entity](p.client, p.relatedItem)
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
	IntervalInMin *uint `json:",omitempty"`
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
		val := uint(math.Round(float64(*toFloat32(t.IntervalInMin))))
		p.IntervalInMin = &val
	}
	p.MaxConsumedWatts = toFloat32(t.MaxConsumedWatts)
	p.MinConsumedWatts = toFloat32(t.MinConsumedWatts)

	return nil
}

// PowerSupply Details of a power supplies associated with this system or
// device.
type PowerSupply struct {
	Entity
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
	IndicatorLED IndicatorLED
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
	Location Location
	// Manufacturer shall contain the name of the organization responsible for
	// producing the power supply. This organization may be the entity from whom
	// the power supply is purchased, but this is not necessarily true.
	//
	// Version added: v1.1.0
	Manufacturer string
	// MemberID shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// Model shall contain the model information as defined by the manufacturer for
	// the associated power supply.
	Model string
	// OEM shall contain the OEM extensions. All values for properties that this
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
	// Redundancy shall contain an array of links to the redundancy groups to which
	// this power supply belongs.
	redundancy []string
	// RedundancyCount
	RedundancyCount int `json:"Redundancy@odata.count"`
	// RelatedItem shall contain an array of links to resources or objects
	// associated with this power supply.
	relatedItem []string
	// RelatedItemCount
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SerialNumber shall contain the serial number as defined by the manufacturer
	// for the associated power supply.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as
	// defined by the manufacturer for the associated power supply.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// powerSupplyResetTarget is the URL to send PowerSupplyReset requests.
	powerSupplyResetTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a PowerSupply object from the raw JSON.
func (p *PowerSupply) UnmarshalJSON(b []byte) error {
	type temp PowerSupply
	type pActions struct {
		PowerSupplyReset ActionTarget `json:"#Power.PowerSupplyReset"`
	}
	var tmp struct {
		temp
		Actions              pActions
		Assembly             Link  `json:"Assembly"`
		Redundancy           Links `json:"Redundancy"`
		RelatedItem          Links `json:"RelatedItem"`
		MemberID             any   `json:"MemberId"`
		LineInputVoltage     any
		LastPowerOutputWatts any
		PowerInputWatts      any
		PowerOutputWatts     any
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PowerSupply(tmp.temp)

	// Extract the links to other entities for later
	p.powerSupplyResetTarget = tmp.Actions.PowerSupplyReset.Target
	p.assembly = tmp.Assembly.String()
	p.redundancy = tmp.Redundancy.ToStrings()
	p.relatedItem = tmp.RelatedItem.ToStrings()

	p.MemberID = parseString(tmp.MemberID)
	p.LineInputVoltage = toFloat32(tmp.LineInputVoltage)
	p.LastPowerOutputWatts = toFloat32(tmp.LastPowerOutputWatts)
	p.PowerInputWatts = toFloat32(tmp.PowerInputWatts)
	p.PowerOutputWatts = toFloat32(tmp.PowerOutputWatts)

	// This is a read/write object, so we need to save the raw object data for later
	p.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *PowerSupply) Update() error {
	readWriteFields := []string{
		"IndicatorLED",
	}

	return p.UpdateFromRawData(p, p.RawData, readWriteFields)
}

// GetPowerSupply will get a PowerSupply instance from the service.
func GetPowerSupply(c Client, uri string) (*PowerSupply, error) {
	return GetObject[PowerSupply](c, uri)
}

// ListReferencedPowerSupplies gets the collection of PowerSupply from
// a provided reference.
func ListReferencedPowerSupplies(c Client, link string) ([]*PowerSupply, error) {
	return GetCollectionObjects[PowerSupply](c, link)
}

// This action shall reset a power supply specified by the 'MemberId' from the
// 'PowerSupplies' array. A 'GracefulRestart' 'ResetType' shall reset the power
// supply but shall not affect the power output. A 'ForceRestart' 'ResetType'
// can affect the power supply output.
// memberID - This parameter shall contain the identifier of the member within
// the 'PowerSupplies' array on which to perform the reset.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and shall perform a
// 'GracefulRestart'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *PowerSupply) PowerSupplyReset(memberID string, resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["MemberId"] = memberID
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(p.client,
		p.powerSupplyResetTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Assembly gets the Assembly linked resource.
func (p *PowerSupply) Assembly() (*Assembly, error) {
	if p.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](p.client, p.assembly)
}

// Redundancy gets the Redundancy linked resources.
func (p *PowerSupply) Redundancy() ([]*Redundancy, error) {
	return GetObjects[Redundancy](p.client, p.redundancy)
}

// RelatedItem gets the RelatedItem linked resources.
func (p *PowerSupply) RelatedItem() ([]*Entity, error) {
	return GetObjects[Entity](p.client, p.relatedItem)
}

// Voltage represents the Voltage type.
type Voltage struct {
	Entity
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
	// MemberID shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// MinReadingRange shall indicate the lowest possible value for the
	// 'ReadingVolts' property. The value of the property shall use the same units
	// as the 'ReadingVolts' property.
	MinReadingRange *float32 `json:",omitempty"`
	// OEM shall contain the OEM extensions. All values for properties that this
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
	// RelatedItemCount
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SensorNumber shall contain a numerical identifier for this voltage sensor
	// that is unique within this resource.
	SensorNumber *int `json:",omitempty"`
	// Status shall contain any status or health properties of the resource.
	Status Status
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
		PowerSupplyReset ActionTarget `json:"#Power.PowerSupplyReset"`
	}
	var tmp struct {
		temp
		Actions     vActions
		RelatedItem Links `json:"RelatedItem"`

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

	v.MemberID = parseString(tmp.MemberID)
	v.UpperThresholdCritical = toFloat32(tmp.UpperThresholdCritical)
	v.UpperThresholdFatal = toFloat32(tmp.UpperThresholdFatal)
	v.UpperThresholdNonCritical = toFloat32(tmp.UpperThresholdNonCritical)
	v.LowerThresholdCritical = toFloat32(tmp.LowerThresholdCritical)
	v.LowerThresholdFatal = toFloat32(tmp.LowerThresholdFatal)
	v.LowerThresholdNonCritical = toFloat32(tmp.LowerThresholdNonCritical)
	v.ReadingVolts = toFloat32(tmp.ReadingVolts)

	return nil
}

// GetVoltage will get a Voltage instance from the service.
func GetVoltage(c Client, uri string) (*Voltage, error) {
	return GetObject[Voltage](c, uri)
}

// ListReferencedVoltages gets the collection of Voltage from
// a provided reference.
func ListReferencedVoltages(c Client, link string) ([]*Voltage, error) {
	return GetCollectionObjects[Voltage](c, link)
}

// This action shall reset a power supply specified by the 'MemberId' from the
// 'PowerSupplies' array. A 'GracefulRestart' 'ResetType' shall reset the power
// supply but shall not affect the power output. A 'ForceRestart' 'ResetType'
// can affect the power supply output.
// memberID - This parameter shall contain the identifier of the member within
// the 'PowerSupplies' array on which to perform the reset.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and shall perform a
// 'GracefulRestart'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (v *Voltage) PowerSupplyReset(memberID string, resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["MemberId"] = memberID
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(v.client,
		v.powerSupplyResetTarget, payload, v.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// RedundancySet returns the power supplies in the specified redundancy group.
// memberID identifies the redundancy group in the Redundancy array.
func (p *Power) RedundancySet(memberID int) []PowerSupply {
	var powerSupplies []PowerSupply
	if len(p.Redundancy) >= memberID+1 {
		for _, psLink := range p.Redundancy[memberID].redundancySet {
			for i := range p.PowerSupplies {
				if p.PowerSupplies[i].ODataID == psLink {
					powerSupplies = append(powerSupplies, p.PowerSupplies[i])
				}
			}
		}
	}
	return powerSupplies
}

// RelatedItem gets the RelatedItem linked resources.
func (v *Voltage) RelatedItem() ([]*Entity, error) {
	return GetObjects[Entity](v.client, v.relatedItem)
}
