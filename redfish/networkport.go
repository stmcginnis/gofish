//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// FlowControl is type of flow control for the port.
type FlowControl string

const (
	// NoneFlowControl No IEEE 802.3x flow control is enabled on this port.
	NoneFlowControl FlowControl = "None"
	// TXFlowControl IEEE 802.3x flow control may be initiated by this
	// station.
	TXFlowControl FlowControl = "TX"
	// RXFlowControl IEEE 802.3x flow control may be initiated by the link
	// partner.
	RXFlowControl FlowControl = "RX"
	// TXRXFlowControl IEEE 802.3x flow control may be initiated by this
	// station or the link partner.
	TXRXFlowControl FlowControl = "TX_RX"
)

// LinkNetworkTechnology is the link type.
type LinkNetworkTechnology string

const (
	// EthernetLinkNetworkTechnology means the port is capable of connecting to an
	// Ethernet network.
	EthernetLinkNetworkTechnology LinkNetworkTechnology = "Ethernet"
	// InfiniBandLinkNetworkTechnology means the port is capable of connecting to
	// an InfiniBand network.
	InfiniBandLinkNetworkTechnology LinkNetworkTechnology = "InfiniBand"
	// FibreChannelLinkNetworkTechnology means the port is capable of connecting to
	// a Fibre Channel network.
	FibreChannelLinkNetworkTechnology LinkNetworkTechnology = "FibreChannel"
)

// PortConnectionType is
type PortConnectionType string

const (
	// NotConnectedPortConnectionType means this port is not connected.
	NotConnectedPortConnectionType PortConnectionType = "NotConnected"
	// NPortPortConnectionType means this port connects via an N-Port to a switch.
	NPortPortConnectionType PortConnectionType = "NPort"
	// PointToPointPortConnectionType means this port connects in a Point-to-point
	// configuration.
	PointToPointPortConnectionType PortConnectionType = "PointToPoint"
	// PrivateLoopPortConnectionType means this port connects in a private loop
	// configuration.
	PrivateLoopPortConnectionType PortConnectionType = "PrivateLoop"
	// PublicLoopPortConnectionType means this port connects in a public
	// configuration.
	PublicLoopPortConnectionType PortConnectionType = "PublicLoop"
	// GenericPortConnectionType means this port connection type is a generic
	// fabric port.
	GenericPortConnectionType PortConnectionType = "Generic"
	// ExtenderFabricPortConnectionType means this port connection type is an
	// extender fabric port.
	ExtenderFabricPortConnectionType PortConnectionType = "ExtenderFabric"
)

// SupportedEthernetCapabilities is
type SupportedEthernetCapabilities string

const (
	// WakeOnLANSupportedEthernetCapabilities Wake on LAN (WoL) is supported
	// on this port.
	WakeOnLANSupportedEthernetCapabilities SupportedEthernetCapabilities = "WakeOnLAN"
	// EEESupportedEthernetCapabilities IEEE 802.3az Energy Efficient
	// Ethernet (EEE) is supported on this port.
	EEESupportedEthernetCapabilities SupportedEthernetCapabilities = "EEE"
)

// NetDevFuncMaxBWAlloc shall describe a maximum bandwidth percentage allocation
// for a network device function associated with a port.
type NetDevFuncMaxBWAlloc struct {
	// MaxBWAllocPercent shall be the maximum bandwidth percentage allocation
	// for the associated network device function.
	MaxBWAllocPercent int
	// NetworkDeviceFunction shall be a reference of type NetworkDeviceFunction
	// that represents the Network Device Function associated with this
	// bandwidth setting of this Network Port.
	NetworkDeviceFunction NetworkDeviceFunction
}

// NetDevFuncMinBWAlloc shall describe a minimum bandwidth percentage allocation
// for a network device function associated with a port.
type NetDevFuncMinBWAlloc struct {
	// MinBWAllocPercent shall be the minimum bandwidth percentage allocation
	// for the associated network device function. The sum total of all minimum
	// percentages shall not exceed 100.
	MinBWAllocPercent int
	// NetworkDeviceFunction shall be a reference of type NetworkDeviceFunction
	// that represents the Network Device Function associated with this
	// bandwidth setting of this Network Port.
	NetworkDeviceFunction NetworkDeviceFunction
}

// PortLinkStatus is the port link status.
type PortLinkStatus string

const (
	// UpPortLinkStatus The port is enabled and link is good (up).
	UpPortLinkStatus PortLinkStatus = "Up"
	// DownPortLinkStatus  The port is enabled but link is down.
	DownPortLinkStatus PortLinkStatus = "Down"
)

// NetworkPort represents a discrete physical port capable of connecting to a
// network.
type NetworkPort struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions shall contain the available actions for this resource.
	// actions string
	// ActiveLinkTechnology shall be the
	// configured link technology of this port.
	ActiveLinkTechnology LinkNetworkTechnology
	// AssociatedNetworkAddresses shall be an array of configured network
	// addresses that are associated with this network port, including the
	// programmed address of the lowest numbered network device function, the
	// configured but not active address if applicable, the address for hardware
	// port teaming, or other network addresses.
	AssociatedNetworkAddresses []string
	// CurrentLinkSpeedMbps shall be the current configured link speed of this
	// port.
	CurrentLinkSpeedMbps int
	// Description provides a description of this resource.
	Description string
	// EEEEnabled shall be a boolean indicating whether IEEE 802.3az Energy
	// Efficient Ethernet (EEE) is enabled for this network port.
	EEEEnabled bool
	// FCFabricName shall indicate the FC Fabric Name provided by the switch.
	FCFabricName string
	// FCPortConnectionType shall be the connection type for this port.
	FCPortConnectionType PortConnectionType
	// FlowControlConfiguration shall be the
	// locally configured 802.3x flow control setting for this network port.
	FlowControlConfiguration FlowControl
	// FlowControlStatus shall be the 802.3x
	// flow control behavior negotiated with the link partner for this
	// network port (Ethernet-only).
	FlowControlStatus FlowControl
	// LinkStatus shall be the link status between this port and its link
	// partner.
	LinkStatus PortLinkStatus
	// MaxFrameSize shall be the maximum frame size supported by the port.
	MaxFrameSize int
	// NetDevFuncMaxBWAlloc shall be an array of maximum bandwidth allocation
	// percentages for the Network Device Functions associated with this port.
	NetDevFuncMaxBWAlloc []NetDevFuncMaxBWAlloc
	// NetDevFuncMinBWAlloc shall be an array of minimum bandwidth percentage
	// allocations for each of the network device functions associated with this
	// port.
	NetDevFuncMinBWAlloc []NetDevFuncMinBWAlloc
	// NumberDiscoveredRemotePorts shall be the number of ports not on this
	// adapter that this port has discovered.
	NumberDiscoveredRemotePorts int
	// PhysicalPortNumber shall be the physical port number on the network
	// adapter hardware that this Network Port corresponds to. This value should
	// match a value visible on the hardware. When HostPortEnabled and
	// ManagementPortEnabled are both "false", the port shall not establish
	// physical link.
	PhysicalPortNumber string
	// PortMaximumMTU shall be the largest maximum transmission unit (MTU) that
	// can be configured for this network port.
	PortMaximumMTU int
	// SignalDetected shall be a boolean
	// indicating whether the port has detected enough signal on enough lanes
	// to establish link.
	SignalDetected bool
	// Status shall contain any status or health properties
	// of the resource.
	Status common.Status
	// SupportedEthernetCapabilities shall be
	// an array of zero or more Ethernet capabilities supported by this port.
	SupportedEthernetCapabilities []SupportedEthernetCapabilities
	// SupportedLinkCapabilities is This object shall describe the static
	// capabilities of the port, irrespective of transient conditions such as
	// cabling, interface module presence, or remote link partner status or
	// configuration.
	SupportedLinkCapabilities []SupportedLinkCapabilities
	// VendorID shall indicate the Vendor Identification string information as
	// provided by the manufacturer of this port.
	VendorID string `json:"VendorId"`
	// WakeOnLANEnabled shall be a boolean
	// indicating whether Wake on LAN (WoL) is enabled for this network port.
	WakeOnLANEnabled bool
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a NetworkPort object from the raw JSON.
func (networkport *NetworkPort) UnmarshalJSON(b []byte) error {
	type temp NetworkPort
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*networkport = NetworkPort(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	networkport.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (networkport *NetworkPort) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(NetworkPort)
	err := original.UnmarshalJSON(networkport.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"ActiveLinkTechnology",
		"CurrentLinkSpeedMbps",
		"EEEEnabled",
		"FlowControlConfiguration",
		"WakeOnLANEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(networkport).Elem()

	return networkport.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetNetworkPort will get a NetworkPort instance from the service.
func GetNetworkPort(c common.Client, uri string) (*NetworkPort, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var networkport NetworkPort
	err = json.NewDecoder(resp.Body).Decode(&networkport)
	if err != nil {
		return nil, err
	}

	networkport.SetClient(c)
	return &networkport, nil
}

// ListReferencedNetworkPorts gets the collection of NetworkPort from
// a provided reference.
func ListReferencedNetworkPorts(c common.Client, link string) ([]*NetworkPort, error) {
	var result []*NetworkPort
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, networkportLink := range links.ItemLinks {
		networkport, err := GetNetworkPort(c, networkportLink)
		if err != nil {
			return result, err
		}
		result = append(result, networkport)
	}

	return result, nil
}

// SupportedLinkCapabilities shall describe the static capabilities of an
// associated port, irrespective of transient conditions such as cabling,
// interface module presence, or remote link partner status or configuration.
type SupportedLinkCapabilities struct {
	// AutoSpeedNegotiation shall be indicate whether the port is capable of
	// auto-negotiating speed.
	AutoSpeedNegotiation bool
	// CapableLinkSpeedMbps shall be all of the possible network link speed
	// capabilities of this port.
	CapableLinkSpeedMbps []int
	// LinkNetworkTechnology shall be a network technology capability of this
	// port.
	LinkNetworkTechnology LinkNetworkTechnology
}
