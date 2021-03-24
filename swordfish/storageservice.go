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
	storageservice.classesOfService = string(t.ClassesOfService)
	storageservice.dataProtectionLoSCapabilities = string(t.DataProtectionLoSCapabilities)
	storageservice.dataSecurityLoSCapabilities = string(t.DataSecurityLoSCapabilities)
	storageservice.dataStorageLoSCapabilities = string(t.DataStorageLoSCapabilities)
	storageservice.defaultClassOfService = string(t.DefaultClassOfService)
	storageservice.drives = string(t.Drives)
	storageservice.endpointGroups = string(t.EndpointGroups)
	storageservice.endpoints = string(t.Endpoints)
	storageservice.fileSystems = string(t.FileSystems)
	storageservice.ioConnectivityLoSCapabilities = string(t.IOConnectivityLoSCapabilities)
	storageservice.ioPerformanceLoSCapabilities = string(t.IOPerformanceLoSCapabilities)
	storageservice.redundancy = t.Redundancy.ToStrings()
	storageservice.spareResourceSets = t.SpareResourceSets.ToStrings()
	storageservice.storageGroups = string(t.StorageGroups)
	storageservice.storagePools = string(t.StoragePools)
	storageservice.storageSubsystems = string(t.StorageSubsystems)
	storageservice.hostingSystem = string(t.Links.HostingSystem)
	storageservice.volumes = string(t.Volumes)
	storageservice.setEncryptionKeyTarget = t.Actions.SetEncryptionKey.Target

	return nil
}

// GetStorageService will get a StorageService instance from the service.
func GetStorageService(c common.Client, uri string) (*StorageService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var storageservice StorageService
	err = json.NewDecoder(resp.Body).Decode(&storageservice)
	if err != nil {
		return nil, err
	}

	storageservice.SetClient(c)
	return &storageservice, nil
}

// ListReferencedStorageServices gets the collection of StorageService from
// a provided reference.
func ListReferencedStorageServices(c common.Client, link string) ([]*StorageService, error) {
	var result []*StorageService
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, storageserviceLink := range links.ItemLinks {
		storageservice, err := GetStorageService(c, storageserviceLink)
		if err != nil {
			return result, err
		}
		result = append(result, storageservice)
	}

	return result, nil
}

// ClassesOfService gets the storage service's classes of service.
func (storageservice *StorageService) ClassesOfService() ([]*ClassOfService, error) {
	return ListReferencedClassOfServices(storageservice.Client, storageservice.classesOfService)
}

// DataProtectionLoSCapabilities gets the storage service's data protection
// capabilities.
func (storageservice *StorageService) DataProtectionLoSCapabilities() (*DataProtectionLoSCapabilities, error) {
	if storageservice.dataProtectionLoSCapabilities == "" {
		return nil, nil
	}
	return GetDataProtectionLoSCapabilities(storageservice.Client, storageservice.dataProtectionLoSCapabilities)
}

// DataSecurityLoSCapabilities gets the storage service's data security
// capabilities.
func (storageservice *StorageService) DataSecurityLoSCapabilities() (*DataSecurityLoSCapabilities, error) {
	if storageservice.dataSecurityLoSCapabilities == "" {
		return nil, nil
	}
	return GetDataSecurityLoSCapabilities(storageservice.Client, storageservice.dataSecurityLoSCapabilities)
}

// DataStorageLoSCapabilities references the data storage capabilities of this service.
func (storageservice *StorageService) DataStorageLoSCapabilities() (*DataStorageLoSCapabilities, error) {
	if storageservice.dataStorageLoSCapabilities == "" {
		return nil, nil
	}
	return GetDataStorageLoSCapabilities(storageservice.Client, storageservice.dataStorageLoSCapabilities)
}

// DefaultClassOfService references the default class of service for entities
// allocated by this storage service.
func (storageservice *StorageService) DefaultClassOfService() (*ClassOfService, error) {
	if storageservice.defaultClassOfService == "" {
		return nil, nil
	}
	return GetClassOfService(storageservice.Client, storageservice.defaultClassOfService)
}

// Drives gets the storage service's drives.
func (storageservice *StorageService) Drives() ([]*redfish.Drive, error) {
	return redfish.ListReferencedDrives(storageservice.Client, storageservice.drives)
}

// EndpointGroups gets the storage service's endpoint groups.
func (storageservice *StorageService) EndpointGroups() ([]*EndpointGroup, error) {
	return ListReferencedEndpointGroups(storageservice.Client, storageservice.endpointGroups)
}

// Endpoints gets the storage service's endpoints.
func (storageservice *StorageService) Endpoints() ([]*redfish.Endpoint, error) {
	return redfish.ListReferencedEndpoints(storageservice.Client, storageservice.endpoints)
}

// FileSystems gets all filesystems available through this storage service.
func (storageservice *StorageService) FileSystems() ([]*FileSystem, error) {
	return ListReferencedFileSystems(storageservice.Client, storageservice.fileSystems)
}

// IOConnectivityLoSCapabilities references the IO connectivity capabilities of this service.
func (storageservice *StorageService) IOConnectivityLoSCapabilities() (*IOConnectivityLoSCapabilities, error) {
	if storageservice.ioConnectivityLoSCapabilities == "" {
		return nil, nil
	}
	return GetIOConnectivityLoSCapabilities(storageservice.Client, storageservice.ioConnectivityLoSCapabilities)
}

// IOPerformanceLoSCapabilities references the IO performance capabilities of this service.
func (storageservice *StorageService) IOPerformanceLoSCapabilities() (*IOPerformanceLoSCapabilities, error) {
	if storageservice.ioConnectivityLoSCapabilities == "" {
		return nil, nil
	}
	return GetIOPerformanceLoSCapabilities(storageservice.Client, storageservice.ioPerformanceLoSCapabilities)
}

// Redundancy gets the redundancy information for the storage subsystem.
func (storageservice *StorageService) Redundancy() ([]*redfish.Redundancy, error) {
	var result []*redfish.Redundancy
	for _, redundancyLink := range storageservice.redundancy {
		redundancy, err := redfish.GetRedundancy(storageservice.Client, redundancyLink)
		if err != nil {
			return result, err
		}
		result = append(result, redundancy)
	}

	return result, nil
}

// SpareResourceSets gets resources that may be utilized to replace the capacity
// provided by a failed resource having a compatible type.
func (storageservice *StorageService) SpareResourceSets() ([]*SpareResourceSet, error) {
	var result []*SpareResourceSet
	for _, srsLink := range storageservice.spareResourceSets {
		srs, err := GetSpareResourceSet(storageservice.Client, srsLink)
		if err != nil {
			return result, err
		}
		result = append(result, srs)
	}

	return result, nil
}

// StorageGroups gets the storage groups that are a part of this storage service.
func (storageservice *StorageService) StorageGroups() ([]*StorageGroup, error) {
	var result []*StorageGroup
	for _, sgLink := range storageservice.spareResourceSets {
		sg, err := GetStorageGroup(storageservice.Client, sgLink)
		if err != nil {
			return result, err
		}
		result = append(result, sg)
	}

	return result, nil
}

// Volumes gets the volumes that are a part of this storage service.
func (storageservice *StorageService) Volumes() ([]*Volume, error) {
	return ListReferencedVolumes(storageservice.Client, storageservice.volumes)
}

// SetEncryptionKey shall set the encryption key for the storage subsystem.
func (storageservice *StorageService) SetEncryptionKey(key string) error {
	type temp struct {
		EncryptionKey string
	}
	t := temp{EncryptionKey: key}

	_, err := storageservice.Client.Post(storageservice.setEncryptionKeyTarget, t)
	return err
}
