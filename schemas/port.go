//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/Port.v1_18_0.json
// 2025.4 - #Port.v1_18_0.Port

package schemas

import (
	"encoding/json"
)

type ConnectedDeviceMode string

const (
	// DisconnectedConnectedDeviceMode shall indicate the connection is not CXL or
	// is disconnected.
	DisconnectedConnectedDeviceMode ConnectedDeviceMode = "Disconnected"
	// RCDConnectedDeviceMode shall indicate the connected device mode is
	// restricted CXL device (RCD).
	RCDConnectedDeviceMode ConnectedDeviceMode = "RCD"
	// CXL68BFlitAndVHConnectedDeviceMode shall indicate the connected device mode
	// is CXL 68B flit and VH.
	CXL68BFlitAndVHConnectedDeviceMode ConnectedDeviceMode = "CXL68BFlitAndVH"
	// Standard256BFlitConnectedDeviceMode shall indicate the connected device mode
	// is standard 256B flit.
	Standard256BFlitConnectedDeviceMode ConnectedDeviceMode = "Standard256BFlit"
	// CXLLatencyOptimized256BFlitConnectedDeviceMode shall indicate the connected
	// device mode is CXL latency-optimized 256B flit.
	CXLLatencyOptimized256BFlitConnectedDeviceMode ConnectedDeviceMode = "CXLLatencyOptimized256BFlit"
	// PBRConnectedDeviceMode shall indicate the connected device mode is
	// port-based routing (PBR).
	PBRConnectedDeviceMode ConnectedDeviceMode = "PBR"
)

type ConnectedDeviceType string

const (
	// NoneConnectedDeviceType shall indicate no device is detected.
	NoneConnectedDeviceType ConnectedDeviceType = "None"
	// PCIeDeviceConnectedDeviceType shall indicate the connected device is a PCIe
	// device.
	PCIeDeviceConnectedDeviceType ConnectedDeviceType = "PCIeDevice"
	// Type1ConnectedDeviceType shall indicate the connected device is a CXL Type 1
	// device.
	Type1ConnectedDeviceType ConnectedDeviceType = "Type1"
	// Type2ConnectedDeviceType shall indicate the connected device is a CXL Type 2
	// device.
	Type2ConnectedDeviceType ConnectedDeviceType = "Type2"
	// Type3SLDConnectedDeviceType shall indicate the connected device is a CXL
	// Type 3 single logical device (SLD).
	Type3SLDConnectedDeviceType ConnectedDeviceType = "Type3SLD"
	// Type3MLDConnectedDeviceType shall indicate the connected device is a CXL
	// Type 3 multi-logical device (MLD).
	Type3MLDConnectedDeviceType ConnectedDeviceType = "Type3MLD"
	// PBRComponentConnectedDeviceType shall indicate the connected device is a
	// port-based routing (PBR) component.
	PBRComponentConnectedDeviceType ConnectedDeviceType = "PBRComponent"
)

type CurrentPortConfigurationState string

const (
	// DisabledCurrentPortConfigurationState shall indicate the port is disabled.
	DisabledCurrentPortConfigurationState CurrentPortConfigurationState = "Disabled"
	// BindInProgressCurrentPortConfigurationState shall indicate a bind is in
	// progress for the port.
	BindInProgressCurrentPortConfigurationState CurrentPortConfigurationState = "BindInProgress"
	// UnbindInProgressCurrentPortConfigurationState shall indicate an unbind is in
	// progress for the port.
	UnbindInProgressCurrentPortConfigurationState CurrentPortConfigurationState = "UnbindInProgress"
	// DSPCurrentPortConfigurationState shall indicate the port is enabled as a
	// downstream port (DSP).
	DSPCurrentPortConfigurationState CurrentPortConfigurationState = "DSP"
	// USPCurrentPortConfigurationState shall indicate the port is enabled as an
	// upstream port (USP).
	USPCurrentPortConfigurationState CurrentPortConfigurationState = "USP"
	// ReservedCurrentPortConfigurationState shall indicate the port is in a
	// reserved state.
	ReservedCurrentPortConfigurationState CurrentPortConfigurationState = "Reserved"
	// FabricLinkCurrentPortConfigurationState shall indicate the port is enabled
	// as a fabric link to another switch. This value is the same as 'Fabric Port'
	// as described in the CXL Specification v3.1 and later. Previous versions of
	// the CXL Specification referred to this value as a 'Fabric Link'.
	FabricLinkCurrentPortConfigurationState CurrentPortConfigurationState = "FabricLink"
)

type FiberConnectionType string

const (
	// SingleModeFiberConnectionType The connection is using single mode operation.
	SingleModeFiberConnectionType FiberConnectionType = "SingleMode"
	// MultiModeFiberConnectionType The connection is using multi mode operation.
	MultiModeFiberConnectionType FiberConnectionType = "MultiMode"
)

type HostDeviceType string

const (
	// NoneHostDeviceType shall indicate the port is not connected to any host
	// device.
	NoneHostDeviceType HostDeviceType = "None"
	// SystemHostDeviceType shall indicate the port is connected to a computer
	// system device.
	SystemHostDeviceType HostDeviceType = "System"
	// ManagerHostDeviceType shall indicate the port is connected to a manager
	// device.
	ManagerHostDeviceType HostDeviceType = "Manager"
)

type IEEE802IDSubtype string

const (
	// ChassisCompIEEE802IDSubtype Chassis component, based on the value of
	// entPhysicalAlias in RFC4133.
	ChassisCompIEEE802IDSubtype IEEE802IDSubtype = "ChassisComp"
	// IfAliasIEEE802IDSubtype Interface alias, based on the ifAlias MIB object.
	IfAliasIEEE802IDSubtype IEEE802IDSubtype = "IfAlias"
	// PortCompIEEE802IDSubtype Port component, based on the value of
	// entPhysicalAlias in RFC4133.
	PortCompIEEE802IDSubtype IEEE802IDSubtype = "PortComp"
	// MacAddrIEEE802IDSubtype MAC address, based on an agent-detected unicast
	// source address as defined in IEEE standard 802.
	MacAddrIEEE802IDSubtype IEEE802IDSubtype = "MacAddr"
	// NetworkAddrIEEE802IDSubtype Network address, based on an agent-detected
	// network address.
	NetworkAddrIEEE802IDSubtype IEEE802IDSubtype = "NetworkAddr"
	// IfNameIEEE802IDSubtype Interface name, based on the ifName MIB object.
	IfNameIEEE802IDSubtype IEEE802IDSubtype = "IfName"
	// AgentIDIEEE802IDSubtype Agent circuit ID, based on the agent-local
	// identifier of the circuit as defined in RFC3046.
	AgentIDIEEE802IDSubtype IEEE802IDSubtype = "AgentId"
	// LocalAssignIEEE802IDSubtype Locally assigned, based on an alphanumeric value
	// locally assigned.
	LocalAssignIEEE802IDSubtype IEEE802IDSubtype = "LocalAssign"
	// NotTransmittedIEEE802IDSubtype No data to be sent to/received from remote
	// partner.
	NotTransmittedIEEE802IDSubtype IEEE802IDSubtype = "NotTransmitted"
)

type LLDPSystemCapabilities string

const (
	// NoneLLDPSystemCapabilities shall indicate the system capabilities are
	// transmitted, but no capabilities are set.
	NoneLLDPSystemCapabilities LLDPSystemCapabilities = "None"
	// BridgeLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined 'bridge'
	// capability.
	BridgeLLDPSystemCapabilities LLDPSystemCapabilities = "Bridge"
	// DOCSISCableDeviceLLDPSystemCapabilities shall indicate the
	// IEEE802.1AB-defined 'DOCSIS cable device' capability.
	DOCSISCableDeviceLLDPSystemCapabilities LLDPSystemCapabilities = "DOCSISCableDevice"
	// OtherLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined 'other'
	// capability.
	OtherLLDPSystemCapabilities LLDPSystemCapabilities = "Other"
	// RepeaterLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined
	// 'repeater' capability.
	RepeaterLLDPSystemCapabilities LLDPSystemCapabilities = "Repeater"
	// RouterLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined 'router'
	// capability.
	RouterLLDPSystemCapabilities LLDPSystemCapabilities = "Router"
	// StationLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined
	// 'station' capability.
	StationLLDPSystemCapabilities LLDPSystemCapabilities = "Station"
	// TelephoneLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined
	// 'telephone' capability.
	TelephoneLLDPSystemCapabilities LLDPSystemCapabilities = "Telephone"
	// WLANAccessPointLLDPSystemCapabilities shall indicate the IEEE802.1AB-defined
	// 'WLAN access point' capability.
	WLANAccessPointLLDPSystemCapabilities LLDPSystemCapabilities = "WLANAccessPoint"
)

type LinkNetworkTechnology string

const (
	// EthernetLinkNetworkTechnology The port is capable of connecting to an
	// Ethernet network.
	EthernetLinkNetworkTechnology LinkNetworkTechnology = "Ethernet"
	// InfiniBandLinkNetworkTechnology The port is capable of connecting to an
	// InfiniBand network.
	InfiniBandLinkNetworkTechnology LinkNetworkTechnology = "InfiniBand"
	// FibreChannelLinkNetworkTechnology The port is capable of connecting to a
	// Fibre Channel network.
	FibreChannelLinkNetworkTechnology LinkNetworkTechnology = "FibreChannel"
	// GenZLinkNetworkTechnology The port is capable of connecting to a Gen-Z
	// fabric.
	GenZLinkNetworkTechnology LinkNetworkTechnology = "GenZ"
	// PCIeLinkNetworkTechnology The port is capable of connecting to PCIe and CXL
	// fabrics.
	PCIeLinkNetworkTechnology LinkNetworkTechnology = "PCIe"
)

type LinkState string

const (
	// EnabledLinkState shall indicate the link is enabled and operational. The
	// port is allowed to establish a connection with the remote port.
	EnabledLinkState LinkState = "Enabled"
	// DisabledLinkState shall indicate the link is disabled and not operational.
	// The port is not allowed to establish a connection with the remote port.
	// However, other types of traffic, such as management traffic, may be sent or
	// received by the port.
	DisabledLinkState LinkState = "Disabled"
)

type PortLinkStatus string

const (
	// LinkUpPortLinkStatus This link on this interface is up.
	LinkUpPortLinkStatus PortLinkStatus = "LinkUp"
	// StartingPortLinkStatus This link on this interface is starting. A physical link
	// has been established, but the port is not able to transfer data.
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

type PortConnectionType string

const (
	// NotConnectedPortConnectionType This port is not connected.
	NotConnectedPortConnectionType PortConnectionType = "NotConnected"
	// NPortPortConnectionType This port connects through an N-port to a switch.
	NPortPortConnectionType PortConnectionType = "NPort"
	// PointToPointPortConnectionType This port connects in a point-to-point
	// configuration.
	PointToPointPortConnectionType PortConnectionType = "PointToPoint"
	// PrivateLoopPortConnectionType This port connects in a private loop
	// configuration.
	PrivateLoopPortConnectionType PortConnectionType = "PrivateLoop"
	// PublicLoopPortConnectionType This port connects in a public configuration.
	PublicLoopPortConnectionType PortConnectionType = "PublicLoop"
	// GenericPortConnectionType This port connection type is a generic fabric
	// port.
	GenericPortConnectionType PortConnectionType = "Generic"
	// ExtenderFabricPortConnectionType This port connection type is an extender
	// fabric port.
	ExtenderFabricPortConnectionType PortConnectionType = "ExtenderFabric"
	// FPortPortConnectionType This port connection type is a fabric port.
	FPortPortConnectionType PortConnectionType = "FPort"
	// EPortPortConnectionType This port connection type is an extender fabric
	// port.
	EPortPortConnectionType PortConnectionType = "EPort"
	// TEPortPortConnectionType This port connection type is a trunking extender
	// fabric port.
	TEPortPortConnectionType PortConnectionType = "TEPort"
	// NPPortPortConnectionType This port connection type is a proxy N-port for
	// N-port virtualization.
	NPPortPortConnectionType PortConnectionType = "NPPort"
	// GPortPortConnectionType This port connection type is a generic fabric port.
	GPortPortConnectionType PortConnectionType = "GPort"
	// NLPortPortConnectionType This port connects in a node loop configuration.
	NLPortPortConnectionType PortConnectionType = "NLPort"
	// FLPortPortConnectionType This port connects in a fabric loop configuration.
	FLPortPortConnectionType PortConnectionType = "FLPort"
	// EXPortPortConnectionType This port connection type is an external fabric
	// port.
	EXPortPortConnectionType PortConnectionType = "EXPort"
	// UPortPortConnectionType This port connection type is unassigned.
	UPortPortConnectionType PortConnectionType = "UPort"
	// DPortPortConnectionType This port connection type is a diagnostic port.
	DPortPortConnectionType PortConnectionType = "DPort"
)

type PortMedium string

const (
	// ElectricalPortMedium This port has an electrical cable connection.
	ElectricalPortMedium PortMedium = "Electrical"
	// OpticalPortMedium This port has an optical cable connection.
	OpticalPortMedium PortMedium = "Optical"
)

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
	// SFPSFPType The transceiver conforms to the SFF Specification for SFP.
	SFPSFPType SFPType = "SFP"
	// SFPPlusSFPType The transceiver conforms to the SFF Specification for SFP+.
	SFPPlusSFPType SFPType = "SFPPlus"
	// SFP28SFPType The transceiver conforms to the SFF Specification for SFP+ and
	// IEEE 802.3by Specification.
	SFP28SFPType SFPType = "SFP28"
	// cSFPSFPType The transceiver conforms to the CSFP MSA Specification.
	CSFPSFPType SFPType = "cSFP"
	// SFPDDSFPType The transceiver conforms to the SFP-DD MSA Specification.
	SFPDDSFPType SFPType = "SFPDD"
	// QSFPSFPType The transceiver conforms to the SFF Specification for QSFP.
	QSFPSFPType SFPType = "QSFP"
	// QSFPPlusSFPType The transceiver conforms to the SFF Specification for QSFP+.
	QSFPPlusSFPType SFPType = "QSFPPlus"
	// QSFP14SFPType The transceiver conforms to the SFF Specification for QSFP14.
	QSFP14SFPType SFPType = "QSFP14"
	// QSFP28SFPType The transceiver conforms to the SFF Specification for QSFP28.
	QSFP28SFPType SFPType = "QSFP28"
	// QSFP56SFPType The transceiver conforms to the SFF Specification for QSFP56.
	QSFP56SFPType SFPType = "QSFP56"
	// MiniSASHDSFPType The transceiver conforms to the SFF Specification SFF-8644.
	MiniSASHDSFPType SFPType = "MiniSASHD"
	// QSFPDDSFPType The transceiver conforms to the QSFP Double Density
	// Specification.
	QSFPDDSFPType SFPType = "QSFPDD"
	// OSFPSFPType The transceiver conforms to the OSFP Specification.
	OSFPSFPType SFPType = "OSFP"
	// CDFPSFPType The transceiver conforms to the SNIA SFF-TA-1032 Specification.
	CDFPSFPType SFPType = "CDFP"
)

type SupportedEthernetCapabilities string

const (
	// WakeOnLANSupportedEthernetCapabilities Wake on LAN (WoL) is supported on
	// this port.
	WakeOnLANSupportedEthernetCapabilities SupportedEthernetCapabilities = "WakeOnLAN"
	// EEESupportedEthernetCapabilities IEEE 802.3az Energy-Efficient Ethernet
	// (EEE) is supported on this port.
	EEESupportedEthernetCapabilities SupportedEthernetCapabilities = "EEE"
)

type TransceiverManagementInterfaceType string

const (
	// SFPTransceiverManagementInterfaceType shall indicate that the transceiver is
	// managed via an interface defined by SFF-8472, SFF-8436, or SFF-8636.
	SFPTransceiverManagementInterfaceType TransceiverManagementInterfaceType = "SFP"
	// CMISTransceiverManagementInterfaceType shall indicate that the transceiver
	// is managed via an interface defined by the Common Management Interface
	// Specification (CMIS).
	CMISTransceiverManagementInterfaceType TransceiverManagementInterfaceType = "CMIS"
)

type UALinkGeneration string

const (
	// UALink128GUALinkGeneration The port conforms to UALink 128G Data Link Layer
	// and Physical Layer requirements.
	UALink128GUALinkGeneration UALinkGeneration = "UALink128G"
	// UALink200GUALinkGeneration The port conforms to UALink 200G Data Link Layer
	// and Physical Layer requirements.
	UALink200GUALinkGeneration UALinkGeneration = "UALink200G"
)

// Port This resource contains a simple port for a Redfish implementation.
type Port struct {
	Entity
	// ActiveWidth shall contain the number of active lanes for this interface.
	//
	// Version added: v1.2.0
	ActiveWidth int
	// AssociatedPhysicalPort shall contain the index of the physical port from
	// which this subport has been split. If 'IsSplit' contains 'false', this
	// property shall not be present.
	//
	// Version added: v1.17.0
	AssociatedPhysicalPort *int `json:",omitempty"`
	// CXL shall contain CXL-specific properties for this port.
	//
	// Version added: v1.8.0
	CXL CXLPort
	// CapableProtocolVersions shall contain the protocol versions capable of being
	// sent over this port. This property should only be used for protocols where
	// the version and not the speed is of primary interest such as USB,
	// DisplayPort, or HDMI.
	//
	// Version added: v1.4.0
	CapableProtocolVersions []string
	// ConfiguredSpeedGbps shall contain the unidirectional speed to which this
	// port is configured to train. This value includes overhead associated with
	// the protocol. If 'AutoSpeedNegotiationEnabled' contains 'true', this
	// property shall be ignored.
	//
	// Version added: v1.16.0
	ConfiguredSpeedGbps *float64 `json:",omitempty"`
	// ConfiguredWidth shall contain the number of physical transport links to
	// which this port is configured to train. If 'AutoSpeedNegotiationEnabled'
	// contains 'true', this property shall be ignored.
	//
	// Version added: v1.16.0
	ConfiguredWidth *int `json:",omitempty"`
	// CurrentProtocolVersion shall contain the protocol version being sent over
	// this port. This property should only be used for protocols where the version
	// and not the speed is of primary interest such as USB, DisplayPort, or HDMI.
	//
	// Version added: v1.4.0
	CurrentProtocolVersion string
	// CurrentSpeedGbps shall contain the unidirectional speed of this port
	// currently negotiated and running. This value includes overhead associated
	// with the protocol.
	CurrentSpeedGbps *float64 `json:",omitempty"`
	// Enabled shall indicate if this port is enabled. Disabling a port will
	// disconnect any devices only connected to the system through this port.
	//
	// Version added: v1.4.0
	//
	// Deprecated: v1.10.0
	// This property has been deprecated in favor of 'InterfaceEnabled'.
	Enabled bool
	// EnvironmentMetrics shall contain a link to a resource of type
	// 'EnvironmentMetrics' that represents the environment metrics for this port
	// or any attached small form-factor pluggable (SFP) device.
	//
	// Version added: v1.4.0
	environmentMetrics string
	// Ethernet shall contain Ethernet-specific properties of the port.
	//
	// Version added: v1.3.0
	Ethernet EthernetProperties
	// FibreChannel shall contain Fibre Channel-specific properties of the port.
	//
	// Version added: v1.3.0
	FibreChannel FibreChannelProperties
	// FirstLane shall contain the first lane that this port is configured to use.
	// If 'IsSplit' contains 'false', this property shall not be present.
	//
	// Version added: v1.17.0
	FirstLane *int `json:",omitempty"`
	// FunctionMaxBandwidth shall contain an array of maximum bandwidth allocation
	// percentages for the functions associated with this port.
	//
	// Version added: v1.4.0
	FunctionMaxBandwidth []FunctionMaxBandwidth
	// FunctionMinBandwidth shall contain an array of minimum bandwidth percentage
	// allocations for each of the functions associated with this port.
	//
	// Version added: v1.4.0
	FunctionMinBandwidth []FunctionMinBandwidth
	// GenZ shall contain Gen-Z specific properties for this interface.
	//
	// Version added: v1.2.0
	GenZ PortGenZ
	// HostDevice shall contain the current host device of port.
	//
	// Version added: v1.15.0
	HostDevice HostDeviceType
	// InfiniBand shall contain InfiniBand-specific properties of the port.
	//
	// Version added: v1.6.0
	InfiniBand InfiniBandProperties
	// InterfaceEnabled shall indicate whether the port is enabled. When disabled,
	// no traffic of any type, such as link protocol traffic and management
	// traffic, is sent or received by the port.
	//
	// Version added: v1.2.0
	InterfaceEnabled bool
	// IsAggregation shall indicate whether this is a logical aggregation.
	//
	// Version added: v1.18.0
	IsAggregation bool
	// IsSplit shall indicate whether this is a subport split from a physical port.
	//
	// Version added: v1.17.0
	IsSplit bool
	// LinkConfiguration shall contain the static capabilities and configuration
	// settings of the port.
	//
	// Version added: v1.3.0
	LinkConfiguration []LinkConfiguration
	// LinkNetworkTechnology shall contain the current network technology for this
	// port.
	//
	// Version added: v1.2.0
	LinkNetworkTechnology LinkNetworkTechnology
	// LinkState shall contain the desired link state for this interface.
	//
	// Version added: v1.2.0
	LinkState LinkState
	// LinkStatus shall contain the link status for this interface.
	//
	// Version added: v1.2.0
	LinkStatus PortLinkStatus
	// LinkTransitionIndicator shall contain the number of link state transitions
	// for this interface.
	//
	// Version added: v1.2.0
	LinkTransitionIndicator int
	// Location shall contain the location information of the associated port.
	//
	// Version added: v1.1.0
	Location Location
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource. A write to this property shall
	// update the value of 'IndicatorLED' in this resource, if supported, to
	// reflect the implementation of the locating function.
	//
	// Version added: v1.3.0
	LocationIndicatorActive bool
	// MaxFrameSize shall contain the maximum frame size supported by the port.
	//
	// Version added: v1.3.0
	MaxFrameSize *int `json:",omitempty"`
	// MaxSpeedGbps shall contain the maximum unidirectional speed of which this
	// port is capable of being configured. If capable of autonegotiation, the
	// system shall attempt to negotiate at the maximum speed set. This value
	// includes overhead associated with the protocol.
	MaxSpeedGbps *float64 `json:",omitempty"`
	// Metrics shall contain a link to the metrics associated with this port. If
	// 'IsAggregation' contains 'true', the metric values in the linked resource
	// represent the combined metrics from all associated physical ports.
	//
	// Version added: v1.2.0
	metrics string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PortID shall contain the hardware-defined identifier of this port. The
	// human-readable name of this port is described in the 'ServiceLabel' property
	// within 'Location' for this port.
	PortID string `json:"PortId"`
	// PortMedium shall contain the physical connection medium for this port.
	//
	// Version added: v1.2.0
	PortMedium PortMedium
	// PortProtocol shall contain the protocol being sent over this port.
	PortProtocol Protocol
	// PortType shall contain the port type for this port.
	PortType PortType
	// RemotePortID shall contain the identifier of the remote port, such as a
	// switch or device, to which this port is connected.
	//
	// Version added: v1.8.0
	RemotePortID string `json:"RemotePortId"`
	// SFP shall contain data about the small form-factor pluggable (SFP) device
	// currently occupying this port.
	//
	// Version added: v1.4.0
	SFP SFP
	// SignalDetected shall indicate whether a signal that is appropriate for this
	// link technology is detected for this port.
	//
	// Version added: v1.2.0
	SignalDetected bool
	// Status shall contain any status or health properties of the resource.
	Status Status
	// UALink shall contain UALink attributes of the port.
	//
	// Version added: v1.18.0
	UALink PortUALink
	// Width shall contain the number of physical transport links that this port
	// contains.
	Width *int `json:",omitempty"`
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// resetPPBTarget is the URL to send ResetPPB requests.
	resetPPBTarget string
	// associatedEndpoints are the URIs for AssociatedEndpoints.
	associatedEndpoints []string
	// associatedPhysicalPorts are the URIs for AssociatedPhysicalPorts.
	associatedPhysicalPorts []string
	// cables are the URIs for Cables.
	cables []string
	// connectedPorts are the URIs for ConnectedPorts.
	connectedPorts []string
	// connectedSwitchPorts are the URIs for ConnectedSwitchPorts.
	connectedSwitchPorts []string
	// connectedSwitches are the URIs for ConnectedSwitches.
	connectedSwitches []string
	// ethernetInterfaces are the URIs for EthernetInterfaces.
	ethernetInterfaces []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Port object from the raw JSON.
func (p *Port) UnmarshalJSON(b []byte) error {
	type temp Port
	type pActions struct {
		Reset    ActionTarget `json:"#Port.Reset"`
		ResetPPB ActionTarget `json:"#Port.ResetPPB"`
	}
	type pLinks struct {
		AssociatedEndpoints     Links `json:"AssociatedEndpoints"`
		AssociatedPhysicalPorts Links `json:"AssociatedPhysicalPorts"`
		Cables                  Links `json:"Cables"`
		ConnectedPorts          Links `json:"ConnectedPorts"`
		ConnectedSwitchPorts    Links `json:"ConnectedSwitchPorts"`
		ConnectedSwitches       Links `json:"ConnectedSwitches"`
		EthernetInterfaces      Links `json:"EthernetInterfaces"`
	}
	var tmp struct {
		temp
		Actions            pActions
		Links              pLinks
		EnvironmentMetrics Link `json:"EnvironmentMetrics"`
		Metrics            Link `json:"Metrics"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = Port(tmp.temp)

	// Extract the links to other entities for later
	p.resetTarget = tmp.Actions.Reset.Target
	p.resetPPBTarget = tmp.Actions.ResetPPB.Target
	p.associatedEndpoints = tmp.Links.AssociatedEndpoints.ToStrings()
	p.associatedPhysicalPorts = tmp.Links.AssociatedPhysicalPorts.ToStrings()
	p.cables = tmp.Links.Cables.ToStrings()
	p.connectedPorts = tmp.Links.ConnectedPorts.ToStrings()
	p.connectedSwitchPorts = tmp.Links.ConnectedSwitchPorts.ToStrings()
	p.connectedSwitches = tmp.Links.ConnectedSwitches.ToStrings()
	p.ethernetInterfaces = tmp.Links.EthernetInterfaces.ToStrings()
	p.environmentMetrics = tmp.EnvironmentMetrics.String()
	p.metrics = tmp.Metrics.String()

	// This is a read/write object, so we need to save the raw object data for later
	p.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (p *Port) Update() error {
	readWriteFields := []string{
		"ConfiguredSpeedGbps",
		"ConfiguredWidth",
		"Enabled",
		"HostDevice",
		"InterfaceEnabled",
		"LinkState",
		"LinkTransitionIndicator",
		"LocationIndicatorActive",
		"PortType",
	}

	return p.UpdateFromRawData(p, p.RawData, readWriteFields)
}

// GetPort will get a Port instance from the service.
func GetPort(c Client, uri string) (*Port, error) {
	return GetObject[Port](c, uri)
}

// ListReferencedPorts gets the collection of Port from
// a provided reference.
func ListReferencedPorts(c Client, link string) ([]*Port, error) {
	return GetCollectionObjects[Port](c, link)
}

// This action shall reset this port.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *Port) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetType"] = resetType
	resp, taskInfo, err := PostWithTask(p.client,
		p.resetTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset the PCI-to-PCI bridge (PPB) for this port.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *Port) ResetPPB() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(p.client,
		p.resetPPBTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// AssociatedEndpoints gets the AssociatedEndpoints linked resources.
func (p *Port) AssociatedEndpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](p.client, p.associatedEndpoints)
}

// AssociatedPhysicalPorts gets the AssociatedPhysicalPorts linked resources.
func (p *Port) AssociatedPhysicalPorts() ([]*Port, error) {
	return GetObjects[Port](p.client, p.associatedPhysicalPorts)
}

// Cables gets the Cables linked resources.
func (p *Port) Cables() ([]*Cable, error) {
	return GetObjects[Cable](p.client, p.cables)
}

// ConnectedPorts gets the ConnectedPorts linked resources.
func (p *Port) ConnectedPorts() ([]*Port, error) {
	return GetObjects[Port](p.client, p.connectedPorts)
}

// ConnectedSwitchPorts gets the ConnectedSwitchPorts linked resources.
func (p *Port) ConnectedSwitchPorts() ([]*Port, error) {
	return GetObjects[Port](p.client, p.connectedSwitchPorts)
}

// ConnectedSwitches gets the ConnectedSwitches linked resources.
func (p *Port) ConnectedSwitches() ([]*Switch, error) {
	return GetObjects[Switch](p.client, p.connectedSwitches)
}

// EthernetInterfaces gets the EthernetInterfaces linked resources.
func (p *Port) EthernetInterfaces() ([]*EthernetInterface, error) {
	return GetObjects[EthernetInterface](p.client, p.ethernetInterfaces)
}

// EnvironmentMetrics gets the EnvironmentMetrics linked resource.
func (p *Port) EnvironmentMetrics() (*EnvironmentMetrics, error) {
	if p.environmentMetrics == "" {
		return nil, nil
	}
	return GetObject[EnvironmentMetrics](p.client, p.environmentMetrics)
}

// Metrics gets the Metrics linked resource.
func (p *Port) Metrics() (*PortMetrics, error) {
	if p.metrics == "" {
		return nil, nil
	}
	return GetObject[PortMetrics](p.client, p.metrics)
}

// CXLPort shall contain CXL-specific properties for a port.
type CXLPort struct {
	// Congestion shall contain the congestion properties for this CXL port.
	//
	// Version added: v1.8.0
	Congestion Congestion
	// ConnectedDeviceMode shall contain the CXL Specification-defined 'Connected
	// Device Mode'.
	//
	// Version added: v1.8.0
	ConnectedDeviceMode ConnectedDeviceMode
	// ConnectedDeviceType shall contain the CXL Specification-defined 'Connected
	// Device Type'.
	//
	// Version added: v1.8.0
	ConnectedDeviceType ConnectedDeviceType
	// CurrentPortConfigurationState shall contain the CXL Specification-defined
	// 'Current Port Configuration State'.
	//
	// Version added: v1.8.0
	CurrentPortConfigurationState CurrentPortConfigurationState
	// MaxLogicalDeviceCount shall contain the CXL Specification-defined 'Supported
	// LD Count'.
	//
	// Version added: v1.8.0
	MaxLogicalDeviceCount *int `json:",omitempty"`
	// QoSTelemetryCapabilities shall contain the quality of service telemetry
	// capabilities for this CXL port.
	//
	// Version added: v1.8.0
	QoSTelemetryCapabilities QoSTelemetryCapabilities
	// SupportedCXLModes shall contain the CXL Specification-defined 'Supported CXL
	// Modes'. This property shall not contain the value 'Disconnected'.
	//
	// Version added: v1.11.0
	SupportedCXLModes []ConnectedDeviceMode
	// TemporaryThroughputReductionEnabled shall indicate whether the CXL
	// Specification-defined 'Temporary Throughput Reduction' mechanism is enabled.
	//
	// Version added: v1.8.0
	//
	// Deprecated: v1.12.0
	// This property has been deprecated in favor of
	// 'TemporaryThroughputReductionEnabled' in 'PCIeDevice'.
	TemporaryThroughputReductionEnabled bool
}

// ConfiguredNetworkLink shall contain a set of link settings that a port is
// configured to use for autonegotiation.
type ConfiguredNetworkLink struct {
	// ConfiguredLinkSpeedGbps shall contain the network link speed per lane this
	// port is configured to allow for autonegotiation purposes. This value
	// includes overhead associated with the protocol.
	//
	// Version added: v1.3.0
	ConfiguredLinkSpeedGbps *float64 `json:",omitempty"`
	// ConfiguredWidth shall contain the network link width this port is configured
	// to use for autonegotiation purposes.
	//
	// Version added: v1.3.0
	ConfiguredWidth *int `json:",omitempty"`
}

// Congestion shall contain the congestion properties for a CXL port.
type Congestion struct {
	// BackpressureSampleInterval shall contain the CXL Specification-defined
	// 'Backpressure Sample Interval' in nanoseconds.
	//
	// Version added: v1.8.0
	BackpressureSampleInterval *int `json:",omitempty"`
	// CompletionCollectionInterval shall contain the CXL Specification-defined
	// 'Completion Collection Interval' in nanoseconds.
	//
	// Version added: v1.8.0
	CompletionCollectionInterval *int `json:",omitempty"`
	// CongestionTelemetryEnabled shall indicate whether congestion telemetry
	// collection is enabled for this port.
	//
	// Version added: v1.8.0
	CongestionTelemetryEnabled bool
	// EgressModeratePercentage shall contain the threshold for moderate egress
	// port congestion for the CXL Specification-defined 'Egress Port Congestion'
	// mechanism as a percentage, '0' to '100'.
	//
	// Version added: v1.8.0
	EgressModeratePercentage *uint `json:",omitempty"`
	// EgressSeverePercentage shall contain the CXL Specification-defined 'Egress
	// Severe Percentage' as a percentage, '0' to '100'.
	//
	// Version added: v1.8.0
	EgressSeverePercentage *uint `json:",omitempty"`
	// MaxSustainedRequestCmpBias shall contain the CXL Specification-defined
	// 'ReqCmpBasis'.
	//
	// Version added: v1.8.0
	MaxSustainedRequestCmpBias *int `json:",omitempty"`
}

// EthernetProperties shall contain Ethernet-specific properties for a port.
type EthernetProperties struct {
	// AssociatedMACAddresses shall contain an array of all configured MAC
	// addresses associated with this network port.
	//
	// Version added: v1.4.0
	AssociatedMACAddresses []string
	// EEEEnabled shall indicate whether IEEE 802.3az Energy-Efficient Ethernet
	// (EEE) is enabled on this port.
	//
	// Version added: v1.5.0
	EEEEnabled bool
	// FlowControlConfiguration shall contain the locally configured 802.3x flow
	// control setting for this port.
	//
	// Version added: v1.3.0
	FlowControlConfiguration FlowControl
	// FlowControlStatus shall contain the 802.3x flow control behavior negotiated
	// with the link partner for this port.
	//
	// Version added: v1.3.0
	FlowControlStatus FlowControl
	// LLDPEnabled shall contain the state indicating whether to enable LLDP for a
	// port. If LLDP is disabled at the adapter level, this property shall be
	// ignored.
	//
	// Version added: v1.4.0
	LLDPEnabled bool
	// LLDPReceive shall contain the LLDP data being received on this link.
	//
	// Version added: v1.4.0
	LLDPReceive LLDPReceive
	// LLDPTransmit shall contain the LLDP data being transmitted on this link.
	//
	// Version added: v1.4.0
	LLDPTransmit LLDPTransmit
	// SupportedEthernetCapabilities shall contain an array of Ethernet
	// capabilities supported by this port.
	//
	// Version added: v1.3.0
	//
	// Deprecated: v1.5.0
	// This property has been deprecated in favor of individual fields for the
	// various properties.
	SupportedEthernetCapabilities []SupportedEthernetCapabilities
	// WakeOnLANEnabled shall indicate whether Wake on LAN (WoL) is enabled on this
	// port.
	//
	// Version added: v1.5.0
	WakeOnLANEnabled bool
}

// FibreChannelProperties shall contain Fibre Channel-specific properties for a
// port.
type FibreChannelProperties struct {
	// AssociatedWorldWideNames shall contain an array of configured World Wide
	// Names (WWN) that are associated with this network port, including the
	// programmed address of the lowest-numbered network device function, the
	// configured but not active address if applicable, the address for hardware
	// port teaming, or other network addresses.
	//
	// Version added: v1.4.0
	AssociatedWorldWideNames []string
	// FabricName shall indicate the Fibre Channel Fabric Name provided by the
	// switch.
	//
	// Version added: v1.3.0
	FabricName string
	// NumberDiscoveredRemotePorts shall contain the number of ports not on this
	// associated device that this port has discovered.
	//
	// Version added: v1.3.0
	NumberDiscoveredRemotePorts *int `json:",omitempty"`
	// PortConnectionType shall contain the connection type for this port.
	//
	// Version added: v1.3.0
	PortConnectionType PortConnectionType
}

// FunctionMaxBandwidth shall describe a maximum bandwidth percentage allocation
// for a network device function associated with a port.
type FunctionMaxBandwidth struct {
	// AllocationPercent shall contain the maximum bandwidth percentage allocation,
	// '0' to '100', for the associated network device function.
	//
	// Version added: v1.4.0
	AllocationPercent *uint `json:",omitempty"`
	// NetworkDeviceFunction shall contain a link to a resource of type
	// 'NetworkDeviceFunction' that represents the network device function
	// associated with this bandwidth setting of this network port.
	//
	// Version added: v1.4.0
	networkDeviceFunction string
}

// UnmarshalJSON unmarshals a FunctionMaxBandwidth object from the raw JSON.
func (f *FunctionMaxBandwidth) UnmarshalJSON(b []byte) error {
	type temp FunctionMaxBandwidth
	var tmp struct {
		temp
		NetworkDeviceFunction Link `json:"NetworkDeviceFunction"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = FunctionMaxBandwidth(tmp.temp)

	// Extract the links to other entities for later
	f.networkDeviceFunction = tmp.NetworkDeviceFunction.String()

	return nil
}

// NetworkDeviceFunction gets the NetworkDeviceFunction linked resource.
func (f *FunctionMaxBandwidth) NetworkDeviceFunction(client Client) (*NetworkDeviceFunction, error) {
	if f.networkDeviceFunction == "" {
		return nil, nil
	}
	return GetObject[NetworkDeviceFunction](client, f.networkDeviceFunction)
}

// FunctionMinBandwidth shall describe a minimum bandwidth percentage allocation
// for a network device function associated with a port.
type FunctionMinBandwidth struct {
	// AllocationPercent shall contain the minimum bandwidth percentage allocation,
	// '0' to '100', for the associated network device function. The sum of all
	// minimum percentages shall not exceed '100'.
	//
	// Version added: v1.4.0
	AllocationPercent *uint `json:",omitempty"`
	// NetworkDeviceFunction shall contain a link to a resource of type
	// 'NetworkDeviceFunction' that represents the network device function
	// associated with this bandwidth setting of this network port.
	//
	// Version added: v1.4.0
	networkDeviceFunction string
}

// UnmarshalJSON unmarshals a FunctionMinBandwidth object from the raw JSON.
func (f *FunctionMinBandwidth) UnmarshalJSON(b []byte) error {
	type temp FunctionMinBandwidth
	var tmp struct {
		temp
		NetworkDeviceFunction Link `json:"NetworkDeviceFunction"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*f = FunctionMinBandwidth(tmp.temp)

	// Extract the links to other entities for later
	f.networkDeviceFunction = tmp.NetworkDeviceFunction.String()

	return nil
}

// NetworkDeviceFunction gets the NetworkDeviceFunction linked resource.
func (f *FunctionMinBandwidth) NetworkDeviceFunction(client Client) (*NetworkDeviceFunction, error) {
	if f.networkDeviceFunction == "" {
		return nil, nil
	}
	return GetObject[NetworkDeviceFunction](client, f.networkDeviceFunction)
}

// PortGenZ shall contain Gen-Z specific port properties.
type PortGenZ struct {
	// LPRT shall contain a link to a resource collection of type
	// 'RouteEntryCollection', and shall represent the Gen-Z Core
	// Specification-defined Linear Packet Relay Table for this port.
	//
	// Version added: v1.2.0
	lPRT string
	// MPRT shall contain a link to a resource collection of type
	// 'RouteEntryCollection', and shall represent the Gen-Z Core
	// Specification-defined Multi-subnet Packet Relay Table for this port.
	//
	// Version added: v1.2.0
	mPRT string
	// VCAT shall contain a link to a resource collection of type
	// 'VCATEntryCollection'.
	//
	// Version added: v1.2.0
	vCAT string
}

// UnmarshalJSON unmarshals a GenZ object from the raw JSON.
func (g *PortGenZ) UnmarshalJSON(b []byte) error {
	type temp PortGenZ
	var tmp struct {
		temp
		LPRT Link `json:"LPRT"`
		MPRT Link `json:"MPRT"`
		VCAT Link `json:"VCAT"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*g = PortGenZ(tmp.temp)

	// Extract the links to other entities for later
	g.lPRT = tmp.LPRT.String()
	g.mPRT = tmp.MPRT.String()
	g.vCAT = tmp.VCAT.String()

	return nil
}

// LPRT gets the LPRT collection.
func (g *PortGenZ) LPRT(client Client) ([]*RouteEntry, error) {
	if g.lPRT == "" {
		return nil, nil
	}
	return GetCollectionObjects[RouteEntry](client, g.lPRT)
}

// MPRT gets the MPRT collection.
func (g *PortGenZ) MPRT(client Client) ([]*RouteEntry, error) {
	if g.mPRT == "" {
		return nil, nil
	}
	return GetCollectionObjects[RouteEntry](client, g.mPRT)
}

// VCAT gets the VCAT collection.
func (g *PortGenZ) VCAT(client Client) ([]*VCATEntry, error) {
	if g.vCAT == "" {
		return nil, nil
	}
	return GetCollectionObjects[VCATEntry](client, g.vCAT)
}

// InfiniBandProperties shall contain InfiniBand-specific properties for a port.
type InfiniBandProperties struct {
	// AssociatedNodeGUIDs shall contain an array of configured node GUIDs that are
	// associated with this network port, including the programmed address of the
	// lowest-numbered network device function, the configured but not active
	// address if applicable, the address for hardware port teaming, or other
	// network addresses.
	//
	// Version added: v1.6.0
	AssociatedNodeGUIDs []string
	// AssociatedPortGUIDs shall contain an array of configured port GUIDs that are
	// associated with this network port, including the programmed address of the
	// lowest-numbered network device function, the configured but not active
	// address if applicable, the address for hardware port teaming, or other
	// network addresses.
	//
	// Version added: v1.6.0
	AssociatedPortGUIDs []string
	// AssociatedSystemGUIDs shall contain an array of configured system GUIDs that
	// are associated with this network port, including the programmed address of
	// the lowest-numbered network device function, the configured but not active
	// address if applicable, the address for hardware port teaming, or other
	// network addresses.
	//
	// Version added: v1.6.0
	AssociatedSystemGUIDs []string
}

// LLDPReceive shall contain the LLDP data from the remote partner across this
// link.
type LLDPReceive struct {
	// AdditionalManagementAddressesIPv4 shall contain additional IPv4 management
	// addresses received from the remote partner across this link.
	// 'ManagementAddressIPv4' shall contain the first IPv4 address received. This
	// property should not be present if the port did not receive more than one
	// IPv4 management address.
	//
	// Version added: v1.18.0
	AdditionalManagementAddressesIPv4 []string
	// AdditionalManagementAddressesIPv6 shall contain additional IPv6 management
	// addresses received from the remote partner across this link.
	// 'ManagementAddressIPv6' shall contain the first IPv6 address received. This
	// property should not be present if the port did not receive more than one
	// IPv6 management address.
	//
	// Version added: v1.18.0
	AdditionalManagementAddressesIPv6 []string
	// ChassisID shall contain the chassis ID received from the remote partner
	// across this link. If no such chassis ID has been received, this property
	// should not be present.
	//
	// Version added: v1.4.0
	ChassisID string `json:"ChassisId"`
	// ChassisIDSubtype shall contain the IEEE 802.1AB-2009 chassis ID subtype
	// received from the remote partner across this link. If no such chassis ID
	// subtype has been received, this property should not be present.
	//
	// Version added: v1.4.0
	ChassisIDSubtype IEEE802IDSubtype `json:"ChassisIdSubtype"`
	// ManagementAddressIPv4 shall contain the IPv4 management address received
	// from the remote partner across this link. If no such management address has
	// been received, this property should not be present.
	//
	// Version added: v1.4.0
	ManagementAddressIPv4 string
	// ManagementAddressIPv6 shall contain the IPv6 management address received
	// from the remote partner across this link. If no such management address has
	// been received, this property should not be present.
	//
	// Version added: v1.4.0
	ManagementAddressIPv6 string
	// ManagementAddressMAC shall contain the management MAC address received from
	// the remote partner across this link. If no such management address has been
	// received, this property should not be present.
	//
	// Version added: v1.4.0
	ManagementAddressMAC string
	// ManagementVlanID shall contain the management VLAN ID received from the
	// remote partner across this link. If no such management VLAN ID has been
	// received, this property should not be present.
	//
	// Version added: v1.4.0
	ManagementVlanID *uint `json:"ManagementVlanId,omitempty"`
	// PortID shall contain a colon-delimited string of hexadecimal octets
	// identifying the port received from the remote partner across this link. If
	// no such port ID has been received, this property should not be present.
	//
	// Version added: v1.4.0
	PortID string `json:"PortId"`
	// PortIDSubtype shall contain the port ID subtype from IEEE 802.1AB-2009 Table
	// 8-3 received from the remote partner across this link. If no such port ID
	// subtype has been received, this property should not be present.
	//
	// Version added: v1.4.0
	PortIDSubtype IEEE802IDSubtype `json:"PortIdSubtype"`
	// SystemCapabilities shall contain the system capabilities received from the
	// remote partner across this link. If no such system capabilities have been
	// received, this property shall not be present. This property shall not
	// contain the value 'None'.
	//
	// Version added: v1.8.0
	SystemCapabilities []LLDPSystemCapabilities
	// SystemDescription shall contain the system description received from the
	// remote partner across this link. If no such system description has been
	// received, this property shall not be present.
	//
	// Version added: v1.8.0
	SystemDescription string
	// SystemName shall contain the system name received from the remote partner
	// across this link. If no such system name has been received, this property
	// shall not be present.
	//
	// Version added: v1.8.0
	SystemName string
}

// LLDPTransmit shall contain the LLDP data to be transmitted from this
// endpoint.
type LLDPTransmit struct {
	// AdditionalManagementAddressesIPv4 shall contain additional IPv4 management
	// addresses transmitted on this link. 'ManagementAddressIPv4' shall contain
	// the first IPv4 address for this port. This property should not be present if
	// the port does not transmit more than one IPv4 management address.
	//
	// Version added: v1.18.0
	AdditionalManagementAddressesIPv4 []string
	// AdditionalManagementAddressesIPv6 shall contain additional IPv6 management
	// addresses transmitted on this link. 'ManagementAddressIPv6' shall contain
	// the first IPv6 address for this port. This property should not be present if
	// the port does not transmit more than one IPv6 management address.
	//
	// Version added: v1.18.0
	AdditionalManagementAddressesIPv6 []string
	// ChassisID shall contain the chassis ID to be transmitted from this endpoint.
	// If no such chassis ID is to be transmitted, this value shall be an empty
	// string.
	//
	// Version added: v1.4.0
	ChassisID string `json:"ChassisId"`
	// ChassisIDSubtype shall contain the IEEE 802.1AB-2009 chassis ID subtype to
	// be transmitted from this endpoint. If no such chassis ID subtype is to be
	// transmitted, this value shall be 'NotTransmitted'.
	//
	// Version added: v1.4.0
	ChassisIDSubtype IEEE802IDSubtype `json:"ChassisIdSubtype"`
	// ManagementAddressIPv4 shall contain the IPv4 management address to be
	// transmitted from this endpoint. If no such management address is to be
	// transmitted, this value shall be an empty string.
	//
	// Version added: v1.4.0
	ManagementAddressIPv4 string
	// ManagementAddressIPv6 shall contain the IPv6 management address to be
	// transmitted from this endpoint. If no such management address is to be
	// transmitted, this value shall be an empty string.
	//
	// Version added: v1.4.0
	ManagementAddressIPv6 string
	// ManagementAddressMAC shall contain the management MAC address to be
	// transmitted from this endpoint. If no such management address is to be
	// transmitted, this value shall be an empty string.
	//
	// Version added: v1.4.0
	ManagementAddressMAC string
	// ManagementVlanID shall contain the management VLAN ID to be transmitted from
	// this endpoint. If no such port ID is to be transmitted, this value shall be
	// '4095'.
	//
	// Version added: v1.4.0
	ManagementVlanID *uint `json:"ManagementVlanId,omitempty"`
	// PortID shall contain a colon-delimited string of hexadecimal octets
	// identifying the port for an LLDP endpoint. If no such port ID is to be
	// transmitted, this value shall be an empty string.
	//
	// Version added: v1.4.0
	PortID string `json:"PortId"`
	// PortIDSubtype shall contain the port ID subtype from IEEE 802.1AB-2009 Table
	// 8-3 to be transmitted from this endpoint. If no such port ID subtype is to
	// be transmitted, this value shall be 'NotTransmitted'.
	//
	// Version added: v1.4.0
	PortIDSubtype IEEE802IDSubtype `json:"PortIdSubtype"`
	// SystemCapabilities shall contain the system capabilities to be transmitted
	// from this endpoint. If no such system capabilities are to be transmitted,
	// this value shall be an empty array. If this property contains the value
	// 'None', an empty set of system capabilities is transmitted from this
	// endpoint.
	//
	// Version added: v1.8.0
	SystemCapabilities []LLDPSystemCapabilities
	// SystemDescription shall contain the system description to be transmitted
	// from this endpoint. If no such system description is to be transmitted, this
	// value shall be an empty string.
	//
	// Version added: v1.8.0
	SystemDescription string
	// SystemName shall contain the system name to be transmitted from this
	// endpoint. If no such system name is to be transmitted, this value shall be
	// an empty string.
	//
	// Version added: v1.8.0
	SystemName string
}

// LinkConfiguration shall contain static capabilities and configuration
// settings for a link.
type LinkConfiguration struct {
	// AutoSpeedNegotiationCapable shall indicate whether the port is capable of
	// autonegotiating speed.
	//
	// Version added: v1.3.0
	AutoSpeedNegotiationCapable bool
	// AutoSpeedNegotiationEnabled shall indicate whether the port is configured to
	// autonegotiate speed.
	//
	// Version added: v1.3.0
	AutoSpeedNegotiationEnabled bool
	// CapableLinkSpeedGbps shall contain all of the possible network link speed
	// capabilities of this port. This value includes overhead associated with the
	// protocol.
	//
	// Version added: v1.3.0
	CapableLinkSpeedGbps []*float64
	// ConfiguredNetworkLinks shall contain the set of link speed and width pairs
	// to which this port is restricted for autonegotiation purposes. An empty
	// array shall indicate autoconfiguration uses any available link speed and
	// width pairs.
	//
	// Version added: v1.3.0
	ConfiguredNetworkLinks []ConfiguredNetworkLink
	// LinkNetworkTechnology shall contain the link network technology for this
	// link configuration.
	//
	// Version added: v1.16.0
	LinkNetworkTechnology LinkNetworkTechnology
}

// QoSTelemetryCapabilities shall contain the quality of service telemetry
// capabilities for a CXL port.
type QoSTelemetryCapabilities struct {
	// EgressPortBackpressureSupported shall indicate whether the port supports the
	// CXL Specification-defined 'Egress Port Backpressure' mechanism.
	//
	// Version added: v1.8.0
	EgressPortBackpressureSupported bool
	// TemporaryThroughputReductionSupported shall indicate whether the port
	// supports the CXL Specification-defined 'Temporary Throughput Reduction'
	// mechanism.
	//
	// Version added: v1.8.0
	//
	// Deprecated: v1.12.0
	// This property has been deprecated in favor of
	// 'TemporaryThroughputReductionSupported' in 'PCIeDevice'.
	TemporaryThroughputReductionSupported bool
}

// SFP shall describe a transceiver, such as a small form-factor pluggable (SFP)
// device, attached to a port.
type SFP struct {
	// DateCode shall contain the manufacturing date code for this SFP as
	// determined by the vendor or supplier.
	//
	// Version added: v1.14.0
	DateCode string
	// FiberConnectionType shall contain the fiber connection type used by the
	// transceiver.
	//
	// Version added: v1.4.0
	FiberConnectionType FiberConnectionType
	// ManagementInterfaceType shall contain the type of management interface for
	// this transceiver.
	//
	// Version added: v1.18.0
	ManagementInterfaceType TransceiverManagementInterfaceType
	// Manufacturer shall contain the name of the organization responsible for
	// producing the transceiver. This organization may be the entity from which
	// the transceiver is purchased, but this is not necessarily true.
	//
	// Version added: v1.4.0
	Manufacturer string
	// MediumType shall contain the medium type used by the transceiver.
	//
	// Version added: v1.4.0
	MediumType MediumType
	// PartNumber shall contain the manufacturer-provided part number for the
	// transceiver.
	//
	// Version added: v1.4.0
	PartNumber string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the transceiver.
	//
	// Version added: v1.4.0
	SerialNumber string
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.4.0
	Status Status
	// SupportedSFPTypes shall contain an array of transceiver types supported by
	// this port.
	//
	// Version added: v1.4.0
	SupportedSFPTypes []SFPType
	// Type shall contain the transceiver type currently attached to this port.
	//
	// Version added: v1.4.0
	Type SFPType
	// VendorOUI shall contain the IEEE organizationally unique identifier (OUI) of
	// the vendor of this SFP.
	//
	// Version added: v1.13.0
	VendorOUI string
	// Version shall contain the hardware version of this SFP as determined by the
	// vendor or supplier.
	//
	// Version added: v1.13.0
	Version string
}

// PortUALink shall contain information about UALink port management.
type PortUALink struct {
	// AuthenticationModeEnabled shall indicate whether UALink-defined
	// authentication mode is enabled on the port. Authentication mode affects
	// Transaction Layer flit packing and unpacking.
	//
	// Version added: v1.18.0
	AuthenticationModeEnabled bool
	// Generation This property contains the generation of the UALink Data Link
	// Layer and Physical Layer of this port.
	//
	// Version added: v1.18.0
	Generation UALinkGeneration
	// NonStrictOrderingModeEnabled shall indicate whether UALink-defined
	// non-strict ordering mode is enabled on the port. Non-strict ordering mode
	// affects request and response delivery ordering and UPLI (UALink Protocol
	// Level Interface) virtual channel handling.
	//
	// Version added: v1.18.0
	NonStrictOrderingModeEnabled bool
	// UPLIWatchdogTimerMicroseconds shall contain time limit, in microseconds,
	// before the UPLI watchdog timer expires, taking the UPLI originator of the
	// port into isolation mode.
	//
	// Version added: v1.18.0
	UPLIWatchdogTimerMicroseconds *uint `json:",omitempty"`
}
