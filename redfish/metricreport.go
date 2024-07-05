//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// MetricReport shall represent a metric report in a Redfish implementation. When a metric report is deleted, the
// historic metric data used to generate the report shall be deleted as well unless other metric reports are
// consuming the data.
type MetricReport struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Context shall contain a client supplied context for the event destination to which this event is being sent.
	// This property shall only be present when sent as a payload in an event.
	Context string
	// Description provides a description of this resource.
	Description string
	// MetricReportDefinition shall contain a link to a resource of type MetricReportDefinition.
	metricReportDefinition string
	// MetricValues shall be metric values for this metric report.
	MetricValues []MetricValue
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Timestamp shall contain the time when the metric report was generated.
	Timestamp string
}

// UnmarshalJSON unmarshals a MetricReport object from the raw JSON.
func (metricreport *MetricReport) UnmarshalJSON(b []byte) error {
	type temp MetricReport
	var t struct {
		temp
		MetricReportDefinition common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*metricreport = MetricReport(t.temp)

	// Extract the links to other entities for later
	metricreport.metricReportDefinition = t.MetricReportDefinition.String()

	return nil
}

// MetricReportDefinition gets the definition for this metric
func (metricreport *MetricReport) MetricReportDefinition() (*MetricReportDefinition, error) {
	if metricreport.metricReportDefinition == "" {
		return nil, nil
	}

	return GetMetricReportDefinition(metricreport.GetClient(), metricreport.metricReportDefinition)
}

// GetMetricReport will get a MetricReport instance from the service.
func GetMetricReport(c common.Client, uri string) (*MetricReport, error) {
	return common.GetObject[MetricReport](c, uri)
}

// ListReferencedMetricReports gets the collection of MetricReport from
// a provided reference.
func ListReferencedMetricReports(c common.Client, link string) ([]*MetricReport, error) {
	return common.GetCollectionObjects[MetricReport](c, link)
}

// MetricValue shall contain properties that capture a metric value and other associated information.
type MetricValue struct {
	// MetricID shall contain the value of the ID property of the MetricDefinition resource that contains additional
	// information for the source metric.
	MetricID string
	// MetricProperty shall contain a URI following RFC6901-specified JSON pointer notation to the property from which
	// this metric is derived. The value of MetricValue may contain additional calculations performed on the property
	// based upon the configuration of the MetricReportDefinition.
	MetricProperty string
	// MetricValue shall contain the metric value, as a string.
	MetricValue string
	// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Timestamp shall time when the metric value was obtained. Note that this value may be different from the time
	// when this instance is created.
	Timestamp string
}
