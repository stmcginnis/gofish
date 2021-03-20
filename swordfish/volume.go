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
	// block device for write-ahead logging to adddress write hole issue. All
	// write operations on the RAID volume are first logged on dedicated
	// journaling device that is not part of the volume.
	JournalingWriteHoleProtectionPolicyType WriteHoleProtectionPolicyType = "Journaling"
	// DistributedLogWriteHoleProtectionPolicyType The policy that
	// distributes additional log (e.q. cheksum of the parity) among the
	// volume's capacity sources to address write hole issue. Additional data
	// is used to detect data corruption on the volume.
	DistributedLogWriteHoleProtectionPolicyType WriteHoleProtectionPolicyType = "DistributedLog"
	// OEMWriteHoleProtectionPolicyType The policy that is Oem specific. The
	// mechanism details are unknown unless provided separatly by the Oem.
	OEMWriteHoleProtectionPolicyType WriteHoleProtectionPolicyType = "Oem"
)

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
	// BlockSizeBytes shall contain size of the smallest
	// addressable unit of the associated volume.
	BlockSizeBytes int
	// Capacity is Information about the utilization of capacity allocated to
	// this storage volume.
	Capacity Capacity
	// CapacityBytes shall contain the size in bytes of the
	// associated volume.
	CapacityBytes int
	// CapacitySources is Fully or partially consumed storage from a source
	// resource. Each entry provides capacity allocation information from a
	// named source resource.
	CapacitySources []CapacitySource
	// CapacitySources@odata.count is
	CapacitySourcesCount int `json:"CapacitySources@odata.count"`
	// Compressed shall contain a boolean indicator if the Volume is currently
	// utilizing compression or not.
	Compressed bool
	// Deduplicated shall contain a boolean indicator if the Volume is currently
	// utilizing deduplication or not.
	Deduplicated bool
	// Description provides a description of this resource.
	Description string
	// Encrypted shall contain a boolean indicator if the
	// Volume is currently utilizing encryption or not.
	Encrypted bool
	// EncryptionTypes is used by this Volume.
	EncryptionTypes []redfish.EncryptionTypes
	// IOStatistics shall represent IO statistics for this volume.
	// IOStatistics IOStatistics
	// Identifiers shall contain a list of all known durable
	// names for the associated volume.
	Identifiers []common.Identifier
	// Links is The Links property, as described by the Redfish
	// Specification, shall contain references to resources that are related
	// to, but not contained by (subordinate to), this resource.
	Links string
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
	// Model is The value is assigned by the manufacturer and shall
	// represents a specific storage volume implementation.
	Model string
	// Operations shall contain a list of all currently
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
	// ReplicaInfo shall describe the replica relationship
	// between this storage volume and a corresponding source volume.
	// ReplicaInfo redfish.ReplicaInfo
	// ReplicaTargets shall reference the target replicas that
	// are sourced by this replica.
	ReplicaTargets []string
	// ReplicaTargets@odata.count is
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// Status is
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
	// classOfService shall contain a reference to the
	// ClassOfService that this storage volume conforms to.
	classOfService string
	// DedicatedSpareDrivesCount is the number of dedicates spare drives
	DedicatedSpareDrivesCount int
	// DrivesCount is the number of associated drives.
	DrivesCount int
	// SpareResourceSetsCount is the number of spare resources sets.
	SpareResourceSetsCount int
	// dedicatedSpareDrives shall be a reference to the resources that this
	// volume is associated with and shall reference resources of type Drive.
	// This property shall only contain references to Drive entities which are
	// currently assigned as a dedicated spare and are able to support this Volume.
	dedicatedSpareDrives []string
	// DisplayName shall contain a user-configurable string to name the volume.
	DisplayName string
	// drives shall be a reference to the resources that this volume is
	// associated with and shall reference resources of type Drive. This
	// property shall only contain references to Drive entities which are
	// currently members of the Volume, not hot spare Drives which are not
	// currently a member of the volume.
	drives []string
	// SpareResourceSets referenced SpareResourceSet shall contain
	// resources that may be utilized to replace the capacity provided by a
	// failed resource having a compatible type.
	spareResourceSets []string
	// allocatedPools shall contain references to all storage pools allocated
	// from this volume.
	allocatedPools []string
	// storageGroups shall contain references to all storage groups that include
	// this volume.
	storageGroups []string
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
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Volume object from the raw JSON.
func (volume *Volume) UnmarshalJSON(b []byte) error {
	type temp Volume

	type links struct {
		// ClassOfService shall contain a reference to the
		// ClassOfService that this storage volume conforms to.
		ClassOfService common.Link
		// DedicatedSpareDrives shall be a
		// reference to the resources that this volume is associated with and
		// shall reference resources of type Drive. This property shall only
		// contain references to Drive entities which are currently assigned as a
		// dedicated spare and are able to support this Volume.
		DedicatedSpareDrives common.Links
		// DedicatedSpareDrives@odata.count is
		DedicatedSpareDrivesCount int `json:"DedicatedSpareDrives@odata.count"`
		// Drives shall be a reference to the
		// resources that this volume is associated with and shall reference
		// resources of type Drive. This property shall only contain references
		// to Drive entities which are currently members of the Volume, not hot
		// spare Drives which are not currently a member of the volume.
		Drives common.Links
		// Drives@odata.count is
		DrivesCount int `json:"Drives@odata.count"`
		// SpareResourceSets is Each referenced SpareResourceSet shall contain
		// resources that may be utilized to replace the capacity provided by a
		// failed resource having a compatible type.
		SpareResourceSets common.Links
		// SpareResourceSets@odata.count is
		SpareResourceSetsCount int `json:"SpareResourceSets@odata.count"`
	}

	type actions struct {
		AssignReplicaTarget struct {
			Target string
		} `json:"#Volume.AssignReplicaTarget"`
		CheckConsistency struct {
			Target string
		} `json:"#Volume.CheckConsistency"`
		CreateReplicaTarget struct {
			Target string
		} `json:"#Volume.CreateReplicaTarget"`
		Initialize struct {
			Target string
		} `json:"#Volume.Initialize"`
		RemoveReplicaRelationship struct {
			Target string
		} `json:"#Volume.RemoveReplicaRelationship"`
		ResumeReplication struct {
			Target string
		} `json:"#Volume.ResumeReplication"`
		ReverseReplicationRelationship struct {
			Target string
		} `json:"#Volume.ReverseReplicationRelationship"`
		SplitReplication struct {
			Target string
		} `json:"#Volume.SplitReplication"`
		SuspendReplication struct {
			Target string
		} `json:"#Volume.SuspendReplication"`
	}

	var t struct {
		temp
		AllocatedPools common.Links
		StorageGroups  common.Links
		Links          links
		Actions        actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*volume = Volume(t.temp)
	volume.allocatedPools = t.AllocatedPools.ToStrings()
	volume.storageGroups = t.StorageGroups.ToStrings()
	volume.classOfService = string(t.Links.ClassOfService)
	volume.dedicatedSpareDrives = t.Links.DedicatedSpareDrives.ToStrings()
	volume.drives = t.Links.Drives.ToStrings()
	volume.spareResourceSets = t.Links.SpareResourceSets.ToStrings()
	volume.DedicatedSpareDrivesCount = t.Links.DedicatedSpareDrivesCount
	volume.DrivesCount = t.Links.DrivesCount
	volume.SpareResourceSetsCount = t.Links.SpareResourceSetsCount
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
		"LowSpaceWarningThresholdPercents",
		"ProvisioningPolicy",
		"ReadCachePolicy",
		"RecoverableCapacitySourceCount",
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
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var volume Volume
	err = json.NewDecoder(resp.Body).Decode(&volume)
	if err != nil {
		return nil, err
	}

	volume.SetClient(c)
	return &volume, nil
}

// ListReferencedVolumes gets the collection of Volume from a provided reference.
func ListReferencedVolumes(c common.Client, link string) ([]*Volume, error) {
	var result []*Volume
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, volumeLink := range links.ItemLinks {
		volume, err := GetVolume(c, volumeLink)
		if err != nil {
			return result, err
		}
		result = append(result, volume)
	}

	return result, nil
}

// ClassOfService gets the class of service that this storage volume conforms to.
func (volume *Volume) ClassOfService() (*ClassOfService, error) {
	if volume.classOfService == "" {
		return nil, nil
	}

	return GetClassOfService(volume.Client, volume.classOfService)
}

// getDrives gets a set of referenced drives.
func (volume *Volume) getDrives(links []string) ([]*redfish.Drive, error) {
	var result []*redfish.Drive

	for _, driveLink := range links {
		drive, err := redfish.GetDrive(volume.Client, driveLink)
		if err != nil {
			return result, err
		}
		result = append(result, drive)
	}

	return result, nil
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

// SpareResourceSets gets the spare resources that can be used for this volume.
func (volume *Volume) SpareResourceSets() ([]*SpareResourceSet, error) {
	var result []*SpareResourceSet
	for _, srsLink := range volume.spareResourceSets {
		srs, err := GetSpareResourceSet(volume.Client, srsLink)
		if err != nil {
			return result, err
		}
		result = append(result, srs)
	}

	return result, nil
}

// StorageGroups gets the storage groups that associated with this volume.
func (volume *Volume) StorageGroups() ([]*StorageGroup, error) {
	var result []*StorageGroup
	for _, sgLink := range volume.storageGroups {
		sg, err := GetStorageGroup(volume.Client, sgLink)
		if err != nil {
			return result, err
		}
		result = append(result, sg)
	}

	return result, nil
}

// StoragePools gets the storage pools that associated with this volume.
func (volume *Volume) StoragePools() ([]*StoragePool, error) {
	var result []*StoragePool
	for _, sgLink := range volume.allocatedPools {
		sg, err := GetStoragePool(volume.Client, sgLink)
		if err != nil {
			return result, err
		}
		result = append(result, sg)
	}

	return result, nil
}

// AssignReplicaTarget is used to establish a replication relationship by
// assigning an existing volume to serve as a target replica for an existing
// source volume.
func (volume *Volume) AssignReplicaTarget(replicaType ReplicaType, updateMode ReplicaUpdateMode, targetVolumeODataID string) error {
	// This action wasn't added until later revisions
	if volume.assignReplicaTargetTarget == "" {
		return fmt.Errorf("AssignReplicaTarget action is not supported by this system") // nolint
	}

	// Define this action's parameters
	type temp struct {
		ReplicaType       ReplicaType
		ReplicaUpdateMode ReplicaUpdateMode
		TargetVolume      string
	}

	// Set the values for the action arguments
	t := temp{
		ReplicaType:       replicaType,
		ReplicaUpdateMode: updateMode,
		TargetVolume:      targetVolumeODataID,
	}

	_, err := volume.Client.Post(volume.assignReplicaTargetTarget, t)
	return err
}

// CheckConsistency is used to force a check of the Volume's parity or redundant
// data to ensure it matches calculated values.
func (volume *Volume) CheckConsistency() error {
	if volume.checkConsistencyTarget == "" {
		return fmt.Errorf("CheckConsistency action is not supported by this system") // nolint
	}

	_, err := volume.Client.Post(volume.checkConsistencyTarget, nil)
	return err
}

// Initialize is used to prepare the contents of the volume for use by the system.
func (volume *Volume) Initialize(initType InitializeType) error {
	if volume.initializeTarget == "" {
		return fmt.Errorf("initialize action is not supported by this system")
	}

	// Define this action's parameters
	type temp struct {
		InitializeType InitializeType
	}

	// Set the values for the action arguments
	t := temp{InitializeType: initType}

	_, err := volume.Client.Post(volume.initializeTarget, t)
	return err
}

// RemoveReplicaRelationship is used to disable data synchronization between a
// source and target volume, remove the replication relationship, and optionally
// delete the target volume.
func (volume *Volume) RemoveReplicaRelationship(deleteTarget bool, targetVolumeODataID string) error {
	// This action wasn't added until later revisions
	if volume.removeReplicaRelationshipTarget == "" {
		return fmt.Errorf("RemoveReplicaRelationship action is not supported by this system") // nolint
	}

	// Define this action's parameters
	type temp struct {
		DeleteTargetVolume bool
		TargetVolume       string
	}

	// Set the values for the action arguments
	t := temp{
		DeleteTargetVolume: deleteTarget,
		TargetVolume:       targetVolumeODataID,
	}

	_, err := volume.Client.Post(volume.removeReplicaRelationshipTarget, t)
	return err
}

// ResumeReplication is used to resume the active data synchronization between a
// source and target volume, without otherwise altering the replication
// relationship.
func (volume *Volume) ResumeReplication(targetVolumeODataID string) error {
	// This action wasn't added until later revisions
	if volume.resumeReplicationTarget == "" {
		return fmt.Errorf("ResumeReplication action is not supported by this system") // nolint
	}

	// Define this action's parameters
	type temp struct {
		TargetVolume string
	}

	// Set the values for the action arguments
	t := temp{TargetVolume: targetVolumeODataID}

	_, err := volume.Client.Post(volume.resumeReplicationTarget, t)
	return err
}

// ReverseReplicationRelationship is used to reverse the replication
// relationship between a source and target volume.
func (volume *Volume) ReverseReplicationRelationship(targetVolumeODataID string) error {
	// This action wasn't added until later revisions
	if volume.reverseReplicationRelationshipTarget == "" {
		return fmt.Errorf("ReverseReplicationRelationship action is not supported by this system") // nolint
	}

	// Define this action's parameters
	type temp struct {
		TargetVolume string
	}

	// Set the values for the action arguments
	t := temp{TargetVolume: targetVolumeODataID}

	_, err := volume.Client.Post(volume.reverseReplicationRelationshipTarget, t)
	return err
}

// SplitReplication is used to split the replication relationship and suspend
// data synchronization between a source and target volume.
func (volume *Volume) SplitReplication(targetVolumeODataID string) error {
	// This action wasn't added until later revisions
	if volume.splitReplicationTarget == "" {
		return fmt.Errorf("SplitReplication action is not supported by this system") // nolint
	}

	// Define this action's parameters
	type temp struct {
		TargetVolume string
	}

	// Set the values for the action arguments
	t := temp{TargetVolume: targetVolumeODataID}

	_, err := volume.Client.Post(volume.splitReplicationTarget, t)
	return err
}

// SuspendReplication is used to suspend active data synchronization between a
// source and target volume, without otherwise altering the replication
// relationship.
func (volume *Volume) SuspendReplication(targetVolumeODataID string) error {
	// This action wasn't added until later revisions
	if volume.suspendReplicationTarget == "" {
		return fmt.Errorf("SuspendReplication action is not supported by this system") // nolint
	}

	// Define this action's parameters
	type temp struct {
		TargetVolume string
	}

	// Set the values for the action arguments
	t := temp{TargetVolume: targetVolumeODataID}

	_, err := volume.Client.Post(volume.suspendReplicationTarget, t)
	return err
}
