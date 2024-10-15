package smc

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/redfish"
)

type Drive struct {
	redfish.Drive
	Oem DriveOem `json:"Oem"`
}

type DriveOem struct {
	Supermicro struct {
		Temperature             int
		PercentageDriveLifeUsed int
		DriveFunctional         bool
	} `json:"Supermicro"`
}

type DriveTarget struct {
	Target     string `json:"target"`
	ActionInfo string `json:"@Redfish.ActionInfo"`
}

type DriveActions struct {
	redfish.DriveActions
	Oem struct {
		DriveIndicate    DriveTarget `json:"#Drive.Indicate"`
		SmcDriveIndicate DriveTarget `json:"#SmcDrive.Indicate"`
	} `json:"Oem"`
}

func FromDrive(drive *redfish.Drive) (Drive, error) {
	var oem DriveOem
	err := json.Unmarshal(drive.Oem, &oem)

	return Drive{
		Drive: *drive,
		Oem:   oem,
	}, err
}

func FromDriveActions(da *redfish.DriveActions) (DriveActions, error) {
	oemActions := DriveActions{
		DriveActions: *da,
	}

	err := json.Unmarshal(da.Oem, &oemActions.Oem)
	return oemActions, err
}

// DriveIndicateTarget checks both the SmcDriveIndicate and DriveIndicateTarget
// Oem entries and returns the first populated target, due to key inconsistencies
func (da DriveActions) DriveIndicateTarget() string {
	if len(da.Oem.SmcDriveIndicate.Target) > 0 {
		return da.Oem.SmcDriveIndicate.Target
	}

	return da.Oem.DriveIndicate.Target
}
