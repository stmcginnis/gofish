//
// SPDX-License-Identifier: BSD-3-Clause
//

package ami

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// ManagerBootMode is is the boot mode of the manager.
type ManagerBootMode string

const (
	// NoneManagerBootMode Added None in Boot Option
	NoneManagerBootMode ManagerBootMode = "None"
	// SoftResetManagerBootMode Added SoftReset in Boot Option
	SoftResetManagerBootMode ManagerBootMode = "SoftReset"
	// ResetTimeoutManagerBootMode ResetTimeout support is Boot Option
	ResetTimeoutManagerBootMode ManagerBootMode = "ResetTimeout"
)

// AMIBIOSInventoryCRC provides the information related to inventory data/
type AMIBIOSInventoryCRC struct {
	// Bios provides the information related to inventory data.
	Bios Bios
	// ManagerBootConfiguration indicates the properties related to ManagerBoot
	ManagerBootConfiguration ManagerBootConfiguration
}

// BiosTableis the root for BiosTable information.
type BiosTable struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// FilesContent contains the contents of the BiosTable file.
	FilesContent string
}

// TableTag contains the TableTag informations.
type TableTag struct {
	// TableType shall contain a string representing the TableTag.
	TableType string
	// Value shall contains the value for the corresponding TableTag.
	Value string
}

// BiosTableTags is the root for TableTags information.
type BiosTableTags struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// NumberofTables contains the number of TableTags present.
	NumberofTables string
	// TableTags contains the TableTags informations.
	TableTags []TableTag
}

// Bios
type Bios struct {
	// BiosTable provides the information related to BiosTable
	BiosTable BiosTable
	// BiosTableTags provides the information related to BiosTableTags.
	BiosTableTags BiosTableTags
	// Inventory provides the information related to inventory data Crc value.
	Inventory Inventory
	// RedfishVersion shall represent the version of the Redfish service. The format of this string shall be of the
	// format majorversion.minorversion.errata in compliance with Protocol Version section of the Redfish
	// specification.
	RedfishVersion string
	// RTPVersion shall represent the version of the RTP Version.
	RTPVersion string
}

// Crc
type Crc struct {
	// GroupCrcList provides the information related to inventory data of GroupCrcList value.
	GroupCrcList []map[string]uint64
}

// Inventory
type Inventory struct {
	// Crc provides the information related to inventory data of Crc value.
	Crc Crc
}

// ManagerBootConfiguration
type ManagerBootConfiguration struct {
	// ManagerBootMode shall specify the enum supported by ManagerBootMode.
	ManagerBootMode ManagerBootMode
}

// ComputerSystem is the update service instance associated with the system.
type ComputerSystem struct {
	redfish.ComputerSystem

	BIOS                     Bios
	ManagerBootConfiguration ManagerBootConfiguration
	SSIFMode                 string
}

// FromComputerSystem gets the OEM instance of the ComputerSystemSystem.
func FromComputerSystem(computerSystem *redfish.ComputerSystem) (*ComputerSystem, error) {
	us := ComputerSystem{
		ComputerSystem: *computerSystem,
	}

	var t struct {
		Oem struct {
			Ami struct {
				BIOS                     Bios                     `json:"BIOS"`
				ManagerBootConfiguration ManagerBootConfiguration `json:"ManagerBootConfiguration"`
				SSIFMode                 string                   `json:"SSIFMode"`
			}
		}
	}

	err := json.Unmarshal(computerSystem.RawData, &t)
	if err != nil {
		return nil, err
	}

	us.BIOS = t.Oem.Ami.BIOS
	us.ManagerBootConfiguration = t.Oem.Ami.ManagerBootConfiguration
	us.SSIFMode = t.Oem.Ami.SSIFMode

	return &us, nil
}
