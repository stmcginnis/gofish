//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// CollectionFunction is if present, the metric value shall be computed according to this function.
type CollectionFunction string

const (
	// AverageCollectionFunction An averaging function.
	AverageCollectionFunction CollectionFunction = "Average"
	// MaximumCollectionFunction A maximum function.
	MaximumCollectionFunction CollectionFunction = "Maximum"
	// MinimumCollectionFunction A minimum function.
	MinimumCollectionFunction CollectionFunction = "Minimum"
	// SummationCollectionFunction A summation function.
	SummationCollectionFunction CollectionFunction = "Summation"
)

// TelemetryMetricValue shall contain properties that capture a metric value and other associated information.
type TelemetryMetricValue struct {
	// MetricDefinition shall contain a link to a resource of type MetricDefinition that describes what this metric
	// value captures.
	MetricDefinition MetricDefinition
	// MetricID shall contain the same value as the ID property of the source metric within the associated metric
	// definition.
	MetricID string
	// MetricProperty shall be the URI to the property following the JSON fragment notation, as defined by RFC6901, to
	// identify an individual property in a Redfish resource.
	MetricProperty string
	// TelemetryMetricValue shall contain the metric value, as a string.
	MetricValue string
	// Timestamp shall time when the metric value was obtained. Note that this value may be different from the time
	// when this instance is created.
	Timestamp string
}

// TelemetryService This resource contains a telemetry service for a Redfish implementation.
type TelemetryService struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// LogService shall contain a link to a resource of type LogService that this telemetry service uses.
	logService string
	// MaxReports shall contain the maximum number of metric reports that this service supports.
	MaxReports int
	// MetricDefinitions shall contain a link to a resource collection of type MetricDefinitionCollection.
	metricDefinitions string
	// MetricReportDefinitions shall contain a link to a resource collection of type MetricReportDefinitionCollection.
	metricReportDefinitions string
	// MetricReports shall contain a link to a resource collection of type MetricReportCollection.
	metricReports string
	// MinCollectionInterval shall contain the minimum time interval between gathering metric data that this service
	// allows.
	MinCollectionInterval string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupportedCollectionFunctions shall contain the function to apply over the collection duration.
	SupportedCollectionFunctions []CollectionFunction
	// Triggers shall contain a link to a resource collection of type TriggersCollection.
	triggers string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	clearMetricReportTarget                      string
	resetMetricReportDefinitionsToDefaultsTarget string
	resetTriggersToDefaultsTarget                string
	submitTestMetricReportTarget                 string
}

// UnmarshalJSON unmarshals a TelemetryService object from the raw JSON.
func (telemetryservice *TelemetryService) UnmarshalJSON(b []byte) error {
	type temp TelemetryService
	type Actions struct {
		ClearMetricReports                     common.ActionTarget `json:"#TelemetryService.ClearMetricReports"`
		ResetMetricReportDefinitionsToDefaults common.ActionTarget `json:"#TelemetryService.ResetMetricReportDefinitionsToDefaults"`
		ResetTriggersToDefaults                common.ActionTarget `json:"#TelemetryService.ResetTriggersToDefaults"`
		SubmitTestMetricReport                 common.ActionTarget `json:"#TelemetryService.SubmitTestMetricReport"`
	}
	var t struct {
		temp
		Actions                 Actions
		LogService              common.Link
		MetricDefinitions       common.Link
		MetricReportDefinitions common.Link
		MetricReports           common.Link
		Triggers                common.Link
		// Bug in Supermicro implementation
		SupportedCollectionFuntions []CollectionFunction
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*telemetryservice = TelemetryService(t.temp)

	// Extract the links to other entities for later
	telemetryservice.logService = t.LogService.String()
	telemetryservice.metricDefinitions = t.MetricDefinitions.String()
	telemetryservice.metricReportDefinitions = t.MetricReportDefinitions.String()
	telemetryservice.metricReports = t.MetricReports.String()
	telemetryservice.triggers = t.Triggers.String()

	telemetryservice.clearMetricReportTarget = t.Actions.ClearMetricReports.Target
	telemetryservice.resetMetricReportDefinitionsToDefaultsTarget = t.Actions.ResetMetricReportDefinitionsToDefaults.Target
	telemetryservice.resetTriggersToDefaultsTarget = t.Actions.ResetTriggersToDefaults.Target
	telemetryservice.submitTestMetricReportTarget = t.Actions.SubmitTestMetricReport.Target

	if len(telemetryservice.SupportedCollectionFunctions) == 0 && len(t.SupportedCollectionFuntions) > 0 {
		telemetryservice.SupportedCollectionFunctions = t.SupportedCollectionFuntions
	}

	// This is a read/write object, so we need to save the raw object data for later
	telemetryservice.rawData = b

	return nil
}

// LogService gets the log service that the telemetry service uses.
func (telemetryservice *TelemetryService) LogService() (*LogService, error) {
	if telemetryservice.logService == "" {
		return nil, nil
	}
	return GetLogService(telemetryservice.GetClient(), telemetryservice.logService)
}

// MetricDefinitions gets the metric definitions.
func (telemetryservice *TelemetryService) MetricDefinitions() ([]*MetricDefinition, error) {
	return ListReferencedMetricDefinitions(telemetryservice.GetClient(), telemetryservice.metricDefinitions)
}

// MetricReportDefinitions gets the metric report definitions.
func (telemetryservice *TelemetryService) MetricReportDefinitions() ([]*MetricReportDefinition, error) {
	return ListReferencedMetricReportDefinitions(telemetryservice.GetClient(), telemetryservice.metricReportDefinitions)
}

// MetricReports gets the metric reports.
func (telemetryservice *TelemetryService) MetricReports() ([]*MetricReport, error) {
	return ListReferencedMetricReports(telemetryservice.GetClient(), telemetryservice.metricReports)
}

// Triggers gets the triggers.
func (telemetryservice *TelemetryService) Triggers() ([]*Triggers, error) {
	return ListReferencedTriggerss(telemetryservice.GetClient(), telemetryservice.triggers)
}

// ClearMetricReports will clear the metric reports for this telemetry service.
func (telemetryservice *TelemetryService) ClearMetricReports() error {
	return telemetryservice.Post(telemetryservice.clearMetricReportTarget, nil)
}

// ResetMetricReportDefinitionsToDefaults will reset the metric report definitions to factory defaults.
func (telemetryservice *TelemetryService) ResetMetricReportDefinitionsToDefaults() error {
	return telemetryservice.Post(telemetryservice.resetMetricReportDefinitionsToDefaultsTarget, nil)
}

// ResetTriggersToDefaults will reset the triggers to factory defaults.
func (telemetryservice *TelemetryService) ResetTriggersToDefaults() error {
	return telemetryservice.Post(telemetryservice.resetTriggersToDefaultsTarget, nil)
}

// SubmitTestMetricReport will immediately generate the metric report as an alert event.
// `reportValues` is the content for the generated metric report.
// `metricReportName` is the name for the metric report.
func (telemetryservice *TelemetryService) SubmitTestMetricReport(reportValues []TelemetryMetricValue, reportName string) error {
	t := struct {
		GeneratedMetricReportValues []TelemetryMetricValue
		MetricReportName            string
	}{
		GeneratedMetricReportValues: reportValues,
		MetricReportName:            reportName,
	}
	return telemetryservice.Post(telemetryservice.submitTestMetricReportTarget, t)
}

// Update commits updates to this object's properties to the running system.
func (telemetryservice *TelemetryService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(TelemetryService)
	original.UnmarshalJSON(telemetryservice.rawData)

	readWriteFields := []string{
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(telemetryservice).Elem()

	return telemetryservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetTelemetryService will get a TelemetryService instance from the service.
func GetTelemetryService(c common.Client, uri string) (*TelemetryService, error) {
	return common.GetObject[TelemetryService](c, uri)
}

// ListReferencedTelemetryServices gets the collection of TelemetryService from
// a provided reference.
func ListReferencedTelemetryServices(c common.Client, link string) ([]*TelemetryService, error) {
	return common.GetCollectionObjects[TelemetryService](c, link)
}
