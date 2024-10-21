//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/common"
)

const serviceRootBody = `{
  "@odata.type": "#ServiceRoot.v1_9_0.ServiceRoot",
  "@odata.id": "/redfish/v1",
  "Id": "ServiceRoot",
  "Name": "Root Service",
  "RedfishVersion": "1.11.0",
  "UUID": "00000000-0000-0000-0000-3CECEFE32D23",
  "Systems": {
    "@odata.id": "/redfish/v1/Systems"
  },
  "Chassis": {
    "@odata.id": "/redfish/v1/Chassis"
  },
  "Managers": {
    "@odata.id": "/redfish/v1/Managers"
  },
  "Tasks": {
    "@odata.id": "/redfish/v1/TaskService"
  },
  "SessionService": {
    "@odata.id": "/redfish/v1/SessionService"
  },
  "AccountService": {
    "@odata.id": "/redfish/v1/AccountService"
  },
  "EventService": {
    "@odata.id": "/redfish/v1/EventService"
  },
  "UpdateService": {
    "@odata.id": "/redfish/v1/UpdateService"
  },
  "CertificateService": {
    "@odata.id": "/redfish/v1/CertificateService"
  },
  "Registries": {
    "@odata.id": "/redfish/v1/Registries"
  },
  "JsonSchemas": {
    "@odata.id": "/redfish/v1/JsonSchemas"
  },
  "TelemetryService": {
    "@odata.id": "/redfish/v1/TelemetryService"
  },
  "Product": null,
  "Links": {
    "Sessions": {
      "@odata.id": "/redfish/v1/SessionService/Sessions"
    }
  },
  "Oem": {
    "Supermicro": {
      "DumpService": {
        "@odata.id": "/redfish/v1/Oem/Supermicro/DumpService"
      }
    }
  },
  "ProtocolFeaturesSupported": {
    "FilterQuery": true,
    "SelectQuery": true,
    "ExcerptQuery": false,
    "OnlyMemberQuery": false,
    "DeepOperations": {
      "DeepPATCH": false,
      "DeepPOST": false,
      "MaxLevels": 1
    },
    "ExpandQuery": {
      "Links": true,
      "NoLinks": true,
      "ExpandAll": true,
      "Levels": true,
      "MaxLevels": 2
    }
  },
  "@odata.etag": "\"1a10733cff76c5506e6903b25ab88e55\""
}`

var updateServiceBody = `{
  "@odata.type": "#UpdateService.v1_8_4.UpdateService",
  "@odata.id": "/redfish/v1/UpdateService",
  "Id": "UpdateService",
  "Name": "Update Service",
  "Description": "Service for updating firmware and includes inventory of firmware",
  "Status": {
    "State": "Enabled",
    "Health": "OK",
    "HealthRollup": "OK"
  },
  "ServiceEnabled": true,
  "MultipartHttpPushUri": "/redfish/v1/UpdateService/upload",
  "FirmwareInventory": {
    "@odata.id": "/redfish/v1/UpdateService/FirmwareInventory"
  },
  "Actions": {
    "Oem": {
      "#SmcUpdateService.Install": {
        "target": "/redfish/v1/UpdateService/Actions/Oem/SmcUpdateService.Install",
        "@Redfish.ActionInfo": "/redfish/v1/UpdateService/Oem/Supermicro/InstallActionInfo"
      }
    },
    "#UpdateService.SimpleUpdate": {
      "target": "/redfish/v1/UpdateService/Actions/UpdateService.SimpleUpdate",
      "@Redfish.ActionInfo": "/redfish/v1/UpdateService/SimpleUpdateActionInfo"
    },
    "#UpdateService.StartUpdate": {
      "target": "/redfish/v1/UpdateService/Actions/UpdateService.StartUpdate"
    }
  },
  "Oem": {
    "Supermicro": {
      "@odata.type": "#SmcUpdateServiceExtensions.v1_0_0.UpdateService",
      "SSLCert": {
        "@odata.id": "/redfish/v1/UpdateService/Oem/Supermicro/SSLCert"
      },
      "IPMIConfig": {
        "@odata.id": "/redfish/v1/UpdateService/Oem/Supermicro/IPMIConfig"
      }
    }
  },
  "@odata.etag": "\"e9b94401dae9992fef2e71ef30cbcfdc\""
}`

const smcSSLCertBody = `{
  "@odata.type": "#SSLCert.v1_0_0.SSLCert",
  "@odata.id": "/redfish/v1/UpdateService/Oem/Supermicro/SSLCert",
  "Id": "SSLCert",
  "Name": "SSLCert",
  "VaildFrom": "Oct  9 11:15:00 2024 GMT",
  "GoodTHRU": "Oct  9 11:15:00 2025 GMT",
  "Actions": {
    "#SmcSSLCert.Upload": {
      "target": "/redfish/v1/UpdateService/Oem/Supermicro/SSLCert/Actions/SmcSSLCert.Upload",
      "UploadKeyWords@Redfish.AllowableValues": [
        "cert_file",
        "key_file"
      ]
    }
  },
  "@odata.etag": "\"e4be24decdd8b293984fb26e1a78e62a\""
}`

const smcSSLCertUploadResponse = `{
  "Success": {
    "code": "Base.v1_10_3.Success",
    "message": "Successfully Completed Request. See ExtendedInfo for more information.",
    "@Message.ExtendedInfo": [
      {
        "MessageId": "SMC.1.0.OemSslcertUploaded",
        "Severity": "OK",
        "Resolution": "No resolution was required.",
        "Message": "SSL certificate and private key were successfully uploaded.",
        "MessageArgs": [
          ""
        ],
        "RelatedProperties": [
          ""
        ]
      }
    ]
  }
}`

const sslCertFile = `-----BEGIN CERTIFICATE-----
MIIDpDCCAoygAwIBAgIUIf
-----END CERTIFICATE-----`

//nolint:gosec
const sslKeyFile = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEAz
-----END RSA PRIVATE KEY-----`

// TestSmcUpdateService tests the parsing of the UpdateService oem field
func TestSmcUpdateService(t *testing.T) {
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

			rw.Write([]byte(serviceRootBody)) //nolint:errcheck
		} else if req.Method == http.MethodGet && // Get event service
			req.URL.String() == "/redfish/v1/UpdateService" &&
			requestCounter == 2 {
			requestCounter++
			rw.Write([]byte(updateServiceBody)) //nolint:errcheck
		} else if req.Method == http.MethodGet &&
			req.URL.String() == "/redfish/v1/UpdateService/Oem/Supermicro/SSLCert" &&
			requestCounter == 3 {
			requestCounter++
			rw.Write([]byte(smcSSLCertBody)) //nolint:errcheck
		} else if req.Method == http.MethodPost && // SubmitTestEvent
			req.URL.String() == "/redfish/v1/UpdateService/Oem/Supermicro/SSLCert/Actions/SmcSSLCert.Upload" &&
			requestCounter == 4 {
			// TODO: Actually check if the request body is correct
			requestCounter++
			rw.Write([]byte(smcSSLCertUploadResponse)) //nolint:errcheck
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
	origUpdateService, err := serviceRoot.UpdateService()
	if err != nil {
		t.Errorf("failed to get update service due to: %v", err)
	}
	updateService, err := FromUpdateService(origUpdateService)
	if err != nil {
		t.Errorf("error getting OEM object: %v", err)
	}

	if updateService.ID != "UpdateService" {
		t.Errorf("unexpected ID: %s", updateService.ID)
	}

	if updateService.installTarget != "/redfish/v1/UpdateService/Actions/Oem/SmcUpdateService.Install" {
		t.Errorf("unexpected install target: %s", updateService.installTarget)
	}

	if updateService.sslCert != "/redfish/v1/UpdateService/Oem/Supermicro/SSLCert" {
		t.Errorf("unexpected ssl cert link: %s", updateService.installTarget)
	}

	if updateService.ipmiConfig != "/redfish/v1/UpdateService/Oem/Supermicro/IPMIConfig" {
		t.Errorf("unexpected ipmi config link: %s", updateService.installTarget)
	}

	cert, err := updateService.SSLCert()
	if err != nil {
		t.Errorf("Failed to get SSL certificate due to: %v", err)
	}

	err = cert.Upload(strings.NewReader(sslCertFile), strings.NewReader(sslKeyFile))
	if err != nil {
		t.Errorf("Failed to upload SSL certificate due to: %v", err)
	}
}
