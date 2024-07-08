//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

type ALUA struct {
	// ANAGroupID shall contain the ANA group id for this volume.
	ANAGroupID int64
}

type LBAFormat struct {
	// LBADataSizeBytes shall be the LBA data size reported in bytes.
	LBADataSizeBytes int
	// LBAFormatType shall be the LBA format type. This property is intended for capabilities instrumentation.
	LBAFormatType LBAFormatType
	// LBAMetadataSizeBytes shall be the LBA metadata size reported in bytes.
	LBAMetadataSizeBytes int
	// RelativePerformance shall be the LBA Relative Performance type. This field indicates the relative performance of
	// the LBA format indicated relative to other LBA formats supported by the controller. This property is intended
	// for capabilities instrumentation.
	RelativePerformance LBARelativePerformanceType
}

type LBAFormatType string

const (
	// LBAFormat0LBAFormatType indicates the LBA data size supported.
	LBAFormat0LBAFormatType LBAFormatType = "LBAFormat0"
	// LBAFormat1LBAFormatType indicates the LBA data size supported.
	LBAFormat1LBAFormatType LBAFormatType = "LBAFormat1"
	// LBAFormat10LBAFormatType indicates the LBA data size supported.
	LBAFormat10LBAFormatType LBAFormatType = "LBAFormat10"
	// LBAFormat11LBAFormatType indicates the LBA data size supported.
	LBAFormat11LBAFormatType LBAFormatType = "LBAFormat11"
	// LBAFormat12LBAFormatType indicates the LBA data size supported.
	LBAFormat12LBAFormatType LBAFormatType = "LBAFormat12"
	// LBAFormat13LBAFormatType indicates the LBA data size supported.
	LBAFormat13LBAFormatType LBAFormatType = "LBAFormat13"
	// LBAFormat14LBAFormatType indicates the LBA data size supported.
	LBAFormat14LBAFormatType LBAFormatType = "LBAFormat14"
	// LBAFormat15LBAFormatType indicates the LBA data size supported.
	LBAFormat15LBAFormatType LBAFormatType = "LBAFormat15"
	// LBAFormat2LBAFormatType indicates the LBA data size supported.
	LBAFormat2LBAFormatType LBAFormatType = "LBAFormat2"
	// LBAFormat3LBAFormatType indicates the LBA data size supported.
	LBAFormat3LBAFormatType LBAFormatType = "LBAFormat3"
	// LBAFormat4LBAFormatType indicates the LBA data size supported.
	LBAFormat4LBAFormatType LBAFormatType = "LBAFormat4"
	// LBAFormat5LBAFormatType indicates the LBA data size supported.
	LBAFormat5LBAFormatType LBAFormatType = "LBAFormat5"
	// LBAFormat6LBAFormatType indicates the LBA data size supported.
	LBAFormat6LBAFormatType LBAFormatType = "LBAFormat6"
	// LBAFormat7LBAFormatType indicates the LBA data size supported.
	LBAFormat7LBAFormatType LBAFormatType = "LBAFormat7"
	// LBAFormat8LBAFormatType indicates the LBA data size supported.
	LBAFormat8LBAFormatType LBAFormatType = "LBAFormat8"
	// LBAFormat9LBAFormatType indicates the LBA data size supported.
	LBAFormat9LBAFormatType LBAFormatType = "LBAFormat9"
)

type LBARelativePerformanceType string

const (
	// BestLBARelativePerformanceType indicates the best performance.
	BestLBARelativePerformanceType LBARelativePerformanceType = "Best"
	// BetterLBARelativePerformanceType indicates the bestbetter performance.
	BetterLBARelativePerformanceType LBARelativePerformanceType = "Better"
	// DegradedLBARelativePerformanceType indicates degraded performance.
	DegradedLBARelativePerformanceType LBARelativePerformanceType = "Degraded"
	// GoodLBARelativePerformanceType indicates good performance.
	GoodLBARelativePerformanceType LBARelativePerformanceType = "Good"
)

// InitializeType is
type InitializeType string

const (

	// FastInitializeType The volume is prepared for use quickly, typically
	// by erasing just the beginning and end of the space so that
	// partitioning can be performed.
	FastInitializeType InitializeType = "Fast"
	// SlowInitializeType The volume is prepared for use slowly, typically by
	// completely erasing the volume.
	SlowInitializeType InitializeType = "Slow"
)

// RAIDType is
type RAIDType string

const (

	// RAID0RAIDType A placement policy where consecutive logical blocks of
	// data are uniformly distributed across a set of independent storage
	// devices without offering any form of redundancy. This is commonly
	// referred to as data striping. This form of RAID will encounter data
	// loss with the failure of any storage device in the set.
	RAID0RAIDType RAIDType = "RAID0"
	// RAID1RAIDType A placement policy where each logical block of data is
	// stored on more than one independent storage device. This is commonly
	// referred to as mirroring. Data stored using this form of RAID is able
	// to survive a single storage device failure without data loss.
	RAID1RAIDType RAIDType = "RAID1"
	// RAID3RAIDType A placement policy using parity-based protection where
	// logical bytes of data are uniformly distributed across a set of
	// independent storage devices and where the parity is stored on a
	// dedicated independent storage device. Data stored using this form of
	// RAID is able to survive a single storage device failure without data
	// loss. If the storage devices use rotating media, they are assumed to
	// be rotationally synchronized, and the data stripe size should be no
	// larger than the exported block size.
	RAID3RAIDType RAIDType = "RAID3"
	// RAID4RAIDType A placement policy using parity-based protection where
	// logical blocks of data are uniformly distributed across a set of
	// independent storage devices and where the parity is stored on a
	// dedicated independent storage device. Data stored using this form of
	// RAID is able to survive a single storage device failure without data
	// loss.
	RAID4RAIDType RAIDType = "RAID4"
	// RAID5RAIDType A placement policy using parity-based protection for
	// storing stripes of 'n' logical blocks of data and one logical block of
	// parity across a set of 'n+1' independent storage devices where the
	// parity and data blocks are interleaved across the storage devices.
	// Data stored using this form of RAID is able to survive a single
	// storage device failure without data loss.
	RAID5RAIDType RAIDType = "RAID5"
	// RAID6RAIDType A placement policy using parity-based protection for
	// storing stripes of 'n' logical blocks of data and two logical blocks
	// of independent parity across a set of 'n+2' independent storage
	// devices where the parity and data blocks are interleaved across the
	// storage devices. Data stored using this form of RAID is able to
	// survive any two independent storage device failures without data loss.
	RAID6RAIDType RAIDType = "RAID6"
	// RAID10RAIDType A placement policy that creates a striped device (RAID
	// 0) over a set of mirrored devices (RAID 1). This is commonly referred
	// to as RAID 1/0. Data stored using this form of RAID is able to survive
	// storage device failures in each RAID 1 set without data loss.
	RAID10RAIDType RAIDType = "RAID10"
	// RAID01RAIDType A data placement policy that creates a mirrored device
	// (RAID 1) over a set of striped devices (RAID 0). This is commonly
	// referred to as RAID 0+1 or RAID 0/1. Data stored using this form of
	// RAID is able to survive a single RAID 0 data set failure without data
	// loss.
	RAID01RAIDType RAIDType = "RAID01"
	// RAID6TPRAIDType A placement policy that uses parity-based protection
	// for storing stripes of 'n' logical blocks of data and three logical
	// blocks of independent parity across a set of 'n+3' independent storage
	// devices where the parity and data blocks are interleaved across the
	// storage devices. This is commonly referred to as Triple Parity RAID.
	// Data stored using this form of RAID is able to survive any three
	// independent storage device failures without data loss.
	RAID6TPRAIDType RAIDType = "RAID6TP"
	// RAID1ERAIDType A placement policy that uses a form of mirroring
	// implemented over a set of independent storage devices where logical
	// blocks are duplicated on a pair of independent storage devices so that
	// data is uniformly distributed across the storage devices. This is
	// commonly referred to as RAID 1 Enhanced. Data stored using this form
	// of RAID is able to survive a single storage device failure without
	// data loss.
	RAID1ERAIDType RAIDType = "RAID1E"
	// RAID50RAIDType A placement policy that uses a RAID 0 stripe set over
	// two or more RAID 5 sets of independent storage devices. Data stored
	// using this form of RAID is able to survive a single storage device
	// failure within each RAID 5 set without data loss.
	RAID50RAIDType RAIDType = "RAID50"
	// RAID60RAIDType A placement policy that uses a RAID 0 stripe set over
	// two or more RAID 6 sets of independent storage devices. Data stored
	// using this form of RAID is able to survive two device failures within
	// each RAID 6 set without data loss.
	RAID60RAIDType RAIDType = "RAID60"
	// RAID00RAIDType A placement policy that creates a RAID 0 stripe set
	// over two or more RAID 0 sets. This is commonly referred to as RAID
	// 0+0. This form of data layout is not fault tolerant; if any storage
	// device fails there will be data loss.
	RAID00RAIDType RAIDType = "RAID00"
	// RAID10ERAIDType A placement policy that uses a RAID 0 stripe set over
	// two or more RAID 10 sets. This is commonly referred to as Enhanced
	// RAID 10. Data stored using this form of RAID is able to survive a
	// single device failure within each nested RAID 1 set without data loss.
	RAID10ERAIDType RAIDType = "RAID10E"
	// RAID1TripleRAIDType A placement policy where each logical block of
	// data is mirrored three times across a set of three independent storage
	// devices. This is commonly referred to as three-way mirroring. This
	// form of RAID can survive two device failures without data loss.
	RAID1TripleRAIDType RAIDType = "RAID1Triple"
	// RAID10TripleRAIDType A placement policy that uses a striped device
	// (RAID 0) over a set of triple mirrored devices (RAID 1Triple). This
	// form of RAID can survive up to two failures in each triple mirror set
	// without data loss.
	RAID10TripleRAIDType RAIDType = "RAID10Triple"
)

// ReadCachePolicyType is the type of read cache policy.
type ReadCachePolicyType string

const (

	// ReadAheadReadCachePolicyType A caching technique in which the
	// controller pre-fetches data anticipating future read requests.
	ReadAheadReadCachePolicyType ReadCachePolicyType = "ReadAhead"
	// AdaptiveReadAheadReadCachePolicyType A caching technique in which the
	// controller dynamically determines whether to pre-fetch data
	// anticipating future read requests, based on previous cache hit ratio.
	AdaptiveReadAheadReadCachePolicyType ReadCachePolicyType = "AdaptiveReadAhead"
	// OffReadCachePolicyType The read cache is disabled.
	OffReadCachePolicyType ReadCachePolicyType = "Off"
)

type InitializeMethod string

const (
	// BackgroundInitializeMethod means the volume will be available for use immediately, with data erasure and preparation to happen as background tasks.
	BackgroundInitializeMethod InitializeMethod = "Background"
	// ForegroundInitializeMethod means data erasure and preparation tasks will complete before the volume is presented as available for use.
	ForegroundInitializeMethod InitializeMethod = "Foreground"
	// SkipInitializeMethod means the volume will be available for use immediately, with no preparation.
	SkipInitializeMethod InitializeMethod = "Skip"
)

// VolumeUsageType is the type of volume usage.
type VolumeUsageType string

const (

	// DataVolumeUsageType shall be allocated for use as a consumable data
	// volume.
	DataVolumeUsageType VolumeUsageType = "Data"
	// SystemDataVolumeUsageType shall be allocated for use as a consumable
	// data volume reserved for system use.
	SystemDataVolumeUsageType VolumeUsageType = "SystemData"
	// CacheOnlyVolumeUsageType shall be allocated for use as a non-
	// consumable cache only volume.
	CacheOnlyVolumeUsageType VolumeUsageType = "CacheOnly"
	// SystemReserveVolumeUsageType shall be allocated for use as a non-
	// consumable system reserved volume.
	SystemReserveVolumeUsageType VolumeUsageType = "SystemReserve"
	// ReplicationReserveVolumeUsageType shall be allocated for use as a non-
	// consumable reserved volume for replication use.
	ReplicationReserveVolumeUsageType VolumeUsageType = "ReplicationReserve"
)

// WriteCachePolicyType is the type of write cache policy.
type WriteCachePolicyType string

const (

	// WriteThroughWriteCachePolicyType A caching technique in which the
	// completion of a write request is not signaled until data is safely
	// stored on non-volatile media.
	WriteThroughWriteCachePolicyType WriteCachePolicyType = "WriteThrough"
	// ProtectedWriteBackWriteCachePolicyType A caching technique in which
	// the completion of a write request is signaled as soon as the data is
	// in cache, and actual writing to non-volatile media is guaranteed to
	// occur at a later time.
	ProtectedWriteBackWriteCachePolicyType WriteCachePolicyType = "ProtectedWriteBack"
	// UnprotectedWriteBackWriteCachePolicyType A caching technique in which
	// the completion of a write request is signaled as soon as the data is
	// in cache; actual writing to non-volatile media is not guaranteed to
	// occur at a later time.
	UnprotectedWriteBackWriteCachePolicyType WriteCachePolicyType = "UnprotectedWriteBack"
	// OffWriteCachePolicyType shall be disabled.
	OffWriteCachePolicyType WriteCachePolicyType = "Off"
)

// WriteCacheStateType is the write cache state.
type WriteCacheStateType string

const (

	// UnprotectedWriteCacheStateType Indicates that the cache state type in
	// use generally does not protect write requests on non-volatile media.
	UnprotectedWriteCacheStateType WriteCacheStateType = "Unprotected"
	// ProtectedWriteCacheStateType Indicates that the cache state type in
	// use generally protects write requests on non-volatile media.
	ProtectedWriteCacheStateType WriteCacheStateType = "Protected"
	// DegradedWriteCacheStateType Indicates an issue with the cache state in
	// which the cache space is diminished or disabled due to a failure or an
	// outside influence such as a discharged battery.
	DegradedWriteCacheStateType WriteCacheStateType = "Degraded"
)

// WriteHoleProtectionPolicyType is the write hole protection policy.
type WriteHoleProtectionPolicyType string

const (

	// OffWriteHoleProtectionPolicyType The support for addressing the write
	// hole issue is disabled. The volume is not performing any additional
	// activities to close the RAID write hole.
	OffWriteHoleProtectionPolicyType WriteHoleProtectionPolicyType = "Off"
	// JournalingWriteHoleProtectionPolicyType The policy that uses separate
	// block device for write-ahead logging to address write hole issue. All
	// write operations on the RAID volume are first logged on dedicated
	// journaling device that is not part of the volume.
	JournalingWriteHoleProtectionPolicyType WriteHoleProtectionPolicyType = "Journaling"
	// DistributedLogWriteHoleProtectionPolicyType The policy that
	// distributes additional log (e.q. checksum of the parity) among the
	// volume's capacity sources to address write hole issue. Additional data
	// is used to detect data corruption on the volume.
	DistributedLogWriteHoleProtectionPolicyType WriteHoleProtectionPolicyType = "DistributedLog"
	// OEMWriteHoleProtectionPolicyType The policy that is Oem specific. The
	// mechanism details are unknown unless provided separately by the Oem.
	OEMWriteHoleProtectionPolicyType WriteHoleProtectionPolicyType = "Oem"
)

type NamespaceType string

const (
	// BlockNamespaceType indicates the namespace is configured for use with a block storage interface.
	BlockNamespaceType NamespaceType = "Block"
	// ComputationalNamespaceType indicates the namespace is configured for use with a computational storage interface.
	ComputationalNamespaceType NamespaceType = "Computational"
	// KeyValueNamespaceType indicates the namespace is configured for use with a KeyValue interface.
	KeyValueNamespaceType NamespaceType = "KeyValue"
	// ZNSNamespaceType indicates the namespace is configured for use with a zoned storage interface.
	ZNSNamespaceType NamespaceType = "ZNS"
)

// NVMeNamespaceProperties This contains properties to use when Volume is used to describe an NVMe Namespace.
type NVMeNamespaceProperties struct {
	// FormattedLBASize shall contain the LBA data size and metadata size combination that the namespace has been
	// formatted with. This is a 4-bit data structure.
	FormattedLBASize string
	// IsShareable shall indicate whether the namespace is shareable.
	IsShareable bool
	// LBAFormat shall describe the current LBA format ID and corresponding detailed properties, such as the LBA data
	// size and metadata size. Use the LBAFormats property to describe namespace capabilities in a collection
	// capabilities annotation.
	LBAFormat LBAFormat
	// LBAFormats shall describe the LBA format IDs and corresponding detailed properties, such as the LBA data size
	// and metadata size. This property is intended for use in a collection capabilities annotation. Use the LBAFormat
	// property on an instance of a namespace.
	LBAFormats []LBAFormat
	// LBAFormatsSupported shall be a list of the LBA formats supported for the namespace, or potential namespaces.
	LBAFormatsSupported []LBAFormatType
	// MetadataTransferredAtEndOfDataLBA shall indicate whether or not the metadata is transferred at the end of the
	// LBA creating an extended data LBA.
	MetadataTransferredAtEndOfDataLBA bool
	// NVMeVersion shall contain the version of the NVMe Base Specification supported.
	NVMeVersion string
	// NamespaceFeatures shall contain a set of Namespace Features.
	NamespaceFeatures NamespaceFeatures
	// NamespaceID shall contain the NVMe Namespace Identifier for this namespace. This property shall be a hex value.
	// Namespace identifiers are not durable and do not have meaning outside the scope of the NVMe subsystem. NSID 0x0,
	// 0xFFFFFFFF, 0xFFFFFFFE are special purpose values.
	NamespaceID string
	// NamespaceType shall identify the type of namespace.
	NamespaceType NamespaceType
	// NumberLBAFormats shall contain the number of LBA data size and metadata size combinations supported by this
	// namespace. The value of this property is between 0 and 16. LBA formats with an index set beyond this value will
	// not be supported.
	NumberLBAFormats int
	// SupportsIOPerformanceHints shall indicate whether the namespace supports IO performance hints.
	SupportsIOPerformanceHints bool
	// SupportsMultipleNamespaceAttachments shall indicate whether the namespace may be attached to two or more
	// controllers.
	SupportsMultipleNamespaceAttachments bool
	// Type shall identify the type of namespace.
	Type NamespaceType
}

// NamespaceFeatures
type NamespaceFeatures struct {
	// SupportsAtomicTransactionSize shall indicate whether or not the NVM fields for Namespace preferred write
	// granularity (NPWG), write alignment (NPWA), deallocate granularity (NPDG), deallocate alignment (NPDA) and
	// optimal write size (NOWS) are defined for this namespace and should be used by the host for I/O optimization.
	SupportsAtomicTransactionSize bool
	// SupportsDeallocatedOrUnwrittenLBError shall indicate that the controller supports deallocated or unwritten
	// logical block error for this namespace.
	SupportsDeallocatedOrUnwrittenLBError bool
	// SupportsIOPerformanceHints shall indicate that the Namespace Atomic Write Unit Normal (NAWUN), Namespace Atomic
	// Write Unit Power Fail (NAWUPF), and Namespace Atomic Compare and Write Unit (NACWU) fields are defined for this
	// namespace and should be used by the host for this namespace instead of the controller-level properties AWUN,
	// AWUPF, and ACWU.
	SupportsIOPerformanceHints bool
	// SupportsNGUIDReuse shall indicate that the namespace supports the use of an NGUID (namespace globally unique
	// identifier) value.
	SupportsNGUIDReuse bool
	// SupportsThinProvisioning shall indicate whether or not the NVMe Namespace supports thin provisioning.
	// Specifically, the namespace capacity reported may be less than the namespace size.
	SupportsThinProvisioning bool
}

type Operation struct {
	// AssociatedFeaturesRegistry A reference to the task associated with the operation if any.
	AssociatedFeaturesRegistry string
	// Operation shall contain the type of the operation.
	Operation OperationType
	// PercentageComplete The percentage of the operation that has been completed.
	PercentageComplete int
}

type OperationType string

const (
	// ChangeRAIDTypeOperationType indicates a ChangeRAIDType operation is being performed.
	ChangeRAIDTypeOperationType OperationType = "ChangeRAIDType"
	// ChangeStripSizeOperationType indicates a ChangeStripSize operation is being performed.
	ChangeStripSizeOperationType OperationType = "ChangeStripSize"
	// CheckConsistencyOperationType indicates a CheckConsistency operation is being performed.
	CheckConsistencyOperationType OperationType = "CheckConsistency"
	// CompressOperationType indicates a Compress operation is being performed.
	CompressOperationType OperationType = "Compress"
	// DecryptOperationType indicates a Decrypt operation is being performed.
	DecryptOperationType OperationType = "Decrypt"
	// DeduplicateOperationType indicates a Deduplicate operation is being performed.
	DeduplicateOperationType OperationType = "Deduplicate"
	// DeleteOperationType indicates a Delete operation is being performed.
	DeleteOperationType OperationType = "Delete"
	// EncryptOperationType indicates a Encrypt operation is being performed.
	EncryptOperationType OperationType = "Encrypt"
	// FormatOperationType indicates an Format operation is being performed.
	FormatOperationType OperationType = "Format"
	// InitializeOperationType indicates a Initialize operation is being performed.
	InitializeOperationType OperationType = "Initialize"
	// RebuildOperationType indicates a Rebuild operation is being performed.
	RebuildOperationType OperationType = "Rebuild"
	// ReplicateOperationType indicates a Replicate operation is being performed.
	ReplicateOperationType OperationType = "Replicate"
	// ResizeOperationType indicates a Resize operation is being performed.
	ResizeOperationType OperationType = "Resize"
	// ChangeRAIDTypeOperationType indicates a Sanitize operation is being performed.
	SanitizeOperationType OperationType = "Sanitize"
)

// Volume is used to represent a volume, virtual disk, logical disk, LUN,
// or other logical storage for a Redfish implementation.
type Volume struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ALUA shall identify the ALUA properties for this volume.
	ALUA ALUA
	// AccessCapabilities shall specify a current storage access capability.
	AccessCapabilities []StorageAccessCapability
	// allocatedPools shall contain references to all storage pools allocated
	// from this volume.
	allocatedPools []string
	// BlockSizeBytes shall contain size of the smallest
	// addressable unit of the associated volume.
	BlockSizeBytes int
	// Capacity is Information about the utilization of capacity allocated to
	// this storage volume.
	Capacity Capacity
	// CapacityBytes shall contain the size in bytes of the
	// associated volume.
	CapacityBytes int64
	// CapacitySources is fully or partially consumed storage from a source
	// resource. Each entry provides capacity allocation information from a
	// named source resource.
	capacitySources []string
	// CapacitySourcesCount is the number of capacity sources.
	CapacitySourcesCount int `json:"CapacitySources@odata.count"`
	// Compressed shall contain a boolean indicator if the Volume is currently
	// utilizing compression or not.
	Compressed bool
	// Connections shall contain references to all Connections that include this volume.
	connections []string
	// ConnectionsCount is the number of connections.
	ConnectionsCount int `json:"Connections@odata.count"`
	// Deduplicated shall contain a boolean indicator if the Volume is currently
	// utilizing deduplication or not.
	Deduplicated bool
	// Description provides a description of this resource.
	Description string
	// DisplayName shall contain a user-configurable string to name the volume.
	DisplayName string
	// Encrypted shall contain a boolean indicator if the
	// Volume is currently utilizing encryption or not.
	Encrypted bool
	// EncryptionTypes is the type of encryption used by this Volume.
	EncryptionTypes []redfish.EncryptionTypes
	// IOPerfModeEnabled shall indicate whether IO performance mode is enabled for the volume.
	IOPerfModeEnabled bool
	// Identifiers shall contain a list of all known durable
	// names for the associated volume.
	Identifiers []common.Identifier
	// InitializeMethod shall indicate the initialization method used for this volume. If InitializeMethod is not
	// specified, the InitializeMethod should be Foreground. This value reflects the most recently used Initialization
	// Method, and may be changed using the Initialize Action.
	InitializeMethod InitializeMethod
	// IsBootCapable shall indicate whether or not the Volume contains a boot image and is capable of booting. This
	// property may be settable by an admin or client with visibility into the contents of the volume. This property
	// should only be set to true when VolumeUsage is either not specified, or when VolumeUsage is set to Data or
	// SystemData.
	IsBootCapable bool
	// LogicalUnitNumber shall contain host-visible LogicalUnitNumber assigned
	// to this Volume. This property shall only be used when in a single connect
	// configuration and no StorageGroup configuration is used.
	LogicalUnitNumber int
	// LowSpaceWarningThresholdPercents is Each time the following value is
	// less than one of the values in the array the
	// LOW_SPACE_THRESHOLD_WARNING event shall be triggered: Across all
	// CapacitySources entries, percent = (SUM(AllocatedBytes) -
	// SUM(ConsumedBytes))/SUM(AllocatedBytes).
	LowSpaceWarningThresholdPercents []int
	// Manufacturer shall contain a value that represents
	// the manufacturer or implementer of the storage volume.
	Manufacturer string
	// MaxBlockSizeBytes shall contain size of the largest
	// addressable unit of this storage volume.
	MaxBlockSizeBytes int
	// MediaSpanCount shall indicate the number of media elements used per span
	// in the secondary RAID for a hierarchical RAID type.
	MediaSpanCount int
	// Metrics shall contain a link to a resource of type VolumeMetrics that specifies the metrics for this volume. IO
	// metrics are reported in the IOStatistics property.
	metrics string
	// Model is The value is assigned by the manufacturer and shall
	// represents a specific storage volume implementation.
	Model string
	// NVMeNamespaceProperties shall contain properties to use when Volume is used to describe an NVMe Namespace.
	NVMeNamespaceProperties NVMeNamespaceProperties
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Operations shall contain a list of all operations currently
	// running on the Volume.
	Operations []common.Operations
	// OptimumIOSizeBytes shall contain the optimum IO size
	// to use when performing IO on this volume. For logical disks, this is
	// the stripe size. For physical disks, this describes the physical
	// sector size.
	OptimumIOSizeBytes int
	// ProvisioningPolicy shall specify the volume's supported storage
	// allocation policy.
	ProvisioningPolicy ProvisioningPolicy
	// RAIDType shall contain the RAID type of the associated Volume.
	RAIDType RAIDType
	// ReadCachePolicy shall contain a boolean indicator of the read cache
	// policy for the Volume.
	ReadCachePolicy ReadCachePolicyType
	// RecoverableCapacitySourceCount is the number of available
	// capacity source resources currently available in the event that an
	// equivalent capacity source resource fails.
	RecoverableCapacitySourceCount int
	// RemainingCapacityPercent is if present, this value shall return
	// {[(SUM(AllocatedBytes) - SUM(ConsumedBytes)]/SUM(AllocatedBytes)}*100
	// represented as an integer value.
	RemainingCapacityPercent int
	// RemoteReplicaTargets shall reference the URIs to the remote target replicas that are sourced by this replica.
	// Remote indicates that the replica is managed by a separate Swordfish service instance.
	RemoteReplicaTargets []string
	// ReplicaInfo shall describe the replica relationship between this storage volume and a corresponding source
	// volume.
	ReplicaInfo ReplicaInfo
	// ReplicaTargets shall reference the target replicas that are sourced by this replica.
	replicaTargets []string
	// ReplicaTargetsCount is the number of replica targets.
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// ReplicationEnabled shall indicate whether or not replication is enabled on the volume. This property shall be
	// consistent with the state reflected at the storage pool level.
	ReplicationEnabled bool
	// Status shall contain the status of the Volume.
	Status common.Status
	// StripSizeBytes is the number of consecutively addressed virtual disk
	// blocks (bytes) mapped to consecutively addressed blocks on a single
	// member extent of a disk array. Synonym for stripe depth and chunk
	// size.
	StripSizeBytes int
	// VolumeUsage shall contain the volume usage type for the Volume.
	VolumeUsage VolumeUsageType
	// WriteCachePolicy shall contain a boolean indicator of the write cache
	// policy for the Volume.
	WriteCachePolicy WriteCachePolicyType
	// WriteCacheState shall contain the WriteCacheState policy setting for the
	// Volume.
	WriteCacheState WriteCacheStateType
	// WriteHoleProtectionPolicy specifies the policy that is enabled to address
	// the write hole issue on the RAID volume. If no policy is enabled at the
	// moment, this property shall be set to 'Off'.
	WriteHoleProtectionPolicy WriteHoleProtectionPolicyType
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	cacheDataVolumes []string
	// CacheDatavolumesCount is the number of cache data volumes.
	CacheDataVolumesCount int
	cacheVolumeSource     string
	// classOfService shall contain a reference to the
	// ClassOfService that this storage volume conforms to.
	classOfService  string
	clientEndpoints []string
	// ClientEndpointsCount is the number of client endpoints.
	ClientEndpointsCount int
	consistencyGroups    []string
	// ConsistencyGroupsCount is the number of consistency groups associated with this volume.
	ConsistencyGroupsCount int
	controllers            []string
	// ControllersCount is the number of storage controllers associated with this volume.
	ControllersCount     int
	dedicatedSpareDrives []string
	// DedicatedSpareDrivesCount is the number of dedicates spare drives
	DedicatedSpareDrivesCount int
	drives                    []string
	// DrivesCount is the number of associated drives.
	DrivesCount           int
	journalingMedia       string
	owningStorageResource string
	owningStorageService  string
	providingStoragePool  string
	serverEndpoints       []string
	// ServerEndpointsCount is the number of server endpoints this volume is associated with.
	ServerEndpointsCount int
	spareResourceSets    []string
	// SpareResourceSetsCount is the number of spare resources sets.
	SpareResourceSetsCount int
	storageGroups          []string
	// StorageGroupsCount is the number of storage groups associated with this volume.
	StorageGroupsCount int

	// assignReplicaTargetTarget is the URL to send AssignReplicaTarget requests.
	assignReplicaTargetTarget string
	// checkConsistencyTarget is the URL to send CheckConsistency requests.
	checkConsistencyTarget string
	// createReplicaTargetTarget is the URL to send CreateReplicaTarget requests.
	createReplicaTargetTarget string
	// initializeTarget is the URL to send Initialize requests.
	initializeTarget string
	// removeReplicaRelationshipTarget is the URL to send RemoveReplicaRelationship requests.
	removeReplicaRelationshipTarget string
	// resumeReplicationTarget is the URL to send ResumeReplication requests.
	resumeReplicationTarget string
	// reverseReplicationRelationshipTarget is the URL to send
	// ReverseReplicationRelationship requests.
	reverseReplicationRelationshipTarget string
	// splitReplicationTarget is the URL to send SplitReplication requests.
	splitReplicationTarget string
	// suspendReplicationTarget is the URL to send SuspendReplication requests.
	suspendReplicationTarget string
}

type volumeLinks struct {
	CacheDataVolumes          common.Links
	CacheDataVolumesCount     int `json:"CacheDataVolumes@odata.count"`
	CacheVolumeSource         common.Link
	ClassOfService            common.Link
	ClientEndpoints           common.Links
	ClientEndpointsCount      int `json:"ClientEndpoints@odata.count"`
	ConsistencyGroups         common.Links
	ConsistencyGroupsCount    int `json:"ConsistencyGroups@odata.count"`
	Controllers               common.Links
	ControllersCount          int `json:"Controllers@odata.count"`
	DedicatedSpareDrives      common.Links
	DedicatedSpareDrivesCount int `json:"DedicatedSpareDrives@odata.count"`
	Drives                    common.Links
	DrivesCount               int `json:"Drives@odata.count"`
	JournalingMedia           common.Link
	OwningStorageResource     common.Link
	OwningStorageService      common.Link
	ProvidingStoragePool      common.Link
	ServerEndpoints           common.Links
	ServerEndpointsCount      int `json:"ServerEndpoints@odata.count"`
	SpareResourceSets         common.Links
	SpareResourceSetsCount    int `json:"SpareResourceSets@odata.count"`
	StorageGroups             common.Links
	StorageGroupsCount        int `json:"StorageGroups@odata.count"`
}

type volumeActions struct {
	AssignReplicaTarget            common.ActionTarget `json:"#Volume.AssignReplicaTarget"`
	CheckConsistency               common.ActionTarget `json:"#Volume.CheckConsistency"`
	CreateReplicaTarget            common.ActionTarget `json:"#Volume.CreateReplicaTarget"`
	Initialize                     common.ActionTarget `json:"#Volume.Initialize"`
	RemoveReplicaRelationship      common.ActionTarget `json:"#Volume.RemoveReplicaRelationship"`
	ResumeReplication              common.ActionTarget `json:"#Volume.ResumeReplication"`
	ReverseReplicationRelationship common.ActionTarget `json:"#Volume.ReverseReplicationRelationship"`
	SplitReplication               common.ActionTarget `json:"#Volume.SplitReplication"`
	SuspendReplication             common.ActionTarget `json:"#Volume.SuspendReplication"`
}

// UnmarshalJSON unmarshals a Volume object from the raw JSON.
func (volume *Volume) UnmarshalJSON(b []byte) error {
	type temp Volume

	var t struct {
		temp
		AllocatedPools  common.Links
		CapacitySources common.Links
		Connections     common.Links
		Metrics         common.Link
		ReplicaTargets  common.Links
		StorageGroups   common.Links
		Links           volumeLinks
		Actions         volumeActions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*volume = Volume(t.temp)
	volume.allocatedPools = t.AllocatedPools.ToStrings()
	volume.capacitySources = t.CapacitySources.ToStrings()
	volume.connections = t.Connections.ToStrings()
	volume.metrics = t.Metrics.String()
	volume.replicaTargets = t.ReplicaTargets.ToStrings()
	volume.storageGroups = t.StorageGroups.ToStrings()

	volume.cacheDataVolumes = t.Links.CacheDataVolumes.ToStrings()
	volume.CacheDataVolumesCount = t.Links.CacheDataVolumesCount
	volume.cacheVolumeSource = t.Links.CacheVolumeSource.String()
	volume.classOfService = t.Links.ClassOfService.String()
	volume.clientEndpoints = t.Links.ClientEndpoints.ToStrings()
	volume.ClientEndpointsCount = t.Links.ClientEndpointsCount
	volume.consistencyGroups = t.Links.ConsistencyGroups.ToStrings()
	volume.ConsistencyGroupsCount = t.Links.ConsistencyGroupsCount
	volume.controllers = t.Links.Controllers.ToStrings()
	volume.ControllersCount = t.Links.ControllersCount
	volume.dedicatedSpareDrives = t.Links.DedicatedSpareDrives.ToStrings()
	volume.DedicatedSpareDrivesCount = t.Links.DedicatedSpareDrivesCount
	volume.drives = t.Links.Drives.ToStrings()
	volume.DrivesCount = t.Links.DrivesCount
	volume.journalingMedia = t.Links.JournalingMedia.String()
	volume.owningStorageResource = t.Links.OwningStorageResource.String()
	volume.owningStorageService = t.Links.OwningStorageService.String()
	volume.providingStoragePool = t.Links.ProvidingStoragePool.String()
	volume.serverEndpoints = t.Links.ServerEndpoints.ToStrings()
	volume.ServerEndpointsCount = t.Links.ServerEndpointsCount
	volume.spareResourceSets = t.Links.SpareResourceSets.ToStrings()
	volume.SpareResourceSetsCount = t.Links.SpareResourceSetsCount

	if len(volume.storageGroups) == 0 {
		volume.storageGroups = t.Links.StorageGroups.ToStrings()
		volume.StorageGroupsCount = t.Links.StorageGroupsCount
	}

	volume.assignReplicaTargetTarget = t.Actions.AssignReplicaTarget.Target
	volume.checkConsistencyTarget = t.Actions.CheckConsistency.Target
	volume.createReplicaTargetTarget = t.Actions.CreateReplicaTarget.Target
	volume.initializeTarget = t.Actions.Initialize.Target
	volume.removeReplicaRelationshipTarget = t.Actions.RemoveReplicaRelationship.Target
	volume.resumeReplicationTarget = t.Actions.ResumeReplication.Target
	volume.reverseReplicationRelationshipTarget = t.Actions.ReverseReplicationRelationship.Target
	volume.splitReplicationTarget = t.Actions.SplitReplication.Target
	volume.suspendReplicationTarget = t.Actions.SuspendReplication.Target

	// This is a read/write object, so we need to save the raw object data for later
	volume.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (volume *Volume) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Volume)
	err := original.UnmarshalJSON(volume.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AccessCapabilities",
		"CapacityBytes",
		"CapacitySources",
		"Compressed",
		"Deduplicated",
		"DisplayName",
		"Encrypted",
		"EncryptionTypes",
		"IOPerfModeEnabled",
		"IsBootCapable",
		"LowSpaceWarningThresholdPercents",
		"ProvisioningPolicy",
		"ReadCachePolicy",
		"RecoverableCapacitySourceCount",
		"ReplicationEnabled",
		"StripSizeBytes",
		"WriteCachePolicy",
		"WriteHoleProtectionPolicy",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(volume).Elem()

	return volume.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetVolume will get a Volume instance from the service.
func GetVolume(c common.Client, uri string) (*Volume, error) {
	return common.GetObject[Volume](c, uri)
}

// ListReferencedVolumes gets the collection of Volume from a provided reference.
func ListReferencedVolumes(c common.Client, link string) ([]*Volume, error) {
	return common.GetCollectionObjects[Volume](c, link)
}

// CacheDataVolumes gets the data volumes this volume serves as a cache volume.
func (volume *Volume) CacheDataVolumes() ([]*Volume, error) {
	return common.GetObjects[Volume](volume.GetClient(), volume.cacheDataVolumes)
}

// CacheVolumeSources gets the cache volume source for this volume.
func (volume *Volume) CacheVolumeSource() (*Volume, error) {
	if volume.cacheVolumeSource == "" {
		return nil, nil
	}

	return GetVolume(volume.GetClient(), volume.cacheVolumeSource)
}

// ClassOfService gets the class of service that this storage volume conforms to.
func (volume *Volume) ClassOfService() (*ClassOfService, error) {
	if volume.classOfService == "" {
		return nil, nil
	}

	return GetClassOfService(volume.GetClient(), volume.classOfService)
}

// ClientEndpoints gets the client Endpoints associated with this volume.
func (volume *Volume) ClientEndpoints() ([]*redfish.Endpoint, error) {
	return common.GetObjects[redfish.Endpoint](volume.GetClient(), volume.clientEndpoints)
}

// ConsistencyGroups gets the ConsistencyGroups associated with this volume.
func (volume *Volume) ConsistencyGroups() ([]*ConsistencyGroup, error) {
	return common.GetObjects[ConsistencyGroup](volume.GetClient(), volume.consistencyGroups)
}

// Controllers gets the controllers (of type StorageController) associated with
// this volume. When the volume is of type NVMe, these may be both the physical
// and logical controller representations.
func (volume *Volume) Controllers() ([]*redfish.StorageController, error) {
	return common.GetObjects[redfish.StorageController](volume.GetClient(), volume.controllers)
}

// getDrives gets a set of referenced drives.
func (volume *Volume) getDrives(links []string) ([]*redfish.Drive, error) {
	return common.GetObjects[redfish.Drive](volume.GetClient(), links)
}

// DedicatedSpareDrives references the Drives that are dedicated spares for this
// volume.
func (volume *Volume) DedicatedSpareDrives() ([]*redfish.Drive, error) {
	return volume.getDrives(volume.dedicatedSpareDrives)
}

// Drives references the Drives that are associated with this volume.
func (volume *Volume) Drives() ([]*redfish.Drive, error) {
	return volume.getDrives(volume.drives)
}

// OwningStorageResource gets the Storage resource that owns or contains this volume.
func (volume *Volume) OwningStorageResource() (*redfish.Storage, error) {
	if volume.owningStorageResource == "" {
		return nil, nil
	}

	return redfish.GetStorage(volume.GetClient(), volume.owningStorageResource)
}

// OwningStorageService gets the StorageService that owns or contains this volume.
func (volume *Volume) OwningStorageService() (*StorageService, error) {
	if volume.owningStorageService == "" {
		return nil, nil
	}

	return GetStorageService(volume.GetClient(), volume.owningStorageService)
}

// ProvidingStoragePool gets the StoragePool resource that provides this volume resource.
func (volume *Volume) ProvidingStoragePool() (*StoragePool, error) {
	if volume.providingStoragePool == "" {
		return nil, nil
	}

	return GetStoragePool(volume.GetClient(), volume.providingStoragePool)
}

// ServerEndpoints gets the server Endpoints associated with this volume.
func (volume *Volume) ServerEndpoints() ([]*redfish.Endpoint, error) {
	return common.GetObjects[redfish.Endpoint](volume.GetClient(), volume.serverEndpoints)
}

// SpareResourceSets gets the spare resources that can be used for this volume.
func (volume *Volume) SpareResourceSets() ([]*SpareResourceSet, error) {
	return common.GetObjects[SpareResourceSet](volume.GetClient(), volume.spareResourceSets)
}

// StorageGroups gets the storage groups that associated with this volume.
// This property is deprecated in favor of the Connections property.
func (volume *Volume) StorageGroups() ([]*StorageGroup, error) {
	return common.GetObjects[StorageGroup](volume.GetClient(), volume.storageGroups)
}

// AllocatedPools gets the storage pools that associated with this volume.
func (volume *Volume) AllocatedPools() ([]*StoragePool, error) {
	return common.GetObjects[StoragePool](volume.GetClient(), volume.allocatedPools)
}

// CapacitySources gets the space allocations to this volume.
func (volume *Volume) CapacitySources() ([]*CapacitySource, error) {
	return common.GetObjects[CapacitySource](volume.GetClient(), volume.capacitySources)
}

// Connections gets the connections that include this volume.
func (volume *Volume) Connections() ([]*redfish.Connection, error) {
	return common.GetObjects[redfish.Connection](volume.GetClient(), volume.connections)
}

// Metrics gets the metrics for this volume. IO metrics are reported in the IOStatistics property.
func (volume *Volume) Metrics() (*VolumeMetrics, error) {
	if volume.metrics == "" {
		return nil, nil
	}

	return GetVolumeMetrics(volume.GetClient(), volume.metrics)
}

// AssignReplicaTarget is used to establish a replication relationship by
// assigning an existing volume to serve as a target replica for an existing
// source volume.
func (volume *Volume) AssignReplicaTarget(replicaType ReplicaType, updateMode ReplicaUpdateMode, targetVolumeODataID string) error {
	// This action wasn't added until later revisions
	if volume.assignReplicaTargetTarget == "" {
		return fmt.Errorf("AssignReplicaTarget action is not supported by this system")
	}

	// Define this action's parameters
	// Set the values for the action arguments
	t := struct {
		ReplicaType       ReplicaType
		ReplicaUpdateMode ReplicaUpdateMode
		TargetVolume      string
	}{
		ReplicaType:       replicaType,
		ReplicaUpdateMode: updateMode,
		TargetVolume:      targetVolumeODataID,
	}

	return volume.Post(volume.assignReplicaTargetTarget, t)
}

// CheckConsistency is used to force a check of the Volume's parity or redundant
// data to ensure it matches calculated values.
func (volume *Volume) CheckConsistency() error {
	if volume.checkConsistencyTarget == "" {
		return fmt.Errorf("CheckConsistency action is not supported by this system")
	}

	return volume.Post(volume.checkConsistencyTarget, nil)
}

// Initialize is used to prepare the contents of the volume for use by the system.
func (volume *Volume) Initialize(initType InitializeType) error {
	if volume.initializeTarget == "" {
		return fmt.Errorf("initialize action is not supported by this system")
	}

	// Define this action's parameters
	// Set the values for the action arguments
	t := struct {
		InitializeType InitializeType
	}{InitializeType: initType}

	return volume.Post(volume.initializeTarget, t)
}

// RemoveReplicaRelationship is used to disable data synchronization between a
// source and target volume, remove the replication relationship, and optionally
// delete the target volume.
func (volume *Volume) RemoveReplicaRelationship(deleteTarget bool, targetVolumeODataID string) error {
	// This action wasn't added until later revisions
	if volume.removeReplicaRelationshipTarget == "" {
		return fmt.Errorf("RemoveReplicaRelationship action is not supported by this system")
	}

	// Define this action's parameters
	// Set the values for the action arguments
	t := struct {
		DeleteTargetVolume bool
		TargetVolume       string
	}{
		DeleteTargetVolume: deleteTarget,
		TargetVolume:       targetVolumeODataID,
	}

	return volume.Post(volume.removeReplicaRelationshipTarget, t)
}

// ResumeReplication is used to resume the active data synchronization between a
// source and target volume, without otherwise altering the replication
// relationship.
func (volume *Volume) ResumeReplication(targetVolumeODataID string) error {
	// This action wasn't added until later revisions
	if volume.resumeReplicationTarget == "" {
		return fmt.Errorf("ResumeReplication action is not supported by this system")
	}

	// Define this action's parameters
	// Set the values for the action arguments
	t := struct {
		TargetVolume string
	}{TargetVolume: targetVolumeODataID}

	return volume.Post(volume.resumeReplicationTarget, t)
}

// ReverseReplicationRelationship is used to reverse the replication
// relationship between a source and target volume.
func (volume *Volume) ReverseReplicationRelationship(targetVolumeODataID string) error {
	// This action wasn't added until later revisions
	if volume.reverseReplicationRelationshipTarget == "" {
		return fmt.Errorf("ReverseReplicationRelationship action is not supported by this system")
	}

	// Define this action's parameters
	// Set the values for the action arguments
	t := struct {
		TargetVolume string
	}{TargetVolume: targetVolumeODataID}

	return volume.Post(volume.reverseReplicationRelationshipTarget, t)
}

// SplitReplication is used to split the replication relationship and suspend
// data synchronization between a source and target volume.
func (volume *Volume) SplitReplication(targetVolumeODataID string) error {
	// This action wasn't added until later revisions
	if volume.splitReplicationTarget == "" {
		return fmt.Errorf("SplitReplication action is not supported by this system")
	}

	// Define this action's parameters
	// Set the values for the action arguments
	t := struct {
		TargetVolume string
	}{TargetVolume: targetVolumeODataID}

	return volume.Post(volume.splitReplicationTarget, t)
}

// SuspendReplication is used to suspend active data synchronization between a
// source and target volume, without otherwise altering the replication
// relationship.
func (volume *Volume) SuspendReplication(targetVolumeODataID string) error {
	// This action wasn't added until later revisions
	if volume.suspendReplicationTarget == "" {
		return fmt.Errorf("SuspendReplication action is not supported by this system")
	}

	// Define this action's parameters
	// Set the values for the action arguments
	t := struct {
		TargetVolume string
	}{TargetVolume: targetVolumeODataID}

	return volume.Post(volume.suspendReplicationTarget, t)
}
