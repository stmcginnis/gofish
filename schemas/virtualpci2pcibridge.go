//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #VirtualPCI2PCIBridge.v1_0_0.VirtualPCI2PCIBridge

package schemas

import (
	"encoding/json"
)

type vPPBStatusTypes string

const (
	// UnboundvPPBStatusTypes shall indicate the CXL Specification-defined
	// 'Unbound' state for the vPPB.
	UnboundvPPBStatusTypes vPPBStatusTypes = "Unbound"
	// BusyvPPBStatusTypes shall indicate the CXL Specification-defined 'Binding or
	// Unbinding' state for the vPPB.
	BusyvPPBStatusTypes vPPBStatusTypes = "Busy"
	// BoundPhysicalPortvPPBStatusTypes shall indicate the CXL
	// Specification-defined 'Bound Physical Port' state for the vPPB.
	BoundPhysicalPortvPPBStatusTypes vPPBStatusTypes = "BoundPhysicalPort"
	// BoundLDvPPBStatusTypes shall indicate the CXL Specification-defined 'Bound
	// LD' state for the vPPB.
	BoundLDvPPBStatusTypes vPPBStatusTypes = "BoundLD"
	// BoundPIDvPPBStatusTypes shall indicate the CXL Specification-defined 'Bound
	// PBR Id' state for the vPPB.
	BoundPIDvPPBStatusTypes vPPBStatusTypes = "BoundPID"
)

// VirtualPCI2PCIBridge shall represent a Virtual PCI-to-PCI Bridge (vPPB)
// inside a CXL switch that is host-owned. This can be bound to a port that is
// either disconnected, connected to a PCIe component, or connected to a CXL
// component.
type VirtualPCI2PCIBridge struct {
	Entity
	// BindingStatus shall contain the binding status of the vPPB.
	BindingStatus vPPBStatusTypes
	// BoundLDID shall contain the identifier of the bound local logical device
	// bound to this vPPB. This property shall only be present if 'BindingStatus'
	// contains 'BoundLD'.
	BoundLDID int `json:"BoundLDId"`
	// BoundPBRID shall contain the identifier of PBR bound to this vPPB. This
	// property shall only be present if 'BindingStatus' contains 'BoundPID'.
	BoundPBRID int `json:"BoundPBRId"`
	// BoundPortID shall contain the physical port number of the port bound to this
	// vPPB. This property shall only be present if 'BindingStatus' contains
	// 'BoundPhysicalPort' or 'BoundLD'.
	BoundPortID int `json:"BoundPortId"`
	// GCXLID shall contain the globally unique CXL logical device identifier
	// (GCXLID) for the CXL logical device that is bound to this vPPB.
	GCXLID string
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
	// VPPBID shall contain the identifier of the vPPB. This property shall contain
	// the same value as the 'Id' property.
	VPPBID string `json:"VPPBId"`
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// port is the URI for Port.
	port string
}

// UnmarshalJSON unmarshals a VirtualPCI2PCIBridge object from the raw JSON.
func (v *VirtualPCI2PCIBridge) UnmarshalJSON(b []byte) error {
	type temp VirtualPCI2PCIBridge
	type vLinks struct {
		Endpoints Links `json:"Endpoints"`
		Port      Link  `json:"Port"`
	}
	var tmp struct {
		temp
		Links vLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*v = VirtualPCI2PCIBridge(tmp.temp)

	// Extract the links to other entities for later
	v.endpoints = tmp.Links.Endpoints.ToStrings()
	v.port = tmp.Links.Port.String()

	return nil
}

// GetVirtualPCI2PCIBridge will get a VirtualPCI2PCIBridge instance from the service.
func GetVirtualPCI2PCIBridge(c Client, uri string) (*VirtualPCI2PCIBridge, error) {
	return GetObject[VirtualPCI2PCIBridge](c, uri)
}

// ListReferencedVirtualPCI2PCIBridges gets the collection of VirtualPCI2PCIBridge from
// a provided reference.
func ListReferencedVirtualPCI2PCIBridges(c Client, link string) ([]*VirtualPCI2PCIBridge, error) {
	return GetCollectionObjects[VirtualPCI2PCIBridge](c, link)
}

// Endpoints gets the Endpoints linked resources.
func (v *VirtualPCI2PCIBridge) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](v.client, v.endpoints)
}

// Port gets the Port linked resource.
func (v *VirtualPCI2PCIBridge) Port() (*Port, error) {
	if v.port == "" {
		return nil, nil
	}
	return GetObject[Port](v.client, v.port)
}
