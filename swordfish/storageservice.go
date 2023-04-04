//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
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
	// hostingSystem shall reference the ComputerSystem or
	// StorageController that hosts this service.
	hostingSystem string
	// ioConnectivityLoSCapabilities shall reference the IO connectivity
	// capabilities of this service.
	ioConnectivityLoSCapabilities string
	// ioPerformanceLoSCapabilities shall reference the IO performance
	// capabilities of this service.
	ioPerformanceLoSCapabilities string
	// IOStatistics shall represent IO statistics for this StorageService.
	IOStatistics IOStatistics
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
	// setEncryptionKeyTarget is the URL to send SetEncryptionKey requests.
	setEncryptionKeyTarget string
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
		SetEncryptionKey struct {
			Target string
		} `json:"#StorageService.SetEncryptionKey"`
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
	storageservice.redundancy = t.Redundancy.ToStrings()
	storageservice.spareResourceSets = t.SpareResourceSets.ToStrings()
	storageservice.storageGroups = t.StorageGroups.String()
	storageservice.storagePools = t.StoragePools.String()
	storageservice.storageSubsystems = t.StorageSubsystems.String()
	storageservice.hostingSystem = t.Links.HostingSystem.String()
	storageservice.volumes = t.Volumes.String()
	storageservice.setEncryptionKeyTarget = t.Actions.SetEncryptionKey.Target

	return nil
}

// GetStorageService will get a StorageService instance from the service.
func GetStorageService(c common.Client, uri string) (*StorageService, error) {
	var storageService StorageService
	return &storageService, storageService.Get(c, uri, &storageService)
}

// ListReferencedStorageServices gets the collection of StorageService from
// a provided reference.
func ListReferencedStorageServices(c common.Client, link string) ([]*StorageService, error) { //nolint:dupl
	var result []*StorageService
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *StorageService
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		storageservice, err := GetStorageService(c, link)
		ch <- GetResult{Item: storageservice, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// ClassesOfService gets the storage service's classes of service.
func (storageservice *StorageService) ClassesOfService() ([]*ClassOfService, error) {
	return ListReferencedClassOfServices(storageservice.GetClient(), storageservice.classesOfService)
}

// DataProtectionLoSCapabilities gets the storage service's data protection
// capabilities.
func (storageservice *StorageService) DataProtectionLoSCapabilities() (*DataProtectionLoSCapabilities, error) {
	if storageservice.dataProtectionLoSCapabilities == "" {
		return nil, nil
	}
	return GetDataProtectionLoSCapabilities(storageservice.GetClient(), storageservice.dataProtectionLoSCapabilities)
}

// DataSecurityLoSCapabilities gets the storage service's data security
// capabilities.
func (storageservice *StorageService) DataSecurityLoSCapabilities() (*DataSecurityLoSCapabilities, error) {
	if storageservice.dataSecurityLoSCapabilities == "" {
		return nil, nil
	}
	return GetDataSecurityLoSCapabilities(storageservice.GetClient(), storageservice.dataSecurityLoSCapabilities)
}

// DataStorageLoSCapabilities references the data storage capabilities of this service.
func (storageservice *StorageService) DataStorageLoSCapabilities() (*DataStorageLoSCapabilities, error) {
	if storageservice.dataStorageLoSCapabilities == "" {
		return nil, nil
	}
	return GetDataStorageLoSCapabilities(storageservice.GetClient(), storageservice.dataStorageLoSCapabilities)
}

// DefaultClassOfService references the default class of service for entities
// allocated by this storage service.
func (storageservice *StorageService) DefaultClassOfService() (*ClassOfService, error) {
	if storageservice.defaultClassOfService == "" {
		return nil, nil
	}
	return GetClassOfService(storageservice.GetClient(), storageservice.defaultClassOfService)
}

// Drives gets the storage service's drives.
func (storageservice *StorageService) Drives() ([]*redfish.Drive, error) {
	return redfish.ListReferencedDrives(storageservice.GetClient(), storageservice.drives)
}

// EndpointGroups gets the storage service's endpoint groups.
func (storageservice *StorageService) EndpointGroups() ([]*EndpointGroup, error) {
	return ListReferencedEndpointGroups(storageservice.GetClient(), storageservice.endpointGroups)
}

// Endpoints gets the storage service's endpoints.
func (storageservice *StorageService) Endpoints() ([]*redfish.Endpoint, error) {
	return redfish.ListReferencedEndpoints(storageservice.GetClient(), storageservice.endpoints)
}

// FileSystems gets all filesystems available through this storage service.
func (storageservice *StorageService) FileSystems() ([]*FileSystem, error) {
	return ListReferencedFileSystems(storageservice.GetClient(), storageservice.fileSystems)
}

// IOConnectivityLoSCapabilities references the IO connectivity capabilities of this service.
func (storageservice *StorageService) IOConnectivityLoSCapabilities() (*IOConnectivityLoSCapabilities, error) {
	if storageservice.ioConnectivityLoSCapabilities == "" {
		return nil, nil
	}
	return GetIOConnectivityLoSCapabilities(storageservice.GetClient(), storageservice.ioConnectivityLoSCapabilities)
}

// IOPerformanceLoSCapabilities references the IO performance capabilities of this service.
func (storageservice *StorageService) IOPerformanceLoSCapabilities() (*IOPerformanceLoSCapabilities, error) {
	if storageservice.ioConnectivityLoSCapabilities == "" {
		return nil, nil
	}
	return GetIOPerformanceLoSCapabilities(storageservice.GetClient(), storageservice.ioPerformanceLoSCapabilities)
}

// Redundancy gets the redundancy information for the storage subsystem.
func (storageservice *StorageService) Redundancy() ([]*redfish.Redundancy, error) {
	var result []*redfish.Redundancy

	collectionError := common.NewCollectionError()
	for _, redundancyLink := range storageservice.redundancy {
		redundancy, err := redfish.GetRedundancy(storageservice.GetClient(), redundancyLink)
		if err != nil {
			collectionError.Failures[redundancyLink] = err
		} else {
			result = append(result, redundancy)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// SpareResourceSets gets resources that may be utilized to replace the capacity
// provided by a failed resource having a compatible type.
func (storageservice *StorageService) SpareResourceSets() ([]*SpareResourceSet, error) {
	var result []*SpareResourceSet

	collectionError := common.NewCollectionError()
	for _, srsLink := range storageservice.spareResourceSets {
		srs, err := GetSpareResourceSet(storageservice.GetClient(), srsLink)
		if err != nil {
			collectionError.Failures[srsLink] = err
		} else {
			result = append(result, srs)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// StorageGroups gets the storage groups that are a part of this storage service.
func (storageservice *StorageService) StorageGroups() ([]*StorageGroup, error) {
	var result []*StorageGroup

	collectionError := common.NewCollectionError()
	for _, sgLink := range storageservice.spareResourceSets {
		sg, err := GetStorageGroup(storageservice.GetClient(), sgLink)
		if err != nil {
			collectionError.Failures[sgLink] = err
		} else {
			result = append(result, sg)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// Volumes gets the volumes that are a part of this storage service.
func (storageservice *StorageService) Volumes() ([]*Volume, error) {
	return ListReferencedVolumes(storageservice.GetClient(), storageservice.volumes)
}

// SetEncryptionKey shall set the encryption key for the storage subsystem.
func (storageservice *StorageService) SetEncryptionKey(key string) error {
	t := struct {
		EncryptionKey string
	}{EncryptionKey: key}

	return storageservice.Post(storageservice.setEncryptionKeyTarget, t)
}
