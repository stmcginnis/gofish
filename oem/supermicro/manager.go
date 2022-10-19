//
// SPDX-License-Identifier: BSD-3-Clause
//

package supermicro

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

type Manager struct {
	redfish.Manager
	Oem ManagerOem
}

type ManagerOem struct {
	Supermicro ManagerOemSupermicro
}

type ManagerOemSupermicro struct {
	OdataType string `json:"@odata.type"`
	NTP       common.Link
	Syslog    common.Link
	IKVM      common.Link
	m         *redfish.Manager
}

func FromManager(manager *redfish.Manager) (Manager, error) {
	var oem ManagerOem
	_ = json.Unmarshal(manager.Oem, &oem)

	oem.Supermicro.m = manager

	return Manager{
		Manager: *manager,
		Oem:     oem,
	}, nil
}

func (manager *ManagerOemSupermicro) NTPs() (*NTP, error) {
	return GetNTP(manager.m.Client, string(manager.NTP))
}

func (manager *ManagerOemSupermicro) Syslogs() (*Syslog, error) {
	return GetSyslog(manager.m.Client, string(manager.Syslog))
}

func (manager *ManagerOemSupermicro) IKVMs() (*IKVM, error) {
	return GetIKVM(manager.m.Client, string(manager.IKVM))
}
