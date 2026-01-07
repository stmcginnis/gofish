//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #TrustedComponent.v1_4_0.TrustedComponent

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type TrustedComponentType string

const (
	// DiscreteTrustedComponentType shall indicate that the entity has a
	// well-defined physical boundary within the chassis.
	DiscreteTrustedComponentType TrustedComponentType = "Discrete"
	// IntegratedTrustedComponentType shall indicate that the entity is integrated
	// into another device.
	IntegratedTrustedComponentType TrustedComponentType = "Integrated"
)

// TrustedComponent shall represent a trusted component in a Redfish
// implementation.
type TrustedComponent struct {
	common.Entity
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains device identity certificates of the
	// trusted component.
	certificates string
	// FirmwareVersion shall contain a version number associated with the active
	// software image on the trusted component.
	FirmwareVersion string
	// Location shall contain the location information of the trusted component.
	//
	// Version added: v1.4.0
	Location common.Location
	// Manufacturer shall contain the name of the organization responsible for
	// producing the trusted component. This organization may be the entity from
	// whom the trusted component is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the name by which the manufacturer generally refers to
	// the trusted component.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain a part number assigned by the organization that is
	// responsible for producing or manufacturing the trusted component.
	PartNumber string
	// SKU shall contain the stock-keeping unit number for this trusted component.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the trusted component.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TPM shall contain TPM-specific information for this trusted component. This
	// property shall only be present for TCG-defined TPM trusted components.
	//
	// Version added: v1.1.0
	TPM TPM
	// TrustedComponentType shall contain the type of trusted component.
	TrustedComponentType TrustedComponentType
	// UUID shall contain a universally unique identifier number for the trusted
	// component.
	UUID string
	// tPMGetEventLogTarget is the URL to send TPMGetEventLog requests.
	tPMGetEventLogTarget string
	// activeSoftwareImage is the URI for ActiveSoftwareImage.
	activeSoftwareImage string
	// componentIntegrity are the URIs for ComponentIntegrity.
	componentIntegrity []string
	// componentsProtected are the URIs for ComponentsProtected.
	componentsProtected []string
	// integratedInto is the URI for IntegratedInto.
	integratedInto string
	// owner is the URI for Owner.
	Owner string
	// softwareImages are the URIs for SoftwareImages.
	softwareImages []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a TrustedComponent object from the raw JSON.
func (t *TrustedComponent) UnmarshalJSON(b []byte) error {
	type temp TrustedComponent
	type tActions struct {
		TPMGetEventLog common.ActionTarget `json:"#TrustedComponent.TPMGetEventLog"`
	}
	type tLinks struct {
		ActiveSoftwareImage common.Link  `json:"ActiveSoftwareImage"`
		ComponentIntegrity  common.Links `json:"ComponentIntegrity"`
		ComponentsProtected common.Links `json:"ComponentsProtected"`
		IntegratedInto      common.Link  `json:"IntegratedInto"`
		SoftwareImages      common.Links `json:"SoftwareImages"`
	}
	var tmp struct {
		temp
		Actions      tActions
		Links        tLinks
		Certificates common.Link `json:"certificates"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = TrustedComponent(tmp.temp)

	// Extract the links to other entities for later
	t.tPMGetEventLogTarget = tmp.Actions.TPMGetEventLog.Target
	t.activeSoftwareImage = tmp.Links.ActiveSoftwareImage.String()
	t.componentIntegrity = tmp.Links.ComponentIntegrity.ToStrings()
	t.componentsProtected = tmp.Links.ComponentsProtected.ToStrings()
	t.integratedInto = tmp.Links.IntegratedInto.String()
	t.softwareImages = tmp.Links.SoftwareImages.ToStrings()
	t.certificates = tmp.Certificates.String()

	// This is a read/write object, so we need to save the raw object data for later
	t.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (t *TrustedComponent) Update() error {
	readWriteFields := []string{
		"Location",
		"Status",
		"TPM",
	}

	return t.UpdateFromRawData(t, t.rawData, readWriteFields)
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

// TPMGetEventLog shall return the event log for TPM 2.0 devices.
func (t *TrustedComponent) TPMGetEventLog() (*TPMGetEventLogResponse, error) {
	payload := make(map[string]any)

	resp, err := t.PostWithResponse(t.tPMGetEventLogTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, common.CleanupHTTPResponse(resp)
	}

	var result TPMGetEventLogResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ActiveSoftwareImage gets the ActiveSoftwareImage linked resource.
func (t *TrustedComponent) ActiveSoftwareImage(client common.Client) (*SoftwareInventory, error) {
	if t.activeSoftwareImage == "" {
		return nil, nil
	}
	return common.GetObject[SoftwareInventory](client, t.activeSoftwareImage)
}

// ComponentIntegrity gets the ComponentIntegrity linked resources.
func (t *TrustedComponent) ComponentIntegrity(client common.Client) ([]*ComponentIntegrity, error) {
	return common.GetObjects[ComponentIntegrity](client, t.componentIntegrity)
}

// ComponentsProtected gets the ComponentsProtected linked resources.
func (t *TrustedComponent) ComponentsProtected(client common.Client) ([]*common.Entity, error) {
	return common.GetObjects[common.Entity](client, t.componentsProtected)
}

// IntegratedInto gets the IntegratedInto linked resource.
func (t *TrustedComponent) IntegratedInto(client common.Client) (*common.Entity, error) {
	if t.integratedInto == "" {
		return nil, nil
	}
	return common.GetObject[common.Entity](client, t.integratedInto)
}

// SoftwareImages gets the SoftwareImages linked resources.
func (t *TrustedComponent) SoftwareImages(client common.Client) ([]*SoftwareInventory, error) {
	return common.GetObjects[SoftwareInventory](client, t.softwareImages)
}

// Certificates gets the Certificates collection.
func (t *TrustedComponent) Certificates(client common.Client) ([]*Certificate, error) {
	if t.certificates == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Certificate](client, t.certificates)
}

// TPM shall contain TPM-specific information for a trusted component.
type TPM struct {
	// CapabilitiesVendorID shall contain an ASCII string of the 4-byte TCG-defined
	// 'TPM Capabilities Vendor ID' for this trusted component.
	//
	// Version added: v1.1.0
	CapabilitiesVendorID string
	// HardwareInterfaceVendorID shall contain the TCG-defined 'TPM Hardware
	// Interface Vendor ID' for this trusted component with the most significant
	// byte shown first.
	//
	// Version added: v1.1.0
	HardwareInterfaceVendorID string
}

// TPMGetEventLogResponse shall contain the TPM event log.
type TPMGetEventLogResponse struct {
	// EventLog shall contain a Base64-encoded string, with padding characters, of
	// the entire event log defined in the 'Event Logging' section of the 'TCG PC
	// Client Platform Firmware Profile Specification'.
	//
	// Version added: v1.3.0
	EventLog string
	// Oem shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.3.0
	OEM json.RawMessage `json:"Oem"`
}
