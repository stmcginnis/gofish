//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var containerBody = `{
	"@odata.type": "#Container.v1_0_0.Container",
	"Id": "WebBusinessLogic",
	"Name": "Internal Web Business Logic",
	"StartTime": "2021-02-06T22:49:02Z",
	"Limits": {
	  "MemoryBytes": 4294967296,
	  "CPUCount": 1.5
	},
	"Status": {
	  "State": "Enabled"
	},
	"MountPoints": [
	  {
		"Source": "/opt/MyContainerStorage/WebConfig",
		"Destination": "/config"
	  }
	],
	"ProgrammaticId": "5584c257aba833892e1841cb77d898edc1f942f3bf901e7e0c390a504b9897a0",
	"EthernetInterfaces": {
	  "@odata.id": "/redfish/v1/Systems/VM1/OperatingSystem/Containers/WebBusinessLogic/EthernetInterfaces"
	},
	"Links": {
	  "ContainerImage": {
		"@odata.id": "/redfish/v1/Systems/VM1/OperatingSystem/ContainerImages/WebLogic"
	  }
	},
	"@odata.id": "/redfish/v1/Systems/VM1/OperatingSystem/Containers/WebBusinessLogic"
  }`

// TestContainer tests the parsing of Container objects.
func TestContainer(t *testing.T) {
	var result Container
	err := json.NewDecoder(strings.NewReader(containerBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "WebBusinessLogic", result.ID)
	assertEquals(t, "Internal Web Business Logic", result.Name)

	if result.Limits.MemoryBytes != 4294967296 {
		t.Errorf("Unexpected memory limit, got %d", result.Limits.MemoryBytes)
	}

	if result.Limits.CPUCount != 1.5 {
		t.Errorf("Unexpected CPU limit, got %.2f", result.Limits.CPUCount)
	}

	assertEquals(t, "/opt/MyContainerStorage/WebConfig", result.MountPoints[0].Source)
	assertEquals(t, "/config", result.MountPoints[0].Destination)
	assertEquals(t, "5584c257aba833892e1841cb77d898edc1f942f3bf901e7e0c390a504b9897a0", result.ProgrammaticID)
	assertEquals(t, "/redfish/v1/Systems/VM1/OperatingSystem/Containers/WebBusinessLogic/EthernetInterfaces", result.ethernetInterfaces)
	assertEquals(t, "/redfish/v1/Systems/VM1/OperatingSystem/ContainerImages/WebLogic", result.containerImage)
}
