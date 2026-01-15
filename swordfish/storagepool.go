//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.6 - #StoragePool.v1_9_2.StoragePool

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

type NVMePoolType string

const (
	// EnduranceGroupNVMePoolType shall be used to specify a pool of type
	// EnduranceGroup, used by NVMe devices.
	EnduranceGroupNVMePoolType NVMePoolType = "EnduranceGroup"
	// NVMSetNVMePoolType shall be used to specify a pool of type NVMSet, used by
	// NVMe devices.
	NVMSetNVMePoolType NVMePoolType = "NVMSet"
)

type PoolType string

const (
	// BlockPoolType shall be used to specify a pool of type block. This is used
	// when the pool serves block storage.
	BlockPoolType PoolType = "Block"
	// FilePoolType shall be used to specify a pool of type file. This setting is
	// used when the pool serves file storage.
	FilePoolType PoolType = "File"
	// ObjectPoolType shall be used to specify a pool of type object.
	ObjectPoolType PoolType = "Object"
	// PoolPoolType shall be used to specify a pool of type pool. This setting is
	// used to indicate a 'pool of pools' hierarchy.
	PoolPoolType PoolType = "Pool"
)

// StoragePool is a container of data storage capable of providing capacity
// conforming to one of its supported classes of service. The storage pool does
// not support IO to its data storage.
type StoragePool struct {
	common.Entity
	// AllocatedPools shall contain a reference to the collection of storage pools
	// allocated from this storage pool.
	allocatedPools string
	// AllocatedVolumes shall contain a reference to the collection of volumes
	// allocated from this storage pool.
	allocatedVolumes string
	// BlockSizeBytes Maximum size in bytes of the blocks which form this Volume.
	// If the block size is variable, then the maximum block size in bytes should
	// be specified. If the block size is unknown or if a block concept is not
	// valid (for example, with Memory), enter a 1.
	//
	// Deprecated
	// This property has been Deprecated in favor of
	// StoragePool.v1_1_1.StoragePool.MaxBlockSizeBytes
	BlockSizeBytes *uint `json:",omitempty"`
	// Capacity shall provide an information about the actual utilization of the
	// capacity within this storage pool.
	Capacity Capacity
	// CapacitySources shall provide capacity allocation data from a named source
	// resource.
	CapacitySources []CapacitySource
	// CapacitySources@odata.count
	CapacitySourcesCount int `json:"CapacitySources@odata.count"`
	// ClassesOfService shall contain references to all classes of service
	// supported by this storage pool. Capacity allocated from this storage pool
	// shall conform to one of the referenced classes of service.
	classesOfService string
	// Compressed shall contain a boolean indicator if the StoragePool is currently
	// utilizing compression or not.
	//
	// Version added: v1.3.0
	//
	// Deprecated: v1.6.0
	// This property has been deprecated in favor of the IsCompressed and
	// DefaultCompressionBehavior properties.
	Compressed bool
	// CompressionEnabled shall indicate whether or not compression is enabled on
	// the storage pool.
	//
	// Version added: v1.6.0
	CompressionEnabled bool
	// Deduplicated shall contain a boolean indicator if the StoragePool is
	// currently utilizing deduplication or not.
	//
	// Version added: v1.3.0
	//
	// Deprecated: v1.6.0
	// This property has been deprecated in favor of the IsDeduplicated and
	// DefaultDedupeBehavior properties.
	Deduplicated bool
	// DeduplicationEnabled shall indicate whether or not deduplication is enabled
	// on the storage pool.
	//
	// Version added: v1.6.0
	DeduplicationEnabled bool
	// DefaultClassOfService shall reference the default class of service for
	// entities allocated from this storage pool. If the ClassesOfService
	// collection is not empty, then the value of this property shall be one of its
	// entries. If not present, the default class of service of the containing
	// StorageService entity shall be used.
	//
	// Version added: v1.2.0
	defaultClassOfService string
	// DefaultCompressionBehavior shall indicate the default dedupe behavior
	// applied to the child resource (E.g., volume or storage pool) created out of
	// the storage pool if the 'Compressed' property is not set on the create
	// request.
	//
	// Version added: v1.6.0
	DefaultCompressionBehavior bool
	// DefaultDeduplicationBehavior shall indicate the default deduplication
	// behavior applied to the child resource (E.g., volume or storage pool)
	// created out of the storage pool if the 'Deduplicated' property is not set on
	// the create request.
	//
	// Version added: v1.6.0
	DefaultDeduplicationBehavior bool
	// DefaultEncryptionBehavior shall indicate the default dedupe behavior applied
	// to the child resource (E.g., volume or storage pool) created out of the
	// storage pool if the 'Encrypted' property is not set on the create request.
	//
	// Version added: v1.6.0
	DefaultEncryptionBehavior bool
	// Encrypted shall contain a boolean indicator if the StoragePool is currently
	// utilizing encryption or not.
	//
	// Version added: v1.3.0
	//
	// Deprecated: v1.6.0
	// This property has been deprecated in favor of the IsEncrypted and
	// DefaultEncryptionBehavior properties.
	Encrypted bool
	// EncryptionEnabled shall indicate whether or not encryption is enabled on the
	// storage pool.
	//
	// Version added: v1.6.0
	EncryptionEnabled bool
	// IOStatistics shall represent IO statistics for this StoragePool.
	//
	// Version added: v1.2.0
	//
	// Deprecated: v1.9.0
	// This property has been deprecated in favor of the IOStatistics property in
	// Metrics.
	IOStatistics IOStatistics
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// LowSpaceWarningThresholdPercents shall be triggered: Across all
	// CapacitySources entries, percent = (SUM(AllocatedBytes) -
	// SUM(ConsumedBytes))/SUM(AllocatedBytes).
	LowSpaceWarningThresholdPercents []*int
	// MaxBlockSizeBytes shall be NULL.
	//
	// Version added: v1.1.1
	MaxBlockSizeBytes *uint `json:",omitempty"`
	// Metrics shall contain a link to a resource of type StoragePoolMetrics that
	// specifies the metrics for this storage pool. IO metrics are reported in the
	// IOStatistics property.
	//
	// Version added: v1.9.0
	metrics string
	// NVMeEnduranceGroupProperties shall contain properties to use when
	// StoragePool is used to describe an NVMe Endurance Group.
	//
	// Version added: v1.4.0
	NVMeEnduranceGroupProperties NVMeEnduranceGroupProperties
	// NVMeProperties shall indicate the type of storage pool.
	//
	// Version added: v1.6.0
	NVMeProperties NVMeProperties
	// NVMeSetProperties shall contain properties to use when StoragePool is used
	// to describe an NVMe Set.
	//
	// Version added: v1.4.0
	NVMeSetProperties NVMeSetProperties
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PoolType shall indicate the type of storage pool.
	//
	// Version added: v1.6.0
	//
	// Deprecated: v1.7.0
	// This property has been deprecated in favor of the SupportedPoolTypes
	// property.
	PoolType []PoolType
	// RecoverableCapacitySourceCount The value is the number of available capacity
	// source resources currently available in the event that an equivalent
	// capacity source resource fails.
	//
	// Version added: v1.2.0
	RecoverableCapacitySourceCount *int `json:",omitempty"`
	// RemainingCapacityPercent shall return {[(SUM(AllocatedBytes) -
	// SUM(ConsumedBytes)]/SUM(AllocatedBytes)}*100 represented as an integer
	// value.
	//
	// Version added: v1.1.0
	RemainingCapacityPercent *int `json:",omitempty"`
	// ReplicationEnabled shall indicate whether or not replication is enabled on
	// the storage pool. If enabled for pool, replication can still be disabled on
	// individual resources (e.g., volumes) within the pool.
	//
	// Version added: v1.8.0
	ReplicationEnabled bool
	// Status shall contain the status of the StoragePool.
	Status common.Status
	// SupportedPoolTypes shall contain all the PoolType values supported by the
	// storage pool.
	//
	// Version added: v1.7.0
	SupportedPoolTypes []PoolType
	// SupportedProvisioningPolicies shall specify all supported storage allocation
	// policies for the Storage Pool.
	//
	// Version added: v1.3.0
	SupportedProvisioningPolicies []ProvisioningPolicy
	// SupportedRAIDTypes shall contain all the RAIDType values supported by the
	// storage pool.
	//
	// Version added: v1.3.0
	SupportedRAIDTypes []RAIDType
	// addDrivesTarget is the URL to send AddDrives requests.
	addDrivesTarget string
	// removeDrivesTarget is the URL to send RemoveDrives requests.
	removeDrivesTarget string
	// setCompressionStateTarget is the URL to send SetCompressionState requests.
	setCompressionStateTarget string
	// setDeduplicationStateTarget is the URL to send SetDeduplicationState requests.
	setDeduplicationStateTarget string
	// setEncryptionStateTarget is the URL to send SetEncryptionState requests.
	setEncryptionStateTarget string
	// dedicatedSpareDrives are the URIs for DedicatedSpareDrives.
	dedicatedSpareDrives []string
	// owningStorageResource is the URI for OwningStorageResource.
	owningStorageResource string
	// spareResourceSets are the URIs for SpareResourceSets.
	spareResourceSets []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a StoragePool object from the raw JSON.
func (s *StoragePool) UnmarshalJSON(b []byte) error {
	type temp StoragePool
	type sActions struct {
		AddDrives             common.ActionTarget `json:"#StoragePool.AddDrives"`
		RemoveDrives          common.ActionTarget `json:"#StoragePool.RemoveDrives"`
		SetCompressionState   common.ActionTarget `json:"#StoragePool.SetCompressionState"`
		SetDeduplicationState common.ActionTarget `json:"#StoragePool.SetDeduplicationState"`
		SetEncryptionState    common.ActionTarget `json:"#StoragePool.SetEncryptionState"`
	}
	type sLinks struct {
		DedicatedSpareDrives  common.Links `json:"DedicatedSpareDrives"`
		DefaultClassOfService common.Link  `json:"DefaultClassOfService"`
		OwningStorageResource common.Link  `json:"OwningStorageResource"`
		SpareResourceSets     common.Links `json:"SpareResourceSets"`
	}
	var tmp struct {
		temp
		Actions          sActions
		Links            sLinks
		AllocatedPools   common.Link `json:"allocatedPools"`
		AllocatedVolumes common.Link `json:"allocatedVolumes"`
		ClassesOfService common.Link `json:"classesOfService"`
		Metrics          common.Link `json:"metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = StoragePool(tmp.temp)

	// Extract the links to other entities for later
	s.addDrivesTarget = tmp.Actions.AddDrives.Target
	s.removeDrivesTarget = tmp.Actions.RemoveDrives.Target
	s.setCompressionStateTarget = tmp.Actions.SetCompressionState.Target
	s.setDeduplicationStateTarget = tmp.Actions.SetDeduplicationState.Target
	s.setEncryptionStateTarget = tmp.Actions.SetEncryptionState.Target
	s.dedicatedSpareDrives = tmp.Links.DedicatedSpareDrives.ToStrings()
	s.defaultClassOfService = tmp.Links.DefaultClassOfService.String()
	s.owningStorageResource = tmp.Links.OwningStorageResource.String()
	s.spareResourceSets = tmp.Links.SpareResourceSets.ToStrings()
	s.allocatedPools = tmp.AllocatedPools.String()
	s.allocatedVolumes = tmp.AllocatedVolumes.String()
	s.classesOfService = tmp.ClassesOfService.String()
	s.metrics = tmp.Metrics.String()

	// This is a read/write object, so we need to save the raw object data for later
	s.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *StoragePool) Update() error {
	readWriteFields := []string{
		"Capacity",
		"CapacitySources",
		"CapacitySources@odata.count",
		"ClassesOfService",
		"Compressed",
		"CompressionEnabled",
		"Deduplicated",
		"DeduplicationEnabled",
		"DefaultClassOfService",
		"DefaultCompressionBehavior",
		"DefaultDeduplicationBehavior",
		"DefaultEncryptionBehavior",
		"Encrypted",
		"EncryptionEnabled",
		"IOStatistics",
		"Identifier",
		"LowSpaceWarningThresholdPercents",
		"NVMeEnduranceGroupProperties",
		"NVMeProperties",
		"NVMeSetProperties",
		"RecoverableCapacitySourceCount",
		"ReplicationEnabled",
		"Status",
		"SupportedProvisioningPolicies",
	}

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
}

// GetStoragePool will get a StoragePool instance from the service.
func GetStoragePool(c common.Client, uri string) (*StoragePool, error) {
	return common.GetObject[StoragePool](c, uri)
}

// ListReferencedStoragePools gets the collection of StoragePool from
// a provided reference.
func ListReferencedStoragePools(c common.Client, link string) ([]*StoragePool, error) {
	return common.GetCollectionObjects[StoragePool](c, link)
}

// AddDrives shall be used to add a drive, or set of drives, to an underlying
// capacity source for the storage pool.
// capacitySource - This parameter shall contain the target capacity source for
// the drive(s). This property does not need to be specified if the storage
// pool only contains one capacity source, or if the implementation is capable
// of automatically selecting the appropriate capacity source.
// drives - This parameter shall contain the Uri to the existing drive or
// drives to be added to a capacity source of the storage pool. The
// implementation may impose restrictions on the number of drives added
// simultaneously.
func (s *StoragePool) AddDrives(capacitySource string, drives []string) error {
	payload := make(map[string]any)
	payload["CapacitySource"] = capacitySource
	payload["Drives"] = drives
	return s.Post(s.addDrivesTarget, payload)
}

// RemoveDrives shall be used to remove a drive from the StoragePool. This
// action is targeted at a graceful drive removal process, such as initiating a
// drive cleanup and data reallocation before drive removal from the pool. The
// implementation may impose restrictions on the number of drives removed
// simultaneously.
// drives - This parameter shall contain the Uri to the drive or drives to be
// removed from the underlying capacity source.
func (s *StoragePool) RemoveDrives(drives []string) error {
	payload := make(map[string]any)
	payload["Drives"] = drives
	return s.Post(s.removeDrivesTarget, payload)
}

// SetCompressionState shall be used to set the compression state of the storage pool.
// This may be both a highly impactful, as well as a long running operation.
// enable - This property shall indicate the desired compression state of the
// storage pool.
func (s *StoragePool) SetCompressionState(enable bool) error {
	payload := make(map[string]any)
	payload["Enable"] = enable
	return s.Post(s.setCompressionStateTarget, payload)
}

// SetDeduplicationState shall be used to set the dedupe state of the storage pool. This
// may be both a highly impactful, as well as a long running operation.
// enable - This property shall indicate the desired deduplication state of the
// storage pool.
func (s *StoragePool) SetDeduplicationState(enable bool) error {
	payload := make(map[string]any)
	payload["Enable"] = enable
	return s.Post(s.setDeduplicationStateTarget, payload)
}

// SetEncryptionState shall be used to set the encryption state of the storage pool.
// This may be both a highly impactful, as well as a long running operation.
// enable - This property shall indicate the desired encryption state of the
// storage pool.
func (s *StoragePool) SetEncryptionState(enable bool) error {
	payload := make(map[string]any)
	payload["Enable"] = enable
	return s.Post(s.setEncryptionStateTarget, payload)
}

// DedicatedSpareDrives gets the DedicatedSpareDrives linked resources.
func (s *StoragePool) DedicatedSpareDrives(client common.Client) ([]*redfish.Drive, error) {
	return common.GetObjects[redfish.Drive](client, s.dedicatedSpareDrives)
}

// DefaultClassOfService gets the DefaultClassOfService linked resource.
func (s *StoragePool) DefaultClassOfService(client common.Client) (*ClassOfService, error) {
	if s.defaultClassOfService == "" {
		return nil, nil
	}
	return common.GetObject[ClassOfService](client, s.defaultClassOfService)
}

// OwningStorageResource gets the OwningStorageResource linked resource.
func (s *StoragePool) OwningStorageResource(client common.Client) (*redfish.Storage, error) {
	if s.owningStorageResource == "" {
		return nil, nil
	}
	return common.GetObject[redfish.Storage](client, s.owningStorageResource)
}

// SpareResourceSets gets the SpareResourceSets linked resources.
func (s *StoragePool) SpareResourceSets(client common.Client) ([]*SpareResourceSet, error) {
	return common.GetObjects[SpareResourceSet](client, s.spareResourceSets)
}

// AllocatedPools gets the AllocatedPools collection.
func (s *StoragePool) AllocatedPools(client common.Client) ([]*StoragePool, error) {
	if s.allocatedPools == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[StoragePool](client, s.allocatedPools)
}

// AllocatedVolumes gets the AllocatedVolumes collection.
func (s *StoragePool) AllocatedVolumes(client common.Client) ([]*common.Volume, error) {
	if s.allocatedVolumes == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[common.Volume](client, s.allocatedVolumes)
}

// ClassesOfService gets the ClassesOfService collection.
func (s *StoragePool) ClassesOfService(client common.Client) ([]*ClassOfService, error) {
	if s.classesOfService == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[ClassOfService](client, s.classesOfService)
}

// Metrics gets the Metrics linked resource.
func (s *StoragePool) Metrics(client common.Client) (*StoragePoolMetrics, error) {
	if s.metrics == "" {
		return nil, nil
	}
	return common.GetObject[StoragePoolMetrics](client, s.metrics)
}

// EndGrpLifetime This contains properties for the Endurance Group Lifetime
// attributes.
type EndGrpLifetime struct {
	// DataUnitsRead shall contain the total number of data units read from this
	// endurance group. This value does not include controller reads due to
	// internal operations such as garbage collection. The value is reported in
	// billions, where a value of 1 corresponds to 1 billion bytes written, and is
	// rounded up. A value of zero indicates the property is unsupported.
	//
	// Version added: v1.4.0
	DataUnitsRead *uint `json:",omitempty"`
	// DataUnitsWritten shall contain the total number of data units written from
	// this endurance group. This value does not include controller writes due to
	// internal operations such as garbage collection. The value is reported in
	// billions, where a value of 1 corresponds to 1 billion bytes written, and is
	// rounded up. A value of zero indicates the property is unsupported.
	//
	// Version added: v1.4.0
	DataUnitsWritten *uint `json:",omitempty"`
	// EnduranceEstimate shall contain an estimate of the total number of data
	// bytes that may be written to the Endurance Group over the lifetime of the
	// Endurance Group assuming a write amplication of 1. The value is reported in
	// billions, where a value of 1 corresponds to 1 billion bytes written, and is
	// rounded up. A value of zero indicates endurance estimates are unsupported.
	//
	// Version added: v1.4.0
	EnduranceEstimate *uint `json:",omitempty"`
	// ErrorInformationLogEntryCount shall contain the number of error information
	// log entries over the life of the controller for the endurance group.
	//
	// Version added: v1.4.0
	ErrorInformationLogEntryCount *uint `json:",omitempty"`
	// HostReadCommandCount shall contain the number of read commands completed by
	// all controllers in the NVM subsystem for the Endurance Group. For the NVM
	// command set, the is the number of compare commands and read commands.
	//
	// Version added: v1.4.0
	HostReadCommandCount *uint `json:",omitempty"`
	// HostWriteCommandCount shall contain the number of write commands completed
	// by all controllers in the NVM subsystem for the Endurance Group. For the NVM
	// command set, the is the number of compare commands and write commands.
	//
	// Version added: v1.4.0
	HostWriteCommandCount *uint `json:",omitempty"`
	// MediaAndDataIntegrityErrorCount shall contain the number of occurrences where
	// the controller detected an unrecovered data integrity error for the
	// Endurance Group. Errors such as uncorrectable ECC, CRC checksum failure, or
	// LBA tag mismatch are included in this field.
	//
	// Version added: v1.4.0
	MediaAndDataIntegrityErrorCount *uint `json:",omitempty"`
	// MediaUnitsWritten shall contain the total number of data units written from
	// this endurance group. This value includes host and controller writes due to
	// internal operations such as garbage collection. The value is reported in
	// billions, where a value of 1 corresponds to 1 billion bytes written, and is
	// rounded up. A value of zero indicates the property is unsupported.
	//
	// Version added: v1.4.0
	MediaUnitsWritten *uint `json:",omitempty"`
	// PercentUsed shall contain a vendor-specific estimate of the percent life
	// used for the endurance group based on the actual usage and the manufacturer
	// prediction of NVM life. A value of 100 indicates that the estimated
	// endurance of the NVM in the Endurance Group has been consumed, but may not
	// indicate an NVM failure. According to the NVMe and JEDEC specs, the value is
	// allowed to exceed 100. Percentages greater than 254 shall be represented as
	// 255.
	//
	// Version added: v1.4.0
	PercentUsed *uint `json:",omitempty"`
}

// NVMeEnduranceGroupProperties This contains properties to use when StoragePool
// is used to describe an NVMe Endurance Group.
type NVMeEnduranceGroupProperties struct {
	// EndGrpLifetime shall contain any Endurance Group Lifetime properties.
	//
	// Version added: v1.4.0
	EndGrpLifetime EndGrpLifetime
	// PredictedMediaLifeLeftPercent shall contain an indicator of the percentage
	// of life remaining in the drive's media.
	//
	// Version added: v1.4.0
	PredictedMediaLifeLeftPercent *float64 `json:",omitempty"`
}

// NVMeProperties This contains properties to use when StoragePool is used to
// describe an NVMe construct.
type NVMeProperties struct {
	// NVMePoolType shall indicate whether the StoragePool is used as an
	// EnduranceGroup or an NVMSet.
	//
	// Version added: v1.6.0
	NVMePoolType NVMePoolType
}

// NVMeSetProperties This contains properties to use when StoragePool is used to
// describe an NVMe Set.
type NVMeSetProperties struct {
	// EnduranceGroupIdentifier shall contain a 16-bit hex value that contains the
	// endurance group identifier. The endurance group identifier is unique within
	// a subsystem. Reserved values include 0.
	//
	// Version added: v1.4.0
	EnduranceGroupIdentifier string
	// OptimalWriteSizeBytes shall contain the Optimal Write Size in Bytes for this
	// NVMe Set.
	//
	// Version added: v1.4.0
	OptimalWriteSizeBytes *uint `json:",omitempty"`
	// Random4kReadTypicalNanoSeconds shall contain the typical time to complete a
	// 4k read in 100 nano-second units when the NVM Set is in a Predictable
	// Latency Mode Deterministic Window and there is 1 outstanding command per NVM
	// Set.
	//
	// Version added: v1.4.0
	Random4kReadTypicalNanoSeconds *uint `json:",omitempty"`
	// SetIdentifier shall contain a 16-bit hex value that contains the NVMe Set
	// group identifier. The NVM Set identifier is unique within a subsystem.
	// Reserved values include 0.
	//
	// Version added: v1.4.0
	SetIdentifier string
	// UnallocatedNVMNamespaceCapacityBytes shall contain the unallocated capacity
	// of the NVMe Set in bytes.
	//
	// Version added: v1.4.0
	UnallocatedNVMNamespaceCapacityBytes *uint `json:",omitempty"`
}
