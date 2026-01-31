//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"github.com/stmcginnis/gofish/schemas"
)

// FixedBootOrder is the fixed boot order information associated with the system.
// The non-OEM ComputerSystem BootOrder property does not support PATCH method
// since X13/H13 platforms Configuring system boot device order should be via
// FixedBootOrder.
// TODO: This is currently read-only in Gofish.
type FixedBootOrder struct {
	schemas.Entity
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
func GetFixedBootOrder(c schemas.Client, uri string) (*FixedBootOrder, error) {
	return schemas.GetObject[FixedBootOrder](c, uri)
}
