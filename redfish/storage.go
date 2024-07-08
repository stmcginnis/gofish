//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type AutoVolumeCreate string

const (
	// DisabledAutoVolumeCreate Do not automatically create volumes.
	DisabledAutoVolumeCreate AutoVolumeCreate = "Disabled"
	// NonRAIDAutoVolumeCreate Automatically create non-RAID volumes.
	NonRAIDAutoVolumeCreate AutoVolumeCreate = "NonRAID"
	// RAID0AutoVolumeCreate Automatically create RAID0 volumes.
	RAID0AutoVolumeCreate AutoVolumeCreate = "RAID0"
	// RAID1AutoVolumeCreate Automatically create RAID1 volumes.
	RAID1AutoVolumeCreate AutoVolumeCreate = "RAID1"
)

type EncryptionMode string

const (
	// DisabledEncryptionMode Encryption is disabled on the storage subsystem.
	DisabledEncryptionMode EncryptionMode = "Disabled"
	// UseExternalKeyEncryptionMode The storage subsystem uses one or more external keys for encryption.
	UseExternalKeyEncryptionMode EncryptionMode = "UseExternalKey"
	// UseLocalKeyEncryptionMode The storage subsystem uses a local key for encryption.
	UseLocalKeyEncryptionMode EncryptionMode = "UseLocalKey"
)

type HotspareActivationPolicy string

const (
	// OnDriveFailureHotspareActivationPolicy The hot spare drive will take over for the original drive when the
	// original drive has been marked as failed by the storage domain.
	OnDriveFailureHotspareActivationPolicy HotspareActivationPolicy = "OnDriveFailure"
	// OnDrivePredictedFailureHotspareActivationPolicy The hot spare drive will take over for the original drive when
	// the original drive has been predicted to fail in the future by the storage domain.
	OnDrivePredictedFailureHotspareActivationPolicy HotspareActivationPolicy = "OnDrivePredictedFailure"
	// OEMHotspareActivationPolicy The hot spare drive will take over for the original drive in an algorithm custom to
	// the OEM.
	OEMHotspareActivationPolicy HotspareActivationPolicy = "OEM"
)

type StorageResetToDefaultsType string

const (
	// ResetAllStorageResetToDefaultsType Reset all settings to factory defaults and remove all volumes.
	ResetAllStorageResetToDefaultsType StorageResetToDefaultsType = "ResetAll"
	// PreserveVolumesStorageResetToDefaultsType Reset all settings to factory defaults but preserve the configured volumes on
	// the controllers.
	PreserveVolumesStorageResetToDefaultsType StorageResetToDefaultsType = "PreserveVolumes"
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

// Rates shall contain all the rate settings available on the controller.
type Rates struct {
	// ConsistencyCheckRatePercent shall contain the percentage of controller resources used for checking data
	// consistency on volumes.
	ConsistencyCheckRatePercent int
	// RebuildRatePercent shall contain the percentage of controller resources used for rebuilding volumes.
	RebuildRatePercent int
	// TransformationRatePercent shall contain the percentage of controller resources used for transforming volumes.
	TransformationRatePercent int
}

// Storage is used to represent resources that represent a storage
// subsystem in the Redfish specification.
type Storage struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AutoVolumeCreate shall indicate if volumes are created automatically for each unassigned drive attached to this
	// storage subsystem.
	AutoVolumeCreate AutoVolumeCreate
	// Connections shall contain a link to a resource collection of type ConnectionCollection. The members of this
	// collection shall reference Connection resources subordinate to Fabric resources.
	connections string
	// ConsistencyGroups shall contain a link to a resource collection of type ConsistencyGroupCollection. The property
	// shall be used when groups of volumes are treated as a single resource by an application or set of applications.
	consistencyGroups string
	// Controllers shall contain a link to a resource collection of type StorageControllerCollection that contains the
	// set of storage controllers allocated to this storage subsystem.
	controllers string
	// Description provides a description of this resource.
	Description string
	// Drives is a collection that indicates all the drives attached to the
	// storage controllers that this resource represents.
	drives []string
	// DrivesCount is the number of drives.
	DrivesCount int `json:"Drives@odata.count"`
	// EncryptionMode shall contain the encryption mode of this storage subsystem.
	EncryptionMode EncryptionMode
	// EndpointGroups shall contain a link to a resource collection of type EndpointGroupCollection. This property
	// shall be implemented when atomic control is needed to perform mapping, masking, and zoning operations.
	endpointGroups string
	// FileSystems shall contain a link to a resource collection of type FileSystemCollection. This property shall be
	// used when file systems are shared or exported by the storage subsystem.
	fileSystems string
	// HotspareActivationPolicy shall contain the policy under which all drives operating as hot spares in this storage
	// domain will activate.
	HotspareActivationPolicy HotspareActivationPolicy
	// Identifiers shall contain a list of all known durable names for the storage subsystem.
	Identifiers []common.Identifier
	// LocalEncryptionKeyIdentifier shall contain the local encryption key identifier used by the storage subsystem
	// when EncryptionMode contains 'UseLocalKey'.
	LocalEncryptionKeyIdentifier string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Redundancy shall contain redundancy information for the storage subsystem.
	Redundancy []Redundancy
	// RedundancyCount is the number of Redundancy objects.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// StorageControllers is a collection that indicates all the storage
	// controllers that this resource represents.
	// This property has been deprecated in favor of Controllers to allow for storage controllers to be represented as their own resources.
	StorageControllers []StorageController
	// StorageControllersCount is the number of storage controllers.
	StorageControllersCount int `json:"StorageControllers@odata.count"`
	// StoragePools shall contain a link to a resource collection of type StoragePoolCollection. This property shall be
	// used when an abstraction of media, rather than references to individual media, are used as the storage data
	// source.
	storagePools string
	// Volumes is a collection that indicates all the volumes produced by the
	// storage controllers that this resource represents.
	volumes string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	// Enclosures shall reference a resource of type Chassis that represents the
	// physical containers attached to this resource.
	enclosures []string
	// EnclosuresCount is the number of enclosures.
	EnclosuresCount       int
	hostingStorageSystems []string
	// HostingStorageSystemsCount is the number of hosting storage systems.
	HostingStorageSystemsCount int
	// NVMeoFDiscoverySubsystems shall contain an array of links to resources of type Storage that represent the
	// discovery subsystems that discovered this subsystem in an NVMe-oF environment.
	nvmeoFDiscoverySubsystems []string
	// NVMeoFDiscoverySubsystemsCount is the number of NVMeoFDiscoverySubsystems.
	NVMeoFDiscoverySubsystemsCount int
	// SimpleStorage shall contain a link to a resource of type SimpleStorage that represents the same storage
	// subsystem as this resource.
	simpleStorage string
	// StorageServices shall contain an array of links to resources of type StorageService with which this storage
	// subsystem is associated.
	storageServices []string
	// StorageServicesCount is the number of storage services.
	StorageServicesCount int

	// setEncryptionKeyTarget is the URL to send SetEncryptionKey requests.
	setEncryptionKeyTarget string
	// Oem are all OEM data under top level storage section
	Oem json.RawMessage

	resetToDefaultsTarget string
}

// UnmarshalJSON unmarshals a Storage object from the raw JSON.
func (storage *Storage) UnmarshalJSON(b []byte) error {
	type temp Storage
	type links struct {
		Enclosures                     common.Links
		EnclosuresCount                int `json:"Enclosures@odata.count"`
		HostingStorageSystems          common.Links
		HostingStorageSystemsCount     int `json:"HostingStorageSystems@odata.count"`
		NVMeoFDiscoverySubsystems      common.Links
		NVMeoFDiscoverySubsystemsCount int `json:"NVMeoFDiscoverySubsystems@odata.count"`
		SimpleStorage                  common.Link
		StorageServices                common.Links
		StorageServicesCount           int `json:"StorageServices@odata.count"`
	}
	type actions struct {
		ResetToDefaults  common.ActionTarget `json:"#/definitions/ResetToDefaults"`
		SetEncryptionKey common.ActionTarget `json:"#Storage.SetEncryptionKey"`
	}
	var t struct {
		temp
		Links             links
		Connections       common.Link
		ConsistencyGroups common.Link
		Controllers       common.Link
		Drives            common.Links
		EndpointGroups    common.Link
		FileSystems       common.Link
		StoragePools      common.Link
		Volumes           common.Link
		Actions           actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storage = Storage(t.temp)

	// Extract the links to other entities for later
	storage.connections = t.Connections.String()
	storage.consistencyGroups = t.ConsistencyGroups.String()
	storage.controllers = t.Controllers.String()
	storage.drives = t.Drives.ToStrings()
	storage.endpointGroups = t.EndpointGroups.String()
	storage.fileSystems = t.FileSystems.String()
	storage.storagePools = t.StoragePools.String()
	storage.volumes = t.Volumes.String()

	storage.enclosures = t.Links.Enclosures.ToStrings()
	storage.EnclosuresCount = t.Links.EnclosuresCount
	storage.hostingStorageSystems = t.Links.HostingStorageSystems.ToStrings()
	storage.HostingStorageSystemsCount = t.Links.HostingStorageSystemsCount
	storage.nvmeoFDiscoverySubsystems = t.Links.NVMeoFDiscoverySubsystems.ToStrings()
	storage.NVMeoFDiscoverySubsystemsCount = t.Links.NVMeoFDiscoverySubsystemsCount
	storage.simpleStorage = t.Links.SimpleStorage.String()
	storage.storageServices = t.Links.StorageServices.ToStrings()
	storage.StorageServicesCount = t.Links.StorageServicesCount

	storage.resetToDefaultsTarget = t.Actions.ResetToDefaults.Target
	storage.setEncryptionKeyTarget = t.Actions.SetEncryptionKey.Target

	// This is a read/write object, so we need to save the raw object data for later
	storage.rawData = b

	return nil
}

// Connection gets the connections that this storage subsystem contains.
func (storage *Storage) Connections() ([]*Connection, error) {
	return ListReferencedConnections(storage.GetClient(), storage.connections)
}

// ConsistencyGroups gets groups of volumes that are treated as a single resource
// by an application or set of applications.
// func (storage *Storage) ConsistencyGroups() ([]*swordfish.ConsistencyGroup, error) {
//	return swordfishListReferencedConsistencyGroups(storage.GetClient(), storage.consistencyGroups)
// }

// Controllers gets the set of storage controllers allocated to this storage subsystem.
func (storage *Storage) Controllers() ([]*StorageController, error) {
	return ListReferencedStorageControllers(storage.GetClient(), storage.controllers)
}

// Drives gets the drives attached to the storage controllers that this
// resource represents.
func (storage *Storage) Drives() ([]*Drive, error) {
	return common.GetObjects[Drive](storage.GetClient(), storage.drives)
}

// EndpointGroups gets the set of endpoints that are used for a common purpose such as an ACL
// or logical identification, that belong to this storage subsystem.
func (storage *Storage) EndpointGroups() ([]*EndpointGroup, error) {
	return ListReferencedEndpointGroups(storage.GetClient(), storage.endpointGroups)
}

// FileSystems gets the file systems that are allocated by this storage subsystem.
// func (storage *Storage) FileSystems() ([]*swordfish.FileSystem, error) {
// return swordfish.ListReferencedFileSystems(storage.GetClient(), storage.fileSystems)
// }

// Volumes gets the volumes associated with this storage subsystem.
func (storage *Storage) Volumes() ([]*Volume, error) {
	return ListReferencedVolumes(storage.GetClient(), storage.volumes)
}

// Enclosures gets the physical containers attached to this resource.
func (storage *Storage) Enclosures() ([]*Chassis, error) {
	return common.GetObjects[Chassis](storage.GetClient(), storage.enclosures)
}

// HostingStorageSystems gets the storage systems that host this storage subsystem.
func (storage *Storage) HostingStorageSystems() ([]*ComputerSystem, error) {
	return common.GetObjects[ComputerSystem](storage.GetClient(), storage.hostingStorageSystems)
}

// NVMeoFDiscoverySubsystems gets the discovery subsystems that discovered this subsystem in an NVMe-oF environment.
func (storage *Storage) NVMeoFDiscoverySubsystems() ([]*Storage, error) {
	return common.GetObjects[Storage](storage.GetClient(), storage.nvmeoFDiscoverySubsystems)
}

// SimpleStorage gets the simple storage instance that corresponds to this storage.
func (storage *Storage) SimpleStorage() (*SimpleStorage, error) {
	if storage.simpleStorage == "" {
		return nil, nil
	}
	return GetSimpleStorage(storage.GetClient(), storage.simpleStorage)
}

// // StorageServices gets the storage services that connect to this storage subsystem.
// func (storage *Storage) StorageServices() ([]*swordfish.StorageService, error) {
//  	return common.GetObjects[StorageService](storage.GetClient(), storage.storageServices)
// }

// ResetToDefaults resets the storage device to factory defaults. This can cause the loss of data.
func (storage *Storage) ResetToDefaults(resetType StorageResetToDefaultsType) error {
	t := struct {
		ResetType StorageResetToDefaultsType
	}{ResetType: resetType}

	return storage.Post(storage.resetToDefaultsTarget, t)
}

// SetEncryptionKey shall set the encryption key for the storage subsystem.
//
// `key` is the local encryption key to set on the storage subsystem.
//
// `currentEncryptionKey` (optional since v1.14.0) is the current local encryption key
// on the storage subsystem.
// Services may reject the action request if this parameter is not provided.
//
// `encryptionKeyIdentifier` (optional) is the local encryption key identifier used by the storage subsystem.
func (storage *Storage) SetEncryptionKey(key, currentEncryptionKey, encryptionKeyIdentifier string) error {
	t := struct {
		CurrentEncryptionKey    string `json:",omitempty"`
		EncryptionKey           string
		EncryptionKeyIdentifier string `json:",omitempty"`
	}{
		CurrentEncryptionKey:    currentEncryptionKey,
		EncryptionKey:           key,
		EncryptionKeyIdentifier: encryptionKeyIdentifier,
	}

	return storage.Post(storage.setEncryptionKeyTarget, t)
}

// Update commits updates to this object's properties to the running system.
func (storage *Storage) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Storage)
	original.UnmarshalJSON(storage.rawData)

	readWriteFields := []string{
		"AutoVolumeCreate",
		"EncryptionMode",
		"HotspareActivationPolicy",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(storage).Elem()

	return storage.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetStorage will get a Storage instance from the service.
func GetStorage(c common.Client, uri string) (*Storage, error) {
	return common.GetObject[Storage](c, uri)
}

// ListReferencedStorages gets the collection of Storage from a provided
// reference.
func ListReferencedStorages(c common.Client, link string) ([]*Storage, error) {
	return common.GetCollectionObjects[Storage](c, link)
}

// GetOperationApplyTimeValues returns the OperationApplyTime values applicable for this storage
func (storage *Storage) GetOperationApplyTimeValues() ([]common.OperationApplyTime, error) {
	return AllowedVolumesUpdateApplyTimes(storage.GetClient(), storage.volumes)
}
