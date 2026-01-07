//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.1 - #GraphicsController.v1_0_2.GraphicsController

package schemas

import (
	"encoding/json"
)

// GraphicsController shall represent a graphics output device in a Redfish
// implementation.
type GraphicsController struct {
	Entity
	// AssetTag shall contain the user-assigned asset tag, which is an identifying
	// string that tracks the drive for inventory purposes.
	AssetTag string
	// BiosVersion shall contain the version string of the currently installed and
	// running BIOS or firmware for the graphics controller.
	BiosVersion string
	// DriverVersion shall contain the version string of the currently loaded
	// driver for this graphics controller.
	DriverVersion string
	// Location shall contain the location information of the associated graphics
	// controller.
	Location Location
	// Manufacturer shall contain the name of the organization responsible for
	// producing the graphics controller. This organization may be the entity from
	// which the graphics controller is purchased, but this is not necessarily
	// true.
	Manufacturer string
	// Model shall contain the manufacturer-provided model information of this
	// graphics controller.
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the manufacturer-provided part number for the
	// graphics controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type
	// 'PortCollection'.
	ports string
	// SKU shall contain the SKU number for this graphics controller.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the graphics controller.
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the graphics
	// controller.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// pCIeDevice is the URI for PCIeDevice.
	pCIeDevice string
	// processors are the URIs for Processors.
	processors []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a GraphicsController object from the raw JSON.
func (g *GraphicsController) UnmarshalJSON(b []byte) error {
	type temp GraphicsController
	type gLinks struct {
		PCIeDevice Link  `json:"PCIeDevice"`
		Processors Links `json:"Processors"`
	}
	var tmp struct {
		temp
		Links gLinks
		Ports Link `json:"Ports"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*g = GraphicsController(tmp.temp)

	// Extract the links to other entities for later
	g.pCIeDevice = tmp.Links.PCIeDevice.String()
	g.processors = tmp.Links.Processors.ToStrings()
	g.ports = tmp.Ports.String()

	// This is a read/write object, so we need to save the raw object data for later
	g.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (g *GraphicsController) Update() error {
	readWriteFields := []string{
		"AssetTag",
	}

	return g.UpdateFromRawData(g, g.RawData, readWriteFields)
}

// GetGraphicsController will get a GraphicsController instance from the service.
func GetGraphicsController(c Client, uri string) (*GraphicsController, error) {
	return GetObject[GraphicsController](c, uri)
}

// ListReferencedGraphicsControllers gets the collection of GraphicsController from
// a provided reference.
func ListReferencedGraphicsControllers(c Client, link string) ([]*GraphicsController, error) {
	return GetCollectionObjects[GraphicsController](c, link)
}

// PCIeDevice gets the PCIeDevice linked resource.
func (g *GraphicsController) PCIeDevice() (*PCIeDevice, error) {
	if g.pCIeDevice == "" {
		return nil, nil
	}
	return GetObject[PCIeDevice](g.client, g.pCIeDevice)
}

// Processors gets the Processors linked resources.
func (g *GraphicsController) Processors() ([]*Processor, error) {
	return GetObjects[Processor](g.client, g.processors)
}

// Ports gets the Ports collection.
func (g *GraphicsController) Ports() ([]*Port, error) {
	if g.ports == "" {
		return nil, nil
	}
	return GetCollectionObjects[Port](g.client, g.ports)
}
