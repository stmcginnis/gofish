//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/redfish"
)

const computerSystemBody = `{
    "Id": "System.Embedded.1",
    "Links": {
        "Oem": {
            "Dell": {
                "DellVideoNetworkCollection@Redfish.Deprecated": "The DellVideoNetwork resource has been deprecated.",
                "DellSoftwareInstallationService": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSoftwareInstallationService"
                },
                "DellBIOSService": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellBIOSService"
                },
                "DellProcessorCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellProcessors"
                },
                "DellSwitchConnectionService": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSwitchConnectionService"
                },
                "DellPhysicalDiskCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellDrives"
                },
                "DellSwitchConnectionCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSwitchConnections"
                },
                "DellRaidService": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellRaidService"
                },
                "DellChassisCollection": {
                    "@odata.id": "/redfish/v1/Chassis/System.Embedded.1/Oem/Dell/DellChassis"
                },
                "DellVideoCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellVideo"
                },
                "DellPCIeSSDCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellPCIeSSDs"
                },
                "DellVideoNetworkCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellVideoNetwork"
                },
                "DellMetricService": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellMetricService"
                },
                "DellAcceleratorCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellAccelerators"
                },
                "DellMemoryCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellMemory"
                },
                "DellSystemManagementService": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSystemManagementService"
                },
                "DellPresenceAndStatusSensorCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellPresenceAndStatusSensors"
                },
                "DellRollupStatusCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellRollupStatus"
                },
                "@odata.type": "#OemComputerSystem.v1_0_0.Links",
                "DellNumericSensorCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellNumericSensors"
                },
                "DellGPUSensorCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellGPUSensors"
                },
                "DellControllerCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellControllers"
                },
                "DellBootSources": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellBootSources"
                },
                "DellSlotCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSlots"
                },
                "DellOSDeploymentService": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellOSDeploymentService"
                },
                "DellPCIeSSDExtenderCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellPCIeSSDExtenders"
                },
                "DellSensorCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSensors"
                },
                "DellVirtualDiskCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellVolumes"
                },
                "@odata.context": "/redfish/v1/$metadata#OemComputerSystem.OemComputerSystem",
                "DellPSNumericSensorCollection": {
                    "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellPSNumericSensors"
                }
            }
        }
    },
    "Oem": {
        "Dell": {
            "@odata.type": "#OemComputerSystem.v1_0_0.ComputerSystem",
            "DellSystem": {
                "BaseBoardChassisSlot": "NA",
                "UUID": "00000000-0000-0000-0000-000000000000",
                "CoolingRollupStatus": "OK",
                "BIOSReleaseDate": "05/20/2025",
                "ManagedSystemSize": "1 U",
                "@odata.etag": "W/\"gen-21832\"",
                "@odata.type": "#DellSystem.v1_4_0.DellSystem",
                "SysMemFailOverState": "NotInUse",
                "EstimatedExhaustTemperatureCelsius": 37,
                "SysMemPrimaryStatus": "OK",
                "ChassisModel": null,
                "SysMemErrorMethodology": "Multi-bitECC",
                "BatteryRollupStatus": "OK",
                "BladeGeometry": "NotApplicable",
                "MaxSystemMemoryMiB": 6291456,
                "SysMemLocation": "SystemBoardOrMotherboard",
                "PopulatedDIMMSlots": 12,
                "@odata.context": "/redfish/v1/$metadata#DellSystem.DellSystem",
                "CPURollupStatus": "OK",
                "MaxPCIeSlots": 5,
                "IsOEMBranded": "False",
                "EstimatedSystemAirflowCFM": 40,
                "PopulatedPCIeSlots": 3,
                "PSRollupStatus": "OK",
                "TempStatisticsRollupStatus": "OK",
                "NodeID": "0000000",
                "SystemHealthRollupStatus": "OK",
                "ChassisName": "Main System Chassis",
                "PowerCapEnabledState": "Disabled",
                "ServerAllocationWatts": null,
                "MaxDIMMSlots": 24,
                "IntrusionRollupStatus": "OK",
                "SDCardRollupStatus": null,
                "Id": "System.Embedded.1",
                "ExpressServiceCode": "00000000000",
                "@odata.id": "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSystem/System.Embedded.1",
                "TempRollupStatus": "OK",
                "MemoryOperationMode": "Unknown",
                "FanRollupStatus": "OK",
                "LastSystemInventoryTime": "2025-08-05T02:55:58+00:00",
                "SELRollupStatus": "OK",
                "SystemRevision": "I",
                "CurrentRollupStatus": "OK",
                "LicensingRollupStatus": "OK",
                "smbiosGUID": "00000000-0000-0000-0000-000000000000",
                "ChassisSystemHeightUnit": 1,
                "VoltRollupStatus": "OK",
                "Description": "An instance of DellSystem will have data representing the overall system devices in the managed system.",
                "MaxCPUSockets": 1,
                "SystemGeneration": "17G Monolithic",
                "LastUpdateTime": "2025-07-19T01:43:33+00:00",
                "SystemID": 3244,
                "Name": "DellSystem",
                "PlatformGUID": "00000000-0000-0000-0000-000000000000",
                "ChassisServiceTag": "0000000",
                "StorageRollupStatus": "OK"
            },
            "@odata.context": "/redfish/v1/$metadata#OemComputerSystem.OemComputerSystem"
        }
    },
    "Actions": {
        "#ComputerSystem.Decommission": {
            "DecommissionTypes@Redfish.AllowableValues": [
                "Logs",
                "ManagerConfig",
                "All"
            ],
            "OEMDecommissionTypes@Redfish.AllowableValues": [
                "DellFwStoreClean",
                "DellFPSPIClean",
                "DellUserCertClean"
            ],
            "target": "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Decommission"
        },
        "#ComputerSystem.Reset": {
            "target": "/redfish/v1/Systems/System.Embedded.1/Actions/ComputerSystem.Reset",
            "ResetType@Redfish.AllowableValues": [
                "On",
                "ForceOff",
                "GracefulRestart",
                "GracefulShutdown",
                "ForceRestart",
                "Nmi",
                "PowerCycle",
                "PushPowerButton"
            ]
        }
    },
    "Name": "System"
}`

func TestComputerSystem(t *testing.T) {
	var s redfish.ComputerSystem
	err := json.NewDecoder(strings.NewReader(computerSystemBody)).Decode(&s)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	result, err := FromComputerSystem(&s)
	if err != nil {
		t.Errorf("Error converting Redfish Manager to SMC Manager: %s", err)
	}

	if result.ID != "System.Embedded.1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "System" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.targetSoftwareInstallationService != "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSoftwareInstallationService" {
		t.Errorf("Invalid ImportSystemConfig link: %s", result.targetSoftwareInstallationService)
	}

	if result.OEMSystem.SystemID != 3244 {
		t.Errorf("Invalid SystemID: %d", result.OEMSystem.SystemID)
	}
}
