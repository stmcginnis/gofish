//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
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
	Redundancy []Redundancy
	// RedundancyCount is the number of Redundancy objects.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// StorageControllers is a collection that indicates all the storage
	// controllers that this resource represents.
	StorageControllers []StorageController
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
	// setEncryptionKeyTarget is the URL to send SetEncryptionKey requests.
	setEncryptionKeyTarget string
}

// UnmarshalJSON unmarshals a Storage object from the raw JSON.
func (storage *Storage) UnmarshalJSON(b []byte) error {
	type temp Storage
	type links struct {
		Enclosures      common.Links
		EnclosuresCount int `json:"Enclosures@odata.count"`
	}
	type actions struct {
		SetEncryptionKey struct {
			Target string
		} `json:"#Storage.SetEncryptionKey"`
	}
	var t struct {
		temp
		Links   links
		Drives  common.Links
		Volumes common.Link
		Actions actions
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
	storage.volumes = string(t.Volumes)
	storage.setEncryptionKeyTarget = t.Actions.SetEncryptionKey.Target

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

// Volumes gets the volumes associated with this storage subsystem.
func (storage *Storage) Volumes() ([]*Volume, error) {
	return ListReferencedVolumes(storage.Client, storage.volumes)
}

// SetEncryptionKey shall set the encryption key for the storage subsystem.
func (storage *Storage) SetEncryptionKey(key string) error {
	type temp struct {
		EncryptionKey string
	}
	t := temp{EncryptionKey: key}

	_, err := storage.Client.Post(storage.setEncryptionKeyTarget, t)
	return err
}

// GetOperationApplyTimeValues returns the OperationApplyTime values applicable for this storage
func (storage *Storage) GetOperationApplyTimeValues() ([]common.OperationApplyTime, error) {
	return AllowedVolumesUpdateApplyTimes(storage.Client, storage.volumes)
}

// StorageController is used to represent a resource that represents a
// storage controller in the Redfish specification.
type StorageController struct {
	common.Entity

	// Assembly shall be a link to a resource of type Assembly.
	assembly string
	// AssetTag is used to track the storage controller for inventory
	// purposes.
	AssetTag string
	// CacheSummary shall contain properties which describe the cache memory for
	// the current resource.
	CacheSummary CacheSummary
	// FirmwareVersion shall contain the firmware version as defined by the
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
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a StorageController object from the raw JSON.
func (storagecontroller *StorageController) UnmarshalJSON(b []byte) error { // nolint:dupl
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

	// This is a read/write object, so we need to save the raw object data for later
	storagecontroller.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (storagecontroller *StorageController) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(StorageController)
	err := original.UnmarshalJSON(storagecontroller.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AssetTag",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(storagecontroller).Elem()

	return storagecontroller.Entity.Update(originalElement, currentElement, readWriteFields)
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
