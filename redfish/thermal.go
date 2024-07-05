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

	// RPMReadingUnits Indicates that the fan reading and thresholds are
	// measured in rotations per minute.
	RPMReadingUnits ReadingUnits = "RPM"
	// PercentReadingUnits Indicates that the fan reading and thresholds are
	// measured in percentage.
	PercentReadingUnits ReadingUnits = "Percent"
)

type ThermalFan struct {
	common.Entity
	// Assembly shall contain a link to a resource of type Assembly.
	assembly string
	// HotPluggable shall indicate whether the device can be inserted or removed while the underlying equipment
	// otherwise remains in its current operational state. Hot-pluggable devices can become operable without altering
	// the operational state of the underlying equipment. Devices that cannot be inserted or removed from equipment in
	// operation, or devices that cannot become operable without affecting the operational state of that equipment,
	// shall not be hot-pluggable.
	HotPluggable bool
	// IndicatorLED shall contain the state of the indicator light associated with this fan.
	IndicatorLED common.IndicatorLED
	// Location shall contain the location information of the associated fan.
	Location common.Location
	// LowerThresholdCritical shall contain the value at which the Reading property is below the normal range but is
	// not yet fatal. The value of the property shall use the same units as the Reading property.
	LowerThresholdCritical int
	// LowerThresholdFatal shall contain the value at which the Reading property is below the normal range and is
	// fatal. The value of the property shall use the same units as the Reading property.
	LowerThresholdFatal int
	// LowerThresholdNonCritical shall contain the value at which the Reading property is below normal range. The value
	// of the property shall use the same units as the Reading property.
	LowerThresholdNonCritical int
	// Manufacturer shall contain the name of the organization responsible for producing the fan. This organization may
	// be the entity from whom the fan is purchased, but this is not necessarily true.
	Manufacturer string
	// MaxReadingRange shall indicate the highest possible value for the Reading property. The value of the property
	// shall use the same units as the Reading property.
	MaxReadingRange int
	// MemberId shall contain the unique identifier for this member within an array. For services supporting Redfish
	// v1.6 or higher, this value shall contain the zero-based array index.
	MemberID string
	// MinReadingRange shall indicate the lowest possible value for the Reading property. The value of the property
	// shall use the same units as the Reading property.
	MinReadingRange int
	// Model shall contain the model information as defined by the manufacturer for the associated fan.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for the associated fan.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region within the chassis with which this
	// fan is associated.
	PhysicalContext PhysicalContext
	// Reading shall contain the fan sensor reading.
	Reading int
	// ReadingUnits shall contain the units in which the fan reading and thresholds are measured.
	ReadingUnits ReadingUnits
	// Redundancy shall contain an array of links to the redundancy groups to which this fan belongs.
	redundancy []string
	// RedundancyCount is the number of Redundancy items.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// RelatedItem shall contain an array of links to resources or objects that this fan services.
	relatedItem []string
	// RelatedItem@odataCount is the number of related items.
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SensorNumber shall contain a numerical identifier for this fan speed sensor that is unique within this resource.
	SensorNumber int
	// SerialNumber shall contain the serial number as defined by the manufacturer for the associated fan.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as defined by the manufacturer for the
	// associated fan.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UpperThresholdCritical shall contain the value at which the Reading property is above the normal range but is
	// not yet fatal. The value of the property shall use the same units as the Reading property.
	UpperThresholdCritical int
	// UpperThresholdFatal shall contain the value at which the Reading property is above the normal range and is
	// fatal. The value of the property shall use the same units as the Reading property.
	UpperThresholdFatal int
	// UpperThresholdNonCritical shall contain the value at which the Reading property is above the normal range. The
	// value of the property shall use the same units as the Reading property.
	UpperThresholdNonCritical int
	// rawData holds the original serialized JSON so we can compare updates.
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
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(ThermalFan)
	original.UnmarshalJSON(fan.rawData)

	readWriteFields := []string{
		"IndicatorLED",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(fan).Elem()

	return fan.Entity.Update(originalElement, currentElement, readWriteFields)
}

type Temperature struct {
	common.Entity
	// AdjustedMaxAllowableOperatingValue shall
	// indicate the adjusted maximum allowable operating temperature for the
	// equipment monitored by this temperature sensor, as specified by a
	// standards body, manufacturer, or a combination, and adjusted based on
	// environmental conditions present. For example, liquid inlet
	// temperature may be adjusted based on the available liquid pressure.
	AdjustedMaxAllowableOperatingValue float32
	// AdjustedMinAllowableOperatingValue shall
	// indicate the adjusted minimum allowable operating temperature for the
	// equipment monitored by this temperature sensor, as specified by a
	// standards body, manufacturer, or a combination, and adjusted based on
	// environmental conditions present. For example, liquid inlet
	// temperature may be adjusted based on the available liquid pressure.
	AdjustedMinAllowableOperatingValue float32
	// DeltaPhysicalContext shall be a description of the affected device or
	// region within the chassis to which the DeltaReadingCelsius temperature
	// measurement applies, relative to PhysicalContext.
	DeltaPhysicalContext PhysicalContext
	// DeltaReadingCelsius shall be the delta of the values of the temperature
	// readings across this sensor and the sensor at DeltaPhysicalContext.
	DeltaReadingCelsius float32
	// LowerThresholdCritical shall indicate
	// the ReadingCelsius is below the normal range but is not yet fatal. The
	// units shall be the same units as the related ReadingCelsius property.
	LowerThresholdCritical float32
	// LowerThresholdFatal shall indicate the
	// ReadingCelsius is below the normal range and is fatal. The units shall
	// be the same units as the related ReadingCelsius property.
	LowerThresholdFatal float32
	// LowerThresholdNonCritical shall indicate
	// the ReadingCelsius is below the normal range but is not critical. The
	// units shall be the same units as the related ReadingCelsius property.
	LowerThresholdNonCritical float32
	// LowerThresholdUser shall contain the value at which
	// the ReadingCelsius property is below the user-defined range. The
	// value of the property shall use the same units as the ReadingCelsius
	// property. The value shall be equal to the value of
	// LowerThresholdNonCritical, LowerThresholdCritical, or
	// LowerThresholdFatal, unless set by a user.
	LowerThresholdUser float32
	// MaxAllowableOperatingValue shall
	// indicate the maximum allowable operating temperature for the equipment
	// monitored by this temperature sensor, as specified by a standards
	// body, manufacturer, or a combination.
	MaxAllowableOperatingValue float32
	// MaxReadingRangeTemp shall indicate the
	// highest possible value for ReadingCelsius. The units shall be the same
	// units as the related ReadingCelsius property.
	MaxReadingRangeTemp float32
	// MemberID shall uniquely identify the member within the collection. For
	// services supporting Redfish v1.6 or higher, this value shall be the
	// zero-based array index.
	MemberID string
	// MinAllowableOperatingValue shall indicate the minimum allowable operating
	// temperature for the equipment monitored by this temperature sensor, as
	// specified by a standards body, manufacturer, or a combination.
	MinAllowableOperatingValue float32
	// MinReadingRangeTemp shall indicate the lowest possible value for
	// ReadingCelsius. The units shall be the same units as the related
	// ReadingCelsius property.
	MinReadingRangeTemp float32
	// PhysicalContext shall be a description of the affected device or region
	// within the chassis to which this temperature measurement applies.
	PhysicalContext PhysicalContext
	// ReadingCelsius shall be the current value of the temperature sensor's reading.
	ReadingCelsius float32
	// RelatedItem shall contain an array of links to resources or objects that represent areas or devices to which
	// this temperature applies.
	relatedItem []string
	// RelatedItemCount is the number of related items.
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SensorNumber shall be a numerical identifier for this temperature sensor
	// that is unique within this resource.
	SensorNumber int
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UpperThresholdCritical shall indicate
	// the ReadingCelsius is above the normal range but is not yet fatal. The
	// units shall be the same units as the related ReadingCelsius property.
	UpperThresholdCritical float32
	// UpperThresholdFatal shall indicate the
	// ReadingCelsius is above the normal range and is fatal. The units shall
	// be the same units as the related ReadingCelsius property.
	UpperThresholdFatal float32
	// UpperThresholdNonCritical shall indicate
	// the ReadingCelsius is above the normal range but is not critical. The
	// units shall be the same units as the related ReadingCelsius property.
	UpperThresholdNonCritical float32
	// UpperThresholdUser shall contain the value at which
	// the ReadingCelsius property is above the user-defined range. The
	// value of the property shall use the same units as the ReadingCelsius
	// property. The value shall be equal to the value of
	// UpperThresholdNonCritical, UpperThresholdCritical, or
	// UpperThresholdFatal, unless set by a user.
	UpperThresholdUser float32
	// rawData holds the original serialized JSON so we can compare updates.
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
	// Get a representation of the object's original state so we can find what
	// to update.
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

// Thermal is used to represent a thermal metrics resource for a Redfish
// implementation.
type Thermal struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Fans shall be the definition for fans for a Redfish implementation.
	Fans []ThermalFan
	// FansCount is the number of ThermalFans.
	FansCount int `json:"Fans@odata.count"`
	// Redundancy is used to show redundancy for fans and other elements in
	// this resource. The use of IDs within these arrays shall reference the
	// members of the redundancy groups.
	redundancy []string
	// RedundancyCount is the number of Redundancy objects.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Temperatures shall be the definition for temperature sensors for a
	// Redfish implementation.
	Temperatures []Temperature
	// TemperaturesCount is the number of Temperature objects
	TemperaturesCount int `json:"Temperatures@odata.count"`
	// Oem shall contain the OEM extensions. All values for properties that
	// this object contains shall conform to the Redfish Specification
	// described requirements.
	Oem json.RawMessage
	// rawData holds the original serialized JSON so we can compare updates.
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
	// Get a representation of the object's original state so we can find what
	// to update.
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
