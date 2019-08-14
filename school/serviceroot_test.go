//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"encoding/json"
	"strings"
	"testing"
)

var serviceRootBody = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#ServiceRoot.ServiceRoot",
		"@odata.type": "#ServiceRoot.v1_5_1.ServiceRoot",
		"@odata.id": "/redfish/v1/ServiceRoot",
		"Id": "ServiceRoot-1",
		"Name": "ServiceRootOne",
		"Description": "ServiceRoot One",
		"AccountService": {
			"@odata.id": "/redfish/v1/Accounts"
		},
		"CertificateService": {
			"@odata.id": "/redfish/v1/Certificates"
		},
		"Chassis": {
			"@odata.id": "/redfish/v1/Chassis"
		},
		"CompositionService": {
			"@odata.id": "/redfish/v1/Compositions"
		},
		"EventService": {
			"@odata.id": "/redfish/v1/Events"
		},
		"Fabrics": {
			"@odata.id": "/redfish/v1/Fabrics"
		},
		"JobService": {
			"@odata.id": "/redfish/v1/Jobs"
		},
		"JsonSchemas": {
			"@odata.id": "/redfish/v1/JsonSchemas"
		},
		"Links": {
			"Sessions": {
				"@odata.id": "/redfish/v1/Sessions"
			}
		},
		"Managers": {
			"@odata.id": "/redfish/v1/Managers"
		},
		"Product": "Product One",
		"ProtocolFeaturesSupported": {
			"ExcerptQuery": true,
			"ExpandQuery": {
				"ExpandAll": true,
				"Levels": true,
				"Links": true,
				"MaxLevels": 21,
				"NoLinks": true
			},
			"FilterQuery": true,
			"OnlyMemberQuery": true,
			"SelectQuery": true
		},
		"RedfishVersion": "1.2.3",
		"Registries": {
			"@odata.id": "/redfish/v1/Registries"
		},
		"ResourceBlocks": {
			"@odata.id": "/redfish/v1/ResourceBlocks"
		},
		"SessionService": {
			"@odata.id": "/redfish/v1/SessionService"
		},
		"StorageServices": {
			"@odata.id": "/redfish/v1/StorageServices"
		},
		"StorageSystems": {
			"@odata.id": "/redfish/v1/StorageSystems"
		},
		"Systems": {
			"@odata.id": "/redfish/v1/Systems"
		},
		"Tasks": {
			"@odata.id": "/redfish/v1/Tasks"
		},
		"TelemetryService": {
			"@odata.id": "/redfish/v1/TelemetryService"
		},
		"UUID": "ae058175-af1d-40fe-ad5b-c1ab79de2c65",
		"UpdateService": {
			"@odata.id": "/redfish/v1/UpdateService"
		},
		"Vendor": "Acme Services"
	}`)

// TestServiceRoot tests the parsing of ServiceRoot objects.
func TestServiceRoot(t *testing.T) {
	var result Service
	err := json.NewDecoder(serviceRootBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "ServiceRoot-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "ServiceRootOne" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.accountService != "/redfish/v1/Accounts" {
		t.Errorf("Invalid AccountService link: %s", result.accountService)
	}

	if result.certificateService != "/redfish/v1/Certificates" {
		t.Errorf("Invalid CertificateService link: %s", result.certificateService)
	}

	if result.chassis != "/redfish/v1/Chassis" {
		t.Errorf("Invalid Chassis link: %s", result.chassis)
	}

	if result.compositionService != "/redfish/v1/Compositions" {
		t.Errorf("Invalid CompositionService link: %s", result.compositionService)
	}

	if result.eventService != "/redfish/v1/Events" {
		t.Errorf("Invalid EventService link: %s", result.eventService)
	}

	if result.fabrics != "/redfish/v1/Fabrics" {
		t.Errorf("Invalid Fabrics link: %s", result.fabrics)
	}

	if result.jobService != "/redfish/v1/Jobs" {
		t.Errorf("Invalid JobService link: %s", result.jobService)
	}

	if result.jsonSchemas != "/redfish/v1/JsonSchemas" {
		t.Errorf("Invalid JsonSchemas link: %s", result.jsonSchemas)
	}

	if result.sessions != "/redfish/v1/Sessions" {
		t.Errorf("Invalid Sessions link: %s", result.sessions)
	}

	if result.managers != "/redfish/v1/Managers" {
		t.Errorf("Invalid Managers link: %s", result.managers)
	}

	if !result.ProtocolFeaturesSupported.ExcerptQuery {
		t.Error("ExcerptQuery should be true")
	}

	if result.registries != "/redfish/v1/Registries" {
		t.Errorf("Invalid Registries link: %s", result.registries)
	}

	if result.resourceBlocks != "/redfish/v1/ResourceBlocks" {
		t.Errorf("Invalid ResourceBlocks link: %s", result.resourceBlocks)
	}

	if result.sessionService != "/redfish/v1/SessionService" {
		t.Errorf("Invalid SessionService link: %s", result.sessionService)
	}

	if result.storageServices != "/redfish/v1/StorageServices" {
		t.Errorf("Invalid StorageServices link: %s", result.storageServices)
	}

	if result.storageSystems != "/redfish/v1/StorageSystems" {
		t.Errorf("Invalid StorageSystems link: %s", result.storageSystems)
	}

	if result.systems != "/redfish/v1/Systems" {
		t.Errorf("Invalid Systems link: %s", result.systems)
	}

	if result.tasks != "/redfish/v1/Tasks" {
		t.Errorf("Invalid Tasks link: %s", result.tasks)
	}

	if result.telemetryService != "/redfish/v1/TelemetryService" {
		t.Errorf("Invalid TelemetryService link: %s", result.telemetryService)
	}

	if result.updateService != "/redfish/v1/UpdateService" {
		t.Errorf("Invalid UpdateService link: %s", result.updateService)
	}
}
