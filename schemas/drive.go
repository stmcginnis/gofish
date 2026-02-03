//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.4 - #Drive.v1_22_0.Drive

package schemas

import (
	"encoding/json"
)

type DataSanitizationType string

const (
	// BlockEraseDataSanitizationType shall indicate sanitization is performed by
	// deleting all logical block addresses, including those that are not currently
	// mapping to active addresses, but leaving the data on the drive.
	BlockEraseDataSanitizationType DataSanitizationType = "BlockErase"
	// CryptographicEraseDataSanitizationType shall indicate sanitization is
	// performed by erasing the target data's encryption key leaving only the
	// ciphertext on the drive. For more information, see NIST800-88 and ISO/IEC
	// 27040.
	CryptographicEraseDataSanitizationType DataSanitizationType = "CryptographicErase"
	// OverwriteDataSanitizationType shall indicate sanitization is performed by
	// overwriting data by writing an implementation-specific pattern onto all
	// sectors of the drive.
	OverwriteDataSanitizationType DataSanitizationType = "Overwrite"
)

type EncryptionAbility string

const (
	// NoneEncryptionAbility The drive is not capable of self-encryption.
	NoneEncryptionAbility EncryptionAbility = "None"
	// SelfEncryptingDriveEncryptionAbility The drive is capable of self-encryption
	// per the Trusted Computing Group's Self Encrypting Drive Standard.
	SelfEncryptingDriveEncryptionAbility EncryptionAbility = "SelfEncryptingDrive"
	// OtherEncryptionAbility The drive is capable of self-encryption through some
	// other means.
	OtherEncryptionAbility EncryptionAbility = "Other"
)

type EncryptionStatus string

const (
	// UnecryptedEncryptionStatus The drive is not currently encrypted.
	UnecryptedEncryptionStatus EncryptionStatus = "Unecrypted"
	// UnlockedEncryptionStatus The drive is currently encrypted but the data is
	// accessible to the user in unencrypted form.
	UnlockedEncryptionStatus EncryptionStatus = "Unlocked"
	// LockedEncryptionStatus The drive is currently encrypted and the data is not
	// accessible to the user. However, the system can unlock the drive
	// automatically.
	LockedEncryptionStatus EncryptionStatus = "Locked"
	// ForeignEncryptionStatus The drive is currently encrypted, the data is not
	// accessible to the user, and the system requires user intervention to expose
	// the data.
	ForeignEncryptionStatus EncryptionStatus = "Foreign"
	// UnencryptedEncryptionStatus The drive is not currently encrypted.
	UnencryptedEncryptionStatus EncryptionStatus = "Unencrypted"
)

type FormFactor string

const (
	// Drive35FormFactor shall indicate the drive is approximately 3.5 inches in
	// width and no more than 1.1 inches in height.
	Drive35FormFactor FormFactor = "Drive3_5"
	// Drive25FormFactor shall indicate the drive is approximately 2.5 inches in
	// width and no more than 0.8 inches in height and is not a U.2 drive.
	Drive25FormFactor FormFactor = "Drive2_5"
	// EDSFFFormFactor shall indicate the drive corresponds to one of SNIA's SFF
	// specifications with an unspecified form factor. The 'SlotFormFactor'
	// property should not contain this value.
	EDSFFFormFactor FormFactor = "EDSFF"
	// EDSFF1ULongFormFactor shall indicate the drive corresponds to the
	// SFF-TA-1007 Specification.
	EDSFF1ULongFormFactor FormFactor = "EDSFF_1U_Long"
	// EDSFF1UShortFormFactor shall indicate the drive corresponds to the
	// SFF-TA-1006 Specification.
	EDSFF1UShortFormFactor FormFactor = "EDSFF_1U_Short"
	// EDSFFE3ShortFormFactor shall indicate the drive corresponds to the
	// SFF-TA-1008 Specification and is approximately 112.75 mm in length.
	EDSFFE3ShortFormFactor FormFactor = "EDSFF_E3_Short"
	// EDSFFE3LongFormFactor shall indicate the drive corresponds to the
	// SFF-TA-1008 Specification and is approximately 142.2 mm in length.
	EDSFFE3LongFormFactor FormFactor = "EDSFF_E3_Long"
	// M2FormFactor shall indicate the drive corresponds to the PCI Express M.2
	// Specification with an unspecified form factor. The 'SlotFormFactor' property
	// should not contain this value.
	M2FormFactor FormFactor = "M2"
	// M22230FormFactor shall indicate the drive corresponds to the PCI Express M.2
	// Specification and is approximately 22 mm in width and 30 mm in length.
	M22230FormFactor FormFactor = "M2_2230"
	// M22242FormFactor shall indicate the drive corresponds to the PCI Express M.2
	// Specification and is approximately 22 mm in width and 42 mm in length.
	M22242FormFactor FormFactor = "M2_2242"
	// M22260FormFactor shall indicate the drive corresponds to the PCI Express M.2
	// Specification and is approximately 22 mm in width and 60 mm in length.
	M22260FormFactor FormFactor = "M2_2260"
	// M22280FormFactor shall indicate the drive corresponds to the PCI Express M.2
	// Specification and is approximately 22 mm in width and 80 mm in length.
	M22280FormFactor FormFactor = "M2_2280"
	// M222110FormFactor shall indicate the drive corresponds to the PCI Express
	// M.2 Specification and is approximately 22 mm in width and 110 mm in length.
	M222110FormFactor FormFactor = "M2_22110"
	// U2FormFactor shall indicate the drive corresponds to the PCI Express
	// SFF-8639 Module Specification.
	U2FormFactor FormFactor = "U2"
	// PCIeSlotFullLengthFormFactor shall indicate the drive is an add-in card
	// greater than 7 inches in length.
	PCIeSlotFullLengthFormFactor FormFactor = "PCIeSlotFullLength"
	// PCIeSlotLowProfileFormFactor shall indicate the drive is an add-in card less
	// than 2.5 inches in height.
	PCIeSlotLowProfileFormFactor FormFactor = "PCIeSlotLowProfile"
	// PCIeHalfLengthFormFactor shall indicate the drive is an add-in card less
	// than 7 inches in length.
	PCIeHalfLengthFormFactor FormFactor = "PCIeHalfLength"
	// OEMFormFactor shall indicate the drive is an OEM-defined form factor.
	OEMFormFactor FormFactor = "OEM"
)

type HotspareReplacementModeType string

const (
	// RevertibleHotspareReplacementModeType The hot spare drive that is
	// commissioned due to a drive failure reverts to a hot spare after the failed
	// drive is replaced and rebuilt.
	RevertibleHotspareReplacementModeType HotspareReplacementModeType = "Revertible"
	// NonRevertibleHotspareReplacementModeType The hot spare drive that is
	// commissioned due to a drive failure remains as a data drive and does not
	// revert to a hot spare if the failed drive is replaced.
	NonRevertibleHotspareReplacementModeType HotspareReplacementModeType = "NonRevertible"
)

type HotspareType string

const (
	// NoneHotspareType The drive is not a hot spare.
	NoneHotspareType HotspareType = "None"
	// GlobalHotspareType The drive is serving as a hot spare for all other drives
	// in this storage domain.
	GlobalHotspareType HotspareType = "Global"
	// ChassisHotspareType The drive is serving as a hot spare for all other drives
	// in this storage domain that are contained in the same chassis.
	ChassisHotspareType HotspareType = "Chassis"
	// DedicatedHotspareType The drive is serving as a hot spare for a user-defined
	// set of drives or volumes. Clients cannot specify this value when modifying
	// the 'HotspareType' property. This value is reported as a result of
	// configuring the spare drives within a volume.
	DedicatedHotspareType HotspareType = "Dedicated"
)

type MediaType string

const (
	// HDDMediaType The drive media type is traditional magnetic platters.
	HDDMediaType MediaType = "HDD"
	// SSDMediaType The drive media type is solid state or flash memory.
	SSDMediaType MediaType = "SSD"
	// SMRMediaType The drive media type is shingled magnetic recording.
	SMRMediaType MediaType = "SMR"
)

type StatusIndicator string

const (
	// OKStatusIndicator The drive is OK.
	OKStatusIndicator StatusIndicator = "OK"
	// FailStatusIndicator The drive has failed.
	FailStatusIndicator StatusIndicator = "Fail"
	// RebuildStatusIndicator The drive is being rebuilt.
	RebuildStatusIndicator StatusIndicator = "Rebuild"
	// PredictiveFailureAnalysisStatusIndicator The drive still works but is
	// predicted to fail soon.
	PredictiveFailureAnalysisStatusIndicator StatusIndicator = "PredictiveFailureAnalysis"
	// HotspareStatusIndicator The drive has been marked to automatically rebuild
	// and replace a failed drive.
	HotspareStatusIndicator StatusIndicator = "Hotspare"
	// InACriticalArrayStatusIndicator The array to which this drive belongs has
	// been degraded.
	InACriticalArrayStatusIndicator StatusIndicator = "InACriticalArray"
	// InAFailedArrayStatusIndicator The array to which this drive belongs has
	// failed.
	InAFailedArrayStatusIndicator StatusIndicator = "InAFailedArray"
)

// Drive shall represent a drive or other physical storage medium for a Redfish
// implementation. It may also represent a location, such as a slot, socket, or
// bay, where a unit may be installed, but the 'State' property within the
// 'Status' property contains 'Absent'.
type Drive struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	//
	// Version added: v1.3.0
	assembly string
	// AssetTag shall track the drive for inventory purposes.
	AssetTag string
	// BlockSecurityIDEnabled shall indicate if establishment of a TCG-defined
	// security ID (SID) on the drive is blocked. The value 'true' shall indicate
	// the TCG-defined 'Block SID' command is sent to the drive during each drive
	// boot sequence.
	//
	// Version added: v1.20.0
	BlockSecurityIDEnabled bool
	// BlockSizeBytes shall contain the size of the smallest addressable unit of
	// the associated drive.
	BlockSizeBytes *int `json:",omitempty"`
	// CapableSpeedGbs shall contain fastest capable bus speed, in gigabits per
	// second (Gbit/s) units, of the associated drive.
	CapableSpeedGbs *float64 `json:",omitempty"`
	// CapacityBytes shall contain the raw size, in bytes, of the associated drive.
	CapacityBytes *int `json:",omitempty"`
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains certificates for device identity and
	// attestation.
	//
	// Version added: v1.12.0
	certificates string
	// ConfigurationLock shall indicate whether configuration requests to the drive
	// are locked. Services shall reject modification requests that contain the
	// value 'Partial'.
	//
	// Version added: v1.19.0
	ConfigurationLock ConfigurationLock
	// DriveFormFactor shall contain the form factor of the drive inserted in this
	// slot.
	//
	// Version added: v1.16.0
	DriveFormFactor FormFactor
	// EncryptionAbility shall contain the encryption ability for the associated
	// drive.
	EncryptionAbility EncryptionAbility
	// EncryptionStatus shall contain the encryption status for the associated
	// drive.
	EncryptionStatus EncryptionStatus
	// EnvironmentMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that specifies the environment metrics for this drive.
	//
	// Version added: v1.12.0
	environmentMetrics string
	// FailurePredicted shall indicate whether this drive currently predicts a
	// manufacturer-defined failure.
	FailurePredicted bool
	// FirmwareVersion shall contain the firmware version as defined by the
	// manufacturer for this drive.
	//
	// Version added: v1.17.0
	FirmwareVersion string
	// HardwareVersion shall contain the hardware version of this drive as
	// determined by the vendor or supplier.
	//
	// Version added: v1.21.0
	HardwareVersion string
	// HotspareReplacementMode shall indicate whether a commissioned hot spare
	// continues to serve as a hot spare after the failed drive is replaced.
	//
	// Version added: v1.5.0
	HotspareReplacementMode HotspareReplacementModeType
	// HotspareType shall contain the hot spare type for the associated drive. If
	// the drive currently serves as a hot spare, the 'State' property in 'Status'
	// shall contain 'StandbySpare' and 'Enabled' when it is part of a volume.
	HotspareType HotspareType
	// Identifiers shall contain a list of all known durable names for the
	// associated drive.
	Identifiers []Identifier
	// IndicatorLED shall contain the state for the indicator light associated with
	// this drive.
	//
	// Deprecated: v1.11.0
	// This property has been deprecated in favor of the 'LocationIndicatorActive'
	// property.
	IndicatorLED IndicatorLED
	// Location shall contain the location information of the associated drive.
	//
	// Deprecated: v1.4.0
	// This property has been deprecated in favor of the singular
	// 'PhysicalLocation' property.
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource. A write to this property shall
	// update the value of 'IndicatorLED' in this resource, if supported, to
	// reflect the implementation of the locating function.
	//
	// Version added: v1.11.0
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for
	// producing the drive. This organization may be the entity from whom the drive
	// is purchased, but this is not necessarily true.
	Manufacturer string
	// Measurements shall contain an array of DSP0274-defined measurement blocks.
	//
	// Version added: v1.12.0
	//
	// Deprecated: v1.14.0
	// This property has been deprecated in favor of the 'ComponentIntegrity'
	// resource.
	Measurements []MeasurementBlock
	// MediaType shall contain the type of media contained in the associated drive.
	MediaType MediaType
	// Metrics shall contain a link to the metrics associated with this drive.
	//
	// Version added: v1.17.0
	metrics string
	// Model shall contain the name by which the manufacturer generally refers to
	// the drive.
	Model string
	// Multipath shall indicate whether the drive is accessible by an initiator
	// from multiple paths allowing for failover capabilities upon a path failure.
	//
	// Version added: v1.9.0
	Multipath bool
	// NVMe shall contain NVMe-specific properties of this drive.
	//
	// Version added: v1.20.0
	NVMe NVMe
	// NegotiatedSpeedGbs shall contain current bus speed, in gigabits per second
	// (Gbit/s) units, of the associated drive.
	NegotiatedSpeedGbs *float64 `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Operations shall contain a list of all operations currently running on the
	// drive.
	//
	// Version added: v1.1.0
	Operations []Operations
	// PartNumber shall contain the part number assigned by the organization that
	// is responsible for producing or manufacturing the drive.
	PartNumber string
	// PhysicalLocation shall contain the location information of the associated
	// drive.
	//
	// Version added: v1.4.0
	PhysicalLocation Location
	// PredictedMediaLifeLeftPercent shall contain an indicator of the percentage,
	// typically '0' to '100', of life remaining in the drive's media.
	PredictedMediaLifeLeftPercent *float64 `json:",omitempty"`
	// Protocol shall contain the protocol that the associated drive currently uses
	// to communicate to the storage controller for this system.
	Protocol Protocol
	// ReadyToRemove shall indicate whether the system is prepared for the removal
	// of this drive.
	//
	// Version added: v1.10.0
	ReadyToRemove bool
	// Revision shall contain the manufacturer-defined revision for a SCSI-based
	// drive, as returned in the 'Product Revision Level' field from the 'Inquiry'
	// command, which can be the firmware or hardware version. For other types of
	// drives, this property should not be present and services should support the
	// 'HardwareVersion' and 'FirmwareVersion' properties in favor of this
	// property. If this property is present for drives that are not SCSI-based,
	// the value may contain a firmware version, hardware version, or a
	// combination.
	Revision string
	// RotationSpeedRPM shall contain the rotation speed, in revolutions per minute
	// (RPM) units, of the associated drive.
	RotationSpeedRPM *float64 `json:",omitempty"`
	// SKU shall contain the stock-keeping unit (SKU) number for this drive.
	SKU string
	// SecurityIDOwned shall indicate if a TCG-defined security ID (SID) differs
	// from the TCG-defined manufacturing security ID (MSID) on the drive, flagging
	// the drive as owned. The value 'true' shall indicate that the MSID and SID
	// differ. This property should not be present for non-self-encrypting (SED)
	// drives.
	//
	// Version added: v1.22.0
	SecurityIDOwned bool
	// SerialNumber shall contain the manufacturer-allocated number that identifies
	// the drive.
	SerialNumber string
	// SlotCapableProtocols shall contain the drive protocols capable in this slot.
	// The value of this property depends upon the connector in this slot, the
	// storage controllers connected to this slot, the configuration of the system,
	// and other constraints that determine if a particular protocol is capable at
	// a given time.
	//
	// Version added: v1.16.0
	SlotCapableProtocols []Protocol
	// SlotFormFactor shall contain the form factor of the slot.
	//
	// Version added: v1.16.0
	SlotFormFactor FormFactor
	// SparePartNumber shall contain the spare part number of the drive.
	//
	// Version added: v1.19.0
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// StatusIndicator shall contain the status indicator state for the status
	// indicator associated with this drive. The '@Redfish.AllowableValues'
	// annotation specifies the valid values for this property.
	StatusIndicator StatusIndicator
	// TargetConfigurationLockLevel shall contain the target configuration lock
	// level for the drive. For NVMe drives, services shall implement the locking
	// requirements specified by SNIA's Swordfish NVMe Model Overview and Mapping
	// Guide.
	//
	// Version added: v1.20.0
	TargetConfigurationLockLevel TargetConfigurationLockLevel
	// WriteCacheEnabled shall indicate whether the drive write cache is enabled.
	//
	// Version added: v1.7.0
	WriteCacheEnabled bool
	// freezePersonalityTarget is the URL to send FreezePersonality requests.
	freezePersonalityTarget string
	// getPersonalityNonceTarget is the URL to send GetPersonalityNonce requests.
	getPersonalityNonceTarget string
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// revertPersonalitiesToDefaultsTarget is the URL to send RevertPersonalitiesToDefaults requests.
	revertPersonalitiesToDefaultsTarget string
	// revertToOriginalFactoryStateTarget is the URL to send RevertToOriginalFactoryState requests.
	revertToOriginalFactoryStateTarget string
	// secureEraseTarget is the URL to send SecureErase requests.
	secureEraseTarget string
	// setPersonalityKeyTarget is the URL to send SetPersonalityKey requests.
	setPersonalityKeyTarget string
	// unfreezePersonalityTarget is the URL to send UnfreezePersonality requests.
	unfreezePersonalityTarget string
	// activeSoftwareImage is the URI for ActiveSoftwareImage.
	activeSoftwareImage string
	// chassis is the URI for Chassis.
	chassis string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// networkDeviceFunctions are the URIs for NetworkDeviceFunctions.
	networkDeviceFunctions []string
	// pCIeFunctions are the URIs for PCIeFunctions.
	pCIeFunctions []string
	// softwareImages are the URIs for SoftwareImages.
	softwareImages []string
	// storage is the URI for Storage.
	storage string
	// storagePools are the URIs for StoragePools.
	storagePools []string
	// volumes are the URIs for Volumes.
	volumes []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Drive object from the raw JSON.
func (d *Drive) UnmarshalJSON(b []byte) error {
	type temp Drive
	type dActions struct {
		FreezePersonality             ActionTarget `json:"#Drive.FreezePersonality"`
		GetPersonalityNonce           ActionTarget `json:"#Drive.GetPersonalityNonce"`
		Reset                         ActionTarget `json:"#Drive.Reset"`
		RevertPersonalitiesToDefaults ActionTarget `json:"#Drive.RevertPersonalitiesToDefaults"`
		RevertToOriginalFactoryState  ActionTarget `json:"#Drive.RevertToOriginalFactoryState"`
		SecureErase                   ActionTarget `json:"#Drive.SecureErase"`
		SetPersonalityKey             ActionTarget `json:"#Drive.SetPersonalityKey"`
		UnfreezePersonality           ActionTarget `json:"#Drive.UnfreezePersonality"`
	}
	type dLinks struct {
		ActiveSoftwareImage    Link  `json:"ActiveSoftwareImage"`
		Chassis                Link  `json:"Chassis"`
		Endpoints              Links `json:"Endpoints"`
		NetworkDeviceFunctions Links `json:"NetworkDeviceFunctions"`
		PCIeFunctions          Links `json:"PCIeFunctions"`
		SoftwareImages         Links `json:"SoftwareImages"`
		Storage                Link  `json:"Storage"`
		StoragePools           Links `json:"StoragePools"`
		Volumes                Links `json:"Volumes"`
	}
	var tmp struct {
		temp
		Actions            dActions
		Links              dLinks
		Assembly           Link `json:"Assembly"`
		Certificates       Link `json:"Certificates"`
		EnvironmentMetrics Link `json:"EnvironmentMetrics"`
		Metrics            Link `json:"Metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*d = Drive(tmp.temp)

	// Extract the links to other entities for later
	d.freezePersonalityTarget = tmp.Actions.FreezePersonality.Target
	d.getPersonalityNonceTarget = tmp.Actions.GetPersonalityNonce.Target
	d.resetTarget = tmp.Actions.Reset.Target
	d.revertPersonalitiesToDefaultsTarget = tmp.Actions.RevertPersonalitiesToDefaults.Target
	d.revertToOriginalFactoryStateTarget = tmp.Actions.RevertToOriginalFactoryState.Target
	d.secureEraseTarget = tmp.Actions.SecureErase.Target
	d.setPersonalityKeyTarget = tmp.Actions.SetPersonalityKey.Target
	d.unfreezePersonalityTarget = tmp.Actions.UnfreezePersonality.Target
	d.activeSoftwareImage = tmp.Links.ActiveSoftwareImage.String()
	d.chassis = tmp.Links.Chassis.String()
	d.endpoints = tmp.Links.Endpoints.ToStrings()
	d.networkDeviceFunctions = tmp.Links.NetworkDeviceFunctions.ToStrings()
	d.pCIeFunctions = tmp.Links.PCIeFunctions.ToStrings()
	d.softwareImages = tmp.Links.SoftwareImages.ToStrings()
	d.storage = tmp.Links.Storage.String()
	d.storagePools = tmp.Links.StoragePools.ToStrings()
	d.volumes = tmp.Links.Volumes.ToStrings()
	d.assembly = tmp.Assembly.String()
	d.certificates = tmp.Certificates.String()
	d.environmentMetrics = tmp.EnvironmentMetrics.String()
	d.metrics = tmp.Metrics.String()

	// This is a read/write object, so we need to save the raw object data for later
	d.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (d *Drive) Update() error {
	readWriteFields := []string{
		"AssetTag",
		"BlockSecurityIDEnabled",
		"ConfigurationLock",
		"HotspareReplacementMode",
		"HotspareType",
		"IndicatorLED",
		"LocationIndicatorActive",
		"ReadyToRemove",
		"StatusIndicator",
		"TargetConfigurationLockLevel",
		"WriteCacheEnabled",
	}

	return d.UpdateFromRawData(d, d.RawData, readWriteFields)
}

// GetDrive will get a Drive instance from the service.
func GetDrive(c Client, uri string) (*Drive, error) {
	return GetObject[Drive](c, uri)
}

// ListReferencedDrives gets the collection of Drive from
// a provided reference.
func ListReferencedDrives(c Client, link string) ([]*Drive, error) {
	return GetCollectionObjects[Drive](c, link)
}

// This action shall freeze a personality for the NVMe subsystem as defined by
// the 'Configurable Device Personality' feature in the NVMe Base
// Specification.
// personality - This parameter shall contain the personality to freeze as
// defined by the 'Personality Identifier List' figure in the NVMe Base
// Specification. The value '255' shall indicate all personalities.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (d *Drive) FreezePersonality(personality uint) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Personality"] = personality
	resp, taskInfo, err := PostWithTask(d.client,
		d.freezePersonalityTarget, payload, d.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall obtain a nonce required to unfreeze personalities with a
// programmed key as defined by the 'CDP Random Nonce Data' clause in the NVMe
// Base Specification.
func (d *Drive) GetPersonalityNonce() (*GetPersonalityNonceResponse, error) {
	payload := make(map[string]any)

	resp, err := d.PostWithResponse(d.getPersonalityNonceTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result GetPersonalityNonceResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// This action shall reset this drive.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (d *Drive) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(d.client,
		d.resetTarget, payload, d.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall revert all personalities to their manufacturing default
// state for the NVMe subsystem as defined by the 'Manufacturing Default
// Personality' clause of the NVMe Base Specification.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (d *Drive) RevertPersonalitiesToDefaults() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(d.client,
		d.revertPersonalitiesToDefaultsTarget, payload, d.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall revert a self-encrypting drive (SED) to the original
// factory state. Upon successful completion of this action, the drive data
// shall be securely erased and the 'EncryptionStatus' property shall contain
// 'Unencrypted'.
// physicalSecureID - This parameter shall contain the physical secure ID
// (PSID). The PSID is generally printed on the drive label and used to revert
// an encrypted SED.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (d *Drive) RevertToOriginalFactoryState(physicalSecureID string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["PhysicalSecureID"] = physicalSecureID
	resp, taskInfo, err := PostWithTask(d.client,
		d.revertToOriginalFactoryStateTarget, payload, d.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall securely erase the drive.
// overwritePasses - This parameter shall contain the number of times to
// overwrite the drive if the 'SanitizationType' parameter contains the value
// 'Overwrite'. This parameter shall be ignored if the 'SanitizationType'
// parameter does not contain the value 'Overwrite'. If the client does not
// provide this parameter, the service shall perform an implementation-specific
// number of passes.
// sanitizationType - This parameter shall contain the type of data
// sanitization to perform for the secure erase request. The service can accept
// a request without the parameter and perform an implementation-specific
// default secure erase.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (d *Drive) SecureErase(overwritePasses int, sanitizationType DataSanitizationType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["OverwritePasses"] = overwritePasses
	payload["SanitizationType"] = sanitizationType
	resp, taskInfo, err := PostWithTask(d.client,
		d.secureEraseTarget, payload, d.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall set the personality key for the NVMe subsystem as defined
// by the 'Programmable Key Authentication Considerations' clause in the NVMe
// Base Specification.
// algorithm - This parameter shall contain the type of authentication
// algorithm for the key as defined by the 'CDP Authentication Algorithm' field
// in the NVMe Base Specification.
// key - This parameter shall contain the personality key to set on the NVMe
// subsystem as a hex-encoded string.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (d *Drive) SetPersonalityKey(algorithm NMVePersonalityKeyAlgorithm, key string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Algorithm"] = algorithm
	payload["Key"] = key
	resp, taskInfo, err := PostWithTask(d.client,
		d.setPersonalityKeyTarget, payload, d.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// DriveUnfreezePersonalityParameters holds the parameters for the UnfreezePersonality action.
type DriveUnfreezePersonalityParameters struct {
	// AuthenticationMethod shall contain the authentication method for the
	// unfreeze operation as defined by the 'Authenticated Unfreeze Support' field
	// in the NVMe Base Specification.
	AuthenticationMethod NMVeUnfreezeAuthMode `json:"AuthenticationMethod,omitempty"`
	// Key shall contain the authentication key for the unfreeze operation as a
	// hex-encoded string. If 'AuthenticationMethod' contains 'ProgrammedKey', this
	// parameter contains a key that is derived from the original key provided in
	// the 'SetPersonalityKey' action. If 'AuthenticationMethod' contains
	// 'PhysicalId', this parameter contains the physical secure ID (PSID) of the
	// drive converted to a hex-encoded string.
	Key string `json:"Key,omitempty"`
	// Nonce shall contain the personality nonce provided by the
	// 'GetPersonalityNonce' action. This parameter shall be required if
	// 'AuthenticationMethod' contains 'ProgrammedKey'.
	Nonce string `json:"Nonce,omitempty"`
	// Personality shall contain the personality to freeze as defined by the
	// 'Personality Identifier List' figure in the NVMe Base Specification. The
	// value '255' shall indicate all personalities.
	Personality uint `json:"Personality,omitempty"`
}

// This action shall unfreeze a personality for the NVMe subsystem as defined
// by the 'Authenticated Unfreeze Operation' clause in the NVMe Base
// Specification.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (d *Drive) UnfreezePersonality(params *DriveUnfreezePersonalityParameters) (*TaskMonitorInfo, error) {
	resp, taskInfo, err := PostWithTask(d.client,
		d.unfreezePersonalityTarget, params, d.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// ActiveSoftwareImage gets the ActiveSoftwareImage linked resource.
func (d *Drive) ActiveSoftwareImage() (*SoftwareInventory, error) {
	if d.activeSoftwareImage == "" {
		return nil, nil
	}
	return GetObject[SoftwareInventory](d.client, d.activeSoftwareImage)
}

// Chassis gets the Chassis linked resource.
func (d *Drive) Chassis() (*Chassis, error) {
	if d.chassis == "" {
		return nil, nil
	}
	return GetObject[Chassis](d.client, d.chassis)
}

// Endpoints gets the Endpoints linked resources.
func (d *Drive) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](d.client, d.endpoints)
}

// NetworkDeviceFunctions gets the NetworkDeviceFunctions linked resources.
func (d *Drive) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	return GetObjects[NetworkDeviceFunction](d.client, d.networkDeviceFunctions)
}

// PCIeFunctions gets the PCIeFunctions linked resources.
func (d *Drive) PCIeFunctions() ([]*PCIeFunction, error) {
	return GetObjects[PCIeFunction](d.client, d.pCIeFunctions)
}

// SoftwareImages gets the SoftwareImages linked resources.
func (d *Drive) SoftwareImages() ([]*SoftwareInventory, error) {
	return GetObjects[SoftwareInventory](d.client, d.softwareImages)
}

// Storage gets the Storage linked resource.
func (d *Drive) Storage() (*Storage, error) {
	if d.storage == "" {
		return nil, nil
	}
	return GetObject[Storage](d.client, d.storage)
}

// StoragePools gets the StoragePools linked resources.
func (d *Drive) StoragePools() ([]*StoragePool, error) {
	return GetObjects[StoragePool](d.client, d.storagePools)
}

// Volumes gets the Volumes linked resources.
func (d *Drive) Volumes() ([]*Volume, error) {
	return GetObjects[Volume](d.client, d.volumes)
}

// Assembly gets the Assembly linked resource.
func (d *Drive) Assembly() (*Assembly, error) {
	if d.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](d.client, d.assembly)
}

// Certificates gets the Certificates collection.
func (d *Drive) Certificates() ([]*Certificate, error) {
	if d.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](d.client, d.certificates)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (d *Drive) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if d.environmentMetrics == "" {
		return nil, nil
	}
	return GetObject[EnvironmentMetrics](d.client, d.environmentMetrics)
}

// Metrics gets the Metrics linked resource.
func (d *Drive) Metrics() (*DriveMetrics, error) {
	if d.metrics == "" {
		return nil, nil
	}
	return GetObject[DriveMetrics](d.client, d.metrics)
}

// NVMe shall contain NVMe-specific properties of a drive.
type NVMe struct {
	// ActivePersonalities shall contain the active personality settings of the
	// NVMe subsystem as defined by the NVMe Device Personalities log page and NVMe
	// Configurable Device Personality feature.
	//
	// Version added: v1.22.0
	ActivePersonalities []NVMePersonality
	// ConfigurationLockState shall contain the configurable features that are able
	// to be locked from in-band usage on an NVMe subsystem and their current lock
	// state.
	//
	// Version added: v1.20.0
	ConfigurationLockState NVMeConfigurationLockState
}

// Operations shall describe a currently running operation on the resource.
type Operations struct {
	// AssociatedTask shall contain a link to a resource of type 'Task' that
	// represents the task associated with the operation.
	//
	// Version added: v1.1.0
	associatedTask string
	// Operation shall contain the type of the operation.
	//
	// Version added: v1.17.0
	Operation OperationType
	// OperationName shall contain a string of the name of the operation.
	//
	// Version added: v1.1.0
	//
	// Deprecated: v1.17.0
	// This property is deprecated in favor of the 'Operation' property using the
	// 'OperationType' enumeration defined in the 'Volume' schema.
	OperationName string
	// PercentageComplete shall contain an integer of the percentage, '0' to '100',
	// of the operation that has been completed.
	//
	// Version added: v1.1.0
	PercentageComplete *uint `json:",omitempty"`
}

// UnmarshalJSON unmarshals a Operations object from the raw JSON.
func (o *Operations) UnmarshalJSON(b []byte) error {
	type temp Operations
	var tmp struct {
		temp
		AssociatedTask Link `json:"AssociatedTask"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*o = Operations(tmp.temp)

	// Extract the links to other entities for later
	o.associatedTask = tmp.AssociatedTask.String()

	return nil
}

// AssociatedTask gets the AssociatedTask linked resource.
func (o *Operations) AssociatedTask(client Client) (*Task, error) {
	if o.associatedTask == "" {
		return nil, nil
	}
	return GetObject[Task](client, o.associatedTask)
}
