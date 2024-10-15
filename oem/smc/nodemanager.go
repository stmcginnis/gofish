//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"fmt"

	"github.com/stmcginnis/gofish/common"
)

type NodeManagerCapabilities struct {
	DomainID              string
	PolicyType            string
	MaxConcurrentSettings int
	MaxValueAfterReset    uint
	MinValueAfterReset    uint
	MaxCorrectionTime     uint64
	MinCorrectionTime     uint64
	MaxReportingPeriod    uint64
	MinReportingPeriod    uint64
	DomainLimitingScope   int
}

type NodeManagerStatistics struct {
	Mode            string
	DomainID        string
	CurrentValue    uint64
	MaximumValue    uint64
	MinimumValue    uint64
	AverageValue    uint64
	ReportingPeriod uint64
}

type NodeManagerPolicy struct {
	DomainID               string
	PolicyID               uint
	PolicyType             uint
	PolicyExceptionActions uint
	PowerLimit             uint
	CorrectionTimeLimit    uint
	PolicyTriggerLimit     uint
	StatReportingPeriod    uint
}

// NodeManager is the node manager instance associated with the system.
// This Redfish API can only be supported on Intel platforms with Intel ME.
type NodeManager struct {
	common.Entity
	Capabilities       []NodeManagerCapabilities
	Statistics         []NodeManagerStatistics
	IntelPsysEnabled   bool
	IntelPsysSupported bool
	Version            struct {
		IntelNMVersion string
		IPMIVersion    string
		PatchVersion   uint
		MajorRevision  uint
		MinorRevision  uint
	}
	SelfTest struct {
		MajorCode  uint
		MinorCode  uint
		ImageFlags string
	}
	Policy []NodeManagerPolicy

	clearAllPoliciesTarget string
}

// UnmarshalJSON unmarshals a NodeManager object from the raw JSON.
func (nm *NodeManager) UnmarshalJSON(b []byte) error {
	type temp NodeManager
	var t struct {
		temp
		Actions struct {
			ClearAllPolicies common.ActionTarget `json:"#SmcNodeManager.ClearAllPolicies"`
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*nm = NodeManager(t.temp)
	nm.clearAllPoliciesTarget = t.Actions.ClearAllPolicies.Target

	return nil
}

// GetNodeManager will get a NodeManager instance from the service.
func GetNodeManager(c common.Client, uri string) (*NodeManager, error) {
	return common.GetObject[NodeManager](c, uri)
}

// ClearAllPolicies clears the configured policies of the NodeManager.
func (nm *NodeManager) ClearAllPolicies() error {
	if nm.clearAllPoliciesTarget == "" {
		return fmt.Errorf("ClearAllPolicies is not supported by this system")
	}

	return nm.Post(nm.clearAllPoliciesTarget, nil)
}
