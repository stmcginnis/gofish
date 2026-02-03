//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.4 - #MetricReportDefinition.v1_4_7.MetricReportDefinition

package schemas

import (
	"encoding/json"
)

// CollectionTimeScope is This type shall specify the time scope of the
// corresponding metric values.
type CollectionTimeScope string

const (
	// PointCollectionTimeScope shall indicate the corresponding metric values
	// apply to a point in time. On the corresponding metric value instances, the
	// 'Timestamp' property value in the metric report shall specify the point in
	// time.
	PointCollectionTimeScope CollectionTimeScope = "Point"
	// IntervalCollectionTimeScope shall indicate the corresponding metric values
	// apply to a time interval. On the corresponding metric value instances, the
	// 'Timestamp' property value in the metric report shall specify the end of the
	// time interval and the 'CollectionDuration' property shall specify its
	// duration.
	IntervalCollectionTimeScope CollectionTimeScope = "Interval"
	// StartupIntervalCollectionTimeScope shall indicate the corresponding metric
	// values apply to a time interval that began at the startup of the measured
	// resource. On the corresponding metric value instances, the 'Timestamp'
	// property value in the metric report shall specify the end of the time
	// interval. The 'CollectionDuration' property value shall specify the duration
	// between the startup of the resource and timestamp.
	StartupIntervalCollectionTimeScope CollectionTimeScope = "StartupInterval"
)

// MetricReportDefinitionType is This type shall specify when the metric report
// is generated.
type MetricReportDefinitionType string

const (
	// PeriodicMetricReportDefinitionType The metric report is generated at a
	// periodic time interval, specified in the 'Schedule' property.
	PeriodicMetricReportDefinitionType MetricReportDefinitionType = "Periodic"
	// OnChangeMetricReportDefinitionType The metric report is generated when any
	// of the metric values change.
	OnChangeMetricReportDefinitionType MetricReportDefinitionType = "OnChange"
	// OnRequestMetricReportDefinitionType The metric report is generated when an
	// HTTP 'GET' is performed on the specified metric report.
	OnRequestMetricReportDefinitionType MetricReportDefinitionType = "OnRequest"
)

// ReportActionsEnum is This type shall specify the actions to perform when a
// metric report is generated.
type ReportActionsEnum string

const (
	// LogToMetricReportsCollectionReportActionsEnum shall indicate the service
	// records the occurrence to the metric report collection found under the
	// telemetry service. The service shall update the metric report based on the
	// setting of the 'ReportUpdates' property.
	LogToMetricReportsCollectionReportActionsEnum ReportActionsEnum = "LogToMetricReportsCollection"
	// RedfishEventReportActionsEnum shall indicate the service sends a Redfish
	// event of type 'MetricReport' to subscribers in the event subscription
	// collection of the event service.
	RedfishEventReportActionsEnum ReportActionsEnum = "RedfishEvent"
)

// ReportUpdatesEnum is This type shall indicate how the service handles
// subsequent metric reports when a metric report exists.
type ReportUpdatesEnum string

const (
	// OverwriteReportUpdatesEnum shall indicate the service overwrites the metric
	// report referenced by the 'MetricReport' property.
	OverwriteReportUpdatesEnum ReportUpdatesEnum = "Overwrite"
	// AppendWrapsWhenFullReportUpdatesEnum shall indicate the service appends new
	// information to the metric report referenced by the 'MetricReport' property.
	// The service shall overwrite entries in the metric report with new entries
	// when the metric report has reached its maximum capacity.
	AppendWrapsWhenFullReportUpdatesEnum ReportUpdatesEnum = "AppendWrapsWhenFull"
	// AppendStopsWhenFullReportUpdatesEnum shall indicate the service appends new
	// information to the metric report referenced by the 'MetricReport' property.
	// The service shall stop adding entries when the metric report has reached its
	// maximum capacity. The 'State' property within 'Status' should be set to
	// 'Disabled' and the 'MetricReportDefinitionEnabled' property should be set to
	// 'false' when the append limit is reached.
	AppendStopsWhenFullReportUpdatesEnum ReportUpdatesEnum = "AppendStopsWhenFull"
	// NewReportReportUpdatesEnum shall indicate the service creates a new metric
	// report resource, whose 'Id' property is a service-defined identifier
	// concatenated with the timestamp. The metric report referenced by the
	// 'MetricReport' property shall reference the metric report most recently
	// created by this metric report definition.
	NewReportReportUpdatesEnum ReportUpdatesEnum = "NewReport"
)

// MetricReportDefinition shall specify a set of metrics that shall be collected
// into a metric report in a Redfish implementation.
type MetricReportDefinition struct {
	Entity
	// AppendLimit shall contain a number that indicates the maximum number of
	// entries that can be appended to a metric report. When the metric report
	// reaches its limit, its behavior shall be dictated by the 'ReportUpdates'
	// property. This property shall be required if 'ReportUpdates' contains
	// 'AppendWrapsWhenFull' or 'AppendStopsWhenFull'.
	AppendLimit uint
	// MetricProperties shall contain a list of URIs with wildcards and property
	// identifiers to include in the metric report. A set of curly braces shall
	// delimit each wildcard in the URI. The corresponding entry in the 'Wildcard'
	// property shall replace each wildcard. After each wildcard is replaced, it
	// shall describe a resource property to include in the metric report. The
	// property identifiers portion of the URI shall follow RFC6901-specified JSON
	// pointer notation rules.
	MetricProperties []string
	// MetricReport shall contain a link to a resource of type 'MetricReport' that
	// represents the most recent metric report produced by this metric report
	// definition.
	metricReport string
	// MetricReportDefinitionEnabled shall indicate whether the generation of new
	// metric reports is enabled.
	//
	// Version added: v1.2.0
	MetricReportDefinitionEnabled bool
	// MetricReportDefinitionType shall specify when the metric report is
	// generated. If the value is 'Periodic', the 'Schedule' property shall be
	// present.
	MetricReportDefinitionType MetricReportDefinitionType
	// MetricReportHeartbeatInterval shall contain a Redfish duration that
	// describes the time interval between generations of the unsuppressed metric
	// report. It shall always be a value greater than the 'RecurrenceInterval'
	// property within 'Schedule' and should only apply when the
	// 'SuppressRepeatedMetricValue' property is 'true'.
	//
	// Version added: v1.2.0
	MetricReportHeartbeatInterval string
	// Metrics shall contain a list of metrics to include in the metric report. The
	// metrics may include calculations to apply to metric properties.
	Metrics []Metric
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ReportActions shall contain the set of actions to perform when the metric
	// report is generated. This property should be ignored if
	// 'MetricReportDefinitionType' contains the value 'OnRequest'.
	ReportActions []ReportActionsEnum
	// ReportTimespan shall contain the maximum timespan that a metric report can
	// cover.
	//
	// Version added: v1.3.0
	ReportTimespan string
	// ReportUpdates shall contain the behavior for how subsequent metric reports
	// are handled in relationship to an existing metric report created from the
	// metric report definition. This property should be ignored if
	// 'MetricReportDefinitionType' contains the value 'OnRequest'.
	ReportUpdates ReportUpdatesEnum
	// Schedule shall contain the schedule of the metric report. The metric report
	// shall be generated at an interval specified by the 'RecurrenceInterval'
	// property within 'Schedule'. If the 'MaxOccurrences' property within
	// 'Schedule' is specified, the metric report shall no longer be generated
	// after the specified number of occurrences. The 'State' property within
	// 'Status' should be set to 'Disabled' and the 'MetricReportDefinitionEnabled'
	// property should be set to 'false' when the specified number of occurrences
	// is reached.
	Schedule Schedule
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SuppressRepeatedMetricValue shall indicate whether any metrics are
	// suppressed from the generated metric report. If 'true', any metric that
	// equals the same value in the previously generated metric report is
	// suppressed from the current report. Also, duplicate metrics are suppressed.
	// If 'false', no metrics are suppressed from the current report. The current
	// report may contain no metrics if all metrics equal the values in the
	// previously generated metric report.
	//
	// Version added: v1.2.0
	SuppressRepeatedMetricValue bool
	// Wildcards shall contain a set of wildcards and their replacement strings,
	// which are applied to the 'MetricProperties' property. Each wildcard
	// expressed in the 'MetricProperties' property shall have a corresponding
	// entry in this property.
	Wildcards []Wildcard
	// triggers are the URIs for Triggers.
	triggers []string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a MetricReportDefinition object from the raw JSON.
func (m *MetricReportDefinition) UnmarshalJSON(b []byte) error {
	type temp MetricReportDefinition
	type mLinks struct {
		Triggers Links `json:"Triggers"`
	}
	var tmp struct {
		temp
		Links        mLinks
		MetricReport Link `json:"MetricReport"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MetricReportDefinition(tmp.temp)

	// Extract the links to other entities for later
	m.triggers = tmp.Links.Triggers.ToStrings()
	m.metricReport = tmp.MetricReport.String()

	// This is a read/write object, so we need to save the raw object data for later
	m.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (m *MetricReportDefinition) Update() error {
	readWriteFields := []string{
		"MetricProperties",
		"MetricReportDefinitionEnabled",
		"MetricReportDefinitionType",
		"MetricReportHeartbeatInterval",
		"ReportActions",
		"ReportTimespan",
		"ReportUpdates",
		"SuppressRepeatedMetricValue",
	}

	return m.UpdateFromRawData(m, m.RawData, readWriteFields)
}

// GetMetricReportDefinition will get a MetricReportDefinition instance from the service.
func GetMetricReportDefinition(c Client, uri string) (*MetricReportDefinition, error) {
	return GetObject[MetricReportDefinition](c, uri)
}

// ListReferencedMetricReportDefinitions gets the collection of MetricReportDefinition from
// a provided reference.
func ListReferencedMetricReportDefinitions(c Client, link string) ([]*MetricReportDefinition, error) {
	return GetCollectionObjects[MetricReportDefinition](c, link)
}

// Triggers gets the Triggers linked resources.
func (m *MetricReportDefinition) Triggers() ([]*Triggers, error) {
	return GetObjects[Triggers](m.client, m.triggers)
}

// MetricReport gets the MetricReport linked resource.
func (m *MetricReportDefinition) MetricReport() (*MetricReport, error) {
	if m.metricReport == "" {
		return nil, nil
	}
	return GetObject[MetricReport](m.client, m.metricReport)
}

// Metric shall specify a set of metrics to include in the metric report. The
// algorithm specified by 'CollectionFunction', if present, shall be applied to
// each of the metric properties listed in the 'MetricProperties' property or
// the metric properties specified in the 'MetricDefinition' referenced by the
// 'MetricId' property prior to being included in the metric report.
type Metric struct {
	// CollectionDuration shall specify the duration over which the function is
	// computed.
	CollectionDuration string
	// CollectionFunction shall specify the function to perform on each of the
	// metric properties listed in the 'MetricProperties' property or the metric
	// properties specified in the 'MetricDefinition' referenced by the 'MetricId'
	// property. If not specified, calculations shall not be performed on the
	// metric properties.
	CollectionFunction CalculationAlgorithmEnum
	// CollectionTimeScope shall specify the scope of time over which the function
	// is applied.
	CollectionTimeScope CollectionTimeScope
	// MetricID shall contain the value of the 'Id' property of the
	// 'MetricDefinition' resource that contains the metric properties to include
	// in the metric report. This property should not be present if
	// 'MetricProperties' is present.
	MetricID string `json:"MetricId"`
	// MetricProperties shall contain a list of URIs with wildcards and property
	// identifiers to include in the metric report. A set of curly braces shall
	// delimit each wildcard in the URI. The corresponding entry in the 'Wildcard'
	// property shall replace each wildcard. After each wildcard is replaced, it
	// shall describe a resource property to include in the metric report. The
	// property identifiers portion of the URI shall follow RFC6901-specified JSON
	// pointer notation rules. This property should not be present if 'MetricId' is
	// present.
	MetricProperties []string
	// OEM shall contain the OEM extensions. All values for properties contained in
	// this object shall conform to the Redfish Specification-described
	// requirements.
	//
	// Version added: v1.4.0
	OEM json.RawMessage `json:"Oem"`
}
