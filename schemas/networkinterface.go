//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/NetworkInterface.v1_2_3.json
// 2020.3 - #NetworkInterface.v1_2_3.NetworkInterface

package schemas

import (
	"encoding/json"
)

// NetworkInterface This resource contains links to the network adapters,
// network ports, and network device functions, and represents the functionality
// available to the containing system.
type NetworkInterface struct {
	Entity
	// NetworkDeviceFunctions shall contain a link to a resource collection of type
	// 'NetworkDeviceFunctionCollection'. The members of this collection shall not
	// contain 'NetworkDeviceFunction' resources whose 'NetDevFuncType' property
	// contains 'Ethernet'.
	networkDeviceFunctions string
	// NetworkPorts shall contain a link to a resource collection of type
	// 'NetworkPortCollection'.
	//
	// Deprecated: v1.2.0
	// This property has been deprecated in favor of the 'Ports' property.
	networkPorts string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Ports shall contain a link to a resource collection of type
	// 'PortCollection'.
	//
	// Version added: v1.2.0
	ports string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// networkAdapter is the URI for NetworkAdapter.
	networkAdapter string
}

// UnmarshalJSON unmarshals a NetworkInterface object from the raw JSON.
func (n *NetworkInterface) UnmarshalJSON(b []byte) error {
	type temp NetworkInterface
	type nLinks struct {
		NetworkAdapter Link `json:"NetworkAdapter"`
	}
	var tmp struct {
		temp
		Links                  nLinks
		NetworkDeviceFunctions Link `json:"NetworkDeviceFunctions"`
		NetworkPorts           Link `json:"NetworkPorts"`
		Ports                  Link `json:"Ports"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*n = NetworkInterface(tmp.temp)

	// Extract the links to other entities for later
	n.networkAdapter = tmp.Links.NetworkAdapter.String()
	n.networkDeviceFunctions = tmp.NetworkDeviceFunctions.String()
	n.networkPorts = tmp.NetworkPorts.String()
	n.ports = tmp.Ports.String()

	return nil
}

// GetNetworkInterface will get a NetworkInterface instance from the service.
func GetNetworkInterface(c Client, uri string) (*NetworkInterface, error) {
	return GetObject[NetworkInterface](c, uri)
}

// ListReferencedNetworkInterfaces gets the collection of NetworkInterface from
// a provided reference.
func ListReferencedNetworkInterfaces(c Client, link string) ([]*NetworkInterface, error) {
	return GetCollectionObjects[NetworkInterface](c, link)
}

// NetworkAdapter gets the NetworkAdapter linked resource.
func (n *NetworkInterface) NetworkAdapter() (*NetworkAdapter, error) {
	if n.networkAdapter == "" {
		return nil, nil
	}
	return GetObject[NetworkAdapter](n.client, n.networkAdapter)
}

// NetworkDeviceFunctions gets the NetworkDeviceFunctions collection.
func (n *NetworkInterface) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	if n.networkDeviceFunctions == "" {
		return nil, nil
	}
	return GetCollectionObjects[NetworkDeviceFunction](n.client, n.networkDeviceFunctions)
}

// NetworkPorts gets the NetworkPorts collection.
func (n *NetworkInterface) NetworkPorts() ([]*NetworkPort, error) {
	if n.networkPorts == "" {
		return nil, nil
	}
	return GetCollectionObjects[NetworkPort](n.client, n.networkPorts)
}

// Ports gets the Ports collection.
func (n *NetworkInterface) Ports() ([]*Port, error) {
	if n.ports == "" {
		return nil, nil
	}
	return GetCollectionObjects[Port](n.client, n.ports)
}
