//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a MetricReport object from the raw JSON.
func (m *MetricReport) UnmarshalJSON(b []byte) error {
	type temp MetricReport
	var tmp struct {
		temp
		MetricReportDefinition ResolvedLink[MetricReportDefinition] `json:"MetricReportDefinition"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MetricReport(tmp.temp)

	// Extract the links to other entities for later
	m.metricReportDefinition = tmp.MetricReportDefinition

	return nil
}

// MetricReportDefinition gets the MetricReportDefinition linked resource.
func (m *MetricReport) MetricReportDefinition() (*MetricReportDefinition, error) {
	return ResolveOrGet[MetricReportDefinition](m.client, m.metricReportDefinition)
}
