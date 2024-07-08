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
	// DisconnectedConnectedDeviceMode shall indicate the connection is not CXL or is disconnected.
	DisconnectedConnectedDeviceMode ConnectedDeviceMode = "Disconnected"
	// RCDConnectedDeviceMode shall indicate the connected device mode is restricted CXL device (RCD).
	RCDConnectedDeviceMode ConnectedDeviceMode = "RCD"
	// CXL68BFlitAndVHConnectedDeviceMode shall indicate the connected device mode is CXL 68B flit and VH.
	CXL68BFlitAndVHConnectedDeviceMode ConnectedDeviceMode = "CXL68BFlitAndVH"
	// Standard256BFlitConnectedDeviceMode shall indicate the connected device mode is standard 256B flit.
	Standard256BFlitConnectedDeviceMode ConnectedDeviceMode = "Standard256BFlit"
	// CXLLatencyOptimized256BFlitConnectedDeviceMode shall indicate the connected device mode is CXL latency-optimized
	// 256B flit.
	CXLLatencyOptimized256BFlitConnectedDeviceMode ConnectedDeviceMode = "CXLLatencyOptimized256BFlit"
	// PBRConnectedDeviceMode shall indicate the connected device mode is port-based routing (PBR).
	PBRConnectedDeviceMode ConnectedDeviceMode = "PBR"
)

type ConnectedDeviceType string

const (
	// NoneConnectedDeviceType shall indicate no device is detected.
	NoneConnectedDeviceType ConnectedDeviceType = "None"
	// PCIeDeviceConnectedDeviceType shall indicate the connected device is a PCIe device.
	PCIeDeviceConnectedDeviceType ConnectedDeviceType = "PCIeDevice"
	// Type1ConnectedDeviceType shall indicate the connected device is a CXL Type 1 device.
	Type1ConnectedDeviceType ConnectedDeviceType = "Type1"
	// Type2ConnectedDeviceType shall indicate the connected device is a CXL Type 2 device.
	Type2ConnectedDeviceType ConnectedDeviceType = "Type2"
	// Type3SLDConnectedDeviceType shall indicate the connected device is a CXL Type 3 single logical device (SLD).
	Type3SLDConnectedDeviceType ConnectedDeviceType = "Type3SLD"
	// Type3MLDConnectedDeviceType shall indicate the connected device is a CXL Type 3 multi-logical device (MLD).
	Type3MLDConnectedDeviceType ConnectedDeviceType = "Type3MLD"
)

type CurrentPortConfigurationState string

const (
	// DisabledCurrentPortConfigurationState shall indicate the port is disabled.
	DisabledCurrentPortConfigurationState CurrentPortConfigurationState = "Disabled"
	// BindInProgressCurrentPortConfigurationState shall indicate a bind is in progress for the port.
	BindInProgressCurrentPortConfigurationState CurrentPortConfigurationState = "BindInProgress"
	// UnbindInProgressCurrentPortConfigurationState shall indicate an unbind is in progress for the port.
	UnbindInProgressCurrentPortConfigurationState CurrentPortConfigurationState = "UnbindInProgress"
	// DSPCurrentPortConfigurationState shall indicate the port is enabled as a downstream port (DSP).
	DSPCurrentPortConfigurationState CurrentPortConfigurationState = "DSP"
	// USPCurrentPortConfigurationState shall indicate the port is enabled as an upstream port (USP).
	USPCurrentPortConfigurationState CurrentPortConfigurationState = "USP"
	// ReservedCurrentPortConfigurationState shall indicate the port is in a reserved state.
	ReservedCurrentPortConfigurationState CurrentPortConfigurationState = "Reserved"
	// FabricLinkCurrentPortConfigurationState shall indicate the port is enabled as a fabric link to another switch.
	FabricLinkCurrentPortConfigurationState CurrentPortConfigurationState = "FabricLink"
)

type FiberConnectionType string

const (
	// SingleModeFiberConnectionType The connection is using single mode operation.
	SingleModeFiberConnectionType FiberConnectionType = "SingleMode"
	// MultiModeFiberConnectionType The connection is using multi mode operation.
	MultiModeFiberConnectionType FiberConnectionType = "MultiMode"
)

type IEEE802IDSubtype string

const (
	// ChassisCompIEEE802IDSubtype Chassis component, based on the value of entPhysicalAlias in RFC4133.
	ChassisCompIEEE802IDSubtype IEEE802IDSubtype = "ChassisComp"
	// IfAliasIEEE802IDSubtype Interface alias, based on the ifAlias MIB object.
	IfAliasIEEE802IDSubtype IEEE802IDSubtype = "IfAlias"
	// PortCompIEEE802IDSubtype Port component, based on the value of entPhysicalAlias in RFC4133.
	PortCompIEEE802IDSubtype IEEE802IDSubtype = "PortComp"
	// MacAddrIEEE802IDSubtype MAC address, based on an agent-detected unicast source address as defined in IEEE
	// standard 802.
	MacAddrIEEE802IDSubtype IEEE802IDSubtype = "MacAddr"
	// NetworkAddrIEEE802IDSubtype Network address, based on an agent-detected network address.
	NetworkAddrIEEE802IDSubtype IEEE802IDSubtype = "NetworkAddr"
	// IfNameIEEE802IDSubtype Interface name, based on the ifName MIB object.
	IfNameIEEE802IDSubtype IEEE802IDSubtype = "IfName"
	// AgentIDIEEE802IDSubtype Agent circuit ID, based on the agent-local identifier of the circuit as defined in
	// RFC3046.
	AgentIDIEEE802IDSubtype IEEE802IDSubtype = "AgentId"
	// LocalAssignIEEE802IDSubtype Locally assigned, based on an alphanumeric value locally assigned.
	LocalAssignIEEE802IDSubtype IEEE802IDSubtype = "LocalAssign"
	// NotTransmittedIEEE802IDSubtype No data to be sent to/received from remote partner.
	NotTransmittedIEEE802IDSubtype IEEE802IDSubtype = "NotTransmitted"
)

type LLDPSystemCapabilities string

const (
	// NoneLLDPSystemCapabilities shall indicate the system capabilities are transmitted, but no capabilities are set.
	NoneLLDPSystemCapabilities LLDPSystemCapabilities = "None"
	// BridgeLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined 'bridge' capability.
	BridgeLLDPSystemCapabilities LLDPSystemCapabilities = "Bridge"
	// DOCSISCableDeviceLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined 'DOCSIS cable device' capability.
	DOCSISCableDeviceLLDPSystemCapabilities LLDPSystemCapabilities = "DOCSISCableDevice"
	// OtherLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined 'other' capability.
	OtherLLDPSystemCapabilities LLDPSystemCapabilities = "Other"
	// RepeaterLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined 'repeater' capability.
	RepeaterLLDPSystemCapabilities LLDPSystemCapabilities = "Repeater"
	// RouterLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined 'router' capability.
	RouterLLDPSystemCapabilities LLDPSystemCapabilities = "Router"
	// StationLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined 'station' capability.
	StationLLDPSystemCapabilities LLDPSystemCapabilities = "Station"
	// TelephoneLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined 'telephone' capability.
	TelephoneLLDPSystemCapabilities LLDPSystemCapabilities = "Telephone"
	// WLANAccessPointLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined 'WLAN access point' capability.
	WLANAccessPointLLDPSystemCapabilities LLDPSystemCapabilities = "WLANAccessPoint"
)

type PortLinkStatus string

const (
	// LinkUpPortLinkStatus This link on this interface is up.
	LinkUpPortLinkStatus PortLinkStatus = "LinkUp"
	// StartingPortLinkStatus This link on this interface is starting. A physical link has been established, but the port
	// is not able to transfer data.
	StartingPortLinkStatus PortLinkStatus = "Starting"
	// TrainingPortLinkStatus This physical link on this interface is training.
	TrainingPortLinkStatus PortLinkStatus = "Training"
	// LinkDownPortLinkStatus The link on this interface is down.
	LinkDownPortLinkStatus PortLinkStatus = "LinkDown"
	// NoLinkPortLinkStatus No physical link detected on this interface.
	NoLinkPortLinkStatus PortLinkStatus = "NoLink"
)

type MediumType string

const (
	// CopperMediumType The medium connected is copper.
	CopperMediumType MediumType = "Copper"
	// FiberOpticMediumType The medium connected is fiber optic.
	FiberOpticMediumType MediumType = "FiberOptic"
)

type PortMedium string

const (
	// ElectricalPortMedium This port has an electrical cable connection.
	ElectricalPortMedium PortMedium = "Electrical"
	// OpticalPortMedium This port has an optical cable connection.
	OpticalPortMedium PortMedium = "Optical"
)

// PortType is
type PortType string

const (
	// UpstreamPortPortType This port connects to a host device.
	UpstreamPortPortType PortType = "UpstreamPort"
	// DownstreamPortPortType This port connects to a target device.
	DownstreamPortPortType PortType = "DownstreamPort"
	// InterswitchPortPortType This port connects to another switch.
	InterswitchPortPortType PortType = "InterswitchPort"
	// ManagementPortPortType This port connects to a switch manager.
	ManagementPortPortType PortType = "ManagementPort"
	// BidirectionalPortPortType This port connects to any type of device.
	BidirectionalPortPortType PortType = "BidirectionalPort"
	// UnconfiguredPortPortType This port has not yet been configured.
	UnconfiguredPortPortType PortType = "UnconfiguredPort"
)

type SFPType string

const (
	// SFPSFPType The SFP conforms to the SFF Specification for SFP.
	SFPSFPType SFPType = "SFP"
	// SFPPlusSFPType The SFP conforms to the SFF Specification for SFP+.
	SFPPlusSFPType SFPType = "SFPPlus"
	// SFP28SFPType The SFP conforms to the SFF Specification for SFP+ and IEEE 802.3by Specification.
	SFP28SFPType SFPType = "SFP28"
	// CSFPSFPType The SFP conforms to the CSFP MSA Specification.
	CSFPSFPType SFPType = "cSFP"
	// SFPDDSFPType The SFP conforms to the SFP-DD MSA Specification.
	SFPDDSFPType SFPType = "SFPDD"
	// QSFPSFPType The SFP conforms to the SFF Specification for QSFP.
	QSFPSFPType SFPType = "QSFP"
	// QSFPPlusSFPType The SFP conforms to the SFF Specification for QSFP+.
	QSFPPlusSFPType SFPType = "QSFPPlus"
	// QSFP14SFPType The SFP conforms to the SFF Specification for QSFP14.
	QSFP14SFPType SFPType = "QSFP14"
	// QSFP28SFPType The SFP conforms to the SFF Specification for QSFP28.
	QSFP28SFPType SFPType = "QSFP28"
	// QSFP56SFPType The SFP conforms to the SFF Specification for QSFP56.
	QSFP56SFPType SFPType = "QSFP56"
	// MiniSASHDSFPType The SFP conforms to the SFF Specification SFF-8644.
	MiniSASHDSFPType SFPType = "MiniSASHD"
	// QSFPDDSFPType The SFP conforms to the QSFP Double Density Specification.
	QSFPDDSFPType SFPType = "QSFPDD"
	// OSFPSFPType The SFP conforms to the OSFP Specification.
	OSFPSFPType SFPType = "OSFP"
)

type QoSTelemetryCapabilities struct {
	// Indicates whether the port supports the CXL Specification-defined 'Egress Port Backpressure' mechanism.
	EgressPortBackpressureSupported bool
	// Indicates whether the port supports the CXL Specification-defined 'Temporary Throughput Reduction' mechanism.
	TemporaryThroughputReductionSupported bool
}

type CXLPort struct {
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

// ConfiguredNetworkLink shall contain a set of link settings that a port is configured to use for autonegotiation.
type ConfiguredNetworkLink struct {
	// ConfiguredLinkSpeedGbps shall contain the network link speed per lane this port is configured to allow for
	// autonegotiation purposes. This value includes overhead associated with the protocol.
	ConfiguredLinkSpeedGbps float64
	// ConfiguredWidth shall contain the network link width this port is configured to use for autonegotiation
	// purposes.
	ConfiguredWidth int
}

// EthernetProperties shall contain Ethernet-specific properties for a port.
type EthernetProperties struct {
	// AssociatedMACAddresses shall contain an array of configured MAC addresses that are associated with this network
	// port, including the programmed address of the lowest-numbered network device function, the configured but not
	// active address if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedMACAddresses []string
	// EEEEnabled shall indicate whether IEEE 802.3az Energy-Efficient Ethernet (EEE) is enabled on this port.
	EEEEnabled bool
	// FlowControlConfiguration shall contain the locally configured 802.3x flow control setting for this port.
	FlowControlConfiguration FlowControl
	// FlowControlStatus shall contain the 802.3x flow control behavior negotiated with the link partner for this port.
	FlowControlStatus FlowControl
	// LLDPEnabled shall contain the state indicating whether to enable LLDP for a port. If LLDP is disabled at the
	// adapter level, this property shall be ignored.
	LLDPEnabled bool
	// LLDPReceive shall contain the LLDP data being received on this link.
	LLDPReceive LLDPReceive
	// LLDPTransmit shall contain the LLDP data being transmitted on this link.
	LLDPTransmit LLDPTransmit
	// WakeOnLANEnabled shall indicate whether Wake on LAN (WoL) is enabled on this port.
	WakeOnLANEnabled bool
}

// FibreChannelProperties shall contain Fibre Channel-specific properties for a port.
type FibreChannelProperties struct {
	// AssociatedWorldWideNames shall contain an array of configured World Wide Names (WWN) that are associated with
	// this network port, including the programmed address of the lowest-numbered network device function, the
	// configured but not active address if applicable, the address for hardware port teaming, or other network
	// addresses.
	AssociatedWorldWideNames []string
	// FabricName shall indicate the Fibre Channel Fabric Name provided by the switch.
	FabricName string
	// NumberDiscoveredRemotePorts shall contain the number of ports not on this associated device that this port has
	// discovered.
	NumberDiscoveredRemotePorts int
	// PortConnectionType shall contain the connection type for this port.
	PortConnectionType PortConnectionType
}

type FunctionMaxBandwidth struct {
	// AllocationPercent shall contain the maximum bandwidth percentage allocation, '0' to '100', for the associated
	// network device function.
	AllocationPercent int
	// NetworkDeviceFunction shall contain a link to a resource of type NetworkDeviceFunction that represents the
	// network device function associated with this bandwidth setting of this network port.
	networkDeviceFunction common.Link //nolint:unused
}

// FunctionMinBandwidth shall describe a minimum bandwidth percentage allocation for a network device function
// associated with a port.
type FunctionMinBandwidth struct {
	// AllocationPercent shall contain the minimum bandwidth percentage allocation, '0' to '100', for the associated
	// network device function. The sum of all minimum percentages shall not exceed '100'.
	AllocationPercent int
	// NetworkDeviceFunction shall contain a link to a resource of type NetworkDeviceFunction that represents the
	// network device function associated with this bandwidth setting of this network port.
	networkDeviceFunction common.Link //nolint:unused
}

// InfiniBandProperties shall contain InfiniBand-specific properties for a port.
type InfiniBandProperties struct {
	// AssociatedNodeGUIDs shall contain an array of configured node GUIDs that are associated with this network port,
	// including the programmed address of the lowest-numbered network device function, the configured but not active
	// address if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedNodeGUIDs []string
	// AssociatedPortGUIDs shall contain an array of configured port GUIDs that are associated with this network port,
	// including the programmed address of the lowest-numbered network device function, the configured but not active
	// address if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedPortGUIDs []string
	// AssociatedSystemGUIDs shall contain an array of configured system GUIDs that are associated with this network
	// port, including the programmed address of the lowest-numbered network device function, the configured but not
	// active address if applicable, the address for hardware port teaming, or other network addresses.
	AssociatedSystemGUIDs []string
}

// LLDPReceive shall contain the LLDP data from the remote partner across this link.
type LLDPReceive struct {
	// ChassisID shall contain the chassis ID received from the remote partner across this link. If no such chassis ID
	// has been received, this property should not be present.
	ChassisID string
	// ChassisIDSubtype shall contain the IEEE 802.1AB-2009 chassis ID subtype received from the remote partner across
	// this link. If no such chassis ID subtype has been received, this property should not be present.
	ChassisIDSubtype IEEE802IDSubtype
	// ManagementAddressIPv4 shall contain the IPv4 management address received from the remote partner across this
	// link. If no such management address has been received, this property should not be present.
	ManagementAddressIPv4 string
	// ManagementAddressIPv6 shall contain the IPv6 management address received from the remote partner across this
	// link. If no such management address has been received, this property should not be present.
	ManagementAddressIPv6 string
	// ManagementAddressMAC shall contain the management MAC address received from the remote partner across this link.
	// If no such management address has been received, this property should not be present.
	ManagementAddressMAC string
	// ManagementVlanID shall contain the management VLAN ID received from the remote partner across this link. If no
	// such management VLAN ID has been received, this property should not be present.
	ManagementVlanID int
	// PortID shall contain a colon-delimited string of hexadecimal octets identifying the port received from the
	// remote partner across this link. If no such port ID has been received, this property should not be present.
	PortID string
	// PortIDSubtype shall contain the port ID subtype from IEEE 802.1AB-2009 Table 8-3 received from the remote
	// partner across this link. If no such port ID subtype has been received, this property should not be present.
	PortIDSubtype IEEE802IDSubtype
	// SystemCapabilities shall contain the system capabilities received from the remote partner across this link. If
	// no such system capabilities have been received, this property shall not be present. This property shall not
	// contain the value 'None'.
	SystemCapabilities []LLDPSystemCapabilities
	// SystemDescription shall contain the system description received from the remote partner across this link. If no
	// such system description has been received, this property shall not be present.
	SystemDescription string
	// SystemName shall contain the system name received from the remote partner across this link. If no such system
	// name has been received, this property shall not be present.
	SystemName string
}

// LLDPTransmit shall contain the LLDP data to be transmitted from this endpoint.
type LLDPTransmit struct {
	// ChassisID shall contain the chassis ID to be transmitted from this endpoint. If no such chassis ID is to be
	// transmitted, this value shall be an empty string.
	ChassisID string
	// ChassisIDSubtype shall contain the IEEE 802.1AB-2009 chassis ID subtype to be transmitted from this endpoint. If
	// no such chassis ID subtype is to be transmitted, this value shall be 'NotTransmitted'.
	ChassisIDSubtype IEEE802IDSubtype
	// ManagementAddressIPv4 shall contain the IPv4 management address to be transmitted from this endpoint. If no such
	// management address is to be transmitted, this value shall be an empty string.
	ManagementAddressIPv4 string
	// ManagementAddressIPv6 shall contain the IPv6 management address to be transmitted from this endpoint. If no such
	// management address is to be transmitted, this value shall be an empty string.
	ManagementAddressIPv6 string
	// ManagementAddressMAC shall contain the management MAC address to be transmitted from this endpoint. If no such
	// management address is to be transmitted, this value shall be an empty string.
	ManagementAddressMAC string
	// ManagementVlanID shall contain the management VLAN ID to be transmitted from this endpoint. If no such port ID
	// is to be transmitted, this value shall be '4095'.
	ManagementVlanID int
	// PortID shall contain a colon-delimited string of hexadecimal octets identifying the port for an LLDP endpoint.
	// If no such port ID is to be transmitted, this value shall be an empty string.
	PortID string
	// PortIDSubtype shall contain the port ID subtype from IEEE 802.1AB-2009 Table 8-3 to be transmitted from this
	// endpoint. If no such port ID subtype is to be transmitted, this value shall be 'NotTransmitted'.
	PortIDSubtype IEEE802IDSubtype
	// SystemCapabilities shall contain the system capabilities to be transmitted from this endpoint. If no such system
	// capabilities are to be transmitted, this value shall be an empty array. If this property contains the value
	// 'None', an empty set of system capabilities is transmitted from this endpoint.
	SystemCapabilities []LLDPSystemCapabilities
	// SystemDescription shall contain the system description to be transmitted from this endpoint. If no such system
	// description is to be transmitted, this value shall be an empty string.
	SystemDescription string
	// SystemName shall contain the system name to be transmitted from this endpoint. If no such system name is to be
	// transmitted, this value shall be an empty string.
	SystemName string
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
	// Locally assigned, based on an alphanumeric value locally assigned.
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
	LLDPReceive LLDPReceive
	// LLDP data being transmitted on this link.
	LLDPTransmit LLDPTransmit
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
	CXL CXLPort
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
	// The GenZ Linear Packet Relay Table for the port.
	genZLPRT []string
	// The Multi-subnet Packet Relay Table for the port.
	genZMPRT []string
	// The Virtual Channel Action Table for the port.
	genZVCAT []string
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

	associatedEndpoints []string
	// AssociatedEndpointsCount gets the number of endpoints on the other end of the link.
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
		LRPT common.LinksCollection
		MRPT common.LinksCollection
		VCAT common.LinksCollection
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
		ResetPort common.ActionTarget `json:"#Port.Reset"`
		ResetPPB  common.ActionTarget `json:"#Port.ResetPPB"`
		Oem       json.RawMessage     // OEM actions will be stored here
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
	port.metrics = t.Metrics.String()
	port.genZLPRT = t.GenZ.LRPT.ToStrings()
	port.genZMPRT = t.GenZ.MRPT.ToStrings()
	port.genZVCAT = t.GenZ.VCAT.ToStrings()

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

// EnvironmentMetrics gets the environment metrics for this port or any attached small form-factor pluggable (SFP) device.
func (port *Port) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if port.environmentMetrics == "" {
		return nil, nil
	}
	return GetEnvironmentMetrics(port.GetClient(), port.environmentMetrics)
}

// Metrics gets the metrics for this port.
func (port *Port) Metrics() (*PortMetrics, error) {
	if port.metrics == "" {
		return nil, nil
	}
	return GetPortMetrics(port.GetClient(), port.metrics)
}

// AssociatedEndpoints gets the endpoints at the other end of the link.
func (port *Port) AssociatedEndpoints() ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](port.GetClient(), port.associatedEndpoints)
}

// Cables gets the cables connected to this port.
func (port *Port) Cables() ([]*Cable, error) {
	return common.GetObjects[Cable](port.GetClient(), port.cables)
}

// ConnectedPorts gets the remote device ports connected to the other end of the link.
func (port *Port) ConnectedPorts() ([]*Port, error) {
	return common.GetObjects[Port](port.GetClient(), port.connectedPorts)
}

// ConnectedSwitchPorts gets the switch ports connected to the other end of the link.
func (port *Port) ConnectedSwitchPorts() ([]*Port, error) {
	return common.GetObjects[Port](port.GetClient(), port.connectedSwitchPorts)
}

// ConnectedSwitches gets the switches connected to the other end of the link.
func (port *Port) ConnectedSwitches() ([]*Switch, error) {
	return common.GetObjects[Switch](port.GetClient(), port.connectedSwitches)
}

// EthernetInterfaces gets the Ethernet interfaces this port provides.
func (port *Port) EthernetInterfaces() ([]*EthernetInterface, error) {
	return common.GetObjects[EthernetInterface](port.GetClient(), port.ethernetInterfaces)
}

// GenZLPRT gets the Gen-Z Core Specification-defined Linear Packet Relay Table for this port.
func (port *Port) GenZLPRT() ([]*RouteEntry, error) {
	return common.GetObjects[RouteEntry](port.GetClient(), port.genZLPRT)
}

// GenZMPRT gets the Gen-Z Core Specification-defined Multi-subnet Packet Relay Table for this port.
func (port *Port) GenZMPRT() ([]*RouteEntry, error) {
	return common.GetObjects[RouteEntry](port.GetClient(), port.genZMPRT)
}

// GenZVCAT gets the Gen-Z Virtual Channel Action Table for the port.
func (port *Port) GenZVCAT() ([]*VCATEntry, error) {
	return common.GetObjects[VCATEntry](port.GetClient(), port.genZVCAT)
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
		"ChassisIDSubtype",
		"ManagementAddressIPv4",
		"ManagementAddressIPv6",
		"ManagementAddressMAC",
		"ManagementVlanId",
		"PortId",
		"PortIDSubtype",
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
	return common.GetObject[Port](c, uri)
}

// ListReferencedPorts gets the collection of Port from
// a provided reference.
func ListReferencedPorts(c common.Client, link string) ([]*Port, error) {
	return common.GetCollectionObjects[Port](c, link)
}

// ResetPort resets this port.
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

// ResetPPB resets the PCI-to-PCI bridge (PPB) for this port.
func (port *Port) ResetPPB() error {
	if port.resetPPBTarget == "" {
		return fmt.Errorf("ResetPPB action is not supported")
	}

	return port.Post(port.resetPPBTarget, nil)
}
