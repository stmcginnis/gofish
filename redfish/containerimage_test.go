//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var containerImageBody = `{
	"@odata.type": "#ContainerImage.v1_0_0.ContainerImage",
	"Id": "WebLogic",
	"Name": "Contoso Internal Web Business Logic 1.0",
	"CreateTime": "2021-02-06T22:49:02Z",
	"Type": "OCI",
	"Version": "1.0.0",
	"Status": {
	  "State": "Enabled"
	},
	"ProgrammaticId": "2fbd319a987e5265aae45b7e786dead51d1aae48b7bea42bcfc91a62934ca37f",
	"SizeBytes": 373293056,
	"Links": {
	  "Containers": [
		{
		  "@odata.id": "/redfish/v1/Systems/VM1/OperatingSystem/Containers/WebBusinessLogic"
		}
	  ]
	},
	"@odata.id": "/redfish/v1/Systems/VM1/OperatingSystem/ContainerImages/WebLogic"
  }`

// TestContainerImage tests the parsing of ContainerImage objects.
func TestContainerImage(t *testing.T) {
	var result ContainerImage
	err := json.NewDecoder(strings.NewReader(containerImageBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "WebLogic", result.ID)
	assertEquals(t, "Contoso Internal Web Business Logic 1.0", result.Name)
	assertEquals(t, "OCI", string(result.Type))
	assertEquals(t, "2fbd319a987e5265aae45b7e786dead51d1aae48b7bea42bcfc91a62934ca37f", result.ProgrammaticID)

	if result.SizeBytes != 373293056 {
		t.Errorf("Unexpected memory size, got %d", result.SizeBytes)
	}

	assertEquals(t, "/redfish/v1/Systems/VM1/OperatingSystem/Containers/WebBusinessLogic", result.containers[0])
}
