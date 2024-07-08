//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

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

	activeSoftwareImage string
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
	// This property has been deprecated in favor of the Ports property.
	NetworkPortsCount int
	// PCIeDevices shall be an array of references of type PCIeDevice that
	// represent the PCI-e Devices associated with this Network Controller.
	pcieDevices []string
	// PCIeDevicesCount is the number of PCIeDevices.
	PCIeDevicesCount int
	ports            []string
	// PortsCount gets the number of ports associated with this network controller.
	PortsCount     int
	softwareImages []string
	// SoftwareImagesCount gets the number of firmware images that apply to this controller.
	SoftwareImagesCount int
}

// UnmarshalJSON unmarshals a Controllers object from the raw JSON.
func (controllers *Controllers) UnmarshalJSON(b []byte) error {
	type temp Controllers
	type links struct {
		ActiveSoftwareImage         common.Link
		NetworkDeviceFunctions      common.Links
		NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
		NetworkPorts                common.Links
		NetworkPortsCount           int `json:"EthernetInterfaces@odata.count"`
		PCIeDevice                  common.Link
		PCIeDevicesCount            int `json:"PCIeDevices@odata.count"`
		// Ports shall contain an array of links to resources of type Port that represent the ports associated with this
		// network controller.
		Ports      common.Links
		PortsCount int `json:"Ports@odata.count"`
		// SoftwareImages shall contain an array of links to resource of type SoftwareInventory that represent the firmware
		// images that apply to this controller.
		SoftwareImages      common.Links
		SoftwareImagesCount int `json:"SoftwareImages@odata.count"`
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
	controllers.activeSoftwareImage = t.Links.ActiveSoftwareImage.String()
	controllers.networkDeviceFunctions = t.Links.NetworkDeviceFunctions.ToStrings()
	controllers.NetworkDeviceFunctionsCount = t.Links.NetworkDeviceFunctionsCount
	controllers.networkPorts = t.Links.NetworkPorts.ToStrings()
	controllers.NetworkPortsCount = t.Links.NetworkPortsCount
	controllers.pcieDevices = t.Links.NetworkDeviceFunctions.ToStrings()
	controllers.PCIeDevicesCount = t.Links.NetworkDeviceFunctionsCount
	controllers.ports = t.Links.Ports.ToStrings()
	controllers.PortsCount = t.Links.PortsCount
	controllers.softwareImages = t.Links.SoftwareImages.ToStrings()
	controllers.SoftwareImagesCount = t.Links.SoftwareImagesCount

	return nil
}

// ActiveSoftwareImage gets the active firmware image for this network controller.
func (controllers *Controllers) ActiveSoftwareImage(c common.Client) (*SoftwareInventory, error) {
	if controllers.activeSoftwareImage == "" {
		return nil, nil
	}
	return GetSoftwareInventory(c, controllers.activeSoftwareImage)
}

// NetworkDeviceFunctions gets the collection of NetworkDeviceFunctions of this network controller.
func (controllers *Controllers) NetworkDeviceFunctions(c common.Client) ([]*NetworkDeviceFunction, error) {
	return common.GetObjects[NetworkDeviceFunction](c, controllers.networkDeviceFunctions)
}

// NetworkPorts gets the collection of NetworkPorts for this network controller.
func (controllers *Controllers) NetworkPorts(c common.Client) ([]*NetworkPort, error) {
	return common.GetObjects[NetworkPort](c, controllers.networkPorts)
}

// PCIeDevices gets the PCIe devices associated with this network controller.
func (controllers *Controllers) PCIeDevices(c common.Client) ([]*PCIeDevice, error) {
	return common.GetObjects[PCIeDevice](c, controllers.pcieDevices)
}

// Ports gets the ports associated with this network controller.
func (controllers *Controllers) Ports(c common.Client) ([]*Port, error) {
	return common.GetObjects[Port](c, controllers.ports)
}

// SoftwareImages gets the firmware images that apply to this controller.
func (controllers *Controllers) SoftwareImages(c common.Client) ([]*SoftwareInventory, error) {
	return common.GetObjects[SoftwareInventory](c, controllers.softwareImages)
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
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Assembly shall be a link to a resource of type Assembly.
	assembly string
	// Certificates shall contain a link to a resource collection of type CertificateCollection that contains
	// certificates for device identity and attestation.
	certificates string
	// Controllers shall contain the set of network controllers ASICs that make
	// up this NetworkAdapter.
	Controllers []Controllers
	// Description provides a description of this resource.
	Description string
	// EnvironmentMetrics shall contain a link to a resource of type EnvironmentMetrics that specifies the environment
	// metrics for this network adapter.
	environmentMetrics string
	// Identifiers shall contain a list of all known durable names for the network adapter.
	Identifiers []common.Identifier
	// LLDPEnabled shall contain the state indicating whether LLDP is globally enabled on a network adapter. If set to
	// 'false', the LLDPEnabled value for the ports associated with this adapter shall be disregarded.
	LLDPEnabled bool
	// Location shall contain the location information of the network adapter.
	Location common.Location
	// Manufacturer shall contain a value that represents the manufacturer of
	// the network adapter.
	Manufacturer string
	// Metrics are the metrics associated with this adapter.
	Metrics NetworkAdapterMetrics
	// Model shall contain the information about how the manufacturer references
	// this network adapter.
	Model string
	// NetworkDeviceFunctions shall be a link to a collection of type
	// NetworkDeviceFunctionCollection.
	networkDeviceFunctions string
	// NetworkPorts shall be a link to a collection of type NetworkPortCollection.
	// This property has been deprecated in favor of the Ports property.
	networkPorts string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain the part number for the network adapter as
	// defined by the manufacturer.
	PartNumber string
	// Ports shall contain a link to a resource collection of type PortCollection.
	ports string
	// Processors shall contain a link to a resource collection of type ProcessorCollection that represent the offload
	// processors contained in this network adapter.
	processors string
	// SKU shall contain the Stock Keeping Unit (SKU) for the network adapter.
	SKU string
	// SerialNumber shall contain the serial number for the network adapter.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	// resetSettingsToDefaultTarget is the URL for sending a ResetSettingsToDefault action
	resetSettingsToDefaultTarget string
}

// UnmarshalJSON unmarshals a NetworkAdapter object from the raw JSON.
func (networkadapter *NetworkAdapter) UnmarshalJSON(b []byte) error {
	type temp NetworkAdapter
	type actions struct {
		ResetSettingsToDefault common.ActionTarget `json:"#NetworkAdapter.ResetSettingsToDefault"`
	}
	var t struct {
		temp
		Assembly               common.Link
		Certificates           common.Link
		EnvironmentMetrics     common.Link
		NetworkDeviceFunctions common.Link
		NetworkPorts           common.Link
		Ports                  common.Link
		Processors             common.Link
		Actions                actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*networkadapter = NetworkAdapter(t.temp)
	networkadapter.assembly = t.Assembly.String()
	networkadapter.certificates = t.Certificates.String()
	networkadapter.environmentMetrics = t.EnvironmentMetrics.String()
	networkadapter.networkDeviceFunctions = t.NetworkDeviceFunctions.String()
	networkadapter.networkPorts = t.NetworkPorts.String()
	networkadapter.ports = t.Ports.String()
	networkadapter.processors = t.Processors.String()

	networkadapter.resetSettingsToDefaultTarget = t.Actions.ResetSettingsToDefault.Target

	return nil
}

// Update commits updates to this object's properties to the running system.
func (networkadapter *NetworkAdapter) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(NetworkAdapter)
	original.UnmarshalJSON(networkadapter.rawData)

	readWriteFields := []string{
		"LLDPEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(networkadapter).Elem()

	return networkadapter.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetNetworkAdapter will get a NetworkAdapter instance from the Redfish service.
func GetNetworkAdapter(c common.Client, uri string) (*NetworkAdapter, error) {
	return common.GetObject[NetworkAdapter](c, uri)
}

// ListReferencedNetworkAdapter gets the collection of Chassis from a provided reference.
func ListReferencedNetworkAdapter(c common.Client, link string) ([]*NetworkAdapter, error) {
	return common.GetCollectionObjects[NetworkAdapter](c, link)
}

// Assembly gets this adapter's assembly.
func (networkadapter *NetworkAdapter) Assembly() (*Assembly, error) {
	if networkadapter.assembly == "" {
		return nil, nil
	}
	return GetAssembly(networkadapter.GetClient(), networkadapter.assembly)
}

// Certificatea gets the certificates for device identity and attestation.
func (networkadapter *NetworkAdapter) Certificates() ([]*Certificate, error) {
	return ListReferencedCertificates(networkadapter.GetClient(), networkadapter.certificates)
}

// EnvironmentMetrics gets the environment metrics for this network adapter.
func (networkadapter *NetworkAdapter) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if networkadapter.environmentMetrics == "" {
		return nil, nil
	}
	return GetEnvironmentMetrics(networkadapter.GetClient(), networkadapter.environmentMetrics)
}

// NetworkDeviceFunctions gets the collection of NetworkDeviceFunctions of this network adapter
func (networkadapter *NetworkAdapter) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	return ListReferencedNetworkDeviceFunctions(networkadapter.GetClient(), networkadapter.networkDeviceFunctions)
}

// NetworkPorts gets the collection of NetworkPorts for this network adapter
func (networkadapter *NetworkAdapter) NetworkPorts() ([]*NetworkPort, error) {
	return ListReferencedNetworkPorts(networkadapter.GetClient(), networkadapter.networkPorts)
}

// Ports gets the ports associated with this network adapter.
func (networkadapter *NetworkAdapter) Ports() ([]*Port, error) {
	return ListReferencedPorts(networkadapter.GetClient(), networkadapter.ports)
}

// Processors gets the offload processors contained in this network adapter.
func (networkadapter *NetworkAdapter) Processors() ([]*Processor, error) {
	return ListReferencedProcessors(networkadapter.GetClient(), networkadapter.processors)
}

// ResetSettingsToDefault shall perform a reset of all active and pending
// settings back to factory default settings upon reset of the network adapter.
func (networkadapter *NetworkAdapter) ResetSettingsToDefault() error {
	return networkadapter.Post(networkadapter.resetSettingsToDefaultTarget, nil)
}
