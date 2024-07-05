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

// AlertCapabilities shall contain the conditions that would generate an alert to the CXL Fabric Manager or host.
type AlertCapabilities struct {
	// CorrectableECCError shall indicate whether correctable ECC errors generate an alert to the CXL Fabric Manager or
	// host.
	CorrectableECCError bool
	// SpareBlock shall indicate whether spare block conditions generate an alert to the CXL Fabric Manager or host.
	SpareBlock bool
	// Temperature shall indicate whether temperature conditions generate an alert to the CXL Fabric Manager or host.
	Temperature bool
	// UncorrectableECCError shall indicate whether uncorrectable ECC errors generate an alert to the CXL Fabric
	// Manager or host.
	UncorrectableECCError bool
}

// CXLMemoryMetrics shall contain the memory metrics specific to CXL devices.
type CXLMemoryMetrics struct {
	// AlertCapabilities shall contain the conditions that would generate an alert to the CXL Fabric Manager or host.
	AlertCapabilities AlertCapabilities
}

// CurrentPeriod shall describe the metrics of the memory since last time the
// ClearCurrentPeriod Action was performed or the system reset.
type CurrentPeriod struct {
	// BlocksRead shall be number of blocks read since reset.
	BlocksRead uint
	// BlocksWritten shall be number of blocks written since reset.
	BlocksWritten uint
	// CorrectableECCErrorCount shall contain the number of correctable errors since reset. When this resource is
	// subordinate to the MemorySummary object, this property shall be the sum of CorrectableECCErrorCount over all
	// memory.
	CorrectableECCErrorCount int
	// IndeterminateCorrectableErrorCount shall contain the number of indeterminate correctable errors since reset.
	// Since the error origin is indeterminate, the same error can be duplicated across multiple MemoryMetrics
	// resources. When this resource is subordinate to the MemorySummary object, this property shall be the sum of
	// indeterminate correctable errors across all memory without duplication, which may not be the sum of all
	// IndeterminateCorrectableErrorCount properties over all memory.
	IndeterminateCorrectableErrorCount int
	// IndeterminateUncorrectableErrorCount shall contain the number of indeterminate uncorrectable errors since reset.
	// Since the error origin is indeterminate, the same error can be duplicated across multiple MemoryMetrics
	// resources. When this resource is subordinate to the MemorySummary object, this property shall be the sum of
	// indeterminate uncorrectable errors across all memory without duplication, which may not be the sum of all
	// IndeterminateUncorrectableErrorCount properties over all memory.
	IndeterminateUncorrectableErrorCount int
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors since reset. When this resource is
	// subordinate to the MemorySummary object, this property shall be the sum of UncorrectableECCErrorCount over all
	// memory.
	UncorrectableECCErrorCount int
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
	// CorrectableECCErrorCount shall contain the number of correctable errors for the lifetime of the memory. When
	// this resource is subordinate to the MemorySummary object, this property shall be the sum of
	// CorrectableECCErrorCount over all memory.
	CorrectableECCErrorCount int
	// IndeterminateCorrectableErrorCount shall contain the number of indeterminate correctable errors for the lifetime
	// of the memory. Since the error origin is indeterminate, the same error can be duplicated across multiple
	// MemoryMetrics resources. When this resource is subordinate to the MemorySummary object, this property shall be
	// the sum of indeterminate correctable errors across all memory without duplication, which may not be the sum of
	// all IndeterminateCorrectableErrorCount properties over all memory.
	IndeterminateCorrectableErrorCount int
	// IndeterminateUncorrectableErrorCount shall contain the number of indeterminate uncorrectable errors for the
	// lifetime of the memory. Since the error origin is indeterminate, the same error can be duplicated across
	// multiple MemoryMetrics resources. When this resource is subordinate to the MemorySummary object, this property
	// shall be the sum of indeterminate uncorrectable errors across all memory without duplication, which may not be
	// the sum of all IndeterminateUncorrectableErrorCount properties over all memory.
	IndeterminateUncorrectableErrorCount int
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors for the lifetime of the memory. When
	// this resource is subordinate to the MemorySummary object, this property shall be the sum of
	// UncorrectableECCErrorCount over all memory.
	UncorrectableECCErrorCount int
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
	// CXL shall contain the memory metrics specific to CXL devices.
	CXL CXLMemoryMetrics
	// CapacityUtilizationPercent shall contain the memory capacity utilization as a percentage, typically '0' to
	// '100'. When this resource is subordinate to the MemorySummary object, this property shall be the memory capacity
	// utilization over all memory as a percentage.
	CapacityUtilizationPercent float64
	// CorrectedPersistentErrorCount shall contain the number of corrected errors in persistent memory.
	CorrectedPersistentErrorCount int
	// CorrectedVolatileErrorCount shall contain the number of corrected errors in volatile memory.
	CorrectedVolatileErrorCount int
	// CurrentPeriod shall contain properties which describe the CurrentPeriod
	// metrics for the current resource.
	CurrentPeriod CurrentPeriod
	// Description provides a description of this resource.
	Description string
	// DirtyShutdownCount shall contain the number of shutdowns while outstanding writes have not completed to
	// persistent memory.
	DirtyShutdownCount int
	// HealthData shall contain properties which describe the HealthData metrics
	// for the current resource.
	HealthData HealthData
	// LifeTime shall contain properties which describe the LifeTime metrics for
	// the current resource.
	LifeTime LifeTime
	// OperatingSpeedMHz is used by the memory device.
	OperatingSpeedMHz int

	clearCurrentPeriodTarget string
}

// UnmarshalJSON unmarshals a MemoryMetrics object from the raw JSON.
func (memorymetrics *MemoryMetrics) UnmarshalJSON(b []byte) error {
	type temp MemoryMetrics
	type Actions struct {
		ClearCurrentPeriod common.ActionTarget `json:"#MemoryMetrics.ClearCurrentPeriod"`
	}
	var t struct {
		temp
		Actions Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*memorymetrics = MemoryMetrics(t.temp)

	// Extract the links to other entities for later
	memorymetrics.clearCurrentPeriodTarget = t.Actions.ClearCurrentPeriod.Target

	return nil
}

// ClearCurrentPeriod sets the CurrentPeriod property's values to 0.
func (memorymetrics *MemoryMetrics) ClearCurrentPeriod() error {
	return memorymetrics.Post(memorymetrics.clearCurrentPeriodTarget, nil)
}

// GetMemoryMetrics will get a MemoryMetrics instance from the service.
func GetMemoryMetrics(c common.Client, uri string) (*MemoryMetrics, error) {
	return common.GetObject[MemoryMetrics](c, uri)
}

// ListReferencedMemoryMetricss gets the collection of MemoryMetrics from
// a provided reference.
func ListReferencedMemoryMetricss(c common.Client, link string) ([]*MemoryMetrics, error) {
	return common.GetCollectionObjects[MemoryMetrics](c, link)
}
