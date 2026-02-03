//
// SPDX-License-Identifier: BSD-3-Clause
//

package ami

import (
	"encoding/json"
	"errors"

	"github.com/stmcginnis/gofish/schemas"
)

type PreserveConfiguration string

const (
	// AUTOMATION_ENGINE To preserve AUTOMATION_ENGINE.
	AutomationEnginePreserveConfiguration PreserveConfiguration = "AUTOMATION_ENGINE"
	// Authentication To preserve Authentication.
	AuthenticationPreserveConfiguration PreserveConfiguration = "Authentication"
	// CMX To preserve CMX.
	CMXPreserveConfiguration PreserveConfiguration = "CMX"
	// EXTLOG To preserve EXTLOG.
	EXTLOGPreserveConfiguration PreserveConfiguration = "EXTLOG"
	// FRU To preserve FRU.
	FRUPreserveConfiguration PreserveConfiguration = "FRU"
	// IPMI To preserve config of IPMI. It will preserve Network automatically if preserve IPMI
	IPMIPreserveConfiguration PreserveConfiguration = "IPMI"
	// KVM To preserve KVM.
	KVMPreserveConfiguration PreserveConfiguration = "KVM"
	// NTP To preserve NTP.
	NTPPreserveConfiguration PreserveConfiguration = "NTP"
	// Network To preserve config of Network. It will preserve IPMI automatically if preserve Network
	NetworkPreserveConfiguration PreserveConfiguration = "Network"
	// REDFISH To preserve REDFISH.
	REDFISHPreserveConfiguration PreserveConfiguration = "REDFISH"
	// SDR To preserve SDR.
	SDRPreserveConfiguration PreserveConfiguration = "SDR"
	// SEL To preserve SEL.
	SELPreserveConfiguration PreserveConfiguration = "SEL"
	// SNMP To preserve SNMP.
	SNMPPreserveConfiguration PreserveConfiguration = "SNMP"
	// SSH To preserve SSH.
	SSHPreserveConfiguration PreserveConfiguration = "SSH"
	// Syslog To preserve Syslog.
	SyslogPreserveConfiguration PreserveConfiguration = "Syslog"
	// WEB To preserve WEB.
	WEBPreserveConfiguration PreserveConfiguration = "WEB"
)

// AMIUpdateService shall be used to represent an Update Service for a Redfish implementation. It represents the
// properties that affect the service itself.
type AMIUpdateService struct {
	// FlashPercentage shall represent the FlashPercentage of the UpdateService. The format of the string shall be the
	// Percentage completed for Flashing of the UpdateService.
	FlashPercentage string
	// PreserveConfiguration is whether the configuration needs to be preserved
	// when doing firmware update or restore factory defaults.
	PreserveConfiguration bool
	// UpdateInformation is the information about the updated firmware.
	UpdateInformation string
	// UpdateStatus shall represent the UpdateStatus of the UpdateService. The format of the string shall be the Status
	// is Preparing or Downloading or Verifying or Flashing.
	UpdateStatus string
	// UpdateTarget shall represent UpdateTarget of the UpdateService. The format of this string shall be BMC.
	UpdateTarget string
}

// BMC is the schema definition for image configurations and preserve configurations.
type BMC struct {
	// DualImageConfigurations is information about dual image handling.
	DualImageConfigurations DualImageConfigurations
}

// DualImageConfigurations contains information about dual image handling.
type DualImageConfigurations struct {
	// ActiveImage is the active image in the BMC.
	ActiveImage string
	// BootImage represents the image to which BMC boots to.
	BootImage string
	// FirmwareImage1Name is the name of image #1.
	FirmwareImage1Name string
	// FirmwareImage1Version is the version of the first firmware image.
	FirmwareImage1Version string
	// FirmwareImage2Name is the name of image #2.
	FirmwareImage2Name string
	// FirmwareImage2Version is the version of the second firmware image.
	FirmwareImage2Version string
}

// UpdateInformation shall contain the available actions for this resource.
type UpdateInformation struct {
	// UpdateComponent The information about the updated firmware.
	UpdateComponent string
}

type BIOS struct {
	BIOSPreserveNVRAM bool
}

// UpdateService is the update service instance associated with the system.
type UpdateService struct {
	schemas.UpdateService

	AMIUpdateService AMIUpdateService
	BIOS             BIOS
	BMC              BMC

	uploadCABundleTarget string
}

// FromUpdateService gets the OEM instance of the UpdateService.
func FromUpdateService(updateService *schemas.UpdateService) (*UpdateService, error) {
	us := UpdateService{
		UpdateService: *updateService,
	}

	var t struct {
		Actions struct {
			Oem struct {
				UploadCABundle schemas.ActionTarget `json:"#UpdateService.UploadCABundle"`
			}
		}
		Oem struct {
			AMIUpdateService AMIUpdateService `json:"AMIUpdateService"`
			BMC              BMC              `json:"BMC"`
			BIOS             BIOS             `json:"BIOS"`
		}
	}

	err := json.Unmarshal(updateService.RawData, &t)
	if err != nil {
		return nil, err
	}

	us.AMIUpdateService = t.Oem.AMIUpdateService
	us.BMC = t.Oem.BMC
	us.BIOS = t.Oem.BIOS

	us.uploadCABundleTarget = t.Actions.Oem.UploadCABundle.Target

	return &us, nil
}

// GetUpdateService will get a UpdateService instance from the service.
func GetUpdateService(c schemas.Client, uri string) (*UpdateService, error) {
	return schemas.GetObject[UpdateService](c, uri)
}

// UploadCABundle uploads CA certificates.
// WARNING: The AMI Redfish service JsonSchema does not define any parameters for
// this action. This is most likely incorrect and will cause this to fail.
func (us *UpdateService) UploadCABundle() error {
	if us.uploadCABundleTarget == "" {
		return errors.New("upload ca bundle is not supported by this system")
	}

	return us.Post(us.uploadCABundleTarget, nil)
}
