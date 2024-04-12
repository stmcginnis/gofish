//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

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

// Temperature is
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
	DeltaPhysicalContext string
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
	MemberID string `json:"MemberID"`
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
	PhysicalContext string
	// ReadingCelsius shall be the current value of the temperature sensor's reading.
	ReadingCelsius float32
	// SensorNumber shall be a numerical identifier for this temperature sensor
	// that is unique within this resource.
	SensorNumber float32
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
	Fans []Fan
	// FansCount is the number of Fans.
	FansCount int `json:"Fans@odata.count"`
	// Redundancy is used to show redundancy for fans and other elements in
	// this resource. The use of IDs within these arrays shall reference the
	// members of the redundancy groups.
	Redundancy []Redundancy
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
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thermal = Thermal(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	thermal.rawData = b

	return nil
}

// // Update commits updates to this object's properties to the running system.
// func (thermal *Thermal) Update() error {

// 	// Get a representation of the object's original state so we can find what
// 	// to update.
// 	original := new(Thermal)
// 	original.UnmarshalJSON(thermal.rawData)

// 	readWriteFields := []string{
// 		"Fans",
// 		"Temperatures",
// 	}

// 	originalElement := reflect.ValueOf(original).Elem()
// 	currentElement := reflect.ValueOf(thermal).Elem()

// 	return thermal.Entity.Update(originalElement, currentElement, readWriteFields)
// }

// GetThermal will get a Thermal instance from the service.
func GetThermal(c common.Client, uri string) (*Thermal, error) {
	var thermal Thermal
	return &thermal, thermal.Get(c, uri, &thermal)
}

// ListReferencedThermals gets the collection of Thermal from a provided reference.
func ListReferencedThermals(c common.Client, link string) ([]*Thermal, error) {
	var result []*Thermal
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *Thermal
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		thermal, err := GetThermal(c, link)
		ch <- GetResult{Item: thermal, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
