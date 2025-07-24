//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/stmcginnis/gofish/common"
)

// InputType is the type of power input.
type InputType string

const (
	// ACInputType indicates an Alternating Current (AC) input range.
	ACInputType InputType = "AC"
	// DCInputType indicates a Direct Current (DC) input range.
	DCInputType InputType = "DC"
)

// LineInputVoltageType is the type of line voltage supported by a power supply.
type LineInputVoltageType string

const (
	// UnknownLineInputVoltageType indicates the power supply line input voltage type cannot be determined.
	UnknownLineInputVoltageType LineInputVoltageType = "Unknown"

	// ACLowLineLineInputVoltageType indicates 100-127V AC input. Deprecated in v1.1.0 in favor of AC120VLineInputVoltageType.
	ACLowLineLineInputVoltageType LineInputVoltageType = "ACLowLine"
	// ACMidLineLineInputVoltageType indicates 200-240V AC input. Deprecated in v1.1.0 in favor of AC240VLineInputVoltageType.
	ACMidLineLineInputVoltageType LineInputVoltageType = "ACMidLine"
	// ACHighLineLineInputVoltageType indicates 277V AC input. Deprecated in v1.1.0 in favor of AC277VLineInputVoltageType.
	ACHighLineLineInputVoltageType LineInputVoltageType = "ACHighLine"

	// DCNeg48VLineInputVoltageType indicates -48V DC input.
	DCNeg48VLineInputVoltageType LineInputVoltageType = "DCNeg48V"
	// DC380VLineInputVoltageType indicates high-voltage DC input (380V).
	DC380VLineInputVoltageType LineInputVoltageType = "DC380V"

	// AC120VLineInputVoltageType indicates AC 120V nominal input. Added in v1.1.0.
	AC120VLineInputVoltageType LineInputVoltageType = "AC120V"
	// AC240VLineInputVoltageType indicates AC 240V nominal input. Added in v1.1.0.
	AC240VLineInputVoltageType LineInputVoltageType = "AC240V"
	// AC277VLineInputVoltageType indicates AC 277V nominal input. Added in v1.1.0.
	AC277VLineInputVoltageType LineInputVoltageType = "AC277V"
	// ACandDCWideRangeLineInputVoltageType indicates wide range AC or DC input. Added in v1.1.0.
	ACandDCWideRangeLineInputVoltageType LineInputVoltageType = "ACandDCWideRange"
	// ACWideRangeLineInputVoltageType indicates wide range AC input. Added in v1.1.0.
	ACWideRangeLineInputVoltageType LineInputVoltageType = "ACWideRange"
	// DC240VLineInputVoltageType indicates DC 240V nominal input. Added in v1.1.0.
	DC240VLineInputVoltageType LineInputVoltageType = "DC240V"
)

// PowerLimitException is the action taken when power cannot be maintained below the limit.
type PowerLimitException string

const (
	// NoActionPowerLimitException indicates no action is taken when the limit is exceeded.
	NoActionPowerLimitException PowerLimitException = "NoAction"
	// HardPowerOffPowerLimitException indicates power is turned off immediately when the limit is exceeded.
	HardPowerOffPowerLimitException PowerLimitException = "HardPowerOff"
	// LogEventOnlyPowerLimitException indicates an event is logged when the limit is exceeded, but no further action is taken.
	LogEventOnlyPowerLimitException PowerLimitException = "LogEventOnly"
	// OemPowerLimitException indicates an OEM-defined action is taken.
	OemPowerLimitException PowerLimitException = "Oem"
)

// PowerSupplyType is the type of power supply.
type PowerSupplyType string

const (
	// UnknownPowerSupplyType indicates the power supply type cannot be determined.
	UnknownPowerSupplyType PowerSupplyType = "Unknown"
	// ACPowerSupplyType indicates an Alternating Current (AC) power supply.
	ACPowerSupplyType PowerSupplyType = "AC"
	// DCPowerSupplyType indicates a Direct Current (DC) power supply.
	DCPowerSupplyType PowerSupplyType = "DC"
	// ACorDCPowerSupplyType indicates a power supply that supports both DC and AC input.
	ACorDCPowerSupplyType PowerSupplyType = "ACorDC"
)

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

// Power represents power metrics for a Redfish implementation.
type Power struct {
	common.Entity
	// ODataContext is the OData context URL.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the OData ETag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the OData type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description *string `json:"Description"`
	// Oem contains OEM-specific extensions.
	OEM json.RawMessage `json:"Oem"`
	// Actions contains the available actions for this resource.
	Actions struct {
		PowerSupplyReset common.ActionTarget `json:"#Power.PowerSupplyReset"`
		OEM              json.RawMessage     `json:"Oem"`
	} `json:"Actions"`
	// PowerControl contains the set of power control readings and settings.
	PowerControl []PowerControl `json:"PowerControl"`
	// PowerControlCount is the number of power control items.
	PowerControlCount int `json:"PowerControl@odata.count"`
	// PowerSupplies contains the set of power supplies associated with this system or device.
	PowerSupplies []PowerSupply `json:"PowerSupplies"`
	// PowerSuppliesCount is the number of power supply items.
	PowerSuppliesCount int `json:"PowerSupplies@odata.count"`
	// Redundancy contains redundancy information for the set of power supplies.
	Redundancy []Redundancy `json:"Redundancy"`
	// RedundancyCount is the number of redundancy items.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Voltages contains the set of voltage sensors for this chassis.
	Voltages []Voltage `json:"Voltages"`
	// VoltagesCount is the number of voltage items.
	VoltagesCount int `json:"Voltages@odata.count"`

	powerSupplyResetTarget string
}

// UnmarshalJSON unmarshals a Power object from the raw JSON.
func (power *Power) UnmarshalJSON(b []byte) error {
	type pwr Power
	var t struct {
		pwr
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*power = Power(t.pwr)
	power.powerSupplyResetTarget = t.Actions.PowerSupplyReset.Target
	return nil
}

// PowerSupplyReset resets the specified power supply.
// memberID identifies the power supply in the PowerSupplies array.
// resetType specifies the type of reset to perform.
func (power *Power) PowerSupplyReset(memberID string, resetType ResetType) error {
	t := struct {
		MemberID  string    `json:"MemberId"`
		ResetType ResetType `json:"ResetType"`
	}{
		MemberID:  memberID,
		ResetType: resetType,
	}
	return power.Post(power.powerSupplyResetTarget, t)
}

// RedundancySet returns the power supplies in the specified redundancy group.
// memberID identifies the redundancy group in the Redundancy array.
func (power *Power) RedundancySet(memberID int) []PowerSupply {
	var powerSupplies []PowerSupply
	if len(power.Redundancy) >= memberID+1 {
		for _, psLink := range power.Redundancy[memberID].redundancySet {
			for i := range power.PowerSupplies {
				if power.PowerSupplies[i].ODataID == psLink {
					powerSupplies = append(powerSupplies, power.PowerSupplies[i])
				}
			}
		}
	}
	return powerSupplies
}

// GetPower retrieves a Power instance from the service.
func GetPower(c common.Client, uri string) (*Power, error) {
	return common.GetObject[Power](c, uri)
}

// ListReferencedPowers retrieves a collection of Power from a reference.
func ListReferencedPowers(c common.Client, link string) ([]*Power, error) {
	return common.GetCollectionObjects[Power](c, link)
}

// PowerControl represents power control functions for a chassis or system.
type PowerControl struct {
	common.Entity

	// ODataID is the OData identifier.
	ODataID string `json:"@odata.id"`
	// MemberID uniquely identifies this member within the collection.
	MemberID string `json:"MemberId"`
	// Name is the name of this power control function.
	Name string `json:"Name,omitempty"`
	// PhysicalContext describes the affected device(s) or region within the chassis.
	PhysicalContext common.PhysicalContext `json:"PhysicalContext,omitempty"`
	// PowerAllocatedWatts is the total power currently allocated to chassis resources.
	PowerAllocatedWatts *float32 `json:"PowerAllocatedWatts,omitempty"`
	// PowerAvailableWatts is the available power capacity (PowerCapacityWatts - PowerAllocatedWatts).
	PowerAvailableWatts *float32 `json:"PowerAvailableWatts,omitempty"`
	// PowerCapacityWatts is the total power capacity available for allocation.
	PowerCapacityWatts *float32 `json:"PowerCapacityWatts,omitempty"`
	// PowerConsumedWatts is the actual power being consumed by the chassis.
	PowerConsumedWatts *float32 `json:"PowerConsumedWatts,omitempty"`
	// PowerLimit contains power limit status and configuration information.
	PowerLimit *PowerLimit `json:"PowerLimit,omitempty"`
	// PowerMetrics contains power metrics including interval, min/max/avg consumption.
	PowerMetrics *PowerMetric `json:"PowerMetrics,omitempty"`
	// PowerRequestedWatts is the power currently requested for future use.
	PowerRequestedWatts *float32 `json:"PowerRequestedWatts,omitempty"`
	// RelatedItem contains links to resources associated with this power limit.
	RelatedItem []common.Link `json:"RelatedItem,omitempty"`
	// Status contains status and health properties of this resource.
	Status common.Status `json:"Status,omitempty"`
	// Actions contains the available actions for this resource.
	Actions struct {
		Oem json.RawMessage `json:"Oem,omitempty"`
	} `json:"Actions,omitempty"`
	// Oem contains OEM-specific extensions.
	OEM json.RawMessage `json:"Oem,omitempty"`
}

// UnmarshalJSON unmarshals a PowerControl object from the raw JSON.
func (powercontrol *PowerControl) UnmarshalJSON(b []byte) error {
	if powercontrol == nil {
		return fmt.Errorf("nil PowerControl receiver")
	}
	if len(b) == 0 {
		return fmt.Errorf("empty input data")
	}

	type pc PowerControl
	type t1 struct {
		pc
	}
	var t t1

	// First try normal unmarshaling where MemberID is a string
	err := json.Unmarshal(b, &t)
	if err == nil {
		*powercontrol = PowerControl(t.pc)
		return nil
	}

	// If first attempt failed, try to handle MemberID as any type
	var raw map[string]any
	if err := json.Unmarshal(b, &raw); err != nil {
		return fmt.Errorf("failed to unmarshal PowerControl: %v", err)
	}

	// Create a new pc struct and copy all fields except MemberId
	var result pc
	for k, v := range raw {
		if k != "MemberId" {
			// This is simplified - you might need more sophisticated field copying
			// depending on your actual struct fields
			continue
		}

		// Handle MemberId conversion
		switch v := v.(type) {
		case string:
			result.MemberID = v
		case float64: // JSON numbers are float64
			result.MemberID = strconv.Itoa(int(v))
		case int:
			result.MemberID = strconv.Itoa(v)
		case bool:
			result.MemberID = strconv.FormatBool(v)
		default:
			// For any other type, convert to string via fmt.Sprint
			result.MemberID = fmt.Sprint(v)
		}
	}

	*powercontrol = PowerControl(result)
	return nil
}

// PowerLimit contains power limit status and configuration information.
type PowerLimit struct {
	// CorrectionInMs is the time required to reduce power consumption below the limit.
	CorrectionInMs *int64 `json:"CorrectionInMs,omitempty"`
	// LimitException is the action taken if power cannot be maintained below the limit.
	LimitException PowerLimitException `json:"LimitException,omitempty"`
	// LimitInWatts is the power cap limit. If null, power capping is disabled.
	LimitInWatts *float32 `json:"LimitInWatts,omitempty"`
}

// PowerMetric contains power metrics including interval, min/max/avg consumption.
type PowerMetric struct {
	// AverageConsumedWatts is the average power over the last IntervalInMin minutes.
	AverageConsumedWatts *float32 `json:"AverageConsumedWatts,omitempty"`
	// IntervalInMin is the measurement window in minutes.
	IntervalInMin *int `json:"IntervalInMin,omitempty"`
	// MaxConsumedWatts is the maximum power within the measurement window.
	MaxConsumedWatts *float32 `json:"MaxConsumedWatts,omitempty"`
	// MinConsumedWatts is the minimum power within the measurement window.
	MinConsumedWatts *float32 `json:"MinConsumedWatts,omitempty"`
}

// PowerSupply represents a power supply associated with a system or device.
type PowerSupply struct {
	common.Entity

	// ODataID is the OData identifier.
	ODataID string `json:"@odata.id"`
	// Assembly is a link to the associated Assembly resource.
	Assembly string `json:"Assembly,omitempty"`
	// EfficiencyPercent is the measured power efficiency percentage.
	EfficiencyPercent *float32 `json:"EfficiencyPercent,omitempty"`
	// FirmwareVersion is the firmware version of this power supply.
	FirmwareVersion string `json:"FirmwareVersion,omitempty"`
	// HotPluggable indicates if this device can be inserted/removed while operating.
	HotPluggable *bool `json:"HotPluggable,omitempty"`
	// IndicatorLED is the state of the indicator LED for this power supply.
	IndicatorLED common.IndicatorLED `json:"IndicatorLED,omitempty"`
	// InputRanges contains the input ranges this power supply can use.
	InputRanges []InputRange `json:"InputRanges,omitempty"`
	// LastPowerOutputWatts is the average power output in Watts.
	LastPowerOutputWatts *float32 `json:"LastPowerOutputWatts,omitempty"`
	// LineInputVoltage is the current line input voltage in Volts.
	LineInputVoltage *float32 `json:"LineInputVoltage,omitempty"`
	// LineInputVoltageType is the type of input line voltage.
	LineInputVoltageType LineInputVoltageType `json:"LineInputVoltageType,omitempty"`
	// Location contains location information for this power supply.
	Location common.Location `json:"Location,omitempty"`
	// Manufacturer is the name of the manufacturer.
	Manufacturer string `json:"Manufacturer,omitempty"`
	// MemberID uniquely identifies this member within the collection.
	MemberID string `json:"MemberId"`
	// Metrics is a link to the power supply metrics resource.
	Metrics string `json:"Metrics,omitempty"`
	// Model is the model information for this power supply.
	Model string `json:"Model,omitempty"`
	// Name is the name of this power supply.
	Name string `json:"Name,omitempty"`
	// Oem contains OEM-specific extensions.
	OEM json.RawMessage `json:"Oem,omitempty"`
	// PartNumber is the part number for this power supply.
	PartNumber string `json:"PartNumber,omitempty"`
	// PowerCapacityWatts is the maximum power capacity in Watts.
	PowerCapacityWatts *float32 `json:"PowerCapacityWatts,omitempty"`
	// PowerInputWatts is the measured input power in Watts.
	PowerInputWatts *float32 `json:"PowerInputWatts,omitempty"`
	// PowerOutputWatts is the measured output power in Watts.
	PowerOutputWatts *float32 `json:"PowerOutputWatts,omitempty"`
	// PowerSupplyType is the input power type (AC or DC).
	PowerSupplyType PowerSupplyType `json:"PowerSupplyType,omitempty"`
	// Redundancy contains redundancy information for this power supply.
	Redundancy []common.Link `json:"Redundancy,omitempty"`
	// RedundancyCount is the number of redundancy items.
	RedundancyCount int `json:"Redundancy@odata.count,omitempty"`
	// RelatedItem contains links to resources associated with this power supply.
	RelatedItem []common.Link `json:"RelatedItem,omitempty"`
	// RelatedItemCount is the number of related items.
	RelatedItemCount int `json:"RelatedItem@odata.count,omitempty"`
	// SerialNumber is the serial number for this power supply.
	SerialNumber string `json:"SerialNumber,omitempty"`
	// SparePartNumber is the spare part number for this power supply.
	SparePartNumber string `json:"SparePartNumber,omitempty"`
	// Status contains status and health properties of this resource.
	Status common.Status `json:"Status,omitempty"`
	// Actions contains the available actions for this resource.
	Actions struct {
		OEM json.RawMessage `json:"Oem,omitempty"`
	} `json:"Actions,omitempty"`

	rawData          []byte
	redundancyLinks  []string
	relateditemLinks []string
}

// UnmarshalJSON unmarshals a PowerSupply object from the raw JSON.
func (powersupply *PowerSupply) UnmarshalJSON(b []byte) error {
	type ps PowerSupply
	var t struct {
		ps
		Assembly    common.Link  `json:"Assembly"`
		Metrics     common.Link  `json:"Metrics"`
		Redundancy  common.Links `json:"Redundancy"`
		RelatedItem common.Links `json:"RelatedItem"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*powersupply = PowerSupply(t.ps)
	powersupply.Assembly = t.Assembly.String()
	powersupply.Metrics = t.Metrics.String()
	powersupply.redundancyLinks = t.Redundancy.ToStrings()
	powersupply.relateditemLinks = t.RelatedItem.ToStrings()
	powersupply.rawData = b

	return nil
}

// GetPowerSupply retrieves a PowerSupply instance from the service.
func GetPowerSupply(c common.Client, uri string) (*PowerSupply, error) {
	var powerSupply PowerSupply
	return &powerSupply, powerSupply.Get(c, uri, &powerSupply)
}

// ListReferencedPowerSupplies retrieves a collection of PowerSupplies from a reference.
func ListReferencedPowerSupplies(c common.Client, link string) ([]*PowerSupply, error) {
	return common.GetCollectionObjects[PowerSupply](c, link)
}

// Update commits updates to this object's properties to the running system.
func (powersupply *PowerSupply) Update() error {
	readWriteFields := []string{"IndicatorLED"}

	return powersupply.UpdateFromRawData(powersupply, powersupply.rawData, readWriteFields)
}

// GetAssembly retrieves the containing Assembly.
func (powersupply *PowerSupply) GetAssembly() (*Assembly, error) {
	if powersupply.Assembly == "" {
		return nil, nil
	}
	return GetAssembly(powersupply.GetClient(), powersupply.Assembly)
}

// GetMetrics retrieves the metrics associated with this power supply.
func (powersupply *PowerSupply) GetMetrics() (*PowerSupplyUnitMetrics, error) {
	if powersupply.Metrics == "" {
		return nil, nil
	}
	return GetPowerSupplyUnitMetrics(powersupply.GetClient(), powersupply.Metrics)
}

// GetRedundancy retrieves the redundancy groups this power supply belongs to.
func (powersupply *PowerSupply) GetRedundancy() ([]*Redundancy, error) {
	return common.GetObjects[Redundancy](powersupply.GetClient(), powersupply.redundancyLinks)
}

// Voltage represents a voltage sensor for a chassis.
type Voltage struct {
	common.Entity

	// ODataID is the OData identifier.
	ODataID string `json:"@odata.id"`
	// Actions contains the available actions for this resource.
	Actions *VoltageActions `json:"Actions,omitempty"`
	// LowerThresholdCritical indicates the reading is below normal range but not yet fatal.
	LowerThresholdCritical *float32 `json:"LowerThresholdCritical,omitempty"`
	// LowerThresholdFatal indicates the reading is below normal range and fatal.
	LowerThresholdFatal *float32 `json:"LowerThresholdFatal,omitempty"`
	// LowerThresholdNonCritical indicates the reading is below normal range but not critical.
	LowerThresholdNonCritical *float32 `json:"LowerThresholdNonCritical,omitempty"`
	// MaxReadingRange indicates the highest possible value for ReadingVolts.
	MaxReadingRange *float32 `json:"MaxReadingRange,omitempty"`
	// MemberID uniquely identifies this member within the collection.
	MemberID string `json:"MemberId"`
	// MinReadingRange indicates the lowest possible value for ReadingVolts.
	MinReadingRange *float32 `json:"MinReadingRange,omitempty"`
	// Name is the name of this voltage sensor.
	Name string `json:"Name,omitempty"`
	// Oem contains OEM-specific extensions.
	OEM json.RawMessage `json:"Oem,omitempty"`
	// PhysicalContext describes the affected device or region within the chassis.
	PhysicalContext common.PhysicalContext `json:"PhysicalContext,omitempty"`
	// ReadingVolts is the present voltage reading.
	ReadingVolts *float32 `json:"ReadingVolts,omitempty"`
	// RelatedItem contains links to resources associated with this voltage measurement.
	RelatedItem []common.Link `json:"RelatedItem,omitempty"`
	// SensorNumber is a numerical identifier unique within this resource.
	SensorNumber *int `json:"SensorNumber,omitempty"`
	// Status contains status and health properties of this resource.
	Status common.Status `json:"Status,omitempty"`
	// UpperThresholdCritical indicates the reading is above normal range but not yet fatal.
	UpperThresholdCritical *float32 `json:"UpperThresholdCritical,omitempty"`
	// UpperThresholdFatal indicates the reading is above normal range and fatal.
	UpperThresholdFatal *float32 `json:"UpperThresholdFatal,omitempty"`
	// UpperThresholdNonCritical indicates the reading is above normal range but not critical.
	UpperThresholdNonCritical *float32 `json:"UpperThresholdNonCritical,omitempty"`
}

// VoltageActions contains the available actions for a Voltage resource.
type VoltageActions struct {
	// Oem contains OEM-specific actions.
	OEM json.RawMessage `json:"Oem,omitempty"`
}

// UnmarshalJSON unmarshals a Voltage object from the raw JSON.
func (voltage *Voltage) UnmarshalJSON(b []byte) error {
	type vlg Voltage
	var t struct {
		vlg
		RelatedItemCount int `json:"RelatedItem@odata.count"`
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		var t2 struct {
			vlg
			RelatedItemCount int `json:"RelatedItem@odata.count"`
			MemberID         int `json:"MemberId"`
		}
		err2 := json.Unmarshal(b, &t2)
		if err2 != nil {
			return err
		}

		t = struct {
			vlg
			RelatedItemCount int `json:"RelatedItem@odata.count"`
		}{
			vlg:              t2.vlg,
			RelatedItemCount: t2.RelatedItemCount,
		}
		t.MemberID = strconv.Itoa(t2.MemberID)
	}

	*voltage = Voltage(t.vlg)
	return nil
}
