//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type PumpType string

const (
	// LiquidPumpType is a water or liquid pump.
	LiquidPumpType PumpType = "Liquid"
	// CompressorPumpType is a compressor pump.
	CompressorPumpType PumpType = "Compressor"
)

// Pump shall represent the management properties for monitoring and management of pumps for a Redfish
// implementation.
type Pump struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Assembly shall contain a link to a resource of type Assembly.
	assembly string
	// AssetTag shall contain the user-assigned asset tag, which is an identifying string that tracks the equipment for
	// inventory purposes.
	AssetTag string
	// Description provides a description of this resource.
	Description string
	// Filters shall contain a link to a resource collection of type FilterCollection that contains a set of filters.
	filters []string
	// FirmwareVersion shall contain a string describing the firmware version of this equipment as provided by the
	// manufacturer.
	FirmwareVersion string
	// Location shall contain the location information of this pump.
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for producing the pump. This organization
	// may be the entity from whom the pump is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the model information as defined by the manufacturer for this pump.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for this pump.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region within the chassis with which this
	// pump is associated.
	PhysicalContext PhysicalContext
	// ProductionDate shall contain the date of production or manufacture for this equipment.
	ProductionDate string
	// PumpSpeedPercent shall contain the current speed, in percent units, of this pump. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Percent'.
	PumpSpeedPercent SensorPumpExcerpt
	// PumpType shall contain the type of pump represented by this resource.
	PumpType PumpType
	// SerialNumber shall contain the serial number as defined by the manufacturer for this pump.
	SerialNumber string
	// ServiceHours shall contain the number of hours of service that the pump has been in operation.
	ServiceHours float64
	// SparePartNumber shall contain the spare or replacement part number as defined by the manufacturer for this pump.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// UserLabel shall contain a user-assigned label used to identify this resource. If a value has not been assigned
	// by a user, the value of this property shall be an empty string.
	UserLabel string
	// Version shall contain the hardware version of this equipment as determined by the vendor or supplier.
	Version string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Pump object from the raw JSON.
func (pump *Pump) UnmarshalJSON(b []byte) error {
	type temp Pump
	var t struct {
		temp
		Assembly common.Link
		Filters  common.LinksCollection
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*pump = Pump(t.temp)

	// Extract the links to other entities for later
	pump.assembly = t.Assembly.String()
	pump.filters = t.Filters.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	pump.rawData = b

	return nil
}

// Assembly gets the containing assembly for this pump.
func (pump *Pump) Assembly() (*Assembly, error) {
	if pump.assembly == "" {
		return nil, nil
	}
	return GetAssembly(pump.GetClient(), pump.assembly)
}

// Filters gets a collection of filters.
func (pump *Pump) Filters() ([]*Filter, error) {
	return common.GetObjects[Filter](pump.GetClient(), pump.filters)
}

// Update commits updates to this object's properties to the running system.
func (pump *Pump) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Pump)
	original.UnmarshalJSON(pump.rawData)

	readWriteFields := []string{
		"AssetTag",
		"LocationIndicatorActive",
		"ServiceHours",
		"UserLabel",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(pump).Elem()

	return pump.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetPump will get a Pump instance from the service.
func GetPump(c common.Client, uri string) (*Pump, error) {
	return common.GetObject[Pump](c, uri)
}

// ListReferencedPumps gets the collection of Pump from
// a provided reference.
func ListReferencedPumps(c common.Client, link string) ([]*Pump, error) {
	return common.GetCollectionObjects[Pump](c, link)
}
