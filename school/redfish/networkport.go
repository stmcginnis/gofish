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
	// EthernetLinkNetworkTechnology The port is capable of connecting to an
	// Ethernet network.
	EthernetLinkNetworkTechnology LinkNetworkTechnology = "Ethernet"
	// InfiniBandLinkNetworkTechnology The port is capable of connecting to
	// an InfiniBand network.
	InfiniBandLinkNetworkTechnology LinkNetworkTechnology = "InfiniBand"
	// FibreChannelLinkNetworkTechnology The port is capable of connecting to
	// a Fibre Channel network.
	FibreChannelLinkNetworkTechnology LinkNetworkTechnology = "FibreChannel"
)

// PortConnectionType is
type PortConnectionType string

const (
	// NotConnectedPortConnectionType This port is not connected.
	NotConnectedPortConnectionType PortConnectionType = "NotConnected"
	// NPortPortConnectionType This port connects via an N-Port to a switch.
	NPortPortConnectionType PortConnectionType = "NPort"
	// PointToPointPortConnectionType This port connects in a Point-to-point
	// configuration.
	PointToPointPortConnectionType PortConnectionType = "PointToPoint"
	// PrivateLoopPortConnectionType This port connects in a private loop
	// configuration.
	PrivateLoopPortConnectionType PortConnectionType = "PrivateLoop"
	// PublicLoopPortConnectionType This port connects in a public
	// configuration.
	PublicLoopPortConnectionType PortConnectionType = "PublicLoop"
	// GenericPortConnectionType This port connection type is a generic
	// fabric port.
	GenericPortConnectionType PortConnectionType = "Generic"
	// ExtenderFabricPortConnectionType This port connection type is an
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

// NetDevFuncMaxBWAlloc is This type shall describe a maximum bandwidth
// percentage allocation for a network device function associated with a
// port.
type NetDevFuncMaxBWAlloc struct {
	common.Entity

	// MaxBWAllocPercent is The value of this property shall be the maximum
	// bandwidth percentage allocation for the associated network device
	// function.
	MaxBWAllocPercent int
	// NetworkDeviceFunction is The value of this property shall be a
	// reference of type NetworkDeviceFunction that represents the Network
	// Device Function associated with this bandwidth setting of this Network
	// Port.
	NetworkDeviceFunction string
}

// UnmarshalJSON unmarshals a NetDevFuncMaxBWAlloc object from the raw JSON.
func (netdevfuncmaxbwalloc *NetDevFuncMaxBWAlloc) UnmarshalJSON(b []byte) error {
	type temp NetDevFuncMaxBWAlloc
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*netdevfuncmaxbwalloc = NetDevFuncMaxBWAlloc(t.temp)

	// Extract the links to other entities for later

	return nil
}

// NetDevFuncMinBWAlloc is This type shall describe a minimum bandwidth
// percentage allocation for a network device function associated with a
// port.
type NetDevFuncMinBWAlloc struct {
	common.Entity

	// MinBWAllocPercent is The value of this property shall be the minimum
	// bandwidth percentage allocation for the associated network device
	// function. The sum total of all minimum percentages shall not exceed
	// 100.
	MinBWAllocPercent int
	// NetworkDeviceFunction is The value of this property shall be a
	// reference of type NetworkDeviceFunction that represents the Network
	// Device Function associated with this bandwidth setting of this Network
	// Port.
	NetworkDeviceFunction string
}

// UnmarshalJSON unmarshals a NetDevFuncMinBWAlloc object from the raw JSON.
func (netdevfuncminbwalloc *NetDevFuncMinBWAlloc) UnmarshalJSON(b []byte) error {
	type temp NetDevFuncMinBWAlloc
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*netdevfuncminbwalloc = NetDevFuncMinBWAlloc(t.temp)

	// Extract the links to other entities for later

	return nil
}

// NetworkPort is A Network Port represents a discrete physical port
// capable of connecting to a network.
type NetworkPort struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Actions is The Actions property shall contain the available actions
	// for this resource.
	Actions string
	// ActiveLinkTechnology is The value of this property shall be the
	// configured link technology of this port.
	ActiveLinkTechnology LinkNetworkTechnology
	// AssociatedNetworkAddresses is The value of this property shall be an
	// array of configured network addresses that are associated with this
	// network port, including the programmed address of the lowest numbered
	// network device function, the configured but not active address if
	// applicable, the address for hardware port teaming, or other network
	// addresses.
	AssociatedNetworkAddresses []string
	// CurrentLinkSpeedMbps is The value of this property shall be the
	// current configured link speed of this port.
	CurrentLinkSpeedMbps int
	// Description provides a description of this resource.
	Description string
	// EEEEnabled is The value of this property shall be a boolean indicating
	// whether IEEE 802.3az Energy Efficient Ethernet (EEE) is enabled for
	// this network port.
	EEEEnabled bool
	// FCFabricName is This property shall indicate the FC Fabric Name
	// provided by the switch.
	FCFabricName string
	// FCPortConnectionType is The value of this property shall be the
	// connection type for this port.
	FCPortConnectionType PortConnectionType
	// FlowControlConfiguration is The value of this property shall be the
	// locally configured 802.3x flow control setting for this network port.
	FlowControlConfiguration FlowControl
	// FlowControlStatus is The value of this property shall be the 802.3x
	// flow control behavior negotiated with the link partner for this
	// network port (Ethernet-only).
	FlowControlStatus FlowControl
	// LinkStatus is The value of this property shall be the link status
	// between this port and its link partner.
	LinkStatus LinkStatus
	// MaxFrameSize is The value of this property shall be the maximum frame
	// size supported by the port.
	MaxFrameSize int
	// NetDevFuncMaxBWAlloc is The value of this property shall be an array
	// of maximum bandwidth allocation percentages for the Network Device
	// Functions associated with this port.
	NetDevFuncMaxBWAlloc []NetDevFuncMaxBWAlloc
	// NetDevFuncMinBWAlloc is The value of this property shall be an array
	// of minimum bandwidth percentage allocations for each of the network
	// device functions associated with this port.
	NetDevFuncMinBWAlloc []NetDevFuncMinBWAlloc
	// NumberDiscoveredRemotePorts is The value of this property shall be the
	// number of ports not on this adapter that this port has discovered.
	NumberDiscoveredRemotePorts int
	// Oem is The value of this string shall be of the format for the
	// reserved word *Oem*.
	OEM string `json:"Oem"`
	// PhysicalPortNumber is The value of this property shall be the physical
	// port number on the network adapter hardware that this Network Port
	// corresponds to. This value should match a value visible on the
	// hardware. When HostPortEnabled and ManagementPortEnabled are both
	// "false", the port shall not establish physical link.
	PhysicalPortNumber string
	// PortMaximumMTU is The value of this property shall be the largest
	// maximum transmission unit (MTU) that can be configured for this
	// network port.
	PortMaximumMTU int
	// SignalDetected is The value of this property shall be a boolean
	// indicating whether the port has detected enough signal on enough lanes
	// to establish link.
	SignalDetected bool
	// Status is This property shall contain any status or health properties
	// of the resource.
	Status common.Status
	// SupportedEthernetCapabilities is The value of this property shall be
	// an array of zero or more Ethernet capabilities supported by this port.
	SupportedEthernetCapabilities []SupportedEthernetCapabilities
	// SupportedLinkCapabilities is This object shall describe the static
	// capabilities of the port, irrespective of transient conditions such as
	// cabling, interface module presence, or remote link parter status or
	// configuration.
	SupportedLinkCapabilities []SupportedLinkCapabilities
	// VendorID shall indicate the Vendor Identification string information as
	// provided by the manufacturer of this port.
	VendorID string `json:"VendorId"`
	// WakeOnLANEnabled is The value of this property shall be a boolean
	// indicating whether Wake on LAN (WoL) is enabled for this network port.
	WakeOnLANEnabled bool
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

	// Extract the links to other entities for later

	return nil
}

// SupportedLinkCapabilities is This type shall describe the static
// capabilities of an associated port, irrespective of transient
// conditions such as cabling, interface module presence, or remote link
// parter status or configuration.
type SupportedLinkCapabilities struct {
	common.Entity

	// AutoSpeedNegotiation is The value of this property shall be indicate
	// whether the port is capable of auto-negotiating speed.
	AutoSpeedNegotiation bool
	// CapableLinkSpeedMbps is The value of this property shall be all of the
	// possible network link speed capabilities of this port.
	CapableLinkSpeedMbps []string
	// LinkNetworkTechnology is The value of this property shall be a network
	// technology capability of this port.
	LinkNetworkTechnology LinkNetworkTechnology
}

// UnmarshalJSON unmarshals a SupportedLinkCapabilities object from the raw JSON.
func (supportedlinkcapabilities *SupportedLinkCapabilities) UnmarshalJSON(b []byte) error {
	type temp SupportedLinkCapabilities
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*supportedlinkcapabilities = SupportedLinkCapabilities(t.temp)

	// Extract the links to other entities for later

	return nil
}
