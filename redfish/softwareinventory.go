//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type ReleaseType string

const (
	// ProductionReleaseType shall indicate that the software is ready for use in production environments.
	ProductionReleaseType ReleaseType = "Production"
	// PrototypeReleaseType shall indicate that the software is designed for development or internal use.
	PrototypeReleaseType ReleaseType = "Prototype"
	// OtherReleaseType shall indicate that the Redfish service cannot determine if this release is validated or
	// prototype.
	OtherReleaseType ReleaseType = "Other"
)

type VersionScheme string

const (
	// SemVerVersionScheme shall indicate that the value of the Version property conforms to the format and rules of
	// the Semantic Version 2.0 specification, and may include pre-release or build metadata.
	SemVerVersionScheme VersionScheme = "SemVer"
	// DotIntegerNotationVersionScheme shall indicate that the value of the Version property contains a sequence of
	// integers separated by period (dot) characters, and shall follow the pattern '^\d+(\.\d+)*$'. Leading zeros in
	// the sequence shall be ignored.
	DotIntegerNotationVersionScheme VersionScheme = "DotIntegerNotation"
	// OEMVersionScheme shall indicate that the value of the Version property follows a format and rules as defined by
	// the vendor or manufacturer.
	OEMVersionScheme VersionScheme = "OEM"
)

// AdditionalVersions shall contain additional versions.
type AdditionalVersions struct {
	// Bootloader shall contain the bootloader version contained in this software.
	Bootloader string
	// Kernel shall contain the kernel version contained in this software. For strict POSIX software, the value shall
	// contain the output of 'uname -srm'. For Microsoft Windows, the value shall contain the output of 'ver', from
	// Command Prompt.
	Kernel string
	// Microcode shall contain the microcode version contained in this software.
	Microcode string
	// OSDistribution shall contain the operating system name of this software.
	OSDistribution string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage
}

// MeasurementBlock shall describe a DSP0274-defined measurement block.
type MeasurementBlock struct {
	// Measurement shall contain the value of the hexadecimal string representation of the numeric value of the
	// DSP0274-defined Measurement field of the measurement block.
	Measurement string
	// MeasurementIndex shall contain the value of the DSP0274-defined Index field of the measurement block.
	MeasurementIndex int
	// MeasurementSize shall contain the value of the DSP0274-defined MeasurementSize field of the measurement block.
	MeasurementSize int
	// MeasurementSpecification shall contain the value of the DSP0274-defined MeasurementSpecification field of the
	// measurement block.
	MeasurementSpecification int
}

// SoftwareInventory is a single software component that this Redfish Service manages.
type SoftwareInventory struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AdditionalVersions shall contain the additional versions of this software.
	AdditionalVersions AdditionalVersions
	// AssociatedPhysicalContext shall contain a description of the physical context for the software inventory data.
	AssociatedPhysicalContext PhysicalContext
	// Description provides a description of this resource.
	Description string
	// LowestSupportedVersion is used for the Version property.
	LowestSupportedVersion string
	// Manufacturer is represents the name of the
	// manufacturer or producer of this software.
	Manufacturer string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage
	// RelatedItem shall contain an array of IDs for pointers consistent with JSON Pointer syntax to the Resource that
	// is associated with this software inventory item.
	relatedItem []string
	// RelatedItemCount is the number of related items.
	RelatedItemCount int `json:"RelatedItem@odata.count"`
	// ReleaseDate is the date of release or
	// production for this software.  If the time of day is unknown, the time
	// of day portion of the property shall contain `00:00:00Z`.
	ReleaseDate string
	// SoftwareID is represents an implementation-specific
	// label that identifies this software.  This string correlates with a
	// component repository or database.
	SoftwareID string
	// Status is any status or health properties of the Resource.
	Status common.Status
	// UefiDevicePaths is a list UEFI device paths of the components associated with this software inventory item.
	// The UEFI device paths shall be formatted as defined by the UEFI Specification.
	UefiDevicePaths []string
	// Updateable is This property shall indicate whether the Update Service
	// can update this software.  If `true`, the Service can update this
	// software.  If `false`, the Service cannot update this software and the
	// software is for reporting purposes only.
	Updateable bool
	// Version is the version of this software.
	Version string
	// VersionScheme shall describe the scheme used to format the value of the Version property.
	VersionScheme VersionScheme
	// WriteProtected is This property shall indicate whether the software
	// image can be overwritten, where a value `true` shall indicate that the
	// software cannot be altered or overwritten.
	WriteProtected bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a SoftwareInventory object from the raw JSON.
func (softwareinventory *SoftwareInventory) UnmarshalJSON(b []byte) error {
	type temp SoftwareInventory
	var t struct {
		temp
		RelatedItem common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*softwareinventory = SoftwareInventory(t.temp)

	// Extract the links to other entities for later
	softwareinventory.relatedItem = t.RelatedItem.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	softwareinventory.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (softwareinventory *SoftwareInventory) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(SoftwareInventory)
	original.UnmarshalJSON(softwareinventory.rawData)

	readWriteFields := []string{
		"WriteProtected",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(softwareinventory).Elem()

	return softwareinventory.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetSoftwareInventory will get a SoftwareInventory instance from the service.
func GetSoftwareInventory(c common.Client, uri string) (*SoftwareInventory, error) {
	return common.GetObject[SoftwareInventory](c, uri)
}

// ListReferencedSoftwareInventories gets the collection of SoftwareInventory from
// a provided reference.
func ListReferencedSoftwareInventories(c common.Client, link string) ([]*SoftwareInventory, error) {
	return common.GetCollectionObjects[SoftwareInventory](c, link)
}
