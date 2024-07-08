//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type ReservoirType string

const (
	// ReserveReservoirType A reservoir providing reserve fluid capacity.
	ReserveReservoirType ReservoirType = "Reserve"
	// OverflowReservoirType An overflow reservoir for excess fluid.
	OverflowReservoirType ReservoirType = "Overflow"
	// InlineReservoirType An inline or integrated reservoir.
	InlineReservoirType ReservoirType = "Inline"
	// ImmersionReservoirType An immersion cooling tank.
	ImmersionReservoirType ReservoirType = "Immersion"
)

// Reservoir shall represent the management properties for monitoring and management of reservoirs for a Redfish
// implementation.
type Reservoir struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Assembly shall contain a link to a resource of type Assembly.
	assembly string
	// CapacityLiters shall contain the capacity of the reservoir in liter units.
	CapacityLiters float64
	// Coolant shall contain details regarding the coolant contained or used by this unit.
	Coolant Coolant
	// Description provides a description of this resource.
	Description string
	// Filters shall contain a link to a resource collection of type FilterCollection that contains a set of filters.
	filters []string
	// FluidLevelPercent shall contain the amount of fluid capacity, in percent units, filled in this reservoir. The
	// value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType
	// property containing the value 'Percent'. Services that support this property shall also return the
	// FluidLevelStatus property.
	FluidLevelPercent SensorExcerpt
	// FluidLevelStatus shall indicate the status of the fluid level in this reservoir.
	FluidLevelStatus common.Health
	// InternalPressurekPa shall contain the internal pressure, measured in kilopascal units, for the reservoir. The
	// value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the ReadingType
	// property containing the value 'PressurekPa'.
	InternalPressurekPa SensorExcerpt
	// Location shall contain the location information of this reservoir.
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for producing the reservoir. This
	// organization may be the entity from whom the reservoir is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the model information as defined by the manufacturer for this reservoir.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for this reservoir.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region within the chassis with which this
	// reservoir is associated.
	PhysicalContext PhysicalContext
	// ReservoirType shall contain the type of reservoir represented by this resource.
	ReservoirType ReservoirType
	// SerialNumber shall contain the serial number as defined by the manufacturer for this reservoir.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as defined by the manufacturer for this
	// reservoir.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UserLabel shall contain a user-assigned label used to identify this resource. If a value has not been assigned
	// by a user, the value of this property shall be an empty string.
	UserLabel string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Reservoir object from the raw JSON.
func (reservoir *Reservoir) UnmarshalJSON(b []byte) error {
	type temp Reservoir
	var t struct {
		temp
		Assembly common.Link
		Filters  common.LinksCollection
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*reservoir = Reservoir(t.temp)

	// Extract the links to other entities for later
	reservoir.assembly = t.Assembly.String()
	reservoir.filters = t.Filters.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	reservoir.rawData = b

	return nil
}

// Assembly gets the containing assembly for this reservoir.
func (reservoir *Reservoir) Assembly() (*Assembly, error) {
	if reservoir.assembly == "" {
		return nil, nil
	}
	return GetAssembly(reservoir.GetClient(), reservoir.assembly)
}

// Filters gets a collection of filters.
func (reservoir *Reservoir) Filters() ([]*Filter, error) {
	return common.GetObjects[Filter](reservoir.GetClient(), reservoir.filters)
}

// Update commits updates to this object's properties to the running system.
func (reservoir *Reservoir) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Reservoir)
	original.UnmarshalJSON(reservoir.rawData)

	readWriteFields := []string{
		"LocationIndicatorActive",
		"UserLabel",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(reservoir).Elem()

	return reservoir.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetReservoir will get a Reservoir instance from the service.
func GetReservoir(c common.Client, uri string) (*Reservoir, error) {
	return common.GetObject[Reservoir](c, uri)
}

// ListReferencedReservoirs gets the collection of Reservoir from
// a provided reference.
func ListReferencedReservoirs(c common.Client, link string) ([]*Reservoir, error) {
	return common.GetCollectionObjects[Reservoir](c, link)
}
