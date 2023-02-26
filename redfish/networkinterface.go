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
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// networkAdapter shall be a reference to a resource of type NetworkAdapter
	// that represents the physical container associated with this NetworkInterface.
	networkAdapter string
	// networkDeviceFunctions shall be a link to a collection of type
	// NetworkDeviceFunctionCollection.
	networkDeviceFunctions string
	// NetworkPorts shall be a link to a collection of type NetworkPortCollection.
	networkPorts string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a NetworkInterface object from the raw JSON.
func (networkinterface *NetworkInterface) UnmarshalJSON(b []byte) error {
	type temp NetworkInterface
	var t struct {
		temp
		NetworkDeviceFunctions common.Link
		NetworkPorts           common.Link
		Links                  NetworkInterfaceLinks
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*networkinterface = NetworkInterface(t.temp)
	networkinterface.networkAdapter = t.Links.NetworkAdapter.String()
	networkinterface.networkDeviceFunctions = t.NetworkDeviceFunctions.String()
	networkinterface.networkPorts = t.NetworkPorts.String()

	return nil
}

// GetNetworkInterface will get a NetworkInterface instance from the service.
func GetNetworkInterface(c common.Client, uri string) (*NetworkInterface, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var networkinterface NetworkInterface
	err = json.NewDecoder(resp.Body).Decode(&networkinterface)
	if err != nil {
		return nil, err
	}

	networkinterface.SetClient(c)
	return &networkinterface, nil
}

// ListReferencedNetworkInterfaces gets the collection of NetworkInterface from
// a provided reference.
func ListReferencedNetworkInterfaces(c common.Client, link string) ([]*NetworkInterface, error) { //nolint:dupl
	var result []*NetworkInterface
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *NetworkInterface
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		networkinterface, err := GetNetworkInterface(c, link)
		ch <- GetResult{Item: networkinterface, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// NetworkAdapter gets the NetworkAdapter for this interface.
func (networkinterface *NetworkInterface) NetworkAdapter() (*NetworkAdapter, error) {
	if networkinterface.networkAdapter == "" {
		return nil, nil
	}

	return GetNetworkAdapter(networkinterface.Client, networkinterface.networkAdapter)
}

// NetworkDeviceFunctions gets the collection of NetworkDeviceFunctions of this network interface
func (networkinterface *NetworkInterface) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	return ListReferencedNetworkDeviceFunctions(
		networkinterface.Client, networkinterface.networkDeviceFunctions)
}

// NetworkPorts gets the collection of NetworkPorts of this network interface
func (networkinterface *NetworkInterface) NetworkPorts() ([]*NetworkPort, error) {
	return ListReferencedNetworkPorts(
		networkinterface.Client, networkinterface.networkPorts)
}
