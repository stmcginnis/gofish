//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Reservoir.v1_0_2.json
// 2023.1 - #Reservoir.v1_0_2.Reservoir

package schemas

import (
	"encoding/json"
)

type ReservoirType string

const (
	// ReserveReservoirType is a reservoir providing reserve fluid capacity.
	ReserveReservoirType ReservoirType = "Reserve"
	// OverflowReservoirType is an overflow reservoir for excess fluid.
	OverflowReservoirType ReservoirType = "Overflow"
	// InlineReservoirType is an inline or integrated reservoir.
	InlineReservoirType ReservoirType = "Inline"
	// ImmersionReservoirType is an immersion cooling tank.
	ImmersionReservoirType ReservoirType = "Immersion"
)

// Reservoir shall represent the management properties for monitoring and
// management of reservoirs for a Redfish implementation.
type Reservoir struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	assembly string
	// CapacityLiters shall contain the capacity of the reservoir in liter units.
	CapacityLiters *float64 `json:",omitempty"`
	// Coolant shall contain details regarding the coolant contained or used by
	// this unit.
	Coolant Coolant
	// Filters shall contain a link to a resource collection of type
	// 'FilterCollection' that contains a set of filters.
	filters string
	// FluidLevelPercent shall contain the amount of fluid capacity, in percent
	// units, filled in this reservoir. The value of the 'DataSourceUri' property,
	// if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'Percent'. Services that support
	// this property shall also return the 'FluidLevelStatus' property.
	FluidLevelPercent SensorExcerpt
	// FluidLevelStatus shall indicate the status of the fluid level in this
	// reservoir.
	FluidLevelStatus Health
	// InternalPressurekPa shall contain the internal pressure, measured in
	// kilopascal units, for the reservoir. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'PressurekPa'.
	InternalPressurekPa SensorExcerpt
	// Location shall contain the location information of this reservoir.
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for
	// producing the reservoir. This organization may be the entity from whom the
	// reservoir is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the model information as defined by the manufacturer for
	// this reservoir.
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
	// this reservoir.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region
	// within the chassis with which this reservoir is associated.
	PhysicalContext PhysicalContext
	// ReservoirType shall contain the type of reservoir represented by this
	// resource.
	ReservoirType ReservoirType
	// SerialNumber shall contain the serial number as defined by the manufacturer
	// for this reservoir.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as
	// defined by the manufacturer for this reservoir.
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

// UnmarshalJSON unmarshals a Reservoir object from the raw JSON.
func (r *Reservoir) UnmarshalJSON(b []byte) error {
	type temp Reservoir
	var tmp struct {
		temp
		Assembly Link `json:"Assembly"`
		Filters  Link `json:"Filters"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*r = Reservoir(tmp.temp)

	// Extract the links to other entities for later
	r.assembly = tmp.Assembly.String()
	r.filters = tmp.Filters.String()

	// This is a read/write object, so we need to save the raw object data for later
	r.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (r *Reservoir) Update() error {
	readWriteFields := []string{
		"LocationIndicatorActive",
		"UserLabel",
	}

	return r.UpdateFromRawData(r, r.RawData, readWriteFields)
}

// GetReservoir will get a Reservoir instance from the service.
func GetReservoir(c Client, uri string) (*Reservoir, error) {
	return GetObject[Reservoir](c, uri)
}

// ListReferencedReservoirs gets the collection of Reservoir from
// a provided reference.
func ListReferencedReservoirs(c Client, link string) ([]*Reservoir, error) {
	return GetCollectionObjects[Reservoir](c, link)
}

// Assembly gets the Assembly linked resource.
func (r *Reservoir) Assembly() (*Assembly, error) {
	if r.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](r.client, r.assembly)
}

// Filters gets the Filters collection.
func (r *Reservoir) Filters() ([]*Filter, error) {
	if r.filters == "" {
		return nil, nil
	}
	return GetCollectionObjects[Filter](r.client, r.filters)
}
