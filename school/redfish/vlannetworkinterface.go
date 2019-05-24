// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
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
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// VLANEnable is used to indicate if this VLAN is enabled for this
	// interface.
	VLANEnable bool
	// VLANID is used to indicate the VLAN identifier for this VLAN.
	VLANID int16 `json:"VLANId"`
}

// GetVLanNetworkInterface will get a VLanNetworkInterface instance from the service.
func GetVLanNetworkInterface(c common.Client, uri string) (*VLanNetworkInterface, error) {
	resp, err := c.Get(uri)
	defer resp.Body.Close()

	var vlannetworkinterface VLanNetworkInterface
	err = json.NewDecoder(resp.Body).Decode(&vlannetworkinterface)
	if err != nil {
		return nil, err
	}

	vlannetworkinterface.SetClient(c)
	return &vlannetworkinterface, nil
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
