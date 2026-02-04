//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/MemoryMetrics.v1_8_0.json
// 2025.3 - #MemoryMetrics.v1_8_0.MemoryMetrics

package schemas

import (
	"encoding/json"
)

// MemoryMetrics shall represent the memory metrics for a memory device or
// system memory summary in a Redfish implementation.
type MemoryMetrics struct {
	Entity
	// BandwidthPercent shall contain memory bandwidth utilization as a percentage.
	// When this resource is subordinate to the 'MemorySummary' object, this
	// property shall be the memory bandwidth utilization over all memory as a
	// percentage, typically '0' to '100'.
	//
	// Version added: v1.2.0
	BandwidthPercent *float64 `json:",omitempty"`
	// BlockSizeBytes shall contain the block size, in bytes, of all structure
	// elements. When this resource is subordinate to the 'MemorySummary' object,
	// this property is not applicable.
	BlockSizeBytes *int `json:",omitempty"`
	// CXL shall contain the memory metrics specific to CXL devices.
	//
	// Version added: v1.6.0
	CXL MemoryMetricsCXL
	// CapacityUtilizationPercent shall contain the memory capacity utilization as
	// a percentage, typically '0' to '100'. When this resource is subordinate to
	// the 'MemorySummary' object, this property shall be the memory capacity
	// utilization over all memory as a percentage.
	//
	// Version added: v1.7.0
	CapacityUtilizationPercent *float64 `json:",omitempty"`
	// CorrectedPersistentErrorCount shall contain the number of corrected errors
	// in persistent memory.
	//
	// Version added: v1.6.0
	CorrectedPersistentErrorCount *int `json:",omitempty"`
	// CorrectedVolatileErrorCount shall contain the number of corrected errors in
	// volatile memory.
	//
	// Version added: v1.6.0
	CorrectedVolatileErrorCount *int `json:",omitempty"`
	// CurrentPeriod shall contain properties that describe the memory metrics for
	// the current period.
	CurrentPeriod MemoryMetricsCurrentPeriod
	// DirtyShutdownCount shall contain the number of shutdowns while outstanding
	// writes have not completed to persistent memory.
	//
	// Version added: v1.6.0
	DirtyShutdownCount *int `json:",omitempty"`
	// HealthData shall contain properties that describe the health data memory
	// metrics for the memory.
	HealthData HealthData
	// LifeTime shall contain properties that describe the memory metrics for the
	// lifetime of the memory.
	LifeTime MemoryMetricsLifeTime
	// LifetimeStartDateTime shall contain the date and time when the memory
	// started accumulating data for the 'LifeTime' property. This might contain
	// the same value as the production date of the memory.
	//
	// Version added: v1.8.0
	LifetimeStartDateTime string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OperatingSpeedMHz shall contain the operating speed of memory in MHz or MT/s
	// (mega-transfers per second) as reported by the memory device. Memory devices
	// that operate at their bus speed shall report the operating speed in MHz (bus
	// speed), while memory devices that transfer data faster than their bus speed,
	// such as DDR memory, shall report the operating speed in MT/s
	// (mega-transfers/second). The reported value shall match the conventionally
	// reported values for the technology used by the memory device.
	//
	// Version added: v1.3.0
	OperatingSpeedMHz *int `json:",omitempty"`
	// clearCurrentPeriodTarget is the URL to send ClearCurrentPeriod requests.
	clearCurrentPeriodTarget string
}

// UnmarshalJSON unmarshals a MemoryMetrics object from the raw JSON.
func (m *MemoryMetrics) UnmarshalJSON(b []byte) error {
	type temp MemoryMetrics
	type mActions struct {
		ClearCurrentPeriod ActionTarget `json:"#MemoryMetrics.ClearCurrentPeriod"`
	}
	var tmp struct {
		temp
		Actions mActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MemoryMetrics(tmp.temp)

	// Extract the links to other entities for later
	m.clearCurrentPeriodTarget = tmp.Actions.ClearCurrentPeriod.Target

	return nil
}

// GetMemoryMetrics will get a MemoryMetrics instance from the service.
func GetMemoryMetrics(c Client, uri string) (*MemoryMetrics, error) {
	return GetObject[MemoryMetrics](c, uri)
}

// ListReferencedMemoryMetricss gets the collection of MemoryMetrics from
// a provided reference.
func ListReferencedMemoryMetricss(c Client, link string) ([]*MemoryMetrics, error) {
	return GetCollectionObjects[MemoryMetrics](c, link)
}

// This action shall set the 'CurrentPeriod' property's values to 0.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *MemoryMetrics) ClearCurrentPeriod() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(m.client,
		m.clearCurrentPeriodTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// AlarmTrips shall contain properties that describe the types of alarms that
// have been raised by the memory. These alarms shall be reset when the system
// resets. Note that if they are re-discovered they can be reasserted.
type AlarmTrips struct {
	// AddressParityError shall indicate whether an address parity error was
	// detected that a retry could not correct.
	AddressParityError bool
	// CorrectableECCError shall indicate whether the correctable error threshold
	// crossing alarm trip was detected.
	CorrectableECCError bool
	// SpareBlock shall indicate whether the spare block capacity crossing alarm
	// trip was detected.
	SpareBlock bool
	// Temperature shall indicate whether a temperature threshold alarm trip was
	// detected.
	Temperature bool
	// UncorrectableECCError shall indicate whether the uncorrectable error
	// threshold alarm trip was detected.
	UncorrectableECCError bool
}

// AlertCapabilities shall contain the conditions that would generate an alert
// to the CXL Fabric Manager or host.
type AlertCapabilities struct {
	// CorrectableECCError shall indicate whether correctable ECC errors generate
	// an alert to the CXL Fabric Manager or host.
	//
	// Version added: v1.6.0
	CorrectableECCError bool
	// SpareBlock shall indicate whether spare block conditions generate an alert
	// to the CXL Fabric Manager or host.
	//
	// Version added: v1.6.0
	SpareBlock bool
	// Temperature shall indicate whether temperature conditions generate an alert
	// to the CXL Fabric Manager or host.
	//
	// Version added: v1.6.0
	Temperature bool
	// UncorrectableECCError shall indicate whether uncorrectable ECC errors
	// generate an alert to the CXL Fabric Manager or host.
	//
	// Version added: v1.6.0
	UncorrectableECCError bool
}

// MemoryMetricsCXL shall contain the memory metrics specific to CXL devices.
type MemoryMetricsCXL struct {
	// AlertCapabilities shall contain the conditions that would generate an alert
	// to the CXL Fabric Manager or host.
	//
	// Version added: v1.6.0
	AlertCapabilities AlertCapabilities
}

// MemoryMetricsCurrentPeriod shall describe the memory metrics since last system reset or
// 'ClearCurrentPeriod' action.
type MemoryMetricsCurrentPeriod struct {
	// BlocksRead shall contain the number of blocks read since reset. When this
	// resource is subordinate to the 'MemorySummary' object, this property shall
	// be the sum of BlocksRead over all memory.
	BlocksRead *int `json:",omitempty"`
	// BlocksWritten shall contain the number of blocks written since reset. When
	// this resource is subordinate to the 'MemorySummary' object, this property
	// shall be the sum of BlocksWritten over all memory.
	BlocksWritten *int `json:",omitempty"`
	// CorrectableECCErrorCount shall contain the number of correctable errors
	// since reset. When this resource is subordinate to the 'MemorySummary'
	// object, this property shall be the sum of CorrectableECCErrorCount over all
	// memory.
	//
	// Version added: v1.4.0
	CorrectableECCErrorCount *int `json:",omitempty"`
	// IndeterminateCorrectableErrorCount shall contain the number of indeterminate
	// correctable errors since reset. Since the error origin is indeterminate, the
	// same error can be duplicated across multiple 'MemoryMetrics' resources. When
	// this resource is subordinate to the 'MemorySummary' object, this property
	// shall be the sum of indeterminate correctable errors across all memory
	// without duplication, which may not be the sum of all
	// 'IndeterminateCorrectableErrorCount' properties over all memory.
	//
	// Version added: v1.5.0
	IndeterminateCorrectableErrorCount *int `json:",omitempty"`
	// IndeterminateUncorrectableErrorCount shall contain the number of
	// indeterminate uncorrectable errors since reset. Since the error origin is
	// indeterminate, the same error can be duplicated across multiple
	// 'MemoryMetrics' resources. When this resource is subordinate to the
	// 'MemorySummary' object, this property shall be the sum of indeterminate
	// uncorrectable errors across all memory without duplication, which may not be
	// the sum of all 'IndeterminateUncorrectableErrorCount' properties over all
	// memory.
	//
	// Version added: v1.5.0
	IndeterminateUncorrectableErrorCount *int `json:",omitempty"`
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors
	// since reset. When this resource is subordinate to the 'MemorySummary'
	// object, this property shall be the sum of UncorrectableECCErrorCount over
	// all memory.
	//
	// Version added: v1.4.0
	UncorrectableECCErrorCount *int `json:",omitempty"`
}

// HealthData shall contain properties that describe the health information
// metrics for this resource.
type HealthData struct {
	// AlarmTrips shall contain properties describe the types of alarms that have
	// been raised by the memory. When this resource is subordinate to the
	// 'MemorySummary' object, this property shall indicate whether an alarm of a
	// given type have been raised by any area of memory.
	AlarmTrips AlarmTrips
	// DataLossDetected shall indicate whether data loss was detected. When this
	// resource is subordinate to the 'MemorySummary' object, this property shall
	// indicate whether any data loss was detected in any area of memory.
	DataLossDetected bool
	// LastShutdownSuccess shall indicate whether the last shutdown succeeded.
	LastShutdownSuccess bool
	// PerformanceDegraded shall indicate whether performance has degraded. When
	// this resource is subordinate to the 'MemorySummary' object, this property
	// shall indicate whether degraded performance mode status is detected in any
	// area of memory.
	PerformanceDegraded bool
	// PredictedMediaLifeLeftPercent shall contain an indicator of the percentage,
	// '0' to '100', of life remaining in the media.
	//
	// Version added: v1.1.0
	PredictedMediaLifeLeftPercent *float64 `json:",omitempty"`
	// RemainingSpareBlockPercentage shall contain the remaining spare blocks as a
	// percentage, '0' to '100'. When this resource is subordinate to the
	// 'MemorySummary' object, this property shall be the
	// RemainingSpareBlockPercentage over all memory.
	RemainingSpareBlockPercentage *float64 `json:",omitempty"`
}

// MemoryMetricsLifeTime shall describe the memory metrics since manufacturing.
type MemoryMetricsLifeTime struct {
	// BlocksRead shall contain the number of blocks read for the lifetime of the
	// memory. When this resource is subordinate to the 'MemorySummary' object,
	// this property shall be the sum of BlocksRead over all memory.
	BlocksRead *int `json:",omitempty"`
	// BlocksWritten shall contain the number of blocks written for the lifetime of
	// the memory. When this resource is subordinate to the 'MemorySummary' object,
	// this property shall be the sum of BlocksWritten over all memory.
	BlocksWritten *int `json:",omitempty"`
	// CorrectableECCErrorCount shall contain the number of correctable errors for
	// the lifetime of the memory. When this resource is subordinate to the
	// 'MemorySummary' object, this property shall be the sum of
	// CorrectableECCErrorCount over all memory.
	//
	// Version added: v1.4.0
	CorrectableECCErrorCount *int `json:",omitempty"`
	// IndeterminateCorrectableErrorCount shall contain the number of indeterminate
	// correctable errors for the lifetime of the memory. Since the error origin is
	// indeterminate, the same error can be duplicated across multiple
	// 'MemoryMetrics' resources. When this resource is subordinate to the
	// 'MemorySummary' object, this property shall be the sum of indeterminate
	// correctable errors across all memory without duplication, which may not be
	// the sum of all 'IndeterminateCorrectableErrorCount' properties over all
	// memory.
	//
	// Version added: v1.5.0
	IndeterminateCorrectableErrorCount *int `json:",omitempty"`
	// IndeterminateUncorrectableErrorCount shall contain the number of
	// indeterminate uncorrectable errors for the lifetime of the memory. Since the
	// error origin is indeterminate, the same error can be duplicated across
	// multiple 'MemoryMetrics' resources. When this resource is subordinate to the
	// 'MemorySummary' object, this property shall be the sum of indeterminate
	// uncorrectable errors across all memory without duplication, which may not be
	// the sum of all 'IndeterminateUncorrectableErrorCount' properties over all
	// memory.
	//
	// Version added: v1.5.0
	IndeterminateUncorrectableErrorCount *int `json:",omitempty"`
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors
	// for the lifetime of the memory. When this resource is subordinate to the
	// 'MemorySummary' object, this property shall be the sum of
	// UncorrectableECCErrorCount over all memory.
	//
	// Version added: v1.4.0
	UncorrectableECCErrorCount *int `json:",omitempty"`
}
