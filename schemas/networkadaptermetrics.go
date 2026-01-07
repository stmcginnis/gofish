//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2024.1 - #NetworkAdapterMetrics.v1_1_0.NetworkAdapterMetrics

package schemas

import (
	"encoding/json"
)

// NetworkAdapterMetrics shall represent the network metrics for a single
// network adapter in a Redfish implementation.
type NetworkAdapterMetrics struct {
	Entity
	// CPUCorePercent shall contain the device CPU core utilization as a
	// percentage, typically '0' to '100'.
	CPUCorePercent *float64 `json:",omitempty"`
	// HostBusRXPercent shall contain the host bus, such as PCIe, RX utilization as
	// a percentage, typically '0' to '100', which is calculated by dividing the
	// total bytes received by the theoretical max.
	HostBusRXPercent *float64 `json:",omitempty"`
	// HostBusTXPercent shall contain the host bus, such as PCIe, TX utilization as
	// a percentage, typically '0' to '100', which is calculated by dividing the
	// total bytes transmitted by the theoretical max.
	HostBusTXPercent *float64 `json:",omitempty"`
	// NCSIRXBytes shall contain the total number of NC-SI bytes received since
	// reset, including both passthrough and non-passthrough traffic.
	NCSIRXBytes *int `json:",omitempty"`
	// NCSIRXFrames shall contain the total number of NC-SI frames received since
	// reset, including both passthrough and non-passthrough traffic.
	NCSIRXFrames *int `json:",omitempty"`
	// NCSITXBytes shall contain the total number of NC-SI bytes sent since reset,
	// including both passthrough and non-passthrough traffic.
	NCSITXBytes *int `json:",omitempty"`
	// NCSITXFrames shall contain the total number of NC-SI frames sent since
	// reset, including both passthrough and non-passthrough traffic.
	NCSITXFrames *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RXBytes shall contain the total number of bytes received since reset,
	// including host and remote management passthrough traffic, and inclusive of
	// all protocol overhead.
	RXBytes *int `json:",omitempty"`
	// RXMulticastFrames shall contain the total number of good multicast frames
	// received since reset.
	RXMulticastFrames *int `json:",omitempty"`
	// RXUnicastFrames shall contain the total number of good unicast frames
	// received since reset.
	RXUnicastFrames *int `json:",omitempty"`
	// TXBytes shall contain the total number of bytes transmitted since reset,
	// including host and remote management passthrough traffic, and inclusive of
	// all protocol overhead.
	TXBytes *int `json:",omitempty"`
	// TXMulticastFrames shall contain the total number of good multicast frames
	// transmitted since reset.
	TXMulticastFrames *int `json:",omitempty"`
	// TXUnicastFrames shall contain the total number of good unicast frames
	// transmitted since reset.
	TXUnicastFrames *int `json:",omitempty"`
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
}

// UnmarshalJSON unmarshals a NetworkAdapterMetrics object from the raw JSON.
func (n *NetworkAdapterMetrics) UnmarshalJSON(b []byte) error {
	type temp NetworkAdapterMetrics
	type nActions struct {
		ResetMetrics ActionTarget `json:"#NetworkAdapterMetrics.ResetMetrics"`
	}
	var tmp struct {
		temp
		Actions nActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*n = NetworkAdapterMetrics(tmp.temp)

	// Extract the links to other entities for later
	n.resetMetricsTarget = tmp.Actions.ResetMetrics.Target

	return nil
}

// GetNetworkAdapterMetrics will get a NetworkAdapterMetrics instance from the service.
func GetNetworkAdapterMetrics(c Client, uri string) (*NetworkAdapterMetrics, error) {
	return GetObject[NetworkAdapterMetrics](c, uri)
}

// ListReferencedNetworkAdapterMetricss gets the collection of NetworkAdapterMetrics from
// a provided reference.
func ListReferencedNetworkAdapterMetricss(c Client, link string) ([]*NetworkAdapterMetrics, error) {
	return GetCollectionObjects[NetworkAdapterMetrics](c, link)
}

// This action shall reset any time intervals or counted values for this
// device.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (n *NetworkAdapterMetrics) ResetMetrics() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(n.client,
		n.resetMetricsTarget, payload, n.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}
