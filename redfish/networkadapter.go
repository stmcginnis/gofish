//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// ControllerCapabilities shall describe the capabilities of a controller.
type ControllerCapabilities struct {
	// DataCenterBridging shall contain capability, status,
	// and configuration values related to Data Center Bridging (DCB) for
	// this controller.
	DataCenterBridging struct {
		Capable bool
	}
	// NPAR shall contain capability, status, and
	// configuration values related to NIC partitioning for this controller.
	NPAR struct {
		// NparCapable shall indicate the ability of a
		// controller to support NIC function partitioning.
		NparCapable bool
		// NparEnabled shall indicate whether or not NIC
		// function partitioning is active on this controller.
		NparEnabled bool
	}
	// NPIV shall contain N_Port ID Virtualization (NPIV)
	// capabilities for this controller.
	NPIV struct {
		// MaxDeviceLogins shall be the maximum
		// number of N_Port ID Virtualization (NPIV) logins allowed
		// simultaneously from all ports on this controller.
		MaxDeviceLogins int
		// MaxPortLogins shall be the maximum
		// number of N_Port ID Virtualization (NPIV) logins allowed per physical
		// port on this controller.
		MaxPortLogins int
	}
	// NetworkDeviceFunctionCount shall be the
	// number of physical functions available on this controller.
	NetworkDeviceFunctionCount int
	// NetworkPortCount shall be the number of
	// physical ports on this controller.
	NetworkPortCount int
	// VirtualizationOffload shall contain capability, status,
	// and configuration values related to virtualization offload for this
	// controller.
	VirtualizationOffload struct {
		// SRIOV shall contain Single-Root Input/Output Virtualization (SR-IOV)
		// capabilities.
		SRIOV struct {
			// SRIOVVEPACapable shall be a boolean indicating whether this
			// controller supports Single Root Input/Output Virtualization
			// (SR-IOV) in Virtual Ethernet Port Aggregator (VEPA) mode.
			SRIOVVEPACapable bool
		}
		// VirtualFunction shall describe the capability, status, and
		// configuration values related to the virtual function for this controller.
		VirtualFunction struct {
			// DeviceMaxCount shall be the maximum number of Virtual Functions
			// (VFs) supported by this controller.
			DeviceMaxCount int
			// MinAssignmentGroupSize shall be the minimum number of Virtual
			// Functions (VFs) that can be allocated or moved between physical
			// functions for this controller.
			MinAssignmentGroupSize int
			// NetworkPortMaxCount shall be the maximum number of Virtual
			// Functions (VFs) supported per network port for this controller.
			NetworkPortMaxCount int
		}
	}
}

// Controllers shall describe a network controller ASIC that makes up part of a
// NetworkAdapter.
type Controllers struct {
	// ControllerCapabilities shall contain the capabilities of this controller.
	ControllerCapabilities ControllerCapabilities
	// FirmwarePackageVersion shall be the version number of the user-facing
	// firmware package.
	FirmwarePackageVersion string
	// Identifiers shall contain a list of all known durable names for the
	// associated network adapter.
	Identifiers []common.Identifier
	// Location shall contain location information of the associated network
	// adapter controller.
	Location common.Location
	// PCIeInterface is used to connect this PCIe-based controller to its host.
	PCIeInterface PCIeInterface
	// NetworkDeviceFunctions shall be an array of references of type
	// NetworkDeviceFunction that represent the Network Device Functions
	// associated with this Network Controller.
	networkDeviceFunctions []string
	// NetworkDeviceFunctionsCount is the number of network device functions.
	NetworkDeviceFunctionsCount int
	// NetworkPorts shall be an array of references of type NetworkPort that
	// represent the Network Ports associated with this Network Controller.
	networkPorts []string
	// NetworkPortsCount is the number of network ports.
	NetworkPortsCount int
	// PCIeDevices shall be an array of references of type PCIeDevice that
	// represent the PCI-e Devices associated with this Network Controller.
	pcieDevices []string
	// PCIeDevicesCount is the number of PCIeDevices.
	PCIeDevicesCount int
}

// UnmarshalJSON unmarshals a Controllers object from the raw JSON.
func (controllers *Controllers) UnmarshalJSON(b []byte) error {
	type temp Controllers
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

	// Extract the links to other entities for later
	*controllers = Controllers(t.temp)
	controllers.networkPorts = t.Links.NetworkPorts.ToStrings()
	controllers.NetworkPortsCount = t.Links.NetworkPortsCount
	controllers.networkDeviceFunctions = t.Links.NetworkDeviceFunctions.ToStrings()
	controllers.NetworkDeviceFunctionsCount = t.Links.NetworkDeviceFunctionsCount
	controllers.pcieDevices = t.Links.NetworkDeviceFunctions.ToStrings()
	controllers.PCIeDevicesCount = t.Links.NetworkDeviceFunctionsCount

	return nil
}

// DataCenterBridging shall describe the capability, status,
// and configuration values related to Data Center Bridging (DCB) for a
// controller.
type DataCenterBridging struct {
	// Capable shall be a boolean indicating whether this controller is capable
	// of Data Center Bridging (DCB).
	Capable bool
}

// NPIV shall contain N_Port ID Virtualization (NPIV) capabilities for a
// controller.
type NPIV struct {

	// MaxDeviceLogins shall be the maximum number of N_Port ID Virtualization
	// (NPIV) logins allowed simultaneously from all ports on this controller.
	MaxDeviceLogins int
	// MaxPortLogins shall be the maximum number of N_Port ID Virtualization
	// (NPIV) logins allowed per physical port on this controller.
	MaxPortLogins int
}

// A NetworkAdapter represents the physical network adapter capable of
// connecting to a computer network. Examples include but are not limited to
// Ethernet, Fibre Channel, and converged network adapters.
type NetworkAdapter struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Assembly shall be a link to a resource of type Assembly.
	assembly string
	// Controllers shall contain the set of network controllers ASICs that make
	// up this NetworkAdapter.
	Controllers []Controllers
	// Description provides a description of this resource.
	Description string
	// Manufacturer shall contain a value that represents the manufacturer of
	// the network adapter.
	Manufacturer string
	// Model shall contain the information about how the manufacturer references
	// this network adapter.
	Model string
	// NetworkDeviceFunctions shall be a link to a collection of type
	// NetworkDeviceFunctionCollection.
	networkDeviceFunctions string
	// NetworkPorts shall be a link to a collection of type NetworkPortCollection.
	networkPorts string
	// PartNumber shall contain the part number for the network adapter as
	// defined by the manufacturer.
	PartNumber string
	// SKU shall contain the Stock Keeping Unit (SKU) for the network adapter.
	SKU string
	// SerialNumber shall contain the serial number for the network adapter.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// resetSettingsToDefaultTarget is the URL for sending a ResetSettingsToDefault action
	resetSettingsToDefaultTarget string
}

// UnmarshalJSON unmarshals a NetworkAdapter object from the raw JSON.
func (networkadapter *NetworkAdapter) UnmarshalJSON(b []byte) error {
	type temp NetworkAdapter
	type actions struct {
		ResetSettingsToDefault struct {
			Target string
		} `json:"#NetworkAdapter.ResetSettingsToDefault"`
	}
	var t struct {
		temp
		Assembly               common.Link
		NetworkDeviceFunctions common.Link
		NetworkPorts           common.Link
		Actions                actions
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
	networkadapter.resetSettingsToDefaultTarget = t.Actions.ResetSettingsToDefault.Target

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

// Assembly gets this adapter's assembly.
func (networkadapter *NetworkAdapter) Assembly() (*Assembly, error) {
	if networkadapter.assembly == "" {
		return nil, nil
	}
	return GetAssembly(networkadapter.Client, networkadapter.assembly)
}

// NetworkDeviceFunctions gets the collection of NetworkDeviceFunctions of this network adapter
func (networkadapter *NetworkAdapter) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	return ListReferencedNetworkDeviceFunctions(networkadapter.Client, networkadapter.networkDeviceFunctions)
}

// NetworkPorts gets the collection of NetworkPorts for this network adapter
func (networkadapter *NetworkAdapter) NetworkPorts() ([]*NetworkPort, error) {
	return ListReferencedNetworkPorts(networkadapter.Client, networkadapter.networkPorts)
}

// ResetSettingsToDefault shall perform a reset of all active and pending
// settings back to factory default settings upon reset of the network adapter.
func (networkadapter *NetworkAdapter) ResetSettingsToDefault() error {
	_, err := networkadapter.Client.Post(networkadapter.resetSettingsToDefaultTarget, nil)
	return err
}
