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

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
	"github.com/stmcginnis/gofish/school/redfish"
)

// DefaultStorageServicePath is the default URI for the StorageService
// object.
const DefaultStorageServicePath = "/redfish/v1/StorageService"

// SSLinks is
type SSLinks struct {
	// HostingSystem shall reference the ComputerSystem or
	// StorageController that hosts this service.
	HostingSystem common.Link
}

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
	ODataEtag string `json:"@odata.etag"`
	// ODataId is
	ODataID string `json:"@odata.id"`
	// ODataType is
	ODataType string `json:"@odata.type"`
	// ClassesOfService shall reference a ClassOfService supported by this service.
	classesOfService string
	// DataProtectionLoSCapabilities shall reference the data
	// protection capabilities of this service.
	dataProtectionLoSCapabilities string
	// DataSecurityLoSCapabilities shall reference the data
	// security capabilities of this service.
	dataSecurityLoSCapabilities string
	// DataStorageLoSCapabilities shall reference the data
	// storage capabilities of this service.
	dataStorageLoSCapabilities string
	// DefaultClassOfService, if present, shall reference the
	// default class of service for entities allocated by this storage
	// service. This default may be overridden by the DefaultClassOfService
	// property values within contained StoragePools.
	defaultClassOfService string
	// Description is a description for this StorageService.
	Description string
	// Drives is a collection that indicates all the drives managed by this
	// storage service.
	drives string
	// EndpointGroups shall reference an EndpointGroup.
	endpointGroups string
	// Endpoints shall reference an Endpoint managed by this service.
	endpoints string
	// FileSystems is an array of references to FileSystems managed by this
	// storage service.
	fileSystems string
	// HostingSystem shall reference the ComputerSystem or
	// StorageController that hosts this service.
	hostingSystem string
	// IOConnectivityLoSCapabilities shall reference the IO connectivity
	// capabilities of this service.
	ioConnectivityLoSCapabilities string
	// IOPerformanceLoSCapabilities shall reference the IO performance
	// capabilities of this service.
	ioPerformanceLoSCapabilities string
	// IOStatistics shall represent IO statistics for this StorageService.
	ioStatistics string
	// ID is the instance identifier.
	ID string `json:"Id"`
	// Identifier identifies this resource. The value shall be
	// unique within the managed ecosystem.
	Identifier common.Identifier
	// Name is the instance name.
	Name string
	// Redundancy collection shall contain the redundancy information
	// for the storage subsystem.
	redundancy string
	// RedundancyCount is the Redundancy collection object count.
	RedundancyCount int `json:"Redundancy@odata.count"`
	// SpareResourceSets shall contain resources that may be utilized to
	// replace the capacity provided by a failed resource having a compatible type.
	spareResourceSets string
	// SpareResourceSetsCount is the number of SpareResourceSet objects.
	SpareResourceSetsCount int `json:"SpareResourceSets@odata.count"`
	// Status is the StorageService status.
	Status common.Status
	// StorageGroups shall reference a StorageGroup.
	storageGroups string
	// StoragePools is an array of references to StoragePools.
	storagePools string
	// StorageSubsystems shall be a link to a collection of type
	// StorageCollection having members that represent storage subsystems
	// managed by this storage service.
	storageSubsystems string
	// Volumes is an array of references to Volumes managed by this storage
	// service.
	volumes string
}

// UnmarshalJSON unmarshals a StorageService object from the raw JSON.
func (storageservice *StorageService) UnmarshalJSON(b []byte) error {
	type temp StorageService
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
		IOStatistics                  common.Link
		Redundancy                    common.Link
		SpareResourceSets             common.Link
		StorageGroups                 common.Link
		StoragePools                  common.Link
		StorageSubsystems             common.Link
		Links                         SSLinks
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storageservice = StorageService(t.temp)

	// Extract the links to other entities for later
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
	storageservice.ioStatistics = string(t.IOStatistics)
	storageservice.redundancy = string(t.Redundancy)
	storageservice.spareResourceSets = string(t.SpareResourceSets)
	storageservice.storageGroups = string(t.StorageGroups)
	storageservice.storagePools = string(t.StoragePools)
	storageservice.storageSubsystems = string(t.StorageSubsystems)
	storageservice.hostingSystem = string(t.Links.HostingSystem)

	return nil
}

// GetStorageService will get a StorageService instance from the service.
func GetStorageService(c common.Client, uri string) (*StorageService, error) {
	resp, err := c.Get(uri)
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

// ListStorageServices gets all StorageService in the system.
func ListStorageServices(c common.Client) ([]*StorageService, error) {
	return ListReferencedStorageServices(c, DefaultStorageServicePath)
}

// ClassesOfService gets the classes of service supported by this storage.
func (storageservice *StorageService) ClassesOfService() (*ClassesOfService, error) {
	if storageservice.classesOfService == "" {
		return nil, nil
	}

	resp, err := storageservice.Client.Get(storageservice.classesOfService)
	defer resp.Body.Close()

	var classofservice ClassesOfService
	err = json.NewDecoder(resp.Body).Decode(&classofservice)
	if err != nil {
		return nil, err
	}

	return &classofservice, nil
}

// Endpoints gets the storage service's endpoints.
func (storageservice *StorageService) Endpoints() ([]*redfish.Endpoint, error) {
	return redfish.ListReferencedEndpoints(storageservice.Client, storageservice.endpoints)
}

// EndpointGroups gets the storage service's endpoint groups.
func (storageservice *StorageService) EndpointGroups() ([]*EndpointGroup, error) {
	return ListReferencedEndpointGroups(storageservice.Client, storageservice.endpointGroups)
}

// DataProtectionLoSCapabilities gets the storage service's data protection
// capabilities.
func (storageservice *StorageService) DataProtectionLoSCapabilities() ([]*DataProtectionLoSCapabilities, error) {
	return ListReferencedDataProtectionLoSCapabilities(storageservice.Client, storageservice.dataProtectionLoSCapabilities)
}
