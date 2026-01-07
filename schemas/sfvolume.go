//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.6 - #Volume.v1_10_2.Volume

package schemas

import (
	"encoding/json"
)

type SFEncryptionTypes string

const (
	// NativeDriveEncryptionEncryptionTypes The volume is utilizing the native
	// drive encryption capabilities of the drive hardware.
	NativeDriveEncryptionSFEncryptionTypes SFEncryptionTypes = "NativeDriveEncryption"
	// ControllerAssistedEncryptionTypes The volume is being encrypted by the
	// storage controller entity.
	ControllerAssistedSFEncryptionTypes SFEncryptionTypes = "ControllerAssisted"
	// SoftwareAssistedEncryptionTypes The volume is being encrypted by software
	// running on the system or the operating system.
	SoftwareAssistedSFEncryptionTypes SFEncryptionTypes = "SoftwareAssisted"
)

type SFInitializeMethod string

const (
	// SkipInitializeMethod The volume will be available for use immediately, with
	// no preparation.
	SkipSFInitializeMethod SFInitializeMethod = "Skip"
	// BackgroundInitializeMethod The volume will be available for use immediately,
	// with data erasure and preparation to happen as background tasks.
	BackgroundSFInitializeMethod SFInitializeMethod = "Background"
	// ForegroundInitializeMethod Data erasure and preparation tasks will complete
	// before the volume is presented as available for use.
	ForegroundSFInitializeMethod SFInitializeMethod = "Foreground"
)

type SFInitializeType string

const (
	// FastInitializeType The volume is prepared for use quickly, typically by
	// erasing just the beginning and end of the space so that partitioning can be
	// performed.
	FastSFInitializeType SFInitializeType = "Fast"
	// SlowInitializeType The volume is prepared for use slowly, typically by
	// completely erasing the volume.
	SlowSFInitializeType SFInitializeType = "Slow"
)

// LBAFormatType is LBAFormatType is defined in the NVMe specification set. This
// field indicates the LBA data size supported; implementations may report up to
// 16 values. For more details refer to the appropriate NVMe specification.
type SFLBAFormatType string

const (
	// LBAFormat0LBAFormatType LBAFormat0 is a required type. Indicates the LBA
	// data size supported.
	LBAFormat0SFLBAFormatType SFLBAFormatType = "LBAFormat0"
	// LBAFormat1LBAFormatType Indicates the LBA data size if supported.
	LBAFormat1SFLBAFormatType SFLBAFormatType = "LBAFormat1"
	// LBAFormat2LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat2SFLBAFormatType SFLBAFormatType = "LBAFormat2"
	// LBAFormat3LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat3SFLBAFormatType SFLBAFormatType = "LBAFormat3"
	// LBAFormat4LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat4SFLBAFormatType SFLBAFormatType = "LBAFormat4"
	// LBAFormat5LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat5SFLBAFormatType SFLBAFormatType = "LBAFormat5"
	// LBAFormat6LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat6SFLBAFormatType SFLBAFormatType = "LBAFormat6"
	// LBAFormat7LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat7SFLBAFormatType SFLBAFormatType = "LBAFormat7"
	// LBAFormat8LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat8SFLBAFormatType SFLBAFormatType = "LBAFormat8"
	// LBAFormat9LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat9SFLBAFormatType SFLBAFormatType = "LBAFormat9"
	// LBAFormat10LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat10SFLBAFormatType SFLBAFormatType = "LBAFormat10"
	// LBAFormat11LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat11SFLBAFormatType SFLBAFormatType = "LBAFormat11"
	// LBAFormat12LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat12SFLBAFormatType SFLBAFormatType = "LBAFormat12"
	// LBAFormat13LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat13SFLBAFormatType SFLBAFormatType = "LBAFormat13"
	// LBAFormat14LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat14SFLBAFormatType SFLBAFormatType = "LBAFormat14"
	// LBAFormat15LBAFormatType Indicates the LBA data size supported if supported.
	LBAFormat15SFLBAFormatType SFLBAFormatType = "LBAFormat15"
)

// LBARelativePerformanceType is Indicate the relative performance of the LBA
// format relative to other LBA formats supported by the controller.
type SFLBARelativePerformanceType string

const (
	// BestLBARelativePerformanceType Best performance.
	BestSFLBARelativePerformanceType SFLBARelativePerformanceType = "Best"
	// BetterLBARelativePerformanceType Better performance.
	BetterSFLBARelativePerformanceType SFLBARelativePerformanceType = "Better"
	// GoodLBARelativePerformanceType Good performance.
	GoodSFLBARelativePerformanceType SFLBARelativePerformanceType = "Good"
	// DegradedLBARelativePerformanceType Degraded performance.
	DegradedSFLBARelativePerformanceType SFLBARelativePerformanceType = "Degraded"
)

type SFNamespaceType string

const (
	// BlockNamespaceType The namespace is configured for use with a block storage
	// interface.
	BlockSFNamespaceType SFNamespaceType = "Block"
	// KeyValueNamespaceType The namespace is configured for use with a KeyValue
	// interface.
	KeyValueSFNamespaceType SFNamespaceType = "KeyValue"
	// ZNSNamespaceType The namespace is configured for use with a zoned storage
	// interface.
	ZNSSFNamespaceType SFNamespaceType = "ZNS"
	// ComputationalNamespaceType The namespace is configured for use with a
	// computational storage interface.
	ComputationalSFNamespaceType SFNamespaceType = "Computational"
)

type SFOperationType string

const (
	// DeduplicateOperationType is a Deduplicate operation is being performed.
	DeduplicateSFOperationType SFOperationType = "Deduplicate"
	// CheckConsistencyOperationType is a CheckConsistency operation is being
	// performed.
	CheckConsistencySFOperationType SFOperationType = "CheckConsistency"
	// InitializeOperationType is an Initialize operation is being performed.
	InitializeSFOperationType SFOperationType = "Initialize"
	// ReplicateOperationType is a Replicate operation is being performed.
	ReplicateSFOperationType SFOperationType = "Replicate"
	// DeleteOperationType is a Delete operation is being performed.
	DeleteSFOperationType SFOperationType = "Delete"
	// ChangeRAIDTypeOperationType is a ChangeRAIDType operation is being
	// performed.
	ChangeRAIDTypeSFOperationType SFOperationType = "ChangeRAIDType"
	// RebuildOperationType is a Rebuild operation is being performed.
	RebuildSFOperationType SFOperationType = "Rebuild"
	// EncryptOperationType is an Encrypt operation is being performed.
	EncryptSFOperationType SFOperationType = "Encrypt"
	// DecryptOperationType is a Decrypt operation is being performed.
	DecryptSFOperationType SFOperationType = "Decrypt"
	// ResizeOperationType is a Resize operation is being performed.
	ResizeSFOperationType SFOperationType = "Resize"
	// CompressOperationType is a Compress operation is being performed.
	CompressSFOperationType SFOperationType = "Compress"
	// SanitizeOperationType is a Sanitize operation is being performed.
	SanitizeSFOperationType SFOperationType = "Sanitize"
	// FormatOperationType is a Format operation is being performed.
	FormatSFOperationType SFOperationType = "Format"
	// ChangeStripSizeOperationType is a ChangeStripSize operation is being
	// performed.
	ChangeStripSizeSFOperationType SFOperationType = "ChangeStripSize"
)

type SFRAIDType string

const (
	// RAID0RAIDType is a placement policy where consecutive logical blocks of data
	// are uniformly distributed across a set of independent storage devices
	// without offering any form of redundancy. This is y referred to as
	// data striping. This form of RAID will encounter data loss with the failure
	// of any storage device in the set.
	RAID0SFRAIDType SFRAIDType = "RAID0"
	// RAID1RAIDType is a placement policy where each logical block of data is
	// stored on more than one independent storage device. This is y
	// referred to as mirroring. Data stored using this form of RAID is able to
	// survive a single storage device failure without data loss.
	RAID1SFRAIDType SFRAIDType = "RAID1"
	// RAID3RAIDType is a placement policy using parity-based protection where
	// logical bytes of data are uniformly distributed across a set of independent
	// storage devices and where the parity is stored on a dedicated independent
	// storage device. Data stored using this form of RAID is able to survive a
	// single storage device failure without data loss. If the storage devices use
	// rotating media, they are assumed to be rotationally synchronized, and the
	// data stripe size should be no larger than the exported block size.
	RAID3SFRAIDType SFRAIDType = "RAID3"
	// RAID4RAIDType is a placement policy using parity-based protection where
	// logical blocks of data are uniformly distributed across a set of independent
	// storage devices and where the parity is stored on a dedicated independent
	// storage device. Data stored using this form of RAID is able to survive a
	// single storage device failure without data loss.
	RAID4SFRAIDType SFRAIDType = "RAID4"
	// RAID5RAIDType is a placement policy using parity-based protection for
	// storing stripes of 'n' logical blocks of data and one logical block of
	// parity across a set of 'n+1' independent storage devices where the parity
	// and data blocks are interleaved across the storage devices. Data stored
	// using this form of RAID is able to survive a single storage device failure
	// without data loss.
	RAID5SFRAIDType SFRAIDType = "RAID5"
	// RAID6RAIDType is a placement policy using parity-based protection for
	// storing stripes of 'n' logical blocks of data and two logical blocks of
	// independent parity across a set of 'n+2' independent storage devices where
	// the parity and data blocks are interleaved across the storage devices. Data
	// stored using this form of RAID is able to survive any two independent
	// storage device failures without data loss.
	RAID6SFRAIDType SFRAIDType = "RAID6"
	// RAID10RAIDType is a placement policy that creates a striped device (RAID 0)
	// over a set of mirrored devices (RAID 1). This is y referred to as
	// RAID 1/0. Data stored using this form of RAID is able to survive storage
	// device failures in each RAID 1 set without data loss.
	RAID10SFRAIDType SFRAIDType = "RAID10"
	// RAID01RAIDType is a data placement policy that creates a mirrored device
	// (RAID 1) over a set of striped devices (RAID 0). This is y referred
	// to as RAID 0+1 or RAID 0/1. Data stored using this form of RAID is able to
	// survive a single RAID 0 data set failure without data loss.
	RAID01SFRAIDType SFRAIDType = "RAID01"
	// RAID6TPRAIDType is a placement policy that uses parity-based protection for
	// storing stripes of 'n' logical blocks of data and three logical blocks of
	// independent parity across a set of 'n+3' independent storage devices where
	// the parity and data blocks are interleaved across the storage devices. This
	// is y referred to as Triple Parity RAID. Data stored using this form
	// of RAID is able to survive any three independent storage device failures
	// without data loss.
	RAID6TPSFRAIDType SFRAIDType = "RAID6TP"
	// RAID1ERAIDType is a placement policy that uses a form of mirroring
	// implemented over a set of independent storage devices where logical blocks
	// are duplicated on a pair of independent storage devices so that data is
	// uniformly distributed across the storage devices. This is y referred
	// to as RAID 1 Enhanced. Data stored using this form of RAID is able to
	// survive a single storage device failure without data loss.
	RAID1ESFRAIDType SFRAIDType = "RAID1E"
	// RAID50RAIDType is a placement policy that uses a RAID 0 stripe set over two
	// or more RAID 5 sets of independent storage devices. Data stored using this
	// form of RAID is able to survive a single storage device failure within each
	// RAID 5 set without data loss.
	RAID50SFRAIDType SFRAIDType = "RAID50"
	// RAID60RAIDType is a placement policy that uses a RAID 0 stripe set over two
	// or more RAID 6 sets of independent storage devices. Data stored using this
	// form of RAID is able to survive two device failures within each RAID 6 set
	// without data loss.
	RAID60SFRAIDType SFRAIDType = "RAID60"
	// RAID00RAIDType is a placement policy that creates a RAID 0 stripe set over
	// two or more RAID 0 sets. This is y referred to as RAID 0+0. This form
	// of data layout is not fault tolerant; if any storage device fails there will
	// be data loss.
	RAID00SFRAIDType SFRAIDType = "RAID00"
	// RAID10ERAIDType is a placement policy that uses a RAID 0 stripe set over two
	// or more RAID 10 sets. This is y referred to as Enhanced RAID 10. Data
	// stored using this form of RAID is able to survive a single device failure
	// within each nested RAID 1 set without data loss.
	RAID10ESFRAIDType SFRAIDType = "RAID10E"
	// RAID1TripleRAIDType is a placement policy where each logical block of data
	// is mirrored three times across a set of three independent storage devices.
	// This is y referred to as three-way mirroring. This form of RAID can
	// survive two device failures without data loss.
	RAID1TripleSFRAIDType SFRAIDType = "RAID1Triple"
	// RAID10TripleRAIDType is a placement policy that uses a striped device (RAID
	// 0) over a set of triple mirrored devices (RAID 1Triple). This form of RAID
	// can survive up to two failures in each triple mirror set without data loss.
	RAID10TripleSFRAIDType SFRAIDType = "RAID10Triple"
	// NoneRAIDType is a placement policy with no redundancy at the device level.
	NoneSFRAIDType SFRAIDType = "None"
)

type SFReadCachePolicyType string

const (
	// ReadAheadReadCachePolicyType is a caching technique in which the controller
	// pre-fetches data anticipating future read requests.
	ReadAheadSFReadCachePolicyType SFReadCachePolicyType = "ReadAhead"
	// AdaptiveReadAheadReadCachePolicyType is a caching technique in which the
	// controller dynamically determines whether to pre-fetch data anticipating
	// future read requests, based on previous cache hit ratio.
	AdaptiveReadAheadSFReadCachePolicyType SFReadCachePolicyType = "AdaptiveReadAhead"
	// OffReadCachePolicyType The read cache is disabled.
	OffSFReadCachePolicyType SFReadCachePolicyType = "Off"
)

type SFVolumeType string

const (
	// RawDeviceVolumeType The volume is a raw physical device without any RAID or
	// other virtualization applied.
	RawDeviceSFVolumeType SFVolumeType = "RawDevice"
	// NonRedundantVolumeType The volume is a non-redundant storage device.
	NonRedundantSFVolumeType SFVolumeType = "NonRedundant"
	// MirroredVolumeType The volume is a mirrored device.
	MirroredSFVolumeType SFVolumeType = "Mirrored"
	// StripedWithParityVolumeType The volume is a device which uses parity to
	// retain redundant information.
	StripedWithParitySFVolumeType SFVolumeType = "StripedWithParity"
	// SpannedMirrorsVolumeType The volume is a spanned set of mirrored devices.
	SpannedMirrorsSFVolumeType SFVolumeType = "SpannedMirrors"
	// SpannedStripesWithParityVolumeType The volume is a spanned set of devices
	// which uses parity to retain redundant information.
	SpannedStripesWithParitySFVolumeType SFVolumeType = "SpannedStripesWithParity"
)

type SFVolumeUsageType string

const (
	// DataVolumeUsageType shall be allocated for use as a consumable data volume.
	DataSFVolumeUsageType SFVolumeUsageType = "Data"
	// SystemDataVolumeUsageType shall be allocated for use as a consumable data
	// volume reserved for system use.
	SystemDataSFVolumeUsageType SFVolumeUsageType = "SystemData"
	// CacheOnlyVolumeUsageType shall be allocated for use as a non-consumable
	// cache only volume.
	CacheOnlySFVolumeUsageType SFVolumeUsageType = "CacheOnly"
	// SystemReserveVolumeUsageType shall be allocated for use as a non-consumable
	// system reserved volume.
	SystemReserveSFVolumeUsageType SFVolumeUsageType = "SystemReserve"
	// ReplicationReserveVolumeUsageType shall be allocated for use as a
	// non-consumable reserved volume for replication use.
	ReplicationReserveSFVolumeUsageType SFVolumeUsageType = "ReplicationReserve"
)

type SFWriteCachePolicyType string

const (
	// WriteThroughWriteCachePolicyType is a caching technique in which the
	// completion of a write request is not signaled until data is safely stored on
	// non-volatile media.
	WriteThroughSFWriteCachePolicyType SFWriteCachePolicyType = "WriteThrough"
	// ProtectedWriteBackWriteCachePolicyType is a caching technique in which the
	// completion of a write request is signaled as soon as the data is in cache,
	// and actual writing to non-volatile media is guaranteed to occur at a later
	// time.
	ProtectedWriteBackSFWriteCachePolicyType SFWriteCachePolicyType = "ProtectedWriteBack"
	// UnprotectedWriteBackWriteCachePolicyType is a caching technique in which the
	// completion of a write request is signaled as soon as the data is in cache;
	// actual writing to non-volatile media is not guaranteed to occur at a later
	// time.
	UnprotectedWriteBackSFWriteCachePolicyType SFWriteCachePolicyType = "UnprotectedWriteBack"
	// OffWriteCachePolicyType shall be disabled.
	OffSFWriteCachePolicyType SFWriteCachePolicyType = "Off"
)

type SFWriteCacheStateType string

const (
	// UnprotectedWriteCacheStateType Indicates that the cache state type in use
	// generally does not protect write requests on non-volatile media.
	UnprotectedSFWriteCacheStateType SFWriteCacheStateType = "Unprotected"
	// ProtectedWriteCacheStateType Indicates that the cache state type in use
	// generally protects write requests on non-volatile media.
	ProtectedSFWriteCacheStateType SFWriteCacheStateType = "Protected"
	// DegradedWriteCacheStateType Indicates an issue with the cache state in which
	// the cache space is diminished or disabled due to a failure or an outside
	// influence such as a discharged battery.
	DegradedSFWriteCacheStateType SFWriteCacheStateType = "Degraded"
)

type SFWriteHoleProtectionPolicyType string

const (
	// OffWriteHoleProtectionPolicyType The support for addressing the write hole
	// issue is disabled. The volume is not performing any additional activities to
	// close the RAID write hole.
	OffSFWriteHoleProtectionPolicyType SFWriteHoleProtectionPolicyType = "Off"
	// JournalingWriteHoleProtectionPolicyType The policy that uses separate block
	// device for write-ahead logging to address write hole issue. All write
	// operations on the RAID volume are first logged on dedicated journaling
	// device that is not part of the volume.
	JournalingSFWriteHoleProtectionPolicyType SFWriteHoleProtectionPolicyType = "Journaling"
	// DistributedLogWriteHoleProtectionPolicyType The policy that distributes
	// additional log (e.q. checksum of the parity) among the volume's capacity
	// sources to address write hole issue. Additional data is used to detect data
	// corruption on the volume.
	DistributedLogSFWriteHoleProtectionPolicyType SFWriteHoleProtectionPolicyType = "DistributedLog"
	// OemWriteHoleProtectionPolicyType The policy that is Oem specific. The
	// mechanism details are unknown unless provided separately by the Oem.
	OemSFWriteHoleProtectionPolicyType SFWriteHoleProtectionPolicyType = "Oem"
)

// Volume shall be used to represent a volume, virtual disk, logical disk, LUN,
// or other logical storage for a Redfish implementation.
//
//nolint:dupl
type SFVolume struct {
	Entity
	// ALUA shall identify the ALUA properties for this volume.
	//
	// Version added: v1.10.0
	ALUA ALUA
	// AccessCapabilities shall specify a current storage access capability.
	//
	// Version added: v1.1.0
	AccessCapabilities []StorageAccessCapability
	// AllocatedPools shall contain references to all storage pools allocated from
	// this volume.
	//
	// Version added: v1.1.0
	allocatedPools string
	// BlockSizeBytes shall contain size of the smallest addressable unit of the
	// associated volume.
	BlockSizeBytes *int `json:",omitempty"`
	// Capacity Information about the utilization of capacity allocated to this
	// storage volume.
	//
	// Version added: v1.1.0
	Capacity Capacity
	// CapacityBytes shall contain the size in bytes of the associated volume.
	CapacityBytes *int `json:",omitempty"`
	// CapacitySources Fully or partially consumed storage from a source resource.
	// Each entry provides capacity allocation information from a named source
	// resource.
	//
	// Version added: v1.1.0
	CapacitySources []CapacitySource
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
	connections []string
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
	metrics string
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
	ProvisioningPolicy ProvisioningPolicy
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

// UnmarshalJSON unmarshals a SFVolume object from the raw JSON.
func (s *SFVolume) UnmarshalJSON(b []byte) error {
	type temp SFVolume
	type sActions struct {
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
	type sLinks struct {
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
		Actions        sActions
		Links          sLinks
		AllocatedPools Link  `json:"AllocatedPools"`
		Connections    Links `json:"Connections"`
		Metrics        Link  `json:"Metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SFVolume(tmp.temp)

	// Extract the links to other entities for later
	s.assignReplicaTargetTarget = tmp.Actions.AssignReplicaTarget.Target
	s.changeRAIDLayoutTarget = tmp.Actions.ChangeRAIDLayout.Target
	s.checkConsistencyTarget = tmp.Actions.CheckConsistency.Target
	s.createReplicaTargetTarget = tmp.Actions.CreateReplicaTarget.Target
	s.forceEnableTarget = tmp.Actions.ForceEnable.Target
	s.initializeTarget = tmp.Actions.Initialize.Target
	s.removeReplicaRelationshipTarget = tmp.Actions.RemoveReplicaRelationship.Target
	s.resumeReplicationTarget = tmp.Actions.ResumeReplication.Target
	s.reverseReplicationRelationshipTarget = tmp.Actions.ReverseReplicationRelationship.Target
	s.splitReplicationTarget = tmp.Actions.SplitReplication.Target
	s.suspendReplicationTarget = tmp.Actions.SuspendReplication.Target
	s.cacheDataVolumes = tmp.Links.CacheDataVolumes.ToStrings()
	s.cacheVolumeSource = tmp.Links.CacheVolumeSource.String()
	s.classOfService = tmp.Links.ClassOfService.String()
	s.clientEndpoints = tmp.Links.ClientEndpoints.ToStrings()
	s.consistencyGroups = tmp.Links.ConsistencyGroups.ToStrings()
	s.controllers = tmp.Links.Controllers.ToStrings()
	s.dedicatedSpareDrives = tmp.Links.DedicatedSpareDrives.ToStrings()
	s.drives = tmp.Links.Drives.ToStrings()
	s.journalingMedia = tmp.Links.JournalingMedia.String()
	s.owningStorageResource = tmp.Links.OwningStorageResource.String()
	s.owningStorageService = tmp.Links.OwningStorageService.String()
	s.providingStoragePool = tmp.Links.ProvidingStoragePool.String()
	s.serverEndpoints = tmp.Links.ServerEndpoints.ToStrings()
	s.spareResourceSets = tmp.Links.SpareResourceSets.ToStrings()
	s.storageGroups = tmp.Links.StorageGroups.ToStrings()
	s.allocatedPools = tmp.AllocatedPools.String()
	s.connections = tmp.Connections.ToStrings()
	s.metrics = tmp.Metrics.String()

	// This is a read/write object, so we need to save the raw object data for later
	s.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *SFVolume) Update() error {
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

	return s.UpdateFromRawData(s, s.rawData, readWriteFields)
}

// GetSFVolume will get a SFVolume instance from the service.
func GetSFVolume(c Client, uri string) (*SFVolume, error) {
	return GetObject[SFVolume](c, uri)
}

// ListReferencedSFVolumes gets the collection of SFVolume from
// a provided reference.
func ListReferencedSFVolumes(c Client, link string) ([]*SFVolume, error) {
	return GetCollectionObjects[SFVolume](c, link)
}

// This action shall be used to establish a replication relationship by
// assigning an existing volume to serve as a target replica for an existing
// source volume.
// replicaType - This parameter shall contain the type of replica relationship
// to be created (e.g., Clone, Mirror, Snap).
// replicaUpdateMode - This parameter shall specify the replica update mode.
// targetVolume - This parameter shall contain the Uri to the existing target
// volume.
func (s *SFVolume) AssignReplicaTarget(replicaType ReplicaType, replicaUpdateMode ReplicaUpdateMode, targetVolume string) error {
	payload := make(map[string]any)
	payload["ReplicaType"] = replicaType
	payload["ReplicaUpdateMode"] = replicaUpdateMode
	payload["TargetVolume"] = targetVolume
	return s.Post(s.assignReplicaTargetTarget, payload)
}

// This action shall request the system to change the RAID layout of the
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
func (s *SFVolume) ChangeRAIDLayout(drives []string, mediaSpanCount int, rAIDType RAIDType, stripSizeBytes int) error {
	payload := make(map[string]any)
	payload["Drives"] = drives
	payload["MediaSpanCount"] = mediaSpanCount
	payload["RAIDType"] = rAIDType
	payload["StripSizeBytes"] = stripSizeBytes
	return s.Post(s.changeRAIDLayoutTarget, payload)
}

// This defines the name of the custom action supported on this resource.
func (s *SFVolume) CheckConsistency() error {
	payload := make(map[string]any)
	return s.Post(s.checkConsistencyTarget, payload)
}

// This action shall be used to create a new volume resource to provide
// expanded data protection through a replica relationship with the specified
// source volume.
// replicaType - This parameter shall contain the type of replica relationship
// to be created (e.g., Clone, Mirror, Snap).
// replicaUpdateMode - This parameter shall specify the replica update mode.
// targetStoragePool - This parameter shall contain the Uri to the existing
// StoragePool in which to create the target volume.
// volumeName - This parameter shall contain the Name for the target volume.
func (s *SFVolume) CreateReplicaTarget(replicaType ReplicaType, replicaUpdateMode ReplicaUpdateMode, targetStoragePool string, volumeName string) error {
	payload := make(map[string]any)
	payload["ReplicaType"] = replicaType
	payload["ReplicaUpdateMode"] = replicaUpdateMode
	payload["TargetStoragePool"] = targetStoragePool
	payload["VolumeName"] = volumeName
	return s.Post(s.createReplicaTargetTarget, payload)
}

// This action shall request the system to force the volume to enabled state
// regardless of data loss scenarios.
func (s *SFVolume) ForceEnable() error {
	payload := make(map[string]any)
	return s.Post(s.forceEnableTarget, payload)
}

// This defines the name of the custom action supported on this resource. If
// InitializeMethod is not specified in the request body, but the property
// InitializeMethod is specified, the property InitializeMethod value should be
// used. If neither is specified, the InitializeMethod should be Foreground.
// initializeMethod - This defines the property name for the action.
// initializeType - This defines the property name for the action.
func (s *SFVolume) Initialize(initializeMethod InitializeMethod, initializeType InitializeType) error {
	payload := make(map[string]any)
	payload["InitializeMethod"] = initializeMethod
	payload["InitializeType"] = initializeType
	return s.Post(s.initializeTarget, payload)
}

// This action shall be used to disable data synchronization between a source
// and target volume, remove the replication relationship, and optionally
// delete the target volume.
// deleteTargetVolume - This parameter shall indicate whether or not to delete
// the target volume as part of the operation. If not defined, the system
// should use its default behavior.
// targetVolume - This parameter shall contain the Uri to the existing target
// volume.
func (s *SFVolume) RemoveReplicaRelationship(deleteTargetVolume bool, targetVolume string) error {
	payload := make(map[string]any)
	payload["DeleteTargetVolume"] = deleteTargetVolume
	payload["TargetVolume"] = targetVolume
	return s.Post(s.removeReplicaRelationshipTarget, payload)
}

// This action shall be used to resume the active data synchronization between
// a source and target volume, without otherwise altering the replication
// relationship.
// targetVolume - This parameter shall contain the Uri to the existing target
// volume.
func (s *SFVolume) ResumeReplication(targetVolume string) error {
	payload := make(map[string]any)
	payload["TargetVolume"] = targetVolume
	return s.Post(s.resumeReplicationTarget, payload)
}

// This action shall be used to reverse the replication relationship between a
// source and target volume.
// targetVolume - This parameter shall contain the Uri to the existing target
// volume.
func (s *SFVolume) ReverseReplicationRelationship(targetVolume string) error {
	payload := make(map[string]any)
	payload["TargetVolume"] = targetVolume
	return s.Post(s.reverseReplicationRelationshipTarget, payload)
}

// This action shall be used to split the replication relationship and suspend
// data synchronization between a source and target volume.
// targetVolume - This parameter shall contain the Uri to the existing target
// volume.
func (s *SFVolume) SplitReplication(targetVolume string) error {
	payload := make(map[string]any)
	payload["TargetVolume"] = targetVolume
	return s.Post(s.splitReplicationTarget, payload)
}

// This action shall be used to suspend active data synchronization between a
// source and target volume, without otherwise altering the replication
// relationship.
// targetVolume - This parameter shall contain the Uri to the existing target
// volume.
func (s *SFVolume) SuspendReplication(targetVolume string) error {
	payload := make(map[string]any)
	payload["TargetVolume"] = targetVolume
	return s.Post(s.suspendReplicationTarget, payload)
}

// CacheDataVolumes gets the CacheDataVolumes linked resources.
func (s *SFVolume) CacheDataVolumes(client Client) ([]*Volume, error) {
	return GetObjects[Volume](client, s.cacheDataVolumes)
}

// CacheVolumeSource gets the CacheVolumeSource linked resource.
func (s *SFVolume) CacheVolumeSource(client Client) (*SFVolume, error) {
	if s.cacheVolumeSource == "" {
		return nil, nil
	}
	return GetObject[SFVolume](client, s.cacheVolumeSource)
}

// ClassOfService gets the ClassOfService linked resource.
func (s *SFVolume) ClassOfService(client Client) (*ClassOfService, error) {
	if s.classOfService == "" {
		return nil, nil
	}
	return GetObject[ClassOfService](client, s.classOfService)
}

// ClientEndpoints gets the ClientEndpoints linked resources.
func (s *SFVolume) ClientEndpoints(client Client) ([]*Endpoint, error) {
	return GetObjects[Endpoint](client, s.clientEndpoints)
}

// ConsistencyGroups gets the ConsistencyGroups linked resources.
func (s *SFVolume) ConsistencyGroups(client Client) ([]*ConsistencyGroup, error) {
	return GetObjects[ConsistencyGroup](client, s.consistencyGroups)
}

// Controllers gets the Controllers linked resources.
func (s *SFVolume) Controllers(client Client) ([]*StorageController, error) {
	return GetObjects[StorageController](client, s.controllers)
}

// DedicatedSpareDrives gets the DedicatedSpareDrives linked resources.
func (s *SFVolume) DedicatedSpareDrives(client Client) ([]*Drive, error) {
	return GetObjects[Drive](client, s.dedicatedSpareDrives)
}

// Drives gets the Drives linked resources.
func (s *SFVolume) Drives(client Client) ([]*Drive, error) {
	return GetObjects[Drive](client, s.drives)
}

// JournalingMedia gets the JournalingMedia linked resource.
func (s *SFVolume) JournalingMedia(client Client) (*Resource, error) {
	if s.journalingMedia == "" {
		return nil, nil
	}
	return GetObject[Resource](client, s.journalingMedia)
}

// OwningStorageResource gets the OwningStorageResource linked resource.
func (s *SFVolume) OwningStorageResource(client Client) (*Storage, error) {
	if s.owningStorageResource == "" {
		return nil, nil
	}
	return GetObject[Storage](client, s.owningStorageResource)
}

// OwningStorageService gets the OwningStorageService linked resource.
func (s *SFVolume) OwningStorageService(client Client) (*StorageService, error) {
	if s.owningStorageService == "" {
		return nil, nil
	}
	return GetObject[StorageService](client, s.owningStorageService)
}

// ProvidingStoragePool gets the ProvidingStoragePool linked resource.
func (s *SFVolume) ProvidingStoragePool(client Client) (*StoragePool, error) {
	if s.providingStoragePool == "" {
		return nil, nil
	}
	return GetObject[StoragePool](client, s.providingStoragePool)
}

// ServerEndpoints gets the ServerEndpoints linked resources.
func (s *SFVolume) ServerEndpoints(client Client) ([]*Endpoint, error) {
	return GetObjects[Endpoint](client, s.serverEndpoints)
}

// SpareResourceSets gets the SpareResourceSets linked resources.
func (s *SFVolume) SpareResourceSets(client Client) ([]*SpareResourceSet, error) {
	return GetObjects[SpareResourceSet](client, s.spareResourceSets)
}

// StorageGroups gets the StorageGroups linked resources.
func (s *SFVolume) StorageGroups(client Client) ([]*StorageGroup, error) {
	return GetObjects[StorageGroup](client, s.storageGroups)
}

// AllocatedPools gets the AllocatedPools collection.
func (s *SFVolume) AllocatedPools(client Client) ([]*StoragePool, error) {
	if s.allocatedPools == "" {
		return nil, nil
	}
	return GetCollectionObjects[StoragePool](client, s.allocatedPools)
}

// Connections gets the Connections linked resources.
func (s *SFVolume) Connections(client Client) ([]*Connection, error) {
	return GetObjects[Connection](client, s.connections)
}

// Metrics gets the Metrics linked resource.
func (s *SFVolume) Metrics(client Client) (*VolumeMetrics, error) {
	if s.metrics == "" {
		return nil, nil
	}
	return GetObject[VolumeMetrics](client, s.metrics)
}

// ALUA represents the ALUA type.
type SFALUA struct {
	// ANAGroupId shall contain the ANA group id for this volume.
	//
	// Version added: v1.10.0
	ANAGroupID *float64 `json:"ANAGroupId,omitempty"`
}

// LBAFormat represents the LBAFormat type.
type SFLBAFormat struct {
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
type SFNVMeNamespaceProperties struct {
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
type SFNamespaceFeatures struct {
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
type SFOperation struct {
	// AssociatedFeaturesRegistry is a reference to the task associated with the
	// operation if any.
	AssociatedFeaturesRegistry FeaturesRegistry
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
