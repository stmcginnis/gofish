//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"github.com/stmcginnis/gofish/common"
)

// EGCriticalWarningSummary shall contain the NVMe-defined 'Endurance Group Critical Warning Summary'.
type EGCriticalWarningSummary struct {
	// NamespacesInReadOnlyMode shall indicate whether namespaces in one or more Endurance Groups are in read-only mode
	// not as a result of a change in the write protection state of a namespace.
	NamespacesInReadOnlyMode bool
	// ReliabilityDegraded shall indicate whether the reliability of one or more Endurance Groups is degraded due to
	// significant media-related errors or any internal error that degrades the NVM subsystem reliability.
	ReliabilityDegraded bool
	// SpareCapacityUnderThreshold shall indicate whether the available spare capacity of one or more Endurance Groups
	// is below the threshold.
	SpareCapacityUnderThreshold bool
}

// NVMeSMARTMetrics shall contain the NVMe SMART metrics as defined by the NVMe SMART/Health Information log page.
type NVMeSMARTMetrics struct {
	// AvailableSparePercent shall contain the NVMe-defined 'Available Spare', which represents the normalized
	// percentage, '0' to '100', of the remaining spare capacity available.
	AvailableSparePercent float64
	// AvailableSpareThresholdPercent shall contain the NVMe-defined 'Available Spare Threshold' as a percentage, '0'
	// to '100'. When the available spare falls below this value, an asynchronous event completion may occur.
	AvailableSpareThresholdPercent float64
	// CompositeTemperatureCelsius shall contain the composite temperature in degree Celsius units for this storage
	// controller. Services shall derive this value from the NVMe-defined 'Composite Temperature', which represents a
	// composite temperature in kelvin units of the controller and namespaces associated with that controller.
	CompositeTemperatureCelsius float64
	// ControllerBusyTimeMinutes shall contain the NVMe-defined 'Controller Busy Time', which represents the total time
	// the controller is busy with I/O commands in minutes.
	ControllerBusyTimeMinutes int
	// CriticalCompositeTempTimeMinutes shall contain the NVMe-defined 'Critical Composite Temperature Time', which
	// represents the amount of time in minutes that the controller has been operational and that the composite
	// temperature has been greater than or equal to the critical composite temperature threshold.
	CriticalCompositeTempTimeMinutes int
	// CriticalWarnings shall contain the NVMe-defined 'Critical Warning'.
	CriticalWarnings NVMeSMARTCriticalWarnings
	// DataUnitsRead shall contain the NVMe-defined 'Data Units Read', which represents the number of 512 byte data
	// units the host has read from the controller as part of processing a SMART Data Units Read Command in units of
	// one thousand.
	DataUnitsRead int
	// DataUnitsWritten shall contain the NVMe-defined 'Data Units Written', which represents the number of 512 byte
	// data units the host has written to the controller as part of processing a User Data Out Command in units of one
	// thousand.
	DataUnitsWritten int
	// EGCriticalWarningSummary shall contain the NVMe-defined 'Endurance Group Critical Warning Summary'.
	EGCriticalWarningSummary EGCriticalWarningSummary
	// HostReadCommands shall contain the NVMe-defined 'Host Read Commands', which represents the number of SMART Host
	// Read Commands completed by the controller.
	HostReadCommands int
	// HostWriteCommands shall contain the NVMe-defined 'Host Write Commands', which represents the number of User Data
	// Out Commands completed by the controller.
	HostWriteCommands int
	// MediaAndDataIntegrityErrors shall contain the NVMe-defined 'Media and Data Integrity Errors', which represents
	// the number of occurrences where the controller detected an unrecovered data integrity error.
	MediaAndDataIntegrityErrors int
	// NumberOfErrorInformationLogEntries shall contain the NVMe-defined 'Number of Error Information Log Entries',
	// which represents the number of error information log entries over the life of the controller.
	NumberOfErrorInformationLogEntries int
	// PercentageUsed shall contain the NVMe-defined 'Percentage Used', which represents a vendor-specific estimate of
	// the percentage of the NVM subsystem life used based on the actual usage and the manufacturer's prediction of NVM
	// life. A value of '100' indicates that the estimated endurance of the NVM in the NVM subsystem has been consumed,
	// but this may not indicate an NVM subsystem failure. The value is allowed to exceed '100'. Percentages greater
	// than '254' shall be represented as '255'.
	PercentageUsed float64
	// PowerCycles shall contain the NVMe-defined 'Power Cycles', which represents the number of power cycles.
	PowerCycles int
	// PowerOnHours shall contain the NVMe-defined 'Power On Hours', which represents the number of power-on hours.
	PowerOnHours float64
	// TemperatureSensorsCelsius shall contain an array of temperature sensor readings in degree Celsius units for this
	// storage controller. Services shall derive each array member from the NVMe-defined 'Temperature Sensor' values,
	// which represent a temperature sensor reading in kelvin units.
	TemperatureSensorsCelsius []int
	// ThermalMgmtTemp1TotalTimeSeconds shall contain the NVMe-defined 'Total Time For Thermal Management Temperature
	// 1', which represents the number of seconds the controller transitioned to lower power states or performed
	// vendor-specific thermal-management actions while minimizing the impact on performance in order to attempt to
	// reduce the composite temperature.
	ThermalMgmtTemp1TotalTimeSeconds int
	// ThermalMgmtTemp1TransitionCount shall contain the NVMe-defined 'Thermal Management Temperature 1 Transition
	// Count', which represents the number of times the controller transitioned to lower power states or performed
	// vendor-specific thermal-management actions while minimizing the impact on performance in order to attempt to
	// reduce the composite temperature.
	ThermalMgmtTemp1TransitionCount int
	// ThermalMgmtTemp2TotalTimeSeconds shall contain the NVMe-defined 'Total Time For Thermal Management Temperature
	// 2', which represents the number of seconds the controller transitioned to lower power states or performed
	// vendor-specific thermal-management actions regardless of the impact on performance in order to attempt to reduce
	// the composite temperature.
	ThermalMgmtTemp2TotalTimeSeconds int
	// ThermalMgmtTemp2TransitionCount shall contain the NVMe-defined 'Thermal Management Temperature 2 Transition
	// Count', which represents the number of times the controller transitioned to lower power states or performed
	// vendor-specific thermal-management actions regardless of the impact on performance in order to attempt to reduce
	// the composite temperature.
	ThermalMgmtTemp2TransitionCount int
	// UnsafeShutdowns shall contain the NVMe-defined 'Unsafe Shutdowns', which represents the number of times when the
	// controller does not report it is safe to power down prior to loss of main power.
	UnsafeShutdowns int
	// WarningCompositeTempTimeMinutes shall contain the NVMe-defined 'Warning Composite Temperature Time', which
	// represents the amount of time in minutes that the controller has been operational and that the composite
	// temperature has been greater than or equal to the warning composite temperature threshold.
	WarningCompositeTempTimeMinutes int
}

// StorageControllerMetrics shall contain the usage and health statistics for a storage controller in a Redfish
// implementation.
type StorageControllerMetrics struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// CorrectableECCErrorCount shall contain the number of correctable errors for the lifetime of memory of the
	// storage controller.
	CorrectableECCErrorCount int
	// CorrectableParityErrorCount shall contain the number of correctable errors for the lifetime of memory of the
	// storage controller.
	CorrectableParityErrorCount int
	// Description provides a description of this resource.
	Description string
	// NVMeSMART shall contain the NVMe SMART metrics for this storage controller as defined by the NVMe SMART/Health
	// Information log page. This property shall only be present for NVMe storage controllers.
	NVMeSMART NVMeSMARTMetrics
	// StateChangeCount shall contain the number of times the State property within the Status property of the parent
	// StorageController resource changed.
	StateChangeCount int
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors for the lifetime of memory of the
	// storage controller.
	UncorrectableECCErrorCount int
	// UncorrectableParityErrorCount shall contain the number of uncorrectable errors for the lifetime of memory of the
	// storage controller.
	UncorrectableParityErrorCount int
}

// GetStorageControllerMetrics will get a StorageControllerMetrics instance from the service.
func GetStorageControllerMetrics(c common.Client, uri string) (*StorageControllerMetrics, error) {
	return common.GetObject[StorageControllerMetrics](c, uri)
}

// ListReferencedStorageControllerMetrics gets the collection of StorageControllerMetrics from
// a provided reference.
func ListReferencedStorageControllerMetrics(c common.Client, link string) ([]*StorageControllerMetrics, error) {
	return common.GetCollectionObjects[StorageControllerMetrics](c, link)
}
