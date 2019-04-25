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
)

// DefaultStorageServicePath is the default URI for StorageService collections.
const DefaultStorageServicePath = "/redfish/v1/StorageServices"

// StorageService is a Swordfish storage system instance.
type StorageService struct {
	common.Entity
	classesOfService              string
	clientEndPointGroups          string
	Description                   string
	drives                        string
	endpoints                     string
	dataProtectionLOSCapabilities string
	dataSecurityLOSCapabilities   string
	dataStorageLOSCapabilities    string
	enclosures                    string
	hostingSystem                 string
	ioConnectivityLOSCapabilities string
	ioPerformanceLOSCapabilities  string
	serverEndpointGroups          string
	Status                        common.Status
	storageGroups                 string
	storagePools                  string
	storageSubsystems             string
	volumes                       []string
}

// UnmarshalJSON unmarshals a StorageService object from the raw JSON.
func (s *StorageService) UnmarshalJSON(b []byte) error {
	type temp StorageService
	type linkReference struct {
		dataProtectionLOSCapabilities common.Link
		dataSecurityLOSCapabilities   common.Link
		dataStorageLOSCapabilities    common.Link
		enclosures                    common.Link
		hostingSystem                 common.Link
		ioConnectivityLOSCapabilities common.Link
		ioPerformanceLOSCapabilities  common.Link
	}
	var t struct {
		temp
		ClassesOfService     common.Link
		ClientEndPointGroups common.Link
		Drives               common.Link
		Endpoints            common.Link
		ServerEndpointGroups common.Link
		StorageGroups        common.Link
		StoragePools         common.Link
		StorageSubsystems    common.Link
		Volumes              common.LinksCollection
		Links                linkReference
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*s = StorageService(t.temp)

	// Extract the links to other entities
	s.classesOfService = string(t.ClassesOfService)
	s.clientEndPointGroups = string(t.ClientEndPointGroups)
	s.drives = string(t.Drives)
	s.endpoints = string(t.Endpoints)
	s.dataProtectionLOSCapabilities = string(t.Links.dataProtectionLOSCapabilities)
	s.dataSecurityLOSCapabilities = string(t.Links.dataSecurityLOSCapabilities)
	s.dataStorageLOSCapabilities = string(t.Links.dataStorageLOSCapabilities)
	s.enclosures = string(t.Links.enclosures)
	s.hostingSystem = string(t.Links.hostingSystem)
	s.ioConnectivityLOSCapabilities = string(t.Links.ioConnectivityLOSCapabilities)
	s.ioPerformanceLOSCapabilities = string(t.Links.ioPerformanceLOSCapabilities)
	s.serverEndpointGroups = string(t.ServerEndpointGroups)
	s.storageGroups = string(t.StorageGroups)
	s.storagePools = string(t.StoragePools)
	s.storageSubsystems = string(t.StorageSubsystems)
	s.volumes = t.Volumes.ToStrings()

	return nil
}

// GetStorageService will get a StorageService instance from the Swordfish service.
func GetStorageService(c common.Client, uri string) (*StorageService, error) {
	resp, err := c.Get(uri)
	defer resp.Body.Close()

	var StorageService StorageService
	err = json.NewDecoder(resp.Body).Decode(&StorageService)
	if err != nil {
		return nil, err
	}

	StorageService.SetClient(c)
	return &StorageService, nil
}

// ListReferencedStorageServices gets the collection of StorageServices
func ListReferencedStorageServices(c common.Client, link string) ([]*StorageService, error) {
	var result []*StorageService
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, StorageServiceLink := range links.ItemLinks {
		StorageService, err := GetStorageService(c, StorageServiceLink)
		if err != nil {
			return result, err
		}
		result = append(result, StorageService)
	}

	return result, nil
}

// ListStorageServices gets all StorageService in the system
func ListStorageServices(c common.Client) ([]*StorageService, error) {
	return ListReferencedStorageServices(c, DefaultStorageServicePath)
}

// ClassesOfService gets the storage service classes of service
func (s *StorageService) ClassesOfService() ([]*ClassesOfService, error) {
	return ListReferencedClassesOfServices(s.Client, s.classesOfService)
}
