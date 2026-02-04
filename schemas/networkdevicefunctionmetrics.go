//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/NetworkDeviceFunctionMetrics.v1_2_0.json
// 2024.1 - #NetworkDeviceFunctionMetrics.v1_2_0.NetworkDeviceFunctionMetrics

package schemas

import (
	"encoding/json"
)

// NetworkDeviceFunctionMetrics shall represent the network metrics for a single
// network function of a network adapter in a Redfish implementation.
type NetworkDeviceFunctionMetrics struct {
	Entity
	// Ethernet shall contain network function metrics specific to Ethernet
	// adapters.
	Ethernet NetworkDeviceFunctionMetricsEthernet
	// FibreChannel shall contain network function metrics specific to Fibre
	// Channel adapters.
	//
	// Version added: v1.1.0
	FibreChannel NetworkDeviceFunctionMetricsFibreChannel
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RXAvgQueueDepthPercent shall contain the average RX queue depth as a
	// percentage, typically '0' to '100'.
	RXAvgQueueDepthPercent *float64 `json:",omitempty"`
	// RXBytes shall contain the total number of bytes received on a network
	// function, inclusive of all protocol overhead.
	RXBytes *int `json:",omitempty"`
	// RXFrames shall contain the total number of frames received on a network
	// function.
	RXFrames *int `json:",omitempty"`
	// RXMulticastFrames shall contain the total number of good multicast frames
	// received on a network function since reset, including host and remote
	// management passthrough traffic.
	RXMulticastFrames *int `json:",omitempty"`
	// RXQueuesEmpty shall indicate whether nothing is in a network function's RX
	// queues to DMA.
	RXQueuesEmpty bool
	// RXQueuesFull shall contain the number of RX queues that are full.
	RXQueuesFull *int `json:",omitempty"`
	// RXUnicastFrames shall contain the total number of good unicast frames
	// received on a network function since reset.
	RXUnicastFrames *int `json:",omitempty"`
	// TXAvgQueueDepthPercent shall contain the average TX queue depth as a
	// percentage, typically '0' to '100'.
	TXAvgQueueDepthPercent *float64 `json:",omitempty"`
	// TXBytes shall contain the total number of bytes sent on a network function,
	// inclusive of all protocol overhead.
	TXBytes *int `json:",omitempty"`
	// TXFrames shall contain the total number of frames sent on a network
	// function.
	TXFrames *int `json:",omitempty"`
	// TXMulticastFrames shall contain the total number of good multicast frames
	// transmitted on a network function since reset, including host and remote
	// management passthrough traffic.
	TXMulticastFrames *int `json:",omitempty"`
	// TXQueuesEmpty shall indicate whether all TX queues for a network function
	// are empty.
	TXQueuesEmpty bool
	// TXQueuesFull shall contain the number of TX queues that are full.
	TXQueuesFull *int `json:",omitempty"`
	// TXUnicastFrames shall contain the total number of good unicast frames
	// transmitted on a network function since reset, including host and remote
	// management passthrough traffic.
	TXUnicastFrames *int `json:",omitempty"`
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
}

// UnmarshalJSON unmarshals a NetworkDeviceFunctionMetrics object from the raw JSON.
func (n *NetworkDeviceFunctionMetrics) UnmarshalJSON(b []byte) error {
	type temp NetworkDeviceFunctionMetrics
	type nActions struct {
		ResetMetrics ActionTarget `json:"#NetworkDeviceFunctionMetrics.ResetMetrics"`
	}
	var tmp struct {
		temp
		Actions nActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*n = NetworkDeviceFunctionMetrics(tmp.temp)

	// Extract the links to other entities for later
	n.resetMetricsTarget = tmp.Actions.ResetMetrics.Target

	return nil
}

// GetNetworkDeviceFunctionMetrics will get a NetworkDeviceFunctionMetrics instance from the service.
func GetNetworkDeviceFunctionMetrics(c Client, uri string) (*NetworkDeviceFunctionMetrics, error) {
	return GetObject[NetworkDeviceFunctionMetrics](c, uri)
}

// ListReferencedNetworkDeviceFunctionMetricss gets the collection of NetworkDeviceFunctionMetrics from
// a provided reference.
func ListReferencedNetworkDeviceFunctionMetricss(c Client, link string) ([]*NetworkDeviceFunctionMetrics, error) {
	return GetCollectionObjects[NetworkDeviceFunctionMetrics](c, link)
}

// This action shall reset any time intervals or counted values for this
// device.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (n *NetworkDeviceFunctionMetrics) ResetMetrics() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(n.client,
		n.resetMetricsTarget, payload, n.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// NetworkDeviceFunctionMetricsEthernet shall describe the Ethernet-related network function metrics.
type NetworkDeviceFunctionMetricsEthernet struct {
	// NumOffloadedIPv4Conns shall contain the total number of offloaded TCP/IPv4
	// connections.
	NumOffloadedIPv4Conns *int `json:",omitempty"`
	// NumOffloadedIPv6Conns shall contain the total number of offloaded TCP/IPv6
	// connections.
	NumOffloadedIPv6Conns *int `json:",omitempty"`
}

// NetworkDeviceFunctionMetricsFibreChannel shall describe the Fibre Channel-related network function
// metrics.
type NetworkDeviceFunctionMetricsFibreChannel struct {
	// PortLoginAccepts shall contain the total number of PLOGI ACC responses
	// received by this Fibre Channel function.
	//
	// Version added: v1.1.0
	PortLoginAccepts *int `json:",omitempty"`
	// PortLoginRejects shall contain the total number of PLOGI RJT responses
	// received by this Fibre Channel function.
	//
	// Version added: v1.1.0
	PortLoginRejects *int `json:",omitempty"`
	// PortLoginRequests shall contain the total number of PLOGI requests sent by
	// this function.
	//
	// Version added: v1.1.0
	PortLoginRequests *int `json:",omitempty"`
	// RXCongestionFPINs shall contain the total number of Congestion FPINs
	// received by this Fibre Channel function.
	//
	// Version added: v1.1.0
	RXCongestionFPINs *int `json:",omitempty"`
	// RXDeliveryFPINs shall contain the total number of Delivery FPINs received by
	// this Fibre Channel function.
	//
	// Version added: v1.1.0
	RXDeliveryFPINs *int `json:",omitempty"`
	// RXExchanges shall contain the total number of Fibre Channel exchanges
	// received.
	//
	// Version added: v1.1.0
	RXExchanges *int `json:",omitempty"`
	// RXLinkIntegrityFPINs shall contain the total number of Link Integrity FPINs
	// received by this Fibre Channel function.
	//
	// Version added: v1.1.0
	RXLinkIntegrityFPINs *int `json:",omitempty"`
	// RXPeerCongestionFPINs shall contain the total number of Peer Congestion
	// FPINs received by this Fibre Channel function.
	//
	// Version added: v1.1.0
	RXPeerCongestionFPINs *int `json:",omitempty"`
	// RXSequences shall contain the total number of Fibre Channel sequences
	// received.
	//
	// Version added: v1.1.0
	RXSequences *int `json:",omitempty"`
	// TXCongestionFPINs shall contain the total number of Congestion FPINs sent by
	// this Fibre Channel function.
	//
	// Version added: v1.1.0
	TXCongestionFPINs *int `json:",omitempty"`
	// TXDeliveryFPINs shall contain the total number of Delivery FPINs sent by
	// this Fibre Channel function.
	//
	// Version added: v1.1.0
	TXDeliveryFPINs *int `json:",omitempty"`
	// TXExchanges shall contain the total number of Fibre Channel exchanges
	// transmitted.
	//
	// Version added: v1.1.0
	TXExchanges *int `json:",omitempty"`
	// TXLinkIntegrityFPINs shall contain the total number of Link Integrity FPINs
	// sent by this Fibre Channel function.
	//
	// Version added: v1.1.0
	TXLinkIntegrityFPINs *int `json:",omitempty"`
	// TXPeerCongestionFPINs shall contain the total number of Peer Congestion
	// FPINs sent by this Fibre Channel function.
	//
	// Version added: v1.1.0
	TXPeerCongestionFPINs *int `json:",omitempty"`
	// TXSequences shall contain the total number of Fibre Channel sequences
	// transmitted.
	//
	// Version added: v1.1.0
	TXSequences *int `json:",omitempty"`
}
