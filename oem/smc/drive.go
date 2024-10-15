package smc

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/redfish"
)

// Drive extends a redfish.Drive for additional OEM fields
type Drive struct {
	redfish.Drive
	smc struct {
		Oem     driveOem     `json:"Oem"`
		Actions driveActions `json:"Actions"`
	}
}

type driveOem struct {
	Supermicro struct {
		Temperature             int
		PercentageDriveLifeUsed int
		DriveFunctional         bool
	} `json:"Supermicro"`
}

type driveTarget struct {
	Target     string `json:"target"`
	ActionInfo string `json:"@Redfish.ActionInfo"`
}

type driveActions struct {
	redfish.DriveActions
	Oem struct {
		DriveIndicate    driveTarget `json:"#Drive.Indicate"`
		SmcDriveIndicate driveTarget `json:"#SmcDrive.Indicate"`
	} `json:"Oem"`
}

// FromDrive returns an OEM-extended redfish drive
func FromDrive(drive *redfish.Drive) (Drive, error) {
	smcDrive := Drive{
		Drive: *drive,
	}
	smcDrive.smc.Actions.DriveActions = drive.Actions

	if err := json.Unmarshal(drive.Oem, &smcDrive.smc.Oem); err != nil {
		return smcDrive, err
	}

	if err := json.Unmarshal(drive.Actions.Oem, &smcDrive.smc.Actions.Oem); err != nil {
		return smcDrive, err
	}

	return smcDrive, nil
}

// Temperature returns the OEM provided temperature for the drive
func (d Drive) Temperature() int {
	return d.smc.Oem.Supermicro.Temperature
}

// PercentageDriveLifeUsed returns the OEM provided drive life estimate as a percentage used
func (d Drive) PercentageDriveLifeUsed() int {
	return d.smc.Oem.Supermicro.PercentageDriveLifeUsed
}

// Functional returns the OEM provided flag that suggests whether a drive is functional or not
func (d Drive) Functional() bool {
	return d.smc.Oem.Supermicro.DriveFunctional
}

// indicateTarget figures out what uri to follow for indicator light actions.
// This is a separate function for testing.
func (d Drive) indicateTarget() string {
	// We check both the SmcDriveIndicate and the DriveIndicate targets
	// in the Oem sections - certain models and bmc firmwares will mix
	// these up, so we check both
	if len(d.smc.Actions.Oem.SmcDriveIndicate.Target) > 0 {
		return d.smc.Actions.Oem.SmcDriveIndicate.Target
	}

	return d.smc.Actions.Oem.DriveIndicate.Target
}

// Indicate will set the indicator light activity, true for on, false for off
func (d Drive) Indicate(active bool) error {
	return d.Post(d.indicateTarget(), map[string]interface{}{"Active": active})
}
