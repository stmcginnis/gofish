//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var operatingSystemBody = `{
	"@odata.type": "#OperatingSystem.v1_0_1.OperatingSystem",
	"Id": "OperatingSystem",
	"Name": "OperatingSystem running on web-srv344",
	"UptimeSeconds": 6720,
	"Kernel": {
	  "Name": "Linux",
	  "Release": "5.10.13-x86_64",
	  "Version": "#1 SMP Thu Feb 4 13:56:42 EST 2021",
	  "Machine": "x86_64"
	},
	"Type": "Linux",
	"Status": {
	  "State": "Enabled",
	  "Health": "OK"
	},
	"Applications": {
	  "@odata.id": "/redfish/v1/Systems/VM1/OperatingSystem/Applications"
	},
	"ContainerEngines": [
	  {
		"Type": "Docker",
		"Version": "20.10.5",
		"SupportedImageTypes": [
		  "DockerV1",
		  "DockerV2",
		  "OCI"
		],
		"ManagementURIs": [
		  "https://192.168.0.12:5555"
		]
	  }
	],
	"ContainerImages": {
	  "@odata.id": "/redfish/v1/Systems/VM1/OperatingSystem/ContainerImages"
	},
	"Containers": {
	  "@odata.id": "/redfish/v1/Systems/VM1/OperatingSystem/Containers"
	},
	"Links": {
	  "SoftwareImage": {
		"@odata.id": "/redfish/v1/UpdateService/SoftwareInventory/ContosoLinux"
	  }
	},
	"@odata.id": "/redfish/v1/Systems/VM1/OperatingSystem"
  }`

// TestOperatingSystem tests the parsing of OperatingSystem objects.
func TestOperatingSystem(t *testing.T) {
	var result OperatingSystem
	err := json.NewDecoder(strings.NewReader(operatingSystemBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "OperatingSystem", result.ID)
	assertEquals(t, "OperatingSystem running on web-srv344", result.Name)
	assertEquals(t, "#1 SMP Thu Feb 4 13:56:42 EST 2021", result.Kernel.Version)
	assertEquals(t, "/redfish/v1/Systems/VM1/OperatingSystem/Applications", result.applications)
	assertEquals(t, "Docker", string(result.ContainerEngines[0].Type))
	assertEquals(t, "https://192.168.0.12:5555", result.ContainerEngines[0].ManagementURIs[0])
	assertEquals(t, "/redfish/v1/Systems/VM1/OperatingSystem/ContainerImages", result.containerImages)
	assertEquals(t, "/redfish/v1/Systems/VM1/OperatingSystem/Containers", result.containers)
	assertEquals(t, "/redfish/v1/UpdateService/SoftwareInventory/ContosoLinux", result.softwareImage)

	if result.UptimeSeconds != 6720 {
		t.Errorf("Unexpected UptimeSeconds value: %d", result.UptimeSeconds)
	}
}
