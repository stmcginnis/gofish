//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var registeredClientBody = `{
	"@odata.type": "#RegisteredClient.v1_1_1.RegisteredClient",
	"Id": "2",
	"Name": "ContosoConfigure",
	"ClientType": "Configure",
	"CreatedDate": "2021-09-25T20:12:24Z",
	"Description": "Contoso manager access",
	"ExpirationDate": "2022-10-03T20:00:00Z",
	"ManagedResources": [
	  {
		"ManagedResourceURI": "/redfish/v1/Systems",
		"PreferExclusive": true,
		"IncludesSubordinates": true
	  },
	  {
		"ManagedResourceURI": "/redfish/v1/Chassis",
		"PreferExclusive": true,
		"IncludesSubordinates": true
	  }
	],
	"ClientURI": "https://4.5.6.2/ContosoManager",
	"@odata.id": "/redfish/v1/RegisteredClients/2"
  }`

// TestRegisteredClient tests the parsing of RegisteredClient objects.
func TestRegisteredClient(t *testing.T) {
	var result RegisteredClient
	err := json.NewDecoder(strings.NewReader(registeredClientBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "2", result.ID)
	assertEquals(t, "ContosoConfigure", result.Name)
	assertEquals(t, "Configure", string(result.ClientType))
	assertEquals(t, "https://4.5.6.2/ContosoManager", result.ClientURI)
	assertEquals(t, "/redfish/v1/Systems", result.ManagedResources[0].ManagedResourceURI)

	if !result.ManagedResources[1].PreferExclusive {
		t.Error("Expected PreferExclusive to be true")
	}

	if !result.ManagedResources[1].IncludesSubordinates {
		t.Error("Expected IncludesSubordinates to be true")
	}
}
