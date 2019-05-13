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

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
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
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
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
