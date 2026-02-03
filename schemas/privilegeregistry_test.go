//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"strings"
	"testing"
)

var privilegeRegistryBody = `{
	"@odata.type": "#PrivilegeRegistry.v1_1_4.PrivilegeRegistry",
	"Id": "Contoso_1.0.1_PrivilegeRegistry",
	"Name": "Privilege Map",
	"PrivilegesUsed": [
	  "Login",
	  "ConfigureManager",
	  "ConfigureUsers",
	  "ConfigureComponents",
	  "ConfigureSelf"
	],
	"OEMPrivilegesUsed": [],
	"Mappings": [
	  {
		"Entity": "Manager",
		"OperationMap": {
		  "GET": [
			{
			  "Privilege": [
				"Login"
			  ]
			}
		  ],
		  "HEAD": [
			{
			  "Privilege": [
				"Login"
			  ]
			}
		  ],
		  "PATCH": [
			{
			  "Privilege": [
				"ConfigureManager"
			  ]
			}
		  ],
		  "POST": [
			{
			  "Privilege": [
				"ConfigureManager"
			  ]
			}
		  ],
		  "PUT": [
			{
			  "Privilege": [
				"ConfigureManager"
			  ]
			}
		  ],
		  "DELETE": [
			{
			  "Privilege": [
				"ConfigureManager"
			  ]
			}
		  ]
		}
	  }
	],
	"@odata.id": "/redfish/v1/JobService"
  }`

// TestPrivilegeRegistry tests the parsing of PrivilegeRegistry objects.
func TestPrivilegeRegistry(t *testing.T) {
	var result PrivilegeRegistry
	err := json.NewDecoder(strings.NewReader(privilegeRegistryBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "Contoso_1.0.1_PrivilegeRegistry", result.ID)
	assertEquals(t, "Privilege Map", result.Name)
	assertEquals(t, "Login", string(result.PrivilegesUsed[0]))
	assertEquals(t, "ConfigureComponents", string(result.PrivilegesUsed[3]))
	assertEquals(t, "ConfigureSelf", string(result.PrivilegesUsed[4]))
	assertEquals(t, "Manager", result.Mappings[0].Entity)
	assertEquals(t, "Login", result.Mappings[0].OperationMap.GET[0].Privilege[0])
	assertEquals(t, "ConfigureManager", result.Mappings[0].OperationMap.DELETE[0].Privilege[0])
}
