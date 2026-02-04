//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/ProcessorMetrics.v1_7_0.json
// 2025.3 - #ProcessorMetrics.v1_7_0.ProcessorMetrics

package schemas

import (
	"encoding/json"
)

// ProcessorMetrics This resource contains the processor metrics for a single
// processor in a Redfish implementation.
type ProcessorMetrics struct {
	Entity
	// AverageFrequencyMHz shall contain average frequency in MHz, across all
	// enabled cores in the processor. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property is not applicable.
	//
	// Deprecated: v1.1.0
	// This property has been deprecated in favor of the 'OperatingSpeedMHz'
	// property.
	AverageFrequencyMHz *float64 `json:",omitempty"`
	// BandwidthPercent shall contain the bandwidth usage of the processor as a
	// percentage, typically '0' to '100'. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property shall be the CPU utilization over
	// all processors as a percentage.
	BandwidthPercent *float64 `json:",omitempty"`
	// Cache shall contain properties that describe this processor's cache. When
	// this resource is subordinate to the 'ProcessorSummary' object, this property
	// is not applicable.
	Cache []CacheMetrics
	// CacheMetricsTotal shall contain properties that describe the metrics for all
	// of the cache memory of this processor.
	//
	// Version added: v1.2.0
	CacheMetricsTotal CacheMetricsTotal
	// ConsumedPowerWatt shall contain the power, in watt units, that the processor
	// has consumed. When this resource is subordinate to the 'ProcessorSummary'
	// object, this property shall be the sum of power, in watt units, that all
	// processors have consumed.
	//
	// Deprecated: v1.2.0
	// This property has been deprecated in favor of the properties in
	// 'EnvironmentMetrics'.
	ConsumedPowerWatt *float64 `json:",omitempty"`
	// CoreMetrics shall contain properties that describe the cores of this
	// processor. When this resource is subordinate to the 'ProcessorSummary'
	// object, this property is not applicable.
	CoreMetrics []CoreMetrics
	// CoreVoltage shall contain the core voltage, in volt units, of this
	// processor. The core voltage of the processor may change more frequently than
	// the manager is able to monitor. The value of the 'DataSourceUri' property,
	// if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'Voltage'.
	//
	// Version added: v1.3.0
	CoreVoltage SensorVoltageExcerpt
	// CorrectableCoreErrorCount shall contain the number of correctable core
	// errors, such as TLB or cache errors. When this resource is subordinate to
	// the 'ProcessorSummary' object, this property shall be the sum of
	// 'CorrectableCoreErrorCount' over all processors.
	//
	// Version added: v1.5.0
	CorrectableCoreErrorCount *int `json:",omitempty"`
	// CorrectableOtherErrorCount shall contain the number of correctable errors of
	// all other components. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property shall be the sum of
	// 'CorrectableOtherErrorCount' over all processors.
	//
	// Version added: v1.5.0
	CorrectableOtherErrorCount *int `json:",omitempty"`
	// FrequencyRatio shall contain the frequency relative to the nominal processor
	// frequency ratio of this processor. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property shall be the average FrequencyRatio
	// over all processors.
	FrequencyRatio *float64 `json:",omitempty"`
	// KernelPercent shall contain total percentage of time, '0' to '100', the
	// processor has spent in kernel mode. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property shall be the average KernelPercent
	// over all processors.
	KernelPercent *float64 `json:",omitempty"`
	// LifetimeStartDateTime shall contain the date and time when the processor
	// started accumulating data for the 'LifeTime' property. This might contain
	// the same value as the production date of the processor.
	//
	// Version added: v1.7.0
	LifetimeStartDateTime string
	// LocalMemoryBandwidthBytes shall contain the local memory bandwidth usage of
	// this processor in bytes. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property shall be the sum of
	// LocalMemoryBandwidthBytes over all processors.
	LocalMemoryBandwidthBytes *int `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OperatingSpeedMHz shall contain the operating speed of the processor in MHz.
	// The operating speed of the processor may change more frequently than the
	// manager is able to monitor.
	//
	// Version added: v1.1.0
	OperatingSpeedMHz *int `json:",omitempty"`
	// PCIeErrors shall contain the PCIe errors associated with this processor.
	//
	// Version added: v1.4.0
	PCIeErrors PCIeErrors
	// PowerLimitThrottleDuration shall contain the total duration of throttling
	// caused by a power limit of the processor since reset.
	//
	// Version added: v1.6.0
	PowerLimitThrottleDuration string
	// RemoteMemoryBandwidthBytes shall contain the remote memory bandwidth usage
	// of this processor in bytes. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property shall be the sum of
	// RemoteMemoryBandwidthBytes over all processors.
	RemoteMemoryBandwidthBytes *int `json:",omitempty"`
	// TemperatureCelsius shall contain the temperature, in degree Celsius units,
	// of the processor. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property shall be the average temperature,
	// in Celsius, over all processors.
	//
	// Deprecated: v1.2.0
	// This property has been deprecated in favor of the properties in
	// 'EnvironmentMetrics'.
	TemperatureCelsius *float64 `json:",omitempty"`
	// ThermalLimitThrottleDuration shall contain the total duration of throttling
	// caused by a thermal limit of the processor since reset.
	//
	// Version added: v1.6.0
	ThermalLimitThrottleDuration string
	// ThrottlingCelsius shall contain the CPU margin to throttle based on an
	// offset between the maximum temperature in which the processor can operate,
	// and the processor's current temperature. When this resource is subordinate
	// to the 'ProcessorSummary' object, this property is not applicable.
	ThrottlingCelsius *float64 `json:",omitempty"`
	// UncorrectableCoreErrorCount shall contain the number of uncorrectable core
	// errors, such as TLB or cache errors. When this resource is subordinate to
	// the 'ProcessorSummary' object, this property shall be the sum of
	// 'UncorrectableCoreErrorCount' over all processors.
	//
	// Version added: v1.5.0
	UncorrectableCoreErrorCount *int `json:",omitempty"`
	// UncorrectableOtherErrorCount shall contain the number of uncorrectable
	// errors of all other components. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property shall be the sum of
	// 'UncorrectableOtherErrorCount' over all processors.
	//
	// Version added: v1.5.0
	UncorrectableOtherErrorCount *int `json:",omitempty"`
	// UserPercent shall contain total percentage of time, '0' to '100', the
	// processor has spent in user mode. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property shall be the average UserPercent
	// over all processors.
	UserPercent *float64 `json:",omitempty"`
	// clearCurrentPeriodTarget is the URL to send ClearCurrentPeriod requests.
	clearCurrentPeriodTarget string
}

// UnmarshalJSON unmarshals a ProcessorMetrics object from the raw JSON.
func (p *ProcessorMetrics) UnmarshalJSON(b []byte) error {
	type temp ProcessorMetrics
	type pActions struct {
		ClearCurrentPeriod ActionTarget `json:"#ProcessorMetrics.ClearCurrentPeriod"`
	}
	var tmp struct {
		temp
		Actions pActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = ProcessorMetrics(tmp.temp)

	// Extract the links to other entities for later
	p.clearCurrentPeriodTarget = tmp.Actions.ClearCurrentPeriod.Target

	return nil
}

// GetProcessorMetrics will get a ProcessorMetrics instance from the service.
func GetProcessorMetrics(c Client, uri string) (*ProcessorMetrics, error) {
	return GetObject[ProcessorMetrics](c, uri)
}

// ListReferencedProcessorMetricss gets the collection of ProcessorMetrics from
// a provided reference.
func ListReferencedProcessorMetricss(c Client, link string) ([]*ProcessorMetrics, error) {
	return GetCollectionObjects[ProcessorMetrics](c, link)
}

// This action shall set the 'CurrentPeriod' property's values to 0.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (p *ProcessorMetrics) ClearCurrentPeriod() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(p.client,
		p.clearCurrentPeriodTarget, payload, p.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// CStateResidency shall contain properties that describe the C-state residency
// of the processor or core.
type CStateResidency struct {
	// Level shall contain the C-state level, such as C0, C1, or C2. When this
	// resource is subordinate to the 'ProcessorSummary' object, this property is
	// not applicable.
	Level string
	// ResidencyPercent shall contain the percentage of time, '0' to '100', that
	// the processor or core has spent in this particular level of C-state. When
	// this resource is subordinate to the 'ProcessorSummary' object, this property
	// is not applicable.
	ResidencyPercent *float64 `json:",omitempty"`
}

// CacheMetrics shall contain properties that describe cache metrics of a
// processor or core.
type CacheMetrics struct {
	// CacheMiss shall contain the number of cache line misses of the processor or
	// core in millions.
	CacheMiss *float64 `json:",omitempty"`
	// CacheMissesPerInstruction shall contain the number of cache misses per
	// instruction of the processor or core.
	CacheMissesPerInstruction *float64 `json:",omitempty"`
	// HitRatio shall contain the cache hit ratio of the processor or core.
	HitRatio *float64 `json:",omitempty"`
	// Level shall contain the level of the cache in the processor or core.
	Level string
	// OccupancyBytes shall contain the total cache occupancy of the processor or
	// core in bytes.
	OccupancyBytes *int `json:",omitempty"`
	// OccupancyPercent shall contain the total cache occupancy percentage, '0' to
	// '100', of the processor or core.
	OccupancyPercent *float64 `json:",omitempty"`
}

// CacheMetricsTotal shall contain properties that describe the metrics for all
// of the cache memory for a processor.
type CacheMetricsTotal struct {
	// CurrentPeriod shall contain properties that describe the metrics for the
	// current period of cache memory for this processor.
	//
	// Version added: v1.2.0
	CurrentPeriod ProcessorMetricsCurrentPeriod
	// LifeTime shall contain properties that describe the metrics for the lifetime
	// of the cache memory for this processor.
	//
	// Version added: v1.2.0
	LifeTime ProcessMetricsLifeTime
}

// CoreMetrics shall contain properties that describe the cores of a processor.
type CoreMetrics struct {
	// CStateResidency shall contain properties that describe the C-state residency
	// of this core in the processor.
	CStateResidency []CStateResidency
	// CoreCache shall contain properties that describe the cache metrics of this
	// core in the processor.
	CoreCache []CacheMetrics
	// CoreID shall contain the identifier of the core within the processor.
	CoreID string `json:"CoreId"`
	// CorrectableCoreErrorCount shall contain the number of correctable core
	// errors, such as TLB or cache errors. When this resource is subordinate to
	// the 'ProcessorSummary' object, this property shall be the sum of
	// 'CorrectableCoreErrorCount' over all processors.
	//
	// Version added: v1.5.0
	CorrectableCoreErrorCount *int `json:",omitempty"`
	// CorrectableOtherErrorCount shall contain the number of correctable errors of
	// all other components. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property shall be the sum of
	// 'CorrectableOtherErrorCount' over all processors.
	//
	// Version added: v1.5.0
	CorrectableOtherErrorCount *int `json:",omitempty"`
	// IOStallCount shall contain the number of stalled cycles due to I/O
	// operations of this core in the processor.
	IOStallCount *int `json:",omitempty"`
	// InstructionsPerCycle shall contain the number of instructions per clock
	// cycle of this core in the processor.
	InstructionsPerCycle *float64 `json:",omitempty"`
	// MemoryStallCount shall contain the number of stalled cycles due to memory
	// operations of this core in the processor.
	MemoryStallCount *int `json:",omitempty"`
	// UncorrectableCoreErrorCount shall contain the number of uncorrectable core
	// errors, such as TLB or cache errors. When this resource is subordinate to
	// the 'ProcessorSummary' object, this property shall be the sum of
	// 'UncorrectableCoreErrorCount' over all processors.
	//
	// Version added: v1.5.0
	UncorrectableCoreErrorCount *int `json:",omitempty"`
	// UncorrectableOtherErrorCount shall contain the number of uncorrectable
	// errors of all other components. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property shall be the sum of
	// 'UncorrectableOtherErrorCount' over all processors.
	//
	// Version added: v1.5.0
	UncorrectableOtherErrorCount *int `json:",omitempty"`
	// UnhaltedCycles shall contain the number of unhalted cycles of this core in
	// the processor.
	UnhaltedCycles *float64 `json:",omitempty"`
}

// ProcessorMetricsCurrentPeriod shall describe the cache memory metrics since last system reset
// or 'ClearCurrentPeriod' action for a processor.
type ProcessorMetricsCurrentPeriod struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors of
	// cache memory since reset or 'ClearCurrentPeriod' action for this processor.
	// When this resource is subordinate to the 'ProcessorSummary' object, this
	// property shall be the sum of 'CorrectableECCErrorCount' over all processors.
	//
	// Version added: v1.2.0
	CorrectableECCErrorCount *int `json:",omitempty"`
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors
	// of cache memory since reset or 'ClearCurrentPeriod' action for this
	// processor. When this resource is subordinate to the 'ProcessorSummary'
	// object, this property shall be the sum of 'UncorrectableECCErrorCount' over
	// all processors.
	//
	// Version added: v1.2.0
	UncorrectableECCErrorCount *int `json:",omitempty"`
}

// ProcessMetricsLifeTime shall describe the cache memory metrics since manufacturing for a
// processor.
type ProcessMetricsLifeTime struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors for
	// the lifetime of the cache memory. When this resource is subordinate to the
	// 'ProcessorSummary' object, this property shall be the sum of
	// 'CorrectableECCErrorCount' over all processors.
	//
	// Version added: v1.2.0
	CorrectableECCErrorCount *int `json:",omitempty"`
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors
	// for the lifetime of the cache memory. When this resource is subordinate to
	// the 'ProcessorSummary' object, this property shall be the sum of
	// 'UncorrectableECCErrorCount' over all processors.
	//
	// Version added: v1.2.0
	UncorrectableECCErrorCount *int `json:",omitempty"`
}
