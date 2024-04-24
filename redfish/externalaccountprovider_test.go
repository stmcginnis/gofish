//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var externalAccountProviderBody = `{
	"@odata.type": "#ExternalAccountProvider.v1_7_1.ExternalAccountProvider",
	"Id": "ExternalRedfishService",
	"Name": "Remote Redfish Service",
	"Description": "Remote Redfish Service providing additional Accounts to this Redfish Service",
	"AccountProviderType": "RedfishService",
	"ServiceAddresses": [
	  "http://redfish.dmtf.org/redfish/v1/AccountService"
	],
	"Authentication": {
	  "AuthenticationType": "Token",
	  "Token": null
	},
	"RemoteRoleMapping": [
	  {
		"RemoteGroup": "Admin",
		"LocalRole": "Administrator"
	  },
	  {
		"RemoteGroup": "Operator",
		"LocalRole": "Operator"
	  },
	  {
		"RemoteGroup": "ReadOnly",
		"LocalRole": "ReadOnly"
	  }
	],
	"@odata.id": "/redfish/v1/AccountService/ExternalAccountProviders/ExternalRedfishService"
  }`

// TestExternalAccountProvider tests the parsing of ExternalAccountProvider objects.
func TestExternalAccountProvider(t *testing.T) {
	var result ExternalAccountProvider
	err := json.NewDecoder(strings.NewReader(externalAccountProviderBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "ExternalRedfishService", result.ID)
	assertEquals(t, "Remote Redfish Service", result.Name)
	assertEquals(t, "RedfishService", string(result.AccountProviderType))
	assertEquals(t, "http://redfish.dmtf.org/redfish/v1/AccountService", result.ServiceAddresses[0])
	assertEquals(t, "Token", string(result.Authentication.AuthenticationType))
	assertEquals(t, "Admin", result.RemoteRoleMapping[0].RemoteGroup)
	assertEquals(t, "Administrator", result.RemoteRoleMapping[0].LocalRole)
}
