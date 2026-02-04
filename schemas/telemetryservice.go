//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/TelemetryService.v1_4_1.json
// 2025.2 - #TelemetryService.v1_4_1.TelemetryService

package schemas

import (
	"encoding/json"
)

// CollectionFunction is If present, the metric value shall be computed
// according to this function.
type CollectionFunction string

const (
	// AverageCollectionFunction is an averaging function.
	AverageCollectionFunction CollectionFunction = "Average"
	// MaximumCollectionFunction is a maximum function.
	MaximumCollectionFunction CollectionFunction = "Maximum"
	// MinimumCollectionFunction is a minimum function.
	MinimumCollectionFunction CollectionFunction = "Minimum"
	// SummationCollectionFunction is a summation function.
	SummationCollectionFunction CollectionFunction = "Summation"
)

// TelemetryService This resource contains a telemetry service for a Redfish
// implementation.
type TelemetryService struct {
	Entity
	// LogService shall contain a link to a resource of type 'LogService' that this
	// telemetry service uses.
	logService string
	// MaxReports shall contain the maximum number of metric reports that this
	// service supports.
	MaxReports *int `json:",omitempty"`
	// MetricDefinitions shall contain a link to a resource collection of type
	// 'MetricDefinitionCollection'.
	metricDefinitions string
	// MetricReportDefinitions shall contain a link to a resource collection of
	// type 'MetricReportDefinitionCollection'.
	metricReportDefinitions string
	// MetricReports shall contain a link to a resource collection of type
	// 'MetricReportCollection'.
	metricReports string
	// MinCollectionInterval shall contain the minimum time interval between
	// gathering metric data that this service allows.
	MinCollectionInterval string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ServiceEnabled shall indicate whether this service is enabled.
	//
	// Version added: v1.2.0
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SupportedCollectionFunctions shall contain the function to apply over the
	// collection duration.
	SupportedCollectionFunctions []CollectionFunction
	// SupportedOEMTelemetryDataTypes shall contain a list of supported OEM-defined
	// telemetry data types.
	//
	// Version added: v1.4.0
	SupportedOEMTelemetryDataTypes []string
	// SupportedTelemetryDataTypes shall contain a list of supported telemetry data
	// types.
	//
	// Version added: v1.4.0
	SupportedTelemetryDataTypes []TelemetryDataTypes
	// TelemetryData shall contain a link to a resource collection of type
	// 'TelemetryDataCollection'.
	//
	// Version added: v1.4.0
	telemetryData string
	// Triggers shall contain a link to a resource collection of type
	// 'TriggersCollection'.
	triggers string
	// clearMetricReportsTarget is the URL to send ClearMetricReports requests.
	clearMetricReportsTarget string
	// clearTelemetryDataTarget is the URL to send ClearTelemetryData requests.
	clearTelemetryDataTarget string
	// collectTelemetryDataTarget is the URL to send CollectTelemetryData requests.
	collectTelemetryDataTarget string
	// resetMetricReportDefinitionsToDefaultsTarget is the URL to send ResetMetricReportDefinitionsToDefaults requests.
	resetMetricReportDefinitionsToDefaultsTarget string
	// resetTriggersToDefaultsTarget is the URL to send ResetTriggersToDefaults requests.
	resetTriggersToDefaultsTarget string
	// submitTestMetricReportTarget is the URL to send SubmitTestMetricReport requests.
	submitTestMetricReportTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a TelemetryService object from the raw JSON.
func (t *TelemetryService) UnmarshalJSON(b []byte) error {
	type temp TelemetryService
	type tActions struct {
		ClearMetricReports                     ActionTarget `json:"#TelemetryService.ClearMetricReports"`
		ClearTelemetryData                     ActionTarget `json:"#TelemetryService.ClearTelemetryData"`
		CollectTelemetryData                   ActionTarget `json:"#TelemetryService.CollectTelemetryData"`
		ResetMetricReportDefinitionsToDefaults ActionTarget `json:"#TelemetryService.ResetMetricReportDefinitionsToDefaults"`
		ResetTriggersToDefaults                ActionTarget `json:"#TelemetryService.ResetTriggersToDefaults"`
		SubmitTestMetricReport                 ActionTarget `json:"#TelemetryService.SubmitTestMetricReport"`
	}
	var tmp struct {
		temp
		Actions                 tActions
		LogService              Link `json:"LogService"`
		MetricDefinitions       Link `json:"MetricDefinitions"`
		MetricReportDefinitions Link `json:"MetricReportDefinitions"`
		MetricReports           Link `json:"MetricReports"`
		TelemetryData           Link `json:"TelemetryData"`
		Triggers                Link `json:"Triggers"`
		// Bug in Supermicro implementation
		SupportedCollectionFuntions []CollectionFunction
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = TelemetryService(tmp.temp)

	// Extract the links to other entities for later
	t.clearMetricReportsTarget = tmp.Actions.ClearMetricReports.Target
	t.clearTelemetryDataTarget = tmp.Actions.ClearTelemetryData.Target
	t.collectTelemetryDataTarget = tmp.Actions.CollectTelemetryData.Target
	t.resetMetricReportDefinitionsToDefaultsTarget = tmp.Actions.ResetMetricReportDefinitionsToDefaults.Target
	t.resetTriggersToDefaultsTarget = tmp.Actions.ResetTriggersToDefaults.Target
	t.submitTestMetricReportTarget = tmp.Actions.SubmitTestMetricReport.Target
	t.logService = tmp.LogService.String()
	t.metricDefinitions = tmp.MetricDefinitions.String()
	t.metricReportDefinitions = tmp.MetricReportDefinitions.String()
	t.metricReports = tmp.MetricReports.String()
	t.telemetryData = tmp.TelemetryData.String()
	t.triggers = tmp.Triggers.String()

	if len(t.SupportedCollectionFunctions) == 0 && len(tmp.SupportedCollectionFuntions) > 0 {
		t.SupportedCollectionFunctions = tmp.SupportedCollectionFuntions
	}

	// This is a read/write object, so we need to save the raw object data for later
	t.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (t *TelemetryService) Update() error {
	readWriteFields := []string{
		"ServiceEnabled",
	}

	return t.UpdateFromRawData(t, t.RawData, readWriteFields)
}

// GetTelemetryService will get a TelemetryService instance from the service.
func GetTelemetryService(c Client, uri string) (*TelemetryService, error) {
	return GetObject[TelemetryService](c, uri)
}

// ListReferencedTelemetryServices gets the collection of TelemetryService from
// a provided reference.
func ListReferencedTelemetryServices(c Client, link string) ([]*TelemetryService, error) {
	return GetCollectionObjects[TelemetryService](c, link)
}

// This action shall delete all entries found in the metric report collection
// for this telemetry service.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (t *TelemetryService) ClearMetricReports() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(t.client,
		t.clearMetricReportsTarget, payload, t.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall delete all entries found in the telemetry data collection
// for this telemetry service.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (t *TelemetryService) ClearTelemetryData() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(t.client,
		t.clearTelemetryDataTarget, payload, t.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall collect the telemetry data from a device or service. The
// 'Location' header in the response shall contain a URI to a resource of type
// 'TelemetryData' that contains the telemetry data. The 'AdditionalDataURI'
// property in the referenced 'TelemetryData' resource shall contain the URI to
// download the telemetry data.
// oEMTelemetryDataType - This parameter shall contain the OEM-defined type of
// telemetry data to collect. This parameter shall be required if
// 'TelemetryDataType' is 'OEM'.
// targetDevices - This parameter shall contain an array of devices from which
// to collect telemetry data.
// telemetryDataType - This parameter shall contain the type of telemetry data
// to collect.
func (t *TelemetryService) CollectTelemetryData(oEMTelemetryDataType string, targetDevices []Entity, telemetryDataType TelemetryDataTypes) (*CollectTelemetryDataResponse, error) {
	payload := make(map[string]any)
	payload["OEMTelemetryDataType"] = oEMTelemetryDataType
	payload["TargetDevices"] = targetDevices
	payload["TelemetryDataType"] = telemetryDataType

	resp, err := t.PostWithResponse(t.collectTelemetryDataTarget, payload)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, CleanupHTTPResponse(resp)
	}

	var result CollectTelemetryDataResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// This action shall reset all entries found in the metric report definition
// collection to factory defaults. This action may delete members of the metric
// report definition collection.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (t *TelemetryService) ResetMetricReportDefinitionsToDefaults() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(t.client,
		t.resetMetricReportDefinitionsToDefaultsTarget, payload, t.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall reset all entries found in the triggers collection to
// factory defaults. This action may delete members of the triggers collection.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (t *TelemetryService) ResetTriggersToDefaults() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(t.client,
		t.resetTriggersToDefaultsTarget, payload, t.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall cause the event service to immediately generate the metric
// report as an alert event. Then, this message should be sent to any
// appropriate event destinations.
// generatedMetricReportValues - This parameter shall contain the contents of
// the 'MetricReportValues' array property in the generated metric report.
// metricReportName - This parameter shall contain the name of the generated
// metric report.
// metricReportValues - This parameter shall contain the contents of the
// 'MetricReportValues' array property in the generated metric report.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (t *TelemetryService) SubmitTestMetricReport(generatedMetricReportValues []MetricValue, metricReportName string, metricReportValues string) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["GeneratedMetricReportValues"] = generatedMetricReportValues
	payload["MetricReportName"] = metricReportName
	payload["MetricReportValues"] = metricReportValues
	resp, taskInfo, err := PostWithTask(t.client,
		t.submitTestMetricReportTarget, payload, t.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// LogService gets the LogService linked resource.
func (t *TelemetryService) LogService() (*LogService, error) {
	if t.logService == "" {
		return nil, nil
	}
	return GetObject[LogService](t.client, t.logService)
}

// MetricDefinitions gets the MetricDefinitions collection.
func (t *TelemetryService) MetricDefinitions() ([]*MetricDefinition, error) {
	if t.metricDefinitions == "" {
		return nil, nil
	}
	return GetCollectionObjects[MetricDefinition](t.client, t.metricDefinitions)
}

// MetricReportDefinitions gets the MetricReportDefinitions collection.
func (t *TelemetryService) MetricReportDefinitions() ([]*MetricReportDefinition, error) {
	if t.metricReportDefinitions == "" {
		return nil, nil
	}
	return GetCollectionObjects[MetricReportDefinition](t.client, t.metricReportDefinitions)
}

// MetricReports gets the MetricReports collection.
func (t *TelemetryService) MetricReports() ([]*MetricReport, error) {
	if t.metricReports == "" {
		return nil, nil
	}
	return GetCollectionObjects[MetricReport](t.client, t.metricReports)
}

// TelemetryData gets the TelemetryData collection.
func (t *TelemetryService) TelemetryData() ([]*TelemetryData, error) {
	if t.telemetryData == "" {
		return nil, nil
	}
	return GetCollectionObjects[TelemetryData](t.client, t.telemetryData)
}

// Triggers gets the Triggers collection.
func (t *TelemetryService) Triggers() ([]*Triggers, error) {
	if t.triggers == "" {
		return nil, nil
	}
	return GetCollectionObjects[Triggers](t.client, t.triggers)
}

// CollectTelemetryDataResponse shall contain the properties found in the
// response body for the 'CollectTelemetryData' action.
type CollectTelemetryDataResponse struct {
	// TelemetryData shall contain an array of links to resources of type
	// 'TelemetryData' that represent the collected telemetry data.
	//
	// Version added: v1.4.0
	telemetryData []string
}

// UnmarshalJSON unmarshals a CollectTelemetryDataResponse object from the raw JSON.
func (c *CollectTelemetryDataResponse) UnmarshalJSON(b []byte) error {
	type temp CollectTelemetryDataResponse
	var tmp struct {
		temp
		TelemetryData Links `json:"TelemetryData"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CollectTelemetryDataResponse(tmp.temp)

	// Extract the links to other entities for later
	c.telemetryData = tmp.TelemetryData.ToStrings()

	return nil
}

// TelemetryData gets the TelemetryData linked resources.
func (c *CollectTelemetryDataResponse) TelemetryData(client Client) ([]*TelemetryData, error) {
	return GetObjects[TelemetryData](client, c.telemetryData)
}

// TelemetryMetricValue shall contain properties that capture a metric value and other
// associated information.
type TelemetryMetricValue struct {
	// MetricDefinition shall contain a link to a resource of type
	// 'MetricDefinition' that describes what this metric value captures.
	//
	// Version added: v1.1.0
	metricDefinition string
	// MetricID shall contain the same value as the 'Id' property of the source
	// metric within the associated metric definition.
	//
	// Version added: v1.1.0
	MetricID string `json:"MetricId"`
	// MetricProperty shall be the URI to the property following the JSON fragment
	// notation, as defined by RFC6901, to identify an individual property in a
	// Redfish resource.
	//
	// Version added: v1.1.0
	MetricProperty string
	// MetricValue shall contain the metric value, as a string.
	//
	// Version added: v1.1.0
	MetricValue string
	// Timestamp shall time when the metric value was obtained. Note that this
	// value may be different from the time when this instance is created.
	//
	// Version added: v1.1.0
	Timestamp string
}

// UnmarshalJSON unmarshals a TelemetryMetricValue object from the raw JSON.
func (m *TelemetryMetricValue) UnmarshalJSON(b []byte) error {
	type temp TelemetryMetricValue
	var tmp struct {
		temp
		MetricDefinition Link `json:"MetricDefinition"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = TelemetryMetricValue(tmp.temp)

	// Extract the links to other entities for later
	m.metricDefinition = tmp.MetricDefinition.String()

	return nil
}

// MetricDefinition gets the MetricDefinition linked resource.
func (m *TelemetryMetricValue) MetricDefinition(client Client) (*MetricDefinition, error) {
	if m.metricDefinition == "" {
		return nil, nil
	}
	return GetObject[MetricDefinition](client, m.metricDefinition)
}
