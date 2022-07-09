//
// SPDX-License-Identifier: BSD-3-Clause
//

package test

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stmcginnis/gofish"
	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/oem/hpe/events"
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

func TestHpeSubmitTestEvent(t *testing.T) {
	const redfishBaseURL = "/redfish/v1/"
	var (
		c              common.Client
		err            error
		requestCounter int // this counter is used to verify that the received requests, are in the expected order
	)

	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if req.Method == http.MethodGet &&
			req.URL.String() == redfishBaseURL &&
			requestCounter == 0 { // ServiceRoot
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

			rw.Write([]byte(serviceRootBody)) // nolint:errcheck
		} else if req.Method == http.MethodPost && // SubmitTestEvent
			req.URL.String() == events.SubmitTestEventTarget &&
			requestCounter == 1 {
			log.Printf("Mock got SubmitTestEvent POST")

			err := json.NewDecoder(req.Body).Decode(&events.PayloadType{})
			if err != nil {
				t.Errorf("error in SubmitTestEvent payload for Dell due to: %v", err)
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

	err = events.SubmitTestEvent(
		c,
		"TestEventId",
		"iLOEvents.2.1.ServerPoweredOff",
		"Alert",
		"Test Event")

	if err != nil {
		log.Printf("Got error %v", err)
	}
}
