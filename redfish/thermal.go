//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// ReadingUnits is the type of units used for a reading.
type ReadingUnits string

const (
	// RPMReadingUnits indicates that the fan reading and thresholds are measured in rotations per minute.
	RPMReadingUnits ReadingUnits = "RPM"
	// PercentReadingUnits indicates that the fan reading and thresholds are measured in percentage.
	PercentReadingUnits ReadingUnits = "Percent"
)

// ThermalFan represents a fan in a Redfish system.
type ThermalFan struct {
	common.Entity
	// HotPluggable indicates if the device can be inserted/removed while equipment is operating.
	HotPluggable *bool `json:"HotPluggable,omitempty"`
	// IndicatorLED contains the state of the indicator light associated with this fan.
	IndicatorLED *common.IndicatorLED `json:"IndicatorLED,omitempty"`
	// Location contains information about the fan's physical location.
	Location *common.Location `json:"Location,omitempty"`
	// LowerThresholdCritical contains the value below which reading is outside normal range but not yet fatal.
	LowerThresholdCritical *int `json:"LowerThresholdCritical,omitempty"`
	// LowerThresholdFatal contains the value below which reading is outside normal range and is fatal.
	LowerThresholdFatal *int `json:"LowerThresholdFatal,omitempty"`
	// LowerThresholdNonCritical contains the value below which reading is outside normal range.
	LowerThresholdNonCritical *int `json:"LowerThresholdNonCritical,omitempty"`
	// Manufacturer contains the name of the organization that produced the fan.
	Manufacturer string `json:"Manufacturer,omitempty"`
	// MaxReadingRange indicates the highest possible value for the Reading property.
	MaxReadingRange *int `json:"MaxReadingRange,omitempty"`
	// MemberID contains the unique identifier for this member within an array.
	MemberID string `json:"MemberId"`
	// MinReadingRange indicates the lowest possible value for the Reading property.
	MinReadingRange *int `json:"MinReadingRange,omitempty"`
	// Model contains the manufacturer-defined model information for the fan.
	Model string `json:"Model,omitempty"`
	// Name contains the name of the fan.
	Name string `json:"Name,omitempty"`
	// OEM contains OEM extensions (vendor-specific properties).
	OEM json.RawMessage `json:"Oem,omitempty"`
	// PartNumber contains the manufacturer-defined part number for the fan.
	PartNumber string `json:"PartNumber,omitempty"`
	// PhysicalContext describes the device or region within the chassis associated with this fan.
	PhysicalContext *PhysicalContext `json:"PhysicalContext,omitempty"`
	// Reading contains the current fan sensor reading.
	Reading *int `json:"Reading,omitempty"`
	// ReadingUnits contains the units in which fan readings and thresholds are measured.
	ReadingUnits ReadingUnits `json:"ReadingUnits,omitempty"`
	// SensorNumber contains a numerical identifier for this fan speed sensor.
	SensorNumber *int `json:"SensorNumber,omitempty"`
	// SerialNumber contains the manufacturer-defined serial number for the fan.
	SerialNumber string `json:"SerialNumber,omitempty"`
	// SparePartNumber contains the manufacturer-defined spare/replacement part number.
	SparePartNumber string `json:"SparePartNumber,omitempty"`
	// Status contains status and health properties of the resource.
	Status *common.Status `json:"Status,omitempty"`
	// UpperThresholdCritical contains the value above which reading is outside normal range but not yet fatal.
	UpperThresholdCritical *int `json:"UpperThresholdCritical,omitempty"`
	// UpperThresholdFatal contains the value above which reading is outside normal range and is fatal.
	UpperThresholdFatal *int `json:"UpperThresholdFatal,omitempty"`
	// UpperThresholdNonCritical contains the value above which reading is outside normal range.
	UpperThresholdNonCritical *int `json:"UpperThresholdNonCritical,omitempty"`

	// Links to related objects (unmarshaled separately)
	// Assembly contains a link to a resource of type Assembly.
	assembly    string
	redundancy  []string
	relatedItem []string

	// Counters for related objects
	RedundancyCount  int `json:"Redundancy@odata.count,omitempty"`
	RelatedItemCount int `json:"RelatedItem@odata.count,omitempty"`

	// rawData holds original JSON for comparison during updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ThermalFan object from the raw JSON.
func (fan *ThermalFan) UnmarshalJSON(b []byte) error {
	type temp ThermalFan
	var t struct {
		temp
		Assembly    common.Link
		Redundancy  common.Links
		RelatedItem common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fan = ThermalFan(t.temp)

	// Extract the links to other entities for later
	fan.assembly = t.Assembly.String()
	fan.redundancy = t.Redundancy.ToStrings()
	fan.relatedItem = t.RelatedItem.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	fan.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (fan *ThermalFan) Update() error {
	// Get a representation of the object's original state so we can find what to update.
	original := new(ThermalFan)
	original.UnmarshalJSON(fan.rawData)

	readWriteFields := []string{
		"IndicatorLED",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(fan).Elem()

	return fan.Entity.Update(originalElement, currentElement, readWriteFields)
}

// Temperature represents a temperature sensor in a Redfish system.
type Temperature struct {
	common.Entity
	// AdjustedMaxAllowableOperatingValue indicates the adjusted maximum allowable operating temperature.
	AdjustedMaxAllowableOperatingValue *float32 `json:"AdjustedMaxAllowableOperatingValue,omitempty"`
	// AdjustedMinAllowableOperatingValue indicates the adjusted minimum allowable operating temperature.
	AdjustedMinAllowableOperatingValue *float32 `json:"AdjustedMinAllowableOperatingValue,omitempty"`
	// DeltaPhysicalContext describes the device/region to which DeltaReadingCelsius measurement applies.
	DeltaPhysicalContext *PhysicalContext `json:"DeltaPhysicalContext,omitempty"`
	// DeltaReadingCelsius contains the temperature difference between this sensor and DeltaPhysicalContext.
	DeltaReadingCelsius *float32 `json:"DeltaReadingCelsius,omitempty"`
	// LowerThresholdCritical indicates when reading is below normal range but not yet fatal.
	LowerThresholdCritical *float32 `json:"LowerThresholdCritical,omitempty"`
	// LowerThresholdFatal indicates when reading is below normal range and is fatal.
	LowerThresholdFatal *float32 `json:"LowerThresholdFatal,omitempty"`
	// LowerThresholdNonCritical indicates when reading is below normal range.
	LowerThresholdNonCritical *float32 `json:"LowerThresholdNonCritical,omitempty"`
	// LowerThresholdUser contains user-defined lower threshold value.
	LowerThresholdUser *float32 `json:"LowerThresholdUser,omitempty"`
	// MaxAllowableOperatingValue indicates the maximum allowable operating temperature.
	MaxAllowableOperatingValue *float32 `json:"MaxAllowableOperatingValue,omitempty"`
	// MaxReadingRangeTemp indicates the highest possible temperature reading.
	MaxReadingRangeTemp *float32 `json:"MaxReadingRangeTemp,omitempty"`
	// MemberID uniquely identifies this member within the collection.
	MemberID string `json:"MemberId"`
	// MinAllowableOperatingValue indicates the minimum allowable operating temperature.
	MinAllowableOperatingValue *float32 `json:"MinAllowableOperatingValue,omitempty"`
	// MinReadingRangeTemp indicates the lowest possible temperature reading.
	MinReadingRangeTemp *float32 `json:"MinReadingRangeTemp,omitempty"`
	// Name contains the name of the temperature sensor.
	Name string `json:"Name,omitempty"`
	// OEM contains OEM extensions (vendor-specific properties).
	OEM json.RawMessage `json:"Oem,omitempty"`
	// PhysicalContext describes the device or region to which this temperature applies.
	PhysicalContext *PhysicalContext `json:"PhysicalContext,omitempty"`
	// ReadingCelsius contains the current temperature reading in Celsius.
	ReadingCelsius *float32 `json:"ReadingCelsius,omitempty"`
	// SensorNumber contains a numerical identifier for this temperature sensor.
	SensorNumber *int `json:"SensorNumber,omitempty"`
	// Status contains status and health properties of the resource.
	Status *common.Status `json:"Status,omitempty"`
	// UpperThresholdCritical indicates when reading is above normal range but not yet fatal.
	UpperThresholdCritical *float32 `json:"UpperThresholdCritical,omitempty"`
	// UpperThresholdFatal indicates when reading is above normal range and is fatal.
	UpperThresholdFatal *float32 `json:"UpperThresholdFatal,omitempty"`
	// UpperThresholdNonCritical indicates when reading is above normal range.
	UpperThresholdNonCritical *float32 `json:"UpperThresholdNonCritical,omitempty"`
	// UpperThresholdUser contains user-defined upper threshold value.
	UpperThresholdUser *float32 `json:"UpperThresholdUser,omitempty"`

	// Links to related items (unmarshaled separately)
	relatedItem []string

	// Counter for related items
	RelatedItemCount int `json:"RelatedItem@odata.count,omitempty"`

	// rawData holds original JSON for comparison during updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Temperature object from the raw JSON.
func (temperature *Temperature) UnmarshalJSON(b []byte) error {
	type temp Temperature
	var t struct {
		temp
		RelatedItem common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*temperature = Temperature(t.temp)

	// Extract the links to other entities for later
	temperature.relatedItem = t.RelatedItem.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	temperature.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (temperature *Temperature) Update() error {
	// Get a representation of the object's original state so we can find what to update.
	original := new(Temperature)
	original.UnmarshalJSON(temperature.rawData)

	readWriteFields := []string{
		"LowerThresholdUser",
		"UpperThresholdUser",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(temperature).Elem()

	return temperature.Entity.Update(originalElement, currentElement, readWriteFields)
}

// Thermal represents thermal management data in a Redfish system.
type Thermal struct {
	common.Entity
	// ODataContext contains the OData context.
	ODataContext string `json:"@odata.context,omitempty"`
	// ODataType contains the OData type.
	ODataType string `json:"@odata.type,omitempty"`
	// Description provides a description of this resource.
	Description string `json:"Description,omitempty"`
	// Fans contains the set of fans in this chassis.
	Fans []ThermalFan `json:"Fans,omitempty"`
	// OEM contains OEM extensions (vendor-specific properties).
	OEM json.RawMessage `json:"Oem,omitempty"`
	// Temperatures contains the set of temperature sensors in this chassis.
	Temperatures []Temperature `json:"Temperatures,omitempty"`

	// Redundancy contains redundancy information for fans and other elements.
	redundancy []string

	// Counters for collections
	FansCount         int `json:"Fans@odata.count,omitempty"`
	TemperaturesCount int `json:"Temperatures@odata.count,omitempty"`
	RedundancyCount   int `json:"Redundancy@odata.count,omitempty"`

	// rawData holds original JSON for comparison during updates.
	rawData []byte
}

// UnmarshalJSON unmarshals an object from the raw JSON.
func (thermal *Thermal) UnmarshalJSON(b []byte) error {
	type temp Thermal
	var t struct {
		temp
		Redundancy common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thermal = Thermal(t.temp)
	thermal.redundancy = t.Redundancy.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	thermal.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (thermal *Thermal) Update() error {
	// Get a representation of the object's original state so we can find what to update.
	original := new(Thermal)
	original.UnmarshalJSON(thermal.rawData)

	readWriteFields := []string{
		"Fans",
		"Temperatures",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(thermal).Elem()

	return thermal.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetThermal will get a Thermal instance from the service.
func GetThermal(c common.Client, uri string) (*Thermal, error) {
	return common.GetObject[Thermal](c, uri)
}

// ListReferencedThermals gets the collection of Thermal from a provided reference.
func ListReferencedThermals(c common.Client, link string) ([]*Thermal, error) {
	return common.GetCollectionObjects[Thermal](c, link)
}
