//
// SPDX-License-Identifier: BSD-3-Clause
//

package zt

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

const serviceRootBody = `{
"@odata.context": "/redfish/v1/$metadata#ServiceRoot.ServiceRoot",
"@odata.etag": "\"0000000000\"",
"@odata.id": "/redfish/v1/",
"@odata.type": "#ServiceRoot.v1_5_2.ServiceRoot",
"AccountService": {
"@odata.id": "/redfish/v1/AccountService"
},
"CertificateService": {
"@odata.id": "/redfish/v1/CertificateService"
},
"Chassis": {
"@odata.id": "/redfish/v1/Chassis"
},
"CompositionService": {
"@odata.id": "/redfish/v1/CompositionService"
},
"Description": "The service root for all Redfish requests on this host",
"EventService": {
"@odata.id": "/redfish/v1/EventService"
},
"Id": "RootService",
"JsonSchemas": {
"@odata.id": "/redfish/v1/JsonSchemas"
},
"Links": {
"Sessions": {
"@odata.id": "/redfish/v1/SessionService/Sessions"
}
},
"Managers": {
"@odata.id": "/redfish/v1/Managers"
},
"Name": "Root Service",
"Oem": {
"Ami": {
"@odata.type": "#AMIServiceRoot.v1_0_0.AMIServiceRoot",
"InventoryDataStatus": {
"@odata.id": "/redfish/v1/Oem/Ami/InventoryData/Status"
},
"RtpVersion": "1.8.a",
"configurations": {
"@odata.id": "/redfish/v1/configurations"
}
}
},
"Product": "AMI Redfish Server",
"ProtocolFeaturesSupported": {
"ExcerptQuery": true,
"ExpandQuery": {
"ExpandAll": true,
"Levels": true,
"Links": true,
"MaxLevels": 5,
"NoLinks": true
},
"FilterQuery": true,
"OnlyMemberQuery": true,
"SelectQuery": true
},
"RedfishVersion": "1.8.0",
"Registries": {
"@odata.id": "/redfish/v1/Registries"
},
"SessionService": {
"@odata.id": "/redfish/v1/SessionService"
},
"Systems": {
"@odata.id": "/redfish/v1/Systems"
},
"Tasks": {
"@odata.id": "/redfish/v1/TaskService"
},
"TelemetryService": {
"@odata.id": "/redfish/v1/TelemetryService"
},
"UUID": "00000000-0000-0000-0000-000000000000",
"UpdateService": {
"@odata.id": "/redfish/v1/UpdateService"
},
"Vendor": "AMI"
}`

const eventServiceBody = `{
    "@odata.context": "/redfish/v1/$metadata#EventService.EventService",
    "@odata.etag": "\"0000000000\"",
    "@odata.id": "/redfish/v1/EventService",
    "@odata.type": "#EventService.v1_5_0.EventService",
    "Actions": {
        "#EventService.SubmitTestEvent": {
            "@Redfish.ActionInfo": "/redfish/v1/EventService/SubmitTestEventActionInfo",
            "target": "/redfish/v1/EventService/Actions/EventService.SubmitTestEvent"
        }
    },
    "DeliveryRetryAttempts": 3,
    "DeliveryRetryIntervalSeconds": 60,
    "Description": "Event Service",
    "EventFormatTypes": ["MetricReport", "Event"],
    "Id": "EventService",
    "Name": "Event Service",
    "RegistryPrefixes": ["SyncAgent", "Security", "IPMI", "EventLog", "Task", "HttpStatus", "Base"],
    "ResourceTypes": ["Systems", "Managers", "EventService", "TelemetryService", "AccountService", "TaskService", "Chassis"],
    "SSEFilterPropertiesSupported": {
        "EventFormatType": true,
        "MessageId": true,
        "MetricReportDefinition": false,
        "OriginResource": true,
        "RegistryPrefix": true,
        "ResourceType": true,
        "SubordinateResources": false
    },
    "ServerSentEventUri": "/redfish/v1/EventService/SSE",
    "ServiceEnabled": true,
    "Status": {
        "Health": "OK",
        "State": "Enabled"
    },
    "SubordinateResourcesSupported": false,
    "Subscriptions": {
        "@odata.id": "/redfish/v1/EventService/Subscriptions"
    }
}`

const subscribeResponseBody = `{
    "@odata.context": "/redfish/v1/$metadata#EventDestination.EventDestination",
    "@odata.etag": "\"0000000000\"",
    "@odata.id": "/redfish/v1/EventService/Subscriptions",
    "@odata.type": "#EventDestination.v1_6_0.EventDestination",
    "Context": "root",
    "DeliveryRetryPolicy": "TerminateAfterRetries",
    "Description": "Event Subscription",
    "Destination": "https://events.receiver/events/",
    "EventFormatType": "Event",
    "Id": 1,
    "Name": "Subscription 1",
    "Protocol": "Redfish",
    "Status": {
        "Health": "OK",
        "HealthRollup": "OK",
        "State": "Enabled"
    },
    "SubordinateResources": false,
    "SubscriptionType": "RedfishEvent"
}`

func TestSubscribeZT(t *testing.T) {
	const redfishBaseURL = "/redfish/v1/"

	var (
		c              common.Client
		requestCounter int // this counter is used to verify that the received requests, are in the expected order
		err            error
	)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet &&
			req.URL.String() == redfishBaseURL &&
			requestCounter < 2 { // ServiceRoot
			log.Printf("Mock received login request")
			contentType := req.Header.Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("gofish connect sent wrong header. Content-Type:"+
					" is %v and not expected application/json", contentType)
			}

			requestCounter++
			// Send response to be tested
			rw.WriteHeader(http.StatusOK)
			rw.Header().Set("Content-Type", "application/json")

			rw.Write([]byte(serviceRootBody)) //nolint:errcheck
		} else if req.Method == http.MethodGet && // Get event service
			req.URL.String() == "/redfish/v1/EventService" &&
			requestCounter == 2 {
			log.Printf("Getting event service")

			requestCounter++

			rw.Write([]byte(eventServiceBody)) //nolint:errcheck
		} else if req.Method == http.MethodPost && // Subscribe
			req.URL.String() == "/redfish/v1/EventService/Subscriptions" &&
			requestCounter == 3 {
			log.Printf("Mock got suscription POST")

			requestCounter++

			rw.Write([]byte(subscribeResponseBody)) //nolint:errcheck
		} else {
			t.Errorf("mock got unexpected %v request to path %v while request counter is %v",
				req.Method, req.URL.String(), requestCounter)
		}
	}))
	// Close the server when test finishes
	defer server.Close()

	c, err = gofish.Connect(gofish.ClientConfig{Endpoint: server.URL, HTTPClient: server.Client()})
	if err != nil {
		t.Errorf("failed to establish redfish session due to: %v", err)
	}

	serviceRoot, err := gofish.ServiceRoot(c)
	if err != nil {
		t.Errorf("failed to get redfish service root due to: %v", err)
	}
	eventService, err := serviceRoot.EventService()
	if err != nil {
		t.Errorf("failed to get event service due to: %v", err)
	}
	ztEventService, err := FromEventService(eventService)
	if err != nil {
		t.Errorf("failed to get zt systems event service due to: %v", err)
	}
	url, err := ztEventService.Subscribe(
		"https://events.receiver/events/",
		redfish.RedfishEventDestinationProtocol)
	if err != nil {
		t.Errorf("failed to Subscribe() due to: %v", err)
	}

	expectedSubURL := "/redfish/v1/EventService/Subscriptions/1"
	if url != expectedSubURL {
		t.Errorf("expected subscription url: %v got %v", expectedSubURL, url)
	}
}
