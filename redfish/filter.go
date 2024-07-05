//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Filter shall represent the management properties for monitoring and management of filters for a Redfish
// implementation.
type Filter struct {
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
	// HotPluggable shall indicate whether the device can be inserted or removed while the underlying equipment
	// otherwise remains in its current operational state. Devices indicated as hot-pluggable shall allow the device to
	// become operable without altering the operational state of the underlying equipment. Devices that cannot be
	// inserted or removed from equipment in operation, or devices that cannot become operable without affecting the
	// operational state of that equipment, shall be indicated as not hot-pluggable.
	HotPluggable bool
	// Location shall contain the location information of this filter.
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for producing the filter. This organization
	// may be the entity from whom the Filter is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the model information as defined by the manufacturer for this filter.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for this filter.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region within the chassis with which this
	// filter is associated.
	PhysicalContext PhysicalContext
	// RatedServiceHours shall contain the number of hours of service that the filter or filter media is rated to
	// provide before servicing or replacement is necessary.
	RatedServiceHours float64
	// Replaceable shall indicate whether this component can be independently replaced as allowed by the vendor's
	// replacement policy. A value of 'false' indicates the component needs to be replaced by policy as part of another
	// component. If the 'LocationType' property of this component contains 'Embedded', this property shall contain
	// 'false'.
	Replaceable bool
	// SerialNumber shall contain the serial number as defined by the manufacturer for this filter.
	SerialNumber string
	// ServiceHours shall contain the number of hours of service that the filter or filter media has provided.
	ServiceHours float64
	// ServicedDate shall contain the date the filter or filter media was put into active service.
	ServicedDate string
	// SparePartNumber shall contain the spare or replacement part number as defined by the manufacturer for this
	// filter.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UserLabel shall contain a user-assigned label used to identify this resource. If a value has not been assigned
	// by a user, the value of this property shall be an empty string.
	UserLabel string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Filter object from the raw JSON.
func (filter *Filter) UnmarshalJSON(b []byte) error {
	type temp Filter
	var t struct {
		temp
		Assembly common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*filter = Filter(t.temp)

	// Extract the links to other entities for later
	filter.assembly = t.Assembly.String()

	// This is a read/write object, so we need to save the raw object data for later
	filter.rawData = b

	return nil
}

// Assembly gets the assembly for this filter.
func (filter *Filter) Assembly() (*Assembly, error) {
	if filter.assembly == "" {
		return nil, nil
	}
	return GetAssembly(filter.GetClient(), filter.assembly)
}

// Update commits updates to this object's properties to the running system.
func (filter *Filter) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Filter)
	original.UnmarshalJSON(filter.rawData)

	readWriteFields := []string{
		"LocationIndicatorActive",
		"ServiceHours",
		"ServicedDate",
		"UserLabel",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(filter).Elem()

	return filter.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetFilter will get a Filter instance from the service.
func GetFilter(c common.Client, uri string) (*Filter, error) {
	return common.GetObject[Filter](c, uri)
}

// ListReferencedFilters gets the collection of Filter from
// a provided reference.
func ListReferencedFilters(c common.Client, link string) ([]*Filter, error) {
	return common.GetCollectionObjects[Filter](c, link)
}
