//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.4 - #NetworkPort.v1_4_3.NetworkPort

package schemas

import (
	"encoding/json"
)

type FlowControl string

const (
	// NoneFlowControl No IEEE 802.3x flow control is enabled on this port.
	NoneFlowControl FlowControl = "None"
	// TXFlowControl This station can initiate IEEE 802.3x flow control.
	TXFlowControl FlowControl = "TX"
	// RXFlowControl The link partner can initiate IEEE 802.3x flow control.
	RXFlowControl FlowControl = "RX"
	// TXRXFlowControl This station or the link partner can initiate IEEE 802.3x
	// flow control.
	TXRXFlowControl FlowControl = "TX_RX"
)

type NetworkPortLinkStatus string

const (
	// DownNetworkPortLinkStatus The port is enabled but link is down.
	DownNetworkPortLinkStatus NetworkPortLinkStatus = "Down"
	// UpNetworkPortLinkStatus The port is enabled and link is good (up).
	UpNetworkPortLinkStatus NetworkPortLinkStatus = "Up"
	// StartingNetworkPortLinkStatus This link on this interface is starting. A physical link
	// has been established, but the port is not able to transfer data.
	StartingNetworkPortLinkStatus NetworkPortLinkStatus = "Starting"
	// TrainingNetworkPortLinkStatus This physical link on this interface is training.
	TrainingNetworkPortLinkStatus NetworkPortLinkStatus = "Training"
)

type NetworkPortConnectionType string

const (
	// NotConnectedNetworkPortConnectionType This port is not connected.
	NotConnectedNetworkPortConnectionType NetworkPortConnectionType = "NotConnected"
	// NPortNetworkPortConnectionType This port connects through an N-port to a switch.
	NPortNetworkPortConnectionType NetworkPortConnectionType = "NPort"
	// PointToPointNetworkPortConnectionType This port connects in a point-to-point
	// configuration.
	PointToPointNetworkPortConnectionType NetworkPortConnectionType = "PointToPoint"
	// PrivateLoopNetworkPortConnectionType This port connects in a private loop
	// configuration.
	PrivateLoopNetworkPortConnectionType NetworkPortConnectionType = "PrivateLoop"
	// PublicLoopNetworkPortConnectionType This port connects in a public configuration.
	PublicLoopNetworkPortConnectionType NetworkPortConnectionType = "PublicLoop"
	// GenericNetworkPortConnectionType This port connection type is a generic fabric
	// port.
	GenericNetworkPortConnectionType NetworkPortConnectionType = "Generic"
	// ExtenderFabricNetworkPortConnectionType This port connection type is an extender
	// fabric port.
	ExtenderFabricNetworkPortConnectionType NetworkPortConnectionType = "ExtenderFabric"
)

// NetworkPort shall represent a discrete physical port that can connect to a
// network in a Redfish implementation.
type NetworkPort struct {
	Entity
	// ActiveLinkTechnology shall contain the configured link technology of this
	// port.
	ActiveLinkTechnology LinkNetworkTechnology
	// AssociatedNetworkAddresses shall contain an array of configured network
	// addresses that are associated with this network port, including the
	// programmed address of the lowest-numbered network device function, the
	// configured but not active address if applicable, the address for hardware
	// port teaming, or other network addresses.
	AssociatedNetworkAddresses []string
	// CurrentLinkSpeedMbps shall contain the current configured link speed of this
	// port.
	//
	// Version added: v1.2.0
	CurrentLinkSpeedMbps *int `json:",omitempty"`
	// EEEEnabled shall indicate whether IEEE 802.3az Energy-Efficient Ethernet
	// (EEE) is enabled for this network port.
	EEEEnabled bool
	// FCFabricName shall indicate the FC Fabric Name provided by the switch.
	//
	// Version added: v1.2.0
	FCFabricName string
	// FCPortConnectionType shall contain the connection type for this port.
	//
	// Version added: v1.2.0
	FCPortConnectionType NetworkPortConnectionType
	// FlowControlConfiguration shall contain the locally configured 802.3x flow
	// control setting for this network port.
	FlowControlConfiguration FlowControl
	// FlowControlStatus shall contain the 802.3x flow control behavior negotiated
	// with the link partner for this network port (Ethernet-only).
	FlowControlStatus FlowControl
	// LinkStatus shall contain the link status between this port and its link
	// partner.
	LinkStatus NetworkPortLinkStatus
	// MaxFrameSize shall contain the maximum frame size supported by the port.
	//
	// Version added: v1.2.0
	MaxFrameSize *int `json:",omitempty"`
	// NetDevFuncMaxBWAlloc shall contain an array of maximum bandwidth allocation
	// percentages for the network device functions associated with this port.
	NetDevFuncMaxBWAlloc []NetDevFuncMaxBWAlloc
	// NetDevFuncMinBWAlloc shall contain an array of minimum bandwidth percentage
	// allocations for each of the network device functions associated with this
	// port.
	NetDevFuncMinBWAlloc []NetDevFuncMinBWAlloc
	// NumberDiscoveredRemotePorts shall contain the number of ports not on this
	// adapter that this port has discovered.
	//
	// Version added: v1.2.0
	NumberDiscoveredRemotePorts *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PhysicalPortNumber shall contain the physical port number on the network
	// adapter hardware that this network port corresponds to. This value should
	// match a value visible on the hardware.
	PhysicalPortNumber string
	// PortMaximumMTU shall contain the largest maximum transmission unit (MTU)
	// that can be configured for this network port.
	PortMaximumMTU *int `json:",omitempty"`
	// SignalDetected shall indicate whether the port has detected enough signal on
	// enough lanes to establish a link.
	SignalDetected bool
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SupportedEthernetCapabilities shall contain an array of zero or more
	// Ethernet capabilities supported by this port.
	SupportedEthernetCapabilities []SupportedEthernetCapabilities
	// SupportedLinkCapabilities shall describe the static capabilities of the
	// port, irrespective of transient conditions such as cabling, interface module
	// presence, or remote link partner status or configuration.
	SupportedLinkCapabilities []SupportedLinkCapabilities
	// VendorID shall indicate the vendor identification string information as
	// provided by the manufacturer of this port.
	//
	// Version added: v1.2.0
	VendorID string `json:"VendorId"`
	// WakeOnLANEnabled shall indicate whether Wake on LAN (WoL) is enabled for
	// this network port.
	WakeOnLANEnabled bool
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a NetworkPort object from the raw JSON.
func (n *NetworkPort) UnmarshalJSON(b []byte) error {
	type temp NetworkPort
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*n = NetworkPort(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	n.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (n *NetworkPort) Update() error {
	readWriteFields := []string{
		"ActiveLinkTechnology",
		"CurrentLinkSpeedMbps",
		"EEEEnabled",
		"FlowControlConfiguration",
		"WakeOnLANEnabled",
	}

	return n.UpdateFromRawData(n, n.RawData, readWriteFields)
}

// GetNetworkPort will get a NetworkPort instance from the service.
func GetNetworkPort(c Client, uri string) (*NetworkPort, error) {
	return GetObject[NetworkPort](c, uri)
}

// ListReferencedNetworkPorts gets the collection of NetworkPort from
// a provided reference.
func ListReferencedNetworkPorts(c Client, link string) ([]*NetworkPort, error) {
	return GetCollectionObjects[NetworkPort](c, link)
}

// NetDevFuncMaxBWAlloc shall describe a maximum bandwidth percentage allocation
// for a network device function associated with a port.
type NetDevFuncMaxBWAlloc struct {
	// MaxBWAllocPercent shall contain the maximum bandwidth percentage allocation
	// for the associated network device function.
	MaxBWAllocPercent *int `json:",omitempty"`
	// NetworkDeviceFunction shall contain a link to a resource of type
	// 'NetworkDeviceFunction' that represents the network device function
	// associated with this bandwidth setting of this network port.
	networkDeviceFunction string
}

// UnmarshalJSON unmarshals a NetDevFuncMaxBWAlloc object from the raw JSON.
func (n *NetDevFuncMaxBWAlloc) UnmarshalJSON(b []byte) error {
	type temp NetDevFuncMaxBWAlloc
	var tmp struct {
		temp
		NetworkDeviceFunction Link `json:"NetworkDeviceFunction"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*n = NetDevFuncMaxBWAlloc(tmp.temp)

	// Extract the links to other entities for later
	n.networkDeviceFunction = tmp.NetworkDeviceFunction.String()

	return nil
}

// NetworkDeviceFunction gets the NetworkDeviceFunction linked resource.
func (n *NetDevFuncMaxBWAlloc) NetworkDeviceFunction(client Client) (*NetworkDeviceFunction, error) {
	if n.networkDeviceFunction == "" {
		return nil, nil
	}
	return GetObject[NetworkDeviceFunction](client, n.networkDeviceFunction)
}

// NetDevFuncMinBWAlloc shall describe a minimum bandwidth percentage allocation
// for a network device function associated with a port.
type NetDevFuncMinBWAlloc struct {
	// MinBWAllocPercent shall contain the minimum bandwidth percentage allocation
	// for the associated network device function. The sum total of all minimum
	// percentages shall not exceed 100.
	MinBWAllocPercent *int `json:",omitempty"`
	// NetworkDeviceFunction shall contain a link to a resource of type
	// 'NetworkDeviceFunction' that represents the network device function
	// associated with this bandwidth setting of this network port.
	networkDeviceFunction string
}

// UnmarshalJSON unmarshals a NetDevFuncMinBWAlloc object from the raw JSON.
func (n *NetDevFuncMinBWAlloc) UnmarshalJSON(b []byte) error {
	type temp NetDevFuncMinBWAlloc
	var tmp struct {
		temp
		NetworkDeviceFunction Link `json:"NetworkDeviceFunction"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*n = NetDevFuncMinBWAlloc(tmp.temp)

	// Extract the links to other entities for later
	n.networkDeviceFunction = tmp.NetworkDeviceFunction.String()

	return nil
}

// NetworkDeviceFunction gets the NetworkDeviceFunction linked resource.
func (n *NetDevFuncMinBWAlloc) NetworkDeviceFunction(client Client) (*NetworkDeviceFunction, error) {
	if n.networkDeviceFunction == "" {
		return nil, nil
	}
	return GetObject[NetworkDeviceFunction](client, n.networkDeviceFunction)
}

// SupportedLinkCapabilities shall describe the static capabilities of an
// associated port, irrespective of transient conditions such as cabling,
// interface module presence, or remote link partner status or configuration.
type SupportedLinkCapabilities struct {
	// AutoSpeedNegotiation shall indicate whether the port is capable of
	// autonegotiating speed.
	//
	// Version added: v1.2.0
	AutoSpeedNegotiation bool
	// CapableLinkSpeedMbps shall contain all of the possible network link speed
	// capabilities of this port.
	//
	// Version added: v1.2.0
	CapableLinkSpeedMbps []*int
	// LinkNetworkTechnology shall contain a network technology capability of this
	// port.
	LinkNetworkTechnology LinkNetworkTechnology
	// LinkSpeedMbps shall contain the speed of the link in megabits per second
	// (Mbit/s) units for this port when this link network technology is active.
	//
	// Deprecated: v1.2.0
	// This property has been deprecated in favor of the 'CapableLinkSpeedMbps'
	// property.
	LinkSpeedMbps *int `json:",omitempty"`
}
