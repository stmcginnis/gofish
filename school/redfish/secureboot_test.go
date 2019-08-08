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

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var secureBootBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#SecureBoot.SecureBoot",
		"@odata.type": "#SecureBoot.v1_0_5.SecureBoot",
		"@odata.id": "/redfish/v1/SecureBoot",
		"Id": "SecureBoot-1",
		"Name": "SecureBootOne",
		"Description": "SecureBoot One",
		"SecureBootCurrentBoot": "Enabled",
		"SecureBootEnable": true,
		"SecureBootMode": "UserMode"
	}`)

// TestSecureBoot tests the parsing of SecureBoot objects.
func TestSecureBoot(t *testing.T) {
	var result SecureBoot
	err := json.NewDecoder(secureBootBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "SecureBoot-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "SecureBootOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.SecureBootCurrentBoot != EnabledSecureBootCurrentBootType {
		t.Errorf("Invalid SecureBootCurrentBoot: %s", result.SecureBootCurrentBoot)
	}

	if !result.SecureBootEnable {
		t.Error("SecureBootEnable should be true")
	}

	if result.SecureBootMode != UserModeSecureBootModeType {
		t.Errorf("Invalid SecureBootMode: %s", result.SecureBootMode)
	}
}
