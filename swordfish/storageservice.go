//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"reflect"

	"github.com/coreweave/gofish/common"
	"github.com/coreweave/gofish/redfish"
)

// StorageService is a collection of resources that the system can make
// available to one or more host systems. The collection can contain:
// block, file, or object storage; local system access points through
// which the collection is made available; hosts, or host access points
// to which the collection is made available.
type StorageService struct {
	common.Entity

	// ODataContext is
	ODataContext string `json:"@odata.context"`
	// ODataEtag is
	// ODataId is
	// ODataType is
	ODataType string `json:"@odata.type"`
	// Description is a description for this StorageService.
	Description string
	// Identifier identifies this resource. The value shall be
	// unique within the managed ecosystem.
	Identifier common.Identifier
	// RedundancyCount is the Redundancy collection object count.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// SpareResourceSetsCount is the number of SpareResourceSet objects.
	SpareResourceSetsCount int `json:"SpareResourceSets@odata.count"`
	// Status is the StorageService status.
	Status common.Status
	// classesOfService shall reference a ClassOfService supported by this service.
	classesOfService string
	// dataProtectionLoSCapabilities shall reference the data
	// protection capabilities of this service.
	dataProtectionLoSCapabilities string
	// dataSecurityLoSCapabilities shall reference the data
	// security capabilities of this service.
	dataSecurityLoSCapabilities string
	// dataStorageLoSCapabilities shall reference the data
	// storage capabilities of this service.
	dataStorageLoSCapabilities string
	// DefaultClassOfService, if present, shall reference the
	// default class of service for entities allocated by this storage
	// service. This default may be overridden by the DefaultClassOfService
	// property values within contained StoragePools.
	defaultClassOfService string
	// drives is a collection that indicates all the drives managed by this
	// storage service.
	drives string
	// endpointGroups shall reference a collection of EndpointGroups.
	endpointGroups string
	// endpoints shall reference a collection of Endpoints managed by this service.
	endpoints string
	// fileSystems is an array of references to FileSystems managed by this
	// storage service.
	fileSystems string
	// ioConnectivityLoSCapabilities shall reference the IO connectivity
	// capabilities of this service.
	ioConnectivityLoSCapabilities string
	// ioPerformanceLoSCapabilities shall reference the IO performance
	// capabilities of this service.
	ioPerformanceLoSCapabilities string
	// IOStatistics shall represent IO statistics for this StorageService.
	IOStatistics IOStatistics
	// LinesOfService shall reference a LineOfService collection defined for this service.
	linesOfService []string
	// LinesOfServiceCount is the number of lines of service.
	LinesOfServiceCount int `json:"LinesOfService@odata.count"`
	// Metrics shall contain a link to a resource of type StorageServiceMetrics that specifies the metrics for this
	// storage service. IO metrics are reported in the IOStatistics property.
	metrics string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// spareResourceSets shall contain resources that may be utilized to
	// replace the capacity provided by a failed resource having a compatible type.
	spareResourceSets []string
	// storageGroups are the associated storage groups for this service.
	storageGroups string
	// StoragePools is an array of references to StoragePools.
	storagePools string
	// StorageSubsystems shall be a link to a collection of type
	// StorageCollection having members that represent storage subsystems
	// managed by this storage service.
	storageSubsystems string
	// Redundancy collection shall contain the redundancy information
	// for the storage subsystem.
	redundancy []string
	// Volumes is an array of references to Volumes managed by this storage
	// service.
	volumes string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	// setEncryptionKeyTarget is the URL to send SetEncryptionKey requests.
	setEncryptionKeyTarget string

	// hostingSystem shall reference the ComputerSystem or
	// StorageController that hosts this service.
	hostingSystem string
}

// UnmarshalJSON unmarshals a StorageService object from the raw JSON.
func (storageservice *StorageService) UnmarshalJSON(b []byte) error {
	type temp StorageService
	type links struct {
		// HostingSystem shall reference the ComputerSystem or
		// StorageController that hosts this service.
		HostingSystem common.Link
	}
	type actions struct {
		SetEncryptionKey common.ActionTarget `json:"#StorageService.SetEncryptionKey"`
	}
	var t struct {
		temp
		ClassesOfService              common.Link
		DataProtectionLoSCapabilities common.Link
		DataSecurityLoSCapabilities   common.Link
		DataStorageLoSCapabilities    common.Link
		DefaultClassOfService         common.Link
		Drives                        common.Link
		EndpointGroups                common.Link
		Endpoints                     common.Link
		FileSystems                   common.Link
		IOConnectivityLoSCapabilities common.Link
		IOPerformanceLoSCapabilities  common.Link
		Redundancy                    common.Links
		LinesOfService                common.Links
		Metrics                       common.Link
		SpareResourceSets             common.Links
		StorageGroups                 common.Link
		StoragePools                  common.Link
		StorageSubsystems             common.Link
		Volumes                       common.Link
		Links                         links
		Actions                       actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*storageservice = StorageService(t.temp)
	storageservice.classesOfService = t.ClassesOfService.String()
	storageservice.dataProtectionLoSCapabilities = t.DataProtectionLoSCapabilities.String()
	storageservice.dataSecurityLoSCapabilities = t.DataSecurityLoSCapabilities.String()
	storageservice.dataStorageLoSCapabilities = t.DataStorageLoSCapabilities.String()
	storageservice.defaultClassOfService = t.DefaultClassOfService.String()
	storageservice.drives = t.Drives.String()
	storageservice.endpointGroups = t.EndpointGroups.String()
	storageservice.endpoints = t.Endpoints.String()
	storageservice.fileSystems = t.FileSystems.String()
	storageservice.ioConnectivityLoSCapabilities = t.IOConnectivityLoSCapabilities.String()
	storageservice.ioPerformanceLoSCapabilities = t.IOPerformanceLoSCapabilities.String()
	storageservice.metrics = t.Metrics.String()
	storageservice.redundancy = t.Redundancy.ToStrings()
	storageservice.linesOfService = t.LinesOfService.ToStrings()
	storageservice.spareResourceSets = t.SpareResourceSets.ToStrings()
	storageservice.storageGroups = t.StorageGroups.String()
	storageservice.storagePools = t.StoragePools.String()
	storageservice.storageSubsystems = t.StorageSubsystems.String()
	storageservice.volumes = t.Volumes.String()

	storageservice.setEncryptionKeyTarget = t.Actions.SetEncryptionKey.Target

	storageservice.hostingSystem = t.Links.HostingSystem.String()

	// This is a read/write object, so we need to save the raw object data for later
	storageservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (storageservice *StorageService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(StorageService)
	original.UnmarshalJSON(storageservice.rawData)

	readWriteFields := []string{
		"ClassesOfService",
		"ConsistencyGroups",
		"DataProtectionLoSCapabilities",
		"DataSecurityLoSCapabilities",
		"DataStorageLoSCapabilities",
		"DefaultClassOfService",
		"EndpointGroups",
		"FileSystems",
		"IOConnectivityLoSCapabilities",
		"IOPerformanceLoSCapabilities",
		"LinesOfService",
		"SpareResourceSets",
		"Volumes",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(storageservice).Elem()

	return storageservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetStorageService will get a StorageService instance from the service.
func GetStorageService(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*StorageService, error) {
	return common.GetObject[StorageService](c, uri, queryOpts...)
}

// ListReferencedStorageServices gets the collection of StorageService from
// a provided reference.
func ListReferencedStorageServices(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*StorageService, error) {
	return common.GetCollectionObjects[StorageService](c, link, queryOpts...)
}

// ClassesOfService gets the storage service's classes of service.
func (storageservice *StorageService) ClassesOfService(queryOpts ...common.QueryGroupOption) ([]*ClassOfService, error) {
	return ListReferencedClassOfServices(storageservice.GetClient(), storageservice.classesOfService, queryOpts...)
}

// DataProtectionLoSCapabilities gets the storage service's data protection
// capabilities.
func (storageservice *StorageService) DataProtectionLoSCapabilities(queryOpts ...common.QueryGroupOption) (*DataProtectionLoSCapabilities, error) {
	if storageservice.dataProtectionLoSCapabilities == "" {
		return nil, nil
	}
	return GetDataProtectionLoSCapabilities(storageservice.GetClient(), storageservice.dataProtectionLoSCapabilities, queryOpts...)
}

// DataSecurityLoSCapabilities gets the storage service's data security
// capabilities.
func (storageservice *StorageService) DataSecurityLoSCapabilities(queryOpts ...common.QueryGroupOption) (*DataSecurityLoSCapabilities, error) {
	if storageservice.dataSecurityLoSCapabilities == "" {
		return nil, nil
	}
	return GetDataSecurityLoSCapabilities(storageservice.GetClient(), storageservice.dataSecurityLoSCapabilities, queryOpts...)
}

// DataStorageLoSCapabilities references the data storage capabilities of this service.
func (storageservice *StorageService) DataStorageLoSCapabilities(queryOpts ...common.QueryGroupOption) (*DataStorageLoSCapabilities, error) {
	if storageservice.dataStorageLoSCapabilities == "" {
		return nil, nil
	}
	return GetDataStorageLoSCapabilities(storageservice.GetClient(), storageservice.dataStorageLoSCapabilities, queryOpts...)
}

// DefaultClassOfService references the default class of service for entities
// allocated by this storage service.
func (storageservice *StorageService) DefaultClassOfService(queryOpts ...common.QueryGroupOption) (*ClassOfService, error) {
	if storageservice.defaultClassOfService == "" {
		return nil, nil
	}
	return GetClassOfService(storageservice.GetClient(), storageservice.defaultClassOfService, queryOpts...)
}

// Drives gets the storage service's drives.
func (storageservice *StorageService) Drives(queryOpts ...common.QueryGroupOption) ([]*redfish.Drive, error) {
	return redfish.ListReferencedDrives(storageservice.GetClient(), storageservice.drives, queryOpts...)
}

// EndpointGroups gets the storage service's endpoint groups.
func (storageservice *StorageService) EndpointGroups(queryOpts ...common.QueryGroupOption) ([]*EndpointGroup, error) {
	return ListReferencedEndpointGroups(storageservice.GetClient(), storageservice.endpointGroups, queryOpts...)
}

// Endpoints gets the storage service's endpoints.
func (storageservice *StorageService) Endpoints(queryOpts ...common.QueryGroupOption) ([]*redfish.Endpoint, error) {
	return redfish.ListReferencedEndpoints(storageservice.GetClient(), storageservice.endpoints, queryOpts...)
}

// FileSystems gets all filesystems available through this storage service.
func (storageservice *StorageService) FileSystems(queryOpts ...common.QueryGroupOption) ([]*FileSystem, error) {
	return ListReferencedFileSystems(storageservice.GetClient(), storageservice.fileSystems, queryOpts...)
}

// IOConnectivityLoSCapabilities references the IO connectivity capabilities of this service.
func (storageservice *StorageService) IOConnectivityLoSCapabilities(queryOpts ...common.QueryGroupOption) (*IOConnectivityLoSCapabilities, error) {
	if storageservice.ioConnectivityLoSCapabilities == "" {
		return nil, nil
	}
	return GetIOConnectivityLoSCapabilities(storageservice.GetClient(), storageservice.ioConnectivityLoSCapabilities, queryOpts...)
}

// IOPerformanceLoSCapabilities references the IO performance capabilities of this service.
func (storageservice *StorageService) IOPerformanceLoSCapabilities(queryOpts ...common.QueryGroupOption) (*IOPerformanceLoSCapabilities, error) {
	if storageservice.ioConnectivityLoSCapabilities == "" {
		return nil, nil
	}
	return GetIOPerformanceLoSCapabilities(storageservice.GetClient(), storageservice.ioPerformanceLoSCapabilities, queryOpts...)
}

// Redundancy gets the redundancy information for the storage subsystem.
func (storageservice *StorageService) Redundancy(queryOpts ...common.QueryGroupOption) ([]*redfish.Redundancy, error) {
	return common.GetObjects[redfish.Redundancy](storageservice.GetClient(), storageservice.redundancy, queryOpts...)
}

// LinesOfService gets lines of service for this service.
func (storageservice *StorageService) LinesOfService(queryOpts ...common.QueryGroupOption) ([]*LineOfService, error) {
	return common.GetObjects[LineOfService](storageservice.GetClient(), storageservice.linesOfService, queryOpts...)
}

// SpareResourceSets gets resources that may be utilized to replace the capacity
// provided by a failed resource having a compatible type.
func (storageservice *StorageService) SpareResourceSets(queryOpts ...common.QueryGroupOption) ([]*SpareResourceSet, error) {
	return common.GetObjects[SpareResourceSet](storageservice.GetClient(), storageservice.spareResourceSets, queryOpts...)
}

// StorageGroups gets the storage groups that are a part of this storage service.
func (storageservice *StorageService) StorageGroups(queryOpts ...common.QueryGroupOption) ([]*StorageGroup, error) {
	return common.GetCollectionObjects[StorageGroup](storageservice.GetClient(), storageservice.storageGroups, queryOpts...)
}

// Volumes gets the volumes that are a part of this storage service.
func (storageservice *StorageService) Volumes(queryOpts ...common.QueryGroupOption) ([]*Volume, error) {
	return ListReferencedVolumes(storageservice.GetClient(), storageservice.volumes, queryOpts...)
}

// SetEncryptionKey shall set the encryption key for the storage subsystem.
func (storageservice *StorageService) SetEncryptionKey(key string) error {
	t := struct {
		EncryptionKey string
	}{EncryptionKey: key}

	return storageservice.Post(storageservice.setEncryptionKeyTarget, t)
}

// Metrics gets the metrics for this storage pool.
func (storageservice *StorageService) Metrics(queryOpts ...common.QueryGroupOption) (*StorageServiceMetrics, error) {
	if storageservice.metrics == "" {
		return nil, nil
	}
	return GetStorageServiceMetrics(storageservice.GetClient(), storageservice.metrics, queryOpts...)
}
