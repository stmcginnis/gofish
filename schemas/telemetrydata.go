//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #TelemetryData.v1_0_0.TelemetryData

package schemas

import (
	"encoding/json"
)

type TelemetryDataTypes string

const (
	// OEMTelemetryDataTypes OEM telemetry data.
	OEMTelemetryDataTypes TelemetryDataTypes = "OEM"
)

// TelemetryData shall represent bulk telemetry data in a Redfish
// implementation.
type TelemetryData struct {
	Entity
	// AdditionalData shall contain a Base64-encoded string, with padding
	// characters, of the associated telemetry data. The contents shall depend on
	// the value of the 'TelemetryDataType' property. The length of the value
	// should not exceed 4 KB. Larger telemetry data payloads should omit this
	// property and use the 'AdditionalDataURI' property to reference the data. If
	// both 'AdditionalData' and 'AdditionalDataURI' are present, 'AdditionalData'
	// shall contain the Base64-encoding of the data retrieved from the URI
	// specified by the 'AdditionalDataURI' property.
	AdditionalData string
	// AdditionalDataSizeBytes shall contain the size of the data available at
	// location specified by 'AdditionalDataURI':.
	AdditionalDataSizeBytes *int `json:",omitempty"`
	// AdditionalDataURI shall contain the URI at which to access the associated
	// telemetry data, using the Redfish protocol and authentication methods. If
	// both 'AdditionalData' and 'AdditionalDataURI' are present, 'AdditionalData'
	// shall contain a Base64-encoded string, with padding characters, of the data
	// retrieved from the URI specified by the 'AdditionalDataURI' property.
	AdditionalDataURI string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OEMTelemetryDataType shall contain the OEM-defined type of telemetry data
	// available in the 'AdditionalData' property or retrieved from the URI
	// specified by the 'AdditionalDataURI' property. The value of this property
	// should begin with the organization identifier. This property shall be
	// present if 'TelemetryDataType' is 'OEM'.
	OEMTelemetryDataType string
	// TelemetryDataType shall contain the type of telemetry data available in the
	// 'AdditionalData' property or retrieved from the URI specified by the
	// 'AdditionalDataURI' property.
	TelemetryDataType TelemetryDataTypes
	// Timestamp shall contain the time when the telemetry data was generated.
	Timestamp string
}

// GetTelemetryData will get a TelemetryData instance from the service.
func GetTelemetryData(c Client, uri string) (*TelemetryData, error) {
	return GetObject[TelemetryData](c, uri)
}

// ListReferencedTelemetryDatas gets the collection of TelemetryData from
// a provided reference.
func ListReferencedTelemetryDatas(c Client, link string) ([]*TelemetryData, error) {
	return GetCollectionObjects[TelemetryData](c, link)
}
