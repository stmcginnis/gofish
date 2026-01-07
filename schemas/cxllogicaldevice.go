//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #CXLLogicalDevice.v1_3_0.CXLLogicalDevice

package schemas

import (
	"encoding/json"
)

type CXLSemantic string

const (
	// CXLioCXLSemantic shall indicate the device conforms with the CXL
	// Specification-defined 'CXL.io' semantic.
	CXLioCXLSemantic CXLSemantic = "CXLio"
	// CXLcacheCXLSemantic shall indicate the device conforms with the CXL
	// Specification-defined 'CXL.cache' semantic.
	CXLcacheCXLSemantic CXLSemantic = "CXLcache"
	// CXLmemCXLSemantic shall indicate the device conforms with the CXL
	// Specification-defined 'CXL.mem' semantic.
	CXLmemCXLSemantic CXLSemantic = "CXLmem"
)

type PassphraseType string

const (
	// UserPassphraseType shall indicate a user-defined passphrase.
	UserPassphraseType PassphraseType = "User"
	// MasterPassphraseType shall indicate an administrator-defined master
	// passphrase.
	MasterPassphraseType PassphraseType = "Master"
)

// CXLLogicalDevice shall represent a CXL logical device that is a part of a
// PCIe device.
type CXLLogicalDevice struct {
	Entity
	// Identifiers shall contain a list of all known durable names for the
	// associated CXL logical device.
	Identifiers []Identifier
	// Log shall contain a link to a resource of type 'LogService'.
	log string
	// MemoryRegions shall contain a link to a resource collection of type
	// 'MemoryRegionCollection' that represents the memory regions associated with
	// this CXL logical device.
	//
	// Version added: v1.1.0
	memoryRegions string
	// MemorySizeMiB shall contain the total memory capacity currently available in
	// this CXL logical device in mebibytes (MiB). This value shall equate to the
	// sum of the dynamic capacity extents and the static capacity assigned to this
	// logical device.
	MemorySizeMiB int
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// QoS shall contain the quality of service configuration for this CXL logical
	// device.
	QoS QoS
	// QoSTelemetryCapabilities shall contain the quality of service telemetry
	// capabilities for this CXL logical device.
	QoSTelemetryCapabilities CXLDeviceQoSTelemetryCapabilities
	// SemanticsSupported shall contain the CXL Specification-defined semantics
	// that are supported by this CXL logical device.
	SemanticsSupported []CXLSemantic
	// Status shall contain any status or health properties of the resource.
	Status Status
	// disablePassphraseTarget is the URL to send DisablePassphrase requests.
	disablePassphraseTarget string
	// freezeSecurityStateTarget is the URL to send FreezeSecurityState requests.
	freezeSecurityStateTarget string
	// passphraseSecureEraseTarget is the URL to send PassphraseSecureErase requests.
	passphraseSecureEraseTarget string
	// setPassphraseTarget is the URL to send SetPassphrase requests.
	setPassphraseTarget string
	// unlockTarget is the URL to send Unlock requests.
	unlockTarget string
	// endpoints are the URIs for Endpoints.
	endpoints []string
	// memoryChunks are the URIs for MemoryChunks.
	memoryChunks []string
	// memoryDomains are the URIs for MemoryDomains.
	memoryDomains []string
	// pCIeFunctions are the URIs for PCIeFunctions.
	pCIeFunctions []string
	// PCIeFunctionsCount is the number of PCIe functions assigned to this CXL logical device.
	PCIeFunctionsCount int
}

// UnmarshalJSON unmarshals a CXLLogicalDevice object from the raw JSON.
func (c *CXLLogicalDevice) UnmarshalJSON(b []byte) error {
	type temp CXLLogicalDevice
	type cActions struct {
		DisablePassphrase     ActionTarget `json:"#CXLLogicalDevice.DisablePassphrase"`
		FreezeSecurityState   ActionTarget `json:"#CXLLogicalDevice.FreezeSecurityState"`
		PassphraseSecureErase ActionTarget `json:"#CXLLogicalDevice.PassphraseSecureErase"`
		SetPassphrase         ActionTarget `json:"#CXLLogicalDevice.SetPassphrase"`
		Unlock                ActionTarget `json:"#CXLLogicalDevice.Unlock"`
	}
	type cLinks struct {
		Endpoints          Links `json:"Endpoints"`
		MemoryChunks       Links `json:"MemoryChunks"`
		MemoryDomains      Links `json:"MemoryDomains"`
		PCIeFunctions      Links `json:"PCIeFunctions"`
		PCIeFunctionsCount int   `json:"PCIeFunctions@odata.count"`
	}
	var tmp struct {
		temp
		Actions       cActions
		Links         cLinks
		Log           Link `json:"Log"`
		MemoryRegions Link `json:"MemoryRegions"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CXLLogicalDevice(tmp.temp)

	// Extract the links to other entities for later
	c.disablePassphraseTarget = tmp.Actions.DisablePassphrase.Target
	c.freezeSecurityStateTarget = tmp.Actions.FreezeSecurityState.Target
	c.passphraseSecureEraseTarget = tmp.Actions.PassphraseSecureErase.Target
	c.setPassphraseTarget = tmp.Actions.SetPassphrase.Target
	c.unlockTarget = tmp.Actions.Unlock.Target
	c.endpoints = tmp.Links.Endpoints.ToStrings()
	c.memoryChunks = tmp.Links.MemoryChunks.ToStrings()
	c.memoryDomains = tmp.Links.MemoryDomains.ToStrings()
	c.pCIeFunctions = tmp.Links.PCIeFunctions.ToStrings()
	c.PCIeFunctionsCount = tmp.Links.PCIeFunctionsCount
	c.log = tmp.Log.String()
	c.memoryRegions = tmp.MemoryRegions.String()

	return nil
}

// GetCXLLogicalDevice will get a CXLLogicalDevice instance from the service.
func GetCXLLogicalDevice(c Client, uri string) (*CXLLogicalDevice, error) {
	return GetObject[CXLLogicalDevice](c, uri)
}

// ListReferencedCXLLogicalDevices gets the collection of CXLLogicalDevice from
// a provided reference.
func ListReferencedCXLLogicalDevices(c Client, link string) ([]*CXLLogicalDevice, error) {
	return GetCollectionObjects[CXLLogicalDevice](c, link)
}

// This action shall disable the passphrase for the CXL logical device.
// passphrase - This property shall contain the passphrase required to complete
// this action.
// passphraseType - This property shall contain the type of passphrase supplied
// for the operation.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *CXLLogicalDevice) DisablePassphrase(passphrase string, passphraseType PassphraseType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["PassphraseType"] = passphraseType
	resp, taskInfo, err := PostWithTask(c.client,
		c.disablePassphraseTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall freeze the security state of the CXL logical device.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *CXLLogicalDevice) FreezeSecurityState() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(c.client,
		c.freezeSecurityStateTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall perform a CXL Specification-defined secure erase of the
// CXL logical device.
// passphrase - This property shall contain the passphrase required to complete
// this action.
// passphraseType - This property shall contain the type of passphrase supplied
// for the operation.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *CXLLogicalDevice) PassphraseSecureErase(passphrase string, passphraseType PassphraseType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["PassphraseType"] = passphraseType
	resp, taskInfo, err := PostWithTask(c.client,
		c.passphraseSecureEraseTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall set the passphrase for the CXL logical device.
// newPassphrase - This parameter shall contain the new passphrase to set for
// the CXL logical device.
// passphrase - This property shall contain the passphrase required to complete
// this action.
// passphraseType - This property shall contain the type of passphrase supplied
// for the operation.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *CXLLogicalDevice) SetPassphrase(newPassphrase string, passphrase string, passphraseType PassphraseType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["NewPassphrase"] = newPassphrase
	payload["Passphrase"] = passphrase
	payload["PassphraseType"] = passphraseType
	resp, taskInfo, err := PostWithTask(c.client,
		c.setPassphraseTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall unlock the CXL logical device.
// passphrase - This property shall contain the passphrase required to complete
// this action.
// passphraseType - This property shall contain the type of passphrase supplied
// for the operation.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *CXLLogicalDevice) Unlock(passphrase string, passphraseType PassphraseType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["PassphraseType"] = passphraseType
	resp, taskInfo, err := PostWithTask(c.client,
		c.unlockTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// Endpoints gets the Endpoints linked resources.
func (c *CXLLogicalDevice) Endpoints() ([]*Endpoint, error) {
	return GetObjects[Endpoint](c.client, c.endpoints)
}

// MemoryChunks gets the MemoryChunks linked resources.
func (c *CXLLogicalDevice) MemoryChunks() ([]*MemoryChunks, error) {
	return GetObjects[MemoryChunks](c.client, c.memoryChunks)
}

// MemoryDomains gets the MemoryDomains linked resources.
func (c *CXLLogicalDevice) MemoryDomains() ([]*MemoryDomain, error) {
	return GetObjects[MemoryDomain](c.client, c.memoryDomains)
}

// PCIeFunctions gets the PCIeFunctions linked resources.
func (c *CXLLogicalDevice) PCIeFunctions() ([]*PCIeFunction, error) {
	return GetObjects[PCIeFunction](c.client, c.pCIeFunctions)
}

// Log gets the Log linked resource.
func (c *CXLLogicalDevice) Log() (*LogService, error) {
	if c.log == "" {
		return nil, nil
	}
	return GetObject[LogService](c.client, c.log)
}

// MemoryRegions gets the MemoryRegions collection.
func (c *CXLLogicalDevice) MemoryRegions() ([]*MemoryRegion, error) {
	if c.memoryRegions == "" {
		return nil, nil
	}
	return GetCollectionObjects[MemoryRegion](c.client, c.memoryRegions)
}

// QoS shall contain the quality of service properties of this CXL logical
// device.
type QoS struct {
	// AllocatedBandwidth shall contain the bandwidth allocated, '0' to '100', for
	// this CXL logical device as a percentage.
	AllocatedBandwidth *uint `json:",omitempty"`
	// LimitPercent shall contain the bandwidth limit, '0' to '100', for this CXL
	// logical device as a percentage.
	LimitPercent *uint `json:",omitempty"`
}

// CXLDeviceQoSTelemetryCapabilities shall contain the quality of service telemetry
// capabilities for a CXL logical device.
type CXLDeviceQoSTelemetryCapabilities struct {
	// EgressPortBackpressureSupported shall indicate whether the device supports
	// the CXL Specification-defined 'Egress Port Backpressure' mechanism.
	EgressPortBackpressureSupported bool
	// TemporaryThroughputReductionSupported shall indicate whether the device
	// supports the CXL Specification-defined 'Temporary Throughput Reduction'
	// mechanism.
	//
	// Deprecated: v1.2.0
	// This property has been deprecated in favor of
	// 'TemporaryThroughputReductionSupported' in 'PCIeDevice'.
	TemporaryThroughputReductionSupported bool
}
