//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type TrustedComponentType string

const (
	// DiscreteTrustedComponentType shall indicate that the entity has a well-defined physical boundary within the
	// chassis.
	DiscreteTrustedComponentType TrustedComponentType = "Discrete"
	// IntegratedTrustedComponentType shall indicate that the entity is integrated into another device.
	IntegratedTrustedComponentType TrustedComponentType = "Integrated"
)

// trustedComponentLinks shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type trustedComponentLinks struct {
	// ActiveSoftwareImage shall contain a link to a resource of type SoftwareInventory that represents the active
	// firmware image for this trusted component.
	ActiveSoftwareImage common.Link
	// ComponentIntegrity shall contain an array of links to resources of type ComponentIntegrity that represent the
	// communication established with the trusted component by other resources. The TargetComponentURI property in the
	// referenced ComponentIntegrity resources shall reference this trusted component.
	ComponentIntegrity common.Links
	// ComponentIntegrity@odata.count
	ComponentIntegrityCount int `json:"ComponentIntegrity@odata.count"`
	// ComponentsProtected shall contain an array of links to resources whose integrity is measured or reported by the
	// trusted component.
	ComponentsProtected common.Links
	// ComponentsProtected@odata.count
	ComponentsProtectedCount int `json:"ComponentsProtected@odata.count"`
	// IntegratedInto shall contain a link to a resource to which this trusted component is physically integrated. This
	// property shall be present if TrustedComponentType contains 'Integrated'.
	IntegratedInto common.Link
	// Owner shall contain a link to the resource that owns this trusted component. In the case of TPMs, particularly
	// in multiple chassis implementations, this is the resource used to establish a new PCR.
	Owner common.Link
	// SoftwareImages shall contain an array of links to resources of type SoftwareInventory that represent the
	// firmware images that apply to this trusted component.
	SoftwareImages common.Links
	// SoftwareImages@odata.count
	SoftwareImagesCount int `json:"SoftwareImages@odata.count"`
}

// TPM shall contain TPM-specific information for a trusted component.
type TPM struct {
	// CapabilitiesVendorID shall contain an ASCII string of the 4-byte TCG-defined 'TPM Capabilities Vendor ID' for
	// this trusted component.
	CapabilitiesVendorID string
	// HardwareInterfaceVendorID shall contain the TCG-defined 'TPM Hardware Interface Vendor ID' for this trusted
	// component with the most significant byte shown first.
	HardwareInterfaceVendorID string
}

// TPMGetEventLogResponse shall contain the TPM event log.
type TPMGetEventLogResponse struct {
	// EventLog shall contain a Base64-encoded string of the entire event log defined in the 'Event Logging' section of
	// the 'TCG PC Client Platform Firmware Profile Specification'.
	EventLog string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
}

// TrustedComponent shall represent a trusted component in a Redfish implementation.
type TrustedComponent struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains device
	// identity certificates of the trusted component.
	certificates string
	// Description provides a description of this resource.
	Description string
	// FirmwareVersion shall contain a version number associated with the active software image on the trusted
	// component.
	FirmwareVersion string
	// Manufacturer shall contain the name of the organization responsible for producing the trusted component. This
	// organization may be the entity from whom the trusted component is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the name by which the manufacturer generally refers to the trusted component.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain a part number assigned by the organization that is responsible for producing or
	// manufacturing the trusted component.
	PartNumber string
	// SKU shall contain the stock-keeping unit number for this trusted component.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the trusted component.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TPM shall contain TPM-specific information for this trusted component. This property shall only be present for
	// TCG-defined TPM trusted components.
	TPM string
	// TrustedComponentType shall contain the type of trusted component.
	TrustedComponentType TrustedComponentType
	// UUID shall contain a universally unique identifier number for the trusted component.
	UUID string

	tpmGetEventLogTarget string

	activeSoftwareImage string
	componentIntegrity  []string
	// ComponentIntegrityCount is the number of trusted component integrity links.
	ComponentIntegrityCount int
	componentsProtected     []string
	// ComponentsProtectedCount is the number of protected components.
	ComponentsProtectedCount int
	integratedInto           string
	owner                    string
	softwareImages           []string
	// SoftwareImagesCount is the number of images associated with this trusted component.
	SoftwareImagesCount int
}

// UnmarshalJSON unmarshals a TrustedComponent object from the raw JSON.
func (trustedcomponent *TrustedComponent) UnmarshalJSON(b []byte) error {
	type temp TrustedComponent
	var t struct {
		temp
		Certificates common.Link
		Links        trustedComponentLinks
		Actions      struct {
			TPMGetEventLog struct {
				Target string
			} `json:"#TrustedComponent.TPMGetEventLog"`
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*trustedcomponent = TrustedComponent(t.temp)
	trustedcomponent.tpmGetEventLogTarget = t.Actions.TPMGetEventLog.Target

	trustedcomponent.certificates = t.Certificates.String()

	trustedcomponent.activeSoftwareImage = t.Links.ActiveSoftwareImage.String()
	trustedcomponent.componentIntegrity = t.Links.ComponentIntegrity.ToStrings()
	trustedcomponent.ComponentIntegrityCount = t.Links.ComponentIntegrityCount
	trustedcomponent.softwareImages = t.Links.SoftwareImages.ToStrings()
	trustedcomponent.SoftwareImagesCount = t.Links.SoftwareImagesCount

	// TODO: Implement accessors for linked objects
	trustedcomponent.componentsProtected = t.Links.ComponentsProtected.ToStrings()
	trustedcomponent.ComponentsProtectedCount = t.Links.ComponentsProtectedCount
	trustedcomponent.integratedInto = t.Links.IntegratedInto.String()
	trustedcomponent.owner = t.Links.Owner.String()

	return nil
}

// ActiveSoftwareImage gets the active firmware image for this trusted component.
func (trustedcomponent *TrustedComponent) ActiveSoftwareImage() (*SoftwareInventory, error) {
	if trustedcomponent.activeSoftwareImage == "" {
		return nil, nil
	}
	return GetSoftwareInventory(trustedcomponent.GetClient(), trustedcomponent.activeSoftwareImage)
}

// ComponentIntegrity gets the resources for which the trusted component is responsible.
func (trustedcomponent *TrustedComponent) ComponentIntegrity() ([]*ComponentIntegrity, error) {
	return common.GetObjects[ComponentIntegrity](trustedcomponent.GetClient(), trustedcomponent.componentIntegrity)
}

// SoftwareImages gets the firmware images that apply to this trusted component.
func (trustedcomponent *TrustedComponent) SoftwareImages() ([]*SoftwareInventory, error) {
	return common.GetObjects[SoftwareInventory](trustedcomponent.GetClient(), trustedcomponent.softwareImages)
}

// Certificates gets the certificates associated with this trusted component.
func (trustedcomponent *TrustedComponent) Certificates() ([]*Certificate, error) {
	return ListReferencedCertificates(trustedcomponent.GetClient(), trustedcomponent.certificates)
}

// TPMGetEventLog gets the event log for TPM 2.0 devices.
func (trustedcomponent *TrustedComponent) TPMGetEventLog() (*TPMGetEventLogResponse, error) {
	resp, err := trustedcomponent.PostWithResponse(trustedcomponent.tpmGetEventLogTarget, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var tpmGetEventLogResponse TPMGetEventLogResponse
	err = json.NewDecoder(resp.Body).Decode(&tpmGetEventLogResponse)
	if err != nil {
		return nil, err
	}

	return &tpmGetEventLogResponse, nil
}

// GetTrustedComponent will get a TrustedComponent instance from the service.
func GetTrustedComponent(c common.Client, uri string) (*TrustedComponent, error) {
	return common.GetObject[TrustedComponent](c, uri)
}

// ListReferencedTrustedComponents gets the collection of TrustedComponent from
// a provided reference.
func ListReferencedTrustedComponents(c common.Client, link string) ([]*TrustedComponent, error) {
	return common.GetCollectionObjects[TrustedComponent](c, link)
}
