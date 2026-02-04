//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/IOConnectivityLineOfService.v1_2_1.json
// 1.1.0 - #IOConnectivityLineOfService.v1_2_1.IOConnectivityLineOfService

package schemas

import (
	"encoding/json"
)

// IOConnectivityLineOfService is an IO connectivity service option may be used
// to specify the characteristics of storage connectivity.
type IOConnectivityLineOfService struct {
	Entity
	// AccessProtocols shall specify the Access protocol for this service option.
	// NOTE: If multiple protocols are specified, the corresponding
	// MaxSupportedIOPS governs the max achieved across all protocol uses. This may
	// be less than the sum of the individual max values, which may be specified by
	// individual Line of Service entries.
	AccessProtocols []Protocol
	// MaxBytesPerSecond shall be the maximum bytes per second that a connection
	// can support.
	//
	// Version added: v1.1.0
	MaxBytesPerSecond *int `json:",omitempty"`
	// MaxIOPS shall be the maximum IOs per second that the connection shall allow
	// for the selected access protocol.
	//
	// Version added: v1.1.0
	MaxIOPS *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a IOConnectivityLineOfService object from the raw JSON.
func (i *IOConnectivityLineOfService) UnmarshalJSON(b []byte) error {
	type temp IOConnectivityLineOfService
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*i = IOConnectivityLineOfService(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *IOConnectivityLineOfService) Update() error {
	readWriteFields := []string{
		"AccessProtocols",
		"MaxBytesPerSecond",
		"MaxIOPS",
	}

	return i.UpdateFromRawData(i, i.RawData, readWriteFields)
}

// GetIOConnectivityLineOfService will get a IOConnectivityLineOfService instance from the service.
func GetIOConnectivityLineOfService(c Client, uri string) (*IOConnectivityLineOfService, error) {
	return GetObject[IOConnectivityLineOfService](c, uri)
}

// ListReferencedIOConnectivityLineOfServices gets the collection of IOConnectivityLineOfService from
// a provided reference.
func ListReferencedIOConnectivityLineOfServices(c Client, link string) ([]*IOConnectivityLineOfService, error) {
	return GetCollectionObjects[IOConnectivityLineOfService](c, link)
}
