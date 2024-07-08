//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Fan shall represent a cooling fan for a Redfish implementation. It may also represent a location, such as a
// slot, socket, or bay, where a unit may be installed, but the State property within the Status property contains
// 'Absent'.
type Fan struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Assembly shall contain a link to a resource of type Assembly.
	assembly string
	// Description provides a description of this resource.
	Description string
	// FanDiameterMm shall contain the diameter of the fan assembly in millimeter units.
	FanDiameterMm int
	// HotPluggable shall indicate whether the device can be inserted or removed while the underlying equipment
	// otherwise remains in its current operational state. Hot-pluggable devices can become operable without altering
	// the operational state of the underlying equipment. Devices that cannot be inserted or removed from equipment in
	// operation, or devices that cannot become operable without affecting the operational state of that equipment,
	// shall not be hot-pluggable.
	HotPluggable bool
	// Location shall contain the location information of this fan.
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for producing the fan. This organization may
	// be the entity from whom the fan is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the model information as defined by the manufacturer for this fan.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for this fan.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region within the chassis with which this
	// fan is associated.
	PhysicalContext PhysicalContext
	// PowerWatts shall contain the total power, in watt units, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Power'.
	PowerWatts SensorPowerExcerpt
	// Replaceable shall indicate whether this component can be independently replaced as allowed by the vendor's
	// replacement policy. A value of 'false' indicates the component needs to be replaced by policy as part of another
	// component. If the 'LocationType' property of this component contains 'Embedded', this property shall contain
	// 'false'.
	Replaceable bool
	// SecondarySpeedPercent shall contain the fan speed, in percent units, for the secondary rotor of this resource.
	// The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'Percent'.
	SecondarySpeedPercent SensorFanExcerpt
	// SerialNumber shall contain the serial number as defined by the manufacturer for this fan.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as defined by the manufacturer for this fan.
	SparePartNumber string
	// SpeedPercent shall contain the fan speed, in percent units, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Percent'.
	SpeedPercent SensorFanExcerpt
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	coolingChassis []string
	// CoolingChassisCount is the number of cooling chassis related to this fan.
	CoolingChassisCount int
}

// UnmarshalJSON unmarshals a Fan object from the raw JSON.
func (fan *Fan) UnmarshalJSON(b []byte) error {
	type temp Fan
	type Links struct {
		// CoolingChassis shall contain an array of links to resources of type Chassis that represent the chassis directly
		// cooled by this fan. This property shall not be present if the fan is only providing cooling to its containing
		// chassis.
		CoolingChassis      common.Links
		CoolingChassisCount int `json:"CoolingChassis@odata.count"`
	}
	var t struct {
		temp
		FanName  string
		Assembly common.Link
		Links    Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*fan = Fan(t.temp)

	// Extract the links to other entities for later
	fan.assembly = t.Assembly.String()
	fan.coolingChassis = t.Links.CoolingChassis.ToStrings()
	fan.CoolingChassisCount = t.Links.CoolingChassisCount

	if t.FanName != "" {
		fan.Name = t.FanName
	}

	// This is a read/write object, so we need to save the raw object data for later
	fan.rawData = b

	return nil
}

// Assembly gets the assembly for this fan.
func (fan *Fan) Assembly() (*Assembly, error) {
	if fan.assembly == "" {
		return nil, nil
	}
	return GetAssembly(fan.GetClient(), fan.assembly)
}

// CoolingChassis get the cooling chassis related to this fan.
func (fan *Fan) CoolingChassis() ([]*Chassis, error) {
	return common.GetObjects[Chassis](fan.GetClient(), fan.coolingChassis)
}

// Update commits updates to this object's properties to the running system.
func (fan *Fan) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Fan)
	original.UnmarshalJSON(fan.rawData)

	readWriteFields := []string{
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(fan).Elem()

	return fan.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetFan will get a Fan instance from the service.
func GetFan(c common.Client, uri string) (*Fan, error) {
	return common.GetObject[Fan](c, uri)
}

// ListReferencedFans gets the collection of Fan from
// a provided reference.
func ListReferencedFans(c common.Client, link string) ([]*Fan, error) {
	return common.GetCollectionObjects[Fan](c, link)
}
