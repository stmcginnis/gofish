//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type SlotTypes string

const (
	// Full-Length PCIe slot.
	FullLength SlotTypes = "FullLength"
	// Half-Length PCIe slot.
	HalfLength SlotTypes = "HalfLength"
	// Low-Profile or Slim PCIe slot.
	LowProfile SlotTypes = "LowProfile"
	// PCIe M.2 slot.
	M2 SlotTypes = "M2"
	// Mini PCIe slot.
	Mini SlotTypes = "Mini"
	// (v1.2+) Open Compute Project 3.0 large form factor slot.
	OCP3Large SlotTypes = "OCP3Large"
	// (v1.2+) Open Compute Project 3.0 small form factor slot.
	OCP3Small SlotTypes = "OCP3Small"
	// An OEM-specific slot.
	OEM SlotTypes = "OEM"
	// (v1.3+) U.2 / SFF-8639 slot or bay.
	U2 SlotTypes = "U2"
)

type Slot struct {
	// HotPluggable is an indication of whether this PCIe slot supports hotplug.
	HotPluggable bool
	// Lanes is the number of PCIe lanes supported by this slot.
	Lanes int
	// Location is the location of the PCIe slot.
	Location common.Location
	// LocationIndicatorActive is an indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool
	// PCIeType is the PCIe specification supported by this slot.
	PCIeType PCIeTypes
	// SlotType is the PCIe slot type for this slot
	SlotType SlotTypes
	// PCIeDevice shall be an array of links to the PCIe devices contained in this slot.
	pcieDevice []string
	// PCIeDeviceCount is the number of PCIe devices contained in this slot.
	PCIeDeviceCount int
	// Processors shall be an array of links to the processors
	// that are directly connected or directly bridged to this PCIe slot.
	processors []string
	// ProcessorsCount is the number of processors
	// that are directly connected or directly bridged to this PCIe slot.
	ProcessorsCount int
	// OEMLinks are all OEM data under link section
	OemLinks json.RawMessage
	// Oem shall contain the OEM extensions. All values for properties that
	// this object contains shall conform to the Redfish Specification
	// described requirements.
	Oem json.RawMessage
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a Slot object from the raw JSON.
func (slot *Slot) UnmarshalJSON(b []byte) error {
	type temp Slot
	type linkReference struct {
		PCIeDevice      common.Links
		PCIeDeviceCount int `json:"PCIeDevice@odata.count"`
		Processors      common.Links
		ProcessorsCount int `json:"Processors@odata.count"`
		Oem             json.RawMessage
	}

	var t struct {
		temp
		Links linkReference
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	*slot = Slot(t.temp)
	slot.pcieDevice = t.Links.PCIeDevice.ToStrings()
	slot.PCIeDeviceCount = t.Links.PCIeDeviceCount
	slot.processors = t.Links.Processors.ToStrings()
	slot.ProcessorsCount = t.Links.ProcessorsCount
	slot.OemLinks = t.Links.Oem

	return nil
}

// PCIeSlots is used to represent a PCIeSlots resource for a Redfish
// implementation.
type PCIeSlots struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Slots is an array of PCI Slot information.
	Slots []Slot
	// Oem shall contain the OEM extensions. All values for properties that
	// this object contains shall conform to the Redfish Specification
	// described requirements.
	Oem json.RawMessage
	// OemActions contains all the vendor specific actions. It is vendor responsibility to parse this field accordingly
	OemActions json.RawMessage
}

// UnmarshalJSON unmarshals a Slot object from the raw JSON.
func (pcieSlots *PCIeSlots) UnmarshalJSON(b []byte) error {
	type temp PCIeSlots
	type actions struct {
		Oem json.RawMessage // OEM actions will be stored here
	}
	var t struct {
		temp
		Actions actions
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	*pcieSlots = PCIeSlots(t.temp)
	pcieSlots.OemActions = t.Actions.Oem

	return nil
}

// GetPCIeSlots will get a PCIeSlots instance from the chassis.
func GetPCIeSlots(c common.Client, uri string) (*PCIeSlots, error) {
	var pcieSlots PCIeSlots
	return &pcieSlots, pcieSlots.Get(c, uri, &pcieSlots)
}
