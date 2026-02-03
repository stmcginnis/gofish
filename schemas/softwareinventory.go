//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #SoftwareInventory.v1_13_0.SoftwareInventory

package schemas

import (
	"encoding/json"
)

type ReleaseType string

const (
	// ProductionReleaseType shall indicate that the software is ready for use in
	// production environments.
	ProductionReleaseType ReleaseType = "Production"
	// PrototypeReleaseType shall indicate that the software is designed for
	// development or internal use.
	PrototypeReleaseType ReleaseType = "Prototype"
	// OtherReleaseType shall indicate that the Redfish service cannot determine if
	// this release is validated or prototype.
	OtherReleaseType ReleaseType = "Other"
)

type VersionScheme string

const (
	// SemVerVersionScheme shall indicate that the value of the 'Version' property
	// conforms to the format and rules of the Semantic Version 2.0 specification,
	// and may include pre-release or build metadata.
	SemVerVersionScheme VersionScheme = "SemVer"
	// DotIntegerNotationVersionScheme shall indicate that the value of the
	// 'Version' property contains a sequence of integers separated by period (dot)
	// characters, and shall follow the pattern '^\d+(\.\d+)*$'. Leading zeros in
	// the sequence shall be ignored.
	DotIntegerNotationVersionScheme VersionScheme = "DotIntegerNotation"
	// OEMVersionScheme shall indicate that the value of the 'Version' property
	// follows a format and rules as defined by the vendor or manufacturer.
	OEMVersionScheme VersionScheme = "OEM"
)

// SoftwareInventory This resource contains a single software component that
// this Redfish service manages.
type SoftwareInventory struct {
	Entity
	// Active shall indicate that the software image is currently in use on one or
	// more devices.
	//
	// Version added: v1.12.0
	Active bool
	// AdditionalVersions shall contain the additional versions of this software.
	//
	// Version added: v1.7.0
	AdditionalVersions AdditionalVersions
	// AssociatedPhysicalContext shall contain a description of the physical
	// context for the software inventory data.
	//
	// Version added: v1.10.0
	AssociatedPhysicalContext PhysicalContext
	// ImageDataURI shall contain the URI at which to access the image data for
	// this software inventory, using the Redfish protocol and authentication
	// methods. This image should be the original vendor-provided image suitable
	// for update operations.
	//
	// Version added: v1.13.0
	ImageDataURI string
	// LowestSupportedVersion shall represent the lowest supported version of this
	// software. This string is formatted using the same format used for the
	// 'Version' property.
	//
	// Version added: v1.1.0
	LowestSupportedVersion string
	// Manufacturer shall represent the name of the manufacturer or producer of
	// this software.
	//
	// Version added: v1.2.0
	Manufacturer string
	// Measurement shall contain a DSP0274-defined measurement block.
	//
	// Version added: v1.4.0
	//
	// Deprecated: v1.6.0
	// This property has been deprecated in favor of the 'ComponentIntegrity'
	// resource.
	Measurement MeasurementBlock
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RelatedItem shall contain an array of links to resources or objects that
	// represent devices to which this software inventory applies.
	//
	// Version added: v1.1.0
	relatedItem []string
	// RelatedItemCount
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// ReleaseDate shall contain the date of release or production for this
	// software. If the time of day is unknown, the time of day portion of the
	// property shall contain '00:00:00Z'.
	//
	// Version added: v1.2.0
	ReleaseDate string
	// ReleaseType shall describe the type of release for the software.
	//
	// Version added: v1.10.0
	ReleaseType ReleaseType
	// ResetRequiredOnUpdate shall indicate whether a reset is required to apply an
	// update to this software. If 'true', a reset is required and clients should
	// expect a disruption in communication to targets utilizing this software
	// while applying an update. If 'false', a reset is not required and
	// communication can be maintained to targets utilizing this software
	// throughout an update.
	//
	// Version added: v1.12.0
	ResetRequiredOnUpdate bool
	// SizeBytes shall contain the size of the software image in bytes.
	//
	// Version added: v1.13.0
	SizeBytes *int `json:",omitempty"`
	// SoftwareID shall represent an implementation-specific label that identifies
	// this software. This string correlates with a component repository or
	// database.
	//
	// Version added: v1.1.0
	SoftwareID string `json:"SoftwareId"`
	// Staged shall indicate that the software image is ready to be activated to
	// one or more devices.
	//
	// Version added: v1.12.0
	Staged bool
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UefiDevicePaths shall contain a list UEFI device paths of the components
	// associated with this software inventory item. The UEFI device paths shall be
	// formatted as defined by the UEFI Specification.
	//
	// Version added: v1.1.0
	UefiDevicePaths []string
	// Updateable shall indicate whether the update service can update this
	// software. If 'true', the service can update this software. If 'false', the
	// service cannot update this software and the software is for reporting
	// purposes only.
	Updateable bool
	// Version shall contain the version of this software.
	Version string
	// VersionScheme shall describe the scheme used to format the value of the
	// 'Version' property.
	//
	// Version added: v1.9.0
	VersionScheme VersionScheme
	// WriteProtected shall indicate whether the software image can be overwritten,
	// where a value 'true' shall indicate that the software cannot be altered or
	// overwritten.
	//
	// Version added: v1.3.0
	WriteProtected bool
	// activateTarget is the URL to send Activate requests.
	activateTarget string
	// activeTargets are the URIs for ActiveTargets.
	activeTargets []string
	// stagedTargets are the URIs for StagedTargets.
	stagedTargets []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a SoftwareInventory object from the raw JSON.
func (s *SoftwareInventory) UnmarshalJSON(b []byte) error {
	type temp SoftwareInventory
	type sActions struct {
		Activate ActionTarget `json:"#SoftwareInventory.Activate"`
	}
	type sLinks struct {
		ActiveTargets Links `json:"ActiveTargets"`
		StagedTargets Links `json:"StagedTargets"`
	}
	var tmp struct {
		temp
		Actions     sActions
		Links       sLinks
		RelatedItem Links `json:"RelatedItem"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SoftwareInventory(tmp.temp)

	// Extract the links to other entities for later
	s.activateTarget = tmp.Actions.Activate.Target
	s.activeTargets = tmp.Links.ActiveTargets.ToStrings()
	s.stagedTargets = tmp.Links.StagedTargets.ToStrings()
	s.relatedItem = tmp.RelatedItem.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	s.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *SoftwareInventory) Update() error {
	readWriteFields := []string{
		"WriteProtected",
	}

	return s.UpdateFromRawData(s, s.RawData, readWriteFields)
}

// GetSoftwareInventory will get a SoftwareInventory instance from the service.
func GetSoftwareInventory(c Client, uri string) (*SoftwareInventory, error) {
	return GetObject[SoftwareInventory](c, uri)
}

// ListReferencedSoftwareInventories gets the collection of SoftwareInventory from
// a provided reference.
func ListReferencedSoftwareInventories(c Client, link string) ([]*SoftwareInventory, error) {
	return GetCollectionObjects[SoftwareInventory](c, link)
}

// This action shall activate this software inventory instance.
// targets - This parameter shall contain an array of target devices to
// activate this software image. If not specified, the service shall activate
// this software image on all applicable devices.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *SoftwareInventory) Activate(targets []Entity) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Targets"] = targets
	resp, taskInfo, err := PostWithTask(s.client,
		s.activateTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// ActiveTargets gets the ActiveTargets linked resources.
func (s *SoftwareInventory) ActiveTargets() ([]*Entity, error) {
	return GetObjects[Entity](s.client, s.activeTargets)
}

// StagedTargets gets the StagedTargets linked resources.
func (s *SoftwareInventory) StagedTargets() ([]*Entity, error) {
	return GetObjects[Entity](s.client, s.stagedTargets)
}

// RelatedItem gets the RelatedItem linked resources.
func (s *SoftwareInventory) RelatedItem() ([]*Entity, error) {
	return GetObjects[Entity](s.client, s.relatedItem)
}

// AdditionalVersions shall contain additional versions.
type AdditionalVersions struct {
	// BootParameters shall contain the version of the configuration file that
	// contains the initial boot parameters of this software.
	//
	// Version added: v1.11.0
	BootParameters string
	// Bootloader shall contain the bootloader version contained in this software.
	//
	// Version added: v1.7.0
	Bootloader string
	// FactoryConfiguration shall contain the version of the configuration that
	// contains the factory default runtime configuration parameters of this
	// software.
	//
	// Version added: v1.11.0
	FactoryConfiguration string
	// Kernel shall contain the kernel version contained in this software. For
	// strict POSIX software, the value shall contain the output of 'uname -srm'.
	// For Microsoft Windows, the value shall contain the output of 'ver', from
	// Command Prompt.
	//
	// Version added: v1.7.0
	Kernel string
	// Microcode shall contain the microcode version contained in this software.
	//
	// Version added: v1.7.0
	Microcode string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.7.0
	OEM json.RawMessage `json:"Oem"`
	// OSDistribution shall contain the operating system name of this software.
	//
	// Version added: v1.8.0
	OSDistribution string
}

// MeasurementBlock shall describe a DSP0274-defined measurement block.
type MeasurementBlock struct {
	// Measurement shall contain the value of the hexadecimal string representation
	// of the numeric value of the DSP0274-defined 'Measurement' field of the
	// measurement block.
	//
	// Version added: v1.4.0
	Measurement string
	// MeasurementIndex shall contain the value of the DSP0274-defined 'Index'
	// field of the measurement block.
	//
	// Version added: v1.5.0
	MeasurementIndex *int `json:",omitempty"`
	// MeasurementSize shall contain the value of the DSP0274-defined
	// 'MeasurementSize' field of the measurement block.
	//
	// Version added: v1.4.0
	MeasurementSize *int `json:",omitempty"`
	// MeasurementSpecification shall contain the value of the DSP0274-defined
	// 'MeasurementSpecification' field of the measurement block.
	//
	// Version added: v1.4.0
	MeasurementSpecification *int `json:",omitempty"`
}
