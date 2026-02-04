//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Filter.v1_1_0.json
// 2025.2 - #Filter.v1_1_0.Filter

package schemas

import (
	"encoding/json"
)

// Filter shall represent the management properties for monitoring and
// management of filters for a Redfish implementation.
type Filter struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	assembly string
	// DeltaLiquidPressurekPa shall contain the pressure, in kilopascal units, for
	// the difference in pressure between the intake and outflow connections on the
	// filter. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'PressurekPa'.
	//
	// Version added: v1.1.0
	DeltaLiquidPressurekPa SensorExcerpt
	// HotPluggable shall indicate whether the device can be inserted or removed
	// while the underlying equipment otherwise remains in its current operational
	// state. Devices indicated as hot-pluggable shall allow the device to become
	// operable without altering the operational state of the underlying equipment.
	// Devices that cannot be inserted or removed from equipment in operation, or
	// devices that cannot become operable without affecting the operational state
	// of that equipment, shall be indicated as not hot-pluggable.
	HotPluggable bool
	// Location shall contain the location information of this filter.
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for
	// producing the filter. This organization may be the entity from whom the
	// filter is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the model information as defined by the manufacturer for
	// this filter.
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
	// this filter.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region
	// within the chassis with which this filter is associated.
	PhysicalContext PhysicalContext
	// RatedServiceHours shall contain the number of hours of service that the
	// filter or filter media is rated to provide before servicing or replacement
	// is necessary.
	RatedServiceHours *float64 `json:",omitempty"`
	// Replaceable shall indicate whether this component can be independently
	// replaced as allowed by the vendor's replacement policy. A value of 'false'
	// indicates the component needs to be replaced by policy as part of another
	// component. If the 'LocationType' property of this component contains
	// 'Embedded', this property shall contain 'false'.
	Replaceable bool
	// SerialNumber shall contain the serial number as defined by the manufacturer
	// for this filter.
	SerialNumber string
	// ServiceHours shall contain the number of hours of service that the filter or
	// filter media has provided. The service may reset or update the value in
	// response to an update of 'ServicedDate'.
	ServiceHours *float64 `json:",omitempty"`
	// ServicedDate shall contain the date the filter or filter media was put into
	// active service. The service may update the value in response to an update of
	// 'ServiceHours'.
	ServicedDate string
	// SparePartNumber shall contain the spare or replacement part number as
	// defined by the manufacturer for this filter.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UserLabel shall contain a user-assigned label used to identify this
	// resource. If a value has not been assigned by a user, the value of this
	// property shall be an empty string.
	UserLabel string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Filter object from the raw JSON.
func (f *Filter) UnmarshalJSON(b []byte) error {
	type temp Filter
	var tmp struct {
		temp
		Assembly Link `json:"Assembly"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = Filter(tmp.temp)

	// Extract the links to other entities for later
	f.assembly = tmp.Assembly.String()

	// This is a read/write object, so we need to save the raw object data for later
	f.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (f *Filter) Update() error {
	readWriteFields := []string{
		"LocationIndicatorActive",
		"ServiceHours",
		"ServicedDate",
		"UserLabel",
	}

	return f.UpdateFromRawData(f, f.RawData, readWriteFields)
}

// GetFilter will get a Filter instance from the service.
func GetFilter(c Client, uri string) (*Filter, error) {
	return GetObject[Filter](c, uri)
}

// ListReferencedFilters gets the collection of Filter from
// a provided reference.
func ListReferencedFilters(c Client, link string) ([]*Filter, error) {
	return GetCollectionObjects[Filter](c, link)
}

// Assembly gets the Assembly linked resource.
func (f *Filter) Assembly() (*Assembly, error) {
	if f.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](f.client, f.assembly)
}
