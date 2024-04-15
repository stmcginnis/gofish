//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// NetworkAdapterMetrics shall represent the network metrics for a single network adapter in a Redfish
// implementation.
type NetworkAdapterMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// CPUCorePercent shall contain the device CPU core utilization as a percentage, typically '0' to '100'.
	CPUCorePercent float64
	// Description provides a description of this resource.
	Description string
	// HostBusRXPercent shall contain the host bus, such as PCIe, RX utilization as a percentage, typically '0' to
	// '100', which is calculated by dividing the total bytes received by the theoretical max.
	HostBusRXPercent float64
	// HostBusTXPercent shall contain the host bus, such as PCIe, TX utilization as a percentage, typically '0' to
	// '100', which is calculated by dividing the total bytes transmitted by the theoretical max.
	HostBusTXPercent float64
	// NCSIRXBytes shall contain the total number of NC-SI bytes received since reset, including both passthrough and
	// non-passthrough traffic.
	NCSIRXBytes int
	// NCSIRXFrames shall contain the total number of NC-SI frames received since reset, including both passthrough and
	// non-passthrough traffic.
	NCSIRXFrames int
	// NCSITXBytes shall contain the total number of NC-SI bytes sent since reset, including both passthrough and non-
	// passthrough traffic.
	NCSITXBytes int
	// NCSITXFrames shall contain the total number of NC-SI frames sent since reset, including both passthrough and
	// non-passthrough traffic.
	NCSITXFrames int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RXBytes shall contain the total number of bytes received since reset, including host and remote management
	// passthrough traffic, and inclusive of all protocol overhead.
	RXBytes int
	// RXMulticastFrames shall contain the total number of good multicast frames received since reset.
	RXMulticastFrames int
	// RXUnicastFrames shall contain the total number of good unicast frames received since reset.
	RXUnicastFrames int
	// TXBytes shall contain the total number of bytes transmitted since reset, including host and remote management
	// passthrough traffic, and inclusive of all protocol overhead.
	TXBytes int
	// TXMulticastFrames shall contain the total number of good multicast frames transmitted since reset.
	TXMulticastFrames int
	// TXUnicastFrames shall contain the total number of good unicast frames transmitted since reset.
	TXUnicastFrames int
}

// GetNetworkAdapterMetrics will get a NetworkAdapterMetrics instance from the service.
func GetNetworkAdapterMetrics(c common.Client, uri string) (*NetworkAdapterMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var networkadaptermetrics NetworkAdapterMetrics
	err = json.NewDecoder(resp.Body).Decode(&networkadaptermetrics)
	if err != nil {
		return nil, err
	}

	networkadaptermetrics.SetClient(c)
	return &networkadaptermetrics, nil
}

// ListReferencedNetworkAdapterMetrics gets the collection of NetworkAdapterMetrics from
// a provided reference.
func ListReferencedNetworkAdapterMetrics(c common.Client, link string) ([]*NetworkAdapterMetrics, error) {
	var result []*NetworkAdapterMetrics
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *NetworkAdapterMetrics
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		networkadaptermetrics, err := GetNetworkAdapterMetrics(c, link)
		ch <- GetResult{Item: networkadaptermetrics, Link: link, Error: err}
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
