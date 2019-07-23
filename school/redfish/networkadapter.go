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
	common.Entity

	// DataCenterBridging is This object shall contain capability, status,
	// and configuration values related to Data Center Bridging (DCB) for
	// this controller.
	DataCenterBridging string
	// NPAR is This object shall contain capability, status, and
	// configuration values related to NIC partitioning for this controller.
	NPAR string
	// NPIV is This object shall contain N_Port ID Virtualization (NPIV)
	// capabilties for this controller.
	NPIV string
	// NetworkDeviceFunctionCount is The value of this property shall be the
	// number of physical functions available on this controller.
	NetworkDeviceFunctionCount int
	// NetworkPortCount is The value of this property shall be the number of
	// physical ports on this controller.
	NetworkPortCount int
	// VirtualizationOffload is This object shall contain capability, status,
	// and configuration values related to virtualization offload for this
	// controller.
	VirtualizationOffload string
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

// ControllerLinks is This type, as described by the Redfish
// Specification, shall contain references to resources that are related
// to, but not contained by (subordinate to), this resource.
type ControllerLinks struct {
	common.Entity

	// NetworkDeviceFunctions is The value of this property shall be an array
	// of references of type NetworkDeviceFunction that represent the Network
	// Device Functions associated with this Network Controller.
	NetworkDeviceFunctions []NetworkDeviceFunction
	// NetworkDeviceFunctions@odata.count is
	NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
	// NetworkPorts is The value of this property shall be an array of
	// references of type NetworkPort that represent the Network Ports
	// associated with this Network Controller.
	NetworkPorts []NetworkPort
	// NetworkPorts@odata.count is
	NetworkPortsCount int `json:"NetworkPorts@odata.count"`
	// Oem is This object represents the Oem property.  All values for
	// resources described by this schema shall comply to the requirements as
	// described in the Redfish specification.
	OEM string `json:"Oem"`
	// PCIeDevices is The value of this property shall be an array of
	// references of type PCIeDevice that represent the PCI-e Devices
	// associated with this Network Controller.
	PCIeDevices []PCIeDevice
	// PCIeDevices@odata.count is
	PCIeDevicesCount int `json:"PCIeDevices@odata.count"`
}

// UnmarshalJSON unmarshals a ControllerLinks object from the raw JSON.
func (controllerlinks *ControllerLinks) UnmarshalJSON(b []byte) error {
	type temp ControllerLinks
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*controllerlinks = ControllerLinks(t.temp)

	// Extract the links to other entities for later

	return nil
}

// Controllers is This type shall describe a network controller ASIC that
// makes up part of a NetworkAdapter.
type Controllers struct {
	common.Entity

	// ControllerCapabilities is The value of this property shall contain the
	// capabilities of this controller.
	ControllerCapabilities string
	// FirmwarePackageVersion is The value of this property shall be the
	// version number of the user-facing firmware package.
	FirmwarePackageVersion string
	// Links is Links for this controller.
	Links string
	// Location is This property shall contain location information of the
	// associated network adapter controller.
	Location string
	// PCIeInterface is used to connect this PCIe-based controller to its
	// host.
	PCIeInterface string
}

// UnmarshalJSON unmarshals a Controllers object from the raw JSON.
func (controllers *Controllers) UnmarshalJSON(b []byte) error {
	type temp Controllers
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*controllers = Controllers(t.temp)

	// Extract the links to other entities for later

	return nil
}

// DataCenterBridging is This type shall describe the capability, status,
// and configuration values related to Data Center Bridging (DCB) for a
// controller.
type DataCenterBridging struct {
	common.Entity

	// Capable is The value of this property shall be a boolean indicating
	// whether this controller is capable of Data Center Bridging (DCB).
	Capable bool
}

// UnmarshalJSON unmarshals a DataCenterBridging object from the raw JSON.
func (datacenterbridging *DataCenterBridging) UnmarshalJSON(b []byte) error {
	type temp DataCenterBridging
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*datacenterbridging = DataCenterBridging(t.temp)

	// Extract the links to other entities for later

	return nil
}

// NPIV is This type shall contain N_Port ID Virtualization (NPIV)
// capabilties for a controller.
type NPIV struct {
	common.Entity

	// MaxDeviceLogins is The value of this property shall be the maximum
	// number of N_Port ID Virtualization (NPIV) logins allowed
	// simultaneously from all ports on this controller.
	MaxDeviceLogins int
	// MaxPortLogins is The value of this property shall be the maximum
	// number of N_Port ID Virtualization (NPIV) logins allowed per physical
	// port on this controller.
	MaxPortLogins int
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
	// Controllers is The value of this property shall contain the set of
	// network controllers ASICs that make up this NetworkAdapter.
	controllers []string
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
	// Oem is The value of this string shall be of the format for the
	// reserved word *Oem*.
	OEM string `json:"Oem"`
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
	type linkReference struct {
		Controllers common.Links
	}
	var t struct {
		temp
		Assembly               common.Link
		NetworkDeviceFunctions common.Link
		NetworkPorts           common.Link
		Links                  linkReference
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*networkadapter = NetworkAdapter(t.temp)
	networkadapter.controllers = t.Links.Controllers.ToStrings()
	networkadapter.assembly = string(t.Assembly)
	networkadapter.networkDeviceFunctions = string(t.NetworkDeviceFunctions)
	networkadapter.networkPorts = string(t.NetworkPorts)

	return nil
}

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

	// NparCapable is This property shall indicate the ability of a
	// controller to support NIC function partitioning.
	NparCapable bool
	// NparEnabled is This property shall indicate whether or not NIC
	// function partitioning is active on this controller.
	NparEnabled bool
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

// SRIOV is This type shall contain Single-Root Input/Output
// Virtualization (SR-IOV) capabilities.
type SRIOV struct {
	common.Entity

	// SRIOVVEPACapable is The value of this property shall be a boolean
	// indicating whether this controller supports Single Root Input/Output
	// Virtualization (SR-IOV) in Virtual Ethernet Port Aggregator (VEPA)
	// mode.
	SRIOVVEPACapable bool
}

// UnmarshalJSON unmarshals a SRIOV object from the raw JSON.
func (sriov *SRIOV) UnmarshalJSON(b []byte) error {
	type temp SRIOV
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*sriov = SRIOV(t.temp)

	// Extract the links to other entities for later

	return nil
}

// VirtualFunction is This type shall describe the capability, status,
// and configuration values related to a virtual function for a
// controller.
type VirtualFunction struct {
	common.Entity

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

// UnmarshalJSON unmarshals a VirtualFunction object from the raw JSON.
func (virtualfunction *VirtualFunction) UnmarshalJSON(b []byte) error {
	type temp VirtualFunction
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*virtualfunction = VirtualFunction(t.temp)

	// Extract the links to other entities for later

	return nil
}

// VirtualizationOffload is This type shall describe the capability,
// status, and configuration values related to a virtualization offload
// for a controller.
type VirtualizationOffload struct {
	common.Entity

	// SRIOV is This object shall contain Single-Root Input/Output
	// Virtualization (SR-IOV) capabilities.
	SRIOV string
	// VirtualFunction is This property shall describe the capability,
	// status, and configuration values related to the virtual function for
	// this controller.
	VirtualFunction string
}

// UnmarshalJSON unmarshals a VirtualizationOffload object from the raw JSON.
func (virtualizationoffload *VirtualizationOffload) UnmarshalJSON(b []byte) error {
	type temp VirtualizationOffload
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*virtualizationoffload = VirtualizationOffload(t.temp)

	// Extract the links to other entities for later

	return nil
}
