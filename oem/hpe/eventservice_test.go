//
// SPDX-License-Identifier: BSD-3-Clause
//

package hpe

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/common"
)

const serviceRootBody = `{
    "@odata.context": "/redfish/v1/$metadata#ServiceRoot.ServiceRoot",
    "@odata.etag": "W/\"00000000\"",
    "@odata.id": "/redfish/v1/",
    "@odata.type": "#ServiceRoot.v1_5_1.ServiceRoot",
    "Id": "RootService",
    "AccountService": {
        "@odata.id": "/redfish/v1/AccountService/"
    },
    "Chassis": {
        "@odata.id": "/redfish/v1/Chassis/"
    },
    "EventService": {
        "@odata.id": "/redfish/v1/EventService/"
    },
    "JsonSchemas": {
        "@odata.id": "/redfish/v1/JsonSchemas/"
    },
    "Links": {
        "Sessions": {
            "@odata.id": "/redfish/v1/SessionService/Sessions/"
        }
    },
    "Managers": {
        "@odata.id": "/redfish/v1/Managers/"
    },
    "Name": "HPE RESTful Root Service",
    "Oem": {
        "Hpe": {
            "@odata.context": "/redfish/v1/$metadata#HpeiLOServiceExt.HpeiLOServiceExt",
            "@odata.type": "#HpeiLOServiceExt.v2_3_0.HpeiLOServiceExt",
            "Links": {
                "ResourceDirectory": {
                    "@odata.id": "/redfish/v1/ResourceDirectory/"
                }
            },
            "Manager": [{
                "DefaultLanguage": "en",
                "FQDN": "bmc.fqdn",
                "HostName": "bmc",
                "Languages": [{
                    "Language": "en",
                    "TranslationName": "English",
                    "Version": "2.10"
                }],
                "ManagerFirmwareVersion": "2.10",
                "ManagerType": "iLO 5",
                "Status": {
                    "Health": "OK"
                }
            }],
            "Moniker": {
                "ADVLIC": "iLO Advanced",
                "BMC": "iLO",
                "BSYS": "BladeSystem",
                "CLASS": "Baseboard Management Controller",
                "FEDGRP": "DEFAULT",
                "IPROV": "Intelligent Provisioning",
                "PRODABR": "iLO",
                "PRODFAM": "Integrated Lights-Out",
                "PRODGEN": "iLO 5",
                "PRODNAM": "Integrated Lights-Out 5",
                "PRODTAG": "HPE iLO 5",
                "STDLIC": "iLO Standard",
                "SUMABR": "SUM",
                "SUMGR": "Smart Update Manager",
                "SYSFAM": "ProLiant",
                "VENDABR": "HPE",
                "VENDNAM": "Hewlett Packard Enterprise",
                "WWW": "www.hpe.com",
                "WWWAHSV": "www.hpe.com/servers/ahsv",
                "WWWBMC": "www.hpe.com/info/ilo",
                "WWWDOC": "www.hpe.com/support/ilo-docs",
                "WWWERS": "www.hpe.com/services/getconnected",
                "WWWGLIS": "reserved for liconf URI",
                "WWWIOL": "www.hpe.com/info/insightonline",
                "WWWLIC": "www.hpe.com/info/ilo",
                "WWWLML": "www.hpe.com/support",
                "WWWPASS": "www.hpe.com/support/hpesc",
                "WWWPRV": "www.hpe.com/info/privacy",
                "WWWQSPEC": "www.hpe.com/info/qs",
                "WWWRESTDOC": "www.hpe.com/support/restfulinterface/docs",
                "WWWSUP": "www.hpe.com/support/ilo5",
                "WWWSWLIC": "www.hpe.com/software/SWLicensing"
            },
            "Sessions": {
                "CertCommonName": "bmc.fqdn",
                "CertificateLoginEnabled": false,
                "KerberosEnabled": false,
                "LDAPAuthLicenced": true,
                "LDAPEnabled": false,
                "LocalLoginEnabled": true,
                "LoginFailureDelay": 0,
                "LoginHint": {
                    "Hint": "POST to /Sessions to login using the following JSON object:",
                    "HintPOSTData": {
                        "Password": "password",
                        "UserName": "username"
                    }
                },
                "SecurityOverride": false,
                "ServerName": "bmc-serverName"
            },
            "System": [{
                "Status": {
                    "Health": "Warning"
                }
            }],
            "Time": "2022-07-08T09:07:17Z"
        }
    },
    "Product": "ProLiant DL380 Gen10",
    "ProtocolFeaturesSupported": {
        "ExpandQuery": {
            "ExpandAll": false,
            "Levels": true,
            "Links": false,
            "MaxLevels": 1,
            "NoLinks": true
        },
        "FilterQuery": true,
        "OnlyMemberQuery": true,
        "SelectQuery": false
    },
    "RedfishVersion": "1.6.0",
    "Registries": {
        "@odata.id": "/redfish/v1/Registries/"
    },
    "SessionService": {
        "@odata.id": "/redfish/v1/SessionService/"
    },
    "Systems": {
        "@odata.id": "/redfish/v1/Systems/"
    },
    "Tasks": {
        "@odata.id": "/redfish/v1/TaskService/"
    },
    "TelemetryService": {
        "@odata.id": "/redfish/v1/TelemetryService/"
    },
    "UUID": "00000000-0000-0000-0000-000000000000",
    "UpdateService": {
        "@odata.id": "/redfish/v1/UpdateService/"
    },
    "Vendor": "HPE"
}`

const eventServiceBody = `{
    "@odata.context": "/redfish/v1/$metadata#EventService.EventService",
    "@odata.id": "/redfish/v1/EventService",
    "@odata.type": "#EventService.v1_7_0.EventService",
    "Actions": {
        "#EventService.SubmitTestEvent": {
            "EventType@Redfish.AllowableValues": ["Alert"],
            "target": "/redfish/v1/EventService/Actions/EventService.SubmitTestEvent"
        }
    },
    "DeliveryRetryAttempts": 3,
    "DeliveryRetryIntervalSeconds": 5,
    "Description": "Event Service represents the properties for the service",
    "EventFormatTypes": ["Event", "MetricReport"],
    "EventFormatTypes@odata.count": 2,
    "EventTypesForSubscription": ["Alert", "MetricReport", "Other"],
    "EventTypesForSubscription@odata.count": 3,
    "Id": "EventService",
    "Name": "Event Service",
    "SMTP": {
        "Authentication": "None",
        "ConnectionProtocol": "StartTLS",
        "FromAddress": "",
        "Password": null,
        "Port": 25,
        "ServerAddress": "0.0.0.0",
        "Username": ""
    },
    "SSEFilterPropertiesSupported": {
        "EventFormatType": true,
        "EventType": true,
        "MessageId": true,
        "MetricReportDefinition": true,
        "OriginResource": true,
        "RegistryPrefix": true,
        "ResourceType": true,
        "SubordinateResources": false
    },
    "ServerSentEventUri": "/redfish/v1/SSE",
    "ServiceEnabled": true,
    "Status": {
        "Health": "OK",
        "HealthRollup": "OK",
        "State": "Enabled"
    },
    "Subscriptions": {
        "@odata.id": "/redfish/v1/EventService/Subscriptions"
    }
}`

// TestEventService tests the parsing of EventService objects.
func TestEventService(t *testing.T) {
	var result EventService
	err := json.NewDecoder(strings.NewReader(eventServiceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "EventService" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Event Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.DeliveryRetryAttempts != 3 {
		t.Errorf("Expected 3 retry attempts, got: %d", result.DeliveryRetryAttempts)
	}

	if result.DeliveryRetryIntervalSeconds != 5 {
		t.Errorf("Expected 5 second retry interval, got: %d", result.DeliveryRetryIntervalSeconds)
	}

	if !result.SSEFilterPropertiesSupported.MetricReportDefinition {
		t.Error("MetricReportDefinition filter should be true")
	}

	if !result.SSEFilterPropertiesSupported.MessageID {
		t.Error("Message ID filter should be true")
	}

	if result.SubmitTestEventTarget != "/redfish/v1/EventService/Actions/EventService.SubmitTestEvent" {
		t.Errorf("Invalid SubmitTestEvent target: %s", result.SubmitTestEventTarget)
	}
}

// TestHpeSubmitTestEvent tests SubmitTestEvent for Hpe using a mock server.
func TestHpeSubmitTestEvent(t *testing.T) {
	const redfishBaseURL = "/redfish/v1/"
	var (
		c              common.Client
		err            error
		requestCounter int // this counter is used to verify that the received requests are in the expected order
	)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet &&
			req.URL.String() == redfishBaseURL &&
			requestCounter < 2 { // ServiceRoot
			contentType := req.Header.Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("gofish connect sent wrong header. Content-Type:"+
					" is %v and not expected application/json", contentType)
			}

			requestCounter++

			// Send response to be tested
			rw.WriteHeader(http.StatusOK)
			rw.Header().Set("Content-Type", "application/json")

			rw.Write([]byte(serviceRootBody)) // nolint:errcheck
		} else if req.Method == http.MethodGet && // Get event service
			req.URL.String() == "/redfish/v1/EventService/" &&
			requestCounter == 2 {
			requestCounter++

			rw.Write([]byte(eventServiceBody)) // nolint:errcheck
		} else if req.Method == http.MethodPost && // SubmitTestEvent
			req.URL.String() == "/redfish/v1/EventService/Actions/EventService.SubmitTestEvent" &&
			requestCounter == 3 {
			err := json.NewDecoder(req.Body).Decode(&PayloadType{})
			if err != nil {
				t.Errorf("error in SubmitTestEvent payload for Hpe due to: %v", err)
			}

			requestCounter++

			rw.WriteHeader(http.StatusCreated)
		} else {
			t.Errorf("mock got unexpected %v request to path %v while request counter is %v",
				req.Method, req.URL.String(), requestCounter)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c, err = gofish.Connect(gofish.ClientConfig{Endpoint: server.URL, HTTPClient: server.Client()})

	if err != nil {
		t.Errorf("failed to establish client to mock http server due to: %v", err)
	}
	serviceRoot, err := gofish.ServiceRoot(c)
	if err != nil {
		t.Errorf("failed to get redfish service root due to: %v", err)
	}
	eventService, err := serviceRoot.EventService()
	if err != nil {
		t.Errorf("failed to get event service due to: %v", err)
	}
	hpeEventService, err := FromEventService(eventService)
	if err != nil {
		t.Errorf("failed to get hpe event service due to: %v", err)
	}

	err = hpeEventService.SubmitTestEvent(
		"TestEventId",
		"iLOEvents.2.1.ServerPoweredOff",
		"Alert",
		"Test Event")
	if err != nil {
		t.Errorf("failed to submit test event due to: %v", err)
	}
}
