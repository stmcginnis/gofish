//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"github.com/LRichi/gofish/common"
)

// ResetType describe the type off reset to be issue by the resource
type VirtualMediaType string

const (
	CdVirtualMediaType       VirtualMediaType = "CD"
	DvdVirtualMediaType      VirtualMediaType = "DVD"
	USBStickVirtualMediaType VirtualMediaType = "USBStick"
)

type VirtualMediaConnectedMethod string

const (
	NotConnectedVirtualMediaConnectedMethod VirtualMediaConnectedMethod = "NotConnected"
	AppletVirtualMediaConnectedMethod       VirtualMediaConnectedMethod = "Applet"
	OemVirtualMediaConnectedMethod          VirtualMediaConnectedMethod = "Oem"
)

// VirtualMedia is used to represent virtual media resource.
type VirtualMedia struct {
	common.Entity

	ODataContext        string                      `json:"@odata.context"` // ODataContext is the odata context.
	ODataEtag           string                      `json:"@odata.etag"`    // ODataEtag is the odata etag.
	ODataType           string                      `json:"@odata.type"`    // ODataType is the odata type.
	Description         string                      `json:"Description"`    // Description provides a description of this resource.
	ConnectedVia        VirtualMediaConnectedMethod `json:"ConnectedVia"`   // ConnectedVia connected via type
	Image               string                      `json:"Image"`          // Image endpoint for get image
	ImageName           string                      `json:"ImageName"`      // ImageName image file name
	WriteProtected      bool                        `json:"WriteProtected"` // WriteProtected ...
	Inserted            bool                        `json:"Inserted"`       // Inserted status of connect image
	SupportedMediaTypes []VirtualMediaType          `json:"MediaTypes"`     // MediaTypes allowed media types
}

// GetVirtualMedia will get a VirtualMedia instance from the service.
func GetVirtualMedia(c common.Client, uri string) (*VirtualMedia, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var virtualMedia VirtualMedia
	err = json.NewDecoder(resp.Body).Decode(&virtualMedia)
	if err != nil {
		return nil, err
	}

	virtualMedia.SetClient(c)
	return &virtualMedia, nil
}

// VirtualMediaCollection is a collection of virtual media resources.
type VirtualMediaCollection struct {
	common.Entity
	common.LinksCollection

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string `json:"Description"`
}

// GetVirtualMediaCollection will get a VirtualMedia collection from the Redfish service.
func GetVirtualMediaCollection(c common.Client, uri string) (*VirtualMediaCollection, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var virtualMediaCollection VirtualMediaCollection
	err = json.NewDecoder(resp.Body).Decode(&virtualMediaCollection)

	if err != nil {
		return nil, err
	}

	virtualMediaCollection.SetClient(c)

	return &virtualMediaCollection, nil
}

// VirtualMedia get this system's virtual media
func (vmCollection *VirtualMediaCollection) VirtualMedia() ([]*VirtualMedia, error) {
	var result []*VirtualMedia

	for _, link := range vmCollection.LinksCollection.Members.ToStrings() {
		virtualMedia, err := GetVirtualMedia(vmCollection.Client, link)
		if err != nil {
			return nil, err
		}

		result = append(result, virtualMedia)
	}

	return result, nil
}
