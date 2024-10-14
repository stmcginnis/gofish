//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"github.com/stmcginnis/gofish/common"
)

// FixedBootOrder is the fixed boot order information associated with the system.
// The non-OEM ComputerSystem BootOrder property does not support PATCH method
// since X13/H13 platforms Configuring system boot device order should be via
// FixedBootOrder.
// TODO: This is currently read-only in Gofish.
type FixedBootOrder struct {
	common.Entity
	BootModeSelected           string
	FixedBootOrder             []string
	FixedBootOrderDisabledItem []string
	UEFINetwork                []string
	UEFINetworkDisabledItem    []string
	UEFIHardDisk               []string
	UEFIAP                     []string
	UEFIAPDisabledItem         []string
}

// GetFixedBootOrder will get a FixedBootOrder instance from the service.
func GetFixedBootOrder(c common.Client, uri string) (*FixedBootOrder, error) {
	return common.GetObject[FixedBootOrder](c, uri)
}
