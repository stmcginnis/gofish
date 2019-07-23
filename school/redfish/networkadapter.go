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

// ControllerCapabilities is This type shall describe the capabilities of
// a controller.
type ControllerCapabilities struct {
	// DataCenterBridging is This object shall contain capability, status,
	// and configuration values related to Data Center Bridging (DCB) for
	// this controller.
	DataCenterBridging struct {
		Capable bool
	}
	// NPAR is This object shall contain capability, status, and
	// configuration values related to NIC partitioning for this controller.
	NPAR struct {
		// NparCapable is This property shall indicate the ability of a
		// controller to support NIC function partitioning.
		NparCapable bool
		// NparEnabled is This property shall indicate whether or not NIC
		// function partitioning is active on this controller.
		NparEnabled bool
	}
	// NPIV is This object shall contain N_Port ID Virtualization (NPIV)
	// capabilties for this controller.
	NPIV struct {
		// MaxDeviceLogins is The value of this property shall be the maximum
		// number of N_Port ID Virtualization (NPIV) logins allowed
		// simultaneously from all ports on this controller.
		MaxDeviceLogins int
		// MaxPortLogins is The value of this property shall be the maximum
		// number of N_Port ID Virtualization (NPIV) logins allowed per physical
		// port on this controller.
		MaxPortLogins int
	}
	// NetworkDeviceFunctionCount is The value of this property shall be the
	// number of physical functions available on this controller.
	NetworkDeviceFunctionCount int
	// NetworkPortCount is The value of this property shall be the number of
	// physical ports on this controller.
	NetworkPortCount int
	// VirtualizationOffload is This object shall contain capability, status,
	// and configuration values related to virtualization offload for this
	// controller.
	VirtualizationOffload struct {
		// SRIOV is This object shall contain Single-Root Input/Output
		// Virtualization (SR-IOV) capabilities.
		SRIOV struct {
			// SRIOVVEPACapable is The value of this property shall be a boolean
			// indicating whether this controller supports Single Root Input/Output
			// Virtualization (SR-IOV) in Virtual Ethernet Port Aggregator (VEPA)
			// mode.
			SRIOVVEPACapable bool
		}
		// VirtualFunction is This property shall describe the capability,
		// status, and configuration values related to the virtual function for
		// this controller.
		VirtualFunction struct {
			// DeviceMaxCount is The value of this property shall be the maximum
			// number of Virtual Functions (VFs) supported by this controller.
			DeviceMaxCount int
			// MinAssignmentGroupSize is The value of this property shall be the
			// minimum number of Virtual Functions (VFs) that can be allocated or
			// moved between physical functions for this controller.
			MinAssignmentGroupSize int
			// NetworkPortMaxCount is The value of this property shall be the maximum
			// number of Virtual Functions (VFs) supported per network port for this
			// controller.
			NetworkPortMaxCount int
		}
	}
}

// UnmarshalJSON unmarshals a ControllerCapabilities object from the raw JSON.
func (controllercapabilities *ControllerCapabilities) UnmarshalJSON(b []byte) error {
	type temp ControllerCapabilities
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*controllercapabilities = ControllerCapabilities(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Controller is This type shall describe a network controller ASIC that
// makes up part of a NetworkAdapter.
type Controller struct {
	common.Entity

	// ControllerCapabilities is The value of this property shall contain the
	// capabilities of this controller.
	ControllerCapabilities ControllerCapabilities
	// FirmwarePackageVersion is The value of this property shall be the
	// version number of the user-facing firmware package.
	FirmwarePackageVersion string
	// NetworkDeviceFunctions is The value of this property shall be an array
	// of references of type NetworkDeviceFunction that represent the Network
	// Device Functions associated with this Network Controller.
	networkDeviceFunctions []string
	// NetworkDeviceFunctions@odata.count is
	NetworkDeviceFunctionsCount int
	// NetworkPorts is The value of this property shall be an array of
	// references of type NetworkPort that represent the Network Ports
	// associated with this Network Controller.
	networkPorts []string
	// NetworkPorts@odata.count is
	NetworkPortsCount int
	// PCIeDevices is The value of this property shall be an array of
	// references of type PCIeDevice that represent the PCI-e Devices
	// associated with this Network Controller.
	pcieDevices []string
	// PCIeDevices@odata.count is
	PCIeDevicesCount int
	// Location is This property shall contain location information of the
	// associated network adapter controller.
	Location string
	// PCIeInterface is used to connect this PCIe-based controller to its
	// host.
	PCIeInterface string
}

// UnmarshalJSON unmarshals a Controller object from the raw JSON.
func (controller *Controller) UnmarshalJSON(b []byte) error {
	type temp Controller
	type links struct {
		NetworkPorts                common.Links
		NetworkPortsCount           int `json:"EthernetInterfaces@odata.count"`
		NetworkDeviceFunctions      common.Links
		NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
		PCIeDevice                  common.Link
		PCIeDevicesCount            int `json:"PCIeDevices@odata.count"`
	}

	var t struct {
		temp
		Links links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*controller = Controller(t.temp)

	// Extract the links to other entities for later
	controller.networkPorts = t.Links.NetworkPorts.ToStrings()
	controller.NetworkPortsCount = t.Links.NetworkPortsCount
	controller.networkDeviceFunctions = t.Links.NetworkDeviceFunctions.ToStrings()
	controller.NetworkDeviceFunctionsCount = t.Links.NetworkDeviceFunctionsCount
	controller.pcieDevices = t.Links.NetworkDeviceFunctions.ToStrings()
	controller.PCIeDevicesCount = t.Links.NetworkDeviceFunctionsCount

	return nil
}

func (controller *Controller) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	var result []*NetworkDeviceFunction
	for _, uri := range controller.networkDeviceFunctions {
		n, err := GetNetworkDeviceFunction(controller.Client, uri)
		if err != nil {
			return nil, err
		}

		result = append(result, n)
	}

	return result, nil
}

// GetControllers will get a NetworkAdapter instance from the Redfish service.
func GetControllers(c common.Client, uri string) (*Controller, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var controller Controller
	err = json.NewDecoder(resp.Body).Decode(&controller)
	if err != nil {
		return nil, err
	}

	controller.SetClient(c)
	return &controller, nil
}

// ListReferencedControllers gets the collection of controllers from a provided reference.
func ListReferencedControllers(c common.Client, link string) ([]*Controller, error) {
	var result []*Controller
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, controllerLink := range links.ItemLinks {
		controller, err := GetControllers(c, controllerLink)
		if err != nil {
			return result, err
		}
		result = append(result, controller)
	}

	return result, nil
}

// NPIV is This type shall contain N_Port ID Virtualization (NPIV)
// capabilties for a controller.
type NPIV struct {
	common.Entity
}

// UnmarshalJSON unmarshals a NPIV object from the raw JSON.
func (npiv *NPIV) UnmarshalJSON(b []byte) error {
	type temp NPIV
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*npiv = NPIV(t.temp)

	// Extract the links to other entities for later

	return nil
}

// NetworkAdapter is A NetworkAdapter represents the physical network
// adapter capable of connecting to a computer network.  Examples include
// but are not limited to Ethernet, Fibre Channel, and converged network
// adapters.
type NetworkAdapter struct {
	common.Entity

	//// ODataContext is the odata context.
	//ODataContext string `json:"@odata.context"`
	//// ODataEtag is the odata etag.
	//ODataEtag string `json:"@odata.etag"`
	//// ODataID is the odata identifier.
	//ODataId string `json:"@odata.id"`
	//// ODataType is the odata type.
	//ODataType string `json:"@odata.type"`
	// Actions is The Actions property shall contain the available actions
	// for this resource.
	Actions string
	// Assembly is The value of this property shall be a link to a resource
	// of type Assembly.
	assembly string
	// Controller is The value of this property shall contain the set of
	// network controllers ASICs that make up this NetworkAdapter.
	Controllers []Controller
	// Description provides a description of this resource.
	Description string
	// Manufacturer is The value of this property shall contain a value that
	// represents the manufacturer of the network adapter.
	Manufacturer string
	// Model is The value of this property shall contain the information
	// about how the manufacturer references this network adapter.
	Model string
	// NetworkDeviceFunctions is The value of this property shall be a link
	// to a collection of type NetworkDeviceFunctionCollection.
	networkDeviceFunctions string
	// NetworkPorts is The value of this property shall be a link to a
	// collection of type NetworkPortCollection.
	networkPorts string
	// PartNumber is The value of this property shall contain the part number
	// for the network adapter as defined by the manufacturer.
	PartNumber string
	// SKU is The value of this property shall contain the Stock Keeping Unit
	// (SKU) for the network adapter.
	SKU string
	// SerialNumber is The value of this property shall contain the serial
	// number for the network adapter.
	SerialNumber string
	// Status is This property shall contain any status or health properties
	// of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a NetworkAdapter object from the raw JSON.
func (networkadapter *NetworkAdapter) UnmarshalJSON(b []byte) error {
	type temp NetworkAdapter
	var t struct {
		temp
		Assembly               common.Link
		NetworkDeviceFunctions common.Link
		NetworkPorts           common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*networkadapter = NetworkAdapter(t.temp)
	networkadapter.assembly = string(t.Assembly)
	networkadapter.networkDeviceFunctions = string(t.NetworkDeviceFunctions)
	networkadapter.networkPorts = string(t.NetworkPorts)

	return nil
}

// NetworkDeviceFunctions gets the collection of NetworkDeviceFunctions of this network adapter
func (networkadapter *NetworkAdapter) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	return ListReferencedNetworkDeviceFunctions(networkadapter.Client, networkadapter.networkDeviceFunctions)
}

//// Controller gets the collection of controllers of this network adapter
//func (networkadapter *NetworkAdapter) Controller() ([]*Controller, error) {
//	var result []*Controller
//	for _, uri := range networkadapter.controllers {
//		c, err := GetControllers(networkadapter.Client, uri)
//		if err != nil {
//			return nil, err
//		}
//
//		result = append(result, c)
//	}
//
//	return result, nil
//}

// GetNetworkAdapter will get a NetworkAdapter instance from the Redfish service.
func GetNetworkAdapter(c common.Client, uri string) (*NetworkAdapter, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var networkAdapter NetworkAdapter
	err = json.NewDecoder(resp.Body).Decode(&networkAdapter)
	if err != nil {
		return nil, err
	}

	networkAdapter.SetClient(c)
	return &networkAdapter, nil
}

// ListReferencedNetworkAdapter gets the collection of Chassis from a provided reference.
func ListReferencedNetworkAdapter(c common.Client, link string) ([]*NetworkAdapter, error) {
	var result []*NetworkAdapter
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, networkAdapterLink := range links.ItemLinks {
		networkAdapter, err := GetNetworkAdapter(c, networkAdapterLink)
		if err != nil {
			return result, err
		}
		result = append(result, networkAdapter)
	}

	return result, nil
}

// NicPartitioning is This type shall contain the capability, status, and
// configuration values for a controller.
type NicPartitioning struct {
	common.Entity
}

// UnmarshalJSON unmarshals a NicPartitioning object from the raw JSON.
func (nicpartitioning *NicPartitioning) UnmarshalJSON(b []byte) error {
	type temp NicPartitioning
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*nicpartitioning = NicPartitioning(t.temp)

	// Extract the links to other entities for later

	return nil
}
