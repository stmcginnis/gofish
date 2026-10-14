//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"errors"
)

// UnmarshalJSON unmarshals a Chassis object from the raw JSON.
//
//nolint:funlen
func (c *Chassis) UnmarshalJSON(b []byte) error {
	type temp Chassis
	type cActions struct {
		Reset struct {
			ActionTarget
			AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
		} `json:"#Chassis.Reset"`
	}
	type cLinks struct {
		AutomationNodes       Links `json:"AutomationNodes"`
		Cables                Links `json:"Cables"`
		ComputerSystems       Links `json:"ComputerSystems"`
		ConnectedCoolingLoops Links `json:"ConnectedCoolingLoops"`
		ContainedBy           Link  `json:"ContainedBy"`
		Contains              Links `json:"Contains"`
		CooledBy              Links `json:"CooledBy"`
		CoolingUnits          Links `json:"CoolingUnits"`
		Drives                Links `json:"Drives"`
		DrivesCount           int   `json:"Drives@odata.count"`
		Facility              Link  `json:"Facility"`
		Fans                  Links `json:"Fans"`
		ManagedBy             Links `json:"ManagedBy"`
		ManagersInChassis     Links `json:"ManagersInChassis"`
		PCIeDevices           Links `json:"PCIeDevices"`
		PowerDistribution     Link  `json:"PowerDistribution"`
		PowerOutlets          Links `json:"PowerOutlets"`
		PowerSupplies         Links `json:"PowerSupplies"`
		PoweredBy             Links `json:"PoweredBy"`
		Processors            Links `json:"Processors"`
		ResourceBlocks        Links `json:"ResourceBlocks"`
		Storage               Links `json:"Storage"`
		Switches              Links `json:"Switches"`
	}
	var tmp struct {
		temp
		Actions            cActions
		Links              cLinks
		Assembly           ResolvedLink[Assembly] `json:"Assembly"`
		Certificates       Link                   `json:"Certificates"`
		Controls           Link                   `json:"Controls"`
		Drives             Link
		EnvironmentMetrics ResolvedLink[EnvironmentMetrics] `json:"EnvironmentMetrics"`
		FabricAdapters     Link                             `json:"FabricAdapters"`
		LeakDetectors      Link                             `json:"LeakDetectors"`
		LogServices        Link                             `json:"LogServices"`
		MediaControllers   Link                             `json:"MediaControllers"`
		Memory             Link                             `json:"Memory"`
		MemoryDomains      Link                             `json:"MemoryDomains"`
		NetworkAdapters    Link                             `json:"NetworkAdapters"`
		PCIeDevices        Link                             `json:"PCIeDevices"`
		PCIeSlots          ResolvedLink[PCIeSlots]          `json:"PCIeSlots"`
		Power              ResolvedLink[Power]              `json:"Power"`
		PowerSubsystem     ResolvedLink[PowerSubsystem]     `json:"PowerSubsystem"`
		Sensors            Link                             `json:"Sensors"`
		Thermal            ResolvedLink[Thermal]            `json:"Thermal"`
		ThermalSubsystem   ResolvedLink[ThermalSubsystem]   `json:"ThermalSubsystem"`
		TrustedComponents  Link                             `json:"TrustedComponents"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = Chassis(tmp.temp)

	// Extract the links to other entities for later
	c.resetTarget = tmp.Actions.Reset.Target
	c.resetActionInfoTarget = tmp.Actions.Reset.ActionInfoTarget
	c.SupportedResetTypes = tmp.Actions.Reset.AllowedResetTypes

	c.automationNodes = tmp.Links.AutomationNodes.ToStrings()
	c.cables = tmp.Links.Cables.ToStrings()
	c.computerSystems = tmp.Links.ComputerSystems.ToStrings()
	c.connectedCoolingLoops = tmp.Links.ConnectedCoolingLoops.ToStrings()
	c.containedBy = tmp.Links.ContainedBy.String()
	c.contains = tmp.Links.Contains.ToStrings()
	c.cooledBy = tmp.Links.CooledBy.ToStrings()
	c.coolingUnits = tmp.Links.CoolingUnits.ToStrings()
	c.drives = tmp.Drives.String()
	c.linkedDrives = tmp.Links.Drives.ToStrings()
	c.facility = tmp.Links.Facility.String()
	c.fans = tmp.Links.Fans.ToStrings()
	c.managedBy = tmp.Links.ManagedBy.ToStrings()
	c.managersInChassis = tmp.Links.ManagersInChassis.ToStrings()
	c.powerDistribution = tmp.Links.PowerDistribution.String()
	c.powerOutlets = tmp.Links.PowerOutlets.ToStrings()
	c.powerSupplies = tmp.Links.PowerSupplies.ToStrings()
	c.poweredBy = tmp.Links.PoweredBy.ToStrings()
	c.processors = tmp.Links.Processors.ToStrings()
	c.resourceBlocks = tmp.Links.ResourceBlocks.ToStrings()
	c.storage = tmp.Links.Storage.ToStrings()
	c.switches = tmp.Links.Switches.ToStrings()
	c.assembly = tmp.Assembly

	c.certificates = tmp.Certificates.String()
	c.controls = tmp.Controls.String()
	c.environmentMetrics = tmp.EnvironmentMetrics

	c.fabricAdapters = tmp.FabricAdapters.String()
	c.leakDetectors = tmp.LeakDetectors.String()
	c.logServices = tmp.LogServices.String()
	c.mediaControllers = tmp.MediaControllers.String()
	c.memory = tmp.Memory.String()
	c.memoryDomains = tmp.MemoryDomains.String()
	c.networkAdapters = tmp.NetworkAdapters.String()
	c.pCIeDevices = tmp.PCIeDevices.String()
	c.linkedPCIeDevices = tmp.Links.PCIeDevices.ToStrings()
	c.pCIeSlots = tmp.PCIeSlots

	c.power = tmp.Power

	c.powerSubsystem = tmp.PowerSubsystem

	c.sensors = tmp.Sensors.String()
	c.thermal = tmp.Thermal

	c.thermalSubsystem = tmp.ThermalSubsystem

	c.trustedComponents = tmp.TrustedComponents.String()
	if c.DrivesCount == 0 && tmp.Links.DrivesCount > 0 {
		c.DrivesCount = tmp.Links.DrivesCount
	}

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// GetSupportedResetTypes returns any reset types that the Chassis declares as supported
// via either ActionInfo or AllowableValues.
func (c *Chassis) GetSupportedResetTypes() ([]ResetType, error) {
	if len(c.SupportedResetTypes) > 0 {
		return c.SupportedResetTypes, nil
	}

	// if we don't have ResetTypes, try to get from ActionInfo
	if c.resetActionInfoTarget != "" {
		resetActionInfo, err := c.ResetActionInfo()
		if err != nil {
			return nil, err
		}

		vals, err := resetActionInfo.GetParamValues("ResetType", StringParameterTypes)
		if err != nil {
			return nil, err
		}

		for _, val := range vals {
			c.SupportedResetTypes = append(c.SupportedResetTypes, ResetType(val))
		}
	}

	return c.SupportedResetTypes, nil
}

// ResetActionInfo returns the ActionInfo for the Chassis reset action if supported
func (c *Chassis) ResetActionInfo() (*ActionInfo, error) {
	if c.resetActionInfoTarget == "" {
		return nil, errors.New("Chassis Reset resetActionInfoTarget not supported")
	}

	return GetObject[ActionInfo](c.GetClient(), c.resetActionInfoTarget)
}

// Drives gets the Drives linked resources.
func (c *Chassis) Drives() ([]*Drive, error) {
	// In version v1.2.0 of the spec, Drives were added to the Chassis.Links
	// property. But in v1.14.0 of the spec, Chassis.Drives was added as a
	// direct property.
	if c.drives != "" {
		return GetCollectionObjects[Drive](c.GetClient(), c.drives)
	}

	if len(c.linkedDrives) > 0 {
		return GetObjects[Drive](c.GetClient(), c.linkedDrives)
	}

	return []*Drive{}, nil
}

// PCIeDevices gets the PCIeDevices collection.
func (c *Chassis) PCIeDevices() ([]*PCIeDevice, error) {
	if c.pCIeDevices != "" {
		return GetCollectionObjects[PCIeDevice](c.GetClient(), c.pCIeDevices)
	}

	if len(c.linkedPCIeDevices) > 0 {
		return GetObjects[PCIeDevice](c.GetClient(), c.linkedPCIeDevices)
	}

	return []*PCIeDevice{}, nil
}
