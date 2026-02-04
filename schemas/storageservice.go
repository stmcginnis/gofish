//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/StorageService.v1_7_0.json
// 1.2.6 - #StorageService.v1_7_0.StorageService

package schemas

import (
	"encoding/json"
)

// StorageService Collection of resources that the system can make available to
// one or more host systems. The collection can contain: block, file, or object
// storage; local system access points through which the collection is made
// available; hosts, or host access points to which the collection is made
// available.
type StorageService struct {
	Entity
	// ClassesOfService shall reference a ClassOfService supported by this service.
	classesOfService string
	// ClientEndpointGroups shall reference an EndpointGroup.
	//
	// Deprecated
	// Deprecated in favor of EndpointGroups. The GroupType property of
	// EndpointGroup already distinguishes between use for Server or Client.
	clientEndpointGroups string
	// Connections shall contain references to all Connections that include this
	// volume.
	//
	// Version added: v1.6.0
	connections string
	// ConsistencyGroups shall reference a ConsistencyGroup.
	//
	// Version added: v1.3.0
	consistencyGroups string
	// DataProtectionLoSCapabilities shall reference the data protection
	// capabilities of this service.
	//
	// Version added: v1.2.0
	dataProtectionLoSCapabilities string
	// DataSecurityLoSCapabilities shall reference the data security capabilities
	// of this service.
	//
	// Version added: v1.2.0
	dataSecurityLoSCapabilities string
	// DataStorageLoSCapabilities shall reference the data storage capabilities of
	// this service.
	//
	// Version added: v1.2.0
	dataStorageLoSCapabilities string
	// DefaultClassOfService shall reference the default class of service for
	// entities allocated by this storage service. This default may be overridden
	// by the DefaultClassOfService property values within contained StoragePools.
	//
	// Version added: v1.2.0
	defaultClassOfService string
	// Drives is a collection that indicates all the drives managed by this storage
	// service.
	drives string
	// EndpointGroups shall reference an EndpointGroup.
	endpointGroups string
	// Endpoints shall reference an Endpoint managed by this service.
	endpoints string
	// FileSystems is an array of references to FileSystems managed by this storage
	// service.
	fileSystems string
	// IOConnectivityLoSCapabilities shall reference the IO connectivity
	// capabilities of this service.
	//
	// Version added: v1.2.0
	iOConnectivityLoSCapabilities string
	// IOPerformanceLoSCapabilities shall reference the IO performance capabilities
	// of this service.
	//
	// Version added: v1.2.0
	iOPerformanceLoSCapabilities string
	// IOStatistics shall represent IO statistics for this StorageService.
	//
	// Version added: v1.2.0
	//
	// Deprecated: v1.7.0
	// This property is deprecated in favor of the IOStatistics property in
	// StorageServiceMetrics.
	IOStatistics IOStatistics
	// Identifier shall be unique within the managed ecosystem.
	Identifier Identifier
	// LinesOfService shall reference a LineOfService collection defined for this
	// service.
	//
	// Version added: v1.4.0
	LinesOfService []LineOfService
	// LinesOfServiceCount
	LinesOfServiceCount int `json:"LinesOfService@odata.count"`
	// Metrics shall contain a link to a resource of type StorageServiceMetrics
	// that specifies the metrics for this storage service. IO metrics are reported
	// in the IOStatistics property.
	//
	// Version added: v1.7.0
	metrics string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Redundancy shall contain the redundancy information for the storage
	// subsystem.
	redundancy string
	// RedundancyCount
	RedundancyCount int `json:"Redundancy@odata.count"`
	// ServerEndpointGroups shall reference a EndpointGroup.
	//
	// Deprecated
	// Deprecated in favor of EndpointGroups. The GroupType property of
	// EndpointGroup already distinguishes between use for Server or Client.
	serverEndpointGroups string
	// SpareResourceSets shall contain resources that may be utilized to replace
	// the capacity provided by a failed resource having a compatible type.
	//
	// Version added: v1.2.0
	spareResourceSets []string
	// SpareResourceSetsCount
	SpareResourceSetsCount int `json:"SpareResourceSets@odata.count"`
	// Status shall contain the status of the StorageService.
	Status Status
	// StorageGroups shall reference a StorageGroup.
	//
	// Deprecated: v1.6.0
	// This property is deprecated in favor of the Connections property.
	storageGroups string
	// StoragePools is an array of references to StoragePools.
	storagePools string
	// StorageSubsystems shall be a link to a collection of type StorageCollection
	// having members that represent storage subsystems managed by this storage
	// service.
	//
	// Version added: v1.0.1
	storageSubsystems []string
	// StorageSubsystemsCount
	StorageSubsystemsCount int `json:"StorageSubsystems@odata.count"`
	// Volumes is an array of references to Volumes managed by this storage
	// service.
	volumes string
	// setEncryptionKeyTarget is the URL to send SetEncryptionKey requests.
	setEncryptionKeyTarget string
	// hostingSystem is the URI for HostingSystem.
	hostingSystem string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a StorageService object from the raw JSON.
func (s *StorageService) UnmarshalJSON(b []byte) error {
	type temp StorageService
	type sActions struct {
		SetEncryptionKey ActionTarget `json:"#StorageService.SetEncryptionKey"`
	}
	type sLinks struct {
		HostingSystem Link `json:"HostingSystem"`
	}
	var tmp struct {
		temp
		Actions                       sActions
		Links                         sLinks
		ClassesOfService              Link  `json:"ClassesOfService"`
		ClientEndpointGroups          Link  `json:"ClientEndpointGroups"`
		Connections                   Link  `json:"Connections"`
		ConsistencyGroups             Link  `json:"ConsistencyGroups"`
		DataProtectionLoSCapabilities Link  `json:"DataProtectionLoSCapabilities"`
		DataSecurityLoSCapabilities   Link  `json:"DataSecurityLoSCapabilities"`
		DataStorageLoSCapabilities    Link  `json:"DataStorageLoSCapabilities"`
		DefaultClassOfService         Link  `json:"DefaultClassOfService"`
		Drives                        Link  `json:"Drives"`
		EndpointGroups                Link  `json:"EndpointGroups"`
		Endpoints                     Link  `json:"Endpoints"`
		FileSystems                   Link  `json:"FileSystems"`
		IOConnectivityLoSCapabilities Link  `json:"IOConnectivityLoSCapabilities"`
		IOPerformanceLoSCapabilities  Link  `json:"IOPerformanceLoSCapabilities"`
		Metrics                       Link  `json:"Metrics"`
		Redundancy                    Link  `json:"Redundancy"`
		ServerEndpointGroups          Link  `json:"ServerEndpointGroups"`
		SpareResourceSets             Links `json:"SpareResourceSets"`
		StorageGroups                 Link  `json:"StorageGroups"`
		StoragePools                  Link  `json:"StoragePools"`
		StorageSubsystems             Links `json:"StorageSubsystems"`
		Volumes                       Link  `json:"Volumes"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = StorageService(tmp.temp)

	// Extract the links to other entities for later
	s.setEncryptionKeyTarget = tmp.Actions.SetEncryptionKey.Target
	s.hostingSystem = tmp.Links.HostingSystem.String()
	s.classesOfService = tmp.ClassesOfService.String()
	s.clientEndpointGroups = tmp.ClientEndpointGroups.String()
	s.connections = tmp.Connections.String()
	s.consistencyGroups = tmp.ConsistencyGroups.String()
	s.dataProtectionLoSCapabilities = tmp.DataProtectionLoSCapabilities.String()
	s.dataSecurityLoSCapabilities = tmp.DataSecurityLoSCapabilities.String()
	s.dataStorageLoSCapabilities = tmp.DataStorageLoSCapabilities.String()
	s.defaultClassOfService = tmp.DefaultClassOfService.String()
	s.drives = tmp.Drives.String()
	s.endpointGroups = tmp.EndpointGroups.String()
	s.endpoints = tmp.Endpoints.String()
	s.fileSystems = tmp.FileSystems.String()
	s.iOConnectivityLoSCapabilities = tmp.IOConnectivityLoSCapabilities.String()
	s.iOPerformanceLoSCapabilities = tmp.IOPerformanceLoSCapabilities.String()
	s.metrics = tmp.Metrics.String()
	s.redundancy = tmp.Redundancy.String()
	s.serverEndpointGroups = tmp.ServerEndpointGroups.String()
	s.spareResourceSets = tmp.SpareResourceSets.ToStrings()
	s.storageGroups = tmp.StorageGroups.String()
	s.storagePools = tmp.StoragePools.String()
	s.storageSubsystems = tmp.StorageSubsystems.ToStrings()
	s.volumes = tmp.Volumes.String()

	// This is a read/write object, so we need to save the raw object data for later
	s.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *StorageService) Update() error {
	readWriteFields := []string{
		"ClassesOfService",
		"ClientEndpointGroups",
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
		"ServerEndpointGroups",
		"SpareResourceSets",
		"Volumes",
	}

	return s.UpdateFromRawData(s, s.RawData, readWriteFields)
}

// GetStorageService will get a StorageService instance from the service.
func GetStorageService(c Client, uri string) (*StorageService, error) {
	return GetObject[StorageService](c, uri)
}

// ListReferencedStorageServices gets the collection of StorageService from
// a provided reference.
func ListReferencedStorageServices(c Client, link string) ([]*StorageService, error) {
	return GetCollectionObjects[StorageService](c, link)
}

// This defines the name of the custom action supported on this resource.
// encryptionKey - This defines the property name for the action.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageService) SetEncryptionKey(encryptionKey string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["EncryptionKey"] = encryptionKey
	resp, taskInfo, err := PostWithTask(s.client,
		s.setEncryptionKeyTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// HostingSystem gets the HostingSystem linked resource.
func (s *StorageService) HostingSystem() (*Resource, error) {
	if s.hostingSystem == "" {
		return nil, nil
	}
	return GetObject[Resource](s.client, s.hostingSystem)
}

// ClassesOfService gets the ClassesOfService collection.
func (s *StorageService) ClassesOfService() ([]*ClassOfService, error) {
	if s.classesOfService == "" {
		return nil, nil
	}
	return GetCollectionObjects[ClassOfService](s.client, s.classesOfService)
}

// ClientEndpointGroups gets the ClientEndpointGroups collection.
func (s *StorageService) ClientEndpointGroups() ([]*EndpointGroup, error) {
	if s.clientEndpointGroups == "" {
		return nil, nil
	}
	return GetCollectionObjects[EndpointGroup](s.client, s.clientEndpointGroups)
}

// Connections gets the Connections collection.
func (s *StorageService) Connections() ([]*Connection, error) {
	if s.connections == "" {
		return nil, nil
	}
	return GetCollectionObjects[Connection](s.client, s.connections)
}

// ConsistencyGroups gets the ConsistencyGroups collection.
func (s *StorageService) ConsistencyGroups() ([]*ConsistencyGroup, error) {
	if s.consistencyGroups == "" {
		return nil, nil
	}
	return GetCollectionObjects[ConsistencyGroup](s.client, s.consistencyGroups)
}

// DataProtectionLoSCapabilities gets the DataProtectionLoSCapabilities linked resource.
func (s *StorageService) DataProtectionLoSCapabilities() (*DataProtectionLoSCapabilities, error) {
	if s.dataProtectionLoSCapabilities == "" {
		return nil, nil
	}
	return GetObject[DataProtectionLoSCapabilities](s.client, s.dataProtectionLoSCapabilities)
}

// DataSecurityLoSCapabilities gets the DataSecurityLoSCapabilities linked resource.
func (s *StorageService) DataSecurityLoSCapabilities() (*DataSecurityLoSCapabilities, error) {
	if s.dataSecurityLoSCapabilities == "" {
		return nil, nil
	}
	return GetObject[DataSecurityLoSCapabilities](s.client, s.dataSecurityLoSCapabilities)
}

// DataStorageLoSCapabilities gets the DataStorageLoSCapabilities linked resource.
func (s *StorageService) DataStorageLoSCapabilities() (*DataStorageLoSCapabilities, error) {
	if s.dataStorageLoSCapabilities == "" {
		return nil, nil
	}
	return GetObject[DataStorageLoSCapabilities](s.client, s.dataStorageLoSCapabilities)
}

// DefaultClassOfService gets the DefaultClassOfService linked resource.
func (s *StorageService) DefaultClassOfService() (*ClassOfService, error) {
	if s.defaultClassOfService == "" {
		return nil, nil
	}
	return GetObject[ClassOfService](s.client, s.defaultClassOfService)
}

// Drives gets the Drives collection.
func (s *StorageService) Drives() ([]*Drive, error) {
	if s.drives == "" {
		return nil, nil
	}
	return GetCollectionObjects[Drive](s.client, s.drives)
}

// EndpointGroups gets the EndpointGroups collection.
func (s *StorageService) EndpointGroups() ([]*EndpointGroup, error) {
	if s.endpointGroups == "" {
		return nil, nil
	}
	return GetCollectionObjects[EndpointGroup](s.client, s.endpointGroups)
}

// Endpoints gets the Endpoints collection.
func (s *StorageService) Endpoints() ([]*Endpoint, error) {
	if s.endpoints == "" {
		return nil, nil
	}
	return GetCollectionObjects[Endpoint](s.client, s.endpoints)
}

// FileSystems gets the FileSystems collection.
func (s *StorageService) FileSystems() ([]*FileSystem, error) {
	if s.fileSystems == "" {
		return nil, nil
	}
	return GetCollectionObjects[FileSystem](s.client, s.fileSystems)
}

// IOConnectivityLoSCapabilities gets the IOConnectivityLoSCapabilities linked resource.
func (s *StorageService) IOConnectivityLoSCapabilities() (*IOConnectivityLoSCapabilities, error) {
	if s.iOConnectivityLoSCapabilities == "" {
		return nil, nil
	}
	return GetObject[IOConnectivityLoSCapabilities](s.client, s.iOConnectivityLoSCapabilities)
}

// IOPerformanceLoSCapabilities gets the IOPerformanceLoSCapabilities linked resource.
func (s *StorageService) IOPerformanceLoSCapabilities() (*IOPerformanceLoSCapabilities, error) {
	if s.iOPerformanceLoSCapabilities == "" {
		return nil, nil
	}
	return GetObject[IOPerformanceLoSCapabilities](s.client, s.iOPerformanceLoSCapabilities)
}

// Metrics gets the Metrics linked resource.
func (s *StorageService) Metrics() (*StorageServiceMetrics, error) {
	if s.metrics == "" {
		return nil, nil
	}
	return GetObject[StorageServiceMetrics](s.client, s.metrics)
}

// Redundancy gets the Redundancy linked resource.
func (s *StorageService) Redundancy() (*Redundancy, error) {
	if s.redundancy == "" {
		return nil, nil
	}
	return GetObject[Redundancy](s.client, s.redundancy)
}

// ServerEndpointGroups gets the ServerEndpointGroups collection.
func (s *StorageService) ServerEndpointGroups() ([]*EndpointGroup, error) {
	if s.serverEndpointGroups == "" {
		return nil, nil
	}
	return GetCollectionObjects[EndpointGroup](s.client, s.serverEndpointGroups)
}

// SpareResourceSets gets the SpareResourceSets linked resources.
func (s *StorageService) SpareResourceSets() ([]*SpareResourceSet, error) {
	return GetObjects[SpareResourceSet](s.client, s.spareResourceSets)
}

// StorageGroups gets the StorageGroups collection.
func (s *StorageService) StorageGroups() ([]*StorageGroup, error) {
	if s.storageGroups == "" {
		return nil, nil
	}
	return GetCollectionObjects[StorageGroup](s.client, s.storageGroups)
}

// StoragePools gets the StoragePools collection.
func (s *StorageService) StoragePools() ([]*StoragePool, error) {
	if s.storagePools == "" {
		return nil, nil
	}
	return GetCollectionObjects[StoragePool](s.client, s.storagePools)
}

// StorageSubsystems gets the StorageSubsystems linked resources.
func (s *StorageService) StorageSubsystems() ([]*Storage, error) {
	return GetObjects[Storage](s.client, s.storageSubsystems)
}

// Volumes gets the Volumes collection.
func (s *StorageService) Volumes() ([]*Volume, error) {
	if s.volumes == "" {
		return nil, nil
	}
	return GetCollectionObjects[Volume](s.client, s.volumes)
}
