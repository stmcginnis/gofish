//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"github.com/stmcginnis/gofish/common"
)

// SoftwareInventory is This Resource contains a single software
// component that this Redfish Service manages.
type SoftwareInventory struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// LowestSupportedVersion is used for the Version property.
	LowestSupportedVersion string
	// Manufacturer is This property shall represent the name of the
	// manufacturer or producer of this software.
	Manufacturer string
	// Oem is This property shall contain the Oem extensions.  All values for
	// properties that this object contains shall conform to the Redfish
	// Specification-described requirements.
	Oem interface{}
	// ReleaseDate is This property shall contain the date of release or
	// production for this software.  If the time of day is unknown, the time
	// of day portion of the property shall contain `00:00:00Z`.
	ReleaseDate string
	// SoftwareID is This property shall represent an implementation-specific
	// label that identifies this software.  This string correlates with a
	// component repository or database.
	SoftwareID string
	// Status is This property shall contain any status or health properties
	// of the Resource.
	Status common.Status
	// UefiDevicePaths is This property shall contain a list UEFI device
	// paths of the components associated with this software inventory item.
	// The UEFI device paths shall be formatted as defined by the UEFI
	// Specification.
	UefiDevicePaths []string
	// Updateable is This property shall indicate whether the Update Service
	// can update this software.  If `true`, the Service can update this
	// software.  If `false`, the Service cannot update this software and the
	// software is for reporting purposes only.
	Updateable bool
	// Version is This property shall contain the version of this software.
	Version string
	// WriteProtected is This property shall indicate whether the software
	// image can be overwritten, where a value `true` shall indicate that the
	// software cannot be altered or overwritten.
	WriteProtected bool
}

// GetSoftwareInventory will get a SoftwareInventory instance from the service.
func GetSoftwareInventory(c common.Client, uri string) (*SoftwareInventory, error) {
	var softwareInventory SoftwareInventory
	return &softwareInventory, softwareInventory.Get(c, uri, &softwareInventory)
}

// ListReferencedSoftwareInventories gets the collection of SoftwareInventory from
// a provided reference.
func ListReferencedSoftwareInventories(c common.Client, link string) ([]*SoftwareInventory, error) {
	var result []*SoftwareInventory
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *SoftwareInventory
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		softwareinventory, err := GetSoftwareInventory(c, link)
		ch <- GetResult{Item: softwareinventory, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
