//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// MetricReportDefinition gets the MetricReportDefinition linked resource.
func (m *MetricReport) MetricReportDefinition() (*MetricReportDefinition, error) {
	return ResolveOrGet[MetricReportDefinition](m.client, m.metricReportDefinition)
}

// postUnmarshal records the MetricReportDefinition link.
func (m *MetricReport) postUnmarshal(b []byte) error {
	var t struct {
		MetricReportDefinition ResolvedLink[MetricReportDefinition]
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	m.metricReportDefinition = t.MetricReportDefinition

	return nil
}
