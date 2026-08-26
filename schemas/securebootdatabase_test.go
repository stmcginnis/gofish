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

var secureBootDatabaseBody = `{
	"@odata.type": "#SecureBootDatabase.v1_0_2.SecureBootDatabase",
	"Id": "PK",
	"Name": "PK - Platform Key",
	"Description": "UEFI PK Secure Boot Database",
	"DatabaseId": "PK",
	"Certificates": {
	  "@odata.id": "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Certificates/"
	},
	"Signatures": {
	  "@odata.id": "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Signatures/"
	},
	"Actions": {
	  "#SecureBootDatabase.ResetKeys": {
		"target": "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Actions/SecureBootDatabase.ResetKeys",
		"ResetKeysType@Redfish.AllowableValues": [
		  "ResetAllKeysToDefault",
		  "DeleteAllKeys"
		]
	  }
	},
	"@odata.id": "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK"
  }`

// TestSecureBootDatabase tests the parsing of SecureBootDatabase objects.
func TestSecureBootDatabase(t *testing.T) {
	var result SecureBootDatabase
	err := json.NewDecoder(strings.NewReader(secureBootDatabaseBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	assertEquals(t, "PK", result.ID)
	assertEquals(t, "PK - Platform Key", result.Name)
	assertEquals(t, "PK", result.DatabaseID)
	assertEquals(t, "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Certificates/", result.certificates)
	assertEquals(t, "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Signatures/", result.signatures)
}

// TestSecureBootDatabaseResetKeys tests the SecureBootDatabase ResetKeys call.
func TestSecureBootDatabaseResetKeys(t *testing.T) {
	var result SecureBootDatabase
	err := json.NewDecoder(strings.NewReader(secureBootDatabaseBody)).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	testClient := &TestClient{}
	result.SetClient(testClient)

	_, err = result.ResetKeys(DeleteAllKeysSecureBootDatabaseResetKeysType)
	if err != nil {
		t.Errorf("Error making Reset call: %s", err)
	}

	calls := testClient.CapturedCalls()

	if len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}

	if !strings.Contains(calls[0].Payload, "ResetKeysType:DeleteAllKeys") {
		t.Errorf("Expected reset type not found in payload: %s", calls[0].Payload)
	}
}

const secureBootCertificateBody = `{
	"@odata.type": "#Certificate.v1_9_0.Certificate",
	"@odata.id": "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Certificates/1",
	"Id": "1",
	"Name": "Platform Key Certificate",
	"CertificateType": "PEM",
	"UefiSignatureOwner": "6e4dfd8f-6b1a-4c66-9c47-9e2e0b6a7d31"
  }`

const secureBootSignatureBody = `{
	"@odata.type": "#Signature.v1_0_2.Signature",
	"@odata.id": "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Signatures/1",
	"Id": "1",
	"Name": "Forbidden Hash",
	"SignatureType": "EFI_CERT_SHA256_GUID",
	"SignatureTypeRegistry": "UEFI"
  }`

const testPEMCertificate = "-----BEGIN CERTIFICATE-----\nMIIB\n-----END CERTIFICATE-----\n"

// newTestSecureBootDatabase decodes the shared test body and wires it to client.
func newTestSecureBootDatabase(t *testing.T, client Client) *SecureBootDatabase {
	t.Helper()

	var result SecureBootDatabase
	if err := json.NewDecoder(strings.NewReader(secureBootDatabaseBody)).Decode(&result); err != nil {
		t.Fatalf("Error decoding JSON: %s", err)
	}
	result.SetClient(client)

	return &result
}

// testResponse builds a canned response for the TestClient.
func testResponse(statusCode int, body string, header http.Header) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     header,
	}
}

// TestSecureBootDatabaseAddCertificate tests enrolling a certificate when the
// service returns the created resource in the response body.
func TestSecureBootDatabaseAddCertificate(t *testing.T) {
	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodPost: {testResponse(http.StatusCreated, secureBootCertificateBody, http.Header{})},
		},
	}
	result := newTestSecureBootDatabase(t, testClient)

	certificate, err := result.AddCertificate(testPEMCertificate, PEMCertificateType, "")
	if err != nil {
		t.Fatalf("Error making AddCertificate call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("Expected one call to be made, captured: %#v", calls)
	}

	assertEquals(t, http.MethodPost, calls[0].Action)
	assertEquals(t, "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Certificates/", calls[0].URL)
	assertContains(t, "CertificateType:PEM", calls[0].Payload)
	assertContains(t, "BEGIN CERTIFICATE", calls[0].Payload)
	assertNotContain(t, "UefiSignatureOwner", calls[0].Payload)

	if certificate == nil {
		t.Fatal("Expected the created certificate to be returned")
	}

	assertEquals(t, "1", certificate.ID)
	assertEquals(t, "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Certificates/1", certificate.ODataID)

	if certificate.GetClient() != testClient {
		t.Error("Expected the client to be set on the returned certificate")
	}
}

// TestSecureBootDatabaseAddCertificateUefiSignatureOwner tests that the
// optional UEFI signature owner is sent when provided.
func TestSecureBootDatabaseAddCertificateUefiSignatureOwner(t *testing.T) {
	testClient := &TestClient{}
	result := newTestSecureBootDatabase(t, testClient)

	if _, err := result.AddCertificate(testPEMCertificate, PEMCertificateType, "77fa9abd-0359-4d32-bd60-28f4e78f784b"); err != nil {
		t.Fatalf("Error making AddCertificate call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("Expected one call to be made, captured: %#v", calls)
	}

	assertContains(t, "UefiSignatureOwner:77fa9abd-0359-4d32-bd60-28f4e78f784b", calls[0].Payload)
}

// TestSecureBootDatabaseAddCertificateLocation tests enrolling a certificate
// when the service returns an empty body and points at the new resource with
// the Location header.
func TestSecureBootDatabaseAddCertificateLocation(t *testing.T) {
	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodPost: {testResponse(http.StatusCreated, "", http.Header{
				"Location": []string{"https://192.168.1.1/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Certificates/1"},
			})},
			http.MethodGet: {testResponse(http.StatusOK, secureBootCertificateBody, http.Header{})},
		},
	}
	result := newTestSecureBootDatabase(t, testClient)

	certificate, err := result.AddCertificate(testPEMCertificate, PEMCertificateType, "")
	if err != nil {
		t.Fatalf("Error making AddCertificate call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 2 {
		t.Fatalf("Expected two calls to be made, captured: %#v", calls)
	}

	// The absolute URL from the header is reduced to a path before it is used.
	assertEquals(t, http.MethodGet, calls[1].Action)
	assertEquals(t, "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Certificates/1", calls[1].URL)

	if certificate == nil {
		t.Fatal("Expected the created certificate to be returned")
	}

	assertEquals(t, "1", certificate.ID)
}

// TestSecureBootDatabaseAddCertificateNoResource tests that a service which
// identifies neither the created resource nor its location is not an error.
func TestSecureBootDatabaseAddCertificateNoResource(t *testing.T) {
	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodPost: {testResponse(http.StatusNoContent, "", http.Header{})},
		},
	}
	result := newTestSecureBootDatabase(t, testClient)

	certificate, err := result.AddCertificate(testPEMCertificate, PEMCertificateType, "")
	if err != nil {
		t.Fatalf("Error making AddCertificate call: %s", err)
	}

	if certificate != nil {
		t.Errorf("Expected no certificate to be returned, got: %#v", certificate)
	}

	if calls := testClient.CapturedCalls(); len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}
}

// TestSecureBootDatabaseAddCertificateAccepted tests that an asynchronously
// handled enrollment is not an error.
func TestSecureBootDatabaseAddCertificateAccepted(t *testing.T) {
	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodPost: {testResponse(http.StatusAccepted, "", http.Header{
				"Location": []string{"/redfish/v1/TaskService/Tasks/1"},
			})},
		},
	}
	result := newTestSecureBootDatabase(t, testClient)

	certificate, err := result.AddCertificate(testPEMCertificate, PEMCertificateType, "")
	if err != nil {
		t.Fatalf("Error making AddCertificate call: %s", err)
	}

	if certificate != nil {
		t.Errorf("Expected no certificate to be returned, got: %#v", certificate)
	}

	if calls := testClient.CapturedCalls(); len(calls) != 1 {
		t.Errorf("Expected one call to be made, captured: %#v", calls)
	}
}

// TestSecureBootDatabaseAddCertificateUnsupported tests that a database without
// a certificate collection reports an error instead of calling the service.
func TestSecureBootDatabaseAddCertificateUnsupported(t *testing.T) {
	testClient := &TestClient{}
	result := newTestSecureBootDatabase(t, testClient)
	result.certificates = ""

	if _, err := result.AddCertificate(testPEMCertificate, PEMCertificateType, ""); err == nil {
		t.Error("Expected an error for a database without a certificate collection")
	}

	if calls := testClient.CapturedCalls(); len(calls) != 0 {
		t.Errorf("Expected no calls to be made, captured: %#v", calls)
	}
}

// TestSecureBootDatabaseAddSignature tests enrolling a signature.
func TestSecureBootDatabaseAddSignature(t *testing.T) {
	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodPost: {testResponse(http.StatusCreated, secureBootSignatureBody, http.Header{})},
		},
	}
	result := newTestSecureBootDatabase(t, testClient)

	signature, err := result.AddSignature(
		"a4a41f1e0a3c4c9b8f2f0f4f9e0d4a1b", UEFISignatureTypeRegistry, "EFI_CERT_SHA256_GUID", "")
	if err != nil {
		t.Fatalf("Error making AddSignature call: %s", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("Expected one call to be made, captured: %#v", calls)
	}

	assertEquals(t, http.MethodPost, calls[0].Action)
	assertEquals(t, "/redfish/v1/Systems/1/SecureBoot/SecureBootDatabases/PK/Signatures/", calls[0].URL)
	assertContains(t, "SignatureTypeRegistry:UEFI", calls[0].Payload)
	assertContains(t, "SignatureType:EFI_CERT_SHA256_GUID", calls[0].Payload)
	assertContains(t, "SignatureString:a4a41f1e0a3c4c9b8f2f0f4f9e0d4a1b", calls[0].Payload)

	if signature == nil {
		t.Fatal("Expected the created signature to be returned")
	}

	assertEquals(t, "1", signature.ID)
	assertEquals(t, "EFI_CERT_SHA256_GUID", signature.SignatureType)
}

// TestSecureBootDatabaseAddSignatureUnsupported tests that a database without a
// signature collection reports an error instead of calling the service.
func TestSecureBootDatabaseAddSignatureUnsupported(t *testing.T) {
	testClient := &TestClient{}
	result := newTestSecureBootDatabase(t, testClient)
	result.signatures = ""

	if _, err := result.AddSignature("abcd", UEFISignatureTypeRegistry, "EFI_CERT_SHA256_GUID", ""); err == nil {
		t.Error("Expected an error for a database without a signature collection")
	}

	if calls := testClient.CapturedCalls(); len(calls) != 0 {
		t.Errorf("Expected no calls to be made, captured: %#v", calls)
	}
}
