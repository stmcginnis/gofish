//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

type ComputerSystem struct {
	redfish.ComputerSystem

	OEMSystem                         OEMSystem
	targetSoftwareInstallationService string
}

type OEMSystem struct {
	common.Entity
	common.Resource

	// BIOSReleaseDate represents the BIOS release date.
	BIOSReleaseDate string
	// This property represents the modular chassis slot numbers that the server blade occupies in the modular enclosure.
	BaseBoardChassisSlot string
	// BatteryRollupStatus provides the battery rollup status of all the system components.
	BatteryRollupStatus string
	// The property represents the geometric dimension of the server blade enclosure in modular enclosure described.
	BladeGeometry string
	// This property represents the IP address for the enclosures Chassis Management Controller (CMC).
	CMCIP string
	// The property contains the rollup status of all the CPUs.
	CPURollupStatus string
	// This property represents the chassis model for the modular enclosure chassis.
	ChassisModel string
	// This property represents name of the Chassis.
	ChassisName string
	// This property represents the Service Tag of the modular enclosure chassis.
	ChassisServiceTag string
	// The property represents the system height in units of rack space (U).
	ChassisSystemHeightUnit int
	// CurrentRollupStatus provides the current rollup status of all the system components.
	CurrentRollupStatus string
	// Calculated, not measured, exhaust temperature in Degree Celsius.
	EstimatedExhaustTemperatureCelsius int
	// EstimatedSystemAirflow provides the estimated airflow over the chassis in Cubic Feet per Minute (CFM).
	EstimatedSystemAirflowCFM int
	// ExpressServiceCode of the system.
	ExpressServiceCode string
	// FanRollupStatus provides the fan rollup status of all the system components.
	FanRollupStatus string
	// IDSDMRollupStatus provides the live status of IDSDM (Internal Dual SD Mode) sensors.
	IDSDMRollupStatus string
	// IntrusionRollupStatus provides the live status of chassis intrusion sensors.
	IntrusionRollupStatus string
	// This property is used to identify if the system is OEM branded.
	IsOEMBranded string
	// This property provides the last time System Inventory Collection On Reboot(CSIOR) was performed or the object was last updated on iDRAC.
	LastSystemInventoryTime string
	// This property provides the last time the data was updated.
	LastUpdateTime string
	// provides the licensing rollup // status of all the system components.
	LicensingRollupStatus string
	// This property provides the systems physical size such as "1 U"
	ManagedSystemSize string
	// Maximum CPU sockets in the system.
	MaxCPUSockets int
	// The number of slots or sockets available for memory devices in the array.
	MaxDIMMSlots int
	// Maximum PCIe slots in the system.
	MaxPCIeSlots int
	// MemoryOperationMode denotes the mode of operation for system memory, such as mirrored, advanced ECC, or optimized mode.
	MemoryOperationMode string
	// NodeID is a unique property of the blade, based on the Service Tag.
	NodeID string
	// The property contains the power-supply rollup status of all the system components.
	PSRollupStatus string
	// Platform GUID uniquely identifies the platform
	PlatformGUID string
	// This property indicates the memory sockets in the system that are populated.
	PopulatedDIMMSlots int
	// Populated PCIe slots in the system.
	PopulatedPCIeSlots int
	// This property indicates the current state of the powercap setting of the associated managed system element.
	PowerCapEnabledState string
	// SDCardRollupStatus provides the SD-card rollup status.
	SDCardRollupStatus string
	// SELRollupStatus provides the SEL rollup status.
	SELRollupStatus string
	// This property represents the power, in Watt, that is allocated by the chassis manager to the blade systems.
	ServerAllocationWatts int
	// System GUID uniquely identifies the system.
	SmbiosGUID string `json:"smbiosGUID"`
	// StorageRollupStatus provides the storage rollup status of all the storage components.
	StorageRollupStatus string
	// The primary hardware error correction or detection method supported by the memory array.
	SysMemErrorMethodology string
	// Represents the failover state of the system memory.
	SysMemFailOverState string
	// The physical location of the memory array; whether on the system board or on an add-in board.
	SysMemLocation string
	// SystemMemoryPrimaryStatus provides a high-level status value that is intended to align with Red-Yellow-Green type representation of status for the system memory.
	SysMemPrimaryStatus string
	// SystemGeneration represents the generation of the Dell EMC system.
	SystemGeneration string
	// System ID describes the model of the system in integer value.
	SystemID int
	// System revision indicates the revision of the system from a hardware perspective.
	SystemRevision string
	// The property contains the temperature rollup status of all the system components.
	TempRollupStatus string
	// TempStatisticsRollupStatus provides the temperature statistics rollup status of all the system components.
	TempStatisticsRollupStatus string
	// UUID uniquely identifies the system.
	UUID string
	// The property contains the voltage rollup status of all the system components.
	VoltRollupStatus string
}

// Unmarshals a redfish.ComputerSystem into a dell.ComputerSystem
func FromComputerSystem(computerSystem *redfish.ComputerSystem) (*ComputerSystem, error) {
	cs := ComputerSystem{
		ComputerSystem: *computerSystem,
	}

	var t struct {
		Oem struct {
			Dell struct {
				DellSystem OEMSystem `json:",omitempty"`
			} `json:",omitempty"`
		} `json:",omitempty"`
		Links struct {
			Oem struct {
				Dell struct {
					DellSoftwareInstallationService common.Link `json:",omitempty"`
				} `json:",omitempty"`
			} `json:",omitempty"`
		} `json:",omitempty"`
	}

	err := json.Unmarshal(computerSystem.RawData, &t)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal Dell OEM data: %w", err)
	}

	cs.targetSoftwareInstallationService = t.Links.Oem.Dell.DellSoftwareInstallationService.String()
	cs.OEMSystem = t.Oem.Dell.DellSystem

	return &cs, nil
}

func (cs *ComputerSystem) SoftwareInstallationService() (*SoftwareInstallationService, error) {
	if cs.targetSoftwareInstallationService == "" {
		return nil, errors.New("software installation service is not supported by this system")
	}
	return GetSoftwareInstallationService(cs.GetClient(), cs.targetSoftwareInstallationService)
}
