//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var keyBody = `{
	"@odata.type": "#Key.v1_4_0.Key",
	"Id": "0",
	"Name": "NVMeoF key 0, target subsystem",
	"KeyType": "NVMeoF",
	"KeyString": "DHHC-1:00:ia6zGodOr4SEG0Zzaw398rpY0wqipUWj4jWjUh4HWUz6aQ2n:",
	"NVMeoF": {
	  "NQN": "nqn.corp.com:nvme:target-subsystem-0001",
	  "SecurityProtocolType": "DHHC",
	  "HostKeyId": "1",
	  "SecureHashAllowList": [
		"SHA384",
		"SHA512"
	  ]
	},
	"@odata.id": "/redfish/v1/KeyService/NVMeoFSecrets/0"
  }`

// TestKey tests the parsing of Key objects.
func TestKey(t *testing.T) {
	var result Key
	err := json.NewDecoder(strings.NewReader(keyBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "0", result.ID)
	assertEquals(t, "NVMeoF key 0, target subsystem", result.Name)
	assertEquals(t, "NVMeoF", string(result.KeyType))
	assertEquals(t, "DHHC-1:00:ia6zGodOr4SEG0Zzaw398rpY0wqipUWj4jWjUh4HWUz6aQ2n:", result.KeyString)
	assertEquals(t, "nqn.corp.com:nvme:target-subsystem-0001", result.NVMeoF.NQN)
	assertEquals(t, "DHHC", string(result.NVMeoF.SecurityProtocolType))
	assertEquals(t, "1", result.NVMeoF.HostKeyID)
	assertEquals(t, "SHA512", string(result.NVMeoF.SecureHashAllowList[1]))
}
