//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/ManagerDiagnosticData.v1_2_3.json
// 2022.3 - #ManagerDiagnosticData.v1_2_3.ManagerDiagnosticData

package schemas

import (
	"encoding/json"
)

// ManagerDiagnosticData shall represent internal diagnostic data for a manager
// for a Redfish implementation. Clients should not make decisions for raising
// alerts, creating service events, or other actions based on information in
// this resource.
type ManagerDiagnosticData struct {
	Entity
	// BootTimeStatistics shall contain the boot-time statistics of the manager.
	BootTimeStatistics BootTimeStatistics
	// FreeStorageSpaceKiB shall contain the available storage space on this
	// manager in kibibytes (KiB).
	FreeStorageSpaceKiB *int `json:",omitempty"`
	// I2CBuses shall contain the statistics of the I2C buses. Services may
	// subdivide a physical bus into multiple entries in this property based on how
	// the manager tracks bus segments, virtual buses from a controller, and other
	// segmentation capabilities.
	I2CBuses []I2CBusStatistics
	// MemoryECCStatistics shall contain the memory ECC statistics of the manager.
	MemoryECCStatistics MemoryECCStatistics
	// MemoryStatistics shall contain the memory statistics of the manager.
	MemoryStatistics MemoryStatistics
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// ProcessorStatistics shall contain the processor statistics of the manager.
	ProcessorStatistics ProcessorStatistics
	// ServiceRootUptimeSeconds shall contain the wall-clock time the service root
	// hosted by this manager has been running in seconds.
	//
	// Version added: v1.2.0
	ServiceRootUptimeSeconds *float64 `json:",omitempty"`
	// TopProcesses shall contain the statistics of the top processes of this
	// manager.
	TopProcesses []ProcessStatistics
	// resetMetricsTarget is the URL to send ResetMetrics requests.
	resetMetricsTarget string
}

// UnmarshalJSON unmarshals a ManagerDiagnosticData object from the raw JSON.
func (m *ManagerDiagnosticData) UnmarshalJSON(b []byte) error {
	type temp ManagerDiagnosticData
	type mActions struct {
		ResetMetrics ActionTarget `json:"#ManagerDiagnosticData.ResetMetrics"`
	}
	var tmp struct {
		temp
		Actions mActions
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ManagerDiagnosticData(tmp.temp)

	// Extract the links to other entities for later
	m.resetMetricsTarget = tmp.Actions.ResetMetrics.Target

	return nil
}

// GetManagerDiagnosticData will get a ManagerDiagnosticData instance from the service.
func GetManagerDiagnosticData(c Client, uri string) (*ManagerDiagnosticData, error) {
	return GetObject[ManagerDiagnosticData](c, uri)
}

// ListReferencedManagerDiagnosticDatas gets the collection of ManagerDiagnosticData from
// a provided reference.
func ListReferencedManagerDiagnosticDatas(c Client, link string) ([]*ManagerDiagnosticData, error) {
	return GetCollectionObjects[ManagerDiagnosticData](c, link)
}

// This action shall reset any time intervals or counted values of the
// diagnostic data for this manager.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (m *ManagerDiagnosticData) ResetMetrics() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(m.client,
		m.resetMetricsTarget, payload, m.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// BootTimeStatistics shall contain the boot-time statistics of a manager.
type BootTimeStatistics struct {
	// FirmwareTimeSeconds shall contain the number of seconds the manager spent in
	// the firmware stage.
	FirmwareTimeSeconds *float64 `json:",omitempty"`
	// InitrdTimeSeconds shall contain the number of seconds the manager spent in
	// the initrd boot stage.
	InitrdTimeSeconds *float64 `json:",omitempty"`
	// KernelTimeSeconds shall contain the number of seconds the manager spent in
	// the kernel stage.
	KernelTimeSeconds *float64 `json:",omitempty"`
	// LoaderTimeSeconds shall contain the number of seconds the manager spent in
	// the loader stage.
	LoaderTimeSeconds *float64 `json:",omitempty"`
	// UserSpaceTimeSeconds shall contain the number of seconds the manager spent
	// in the user space boot stage.
	UserSpaceTimeSeconds *float64 `json:",omitempty"`
}

// I2CBusStatistics shall contain statistics of an I2C bus.
type I2CBusStatistics struct {
	// BusErrorCount shall contain the number of bus errors on this I2C bus. Bus
	// errors include, but are not limited to, an SDA rising or falling edge while
	// SCL is high or a stuck bus signal.
	BusErrorCount *int `json:",omitempty"`
	// I2CBusName shall contain the name of the I2C bus.
	I2CBusName string
	// NACKCount shall contain the number of NACKs on this I2C bus.
	NACKCount *int `json:",omitempty"`
	// TotalTransactionCount shall contain the total number of transactions on this
	// I2C bus. The count shall include the number of I2C transactions initiated by
	// the manager and the number of I2C transactions where the manager is the
	// target device.
	TotalTransactionCount *int `json:",omitempty"`
}

// MemoryECCStatistics shall contain the memory ECC statistics of a manager.
type MemoryECCStatistics struct {
	// CorrectableECCErrorCount shall contain the number of correctable errors
	// since reset.
	CorrectableECCErrorCount *int `json:",omitempty"`
	// UncorrectableECCErrorCount shall contain the number of uncorrectable errors
	// since reset.
	UncorrectableECCErrorCount *int `json:",omitempty"`
}

// MemoryStatistics shall contain the memory statistics of a manager.
type MemoryStatistics struct {
	// AvailableBytes shall contain the amount of memory available in bytes for
	// starting new processes without swapping. This includes free memory and
	// reclaimable cache and buffers.
	AvailableBytes *int `json:",omitempty"`
	// BuffersAndCacheBytes shall contain the amount of memory used in bytes by
	// kernel buffers, page caches, and slabs.
	BuffersAndCacheBytes *int `json:",omitempty"`
	// FreeBytes shall contain the amount of free memory in bytes.
	FreeBytes *int `json:",omitempty"`
	// SharedBytes shall contain the amount of shared memory in bytes. This
	// includes things such as memory consumed by temporary file systems.
	SharedBytes *int `json:",omitempty"`
	// TotalBytes shall contain the total amount of memory in bytes.
	TotalBytes *int `json:",omitempty"`
	// UsedBytes shall contain the amount of used memory in bytes. This value is
	// calculated as 'TotalBytes' minus 'FreeBytes' minus 'BuffersAndCacheBytes'.
	UsedBytes *int `json:",omitempty"`
}

// ProcessStatistics shall contain the statistics of a process running on a
// manager.
type ProcessStatistics struct {
	// CommandLine shall contain the command line with parameters of this process.
	CommandLine string
	// KernelTimeSeconds shall contain the number of seconds this process executed
	// in kernel space.
	KernelTimeSeconds *float64 `json:",omitempty"`
	// ResidentSetSizeBytes shall contain the resident set size of this process in
	// bytes, which is the amount of memory allocated to the process and is in RAM.
	ResidentSetSizeBytes *int `json:",omitempty"`
	// RestartAfterFailureCount shall contain the number of times this process has
	// restarted unexpectedly, such as due to unintentional failures, restarts, or
	// shutdowns, with the same command line including arguments.
	//
	// Version added: v1.1.0
	RestartAfterFailureCount *int `json:",omitempty"`
	// RestartCount shall contain the number of times this process has restarted
	// with the same command line including arguments.
	//
	// Version added: v1.1.0
	RestartCount *int `json:",omitempty"`
	// UptimeSeconds shall contain the wall-clock time this process has been
	// running in seconds.
	//
	// Version added: v1.1.0
	UptimeSeconds *float64 `json:",omitempty"`
	// UserTimeSeconds shall contain the number of seconds this process executed in
	// user space.
	UserTimeSeconds *float64 `json:",omitempty"`
}

// ProcessorStatistics shall contain the processor statistics of a manager.
type ProcessorStatistics struct {
	// KernelPercent shall contain the percentage of CPU time, '0' to '100', spent
	// in kernel mode.
	KernelPercent *float64 `json:",omitempty"`
	// UserPercent shall contain the percentage of CPU time, '0' to '100', spent in
	// user mode.
	UserPercent *float64 `json:",omitempty"`
}
