//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// BootTimeStatistics shall contain the boot-time statistics of a manager.
type BootTimeStatistics struct {
	// FirmwareTimeSeconds shall contain the number of seconds the manager spent in the firmware stage.
	FirmwareTimeSeconds float64
	// InitrdTimeSeconds shall contain the number of seconds the manager spent in the initrd boot stage.
	InitrdTimeSeconds float64
	// KernelTimeSeconds shall contain the number of seconds the manager spent in the kernel stage.
	KernelTimeSeconds float64
	// LoaderTimeSeconds shall contain the number of seconds the manager spent in the loader stage.
	LoaderTimeSeconds float64
	// UserSpaceTimeSeconds shall contain the number of seconds the manager spent in the user space boot stage.
	UserSpaceTimeSeconds float64
}

// I2CBusStatistics shall contain statistics of an I2C bus.
type I2CBusStatistics struct {
	// BusErrorCount shall contain the number of bus errors on this I2C bus. Bus errors include, but are not limited
	// to, an SDA rising or falling edge while SCL is high or a stuck bus signal.
	BusErrorCount int
	// I2CBusName shall contain the name of the I2C bus.
	I2CBusName string
	// NACKCount shall contain the number of NACKs on this I2C bus.
	NACKCount int
	// TotalTransactionCount shall contain the total number of transactions on this I2C bus. The count shall include
	// the number of I2C transactions initiated by the manager and the number of I2C transactions where the manager is
	// the target device.
	TotalTransactionCount int
}

// ManagerDiagnosticData shall represent internal diagnostic data for a manager for a Redfish implementation.
// Clients should not make decisions for raising alerts, creating service events, or other actions based on
// information in this resource.
type ManagerDiagnosticData struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// BootTimeStatistics shall contain the boot-time statistics of the manager.
	BootTimeStatistics BootTimeStatistics
	// Description provides a description of this resource.
	Description string
	// FreeStorageSpaceKiB shall contain the available storage space on this manager in kibibytes (KiB).
	FreeStorageSpaceKiB int
	// I2CBuses shall contain the statistics of the I2C buses. Services may subdivide a physical bus into multiple
	// entries in this property based on how the manager tracks bus segments, virtual buses from a controller, and
	// other segmentation capabilities.
	I2CBuses []I2CBusStatistics
	// MemoryECCStatistics shall contain the memory ECC statistics of the manager.
	MemoryECCStatistics MemoryECCStatistics
	// MemoryStatistics shall contain the memory statistics of the manager.
	MemoryStatistics MemoryStatistics
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// ProcessorStatistics shall contain the processor statistics of the manager.
	ProcessorStatistics ProcessorStatistics
	// ServiceRootUptimeSeconds shall contain the wall-clock time the service root hosted by this manager has been
	// running in seconds.
	ServiceRootUptimeSeconds float64
	// TopProcesses shall contain the statistics of the top processes of this manager.
	TopProcesses []ProcessStatistics

	resetMetricsTarget string
}

// UnmarshalJSON unmarshals a ManagerDiagnosticData object from the raw JSON.
func (managerdiagnosticdata *ManagerDiagnosticData) UnmarshalJSON(b []byte) error {
	type temp ManagerDiagnosticData
	type Actions struct {
		ResetMetrics common.ActionTarget `json:"#ManagerDiagnosticData.ResetMetrics"`
	}
	var t struct {
		temp
		Actions Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*managerdiagnosticdata = ManagerDiagnosticData(t.temp)

	// Extract the links to other entities for later
	managerdiagnosticdata.resetMetricsTarget = t.Actions.ResetMetrics.Target

	return nil
}

// ResetMetrics resets time intervals or counted values of the diagnostic data for this manager.
func (manager *Manager) ResetMetrics() error {
	return manager.Post(manager.resetToDefaultsTarget, nil)
}

// GetManagerDiagnosticData will get a ManagerDiagnosticData instance from the service.
func GetManagerDiagnosticData(c common.Client, uri string) (*ManagerDiagnosticData, error) {
	return common.GetObject[ManagerDiagnosticData](c, uri)
}

// ListReferencedManagerDiagnosticDatas gets the collection of ManagerDiagnosticData from
// a provided reference.
func ListReferencedManagerDiagnosticDatas(c common.Client, link string) ([]*ManagerDiagnosticData, error) {
	return common.GetCollectionObjects[ManagerDiagnosticData](c, link)
}

// MemoryECCStatistics shall contain the memory ECC statistics of a manager.
type MemoryECCStatistics struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors since reset.
	CorrectableECCErrorCount int
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors since reset.
	UncorrectableECCErrorCount int
}

// MemoryStatistics shall contain the memory statistics of a manager.
type MemoryStatistics struct {
	// AvailableBytes shall contain the amount of memory available in bytes for starting new processes without
	// swapping. This includes free memory and reclaimable cache and buffers.
	AvailableBytes int
	// BuffersAndCacheBytes shall contain the amount of memory used in bytes by kernel buffers, page caches, and slabs.
	BuffersAndCacheBytes int
	// FreeBytes shall contain the amount of free memory in bytes.
	FreeBytes int
	// SharedBytes shall contain the amount of shared memory in bytes. This includes things such as memory consumed by
	// temporary file systems.
	SharedBytes int
	// TotalBytes shall contain the total amount of memory in bytes.
	TotalBytes int
	// UsedBytes shall contain the amount of used memory in bytes. This value is calculated as TotalBytes minus
	// FreeBytes minus BuffersAndCacheBytes.
	UsedBytes int
}

// ProcessStatistics shall contain the statistics of a process running on a manager.
type ProcessStatistics struct {
	// CommandLine shall contain the command line with parameters of this process.
	CommandLine string
	// KernelTimeSeconds shall contain the number of seconds this process executed in kernel space.
	KernelTimeSeconds float64
	// ResidentSetSizeBytes shall contain the resident set size of this process in bytes, which is the amount of memory
	// allocated to the process and is in RAM.
	ResidentSetSizeBytes int
	// RestartAfterFailureCount shall contain the number of times this process has restarted unexpectedly, such as due
	// to unintentional failures, restarts, or shutdowns, with the same command line including arguments.
	RestartAfterFailureCount int
	// RestartCount shall contain the number of times this process has restarted with the same command line including
	// arguments.
	RestartCount int
	// UptimeSeconds shall contain the wall-clock time this process has been running in seconds.
	UptimeSeconds float64
	// UserTimeSeconds shall contain the number of seconds this process executed in user space.
	UserTimeSeconds float64
}

// ProcessorStatistics shall contain the processor statistics of a manager.
type ProcessorStatistics struct {
	// KernelPercent shall contain the percentage of CPU time, '0' to '100', spent in kernel mode.
	KernelPercent float64
	// UserPercent shall contain the percentage of CPU time, '0' to '100', spent in user mode.
	UserPercent float64
}
