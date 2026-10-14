//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

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
		LogService              ResolvedLink[LogService] `json:"LogService"`
		MetricDefinitions       Link                     `json:"MetricDefinitions"`
		MetricReportDefinitions Link                     `json:"MetricReportDefinitions"`
		MetricReports           Link                     `json:"MetricReports"`
		TelemetryData           Link                     `json:"TelemetryData"`
		Triggers                Link                     `json:"Triggers"`
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
	t.logService = tmp.LogService

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
