//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type Congestion struct {
	// The CXL Specification-defined 'Backpressure Sample Interval' in nanoseconds.
	BackpressureSampleInterval int
	// The CXL Specification-defined 'Completion Collection Interval' in nanoseconds.
	CompletionCollectionInterval int
	// Indicates whether congestion telemetry collection is enabled for this port.
	CongestionTelemetryEnabled bool
	// The threshold for moderate egress port congestion as a percentage.
	EgressModeratePercentage int
	// The threshold for severe egress port congestion as a percentage.
	EgressSeverePercentage int
	// The CXL Specification-defined 'ReqCmpBasis'.
	MaxSustainedRequestCmpBias int
}

type ConnectedDeviceMode string

const (
	// CXL 68B flit and VH.
	CXL68BFlitAndVHDeviceMode ConnectedDeviceMode = "CXL68BFlitAndVH"
	// CXL latency-optimized 256B flit.
	CXLLatencyOptimized256BFlitDeviceMode ConnectedDeviceMode = "CXLLatencyOptimized256BFlit"
	// The connection is not CXL or is disconnected.
	DisconnectedDeviceMode ConnectedDeviceMode = "Disconnected"
	// Port-based routing (PBR).
	PBRDeviceMode ConnectedDeviceMode = "PBR"
	// Restricted CXL device (RCD).
	RCDDeviceMode ConnectedDeviceMode = "RCD"
	// Standard 256B flit.
	Standard256BFlitDeviceMode ConnectedDeviceMode = "Standard256BFlit"
)

type ConnectedDeviceType string

const (
	// No device detected.
	NoneDeviceType ConnectedDeviceType = "None"
	// PCIe device.
	PCIeDeviceDeviceType ConnectedDeviceType = "PCIeDevice"
	// CXL Type 1 device.
	Type1DeviceType ConnectedDeviceType = "Type1"
	// CXL Type 2 device.
	Type2DeviceType ConnectedDeviceType = "Type2"
	// CXL Type 3 multi-logical device (MLD).
	Type3MLDDeviceType ConnectedDeviceType = "Type3MLD"
	// CXL Type 3 single logical device (SLD).
	Type3SLDDeviceType ConnectedDeviceType = "Type3SLD"
)

type CurrentPortConfigurationState string

const (
	// Bind in progress.
	BindInProgressPortConfigurationState CurrentPortConfigurationState = "BindInProgress"
	// Disabled.
	DisabledPortConfigurationState CurrentPortConfigurationState = "Disabled"
	// Downstream port (DSP).
	DSPPortConfigurationState CurrentPortConfigurationState = "DSP"
	// Reserved.
	ReservedPortConfigurationState CurrentPortConfigurationState = "Reserved"
	// Unbind in progress.
	UnbindInProgressPortConfigurationState CurrentPortConfigurationState = "UnbindInProgress"
	// Upstream port (USP).
	USPPortConfigurationState CurrentPortConfigurationState = "USP"
)

type QoSTelemetryCapabilities struct {
	// Indicates whether the port supports the CXL Specification-defined 'Egress Port Backpressure' mechanism.
	EgressPortBackpressureSupported bool
	// Indicates whether the port supports the CXL Specification-defined 'Temporary Throughput Reduction' mechanism.
	TemporaryThroughputReductionSupported bool
}

type CXL struct {
	// The congestion properties for this CXL port.
	Congestion Congestion
	// The connected device mode.
	ConnectedDeviceMode ConnectedDeviceMode
	// The connected device type.
	ConnectedDeviceType ConnectedDeviceType
	// The current port configuration state.
	CurrentPortConfigurationState CurrentPortConfigurationState
	// The maximum number of logical devices supported.
	MaxLogicalDeviceCount int
	// The quality of service telemetry capabilities for this CXL port.
	QoSTelemetryCapabilities QoSTelemetryCapabilities
	// Indicates whether temporary throughput reduction is enabled.
	TemporaryThroughputReductionEnabled bool
}

type IDSubtype string

const (
	// Agent circuit ID, based on the agent-local identifier of the circuit as defined in RFC3046.
	AgentIDIDSubtype IDSubtype = "AgentId"
	// Chassis component, based in the value of entPhysicalAlias in RFC4133.
	ChassisCompIDSubtype IDSubtype = "ChassisComp"
	// Interface alias, based on the ifAlias MIB object.
	IfAliasIDSubtype IDSubtype = "IfAlias"
	// Interface name, based on the ifName MIB object.
	IfNameIDSubtype IDSubtype = "IfName"
	// Locally assigned, based on a alpha-numeric value locally assigned.
	LocalAssignIDSubtype IDSubtype = "LocalAssign"
	// MAC address, based on an agent detected unicast source address as defined in IEEE standard 802.
	MacAddrIDSubtype IDSubtype = "MacAddr"
	// Network address, based on an agent detected network address.
	NetworkAddrIDSubtype IDSubtype = "NetworkAddr"
	// No data to be sent to/received from remote partner.
	NotTransmittedIDSubtype IDSubtype = "NotTransmitted"
	// Port component, based in the value of entPhysicalAlias in RFC4133.
	PortCompIDSubtype IDSubtype = "PortComp"
)

type SystemCapability string

const (
	// Bridge.
	BridgeSystemCapability SystemCapability = "Bridge"
	// DOCSIS cable device.
	DOCSISCableDeviceSystemCapability SystemCapability = "DOCSISCableDevice"
	// The system capabilities are transmitted, but no capabilities are set.
	NoneSystemCapability SystemCapability = "None"
	// Other.
	OtherSystemCapability SystemCapability = "Other"
	// Repeater.
	RepeaterSystemCapability SystemCapability = "Repeater"
	// Router.
	RouterSystemCapability SystemCapability = "Router"
	// Station.
	StationSystemCapability SystemCapability = "Station"
	// Telephone
	TelephoneSystemCapability SystemCapability = "Telephone"
	// WLAN access point.
	WLANAccessPointSystemCapability SystemCapability = "WLANAccessPoint"
)

type LLDP struct {
	// Link Layer Data Protocol (LLDP) chassis ID.
	ChassisID string `json:"ChassisId"`
	// The type of identifier used for the chassis ID.
	ChassisIDSubtype IDSubtype `json:"ChassisIdSubtype"`
	// The IPv4 management address to be transmitted from this endpoint.
	ManagementAddressIPv4 string
	// The IPv6 management address to be transmitted from this endpoint.
	ManagementAddressIPv6 string
	// The management MAC address to be transmitted from this endpoint.
	ManagementAddressMAC string
	// The management VLAN ID to be transmitted from this endpoint.
	ManagementVlanID int `json:"ManagementVlanId"`
	// A colon delimited string of hexadecimal octets identifying a port to be transmitted from this endpoint.
	PortID string `json:"PortId"`
	// The port ID subtype to be transmitted from this endpoint.
	PortIDSubtype IDSubtype `json:"PortIdSubtype"`
	// The system capabilities to be transmitted from this endpoint.
	SystemCapabilities []SystemCapability
	// The system description to be transmitted from this endpoint.
	SystemDescription string
	// The system name to be transmitted from this endpoint.
	SystemName string
}

type PortEthernet struct {
	// An array of configured MAC addresses that are associated with this network port.
	AssociatedMACAddresses []string
	// Indicates whether IEEE 802.3az Energy-Efficient Ethernet (EEE) is enabled on this port.
	EEEEnabled bool
	// The locally configured 802.3x flow control setting for this port.
	FlowControlConfiguration FlowControl
	// The 802.3x flow control behavior negotiated with the link partner for this port.
	FlowControlStatus FlowControl
	// Enable/disable LLDP for this port.
	LLDPEnabled bool
	// LLDP data being received on this link.
	LLDPReceive LLDP
	// LLDP data being transmitted on this link.
	LLDPTransmit LLDP
	// Deprecated (v1.5+): The set of Ethernet capabilities that this port supports.
	SupportedEthernetCapabilities []SupportedEthernetCapabilities
	// Indicates whether Wake on LAN (WoL) is enabled on this port.
	WakeOnLANEnabled bool
}

type PortFibreChannel struct {
	// An array of configured World Wide Names (WWN) that are associated with this network port.
	AssociatedWorldWideNames []string
	// The Fibre Channel Fabric Name provided by the switch.
	FabricName string
	// The number of ports not on the associated device that the associated device has discovered through this port.
	NumberDiscoveredRemotePorts int
	// The connection type of this port.
	PortConnectionType PortConnectionType
}

type FunctionBandwidth struct {
	// The bandwidth allocation percentage allocated to the corresponding network device function instance.
	AllocationPercent int
	// The link to the network device function associated with this bandwidth setting of this network port.
	networkDeviceFunction string
}

type PortGenZ struct {
	// The Linear Packet Relay Table for the port.
	lrpt string
	// The Multi-subnet Packet Relay Table for the port.
	mrpt string
	// The Virtual Channel Action Table for the port.
	vcat string
}

type InfiniBand struct {
	// An array of configured node GUIDs that are associated with this network port,
	// including the programmed address of the lowest numbered network device function,
	// the configured but not active address, if applicable,
	// the address for hardware port teaming, or other network addresses.
	AssociatedNodeGUIDs []string
	// An array of configured port GUIDs that are associated with this network port,
	// including the programmed address of the lowest numbered network device function,
	// the configured but not active address, if applicable,
	// the address for hardware port teaming, or other network addresses.
	AssociatedPortGUIDs []string
	// An array of configured system GUIDs that are associated with this network port,
	// including the programmed address of the lowest numbered network device function,
	// the configured but not active address, if applicable,
	// the address for hardware port teaming, or other network addresses.
	AssociatedSystemGUIDs []string
}

type ConfiguredNetworkLink struct {
	// The link speed per lane this port is configured to use for autonegotiation.
	ConfiguredLinkSpeedGbps float32
	// The link width this port is configured to use for autonegotiation in conjunction with the link speed.
	ConfiguredWidth int
}

type LinkConfiguration struct {
	// An indication of whether the port is capable of autonegotiating speed.
	AutoSpeedNegotiationCapable bool
	// Controls whether this port is configured to enable autonegotiating speed.
	AutoSpeedNegotiationEnabled bool
	// The set of link speed capabilities of this port.
	CapableLinkSpeedGbps []float32
	// 	The set of link speed and width pairs this port is configured to use for autonegotiation.
	ConfiguredNetworkLinks []ConfiguredNetworkLink
}

type LinkState string

const (
	// The link is disabled and not operational.
	DisabledLinkState LinkState = "Disabled"
	// The link is enabled and operational.
	EnabledLinkState LinkState = "Enabled"
)

type PortLinkStatus string

const (
	// The link on this interface is down.
	LinkDownPortLinkStatus PortLinkStatus = "LinkDown"
	// The link on this interface is up.
	LinkUpPortLinkStatus PortLinkStatus = "LinkUp"
	// No physical link detected on this interface.
	NoLinkPortLinkStatus PortLinkStatus = "NoLink"
	// This link on this interface is starting. A physical link has been established,
	// but the port is not able to transfer data.
	StartingPortLinkStatus PortLinkStatus = "Starting"
	// This physical link on this interface is training.
	TrainingPortLinkStatus PortLinkStatus = "Training"
)

type PortMedium string

const (
	// This port has an electrical cable connection.
	ElectricalPortMedium PortMedium = "Electrical"
	// This port has an optical cable connection.
	OpticalPortMedium PortMedium = "Optical"
)

type PortProtocol string

const (
	// Advanced Host Controller Interface (AHCI).
	AHCIPortProtocol PortProtocol = "AHCI"
	// Compute Express Link.
	CXLPortProtocol PortProtocol = "CXL"
	// DisplayPort.
	DisplayPortPortProtocol PortProtocol = "DisplayPort"
	// DVI.
	DVIPortProtocol PortProtocol = "DVI"
	// Ethernet.
	EthernetPortProtocol PortProtocol = "Ethernet"
	// Fibre Channel.
	FCPortProtocol PortProtocol = "FC"
	// Fibre Channel over Ethernet (FCoE).
	FCoEPortProtocol PortProtocol = "FCoE"
	// Fibre Channel Protocol for SCSI.
	FCPPortProtocol PortProtocol = "FCP"
	// FIbre CONnection (FICON).
	FICONPortProtocol PortProtocol = "FICON"
	// File Transfer Protocol (FTP).
	FTPPortProtocol PortProtocol = "FTP"
	// GenZ.
	GenZPortProtocol PortProtocol = "GenZ"
	// HDMI.
	HDMIPortProtocol PortProtocol = "HDMI"
	// Hypertext Transport Protocol (HTTP).
	HTTPPortProtocol PortProtocol = "HTTP"
	// Hypertext Transfer Protocol Secure (HTTPS).
	HTTPSPortProtocol PortProtocol = "HTTPS"
	// Inter-Integrated Circuit Bus.
	I2CPortProtocol PortProtocol = "I2C"
	// InfiniBand.
	InfiniBandPortProtocol PortProtocol = "InfiniBand"
	// Internet SCSI.
	ISCSIPortProtocol PortProtocol = "iSCSI"
	// Internet Wide Area RDMA Protocol (iWARP).
	IWARPPortProtocol PortProtocol = "iWARP"
	// Multiple Protocols.
	MultiProtocolPortProtocol PortProtocol = "MultiProtocol"
	// Network File System (NFS) version 3.
	NFSv3PortProtocol PortProtocol = "NFSv3"
	// Network File System (NFS) version 4.
	NFSv4PortProtocol PortProtocol = "NFSv4"
	// NVLink.
	NVLinkPortProtocol PortProtocol = "NVLink"
	// Non-Volatile Memory Express (NVMe).
	NVMePortProtocol PortProtocol = "NVMe"
	// NVMe over Fabrics.
	NVMeOverFabricsPortProtocol PortProtocol = "NVMeOverFabrics"
	// OEM-specific.
	OEMPortProtocol PortProtocol = "OEM"
	// PCI Express.
	PCIePortProtocol PortProtocol = "PCIe"
	// Intel QuickPath Interconnect (QPI).
	QPIPortProtocol PortProtocol = "QPI"
	// RDMA over Converged Ethernet Protocol.
	RoCEPortProtocol PortProtocol = "RoCE"
	// RDMA over Converged Ethernet Protocol Version 2.
	RoCEv2PortProtocol PortProtocol = "RoCEv2"
	// Serial Attached SCSI.
	SASPortProtocol PortProtocol = "SAS"
	// Serial AT Attachment.
	SATAPortProtocol PortProtocol = "SATA"
	// SSH File Transfer Protocol (SFTP).
	SFTPPortProtocol PortProtocol = "SFTP"
	// Server Message Block (SMB).
	// Also known as the Common Internet File System (CIFS).
	SMBPortProtocol PortProtocol = "SMB"
	// Transmission Control Protocol (TCP).
	TCPPortProtocol PortProtocol = "TCP"
	// Trivial File Transfer Protocol (TFTP).
	TFTPPortProtocol PortProtocol = "TFTP"
	// User Datagram Protocol (UDP).
	UDPPortProtocol PortProtocol = "UDP"
	// Universal Host Controller Interface (UHCI).
	UHCIPortProtocol PortProtocol = "UHCI"
	// Intel UltraPath Interconnect (UPI).
	UPIPortProtocol PortProtocol = "UPI"
	// Universal Serial Bus (USB).
	USBPortProtocol PortProtocol = "USB"
	// VGA.
	VGAPortProtocol PortProtocol = "VGA"
)

type PortType string

const (
	// This port connects to any type of device.
	BidirectionalPortType PortType = "BidirectionalPort"
	// This port connects to a target device.
	DownstreamPortType PortType = "DownstreamPort"
	// This port connects to another switch.
	InterswitchPortType PortType = "InterswitchPort"
	// This port connects to a switch manager.
	ManagementPortType PortType = "ManagementPort"
	// This port has not yet been configured.
	UnconfiguredPortType PortType = "UnconfiguredPort"
	// This port connects to a host device.
	UpstreamPortType PortType = "UpstreamPort"
)

type FiberConnectionType string

const (
	// The connection is using multi mode operation.
	MultiModeFiberConnectionType FiberConnectionType = "MultiMode"
	// The connection is using single mode operation.
	SingleModeFiberConnectionType FiberConnectionType = "SingleMode"
)

type MediumType string

const (
	// The medium connected is copper.
	CopperMediumType MediumType = "Copper"
	// The medium connected is fiber optic.
	FiberOpticMediumType MediumType = "FiberOptic"
)

type SFPDeviceType string

const (
	// The SFP conforms to the CSFP MSA Specification.
	CSFPSFPDeviceType SFPDeviceType = "cSFP"
	// The SFP conforms to the SFF Specification SFF-8644.
	MiniSASHDSFPDeviceType SFPDeviceType = "MiniSASHD"
	// The SFP conforms to the OSFP Specification.
	OSFPSFPDeviceType SFPDeviceType = "OSFP"
	// The SFP conforms to the SFF Specification for QSFP.
	QSFPSFPDeviceType SFPDeviceType = "QSFP"
	// The SFP conforms to the SFF Specification for QSFP14.
	QSFP14SFPDeviceType SFPDeviceType = "QSFP14"
	// The SFP conforms to the SFF Specification for QSFP28.
	QSFP28SFPDeviceType SFPDeviceType = "QSFP28"
	// The SFP conforms to the SFF Specification for QSFP56.
	QSFP56SFPDeviceType SFPDeviceType = "QSFP56"
	// The SFP conforms to the QSFP Double Density Specification.
	QSFPDDSFPDeviceType SFPDeviceType = "QSFPDD"
	// The SFP conforms to the SFF Specification for QSFP+.
	QSFPPlusSFPDeviceType SFPDeviceType = "QSFPPlus"
	// The SFP conforms to the SFF Specification for SFP.
	SFPSFPDeviceType SFPDeviceType = "SFP"
	// The SFP conforms to the SFF Specification for SFP+ and IEEE 802.3by Specification.
	SFP28SFPDeviceType SFPDeviceType = "SFP28"
	// The SFP conforms to the SFP-DD MSA Specification.
	SFPDDSFPDeviceType SFPDeviceType = "SFPDD"
	// The SFP conforms to the SFF Specification for SFP+.
	SFPPlusSFPDeviceType SFPDeviceType = "SFPPlus"
)

type SFP struct {
	// The type of fiber connection currently used by this SFP.
	FiberConnectionType FiberConnectionType
	// The manufacturer of this SFP.
	Manufacturer string
	// The medium type connected to this SFP.
	MediumType MediumType
	// The part number for this SFP.
	PartNumber string
	// The serial number for this SFP.
	SerialNumber string
	// The status and health of the resource and
	// its subordinate or dependent resources.
	Status common.Status
	// The types of SFP devices that can be attached to this port.
	SupportedSFPTypes []SFPDeviceType
	// The type of SFP device that is attached to this port.
	Type SFPDeviceType
}

// Port represents a port of a switch, controller, chassis,
// or any other device that could be connected to another entity.
type Port struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`

	// The number of active lanes for this interface.
	ActiveWidth int
	// The protocol versions capable of being sent over this port.
	CapableProtocolVersions []string
	// The protocol version being sent over this port.
	CurrentProtocolVersion string
	// The current speed of this port.
	CurrentSpeedGbps float32
	// CXL properties for this port.
	CXL CXL
	// Description provides a description of this resource.
	Description string
	// Deprecated (v1.10+): An indication of whether this port is enabled.
	Enabled bool
	// The link to the environment metrics for this port
	// or any attached small form-factor pluggable (SFP) device.
	environmentMetrics string
	// Ethernet properties for this port.
	Ethernet PortEthernet
	// Fibre Channel properties for this port.
	FibreChannel PortFibreChannel
	// An array of maximum bandwidth allocation percentages for the functions
	// associated with this port.
	FunctionMaxBandwidth []FunctionBandwidth
	// An array of minimum bandwidth allocation percentages for the functions
	// associated with this port.
	FunctionMinBandwidth []FunctionBandwidth
	// Gen-Z specific properties.
	GenZ PortGenZ
	// InfiniBand properties for this port.
	InfiniBand InfiniBand
	// An indication of whether the port is enabled.
	InterfaceEnabled bool
	// The link configuration of this port.
	LinkConfiguration []LinkConfiguration
	// The link network technology capabilities of this port.
	LinkNetworkTechnology LinkNetworkTechnology
	// The desired link state for this interface.
	LinkState LinkState
	// The link status for this interface.
	LinkStatus PortLinkStatus
	// The number of link state transitions for this interface.
	LinkTransitionIndicator int
	// The location of the port.
	Location common.Location
	// An indicator allowing an operator to physically locate this resource.
	LocationIndicatorActive bool
	// The maximum frame size supported by the port (bytes).
	MaxFrameSize int
	// The maximum speed of this port as currently configured.
	MaxSpeedGbps float32
	// The link to the metrics associated with this port.
	metrics string
	// The label of this port on the physical package for this port.
	PortID string `json:"PortId"`
	// The physical connection medium for this port.
	PortMedium PortMedium
	// The protocol being sent over this port.
	PortProtocol PortProtocol
	// The type of this port.
	PortType PortType
	// The identifier of the remote port to which this port is connected.
	RemotePortID string `json:"RemotePortId"`
	// The small form-factor pluggable (SFP) device associated with this port.
	SFP SFP
	// An indication of whether a signal is detected on this interface.
	SignalDetected bool
	// The status and health of the resource and its subordinate or dependent resources.
	Status common.Status
	// The number of lanes, phys, or other physical transport links that this port contains.
	Width int

	// An array of links to the endpoints at the other end of the link.
	associatedEndpoints      []string
	AssociatedEndpointsCount int
	// An array of links to the cables connected to this port.
	cables      []string
	CablesCount int
	// An array of links to the remote device ports at the other end of the link.
	connectedPorts      []string
	ConnectedPortsCount int
	// An array of links to the switches at the other end of the link.
	connectedSwitches      []string
	ConnectedSwitchesCount int
	// An array of links to the switch ports at the other end of the link.
	connectedSwitchPorts      []string
	ConnectedSwitchPortsCount int
	// The links to the Ethernet interfaces this port provides.
	ethernetInterfaces      []string
	EthernetInterfacesCount int
	OemLinks                json.RawMessage

	// Reset port.
	resetTarget string
	// Reset the PCI-to-PCI bridge (PPB) for this port.
	resetPPBTarget string
	OemActions     json.RawMessage

	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Port object from the raw JSON.
func (port *Port) UnmarshalJSON(b []byte) error {
	type temp Port
	type functionBandwidth struct {
		AllocationPercent     int
		NetworkDeviceFunction common.Link
	}
	type genZ struct {
		LRPT common.Link
		MRPT common.Link
		VCAT common.Link
	}
	type links struct {
		AssociatedEndpoints       common.Links
		AssociatedEndpointsCount  int `json:"AssociatedEndpoints@odata.count"`
		Cables                    common.Links
		CablesCount               int `json:"Cables@odata.count"`
		ConnectedPorts            common.Links
		ConnectedPortsCount       int `json:"ConnectedPorts@odata.count"`
		ConnectedSwitches         common.Links
		ConnectedSwitchesCount    int `json:"ConnectedSwitches@odata.count"`
		ConnectedSwitchPorts      common.Links
		ConnectedSwitchPortsCount int `json:"ConnectedSwitchPorts@odata.count"`
		EthernetInterfaces        common.Links
		EthernetInterfacesCount   int `json:"EthernetInterfaces@odata.count"`
		Oem                       json.RawMessage
	}
	type actions struct {
		ResetPort struct {
			Target string
		} `json:"#Port.Reset"`
		ResetPPB struct {
			Target string
		} `json:"#Port.ResetPPB"`
		Oem json.RawMessage // OEM actions will be stored here
	}
	var t struct {
		temp
		EnvironmentMetrics   common.Link
		FunctionMaxBandwidth []functionBandwidth
		FunctionMinBandwidth []functionBandwidth
		GenZ                 genZ
		Metrics              common.Link
		Links                links
		Actions              actions
	}

	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	// Extract the links to other entities for later
	*port = Port(t.temp)
	port.environmentMetrics = t.EnvironmentMetrics.String()
	port.GenZ.lrpt = t.GenZ.LRPT.String()
	port.GenZ.mrpt = t.GenZ.MRPT.String()
	port.GenZ.vcat = t.GenZ.VCAT.String()
	port.metrics = t.Metrics.String()
	port.associatedEndpoints = t.Links.AssociatedEndpoints.ToStrings()
	port.AssociatedEndpointsCount = t.Links.AssociatedEndpointsCount
	port.cables = t.Links.Cables.ToStrings()
	port.CablesCount = t.Links.CablesCount
	port.connectedPorts = t.Links.ConnectedPorts.ToStrings()
	port.ConnectedPortsCount = t.Links.ConnectedPortsCount
	port.connectedSwitches = t.Links.ConnectedSwitches.ToStrings()
	port.ConnectedSwitchesCount = t.Links.ConnectedSwitchesCount
	port.connectedSwitchPorts = t.Links.ConnectedSwitchPorts.ToStrings()
	port.ConnectedSwitchPortsCount = t.Links.ConnectedPortsCount
	port.ethernetInterfaces = t.Links.EthernetInterfaces.ToStrings()
	port.EthernetInterfacesCount = t.Links.EthernetInterfacesCount
	port.OemLinks = t.Links.Oem
	port.resetTarget = t.Actions.ResetPort.Target
	port.resetPPBTarget = t.Actions.ResetPPB.Target
	port.OemActions = t.Actions.Oem

	port.FunctionMaxBandwidth = make([]FunctionBandwidth, len(t.FunctionMaxBandwidth))
	for i, fb := range t.FunctionMaxBandwidth {
		port.FunctionMaxBandwidth[i].AllocationPercent = fb.AllocationPercent
		port.FunctionMaxBandwidth[i].networkDeviceFunction = fb.NetworkDeviceFunction.String()
	}

	port.FunctionMinBandwidth = make([]FunctionBandwidth, len(t.FunctionMinBandwidth))
	for i, fb := range t.FunctionMinBandwidth {
		port.FunctionMinBandwidth[i].AllocationPercent = fb.AllocationPercent
		port.FunctionMinBandwidth[i].networkDeviceFunction = fb.NetworkDeviceFunction.String()
	}

	return nil
}

// Update commits updates to this object's properties to the running system.
func (port *Port) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Port)
	err := original.UnmarshalJSON(port.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"BackpressureSampleInterval",
		"CompletionCollectionInterval",
		"CongestionTelemetryEnabled",
		"EgressModeratePercentage",
		"EgressSeverePercentage",
		"MaxSustainedRequestCmpBias",
		"TemporaryThroughputReductionEnabled",
		"Enabled",
		"EEEEnabled",
		"FlowControlConfiguration",
		"LLDPEnabled",
		"ChassisId",
		"ChassisIdSubtype",
		"ManagementAddressIPv4",
		"ManagementAddressIPv6",
		"ManagementAddressMAC",
		"ManagementVlanId",
		"PortId",
		"PortIdSubtype",
		"SystemCapabilities",
		"SystemDescription",
		"SystemName",
		"WakeOnLANEnabled",
		"AllocationPercent",
		"InterfaceEnabled",
		"AutoSpeedNegotiationEnabled",
		"ConfiguredLinkSpeedGbps",
		"ConfiguredWidth",
		"LinkState",
		"LinkTransitionIndicator",
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(port).Elem()

	return port.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetPort will get a Port instance from the service.
func GetPort(c common.Client, uri string) (*Port, error) {
	var port Port
	return &port, port.Get(c, uri, &port)
}

// ListReferencedPorts gets the collection of Port from
// a provided reference.
func ListReferencedPorts(c common.Client, link string) ([]*Port, error) { //nolint:dupl
	var result []*Port
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *Port
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		port, err := GetPort(c, link)
		ch <- GetResult{Item: port, Link: link, Error: err}
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

func (port *Port) ResetPort(resetType ResetType) error {
	if port.resetTarget == "" {
		return fmt.Errorf("ResetPort action is not supported")
	}

	t := struct {
		ResetType ResetType
	}{
		ResetType: resetType,
	}
	return port.Post(port.resetTarget, t)
}

func (port *Port) ResetPPB() error {
	if port.resetPPBTarget == "" {
		return fmt.Errorf("ResetPPB action is not supported")
	}

	return port.Post(port.resetPPBTarget, nil)
}
