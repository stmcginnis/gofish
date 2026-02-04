//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/TrustedComponent.v1_4_0.json
// 2025.2 - #TrustedComponent.v1_4_0.TrustedComponent

package schemas

import (
	"encoding/json"
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
	Entity
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
	Location Location
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
	// OEM shall contain the OEM extensions. All values for properties that this
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
	Status Status
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
	owner string
	// softwareImages are the URIs for SoftwareImages.
	softwareImages []string
}

// UnmarshalJSON unmarshals a TrustedComponent object from the raw JSON.
func (t *TrustedComponent) UnmarshalJSON(b []byte) error {
	type temp TrustedComponent
	type tActions struct {
		TPMGetEventLog ActionTarget `json:"#TrustedComponent.TPMGetEventLog"`
	}
	type tLinks struct {
		ActiveSoftwareImage Link  `json:"ActiveSoftwareImage"`
		ComponentIntegrity  Links `json:"ComponentIntegrity"`
		ComponentsProtected Links `json:"ComponentsProtected"`
		IntegratedInto      Link  `json:"IntegratedInto"`
		Owner               Link  `json:"Owner"`
		SoftwareImages      Links `json:"SoftwareImages"`
	}
	var tmp struct {
		temp
		Actions      tActions
		Links        tLinks
		Certificates Link `json:"Certificates"`
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
	t.owner = tmp.Links.Owner.String()
	t.softwareImages = tmp.Links.SoftwareImages.ToStrings()
	t.certificates = tmp.Certificates.String()

	return nil
}

// GetTrustedComponent will get a TrustedComponent instance from the service.
func GetTrustedComponent(c Client, uri string) (*TrustedComponent, error) {
	return GetObject[TrustedComponent](c, uri)
}

// ListReferencedTrustedComponents gets the collection of TrustedComponent from
// a provided reference.
func ListReferencedTrustedComponents(c Client, link string) ([]*TrustedComponent, error) {
	return GetCollectionObjects[TrustedComponent](c, link)
}

// This action shall return the event log for TPM 2.0 devices.
func (t *TrustedComponent) TPMGetEventLog() (*TPMGetEventLogResponse, error) {
	payload := make(map[string]any)

	resp, err := t.PostWithResponse(t.tPMGetEventLogTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result TPMGetEventLogResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// ActiveSoftwareImage gets the ActiveSoftwareImage linked resource.
func (t *TrustedComponent) ActiveSoftwareImage() (*SoftwareInventory, error) {
	if t.activeSoftwareImage == "" {
		return nil, nil
	}
	return GetObject[SoftwareInventory](t.client, t.activeSoftwareImage)
}

// ComponentIntegrity gets the ComponentIntegrity linked resources.
func (t *TrustedComponent) ComponentIntegrity() ([]*ComponentIntegrity, error) {
	return GetObjects[ComponentIntegrity](t.client, t.componentIntegrity)
}

// ComponentsProtected gets the ComponentsProtected linked resources.
func (t *TrustedComponent) ComponentsProtected() ([]*Entity, error) {
	return GetObjects[Entity](t.client, t.componentsProtected)
}

// IntegratedInto gets the IntegratedInto linked resource.
func (t *TrustedComponent) IntegratedInto() (*Entity, error) {
	if t.integratedInto == "" {
		return nil, nil
	}
	return GetObject[Entity](t.client, t.integratedInto)
}

// Owner gets the Owner linked resource.
func (t *TrustedComponent) Owner() (*Entity, error) {
	if t.owner == "" {
		return nil, nil
	}
	return GetObject[Entity](t.client, t.owner)
}

// SoftwareImages gets the SoftwareImages linked resources.
func (t *TrustedComponent) SoftwareImages() ([]*SoftwareInventory, error) {
	return GetObjects[SoftwareInventory](t.client, t.softwareImages)
}

// Certificates gets the Certificates collection.
func (t *TrustedComponent) Certificates() ([]*Certificate, error) {
	if t.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](t.client, t.certificates)
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
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.3.0
	OEM json.RawMessage `json:"Oem"`
}
