//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.6 - #NVMeDomain.v1_2_0.NVMeDomain

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// NVMeDomain Properties for the Domain.
type NVMeDomain struct {
	common.Entity
	// ANAGroupId shall contain the ANA group id which applies to all namespaces
	// within the domain. This corresponds to the value in the ANAGroupId field in
	// volume.
	//
	// Version added: v1.2.0
	ANAGroupID *float64 `json:"ANAGroupId,omitempty"`
	// AvailableFirmwareImages is a collection of available firmware images.
	AvailableFirmwareImages []NVMeFirmwareImage
	// AvailableFirmwareImages@odata.count
	AvailableFirmwareImagesCount int `json:"AvailableFirmwareImages@odata.count"`
	// DomainContents shall contain the members of the domain.
	//
	// Version added: v1.2.0
	DomainContents DomainContents
	// DomainMembers The members of the domain.
	DomainMembers []common.Resource
	// DomainMembers@odata.count
	DomainMembersCount int `json:"DomainMembers@odata.count"`
	// FirmwareImages shall contain an array of pointers to available firmware
	// images.
	//
	// Version added: v1.2.0
	FirmwareImages []redfish.SoftwareInventory
	// FirmwareImages@odata.count
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
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TotalDomainCapacityBytes shall contain the total capacity in bytes of this
	// NVMe Domain.
	TotalDomainCapacityBytes *int `json:",omitempty"`
	// UnallocatedDomainCapacityBytes shall contain the total unallocated capacity
	// in bytes of this NVMe Domain.
	UnallocatedDomainCapacityBytes *int `json:",omitempty"`
	// associatedDomains are the URIs for AssociatedDomains.
	associatedDomains []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a NVMeDomain object from the raw JSON.
func (n *NVMeDomain) UnmarshalJSON(b []byte) error {
	type temp NVMeDomain
	type nLinks struct {
		AssociatedDomains common.Links `json:"AssociatedDomains"`
	}
	var tmp struct {
		temp
		Links nLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*n = NVMeDomain(tmp.temp)

	// Extract the links to other entities for later
	n.associatedDomains = tmp.Links.AssociatedDomains.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	n.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (n *NVMeDomain) Update() error {
	readWriteFields := []string{
		"AvailableFirmwareImages",
		"AvailableFirmwareImages@odata.count",
		"DomainContents",
		"DomainMembers",
		"DomainMembers@odata.count",
		"FirmwareImages@odata.count",
		"Status",
	}

	return n.UpdateFromRawData(n, n.rawData, readWriteFields)
}

// GetNVMeDomain will get a NVMeDomain instance from the service.
func GetNVMeDomain(c common.Client, uri string) (*NVMeDomain, error) {
	return common.GetObject[NVMeDomain](c, uri)
}

// ListReferencedNVMeDomains gets the collection of NVMeDomain from
// a provided reference.
func ListReferencedNVMeDomains(c common.Client, link string) ([]*NVMeDomain, error) {
	return common.GetCollectionObjects[NVMeDomain](c, link)
}

// AssociatedDomains gets the AssociatedDomains linked resources.
func (n *NVMeDomain) AssociatedDomains(client common.Client) ([]*NVMeDomain, error) {
	return common.GetObjects[NVMeDomain](client, n.associatedDomains)
}

// DomainContents shall contain properties that define the contents of the
// domain.
type DomainContents struct {
	// Controllers Contains the current controllers that are part of this domain.
	// These can be IO, Admin, or discovery controllers.
	//
	// Version added: v1.2.0
	Controllers []redfish.StorageController
	// Controllers@odata.count
	ControllersCount int `json:"Controllers@odata.count"`
	// Namespaces Contains the current namespaces that are part of this domain.
	// These can be IO, Admin, or discovery controllers.
	//
	// Version added: v1.2.0
	Namespaces []Volume
	// Namespaces@odata.count
	NamespacesCount int `json:"Namespaces@odata.count"`
}
