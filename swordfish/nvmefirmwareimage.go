//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// NVMeDeviceType is the type of NVMe device.
type NVMeDeviceType string

const (
	// DriveNVMeDeviceType specifies a device type of Drive, indicating a NVMe device that presents as an NVMe SSD device.
	DriveNVMeDeviceType NVMeDeviceType = "Drive"
	// FabricAttachArrayNVMeDeviceType specifies an NVMe device type of FabricAttachArray,
	// indicating a NVMe device that presents an NVMe front-end that abstracts the back end
	// storage, typically with multiple options for availability and protection.
	FabricAttachArrayNVMeDeviceType NVMeDeviceType = "FabricAttachArray"
	// JBOFNVMeDeviceType specifies a device type of JBOF, indicating a NVMe device that
	// presents as an NVMe smart enclosure for NVMe devices, typically NVMe Drives.
	JBOFNVMeDeviceType NVMeDeviceType = "JBOF"
)

// NVMeFirmwareImage NVMe Domain firmware image information.
type NVMeFirmwareImage struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// FirmwareVersion shall contain the firmware version of the available NVMe firmware image.
	FirmwareVersion string
	// NVMeDeviceType shall specify the type of NVMe device for this NVMe firmware image.
	NVMeDeviceType NVMeDeviceType
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Vendor shall include the name of the manufacturer or vendor associate with this NVMe firmware image.
	Vendor string
}

// GetNVMeFirmwareImage will get a NVMeFirmwareImage instance from the service.
func GetNVMeFirmwareImage(c common.Client, uri string) (*NVMeFirmwareImage, error) {
	return common.GetObject[NVMeFirmwareImage](c, uri)
}

// ListReferencedNVMeFirmwareImages gets the collection of NVMeFirmwareImage from
// a provided reference.
func ListReferencedNVMeFirmwareImages(c common.Client, link string) ([]*NVMeFirmwareImage, error) {
	return common.GetCollectionObjects[NVMeFirmwareImage](c, link)
}
