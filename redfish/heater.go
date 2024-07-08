//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// Heater shall represent the management properties for monitoring and management of heaters for a Redfish
// implementation.
type Heater struct {
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
	// otherwise remains in its current operational state. Hot-pluggable devices can become operable without altering
	// the operational state of the underlying equipment. Devices that cannot be inserted or removed from equipment in
	// operation, or devices that cannot become operable without affecting the operational state of that equipment,
	// shall not be hot-pluggable.
	HotPluggable bool
	// Location shall contain the location information of this heater.
	Location common.Location
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for producing the heater. This organization
	// may be the entity from whom the heater is purchased, but this is not necessarily true.
	Manufacturer string
	// Metrics shall contain a link to a resource of type HeaterMetrics.
	metrics string
	// Model shall contain the model information as defined by the manufacturer for this heater.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number as defined by the manufacturer for this heater.
	PartNumber string
	// PhysicalContext shall contain a description of the affected device or region within the chassis with which this
	// heater is associated.
	PhysicalContext PhysicalContext
	// SerialNumber shall contain the serial number as defined by the manufacturer for this heater.
	SerialNumber string
	// SparePartNumber shall contain the spare or replacement part number as defined by the manufacturer for this
	// heater.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	managers []string
	// ManagersCount gets the number of managers for this heater.
	ManagersCount int
	memory        []string
	// MemoryCount gets the number of memory units associated with this heater.
	MemoryCount     int
	networkAdapters []string
	// NetworkAdaptersCount gets the number of network adapters associated with this heater.
	NetworkAdaptersCount int
	processors           []string
	// ProcessorsCount gets the number of processors associated with this heater.
	ProcessorsCount    int
	storageControllers []string
	// StorageControllersCount gets the number of storage controllers associated with this heater.
	StorageControllersCount int
}

// UnmarshalJSON unmarshals a Heater object from the raw JSON.
func (heater *Heater) UnmarshalJSON(b []byte) error {
	type temp Heater
	type Links struct {
		// Managers shall contain an array of links to the managers which this heater heats.
		Managers      common.Links
		ManagersCount int `json:"Managers@odata.count"`
		// Memory shall contain an array of links to the memory devices which this heater heats.
		Memory      common.Links
		MemoryCount int `json:"Memory@odata.count"`
		// NetworkAdapters shall contain an array of links to the network adapters which this heater heats.
		NetworkAdapters      common.Links
		NetworkAdaptersCount int `json:"NetworkAdapters@odata.count"`
		// Processors shall contain an array of links to the processors which this heater heats.
		Processors      common.Links
		ProcessorsCount int `json:"Processors@odata.count"`
		// StorageControllers shall contain an array of links to the storage controllers which this heater heats.
		StorageControllers      common.Links
		StorageControllersCount int `json:"StorageControllers@odata.count"`
	}
	var t struct {
		temp
		Assembly common.Link
		Metrics  common.Link
		Links    Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*heater = Heater(t.temp)

	// Extract the links to other entities for later
	heater.assembly = t.Assembly.String()
	heater.metrics = t.Metrics.String()

	heater.managers = t.Links.Managers.ToStrings()
	heater.ManagersCount = t.Links.ManagersCount
	heater.memory = t.Links.Memory.ToStrings()
	heater.MemoryCount = t.Links.MemoryCount
	heater.networkAdapters = t.Links.NetworkAdapters.ToStrings()
	heater.NetworkAdaptersCount = t.Links.NetworkAdaptersCount
	heater.processors = t.Links.Processors.ToStrings()
	heater.ProcessorsCount = t.Links.ProcessorsCount
	heater.storageControllers = t.Links.StorageControllers.ToStrings()
	heater.StorageControllersCount = t.Links.StorageControllersCount

	// This is a read/write object, so we need to save the raw object data for later
	heater.rawData = b

	return nil
}

// Assembly gets the assembly for this heater.
func (heater *Heater) Assembly() (*Assembly, error) {
	if heater.assembly == "" {
		return nil, nil
	}
	return GetAssembly(heater.GetClient(), heater.assembly)
}

// Managers gets the managers for this heater.
func (heater *Heater) Managers() ([]*Manager, error) {
	return common.GetObjects[Manager](heater.GetClient(), heater.managers)
}

// Memory gets the memory associated with this heater.
func (heater *Heater) Memory() ([]*Memory, error) {
	return common.GetObjects[Memory](heater.GetClient(), heater.memory)
}

// NetworkAdapters gets the network adapters associated with this heater.
func (heater *Heater) NetworkAdapters() ([]*NetworkAdapter, error) {
	return common.GetObjects[NetworkAdapter](heater.GetClient(), heater.networkAdapters)
}

// Processors gets this heater's processors.
func (heater *Heater) Processors() ([]*Processor, error) {
	return common.GetObjects[Processor](heater.GetClient(), heater.processors)
}

// StorageControllers gets the storage controllers associated with this heater.
func (heater *Heater) StorageControllers() ([]*StorageController, error) {
	return common.GetObjects[StorageController](heater.GetClient(), heater.storageControllers)
}

// Metrics gets the heater metrics for this heater.
func (heater *Heater) Metrics() (*HeaterMetrics, error) {
	if heater.metrics == "" {
		return nil, nil
	}
	return GetHeaterMetrics(heater.GetClient(), heater.metrics)
}

// Update commits updates to this object's properties to the running system.
func (heater *Heater) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Heater)
	original.UnmarshalJSON(heater.rawData)

	readWriteFields := []string{
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(heater).Elem()

	return heater.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetHeater will get a Heater instance from the service.
func GetHeater(c common.Client, uri string) (*Heater, error) {
	return common.GetObject[Heater](c, uri)
}

// ListReferencedHeaters gets the collection of Heater from
// a provided reference.
func ListReferencedHeaters(c common.Client, link string) ([]*Heater, error) {
	return common.GetCollectionObjects[Heater](c, link)
}
