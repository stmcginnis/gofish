//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/NetworkAdapter.v1_14_0.json
// 2025.4 - #NetworkAdapter.v1_14_0.NetworkAdapter

package schemas

import (
	"encoding/json"
)

// NetworkAdapter shall represent a physical network adapter capable of
// connecting to a computer network in a Redfish implementation. Services should
// represent adapters that contain multiple controllers with independent
// management interfaces as multiple 'NetworkAdapter' resources.
type NetworkAdapter struct {
	Entity
	// Assembly shall contain a link to a resource of type 'Assembly'.
	//
	// Version added: v1.1.0
	assembly string
	// Certificates shall contain a link to a resource collection of type
	// 'CertificateCollection' that contains certificates for device identity and
	// attestation.
	//
	// Version added: v1.6.0
	certificates string
	// Controllers shall contain the set of network controllers ASICs that make up
	// this network adapter.
	Controllers []Controllers
	// EnvironmentMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that specifies the environment metrics for this network
	// adapter.
	//
	// Version added: v1.7.0
	environmentMetrics string
	// Identifiers shall contain a list of all known durable names for the network
	// adapter.
	//
	// Version added: v1.4.0
	Identifiers []Identifier
	// LLDPEnabled shall contain the state indicating whether LLDP is globally
	// enabled on a network adapter. If set to 'false', the 'LLDPEnabled' value for
	// the ports associated with this adapter shall be disregarded.
	//
	// Version added: v1.7.0
	LLDPEnabled bool
	// Location shall contain the location information of the network adapter.
	//
	// Version added: v1.4.0
	Location Location
	// Manufacturer shall contain a value that represents the manufacturer of the
	// network adapter.
	Manufacturer string
	// Measurements shall contain an array of DSP0274-defined measurement blocks.
	//
	// Version added: v1.6.0
	//
	// Deprecated: v1.9.0
	// This property has been deprecated in favor of the 'ComponentIntegrity'
	// resource.
	Measurements []MeasurementBlock
	// Metrics shall contain a link to a resource of type 'NetworkAdapterMetrics'
	// that contains the metrics associated with this adapter.
	//
	// Version added: v1.7.0
	metrics string
	// Model shall contain the information about how the manufacturer refers to
	// this network adapter.
	Model string
	// NetworkDeviceFunctions shall contain a link to a resource collection of type
	// 'NetworkDeviceFunctionCollection'.
	networkDeviceFunctions string
	// NetworkPorts shall contain a link to a resource collection of type
	// 'NetworkPortCollection'.
	//
	// Deprecated: v1.5.0
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
	// PartNumber shall contain the part number for the network adapter as defined
	// by the manufacturer.
	PartNumber string
	// PortAggregation shall contain capability, status, and configuration values
	// related to aggregating ports on this controller.
	//
	// Version added: v1.14.0
	PortAggregation PortAggregation
	// PortSplitting shall contain capability, status, and configuration values
	// related to physically subdividing the lanes of ports on this controller.
	//
	// Version added: v1.13.0
	PortSplitting PortSplitting
	// Ports shall contain a link to a resource collection of type
	// 'PortCollection'.
	//
	// Version added: v1.5.0
	ports string
	// Processors shall contain a link to a resource collection of type
	// 'ProcessorCollection' that represent the offload processors contained in
	// this network adapter.
	//
	// Version added: v1.8.0
	processors string
	// SKU shall contain the SKU for the network adapter.
	SKU string
	// SerialNumber shall contain the serial number for the network adapter.
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// resetSettingsToDefaultTarget is the URL to send ResetSettingsToDefault requests.
	resetSettingsToDefaultTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a NetworkAdapter object from the raw JSON.
func (n *NetworkAdapter) UnmarshalJSON(b []byte) error {
	type temp NetworkAdapter
	type nActions struct {
		Reset                  ActionTarget `json:"#NetworkAdapter.Reset"`
		ResetSettingsToDefault ActionTarget `json:"#NetworkAdapter.ResetSettingsToDefault"`
	}
	var tmp struct {
		temp
		Actions                nActions
		Assembly               Link `json:"Assembly"`
		Certificates           Link `json:"Certificates"`
		EnvironmentMetrics     Link `json:"EnvironmentMetrics"`
		Metrics                Link `json:"Metrics"`
		NetworkDeviceFunctions Link `json:"NetworkDeviceFunctions"`
		NetworkPorts           Link `json:"NetworkPorts"`
		Ports                  Link `json:"Ports"`
		Processors             Link `json:"Processors"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*n = NetworkAdapter(tmp.temp)

	// Extract the links to other entities for later
	n.resetTarget = tmp.Actions.Reset.Target
	n.resetSettingsToDefaultTarget = tmp.Actions.ResetSettingsToDefault.Target
	n.assembly = tmp.Assembly.String()
	n.certificates = tmp.Certificates.String()
	n.environmentMetrics = tmp.EnvironmentMetrics.String()
	n.metrics = tmp.Metrics.String()
	n.networkDeviceFunctions = tmp.NetworkDeviceFunctions.String()
	n.networkPorts = tmp.NetworkPorts.String()
	n.ports = tmp.Ports.String()
	n.processors = tmp.Processors.String()

	// This is a read/write object, so we need to save the raw object data for later
	n.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (n *NetworkAdapter) Update() error {
	readWriteFields := []string{
		"LLDPEnabled",
	}

	return n.UpdateFromRawData(n, n.RawData, readWriteFields)
}

// GetNetworkAdapter will get a NetworkAdapter instance from the service.
func GetNetworkAdapter(c Client, uri string) (*NetworkAdapter, error) {
	return GetObject[NetworkAdapter](c, uri)
}

// ListReferencedNetworkAdapters gets the collection of NetworkAdapter from
// a provided reference.
func ListReferencedNetworkAdapters(c Client, link string) ([]*NetworkAdapter, error) {
	return GetCollectionObjects[NetworkAdapter](c, link)
}

// This action shall reset a network adapter.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and shall perform a
// 'GracefulRestart'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (n *NetworkAdapter) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(n.client,
		n.resetTarget, payload, n.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset of all active and pending settings back to factory
// default settings upon reset of the network adapter.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (n *NetworkAdapter) ResetSettingsToDefault() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(n.client,
		n.resetSettingsToDefaultTarget, payload, n.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Assembly gets the Assembly linked resource.
func (n *NetworkAdapter) Assembly() (*Assembly, error) {
	if n.assembly == "" {
		return nil, nil
	}
	return GetObject[Assembly](n.client, n.assembly)
}

// Certificates gets the Certificates collection.
func (n *NetworkAdapter) Certificates() ([]*Certificate, error) {
	if n.certificates == "" {
		return nil, nil
	}
	return GetCollectionObjects[Certificate](n.client, n.certificates)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (n *NetworkAdapter) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if n.environmentMetrics == "" {
		return nil, nil
	}
	return GetObject[EnvironmentMetrics](n.client, n.environmentMetrics)
}

// Metrics gets the Metrics linked resource.
func (n *NetworkAdapter) Metrics() (*NetworkAdapterMetrics, error) {
	if n.metrics == "" {
		return nil, nil
	}
	return GetObject[NetworkAdapterMetrics](n.client, n.metrics)
}

// NetworkDeviceFunctions gets the NetworkDeviceFunctions collection.
func (n *NetworkAdapter) NetworkDeviceFunctions() ([]*NetworkDeviceFunction, error) {
	if n.networkDeviceFunctions == "" {
		return nil, nil
	}
	return GetCollectionObjects[NetworkDeviceFunction](n.client, n.networkDeviceFunctions)
}

// NetworkPorts gets the NetworkPorts collection.
func (n *NetworkAdapter) NetworkPorts() ([]*NetworkPort, error) {
	if n.networkPorts == "" {
		return nil, nil
	}
	return GetCollectionObjects[NetworkPort](n.client, n.networkPorts)
}

// Ports gets the Ports collection.
func (n *NetworkAdapter) Ports() ([]*Port, error) {
	if n.ports == "" {
		return nil, nil
	}
	return GetCollectionObjects[Port](n.client, n.ports)
}

// Processors gets the Processors collection.
func (n *NetworkAdapter) Processors() ([]*Processor, error) {
	if n.processors == "" {
		return nil, nil
	}
	return GetCollectionObjects[Processor](n.client, n.processors)
}

// ControllerCapabilities shall describe the capabilities of a controller.
type ControllerCapabilities struct {
	// DataCenterBridging shall contain capability, status, and configuration
	// values related to data center bridging (DCB) for this controller.
	DataCenterBridging DataCenterBridging
	// NPAR shall contain capability, status, and configuration values related to
	// NIC partitioning for this controller.
	//
	// Version added: v1.2.0
	NPAR NicPartitioning
	// NPIV shall contain N_Port ID Virtualization (NPIV) capabilities for this
	// controller.
	NPIV NPIV
	// NetworkDeviceFunctionCount shall contain the number of physical functions
	// available on this controller.
	NetworkDeviceFunctionCount *int `json:",omitempty"`
	// NetworkPortCount shall contain the number of physical ports on this
	// controller.
	NetworkPortCount *int `json:",omitempty"`
	// VirtualizationOffload shall contain capability, status, and configuration
	// values related to virtualization offload for this controller.
	VirtualizationOffload VirtualizationOffload
}

// ControllerLinks shall contain links to resources that are related to but are
// not contained by, or subordinate to, this resource.
type ControllerLinks struct {
	// ActiveSoftwareImage shall contain a link to a resource of type
	// 'SoftwareInventory' that represents the active firmware image for this
	// controller.
	//
	// Version added: v1.10.0
	activeSoftwareImage string
	// NetworkDeviceFunctions shall contain an array of links to resources of type
	// 'NetworkDeviceFunction' that represent the network device functions
	// associated with this network controller.
	networkDeviceFunctions []string
	// NetworkDeviceFunctionsCount
	NetworkDeviceFunctionsCount int `json:"NetworkDeviceFunctions@odata.count"`
	// NetworkPorts shall contain an array of links to resources of type
	// 'NetworkPort' that represent the network ports associated with this network
	// controller.
	//
	// Deprecated: v1.5.0
	// This property has been deprecated in favor of the 'Ports' property.
	networkPorts []string
	// NetworkPortsCount
	NetworkPortsCount int `json:"NetworkPorts@odata.count"`
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PCIeDevices shall contain an array of links to resources of type
	// 'PCIeDevice' that represent the PCIe devices associated with this network
	// controller.
	pCIeDevices []string
	// PCIeDevicesCount
	PCIeDevicesCount int `json:"PCIeDevices@odata.count"`
	// Ports shall contain an array of links to resources of type 'Port' that
	// represent the ports associated with this network controller.
	//
	// Version added: v1.5.0
	ports []string
	// PortsCount
	PortsCount int `json:"Ports@odata.count"`
	// SoftwareImages shall contain an array of links to resource of type
	// 'SoftwareInventory' that represent the firmware images that apply to this
	// controller.
	//
	// Version added: v1.10.0
	softwareImages []string
	// SoftwareImagesCount
	SoftwareImagesCount int `json:"SoftwareImages@odata.count"`
}

// UnmarshalJSON unmarshals a ControllerLinks object from the raw JSON.
func (c *ControllerLinks) UnmarshalJSON(b []byte) error {
	type temp ControllerLinks
	var tmp struct {
		temp
		ActiveSoftwareImage    Link  `json:"ActiveSoftwareImage"`
		NetworkDeviceFunctions Links `json:"NetworkDeviceFunctions"`
		NetworkPorts           Links `json:"NetworkPorts"`
		PCIeDevices            Links `json:"PCIeDevices"`
		Ports                  Links `json:"Ports"`
		SoftwareImages         Links `json:"SoftwareImages"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = ControllerLinks(tmp.temp)

	// Extract the links to other entities for later
	c.activeSoftwareImage = tmp.ActiveSoftwareImage.String()
	c.networkDeviceFunctions = tmp.NetworkDeviceFunctions.ToStrings()
	c.networkPorts = tmp.NetworkPorts.ToStrings()
	c.pCIeDevices = tmp.PCIeDevices.ToStrings()
	c.ports = tmp.Ports.ToStrings()
	c.softwareImages = tmp.SoftwareImages.ToStrings()

	return nil
}

// ActiveSoftwareImage gets the ActiveSoftwareImage linked resource.
func (c *ControllerLinks) ActiveSoftwareImage(client Client) (*SoftwareInventory, error) {
	if c.activeSoftwareImage == "" {
		return nil, nil
	}
	return GetObject[SoftwareInventory](client, c.activeSoftwareImage)
}

// NetworkDeviceFunctions gets the NetworkDeviceFunctions linked resources.
func (c *ControllerLinks) NetworkDeviceFunctions(client Client) ([]*NetworkDeviceFunction, error) {
	return GetObjects[NetworkDeviceFunction](client, c.networkDeviceFunctions)
}

// NetworkPorts gets the NetworkPorts linked resources.
func (c *ControllerLinks) NetworkPorts(client Client) ([]*NetworkPort, error) {
	return GetObjects[NetworkPort](client, c.networkPorts)
}

// PCIeDevices gets the PCIeDevices linked resources.
func (c *ControllerLinks) PCIeDevices(client Client) ([]*PCIeDevice, error) {
	return GetObjects[PCIeDevice](client, c.pCIeDevices)
}

// Ports gets the Ports linked resources.
func (c *ControllerLinks) Ports(client Client) ([]*Port, error) {
	return GetObjects[Port](client, c.ports)
}

// SoftwareImages gets the SoftwareImages linked resources.
func (c *ControllerLinks) SoftwareImages(client Client) ([]*SoftwareInventory, error) {
	return GetObjects[SoftwareInventory](client, c.softwareImages)
}

// Controllers shall describe a network controller ASIC that makes up part of a
// network adapter.
type Controllers struct {
	// ControllerCapabilities shall contain the capabilities of this controller.
	ControllerCapabilities ControllerCapabilities
	// FirmwarePackageVersion shall contain the version number of the user-facing
	// firmware package.
	FirmwarePackageVersion string
	// Identifiers shall contain a list of all known durable names for the
	// controller associated with the network adapter.
	//
	// Version added: v1.3.0
	Identifiers []Identifier
	// Location shall contain the location information of the controller associated
	// with the network adapter.
	//
	// Version added: v1.1.0
	Location Location
	// PCIeInterface shall contain details for the PCIe interface that connects
	// this PCIe-based controller to its host.
	//
	// Version added: v1.2.0
	PCIeInterface PCIeInterface
}

// DataCenterBridging shall describe the capability, status, and configuration
// values related to data center bridging (DCB) for a controller.
type DataCenterBridging struct {
	// Capable shall indicate whether this controller is capable of data center
	// bridging (DCB).
	Capable bool
}

// NPIV shall contain N_Port ID Virtualization (NPIV) capabilities for a
// controller.
type NPIV struct {
	// MaxDeviceLogins shall contain the maximum number of N_Port ID Virtualization
	// (NPIV) logins allowed simultaneously from all ports on this controller.
	MaxDeviceLogins *int `json:",omitempty"`
	// MaxPortLogins shall contain the maximum number of N_Port ID Virtualization
	// (NPIV) logins allowed per physical port on this controller.
	MaxPortLogins *int `json:",omitempty"`
}

// NicPartitioning shall contain the capability, status, and configuration
// values for a controller.
type NicPartitioning struct {
	// NparCapable shall indicate whether the controller supports NIC function
	// partitioning.
	//
	// Version added: v1.2.0
	NparCapable bool
	// NparEnabled shall indicate whether NIC function partitioning is active on
	// this controller.
	//
	// Version added: v1.2.0
	NparEnabled bool
}

// PortAggregation shall contain capability, status, and configuration values
// related to aggregating ports on a controller.
type PortAggregation struct {
	// AggregationEnabled shall indicate whether port aggregation is enabled.
	//
	// Version added: v1.14.0
	AggregationEnabled bool
	// AllowablePhysicalPortsPerAggregation shall contain the allowable number of
	// physical ports per aggregation.
	//
	// Version added: v1.14.0
	AllowablePhysicalPortsPerAggregation []*int
	// ConfiguredAggregatedPorts shall contain the configured number of aggregated
	// ports.
	//
	// Version added: v1.14.0
	ConfiguredAggregatedPorts *int `json:",omitempty"`
	// ConfiguredPhysicalPortsPerAggregation shall contain the configured number of
	// physical ports per aggregation.
	//
	// Version added: v1.14.0
	ConfiguredPhysicalPortsPerAggregation *int `json:",omitempty"`
	// TotalPhysicalPorts shall contain the total number of physical ports.
	//
	// Version added: v1.14.0
	TotalPhysicalPorts *int `json:",omitempty"`
}

// PortSplitting shall contain capability, status, and configuration values
// related to physically subdividing the lanes of ports on a controller.
type PortSplitting struct {
	// CurrentConfiguration shall contain the current port splitting configuration
	// for this controller.
	//
	// Version added: v1.13.0
	CurrentConfiguration []PortSplittingSubconfiguration
	// MaximumSubports shall contain the maximum number of subdivided ports that
	// this controller supports.
	//
	// Version added: v1.13.0
	MaximumSubports *int `json:",omitempty"`
	// MaximumSubportsPerPort shall contain the maximum number of subdivided ports
	// split from a single physical port that this controller supports.
	//
	// Version added: v1.13.0
	MaximumSubportsPerPort *int `json:",omitempty"`
	// SupportedConfigurations shall contain the port splitting configurations that
	// this controller supports. Properties contained in this property shall be
	// read-only.
	//
	// Version added: v1.13.0
	SupportedConfigurations []PortSplittingSubconfigurationList
}

// PortSplittingSubconfiguration shall contain a port splitting subconfiguration
// for one or more physical ports on a controller.
type PortSplittingSubconfiguration struct {
	// EndingPhysicalPort shall contain the last physical port to which this
	// subconfiguration applies. Specifically, the splitting for ports 'StartPort'
	// through 'EndPort', inclusive, is characterized by this subconfiguration.
	// Each set of subconfigurations shall cover all physical ports on the
	// controller and shall describe splitting for each port exactly once.
	//
	// Version added: v1.13.0
	EndingPhysicalPort *int `json:",omitempty"`
	// FirstSubportID shall contain the first identifier to assign to subports in
	// this subconfiguration. Subport identifiers shall be assigned sequentially to
	// the subports starting with those for 'StartPort' and working through to
	// those for 'EndPort'.
	//
	// Version added: v1.13.0
	FirstSubportID *int `json:"FirstSubportId,omitempty"`
	// Lanes shall contain the number of lanes for each subport. The number of
	// members in this array shall equal the value contained in 'SubportsPerPort'
	// in the enclosing subconfiguration.
	//
	// Version added: v1.13.0
	Lanes []*int
	// LinkSpeedGbps shall contain the configured link speed for each subport. The
	// number of members in this array shall equal the value contained in
	// 'SubportsPerPort' in the enclosing subconfiguration.
	//
	// Version added: v1.13.0
	LinkSpeedGbps []*int
	// StartingPhysicalPort shall contain the first physical port to which this
	// subconfiguration applies.
	//
	// Version added: v1.13.0
	StartingPhysicalPort *int `json:",omitempty"`
	// SubportsPerPort shall contain the number of subports created from each port
	// in this subconfiguration.
	//
	// Version added: v1.13.0
	SubportsPerPort *int `json:",omitempty"`
}

// PortSplittingSubconfigurationList shall contain a port splitting
// configuration for a controller.
type PortSplittingSubconfigurationList struct {
	// Subconfigurations shall contain the set of subconfigurations that
	// collectively define a port splitting configuration for this controller.
	//
	// Version added: v1.13.0
	Subconfigurations []PortSplittingSubconfiguration
}

// SRIOV shall contain single-root input/output virtualization (SR-IOV)
// capabilities.
type SRIOV struct {
	// SRIOVEnabled shall indicate whether single root input/output virtualization
	// (SR-IOV) is enabled for this controller.
	//
	// Version added: v1.12.0
	SRIOVEnabled bool
	// SRIOVVEPACapable shall indicate whether this controller supports single root
	// input/output virtualization (SR-IOV) in Virtual Ethernet Port Aggregator
	// (VEPA) mode.
	SRIOVVEPACapable bool
}

// VirtualFunction shall describe the capability, status, and configuration
// values related to a virtual function for a controller.
type VirtualFunction struct {
	// DeviceMaxCount shall contain the maximum number of virtual functions
	// supported by this controller.
	DeviceMaxCount *int `json:",omitempty"`
	// MinAssignmentGroupSize shall contain the minimum number of virtual functions
	// that can be allocated or moved between physical functions for this
	// controller.
	MinAssignmentGroupSize *int `json:",omitempty"`
	// NetworkPortMaxCount shall contain the maximum number of virtual functions
	// supported per network port for this controller.
	NetworkPortMaxCount *int `json:",omitempty"`
}

// VirtualizationOffload shall describe the capability, status, and
// configuration values related to a virtualization offload for a controller.
type VirtualizationOffload struct {
	// SRIOV shall contain single-root input/output virtualization (SR-IOV)
	// capabilities.
	SRIOV SRIOV
	// VirtualFunction shall describe the capability, status, and configuration
	// values related to the virtual function for this controller.
	VirtualFunction VirtualFunction
}
