//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// VLAN shall contain any attributes of a Virtual LAN.
type VLAN struct {
	// Tagged shall indicate whether this VLAN is tagged or untagged for this interface.
	Tagged bool
	// VLANEnable is used to indicate if this VLAN is enabled for this
	// interface.
	VLANEnable bool
	// VLANID is used to indicate the VLAN identifier for this VLAN.
	VLANID int16 `json:"VLANId"`
	// VLANPriority shall contain the priority for this VLAN (0-7).
	VLANPriority int
}

// VLanNetworkInterface shall contain any attributes of a Virtual LAN.
type VLanNetworkInterface struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// VLANEnable is used to indicate if this VLAN is enabled for this
	// interface.
	VLANEnable bool
	// VLANID is used to indicate the VLAN identifier for this VLAN.
	VLANID int16 `json:"VLANId"`
	// VLANPriority shall contain the priority for this VLAN (0-7).
	VLANPriority int
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals an object from the raw JSON.
func (vlannetworkinterface *VLanNetworkInterface) UnmarshalJSON(b []byte) error {
	type temp VLanNetworkInterface
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*vlannetworkinterface = VLanNetworkInterface(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	vlannetworkinterface.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (vlannetworkinterface *VLanNetworkInterface) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(VLanNetworkInterface)
	err := original.UnmarshalJSON(vlannetworkinterface.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"VLANEnable",
		"VLANId",
		"VLANPriority",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(vlannetworkinterface).Elem()

	return vlannetworkinterface.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetVLanNetworkInterface will get a VLanNetworkInterface instance from the service.
func GetVLanNetworkInterface(c common.Client, uri string) (*VLanNetworkInterface, error) {
	return common.GetObject[VLanNetworkInterface](c, uri)
}

// ListReferencedVLanNetworkInterfaces gets the collection of VLanNetworkInterface from
// a provided reference.
func ListReferencedVLanNetworkInterfaces(c common.Client, link string) ([]*VLanNetworkInterface, error) {
	return common.GetCollectionObjects[VLanNetworkInterface](c, link)
}
