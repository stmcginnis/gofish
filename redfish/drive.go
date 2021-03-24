//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// EncryptionAbility is the drive's encryption ability.
type EncryptionAbility string

const (

	// NoneEncryptionAbility indicates the drive is not capable of self encryption.
	NoneEncryptionAbility EncryptionAbility = "None"
	// SelfEncryptingDriveEncryptionAbility indicates the drive is capable of self
	// encryption per the Trusted Computing Group's Self Encrypting Drive
	// Standard.
	SelfEncryptingDriveEncryptionAbility EncryptionAbility = "SelfEncryptingDrive"
	// OtherEncryptionAbility indicates the drive is capable of self encryption through
	// some other means.
	OtherEncryptionAbility EncryptionAbility = "Other"
)

// EncryptionStatus is the drive's encryption state.
type EncryptionStatus string

const (
	// UnecryptedEncryptionStatus indicates the drive is not currently encrypted.
	// note: this typo occurred in the spec and was deprecated in redfish v1.1
	UnecryptedEncryptionStatus EncryptionStatus = "Unecrypted"
	// UnlockedEncryptionStatus indicates the drive is currently encrypted but the data
	// is accessible to the user unencrypted.
	UnlockedEncryptionStatus EncryptionStatus = "Unlocked"
	// LockedEncryptionStatus indicates the drive is currently encrypted and the data
	// is not accessible to the user, however the system has the ability to
	// unlock the drive automatically.
	LockedEncryptionStatus EncryptionStatus = "Locked"
	// ForeignEncryptionStatus indicates the drive is currently encrypted, the data is
	// not accessible to the user, and the system requires user intervention
	// to expose the data.
	ForeignEncryptionStatus EncryptionStatus = "Foreign"
	// UnencryptedEncryptionStatus indicates the drive is not currently encrypted.
	UnencryptedEncryptionStatus EncryptionStatus = "Unencrypted"
)

// HotspareReplacementModeType is the replacement operation mode of a hot spare.
type HotspareReplacementModeType string

const (
	// RevertibleHotspareReplacementModeType indicates the hot spare is drive that is
	// commissioned due to a drive failure will revert to being a hotspare
	// once the failed drive is replaced and rebuilt.
	RevertibleHotspareReplacementModeType HotspareReplacementModeType = "Revertible"
	// NonRevertibleHotspareReplacementModeType indicates the hot spare is drive that is
	// commissioned due to a drive failure will remain as a data drive and
	// will not revert to a hotspare if the failed drive is replaced.
	NonRevertibleHotspareReplacementModeType HotspareReplacementModeType = "NonRevertible"
)

// HotspareType is the type of hot spare.
type HotspareType string

const (
	// NoneHotspareType indicates the drive is not currently a hotspare.
	NoneHotspareType HotspareType = "None"
	// GlobalHotspareType indicates the drive is currently serving as a hotspare for
	// all other drives in the storage system.
	GlobalHotspareType HotspareType = "Global"
	// ChassisHotspareType indicates the drive is currently serving as a hotspare for
	// all other drives in the chassis.
	ChassisHotspareType HotspareType = "Chassis"
	// DedicatedHotspareType indicates the drive is currently serving as a hotspare for
	// a user defined set of drives.
	DedicatedHotspareType HotspareType = "Dedicated"
)

// MediaType is the drive's type.
type MediaType string

const (
	// HDDMediaType The drive media type is traditional magnetic platters.
	HDDMediaType MediaType = "HDD"
	// SSDMediaType The drive media type is solid state or flash memory.
	SSDMediaType MediaType = "SSD"
	// SMRMediaType The drive media type is shingled magnetic recording.
	SMRMediaType MediaType = "SMR"
)

// StatusIndicator is the drive's status.
type StatusIndicator string

const (
	// OKStatusIndicator indicates the drive is OK.
	OKStatusIndicator StatusIndicator = "OK"
	// FailStatusIndicator The drive has failed.
	FailStatusIndicator StatusIndicator = "Fail"
	// RebuildStatusIndicator indicates the drive is being rebuilt.
	RebuildStatusIndicator StatusIndicator = "Rebuild"
	// PredictiveFailureAnalysisStatusIndicator indicates the drive is still working
	// but predicted to fail soon.
	PredictiveFailureAnalysisStatusIndicator StatusIndicator = "PredictiveFailureAnalysis"
	// HotspareStatusIndicator indicates the drive is marked to be automatically
	// rebuilt and used as a replacement for a failed drive.
	HotspareStatusIndicator StatusIndicator = "Hotspare"
	// InACriticalArrayStatusIndicator The array that this drive is a part of
	// is degraded.
	InACriticalArrayStatusIndicator StatusIndicator = "InACriticalArray"
	// InAFailedArrayStatusIndicator The array that this drive is a part of
	// is failed.
	InAFailedArrayStatusIndicator StatusIndicator = "InAFailedArray"
)

// Drive is used to represent a disk drive or other physical storage
// medium for a Redfish implementation.
type Drive struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// assembly shall be a link to a resource of type Assembly.
	assembly string
	// AssetTag is used to track the drive for inventory purposes.
	AssetTag string
	// BlockSizeBytes shall contain size of the smallest addressable unit of the
	// associated drive.
	BlockSizeBytes int
	// CapableSpeedGbs shall contain fastest capable bus speed of the associated
	// drive.
	CapableSpeedGbs float32
	// CapacityBytes shall contain the raw size in bytes of the associated drive.
	CapacityBytes int64
	// Description provides a description of this resource.
	Description string
	// EncryptionAbility shall contain the encryption ability for the associated
	// drive.
	EncryptionAbility EncryptionAbility
	// EncryptionStatus shall contain the encryption status for the associated
	// drive.
	EncryptionStatus EncryptionStatus
	// FailurePredicted shall contain failure information as defined by the
	// manufacturer for the associated drive.
	FailurePredicted bool
	// HotspareReplacementMode shall specify if a commissioned hotspare will
	// continue to serve as a hotspare once the failed drive is replaced.
	HotspareReplacementMode HotspareReplacementModeType
	// HotspareType is used as part of a Volume.
	HotspareType HotspareType
	// Identifiers shall contain a list of all known durable
	// names for the associated drive.
	Identifiers []common.Identifier
	// IndicatorLED shall contain the indicator light state for the indicator
	// light associated with this drive.
	IndicatorLED common.IndicatorLED
	// Location shall contain location information of the associated drive.
	Location []common.Location
	// Manufacturer shall be the name of the organization responsible for
	// producing the drive. This organization might be the entity from whom the
	// drive is purchased, but this is not necessarily true.
	Manufacturer string
	// MediaType shall contain the type of media contained in the associated
	// drive.
	MediaType MediaType
	// Model shall be the name by which the manufacturer generally refers to the
	// drive.
	Model string
	// Multipath shall indicate whether the drive is
	// accessible by an initiator from multiple paths allowing for failover
	// capabilities upon a path failure.
	Multipath bool
	// NegotiatedSpeedGbs shall contain current bus speed of the associated
	// drive.
	NegotiatedSpeedGbs float32
	// Operations shall contain a list of all operations currently running on
	// the Drive.
	Operations []common.Operations
	// PartNumber shall be a part number assigned by the organization that is
	// responsible for producing or manufacturing the drive.
	PartNumber string
	// PhysicalLocation shall contain location information of the associated
	// drive.
	PhysicalLocation common.Location
	// PredictedMediaLifeLeftPercent shall contain an indicator of the
	// percentage of life remaining in the Drive's media.
	PredictedMediaLifeLeftPercent float32
	// Protocol shall contain the protocol the associated drive is using to
	// communicate to the storage controller for this system.
	Protocol common.Protocol
	// Revision shall contain the revision as defined by the manufacturer for
	// the associated drive.
	Revision string
	// RotationSpeedRPM shall contain rotation speed of the associated drive.
	RotationSpeedRPM float32
	// SKU shall be the stock-keeping unit number for this drive.
	SKU string
	// SerialNumber is used to identify the drive.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// StatusIndicator shall contain the status indicator state for the status
	// indicator associated with this drive. The valid values for this property
	// are specified through the Redfish.AllowableValues annotation.
	StatusIndicator StatusIndicator
	// WriteCacheEnabled shall indicate whether the drive
	// write cache is enabled.
	WriteCacheEnabled bool
	// chassis shall be a reference to a resource of type Chassis that represent
	// the physical container associated with this Drive.
	chassis string
	// endpoints shall be a reference to the resources that this drive is
	// associated with and shall reference a resource of type Endpoint.
	endpoints []string
	// EndpointsCount is the number of endpoints.
	EndpointsCount int `json:"Endpoints@odata.count"`
	// volumes are the associated volumes.
	volumes []string
	// Volumes is the number of associated volumes.
	VolumesCount int
	// pcieFunctions are the associated PCIeFunction objects.
	pcieFunctions []string
	// PCIeFunctionCount is the number of PCIeFunctions.
	PCIeFunctionCount int
	// storagePools      []string
	StoragePoolsCount int
	// secureEraseTarget is the URL for SecureErase actions.
	secureEraseTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Drive object from the raw JSON.
func (drive *Drive) UnmarshalJSON(b []byte) error {
	type temp Drive
	type links struct {
		Chassis       common.Link
		Endpoints     common.Links
		EndpointCount int `json:"Endpoints@odata.count"`
		// PCIeFunctions is The value of this property shall reference a resource
		// of type PCIeFunction that represents the PCIe functions associated
		// with this resource.
		PCIeFunctions common.Links
		// PCIeFunctions@odata.count is
		PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
		StoragePools       common.Links
		StoragePoolsCount  int `json:"StoragePools@odata.count"`
		Volumes            common.Links
		VolumeCount        int `json:"Volumes@odata.count"`
	}
	type Actions struct {
		SecureErase struct {
			Target string
		} `json:"#Drive.SecureErase"`
	}
	var t struct {
		temp
		Links    links
		Actions  Actions
		Assembly common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*drive = Drive(t.temp)
	drive.assembly = string(t.Assembly)
	drive.chassis = string(t.Links.Chassis)
	drive.endpoints = t.Links.Endpoints.ToStrings()
	drive.EndpointsCount = t.Links.EndpointCount
	drive.volumes = t.Links.Volumes.ToStrings()
	drive.VolumesCount = t.Links.VolumeCount
	drive.pcieFunctions = t.Links.PCIeFunctions.ToStrings()
	drive.PCIeFunctionCount = t.Links.PCIeFunctionsCount
	drive.secureEraseTarget = t.Actions.SecureErase.Target

	// This is a read/write object, so we need to save the raw object data for later
	drive.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (drive *Drive) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Drive)
	err := original.UnmarshalJSON(drive.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AssetTag",
		"HotspareReplacementMode",
		"IndicatorLED",
		"StatusIndicator",
		"WriteCacheEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(drive).Elem()

	return drive.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetDrive will get a Drive instance from the service.
func GetDrive(c common.Client, uri string) (*Drive, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var drive Drive
	err = json.NewDecoder(resp.Body).Decode(&drive)
	if err != nil {
		return nil, err
	}

	drive.SetClient(c)
	return &drive, nil
}

// ListReferencedDrives gets the collection of Drives from a provided reference.
func ListReferencedDrives(c common.Client, link string) ([]*Drive, error) {
	var result []*Drive
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, driveLink := range links.ItemLinks {
		drive, err := GetDrive(c, driveLink)
		if err != nil {
			return result, err
		}
		result = append(result, drive)
	}

	return result, nil
}

// Assembly gets the Assembly for this drive.
func (drive *Drive) Assembly() (*Assembly, error) {
	if drive.assembly == "" {
		return nil, nil
	}

	return GetAssembly(drive.Client, drive.assembly)
}

// Chassis gets the containing chassis for this drive.
func (drive *Drive) Chassis() (*Chassis, error) {
	if drive.chassis == "" {
		return nil, nil
	}

	return GetChassis(drive.Client, drive.chassis)
}

// Endpoints references the Endpoints that this drive is associated with.
func (drive *Drive) Endpoints() ([]*Endpoint, error) {
	var result []*Endpoint

	for _, endpointLink := range drive.endpoints {
		endpoint, err := GetEndpoint(drive.Client, endpointLink)
		if err != nil {
			return result, err
		}
		result = append(result, endpoint)
	}

	return result, nil
}

// Volumes references the Volumes that this drive is associated with.
func (drive *Drive) Volumes() ([]*Volume, error) {
	var result []*Volume

	for _, volumeLink := range drive.volumes {
		volume, err := GetVolume(drive.Client, volumeLink)
		if err != nil {
			return result, err
		}
		result = append(result, volume)
	}

	return result, nil
}

// PCIeFunctions references the PCIeFunctions that this drive is associated with.
func (drive *Drive) PCIeFunctions() ([]*PCIeFunction, error) {
	var result []*PCIeFunction

	for _, pcieFunctionLink := range drive.pcieFunctions {
		pcieFunction, err := GetPCIeFunction(drive.Client, pcieFunctionLink)
		if err != nil {
			return result, err
		}
		result = append(result, pcieFunction)
	}

	return result, nil
}

// // StoragePools references the StoragePools that this drive is associated with.
// func (drive *Drive) StoragePools() ([]*StoragePools, error) {
// 	var result []*StoragePools

// 	for _, storagePoolLink := range drive.storagePools {
// 		storagePool, err := GetStoragePools(drive.Client, storagePoolLink)
// 		if err != nil {
// 			return result, err
// 		}
// 		result = append(result, storagePool)
// 	}

// 	return result, nil
// }

// SecureErase shall perform a secure erase of the drive.
func (drive *Drive) SecureErase() error {
	_, err := drive.Client.Post(drive.secureEraseTarget, nil)
	return err
}
