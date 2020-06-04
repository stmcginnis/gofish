//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"io/ioutil"

	"github.com/stmcginnis/gofish/common"
)

// ResetType describe the type off reset to be issue by the resource
type VirtualMediaType string

const (
	CdVirtualMediaType       VirtualMediaType = "CD"
	DvdVirtualMediaType      VirtualMediaType = "DVD"
	USBStickVirtualMediaType VirtualMediaType = "USBStick"
)

// VirtualMediaConnectedMethod this is the type of virtual media connection method
type VirtualMediaConnectedMethod string

const (
	NotConnectedVirtualMediaConnectedMethod VirtualMediaConnectedMethod = "NotConnected"
	AppletVirtualMediaConnectedMethod       VirtualMediaConnectedMethod = "Applet"
	OemVirtualMediaConnectedMethod          VirtualMediaConnectedMethod = "Oem"
)

// VirtualMedia used to represent a virtual media resource.
// Allows you to connect media types such as CD, DVD, USB Stick
// and others. Media types and methods depend on the hardware manufacturer.
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
	rawData             []byte                      // rawData holds the original serialized JSON
}

// GetRawData get raw data json
func (virtualMedia *VirtualMedia) GetRawData() []byte {
	return virtualMedia.rawData
}

// GetVirtualMedia will get a VirtualMedia instance from the service.
func GetVirtualMedia(c common.Client, uri string) (*VirtualMedia, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var virtualMedia VirtualMedia
	virtualMedia.rawData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(virtualMedia.rawData, &virtualMedia)
	if err != nil {
		return nil, err
	}

	virtualMedia.SetClient(c)
	return &virtualMedia, nil
}

// ListReferencedVirtualMedia gets the collection of VirtualMedia
func ListReferencedVirtualMedia(c common.Client, link string) ([]*VirtualMedia, error) {
	var result []*VirtualMedia
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, virtualMediaLink := range links.ItemLinks {
		virtualMedia, err := GetVirtualMedia(c, virtualMediaLink)
		if err != nil {
			return result, err
		}

		result = append(result, virtualMedia)
	}

	return result, nil
}
