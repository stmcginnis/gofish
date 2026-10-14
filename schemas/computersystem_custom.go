//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// UnmarshalJSON unmarshals a ComputerSystem object from the raw JSON.
//
//nolint:funlen
func (c *ComputerSystem) UnmarshalJSON(b []byte) error {
	type temp ComputerSystem
	type cActions struct {
		AddResourceBlock                   ActionTarget `json:"#ComputerSystem.AddResourceBlock"`
		Decommission                       ActionTarget `json:"#ComputerSystem.Decommission"`
		ExportConfiguration                ActionTarget `json:"#ComputerSystem.ExportConfiguration"`
		GraphicalConsoleGetOneTimeURL      ActionTarget `json:"#ComputerSystem.GraphicalConsoleGetOneTimeURL"`
		GraphicalConsoleSetOneTimePassword ActionTarget `json:"#ComputerSystem.GraphicalConsoleSetOneTimePassword"`
		GraphicalConsoleSetPassword        ActionTarget `json:"#ComputerSystem.GraphicalConsoleSetPassword"`
		RemoveResourceBlock                ActionTarget `json:"#ComputerSystem.RemoveResourceBlock"`
		Reset                              struct {
			ActionTarget
			AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
		} `json:"#ComputerSystem.Reset"`
		SetDefaultBootOrder ActionTarget `json:"#ComputerSystem.SetDefaultBootOrder"`
	}
	type cLinks struct {
		Chassis                         Links `json:"Chassis"`
		ConsumingComputerSystems        Links `json:"ConsumingComputerSystems"`
		CooledBy                        Links `json:"CooledBy"`
		Endpoints                       Links `json:"Endpoints"`
		HostingComputerSystem           Link  `json:"HostingComputerSystem"`
		ManagedBy                       Links `json:"ManagedBy"`
		OffloadedNetworkDeviceFunctions Links `json:"OffloadedNetworkDeviceFunctions"`
		PoweredBy                       Links `json:"PoweredBy"`
		ResourceBlocks                  Links `json:"ResourceBlocks"`
		SupplyingComputerSystems        Links `json:"SupplyingComputerSystems"`
		TrustedComponents               Links `json:"TrustedComponents"`
		VirtualMachines                 Links `json:"VirtualMachines"`
	}
	var tmp struct {
		temp
		Actions             cActions
		Links               cLinks
		Bios                ResolvedLink[Bios]            `json:"Bios"`
		Certificates        Link                          `json:"Certificates"`
		EthernetInterfaces  Link                          `json:"EthernetInterfaces"`
		FabricAdapters      Link                          `json:"FabricAdapters"`
		GraphicsControllers Link                          `json:"GraphicsControllers"`
		LogServices         Link                          `json:"LogServices"`
		Memory              Link                          `json:"Memory"`
		MemoryDomains       Link                          `json:"MemoryDomains"`
		NetworkInterfaces   Link                          `json:"NetworkInterfaces"`
		OperatingSystem     ResolvedLink[OperatingSystem] `json:"OperatingSystem"`
		PCIeDevices         Links                         `json:"PCIeDevices"`
		PCIeFunctions       Links                         `json:"PCIeFunctions"`
		Processors          Link                          `json:"Processors"`
		Redundancy          ResolvedLink[Redundancy]      `json:"Redundancy"`
		SecureBoot          ResolvedLink[SecureBoot]      `json:"SecureBoot"`
		SimpleStorage       Link                          `json:"SimpleStorage"`
		Storage             Link                          `json:"Storage"`
		USBControllers      Link                          `json:"USBControllers"`
		VirtualMedia        Link                          `json:"VirtualMedia"`
		Settings            Settings                      `json:"@Redfish.Settings"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = ComputerSystem(tmp.temp)

	// Extract the links to other entities for later
	c.addResourceBlockTarget = tmp.Actions.AddResourceBlock.Target
	c.decommissionTarget = tmp.Actions.Decommission.Target
	c.exportConfigurationTarget = tmp.Actions.ExportConfiguration.Target
	c.graphicalConsoleGetOneTimeURLTarget = tmp.Actions.GraphicalConsoleGetOneTimeURL.Target
	c.graphicalConsoleSetOneTimePasswordTarget = tmp.Actions.GraphicalConsoleSetOneTimePassword.Target
	c.graphicalConsoleSetPasswordTarget = tmp.Actions.GraphicalConsoleSetPassword.Target
	c.removeResourceBlockTarget = tmp.Actions.RemoveResourceBlock.Target
	c.resetTarget = tmp.Actions.Reset.Target
	c.resetActionInfoTarget = tmp.Actions.Reset.ActionInfoTarget
	c.supportedResetTypes = tmp.Actions.Reset.AllowedResetTypes

	// Some implementations use a @Redfish.Settings object to direct settings updates to a
	// different URL than the object being updated. Others don't, so handle both.
	c.settingsTarget = tmp.Settings.SettingsObject
	if c.settingsTarget == "" {
		c.settingsTarget = c.ODataID
	}
	c.SettingsApplyTimes = tmp.Settings.SupportedApplyTimes

	c.setDefaultBootOrderTarget = tmp.Actions.SetDefaultBootOrder.Target
	c.chassis = tmp.Links.Chassis.ToStrings()
	c.consumingComputerSystems = tmp.Links.ConsumingComputerSystems.ToStrings()
	c.cooledBy = tmp.Links.CooledBy.ToStrings()
	c.endpoints = tmp.Links.Endpoints.ToStrings()
	c.hostingComputerSystem = tmp.Links.HostingComputerSystem.String()
	c.managedBy = tmp.Links.ManagedBy.ToStrings()
	c.offloadedNetworkDeviceFunctions = tmp.Links.OffloadedNetworkDeviceFunctions.ToStrings()
	c.poweredBy = tmp.Links.PoweredBy.ToStrings()
	c.resourceBlocks = tmp.Links.ResourceBlocks.ToStrings()
	c.supplyingComputerSystems = tmp.Links.SupplyingComputerSystems.ToStrings()
	c.trustedComponents = tmp.Links.TrustedComponents.ToStrings()
	c.virtualMachines = tmp.Links.VirtualMachines.ToStrings()
	c.bios = tmp.Bios

	c.certificates = tmp.Certificates.String()
	c.ethernetInterfaces = tmp.EthernetInterfaces.String()
	c.fabricAdapters = tmp.FabricAdapters.String()
	c.graphicsControllers = tmp.GraphicsControllers.String()
	c.logServices = tmp.LogServices.String()
	c.memory = tmp.Memory.String()
	c.memoryDomains = tmp.MemoryDomains.String()
	c.networkInterfaces = tmp.NetworkInterfaces.String()
	c.operatingSystem = tmp.OperatingSystem

	c.pCIeDevices = tmp.PCIeDevices.ToStrings()
	c.pCIeFunctions = tmp.PCIeFunctions.ToStrings()
	c.processors = tmp.Processors.String()
	c.redundancy = tmp.Redundancy

	c.secureBoot = tmp.SecureBoot

	c.simpleStorage = tmp.SimpleStorage.String()
	c.storage = tmp.Storage.String()
	c.uSBControllers = tmp.USBControllers.String()
	c.virtualMedia = tmp.VirtualMedia.String()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// If @Redfish.Settings is present, PATCH goes to the System resource first.
// Falls back to the Settings URI with ETag concurrency control only if the
// direct PATCH fails (e.g., Dell iDRAC10 returns read-only).
func (c *ComputerSystem) SetBoot(b *Boot) error {
	t := struct {
		Boot *Boot
	}{Boot: b}

	// Try the main System resource first.
	err := c.Patch(c.ODataID, t)
	if err == nil {
		return nil
	}

	// Fall back to the Settings resource with ETag concurrency control.
	if c.settingsTarget != c.ODataID {
		resp, getErr := c.GetClient().Get(c.settingsTarget)
		defer DeferredCleanupHTTPResponse(resp)
		if getErr != nil {
			return fmt.Errorf("system resource: %w, settings resource: %w", err, getErr)
		}

		var header = make(map[string]string)
		if resp.Header["Etag"] != nil {
			header["If-Match"] = sanitizeETag(resp.Header["Etag"][0])
		}

		resp, patchErr := c.GetClient().PatchWithHeaders(c.settingsTarget, t, header)
		defer DeferredCleanupHTTPResponse(resp)
		if patchErr != nil {
			return fmt.Errorf("system resource: %w, settings resource: %w", err, patchErr)
		}

		return nil
	}

	return err
}

// This action shall reset the system represented by the resource. For systems
// that implement ACPI Power Button functionality, the 'PushPowerButton' value
// shall perform or emulate an ACPI Power Button Push, and the 'ForceOff' value
// shall perform an ACPI Power Button Override, commonly known as a four-second
// hold of the power button.
// resetType - This parameter shall contain the type of reset. The service can
// accept a request without the parameter and perform an
// implementation-specific default reset. Services should include the
// '@Redfish.AllowableValues' annotation for this parameter to ensure
// compatibility with clients, even when 'ActionInfo' has been implemented.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *ComputerSystem) Reset(resetType ResetType) (*TaskMonitorInfo, error) {
	// Make sure the requested reset type is supported by the system
	valid := false
	if len(c.supportedResetTypes) > 0 {
		for _, allowed := range c.supportedResetTypes {
			if resetType == allowed {
				valid = true
				break
			}
		}
	} else {
		// No allowed values supplied, assume we are OK
		valid = true
	}

	if !valid {
		return nil, fmt.Errorf("reset type '%s' is not supported by this service",
			resetType)
	}

	t := struct {
		ResetType ResetType
	}{ResetType: resetType}
	resp, taskInfo, err := PostWithTask(c.client,
		c.resetTarget, t, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// GetSupportedResetTypes returns any reset types that the ComputerSystem declares as supported
// via either ActionInfo or AllowableValues.
func (c *ComputerSystem) GetSupportedResetTypes() ([]ResetType, error) {
	if len(c.supportedResetTypes) > 0 {
		return c.supportedResetTypes, nil
	}

	// if we don't have ResetTypes, try to get from ActionInfo
	if c.resetActionInfoTarget != "" {
		resetActionInfo, err := c.ResetActionInfo()
		if err != nil {
			return c.supportedResetTypes, err
		}

		vals, err := resetActionInfo.GetParamValues("ResetType", StringParameterTypes)
		if err != nil {
			return c.supportedResetTypes, err
		}

		for _, val := range vals {
			c.supportedResetTypes = append(c.supportedResetTypes, ResetType(val))
		}
	}

	return c.supportedResetTypes, nil
}

// ResetActionInfo returns the ActionInfo for the ComputerSystem reset action if supported
func (c *ComputerSystem) ResetActionInfo() (*ActionInfo, error) {
	if c.resetActionInfoTarget == "" {
		return nil, errors.New("ComputerSystem Reset ActionInfo not supported")
	}

	return GetObject[ActionInfo](c.GetClient(), c.resetActionInfoTarget)
}

// UpdateBootAttributesApplyAt is used to update attribute values and set apply time together
func (c *ComputerSystem) UpdateBootAttributesApplyAt(attrs SettingsAttributes, applyTime SettingsApplyTime) error {
	payload := make(map[string]any)

	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Bios)
	err := original.UnmarshalJSON(c.RawData)
	if err != nil {
		return err
	}

	for key := range attrs {
		if strings.HasPrefix(key, "BootTypeOrder") ||
			original.Attributes[key] != attrs[key] {
			payload[key] = attrs[key]
		}
	}

	resp, err := c.GetClient().Get(c.settingsTarget)
	defer DeferredCleanupHTTPResponse(resp)
	if err != nil {
		return err
	}

	// If there are any allowed updates, try to send updates to the system and
	// return the result.
	if len(payload) > 0 {
		data := map[string]any{"Boot": payload}
		if applyTime != "" {
			data["@Redfish.SettingsApplyTime"] = map[string]string{"ApplyTime": string(applyTime)}
		}

		var header = make(map[string]string)
		if resp.Header["Etag"] != nil {
			header["If-Match"] = resp.Header["Etag"][0]
		}

		resp, err = c.GetClient().PatchWithHeaders(c.settingsTarget, data, header)
		defer DeferredCleanupHTTPResponse(resp)
		if err != nil {
			return err
		}
	}

	return nil
}

// UpdateBootAttributes is used to update attribute values.
func (c *ComputerSystem) UpdateBootAttributes(attrs SettingsAttributes) error {
	return c.UpdateBootAttributesApplyAt(attrs, "")
}

// BootOptions gets all BootOption items for this system.
func (c *ComputerSystem) BootOptions() ([]*BootOption, error) {
	return GetCollectionObjects[BootOption](
		c.GetClient(),
		c.Boot.bootOptions)
}

// UnmarshalJSON unmarshals a HostedServices object from the raw JSON.
func (h *HostedServices) UnmarshalJSON(b []byte) error {
	type temp HostedServices
	var tmp struct {
		temp
		StorageServices Link `json:"StorageServices"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*h = HostedServices(tmp.temp)

	// Extract the links to other entities for later
	h.storageServices = tmp.StorageServices.String()

	return nil
}

// StorageServices gets the StorageServices linked resource.
func (h *HostedServices) StorageServices(client Client) (*StorageService, error) {
	if h.storageServices == "" {
		return nil, nil
	}
	return GetObject[StorageService](client, h.storageServices)
}

// UnmarshalJSON unmarshals an IPMIHostInterface object from the raw JSON.
func (i *IPMIHostInterface) UnmarshalJSON(b []byte) error {
	type temp IPMIHostInterface
	var t temp
	err := json.Unmarshal(b, &t)
	if err == nil {
		*i = IPMIHostInterface(t)
		return nil
	}

	// Some BMCs emit a bare bool rather than the object the schema defines.
	var enabled bool
	if json.Unmarshal(b, &enabled) != nil {
		return err
	}
	i.ServiceEnabled = enabled
	return nil
}
