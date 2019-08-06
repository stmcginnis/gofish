// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// CacheSummary shall contain properties which describe the cache memory for a
// storage controller.
type CacheSummary struct {
	// PersistentCacheSizeMiB shall contain the amount of
	// cache memory that is persistent as measured in mebibytes. This size
	// shall be less than or equal to the TotalCacheSizeMib.
	PersistentCacheSizeMiB int
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// TotalCacheSizeMiB shall contain the amount of
	// configured cache memory as measured in mebibytes.
	TotalCacheSizeMiB int
}

// Storage is used to represent resources that represent a storage
// subsystem in the Redfish specification.
type Storage struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Drives is a collection that indicates all the drives attached to the
	// storage controllers that this resource represents.
	drives []string
	// DrivesCount is the number of drives.
	DrivesCount int `json:"Drives@odata.count"`
	// Redundancy shall contain redundancy information for the storage subsystem.
	redundancy []string
	// RedundancyCount is the number of Redundancy objects.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// StorageControllers is a collection that indicates all the storage
	// controllers that this resource represents.
	storageControllers []string
	// StorageControllersCount is the number of
	StorageControllersCount int `json:"StorageControllers@odata.count"`
	// Volumes is a collection that indicates all the volumes produced by the
	// storage controllers that this resource represents.
	volumes string
	// Enclosures shall reference a resource of type Chassis that represents the
	// physical containers attached to this resource.
	enclosures []string
	// EnclosuresCount is the number of enclosures.
	EnclosuresCount int
}

// UnmarshalJSON unmarshals a Storage object from the raw JSON.
func (storage *Storage) UnmarshalJSON(b []byte) error {
	type temp Storage
	type links struct {
		Enclosures      common.Links
		EnclosuresCount int `json:"Enclosures@odata.count"`
	}
	var t struct {
		temp
		Links              links
		Drives             common.Links
		Redundancy         common.Links
		StorageControllers common.Links
		Volumes            common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storage = Storage(t.temp)

	// Extract the links to other entities for later
	storage.enclosures = t.Links.Enclosures.ToStrings()
	storage.EnclosuresCount = t.Links.EnclosuresCount
	storage.drives = t.Drives.ToStrings()
	storage.redundancy = t.Redundancy.ToStrings()
	storage.storageControllers = t.StorageControllers.ToStrings()
	storage.volumes = string(t.Volumes)

	return nil
}

// GetStorage will get a Storage instance from the service.
func GetStorage(c common.Client, uri string) (*Storage, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storage Storage
	err = json.NewDecoder(resp.Body).Decode(&storage)
	if err != nil {
		return nil, err
	}

	storage.SetClient(c)
	return &storage, nil
}

// ListReferencedStorages gets the collection of Storage from a provided
// reference.
func ListReferencedStorages(c common.Client, link string) ([]*Storage, error) {
	var result []*Storage
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, storageLink := range links.ItemLinks {
		storage, err := GetStorage(c, storageLink)
		if err != nil {
			return result, err
		}
		result = append(result, storage)
	}

	return result, nil
}

// Enclosures gets the physical containers attached to this resource.
func (storage *Storage) Enclosures() ([]*Chassis, error) {
	var result []*Chassis
	for _, chassisLink := range storage.enclosures {
		chassis, err := GetChassis(storage.Client, chassisLink)
		if err != nil {
			return result, nil
		}
		result = append(result, chassis)
	}
	return result, nil
}

// Drives gets the drives attached to the storage controllers that this
// resource represents.
func (storage *Storage) Drives() ([]*Drive, error) {
	var result []*Drive
	for _, driveLink := range storage.drives {
		drive, err := GetDrive(storage.Client, driveLink)
		if err != nil {
			return result, nil
		}
		result = append(result, drive)
	}
	return result, nil
}

// Redundancies gets the redundancy information for the storage subsystem.
func (storage *Storage) Redundancies() ([]*Redundancy, error) {
	var result []*Redundancy
	for _, redundancyLink := range storage.redundancy {
		redundancy, err := GetRedundancy(storage.Client, redundancyLink)
		if err != nil {
			return result, nil
		}
		result = append(result, redundancy)
	}
	return result, nil
}

// StorageControllers gets all the storage controllers that this resource represents.
func (storage *Storage) StorageControllers() ([]*StorageController, error) {
	var result []*StorageController
	for _, storageControllerLink := range storage.storageControllers {
		storageController, err := GetStorageController(storage.Client, storageControllerLink)
		if err != nil {
			return result, nil
		}
		result = append(result, storageController)
	}
	return result, nil
}

// Volumes gets the volumes associated with this storage subsystem.
func (storage *Storage) Volumes() ([]*Volume, error) {
	return ListReferencedVolumes(storage.Client, storage.volumes)
}

// StorageController is used to represent a resource that represents a
// storage controller in the Redfish specification.
type StorageController struct {
	common.Entity

	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// Assembly shall be a link to a resource of type Assembly.
	assembly string
	// AssetTag is used to track the storage controller for inventory
	// purposes.
	AssetTag string
	// CacheSummary shall contain properties which describe the cache memory for
	// the current resource.
	CacheSummary CacheSummary
	// FirmwareVersion shall contain the firwmare version as defined by the
	// manufacturer for the associated storage controller.
	FirmwareVersion string
	// Identifiers shall contain a list of all known durable names for the
	// associated storage controller.
	Identifiers []common.Identifier
	// Location shall contain location information of the
	// associated storage controller.
	Location common.Location
	// Manufacturer shall be the name of the organization responsible for
	// producing the storage controller. This organization might be the entity
	// from whom the storage controller is purchased, but this is not
	// necessarily true.
	Manufacturer string
	// MemberID shall uniquely identify the member within the collection.
	MemberID string
	// Model shall be the name by which the manufacturer generally refers to the
	// storage controller.
	Model string
	// PCIeInterface is used to connect this PCIe-based controller to its host.
	PCIeInterface PCIeInterface
	// PartNumber shall be a part number assigned by the organization that is
	// responsible for producing or manufacturing the storage controller.
	PartNumber string
	// SKU shall be the stock-keeping unit number for this storage storage
	// controller.
	SKU string
	// SerialNumber is used to identify the storage controller.
	SerialNumber string
	// SpeedGbps shall represent the maximum supported speed of the Storage bus
	// interface (in Gigabits per second). The interface specified connects the
	// controller to the storage devices, not the controller to a host (e.g. SAS
	// bus, not PCIe host bus).
	SpeedGbps int
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupportedControllerProtocols shall be the set of protocols this storage
	// controller can be communicated to.
	SupportedControllerProtocols []common.Protocol
	// SupportedDeviceProtocols shall be the set of protocols this storage
	// controller can use to communicate with attached devices.
	SupportedDeviceProtocols []common.Protocol
	// SupportedRAIDTypes shall contain all the RAIDType values supported by the
	// current resource.
	SupportedRAIDTypes []RAIDType
	// Endpoints shall be a reference to the resources that this controller is
	// associated with and shall reference a resource of type Endpoint.
	endpoints []string
	// EndpointsCount is the number of enclosures.
	EndpointsCount int
	// StorageServices shall be a reference to the resources that this
	// controller is associated with and shall reference a resource of type
	// StorageService.
	storageServices []string
	// StorageServicesCount is the number of storage services.
	StorageServicesCount int
}

// UnmarshalJSON unmarshals a StorageController object from the raw JSON.
func (storagecontroller *StorageController) UnmarshalJSON(b []byte) error {
	type temp StorageController
	type links struct {
		Endpoints            common.Links
		EndpointsCount       int `json:"Endpoints@odata.count"`
		StorageServices      common.Links
		StorageServicesCount int `json:"StorageServices@odata.count"`
	}
	var t struct {
		temp
		Assembly common.Link
		Links    links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storagecontroller = StorageController(t.temp)

	// Extract the links to other entities for later
	storagecontroller.assembly = string(t.Assembly)
	storagecontroller.endpoints = t.Links.StorageServices.ToStrings()
	storagecontroller.EndpointsCount = t.Links.EndpointsCount
	storagecontroller.storageServices = t.Links.StorageServices.ToStrings()
	storagecontroller.StorageServicesCount = t.Links.StorageServicesCount

	return nil
}

// GetStorageController will get a Storage controller instance from the service.
func GetStorageController(c common.Client, uri string) (*StorageController, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storage StorageController
	err = json.NewDecoder(resp.Body).Decode(&storage)
	if err != nil {
		return nil, err
	}

	storage.SetClient(c)
	return &storage, nil
}

// ListReferencedStorageControllers gets the collection of StorageControllers
// from a provided reference.
func ListReferencedStorageControllers(c common.Client, link string) ([]*StorageController, error) {
	var result []*StorageController
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, storageLink := range links.ItemLinks {
		storage, err := GetStorageController(c, storageLink)
		if err != nil {
			return result, err
		}
		result = append(result, storage)
	}

	return result, nil
}

// Assembly gets the storage controller's assembly.
func (storagecontroller *StorageController) Assembly() (*Assembly, error) {
	if storagecontroller.assembly == "" {
		return nil, nil
	}
	return GetAssembly(storagecontroller.Client, storagecontroller.assembly)
}

// Endpoints gets the storage controller's endpoints.
func (storagecontroller *StorageController) Endpoints() ([]*Endpoint, error) {
	var result []*Endpoint
	for _, endpointLink := range storagecontroller.endpoints {
		endpoint, err := GetEndpoint(storagecontroller.Client, endpointLink)
		if err != nil {
			return result, err
		}
		result = append(result, endpoint)
	}
	return result, nil
}
