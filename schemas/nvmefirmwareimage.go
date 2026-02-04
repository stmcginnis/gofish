//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/NVMeFirmwareImage.v1_2_0.json
// 1.2.6 - #NVMeFirmwareImage.v1_2_0.NVMeFirmwareImage

package schemas

import (
	"encoding/json"
)

type NVMeDeviceType string

const (
	// DriveNVMeDeviceType Specifies an device type of Drive, indicating a NVMe
	// device that presents as an NVMe SSD device.
	DriveNVMeDeviceType NVMeDeviceType = "Drive"
	// JBOFNVMeDeviceType Specifies an device type of JBOF, indicating a NVMe
	// device that presents as an NVMe smart enclosure for NVMe devices, typically
	// NVMe Drives.
	JBOFNVMeDeviceType NVMeDeviceType = "JBOF"
	// FabricAttachArrayNVMeDeviceType Specifies an NVMe device type of
	// FabricAttachArray, indicating a NVMe device that presents an NVMe front-end
	// that abstracts the back end storage, typically with multiple options for
	// availability and protection.
	FabricAttachArrayNVMeDeviceType NVMeDeviceType = "FabricAttachArray"
)

// NVMeFirmwareImage NVMe Domain firmware image information.
type NVMeFirmwareImage struct {
	Entity
	// FirmwareVersion shall contain the firmware version of the available NVMe
	// firmware image.
	FirmwareVersion string
	// NVMeDeviceType shall specify the type of NVMe device for this NVMe firmware
	// image.
	NVMeDeviceType NVMeDeviceType
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Vendor shall include the name of the manufacturer or vendor associate with
	// this NVMe firmware image.
	Vendor string
}

// GetNVMeFirmwareImage will get a NVMeFirmwareImage instance from the service.
func GetNVMeFirmwareImage(c Client, uri string) (*NVMeFirmwareImage, error) {
	return GetObject[NVMeFirmwareImage](c, uri)
}

// ListReferencedNVMeFirmwareImages gets the collection of NVMeFirmwareImage from
// a provided reference.
func ListReferencedNVMeFirmwareImages(c Client, link string) ([]*NVMeFirmwareImage, error) {
	return GetCollectionObjects[NVMeFirmwareImage](c, link)
}
