//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"errors"
)

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

// postUnmarshal records the allowed reset types and the Drives and PCIeDevices
// links, which services report either as a collection or as a Links array.
func (c *Chassis) postUnmarshal(b []byte) error {
	var t struct {
		Actions struct {
			Reset struct {
				ActionTarget
				AllowedResetTypes []ResetType `json:"ResetType@Redfish.AllowableValues"`
			} `json:"#Chassis.Reset"`
		}
		Links struct {
			Drives      Links
			DrivesCount int `json:"Drives@odata.count"`
			PCIeDevices Links
		}
		Drives Link
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	c.resetActionInfoTarget = t.Actions.Reset.ActionInfoTarget
	c.SupportedResetTypes = t.Actions.Reset.AllowedResetTypes
	c.drives = t.Drives.String()
	c.linkedDrives = t.Links.Drives.ToStrings()
	c.linkedPCIeDevices = t.Links.PCIeDevices.ToStrings()
	if c.DrivesCount == 0 && t.Links.DrivesCount > 0 {
		c.DrivesCount = t.Links.DrivesCount
	}

	return nil
}
