//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"
	"strconv"

	"github.com/stmcginnis/gofish/common"
)

// InputType is the type of power input.
type InputType string

const (

	// ACInputType Alternating Current (AC) input range.
	ACInputType InputType = "AC"
	// DCInputType Direct Current (DC) input range.
	DCInputType InputType = "DC"
)

// LineInputVoltageType is the type of line voltage.
type LineInputVoltageType string

const (

	// UnknownLineInputVoltageType The power supply line input voltage type
	// cannot be determined.
	UnknownLineInputVoltageType LineInputVoltageType = "Unknown"
	// ACLowLineLineInputVoltageType 100-127V AC input.
	ACLowLineLineInputVoltageType LineInputVoltageType = "ACLowLine"
	// ACMidLineLineInputVoltageType 200-240V AC input.
	ACMidLineLineInputVoltageType LineInputVoltageType = "ACMidLine"
	// ACHighLineLineInputVoltageType 277V AC input.
	ACHighLineLineInputVoltageType LineInputVoltageType = "ACHighLine"
	// DCNeg48VLineInputVoltageType -48V DC input.
	DCNeg48VLineInputVoltageType LineInputVoltageType = "DCNeg48V"
	// DC380VLineInputVoltageType High Voltage DC input (380V).
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

// PowerLimitException is the type of power limit exception.
type PowerLimitException string

const (

	// NoActionPowerLimitException Take no action when the limit is exceeded.
	NoActionPowerLimitException PowerLimitException = "NoAction"
	// HardPowerOffPowerLimitException Turn the power off immediately when
	// the limit is exceeded.
	HardPowerOffPowerLimitException PowerLimitException = "HardPowerOff"
	// LogEventOnlyPowerLimitException Log an event when the limit is
	// exceeded, but take no further action.
	LogEventOnlyPowerLimitException PowerLimitException = "LogEventOnly"
	// OemPowerLimitException Take an OEM-defined action.
	OemPowerLimitException PowerLimitException = "Oem"
)

// PowerSupplyType is the type of power supply.
type PowerSupplyType string

const (

	// UnknownPowerSupplyType The power supply type cannot be determined.
	UnknownPowerSupplyType PowerSupplyType = "Unknown"
	// ACPowerSupplyType Alternating Current (AC) power supply.
	ACPowerSupplyType PowerSupplyType = "AC"
	// DCPowerSupplyType Direct Current (DC) power supply.
	DCPowerSupplyType PowerSupplyType = "DC"
	// ACorDCPowerSupplyType Power Supply supports both DC or AC.
	ACorDCPowerSupplyType PowerSupplyType = "ACorDC"
)

// InputRange shall describe an input range that the associated power supply is
// able to utilize.
type InputRange struct {
	// InputType shall contain the input type (AC or DC) of the associated range.
	InputType InputType
	// MaximumFrequencyHz shall contain the value in Hertz of the maximum line
	// input frequency which the power supply is capable of consuming for this range.
	MaximumFrequencyHz float32
	// MaximumVoltage shall contain the value in Volts of the maximum line input
	// voltage which the power supply is capable of consuming for this range.
	MaximumVoltage float32
	// MinimumFrequencyHz shall contain the value in Hertz of the minimum line
	// input frequency which the power supply is capable of consuming for this range.
	MinimumFrequencyHz float32
	// MinimumVoltage shall contain the value in Volts of the minimum line input
	// voltage which the power supply is capable of consuming for this range.
	MinimumVoltage float32
	// OutputWattage shall contain the maximum amount of power, in Watts, that
	// the associated power supply is rated to deliver while operating in this input range.
	OutputWattage float32
}

// Power is used to represent a power metrics resource for a Redfish
// implementation.
type Power struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// IndicatorLED shall contain the indicator light state for the indicator
	// light associated with this power supply.
	IndicatorLED common.IndicatorLED
	// PowerControl shall be the definition for power control (power reading and
	// limiting) for a Redfish implementation.
	PowerControl []PowerControl
	// PowerControlCount is the number of objects.
	PowerControlCount int `json:"PowerControl@odata.count"`
	// PowerSupplies shall contain details of the power supplies associated with
	// this system or device.
	PowerSupplies []PowerSupply
	// PowerSuppliesCount is the number of objects.
	PowerSuppliesCount int `json:"PowerSupplies@odata.count"`
	// Redundancy shall contain redundancy information for the power subsystem
	// of this system or device.
	Redundancy []Redundancy
	// RedundancyCount is the number of objects.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Voltages shall be the definition for voltage
	// sensors for a Redfish implementation.
	Voltages []Voltage
	// VoltagesCount is the number of objects.
	VoltagesCount int `json:"Voltages@odata.count"`
}

// GetPower will get a Power instance from the service.
func GetPower(c common.Client, uri string) (*Power, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var power Power
	err = json.NewDecoder(resp.Body).Decode(&power)
	if err != nil {
		return nil, err
	}

	power.SetClient(c)
	return &power, nil
}

// ListReferencedPowers gets the collection of Power from
// a provided reference.
func ListReferencedPowers(c common.Client, link string) ([]*Power, error) {
	var result []*Power
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, powerLink := range links.ItemLinks {
		power, err := GetPower(c, powerLink)
		if err != nil {
			return result, err
		}
		result = append(result, power)
	}

	return result, nil
}

// PowerControl is
type PowerControl struct {
	common.Entity

	// MemberID shall uniquely identify the member within the collection. For
	// services supporting Redfish v1.6 or higher, this value shall be the
	// zero-based array index.
	MemberID string `json:"MemberId"`
	// PhysicalContext shall be a description of the affected device(s) or region
	// within the chassis to which this power control applies.
	PhysicalContext common.PhysicalContext
	// PowerAllocatedWatts shall represent the total power currently allocated
	// to chassis resources.
	PowerAllocatedWatts float32
	// PowerAvailableWatts shall represent the amount of power capacity (in
	// Watts) not already allocated and shall equal PowerCapacityWatts -
	// PowerAllocatedWatts.
	PowerAvailableWatts float32
	// PowerCapacityWatts shall represent the total power capacity that is
	// available for allocation to the chassis resources.
	PowerCapacityWatts float32
	// PowerConsumedWatts shall represent the actual power being consumed (in
	// Watts) by the chassis.
	PowerConsumedWatts float32
	// PowerLimit shall contain power limit status and configuration information
	// for this chassis.
	PowerLimit PowerLimit
	// PowerMetrics shall contain power metrics for power readings (interval,
	// minimum/maximum/average power consumption) for the chassis.
	PowerMetrics PowerMetric
	// PowerRequestedWatts shall represent the
	// amount of power (in Watts) that the chassis resource is currently
	// requesting be budgeted to it for future use.
	PowerRequestedWatts float32
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a PowerControl object from the raw JSON.
func (powercontrol *PowerControl) UnmarshalJSON(b []byte) error { // nolint:dupl
	type temp PowerControl
	type t1 struct {
		temp
	}
	var t t1

	err := json.Unmarshal(b, &t)
	if err != nil {
		// See if we need to handle converting MemberID
		var t2 struct {
			t1
			MemberID int `json:"MemberId"`
		}
		err2 := json.Unmarshal(b, &t2)

		if err2 != nil {
			// Return the original error
			return err
		}

		// Convert the numeric member ID to a string
		t = t2.t1
		t.temp.MemberID = strconv.Itoa(t2.MemberID)
	}

	// Extract the links to other entities for later
	*powercontrol = PowerControl(t.temp)

	return nil
}

// PowerLimit shall contain power limit status and
// configuration information for this chassis.
type PowerLimit struct {
	// CorrectionInMs shall represent the time
	// interval in ms required for the limiting process to react and reduce
	// the power consumption below the limit.
	CorrectionInMs int64
	// LimitException shall represent the
	// action to be taken if the resource power consumption can not be
	// limited below the specified limit after several correction time
	// periods.
	LimitException PowerLimitException
	// LimitInWatts shall represent the power
	// cap limit in watts for the resource. If set to null, power capping
	// shall be disabled.
	LimitInWatts float32
}

// PowerMetric shall contain power metrics for power
// readings (interval, minimum/maximum/average power consumption) for a
// resource.
type PowerMetric struct {
	// AverageConsumedWatts shall represent the
	// average power level that occurred averaged over the last IntervalInMin
	// minutes.
	AverageConsumedWatts float32
	// IntervalInMin shall represent the time
	// interval (or window), in minutes, in which the PowerMetrics properties
	// are measured over.
	// Should be an integer, but some Dell implementations return as a float.
	IntervalInMin float32
	// MaxConsumedWatts shall represent the
	// maximum power level in watts that occurred within the last
	// IntervalInMin minutes.
	MaxConsumedWatts float32
	// MinConsumedWatts shall represent the
	// minimum power level in watts that occurred within the last
	// IntervalInMin minutes.
	MinConsumedWatts float32
}

// PowerSupply is Details of a power supplies associated with this system
// or device.
type PowerSupply struct {
	common.Entity

	// assembly shall be a link to a resource of type Assembly.
	assembly string
	// EfficiencyPercent shall contain the value of the measured power
	// efficiency, as a percentage, of the associated power supply.
	EfficiencyPercent float32
	// FirmwareVersion shall contain the firmware version as
	// defined by the manufacturer for the associated power supply.
	FirmwareVersion string
	// HotPluggable shall indicate whether the
	// device can be inserted or removed while the underlying equipment
	// otherwise remains in its current operational state. Devices indicated
	// as hot-pluggable shall allow the device to become operable without
	// altering the operational state of the underlying equipment. Devices
	// that cannot be inserted or removed from equipment in operation, or
	// devices that cannot become operable without affecting the operational
	// state of that equipment, shall be indicated as not hot-pluggable.
	HotPluggable bool
	// IndicatorLED shall contain the indicator
	// light state for the indicator light associated with this power supply.
	IndicatorLED common.IndicatorLED
	// InputRanges shall be a collection of ranges usable by the power supply unit.
	InputRanges []InputRange
	// LastPowerOutputWatts shall contain the average power
	// output, measured in Watts, of the associated power supply.
	LastPowerOutputWatts float32
	// LineInputVoltage shall contain the value in Volts of
	// the line input voltage (measured or configured for) that the power
	// supply has been configured to operate with or is currently receiving.
	LineInputVoltage float32
	// LineInputVoltageType shall contain the type of input
	// line voltage supported by the associated power supply.
	LineInputVoltageType LineInputVoltageType
	// Location shall contain location information of the
	// associated power supply.
	Location common.Location
	// Manufacturer shall be the name of the
	// organization responsible for producing the power supply. This
	// organization might be the entity from whom the power supply is
	// purchased, but this is not necessarily true.
	Manufacturer string
	// MemberID shall uniquely identify the
	// member within the collection. For services supporting Redfish v1.6 or
	// higher, this value shall be the zero-based array index.
	MemberID string `json:"MemberId"`
	// Model shall contain the model information as defined
	// by the manufacturer for the associated power supply.
	Model string
	// PartNumber shall contain the part number as defined
	// by the manufacturer for the associated power supply.
	PartNumber string
	// PowerCapacityWatts shall contain the maximum amount
	// of power, in Watts, that the associated power supply is rated to
	// deliver.
	PowerCapacityWatts float32
	// PowerInputWatts shall contain the value of the
	// measured input power, in Watts, of the associated power supply.
	PowerInputWatts float32
	// PowerOutputWatts shall contain the value of the
	// measured output power, in Watts, of the associated power supply.
	PowerOutputWatts float32
	// PowerSupplyType shall contain the input power type
	// (AC or DC) of the associated power supply.
	PowerSupplyType PowerSupplyType
	// Redundancy is used to show redundancy for power supplies and other
	// elements in this resource. The use of IDs within these arrays shall
	// reference the members of the redundancy groups.
	Redundancy []Redundancy
	// RedundancyCount is the number of objects.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// SerialNumber shall contain the serial number as
	// defined by the manufacturer for the associated power supply.
	SerialNumber string
	// SparePartNumber shall contain the spare or
	// replacement part number as defined by the manufacturer for the
	// associated power supply.
	SparePartNumber string
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PowerSupply object from the raw JSON.
func (powersupply *PowerSupply) UnmarshalJSON(b []byte) error {
	type temp PowerSupply
	var t struct {
		temp
		Assembly common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*powersupply = PowerSupply(t.temp)
	powersupply.assembly = string(t.Assembly)

	// This is a read/write object, so we need to save the raw object data for later
	powersupply.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (powersupply *PowerSupply) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(PowerSupply)
	err := original.UnmarshalJSON(powersupply.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"IndicatorLED",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(powersupply).Elem()

	return powersupply.Entity.Update(originalElement, currentElement, readWriteFields)
}

// Voltage is a voltage representation.
type Voltage struct {
	common.Entity

	// LowerThresholdCritical shall indicate
	// the present reading is below the normal range but is not yet fatal.
	// Units shall use the same units as the related ReadingVolts property.
	LowerThresholdCritical float32
	// LowerThresholdFatal shall indicate the
	// present reading is below the normal range and is fatal. Units shall
	// use the same units as the related ReadingVolts property.
	LowerThresholdFatal float32
	// LowerThresholdNonCritical shall indicate
	// the present reading is below the normal range but is not critical.
	// Units shall use the same units as the related ReadingVolts property.
	LowerThresholdNonCritical float32
	// MaxReadingRange shall indicate the
	// highest possible value for ReadingVolts. Units shall use the same
	// units as the related ReadingVolts property.
	MaxReadingRange float32
	// MemberID shall uniquely identify the member within the collection. For
	// services supporting Redfish v1.6 or higher, this value shall be the
	// zero-based array index.
	MemberID string `json:"MemberId"`
	// MinReadingRange shall indicate the lowest possible value for ReadingVolts.
	// Units shall use the same units as the related ReadingVolts property.
	MinReadingRange float32
	// PhysicalContext shall be a description
	// of the affected device or region within the chassis to which this
	// voltage measurement applies.
	PhysicalContext string
	// ReadingVolts shall be the present
	// reading of the voltage sensor's reading.
	ReadingVolts float32
	// SensorNumber shall be a numerical
	// identifier for this voltage sensor that is unique within this
	// resource.
	SensorNumber int
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// UpperThresholdCritical shall indicate
	// the present reading is above the normal range but is not yet fatal.
	// Units shall use the same units as the related ReadingVolts property.
	UpperThresholdCritical float32
	// UpperThresholdFatal shall indicate the
	// present reading is above the normal range and is fatal. Units shall
	// use the same units as the related ReadingVolts property.
	UpperThresholdFatal float32
	// UpperThresholdNonCritical shall indicate
	// the present reading is above the normal range but is not critical.
	// Units shall use the same units as the related ReadingVolts property.
	UpperThresholdNonCritical float32
}

// UnmarshalJSON unmarshals a Voltage object from the raw JSON.
func (voltage *Voltage) UnmarshalJSON(b []byte) error { // nolint:dupl
	type temp Voltage
	type t1 struct {
		temp
	}
	var t t1

	err := json.Unmarshal(b, &t)
	if err != nil {
		// See if we need to handle converting MemberID
		var t2 struct {
			t1
			MemberID int `json:"MemberId"`
		}
		err2 := json.Unmarshal(b, &t2)

		if err2 != nil {
			// Return the original error
			return err
		}

		// Convert the numeric member ID to a string
		t = t2.t1
		t.temp.MemberID = strconv.Itoa(t2.MemberID)
	}

	// Extract the links to other entities for later
	*voltage = Voltage(t.temp)

	return nil
}
