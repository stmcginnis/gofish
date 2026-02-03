//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/schemas"
)

// USBCDBootSourceOverrideTarget is a Supermicro-specific boot target that is not
// part of the Redfish specification.
const USBCDBootSourceOverrideTarget schemas.BootSource = "UsbCd"

// ComputerSystem is a Supermicro OEM instance of a ComputerSystem.
type ComputerSystem struct {
	schemas.ComputerSystem
	nodeManager    string
	fixedBootOrder string
}

// NodeManager gets the NodeManager for the system.
func (cs *ComputerSystem) NodeManager() (*NodeManager, error) {
	return GetNodeManager(cs.GetClient(), cs.nodeManager)
}

// FixedBootOrder gets the FixedBootOrder instance for the system.
func (cs *ComputerSystem) FixedBootOrder() (*FixedBootOrder, error) {
	return GetFixedBootOrder(cs.GetClient(), cs.fixedBootOrder)
}

// FromComputerSystem converts a standard ComputerSystem object to the OEM implementation.
func FromComputerSystem(system *schemas.ComputerSystem) (*ComputerSystem, error) {
	type Oem struct {
		Supermicro struct {
			NodeManager    schemas.Link `json:"NodeManager"`
			FixedBootOrder schemas.Link `json:"FixedBootOrder"`
		} `json:"Supermicro"`
	}

	cs := &ComputerSystem{}
	err := json.Unmarshal(system.RawData, cs)
	if err != nil {
		return nil, err
	}

	oem := &Oem{}
	err = json.Unmarshal(cs.OEM, oem)
	if err != nil {
		return nil, err
	}
	cs.nodeManager = oem.Supermicro.NodeManager.String()
	cs.fixedBootOrder = oem.Supermicro.FixedBootOrder.String()

	cs.SetClient(system.GetClient())
	return cs, nil
}
