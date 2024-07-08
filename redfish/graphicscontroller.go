//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// GraphicsController shall represent a graphics output device in a Redfish implementation.
type GraphicsController struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AssetTag shall contain the user-assigned asset tag, which is an identifying string that tracks the drive for
	// inventory purposes.
	AssetTag string
	// BiosVersion shall contain the version string of the currently installed and running BIOS or firmware for the
	// graphics controller.
	BiosVersion string
	// Description provides a description of this resource.
	Description string
	// DriverVersion shall contain the version string of the currently loaded driver for this graphics controller.
	DriverVersion string
	// Location shall contain the location information of the associated graphics controller.
	Location common.Location
	// Manufacturer shall contain the name of the organization responsible for producing the graphics controller. This
	// organization may be the entity from which the graphics controller is purchased, but this is not necessarily
	// true.
	Manufacturer string
	// Model shall contain the manufacturer-provided model information of this graphics controller.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the manufacturer-provided part number for the graphics controller.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection.
	ports string
	// SKU shall contain the SKU number for this graphics controller.
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies the graphics controller.
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the graphics controller.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	pcieDevice string
	processors []string
	// ProcessorsCount is the number of processors trhat this graphics controller contains.
	ProcessorsCount int
}

// UnmarshalJSON unmarshals a GraphicsController object from the raw JSON.
func (graphicscontroller *GraphicsController) UnmarshalJSON(b []byte) error {
	type temp GraphicsController
	type Links struct {
		// PCIeDevice shall contain a link to a resource of type PCIeDevice that represents this graphics controller.
		PCIeDevice common.Link
		// Processors shall contain an array of links to resources of type Processor that represent the processors that
		// this graphics controller contains.
		Processors      common.Links
		ProcessorsCount int `json:"Processors@odata.count"`
	}
	var t struct {
		temp
		Ports common.Link
		Links Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*graphicscontroller = GraphicsController(t.temp)

	// Extract the links to other entities for later
	graphicscontroller.ports = t.Ports.String()
	graphicscontroller.pcieDevice = t.Links.PCIeDevice.String()
	graphicscontroller.processors = t.Links.Processors.ToStrings()
	graphicscontroller.ProcessorsCount = t.Links.ProcessorsCount

	// This is a read/write object, so we need to save the raw object data for later
	graphicscontroller.rawData = b

	return nil
}

// Ports get the ports associated with this graphics controller.
func (graphicscontroller *GraphicsController) Ports() ([]*Port, error) {
	return ListReferencedPorts(graphicscontroller.GetClient(), graphicscontroller.ports)
}

// PCIeDevice gets the PCIeDevice for this graphics controller.
func (graphicscontroller *GraphicsController) PCIeDevice() (*PCIeDevice, error) {
	if graphicscontroller.pcieDevice == "" {
		return nil, nil
	}
	return GetPCIeDevice(graphicscontroller.GetClient(), graphicscontroller.pcieDevice)
}

// Processors gets this graphics controllers processors.
func (graphicscontroller *GraphicsController) Processors() ([]*Processor, error) {
	return common.GetObjects[Processor](graphicscontroller.GetClient(), graphicscontroller.processors)
}

// Update commits updates to this object's properties to the running system.
func (graphicscontroller *GraphicsController) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(GraphicsController)
	original.UnmarshalJSON(graphicscontroller.rawData)

	readWriteFields := []string{
		"AssetTag",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(graphicscontroller).Elem()

	return graphicscontroller.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetGraphicsController will get a GraphicsController instance from the service.
func GetGraphicsController(c common.Client, uri string) (*GraphicsController, error) {
	return common.GetObject[GraphicsController](c, uri)
}

// ListReferencedGraphicsControllers gets the collection of GraphicsController from
// a provided reference.
func ListReferencedGraphicsControllers(c common.Client, link string) ([]*GraphicsController, error) {
	return common.GetCollectionObjects[GraphicsController](c, link)
}
