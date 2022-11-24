//
// SPDX-License-Identifier: BSD-3-Clause
//

package supermicro

import (
	"encoding/json"
	"fmt"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

type Manager struct {
	redfish.Manager
	Oem ManagerOem
}

type ManagerOem struct {
	Supermicro struct {
		OdataType string `json:"@odata.type"`
		NTP       common.Link
		Syslog    common.Link
		IKVM      common.Link
	} `json:"Supermicro"`
}

func FromManager(manager *redfish.Manager) (Manager, error) {
	var oem ManagerOem
	err := json.Unmarshal(manager.Oem, &oem)
	if err != nil {
		return Manager{}, fmt.Errorf("can't unmarshal OEM Manager: %v", err)
	}

	return Manager{
		Manager: *manager,
		Oem:     oem,
	}, nil
}

func (manager *Manager) NTP() (*NTP, error) {
	return GetNTP(manager.Client, string(manager.Oem.Supermicro.NTP))
}

func (manager *Manager) Syslog() (*Syslog, error) {
	return GetSyslog(manager.Client, string(manager.Oem.Supermicro.Syslog))
}

func (manager *Manager) IKVM() (*IKVM, error) {
	return GetIKVM(manager.Client, string(manager.Oem.Supermicro.IKVM))
}
