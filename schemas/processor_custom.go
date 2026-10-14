//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// Metrics gets the Metrics linked resource.
func (p *Processor) Metrics() (*ProcessorMetrics, error) {
	if p.metrics == "" {
		return nil, nil
	}
	return GetObject[ProcessorMetrics](p.client, p.metrics)
}

// Metrics gets the Metrics linked resource.
func (m *ProcessorMemorySummary) Metrics(client Client) (*MemoryMetrics, error) {
	if m.metrics == "" {
		return nil, nil
	}
	return GetObject[MemoryMetrics](client, m.metrics)
}

// postUnmarshal records the Metrics link.
func (m *ProcessorMemorySummary) postUnmarshal(b []byte) error {
	var t struct {
		Metrics Link
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	m.metrics = t.Metrics.String()

	return nil
}

// postUnmarshal records the Metrics link.
func (p *Processor) postUnmarshal(b []byte) error {
	var t struct {
		Metrics Link
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	p.metrics = t.Metrics.String()

	return nil
}
