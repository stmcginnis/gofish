//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #Storage.v1_20_0.Storage

package redfish

import (
	"encoding/json"

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
	// UseExternalKeyEncryptionMode The storage subsystem uses one or more external
	// keys for encryption.
	UseExternalKeyEncryptionMode EncryptionMode = "UseExternalKey"
	// UseLocalKeyEncryptionMode The storage subsystem uses a local key for
	// encryption.
	UseLocalKeyEncryptionMode EncryptionMode = "UseLocalKey"
	// PasswordOnlyEncryptionMode The storage subsystem uses a password, but no
	// keys for encryption.
	PasswordOnlyEncryptionMode EncryptionMode = "PasswordOnly"
	// PasswordWithExternalKeyEncryptionMode The storage subsystem uses a password
	// and one or more external keys for encryption.
	PasswordWithExternalKeyEncryptionMode EncryptionMode = "PasswordWithExternalKey"
	// PasswordWithLocalKeyEncryptionMode The storage subsystem uses a password and
	// a local key for encryption.
	PasswordWithLocalKeyEncryptionMode EncryptionMode = "PasswordWithLocalKey"
)

type HotspareActivationPolicy string

const (
	// OnDriveFailureHotspareActivationPolicy The hot spare drive will take over
	// for the original drive when the original drive has been marked as failed by
	// the storage domain.
	OnDriveFailureHotspareActivationPolicy HotspareActivationPolicy = "OnDriveFailure"
	// OnDrivePredictedFailureHotspareActivationPolicy The hot spare drive will
	// take over for the original drive when the original drive has been predicted
	// to fail in the future by the storage domain.
	OnDrivePredictedFailureHotspareActivationPolicy HotspareActivationPolicy = "OnDrivePredictedFailure"
	// OEMHotspareActivationPolicy The hot spare drive will take over for the
	// original drive in an algorithm custom to the OEM.
	OEMHotspareActivationPolicy HotspareActivationPolicy = "OEM"
)

type StorageResetToDefaultsType string

const (
	// ResetAllStorageResetToDefaultsType Reset all settings to factory defaults and
	// remove all volumes.
	ResetAllStorageResetToDefaultsType StorageResetToDefaultsType = "ResetAll"
	// PreserveVolumesStorageResetToDefaultsType Reset all settings to factory defaults
	// but preserve the configured volumes on the controllers.
	PreserveVolumesStorageResetToDefaultsType StorageResetToDefaultsType = "PreserveVolumes"
)

type VolumeAssignmentPolicy string

const (
	// UnassignedVolumeAssignmentPolicy Newly created volumes are not assigned to
	// physical functions, but are added to 'UnassignedVolumes' in the 'Links'
	// property within this resource.
	UnassignedVolumeAssignmentPolicy VolumeAssignmentPolicy = "Unassigned"
	// SupervisorVolumeAssignmentPolicy Newly created volumes are assigned to the
	// supervisor function.
	SupervisorVolumeAssignmentPolicy VolumeAssignmentPolicy = "Supervisor"
	// WeightedRoundRobinVolumeAssignmentPolicy Newly created volumes are assigned
	// to the physical functions in a round-robin fashion.
	WeightedRoundRobinVolumeAssignmentPolicy VolumeAssignmentPolicy = "WeightedRoundRobin"
)

// Storage shall represent a storage subsystem in the Redfish Specification.
type Storage struct {
	common.Entity
	// AutoVolumeCreate shall indicate if volumes are created automatically for
	// each unassigned drive attached to this storage subsystem.
	//
	// Version added: v1.15.0
	AutoVolumeCreate AutoVolumeCreate
	// BlockSecurityIDPolicy shall indicate if the storage controller sends the
	// TCG-defined 'Block SID' command to block establishment of a TCG-defined
	// security ID (SID) during each drive boot sequence for drives that support
	// it. The value 'true' shall indicate the TCG-defined 'Block SID' command is
	// sent to supporting drives during each drive boot sequence.
	//
	// Version added: v1.18.0
	BlockSecurityIDPolicy bool
	// ConfigurationLock shall indicate whether configuration requests to the
	// storage subsystem are locked. Services shall reject modification requests
	// that contain the value 'Partial'. Modifying the value of this property may
	// affect the 'ConfigurationLock' property in 'Drive' resources referenced by
	// the 'Drives' property.
	//
	// Version added: v1.16.0
	ConfigurationLock ConfigurationLock
	// Connections shall contain a link to a resource collection of type
	// 'ConnectionCollection'. The members of this collection shall reference
	// Connection resources subordinate to Fabric resources.
	//
	// Version added: v1.15.0
	connections string
	// ConsistencyGroups shall contain a link to a resource collection of type
	// 'ConsistencyGroupCollection'. The property shall be used when groups of
	// volumes are treated as a single resource by an application or set of
	// applications.
	//
	// Version added: v1.8.0
	consistencyGroups string
	// Controllers shall contain a link to a resource collection of type
	// 'StorageControllerCollection' that contains the set of storage controllers
	// allocated to this storage subsystem.
	//
	// Version added: v1.9.0
	controllers string
	// Drives shall contain a set of the drives attached to the storage controllers
	// that this resource represents.
	drives []string
	// Drives@odata.count
	DrivesCount int `json:"Drives@odata.count"`
	// EncryptionMode shall contain the encryption mode of this storage subsystem.
	//
	// Version added: v1.14.0
	EncryptionMode EncryptionMode
	// EndpointGroups shall contain a link to a resource collection of type
	// 'EndpointGroupCollection'. This property shall be implemented when atomic
	// control is needed to perform mapping, masking, and zoning operations.
	//
	// Version added: v1.8.0
	endpointGroups string
	// FileSystems shall contain a link to a resource collection of type
	// 'FileSystemCollection'. This property shall be used when file systems are
	// shared or exported by the storage subsystem.
	//
	// Version added: v1.8.0
	fileSystems string
	// HotspareActivationPolicy shall contain the policy under which all drives
	// operating as hot spares in this storage domain will activate.
	//
	// Version added: v1.14.0
	HotspareActivationPolicy HotspareActivationPolicy
	// Identifiers shall contain a list of all known durable names for the storage
	// subsystem.
	//
	// Version added: v1.9.0
	Identifiers []Identifier
	// LocalEncryptionKeyIdentifier shall contain the local encryption key
	// identifier used by the storage subsystem when 'EncryptionMode' contains
	// 'UseLocalKey'.
	//
	// Version added: v1.14.0
	LocalEncryptionKeyIdentifier string
	// MPF shall contain multiple physical function-related properties for storage
	// controllers in this storage subsystem. This property should only be present
	// if the storage subsystem supports the use of multiple physical functions to
	// parallelize data transfer.
	//
	// Version added: v1.20.0
	MPF MPF
	// Metrics shall contain a link to a resource of type 'StorageMetrics' that
	// represents the metrics associated with this storage subsystem.
	//
	// Version added: v1.18.0
	metrics string
	// NVMeSubsystemProperties shall contain information specific to NVMe
	// Subsystems. This property shall only be present if this resource represents
	// an NVMe Subsystem.
	//
	// Version added: v1.16.0
	NVMeSubsystemProperties NVMeSubsystemProperties
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Redundancy shall contain redundancy information for the storage subsystem.
	Redundancy []common.Redundancy
	// Redundancy@odata.count
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// StorageControllers shall contain a set of the storage controllers that this
	// resource represents.
	//
	// Deprecated: v1.13.0
	// This property has been deprecated in favor of 'Controllers' to allow for
	// storage controllers to be represented as their own resources.
	StorageControllers []StorageStorageController
	// StorageControllers@odata.count
	StorageControllersCount int `json:"StorageControllers@odata.count"`
	// StorageGroups shall contain a link to a resource collection of type
	// 'StorageGroupsCollection'. This property shall be used when implementing
	// mapping and masking.
	//
	// Version added: v1.8.0
	//
	// Deprecated: v1.15.0
	// This property was deprecated in favor of the 'Connections' property.
	storageGroups string
	// StoragePools shall contain a link to a resource collection of type
	// 'StoragePoolCollection'. This property shall be used when an abstraction of
	// media, rather than references to individual media, are used as the storage
	// data source.
	//
	// Version added: v1.8.0
	storagePools string
	// TargetConfigurationLockLevel shall contain the target configuration lock
	// level for the drive. For NVMe subsystems, services shall implement the
	// locking requirements specified by SNIA's Swordfish NVMe Model Overview and
	// Mapping Guide.
	//
	// Version added: v1.17.0
	TargetConfigurationLockLevel TargetConfigurationLockLevel
	// Volumes shall contain a link to a resource collection of type
	// 'VolumeCollection'.
	volumes string
	// importForeignDrivesTarget is the URL to send ImportForeignDrives requests.
	importForeignDrivesTarget string
	// rekeyExternalKeyTarget is the URL to send RekeyExternalKey requests.
	rekeyExternalKeyTarget string
	// resetToDefaultsTarget is the URL to send ResetToDefaults requests.
	resetToDefaultsTarget string
	// setControllerPasswordTarget is the URL to send SetControllerPassword requests.
	setControllerPasswordTarget string
	// setEncryptionKeyTarget is the URL to send SetEncryptionKey requests.
	setEncryptionKeyTarget string
	// blockSecurityIDUnsupportedDrives are the URIs for BlockSecurityIDUnsupportedDrives.
	blockSecurityIDUnsupportedDrives []string
	// blockSecurityIDUpdateUnsuccessfulDrives are the URIs for BlockSecurityIDUpdateUnsuccessfulDrives.
	blockSecurityIDUpdateUnsuccessfulDrives []string
	// enclosures are the URIs for Enclosures.
	enclosures []string
	// hostingStorageSystems are the URIs for HostingStorageSystems.
	hostingStorageSystems []string
	// nVMeoFDiscoverySubsystems are the URIs for NVMeoFDiscoverySubsystems.
	nVMeoFDiscoverySubsystems []string
	// simpleStorage is the URI for SimpleStorage.
	simpleStorage string
	// storageServices are the URIs for StorageServices.
	storageServices []string
	// unassignedVolumes are the URIs for UnassignedVolumes.
	unassignedVolumes []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Storage object from the raw JSON.
func (s *Storage) UnmarshalJSON(b []byte) error {
	type temp Storage
	type sActions struct {
		ImportForeignDrives   common.ActionTarget `json:"#Storage.ImportForeignDrives"`
		RekeyExternalKey      common.ActionTarget `json:"#Storage.RekeyExternalKey"`
		ResetToDefaults       common.ActionTarget `json:"#Storage.ResetToDefaults"`
		SetControllerPassword common.ActionTarget `json:"#Storage.SetControllerPassword"`
		SetEncryptionKey      common.ActionTarget `json:"#Storage.SetEncryptionKey"`
	}
	type sLinks struct {
		BlockSecurityIDUnsupportedDrives        common.Links `json:"BlockSecurityIDUnsupportedDrives"`
		BlockSecurityIDUpdateUnsuccessfulDrives common.Links `json:"BlockSecurityIDUpdateUnsuccessfulDrives"`
		Enclosures                              common.Links `json:"Enclosures"`
		HostingStorageSystems                   common.Links `json:"HostingStorageSystems"`
		NVMeoFDiscoverySubsystems               common.Links `json:"NVMeoFDiscoverySubsystems"`
		SimpleStorage                           common.Link  `json:"SimpleStorage"`
		StorageServices                         common.Links `json:"StorageServices"`
		UnassignedVolumes                       common.Links `json:"UnassignedVolumes"`
	}
	var tmp struct {
		temp
		Actions           sActions
		Links             sLinks
		Connections       common.Link `json:"connections"`
		ConsistencyGroups common.Link `json:"consistencyGroups"`
		Controllers       common.Link `json:"controllers"`
		Drives            common.Links
		EndpointGroups    common.Link `json:"endpointGroups"`
		FileSystems       common.Link `json:"fileSystems"`
		Metrics           common.Link `json:"metrics"`
		StorageGroups     common.Link `json:"storageGroups"`
		StoragePools      common.Link `json:"storagePools"`
		Volumes           common.Link `json:"volumes"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = Storage(tmp.temp)

	// Extract the links to other entities for later
	s.importForeignDrivesTarget = tmp.Actions.ImportForeignDrives.Target
	s.rekeyExternalKeyTarget = tmp.Actions.RekeyExternalKey.Target
	s.resetToDefaultsTarget = tmp.Actions.ResetToDefaults.Target
	s.setControllerPasswordTarget = tmp.Actions.SetControllerPassword.Target
	s.setEncryptionKeyTarget = tmp.Actions.SetEncryptionKey.Target
	s.blockSecurityIDUnsupportedDrives = tmp.Links.BlockSecurityIDUnsupportedDrives.ToStrings()
	s.blockSecurityIDUpdateUnsuccessfulDrives = tmp.Links.BlockSecurityIDUpdateUnsuccessfulDrives.ToStrings()
	s.enclosures = tmp.Links.Enclosures.ToStrings()
	s.hostingStorageSystems = tmp.Links.HostingStorageSystems.ToStrings()
	s.nVMeoFDiscoverySubsystems = tmp.Links.NVMeoFDiscoverySubsystems.ToStrings()
	s.simpleStorage = tmp.Links.SimpleStorage.String()
	s.storageServices = tmp.Links.StorageServices.ToStrings()
	s.unassignedVolumes = tmp.Links.UnassignedVolumes.ToStrings()
	s.connections = tmp.Connections.String()
	s.consistencyGroups = tmp.ConsistencyGroups.String()
	s.controllers = tmp.Controllers.String()
	s.drives = tmp.Drives.ToStrings()
	s.endpointGroups = tmp.EndpointGroups.String()
	s.fileSystems = tmp.FileSystems.String()
	s.metrics = tmp.Metrics.String()
	s.storageGroups = tmp.StorageGroups.String()
	s.storagePools = tmp.StoragePools.String()
	s.volumes = tmp.Volumes.String()

	// This is a read/write object, so we need to save the raw object data for later
	s.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *Storage) Update() error {
	readWriteFields := []string{
		"AutoVolumeCreate",
		"BlockSecurityIDPolicy",
		"ConfigurationLock",
		"Drives@odata.count",
		"EncryptionMode",
		"HotspareActivationPolicy",
		"Identifiers",
		"MPF",
		"NVMeSubsystemProperties",
		"Redundancy",
		"Redundancy@odata.count",
		"Status",
		"StorageControllers@odata.count",
		"TargetConfigurationLockLevel",
	}

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
}

// GetStorage will get a Storage instance from the service.
func GetStorage(c common.Client, uri string) (*Storage, error) {
	return common.GetObject[Storage](c, uri)
}

// ListReferencedStorages gets the collection of Storage from
// a provided reference.
func ListReferencedStorages(c common.Client, link string) ([]*Storage, error) {
	return common.GetCollectionObjects[Storage](c, link)
}

// ImportForeignDrives shall import encrypted foreign drives to a host-based storage
// controller by supplying encryption access credentials. Upon successful
// completion, the 'EncryptionStatus' property of the affected 'Drive' shall
// change from 'Foreign' to 'Unlocked'.
// controllerPassword - This parameter shall contain the current controller
// password. Services may reject the action request if this parameter is not
// provided or the value supplied does not match the current password.
// driveEncryptionKey - This parameter shall contain the encryption key to
// unlock the drives.
// driveEncryptionKeyIdentifier - This parameter shall contain an encryption
// key identifier to filter the drives that are imported. If this parameter is
// not provided, the service shall not perform filtering of foreign drives.
func (s *Storage) ImportForeignDrives(controllerPassword string, driveEncryptionKey string, driveEncryptionKeyIdentifier string) error {
	payload := make(map[string]any)
	payload["ControllerPassword"] = controllerPassword
	payload["DriveEncryptionKey"] = driveEncryptionKey
	payload["DriveEncryptionKeyIdentifier"] = driveEncryptionKeyIdentifier
	return s.Post(s.importForeignDrivesTarget, payload)
}

// RekeyExternalKey shall cause the controllers of the storage subsystem to request
// new encryption keys managed by an external key service.
func (s *Storage) RekeyExternalKey() error {
	payload := make(map[string]any)
	return s.Post(s.rekeyExternalKeyTarget, payload)
}

// ResetToDefaults shall reset the storage device. This action can impact other
// resources.
// resetType - This parameter shall contain the type of reset to defaults.
func (s *Storage) ResetToDefaults(resetType StorageResetToDefaultsType) error {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	return s.Post(s.resetToDefaultsTarget, payload)
}

// SetControllerPassword shall set the controller boot password for a host-based storage
// controller.
// currentPassword - This parameter shall contain the current controller
// password. Services may reject the action request if this parameter is not
// provided or the value supplied does not match the current password.
// newPassword - This parameter shall contain the new password to set for the
// controller.
// securityKey - This parameter shall contain the controller security key.
// Services may reject the action request if this parameter is not provided or
// the value provided does not match the security key for the controller.
func (s *Storage) SetControllerPassword(currentPassword string, newPassword string, securityKey string) error {
	payload := make(map[string]any)
	payload["CurrentPassword"] = currentPassword
	payload["NewPassword"] = newPassword
	payload["SecurityKey"] = securityKey
	return s.Post(s.setControllerPasswordTarget, payload)
}

// SetEncryptionKey shall set the local encryption key for the storage subsystem.
// currentEncryptionKey - This parameter shall contain the current local
// encryption key on the storage subsystem. Services may reject the action
// request if this parameter is not provided or the value supplied does not
// match the current encryption key.
// encryptionKey - This parameter shall contain the local encryption key to set
// on the storage subsystem.
// encryptionKeyIdentifier - This property shall contain the local encryption
// key identifier used by the storage subsystem.
func (s *Storage) SetEncryptionKey(currentEncryptionKey string, encryptionKey string, encryptionKeyIdentifier string) error {
	payload := make(map[string]any)
	payload["CurrentEncryptionKey"] = currentEncryptionKey
	payload["EncryptionKey"] = encryptionKey
	payload["EncryptionKeyIdentifier"] = encryptionKeyIdentifier
	return s.Post(s.setEncryptionKeyTarget, payload)
}

// BlockSecurityIDUnsupportedDrives gets the BlockSecurityIDUnsupportedDrives linked resources.
func (s *Storage) BlockSecurityIDUnsupportedDrives(client common.Client) ([]*Drive, error) {
	return common.GetObjects[Drive](client, s.blockSecurityIDUnsupportedDrives)
}

// BlockSecurityIDUpdateUnsuccessfulDrives gets the BlockSecurityIDUpdateUnsuccessfulDrives linked resources.
func (s *Storage) BlockSecurityIDUpdateUnsuccessfulDrives(client common.Client) ([]*Drive, error) {
	return common.GetObjects[Drive](client, s.blockSecurityIDUpdateUnsuccessfulDrives)
}

// Enclosures gets the Enclosures linked resources.
func (s *Storage) Enclosures(client common.Client) ([]*Chassis, error) {
	return common.GetObjects[Chassis](client, s.enclosures)
}

// HostingStorageSystems gets the HostingStorageSystems linked resources.
func (s *Storage) HostingStorageSystems(client common.Client) ([]*ComputerSystem, error) {
	return common.GetObjects[ComputerSystem](client, s.hostingStorageSystems)
}

// NVMeoFDiscoverySubsystems gets the NVMeoFDiscoverySubsystems linked resources.
func (s *Storage) NVMeoFDiscoverySubsystems(client common.Client) ([]*Storage, error) {
	return common.GetObjects[Storage](client, s.nVMeoFDiscoverySubsystems)
}

// SimpleStorage gets the SimpleStorage linked resource.
func (s *Storage) SimpleStorage(client common.Client) (*SimpleStorage, error) {
	if s.simpleStorage == "" {
		return nil, nil
	}
	return common.GetObject[SimpleStorage](client, s.simpleStorage)
}

// // StorageServices gets the StorageServices linked resources.
// func (s *Storage) StorageServices(client common.Client) ([]*StorageService, error) {
// 	return common.GetObjects[StorageService](client, s.storageServices)
// }

// UnassignedVolumes gets the UnassignedVolumes linked resources.
func (s *Storage) UnassignedVolumes(client common.Client) ([]*common.Volume, error) {
	return common.GetObjects[common.Volume](client, s.unassignedVolumes)
}

// Connections gets the Connections collection.
func (s *Storage) Connections(client common.Client) ([]*Connection, error) {
	if s.connections == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Connection](client, s.connections)
}

// // ConsistencyGroups gets the ConsistencyGroups collection.
// func (s *Storage) ConsistencyGroups(client common.Client) ([]*ConsistencyGroup, error) {
// 	if s.consistencyGroups == "" {
// 		return nil, nil
// 	}
// 	return common.GetCollectionObjects[ConsistencyGroup](client, s.consistencyGroups)
// }

// Controllers gets the Controllers collection.
func (s *Storage) Controllers(client common.Client) ([]*StorageController, error) {
	if s.controllers == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[StorageController](client, s.controllers)
}

// Drives gets the drives attached to the storage controllers that this
// resource represents.
func (s *Storage) Drives() ([]*Drive, error) {
	return common.GetObjects[Drive](s.GetClient(), s.drives)
}

// EndpointGroups gets the EndpointGroups collection.
func (s *Storage) EndpointGroups(client common.Client) ([]*EndpointGroup, error) {
	if s.endpointGroups == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[EndpointGroup](client, s.endpointGroups)
}

// // FileSystems gets the FileSystems collection.
// func (s *Storage) FileSystems(client common.Client) ([]*FileSystem, error) {
// 	if s.fileSystems == "" {
// 		return nil, nil
// 	}
// 	return common.GetCollectionObjects[FileSystem](client, s.fileSystems)
// }

// Metrics gets the Metrics linked resource.
func (s *Storage) Metrics(client common.Client) (*StorageMetrics, error) {
	if s.metrics == "" {
		return nil, nil
	}
	return common.GetObject[StorageMetrics](client, s.metrics)
}

// // StorageGroups gets the StorageGroups collection.
// func (s *Storage) StorageGroups(client common.Client) ([]*StorageGroup, error) {
// 	if s.storageGroups == "" {
// 		return nil, nil
// 	}
// 	return common.GetCollectionObjects[StorageGroup](client, s.storageGroups)
// }

// // StoragePools gets the StoragePools collection.
// func (s *Storage) StoragePools(client common.Client) ([]*StoragePool, error) {
// 	if s.storagePools == "" {
// 		return nil, nil
// 	}
// 	return common.GetCollectionObjects[StoragePool](client, s.storagePools)
// }

// Volumes gets the Volumes collection.
func (s *Storage) Volumes(client common.Client) ([]*common.Volume, error) {
	if s.volumes == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[common.Volume](client, s.volumes)
}

// MPF shall contain multiple physical function-related properties for storage
// controllers in a storage subsystem.
type MPF struct {
	// ConfiguredPhysicalFunctions shall contain the current number of physical
	// functions configured for the storage controller in this storage subsystem.
	//
	// Version added: v1.20.0
	ConfiguredPhysicalFunctions *int `json:",omitempty"`
	// MaximumSupportedPhysicalFunctions shall contain the maximum number of
	// physical functions supported by the storage controller in this storage
	// subsystem.
	//
	// Version added: v1.20.0
	MaximumSupportedPhysicalFunctions *int `json:",omitempty"`
	// VolumeAssignmentPolicy shall contain the current volume assignment policy
	// configured for the storage controller in this storage subsystem.
	//
	// Version added: v1.20.0
	VolumeAssignmentPolicy VolumeAssignmentPolicy
}

// NVMeSubsystemProperties shall contain information specific to NVMe
// Subsystems.
type NVMeSubsystemProperties struct {
	// ConfigurationLockState shall contain the configurable features that are able
	// to be locked from in-band usage on an NVMe subsystem and their current lock
	// state.
	//
	// Version added: v1.17.0
	ConfigurationLockState NVMeConfigurationLockState
	// MaxNamespacesSupported shall contain the maximum number of namespace
	// attachments supported by this NVMe Subsystem. If no maximum is specified,
	// this property should not be implemented.
	//
	// Version added: v1.16.0
	MaxNamespacesSupported *float64 `json:",omitempty"`
	// SharedNamespaceControllerAttachmentSupported shall indicate whether the
	// subsystem supports shared namespace controller attachment, allowing a shared
	// namespace to be attached concurrently to two or more controllers in an NVMe
	// Subsystem.
	//
	// Version added: v1.16.0
	SharedNamespaceControllerAttachmentSupported bool
}

// StorageStorageController shall represent a storage controller in the Redfish
// Specification.
type StorageStorageController struct {
	common.Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	//
	// Version added: v1.4.0
	assembly string
	// AssetTag shall track the storage controller for inventory purposes.
	AssetTag string
	// CacheSummary shall contain properties that describe the cache memory for
	// this resource.
	//
	// Version added: v1.5.0
	CacheSummary CacheSummary
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains certificates for device identity and
	// attestation.
	//
	// Version added: v1.10.0
	certificates string
	// ControllerRates shall contain all the rate settings available on the
	// controller.
	//
	// Version added: v1.7.0
	ControllerRates Rates
	// FirmwareVersion shall contain the firmware version as defined by the
	// manufacturer for the associated storage controller.
	FirmwareVersion string
	// Identifiers shall contain a list of all known durable names for the
	// associated storage controller.
	Identifiers []Identifier
	// Location shall contain the location information of the associated storage
	// controller.
	//
	// Version added: v1.4.0
	Location common.Location
	// Manufacturer shall contain the name of the organization responsible for
	// producing the storage controller. This organization may be the entity from
	// which the storage controller is purchased, but this is not necessarily true.
	Manufacturer string
	// Measurements shall contain an array of DSP0274-defined measurement blocks.
	//
	// Version added: v1.10.0
	//
	// Deprecated: v1.12.0
	// This property has been deprecated in favor of the 'ComponentIntegrity'
	// resource.
	Measurements []MeasurementBlock
	// MemberId shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// Model shall contain the name by which the manufacturer generally refers to
	// the storage controller.
	Model string
	// Name is the name of the resource or array element.
	//
	// Version added: v1.3.0
	Name string
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeInterface shall contain details on the PCIe interface that connects this
	// PCIe-based controller to its host.
	//
	// Version added: v1.5.0
	PCIeInterface PCIeInterface
	// PartNumber shall contain a part number assigned by the organization that is
	// responsible for producing or manufacturing the storage controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type
	// 'PortCollection'.
	//
	// Version added: v1.7.0
	ports string
	// SKU shall contain the stock-keeping unit number for this storage controller.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the storage controller.
	SerialNumber string
	// SpeedGbps shall represent the maximum supported speed of the storage bus
	// interface, in Gbit/s. The specified interface connects the controller to the
	// storage devices, not the controller to a host. For example, SAS bus not PCIe
	// host bus.
	SpeedGbps *float64 `json:",omitempty"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupportedControllerProtocols shall contain the supported set of protocols
	// for communicating with this storage controller.
	SupportedControllerProtocols []common.Protocol
	// SupportedDeviceProtocols shall contain the set of protocols this storage
	// controller can use to communicate with attached devices.
	SupportedDeviceProtocols []common.Protocol
	// SupportedRAIDTypes shall contain an array of all the RAID types supported by
	// this controller.
	//
	// Version added: v1.6.0
	SupportedRAIDTypes []RAIDType
	// importForeignDrivesTarget is the URL to send ImportForeignDrives requests.
	importForeignDrivesTarget string
	// rekeyExternalKeyTarget is the URL to send RekeyExternalKey requests.
	rekeyExternalKeyTarget string
	// resetToDefaultsTarget is the URL to send ResetToDefaults requests.
	resetToDefaultsTarget string
	// setControllerPasswordTarget is the URL to send SetControllerPassword requests.
	setControllerPasswordTarget string
	// setEncryptionKeyTarget is the URL to send SetEncryptionKey requests.
	setEncryptionKeyTarget string
	// blockSecurityIDUnsupportedDrives are the URIs for BlockSecurityIDUnsupportedDrives.
	blockSecurityIDUnsupportedDrives []string
	// blockSecurityIDUpdateUnsuccessfulDrives are the URIs for BlockSecurityIDUpdateUnsuccessfulDrives.
	blockSecurityIDUpdateUnsuccessfulDrives []string
	// enclosures are the URIs for Enclosures.
	enclosures []string
	// hostingStorageSystems are the URIs for HostingStorageSystems.
	hostingStorageSystems []string
	// nVMeoFDiscoverySubsystems are the URIs for NVMeoFDiscoverySubsystems.
	nVMeoFDiscoverySubsystems []string
	// simpleStorage is the URI for SimpleStorage.
	simpleStorage string
	// storageServices are the URIs for StorageServices.
	storageServices []string
	// unassignedVolumes are the URIs for UnassignedVolumes.
	unassignedVolumes []string
}

// UnmarshalJSON unmarshals a StorageStorageController object from the raw JSON.
func (s *StorageStorageController) UnmarshalJSON(b []byte) error {
	type temp StorageStorageController
	type sActions struct {
		ImportForeignDrives   common.ActionTarget `json:"#Storage.ImportForeignDrives"`
		RekeyExternalKey      common.ActionTarget `json:"#Storage.RekeyExternalKey"`
		ResetToDefaults       common.ActionTarget `json:"#Storage.ResetToDefaults"`
		SetControllerPassword common.ActionTarget `json:"#Storage.SetControllerPassword"`
		SetEncryptionKey      common.ActionTarget `json:"#Storage.SetEncryptionKey"`
	}
	type sLinks struct {
		BlockSecurityIDUnsupportedDrives        common.Links `json:"BlockSecurityIDUnsupportedDrives"`
		BlockSecurityIDUpdateUnsuccessfulDrives common.Links `json:"BlockSecurityIDUpdateUnsuccessfulDrives"`
		Enclosures                              common.Links `json:"Enclosures"`
		HostingStorageSystems                   common.Links `json:"HostingStorageSystems"`
		NVMeoFDiscoverySubsystems               common.Links `json:"NVMeoFDiscoverySubsystems"`
		SimpleStorage                           common.Link  `json:"SimpleStorage"`
		StorageServices                         common.Links `json:"StorageServices"`
		UnassignedVolumes                       common.Links `json:"UnassignedVolumes"`
	}
	var tmp struct {
		temp
		Actions      sActions
		Links        sLinks
		Assembly     common.Link `json:"assembly"`
		Certificates common.Link `json:"certificates"`
		Ports        common.Link `json:"ports"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = StorageStorageController(tmp.temp)

	// Extract the links to other entities for later
	s.importForeignDrivesTarget = tmp.Actions.ImportForeignDrives.Target
	s.rekeyExternalKeyTarget = tmp.Actions.RekeyExternalKey.Target
	s.resetToDefaultsTarget = tmp.Actions.ResetToDefaults.Target
	s.setControllerPasswordTarget = tmp.Actions.SetControllerPassword.Target
	s.setEncryptionKeyTarget = tmp.Actions.SetEncryptionKey.Target
	s.blockSecurityIDUnsupportedDrives = tmp.Links.BlockSecurityIDUnsupportedDrives.ToStrings()
	s.blockSecurityIDUpdateUnsuccessfulDrives = tmp.Links.BlockSecurityIDUpdateUnsuccessfulDrives.ToStrings()
	s.enclosures = tmp.Links.Enclosures.ToStrings()
	s.hostingStorageSystems = tmp.Links.HostingStorageSystems.ToStrings()
	s.nVMeoFDiscoverySubsystems = tmp.Links.NVMeoFDiscoverySubsystems.ToStrings()
	s.simpleStorage = tmp.Links.SimpleStorage.String()
	s.storageServices = tmp.Links.StorageServices.ToStrings()
	s.unassignedVolumes = tmp.Links.UnassignedVolumes.ToStrings()
	s.assembly = tmp.Assembly.String()
	s.certificates = tmp.Certificates.String()
	s.ports = tmp.Ports.String()

	return nil
}

// ImportForeignDrives shall import encrypted foreign drives to a host-based storage
// controller by supplying encryption access credentials. Upon successful
// completion, the 'EncryptionStatus' property of the affected 'Drive' shall
// change from 'Foreign' to 'Unlocked'.
// controllerPassword - This parameter shall contain the current controller
// password. Services may reject the action request if this parameter is not
// provided or the value supplied does not match the current password.
// driveEncryptionKey - This parameter shall contain the encryption key to
// unlock the drives.
// driveEncryptionKeyIdentifier - This parameter shall contain an encryption
// key identifier to filter the drives that are imported. If this parameter is
// not provided, the service shall not perform filtering of foreign drives.
func (s *StorageStorageController) ImportForeignDrives(controllerPassword string, driveEncryptionKey string, driveEncryptionKeyIdentifier string) error {
	payload := make(map[string]any)
	payload["ControllerPassword"] = controllerPassword
	payload["DriveEncryptionKey"] = driveEncryptionKey
	payload["DriveEncryptionKeyIdentifier"] = driveEncryptionKeyIdentifier
	return s.Post(s.importForeignDrivesTarget, payload)
}

// RekeyExternalKey shall cause the controllers of the storage subsystem to request
// new encryption keys managed by an external key service.
func (s *StorageStorageController) RekeyExternalKey() error {
	payload := make(map[string]any)
	return s.Post(s.rekeyExternalKeyTarget, payload)
}

// ResetToDefaults shall reset the storage device. This action can impact other
// resources.
// resetType - This parameter shall contain the type of reset to defaults.
func (s *StorageStorageController) ResetToDefaults(resetType StorageResetToDefaultsType) error {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	return s.Post(s.resetToDefaultsTarget, payload)
}

// SetControllerPassword shall set the controller boot password for a host-based storage
// controller.
// currentPassword - This parameter shall contain the current controller
// password. Services may reject the action request if this parameter is not
// provided or the value supplied does not match the current password.
// newPassword - This parameter shall contain the new password to set for the
// controller.
// securityKey - This parameter shall contain the controller security key.
// Services may reject the action request if this parameter is not provided or
// the value provided does not match the security key for the controller.
func (s *StorageStorageController) SetControllerPassword(currentPassword string, newPassword string, securityKey string) error {
	payload := make(map[string]any)
	payload["CurrentPassword"] = currentPassword
	payload["NewPassword"] = newPassword
	payload["SecurityKey"] = securityKey
	return s.Post(s.setControllerPasswordTarget, payload)
}

// SetEncryptionKey shall set the local encryption key for the storage subsystem.
// currentEncryptionKey - This parameter shall contain the current local
// encryption key on the storage subsystem. Services may reject the action
// request if this parameter is not provided or the value supplied does not
// match the current encryption key.
// encryptionKey - This parameter shall contain the local encryption key to set
// on the storage subsystem.
// encryptionKeyIdentifier - This property shall contain the local encryption
// key identifier used by the storage subsystem.
func (s *StorageStorageController) SetEncryptionKey(currentEncryptionKey string, encryptionKey string, encryptionKeyIdentifier string) error {
	payload := make(map[string]any)
	payload["CurrentEncryptionKey"] = currentEncryptionKey
	payload["EncryptionKey"] = encryptionKey
	payload["EncryptionKeyIdentifier"] = encryptionKeyIdentifier
	return s.Post(s.setEncryptionKeyTarget, payload)
}

// BlockSecurityIDUnsupportedDrives gets the BlockSecurityIDUnsupportedDrives linked resources.
func (s *StorageStorageController) BlockSecurityIDUnsupportedDrives(client common.Client) ([]*Drive, error) {
	return common.GetObjects[Drive](client, s.blockSecurityIDUnsupportedDrives)
}

// BlockSecurityIDUpdateUnsuccessfulDrives gets the BlockSecurityIDUpdateUnsuccessfulDrives linked resources.
func (s *StorageStorageController) BlockSecurityIDUpdateUnsuccessfulDrives(client common.Client) ([]*Drive, error) {
	return common.GetObjects[Drive](client, s.blockSecurityIDUpdateUnsuccessfulDrives)
}

// Enclosures gets the Enclosures linked resources.
func (s *StorageStorageController) Enclosures(client common.Client) ([]*Chassis, error) {
	return common.GetObjects[Chassis](client, s.enclosures)
}

// HostingStorageSystems gets the HostingStorageSystems linked resources.
func (s *StorageStorageController) HostingStorageSystems(client common.Client) ([]*ComputerSystem, error) {
	return common.GetObjects[ComputerSystem](client, s.hostingStorageSystems)
}

// NVMeoFDiscoverySubsystems gets the NVMeoFDiscoverySubsystems linked resources.
func (s *StorageStorageController) NVMeoFDiscoverySubsystems(client common.Client) ([]*Storage, error) {
	return common.GetObjects[Storage](client, s.nVMeoFDiscoverySubsystems)
}

// SimpleStorage gets the SimpleStorage linked resource.
func (s *StorageStorageController) SimpleStorage(client common.Client) (*SimpleStorage, error) {
	if s.simpleStorage == "" {
		return nil, nil
	}
	return common.GetObject[SimpleStorage](client, s.simpleStorage)
}

// StorageServices gets the StorageServices linked resources.
// func (s *StorageStorageController) StorageServices(client common.Client) ([]*StorageService, error) {
// 	return common.GetObjects[StorageService](client, s.storageServices)
// }

// UnassignedVolumes gets the UnassignedVolumes linked resources.
func (s *StorageStorageController) UnassignedVolumes(client common.Client) ([]*common.Volume, error) {
	return common.GetObjects[common.Volume](client, s.unassignedVolumes)
}

// Assembly gets the Assembly linked resource.
func (s *StorageStorageController) Assembly(client common.Client) (*Assembly, error) {
	if s.assembly == "" {
		return nil, nil
	}
	return common.GetObject[Assembly](client, s.assembly)
}

// Certificates gets the Certificates collection.
func (s *StorageStorageController) Certificates(client common.Client) ([]*Certificate, error) {
	if s.certificates == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Certificate](client, s.certificates)
}

// Ports gets the Ports collection.
func (s *StorageStorageController) Ports(client common.Client) ([]*Port, error) {
	if s.ports == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Port](client, s.ports)
}
