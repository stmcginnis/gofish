//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #VirtualCXLSwitch.v1_0_0.VirtualCXLSwitch

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// VirtualCXLSwitch shall represent a VCS entity within a CXL switch. The CXL
// Specification contains the complete definition of a Virtual CXL Switch.
type VirtualCXLSwitch struct {
	common.Entity
	// HDMDecoders The number of HDM (Host Device Memory) decoders present in this
	// Virtual CXL Switch.
	HDMDecoders int
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// VCSId shall contain the identifier of the Virtual CXL Switch.
	VCSID string `json:"VCSId"`
	// VPPBs shall contain a link to a resource collection of type
	// 'VirtualPCI2PCIBridgeCollection'.
	vPPBs string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// port is the URI for Port.
	port string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a VirtualCXLSwitch object from the raw JSON.
func (v *VirtualCXLSwitch) UnmarshalJSON(b []byte) error {
	type temp VirtualCXLSwitch
	type vLinks struct {
		Endpoints common.Links `json:"Endpoints"`
		Port      common.Link  `json:"Port"`
	}
	var tmp struct {
		temp
		Links vLinks
		VPPBs common.Link `json:"vPPBs"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*v = VirtualCXLSwitch(tmp.temp)

	// Extract the links to other entities for later
	v.endpoints = tmp.Links.Endpoints.ToStrings()
	v.port = tmp.Links.Port.String()
	v.vPPBs = tmp.VPPBs.String()

	// This is a read/write object, so we need to save the raw object data for later
	v.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (v *VirtualCXLSwitch) Update() error {
	readWriteFields := []string{
		"Status",
	}

	return v.UpdateFromRawData(v, v.rawData, readWriteFields)
}

// GetVirtualCXLSwitch will get a VirtualCXLSwitch instance from the service.
func GetVirtualCXLSwitch(c common.Client, uri string) (*VirtualCXLSwitch, error) {
	return common.GetObject[VirtualCXLSwitch](c, uri)
}

// ListReferencedVirtualCXLSwitchs gets the collection of VirtualCXLSwitch from
// a provided reference.
func ListReferencedVirtualCXLSwitchs(c common.Client, link string) ([]*VirtualCXLSwitch, error) {
	return common.GetCollectionObjects[VirtualCXLSwitch](c, link)
}

// Endpoints gets the Endpoints linked resources.
func (v *VirtualCXLSwitch) Endpoints(client common.Client) ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](client, v.endpoints)
}

// Port gets the Port linked resource.
func (v *VirtualCXLSwitch) Port(client common.Client) (*Port, error) {
	if v.port == "" {
		return nil, nil
	}
	return common.GetObject[Port](client, v.port)
}

// VPPBs gets the VPPBs collection.
func (v *VirtualCXLSwitch) VPPBs(client common.Client) ([]*VirtualPCI2PCIBridge, error) {
	if v.vPPBs == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[VirtualPCI2PCIBridge](client, v.vPPBs)
}
