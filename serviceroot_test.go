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
		"ComponentIntegrity": {
			"@odata.id": "/redfish/v1/ComponentIntegrity"
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

	if result.AccountServiceLink != "/redfish/v1/Accounts" {
		t.Errorf("Invalid AccountService link: %s", result.AccountServiceLink)
	}

	if result.CertificateServiceLink != "/redfish/v1/Certificates" {
		t.Errorf("Invalid CertificateService link: %s", result.CertificateServiceLink)
	}

	if result.ChassisLink != "/redfish/v1/Chassis" {
		t.Errorf("Invalid Chassis link: %s", result.ChassisLink)
	}

	if result.CompositionServiceLink != "/redfish/v1/Compositions" {
		t.Errorf("Invalid CompositionService link: %s", result.CompositionServiceLink)
	}

	if result.EventServiceLink != "/redfish/v1/Events" {
		t.Errorf("Invalid EventService link: %s", result.EventServiceLink)
	}

	if result.FabricsLink != "/redfish/v1/Fabrics" {
		t.Errorf("Invalid Fabrics link: %s", result.FabricsLink)
	}

	if result.JobServiceLink != "/redfish/v1/Jobs" {
		t.Errorf("Invalid JobService link: %s", result.JobServiceLink)
	}

	if result.JSONSchemasLink != "/redfish/v1/JsonSchemas" {
		t.Errorf("Invalid JsonSchemas link: %s", result.JSONSchemasLink)
	}

	if result.Links.Sessions != "/redfish/v1/Sessions" {
		t.Errorf("Invalid Sessions link: %s", result.Links.Sessions)
	}

	if result.ManagersLink != "/redfish/v1/Managers" {
		t.Errorf("Invalid Managers link: %s", result.ManagersLink)
	}

	if !result.ProtocolFeaturesSupported.ExcerptQuery {
		t.Error("ExcerptQuery should be true")
	}

	if result.RegistriesLink != "/redfish/v1/Registries" {
		t.Errorf("Invalid Registries link: %s", result.RegistriesLink)
	}

	if result.ResourceBlocksLink != "/redfish/v1/ResourceBlocks" {
		t.Errorf("Invalid ResourceBlocks link: %s", result.ResourceBlocksLink)
	}

	if result.SessionServiceLink != "/redfish/v1/SessionService" {
		t.Errorf("Invalid SessionService link: %s", result.SessionServiceLink)
	}

	if result.StorageServicesLink != "/redfish/v1/StorageServices" {
		t.Errorf("Invalid StorageServices link: %s", result.StorageServicesLink)
	}

	if result.StorageSystemsLink != "/redfish/v1/StorageSystems" {
		t.Errorf("Invalid StorageSystems link: %s", result.StorageSystemsLink)
	}

	if result.SystemsLink != "/redfish/v1/Systems" {
		t.Errorf("Invalid Systems link: %s", result.SystemsLink)
	}

	if result.TasksLink != "/redfish/v1/Tasks" {
		t.Errorf("Invalid Tasks link: %s", result.TasksLink)
	}

	if result.TelemetryServiceLink != "/redfish/v1/TelemetryService" {
		t.Errorf("Invalid TelemetryService link: %s", result.TelemetryServiceLink)
	}

	if result.UpdateServiceLink != "/redfish/v1/UpdateService" {
		t.Errorf("Invalid UpdateService link: %s", result.UpdateServiceLink)
	}
}

const oem = `{
			"Dell": {
			  "@odata.context": "/redfish/v1/$metadata#DellServiceRoot.DellServiceRoot",
			  "@odata.type": "#DellServiceRoot.v1_0_0.DellServiceRoot",
				"ManagerMACAddress": "4c:d9:8f:00:11:34",
				"ServiceTag": "C3H3853"
			}
		}`

var serviceRootBodyOEM = strings.NewReader(
	`{
		"@odata.context": "/redfish/v1/$metadata#ServiceRoot.ServiceRoot",
		"@odata.type": "#ServiceRoot.v1_5_1.ServiceRoot",
		"@odata.id": "/redfish/v1/ServiceRoot",
		"Id": "ServiceRoot-1",
		"Name": "ServiceRootOne",
		"Description": "ServiceRoot One",
		"Oem":` + oem + `,
		"Vendor": "Dell"
	}`)

// TestServiceRootOEM tests the parsing of ServiceRoot objects with OEM field.
func TestServiceRootOEM(t *testing.T) {
	var result Service
	err := json.NewDecoder(serviceRootBodyOEM).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "ServiceRoot-1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Oem == nil {
		t.Error("Invalid Oem link")
	}
	oemExp := strings.TrimSpace(oem)
	oemObt := strings.TrimSpace(string(result.Oem))
	if oemExp != oemObt {
		t.Errorf("Expect\n%s\n,Obtain\n%s", oemExp, oemObt)
	}
}
