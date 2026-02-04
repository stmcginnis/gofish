//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Thermal.v1_7_3.json
// 2020.4 - #Thermal.v1_7_3.Thermal

package schemas

import (
	"encoding/json"
)

type ReadingUnits string

const (
	// RPMReadingUnits The fan reading and thresholds are measured in revolutions
	// per minute.
	RPMReadingUnits ReadingUnits = "RPM"
	// PercentReadingUnits The fan reading and thresholds are measured as a
	// percentage.
	PercentReadingUnits ReadingUnits = "Percent"
)

// Thermal shall contain the thermal management properties for temperature
// monitoring and management of cooling fans for a Redfish implementation.
type Thermal struct {
	Entity
	// Fans shall contain the set of fans for this chassis.
	Fans []ThermalFan
	// FansCount
	FansCount int `json:"Fans@odata.count"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Redundancy shall contain redundancy information for the fans in this
	// chassis.
	redundancy string
	// RedundancyCount
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// Temperatures shall contain the set of temperature sensors for this chassis.
	Temperatures []Temperature
	// TemperaturesCount
	TemperaturesCount int `json:"Temperatures@odata.count"`
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Thermal object from the raw JSON.
func (t *Thermal) UnmarshalJSON(b []byte) error {
	type temp Thermal
	var tmp struct {
		temp
		Redundancy Link `json:"Redundancy"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = Thermal(tmp.temp)

	// Extract the links to other entities for later
	t.redundancy = tmp.Redundancy.String()

	// This is a read/write object, so we need to save the raw object data for later
	t.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (t *Thermal) Update() error {
	readWriteFields := []string{
		"Fans",
		"Temperatures",
	}

	return t.UpdateFromRawData(t, t.RawData, readWriteFields)
}

// GetThermal will get a Thermal instance from the service.
func GetThermal(c Client, uri string) (*Thermal, error) {
	return GetObject[Thermal](c, uri)
}

// ListReferencedThermals gets the collection of Thermal from
// a provided reference.
func ListReferencedThermals(c Client, link string) ([]*Thermal, error) {
	return GetCollectionObjects[Thermal](c, link)
}

// Redundancy gets the Redundancy linked resource.
func (t *Thermal) Redundancy() (*Redundancy, error) {
	if t.redundancy == "" {
		return nil, nil
	}
	return GetObject[Redundancy](t.client, t.redundancy)
}

// ThermalFan represents the Fan type.
type ThermalFan struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	//
	// Version added: v1.4.0
	assembly string
	// FanName shall contain the name of the fan.
	//
	// Deprecated: v1.1.0
	// This property has been deprecated in favor of the 'Name' property.
	FanName string
	// HotPluggable shall indicate whether the device can be inserted or removed
	// while the underlying equipment otherwise remains in its current operational
	// state. Hot-pluggable devices can become operable without altering the
	// operational state of the underlying equipment. Devices that cannot be
	// inserted or removed from equipment in operation, or devices that cannot
	// become operable without affecting the operational state of that equipment,
	// shall not be hot-pluggable.
	//
	// Version added: v1.4.0
	HotPluggable bool
	// IndicatorLED shall contain the state of the indicator light associated with
	// this fan.
	//
	// Version added: v1.2.0
	IndicatorLED IndicatorLED
	// Location shall contain the location information of the associated fan.
	//
	// Version added: v1.4.0
	Location Location
	// LowerThresholdCritical shall contain the value at which the 'Reading'
	// property is below the normal range but is not yet fatal. The value of the
	// property shall use the same units as the 'Reading' property.
	LowerThresholdCritical *int `json:",omitempty"`
	// LowerThresholdFatal shall contain the value at which the 'Reading' property
	// is below the normal range and is fatal. The value of the property shall use
	// the same units as the 'Reading' property.
	LowerThresholdFatal *int `json:",omitempty"`
	// LowerThresholdNonCritical shall contain the value at which the 'Reading'
	// property is below normal range. The value of the property shall use the same
	// units as the 'Reading' property.
	LowerThresholdNonCritical *int `json:",omitempty"`
	// Manufacturer shall contain the name of the organization responsible for
	// producing the fan. This organization may be the entity from whom the fan is
	// purchased, but this is not necessarily true.
	//
	// Version added: v1.2.0
	Manufacturer string
	// MaxReadingRange shall indicate the highest possible value for the 'Reading'
	// property. The value of the property shall use the same units as the
	// 'Reading' property.
	MaxReadingRange *int `json:",omitempty"`
	// MemberID shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// MinReadingRange shall indicate the lowest possible value for the 'Reading'
	// property. The value of the property shall use the same units as the
	// 'Reading' property.
	MinReadingRange *int `json:",omitempty"`
	// Model shall contain the model information as defined by the manufacturer for
	// the associated fan.
	//
	// Version added: v1.2.0
	Model string
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for
	// the associated fan.
	//
	// Version added: v1.2.0
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region
	// within the chassis with which this fan is associated.
	PhysicalContext PhysicalContext
	// Reading shall contain the fan sensor reading.
	Reading *int `json:",omitempty"`
	// ReadingUnits shall contain the units in which the fan reading and thresholds
	// are measured.
	//
	// Version added: v1.0.1
	ReadingUnits ReadingUnits
	// Redundancy shall contain an array of links to the redundancy groups to which
	// this fan belongs.
	redundancy string
	// RedundancyCount
	RedundancyCount int `json:"Redundancy@odata.count"`
	// RelatedItem shall contain an array of links to resources or objects that
	// this fan services.
	relatedItem []string
	// RelatedItemCount
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SensorNumber shall contain a numerical identifier for this fan speed sensor
	// that is unique within this resource.
	//
	// Version added: v1.5.0
	SensorNumber *int `json:",omitempty"`
	// SerialNumber shall contain the serial number as defined by the manufacturer
	// for the associated fan.
	//
	// Version added: v1.2.0
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as
	// defined by the manufacturer for the associated fan.
	//
	// Version added: v1.2.0
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UpperThresholdCritical shall contain the value at which the 'Reading'
	// property is above the normal range but is not yet fatal. The value of the
	// property shall use the same units as the 'Reading' property.
	UpperThresholdCritical *int `json:",omitempty"`
	// UpperThresholdFatal shall contain the value at which the 'Reading' property
	// is above the normal range and is fatal. The value of the property shall use
	// the same units as the 'Reading' property.
	UpperThresholdFatal *int `json:",omitempty"`
	// UpperThresholdNonCritical shall contain the value at which the 'Reading'
	// property is above the normal range. The value of the property shall use the
	// same units as the 'Reading' property.
	UpperThresholdNonCritical *int `json:",omitempty"`
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Fan object from the raw JSON.
func (f *ThermalFan) UnmarshalJSON(b []byte) error {
	type temp ThermalFan
	var tmp struct {
		temp
		Assembly    Link  `json:"Assembly"`
		Redundancy  Link  `json:"Redundancy"`
		RelatedItem Links `json:"RelatedItem"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = ThermalFan(tmp.temp)

	// Extract the links to other entities for later
	f.assembly = tmp.Assembly.String()
	f.redundancy = tmp.Redundancy.String()
	f.relatedItem = tmp.RelatedItem.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	f.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (f *ThermalFan) Update() error {
	readWriteFields := []string{
		"IndicatorLED",
	}

	return f.UpdateFromRawData(f, f.RawData, readWriteFields)
}

// GetThermalFan will get a ThermalFan instance from the service.
func GetThermalFan(c Client, uri string) (*ThermalFan, error) {
	return GetObject[ThermalFan](c, uri)
}

// ListReferencedThermalFans gets the collection of ThermalFan from
// a provided reference.
func ListReferencedThermalFans(c Client, link string) ([]*ThermalFan, error) {
	return GetCollectionObjects[ThermalFan](c, link)
}

// Assembly gets the Assembly linked resource.
func (f *ThermalFan) Assembly() (*Assembly, error) {
	if f.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](f.client, f.assembly)
}

// Redundancy gets the Redundancy linked resource.
func (f *ThermalFan) Redundancy() (*Redundancy, error) {
	if f.redundancy == "" {
		return nil, nil
	}
	return GetObject[Redundancy](f.client, f.redundancy)
}

// RelatedItem gets the RelatedItem linked resources.
func (f *ThermalFan) RelatedItem() ([]*Entity, error) {
	return GetObjects[Entity](f.client, f.relatedItem)
}

// Temperature represents the Temperature type.
type Temperature struct {
	Entity
	// AdjustedMaxAllowableOperatingValue shall indicate the adjusted maximum
	// allowable operating temperature for the equipment monitored by this
	// temperature sensor, as specified by a standards body, manufacturer, or a
	// combination, and adjusted based on environmental conditions present. For
	// example, liquid inlet temperature can be adjusted based on the available
	// liquid pressure.
	//
	// Version added: v1.4.0
	AdjustedMaxAllowableOperatingValue *int `json:",omitempty"`
	// AdjustedMinAllowableOperatingValue shall indicate the adjusted minimum
	// allowable operating temperature for the equipment monitored by this
	// temperature sensor, as specified by a standards body, manufacturer, or a
	// combination, and adjusted based on environmental conditions present. For
	// example, liquid inlet temperature can be adjusted based on the available
	// liquid pressure.
	//
	// Version added: v1.4.0
	AdjustedMinAllowableOperatingValue *int `json:",omitempty"`
	// DeltaPhysicalContext shall contain a description of the affected device or
	// region within the chassis to which the 'DeltaReadingCelsius' temperature
	// measurement applies, relative to 'PhysicalContext'.
	//
	// Version added: v1.4.0
	DeltaPhysicalContext PhysicalContext
	// DeltaReadingCelsius shall contain the delta of the values of the temperature
	// readings across this sensor and the sensor at 'DeltaPhysicalContext'.
	//
	// Version added: v1.4.0
	DeltaReadingCelsius *float64 `json:",omitempty"`
	// LowerThresholdCritical shall contain the value at which the 'ReadingCelsius'
	// property is below the normal range but is not yet fatal. The value of the
	// property shall use the same units as the 'ReadingCelsius' property.
	LowerThresholdCritical *float64 `json:",omitempty"`
	// LowerThresholdFatal shall contain the value at which the 'ReadingCelsius'
	// property is below the normal range and is fatal. The value of the property
	// shall use the same units as the 'ReadingCelsius' property.
	LowerThresholdFatal *float64 `json:",omitempty"`
	// LowerThresholdNonCritical shall contain the value at which the
	// 'ReadingCelsius' property is below normal range. The value of the property
	// shall use the same units as the 'ReadingCelsius' property.
	LowerThresholdNonCritical *float64 `json:",omitempty"`
	// LowerThresholdUser shall contain the value at which the 'ReadingCelsius'
	// property is below the user-defined range. The value of the property shall
	// use the same units as the 'ReadingCelsius' property. The value shall be
	// equal to the value of 'LowerThresholdNonCritical', 'LowerThresholdCritical',
	// or 'LowerThresholdFatal', unless set by a user.
	//
	// Version added: v1.6.0
	LowerThresholdUser *float32 `json:",omitempty"`
	// MaxAllowableOperatingValue shall indicate the maximum allowable operating
	// temperature for the equipment monitored by this temperature sensor, as
	// specified by a standards body, manufacturer, or a combination.
	//
	// Version added: v1.4.0
	MaxAllowableOperatingValue *int `json:",omitempty"`
	// MaxReadingRangeTemp shall indicate the highest possible value for the
	// 'ReadingCelsius' property. The value of the property shall use the same
	// units as the 'ReadingCelsius' property.
	MaxReadingRangeTemp *float64 `json:",omitempty"`
	// MemberID shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// MinAllowableOperatingValue shall indicate the minimum allowable operating
	// temperature for the equipment monitored by this temperature sensor, as
	// specified by a standards body, manufacturer, or a combination.
	//
	// Version added: v1.4.0
	MinAllowableOperatingValue *int `json:",omitempty"`
	// MinReadingRangeTemp shall indicate the lowest possible value for the
	// 'ReadingCelsius' property. The value of the property shall use the same
	// units as the 'ReadingCelsius' property.
	MinReadingRangeTemp *float64 `json:",omitempty"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhysicalContext shall contain a description of the affected device or region
	// within the chassis to which this temperature applies.
	PhysicalContext PhysicalContext
	// ReadingCelsius shall contain the temperature in degree Celsius units.
	ReadingCelsius *float64 `json:",omitempty"`
	// RelatedItem shall contain an array of links to resources or objects that
	// represent areas or devices to which this temperature applies.
	relatedItem []string
	// RelatedItemCount
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SensorNumber shall contain a numerical identifier for this temperature
	// sensor that is unique within this resource.
	SensorNumber *int `json:",omitempty"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UpperThresholdCritical shall contain the value at which the 'ReadingCelsius'
	// property is above the normal range but is not yet fatal. The value of the
	// property shall use the same units as the 'ReadingCelsius' property.
	UpperThresholdCritical *float64 `json:",omitempty"`
	// UpperThresholdFatal shall contain the value at which the 'ReadingCelsius'
	// property is above the normal range and is fatal. The value of the property
	// shall use the same units as the 'ReadingCelsius' property.
	UpperThresholdFatal *float64 `json:",omitempty"`
	// UpperThresholdNonCritical shall contain the value at which the
	// 'ReadingCelsius' property is above the normal range. The value of the
	// property shall use the same units as the 'ReadingCelsius' property.
	UpperThresholdNonCritical *float64 `json:",omitempty"`
	// UpperThresholdUser shall contain the value at which the 'ReadingCelsius'
	// property is above the user-defined range. The value of the property shall
	// use the same units as the 'ReadingCelsius' property. The value shall be
	// equal to the value of 'UpperThresholdNonCritical', 'UpperThresholdCritical',
	// or 'UpperThresholdFatal', unless set by a user.
	//
	// Version added: v1.6.0
	UpperThresholdUser *float32 `json:",omitempty"`
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Temperature object from the raw JSON.
func (t *Temperature) UnmarshalJSON(b []byte) error {
	type temp Temperature
	var tmp struct {
		temp
		RelatedItem Links `json:"RelatedItem"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = Temperature(tmp.temp)

	// Extract the links to other entities for later
	t.relatedItem = tmp.RelatedItem.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	t.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (t *Temperature) Update() error {
	readWriteFields := []string{
		"LowerThresholdUser",
		"UpperThresholdUser",
	}

	return t.UpdateFromRawData(t, t.RawData, readWriteFields)
}

// GetTemperature will get a Temperature instance from the service.
func GetTemperature(c Client, uri string) (*Temperature, error) {
	return GetObject[Temperature](c, uri)
}

// ListReferencedTemperatures gets the collection of Temperature from
// a provided reference.
func ListReferencedTemperatures(c Client, link string) ([]*Temperature, error) {
	return GetCollectionObjects[Temperature](c, link)
}

// RelatedItem gets the RelatedItem linked resources.
func (t *Temperature) RelatedItem() ([]*Entity, error) {
	return GetObjects[Entity](t.client, t.relatedItem)
}
