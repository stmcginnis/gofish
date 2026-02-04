//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Fan.v1_6_0.json
// 2025.3 - #Fan.v1_6_0.Fan

package schemas

import (
	"encoding/json"
)

// Fan shall represent a cooling fan for a Redfish implementation. It may also
// represent a location, such as a slot, socket, or bay, where a unit may be
// installed, but the 'State' property within the 'Status' property contains
// 'Absent'.
type Fan struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	assembly string
	// FanDiameterMm shall contain the diameter of the fan assembly in millimeter
	// units.
	//
	// Version added: v1.4.0
	FanDiameterMm *uint `json:",omitempty"`
	// HotPluggable shall indicate whether the device can be inserted or removed
	// while the underlying equipment otherwise remains in its current operational
	// state. Hot-pluggable devices can become operable without altering the
	// operational state of the underlying equipment. Devices that cannot be
	// inserted or removed from equipment in operation, or devices that cannot
	// become operable without affecting the operational state of that equipment,
	// shall not be hot-pluggable.
	HotPluggable bool
	// Location shall contain the location information of this fan.
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for
	// producing the fan. This organization may be the entity from whom the fan is
	// purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the model information as defined by the manufacturer for
	// this fan.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for
	// this fan.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region
	// within the chassis with which this fan is associated.
	PhysicalContext PhysicalContext
	// PowerWatts shall contain the total power, in watt units, for this resource.
	// The value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Power'.
	//
	// Version added: v1.1.0
	PowerWatts SensorPowerExcerpt
	// RatedSecondarySpeedRPM shall contain the rated maximum rotational speed of
	// the second rotor in a multi-rotor fan in revolutions per minute (RPM) units.
	//
	// Version added: v1.6.0
	RatedSecondarySpeedRPM *int `json:",omitempty"`
	// RatedSpeedRPM shall contain the rated maximum rotational speed of the fan in
	// revolutions per minute (RPM) units.
	//
	// Version added: v1.6.0
	RatedSpeedRPM *int `json:",omitempty"`
	// Replaceable shall indicate whether this component can be independently
	// replaced as allowed by the vendor's replacement policy. A value of 'false'
	// indicates the component needs to be replaced by policy as part of another
	// component. If the 'LocationType' property of this component contains
	// 'Embedded', this property shall contain 'false'.
	//
	// Version added: v1.3.0
	Replaceable bool
	// SecondarySpeedPercent shall contain the fan speed, in percent units, for the
	// secondary rotor of this resource. Services should calculate the value of
	// 'Reading' by dividing 'SpeedRPM' by 'RatedSecondarySpeedRPM'. The value of
	// the 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value 'Percent'.
	//
	// Version added: v1.5.0
	SecondarySpeedPercent SensorFanExcerpt
	// SerialNumber shall contain the serial number as defined by the manufacturer
	// for this fan.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as
	// defined by the manufacturer for this fan.
	SparePartNumber string
	// SpeedPercent shall contain the fan speed, in percent units, for this
	// resource. Services should calculate the value of 'Reading' by dividing
	// 'SpeedRPM' by 'RatedSpeedRPM'. The value of the 'DataSourceUri' property, if
	// present, shall reference a resource of type 'Sensor' with the 'ReadingType'
	// property containing the value 'Percent'.
	SpeedPercent SensorFanExcerpt
	// Status shall contain any status or health properties of the resource.
	Status Status
	// coolingChassis are the URIs for CoolingChassis.
	coolingChassis []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Fan object from the raw JSON.
func (f *Fan) UnmarshalJSON(b []byte) error {
	type temp Fan
	type fLinks struct {
		CoolingChassis Links `json:"CoolingChassis"`
	}
	var tmp struct {
		temp
		Links    fLinks
		Assembly Link `json:"Assembly"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = Fan(tmp.temp)

	// Extract the links to other entities for later
	f.coolingChassis = tmp.Links.CoolingChassis.ToStrings()
	f.assembly = tmp.Assembly.String()

	// This is a read/write object, so we need to save the raw object data for later
	f.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (f *Fan) Update() error {
	readWriteFields := []string{
		"LocationIndicatorActive",
	}

	return f.UpdateFromRawData(f, f.RawData, readWriteFields)
}

// GetFan will get a Fan instance from the service.
func GetFan(c Client, uri string) (*Fan, error) {
	return GetObject[Fan](c, uri)
}

// ListReferencedFans gets the collection of Fan from
// a provided reference.
func ListReferencedFans(c Client, link string) ([]*Fan, error) {
	return GetCollectionObjects[Fan](c, link)
}

// CoolingChassis gets the CoolingChassis linked resources.
func (f *Fan) CoolingChassis() ([]*Chassis, error) {
	return GetObjects[Chassis](f.client, f.coolingChassis)
}

// Assembly gets the Assembly linked resource.
func (f *Fan) Assembly() (*Assembly, error) {
	if f.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](f.client, f.assembly)
}
