//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

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

// TelemetryData gets the TelemetryData linked resources.
func (c *CollectTelemetryDataResponse) TelemetryData(client Client) ([]*TelemetryData, error) {
	return GetObjects[TelemetryData](client, c.telemetryData)
}

// MetricDefinition gets the MetricDefinition linked resource.
func (m *TelemetryMetricValue) MetricDefinition(client Client) (*MetricDefinition, error) {
	if m.metricDefinition == "" {
		return nil, nil
	}
	return GetObject[MetricDefinition](client, m.metricDefinition)
}

// postUnmarshal accepts SupportedCollectionFuntions, a misspelling some
// services send instead of SupportedCollectionFunctions.
func (t *TelemetryService) postUnmarshal(b []byte) error {
	var v struct {
		SupportedCollectionFuntions []CollectionFunction
	}
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}

	if len(t.SupportedCollectionFunctions) == 0 && len(v.SupportedCollectionFuntions) > 0 {
		t.SupportedCollectionFunctions = v.SupportedCollectionFuntions
	}

	return nil
}

// postUnmarshal records the TelemetryData links.
func (c *CollectTelemetryDataResponse) postUnmarshal(b []byte) error {
	var t struct {
		TelemetryData Links
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	c.telemetryData = t.TelemetryData.ToStrings()

	return nil
}

// postUnmarshal records the MetricDefinition link.
func (m *TelemetryMetricValue) postUnmarshal(b []byte) error {
	var t struct {
		MetricDefinition Link
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	m.metricDefinition = t.MetricDefinition.String()

	return nil
}
