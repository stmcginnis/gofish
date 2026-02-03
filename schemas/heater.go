//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2022.3 - #Heater.v1_0_2.Heater

package schemas

import (
	"encoding/json"
)

// Heater shall represent the management properties for monitoring and
// management of heaters for a Redfish implementation.
type Heater struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	assembly string
	// HotPluggable shall indicate whether the device can be inserted or removed
	// while the underlying equipment otherwise remains in its current operational
	// state. Hot-pluggable devices can become operable without altering the
	// operational state of the underlying equipment. Devices that cannot be
	// inserted or removed from equipment in operation, or devices that cannot
	// become operable without affecting the operational state of that equipment,
	// shall not be hot-pluggable.
	HotPluggable bool
	// Location shall contain the location information of this heater.
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for
	// producing the heater. This organization may be the entity from whom the
	// heater is purchased, but this is not necessarily true.
	Manufacturer string
	// Metrics shall contain a link to a resource of type 'HeaterMetrics'.
	metrics string
	// Model shall contain the model information as defined by the manufacturer for
	// this heater.
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
	// this heater.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region
	// within the chassis with which this heater is associated.
	PhysicalContext PhysicalContext
	// SerialNumber shall contain the serial number as defined by the manufacturer
	// for this heater.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as
	// defined by the manufacturer for this heater.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// managers are the URIs for Managers.
	managers []string
	// memory are the URIs for Memory.
	memory []string
	// networkAdapters are the URIs for NetworkAdapters.
	networkAdapters []string
	// processors are the URIs for Processors.
	processors []string
	// storageControllers are the URIs for StorageControllers.
	storageControllers []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Heater object from the raw JSON.
func (h *Heater) UnmarshalJSON(b []byte) error {
	type temp Heater
	type hLinks struct {
		Managers           Links `json:"Managers"`
		Memory             Links `json:"Memory"`
		NetworkAdapters    Links `json:"NetworkAdapters"`
		Processors         Links `json:"Processors"`
		StorageControllers Links `json:"StorageControllers"`
	}
	var tmp struct {
		temp
		Links    hLinks
		Assembly Link `json:"Assembly"`
		Metrics  Link `json:"Metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*h = Heater(tmp.temp)

	// Extract the links to other entities for later
	h.managers = tmp.Links.Managers.ToStrings()
	h.memory = tmp.Links.Memory.ToStrings()
	h.networkAdapters = tmp.Links.NetworkAdapters.ToStrings()
	h.processors = tmp.Links.Processors.ToStrings()
	h.storageControllers = tmp.Links.StorageControllers.ToStrings()
	h.assembly = tmp.Assembly.String()
	h.metrics = tmp.Metrics.String()

	// This is a read/write object, so we need to save the raw object data for later
	h.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (h *Heater) Update() error {
	readWriteFields := []string{
		"LocationIndicatorActive",
	}

	return h.UpdateFromRawData(h, h.RawData, readWriteFields)
}

// GetHeater will get a Heater instance from the service.
func GetHeater(c Client, uri string) (*Heater, error) {
	return GetObject[Heater](c, uri)
}

// ListReferencedHeaters gets the collection of Heater from
// a provided reference.
func ListReferencedHeaters(c Client, link string) ([]*Heater, error) {
	return GetCollectionObjects[Heater](c, link)
}

// Managers gets the Managers linked resources.
func (h *Heater) Managers() ([]*Manager, error) {
	return GetObjects[Manager](h.client, h.managers)
}

// Memory gets the Memory linked resources.
func (h *Heater) Memory() ([]*Memory, error) {
	return GetObjects[Memory](h.client, h.memory)
}

// NetworkAdapters gets the NetworkAdapters linked resources.
func (h *Heater) NetworkAdapters() ([]*NetworkAdapter, error) {
	return GetObjects[NetworkAdapter](h.client, h.networkAdapters)
}

// Processors gets the Processors linked resources.
func (h *Heater) Processors() ([]*Processor, error) {
	return GetObjects[Processor](h.client, h.processors)
}

// StorageControllers gets the StorageControllers linked resources.
func (h *Heater) StorageControllers() ([]*StorageController, error) {
	return GetObjects[StorageController](h.client, h.storageControllers)
}

// Assembly gets the Assembly linked resource.
func (h *Heater) Assembly() (*Assembly, error) {
	if h.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](h.client, h.assembly)
}

// Metrics gets the Metrics linked resource.
func (h *Heater) Metrics() (*HeaterMetrics, error) {
	if h.metrics == "" {
		return nil, nil
	}
	return GetObject[HeaterMetrics](h.client, h.metrics)
}
