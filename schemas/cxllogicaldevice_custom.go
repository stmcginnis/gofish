//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a CXLLogicalDevice object from the raw JSON.
func (c *CXLLogicalDevice) UnmarshalJSON(b []byte) error {
	type temp CXLLogicalDevice
	type cActions struct {
		DisablePassphrase      ActionTarget `json:"#CXLLogicalDevice.DisablePassphrase"`
		DynamicCapacityAdd     ActionTarget `json:"#CXLLogicalDevice.DynamicCapacityAdd"`
		DynamicCapacityRelease ActionTarget `json:"#CXLLogicalDevice.DynamicCapacityRelease"`
		FreezeSecurityState    ActionTarget `json:"#CXLLogicalDevice.FreezeSecurityState"`
		PassphraseSecureErase  ActionTarget `json:"#CXLLogicalDevice.PassphraseSecureErase"`
		SetPassphrase          ActionTarget `json:"#CXLLogicalDevice.SetPassphrase"`
		Unlock                 ActionTarget `json:"#CXLLogicalDevice.Unlock"`
	}
	type cLinks struct {
		Endpoints          Links `json:"Endpoints"`
		MemoryChunks       Links `json:"MemoryChunks"`
		MemoryDomains      Links `json:"MemoryDomains"`
		PCIeFunctions      Links `json:"PCIeFunctions"`
		PCIeFunctionsCount int   `json:"PCIeFunctions@odata.count"`
		Port               Link  `json:"Port"`
	}
	var tmp struct {
		temp
		Actions       cActions
		Links         cLinks
		Log           ResolvedLink[LogService] `json:"Log"`
		MemoryRegions Link                     `json:"MemoryRegions"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CXLLogicalDevice(tmp.temp)

	// Extract the links to other entities for later
	c.disablePassphraseTarget = tmp.Actions.DisablePassphrase.Target
	c.dynamicCapacityAddTarget = tmp.Actions.DynamicCapacityAdd.Target
	c.dynamicCapacityReleaseTarget = tmp.Actions.DynamicCapacityRelease.Target
	c.freezeSecurityStateTarget = tmp.Actions.FreezeSecurityState.Target
	c.passphraseSecureEraseTarget = tmp.Actions.PassphraseSecureErase.Target
	c.setPassphraseTarget = tmp.Actions.SetPassphrase.Target
	c.unlockTarget = tmp.Actions.Unlock.Target
	c.endpoints = tmp.Links.Endpoints.ToStrings()
	c.memoryChunks = tmp.Links.MemoryChunks.ToStrings()
	c.memoryDomains = tmp.Links.MemoryDomains.ToStrings()
	c.pCIeFunctions = tmp.Links.PCIeFunctions.ToStrings()
	c.PCIeFunctionsCount = tmp.Links.PCIeFunctionsCount
	c.port = tmp.Links.Port.String()
	c.log = tmp.Log

	c.memoryRegions = tmp.MemoryRegions.String()

	return nil
}
