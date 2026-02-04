//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/NVMeDomain.v1_2_0.json
// 1.2.6 - #NVMeDomain.v1_2_0.NVMeDomain

package schemas

import (
	"encoding/json"
)

// NVMeDomain Properties for the Domain.
type NVMeDomain struct {
	Entity
	// ANAGroupID shall contain the ANA group id which applies to all namespaces
	// within the domain. This corresponds to the value in the ANAGroupId field in
	// volume.
	//
	// Version added: v1.2.0
	ANAGroupID *float64 `json:"ANAGroupId,omitempty"`
	// AvailableFirmwareImages is a collection of available firmware images.
	availableFirmwareImages []string
	// AvailableFirmwareImagesCount
	AvailableFirmwareImagesCount int `json:"AvailableFirmwareImages@odata.count"`
	// DomainContents shall contain the members of the domain.
	//
	// Version added: v1.2.0
	DomainContents DomainContents
	// DomainMembers The members of the domain.
	domainMembers []string
	// DomainMembersCount
	DomainMembersCount int `json:"DomainMembers@odata.count"`
	// FirmwareImages shall contain an array of pointers to available firmware
	// images.
	//
	// Version added: v1.2.0
	firmwareImages []string
	// FirmwareImagesCount
	FirmwareImagesCount int `json:"FirmwareImages@odata.count"`
	// MaxNamespacesSupportedPerController shall contain the maximum number of
	// namespace attachments supported in this NVMe Domain. If there are no limits
	// imposed, this property should not be implemented.
	//
	// Version added: v1.2.0
	MaxNamespacesSupportedPerController *float64 `json:",omitempty"`
	// MaximumCapacityPerEnduranceGroupBytes shall contain the maximum capacity per
	// endurance group in bytes of this NVMe Domain.
	MaximumCapacityPerEnduranceGroupBytes *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// TotalDomainCapacityBytes shall contain the total capacity in bytes of this
	// NVMe Domain.
	TotalDomainCapacityBytes *int `json:",omitempty"`
	// UnallocatedDomainCapacityBytes shall contain the total unallocated capacity
	// in bytes of this NVMe Domain.
	UnallocatedDomainCapacityBytes *int `json:",omitempty"`
	// associatedDomains are the URIs for AssociatedDomains.
	associatedDomains []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a NVMeDomain object from the raw JSON.
func (n *NVMeDomain) UnmarshalJSON(b []byte) error {
	type temp NVMeDomain
	type nLinks struct {
		AssociatedDomains Links `json:"AssociatedDomains"`
	}
	var tmp struct {
		temp
		Links                   nLinks
		AvailableFirmwareImages Links `json:"AvailableFirmwareImages"`
		DomainMembers           Links `json:"DomainMembers"`
		FirmwareImages          Links `json:"FirmwareImages"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*n = NVMeDomain(tmp.temp)

	// Extract the links to other entities for later
	n.associatedDomains = tmp.Links.AssociatedDomains.ToStrings()
	n.availableFirmwareImages = tmp.AvailableFirmwareImages.ToStrings()
	n.domainMembers = tmp.DomainMembers.ToStrings()
	n.firmwareImages = tmp.FirmwareImages.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	n.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (n *NVMeDomain) Update() error {
	readWriteFields := []string{
		"DomainMembers",
	}

	return n.UpdateFromRawData(n, n.RawData, readWriteFields)
}

// GetNVMeDomain will get a NVMeDomain instance from the service.
func GetNVMeDomain(c Client, uri string) (*NVMeDomain, error) {
	return GetObject[NVMeDomain](c, uri)
}

// ListReferencedNVMeDomains gets the collection of NVMeDomain from
// a provided reference.
func ListReferencedNVMeDomains(c Client, link string) ([]*NVMeDomain, error) {
	return GetCollectionObjects[NVMeDomain](c, link)
}

// AssociatedDomains gets the AssociatedDomains linked resources.
func (n *NVMeDomain) AssociatedDomains() ([]*NVMeDomain, error) {
	return GetObjects[NVMeDomain](n.client, n.associatedDomains)
}

// AvailableFirmwareImages gets the AvailableFirmwareImages linked resources.
func (n *NVMeDomain) AvailableFirmwareImages() ([]*NVMeFirmwareImage, error) {
	return GetObjects[NVMeFirmwareImage](n.client, n.availableFirmwareImages)
}

// DomainMembers gets the DomainMembers linked resources.
func (n *NVMeDomain) DomainMembers() ([]*Resource, error) {
	return GetObjects[Resource](n.client, n.domainMembers)
}

// FirmwareImages gets the FirmwareImages linked resources.
func (n *NVMeDomain) FirmwareImages() ([]*SoftwareInventory, error) {
	return GetObjects[SoftwareInventory](n.client, n.firmwareImages)
}

// DomainContents shall contain properties that define the contents of the
// domain.
type DomainContents struct {
	// Controllers Contains the current controllers that are part of this domain.
	// These can be IO, Admin, or discovery controllers.
	//
	// Version added: v1.2.0
	controllers []string
	// ControllersCount
	ControllersCount int `json:"Controllers@odata.count"`
	// Namespaces Contains the current namespaces that are part of this domain.
	// These can be IO, Admin, or discovery controllers.
	//
	// Version added: v1.2.0
	namespaces []string
	// NamespacesCount
	NamespacesCount int `json:"Namespaces@odata.count"`
}

// UnmarshalJSON unmarshals a DomainContents object from the raw JSON.
func (d *DomainContents) UnmarshalJSON(b []byte) error {
	type temp DomainContents
	var tmp struct {
		temp
		Controllers Links `json:"Controllers"`
		Namespaces  Links `json:"Namespaces"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*d = DomainContents(tmp.temp)

	// Extract the links to other entities for later
	d.controllers = tmp.Controllers.ToStrings()
	d.namespaces = tmp.Namespaces.ToStrings()

	return nil
}

// Controllers gets the Controllers linked resources.
func (d *DomainContents) Controllers(client Client) ([]*StorageController, error) {
	return GetObjects[StorageController](client, d.controllers)
}

// Namespaces gets the Namespaces linked resources.
func (d *DomainContents) Namespaces(client Client) ([]*Volume, error) {
	return GetObjects[Volume](client, d.namespaces)
}
