//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2023.2 - #PCIeSlots.v1_6_1.PCIeSlots

package schemas

import (
	"encoding/json"
)

type SlotTypes string

const (
	// FullLengthSlotTypes Full-Length PCIe slot.
	FullLengthSlotTypes SlotTypes = "FullLength"
	// HalfLengthSlotTypes Half-Length PCIe slot.
	HalfLengthSlotTypes SlotTypes = "HalfLength"
	// LowProfileSlotTypes Low-Profile or Slim PCIe slot.
	LowProfileSlotTypes SlotTypes = "LowProfile"
	// MiniSlotTypes Mini PCIe slot.
	MiniSlotTypes SlotTypes = "Mini"
	// M2SlotTypes PCIe M.2 slot.
	M2SlotTypes SlotTypes = "M2"
	// OEMSlotTypes is an OEM-specific slot.
	OEMSlotTypes SlotTypes = "OEM"
	// OCP3SmallSlotTypes Open Compute Project 3.0 small form factor slot.
	OCP3SmallSlotTypes SlotTypes = "OCP3Small"
	// OCP3LargeSlotTypes Open Compute Project 3.0 large form factor slot.
	OCP3LargeSlotTypes SlotTypes = "OCP3Large"
	// U2SlotTypes U.2 / SFF-8639 slot or bay.
	U2SlotTypes SlotTypes = "U2"
)

// PCIeSlots shall represent a set of PCIe slot information for a Redfish
// implementation.
type PCIeSlots struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Slots shall contain an entry for each PCIe slot, including empty slots (with
	// no device or card installed).
	Slots []PCIeSlot
}

// GetPCIeSlots will get a PCIeSlots instance from the service.
func GetPCIeSlots(c Client, uri string) (*PCIeSlots, error) {
	return GetObject[PCIeSlots](c, uri)
}

// ListReferencedPCIeSlotss gets the collection of PCIeSlots from
// a provided reference.
func ListReferencedPCIeSlotss(c Client, link string) ([]*PCIeSlots, error) {
	return GetCollectionObjects[PCIeSlots](c, link)
}

// PCIeLinks shall contain links to resources that are related to but are not
// contained by, or subordinate to, this resource.
type PCIeLinks struct {
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevice shall contain an array of links to resources of type 'PCIeDevice'
	// with which this physical slot is associated. If the 'State' property in
	// 'Status' of this slot is 'Absent', this property shall not appear in the
	// resource.
	pCIeDevice []string
	// PCIeDeviceCount
	PCIeDeviceCount int `json:"PCIeDevice@odata.count"`
	// Processors shall contain an array of links to resources of type 'Processor'
	// that represent processors that are directly connected or directly bridged to
	// this PCIe slot.
	//
	// Version added: v1.5.0
	processors []string
	// ProcessorsCount
	ProcessorsCount int `json:"Processors@odata.count"`
}

// UnmarshalJSON unmarshals a PCIeLinks object from the raw JSON.
func (p *PCIeLinks) UnmarshalJSON(b []byte) error {
	type temp PCIeLinks
	var tmp struct {
		temp
		PCIeDevice Links `json:"PCIeDevice"`
		Processors Links `json:"Processors"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PCIeLinks(tmp.temp)

	// Extract the links to other entities for later
	p.pCIeDevice = tmp.PCIeDevice.ToStrings()
	p.processors = tmp.Processors.ToStrings()

	return nil
}

// PCIeDevice gets the PCIeDevice linked resources.
func (p *PCIeLinks) PCIeDevice(client Client) ([]*PCIeDevice, error) {
	return GetObjects[PCIeDevice](client, p.pCIeDevice)
}

// Processors gets the Processors linked resources.
func (p *PCIeLinks) Processors(client Client) ([]*Processor, error) {
	return GetObjects[Processor](client, p.processors)
}

// PCIeSlot shall contain the definition for a PCIe slot for a Redfish
// implementation.
type PCIeSlot struct {
	// HotPluggable shall contain indicating whether this PCIe slot supports
	// hotplug.
	//
	// Version added: v1.1.0
	HotPluggable bool
	// Lanes shall contain the maximum number of PCIe lanes supported by the slot.
	Lanes *int `json:",omitempty"`
	// Location shall contain part location information, including a 'ServiceLabel'
	// of the associated PCIe slot.
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource. A write to this property shall
	// update the value of 'IndicatorLED' in this resource, if supported, to
	// reflect the implementation of the locating function.
	//
	// Version added: v1.4.0
	LocationIndicatorActive bool
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeType shall contain the maximum PCIe specification that this slot
	// supports.
	PCIeType PCIeTypes
	// SlotType shall contain the slot type as specified by the PCIe specification.
	SlotType SlotTypes
	// Status shall contain any status or health properties of the resource.
	Status Status
}
