//
// SPDX-License-Identifier: BSD-3-Clause
//

package oem

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/redfish"
)

const serviceRootBodyDell = `{
    "@odata.context": "/redfish/v1/$metadata#ServiceRoot.ServiceRoot",
    "@odata.id": "/redfish/v1",
    "@odata.type": "#ServiceRoot.v1_8_0.ServiceRoot",
    "AccountService": {
        "@odata.id": "/redfish/v1/AccountService"
    },
    "CertificateService": {
        "@odata.id": "/redfish/v1/CertificateService"
    },
    "Chassis": {
        "@odata.id": "/redfish/v1/Chassis"
    },
    "Description": "Root Service",
    "EventService": {
        "@odata.id": "/redfish/v1/EventService"
    },
    "Fabrics": {
        "@odata.id": "/redfish/v1/Fabrics"
    },
    "Id": "RootService",
    "JobService": {
        "@odata.id": "/redfish/v1/JobService"
    },
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
        "Dell": {
            "@odata.context": "/redfish/v1/$metadata#DellServiceRoot.DellServiceRoot",
            "@odata.type": "#DellServiceRoot.v1_0_0.DellServiceRoot",
            "IsBranded": 0,
            "ManagerMACAddress": "00:00:00:00:00:00",
            "ServiceTag": "0000000"
        }
    },
    "Product": "Integrated Dell Remote Access Controller",
    "ProtocolFeaturesSupported": {
        "DeepOperations": {
            "DeepPATCH": false,
            "DeepPOST": false
        },
        "ExcerptQuery": false,
        "ExpandQuery": {
            "ExpandAll": true,
            "Levels": true,
            "Links": true,
            "MaxLevels": 1,
            "NoLinks": true
        },
        "FilterQuery": true,
        "OnlyMemberQuery": true,
        "SelectQuery": true
    },
    "RedfishVersion": "1.11.0",
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
    "UpdateService": {
        "@odata.id": "/redfish/v1/UpdateService"
    },
    "Vendor": "Dell"
}`

const eventServiceBodyDell = `{
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

func DellEventServiceMock(t *testing.T) (server *httptest.Server) {
	const redfishBaseURL = "/redfish/v1/"
	var requestCounter int // this counter is used to verify that the received requests are in the expected order

	// Start a local HTTP server
	server = httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		switch {
		// ServiceRoot
		case req.Method == http.MethodGet &&
			req.URL.String() == redfishBaseURL &&
			requestCounter < 4:
			contentType := req.Header.Get("Content-Type")
			if contentType != "application/json" {
				t.Errorf("gofish connect sent wrong header. Content-Type:"+
					" is %v and not expected application/json", contentType)
			}
			// Send response to be tested
			rw.WriteHeader(http.StatusOK)
			rw.Header().Set("Content-Type", "application/json")

			rw.Write([]byte(serviceRootBodyDell)) //nolint:errcheck

		// Get event service
		case req.Method == http.MethodGet &&
			req.URL.String() == "/redfish/v1/EventService" &&
			requestCounter == 2:
			rw.Write([]byte(eventServiceBodyDell)) //nolint:errcheck

		default:
			t.Errorf("mock got unexpected %v request to path %v while request counter is %v",
				req.Method, req.URL.String(), requestCounter)
		}
		requestCounter++
	}))
	return server
}

var server *httptest.Server

func GetMockedEventService(t *testing.T) *redfish.EventService {
	server = DellEventServiceMock(t) // For Dell

	c, err := gofish.Connect(gofish.ClientConfig{Endpoint: server.URL, HTTPClient: server.Client()})

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
	return eventService
}

func TestGetRedfishVendor(t *testing.T) {
	eventService := GetMockedEventService(t)
	// Close the server when test finishes
	defer server.Close()

	vendor, err := GetRedfishVendor(eventService)
	if err != nil {
		t.Errorf("error received from GetRedfishVendor() due to: %v", err)
	}
	if vendor != Dell {
		t.Errorf("expected to get Dell vendor but got: %v", vendor)
	}
}

func TestGetEventServiceByVendor(t *testing.T) {
	eventService := GetMockedEventService(t)
	// Close the server when test finishes
	defer server.Close()
	_, err := GetEventServiceByVendor(eventService)
	if err != nil {
		t.Errorf("failed to get event service by vendor due to: %v", err)
	}
}
