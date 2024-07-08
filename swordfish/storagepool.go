//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// EndGrpLifetime This contains properties for the Endurance Group Lifetime attributes.
type EndGrpLifetime struct {
	// DataUnitsRead shall contain the total number of data units read from this endurance group. This value does not
	// include controller reads due to internal operations such as garbage collection. The value is reported in
	// billions, where a value of 1 corresponds to 1 billion bytes written, and is rounded up. A value of zero
	// indicates the property is unsupported.
	DataUnitsRead int
	// DataUnitsWritten shall contain the total number of data units written from this endurance group. This value does
	// not include controller writes due to internal operations such as garbage collection. The value is reported in
	// billions, where a value of 1 corresponds to 1 billion bytes written, and is rounded up. A value of zero
	// indicates the property is unsupported.
	DataUnitsWritten int
	// EnduranceEstimate shall contain an estimate of the total number of data bytes that may be written to the
	// Endurance Group over the lifetime of the Endurance Group assuming a write amplication of 1. The value is
	// reported in billions, where a value of 1 corresponds to 1 billion bytes written, and is rounded up. A value of
	// zero indicates endurance estimates are unsupported.
	EnduranceEstimate int
	// ErrorInformationLogEntryCount shall contain the number of error information log entries over the life of the
	// controller for the endurance group.
	ErrorInformationLogEntryCount int
	// HostReadCommandCount shall contain the number of read commands completed by all controllers in the NVM subsystem
	// for the Endurance Group. For the NVM command set, the is the number of compare commands and read commands.
	HostReadCommandCount int
	// HostWriteCommandCount shall contain the number of write commands completed by all controllers in the NVM
	// subsystem for the Endurance Group. For the NVM command set, the is the number of compare commands and write
	// commands.
	HostWriteCommandCount int
	// MediaAndDataIntegrityErrorCount shall contain the number of occurrences where the controller detected an
	// unrecovered data integrity error for the Endurance Group. Errors such as uncorrectable ECC, CRC checksum
	// failure, or LBA tag mismatch are included in this field.
	MediaAndDataIntegrityErrorCount int
	// MediaUnitsWritten shall contain the total number of data units written from this endurance group. This value
	// includes host and controller writes due to internal operations such as garbage collection. The value is reported
	// in billions, where a value of 1 corresponds to 1 billion bytes written, and is rounded up. A value of zero
	// indicates the property is unsupported.
	MediaUnitsWritten int
	// PercentUsed shall contain a vendor-specific estimate of the percent life used for the endurance group based on
	// the actual usage and the manufacturer prediction of NVM life. A value of 100 indicates that the estimated
	// endurance of the NVM in the Endurance Group has been consumed, but may not indicate an NVM failure. According to
	// the NVMe and JEDEC specs, the value is allowed to exceed 100. Percentages greater than 254 shall be represented
	// as 255.
	PercentUsed int
}

// NVMeEnduranceGroupProperties contains properties to use when StoragePool is used to describe an NVMe
// Endurance Group.
type NVMeEnduranceGroupProperties struct {
	// EndGrpLifetime shall contain any Endurance Group Lifetime properties.
	EndGrpLifetime EndGrpLifetime
	// PredictedMediaLifeLeftPercent shall contain an indicator of the percentage of life remaining in the drive's
	// media.
	PredictedMediaLifeLeftPercent float64
}

type NVMePoolType string

const (
	// EnduranceGroupNVMePoolType is of type EnduranceGroup, used by NVMe devices.
	EnduranceGroupNVMePoolType NVMePoolType = "EnduranceGroup"
	// NVMSetNVMePoolType is of type NVMSet, used by NVMe devices.
	NVMSetNVMePoolType NVMePoolType = "NVMSet"
)

// NVMeProperties contains properties to use when StoragePool is used to describe an NVMe construct.
type NVMeProperties struct {
	// NVMePoolType shall indicate whether the StoragePool is used as an EnduranceGroup or an NVMSet.
	NVMePoolType NVMePoolType
}

// NVMeSetProperties contains properties to use when StoragePool is used to describe an NVMe Set.
type NVMeSetProperties struct {
	// EnduranceGroupIdentifier shall contain a 16-bit hex value that contains the endurance group identifier. The
	// endurance group identifier is unique within a subsystem. Reserved values include 0.
	EnduranceGroupIdentifier string
	// OptimalWriteSizeBytes shall contain the Optimal Write Size in Bytes for this NVMe Set.
	OptimalWriteSizeBytes int
	// Random4kReadTypicalNanoSeconds shall contain the typical time to complete a 4k read in 100 nano-second units
	// when the NVM Set is in a Predictable Latency Mode Deterministic Window and there is 1 outstanding command per
	// NVM Set.
	Random4kReadTypicalNanoSeconds int
	// SetIdentifier shall contain a 16-bit hex value that contains the NVMe Set group identifier. The NVM Set
	// identifier is unique within a subsystem. Reserved values include 0.
	SetIdentifier string
	// UnallocatedNVMNamespaceCapacityBytes shall contain the unallocated capacity of the NVMe Set in bytes.
	UnallocatedNVMNamespaceCapacityBytes int
}

type PoolType string

const (
	// BlockPoolType is of type block.
	BlockPoolType PoolType = "Block"
	// FilePoolType is of type file.
	FilePoolType PoolType = "File"
	// ObjectPoolType is of type object.
	ObjectPoolType PoolType = "Object"
	// PoolPoolType is of type pool, indicating a hierarchy.
	PoolPoolType PoolType = "Pool"
)

// StoragePool is a container of data storage capable of providing
// capacity conforming to one of its supported classes of service. The
// storage pool does not support IO to its data storage.
type StoragePool struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AllocatedPools shall contain a reference
	// to the collection of storage pools allocated from this storage pool.
	allocatedPools string
	// AllocatedVolumes shall contain a
	// reference to the collection of volumes allocated from this storage
	// pool.
	allocatedVolumes string
	// Capacity shall provide information about the actual utilization of the
	// capacity within this storage pool.
	Capacity Capacity
	// CapacitySources is fully or partially consumed storage from a source
	// resource. Each entry shall provide capacity allocation data from a
	// named source resource.
	capacitySources []string
	// CapacitySourcesCount is the number of capacity sources.
	CapacitySourcesCount int `json:"CapacitySources@odata.count"`
	// ClassesOfService shall contain references to all classes of service
	// supported by this storage pool. Capacity allocated from this storage pool
	// shall conform to one of the referenced classes of service.
	classesOfService string
	// Compressed shall contain a boolean indicator if the StoragePool is
	// currently utilizing compression or not.
	// This property has been deprecated in favor of the IsCompressed and
	// DefaultCompressionBehavior properties.
	Compressed bool
	// CompressionEnabled shall indicate whether or not compression is enabled on the storage pool.
	CompressionEnabled bool
	// Deduplicted shall contain a boolean indicator if the StoragePool is
	// currently utilizing deduplication or not.
	// This property has been deprecated in favor of the IsDeduplicated and
	// DefaultDedupeBehavior properties.
	Deduplicated bool
	// DeduplicationEnabled shall indicate whether or not deduplication is enabled on the storage pool.
	DeduplicationEnabled bool
	// DefaultClassOfService is used.
	defaultClassOfService string
	// DefaultCompressionBehavior shall indicate the default dedupe behavior applied to the child resource (E.g.,
	// volume or storage pool) created out of the storage pool if the 'Compressed' property is not set on the create
	// request.
	DefaultCompressionBehavior bool
	// DefaultDeduplicationBehavior shall indicate the default deduplication behavior applied to the child resource
	// (E.g., volume or storage pool) created out of the storage pool if the 'Deduplicated' property is not set on the
	// create request.
	DefaultDeduplicationBehavior bool
	// DefaultEncryptionBehavior shall indicate the default dedupe behavior applied to the child resource (E.g., volume
	// or storage pool) created out of the storage pool if the 'Encrypted' property is not set on the create request.
	DefaultEncryptionBehavior bool
	// Description provides a description of this resource.
	Description string
	// EncryptionEnabled shall indicate whether or not encryption is enabled on the storage pool.
	EncryptionEnabled bool
	// Encrypted shall contain a boolean indicator if the
	// StoragePool is currently utilizing encryption or not.
	// This property has been deprecated in favor of the IsEncrypted and DefaultEncryptionBehavior properties.
	Encrypted bool
	// IOStatistics is the value shall represent IO statistics for this
	// StoragePool.
	IOStatistics IOStatistics
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// LowSpaceWarningThresholdPercents is each time the following value is
	// less than one of the values in the array the
	// LOW_SPACE_THRESHOLD_WARNING event shall be triggered: Across all
	// CapacitySources entries, percent = (SUM(AllocatedBytes) -
	// SUM(ConsumedBytes))/SUM(AllocatedBytes).
	LowSpaceWarningThresholdPercents []int
	// MaxBlockSizeBytes if present, the value is the maximum block size
	// of an allocated resource. If the block size is unknown or if a block
	// concept is not valid (for example, with Memory), this property shall
	// be NULL.
	MaxBlockSizeBytes int64
	// Metrics shall contain a link to a resource of type StoragePoolMetrics that specifies the metrics for this
	// storage pool. IO metrics are reported in the IOStatistics property.
	metrics string
	// NVMeEnduranceGroupProperties shall contain properties to use when StoragePool is used to describe an NVMe
	// Endurance Group.
	NVMeEnduranceGroupProperties NVMeEnduranceGroupProperties
	// NVMeProperties shall indicate the type of storage pool.
	NVMeProperties NVMeProperties
	// NVMeSetProperties shall contain properties to use when StoragePool is used to describe an NVMe Set.
	NVMeSetProperties NVMeSetProperties
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RecoverableCapacitySourceCount is the value of the number of available
	// capacity source resources currently available in the event that an
	// equivalent capacity source resource fails.
	RecoverableCapacitySourceCount int
	// RemainingCapacityPercent if present, this value shall return
	// {[(SUM(AllocatedBytes) - SUM(ConsumedBytes)]/SUM(AllocatedBytes)}*100
	// represented as an integer value.
	RemainingCapacityPercent int
	// ReplicationEnabled shall indicate whether or not replication is enabled on the storage pool. If enabled for
	// pool, replication can still be disabled on individual resources (e.g., volumes) within the pool.
	ReplicationEnabled bool
	// Status is the storage pool status.
	Status common.Status
	// SupportedPoolTypes shall contain all the PoolType values supported by the storage pool.
	SupportedPoolTypes []PoolType
	// SupportedProvisioningPolicies shall specify all supported storage allocation policies for the Storage Pool.
	SupportedProvisioningPolicies []ProvisioningPolicy
	// SupportedRAIDTypes shall contain all the RAIDType values supported by the storage pool.
	SupportedRAIDTypes []RAIDType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	// DedicatedSpareDrives shall be a reference to the resources that this
	// StoragePool is associated with and shall reference resources of type
	// Drive. This property shall only contain references to Drive entities
	// which are currently assigned as a dedicated spare and are able to support
	// this StoragePool.
	dedicatedSpareDrives []string
	// DedicatedSpareDrivesCount is the number of drives.
	DedicatedSpareDrivesCount int
	// OwningStorageResource shall be a pointer to the Storage resource that owns or contains this StoragePool.
	owningStorageResource string
	// SpareResourceSets shall contain resources that may be utilized to replace
	// the capacity provided by a failed resource having a compatible type.
	spareResourceSets []string
	// SpareResourceSetsCount is the number of spare resource sets.
	SpareResourceSetsCount int

	addDrivesTarget             string
	removeDrivesTarget          string
	setCompressionStateTarget   string
	setDeduplicationStateTarget string
	setEncryptionStateTarget    string
}

// UnmarshalJSON unmarshals a StoragePool object from the raw JSON.
func (storagepool *StoragePool) UnmarshalJSON(b []byte) error {
	type temp StoragePool
	type links struct {
		DedicatedSpareDrives      common.Links
		DedicatedSpareDrivesCount int `json:"DedicatedSpareDrives@odata.count"`
		OwningStorageResource     common.Link
		SpareResourceSets         common.Links
		SpareResourceSetsCount    int `json:"SpareResourceSets@odata.count"`
	}
	var t struct {
		temp
		Links                 links
		AllocatedPools        common.Link
		AllocatedVolumes      common.Link
		CapacitySource        common.Links
		ClassesOfService      common.Link
		DefaultClassOfService common.Link
		Metrics               common.Link
		Actions               struct {
			AddDrives             common.ActionTarget `json:"#StoragePool.AddDrives"`
			RemoveDrives          common.ActionTarget `json:"#StoragePool.RemoveDrives"`
			SetCompressionState   common.ActionTarget `json:"#StoragePool.SetCompressionState"`
			SetDeduplicationState common.ActionTarget `json:"#StoragePool.SetDeduplicationState"`
			SetEncryptionState    common.ActionTarget `json:"#StoragePool.SetEncryptionState"`
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*storagepool = StoragePool(t.temp)

	// Extract the links to other entities for later
	storagepool.dedicatedSpareDrives = t.Links.DedicatedSpareDrives.ToStrings()
	storagepool.DedicatedSpareDrivesCount = t.Links.DedicatedSpareDrivesCount
	storagepool.owningStorageResource = t.Links.OwningStorageResource.String()
	storagepool.spareResourceSets = t.Links.SpareResourceSets.ToStrings()
	storagepool.SpareResourceSetsCount = t.Links.SpareResourceSetsCount

	storagepool.allocatedPools = t.AllocatedPools.String()
	storagepool.allocatedVolumes = t.AllocatedVolumes.String()
	storagepool.capacitySources = t.CapacitySource.ToStrings()
	storagepool.classesOfService = t.ClassesOfService.String()
	storagepool.defaultClassOfService = t.DefaultClassOfService.String()
	storagepool.metrics = t.Metrics.String()

	storagepool.addDrivesTarget = t.Actions.AddDrives.Target
	storagepool.removeDrivesTarget = t.Actions.RemoveDrives.Target
	storagepool.setCompressionStateTarget = t.Actions.SetCompressionState.Target
	storagepool.setDeduplicationStateTarget = t.Actions.SetDeduplicationState.Target
	storagepool.setEncryptionStateTarget = t.Actions.SetEncryptionState.Target

	// This is a read/write object, so we need to save the raw object data for later
	storagepool.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (storagepool *StoragePool) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(StoragePool)
	err := original.UnmarshalJSON(storagepool.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"CapacitySources",
		"ClassesOfService",
		"Compressed",
		"Deduplicated",
		"DefaultClassOfService",
		"Encrypted",
		"LowSpaceWarningThresholdPercents",
		"RecoverableCapacitySourceCount",
		"SupportedProvisioningPolicies",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(storagepool).Elem()

	return storagepool.Entity.Update(originalElement, currentElement, readWriteFields)
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

// DedicatedSpareDrives gets the Drive entities which are currently assigned as
// a dedicated spare and are able to support this StoragePool.
func (storagepool *StoragePool) DedicatedSpareDrives() ([]*redfish.Drive, error) {
	return common.GetObjects[redfish.Drive](storagepool.GetClient(), storagepool.dedicatedSpareDrives)
}

// SpareResourceSets gets resources that may be utilized to replace the capacity
// provided by a failed resource having a compatible type.
func (storagepool *StoragePool) SpareResourceSets() ([]*SpareResourceSet, error) {
	return common.GetObjects[SpareResourceSet](storagepool.GetClient(), storagepool.spareResourceSets)
}

// AllocatedPools gets the storage pools allocated from this storage pool.
func (storagepool *StoragePool) AllocatedPools() ([]*StoragePool, error) {
	return ListReferencedStoragePools(storagepool.GetClient(), storagepool.allocatedPools)
}

// AllocatedVolumes gets the volumes allocated from this storage pool.
func (storagepool *StoragePool) AllocatedVolumes() ([]*Volume, error) {
	return ListReferencedVolumes(storagepool.GetClient(), storagepool.allocatedVolumes)
}

// CapacitySources gets space allocations to this pool.
func (storagepool *StoragePool) CapacitySources() ([]*CapacitySource, error) {
	return common.GetObjects[CapacitySource](storagepool.GetClient(), storagepool.capacitySources)
}

// ClassesOfService gets references to all classes of service supported by this
// storage pool. Capacity allocated from this storage pool shall conform to one
// of the referenced classes of service.
func (storagepool *StoragePool) ClassesOfService() ([]*ClassOfService, error) {
	return ListReferencedClassOfServices(storagepool.GetClient(), storagepool.classesOfService)
}

// DefaultClassOfService gets the default ClassOfService for this pool.
func (storagepool *StoragePool) DefaultClassOfService() (*ClassOfService, error) {
	if storagepool.defaultClassOfService == "" {
		return nil, nil
	}
	return GetClassOfService(storagepool.GetClient(), storagepool.defaultClassOfService)
}

// OwningStorageResource gets the Storage resource that owns or contains this StoragePool.
func (storagepool *StoragePool) OwningStorageResource() (*redfish.Storage, error) {
	if storagepool.owningStorageResource == "" {
		return nil, nil
	}

	return redfish.GetStorage(storagepool.GetClient(), storagepool.owningStorageResource)
}

// Metrics gets the metrics for this storage pool.
func (storagepool *StoragePool) Metrics() (*StoragePoolMetrics, error) {
	if storagepool.metrics == "" {
		return nil, nil
	}
	return GetStoragePoolMetrics(storagepool.GetClient(), storagepool.metrics)
}

// AddDrives will add an additional drive, or set of drives, to a capacity source for the storage pool.
//
// `capacitySource` is the target capacity source for the drive(s). This property does not need to be
// specified if the storage pool only contains one capacity source, or if the implementation is
// capable of automatically selecting the appropriate capacity source.
// `drives` is the existing drive or drives to be added to a capacity source of the storage pool. The
// implementation may impose restrictions on the number of drives added simultaneously.
func (storagepool *StoragePool) AddDrives(capacitySource *CapacitySource, drives []*redfish.Drive) error {
	if storagepool.addDrivesTarget == "" {
		return errors.New("action not supported by this service")
	}

	payload := struct {
		CapacitySource string
		Drives         []string
	}{}

	if capacitySource != nil {
		payload.CapacitySource = capacitySource.ODataID
	}

	for _, drive := range drives {
		payload.Drives = append(payload.Drives, drive.ODataID)
	}

	return storagepool.Post(storagepool.addDrivesTarget, payload)
}

// RemoveDrives will remove drive(s) from the capacity source for the StoragePool.
//
// `drives` is the drive or drives to be removed from the underlying capacity source.
func (storagepool *StoragePool) RemoveDrives(drives []*redfish.Drive) error {
	if storagepool.removeDrivesTarget == "" {
		return errors.New("action not supported by this service")
	}

	payload := struct {
		Drives []string
	}{}

	for _, drive := range drives {
		payload.Drives = append(payload.Drives, drive.ODataID)
	}

	return storagepool.Post(storagepool.removeDrivesTarget, payload)
}

// SetCompressionState will set the compression state of the storage pool.
// This may be both a highly impactful, as well as a long running operation.
//
// `enable` indicates the desired compression state of the storage pool.
func (storagepool *StoragePool) SetCompressionState(enable bool) error {
	if storagepool.setCompressionStateTarget == "" {
		return errors.New("action not supported by this service")
	}

	payload := struct {
		Enable bool
	}{Enable: enable}

	return storagepool.Post(storagepool.setCompressionStateTarget, payload)
}

// SetDeduplicationState will set the dedupe state of the storage pool.
// This may be both a highly impactful, as well as a long running operation.
//
// `enable` indicates the desired deduplication state of the storage pool.
func (storagepool *StoragePool) SetDeduplicationState(enable bool) error {
	if storagepool.setCompressionStateTarget == "" {
		return errors.New("action not supported by this service")
	}

	payload := struct {
		Enable bool
	}{Enable: enable}

	return storagepool.Post(storagepool.setCompressionStateTarget, payload)
}

// SetEncryptionState set the encryption state of the storage pool.
// This may be both a highly impactful, as well as a long running operation.
//
// `enable` indicates the desired encryption state of the storage pool.
func (storagepool *StoragePool) SetEncryptionState(enable bool) error {
	if storagepool.setEncryptionStateTarget == "" {
		return errors.New("action not supported by this service")
	}

	payload := struct {
		Enable bool
	}{Enable: enable}

	return storagepool.Post(storagepool.setEncryptionStateTarget, payload)
}
