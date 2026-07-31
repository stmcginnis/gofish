//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
	"github.com/coreweave/gofish/redfish"
)

// USBCDBootSourceOverrideTarget is a Supermicro-specific boot target that is not
// part of the Redfish specification.
const USBCDBootSourceOverrideTarget redfish.BootSourceOverrideTarget = "UsbCd"

// ComputerSystem is a Supermicro OEM instance of a ComputerSystem.
type ComputerSystem struct {
	redfish.ComputerSystem
	nodeManager    string
	fixedBootOrder string
}

// NodeManager gets the NodeManager for the system.
func (cs *ComputerSystem) NodeManager(queryOpts ...common.QueryGroupOption) (*NodeManager, error) {
	return cs.NodeManagerWithContext(common.ContextOf(cs.GetClient()), queryOpts...)
}

// NodeManagerWithContext gets the NodeManager for the system.
func (cs *ComputerSystem) NodeManagerWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*NodeManager, error) {
	return GetNodeManagerWithContext(ctx, cs.GetClient(), cs.nodeManager, queryOpts...)
}

// FixedBootOrder gets the FixedBootOrder instance for the system.
func (cs *ComputerSystem) FixedBootOrder(queryOpts ...common.QueryGroupOption) (*FixedBootOrder, error) {
	return cs.FixedBootOrderWithContext(common.ContextOf(cs.GetClient()), queryOpts...)
}

// FixedBootOrderWithContext gets the FixedBootOrder instance for the system.
func (cs *ComputerSystem) FixedBootOrderWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*FixedBootOrder, error) {
	return GetFixedBootOrderWithContext(ctx, cs.GetClient(), cs.fixedBootOrder, queryOpts...)
}

// FromComputerSystem converts a standard ComputerSystem object to the OEM implementation.
func FromComputerSystem(system *redfish.ComputerSystem) (*ComputerSystem, error) {
	type Oem struct {
		Supermicro struct {
			NodeManager    common.Link `json:"NodeManager"`
			FixedBootOrder common.Link `json:"FixedBootOrder"`
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
