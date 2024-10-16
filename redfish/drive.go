//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type DataSanitizationType string

const (
	// BlockEraseDataSanitizationType shall indicate sanitization is performed by deleting all logical block addresses,
	// including those that are not currently mapping to active addresses, but leaving the data on the drive.
	BlockEraseDataSanitizationType DataSanitizationType = "BlockErase"
	// CryptographicEraseDataSanitizationType shall indicate sanitization is performed by erasing the target data's
	// encryption key leaving only the ciphertext on the drive. For more information, see NIST800-88 and ISO/IEC 27040.
	CryptographicEraseDataSanitizationType DataSanitizationType = "CryptographicErase"
	// OverwriteDataSanitizationType shall indicate sanitization is performed by overwriting data by writing an
	// implementation-specific pattern onto all sectors of the drive.
	OverwriteDataSanitizationType DataSanitizationType = "Overwrite"
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

type FormFactor string

const (
	// Drive35FormFactor shall indicate the drive is approximately 3.5 inches in width and no more than 1.1 inches in
	// height.
	Drive35FormFactor FormFactor = "Drive3_5"
	// Drive25FormFactor shall indicate the drive is approximately 2.5 inches in width and no more than 0.8 inches in
	// height and is not a U.2 drive.
	Drive25FormFactor FormFactor = "Drive2_5"
	// EDSFFFormFactor shall indicate the drive corresponds to one of SNIA's SFF specifications with an unspecified
	// form factor. The SlotFormFactor property should not contain this value.
	EDSFFFormFactor FormFactor = "EDSFF"
	// EDSFF1ULongFormFactor shall indicate the drive corresponds to the SFF-TA-1007 Specification.
	EDSFF1ULongFormFactor FormFactor = "EDSFF_1U_Long"
	// EDSFF1UShortFormFactor shall indicate the drive corresponds to the SFF-TA-1006 Specification.
	EDSFF1UShortFormFactor FormFactor = "EDSFF_1U_Short"
	// EDSFFE3ShortFormFactor shall indicate the drive corresponds to the SFF-TA-1008 Specification and is
	// approximately 112.75 mm in length.
	EDSFFE3ShortFormFactor FormFactor = "EDSFF_E3_Short"
	// EDSFFE3LongFormFactor shall indicate the drive corresponds to the SFF-TA-1008 Specification and is approximately
	// 142.2 mm in length.
	EDSFFE3LongFormFactor FormFactor = "EDSFF_E3_Long"
	// M2FormFactor shall indicate the drive corresponds to the PCI Express M.2 Specification with an unspecified form
	// factor. The SlotFormFactor property should not contain this value.
	M2FormFactor FormFactor = "M2"
	// M22230FormFactor shall indicate the drive corresponds to the PCI Express M.2 Specification and is approximately
	// 22 mm in width and 30 mm in length.
	M22230FormFactor FormFactor = "M2_2230"
	// M22242FormFactor shall indicate the drive corresponds to the PCI Express M.2 Specification and is approximately
	// 22 mm in width and 42 mm in length.
	M22242FormFactor FormFactor = "M2_2242"
	// M22260FormFactor shall indicate the drive corresponds to the PCI Express M.2 Specification and is approximately
	// 22 mm in width and 60 mm in length.
	M22260FormFactor FormFactor = "M2_2260"
	// M22280FormFactor shall indicate the drive corresponds to the PCI Express M.2 Specification and is approximately
	// 22 mm in width and 80 mm in length.
	M22280FormFactor FormFactor = "M2_2280"
	// M222110FormFactor shall indicate the drive corresponds to the PCI Express M.2 Specification and is approximately
	// 22 mm in width and 110 mm in length.
	M222110FormFactor FormFactor = "M2_22110"
	// U2FormFactor shall indicate the drive corresponds to the PCI Express SFF-8639 Module Specification.
	U2FormFactor FormFactor = "U2"
	// PCIeSlotFullLengthFormFactor shall indicate the drive is an add-in card greater than 7 inches in length.
	PCIeSlotFullLengthFormFactor FormFactor = "PCIeSlotFullLength"
	// PCIeSlotLowProfileFormFactor shall indicate the drive is an add-in card less than 2.5 inches in height.
	PCIeSlotLowProfileFormFactor FormFactor = "PCIeSlotLowProfile"
	// PCIeHalfLengthFormFactor shall indicate the drive is an add-in card less than 7 inches in length.
	PCIeHalfLengthFormFactor FormFactor = "PCIeHalfLength"
	// OEMFormFactor shall indicate the drive is an OEM-defined form factor.
	OEMFormFactor FormFactor = "OEM"
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
	// DriveFormFactor shall contain the form factor of the drive inserted in this slot.
	// Added in v1.18.0.
	DriveFormFactor FormFactor
	// EncryptionAbility shall contain the encryption ability for the associated
	// drive.
	EncryptionAbility EncryptionAbility
	// EncryptionStatus shall contain the encryption status for the associated
	// drive.
	EncryptionStatus EncryptionStatus
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this drive.
	environmentMetrics string
	// FailurePredicted shall contain failure information as defined by the
	// manufacturer for the associated drive.
	FailurePredicted bool
	// FirmwareVersion shall contain the firmware version as defined by the manufacturer for this drive.
	FirmwareVersion string
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
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource. A write to this property shall update the value of IndicatorLED in this resource, if supported, to
	// reflect the implementation of the locating function.
	LocationIndicatorActive bool
	// Manufacturer shall be the name of the organization responsible for
	// producing the drive. This organization might be the entity from whom the
	// drive is purchased, but this is not necessarily true.
	Manufacturer string
	// MediaType shall contain the type of media contained in the associated
	// drive.
	MediaType MediaType
	// Metrics shall contain a link to the metrics associated with this drive.
	Metrics DriveMetrics
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
	// ReadyToRemove shall indicate whether the system is prepared for the removal of this drive.
	ReadyToRemove bool
	// Revision shall contain the revision as defined by the manufacturer for
	// the associated drive.
	Revision string
	// RotationSpeedRPM shall contain rotation speed of the associated drive.
	RotationSpeedRPM float32
	// SKU shall be the stock-keeping unit number for this drive.
	SKU string
	// SerialNumber is used to identify the drive.
	SerialNumber string
	// SlotCapableProtocols shall contain the drive protocols capable in this slot. The value of this property depends
	// upon the connector in this slot, the storage controllers connected to this slot, the configuration of the
	// system, and other constraints that determine if a particular protocol is capable at a given time.
	SlotCapableProtocols []common.Protocol
	// SlotFormFactor shall contain the form factor of the slot.
	SlotFormFactor FormFactor
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// StatusIndicator shall contain the status indicator state for the status
	// indicator associated with this drive. The valid values for this property
	// are specified through the Redfish.AllowableValues annotation.
	StatusIndicator StatusIndicator
	// WriteCacheEnabled shall indicate whether the drive
	// write cache is enabled.
	WriteCacheEnabled bool
	// Oem is all the available OEM information for the drive
	Oem json.RawMessage

	// ActiveSoftwareImage shall contain a link a resource of type SoftwareInventory that represents the active drive
	// firmware image.
	activeSoftwareImage string
	// chassis shall be a reference to a resource of type Chassis that represent
	// the physical container associated with this Drive.
	chassis string
	// endpoints shall be a reference to the resources that this drive is
	// associated with and shall reference a resource of type Endpoint.
	endpoints []string
	// EndpointsCount is the number of endpoints.
	EndpointsCount         int `json:"Endpoints@odata.count"`
	networkDeviceFunctions []string
	// NetworkDeviceFunctionsCount is the number of network device functions related to this drive.
	NetworkDeviceFunctionsCount int
	// volumes are the associated volumes.
	volumes []string
	// Volumes is the number of associated volumes.
	VolumesCount int
	// pcieFunctions are the associated PCIeFunction objects.
	pcieFunctions []string
	// PCIeFunctionCount is the number of PCIeFunctions.
	PCIeFunctionCount int
	softwareImages    []string
	// SoftwareImagesCount is the number of software images related to this drive.
	SoftwareImagesCount int
	storage             string
	storagePools        []string
	// storagePools      []string
	StoragePoolsCount int
	// secureEraseTarget is the URL for SecureErase actions.
	secureEraseTarget string
	// RawData holds the original serialized JSON so we can compare updates
	// as well as access Oem values in the oem package.
	RawData []byte
}

// UnmarshalJSON unmarshals a Drive object from the raw JSON.
func (drive *Drive) UnmarshalJSON(b []byte) error {
	type temp Drive
	type links struct {
		ActiveSoftwareImage common.Link
		Chassis             common.Link
		Endpoints           common.Links
		EndpointCount       int `json:"Endpoints@odata.count"`
		// NetworkDeviceFunctions shall contain the array of links to resources of type NetworkDeviceFunction. This
		// property should only be present for drives with network connectivity, such as Ethernet attached drives.
		NetworkDeviceFunctions common.Links
		// NetworkDeviceFunctions@odata.count
		NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
		// PCIeFunctions shall reference a resource of type PCIeFunction that represents the PCIe functions associated
		// with this resource.
		PCIeFunctions      common.Links
		PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
		// SoftwareImages shall contain an array of links to resources of type SoftwareInventory that represent the
		// firmware images that apply to this drive.
		SoftwareImages      common.Links
		SoftwareImagesCount int `json:"SoftwareImages@odata.count"`
		// Storage shall contain a link to a resource of type Storage that represents the storage subsystem to which this
		// drive belongs.
		Storage           common.Link
		StoragePools      common.Links
		StoragePoolsCount int `json:"StoragePools@odata.count"`
		Volumes           common.Links
		VolumeCount       int `json:"Volumes@odata.count"`
	}
	type Actions struct {
		SecureErase common.ActionTarget `json:"#Drive.SecureErase"`
	}
	var t struct {
		temp
		Links              links
		Actions            Actions
		Assembly           common.Link
		EnvironmentMetrics common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*drive = Drive(t.temp)
	drive.assembly = t.Assembly.String()
	drive.environmentMetrics = t.EnvironmentMetrics.String()

	drive.activeSoftwareImage = t.Links.ActiveSoftwareImage.String()
	drive.chassis = t.Links.Chassis.String()
	drive.endpoints = t.Links.Endpoints.ToStrings()
	drive.EndpointsCount = t.Links.EndpointCount
	drive.networkDeviceFunctions = t.Links.NetworkDeviceFunctions.ToStrings()
	drive.NetworkDeviceFunctionsCount = t.Links.NetworkDeviceFunctionsCount
	drive.Oem = t.Oem
	drive.pcieFunctions = t.Links.PCIeFunctions.ToStrings()
	drive.PCIeFunctionCount = t.Links.PCIeFunctionsCount
	drive.softwareImages = t.Links.SoftwareImages.ToStrings()
	drive.SoftwareImagesCount = t.Links.SoftwareImagesCount
	drive.storage = t.Links.Storage.String()
	drive.storagePools = t.Links.StoragePools.ToStrings()
	drive.StoragePoolsCount = t.Links.StoragePoolsCount
	drive.volumes = t.Links.Volumes.ToStrings()
	drive.VolumesCount = t.Links.VolumeCount

	drive.secureEraseTarget = t.Actions.SecureErase.Target

	// This is a read/write object, so we need to save the raw object data for later
	drive.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (drive *Drive) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Drive)
	err := original.UnmarshalJSON(drive.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AssetTag",
		"HotspareReplacementMode",
		"HotspareType",
		"LocationIndicatorActive",
		"ReadyToRemove",
		"StatusIndicator",
		"WriteCacheEnabled",
		"IndicatorLED",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(drive).Elem()

	return drive.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetDrive will get a Drive instance from the service.
func GetDrive(c common.Client, uri string) (*Drive, error) {
	return common.GetObject[Drive](c, uri)
}

// ListReferencedDrives gets the collection of Drives from a provided reference.
func ListReferencedDrives(c common.Client, link string) ([]*Drive, error) {
	return common.GetCollectionObjects[Drive](c, link)
}

// Assembly gets the Assembly for this drive.
func (drive *Drive) Assembly() (*Assembly, error) {
	if drive.assembly == "" {
		return nil, nil
	}

	return GetAssembly(drive.GetClient(), drive.assembly)
}

// Chassis gets the containing chassis for this drive.
func (drive *Drive) Chassis() (*Chassis, error) {
	if drive.chassis == "" {
		return nil, nil
	}

	return GetChassis(drive.GetClient(), drive.chassis)
}

// Endpoints references the Endpoints that this drive is associated with.
func (drive *Drive) Endpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](drive.GetClient(), drive.endpoints)
}

// EnvironmentMetrics gets the environment metrics for this drive.
// If no metrics are available the EnvironmentMetrics reference will be nil but
// no error will be returned unless it was due to a problem fetching the data.
func (drive *Drive) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if drive.environmentMetrics == "" {
		return nil, nil
	}

	return GetEnvironmentMetrics(drive.GetClient(), drive.environmentMetrics)
}

// Volumes references the Volumes that this drive is associated with.
func (drive *Drive) Volumes() ([]*Volume, error) {
	return common.GetObjects[Volume](drive.GetClient(), drive.volumes)
}

// PCIeFunctions references the PCIeFunctions that this drive is associated with.
func (drive *Drive) PCIeFunctions() ([]*PCIeFunction, error) {
	return common.GetObjects[PCIeFunction](drive.GetClient(), drive.pcieFunctions)
}

// // StoragePools references the StoragePools that this drive is associated with.
// func (drive *Drive) StoragePools() ([]*StoragePools, error) {
//	return common.GetObjects[StoragePools](drive.GetClient(), drive.storagePools)
// }

// SecureErase shall perform a secure erase of the drive.
func (drive *Drive) SecureErase() error {
	return drive.Post(drive.secureEraseTarget, nil)
}
