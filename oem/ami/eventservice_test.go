//
// SPDX-License-Identifier: BSD-3-Clause
//

package ami

import (
	"encoding/json"
	"testing"

	"github.com/stmcginnis/gofish/schemas"
)

var eventServiceBody = `{
  "@odata.context": "/redfish/v1/$metadata#EventService.EventService",
  "@odata.etag": "\"1729105654\"",
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
  "EventFormatTypes": [
    "MetricReport",
    "Event"
  ],
  "Id": "EventService",
  "Name": "Event Service",
  "Oem": {
    "Ami": {
      "@odata.type": "#AMIEventService.v1_0_0.AMIEventService",
      "Certificates": {
        "@odata.id": "/redfish/v1/EventService/Oem/Ami/SMTP/Certificates"
      },
      "SecondarySMTP": {
        "Authentication": "None",
        "ConnectionProtocol": "None",
        "FromAddress": null,
        "Password": null,
        "Port": 25,
        "ServerAddress": null,
        "ServiceEnabled": false,
        "Username": null
      }
    }
  },
  "RegistryPrefixes": [
    "Base",
    "SyncAgent",
    "HttpStatus",
    "EventLog",
    "Security",
    "Task",
    "IPMI"
  ],
  "ResourceTypes": [
    "Systems",
    "Managers",
    "AccountService",
    "Chassis",
    "TelemetryService",
    "EventService",
    "TaskService"
  ],
  "SMTP": {
    "Authentication": "None",
    "ConnectionProtocol": "None",
    "FromAddress": null,
    "Password": null,
    "Port": 25,
    "ServerAddress": null,
    "ServiceEnabled": false,
    "Username": null
  },
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
    "State": "Enabled"
  },
  "SubordinateResourcesSupported": false,
  "Subscriptions": {
    "@odata.id": "/redfish/v1/EventService/Subscriptions"
  }
}`

// TestAMIEventService tests the parsing of the EventService.
func TestAMIEventService(t *testing.T) {
	es := &schemas.EventService{}
	if err := json.Unmarshal([]byte(eventServiceBody), es); err != nil {
		t.Fatalf("error decoding json: %v", err)
	}

	eventService, err := FromEventService(es)
	if err != nil {
		t.Fatalf("error getting oem info: %v", err)
	}

	if eventService.ID != "EventService" {
		t.Errorf("unexpected ID: %s", eventService.ID)
	}

	if eventService.certificates != "/redfish/v1/EventService/Oem/Ami/SMTP/Certificates" {
		t.Errorf("unexpected certificates link: %s", eventService.certificates)
	}

	if *eventService.SecondarySMTP.Port != 25 {
		t.Errorf("unexpected port: %d", eventService.SecondarySMTP.Port)
	}
}
