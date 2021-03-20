//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
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

// Volume is used to represent a volume, virtual disk, logical disk, LUN,
// or other logical storage for a Redfish implementation.
type Volume struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Status is
	Status common.Status
	// CapacityBytes shall contain the size in bytes of the associated volume.
	CapacityBytes int
	// VolumeType shall contain the type of the associated Volume.
	VolumeType VolumeType
	// Encrypted shall contain a boolean indicator if the Volume is currently
	// utilizing encryption or not.
	Encrypted bool
	// EncryptionTypes is used by this Volume.
	EncryptionTypes []EncryptionTypes
	// Identifiers shall contain a list of all known durable names for the
	// associated volume.
	Identifiers []common.Identifier
	// BlockSizeBytes shall contain size of the smallest addressable unit of the
	// associated volume.
	BlockSizeBytes int
	// Operations shall contain a list of all currently running on the Volume.
	Operations []common.Operations
	// OptimumIOSizeBytes shall contain the optimum IO size to use when
	// performing IO on this volume. For logical disks, this is the stripe size.
	// For physical disks, this describes the physical sector size.
	OptimumIOSizeBytes int
	// DrivesCount is the number of associated drives.
	DrivesCount int
	// drives contains references to associated drives.
	drives []string
}

// UnmarshalJSON unmarshals a Volume object from the raw JSON.
func (volume *Volume) UnmarshalJSON(b []byte) error {
	type temp Volume
	type links struct {
		DriveCount int `json:"Drives@odata.count"`
		Drives     common.Links
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
	volume.DrivesCount = t.DrivesCount
	volume.drives = t.Links.Drives.ToStrings()

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

// ListReferencedVolumes gets the collection of Volumes from a provided reference.
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

// Drives references the Drives that this volume is associated with.
func (volume *Volume) Drives() ([]*Drive, error) {
	var result []*Drive

	for _, driveLink := range volume.drives {
		drive, err := GetDrive(volume.Client, driveLink)
		if err != nil {
			return result, err
		}
		result = append(result, drive)
	}

	return result, nil
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
