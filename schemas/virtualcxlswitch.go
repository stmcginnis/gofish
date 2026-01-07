//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #VirtualCXLSwitch.v1_0_0.VirtualCXLSwitch

package schemas

import (
	"encoding/json"
)

// VirtualCXLSwitch shall represent a VCS entity within a CXL switch. The CXL
// Specification contains the complete definition of a Virtual CXL Switch.
type VirtualCXLSwitch struct {
	Entity
	// HDMDecoders The number of HDM (Host Device Memory) decoders present in this
	// Virtual CXL Switch.
	HDMDecoders int
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status Status
	// VCSID shall contain the identifier of the Virtual CXL Switch.
	VCSID string `json:"VCSId"`
	// VPPBs shall contain a link to a resource collection of type
	// 'VirtualPCI2PCIBridgeCollection'.
	vPPBs string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// port is the URI for Port.
	port string
}

// UnmarshalJSON unmarshals a VirtualCXLSwitch object from the raw JSON.
func (v *VirtualCXLSwitch) UnmarshalJSON(b []byte) error {
	type temp VirtualCXLSwitch
	type vLinks struct {
		Endpoints Links `json:"Endpoints"`
		Port      Link  `json:"Port"`
	}
	var tmp struct {
		temp
		Links vLinks
		VPPBs Link `json:"VPPBs"`
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

	return nil
}

// GetVirtualCXLSwitch will get a VirtualCXLSwitch instance from the service.
func GetVirtualCXLSwitch(c Client, uri string) (*VirtualCXLSwitch, error) {
	return GetObject[VirtualCXLSwitch](c, uri)
}

// ListReferencedVirtualCXLSwitchs gets the collection of VirtualCXLSwitch from
// a provided reference.
func ListReferencedVirtualCXLSwitchs(c Client, link string) ([]*VirtualCXLSwitch, error) {
	return GetCollectionObjects[VirtualCXLSwitch](c, link)
}

// Endpoints gets the Endpoints linked resources.
func (v *VirtualCXLSwitch) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](v.client, v.endpoints)
}

// Port gets the Port linked resource.
func (v *VirtualCXLSwitch) Port() (*Port, error) {
	if v.port == "" {
		return nil, nil
	}
	return GetObject[Port](v.client, v.port)
}

// VPPBs gets the VPPBs collection.
func (v *VirtualCXLSwitch) VPPBs() ([]*VirtualPCI2PCIBridge, error) {
	if v.vPPBs == "" {
		return nil, nil
	}
	return GetCollectionObjects[VirtualPCI2PCIBridge](v.client, v.vPPBs)
}
