//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

type TrustedComponentType string

const (
	// DiscreteTrustedComponentType shall indicate that the entity has a well-defined physical boundary within the
	// chassis.
	DiscreteTrustedComponentType TrustedComponentType = "Discrete"
	// IntegratedTrustedComponentType shall indicate that the entity is integrated into another device.
	IntegratedTrustedComponentType TrustedComponentType = "Integrated"
)

// TrustedComponentLinks shall contain links to resources that are related to but are not contained by, or subordinate to, this
// resource.
type TrustedComponentLinks struct {
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

type TPMActions struct {
	TPMClear       common.ActionTarget `json:"#TrustedComponent.TPMClear"`
	TPMGetEventLog common.ActionTarget `json:"#TrustedComponent.TPMGetEventLog"`
}

// TrustedComponent shall represent a trusted component in a Redfish implementation.
type TrustedComponent struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
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
	TPM TPM
	// TrustedComponentType shall contain the type of trusted component.
	TrustedComponentType TrustedComponentType
	// UUID shall contain a universally unique identifier number for the trusted component.
	UUID string

	Actions          TPMActions  `json:"Actions"`
	CertificatesLink common.Link `json:"Certificates"`
	Links            TrustedComponentLinks
}

// ActiveSoftwareImage gets the active firmware image for this trusted component.
func (trustedComponent *TrustedComponent) ActiveSoftwareImage(queryOpts ...common.QueryGroupOption) (*SoftwareInventory, error) {
	return trustedComponent.ActiveSoftwareImageWithContext(common.ContextOf(trustedComponent.GetClient()), queryOpts...)
}

// ActiveSoftwareImageWithContext gets the active firmware image for this trusted component.
func (trustedComponent *TrustedComponent) ActiveSoftwareImageWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*SoftwareInventory, error) {
	if trustedComponent.Links.ActiveSoftwareImage.IsZero() {
		return nil, nil
	}
	return GetSoftwareInventoryWithContext(ctx, trustedComponent.GetClient(), trustedComponent.Links.ActiveSoftwareImage.String(), queryOpts...)
}

// ComponentIntegrity gets the resources for which the trusted component is responsible.
func (trustedComponent *TrustedComponent) ComponentIntegrity(queryOpts ...common.QueryGroupOption) ([]*ComponentIntegrity, error) {
	return trustedComponent.ComponentIntegrityWithContext(common.ContextOf(trustedComponent.GetClient()), queryOpts...)
}

// ComponentIntegrityWithContext gets the resources for which the trusted component is responsible.
func (trustedComponent *TrustedComponent) ComponentIntegrityWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*ComponentIntegrity, error) {
	if len(trustedComponent.Links.ComponentIntegrity) == 0 {
		return nil, nil
	}
	return common.GetObjectsWithContext[ComponentIntegrity](ctx, trustedComponent.GetClient(), trustedComponent.Links.ComponentIntegrity.ToStrings(), queryOpts...)
}

// SoftwareImages gets the firmware images that apply to this trusted component.
func (trustedComponent *TrustedComponent) SoftwareImages(queryOpts ...common.QueryGroupOption) ([]*SoftwareInventory, error) {
	return trustedComponent.SoftwareImagesWithContext(common.ContextOf(trustedComponent.GetClient()), queryOpts...)
}

// SoftwareImagesWithContext gets the firmware images that apply to this trusted component.
func (trustedComponent *TrustedComponent) SoftwareImagesWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*SoftwareInventory, error) {
	if len(trustedComponent.Links.SoftwareImages) == 0 {
		return nil, nil
	}
	return common.GetObjectsWithContext[SoftwareInventory](ctx, trustedComponent.GetClient(), trustedComponent.Links.SoftwareImages.ToStrings(), queryOpts...)
}

// Certificates gets the certificates associated with this trusted component.
func (trustedComponent *TrustedComponent) Certificates(queryOpts ...common.QueryGroupOption) ([]*Certificate, error) {
	return trustedComponent.CertificatesWithContext(common.ContextOf(trustedComponent.GetClient()), queryOpts...)
}

// CertificatesWithContext gets the certificates associated with this trusted component.
func (trustedComponent *TrustedComponent) CertificatesWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Certificate, error) {
	if trustedComponent.CertificatesLink.IsZero() {
		return nil, nil
	}
	return ListReferencedCertificatesWithContext(ctx, trustedComponent.GetClient(), trustedComponent.CertificatesLink.String(), queryOpts...)
}

// TPMGetEventLog gets the event log for TPM 2.0 devices.
func (trustedComponent *TrustedComponent) TPMGetEventLog() (*TPMGetEventLogResponse, error) {
	return trustedComponent.TPMGetEventLogWithContext(common.ContextOf(trustedComponent.GetClient()))
}

// TPMGetEventLogWithContext gets the event log for TPM 2.0 devices.
func (trustedComponent *TrustedComponent) TPMGetEventLogWithContext(ctx context.Context) (*TPMGetEventLogResponse, error) {
	if trustedComponent.Actions.TPMGetEventLog.Target == "" {
		return nil, ErrActionNotSupported
	}

	resp, err := trustedComponent.PostWithResponseWithContext(ctx, trustedComponent.Actions.TPMGetEventLog.Target, nil)
	defer common.DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return nil, err
	}

	var tpmGetEventLogResponse TPMGetEventLogResponse
	err = json.NewDecoder(resp.Body).Decode(&tpmGetEventLogResponse)
	if err != nil {
		return nil, err
	}

	return &tpmGetEventLogResponse, nil
}

// GetTrustedComponent will get a TrustedComponent instance from the service.
func GetTrustedComponent(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*TrustedComponent, error) {
	return GetTrustedComponentWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetTrustedComponentWithContext will get a TrustedComponent instance from the service.
func GetTrustedComponentWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*TrustedComponent, error) {
	return common.GetObjectWithContext[TrustedComponent](ctx, c, uri, queryOpts...)
}

// ListReferencedTrustedComponents gets the collection of TrustedComponent from
// a provided reference.
func ListReferencedTrustedComponents(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*TrustedComponent, error) {
	return ListReferencedTrustedComponentsWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedTrustedComponentsWithContext gets the collection of TrustedComponent from
// a provided reference.
func ListReferencedTrustedComponentsWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*TrustedComponent, error) {
	return common.GetCollectionObjectsWithContext[TrustedComponent](ctx, c, link, queryOpts...)
}
