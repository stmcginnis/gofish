//
// SPDX-License-Identifier: BSD-3-Clause
//

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
