// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
	"github.com/stmcginnis/gofish/school/redfish"
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

// Volume is used to represent a volume, virtual disk, logical disk, LUN,
// or other logical storage for a Redfish implementation.
type Volume struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
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
	// Description provides a description of this resource.
	Description string
	// Encrypted shall contain a boolean indicator if the
	// Volume is currently utilizing encryption or not.
	Encrypted bool
	// EncryptionTypes is used by this Volume.
	EncryptionTypes []redfish.EncryptionTypes
	// IOStatistics shall represent IO statistics for this volume.
	//IOStatistics IOStatistics
	// Identifiers shall contain a list of all known durable
	// names for the associated volume.
	Identifiers []common.Identifier
	// Links is The Links property, as described by the Redfish
	// Specification, shall contain references to resources that are related
	// to, but not contained by (subordinate to), this resource.
	Links string
	// LowSpaceWarningThresholdPercents is Each time the following value is
	// less than one of the values in the array the
	// LOW_SPACE_THRESHOLD_WARNING event shall be triggered: Across all
	// CapacitySources entries, percent = (SUM(AllocatedBytes) -
	// SUM(ConsumedBytes))/SUM(AllocatedBytes).
	LowSpaceWarningThresholdPercents []string
	// Manufacturer shall contain a value that represents
	// the manufacturer or implementer of the storage volume.
	Manufacturer string
	// MaxBlockSizeBytes shall contain size of the largest
	// addressable unit of this storage volume.
	MaxBlockSizeBytes int
	// Model is The value is assigned by the manufacturer and shall
	// represents a specific storage volume implementation.
	Model string
	// Oem is The value of this string shall be of the format for the
	// reserved word *Oem*.
	OEM string `json:"Oem"`
	// Operations shall contain a list of all currently
	// running on the Volume.
	Operations []common.Operations
	// OptimumIOSizeBytes shall contain the optimum IO size
	// to use when performing IO on this volume. For logical disks, this is
	// the stripe size. For physical disks, this describes the physical
	// sector size.
	OptimumIOSizeBytes int
	// RAIDType shall contain the RAID type of the
	// associated Volume.
	RAIDType RAIDType
	// RecoverableCapacitySourceCount is The value is the number of available
	// capacity source resources currently available in the event that an
	// equivalent capacity source resource fails.
	RecoverableCapacitySourceCount int
	// RemainingCapacityPercent is If present, this value shall return
	// {[(SUM(AllocatedBytes) - SUM(ConsumedBytes)]/SUM(AllocatedBytes)}*100
	// represented as an integer value.
	RemainingCapacityPercent int
	// ReplicaInfo shall describe the replica relationship
	// between this storage volume and a corresponding source volume.
	//ReplicaInfo redfish.ReplicaInfo
	// ReplicaTargets shall reference the target replicas that
	// are sourced by this replica.
	ReplicaTargets []string
	// ReplicaTargets@odata.count is
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// Status is
	Status common.Status
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
	// drives shall be a reference to the resources that this volume is
	// associated with and shall reference resources of type Drive. This
	// property shall only contain references to Drive entities which are
	// currently members of the Volume, not hot spare Drives which are not
	// currently a member of the volume.
	drives []string
	// SpareResourceSets referenced SpareResourceSet shall contain
	// resources that may be utilized to replace the capacity provided by a
	// failed resource having a compatible type.
	spareResourceSets string
	// allocatedPools shall contain references
	// to all storage pools allocated from this volume.
	allocatedPools string
	// storageGroups shall contain references to all storage groups that include
	// this volume.
	storageGroups string
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
		SpareResourceSets common.Link
		// SpareResourceSets@odata.count is
		SpareResourceSetsCount int `json:"SpareResourceSets@odata.count"`
	}

	var t struct {
		temp
		Links links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*volume = Volume(t.temp)

	// Extract the links to other entities for later
	volume.classOfService = string(t.Links.ClassOfService)
	volume.dedicatedSpareDrives = t.Links.DedicatedSpareDrives.ToStrings()
	volume.drives = t.Links.Drives.ToStrings()
	volume.spareResourceSets = string(t.Links.SpareResourceSets)
	volume.DedicatedSpareDrivesCount = t.Links.DedicatedSpareDrivesCount
	volume.DrivesCount = t.Links.DrivesCount
	volume.SpareResourceSetsCount = t.Links.SpareResourceSetsCount

	return nil
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
	return ListReferencedSpareResourceSets(volume.Client, volume.spareResourceSets)
}
