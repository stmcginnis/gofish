//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"context"
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
	return m.RADIUSWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// RADIUSWithContext gets the RADIUS instance associated with this manager.
func (m *Manager) RADIUSWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*RADIUS, error) {
	return GetRADIUSWithContext(ctx, m.GetClient(), m.radius, queryOpts...)
}

// MouseMode gets the MouseMode instance associated with this manager.
func (m *Manager) MouseMode(queryOpts ...common.QueryGroupOption) (*MouseMode, error) {
	return m.MouseModeWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// MouseModeWithContext gets the MouseMode instance associated with this manager.
func (m *Manager) MouseModeWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*MouseMode, error) {
	return GetMouseModeWithContext(ctx, m.GetClient(), m.mouseMode, queryOpts...)
}

// NTP gets the NTP instance associated with this manager.
func (m *Manager) NTP(queryOpts ...common.QueryGroupOption) (*NTP, error) {
	return m.NTPWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// NTPWithContext gets the NTP instance associated with this manager.
func (m *Manager) NTPWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*NTP, error) {
	return GetNTPWithContext(ctx, m.GetClient(), m.ntp, queryOpts...)
}

// SMCRAKP gets the SMCRAKP instance associated with this manager.
func (m *Manager) SMCRAKP(queryOpts ...common.QueryGroupOption) (*SMCRAKP, error) {
	return m.SMCRAKPWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// SMCRAKPWithContext gets the SMCRAKP instance associated with this manager.
func (m *Manager) SMCRAKPWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*SMCRAKP, error) {
	return GetSMCRAKPWithContext(ctx, m.GetClient(), m.smcRAKP, queryOpts...)
}

// Syslog gets the Syslog instance associated with this manager.
func (m *Manager) Syslog(queryOpts ...common.QueryGroupOption) (*Syslog, error) {
	return m.SyslogWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// SyslogWithContext gets the Syslog instance associated with this manager.
func (m *Manager) SyslogWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*Syslog, error) {
	return GetSyslogWithContext(ctx, m.GetClient(), m.syslog, queryOpts...)
}

// SysLockdown gets the SysLockdown instance associated with this manager.
func (m *Manager) SysLockdown(queryOpts ...common.QueryGroupOption) (*SysLockdown, error) {
	return m.SysLockdownWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// SysLockdownWithContext gets the SysLockdown instance associated with this manager.
func (m *Manager) SysLockdownWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*SysLockdown, error) {
	return GetSysLockdownWithContext(ctx, m.GetClient(), m.sysLockdown, queryOpts...)
}

// MemoryPFA gets the MemoryPFA instance associated with this manager.
func (m *Manager) MemoryPFA(queryOpts ...common.QueryGroupOption) (*MemoryPFA, error) {
	return m.MemoryPFAWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// MemoryPFAWithContext gets the MemoryPFA instance associated with this manager.
func (m *Manager) MemoryPFAWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*MemoryPFA, error) {
	return GetMemoryPFAWithContext(ctx, m.GetClient(), m.memoryPFA, queryOpts...)
}

// MemoryHealthComp gets the MemoryHealthComp instance associated with this manager.
func (m *Manager) MemoryHealthComp(queryOpts ...common.QueryGroupOption) (*MemoryHealthComp, error) {
	return m.MemoryHealthCompWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// MemoryHealthCompWithContext gets the MemoryHealthComp instance associated with this manager.
func (m *Manager) MemoryHealthCompWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*MemoryHealthComp, error) {
	return GetMemoryHealthCompWithContext(ctx, m.GetClient(), m.memoryHealthComp, queryOpts...)
}

// Snooping gets the Snooping instance associated with this manager.
func (m *Manager) Snooping(queryOpts ...common.QueryGroupOption) (*Snooping, error) {
	return m.SnoopingWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// SnoopingWithContext gets the Snooping instance associated with this manager.
func (m *Manager) SnoopingWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*Snooping, error) {
	return GetSnoopingWithContext(ctx, m.GetClient(), m.snooping, queryOpts...)
}

// FanMode gets the FanMode instance associated with this manager.
func (m *Manager) FanMode(queryOpts ...common.QueryGroupOption) (*FanMode, error) {
	return m.FanModeWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// FanModeWithContext gets the FanMode instance associated with this manager.
func (m *Manager) FanModeWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*FanMode, error) {
	return GetFanModeWithContext(ctx, m.GetClient(), m.fanMode, queryOpts...)
}

// IKVM gets the IKVM instance associated with this manager.
func (m *Manager) IKVM(queryOpts ...common.QueryGroupOption) (*IKVM, error) {
	return m.IKVMWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// IKVMWithContext gets the IKVM instance associated with this manager.
func (m *Manager) IKVMWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*IKVM, error) {
	return GetIKVMWithContext(ctx, m.GetClient(), m.iKVM, queryOpts...)
}

// KCSInterface gets the KCSInterface instance associated with this manager.
func (m *Manager) KCSInterface(queryOpts ...common.QueryGroupOption) (*KCSInterface, error) {
	return m.KCSInterfaceWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// KCSInterfaceWithContext gets the KCSInterface instance associated with this manager.
func (m *Manager) KCSInterfaceWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*KCSInterface, error) {
	return GetKCSInterfaceWithContext(ctx, m.GetClient(), m.kcsInterface, queryOpts...)
}

// LLDP gets the LLDP instance associated with this manager.
func (m *Manager) LLDP(queryOpts ...common.QueryGroupOption) (*LLDP, error) {
	return m.LLDPWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// LLDPWithContext gets the LLDP instance associated with this manager.
func (m *Manager) LLDPWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*LLDP, error) {
	return GetLLDPWithContext(ctx, m.GetClient(), m.lldp, queryOpts...)
}

// LicenseManager gets the LicenseManager instance associated with this manager.
func (m *Manager) LicenseManager(queryOpts ...common.QueryGroupOption) (*LicenseManager, error) {
	return m.LicenseManagerWithContext(common.ContextOf(m.GetClient()), queryOpts...)
}

// LicenseManagerWithContext gets the LicenseManager instance associated with this manager.
func (m *Manager) LicenseManagerWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*LicenseManager, error) {
	return GetLicenseManagerWithContext(ctx, m.GetClient(), m.licenseManager, queryOpts...)
}

// ManagerConfigReset resets the BMC to factory defaults.
func (m *Manager) ManagerConfigReset(option ManagerConfigResetOption) error {
	return m.ManagerConfigResetWithContext(common.ContextOf(m.GetClient()), option)
}

// ManagerConfigResetWithContext resets the BMC to factory defaults.
func (m *Manager) ManagerConfigResetWithContext(ctx context.Context, option ManagerConfigResetOption) error {
	if m.managerConfigResetTarget == "" {
		return errors.New("manager config reset not supported by this system")
	}

	return m.PostWithContext(ctx, m.managerConfigResetTarget, map[string]interface{}{"Option": option})
}
