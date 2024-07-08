//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type InitializeMethod string

const (
	// BackgroundInitializeMethod volume will be available for use immediately, with data erasure and preparation to happen as background tasks.
	BackgroundInitializeMethod InitializeMethod = "Background"
	// ForegroundInitializeMethod Data erasure and preparation tasks will complete before the volume is presented as available for use.
	ForegroundInitializeMethod InitializeMethod = "Foreground"
	// SkipInitializeMethod volume will be available for use immediately, with no preparation.
	SkipInitializeMethod InitializeMethod = "Skip"
)

type InitializeType string

const (
	// FastInitializeType volume is prepared for use quickly, typically by erasing just the beginning and end of the space so that partitioning can be performed.
	FastInitializeType InitializeType = "Fast"
	// SlowInitializeType volume is prepared for use slowly, typically by completely erasing the volume.
	SlowInitializeType InitializeType = "Slow"
)

type LBAFormat struct {
	// LBADataSizeBytes shall be the LBA data size reported in bytes.
	LBADataSizeBytes int
	// LBAFormatType shall be the LBA format type. This property is intended for capabilities instrumentation.
	LBAFormatType string
	// LBAMetadataSizeBytes shall be the LBA metadata size reported in bytes.
	LBAMetadataSizeBytes int
	// RelativePerformance shall be the LBA Relative Performance type. This field indicates the relative performance of
	// the LBA format indicated relative to other LBA formats supported by the controller. This property is intended
	// for capabilities instrumentation.
	RelativePerformance string
}

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
	LBAFormatsSupported []string
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
	NamespaceType string
	// NumberLBAFormats shall contain the number of LBA data size and metadata size combinations supported by this
	// namespace. The value of this property is between 0 and 16. LBA formats with an index set beyond this value will
	// not be supported.
	NumberLBAFormats int
	// Type shall identify the type of namespace.
	Type string
}

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
	Operation string
	// PercentageComplete The percentage of the operation that has been completed.
	PercentageComplete int
}

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

// EncryptionTypes is the type of encryption used by the volume.
type EncryptionTypes string

const (
	// NativeDriveEncryptionEncryptionTypes indicates the volume is is utilizing the
	// native drive encryption capabilities of the drive hardware.
	NativeDriveEncryptionEncryptionTypes EncryptionTypes = "NativeDriveEncryption"
	// ControllerAssistedEncryptionTypes indicates the volume is is being encrypted by the
	// storage controller entity.
	ControllerAssistedEncryptionTypes EncryptionTypes = "ControllerAssisted"
	// SoftwareAssistedEncryptionTypes indicates the volume is is being encrypted by
	// software running on the system or the operating system.
	SoftwareAssistedEncryptionTypes EncryptionTypes = "SoftwareAssisted"
)

// VolumeType is the type of volume.
type VolumeType string

const (
	// RawDeviceVolumeType indicates the volume is is a raw physical device without any
	// RAID or other virtualization applied.
	RawDeviceVolumeType VolumeType = "RawDevice"
	// NonRedundantVolumeType indicates the volume is is a non-redundant storage device.
	NonRedundantVolumeType VolumeType = "NonRedundant"
	// MirroredVolumeType indicates the volume is is a mirrored device.
	MirroredVolumeType VolumeType = "Mirrored"
	// StripedWithParityVolumeType indicates the volume is is a device which uses parity
	// to retain redundant information.
	StripedWithParityVolumeType VolumeType = "StripedWithParity"
	// SpannedMirrorsVolumeType indicates the volume is is a spanned set of mirrored
	// devices.
	SpannedMirrorsVolumeType VolumeType = "SpannedMirrors"
	// SpannedStripesWithParityVolumeType indicates the volume is is a spanned set of
	// devices which uses parity to retain redundant information.
	SpannedStripesWithParityVolumeType VolumeType = "SpannedStripesWithParity"
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

type StorageAccessCapability string

const (
	// AppendStorageAccessCapability is AppendOnly.
	AppendStorageAccessCapability StorageAccessCapability = "Append"
	// ExecuteStorageAccessCapability is Execute access is allowed by the file share.
	ExecuteStorageAccessCapability StorageAccessCapability = "Execute"
	// ReadStorageAccessCapability is Read.
	ReadStorageAccessCapability StorageAccessCapability = "Read"
	// StreamingStorageAccessCapability is Streaming.
	StreamingStorageAccessCapability StorageAccessCapability = "Streaming"
	// WriteStorageAccessCapability is Write Many.
	WriteStorageAccessCapability StorageAccessCapability = "Write"
	// WriteOnceStorageAccessCapability is write once.
	WriteOnceStorageAccessCapability StorageAccessCapability = "WriteOnce"
)

type ProvisioningPolicy string

const (
	FixedProvisioningPolicy ProvisioningPolicy = "Fixed"
	ThinProvisioningPolicy  ProvisioningPolicy = "Thin"
)

type WriteCacheStateType string

const (
	// DegradedWriteCacheStateType indicates an issue with the cache state in which the
	// cache space is diminished or disabled due to a failure or an outside influence
	// such as a discharged battery.
	DegradedWriteCacheStateType WriteCacheStateType = "Degraded"
	// ProtectedWriteCacheStateType indicates that the cache state type in use generally
	// protects write requests on non-volatile media.
	ProtectedWriteCacheStateType WriteCacheStateType = "Protected"
	// UnprotectedWriteCacheStateType indicates that the cache state type in use generally
	// does not protect write requests on non-volatile media.
	UnprotectedWriteCacheStateType WriteCacheStateType = "Unprotected"
)

type volumeLinks struct {
	// CacheDataVolumes shall be a pointer to the cache data volumes this volume serves as a cache volume. The
	// corresponding VolumeUsage property shall be set to CacheOnly when this property is used.
	CacheDataVolumes common.Links
	// CacheDataVolumes@odata.count
	CacheDataVolumesCount int `json:"CacheDataVolumes@odata.count"`
	// CacheVolumeSource shall be a pointer to the cache volume source for this volume. The corresponding VolumeUsage
	// property shall be set to Data when this property is used.
	CacheVolumeSource common.Link
	// ClassOfService shall contain a reference to the ClassOfService that this storage volume conforms to.
	ClassOfService common.Link
	// ClientEndpoints shall be references to the client Endpoints this volume is associated with.
	ClientEndpoints common.Links
	// ClientEndpoints@odata.count
	ClientEndpointsCount int `json:"ClientEndpoints@odata.count"`
	// ConsistencyGroups shall be references to the ConsistencyGroups this volume is associated with.
	ConsistencyGroups common.Links
	// ConsistencyGroupsCount is the number of ConsistencyGroups associated with this volume.
	ConsistencyGroupsCount int `json:"ConsistencyGroups@odata.count"`
	// Controllers shall contain an array of the controllers (of type StorageController) associated with this volume.
	// When the volume is of type NVMe, these may be both the physical and logical controller representations.
	Controllers common.Links
	// Controllers@odata.count
	ControllersCount int `json:"Controllers@odata.count"`
	// DedicatedSpareDrives shall be a reference to the resources that this volume is associated with and shall
	// reference resources of type Drive. This property shall only contain references to Drive entities which are
	// currently assigned as a dedicated spare and are able to support this Volume.
	DedicatedSpareDrives common.Links
	// DedicatedSpareDrives@odata.count
	DedicatedSpareDrivesCount int `json:"DedicatedSpareDrives@odata.count"`
	// Drives shall be a reference to the resources that this volume is associated with and shall reference resources
	// of type Drive. This property shall only contain references to Drive entities which are currently members of the
	// Volume, not hot spare Drives which are not currently a member of the volume.
	Drives common.Links
	// Drives@odata.count
	DrivesCount int `json:"Drives@odata.count"`
	// JournalingMedia shall be a pointer to the journaling media used for this Volume to address the write hole issue.
	// Valid when WriteHoleProtectionPolicy property is set to 'Journaling'.
	JournalingMedia common.Link
	// OwningStorageResource shall be a pointer to the Storage resource that owns or contains this volume.
	OwningStorageResource common.Link
	// OwningStorageService shall be a pointer to the StorageService that owns or contains this volume.
	OwningStorageService common.Link
	// ServerEndpoints shall be references to the server Endpoints this volume is associated with.
	ServerEndpoints common.Links
	// ServerEndpoints@odata.count
	ServerEndpointsCount int `json:"ServerEndpoints@odata.count"`
	// SpareResourceSets shall contain resources that may be utilized to replace the capacity provided by a failed
	// resource having a compatible type.
	SpareResourceSets common.Links
	// SpareResourceSets@odata.count
	SpareResourceSetsCount int `json:"SpareResourceSets@odata.count"`
	// StorageGroups shall be references to the StorageGroups this volume is associated with.
	StorageGroups common.Links
	// StorageGroups@odata.count
	StorageGroupsCount int `json:"StorageGroups@odata.count"`
}

type volumeActions struct {
	AssignReplicaTarget            common.ActionTarget `json:"#Volume.AssignReplicaTarget"`
	ChangeRAIDLayout               common.ActionTarget `json:"#Volume.ChangeRAIDLayout"`
	CheckConsistency               common.ActionTarget `json:"#Volume.CheckConsistency"`
	CreateReplicaTarget            common.ActionTarget `json:"#Volume.CreateReplicaTarget"`
	ForceEnable                    common.ActionTarget `json:"#Volume.ForceEnable"`
	Initialize                     common.ActionTarget `json:"#Volume.Initialize"`
	RemoveReplicaRelationship      common.ActionTarget `json:"#Volume.RemoveReplicaRelationship"`
	ResumeReplication              common.ActionTarget `json:"#Volume.ResumeReplication"`
	ReverseReplicationRelationship common.ActionTarget `json:"#Volume.ReverseReplicationRelationship"`
	SplitReplication               common.ActionTarget `json:"#Volume.SplitReplication"`
	SuspendReplication             common.ActionTarget `json:"#Volume.SuspendReplication"`
}

// Volume is used to represent a volume, virtual disk, logical disk, LUN,
// or other logical storage for a Redfish implementation.
type Volume struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccessCapabilities shall specify a current storage access capability.
	AccessCapabilities []StorageAccessCapability
	// AllocatedPools shall contain references to all storage pools allocated from this volume.
	AllocatedPools string
	// BlockSizeBytes shall contain size of the smallest addressable unit of the associated volume.
	BlockSizeBytes int
	// Capacity Information about the utilization of capacity allocated to this storage volume.
	Capacity string
	// CapacityBytes shall contain the size in bytes of the associated volume.
	CapacityBytes int
	// CapacitySources Fully or partially consumed storage from a source resource. Each entry provides capacity
	// allocation information from a named source resource. Uses the Swordfish CapacitySource.
	CapacitySources json.RawMessage
	// CapacitySourcesCount is the number of capacity sources.
	CapacitySourcesCount int `json:"CapacitySources@odata.count"`
	// Compressed shall contain a boolean indicator if the Volume is currently utilizing compression or not.
	Compressed bool
	// Connections shall contain references to all Connections that include this volume.
	Connections []Connection
	// ConnectionsCount is the number of connections.
	ConnectionsCount int `json:"Connections@odata.count"`
	// Deduplicated shall contain a boolean indicator if the Volume is currently utilizing deduplication or not.
	Deduplicated bool
	// Actions shall contain the available actions for this resource.
	// Description provides a description of this resource.
	Description string
	// DisplayName shall contain a user-configurable string to name the volume.
	DisplayName string
	// Encrypted shall contain a boolean indicator if the Volume is currently utilizing encryption or not.
	Encrypted bool
	// EncryptionTypes shall contain the types of encryption used by this Volume.
	EncryptionTypes []EncryptionTypes
	// IOPerfModeEnabled shall indicate whether IO performance mode is enabled for the volume.
	IOPerfModeEnabled bool
	// IOStatistics shall represent IO statistics for this volume.
	IOStatistics string
	// Identifiers shall contain a list of all known durable names for the associated volume.
	Identifiers []common.Identifier
	// InitializeMethod shall indicate the initialization method used for this volume. If InitializeMethod is not
	// specified, the InitializeMethod should be Foreground. This value reflects the most recently used Initialization
	// Method, and may be changed using the Initialize Action.
	InitializeMethod string
	// IsBootCapable shall indicate whether or not the Volume contains a boot image and is capable of booting. This
	// property may be settable by an admin or client with visibility into the contents of the volume. This property
	// should only be set to true when VolumeUsage is either not specified, or when VolumeUsage is set to Data or
	// SystemData.
	IsBootCapable bool
	// LogicalUnitNumber shall contain host-visible LogicalUnitNumber assigned to this Volume. This property shall only
	// be used when in a single connect configuration and no StorageGroup configuration is used.
	LogicalUnitNumber int
	// LowSpaceWarningThresholdPercents shall be triggered: Across all CapacitySources entries, percent =
	// (SUM(AllocatedBytes) - SUM(ConsumedBytes))/SUM(AllocatedBytes).
	LowSpaceWarningThresholdPercents []string
	// Manufacturer shall contain a value that represents the manufacturer or implementer of the storage volume.
	Manufacturer string
	// MaxBlockSizeBytes shall contain size of the largest addressable unit of this storage volume.
	MaxBlockSizeBytes int
	// MediaSpanCount shall indicate the number of media elements used per span in the secondary RAID for a
	// hierarchical RAID type.
	MediaSpanCount int
	// Metrics shall contain a link to a resource of type VolumeMetrics that specifies the metrics for this volume. IO
	// metrics are reported in the IOStatistics property.
	// Uses Swordfish VolumeMetrics.
	Metrics json.RawMessage
	// Model shall represents a specific storage volume implementation.
	Model string
	// NVMeNamespaceProperties shall contain properties to use when Volume is used to describe an NVMe Namespace.
	NVMeNamespaceProperties NVMeNamespaceProperties
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Operations shall contain a list of all currently running on the Volume.
	Operations []common.Operations
	// OptimumIOSizeBytes shall contain the optimum IO size to use when performing IO on this volume. For logical
	// disks, this is the stripe size. For physical disks, this describes the physical sector size.
	OptimumIOSizeBytes int
	// ProvisioningPolicy shall specify the volume's supported storage allocation policy.
	ProvisioningPolicy ProvisioningPolicy
	// RAIDType shall contain the RAID type of the associated Volume.
	RAIDType RAIDType
	// ReadCachePolicy shall contain a boolean indicator of the read cache policy for the Volume.
	ReadCachePolicy ReadCachePolicyType
	// RecoverableCapacitySourceCount The value is the number of available capacity source resources currently
	// available in the event that an equivalent capacity source resource fails.
	RecoverableCapacitySourceCount int
	// RemainingCapacityPercent shall return {[(SUM(AllocatedBytes) - SUM(ConsumedBytes)]/SUM(AllocatedBytes)}*100
	// represented as an integer value.
	RemainingCapacityPercent int
	// RemoteReplicaTargets shall reference the URIs to the remote target replicas that are sourced by this replica.
	// Remote indicates that the replica is managed by a separate Swordfish service instance.
	RemoteReplicaTargets []string
	// ReplicaInfo shall describe the replica relationship between this storage volume and a corresponding source
	// volume.
	// Uses swordfish ReplicaInfo.
	ReplicaInfo json.RawMessage
	// ReplicaTargets shall reference the target replicas that are sourced by this replica.
	ReplicaTargets []string
	// ReplicaTargets@odata.count
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// ReplicationEnabled shall indicate whether or not replication is enabled on the volume. This property shall be
	// consistent with the state reflected at the storage pool level.
	ReplicationEnabled bool
	// Status shall contain the status of the Volume.
	Status common.Status
	// StripSizeBytes The number of consecutively addressed virtual disk blocks (bytes) mapped to consecutively
	// addressed blocks on a single member extent of a disk array. Synonym for stripe depth and chunk size.
	StripSizeBytes int
	// VolumeUsage shall contain the volume usage type for the Volume.
	VolumeUsage string
	// WriteCachePolicy shall contain a boolean indicator of the write cache policy for the Volume.
	WriteCachePolicy WriteCachePolicyType
	// WriteCacheState shall contain the WriteCacheState policy setting for the Volume.
	WriteCacheState WriteCacheStateType
	// WriteHoleProtectionPolicy shall be set to 'Off'.
	WriteHoleProtectionPolicy string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	// VolumeType shall contain the type of the associated Volume.
	// Deprecated in favor of explicit use of RAIDType.
	VolumeType VolumeType

	cacheDataVolumes []string
	// CacheDataVolumesCount is the number of cache data volumes this volume serves
	// as a cache for.
	CacheDataVolumesCount int
	cacheVolumeSource     string
	classOfService        string
	clientEndpoints       []string
	// ClientEndpointsCount is the number of client endpoints this volume is associated with.
	ClientEndpointsCount int
	consistencyGroups    []string
	// ConsistencyGroupsCount is the number of consistency groups this volume is a part of.
	ConsistencyGroupsCount int
	controllers            []string
	// ControllersCount is the number of storage controllers associated with this volume.
	ControllersCount     int
	dedicatedSpareDrives []string
	// DedicatedSpareDrivesCount is the number of spare drives which are dedicated for this volume.
	DedicatedSpareDrivesCount int
	// DrivesCount is the number of associated drives.
	DrivesCount int
	drives      []string
	// JournalingMedia is a pointer to the journaling media used for this Volume to address
	// the write hole issue. Valid when WriteHoleProtectionPolicy property is set to 'Journaling'.
	JournalingMedia       string
	owningStorageResource string
	owningStorageService  string
	serverEndpoints       []string
	// ServerEndpointsCount is the number of server endpoints associated with this volume.
	ServerEndpointsCount int
	spareResourceSets    []string
	// SpareResourceSetsCount is the number of spare resource sets available for this volume.
	SpareResourceSetsCount int
	storageGroups          []string
	// StorageGroupsCount is the number of storage groups associated with this volume.
	StorageGroupsCount int

	assignReplicaTargetTarget            string
	changeRAIDLayoutTarget               string
	checkConsistencyTarget               string
	createReplicaTargetTarget            string
	forceEnableTarget                    string
	initializeTarget                     string
	removeReplicaRelationshipTarget      string
	resumeReplicationTarget              string
	reverseReplicationRelationshipTarget string
	splitReplicationTarget               string
	suspendReplicationTarget             string
}

// UnmarshalJSON unmarshals a Volume object from the raw JSON.
func (volume *Volume) UnmarshalJSON(b []byte) error {
	type temp Volume
	var t struct {
		temp
		Actions volumeActions
		Links   volumeLinks
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*volume = Volume(t.temp)

	// Extract the links to other entities for later
	volume.assignReplicaTargetTarget = t.Actions.AssignReplicaTarget.Target
	volume.changeRAIDLayoutTarget = t.Actions.ChangeRAIDLayout.Target
	volume.checkConsistencyTarget = t.Actions.CheckConsistency.Target
	volume.createReplicaTargetTarget = t.Actions.CreateReplicaTarget.Target
	volume.forceEnableTarget = t.Actions.ForceEnable.Target
	volume.initializeTarget = t.Actions.Initialize.Target
	volume.removeReplicaRelationshipTarget = t.Actions.RemoveReplicaRelationship.Target
	volume.resumeReplicationTarget = t.Actions.ResumeReplication.Target
	volume.reverseReplicationRelationshipTarget = t.Actions.ReverseReplicationRelationship.Target
	volume.splitReplicationTarget = t.Actions.SplitReplication.Target
	volume.suspendReplicationTarget = t.Actions.SuspendReplication.Target

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
	volume.DrivesCount = t.Links.DrivesCount
	volume.drives = t.Links.Drives.ToStrings()
	volume.JournalingMedia = t.Links.JournalingMedia.String()
	volume.owningStorageResource = t.Links.OwningStorageResource.String()
	volume.owningStorageService = t.Links.OwningStorageService.String()
	volume.serverEndpoints = t.Links.ServerEndpoints.ToStrings()
	volume.ServerEndpointsCount = t.Links.ServerEndpointsCount
	volume.spareResourceSets = t.Links.SpareResourceSets.ToStrings()
	volume.SpareResourceSetsCount = t.Links.SpareResourceSetsCount
	volume.storageGroups = t.Links.StorageGroups.ToStrings()
	volume.StorageGroupsCount = t.Links.StorageGroupsCount

	// This is a read/write object, so we need to save the raw object data for later
	volume.rawData = b

	return nil
}

// AssignReplicaTarget is used to establish a replication relationship by assigning an
// existing volume to serve as a target replica for an existing source volume.
//
// `replicaType` is the Swordfish-defined ReplicaType to be created.
//
// `replicaUpdateMode` is the Swordfish-defined ReplicaUpdateMode (synchronous vs asynchronous).
//
// `targetVolumeURI` is the URI to the existing target volume.
func (volume *Volume) AssignReplicaTarget(replicaType, replicaUpdateMode, targetVolumeURI string) error {
	if volume.assignReplicaTargetTarget == "" {
		return errors.New("assignReplicaTarget is not supported by this volume")
	}
	t := struct {
		ReplicaType       string
		ReplicaUpdateMode string
		TargetVolume      string
	}{
		ReplicaType:       replicaType,
		ReplicaUpdateMode: replicaUpdateMode,
		TargetVolume:      targetVolumeURI,
	}
	return volume.Post(volume.assignReplicaTargetTarget, t)
}

// ChangeRAIDLayout will request the system to change the RAID layout of the volume.
// Depending on the combination of the submitted parameters, this could be changing
// the RAID type, changing the span count, changing the number of drives used by the
// volume, or another configuration change supported by the system. Note that usage
// of this action while online may potentially cause data loss if the available
// capacity is reduced.
//
// `drives` is an array of the drives to be used by the volume.
//
// `mediaSpanCount` is the number of media elements used per span in the secondary RAID for a hierarchical RAID type.
//
// `raidType` is the Swordfish-defined RAIDType for the volume.
//
// `stripSizeBytes` is the number of blocks (bytes) requested for the new strip size.
func (volume *Volume) ChangeRAIDLayout(drives []*Drive, mediaSpanCount int, raidType string, stripSizeBytes int) error {
	if volume.changeRAIDLayoutTarget == "" {
		return errors.New("changeRAIDLayout is not supported by this volume")
	}

	t := struct {
		Drives         []*Drive
		MediaSpanCount int
		RAIDType       string
		StripSizeBytes int
	}{
		Drives:         drives,
		MediaSpanCount: mediaSpanCount,
		RAIDType:       raidType,
		StripSizeBytes: stripSizeBytes,
	}
	return volume.Post(volume.changeRAIDLayoutTarget, t)
}

// CheckConsistency is used to force a check of the Volume's parity or redundant
// data to ensure it matches calculated values.
func (volume *Volume) CheckConsistency() error {
	if volume.checkConsistencyTarget == "" {
		return errors.New("checkConsistency is not supported by this volume")
	}
	return volume.Post(volume.checkConsistencyTarget, nil)
}

// CreateReplicaTarget is used to create a new volume resource to provide expanded
// data protection through a replica relationship with the specified source volume.
//
// `replicaType` is the Swordfish-defined ReplicaType to be created.
//
// `replicaUpdateMode` is the Swordfish-defined ReplicaUpdateMode (synchronous vs asynchronous).
//
// `targetStoragePoolURI` is the URI to the existing target storage pool.
//
// `volumeName` is the name for the new target volume.
func (volume *Volume) CreateReplicaTarget(replicaType, replicaUpdateMode, targetStoragePoolURI, volumeName string) error {
	if volume.createReplicaTargetTarget == "" {
		return errors.New("createReplicaTarget is not supported by this volume")
	}
	t := struct {
		ReplicaType       string
		ReplicaUpdateMode string
		TargetStoragePool string
		VolumeName        string
	}{
		ReplicaType:       replicaType,
		ReplicaUpdateMode: replicaUpdateMode,
		TargetStoragePool: targetStoragePoolURI,
		VolumeName:        volumeName,
	}
	return volume.Post(volume.createReplicaTargetTarget, t)
}

// ForceEnable is used to force the volume to an enabled state regardless of data loss.
func (volume *Volume) ForceEnable() error {
	if volume.forceEnableTarget == "" {
		return errors.New("forceEnable is not supported by this volume")
	}
	return volume.Post(volume.forceEnableTarget, nil)
}

// Initialize is used to prepare the contents of the volume for use by the system.
// If InitializeMethod is not specified in the request body, but the property
// InitializeMethod is specified, the property InitializeMethod value should be
// used. If neither is specified, the InitializeMethod should be Foreground.
//
// `initializeMethod` is the Swordfish-defined InitializeMethod to be performed.
//
// `initializeType` is the Swordfish-defined InitializeType to be performed.
func (volume *Volume) Initialize(initializeMethod InitializeMethod, initializeType InitializeType) error {
	if volume.initializeTarget == "" {
		return errors.New("initialize is not supported by this volume")
	}
	t := struct {
		InitializeMethod InitializeMethod
		InitializeType   InitializeType
	}{
		InitializeMethod: initializeMethod,
		InitializeType:   initializeType,
	}
	return volume.Post(volume.initializeTarget, t)
}

// RemoveReplicaRelationship is used to disable data synchronization between a source and
// target volume, remove the replication relationship, and optionally delete the target volume.
//
// `deleteTargetVolume` indicates whether to delete the target volume as part of the operation.
//
// `targetVolumeURI` is the URI to the existing target volume.
func (volume *Volume) RemoveReplicaRelationship(deleteTargetVolume bool, targetVolumeURI string) error {
	if volume.removeReplicaRelationshipTarget == "" {
		return errors.New("removeReplicaRelationship not supported by this volume")
	}

	t := struct {
		DeleteTargetVolume bool
		TargetVolume       string
	}{
		DeleteTargetVolume: deleteTargetVolume,
		TargetVolume:       targetVolumeURI,
	}
	return volume.Post(volume.removeReplicaRelationshipTarget, t)
}

// ResumeReplication is used to resume the active data synchronization between a source
// and target volume, without otherwise altering the replication relationship.
//
// `targetVolumeURI` is the URI to the existing target volume.
func (volume *Volume) ResumeReplication(targetVolumeURI string) error {
	if volume.resumeReplicationTarget == "" {
		return errors.New("resumeReplication not supported by this volume")
	}

	t := struct {
		TargetVolume string
	}{
		TargetVolume: targetVolumeURI,
	}
	return volume.Post(volume.resumeReplicationTarget, t)
}

// ReverseReplicationRelationship is used to reverse the replication relationship
// between a source and target volume.
//
// `targetVolumeURI` is the URI to the existing target volume.
func (volume *Volume) ReverseReplicationRelationship(targetVolumeURI string) error {
	if volume.reverseReplicationRelationshipTarget == "" {
		return errors.New("reverseReplicationRelationship not supported by this volume")
	}

	t := struct {
		TargetVolume string
	}{
		TargetVolume: targetVolumeURI,
	}
	return volume.Post(volume.reverseReplicationRelationshipTarget, t)
}

// SplitReplication is used to split the replication relationship and suspend data
// synchronization between a source and target volume.
//
// `targetVolumeURI` is the URI to the existing target volume.
func (volume *Volume) SplitReplication(targetVolumeURI string) error {
	if volume.splitReplicationTarget == "" {
		return errors.New("splitReplication not supported by this volume")
	}

	t := struct {
		TargetVolume string
	}{
		TargetVolume: targetVolumeURI,
	}
	return volume.Post(volume.splitReplicationTarget, t)
}

// SuspendReplication is used to suspend active data synchronization between a source
// and target volume, without otherwise altering the replication relationship.
//
// `targetVolumeURI` is the URI to the existing target volume.
func (volume *Volume) SuspendReplication(targetVolumeURI string) error {
	if volume.suspendReplicationTarget == "" {
		return errors.New("suspendReplication not supported by this volume")
	}

	t := struct {
		TargetVolume string
	}{
		TargetVolume: targetVolumeURI,
	}
	return volume.Post(volume.suspendReplicationTarget, t)
}

// CacheDataVolumes gets the cache data volumes this volume serves as a cache volume. The
// corresponding VolumeUsage property shall be set to CacheOnly when this property is used.
func (volume *Volume) CacheDataVolumes() ([]*Volume, error) {
	return common.GetObjects[Volume](volume.GetClient(), volume.cacheDataVolumes)
}

// CacheVolumeSource gets the cache volume source for this volume. The corresponding VolumeUsage
// property shall be set to Data when this property is used.
func (volume *Volume) CacheVolumeSource() (*Volume, error) {
	if volume.cacheVolumeSource == "" {
		return nil, nil
	}
	return GetVolume(volume.GetClient(), volume.cacheVolumeSource)
}

// // ClassOfService gets the cache volume source for this volume. The corresponding VolumeUsage
// // property shall be set to Data when this property is used.
// func (volume *Volume) ClassOfService() (*swordfish.ClassOfService, error) {
// 	if volume.classOfService == "" {
// 		return nil, nil
// 	}
// 	return swordfish.GetClassOfService(volume.GetClient(), volume.classOfService)
// }

// ClientEndpoints gets the client Endpoints this volume is associated with.
func (volume *Volume) ClientEndpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](volume.GetClient(), volume.clientEndpoints)
}

// Controllers gets the controllers (of type StorageController) associated with this volume.
// When the volume is of type NVMe, these may be both the physical and logical controller
// representations.
func (volume *Volume) Controllers() ([]*StorageController, error) {
	return common.GetObjects[StorageController](volume.GetClient(), volume.controllers)
}

// DedicatedSpareDrives gets the drives which are dedicated spares for this volume.
func (volume *Volume) DedicatedSpareDrives() ([]*Drive, error) {
	return common.GetObjects[Drive](volume.GetClient(), volume.dedicatedSpareDrives)
}

// Drives references the Drives that this volume is associated with.
func (volume *Volume) Drives() ([]*Drive, error) {
	return common.GetObjects[Drive](volume.GetClient(), volume.drives)
}

// OwningStorageResource gets the Storage resource that owns or contains this volume.
func (volume *Volume) OwningStorageResource() (*Storage, error) {
	if volume.owningStorageResource == "" {
		return nil, nil
	}
	return GetStorage(volume.GetClient(), volume.owningStorageResource)
}

// // OwningStorageService gets the StorageService that owns or contains this volume.
// func (volume *Volume) OwningStorageService() (*StorageService, error) {
// 	if volume.owningStorageService == "" {
// 		return nil, nil
// 	}
// 	return GetStorageService(volume.GetClient(), volume.owningStorageService)
// }

// ServerEndpoints gets the server Endpoints this volume is associated with.
func (volume *Volume) ServerEndpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](volume.GetClient(), volume.serverEndpoints)
}

// Update commits updates to this object's properties to the running system.
func (volume *Volume) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Volume)
	original.UnmarshalJSON(volume.rawData)

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

// ListReferencedVolumes gets the collection of Volumes from a provided reference.
func ListReferencedVolumes(c common.Client, link string) ([]*Volume, error) {
	return common.GetCollectionObjects[Volume](c, link)
}

// AllowedVolumesUpdateApplyTimes returns the set of allowed apply times to request when setting the volumes values
func AllowedVolumesUpdateApplyTimes(c common.Client, link string) ([]common.OperationApplyTime, error) {
	resp, err := c.Get(link)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var temp struct {
		OperationApplyTimeSupport common.OperationApplyTimeSupport `json:"@Redfish.OperationApplyTimeSupport"`
	}

	err = json.NewDecoder(resp.Body).Decode(&temp)
	if err != nil {
		return nil, err
	}

	var applyTimes []common.OperationApplyTime
	applyTimes = append(applyTimes, temp.OperationApplyTimeSupport.SupportedValues...)
	return applyTimes, nil
}
