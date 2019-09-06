//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var dataSecurityLoSCapabilitiesBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#DataSecurityLoSCapabilities.DataSecurityLoSCapabilities",
		"@odata.type": "#DataSecurityLoSCapabilities.v1_1_2.DataSecurityLoSCapabilities",
		"@odata.id": "/redfish/v1/DataSecurityLoSCapabilities",
		"Id": "DataSecurityLoSCapabilities-1",
		"Name": "DataSecurityLoSCapabilitiesOne",
		"Description": "DataSecurityLoSCapabilities One",
		"SupportedAntivirusEngineProviders": [
			"Acme Antivirus",
			"MalwareRUs"
		],
		"SupportedAntivirusScanPolicies": [
			"OnFirstRead",
			"OnPatternUpdate",
			"OnUpdate"
		],
		"SupportedChannelEncryptionStrengths": [
			"Bits_128",
			"Bits_256"
		],
		"SupportedDataSanitizationPolicies": [
			"Clear",
			"CryptographicErase"
		],
		"SupportedHostAuthenticationTypes": [
			"PKI",
			"Ticket",
			"Password"
		],
		"SupportedLinesOfService": [
			{
				"@odata.id": "/redfish/v1/DataSecurityLoSCapabilities"
			}
		],
		"SupportedLinesOfService@odata.count": 1,
		"SupportedMediaEncryptionStrengths": [
			"Bits_128",
			"Bits_256"
		],
		"SupportedSecureChannelProtocols": [
			"TLS",
			"IPsec"
		],
		"SupportedUserAuthenticationTypes": [
			"PKI",
			"Ticket",
			"Password"
		]
	}`)

// TestDataSecurityLoSCapabilities tests the parsing of DataSecurityLoSCapabilities objects.
func TestDataSecurityLoSCapabilities(t *testing.T) {
	var result DataSecurityLoSCapabilities
	err := json.NewDecoder(dataSecurityLoSCapabilitiesBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "DataSecurityLoSCapabilities-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "DataSecurityLoSCapabilitiesOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.SupportedAntivirusScanPolicies[0] != OnFirstReadAntiVirusScanTrigger {
		t.Errorf("Invalid SupportedAntivirusScanPolicy: %s",
			result.SupportedAntivirusScanPolicies[0])
	}

	if result.SupportedChannelEncryptionStrengths[0] != Bits128KeySize {
		t.Errorf("Invalid SupportedChannelEncryptionStrength: %s",
			result.SupportedChannelEncryptionStrengths[0])
	}

	if result.SupportedMediaEncryptionStrengths[0] != Bits128KeySize {
		t.Errorf("Invalid SupportedMediaEncryptionStrength: %s",
			result.SupportedMediaEncryptionStrengths[0])
	}

	if result.SupportedDataSanitizationPolicies[1] != CryptographicEraseDataSanitizationPolicy {
		t.Errorf("Invalid SupportedDataSanitizationPolicy: %s", result.SupportedDataSanitizationPolicies[1])
	}

	if result.SupportedHostAuthenticationTypes[0] != PKIAuthenticationType {
		t.Errorf("Invalid SupportedHostAuthType: %s", result.SupportedHostAuthenticationTypes[0])
	}

	if result.SupportedUserAuthenticationTypes[0] != PKIAuthenticationType {
		t.Errorf("Invalid SupportedUserAuthType: %s", result.SupportedUserAuthenticationTypes[0])
	}
}
