//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// DeviceType is
type DeviceType string

const (
	// SingleFunctionDeviceType A single-function PCIe device.
	SingleFunctionDeviceType DeviceType = "SingleFunction"
	// MultiFunctionDeviceType A multi-function PCIe device.
	MultiFunctionDeviceType DeviceType = "MultiFunction"
	// SimulatedDeviceType A PCIe device which is not currently physically
	// present, but is being simulated by the PCIe infrastructure.
	SimulatedDeviceType DeviceType = "Simulated"
)

// PCIeTypes is the type of PCIe device.
type PCIeTypes string

const (
	// Gen1PCIeTypes A PCIe v1.0 slot.
	Gen1PCIeTypes PCIeTypes = "Gen1"
	// Gen2PCIeTypes A PCIe v2.0 slot.
	Gen2PCIeTypes PCIeTypes = "Gen2"
	// Gen3PCIeTypes A PCIe v3.0 slot.
	Gen3PCIeTypes PCIeTypes = "Gen3"
	// Gen4PCIeTypes A PCIe v4.0 slot.
	Gen4PCIeTypes PCIeTypes = "Gen4"
	// Gen5PCIeTypes A PCIe v5.0 slot.
	Gen5PCIeTypes PCIeTypes = "Gen5"
)

// PCIeDevice is used to represent a PCIeDevice attached to a System.
type PCIeDevice struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Assembly shall be a link to a resource of type Assembly.
	assembly string
	// AssetTag is used to track the PCIe device for inventory purposes.
	AssetTag string
	// Description provides a description of this resource.
	Description string
	// DeviceType shall be the device type of the PCIe device such as
	// SingleFunction or MultiFunction.
	DeviceType DeviceType
	// FirmwareVersion shall be the firmware version of the PCIe device.
	FirmwareVersion string
	// Manufacturer shall be the name of the organization responsible for
	// producing the PCIe device. This organization might be the entity from
	// whom the PCIe device is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall be the name by which the manufacturer generally refers to the
	// PCIe device.
	Model string
	// PCIeInterface is used to connect this PCIe Device to its host or
	// upstream switch.
	PCIeInterface PCIeInterface
	// PartNumber shall be a part number assigned by the organization that is
	// responsible for producing or manufacturing the PCIe device.
	PartNumber string
	// SKU shall be the stock-keeping unit number for this PCIe device.
	SKU string
	// SerialNumber is used to identify the PCIe device.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// Chassis shall reference a resource of type Chassis that represents the
	// physical container associated with this resource.
	chassis []string
	// ChassisCount is the number of number of associated chassis.
	ChassisCount int
	// PCIeFunctions shall be a reference to the resources that this device
	// exposes and shall reference a resource of type PCIeFunction.
	pcieFunctions []string
	// PCIeFunctionsCount is the number of PCIeFunctions.
	PCIeFunctionsCount int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a PCIeDevice object from the raw JSON.
func (pciedevice *PCIeDevice) UnmarshalJSON(b []byte) error {
	type temp PCIeDevice
	type links struct {
		Chassis            common.Links
		ChassisCount       int `json:"Chassis@odata.count"`
		PCIeFunctions      common.Links
		PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
	}
	var t struct {
		temp
		Assembly common.Link
		Links    links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*pciedevice = PCIeDevice(t.temp)

	// Extract the links to other entities for later
	pciedevice.assembly = t.Assembly.String()
	pciedevice.chassis = t.Links.Chassis.ToStrings()
	pciedevice.ChassisCount = t.Links.ChassisCount
	pciedevice.pcieFunctions = t.Links.PCIeFunctions.ToStrings()
	pciedevice.PCIeFunctionsCount = t.Links.PCIeFunctionsCount

	// This is a read/write object, so we need to save the raw object data for later
	pciedevice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (pciedevice *PCIeDevice) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(PCIeDevice)
	err := original.UnmarshalJSON(pciedevice.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"AssetTag",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(pciedevice).Elem()

	return pciedevice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetPCIeDevice will get a PCIeDevice instance from the service.
func GetPCIeDevice(c common.Client, uri string) (*PCIeDevice, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var pciedevice PCIeDevice
	err = json.NewDecoder(resp.Body).Decode(&pciedevice)
	if err != nil {
		return nil, err
	}

	pciedevice.SetClient(c)
	return &pciedevice, nil
}

// ListReferencedPCIeDevices gets the collection of PCIeDevice from
// a provided reference.
func ListReferencedPCIeDevices(c common.Client, link string) ([]*PCIeDevice, error) { //nolint:dupl
	var result []*PCIeDevice
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *PCIeDevice
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		pciedevice, err := GetPCIeDevice(c, link)
		ch <- GetResult{Item: pciedevice, Link: link, Error: err}
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

// PCIeInterface properties shall be the definition for a PCIe Interface for a
// Redfish implementation.
type PCIeInterface struct {
	// LanesInUse shall be the number of PCIe lanes in use by this device, which
	// shall be equal or less than the value of MaxLanes.
	LanesInUse int
	// MaxLanes shall be the maximum number of PCIe lanes supported by this device.
	MaxLanes int
	// MaxPCIeType shall be the maximum PCIe specification that this device supports.
	MaxPCIeType PCIeTypes
	// PCIeType shall be the negotiated PCIe interface version in use by this device.
	PCIeType PCIeTypes
}

// Assembly gets the assembly for this device.
func (pciedevice *PCIeDevice) Assembly() (*Assembly, error) {
	if pciedevice.assembly == "" {
		return nil, nil
	}
	return GetAssembly(pciedevice.Client, pciedevice.assembly)
}

// Chassis gets the chassis in which the PCIe device is contained.
func (pciedevice *PCIeDevice) Chassis() ([]*Chassis, error) {
	var result []*Chassis

	collectionError := common.NewCollectionError()
	for _, chassisLink := range pciedevice.chassis {
		chassis, err := GetChassis(pciedevice.Client, chassisLink)
		if err != nil {
			collectionError.Failures[chassisLink] = err
		} else {
			result = append(result, chassis)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// PCIeFunctions get the PCIe functions that this device exposes.
func (pciedevice *PCIeDevice) PCIeFunctions() ([]*PCIeDevice, error) {
	var result []*PCIeDevice

	collectionError := common.NewCollectionError()
	for _, funcLink := range pciedevice.pcieFunctions {
		pciFunction, err := GetPCIeDevice(pciedevice.Client, funcLink)
		if err != nil {
			collectionError.Failures[funcLink] = err
		} else {
			result = append(result, pciFunction)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
