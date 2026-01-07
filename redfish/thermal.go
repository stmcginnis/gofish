//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.4 - #Thermal.v1_7_3.Thermal

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
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
	common.Entity
	// Fans shall contain the set of fans for this chassis.
	Fans []ThermalFan
	// Fans@odata.count
	FansCount int `json:"Fans@odata.count"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Redundancy shall contain redundancy information for the fans in this
	// chassis.
	Redundancy []common.Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Temperatures shall contain the set of temperature sensors for this chassis.
	Temperatures []Temperature
	// Temperatures@odata.count
	TemperaturesCount int `json:"Temperatures@odata.count"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Thermal object from the raw JSON.
func (t *Thermal) UnmarshalJSON(b []byte) error {
	type temp Thermal
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = Thermal(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	t.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (t *Thermal) Update() error {
	readWriteFields := []string{
		"Fans",
		"Fans@odata.count",
		"Redundancy",
		"Redundancy@odata.count",
		"Status",
		"Temperatures",
		"Temperatures@odata.count",
	}

	return t.UpdateFromRawData(t, t.rawData, readWriteFields)
}

// GetThermal will get a Thermal instance from the service.
func GetThermal(c common.Client, uri string) (*Thermal, error) {
	return common.GetObject[Thermal](c, uri)
}

// ListReferencedThermals gets the collection of Thermal from
// a provided reference.
func ListReferencedThermals(c common.Client, link string) ([]*Thermal, error) {
	return common.GetCollectionObjects[Thermal](c, link)
}

// ThermalFan represents the ThermalFan type.
type ThermalFan struct {
	common.Entity
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
	IndicatorLED common.IndicatorLED
	// Location shall contain the location information of the associated fan.
	//
	// Version added: v1.4.0
	Location common.Location
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
	// MemberId shall contain the unique identifier for this member within an
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
	// Name is the name of the resource or array element.
	//
	// Version added: v1.1.0
	Name string
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// Oem shall contain the OEM extensions. All values for properties that this
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
	redundancy []string
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// RelatedItem shall contain an array of links to resources or objects that
	// this fan services.
	relatedItem []string
	// RelatedItem@odata.count
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
	Status common.Status
	// UpperThresholdCritical shall contain the value at which the 'Reading'
	// property is above the normal range but is not yet fatal. The value of the
	// property shall use the same units as the 'Reading' property.
	UpperThresholdCritical *float32 `json:",omitempty"`
	// UpperThresholdFatal shall contain the value at which the 'Reading' property
	// is above the normal range and is fatal. The value of the property shall use
	// the same units as the 'Reading' property.
	UpperThresholdFatal *float32 `json:",omitempty"`
	// UpperThresholdNonCritical shall contain the value at which the 'Reading'
	// property is above the normal range. The value of the property shall use the
	// same units as the 'Reading' property.
	UpperThresholdNonCritical *float32 `json:",omitempty"`

	// rawData holds original JSON for comparison during updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ThermalFan object from the raw JSON.
func (f *ThermalFan) UnmarshalJSON(b []byte) error {
	type temp ThermalFan
	var tmp struct {
		temp
		Assembly    common.Link `json:"assembly"`
		Redundancy  common.Links
		RelatedItem common.Links
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = ThermalFan(tmp.temp)

	// Extract the links to other entities for later
	f.assembly = tmp.Assembly.String()
	f.redundancy = tmp.Redundancy.ToStrings()
	f.relatedItem = tmp.RelatedItem.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	f.rawData = b

	return nil
}

// Assembly gets the Assembly linked resource.
func (f *ThermalFan) Assembly(client common.Client) (*Assembly, error) {
	if f.assembly == "" {
		return nil, nil
	}
	return common.GetObject[Assembly](client, f.assembly)
}

// Update commits updates to this object's properties to the running system.
func (f *ThermalFan) Update() error {
	readWriteFields := []string{"IndicatorLED"}

	return f.UpdateFromRawData(f, f.rawData, readWriteFields)
}

// Temperature represents the Temperature type.
type Temperature struct {
	common.Entity
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
	LowerThresholdCritical *float32 `json:",omitempty"`
	// LowerThresholdFatal shall contain the value at which the 'ReadingCelsius'
	// property is below the normal range and is fatal. The value of the property
	// shall use the same units as the 'ReadingCelsius' property.
	LowerThresholdFatal *float32 `json:",omitempty"`
	// LowerThresholdNonCritical shall contain the value at which the
	// 'ReadingCelsius' property is below normal range. The value of the property
	// shall use the same units as the 'ReadingCelsius' property.
	LowerThresholdNonCritical *float32 `json:",omitempty"`
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
	MaxAllowableOperatingValue *float32 `json:",omitempty"`
	// MaxReadingRangeTemp shall indicate the highest possible value for the
	// 'ReadingCelsius' property. The value of the property shall use the same
	// units as the 'ReadingCelsius' property.
	MaxReadingRangeTemp *float32 `json:",omitempty"`
	// MemberId shall contain the unique identifier for this member within an
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
	// Name is the name of the resource or array element.
	Name string
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// Oem shall contain the OEM extensions. All values for properties that this
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
	// RelatedItem@odata.count
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// SensorNumber shall contain a numerical identifier for this temperature
	// sensor that is unique within this resource.
	SensorNumber *int `json:",omitempty"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
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
	UpperThresholdUser *float32

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
	readWriteFields := []string{
		"LowerThresholdUser",
		"UpperThresholdUser",
	}

	return temperature.UpdateFromRawData(temperature, temperature.rawData, readWriteFields)
}
