//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
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
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
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
func GetNetworkInterface(c common.Client, uri string) (*NetworkInterface, error) {
	var networkInterface NetworkInterface
	return &networkInterface, networkInterface.Get(c, uri, &networkInterface)
}

// ListReferencedNetworkInterfaces gets the collection of NetworkInterface from
// a provided reference.
func ListReferencedNetworkInterfaces(c common.Client, link string) ([]*NetworkInterface, error) {
	return common.GetCollectionObjects(c, link, GetNetworkInterface)
}

// NetworkAdapter gets the NetworkAdapter for this interface.
func (networkinterface *NetworkInterface) NetworkAdapter() (*NetworkAdapter, error) {
	if networkinterface.networkAdapter == "" {
		return nil, nil
	}

	return GetNetworkAdapter(networkinterface.GetClient(), networkinterface.networkAdapter)
}

// NetworkDeviceFunctions gets the collection of NetworkDeviceFunctions of this network interface
func (networkinterface *NetworkInterface) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	var result []*NetworkDeviceFunction

	collectionError := common.NewCollectionError()
	for _, uri := range networkinterface.networkDeviceFunctions {
		unit, err := GetNetworkDeviceFunction(networkinterface.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, unit)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// NetworkPorts gets the collection of NetworkPorts of this network interface
// This property has been deprecated in favor of the Ports property.
func (networkinterface *NetworkInterface) NetworkPorts() ([]*NetworkPort, error) {
	var result []*NetworkPort

	collectionError := common.NewCollectionError()
	for _, uri := range networkinterface.networkPorts {
		unit, err := GetNetworkPort(networkinterface.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, unit)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// Ports gets the ports associated with this network interface.
func (networkinterface *NetworkInterface) Ports() ([]*Port, error) {
	var result []*Port

	collectionError := common.NewCollectionError()
	for _, uri := range networkinterface.ports {
		unit, err := GetPort(networkinterface.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, unit)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
