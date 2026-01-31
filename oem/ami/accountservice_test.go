//
// SPDX-License-Identifier: BSD-3-Clause
//

package ami

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/schemas"
)

var accountServiceBody = `{
  "@odata.type": "#AccountService.v1_7_2.AccountService",
  "@odata.id": "/redfish/v1/AccountService",
  "Id": "AccountService",
  "Name": "Account Service",
  "Oem": {
    "Ami": {
      "@odata.type": "#AMIAccountService.v1_0_0.AMIAccountService",
      "Configuration": {
        "@odata.id": "/redfish/v1/AccountService/Oem/Ami/Configurations"
      }
    }
  },
  "PrivilegeMap": {
    "@odata.id": "/redfish/v1/Registries/Redfish_1.4.0_PrivilegeRegistry.json"
  },
  "Roles": {
    "@odata.id": "/redfish/v1/AccountService/Roles"
  },
  "ServiceEnabled": true,
  "Status": {
    "Health": "OK",
    "State": "Enabled"
  }
}`

// TestAMIAccountService tests the parsing of the AccountService.
func TestAMIAccountService(t *testing.T) {
	as := &schemas.AccountService{}
	if err := json.Unmarshal([]byte(accountServiceBody), as); err != nil {
		t.Fatalf("error decoding json: %v", err)
	}

	accountService, err := FromAccountService(as)
	if err != nil {
		t.Fatalf("error getting oem info: %v", err)
	}

	if accountService.ID != "AccountService" {
		t.Errorf("unexpected ID: %s", accountService.ID)
	}

	if accountService.configuration != "/redfish/v1/AccountService/Oem/Ami/Configurations" {
		t.Errorf("unexpected configuration link: %s", accountService.configuration)
	}
}

var accountServiceConfigurationsBody = `{
  "@odata.context": "/redfish/v1/$metadata#AMIAccountServiceConfigurations.AMIAccountServiceConfigurations",
  "@odata.etag": "\"1729105654\"",
  "@odata.id": "/redfish/v1/AccountService/Oem/Ami/Configurations",
  "@odata.type": "#AMIAccountServiceConfigurations.v1_0_0.AMIAccountServiceConfigurations",
  "Id": "Configurations",
  "Name": "AccountService Configurations",
  "PAMEnabled": true,
  "PAMOrder": [
    "IPMI",
    "LDAP",
    "ACTIVE DIRECTORY",
    "RADIUS"
  ]
}`

// TestAMIAccountServiceConfigurations tests the parsing of the AccountServiceConfigurations.
func TestAMIAccountServiceConfigurations(t *testing.T) {
	var result AccountServiceConfigurations
	err := json.NewDecoder(strings.NewReader(accountServiceConfigurationsBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "Configurations" {
		t.Errorf("unexpected ID: %s", result.ID)
	}

	if !result.PAMEnabled {
		t.Errorf("unexpected PAMEnabled: %t", result.PAMEnabled)
	}

	if len(result.PAMOrder) != 4 {
		t.Errorf("unexpected PAMOrder length: %d", len(result.PAMOrder))
	}

	if result.PAMOrder[0] != "IPMI" {
		t.Errorf("unexpected PAMOrder[0]: %s", result.PAMOrder[0])
	}
}
