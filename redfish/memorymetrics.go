//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// AlarmTrips shall contain properties describing the types of alarms that have
// been raised by the memory. These alarms shall be reset when the system
// resets. Note that if they are re-discovered they can be reasserted.
type AlarmTrips struct {
	// AddressParityError shall be true if an Address Parity Error was detected
	// which could not be corrected by retry.
	AddressParityError bool
	// CorrectableECCError shall be true if the correctable error threshold
	// crossing alarm trip was detected.
	CorrectableECCError bool
	// SpareBlock shall be true if the spare block capacity crossing alarm trip
	// was detected.
	SpareBlock bool
	// Temperature shall be true if a temperature threshold alarm trip was detected.
	Temperature bool
	// UncorrectableECCError shall be true if the uncorrectable error threshold
	// alarm trip was detected.
	UncorrectableECCError bool
}

// CurrentPeriod shall describe the metrics of the memory since last time the
// ClearCurrentPeriod Action was performed or the system reset.
type CurrentPeriod struct {
	// BlocksRead shall be number of blocks read since reset.
	BlocksRead uint
	// BlocksWritten shall be number of blocks written since reset.
	BlocksWritten uint
}

// HealthData shall contain properties which describe the HealthData metrics for
// the current resource.
type HealthData struct {
	// AlarmTrips shall contain properties describe the types of alarms that
	// have been raised by the memory.
	AlarmTrips AlarmTrips
	// DataLossDetected shall be data loss detection status, with true
	// indicating data loss detected.
	DataLossDetected bool
	// LastShutdownSuccess shall be the status of the  last shutdown, with true
	// indicating success.
	LastShutdownSuccess bool
	// PerformanceDegraded shall be performance degraded mode status, with true
	// indicating performance degraded.
	PerformanceDegraded bool
	// PredictedMediaLifeLeftPercent shall contain an indicator
	// of the percentage of life remaining in the media.
	PredictedMediaLifeLeftPercent float32
	// RemainingSpareBlockPercentage shall be the remaining spare blocks in percentage.
	RemainingSpareBlockPercentage float32
}

// LifeTime shall describe the metrics of the memory since manufacturing.
type LifeTime struct {
	// BlocksRead shall be number of blocks read for the lifetime of the Memory.
	BlocksRead uint64
	// BlocksWritten shall be number of blocks written for the lifetime of the Memory.
	BlocksWritten uint64
}

// MemoryMetrics is used to represent the Memory Metrics for a single Memory
// device in a Redfish implementation.
type MemoryMetrics struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// BandwidthPercent shall contain memory bandwidth utilization as a
	// percentage.  When this resource is subordinate to the MemorySummary
	// object, this property shall be the memory bandwidth utilization over all
	// memory as a percentage.
	BandwidthPercent float32
	// BlockSizeBytes shall be the block size in bytes of all structure elements.
	BlockSizeBytes int
	// CurrentPeriod shall contain properties which describe the CurrentPeriod
	// metrics for the current resource.
	CurrentPeriod CurrentPeriod
	// Description provides a description of this resource.
	Description string
	// HealthData shall contain properties which describe the HealthData metrics
	// for the current resource.
	HealthData HealthData
	// LifeTime shall contain properties which describe the LifeTime metrics for
	// the current resource.
	LifeTime LifeTime
	// OperatingSpeedMHz is used by the memory device.
	OperatingSpeedMHz int
}

// GetMemoryMetrics will get a MemoryMetrics instance from the service.
func GetMemoryMetrics(c common.Client, uri string) (*MemoryMetrics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var memorymetrics MemoryMetrics
	err = json.NewDecoder(resp.Body).Decode(&memorymetrics)
	if err != nil {
		return nil, err
	}

	memorymetrics.SetClient(c)
	return &memorymetrics, nil
}

// ListReferencedMemoryMetricss gets the collection of MemoryMetrics from
// a provided reference.
func ListReferencedMemoryMetricss(c common.Client, link string) ([]*MemoryMetrics, error) {
	var result []*MemoryMetrics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, memorymetricsLink := range links.ItemLinks {
		memorymetrics, err := GetMemoryMetrics(c, memorymetricsLink)
		if err != nil {
			return result, err
		}
		result = append(result, memorymetrics)
	}

	return result, nil
}
