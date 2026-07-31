//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

// NetworkInterfaceLinks references to resources that are related to, but not
// contained by (subordinate to), this resource.
type NetworkInterfaceLinks struct {
	// NetworkAdapter shall be a reference to a
	// resource of type NetworkAdapter that represents the physical container
	// associated with this NetworkInterface.
	NetworkAdapter common.Link
}

// A NetworkInterface contains references linking NetworkAdapter, NetworkPort,
// and NetworkDeviceFunction resources and represents the functionality
// available to the containing system.
type NetworkInterface struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// networkDeviceFunctions shall be a link to a collection of type
	// NetworkDeviceFunctionCollection.
	networkDeviceFunctions []string
	// NetworkPorts shall be a link to a collection of type NetworkPortCollection.
	// This property has been deprecated in favor of the Ports property.
	networkPorts []string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Ports shall contain a link to a resource collection of type PortCollection.
	ports []string
	// Status shall contain any status or health properties of the resource.
	Status common.Status

	// networkAdapter shall be a reference to a resource of type NetworkAdapter
	// that represents the physical container associated with this NetworkInterface.
	networkAdapter string
}

// UnmarshalJSON unmarshals a NetworkInterface object from the raw JSON.
func (networkinterface *NetworkInterface) UnmarshalJSON(b []byte) error {
	type temp NetworkInterface
	var t struct {
		temp
		NetworkDeviceFunctions common.LinksCollection
		NetworkPorts           common.LinksCollection
		Ports                  common.LinksCollection
		Links                  NetworkInterfaceLinks
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*networkinterface = NetworkInterface(t.temp)
	networkinterface.networkAdapter = t.Links.NetworkAdapter.String()

	networkinterface.networkDeviceFunctions = t.NetworkDeviceFunctions.ToStrings()
	networkinterface.networkPorts = t.NetworkPorts.ToStrings()
	networkinterface.ports = t.Ports.ToStrings()

	return nil
}

// GetNetworkInterface will get a NetworkInterface instance from the service.
func GetNetworkInterface(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*NetworkInterface, error) {
	return GetNetworkInterfaceWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetNetworkInterfaceWithContext will get a NetworkInterface instance from the service.
func GetNetworkInterfaceWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*NetworkInterface, error) {
	return common.GetObjectWithContext[NetworkInterface](ctx, c, uri, queryOpts...)
}

// ListReferencedNetworkInterfaces gets the collection of NetworkInterface from
// a provided reference.
func ListReferencedNetworkInterfaces(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*NetworkInterface, error) {
	return ListReferencedNetworkInterfacesWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedNetworkInterfacesWithContext gets the collection of NetworkInterface from
// a provided reference.
func ListReferencedNetworkInterfacesWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*NetworkInterface, error) {
	return common.GetCollectionObjectsWithContext[NetworkInterface](ctx, c, link, queryOpts...)
}

// NetworkAdapter gets the NetworkAdapter for this interface.
func (networkinterface *NetworkInterface) NetworkAdapter(queryOpts ...common.QueryGroupOption) (*NetworkAdapter, error) {
	return networkinterface.NetworkAdapterWithContext(common.ContextOf(networkinterface.GetClient()), queryOpts...)
}

// NetworkAdapterWithContext gets the NetworkAdapter for this interface.
func (networkinterface *NetworkInterface) NetworkAdapterWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*NetworkAdapter, error) {
	if networkinterface.networkAdapter == "" {
		return nil, nil
	}

	return GetNetworkAdapterWithContext(ctx, networkinterface.GetClient(), networkinterface.networkAdapter, queryOpts...)
}

// NetworkDeviceFunctions gets the collection of NetworkDeviceFunctions of this network interface
func (networkinterface *NetworkInterface) NetworkDeviceFunctions(queryOpts ...common.QueryGroupOption) ([]*NetworkDeviceFunction, error) {
	return networkinterface.NetworkDeviceFunctionsWithContext(common.ContextOf(networkinterface.GetClient()), queryOpts...)
}

// NetworkDeviceFunctionsWithContext gets the collection of NetworkDeviceFunctions of this network interface
func (networkinterface *NetworkInterface) NetworkDeviceFunctionsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*NetworkDeviceFunction, error) {
	return common.GetObjectsWithContext[NetworkDeviceFunction](ctx, networkinterface.GetClient(), networkinterface.networkDeviceFunctions, queryOpts...)
}

// NetworkPorts gets the collection of NetworkPorts of this network interface
// This property has been deprecated in favor of the Ports property.
func (networkinterface *NetworkInterface) NetworkPorts(queryOpts ...common.QueryGroupOption) ([]*NetworkPort, error) {
	return networkinterface.NetworkPortsWithContext(common.ContextOf(networkinterface.GetClient()), queryOpts...)
}

// NetworkPortsWithContext gets the collection of NetworkPorts of this network interface
// This property has been deprecated in favor of the Ports property.
func (networkinterface *NetworkInterface) NetworkPortsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*NetworkPort, error) {
	return common.GetObjectsWithContext[NetworkPort](ctx, networkinterface.GetClient(), networkinterface.networkPorts, queryOpts...)
}

// Ports gets the ports associated with this network interface.
func (networkinterface *NetworkInterface) Ports(queryOpts ...common.QueryGroupOption) ([]*Port, error) {
	return networkinterface.PortsWithContext(common.ContextOf(networkinterface.GetClient()), queryOpts...)
}

// PortsWithContext gets the ports associated with this network interface.
func (networkinterface *NetworkInterface) PortsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Port, error) {
	return common.GetObjectsWithContext[Port](ctx, networkinterface.GetClient(), networkinterface.ports, queryOpts...)
}
