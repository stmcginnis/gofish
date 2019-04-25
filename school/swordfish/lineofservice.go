// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// LineOfService is the common attributes of LOS objects
type LineOfService struct {
	Name string
}

// LocationInfo has device address and location details
type LocationInfo struct {
	Address1          string
	Address2          string
	Address3          string
	Building          string
	City              string
	Country           string
	GPSCoords         string
	Item              string
	OtherLocationInfo string
	PostalCode        string
	Rack              string
	Room              string
	Row               string
	Shelf             string
	State             string
	Territory         string
}

// DataProtectionLineOfService is a line of service for data protection
type DataProtectionLineOfService struct {
	LineOfService
	RecoveryGeographicObjective string
	RecoveryPointObjective      string
	RecoveryTimeObjective       string
	ReplicaAccessLocation       LocationInfo
	ReplicaType                 string
	Schedule                    string
	replicaOption               string
}

// UnmarshalJSON unmarshals a DataProtectionLineOfService object
func (s *DataProtectionLineOfService) UnmarshalJSON(b []byte) error {
	type temp DataProtectionLineOfService
	type linkReference struct {
		replicaOption common.Link
	}
	var t struct {
		temp
		Links linkReference
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*s = DataProtectionLineOfService(t.temp)

	// Extract the links to other entities
	s.replicaOption = string(t.Links.replicaOption)

	return nil
}

// DataSecurityLineOfService is a line of service for data security
type DataSecurityLineOfService struct {
	LineOfService
	AntivirusEngineProvider   string
	AntivirusScanPolicies     []string
	ChannelEncruptionStrength string
	DataSantizationPolicy     string
	HostAuthenticationType    string
	MediaEncryptionStrength   string
	SecureChannelProtocol     string
	UserAuthenticationType    string
}

// DataStorageLineOfService has information about data storage attributes
type DataStorageLineOfService struct {
	LineOfService
	ProvisioningPolicy    string
	RecoveryTimeObjective int
	SpaceEfficient        bool
}

// IOConnectivityLineOfService has information about IO connectivity
type IOConnectivityLineOfService struct {
	LineOfService
	AccessProtocol                    string
	MaxSupportedIOOperationsPerSecond string `json:"MaxSupportedIoOperationsPerSecond"`
}

// IOWorkload has workload details
type IOWorkload struct {
	Name string
}

// IOPerformanceLineOfService has IO performance characteristics
type IOPerformanceLineOfService struct {
	LineOfService
	AverageIOOperationLatencyMicroseconds int `json:"AverageIoOperationLatencyMicroseconds"`
	IOWorkload                            IOWorkload
	IOOperationsPerSecondIsLimitedBoolean bool `json:"IoOperationsPerSecondIsLimitedBoolean"`
	MaxIOOperationsPerSecondPerTerabyte   int  `json:"MaxIoOperationsPerSecondPerTerabyte"`
	SamplePeriod                          string
}
