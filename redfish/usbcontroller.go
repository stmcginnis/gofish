//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.1 - #USBController.v1_0_1.USBController

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// USBController shall represent a USB controller in a Redfish implementation.
type USBController struct {
	common.Entity
	// Manufacturer shall contain the name of the organization responsible for
	// producing the USB controller. This organization may be the entity from which
	// the USB controller is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the manufacturer-provided model information of this USB
	// controller.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the manufacturer-provided part number for the USB
	// controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type
	// 'PortCollection'.
	ports string
	// SKU shall contain the SKU number for this USB controller.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the USB controller.
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the USB controller.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// pCIeDevice is the URI for PCIeDevice.
	pCIeDevice string
	// processors are the URIs for Processors.
	processors []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a USBController object from the raw JSON.
func (u *USBController) UnmarshalJSON(b []byte) error {
	type temp USBController
	type uLinks struct {
		PCIeDevice common.Link  `json:"PCIeDevice"`
		Processors common.Links `json:"Processors"`
	}
	var tmp struct {
		temp
		Links uLinks
		Ports common.Link `json:"ports"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*u = USBController(tmp.temp)

	// Extract the links to other entities for later
	u.pCIeDevice = tmp.Links.PCIeDevice.String()
	u.processors = tmp.Links.Processors.ToStrings()
	u.ports = tmp.Ports.String()

	// This is a read/write object, so we need to save the raw object data for later
	u.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (u *USBController) Update() error {
	readWriteFields := []string{
		"Ports",
		"Status",
	}

	return u.UpdateFromRawData(u, u.rawData, readWriteFields)
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

// PCIeDevice gets the PCIeDevice linked resource.
func (u *USBController) PCIeDevice(client common.Client) (*PCIeDevice, error) {
	if u.pCIeDevice == "" {
		return nil, nil
	}
	return common.GetObject[PCIeDevice](client, u.pCIeDevice)
}

// Processors gets the Processors linked resources.
func (u *USBController) Processors(client common.Client) ([]*Processor, error) {
	return common.GetObjects[Processor](client, u.processors)
}

// Ports gets the Ports collection.
func (u *USBController) Ports(client common.Client) ([]*Port, error) {
	if u.ports == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Port](client, u.ports)
}
