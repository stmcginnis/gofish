//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var keyPolicyBody = `{
	"@odata.type": "#KeyPolicy.v1_0_0.KeyPolicy",
	"Id": "0",
	"Name": "Default NVMeoF Key Policy",
	"IsDefault": true,
	"KeyPolicyType": "NVMeoF",
	"NVMeoF": {
	  "SecurityTransportAllowList": [
		"TLSv2",
		"TLSv3"
	  ],
	  "CipherSuiteAllowList": [
		"TLS_AES_128_GCM_SHA256",
		"TLS_AES_256_GCM_SHA384"
	  ],
	  "SecurityProtocolAllowList": [
		"DHHC",
		"TLS_PSK"
	  ],
	  "DHGroupAllowList": [
		"FFDHE2048",
		"FFDHE3072",
		"FFDHE4096",
		"FFDHE6144",
		"FFDHE8192"
	  ],
	  "SecureHashAllowList": [
		"SHA384",
		"SHA512"
	  ]
	},
	"@odata.id": "/redfish/v1/KeyService/NVMeoFKeyPolicies/0"
  }`

// TestKeyPolicy tests the parsing of KeyPolicy objects.
func TestKeyPolicy(t *testing.T) {
	var result KeyPolicy
	err := json.NewDecoder(strings.NewReader(keyPolicyBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "0", result.ID)
	assertEquals(t, "Default NVMeoF Key Policy", result.Name)
	assertEquals(t, "NVMeoF", string(result.KeyPolicyType))
	assertEquals(t, "TLSv3", string(result.NVMeoF.SecurityTransportAllowList[1]))
	assertEquals(t, "TLS_AES_256_GCM_SHA384", string(result.NVMeoF.CipherSuiteAllowList[1]))
	assertEquals(t, "DHHC", string(result.NVMeoF.SecurityProtocolAllowList[0]))
	assertEquals(t, "SHA512", string(result.NVMeoF.SecureHashAllowList[1]))

	if !result.IsDefault {
		t.Error("Expected to be the default")
	}
}
