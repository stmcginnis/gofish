//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
)

var serviceBody = `{
	"@odata.type": "#CertificateService.v1_0_1.CertificateService",
	"@odata.id": "/redfish/v1/CertificateService",
	"Id": "CertificateService",
	"Name": "Certificate Service",
	"CertificateLocations": {
	  "@odata.id": "/redfish/v1/CertificateService/CertificateLocations"
	},
	"Actions": {
	  "Oem": {},
	  "#CertificateService.GenerateCSR": {
		"target": "/redfish/v1/CertificateService/Actions/CertificateService.GenerateCSR",
		"@Redfish.ActionInfo": "/redfish/v1/CertificateService/GenerateCSRActionInfo"
	  },
	  "#CertificateService.ReplaceCertificate": {
		"target": "/redfish/v1/CertificateService/Actions/CertificateService.ReplaceCertificate",
		"@Redfish.ActionInfo": "/redfish/v1/CertificateService/ReplaceCertificateActionInfo"
	  }
	}
  }`

// TestCertificateService tests the parsing of CertificateService objects.
func TestCertificateService(t *testing.T) {
	var result CertificateService
	err := json.NewDecoder(strings.NewReader(serviceBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "CertificateService" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Certificate Service" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.certificateLocations.uri != "/redfish/v1/CertificateService/CertificateLocations" {
		t.Errorf("Received invalid certificate locations: %s",
			result.certificateLocations.uri)
	}
}

// TestCertificateServiceGenerateCSR tests that the GenerateCSR action sends the target
// certificate collection as a link object.
func TestCertificateServiceGenerateCSR(t *testing.T) {
	var result CertificateService
	err := json.NewDecoder(strings.NewReader(serviceBody)).Decode(&result)
	if err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodPost: {
				&http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(strings.NewReader(`{"CSRString": "-----BEGIN CERTIFICATE REQUEST-----\nMIIC\n-----END CERTIFICATE REQUEST-----", "CertificateCollection": {"@odata.id": "/redfish/v1/Managers/1/NetworkProtocol/HTTPS/Certificates"}}`)),
				},
			},
		},
	}
	result.SetClient(testClient)

	response, err := result.GenerateCSR(&CertificateServiceGenerateCSRParameters{
		CertificateCollection: "/redfish/v1/Managers/1/NetworkProtocol/HTTPS/Certificates",
		CommonName:            "bmc.example.com",
		Country:               "DE",
	})
	if err != nil {
		t.Fatalf("Error making GenerateCSR call: %s", err)
	}
	if response.CSRString == "" {
		t.Error("Expected a CSRString in the response")
	}
	if response.certificateCollection != "/redfish/v1/Managers/1/NetworkProtocol/HTTPS/Certificates" {
		t.Errorf("Unexpected response certificate collection: %s", response.certificateCollection)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 1 || calls[0].URL != "/redfish/v1/CertificateService/Actions/CertificateService.GenerateCSR" {
		t.Fatalf("Unexpected calls: %+v", calls)
	}
	if !strings.Contains(calls[0].Payload, "CertificateCollection:map[@odata.id:/redfish/v1/Managers/1/NetworkProtocol/HTTPS/Certificates]") {
		t.Errorf("Expected CertificateCollection to be sent as a link object, got payload: %s", calls[0].Payload)
	}
	if !strings.Contains(calls[0].Payload, "CommonName:bmc.example.com") {
		t.Errorf("Unexpected payload: %s", calls[0].Payload)
	}
}
