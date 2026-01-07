//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.6 - #StorageService.v1_7_0.StorageService

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// StorageService Collection of resources that the system can make available to
// one or more host systems. The collection can contain: block, file, or object
// storage; local system access points through which the collection is made
// available; hosts, or host access points to which the collection is made
// available.
type StorageService struct {
	common.Entity
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
	// // DataProtectionLoSCapabilities shall reference the data protection
	// // capabilities of this service.
	// //
	// // Version added: v1.2.0
	// DataProtectionLoSCapabilities DataProtectionLoSCapabilities
	// // DataSecurityLoSCapabilities shall reference the data security capabilities
	// // of this service.
	// //
	// // Version added: v1.2.0
	// DataSecurityLoSCapabilities DataSecurityLoSCapabilities
	// // DataStorageLoSCapabilities shall reference the data storage capabilities of
	// // this service.
	// //
	// // Version added: v1.2.0
	// DataStorageLoSCapabilities DataStorageLoSCapabilities
	// // DefaultClassOfService shall reference the default class of service for
	// // entities allocated by this storage service. This default may be overridden
	// // by the DefaultClassOfService property values within contained StoragePools.
	// //
	// // Version added: v1.2.0
	// DefaultClassOfService ClassOfService
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
	// // IOConnectivityLoSCapabilities shall reference the IO connectivity
	// // capabilities of this service.
	// //
	// // Version added: v1.2.0
	// IOConnectivityLoSCapabilities IOConnectivityLoSCapabilities
	// // IOPerformanceLoSCapabilities shall reference the IO performance capabilities
	// // of this service.
	// //
	// // Version added: v1.2.0
	// IOPerformanceLoSCapabilities IOPerformanceLoSCapabilities
	// IOStatistics shall represent IO statistics for this StorageService.
	//
	// Version added: v1.2.0
	//
	// Deprecated: v1.7.0
	// This property is deprecated in favor of the IOStatistics property in
	// StorageServiceMetrics.
	IOStatistics IOStatistics
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// LinesOfService shall reference a LineOfService collection defined for this
	// service.
	//
	// Version added: v1.4.0
	LinesOfService []LineOfService
	// LinesOfService@odata.count
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
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Redundancy shall contain the redundancy information for the storage
	// subsystem.
	Redundancy []common.Redundancy
	// Redundancy@odata.count
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
	SpareResourceSets []SpareResourceSet
	// SpareResourceSets@odata.count
	SpareResourceSetsCount int `json:"SpareResourceSets@odata.count"`
	// Status shall contain the status of the StorageService.
	Status common.Status
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
	// StorageSubsystems@odata.count
	StorageSubsystemsCount int `json:"StorageSubsystems@odata.count"`
	// Volumes is an array of references to Volumes managed by this storage
	// service.
	volumes string
	// setEncryptionKeyTarget is the URL to send SetEncryptionKey requests.
	setEncryptionKeyTarget string
	// dataProtectionLoSCapabilities is the URI for DataProtectionLoSCapabilities.
	dataProtectionLoSCapabilities string
	// dataSecurityLoSCapabilities is the URI for DataSecurityLoSCapabilities.
	dataSecurityLoSCapabilities string
	// dataStorageLoSCapabilities is the URI for DataStorageLoSCapabilities.
	dataStorageLoSCapabilities string
	// defaultClassOfService is the URI for DefaultClassOfService.
	defaultClassOfService string
	// hostingSystem is the URI for HostingSystem.
	hostingSystem string
	// iOConnectivityLoSCapabilities is the URI for IOConnectivityLoSCapabilities.
	iOConnectivityLoSCapabilities string
	// iOPerformanceLoSCapabilities is the URI for IOPerformanceLoSCapabilities.
	iOPerformanceLoSCapabilities string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a StorageService object from the raw JSON.
func (s *StorageService) UnmarshalJSON(b []byte) error {
	type temp StorageService
	type sActions struct {
		SetEncryptionKey common.ActionTarget `json:"#StorageService.SetEncryptionKey"`
	}
	type sLinks struct {
		DataProtectionLoSCapabilities common.Link `json:"DataProtectionLoSCapabilities"`
		DataSecurityLoSCapabilities   common.Link `json:"DataSecurityLoSCapabilities"`
		DataStorageLoSCapabilities    common.Link `json:"DataStorageLoSCapabilities"`
		DefaultClassOfService         common.Link `json:"DefaultClassOfService"`
		HostingSystem                 common.Link `json:"HostingSystem"`
		IOConnectivityLoSCapabilities common.Link `json:"IOConnectivityLoSCapabilities"`
		IOPerformanceLoSCapabilities  common.Link `json:"IOPerformanceLoSCapabilities"`
	}
	var tmp struct {
		temp
		Actions              sActions
		Links                sLinks
		ClassesOfService     common.Link  `json:"classesOfService"`
		ClientEndpointGroups common.Link  `json:"clientEndpointGroups"`
		Connections          common.Link  `json:"connections"`
		ConsistencyGroups    common.Link  `json:"consistencyGroups"`
		Drives               common.Link  `json:"drives"`
		EndpointGroups       common.Link  `json:"endpointGroups"`
		Endpoints            common.Link  `json:"endpoints"`
		FileSystems          common.Link  `json:"fileSystems"`
		Metrics              common.Link  `json:"metrics"`
		ServerEndpointGroups common.Link  `json:"serverEndpointGroups"`
		StorageGroups        common.Link  `json:"storageGroups"`
		StoragePools         common.Link  `json:"storagePools"`
		StorageSubsystems    common.Links `json:"storageSubsystems"`
		Volumes              common.Link  `json:"volumes"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = StorageService(tmp.temp)

	// Extract the links to other entities for later
	s.setEncryptionKeyTarget = tmp.Actions.SetEncryptionKey.Target
	s.dataProtectionLoSCapabilities = tmp.Links.DataProtectionLoSCapabilities.String()
	s.dataSecurityLoSCapabilities = tmp.Links.DataSecurityLoSCapabilities.String()
	s.dataStorageLoSCapabilities = tmp.Links.DataStorageLoSCapabilities.String()
	s.defaultClassOfService = tmp.Links.DefaultClassOfService.String()
	s.hostingSystem = tmp.Links.HostingSystem.String()
	s.iOConnectivityLoSCapabilities = tmp.Links.IOConnectivityLoSCapabilities.String()
	s.iOPerformanceLoSCapabilities = tmp.Links.IOPerformanceLoSCapabilities.String()
	s.classesOfService = tmp.ClassesOfService.String()
	s.clientEndpointGroups = tmp.ClientEndpointGroups.String()
	s.connections = tmp.Connections.String()
	s.consistencyGroups = tmp.ConsistencyGroups.String()
	s.drives = tmp.Drives.String()
	s.endpointGroups = tmp.EndpointGroups.String()
	s.endpoints = tmp.Endpoints.String()
	s.fileSystems = tmp.FileSystems.String()
	s.metrics = tmp.Metrics.String()
	s.serverEndpointGroups = tmp.ServerEndpointGroups.String()
	s.storageGroups = tmp.StorageGroups.String()
	s.storagePools = tmp.StoragePools.String()
	s.storageSubsystems = tmp.StorageSubsystems.ToStrings()
	s.volumes = tmp.Volumes.String()

	// This is a read/write object, so we need to save the raw object data for later
	s.rawData = b

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
		"IOStatistics",
		"Identifier",
		"LinesOfService",
		"LinesOfService@odata.count",
		"Redundancy@odata.count",
		"ServerEndpointGroups",
		"SpareResourceSets",
		"SpareResourceSets@odata.count",
		"Status",
		"StorageSubsystems@odata.count",
		"Volumes",
	}

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
}

// GetStorageService will get a StorageService instance from the service.
func GetStorageService(c common.Client, uri string) (*StorageService, error) {
	return common.GetObject[StorageService](c, uri)
}

// ListReferencedStorageServices gets the collection of StorageService from
// a provided reference.
func ListReferencedStorageServices(c common.Client, link string) ([]*StorageService, error) {
	return common.GetCollectionObjects[StorageService](c, link)
}

// SetEncryptionKey defines the name of the custom action supported on this resource.
// encryptionKey - This defines the property name for the action.
func (s *StorageService) SetEncryptionKey(encryptionKey string) error {
	payload := make(map[string]any)
	payload["EncryptionKey"] = encryptionKey
	return s.Post(s.setEncryptionKeyTarget, payload)
}

// DataProtectionLoSCapabilities gets the DataProtectionLoSCapabilities linked resource.
func (s *StorageService) DataProtectionLoSCapabilities(client common.Client) (*DataProtectionLoSCapabilities, error) {
	if s.dataProtectionLoSCapabilities == "" {
		return nil, nil
	}
	return common.GetObject[DataProtectionLoSCapabilities](client, s.dataProtectionLoSCapabilities)
}

// DataSecurityLoSCapabilities gets the DataSecurityLoSCapabilities linked resource.
func (s *StorageService) DataSecurityLoSCapabilities(client common.Client) (*DataSecurityLoSCapabilities, error) {
	if s.dataSecurityLoSCapabilities == "" {
		return nil, nil
	}
	return common.GetObject[DataSecurityLoSCapabilities](client, s.dataSecurityLoSCapabilities)
}

// DataStorageLoSCapabilities gets the DataStorageLoSCapabilities linked resource.
func (s *StorageService) DataStorageLoSCapabilities(client common.Client) (*DataStorageLoSCapabilities, error) {
	if s.dataStorageLoSCapabilities == "" {
		return nil, nil
	}
	return common.GetObject[DataStorageLoSCapabilities](client, s.dataStorageLoSCapabilities)
}

// DefaultClassOfService gets the DefaultClassOfService linked resource.
func (s *StorageService) DefaultClassOfService(client common.Client) (*ClassOfService, error) {
	if s.defaultClassOfService == "" {
		return nil, nil
	}
	return common.GetObject[ClassOfService](client, s.defaultClassOfService)
}

// HostingSystem gets the HostingSystem linked resource.
func (s *StorageService) HostingSystem(client common.Client) (*common.Resource, error) {
	if s.hostingSystem == "" {
		return nil, nil
	}
	return common.GetObject[common.Resource](client, s.hostingSystem)
}

// IOConnectivityLoSCapabilities gets the IOConnectivityLoSCapabilities linked resource.
func (s *StorageService) IOConnectivityLoSCapabilities(client common.Client) (*IOConnectivityLoSCapabilities, error) {
	if s.iOConnectivityLoSCapabilities == "" {
		return nil, nil
	}
	return common.GetObject[IOConnectivityLoSCapabilities](client, s.iOConnectivityLoSCapabilities)
}

// IOPerformanceLoSCapabilities gets the IOPerformanceLoSCapabilities linked resource.
func (s *StorageService) IOPerformanceLoSCapabilities(client common.Client) (*IOPerformanceLoSCapabilities, error) {
	if s.iOPerformanceLoSCapabilities == "" {
		return nil, nil
	}
	return common.GetObject[IOPerformanceLoSCapabilities](client, s.iOPerformanceLoSCapabilities)
}

// ClassesOfService gets the ClassesOfService collection.
func (s *StorageService) ClassesOfService(client common.Client) ([]*ClassOfService, error) {
	if s.classesOfService == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[ClassOfService](client, s.classesOfService)
}

// ClientEndpointGroups gets the ClientEndpointGroups collection.
func (s *StorageService) ClientEndpointGroups(client common.Client) ([]*EndpointGroup, error) {
	if s.clientEndpointGroups == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[EndpointGroup](client, s.clientEndpointGroups)
}

// Connections gets the Connections collection.
func (s *StorageService) Connections(client common.Client) ([]*redfish.Connection, error) {
	if s.connections == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[redfish.Connection](client, s.connections)
}

// ConsistencyGroups gets the ConsistencyGroups collection.
func (s *StorageService) ConsistencyGroups(client common.Client) ([]*ConsistencyGroup, error) {
	if s.consistencyGroups == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[ConsistencyGroup](client, s.consistencyGroups)
}

// Drives gets the Drives collection.
func (s *StorageService) Drives(client common.Client) ([]*redfish.Drive, error) {
	if s.drives == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[redfish.Drive](client, s.drives)
}

// EndpointGroups gets the EndpointGroups collection.
func (s *StorageService) EndpointGroups(client common.Client) ([]*EndpointGroup, error) {
	if s.endpointGroups == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[EndpointGroup](client, s.endpointGroups)
}

// Endpoints gets the Endpoints collection.
func (s *StorageService) Endpoints(client common.Client) ([]*redfish.Endpoint, error) {
	if s.endpoints == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[redfish.Endpoint](client, s.endpoints)
}

// FileSystems gets the FileSystems collection.
func (s *StorageService) FileSystems(client common.Client) ([]*FileSystem, error) {
	if s.fileSystems == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[FileSystem](client, s.fileSystems)
}

// Metrics gets the Metrics linked resource.
func (s *StorageService) Metrics(client common.Client) (*StorageServiceMetrics, error) {
	if s.metrics == "" {
		return nil, nil
	}
	return common.GetObject[StorageServiceMetrics](client, s.metrics)
}

// ServerEndpointGroups gets the ServerEndpointGroups collection.
func (s *StorageService) ServerEndpointGroups(client common.Client) ([]*EndpointGroup, error) {
	if s.serverEndpointGroups == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[EndpointGroup](client, s.serverEndpointGroups)
}

// StorageGroups gets the StorageGroups collection.
func (s *StorageService) StorageGroups(client common.Client) ([]*StorageGroup, error) {
	if s.storageGroups == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[StorageGroup](client, s.storageGroups)
}

// StoragePools gets the StoragePools collection.
func (s *StorageService) StoragePools(client common.Client) ([]*StoragePool, error) {
	if s.storagePools == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[StoragePool](client, s.storagePools)
}

// StorageSubsystems gets the StorageSubsystems linked resources.
func (s *StorageService) StorageSubsystems(client common.Client) ([]*redfish.Storage, error) {
	return common.GetObjects[redfish.Storage](client, s.storageSubsystems)
}

// Volumes gets the Volumes collection.
func (s *StorageService) Volumes(client common.Client) ([]*common.Volume, error) {
	if s.volumes == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[common.Volume](client, s.volumes)
}
