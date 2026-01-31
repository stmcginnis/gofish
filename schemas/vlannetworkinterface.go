//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2021.2 - #VLanNetworkInterface.v1_3_1.VLanNetworkInterface

package schemas

import (
	"encoding/json"
)

// VLanNetworkInterface This resource contains information for a VLAN network
// instance that is available on a manager, system, or other device for a
// Redfish implementation.
type VLanNetworkInterface struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// VLANEnable shall indicate whether this VLAN is enabled for this interface.
	VLANEnable bool
	// VLANID shall contain the ID for this VLAN.
	VLANID uint16 `json:"VLANId"`
	// VLANPriority shall contain the priority for this VLAN.
	//
	// Version added: v1.2.0
	VLANPriority uint8
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a VLanNetworkInterface object from the raw JSON.
func (v *VLanNetworkInterface) UnmarshalJSON(b []byte) error {
	type temp VLanNetworkInterface
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*v = VLanNetworkInterface(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	v.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (v *VLanNetworkInterface) Update() error {
	readWriteFields := []string{
		"VLANEnable",
		"VLANId",
		"VLANPriority",
	}

	return v.UpdateFromRawData(v, v.RawData, readWriteFields)
}

// GetVLanNetworkInterface will get a VLanNetworkInterface instance from the service.
func GetVLanNetworkInterface(c Client, uri string) (*VLanNetworkInterface, error) {
	return GetObject[VLanNetworkInterface](c, uri)
}

// ListReferencedVLanNetworkInterfaces gets the collection of VLanNetworkInterface from
// a provided reference.
func ListReferencedVLanNetworkInterfaces(c Client, link string) ([]*VLanNetworkInterface, error) {
	return GetCollectionObjects[VLanNetworkInterface](c, link)
}

// VLAN shall contain any attributes of a VLAN.
type VLAN struct {
	// Tagged shall indicate whether this VLAN is tagged or untagged for this
	// interface.
	//
	// Version added: v1.3.0
	Tagged bool
	// VLANEnable shall indicate whether this VLAN is enabled for this VLAN network
	// interface.
	VLANEnable bool
	// VLANID shall contain the ID for this VLAN.
	VLANID uint16 `json:"VLANId"`
	// VLANPriority shall contain the priority for this VLAN.
	//
	// Version added: v1.2.0
	VLANPriority uint8
}
