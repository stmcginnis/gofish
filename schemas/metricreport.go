//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/MetricReport.v1_5_2.json
// 2022.2 - #MetricReport.v1_5_2.MetricReport

package schemas

import (
	"encoding/json"
)

// MetricReport shall represent a metric report in a Redfish implementation.
// When a metric report is deleted, the historic metric data used to generate
// the report shall be deleted as well unless other metric reports are consuming
// the data.
type MetricReport struct {
	Entity
	// Context shall contain a client supplied context for the event destination to
	// which this event is being sent. This property shall only be present when
	// sent as a payload in an event.
	//
	// Version added: v1.4.0
	Context string
	// MetricReportDefinition shall contain a link to a resource of type
	// 'MetricReportDefinition'.
	metricReportDefinition string
	// MetricValues shall be metric values for this metric report.
	MetricValues []MetricValue
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ReportSequence shall contain the current sequence identifier for this metric
	// report. The sequence identifier is a unique identifier assigned by the
	// service for serializing metric reports as they are produced.
	//
	// Deprecated: v1.3.0
	// This property has been deprecated due to specification changes with regards
	// to Server-Sent Events.
	ReportSequence string
	// Timestamp shall contain the time when the metric report was generated.
	//
	// Version added: v1.1.0
	Timestamp string
}

// UnmarshalJSON unmarshals a MetricReport object from the raw JSON.
func (m *MetricReport) UnmarshalJSON(b []byte) error {
	type temp MetricReport
	var tmp struct {
		temp
		MetricReportDefinition Link `json:"MetricReportDefinition"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MetricReport(tmp.temp)

	// Extract the links to other entities for later
	m.metricReportDefinition = tmp.MetricReportDefinition.String()

	return nil
}

// GetMetricReport will get a MetricReport instance from the service.
func GetMetricReport(c Client, uri string) (*MetricReport, error) {
	return GetObject[MetricReport](c, uri)
}

// ListReferencedMetricReports gets the collection of MetricReport from
// a provided reference.
func ListReferencedMetricReports(c Client, link string) ([]*MetricReport, error) {
	return GetCollectionObjects[MetricReport](c, link)
}

// MetricReportDefinition gets the MetricReportDefinition linked resource.
func (m *MetricReport) MetricReportDefinition() (*MetricReportDefinition, error) {
	if m.metricReportDefinition == "" {
		return nil, nil
	}
	return GetObject[MetricReportDefinition](m.client, m.metricReportDefinition)
}

// MetricValue shall contain properties that capture a metric value and other
// associated information.
type MetricValue struct {
	// MetricDefinition shall contain a link to a resource of type
	// 'MetricDefinition' that describes what this metric value captures.
	//
	// Deprecated: v1.5.0
	// This property has been deprecated in favor of the 'MetricId' property.
	metricDefinition string
	// MetricID shall contain the value of the 'Id' property of the
	// 'MetricDefinition' resource that contains additional information for the
	// source metric.
	MetricID string `json:"MetricId"`
	// MetricProperty shall contain a URI following RFC6901-specified JSON pointer
	// notation to the property from which this metric is derived. The value of
	// 'MetricValue' may contain additional calculations performed on the property
	// based upon the configuration of the 'MetricReportDefinition'.
	MetricProperty string
	// MetricValue shall contain the metric value, as a string. For numeric
	// metrics, the service shall convert the number to a string representation of
	// the number. For array metrics, the service shall convert the array to an
	// RFC8259-defined JSON string. For boolean metrics, this property shall
	// contain the strings 'true' or 'false'. If the metric value is 'null', this
	// property shall contain 'null'.
	MetricValue string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.2.0
	OEM json.RawMessage `json:"Oem"`
	// Timestamp shall time when the metric value was obtained. Note that this
	// value may be different from the time when this instance is created.
	Timestamp string
}

// UnmarshalJSON unmarshals a MetricValue object from the raw JSON.
func (m *MetricValue) UnmarshalJSON(b []byte) error {
	type temp MetricValue
	var tmp struct {
		temp
		MetricDefinition Link `json:"MetricDefinition"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MetricValue(tmp.temp)

	// Extract the links to other entities for later
	m.metricDefinition = tmp.MetricDefinition.String()

	return nil
}

// MetricDefinition gets the MetricDefinition linked resource.
func (m *MetricValue) MetricDefinition(client Client) (*MetricDefinition, error) {
	if m.metricDefinition == "" {
		return nil, nil
	}
	return GetObject[MetricDefinition](client, m.metricDefinition)
}
