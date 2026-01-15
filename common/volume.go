//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.6 - #Volume.v1_10_2.Volume

package common

import (
	"encoding/json"
)

type EncryptionTypes string

const (
	// NativeDriveEncryptionEncryptionTypes The volume is utilizing the native
	// drive encryption capabilities of the drive hardware.
	NativeDriveEncryptionEncryptionTypes EncryptionTypes = "NativeDriveEncryption"
	// ControllerAssistedEncryptionTypes The volume is being encrypted by the
	// storage controller entity.
	ControllerAssistedEncryptionTypes EncryptionTypes = "ControllerAssisted"
	// SoftwareAssistedEncryptionTypes The volume is being encrypted by software
	// running on the system or the operating system.
	SoftwareAssistedEncryptionTypes EncryptionTypes = "SoftwareAssisted"
)

type InitializeMethod string

const (
	// SkipInitializeMethod The volume will be available for use immediately, with
	// no preparation.
	SkipInitializeMethod InitializeMethod = "Skip"
	// BackgroundInitializeMethod The volume will be available for use immediately,
	// with data erasure and preparation to happen as background tasks.
	BackgroundInitializeMethod InitializeMethod = "Background"
	// ForegroundInitializeMethod Data erasure and preparation tasks will complete
	// before the volume is presented as available for use.
	ForegroundInitializeMethod InitializeMethod = "Foreground"
)

type InitializeType string

const (
	// FastInitializeType The volume is prepared for use quickly, typically by
	// erasing just the beginning and end of the space so that partitioning can be
	// performed.
	FastInitializeType InitializeType = "Fast"
	// SlowInitializeType The volume is prepared for use slowly, typically by
	// completely erasing the volume.
	SlowInitializeType InitializeType = "Slow"
)

// LBAFormatType is LBAFormatType is defined in the NVMe specification set. This
// field indicates the LBA data size supported; implementations may report up to
// 16 values. For more details refer to the appropriate NVMe specification.
type LBAFormatType string

const (
	// LBAFormat0LBAFormatType LBAFormat0 is a required type. Indicates the LBA
	// data size supported.
	LBAFormat0LBAFormatType LBAFormatType = "LBAFormat0"
	// LBAFormat1LBAFormatType Indicates the LBA data size if supported.
	LBAFormat1LBAFormatType LBAFormatType = "LBAFormat1"
	// LBAFormat2LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat2LBAFormatType LBAFormatType = "LBAFormat2"
	// LBAFormat3LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat3LBAFormatType LBAFormatType = "LBAFormat3"
	// LBAFormat4LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat4LBAFormatType LBAFormatType = "LBAFormat4"
	// LBAFormat5LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat5LBAFormatType LBAFormatType = "LBAFormat5"
	// LBAFormat6LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat6LBAFormatType LBAFormatType = "LBAFormat6"
	// LBAFormat7LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat7LBAFormatType LBAFormatType = "LBAFormat7"
	// LBAFormat8LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat8LBAFormatType LBAFormatType = "LBAFormat8"
	// LBAFormat9LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat9LBAFormatType LBAFormatType = "LBAFormat9"
	// LBAFormat10LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat10LBAFormatType LBAFormatType = "LBAFormat10"
	// LBAFormat11LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat11LBAFormatType LBAFormatType = "LBAFormat11"
	// LBAFormat12LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat12LBAFormatType LBAFormatType = "LBAFormat12"
	// LBAFormat13LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat13LBAFormatType LBAFormatType = "LBAFormat13"
	// LBAFormat14LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat14LBAFormatType LBAFormatType = "LBAFormat14"
	// LBAFormat15LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat15LBAFormatType LBAFormatType = "LBAFormat15"
)

// LBARelativePerformanceType is Indicate the relative performance of the LBA
// format relative to other LBA formats supported by the controller.
type LBARelativePerformanceType string

const (
	// BestLBARelativePerformanceType Best performance.
	BestLBARelativePerformanceType LBARelativePerformanceType = "Best"
	// BetterLBARelativePerformanceType Better performance.
	BetterLBARelativePerformanceType LBARelativePerformanceType = "Better"
	// GoodLBARelativePerformanceType Good performance.
	GoodLBARelativePerformanceType LBARelativePerformanceType = "Good"
	// DegradedLBARelativePerformanceType Degraded performance.
	DegradedLBARelativePerformanceType LBARelativePerformanceType = "Degraded"
)

type NamespaceType string

const (
	// BlockNamespaceType The namespace is configured for use with a block storage
	// interface.
	BlockNamespaceType NamespaceType = "Block"
	// KeyValueNamespaceType The namespace is configured for use with a KeyValue
	// interface.
	KeyValueNamespaceType NamespaceType = "KeyValue"
	// ZNSNamespaceType The namespace is configured for use with a zoned storage
	// interface.
	ZNSNamespaceType NamespaceType = "ZNS"
	// ComputationalNamespaceType The namespace is configured for use with a
	// computational storage interface.
	ComputationalNamespaceType NamespaceType = "Computational"
)

type OperationType string

const (
	// DeduplicateOperationType is a Deduplicate operation is being performed.
	DeduplicateOperationType OperationType = "Deduplicate"
	// CheckConsistencyOperationType is a CheckConsistency operation is being
	// performed.
	CheckConsistencyOperationType OperationType = "CheckConsistency"
	// InitializeOperationType is an Initialize operation is being performed.
	InitializeOperationType OperationType = "Initialize"
	// ReplicateOperationType is a Replicate operation is being performed.
	ReplicateOperationType OperationType = "Replicate"
	// DeleteOperationType is a Delete operation is being performed.
	DeleteOperationType OperationType = "Delete"
	// ChangeRAIDTypeOperationType is a ChangeRAIDType operation is being
	// performed.
	ChangeRAIDTypeOperationType OperationType = "ChangeRAIDType"
	// RebuildOperationType is a Rebuild operation is being performed.
	RebuildOperationType OperationType = "Rebuild"
	// EncryptOperationType is an Encrypt operation is being performed.
	EncryptOperationType OperationType = "Encrypt"
	// DecryptOperationType is a Decrypt operation is being performed.
	DecryptOperationType OperationType = "Decrypt"
	// ResizeOperationType is a Resize operation is being performed.
	ResizeOperationType OperationType = "Resize"
	// CompressOperationType is a Compress operation is being performed.
	CompressOperationType OperationType = "Compress"
	// SanitizeOperationType is a Sanitize operation is being performed.
	SanitizeOperationType OperationType = "Sanitize"
	// FormatOperationType is a Format operation is being performed.
	FormatOperationType OperationType = "Format"
	// ChangeStripSizeOperationType is a ChangeStripSize operation is being
	// performed.
	ChangeStripSizeOperationType OperationType = "ChangeStripSize"
)

type RAIDType string

const (
	// RAID0RAIDType is a placement policy where consecutive logical blocks of data
	// are uniformly distributed across a set of independent storage devices
	// without offering any form of redundancy. This is commonly referred to as
	// data striping. This form of RAID will encounter data loss with the failure
	// of any storage device in the set.
	RAID0RAIDType RAIDType = "RAID0"
	// RAID1RAIDType is a placement policy where each logical block of data is
	// stored on more than one independent storage device. This is commonly
	// referred to as mirroring. Data stored using this form of RAID is able to
	// survive a single storage device failure without data loss.
	RAID1RAIDType RAIDType = "RAID1"
	// RAID3RAIDType is a placement policy using parity-based protection where
	// logical bytes of data are uniformly distributed across a set of independent
	// storage devices and where the parity is stored on a dedicated independent
	// storage device. Data stored using this form of RAID is able to survive a
	// single storage device failure without data loss. If the storage devices use
	// rotating media, they are assumed to be rotationally synchronized, and the
	// data stripe size should be no larger than the exported block size.
	RAID3RAIDType RAIDType = "RAID3"
	// RAID4RAIDType is a placement policy using parity-based protection where
	// logical blocks of data are uniformly distributed across a set of independent
	// storage devices and where the parity is stored on a dedicated independent
	// storage device. Data stored using this form of RAID is able to survive a
	// single storage device failure without data loss.
	RAID4RAIDType RAIDType = "RAID4"
	// RAID5RAIDType is a placement policy using parity-based protection for
	// storing stripes of 'n' logical blocks of data and one logical block of
	// parity across a set of 'n+1' independent storage devices where the parity
	// and data blocks are interleaved across the storage devices. Data stored
	// using this form of RAID is able to survive a single storage device failure
	// without data loss.
	RAID5RAIDType RAIDType = "RAID5"
	// RAID6RAIDType is a placement policy using parity-based protection for
	// storing stripes of 'n' logical blocks of data and two logical blocks of
	// independent parity across a set of 'n+2' independent storage devices where
	// the parity and data blocks are interleaved across the storage devices. Data
	// stored using this form of RAID is able to survive any two independent
	// storage device failures without data loss.
	RAID6RAIDType RAIDType = "RAID6"
	// RAID10RAIDType is a placement policy that creates a striped device (RAID 0)
	// over a set of mirrored devices (RAID 1). This is commonly referred to as
	// RAID 1/0. Data stored using this form of RAID is able to survive storage
	// device failures in each RAID 1 set without data loss.
	RAID10RAIDType RAIDType = "RAID10"
	// RAID01RAIDType is a data placement policy that creates a mirrored device
	// (RAID 1) over a set of striped devices (RAID 0). This is commonly referred
	// to as RAID 0+1 or RAID 0/1. Data stored using this form of RAID is able to
	// survive a single RAID 0 data set failure without data loss.
	RAID01RAIDType RAIDType = "RAID01"
	// RAID6TPRAIDType is a placement policy that uses parity-based protection for
	// storing stripes of 'n' logical blocks of data and three logical blocks of
	// independent parity across a set of 'n+3' independent storage devices where
	// the parity and data blocks are interleaved across the storage devices. This
	// is commonly referred to as Triple Parity RAID. Data stored using this form
	// of RAID is able to survive any three independent storage device failures
	// without data loss.
	RAID6TPRAIDType RAIDType = "RAID6TP"
	// RAID1ERAIDType is a placement policy that uses a form of mirroring
	// implemented over a set of independent storage devices where logical blocks
	// are duplicated on a pair of independent storage devices so that data is
	// uniformly distributed across the storage devices. This is commonly referred
	// to as RAID 1 Enhanced. Data stored using this form of RAID is able to
	// survive a single storage device failure without data loss.
	RAID1ERAIDType RAIDType = "RAID1E"
	// RAID50RAIDType is a placement policy that uses a RAID 0 stripe set over two
	// or more RAID 5 sets of independent storage devices. Data stored using this
	// form of RAID is able to survive a single storage device failure within each
	// RAID 5 set without data loss.
	RAID50RAIDType RAIDType = "RAID50"
	// RAID60RAIDType is a placement policy that uses a RAID 0 stripe set over two
	// or more RAID 6 sets of independent storage devices. Data stored using this
	// form of RAID is able to survive two device failures within each RAID 6 set
	// without data loss.
	RAID60RAIDType RAIDType = "RAID60"
	// RAID00RAIDType is a placement policy that creates a RAID 0 stripe set over
	// two or more RAID 0 sets. This is commonly referred to as RAID 0+0. This form
	// of data layout is not fault tolerant; if any storage device fails there will
	// be data loss.
	RAID00RAIDType RAIDType = "RAID00"
	// RAID10ERAIDType is a placement policy that uses a RAID 0 stripe set over two
	// or more RAID 10 sets. This is commonly referred to as Enhanced RAID 10. Data
	// stored using this form of RAID is able to survive a single device failure
	// within each nested RAID 1 set without data loss.
	RAID10ERAIDType RAIDType = "RAID10E"
	// RAID1TripleRAIDType is a placement policy where each logical block of data
	// is mirrored three times across a set of three independent storage devices.
	// This is commonly referred to as three-way mirroring. This form of RAID can
	// survive two device failures without data loss.
	RAID1TripleRAIDType RAIDType = "RAID1Triple"
	// RAID10TripleRAIDType is a placement policy that uses a striped device (RAID
	// 0) over a set of triple mirrored devices (RAID 1Triple). This form of RAID
	// can survive up to two failures in each triple mirror set without data loss.
	RAID10TripleRAIDType RAIDType = "RAID10Triple"
	// NoneRAIDType is a placement policy with no redundancy at the device level.
	NoneRAIDType RAIDType = "None"
)

type ReadCachePolicyType string

const (
	// ReadAheadReadCachePolicyType is a caching technique in which the controller
	// pre-fetches data anticipating future read requests.
	ReadAheadReadCachePolicyType ReadCachePolicyType = "ReadAhead"
	// AdaptiveReadAheadReadCachePolicyType is a caching technique in which the
	// controller dynamically determines whether to pre-fetch data anticipating
	// future read requests, based on previous cache hit ratio.
	AdaptiveReadAheadReadCachePolicyType ReadCachePolicyType = "AdaptiveReadAhead"
	// OffReadCachePolicyType The read cache is disabled.
	OffReadCachePolicyType ReadCachePolicyType = "Off"
)

type VolumeType string

const (
	// RawDeviceVolumeType The volume is a raw physical device without any RAID or
	// other virtualization applied.
	RawDeviceVolumeType VolumeType = "RawDevice"
	// NonRedundantVolumeType The volume is a non-redundant storage device.
	NonRedundantVolumeType VolumeType = "NonRedundant"
	// MirroredVolumeType The volume is a mirrored device.
	MirroredVolumeType VolumeType = "Mirrored"
	// StripedWithParityVolumeType The volume is a device which uses parity to
	// retain redundant information.
	StripedWithParityVolumeType VolumeType = "StripedWithParity"
	// SpannedMirrorsVolumeType The volume is a spanned set of mirrored devices.
	SpannedMirrorsVolumeType VolumeType = "SpannedMirrors"
	// SpannedStripesWithParityVolumeType The volume is a spanned set of devices
	// which uses parity to retain redundant information.
	SpannedStripesWithParityVolumeType VolumeType = "SpannedStripesWithParity"
)

type VolumeUsageType string

const (
	// DataVolumeUsageType shall be allocated for use as a consumable data volume.
	DataVolumeUsageType VolumeUsageType = "Data"
	// SystemDataVolumeUsageType shall be allocated for use as a consumable data
	// volume reserved for system use.
	SystemDataVolumeUsageType VolumeUsageType = "SystemData"
	// CacheOnlyVolumeUsageType shall be allocated for use as a non-consumable
	// cache only volume.
	CacheOnlyVolumeUsageType VolumeUsageType = "CacheOnly"
	// SystemReserveVolumeUsageType shall be allocated for use as a non-consumable
	// system reserved volume.
	SystemReserveVolumeUsageType VolumeUsageType = "SystemReserve"
	// ReplicationReserveVolumeUsageType shall be allocated for use as a
	// non-consumable reserved volume for replication use.
	ReplicationReserveVolumeUsageType VolumeUsageType = "ReplicationReserve"
)

type WriteCachePolicyType string

const (
	// WriteThroughWriteCachePolicyType is a caching technique in which the
	// completion of a write request is not signaled until data is safely stored on
	// non-volatile media.
	WriteThroughWriteCachePolicyType WriteCachePolicyType = "WriteThrough"
	// ProtectedWriteBackWriteCachePolicyType is a caching technique in which the
	// completion of a write request is signaled as soon as the data is in cache,
	// and actual writing to non-volatile media is guaranteed to occur at a later
	// time.
	ProtectedWriteBackWriteCachePolicyType WriteCachePolicyType = "ProtectedWriteBack"
	// UnprotectedWriteBackWriteCachePolicyType is a caching technique in which the
	// completion of a write request is signaled as soon as the data is in cache;
	// actual writing to non-volatile media is not guaranteed to occur at a later
	// time.
	UnprotectedWriteBackWriteCachePolicyType WriteCachePolicyType = "UnprotectedWriteBack"
	// OffWriteCachePolicyType shall be disabled.
	OffWriteCachePolicyType WriteCachePolicyType = "Off"
)

type WriteCacheStateType string

const (
	// UnprotectedWriteCacheStateType Indicates that the cache state type in use
	// generally does not protect write requests on non-volatile media.
	UnprotectedWriteCacheStateType WriteCacheStateType = "Unprotected"
	// ProtectedWriteCacheStateType Indicates that the cache state type in use
	// generally protects write requests on non-volatile media.
	ProtectedWriteCacheStateType WriteCacheStateType = "Protected"
	// DegradedWriteCacheStateType Indicates an issue with the cache state in which
	// the cache space is diminished or disabled due to a failure or an outside
	// influence such as a discharged battery.
	DegradedWriteCacheStateType WriteCacheStateType = "Degraded"
)

type WriteHoleProtectionPolicyType string

const (
	// OffWriteHoleProtectionPolicyType The support for addressing the write hole
	// issue is disabled. The volume is not performing any additional activities to
	// close the RAID write hole.
	OffWriteHoleProtectionPolicyType WriteHoleProtectionPolicyType = "Off"
	// JournalingWriteHoleProtectionPolicyType The policy that uses separate block
	// device for write-ahead logging to address write hole issue. All write
	// operations on the RAID volume are first logged on dedicated journaling
	// device that is not part of the volume.
	JournalingWriteHoleProtectionPolicyType WriteHoleProtectionPolicyType = "Journaling"
	// DistributedLogWriteHoleProtectionPolicyType The policy that distributes
	// additional log (e.q. checksum of the parity) among the volume's capacity
	// sources to address write hole issue. Additional data is used to detect data
	// corruption on the volume.
	DistributedLogWriteHoleProtectionPolicyType WriteHoleProtectionPolicyType = "DistributedLog"
	// OemWriteHoleProtectionPolicyType The policy that is Oem specific. The
	// mechanism details are unknown unless provided separately by the Oem.
	OemWriteHoleProtectionPolicyType WriteHoleProtectionPolicyType = "Oem"
)

// Volume shall be used to represent a volume, virtual disk, logical disk, LUN,
// or other logical storage for a Redfish implementation.
type Volume struct {
	Entity
	// ALUA shall identify the ALUA properties for this volume.
	//
	// Version added: v1.10.0
	ALUA ALUA
	// AccessCapabilities shall specify a current storage access capability.
	//
	// Version added: v1.1.0
	// AccessCapabilities []StorageAccessCapability
	// AllocatedPools shall contain references to all storage pools allocated from
	// this volume.
	//
	// Version added: v1.1.0
	AllocatedPools string
	// BlockSizeBytes shall contain size of the smallest addressable unit of the
	// associated volume.
	BlockSizeBytes *int `json:",omitempty"`
	// Capacity Information about the utilization of capacity allocated to this
	// storage volume.
	//
	// Version added: v1.1.0
	// Capacity Capacity
	// CapacityBytes shall contain the size in bytes of the associated volume.
	CapacityBytes *int `json:",omitempty"`
	// CapacitySources Fully or partially consumed storage from a source resource.
	// Each entry provides capacity allocation information from a named source
	// resource.
	//
	// Version added: v1.1.0
	// CapacitySources []CapacitySource
	// CapacitySources@odata.count
	CapacitySourcesCount int `json:"CapacitySources@odata.count"`
	// Compressed shall contain a boolean indicator if the Volume is currently
	// utilizing compression or not.
	//
	// Version added: v1.4.0
	Compressed bool
	// Connections shall contain references to all Connections that include this
	// volume.
	//
	// Version added: v1.9.0
	// Connections []Connection
	// Connections@odata.count
	ConnectionsCount int `json:"Connections@odata.count"`
	// Deduplicated shall contain a boolean indicator if the Volume is currently
	// utilizing deduplication or not.
	//
	// Version added: v1.4.0
	Deduplicated bool
	// DisplayName shall contain a user-configurable string to name the volume.
	//
	// Version added: v1.4.0
	DisplayName string
	// Encrypted shall contain a boolean indicator if the Volume is currently
	// utilizing encryption or not.
	Encrypted bool
	// EncryptionTypes shall contain the types of encryption used by this Volume.
	EncryptionTypes []EncryptionTypes
	// IOPerfModeEnabled shall indicate whether IO performance mode is enabled for
	// the volume.
	//
	// Version added: v1.5.0
	IOPerfModeEnabled bool
	// IOStatistics shall represent IO statistics for this volume.
	//
	// Version added: v1.2.0
	//
	// Deprecated: v1.10.0
	// This property is deprecated in favor of the IOStatistics property in
	// VolumeMetrics.
	IOStatistics IOStatistics
	// Identifiers shall contain a list of all known durable names for the
	// associated volume.
	Identifiers []Identifier
	// InitializeMethod shall indicate the initialization method used for this
	// volume. If InitializeMethod is not specified, the InitializeMethod should be
	// Foreground. This value reflects the most recently used Initialization
	// Method, and may be changed using the Initialize Action.
	//
	// Version added: v1.6.0
	InitializeMethod InitializeMethod
	// IsBootCapable shall indicate whether or not the Volume contains a boot image
	// and is capable of booting. This property may be settable by an admin or
	// client with visibility into the contents of the volume. This property should
	// only be set to true when VolumeUsage is either not specified, or when
	// VolumeUsage is set to Data or SystemData.
	//
	// Version added: v1.7.0
	IsBootCapable bool
	// LogicalUnitNumber shall contain host-visible LogicalUnitNumber assigned to
	// this Volume. This property shall only be used when in a single connect
	// configuration and no StorageGroup configuration is used.
	//
	// Version added: v1.4.0
	LogicalUnitNumber *int `json:",omitempty"`
	// LowSpaceWarningThresholdPercents shall be triggered: Across all
	// CapacitySources entries, percent = (SUM(AllocatedBytes) -
	// SUM(ConsumedBytes))/SUM(AllocatedBytes).
	//
	// Version added: v1.1.0
	LowSpaceWarningThresholdPercents []*int
	// Manufacturer shall contain a value that represents the manufacturer or
	// implementer of the storage volume.
	//
	// Version added: v1.1.0
	Manufacturer string
	// MaxBlockSizeBytes shall contain size of the largest addressable unit of this
	// storage volume.
	//
	// Version added: v1.1.0
	MaxBlockSizeBytes *int `json:",omitempty"`
	// MediaSpanCount shall indicate the number of media elements used per span in
	// the secondary RAID for a hierarchical RAID type.
	//
	// Version added: v1.4.0
	MediaSpanCount *int `json:",omitempty"`
	// Metrics shall contain a link to a resource of type VolumeMetrics that
	// specifies the metrics for this volume. IO metrics are reported in the
	// IOStatistics property.
	//
	// Version added: v1.9.0
	Metrics string
	// Model shall represents a specific storage volume implementation.
	//
	// Version added: v1.1.0
	Model string
	// NVMeNamespaceProperties shall contain properties to use when Volume is used
	// to describe an NVMe Namespace.
	//
	// Version added: v1.5.0
	NVMeNamespaceProperties NVMeNamespaceProperties
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Operations shall contain a list of all currently running on the Volume.
	Operations []Operation
	// OptimumIOSizeBytes shall contain the optimum IO size to use when performing
	// IO on this volume. For logical disks, this is the stripe size. For physical
	// disks, this describes the physical sector size.
	OptimumIOSizeBytes *int `json:",omitempty"`
	// ProvisioningPolicy shall specify the volume's supported storage allocation
	// policy.
	//
	// Version added: v1.4.0
	// ProvisioningPolicy ProvisioningPolicy
	// RAIDType shall contain the RAID type of the associated Volume.
	//
	// Version added: v1.3.1
	RAIDType RAIDType
	// ReadCachePolicy shall contain a boolean indicator of the read cache policy
	// for the Volume.
	//
	// Version added: v1.4.0
	ReadCachePolicy ReadCachePolicyType
	// RecoverableCapacitySourceCount The value is the number of available capacity
	// source resources currently available in the event that an equivalent
	// capacity source resource fails.
	//
	// Version added: v1.3.0
	RecoverableCapacitySourceCount *int `json:",omitempty"`
	// RemainingCapacityPercent shall return {[(SUM(AllocatedBytes) -
	// SUM(ConsumedBytes)]/SUM(AllocatedBytes)}*100 represented as an integer
	// value.
	//
	// Version added: v1.2.0
	RemainingCapacityPercent *int `json:",omitempty"`
	// RemoteReplicaTargets shall reference the URIs to the remote target replicas
	// that are sourced by this replica. Remote indicates that the replica is
	// managed by a separate Swordfish service instance.
	//
	// Version added: v1.8.0
	RemoteReplicaTargets []string
	// ReplicaInfo shall describe the replica relationship between this storage
	// volume and a corresponding source volume.
	//
	// Version added: v1.1.0
	ReplicaInfo ReplicaInfo
	// ReplicaTargets shall reference the target replicas that are sourced by this
	// replica.
	//
	// Version added: v1.3.0
	ReplicaTargets []Entity
	// ReplicaTargets@odata.count
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// ReplicationEnabled shall indicate whether or not replication is enabled on
	// the volume. This property shall be consistent with the state reflected at
	// the storage pool level.
	//
	// Version added: v1.9.0
	ReplicationEnabled bool
	// Status shall contain the status of the Volume.
	Status Status
	// StorageGroups shall contain references to all storage groups that include
	// this volume.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.9.0
	// This property is deprecated in favor of the Connections property.
	// StorageGroups string
	// StripSizeBytes The number of consecutively addressed virtual disk blocks
	// (bytes) mapped to consecutively addressed blocks on a single member extent
	// of a disk array. Synonym for stripe depth and chunk size.
	//
	// Version added: v1.4.0
	StripSizeBytes *int `json:",omitempty"`
	// VolumeType shall contain the type of the associated Volume.
	//
	// Deprecated
	// Deprecated in favor of explicit use of RAIDType.
	VolumeType VolumeType
	// VolumeUsage shall contain the volume usage type for the Volume.
	//
	// Version added: v1.4.0
	VolumeUsage VolumeUsageType
	// WriteCachePolicy shall contain a boolean indicator of the write cache policy
	// for the Volume.
	//
	// Version added: v1.4.0
	WriteCachePolicy WriteCachePolicyType
	// WriteCacheState shall contain the WriteCacheState policy setting for the
	// Volume.
	//
	// Version added: v1.4.0
	WriteCacheState WriteCacheStateType
	// WriteHoleProtectionPolicy shall be set to 'Off'.
	//
	// Version added: v1.4.0
	WriteHoleProtectionPolicy WriteHoleProtectionPolicyType
	// assignReplicaTargetTarget is the URL to send AssignReplicaTarget requests.
	assignReplicaTargetTarget string
	// changeRAIDLayoutTarget is the URL to send ChangeRAIDLayout requests.
	changeRAIDLayoutTarget string
	// checkConsistencyTarget is the URL to send CheckConsistency requests.
	checkConsistencyTarget string
	// createReplicaTargetTarget is the URL to send CreateReplicaTarget requests.
	createReplicaTargetTarget string
	// forceEnableTarget is the URL to send ForceEnable requests.
	forceEnableTarget string
	// initializeTarget is the URL to send Initialize requests.
	initializeTarget string
	// removeReplicaRelationshipTarget is the URL to send RemoveReplicaRelationship requests.
	removeReplicaRelationshipTarget string
	// resumeReplicationTarget is the URL to send ResumeReplication requests.
	resumeReplicationTarget string
	// reverseReplicationRelationshipTarget is the URL to send ReverseReplicationRelationship requests.
	reverseReplicationRelationshipTarget string
	// splitReplicationTarget is the URL to send SplitReplication requests.
	splitReplicationTarget string
	// suspendReplicationTarget is the URL to send SuspendReplication requests.
	suspendReplicationTarget string
	// cacheDataVolumes are the URIs for CacheDataVolumes.
	cacheDataVolumes []string
	// cacheVolumeSource is the URI for CacheVolumeSource.
	cacheVolumeSource string
	// classOfService is the URI for ClassOfService.
	classOfService string
	// clientEndpoints are the URIs for ClientEndpoints.
	clientEndpoints []string
	// consistencyGroups are the URIs for ConsistencyGroups.
	consistencyGroups []string
	// controllers are the URIs for Controllers.
	controllers []string
	// dedicatedSpareDrives are the URIs for DedicatedSpareDrives.
	dedicatedSpareDrives []string
	// drives are the URIs for Drives.
	drives []string
	// journalingMedia is the URI for JournalingMedia.
	journalingMedia string
	// owningStorageResource is the URI for OwningStorageResource.
	owningStorageResource string
	// owningStorageService is the URI for OwningStorageService.
	owningStorageService string
	// providingStoragePool is the URI for ProvidingStoragePool.
	providingStoragePool string
	// serverEndpoints are the URIs for ServerEndpoints.
	serverEndpoints []string
	// spareResourceSets are the URIs for SpareResourceSets.
	spareResourceSets []string
	// storageGroups are the URIs for StorageGroups.
	storageGroups []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Volume object from the raw JSON.
func (v *Volume) UnmarshalJSON(b []byte) error {
	type temp Volume
	type vActions struct {
		AssignReplicaTarget            ActionTarget `json:"#Volume.AssignReplicaTarget"`
		ChangeRAIDLayout               ActionTarget `json:"#Volume.ChangeRAIDLayout"`
		CheckConsistency               ActionTarget `json:"#Volume.CheckConsistency"`
		CreateReplicaTarget            ActionTarget `json:"#Volume.CreateReplicaTarget"`
		ForceEnable                    ActionTarget `json:"#Volume.ForceEnable"`
		Initialize                     ActionTarget `json:"#Volume.Initialize"`
		RemoveReplicaRelationship      ActionTarget `json:"#Volume.RemoveReplicaRelationship"`
		ResumeReplication              ActionTarget `json:"#Volume.ResumeReplication"`
		ReverseReplicationRelationship ActionTarget `json:"#Volume.ReverseReplicationRelationship"`
		SplitReplication               ActionTarget `json:"#Volume.SplitReplication"`
		SuspendReplication             ActionTarget `json:"#Volume.SuspendReplication"`
	}
	type vLinks struct {
		CacheDataVolumes      Links `json:"CacheDataVolumes"`
		CacheVolumeSource     Link  `json:"CacheVolumeSource"`
		ClassOfService        Link  `json:"ClassOfService"`
		ClientEndpoints       Links `json:"ClientEndpoints"`
		ConsistencyGroups     Links `json:"ConsistencyGroups"`
		Controllers           Links `json:"Controllers"`
		DedicatedSpareDrives  Links `json:"DedicatedSpareDrives"`
		Drives                Links `json:"Drives"`
		JournalingMedia       Link  `json:"JournalingMedia"`
		OwningStorageResource Link  `json:"OwningStorageResource"`
		OwningStorageService  Link  `json:"OwningStorageService"`
		ProvidingStoragePool  Link  `json:"ProvidingStoragePool"`
		ServerEndpoints       Links `json:"ServerEndpoints"`
		SpareResourceSets     Links `json:"SpareResourceSets"`
		StorageGroups         Links `json:"StorageGroups"`
	}
	var tmp struct {
		temp
		Actions vActions
		Links   vLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*v = Volume(tmp.temp)

	// Extract the links to other entities for later
	v.assignReplicaTargetTarget = tmp.Actions.AssignReplicaTarget.Target
	v.changeRAIDLayoutTarget = tmp.Actions.ChangeRAIDLayout.Target
	v.checkConsistencyTarget = tmp.Actions.CheckConsistency.Target
	v.createReplicaTargetTarget = tmp.Actions.CreateReplicaTarget.Target
	v.forceEnableTarget = tmp.Actions.ForceEnable.Target
	v.initializeTarget = tmp.Actions.Initialize.Target
	v.removeReplicaRelationshipTarget = tmp.Actions.RemoveReplicaRelationship.Target
	v.resumeReplicationTarget = tmp.Actions.ResumeReplication.Target
	v.reverseReplicationRelationshipTarget = tmp.Actions.ReverseReplicationRelationship.Target
	v.splitReplicationTarget = tmp.Actions.SplitReplication.Target
	v.suspendReplicationTarget = tmp.Actions.SuspendReplication.Target
	v.cacheDataVolumes = tmp.Links.CacheDataVolumes.ToStrings()
	v.cacheVolumeSource = tmp.Links.CacheVolumeSource.String()
	v.classOfService = tmp.Links.ClassOfService.String()
	v.clientEndpoints = tmp.Links.ClientEndpoints.ToStrings()
	v.consistencyGroups = tmp.Links.ConsistencyGroups.ToStrings()
	v.controllers = tmp.Links.Controllers.ToStrings()
	v.dedicatedSpareDrives = tmp.Links.DedicatedSpareDrives.ToStrings()
	v.drives = tmp.Links.Drives.ToStrings()
	v.journalingMedia = tmp.Links.JournalingMedia.String()
	v.owningStorageResource = tmp.Links.OwningStorageResource.String()
	v.owningStorageService = tmp.Links.OwningStorageService.String()
	v.providingStoragePool = tmp.Links.ProvidingStoragePool.String()
	v.serverEndpoints = tmp.Links.ServerEndpoints.ToStrings()
	v.spareResourceSets = tmp.Links.SpareResourceSets.ToStrings()
	v.storageGroups = tmp.Links.StorageGroups.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	v.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (v *Volume) Update() error {
	readWriteFields := []string{
		"ALUA",
		"AccessCapabilities",
		"Capacity",
		"CapacityBytes",
		"CapacitySources",
		"CapacitySources@odata.count",
		"Compressed",
		"Connections@odata.count",
		"Deduplicated",
		"DisplayName",
		"Encrypted",
		"EncryptionTypes",
		"IOPerfModeEnabled",
		"IOStatistics",
		"Identifiers",
		"IsBootCapable",
		"LowSpaceWarningThresholdPercents",
		"NVMeNamespaceProperties",
		"Operations",
		"ProvisioningPolicy",
		"ReadCachePolicy",
		"RecoverableCapacitySourceCount",
		"ReplicaInfo",
		"ReplicaTargets@odata.count",
		"ReplicationEnabled",
		"Status",
		"StripSizeBytes",
		"WriteCachePolicy",
		"WriteHoleProtectionPolicy",
	}

	return v.UpdateFromRawData(v, v.rawData, readWriteFields)
}

// GetVolume will get a Volume instance from the service.
func GetVolume(c Client, uri string) (*Volume, error) {
	return GetObject[Volume](c, uri)
}

// ListReferencedVolumes gets the collection of Volume from
// a provided reference.
func ListReferencedVolumes(c Client, link string) ([]*Volume, error) {
	return GetCollectionObjects[Volume](c, link)
}

// AssignReplicaTarget shall be used to establish a replication relationship by
// assigning an existing volume to serve as a target replica for an existing
// source volume.
// replicaType - This parameter shall contain the type of replica relationship
// to be created (e.g., Clone, Mirror, Snap).
// replicaUpdateMode - This parameter shall specify the replica update mode.
// targetVolume - This parameter shall contain the Uri to the existing target
// volume.
func (v *Volume) AssignReplicaTarget(replicaType ReplicaType, replicaUpdateMode ReplicaUpdateMode, targetVolume string) error {
	payload := make(map[string]any)
	payload["ReplicaType"] = replicaType
	payload["ReplicaUpdateMode"] = replicaUpdateMode
	payload["TargetVolume"] = targetVolume
	return v.Post(v.assignReplicaTargetTarget, payload)
}

// ChangeRAIDLayout shall request the system to change the RAID layout of the
// volume. Depending on the combination of the submitted parameters, this could
// be changing the RAID type, changing the span count, changing the number of
// drives used by the volume, or another configuration change supported by the
// system. Note that usage of this action while online may potentially cause
// data loss if the available capacity is reduced.
// drives - This parameter shall contain an array of the drives to be used by
// the volume.
// mediaSpanCount - This parameter shall contain the requested number of media
// elements used per span in the secondary RAID for a hierarchical RAID type.
// rAIDType - This parameter shall contain the requested RAID type for the
// volume.
// stripSizeBytes - This parameter shall contain the number of blocks (bytes)
// requested for the strip size.
func (v *Volume) ChangeRAIDLayout(drives []string, mediaSpanCount int, rAIDType RAIDType, stripSizeBytes int) error {
	payload := make(map[string]any)
	payload["Drives"] = drives
	payload["MediaSpanCount"] = mediaSpanCount
	payload["RAIDType"] = rAIDType
	payload["StripSizeBytes"] = stripSizeBytes
	return v.Post(v.changeRAIDLayoutTarget, payload)
}

// CheckConsistency defines the name of the custom action supported on this resource.
func (v *Volume) CheckConsistency() error {
	payload := make(map[string]any)
	return v.Post(v.checkConsistencyTarget, payload)
}

// CreateReplicaTarget shall be used to create a new volume resource to provide
// expanded data protection through a replica relationship with the specified
// source volume.
// replicaType - This parameter shall contain the type of replica relationship
// to be created (e.g., Clone, Mirror, Snap).
// replicaUpdateMode - This parameter shall specify the replica update mode.
// targetStoragePool - This parameter shall contain the Uri to the existing
// StoragePool in which to create the target volume.
// volumeName - This parameter shall contain the Name for the target volume.
func (v *Volume) CreateReplicaTarget(replicaType ReplicaType, replicaUpdateMode ReplicaUpdateMode, targetStoragePool string, volumeName string) error {
	payload := make(map[string]any)
	payload["ReplicaType"] = replicaType
	payload["ReplicaUpdateMode"] = replicaUpdateMode
	payload["TargetStoragePool"] = targetStoragePool
	payload["VolumeName"] = volumeName
	return v.Post(v.createReplicaTargetTarget, payload)
}

// ForceEnable shall request the system to force the volume to enabled state
// regardless of data loss scenarios.
func (v *Volume) ForceEnable() error {
	payload := make(map[string]any)
	return v.Post(v.forceEnableTarget, payload)
}

// Initialize defines the name of the custom action supported on this resource. If
// InitializeMethod is not specified in the request body, but the property
// InitializeMethod is specified, the property InitializeMethod value should be
// used. If neither is specified, the InitializeMethod should be Foreground.
// initializeMethod - This defines the property name for the action.
// initializeType - This defines the property name for the action.
func (v *Volume) Initialize(initializeMethod InitializeMethod, initializeType InitializeType) error {
	payload := make(map[string]any)
	payload["InitializeMethod"] = initializeMethod
	payload["InitializeType"] = initializeType
	return v.Post(v.initializeTarget, payload)
}

// RemoveReplicaRelationship shall be used to disable data synchronization between a source
// and target volume, remove the replication relationship, and optionally
// delete the target volume.
// deleteTargetVolume - This parameter shall indicate whether or not to delete
// the target volume as part of the operation. If not defined, the system
// should use its default behavior.
// targetVolume - This parameter shall contain the Uri to the existing target
// volume.
func (v *Volume) RemoveReplicaRelationship(deleteTargetVolume bool, targetVolume string) error {
	payload := make(map[string]any)
	payload["DeleteTargetVolume"] = deleteTargetVolume
	payload["TargetVolume"] = targetVolume
	return v.Post(v.removeReplicaRelationshipTarget, payload)
}

// ResumeReplication shall be used to resume the active data synchronization between
// a source and target volume, without otherwise altering the replication
// relationship.
// targetVolume - This parameter shall contain the Uri to the existing target
// volume.
func (v *Volume) ResumeReplication(targetVolume string) error {
	payload := make(map[string]any)
	payload["TargetVolume"] = targetVolume
	return v.Post(v.resumeReplicationTarget, payload)
}

// ReverseReplicationRelationship shall be used to reverse the replication relationship between a
// source and target volume.
// targetVolume - This parameter shall contain the Uri to the existing target
// volume.
func (v *Volume) ReverseReplicationRelationship(targetVolume string) error {
	payload := make(map[string]any)
	payload["TargetVolume"] = targetVolume
	return v.Post(v.reverseReplicationRelationshipTarget, payload)
}

// SplitReplication shall be used to split the replication relationship and suspend
// data synchronization between a source and target volume.
// targetVolume - This parameter shall contain the Uri to the existing target
// volume.
func (v *Volume) SplitReplication(targetVolume string) error {
	payload := make(map[string]any)
	payload["TargetVolume"] = targetVolume
	return v.Post(v.splitReplicationTarget, payload)
}

// SuspendReplication shall be used to suspend active data synchronization between a
// source and target volume, without otherwise altering the replication
// relationship.
// targetVolume - This parameter shall contain the Uri to the existing target
// volume.
func (v *Volume) SuspendReplication(targetVolume string) error {
	payload := make(map[string]any)
	payload["TargetVolume"] = targetVolume
	return v.Post(v.suspendReplicationTarget, payload)
}

// CacheDataVolumes gets the CacheDataVolumes linked resources.
func (v *Volume) CacheDataVolumes(client Client) ([]*Volume, error) {
	return GetObjects[Volume](client, v.cacheDataVolumes)
}

// // CacheVolumeSource gets the CacheVolumeSource linked resource.
// func (v *Volume) CacheVolumeSource(client Client) (*CacheVolumeSource, error) {
// 	if v.cacheVolumeSource == "" {
// 		return nil, nil
// 	}
// 	return GetObject[CacheVolumeSource](client, v.cacheVolumeSource)
// }

// // ClassOfService gets the ClassOfService linked resource.
// func (v *Volume) ClassOfService(client Client) (*ClassOfService, error) {
// 	if v.classOfService == "" {
// 		return nil, nil
// 	}
// 	return GetObject[ClassOfService](client, v.classOfService)
// }

// // ClientEndpoints gets the ClientEndpoints linked resources.
// func (v *Volume) ClientEndpoints(client Client) ([]*Endpoint, error) {
// 	return GetObjects[Endpoint](client, v.clientEndpoints)
// }

// // ConsistencyGroups gets the ConsistencyGroups linked resources.
// func (v *Volume) ConsistencyGroups(client Client) ([]*ConsistencyGroup, error) {
// 	return GetObjects[ConsistencyGroup](client, v.consistencyGroups)
// }

// // Controllers gets the Controllers linked resources.
// func (v *Volume) Controllers(client Client) ([]*StorageController, error) {
// 	return GetObjects[StorageController](client, v.controllers)
// }

// // DedicatedSpareDrives gets the DedicatedSpareDrives linked resources.
// func (v *Volume) DedicatedSpareDrives(client Client) ([]*Drive, error) {
// 	return GetObjects[Drive](client, v.dedicatedSpareDrives)
// }

// // Drives gets the Drives linked resources.
// func (v *Volume) Drives(client Client) ([]*Drive, error) {
// 	return GetObjects[Drive](client, v.drives)
// }

// // JournalingMedia gets the JournalingMedia linked resource.
// func (v *Volume) JournalingMedia(client Client) (*JournalingMedia, error) {
// 	if v.journalingMedia == "" {
// 		return nil, nil
// 	}
// 	return GetObject[JournalingMedia](client, v.journalingMedia)
// }

// // OwningStorageResource gets the OwningStorageResource linked resource.
// func (v *Volume) OwningStorageResource(client Client) (*Storage, error) {
// 	if v.owningStorageResource == "" {
// 		return nil, nil
// 	}
// 	return GetObject[Storage](client, v.owningStorageResource)
// }

// // OwningStorageService gets the OwningStorageService linked resource.
// func (v *Volume) OwningStorageService(client Client) (*StorageService, error) {
// 	if v.owningStorageService == "" {
// 		return nil, nil
// 	}
// 	return GetObject[StorageService](client, v.owningStorageService)
// }

// // ProvidingStoragePool gets the ProvidingStoragePool linked resource.
// func (v *Volume) ProvidingStoragePool(client Client) (*ProvidingStoragePool, error) {
// 	if v.providingStoragePool == "" {
// 		return nil, nil
// 	}
// 	return GetObject[ProvidingStoragePool](client, v.providingStoragePool)
// }

// // ServerEndpoints gets the ServerEndpoints linked resources.
// func (v *Volume) ServerEndpoints(client Client) ([]*Endpoint, error) {
// 	return GetObjects[Endpoint](client, v.serverEndpoints)
// }

// // SpareResourceSets gets the SpareResourceSets linked resources.
// func (v *Volume) SpareResourceSets(client Client) ([]*SpareResourceSet, error) {
// 	return GetObjects[SpareResourceSet](client, v.spareResourceSets)
// }

// // StorageGroups gets the StorageGroups linked resources.
// func (v *Volume) StorageGroups(client Client) ([]*StorageGroup, error) {
// 	return GetObjects[StorageGroup](client, v.storageGroups)
// }

// ALUA represents the ALUA type.
type ALUA struct {
	// ANAGroupId shall contain the ANA group id for this volume.
	//
	// Version added: v1.10.0
	ANAGroupID *int `json:"ANAGroupId,omitempty"`
}

// LBAFormat represents the LBAFormat type.
type LBAFormat struct {
	// LBADataSizeBytes shall be the LBA data size reported in bytes.
	//
	// Version added: v1.9.0
	LBADataSizeBytes *int `json:",omitempty"`
	// LBAFormatType shall be the LBA format type. This property is intended for
	// capabilities instrumentation.
	//
	// Version added: v1.9.0
	LBAFormatType LBAFormatType
	// LBAMetadataSizeBytes shall be the LBA metadata size reported in bytes.
	//
	// Version added: v1.9.0
	LBAMetadataSizeBytes *int `json:",omitempty"`
	// RelativePerformance shall be the LBA Relative Performance type. This field
	// indicates the relative performance of the LBA format indicated relative to
	// other LBA formats supported by the controller. This property is intended for
	// capabilities instrumentation.
	//
	// Version added: v1.9.0
	RelativePerformance LBARelativePerformanceType
}

// NVMeNamespaceProperties This contains properties to use when Volume is used
// to describe an NVMe Namespace.
type NVMeNamespaceProperties struct {
	// FormattedLBASize shall contain the LBA data size and metadata size
	// combination that the namespace has been formatted with. This is a 4-bit data
	// structure.
	//
	// Version added: v1.5.0
	FormattedLBASize string
	// IsShareable shall indicate whether the namespace is shareable.
	//
	// Version added: v1.5.0
	IsShareable bool
	// LBAFormat shall describe the current LBA format ID and corresponding
	// detailed properties, such as the LBA data size and metadata size. Use the
	// LBAFormats property to describe namespace capabilities in a collection
	// capabilities annotation.
	//
	// Version added: v1.9.0
	LBAFormat LBAFormat
	// LBAFormats shall describe the LBA format IDs and corresponding detailed
	// properties, such as the LBA data size and metadata size. This property is
	// intended for use in a collection capabilities annotation. Use the LBAFormat
	// property on an instance of a namespace.
	//
	// Version added: v1.9.0
	LBAFormats []LBAFormat
	// LBAFormatsSupported shall be a list of the LBA formats supported for the
	// namespace, or potential namespaces.
	//
	// Version added: v1.8.0
	LBAFormatsSupported []LBAFormatType
	// MetadataTransferredAtEndOfDataLBA shall indicate whether or not the metadata
	// is transferred at the end of the LBA creating an extended data LBA.
	//
	// Version added: v1.5.0
	MetadataTransferredAtEndOfDataLBA bool
	// NVMeVersion shall contain the version of the NVMe Base Specification
	// supported.
	//
	// Version added: v1.5.0
	NVMeVersion string
	// NamespaceFeatures shall contain a set of Namespace Features.
	//
	// Version added: v1.5.0
	NamespaceFeatures NamespaceFeatures
	// NamespaceId shall contain the NVMe Namespace Identifier for this namespace.
	// This property shall be a hex value. Namespace identifiers are not durable
	// and do not have meaning outside the scope of the NVMe subsystem. NSID 0x0,
	// 0xFFFFFFFF, 0xFFFFFFFE are special purpose values.
	//
	// Version added: v1.5.0
	NamespaceID string `json:"NamespaceId"`
	// NamespaceType shall identify the type of namespace.
	//
	// Version added: v1.9.0
	NamespaceType NamespaceType
	// NumberLBAFormats shall contain the number of LBA data size and metadata size
	// combinations supported by this namespace. The value of this property is
	// between 0 and 16. LBA formats with an index set beyond this value will not
	// be supported.
	//
	// Version added: v1.5.0
	NumberLBAFormats *uint `json:",omitempty"`
	// SupportsIOPerformanceHints shall indicate whether the namespace supports IO
	// performance hints.
	//
	// Version added: v1.10.0
	SupportsIOPerformanceHints bool
	// SupportsMultipleNamespaceAttachments shall indicate whether the namespace
	// may be attached to two or more controllers.
	//
	// Version added: v1.10.0
	SupportsMultipleNamespaceAttachments bool
	// Type shall identify the type of namespace.
	//
	// Version added: v1.8.0
	Type NamespaceType
}

// NamespaceFeatures represents the NamespaceFeatures type.
type NamespaceFeatures struct {
	// SupportsAtomicTransactionSize shall indicate whether or not the NVM fields
	// for Namespace preferred write granularity (NPWG), write alignment (NPWA),
	// deallocate granularity (NPDG), deallocate alignment (NPDA) and optimal write
	// size (NOWS) are defined for this namespace and should be used by the host
	// for I/O optimization.
	//
	// Version added: v1.5.0
	SupportsAtomicTransactionSize bool
	// SupportsDeallocatedOrUnwrittenLBError shall indicate that the controller
	// supports deallocated or unwritten logical block error for this namespace.
	//
	// Version added: v1.5.0
	SupportsDeallocatedOrUnwrittenLBError bool
	// SupportsIOPerformanceHints shall indicate that the Namespace Atomic Write
	// Unit Normal (NAWUN), Namespace Atomic Write Unit Power Fail (NAWUPF), and
	// Namespace Atomic Compare and Write Unit (NACWU) fields are defined for this
	// namespace and should be used by the host for this namespace instead of the
	// controller-level properties AWUN, AWUPF, and ACWU.
	//
	// Version added: v1.5.0
	SupportsIOPerformanceHints bool
	// SupportsNGUIDReuse shall indicate that the namespace supports the use of an
	// NGUID (namespace globally unique identifier) value.
	//
	// Version added: v1.5.0
	SupportsNGUIDReuse bool
	// SupportsThinProvisioning shall indicate whether or not the NVMe Namespace
	// supports thin provisioning. Specifically, the namespace capacity reported
	// may be less than the namespace size.
	//
	// Version added: v1.5.0
	SupportsThinProvisioning bool
}

// Operation represents the Operation type.
type Operation struct {
	// AssociatedFeaturesRegistry is a reference to the task associated with the
	// operation if any.
	// AssociatedFeaturesRegistry FeaturesRegistry
	// Operation shall contain the type of the operation.
	//
	// Version added: v1.9.0
	Operation OperationType
	// OperationName The name of the operation.
	//
	// Deprecated: v1.9.0
	// This property is deprecated in favor of the Operation property using the
	// Operation enum.
	OperationName string
	// PercentageComplete The percentage of the operation that has been completed.
	PercentageComplete *int `json:",omitempty"`
}
