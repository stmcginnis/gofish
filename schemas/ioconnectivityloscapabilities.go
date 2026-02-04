//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/swordfish/v1/IOConnectivityLoSCapabilities.v1_2_0.json
// 1.2.1c - #IOConnectivityLoSCapabilities.v1_2_0.IOConnectivityLoSCapabilities

package schemas

import (
	"encoding/json"
)

// IOConnectivityLoSCapabilities Each instance of IOConnectivityLoSCapabilities
// describes capabilities of the system to support various IO Connectivity
// service options.
type IOConnectivityLoSCapabilities struct {
	Entity
	// Identifier shall be unique within the managed ecosystem.
	Identifier Identifier
	// MaxSupportedBytesPerSecond shall be the maximum bytes per second that a
	// connection can support.
	MaxSupportedBytesPerSecond *int `json:",omitempty"`
	// MaxSupportedIOPS shall be the maximum IOPS that a connection can support.
	//
	// Version added: v1.1.0
	MaxSupportedIOPS *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// SupportedAccessProtocols Access protocols supported by this service option.
	// NOTE: SMB+NFS* requires that SMB and at least one of NFSv3 or NFXv4 are also
	// selected, (i.e. {'SMB', 'NFSv4', 'SMB+NFS*'}).
	SupportedAccessProtocols []Protocol
	// SupportedLinesOfService shall contain known and supported
	// IOConnectivityLinesOfService.
	SupportedLinesOfService []IOConnectivityLineOfService
	// SupportedLinesOfServiceCount
	SupportedLinesOfServiceCount int `json:"SupportedLinesOfService@odata.count"`
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a IOConnectivityLoSCapabilities object from the raw JSON.
func (i *IOConnectivityLoSCapabilities) UnmarshalJSON(b []byte) error {
	type temp IOConnectivityLoSCapabilities
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*i = IOConnectivityLoSCapabilities(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *IOConnectivityLoSCapabilities) Update() error {
	readWriteFields := []string{
		"MaxSupportedBytesPerSecond",
		"MaxSupportedIOPS",
		"SupportedAccessProtocols",
		"SupportedLinesOfService",
	}

	return i.UpdateFromRawData(i, i.RawData, readWriteFields)
}

// GetIOConnectivityLoSCapabilities will get a IOConnectivityLoSCapabilities instance from the service.
func GetIOConnectivityLoSCapabilities(c Client, uri string) (*IOConnectivityLoSCapabilities, error) {
	return GetObject[IOConnectivityLoSCapabilities](c, uri)
}

// ListReferencedIOConnectivityLoSCapabilitiess gets the collection of IOConnectivityLoSCapabilities from
// a provided reference.
func ListReferencedIOConnectivityLoSCapabilitiess(c Client, link string) ([]*IOConnectivityLoSCapabilities, error) {
	return GetCollectionObjects[IOConnectivityLoSCapabilities](c, link)
}
