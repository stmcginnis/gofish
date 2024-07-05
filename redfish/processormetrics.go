//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// CStateResidency shall contain properties that describe the C-state residency of the processor or core.
type CStateResidency struct {
	// Level shall contain the C-state level, such as C0, C1, or C2. When this resource is subordinate to the
	// ProcessorSummary object, this property is not applicable.
	Level string
	// ResidencyPercent shall contain the percentage of time, '0' to '100', that the processor or core has spent in
	// this particular level of C-state. When this resource is subordinate to the ProcessorSummary object, this
	// property is not applicable.
	ResidencyPercent float64
}

// CacheMetrics shall contain properties that describe cache metrics of a processor or core.
type CacheMetrics struct {
	// CacheMiss shall contain the number of cache line misses of the processor or core in millions.
	CacheMiss float64
	// CacheMissesPerInstruction shall contain the number of cache misses per instruction of the processor or core.
	CacheMissesPerInstruction float64
	// HitRatio shall contain the cache hit ratio of the processor or core.
	HitRatio float64
	// Level shall contain the level of the cache in the processor or core.
	Level string
	// OccupancyBytes shall contain the total cache occupancy of the processor or core in bytes.
	OccupancyBytes int
	// OccupancyPercent shall contain the total cache occupancy percentage, '0' to '100', of the processor or core.
	OccupancyPercent float64
}

// CacheMetricsTotal shall contain properties that describe the metrics for all of the cache memory for a
// processor.
type CacheMetricsTotal struct {
	// CurrentPeriod shall contain properties that describe the metrics for the current period of cache memory for this
	// processor.
	CurrentPeriod CurrentPeriod
	// LifeTime shall contain properties that describe the metrics for the lifetime of the cache memory for this
	// processor.
	LifeTime ProcessorMetricsLifeTime
}

// CoreMetrics shall contain properties that describe the cores of a processor.
type CoreMetrics struct {
	// CStateResidency shall contain properties that describe the C-state residency of this core in the processor.
	CStateResidency []CStateResidency
	// CoreCache shall contain properties that describe the cache metrics of this core in the processor.
	CoreCache []CacheMetrics
	// CoreID shall contain the identifier of the core within the processor.
	CoreID string
	// CorrectableCoreErrorCount shall contain the number of correctable core errors, such as TLB or cache errors. When
	// this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// CorrectableCoreErrorCount over all processors.
	CorrectableCoreErrorCount int
	// CorrectableOtherErrorCount shall contain the number of correctable errors of all other components. When this
	// resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// CorrectableOtherErrorCount over all processors.
	CorrectableOtherErrorCount int
	// IOStallCount shall contain the number of stalled cycles due to I/O operations of this core in the processor.
	IOStallCount int
	// InstructionsPerCycle shall contain the number of instructions per clock cycle of this core in the processor.
	InstructionsPerCycle float64
	// MemoryStallCount shall contain the number of stalled cycles due to memory operations of this core in the
	// processor.
	MemoryStallCount int
	// UncorrectableCoreErrorCount shall contain the number of uncorrectable core errors, such as TLB or cache errors.
	// When this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// UncorrectableCoreErrorCount over all processors.
	UncorrectableCoreErrorCount int
	// UncorrectableOtherErrorCount shall contain the number of uncorrectable errors of all other components. When this
	// resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// UncorrectableOtherErrorCount over all processors.
	UncorrectableOtherErrorCount int
	// UnhaltedCycles shall contain the number of unhalted cycles of this core in the processor.
	UnhaltedCycles float64
}

// ProcessorMetricsCurrentPeriod shall describe the cache memory metrics since last system reset or ClearCurrentPeriod action for a
// processor.
type ProcessorMetricsCurrentPeriod struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors of cache memory since reset or
	// ClearCurrentPeriod action for this processor. When this resource is subordinate to the ProcessorSummary object,
	// this property shall be the sum of CorrectableECCErrorCount over all processors.
	CorrectableECCErrorCount int
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors of cache memory since reset or
	// ClearCurrentPeriod action for this processor. When this resource is subordinate to the ProcessorSummary object,
	// this property shall be the sum of UncorrectableECCErrorCount over all processors.
	UncorrectableECCErrorCount int
}

// ProcessorMetricsLifeTime shall describe the cache memory metrics since manufacturing for a processor.
type ProcessorMetricsLifeTime struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors for the lifetime of the cache memory.
	// When this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// CorrectableECCErrorCount over all processors.
	CorrectableECCErrorCount int
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors for the lifetime of the cache
	// memory. When this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// UncorrectableECCErrorCount over all processors.
	UncorrectableECCErrorCount int
}

// ProcessorMetrics This resource contains the processor metrics for a single processor in a Redfish
// implementation.
type ProcessorMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// BandwidthPercent shall contain the bandwidth usage of the processor as a percentage, typically '0' to '100'.
	// When this resource is subordinate to the ProcessorSummary object, this property shall be the CPU utilization
	// over all processors as a percentage.
	BandwidthPercent float64
	// Cache shall contain properties that describe this processor's cache. When this resource is subordinate to the
	// ProcessorSummary object, this property is not applicable.
	Cache []CacheMetrics
	// CacheMetricsTotal shall contain properties that describe the metrics for all of the cache memory of this
	// processor.
	CacheMetricsTotal CacheMetricsTotal
	// CoreMetrics shall contain properties that describe the cores of this processor. When this resource is
	// subordinate to the ProcessorSummary object, this property is not applicable.
	CoreMetrics []CoreMetrics
	// CoreVoltage shall contain the core voltage, in volt units, of this processor. The core voltage of the processor
	// may change more frequently than the manager is able to monitor. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value 'Voltage'.
	CoreVoltage SensorVoltageExcerpt
	// CorrectableCoreErrorCount shall contain the number of correctable core errors, such as TLB or cache errors. When
	// this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// CorrectableCoreErrorCount over all processors.
	CorrectableCoreErrorCount int
	// CorrectableOtherErrorCount shall contain the number of correctable errors of all other components. When this
	// resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// CorrectableOtherErrorCount over all processors.
	CorrectableOtherErrorCount int
	// Description provides a description of this resource.
	Description string
	// FrequencyRatio shall contain the frequency relative to the nominal processor frequency ratio of this processor.
	// When this resource is subordinate to the ProcessorSummary object, this property shall be the average
	// FrequencyRatio over all processors.
	FrequencyRatio float64
	// KernelPercent shall contain total percentage of time, '0' to '100', the processor has spent in kernel mode. When
	// this resource is subordinate to the ProcessorSummary object, this property shall be the average KernelPercent
	// over all processors.
	KernelPercent float64
	// LocalMemoryBandwidthBytes shall contain the local memory bandwidth usage of this processor in bytes. When this
	// resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// LocalMemoryBandwidthBytes over all processors.
	LocalMemoryBandwidthBytes int
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OperatingSpeedMHz shall contain the operating speed of the processor in MHz. The operating speed of the
	// processor may change more frequently than the manager is able to monitor.
	OperatingSpeedMHz int
	// PCIeErrors shall contain the PCIe errors associated with this processor.
	PCIeErrors PCIeErrors
	// PowerLimitThrottleDuration shall contain the total duration of throttling caused by a power limit of the
	// processor since reset.
	PowerLimitThrottleDuration string
	// RemoteMemoryBandwidthBytes shall contain the remote memory bandwidth usage of this processor in bytes. When this
	// resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// RemoteMemoryBandwidthBytes over all processors.
	RemoteMemoryBandwidthBytes int
	// ThermalLimitThrottleDuration shall contain the total duration of throttling caused by a thermal limit of the
	// processor since reset.
	ThermalLimitThrottleDuration string
	// ThrottlingCelsius shall contain the CPU margin to throttle based on an offset between the maximum temperature in
	// which the processor can operate, and the processor's current temperature. When this resource is subordinate to
	// the ProcessorSummary object, this property is not applicable.
	ThrottlingCelsius float64
	// UncorrectableCoreErrorCount shall contain the number of uncorrectable core errors, such as TLB or cache errors.
	// When this resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// UncorrectableCoreErrorCount over all processors.
	UncorrectableCoreErrorCount int
	// UncorrectableOtherErrorCount shall contain the number of uncorrectable errors of all other components. When this
	// resource is subordinate to the ProcessorSummary object, this property shall be the sum of
	// UncorrectableOtherErrorCount over all processors.
	UncorrectableOtherErrorCount int
	// UserPercent shall contain total percentage of time, '0' to '100', the processor has spent in user mode. When
	// this resource is subordinate to the ProcessorSummary object, this property shall be the average UserPercent over
	// all processors.
	UserPercent float64

	clearCurrentPeriodTarget string
}

// UnmarshalJSON unmarshals a ProcessorMetrics object from the raw JSON.
func (processormetrics *ProcessorMetrics) UnmarshalJSON(b []byte) error {
	type temp ProcessorMetrics
	type Actions struct {
		ClearCurrentPeriod common.ActionTarget `json:"#ProcessorMetrics.ClearCurrentPeriod"`
	}
	var t struct {
		temp
		Actions Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*processormetrics = ProcessorMetrics(t.temp)

	// Extract the links to other entities for later
	processormetrics.clearCurrentPeriodTarget = t.Actions.ClearCurrentPeriod.Target

	return nil
}

// ClearCurrentPeriod sets the CurrentPeriod property's values to 0.
func (processormetrics *ProcessorMetrics) ClearCurrentPeriod() error {
	return processormetrics.Post(processormetrics.clearCurrentPeriodTarget, nil)
}

// GetProcessorMetrics will get a ProcessorMetrics instance from the service.
func GetProcessorMetrics(c common.Client, uri string) (*ProcessorMetrics, error) {
	return common.GetObject[ProcessorMetrics](c, uri)
}

// ListReferencedProcessorMetricss gets the collection of ProcessorMetrics from
// a provided reference.
func ListReferencedProcessorMetricss(c common.Client, link string) ([]*ProcessorMetrics, error) {
	return common.GetCollectionObjects[ProcessorMetrics](c, link)
}
