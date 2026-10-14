//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a NetworkDeviceFunction object from the raw JSON.
func (n *NetworkDeviceFunction) UnmarshalJSON(b []byte) error {
	type temp NetworkDeviceFunction
	type nLinks struct {
		Endpoints                     Links `json:"Endpoints"`
		EthernetInterface             Link  `json:"EthernetInterface"`
		EthernetInterfaces            Links `json:"EthernetInterfaces"`
		OffloadProcessors             Links `json:"OffloadProcessors"`
		OffloadSystem                 Link  `json:"OffloadSystem"`
		PCIeFunction                  Link  `json:"PCIeFunction"`
		PhysicalNetworkPortAssignment Link  `json:"PhysicalNetworkPortAssignment"`
		PhysicalPortAssignment        Link  `json:"PhysicalPortAssignment"`
	}
	var tmp struct {
		temp
		Links                          nLinks
		AllowDeny                      Link                                       `json:"AllowDeny"`
		AssignablePhysicalNetworkPorts Links                                      `json:"AssignablePhysicalNetworkPorts"`
		AssignablePhysicalPorts        Links                                      `json:"AssignablePhysicalPorts"`
		Metrics                        ResolvedLink[NetworkDeviceFunctionMetrics] `json:"Metrics"`
		PhysicalPortAssignment         Link                                       `json:"PhysicalPortAssignment"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*n = NetworkDeviceFunction(tmp.temp)

	// Extract the links to other entities for later
	n.endpoints = tmp.Links.Endpoints.ToStrings()
	n.ethernetInterface = tmp.Links.EthernetInterface.String()
	n.ethernetInterfaces = tmp.Links.EthernetInterfaces.ToStrings()
	n.offloadProcessors = tmp.Links.OffloadProcessors.ToStrings()
	n.offloadSystem = tmp.Links.OffloadSystem.String()
	n.pCIeFunction = tmp.Links.PCIeFunction.String()
	n.physicalNetworkPortAssignment = tmp.Links.PhysicalNetworkPortAssignment.String()
	n.physicalPortAssignment = tmp.Links.PhysicalPortAssignment.String()
	n.allowDeny = tmp.AllowDeny.String()
	n.assignablePhysicalNetworkPorts = tmp.AssignablePhysicalNetworkPorts.ToStrings()
	n.assignablePhysicalPorts = tmp.AssignablePhysicalPorts.ToStrings()
	n.metrics = tmp.Metrics

	if tmp.Links.PhysicalPortAssignment.String() == "" {
		n.physicalPortAssignment = tmp.PhysicalPortAssignment.String()
	}

	// This is a read/write object, so we need to save the raw object data for later
	n.RawData = b

	return nil
}

// PhysicalPortAssignment gets the PhysicalPortAssignment linked resource.
func (n *NetworkDeviceFunction) PhysicalPortAssignment() (*NetworkPort, error) {
	if n.physicalPortAssignment == "" {
		return nil, nil
	}
	return GetObject[NetworkPort](n.client, n.physicalPortAssignment)
}
