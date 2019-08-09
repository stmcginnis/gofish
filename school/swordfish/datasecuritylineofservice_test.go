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
	"strings"
	"testing"
)

var dataSecurityLineOfServiceBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#DataSecurityLineOfService.DataSecurityLineOfService",
		"@odata.type": "#DataSecurityLineOfService.v1_0_1.DataSecurityLineOfService",
		"@odata.id": "/redfish/v1/DataSecurityLineOfService",
		"Id": "DataSecurityLineOfService-1",
		"Name": "DataSecurityLineOfServiceOne",
		"Description": "DataSecurityLineOfService One",
		"AntivirusEngineProvider": "Acme Antivirus",
		"AntivirusScanPolicies": [
			"OnUpdate",
			"OnPatternUpdate"
		],
		"ChannelEncryptionStrength": "Bits_256",
		"DataSanitizationPolicy": "Clear",
		"HostAuthenticationType": "PKI",
		"MediaEncryptionStrength": "Bits_256",
		"SecureChannelProtocol": "TLS",
		"UserAuthenticationType": "Password"
	}`)

// TestDataSecurityLineOfService tests the parsing of DataSecurityLineOfService objects.
func TestDataSecurityLineOfService(t *testing.T) {
	var result DataSecurityLineOfService
	err := json.NewDecoder(dataSecurityLineOfServiceBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "DataSecurityLineOfService-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "DataSecurityLineOfServiceOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.AntivirusScanPolicies[0] != OnUpdateAntiVirusScanTrigger {
		t.Errorf("Invalid AntivirusScanPolicy: %s", result.AntivirusScanPolicies[0])
	}

	if result.ChannelEncryptionStrength != Bits256KeySize {
		t.Errorf("Invalid ChannelEncryptionStrength: %s", result.ChannelEncryptionStrength)
	}

	if result.SecureChannelProtocol != TLSSecureChannelProtocol {
		t.Errorf("Invalid SecureChannelProtocol: %s", result.SecureChannelProtocol)
	}
}
