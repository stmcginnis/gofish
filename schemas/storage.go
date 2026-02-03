//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #Storage.v1_21_0.Storage

package schemas

import (
	"encoding/json"
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

type ConfigLockOptions string

const (
	// UnlockedConfigLockOptions The command is supported, able to be locked, and
	// is currently unlocked.
	UnlockedConfigLockOptions ConfigLockOptions = "Unlocked"
	// LockedConfigLockOptions The command is supported and is currently locked.
	LockedConfigLockOptions ConfigLockOptions = "Locked"
	// LockdownUnsupportedConfigLockOptions The command is supported but is not
	// able to be locked.
	LockdownUnsupportedConfigLockOptions ConfigLockOptions = "LockdownUnsupported"
	// CommandUnsupportedConfigLockOptions The command is not supported, therefore
	// lockdown does not apply.
	CommandUnsupportedConfigLockOptions ConfigLockOptions = "CommandUnsupported"
)

type ConfigurationLock string

const (
	// EnabledConfigurationLock shall indicate in-band configuration requests are
	// locked as specified by 'TargetConfigurationLockLevel'.
	EnabledConfigurationLock ConfigurationLock = "Enabled"
	// DisabledConfigurationLock shall indicate in-band configuration requests are
	// not locked.
	DisabledConfigurationLock ConfigurationLock = "Disabled"
	// PartialConfigurationLock shall indicate some in-band configuration requests
	// are not locked while others are locked.
	PartialConfigurationLock ConfigurationLock = "Partial"
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

type NMVeMinimumRequiredResetType string

const (
	// NoResetNMVeMinimumRequiredResetType No reset required.
	NoResetNMVeMinimumRequiredResetType NMVeMinimumRequiredResetType = "NoReset"
	// ControllerLevelNMVeMinimumRequiredResetType is a Controller Level Reset is
	// required.
	ControllerLevelNMVeMinimumRequiredResetType NMVeMinimumRequiredResetType = "ControllerLevel"
	// LimitedControllerLevelNMVeMinimumRequiredResetType is a Controller Level
	// Reset other than a Controller Level Reset initiated by a Controller Reset is
	// required.
	LimitedControllerLevelNMVeMinimumRequiredResetType NMVeMinimumRequiredResetType = "LimitedControllerLevel"
	// NVMSubsystemResetNMVeMinimumRequiredResetType is an NVM Subsystem Reset is
	// required.
	NVMSubsystemResetNMVeMinimumRequiredResetType NMVeMinimumRequiredResetType = "NVMSubsystemReset"
	// PowerCycleNMVeMinimumRequiredResetType is a power cycle is required.
	PowerCycleNMVeMinimumRequiredResetType NMVeMinimumRequiredResetType = "PowerCycle"
)

type NMVePersonalityKeyAlgorithm string

const (
	// HMACSHA384NMVePersonalityKeyAlgorithm HMAC-SHA-384.
	HMACSHA384NMVePersonalityKeyAlgorithm NMVePersonalityKeyAlgorithm = "HMAC_SHA384"
)

type NMVeUnfreezeAuthMode string

const (
	// ProgrammedKeyNMVeUnfreezeAuthMode Programmable Key Authentication.
	ProgrammedKeyNMVeUnfreezeAuthMode NMVeUnfreezeAuthMode = "ProgrammedKey"
	// PhysicalIDNMVeUnfreezeAuthMode Physical Credential Authentication.
	PhysicalIDNMVeUnfreezeAuthMode NMVeUnfreezeAuthMode = "PhysicalId"
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

type TargetConfigurationLockLevel string

const (
	// BaselineTargetConfigurationLockLevel The standard configuration lock level,
	// corresponding to applying firmware, updating security keys, and modifying
	// other hardware settings. It does not include managing the volumes or data
	// within the storage subsystem.
	BaselineTargetConfigurationLockLevel TargetConfigurationLockLevel = "Baseline"
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
	Entity
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
	// DrivesCount
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
	// subsystems. This property shall only be present if this resource represents
	// an NVMe subsystem.
	//
	// Version added: v1.16.0
	NVMeSubsystemProperties NVMeSubsystemProperties
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Redundancy shall contain redundancy information for the storage subsystem.
	redundancy string
	// RedundancyCount
	RedundancyCount int `json:"Redundancy@odata.count"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// StorageControllers shall contain a set of the storage controllers that this
	// resource represents.
	//
	// Deprecated: v1.13.0
	// This property has been deprecated in favor of 'Controllers' to allow for
	// storage controllers to be represented as their own resources.
	StorageControllers []StorageController
	// StorageControllersCount
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
	// freezePersonalityTarget is the URL to send FreezePersonality requests.
	freezePersonalityTarget string
	// getPersonalityNonceTarget is the URL to send GetPersonalityNonce requests.
	getPersonalityNonceTarget string
	// importForeignDrivesTarget is the URL to send ImportForeignDrives requests.
	importForeignDrivesTarget string
	// rekeyExternalKeyTarget is the URL to send RekeyExternalKey requests.
	rekeyExternalKeyTarget string
	// resetToDefaultsTarget is the URL to send ResetToDefaults requests.
	resetToDefaultsTarget string
	// revertPersonalitiesToDefaultsTarget is the URL to send RevertPersonalitiesToDefaults requests.
	revertPersonalitiesToDefaultsTarget string
	// setControllerPasswordTarget is the URL to send SetControllerPassword requests.
	setControllerPasswordTarget string
	// setEncryptionKeyTarget is the URL to send SetEncryptionKey requests.
	setEncryptionKeyTarget string
	// setPersonalityKeyTarget is the URL to send SetPersonalityKey requests.
	setPersonalityKeyTarget string
	// unfreezePersonalityTarget is the URL to send UnfreezePersonality requests.
	unfreezePersonalityTarget string
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
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Storage object from the raw JSON.
func (s *Storage) UnmarshalJSON(b []byte) error {
	type temp Storage
	type sActions struct {
		FreezePersonality             ActionTarget `json:"#Storage.FreezePersonality"`
		GetPersonalityNonce           ActionTarget `json:"#Storage.GetPersonalityNonce"`
		ImportForeignDrives           ActionTarget `json:"#Storage.ImportForeignDrives"`
		RekeyExternalKey              ActionTarget `json:"#Storage.RekeyExternalKey"`
		ResetToDefaults               ActionTarget `json:"#Storage.ResetToDefaults"`
		RevertPersonalitiesToDefaults ActionTarget `json:"#Storage.RevertPersonalitiesToDefaults"`
		SetControllerPassword         ActionTarget `json:"#Storage.SetControllerPassword"`
		SetEncryptionKey              ActionTarget `json:"#Storage.SetEncryptionKey"`
		SetPersonalityKey             ActionTarget `json:"#Storage.SetPersonalityKey"`
		UnfreezePersonality           ActionTarget `json:"#Storage.UnfreezePersonality"`
	}
	type sLinks struct {
		BlockSecurityIDUnsupportedDrives        Links `json:"BlockSecurityIDUnsupportedDrives"`
		BlockSecurityIDUpdateUnsuccessfulDrives Links `json:"BlockSecurityIDUpdateUnsuccessfulDrives"`
		Enclosures                              Links `json:"Enclosures"`
		HostingStorageSystems                   Links `json:"HostingStorageSystems"`
		NVMeoFDiscoverySubsystems               Links `json:"NVMeoFDiscoverySubsystems"`
		SimpleStorage                           Link  `json:"SimpleStorage"`
		StorageServices                         Links `json:"StorageServices"`
		UnassignedVolumes                       Links `json:"UnassignedVolumes"`
	}
	var tmp struct {
		temp
		Actions           sActions
		Links             sLinks
		Connections       Link  `json:"Connections"`
		ConsistencyGroups Link  `json:"ConsistencyGroups"`
		Controllers       Link  `json:"Controllers"`
		Drives            Links `json:"Drives"`
		EndpointGroups    Link  `json:"EndpointGroups"`
		FileSystems       Link  `json:"FileSystems"`
		Metrics           Link  `json:"Metrics"`
		Redundancy        Link  `json:"Redundancy"`
		StorageGroups     Link  `json:"StorageGroups"`
		StoragePools      Link  `json:"StoragePools"`
		Volumes           Link  `json:"Volumes"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = Storage(tmp.temp)

	// Extract the links to other entities for later
	s.freezePersonalityTarget = tmp.Actions.FreezePersonality.Target
	s.getPersonalityNonceTarget = tmp.Actions.GetPersonalityNonce.Target
	s.importForeignDrivesTarget = tmp.Actions.ImportForeignDrives.Target
	s.rekeyExternalKeyTarget = tmp.Actions.RekeyExternalKey.Target
	s.resetToDefaultsTarget = tmp.Actions.ResetToDefaults.Target
	s.revertPersonalitiesToDefaultsTarget = tmp.Actions.RevertPersonalitiesToDefaults.Target
	s.setControllerPasswordTarget = tmp.Actions.SetControllerPassword.Target
	s.setEncryptionKeyTarget = tmp.Actions.SetEncryptionKey.Target
	s.setPersonalityKeyTarget = tmp.Actions.SetPersonalityKey.Target
	s.unfreezePersonalityTarget = tmp.Actions.UnfreezePersonality.Target
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
	s.redundancy = tmp.Redundancy.String()
	s.storageGroups = tmp.StorageGroups.String()
	s.storagePools = tmp.StoragePools.String()
	s.volumes = tmp.Volumes.String()

	// This is a read/write object, so we need to save the raw object data for later
	s.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *Storage) Update() error {
	readWriteFields := []string{
		"AutoVolumeCreate",
		"BlockSecurityIDPolicy",
		"ConfigurationLock",
		"EncryptionMode",
		"HotspareActivationPolicy",
		"TargetConfigurationLockLevel",
	}

	return s.UpdateFromRawData(s, s.RawData, readWriteFields)
}

// GetStorage will get a Storage instance from the service.
func GetStorage(c Client, uri string) (*Storage, error) {
	return GetObject[Storage](c, uri)
}

// ListReferencedStorages gets the collection of Storage from
// a provided reference.
func ListReferencedStorages(c Client, link string) ([]*Storage, error) {
	return GetCollectionObjects[Storage](c, link)
}

// This action shall freeze a personality for the NVMe subsystem as defined by
// the 'Configurable Device Personality' feature in the NVMe Base
// Specification.
// personality - This parameter shall contain the personality to freeze as
// defined by the 'Personality Identifier List' figure in the NVMe Base
// Specification. The value '255' shall indicate all personalities.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *Storage) FreezePersonality(personality uint) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Personality"] = personality
	resp, taskInfo, err := PostWithTask(s.client,
		s.freezePersonalityTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall obtain a nonce required to unfreeze personalities with a
// programmed key as defined by the 'CDP Random Nonce Data' clause in the NVMe
// Base Specification.
func (s *Storage) GetPersonalityNonce() (*GetPersonalityNonceResponse, error) {
	payload := make(map[string]any)

	resp, err := s.PostWithResponse(s.getPersonalityNonceTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result GetPersonalityNonceResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// This action shall import encrypted foreign drives to a host-based storage
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
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *Storage) ImportForeignDrives(controllerPassword string, driveEncryptionKey string, driveEncryptionKeyIdentifier string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ControllerPassword"] = controllerPassword
	payload["DriveEncryptionKey"] = driveEncryptionKey
	payload["DriveEncryptionKeyIdentifier"] = driveEncryptionKeyIdentifier
	resp, taskInfo, err := PostWithTask(s.client,
		s.importForeignDrivesTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall cause the controllers of the storage subsystem to request
// new encryption keys managed by an external key service.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *Storage) RekeyExternalKey() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(s.client,
		s.rekeyExternalKeyTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset the storage device. This action can impact other
// resources.
// resetType - This parameter shall contain the type of reset to defaults.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *Storage) ResetToDefaults(resetType ResetToDefaultsType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(s.client,
		s.resetToDefaultsTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall revert all personalities to their manufacturing default
// state for the NVMe subsystem as defined by the 'Manufacturing Default
// Personality' clause of the NVMe Base Specification.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *Storage) RevertPersonalitiesToDefaults() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(s.client,
		s.revertPersonalitiesToDefaultsTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall set the controller boot password for a host-based storage
// controller.
// currentPassword - This parameter shall contain the current controller
// password. Services may reject the action request if this parameter is not
// provided or the value supplied does not match the current password.
// newPassword - This parameter shall contain the new password to set for the
// controller.
// securityKey - This parameter shall contain the controller security key.
// Services may reject the action request if this parameter is not provided or
// the value provided does not match the security key for the controller.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *Storage) SetControllerPassword(currentPassword string, newPassword string, securityKey string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["CurrentPassword"] = currentPassword
	payload["NewPassword"] = newPassword
	payload["SecurityKey"] = securityKey
	resp, taskInfo, err := PostWithTask(s.client,
		s.setControllerPasswordTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall set the local encryption key for the storage subsystem.
// currentEncryptionKey - This parameter shall contain the current local
// encryption key on the storage subsystem. Services may reject the action
// request if this parameter is not provided or the value supplied does not
// match the current encryption key.
// encryptionKey - This parameter shall contain the local encryption key to set
// on the storage subsystem.
// encryptionKeyIdentifier - This property shall contain the local encryption
// key identifier used by the storage subsystem.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *Storage) SetEncryptionKey(currentEncryptionKey string, encryptionKey string, encryptionKeyIdentifier string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["CurrentEncryptionKey"] = currentEncryptionKey
	payload["EncryptionKey"] = encryptionKey
	payload["EncryptionKeyIdentifier"] = encryptionKeyIdentifier
	resp, taskInfo, err := PostWithTask(s.client,
		s.setEncryptionKeyTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall set the personality key for the NVMe subsystem as defined
// by the 'Programmable Key Authentication Considerations' clause in the NVMe
// Base Specification.
// algorithm - This parameter shall contain the type of authentication
// algorithm for the key as defined by the 'CDP Authentication Algorithm' field
// in the NVMe Base Specification.
// key - This parameter shall contain the personality key to set on the NVMe
// subsystem as a hex-encoded string.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *Storage) SetPersonalityKey(algorithm NMVePersonalityKeyAlgorithm, key string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Algorithm"] = algorithm
	payload["Key"] = key
	resp, taskInfo, err := PostWithTask(s.client,
		s.setPersonalityKeyTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// StorageUnfreezePersonalityParameters holds the parameters for the UnfreezePersonality action.
type StorageUnfreezePersonalityParameters struct {
	// AuthenticationMethod shall contain the authentication method for the
	// unfreeze operation as defined by the 'Authenticated Unfreeze Support' field
	// in the NVMe Base Specification.
	AuthenticationMethod NMVeUnfreezeAuthMode `json:"AuthenticationMethod,omitempty"`
	// Key shall contain the authentication key for the unfreeze operation as a
	// hex-encoded string. If 'AuthenticationMethod' contains 'ProgrammedKey', this
	// parameter contains a key that is derived from the original key provided in
	// the 'SetPersonalityKey' action. If 'AuthenticationMethod' contains
	// 'PhysicalId', this parameter contains the physical secure ID (PSID) of the
	// drive converted to a hex-encoded string.
	Key string `json:"Key,omitempty"`
	// Nonce shall contain the personality nonce provided by the
	// 'GetPersonalityNonce' action. This parameter shall be required if
	// 'AuthenticationMethod' contains 'ProgrammedKey'.
	Nonce string `json:"Nonce,omitempty"`
	// Personality shall contain the personality to freeze as defined by the
	// 'Personality Identifier List' figure in the NVMe Base Specification. The
	// value '255' shall indicate all personalities.
	Personality uint `json:"Personality,omitempty"`
}

// This action shall unfreeze a personality for the NVMe subsystem as defined
// by the 'Authenticated Unfreeze Operation' clause in the NVMe Base
// Specification.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *Storage) UnfreezePersonality(params *StorageUnfreezePersonalityParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(s.client,
		s.unfreezePersonalityTarget, params, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// BlockSecurityIDUnsupportedDrives gets the BlockSecurityIDUnsupportedDrives linked resources.
func (s *Storage) BlockSecurityIDUnsupportedDrives() ([]*Drive, error) {
	return GetObjects[Drive](s.client, s.blockSecurityIDUnsupportedDrives)
}

// BlockSecurityIDUpdateUnsuccessfulDrives gets the BlockSecurityIDUpdateUnsuccessfulDrives linked resources.
func (s *Storage) BlockSecurityIDUpdateUnsuccessfulDrives() ([]*Drive, error) {
	return GetObjects[Drive](s.client, s.blockSecurityIDUpdateUnsuccessfulDrives)
}

// Enclosures gets the Enclosures linked resources.
func (s *Storage) Enclosures() ([]*Chassis, error) {
	return GetObjects[Chassis](s.client, s.enclosures)
}

// HostingStorageSystems gets the HostingStorageSystems linked resources.
func (s *Storage) HostingStorageSystems() ([]*ComputerSystem, error) {
	return GetObjects[ComputerSystem](s.client, s.hostingStorageSystems)
}

// NVMeoFDiscoverySubsystems gets the NVMeoFDiscoverySubsystems linked resources.
func (s *Storage) NVMeoFDiscoverySubsystems() ([]*Storage, error) {
	return GetObjects[Storage](s.client, s.nVMeoFDiscoverySubsystems)
}

// SimpleStorage gets the SimpleStorage linked resource.
func (s *Storage) SimpleStorage() (*SimpleStorage, error) {
	if s.simpleStorage == "" {
		return nil, nil
	}
	return GetObject[SimpleStorage](s.client, s.simpleStorage)
}

// StorageServices gets the StorageServices linked resources.
func (s *Storage) StorageServices() ([]*StorageService, error) {
	return GetObjects[StorageService](s.client, s.storageServices)
}

// UnassignedVolumes gets the UnassignedVolumes linked resources.
func (s *Storage) UnassignedVolumes() ([]*Volume, error) {
	return GetObjects[Volume](s.client, s.unassignedVolumes)
}

// Connections gets the Connections collection.
func (s *Storage) Connections() ([]*Connection, error) {
	if s.connections == "" {
		return nil, nil
	}
	return GetCollectionObjects[Connection](s.client, s.connections)
}

// ConsistencyGroups gets the ConsistencyGroups collection.
func (s *Storage) ConsistencyGroups() ([]*ConsistencyGroup, error) {
	if s.consistencyGroups == "" {
		return nil, nil
	}
	return GetCollectionObjects[ConsistencyGroup](s.client, s.consistencyGroups)
}

// Controllers gets the Controllers collection.
func (s *Storage) Controllers() ([]*StorageController, error) {
	if s.controllers == "" {
		return nil, nil
	}
	return GetCollectionObjects[StorageController](s.client, s.controllers)
}

// Drives gets the Drives linked resources.
func (s *Storage) Drives() ([]*Drive, error) {
	return GetObjects[Drive](s.client, s.drives)
}

// EndpointGroups gets the EndpointGroups collection.
func (s *Storage) EndpointGroups() ([]*EndpointGroup, error) {
	if s.endpointGroups == "" {
		return nil, nil
	}
	return GetCollectionObjects[EndpointGroup](s.client, s.endpointGroups)
}

// FileSystems gets the FileSystems collection.
func (s *Storage) FileSystems() ([]*FileSystem, error) {
	if s.fileSystems == "" {
		return nil, nil
	}
	return GetCollectionObjects[FileSystem](s.client, s.fileSystems)
}

// Metrics gets the Metrics linked resource.
func (s *Storage) Metrics() (*StorageMetrics, error) {
	if s.metrics == "" {
		return nil, nil
	}
	return GetObject[StorageMetrics](s.client, s.metrics)
}

// Redundancy gets the Redundancy linked resource.
func (s *Storage) Redundancy() (*Redundancy, error) {
	if s.redundancy == "" {
		return nil, nil
	}
	return GetObject[Redundancy](s.client, s.redundancy)
}

// StorageGroups gets the StorageGroups collection.
func (s *Storage) StorageGroups() ([]*StorageGroup, error) {
	if s.storageGroups == "" {
		return nil, nil
	}
	return GetCollectionObjects[StorageGroup](s.client, s.storageGroups)
}

// StoragePools gets the StoragePools collection.
func (s *Storage) StoragePools() ([]*StoragePool, error) {
	if s.storagePools == "" {
		return nil, nil
	}
	return GetCollectionObjects[StoragePool](s.client, s.storagePools)
}

// Volumes gets the Volumes collection.
func (s *Storage) Volumes() ([]*Volume, error) {
	if s.volumes == "" {
		return nil, nil
	}
	return GetCollectionObjects[Volume](s.client, s.volumes)
}

// CacheSummary shall contain properties that describe the cache memory for a
// storage controller.
type CacheSummary struct {
	// PersistentCacheSizeMiB shall contain the amount of cache memory that is
	// persistent as measured in mebibytes. This size shall be less than or equal
	// to the 'TotalCacheSizeMiB'.
	//
	// Version added: v1.5.0
	PersistentCacheSizeMiB *uint `json:",omitempty"`
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.5.0
	Status Status
	// TotalCacheSizeMiB shall contain the amount of configured cache memory as
	// measured in mebibytes.
	//
	// Version added: v1.5.0
	TotalCacheSizeMiB *uint `json:",omitempty"`
}

// GetPersonalityNonceResponse shall contain the personality nonce from a drive.
type GetPersonalityNonceResponse struct {
	// Nonce shall contain the personality nonce as a hex-encoded string as defined
	// by the 'CDP Random Nonce Data' clause of the NVMe Base Specification.
	//
	// Version added: v1.21.0
	Nonce string
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

// NVMeConfigurationLockState shall contain the configurable features that are
// able to be locked on an NVMe subsystem and their current lock state.
type NVMeConfigurationLockState struct {
	// FirmwareCommit shall contain the lock state of the NVMe-defined Firmware
	// Commit command.
	//
	// Version added: v1.17.0
	FirmwareCommit ConfigLockOptions
	// FirmwareImageDownload shall contain the lock state of the NVMe-defined
	// Firmware Image Download command.
	//
	// Version added: v1.17.0
	FirmwareImageDownload ConfigLockOptions
	// Lockdown shall contain the lock state of the NVMe-defined Lockdown command.
	//
	// Version added: v1.17.0
	Lockdown ConfigLockOptions
	// SecuritySend shall contain the lock state of the NVMe-defined Security Send
	// command.
	//
	// Version added: v1.17.0
	SecuritySend ConfigLockOptions
	// VPDWrite shall contain the lock state of the NVMe-MI-defined VPD Write
	// command.
	//
	// Version added: v1.17.0
	VPDWrite ConfigLockOptions
}

// NVMePersonality shall contain an active personality setting of an NVMe
// subsystem as defined by the NVMe Device Personalities log page and NVMe
// Configurable Device Personality feature.
type NVMePersonality struct {
	// ChangeAffectsUserData shall indicate whether changing this personality
	// affects user data as defined by the 'Personality Settings Change User Data
	// Effect' field in the NVMe Base Specification.
	//
	// Version added: v1.21.0
	ChangeAffectsUserData bool
	// Data shall contain the personality data as a hex-encoded string.
	//
	// Version added: v1.21.0
	Data string
	// Default shall indicate whether this personality is set to the manufacturing
	// default values as defined by the 'Personality Manufacturing Default Settings
	// State' field in the NVMe Base Specification.
	//
	// Version added: v1.21.0
	Default bool
	// Frozen shall indicate whether this personality is frozen as defined by the
	// 'Personality Freeze State' field in the NVMe Base Specification.
	//
	// Version added: v1.21.0
	Frozen bool
	// Identifier shall be unique within the managed ecosystem.
	//
	// Version added: v1.21.0
	Identifier Identifier
	// MinimumRequiredResetType shall contain the minimum reset type required to
	// change the personality settings as defined by the 'Minimum Required Reset
	// Type' field in the NVMe Base Specification.
	//
	// Version added: v1.21.0
	MinimumRequiredResetType NMVeMinimumRequiredResetType
	// Pending shall indicate whether this personality contains pending changes as
	// defined by the 'Pending Personality Settings Change' field in the NVMe Base
	// Specification.
	//
	// Version added: v1.21.0
	Pending bool
	// UnfreezeAuthenticationModes shall contain the authentication modes supported
	// to unfreeze this personality as defined by the 'Authenticated Unfreeze
	// Support' field in the NVMe Base Specification.
	//
	// Version added: v1.21.0
	UnfreezeAuthenticationModes []NMVeUnfreezeAuthMode
}

// NVMeSubsystemProperties shall contain information specific to NVMe
// subsystems.
type NVMeSubsystemProperties struct {
	// ActivePersonalities shall contain the active personality settings of the
	// NVMe subsystem as defined by the NVMe Device Personalities log page and NVMe
	// Configurable Device Personality feature.
	//
	// Version added: v1.21.0
	ActivePersonalities []NVMePersonality
	// ConfigurationLockState shall contain the configurable features that are able
	// to be locked from in-band usage on an NVMe subsystem and their current lock
	// state.
	//
	// Version added: v1.17.0
	ConfigurationLockState NVMeConfigurationLockState
	// MaxNamespacesSupported shall contain the maximum number of namespace
	// attachments supported by this NVMe subsystem. If no maximum is specified,
	// this property should not be implemented.
	//
	// Version added: v1.16.0
	MaxNamespacesSupported *float64 `json:",omitempty"`
	// SharedNamespaceControllerAttachmentSupported shall indicate whether the
	// subsystem supports shared namespace controller attachment, allowing a shared
	// namespace to be attached concurrently to two or more controllers in an NVMe
	// subsystem.
	//
	// Version added: v1.16.0
	SharedNamespaceControllerAttachmentSupported bool
}

// Rates shall contain all the rate settings available on the controller.
type Rates struct {
	// ConsistencyCheckRatePercent shall contain the percentage of controller
	// resources used for checking data consistency on volumes.
	//
	// Version added: v1.7.0
	ConsistencyCheckRatePercent *uint `json:",omitempty"`
	// RebuildRatePercent shall contain the percentage of controller resources used
	// for rebuilding volumes.
	//
	// Version added: v1.7.0
	RebuildRatePercent *uint `json:",omitempty"`
	// TransformationRatePercent shall contain the percentage of controller
	// resources used for transforming volumes.
	//
	// Version added: v1.7.0
	TransformationRatePercent *uint `json:",omitempty"`
}

// StorageStorageController shall represent a storage controller in the Redfish
// Specification.
type StorageStorageController struct {
	Entity
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
	Location Location
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
	// MemberID shall contain the unique identifier for this member within an
	// array. For services supporting Redfish v1.6 or higher, this value shall
	// contain the zero-based array index.
	MemberID string `json:"MemberId"`
	// Model shall contain the name by which the manufacturer generally refers to
	// the storage controller.
	Model string
	// OEM shall contain the OEM extensions. All values for properties that this
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
	Status Status
	// SupportedControllerProtocols shall contain the supported set of protocols
	// for communicating with this storage controller.
	SupportedControllerProtocols []Protocol
	// SupportedDeviceProtocols shall contain the set of protocols this storage
	// controller can use to communicate with attached devices.
	SupportedDeviceProtocols []Protocol
	// SupportedRAIDTypes shall contain an array of all the RAID types supported by
	// this controller.
	//
	// Version added: v1.6.0
	SupportedRAIDTypes []RAIDType
	// freezePersonalityTarget is the URL to send FreezePersonality requests.
	freezePersonalityTarget string
	// getPersonalityNonceTarget is the URL to send GetPersonalityNonce requests.
	getPersonalityNonceTarget string
	// importForeignDrivesTarget is the URL to send ImportForeignDrives requests.
	importForeignDrivesTarget string
	// rekeyExternalKeyTarget is the URL to send RekeyExternalKey requests.
	rekeyExternalKeyTarget string
	// resetToDefaultsTarget is the URL to send ResetToDefaults requests.
	resetToDefaultsTarget string
	// revertPersonalitiesToDefaultsTarget is the URL to send RevertPersonalitiesToDefaults requests.
	revertPersonalitiesToDefaultsTarget string
	// setControllerPasswordTarget is the URL to send SetControllerPassword requests.
	setControllerPasswordTarget string
	// setEncryptionKeyTarget is the URL to send SetEncryptionKey requests.
	setEncryptionKeyTarget string
	// setPersonalityKeyTarget is the URL to send SetPersonalityKey requests.
	setPersonalityKeyTarget string
	// unfreezePersonalityTarget is the URL to send UnfreezePersonality requests.
	unfreezePersonalityTarget string
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
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a StorageStorageController object from the raw JSON.
func (s *StorageStorageController) UnmarshalJSON(b []byte) error {
	type temp StorageStorageController
	type sActions struct {
		FreezePersonality             ActionTarget `json:"#Storage.FreezePersonality"`
		GetPersonalityNonce           ActionTarget `json:"#Storage.GetPersonalityNonce"`
		ImportForeignDrives           ActionTarget `json:"#Storage.ImportForeignDrives"`
		RekeyExternalKey              ActionTarget `json:"#Storage.RekeyExternalKey"`
		ResetToDefaults               ActionTarget `json:"#Storage.ResetToDefaults"`
		RevertPersonalitiesToDefaults ActionTarget `json:"#Storage.RevertPersonalitiesToDefaults"`
		SetControllerPassword         ActionTarget `json:"#Storage.SetControllerPassword"`
		SetEncryptionKey              ActionTarget `json:"#Storage.SetEncryptionKey"`
		SetPersonalityKey             ActionTarget `json:"#Storage.SetPersonalityKey"`
		UnfreezePersonality           ActionTarget `json:"#Storage.UnfreezePersonality"`
	}
	type sLinks struct {
		BlockSecurityIDUnsupportedDrives        Links `json:"BlockSecurityIDUnsupportedDrives"`
		BlockSecurityIDUpdateUnsuccessfulDrives Links `json:"BlockSecurityIDUpdateUnsuccessfulDrives"`
		Enclosures                              Links `json:"Enclosures"`
		HostingStorageSystems                   Links `json:"HostingStorageSystems"`
		NVMeoFDiscoverySubsystems               Links `json:"NVMeoFDiscoverySubsystems"`
		SimpleStorage                           Link  `json:"SimpleStorage"`
		StorageServices                         Links `json:"StorageServices"`
		UnassignedVolumes                       Links `json:"UnassignedVolumes"`
	}
	var tmp struct {
		temp
		Actions                      sActions
		Links                        sLinks
		Assembly                     Link  `json:"Assembly"`
		Certificates                 Link  `json:"Certificates"`
		Ports                        Link  `json:"Ports"`
		SupportedControllerProtocols Links `json:"SupportedControllerProtocols"`
		SupportedDeviceProtocols     Links `json:"SupportedDeviceProtocols"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = StorageStorageController(tmp.temp)

	// Extract the links to other entities for later
	s.freezePersonalityTarget = tmp.Actions.FreezePersonality.Target
	s.getPersonalityNonceTarget = tmp.Actions.GetPersonalityNonce.Target
	s.importForeignDrivesTarget = tmp.Actions.ImportForeignDrives.Target
	s.rekeyExternalKeyTarget = tmp.Actions.RekeyExternalKey.Target
	s.resetToDefaultsTarget = tmp.Actions.ResetToDefaults.Target
	s.revertPersonalitiesToDefaultsTarget = tmp.Actions.RevertPersonalitiesToDefaults.Target
	s.setControllerPasswordTarget = tmp.Actions.SetControllerPassword.Target
	s.setEncryptionKeyTarget = tmp.Actions.SetEncryptionKey.Target
	s.setPersonalityKeyTarget = tmp.Actions.SetPersonalityKey.Target
	s.unfreezePersonalityTarget = tmp.Actions.UnfreezePersonality.Target
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

	// This is a read/write object, so we need to save the raw object data for later
	s.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *StorageStorageController) Update() error {
	readWriteFields := []string{
		"AssetTag",
	}

	return s.UpdateFromRawData(s, s.RawData, readWriteFields)
}

// GetStorageStorageController will get a StorageStorageController instance from the service.
func GetStorageStorageController(c Client, uri string) (*StorageStorageController, error) {
	return GetObject[StorageStorageController](c, uri)
}

// ListReferencedStorageStorageControllers gets the collection of StorageStorageController from
// a provided reference.
func ListReferencedStorageStorageControllers(c Client, link string) ([]*StorageStorageController, error) {
	return GetCollectionObjects[StorageStorageController](c, link)
}

// This action shall freeze a personality for the NVMe subsystem as defined by
// the 'Configurable Device Personality' feature in the NVMe Base
// Specification.
// personality - This parameter shall contain the personality to freeze as
// defined by the 'Personality Identifier List' figure in the NVMe Base
// Specification. The value '255' shall indicate all personalities.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageStorageController) FreezePersonality(personality uint) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Personality"] = personality
	resp, taskInfo, err := PostWithTask(s.client,
		s.freezePersonalityTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall obtain a nonce required to unfreeze personalities with a
// programmed key as defined by the 'CDP Random Nonce Data' clause in the NVMe
// Base Specification.
func (s *StorageStorageController) GetPersonalityNonce() (*GetPersonalityNonceResponse, error) {
	payload := make(map[string]any)

	resp, err := s.PostWithResponse(s.getPersonalityNonceTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result GetPersonalityNonceResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// This action shall import encrypted foreign drives to a host-based storage
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
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageStorageController) ImportForeignDrives(controllerPassword string, driveEncryptionKey string, driveEncryptionKeyIdentifier string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ControllerPassword"] = controllerPassword
	payload["DriveEncryptionKey"] = driveEncryptionKey
	payload["DriveEncryptionKeyIdentifier"] = driveEncryptionKeyIdentifier
	resp, taskInfo, err := PostWithTask(s.client,
		s.importForeignDrivesTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall cause the controllers of the storage subsystem to request
// new encryption keys managed by an external key service.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageStorageController) RekeyExternalKey() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(s.client,
		s.rekeyExternalKeyTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset the storage device. This action can impact other
// resources.
// resetType - This parameter shall contain the type of reset to defaults.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageStorageController) ResetToDefaults(resetType ResetToDefaultsType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(s.client,
		s.resetToDefaultsTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall revert all personalities to their manufacturing default
// state for the NVMe subsystem as defined by the 'Manufacturing Default
// Personality' clause of the NVMe Base Specification.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageStorageController) RevertPersonalitiesToDefaults() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(s.client,
		s.revertPersonalitiesToDefaultsTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall set the controller boot password for a host-based storage
// controller.
// currentPassword - This parameter shall contain the current controller
// password. Services may reject the action request if this parameter is not
// provided or the value supplied does not match the current password.
// newPassword - This parameter shall contain the new password to set for the
// controller.
// securityKey - This parameter shall contain the controller security key.
// Services may reject the action request if this parameter is not provided or
// the value provided does not match the security key for the controller.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageStorageController) SetControllerPassword(currentPassword string, newPassword string, securityKey string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["CurrentPassword"] = currentPassword
	payload["NewPassword"] = newPassword
	payload["SecurityKey"] = securityKey
	resp, taskInfo, err := PostWithTask(s.client,
		s.setControllerPasswordTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall set the local encryption key for the storage subsystem.
// currentEncryptionKey - This parameter shall contain the current local
// encryption key on the storage subsystem. Services may reject the action
// request if this parameter is not provided or the value supplied does not
// match the current encryption key.
// encryptionKey - This parameter shall contain the local encryption key to set
// on the storage subsystem.
// encryptionKeyIdentifier - This property shall contain the local encryption
// key identifier used by the storage subsystem.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageStorageController) SetEncryptionKey(currentEncryptionKey string, encryptionKey string, encryptionKeyIdentifier string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["CurrentEncryptionKey"] = currentEncryptionKey
	payload["EncryptionKey"] = encryptionKey
	payload["EncryptionKeyIdentifier"] = encryptionKeyIdentifier
	resp, taskInfo, err := PostWithTask(s.client,
		s.setEncryptionKeyTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall set the personality key for the NVMe subsystem as defined
// by the 'Programmable Key Authentication Considerations' clause in the NVMe
// Base Specification.
// algorithm - This parameter shall contain the type of authentication
// algorithm for the key as defined by the 'CDP Authentication Algorithm' field
// in the NVMe Base Specification.
// key - This parameter shall contain the personality key to set on the NVMe
// subsystem as a hex-encoded string.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageStorageController) SetPersonalityKey(algorithm NMVePersonalityKeyAlgorithm, key string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Algorithm"] = algorithm
	payload["Key"] = key
	resp, taskInfo, err := PostWithTask(s.client,
		s.setPersonalityKeyTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// StorageStorageControllerUnfreezePersonalityParameters holds the parameters for the UnfreezePersonality action.
type StorageStorageControllerUnfreezePersonalityParameters struct {
	// AuthenticationMethod shall contain the authentication method for the
	// unfreeze operation as defined by the 'Authenticated Unfreeze Support' field
	// in the NVMe Base Specification.
	AuthenticationMethod NMVeUnfreezeAuthMode `json:"AuthenticationMethod,omitempty"`
	// Key shall contain the authentication key for the unfreeze operation as a
	// hex-encoded string. If 'AuthenticationMethod' contains 'ProgrammedKey', this
	// parameter contains a key that is derived from the original key provided in
	// the 'SetPersonalityKey' action. If 'AuthenticationMethod' contains
	// 'PhysicalId', this parameter contains the physical secure ID (PSID) of the
	// drive converted to a hex-encoded string.
	Key string `json:"Key,omitempty"`
	// Nonce shall contain the personality nonce provided by the
	// 'GetPersonalityNonce' action. This parameter shall be required if
	// 'AuthenticationMethod' contains 'ProgrammedKey'.
	Nonce string `json:"Nonce,omitempty"`
	// Personality shall contain the personality to freeze as defined by the
	// 'Personality Identifier List' figure in the NVMe Base Specification. The
	// value '255' shall indicate all personalities.
	Personality uint `json:"Personality,omitempty"`
}

// This action shall unfreeze a personality for the NVMe subsystem as defined
// by the 'Authenticated Unfreeze Operation' clause in the NVMe Base
// Specification.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *StorageStorageController) UnfreezePersonality(params *StorageStorageControllerUnfreezePersonalityParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(s.client,
		s.unfreezePersonalityTarget, params, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// BlockSecurityIDUnsupportedDrives gets the BlockSecurityIDUnsupportedDrives linked resources.
func (s *StorageStorageController) BlockSecurityIDUnsupportedDrives() ([]*Drive, error) {
	return GetObjects[Drive](s.client, s.blockSecurityIDUnsupportedDrives)
}

// BlockSecurityIDUpdateUnsuccessfulDrives gets the BlockSecurityIDUpdateUnsuccessfulDrives linked resources.
func (s *StorageStorageController) BlockSecurityIDUpdateUnsuccessfulDrives() ([]*Drive, error) {
	return GetObjects[Drive](s.client, s.blockSecurityIDUpdateUnsuccessfulDrives)
}

// Enclosures gets the Enclosures linked resources.
func (s *StorageStorageController) Enclosures() ([]*Chassis, error) {
	return GetObjects[Chassis](s.client, s.enclosures)
}

// HostingStorageSystems gets the HostingStorageSystems linked resources.
func (s *StorageStorageController) HostingStorageSystems() ([]*ComputerSystem, error) {
	return GetObjects[ComputerSystem](s.client, s.hostingStorageSystems)
}

// NVMeoFDiscoverySubsystems gets the NVMeoFDiscoverySubsystems linked resources.
func (s *StorageStorageController) NVMeoFDiscoverySubsystems() ([]*Storage, error) {
	return GetObjects[Storage](s.client, s.nVMeoFDiscoverySubsystems)
}

// SimpleStorage gets the SimpleStorage linked resource.
func (s *StorageStorageController) SimpleStorage() (*SimpleStorage, error) {
	if s.simpleStorage == "" {
		return nil, nil
	}
	return GetObject[SimpleStorage](s.client, s.simpleStorage)
}

// StorageServices gets the StorageServices linked resources.
func (s *StorageStorageController) StorageServices() ([]*StorageService, error) {
	return GetObjects[StorageService](s.client, s.storageServices)
}

// UnassignedVolumes gets the UnassignedVolumes linked resources.
func (s *StorageStorageController) UnassignedVolumes() ([]*Volume, error) {
	return GetObjects[Volume](s.client, s.unassignedVolumes)
}

// Assembly gets the Assembly linked resource.
func (s *StorageStorageController) Assembly() (*Assembly, error) {
	if s.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](s.client, s.assembly)
}

// Certificates gets the Certificates collection.
func (s *StorageStorageController) Certificates() ([]*Certificate, error) {
	if s.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](s.client, s.certificates)
}

// Ports gets the Ports collection.
func (s *StorageStorageController) Ports() ([]*Port, error) {
	if s.ports == "" {
		return nil, nil
	}
	return GetCollectionObjects[Port](s.client, s.ports)
}
