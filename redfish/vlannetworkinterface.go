//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"io/ioutil"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// VLAN shall contain any attributes of a Virtual LAN.
type VLAN struct {
	// VLANEnable is used to indicate if this VLAN is enabled for this
	// interface.
	VLANEnable bool
	// VLANID is used to indicate the VLAN identifier for this VLAN.
	VLANID int16 `json:"VLANId"`
}

// VLanNetworkInterface shall contain any attributes of a Virtual LAN.
type VLanNetworkInterface struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// VLANEnable is used to indicate if this VLAN is enabled for this
	// interface.
	VLANEnable bool
	// VLANID is used to indicate the VLAN identifier for this VLAN.
	VLANID int16 `json:"VLANId"`
	// rawData holds the original serialized JSON
	rawData             []byte
}

// GetRawData get raw data json
func (vlannetworkinterface *VLanNetworkInterface) GetRawData() []byte {
	return vlannetworkinterface.rawData
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
	original.UnmarshalJSON(vlannetworkinterface.rawData)

	readWriteFields := []string{
		"VLANEnable",
		"VLANId",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(vlannetworkinterface).Elem()

	return vlannetworkinterface.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetVLanNetworkInterface will get a VLanNetworkInterface instance from the service.
func GetVLanNetworkInterface(c common.Client, uri string) (*VLanNetworkInterface, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var vlanNetworkInterface VLanNetworkInterface
	vlanNetworkInterface.rawData, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(vlanNetworkInterface.rawData, &vlanNetworkInterface)
	if err != nil {
		return nil, err
	}

	vlanNetworkInterface.SetClient(c)
	return &vlanNetworkInterface, nil
}

// ListReferencedVLanNetworkInterfaces gets the collection of VLanNetworkInterface from
// a provided reference.
func ListReferencedVLanNetworkInterfaces(c common.Client, link string) ([]*VLanNetworkInterface, error) {
	var result []*VLanNetworkInterface
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, vlannetworkinterfaceLink := range links.ItemLinks {
		vlannetworkinterface, err := GetVLanNetworkInterface(c, vlannetworkinterfaceLink)
		if err != nil {
			return result, err
		}
		result = append(result, vlannetworkinterface)
	}

	return result, nil
}
