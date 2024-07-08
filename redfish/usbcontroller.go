//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// USBController shall represent a USB controller in a Redfish implementation.
type USBController struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Manufacturer shall contain the name of the organization responsible for producing the USB controller. This
	// organization may be the entity from which the USB controller is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the manufacturer-provided model information of this USB controller.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the manufacturer-provided part number for the USB controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection.
	ports string
	// SKU shall contain the SKU number for this USB controller.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the USB controller.
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the USB controller.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status

	pcieDevice string
	processors []string
	// ProcessorsCount is the number of processors that can use this USB controller.
	ProcessorsCount int
}

// UnmarshalJSON unmarshals a USBController object from the raw JSON.
func (usbcontroller *USBController) UnmarshalJSON(b []byte) error {
	type temp USBController
	type Links struct {
		PCIeDevice      common.Link
		Processors      common.Links
		ProcessorsCount int `json:"Processors@odata.count"`
	}
	var t struct {
		temp
		Links Links
		Ports common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*usbcontroller = USBController(t.temp)
	usbcontroller.pcieDevice = t.Links.PCIeDevice.String()
	usbcontroller.processors = t.Links.Processors.ToStrings()

	usbcontroller.ports = t.Ports.String()

	return nil
}

// PCIeDevice gets the PCIeDevice for this USB controller.
func (usbcontroller *USBController) PCIeDevice() (*PCIeDevice, error) {
	if usbcontroller.pcieDevice == "" {
		return nil, nil
	}
	return GetPCIeDevice(usbcontroller.GetClient(), usbcontroller.pcieDevice)
}

// Processors gets the processors that can utilize this USB controller.
func (usbcontroller *USBController) Processors() ([]*Processor, error) {
	return common.GetObjects[Processor](usbcontroller.GetClient(), usbcontroller.processors)
}

// Ports gets the ports of the USB controller.
func (usbcontroller *USBController) Ports() ([]*Port, error) {
	return ListReferencedPorts(usbcontroller.GetClient(), usbcontroller.ports)
}

// GetUSBController will get a USBController instance from the service.
func GetUSBController(c common.Client, uri string) (*USBController, error) {
	return common.GetObject[USBController](c, uri)
}

// ListReferencedUSBControllers gets the collection of USBController from
// a provided reference.
func ListReferencedUSBControllers(c common.Client, link string) ([]*USBController, error) {
	return common.GetCollectionObjects[USBController](c, link)
}
