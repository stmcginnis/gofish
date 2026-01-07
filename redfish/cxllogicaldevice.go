//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #CXLLogicalDevice.v1_3_0.CXLLogicalDevice

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
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
	common.Entity
	// Identifiers shall contain a list of all known durable names for the
	// associated CXL logical device.
	Identifiers []common.Identifier
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
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// QoS shall contain the quality of service configuration for this CXL logical
	// device.
	QoS QoS
	// QoSTelemetryCapabilities shall contain the quality of service telemetry
	// capabilities for this CXL logical device.
	QoSTelemetryCapabilities QoSTelemetryCapabilities
	// SemanticsSupported shall contain the CXL Specification-defined semantics
	// that are supported by this CXL logical device.
	SemanticsSupported []CXLSemantic
	// Status shall contain any status or health properties of the resource.
	Status common.Status
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
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a CXLLogicalDevice object from the raw JSON.
func (c *CXLLogicalDevice) UnmarshalJSON(b []byte) error {
	type temp CXLLogicalDevice
	type cActions struct {
		DisablePassphrase     common.ActionTarget `json:"#CXLLogicalDevice.DisablePassphrase"`
		FreezeSecurityState   common.ActionTarget `json:"#CXLLogicalDevice.FreezeSecurityState"`
		PassphraseSecureErase common.ActionTarget `json:"#CXLLogicalDevice.PassphraseSecureErase"`
		SetPassphrase         common.ActionTarget `json:"#CXLLogicalDevice.SetPassphrase"`
		Unlock                common.ActionTarget `json:"#CXLLogicalDevice.Unlock"`
	}
	type cLinks struct {
		Endpoints     common.Links `json:"Endpoints"`
		MemoryChunks  common.Links `json:"MemoryChunks"`
		MemoryDomains common.Links `json:"MemoryDomains"`
		PCIeFunctions common.Links `json:"PCIeFunctions"`
	}
	var tmp struct {
		temp
		Actions       cActions
		Links         cLinks
		Log           common.Link `json:"log"`
		MemoryRegions common.Link `json:"memoryRegions"`
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
	c.log = tmp.Log.String()
	c.memoryRegions = tmp.MemoryRegions.String()

	// This is a read/write object, so we need to save the raw object data for later
	c.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *CXLLogicalDevice) Update() error {
	readWriteFields := []string{
		"Identifiers",
		"QoS",
		"QoSTelemetryCapabilities",
		"Status",
	}

	return c.UpdateFromRawData(c, c.rawData, readWriteFields)
}

// GetCXLLogicalDevice will get a CXLLogicalDevice instance from the service.
func GetCXLLogicalDevice(c common.Client, uri string) (*CXLLogicalDevice, error) {
	return common.GetObject[CXLLogicalDevice](c, uri)
}

// ListReferencedCXLLogicalDevices gets the collection of CXLLogicalDevice from
// a provided reference.
func ListReferencedCXLLogicalDevices(c common.Client, link string) ([]*CXLLogicalDevice, error) {
	return common.GetCollectionObjects[CXLLogicalDevice](c, link)
}

// DisablePassphrase shall disable the passphrase for the CXL logical device.
// passphrase - This property shall contain the passphrase required to complete
// this action.
// passphraseType - This property shall contain the type of passphrase supplied
// for the operation.
func (c *CXLLogicalDevice) DisablePassphrase(passphrase string, passphraseType PassphraseType) error {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["PassphraseType"] = passphraseType
	return c.Post(c.disablePassphraseTarget, payload)
}

// FreezeSecurityState shall freeze the security state of the CXL logical device.
func (c *CXLLogicalDevice) FreezeSecurityState() error {
	payload := make(map[string]any)
	return c.Post(c.freezeSecurityStateTarget, payload)
}

// PassphraseSecureErase shall perform a CXL Specification-defined secure erase of the
// CXL logical device.
// passphrase - This property shall contain the passphrase required to complete
// this action.
// passphraseType - This property shall contain the type of passphrase supplied
// for the operation.
func (c *CXLLogicalDevice) PassphraseSecureErase(passphrase string, passphraseType PassphraseType) error {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["PassphraseType"] = passphraseType
	return c.Post(c.passphraseSecureEraseTarget, payload)
}

// SetPassphrase shall set the passphrase for the CXL logical device.
// newPassphrase - This parameter shall contain the new passphrase to set for
// the CXL logical device.
// passphrase - This property shall contain the passphrase required to complete
// this action.
// passphraseType - This property shall contain the type of passphrase supplied
// for the operation.
func (c *CXLLogicalDevice) SetPassphrase(newPassphrase string, passphrase string, passphraseType PassphraseType) error {
	payload := make(map[string]any)
	payload["NewPassphrase"] = newPassphrase
	payload["Passphrase"] = passphrase
	payload["PassphraseType"] = passphraseType
	return c.Post(c.setPassphraseTarget, payload)
}

// Unlock shall unlock the CXL logical device.
// passphrase - This property shall contain the passphrase required to complete
// this action.
// passphraseType - This property shall contain the type of passphrase supplied
// for the operation.
func (c *CXLLogicalDevice) Unlock(passphrase string, passphraseType PassphraseType) error {
	payload := make(map[string]any)
	payload["Passphrase"] = passphrase
	payload["PassphraseType"] = passphraseType
	return c.Post(c.unlockTarget, payload)
}

// Endpoints gets the Endpoints linked resources.
func (c *CXLLogicalDevice) Endpoints(client common.Client) ([]*Endpoint, error) {
	return common.GetObjects[Endpoint](client, c.endpoints)
}

// MemoryChunks gets the MemoryChunks linked resources.
func (c *CXLLogicalDevice) MemoryChunks(client common.Client) ([]*MemoryChunks, error) {
	return common.GetObjects[MemoryChunks](client, c.memoryChunks)
}

// MemoryDomains gets the MemoryDomains linked resources.
func (c *CXLLogicalDevice) MemoryDomains(client common.Client) ([]*MemoryDomain, error) {
	return common.GetObjects[MemoryDomain](client, c.memoryDomains)
}

// PCIeFunctions gets the PCIeFunctions linked resources.
func (c *CXLLogicalDevice) PCIeFunctions(client common.Client) ([]*PCIeFunction, error) {
	return common.GetObjects[PCIeFunction](client, c.pCIeFunctions)
}

// Log gets the Log linked resource.
func (c *CXLLogicalDevice) Log(client common.Client) (*LogService, error) {
	if c.log == "" {
		return nil, nil
	}
	return common.GetObject[LogService](client, c.log)
}

// MemoryRegions gets the MemoryRegions collection.
func (c *CXLLogicalDevice) MemoryRegions(client common.Client) ([]*MemoryRegion, error) {
	if c.memoryRegions == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[MemoryRegion](client, c.memoryRegions)
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
