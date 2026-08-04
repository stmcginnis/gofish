//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"errors"

	"github.com/coreweave/gofish/common"
	"github.com/coreweave/gofish/redfish"
)

type ManagerConfigResetOption string

const (
	PreserveUserManagerConfigResetOption ManagerConfigResetOption = "PreserveUser"
	ClearConfigManagerConfigResetOption  ManagerConfigResetOption = "ClearConfig"
	ResetToAdminManagerConfigResetOption ManagerConfigResetOption = "ResetToADMIN"
)

// Manager is a Supermicro OEM instance of a Manager.
type Manager struct {
	redfish.Manager

	radius           string
	mouseMode        string
	ntp              string
	ipAccessControl  string
	smcRAKP          string
	syslog           string
	sysLockdown      string
	memoryPFA        string
	memoryHealthComp string
	snooping         string
	fanMode          string
	iKVM             string
	kcsInterface     string
	lldp             string
	licenseManager   string

	managerConfigResetTarget string
}

// FromManager converts a standard Manager object to the OEM implementation.
func FromManager(manager *redfish.Manager) (*Manager, error) {
	m := Manager{
		Manager: *manager,
	}

	var t struct {
		Oem struct {
			Supermicro struct {
				RADIUS           common.Link `json:"RADIUS"`
				MouseMode        common.Link `json:"MouseMode"`
				NTP              common.Link `json:"NTP"`
				IPAccessControl  common.Link `json:"IPAccessControl"`
				SMCRAKP          common.Link `json:"SMCRAKP"`
				Syslog           common.Link `json:"Syslog"`
				SysLockdown      common.Link `json:"SysLockdown"`
				MemoryPFA        common.Link `json:"MemoryPFA"`
				MemoryHealthComp common.Link `json:"MemoryHealthComp"`
				Snooping         common.Link `json:"Snooping"`
				FanMode          common.Link `json:"FanMode"`
				IKVM             common.Link `json:"IKVM"`
				KCSInterface     common.Link `json:"KCSInterface"`
				LLDP             common.Link `json:"LLDP"`
				LicenseManager   common.Link `json:"LicenseManager"`
			} `json:"Supermicro"`
		} `json:"Oem"`
		Actions struct {
			Oem struct {
				ManagerConfigReset common.ActionTarget `json:"#SmcManagerConfig.Reset"`
			} `json:"Oem"`
		} `json:"Actions"`
	}

	err := json.Unmarshal(manager.RawData, &t)
	if err != nil {
		return nil, err
	}

	m.radius = t.Oem.Supermicro.RADIUS.String()
	m.mouseMode = t.Oem.Supermicro.MouseMode.String()
	m.ntp = t.Oem.Supermicro.NTP.String()
	m.ipAccessControl = t.Oem.Supermicro.IPAccessControl.String()
	m.smcRAKP = t.Oem.Supermicro.SMCRAKP.String()
	m.syslog = t.Oem.Supermicro.Syslog.String()
	m.sysLockdown = t.Oem.Supermicro.SysLockdown.String()
	m.memoryPFA = t.Oem.Supermicro.MemoryPFA.String()
	m.memoryHealthComp = t.Oem.Supermicro.MemoryHealthComp.String()
	m.snooping = t.Oem.Supermicro.Snooping.String()
	m.fanMode = t.Oem.Supermicro.FanMode.String()
	m.iKVM = t.Oem.Supermicro.IKVM.String()
	m.kcsInterface = t.Oem.Supermicro.KCSInterface.String()
	m.lldp = t.Oem.Supermicro.LLDP.String()
	m.licenseManager = t.Oem.Supermicro.LicenseManager.String()

	m.managerConfigResetTarget = t.Actions.Oem.ManagerConfigReset.Target

	m.SetClient(manager.GetClient())
	return &m, nil
}

// RADIUS gets the RADIUS instance associated with this manager.
func (m *Manager) RADIUS(queryOpts ...common.QueryGroupOption) (*RADIUS, error) {
	return GetRADIUS(m.GetClient(), m.radius, queryOpts...)
}

// MouseMode gets the MouseMode instance associated with this manager.
func (m *Manager) MouseMode(queryOpts ...common.QueryGroupOption) (*MouseMode, error) {
	return GetMouseMode(m.GetClient(), m.mouseMode, queryOpts...)
}

// NTP gets the NTP instance associated with this manager.
func (m *Manager) NTP(queryOpts ...common.QueryGroupOption) (*NTP, error) {
	return GetNTP(m.GetClient(), m.ntp, queryOpts...)
}

// SMCRAKP gets the SMCRAKP instance associated with this manager.
func (m *Manager) SMCRAKP(queryOpts ...common.QueryGroupOption) (*SMCRAKP, error) {
	return GetSMCRAKP(m.GetClient(), m.smcRAKP, queryOpts...)
}

// Syslog gets the Syslog instance associated with this manager.
func (m *Manager) Syslog(queryOpts ...common.QueryGroupOption) (*Syslog, error) {
	return GetSyslog(m.GetClient(), m.syslog, queryOpts...)
}

// SysLockdown gets the SysLockdown instance associated with this manager.
func (m *Manager) SysLockdown(queryOpts ...common.QueryGroupOption) (*SysLockdown, error) {
	return GetSysLockdown(m.GetClient(), m.sysLockdown, queryOpts...)
}

// MemoryPFA gets the MemoryPFA instance associated with this manager.
func (m *Manager) MemoryPFA(queryOpts ...common.QueryGroupOption) (*MemoryPFA, error) {
	return GetMemoryPFA(m.GetClient(), m.memoryPFA, queryOpts...)
}

// MemoryHealthComp gets the MemoryHealthComp instance associated with this manager.
func (m *Manager) MemoryHealthComp(queryOpts ...common.QueryGroupOption) (*MemoryHealthComp, error) {
	return GetMemoryHealthComp(m.GetClient(), m.memoryHealthComp, queryOpts...)
}

// Snooping gets the Snooping instance associated with this manager.
func (m *Manager) Snooping(queryOpts ...common.QueryGroupOption) (*Snooping, error) {
	return GetSnooping(m.GetClient(), m.snooping, queryOpts...)
}

// FanMode gets the FanMode instance associated with this manager.
func (m *Manager) FanMode(queryOpts ...common.QueryGroupOption) (*FanMode, error) {
	return GetFanMode(m.GetClient(), m.fanMode, queryOpts...)
}

// IKVM gets the IKVM instance associated with this manager.
func (m *Manager) IKVM(queryOpts ...common.QueryGroupOption) (*IKVM, error) {
	return GetIKVM(m.GetClient(), m.iKVM, queryOpts...)
}

// KCSInterface gets the KCSInterface instance associated with this manager.
func (m *Manager) KCSInterface(queryOpts ...common.QueryGroupOption) (*KCSInterface, error) {
	return GetKCSInterface(m.GetClient(), m.kcsInterface, queryOpts...)
}

// LLDP gets the LLDP instance associated with this manager.
func (m *Manager) LLDP(queryOpts ...common.QueryGroupOption) (*LLDP, error) {
	return GetLLDP(m.GetClient(), m.lldp, queryOpts...)
}

// LicenseManager gets the LicenseManager instance associated with this manager.
func (m *Manager) LicenseManager(queryOpts ...common.QueryGroupOption) (*LicenseManager, error) {
	return GetLicenseManager(m.GetClient(), m.licenseManager, queryOpts...)
}

// ManagerConfigReset resets the BMC to factory defaults.
func (m *Manager) ManagerConfigReset(option ManagerConfigResetOption) error {
	if m.managerConfigResetTarget == "" {
		return errors.New("manager config reset not supported by this system")
	}

	return m.Post(m.managerConfigResetTarget, map[string]interface{}{"Option": option})
}
