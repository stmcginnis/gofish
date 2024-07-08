//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// DomainContents shall contain properties that define the contents of the domain.
type DomainContents struct {
	// Controllers contains the current controllers that are part of this domain. These can be IO, Admin, or discovery
	// controllers.
	controllers []string
	// ControllersCount is the number of controllers
	ControllersCount int `json:"Controllers@odata.count"`
	// Namespaces contains the current namespaces that are part of this domain. These can be IO, Admin, or discovery
	// controllers.
	Namespaces []Volume
	// NamespacesCount is the number of namespaces.
	NamespacesCount int `json:"Namespaces@odata.count"`
}

// UnmarshalJSON unmarshals a DomainContents object from the raw JSON.
func (domaincontents *DomainContents) UnmarshalJSON(b []byte) error {
	type temp DomainContents
	var t struct {
		temp
		Controllers common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*domaincontents = DomainContents(t.temp)

	// Extract the links to other entities for later
	domaincontents.controllers = t.Controllers.ToStrings()

	return nil
}

// NVMeDomain Properties for the Domain.
type NVMeDomain struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ANAGroupID shall contain the ANA group id which applies to all namespaces within the domain. This corresponds to
	// the value in the ANAGroupID field in volume.
	ANAGroupID float64
	// AvailableFirmwareImages is a collection of available firmware images.
	AvailableFirmwareImages []NVMeFirmwareImage
	// AvailableFirmwareImagesCount is the number of available firmware images.
	AvailableFirmwareImagesCount int `json:"AvailableFirmwareImages@odata.count"`
	// Description provides a description of this resource.
	Description string
	// DomainContents shall contain the members of the domain.
	DomainContents DomainContents
	// DomainMembers are the members of the domain.
	domainMembers []string
	// DomainMembers@odata.count
	DomainMembersCount int `json:"DomainMembers@odata.count"`
	// FirmwareImages shall contain an array of pointers to available firmware images.
	firmwareImages []string
	// FirmwareImagesCount is the number of firmware images.
	FirmwareImagesCount int `json:"FirmwareImages@odata.count"`
	// MaxNamespacesSupportedPerController shall contain the maximum number of namespace attachments supported in this
	// NVMe Domain. If there are no limits imposed, this property should not be implemented.
	MaxNamespacesSupportedPerController float64
	// MaximumCapacityPerEnduranceGroupBytes shall contain the maximum capacity per endurance group in bytes of this
	// NVMe Domain.
	MaximumCapacityPerEnduranceGroupBytes int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TotalDomainCapacityBytes shall contain the total capacity in bytes of this NVMe Domain.
	TotalDomainCapacityBytes int64
	// UnallocatedDomainCapacityBytes shall contain the total unallocated capacity in bytes of this NVMe Domain.
	UnallocatedDomainCapacityBytes int64
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	associatedDomains []string
	// AssociatedDomainsCount is the number of associated domains.
	AssociatedDomainsCount int
}

// UnmarshalJSON unmarshals a NVMeDomain object from the raw JSON.
func (nvmedomain *NVMeDomain) UnmarshalJSON(b []byte) error {
	type temp NVMeDomain
	var t struct {
		temp
		Links struct {
			AssociatedDomains      common.Links
			AssociatedDomainsCount int `json:"AssociatedDomains@odata.count"`
		}
		DomainMembers  common.Links
		FirmwareImages common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*nvmedomain = NVMeDomain(t.temp)

	// Extract the links to other entities for later
	nvmedomain.domainMembers = t.DomainMembers.ToStrings()
	nvmedomain.firmwareImages = t.FirmwareImages.ToStrings()
	nvmedomain.associatedDomains = t.Links.AssociatedDomains.ToStrings()
	nvmedomain.AssociatedDomainsCount = t.Links.AssociatedDomainsCount

	// This is a read/write object, so we need to save the raw object data for later
	nvmedomain.rawData = b

	return nil
}

// AssociatedDomains gets the NVMeDomains associated with this domain.
func (nvmedomain *NVMeDomain) AssociatedDomains() ([]*NVMeDomain, error) {
	return common.GetObjects[NVMeDomain](nvmedomain.GetClient(), nvmedomain.associatedDomains)
}

// Update commits updates to this object's properties to the running system.
func (nvmedomain *NVMeDomain) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(NVMeDomain)
	original.UnmarshalJSON(nvmedomain.rawData)

	readWriteFields := []string{
		"DomainMembers",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(nvmedomain).Elem()

	return nvmedomain.Entity.Update(originalElement, currentElement, readWriteFields)
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
