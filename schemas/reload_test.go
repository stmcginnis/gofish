//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"errors"
	"io"
	"net/http"
	"strings"
	"testing"
)

const reloadThermalURI = "/redfish/v1/Chassis/1/Thermal"

// notModifiedCall returns a 304 Not Modified response carrying an Etag header,
// as a service would answer a conditional GET whose If-None-Match still matches.
func notModifiedCall(etag string) *http.Response {
	header := make(http.Header)
	header.Set("Etag", etag)
	return &http.Response{
		Status:     "304 Not Modified",
		StatusCode: http.StatusNotModified,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Body:       io.NopCloser(strings.NewReader("")),
		Header:     header,
	}
}

// okCall returns a 200 response with the given body and Etag header.
func okCall(body, etag string) *http.Response {
	header := make(http.Header)
	if etag != "" {
		header.Set("Etag", etag)
	}
	return &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          io.NopCloser(strings.NewReader(body)),
		ContentLength: int64(len(body)),
		Header:        header,
	}
}

// TestRefreshNotModified verifies Refresh sends the object's ETag as
// If-None-Match and, on a 304, returns the original object unchanged along with
// ErrNotModified.
func TestRefreshNotModified(t *testing.T) {
	const etag = `"v1"`
	obj := &Entity{ODataID: reloadThermalURI, Name: "Thermal", ODataEtag: etag}

	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {notModifiedCall(etag)},
		},
	}
	obj.SetClient(testClient)

	got, err := Refresh(obj)
	if !errors.Is(err, ErrNotModified) {
		t.Fatalf("expected ErrNotModified, got %v", err)
	}
	if got != obj {
		t.Errorf("expected the original object back on 304, got a different value")
	}
	if got.Name != "Thermal" || got.GetETag() != etag {
		t.Errorf("object should be unchanged on 304, got Name=%q ETag=%q", got.Name, got.GetETag())
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(calls))
	}
	if calls[0].URL != reloadThermalURI {
		t.Errorf("expected GET %q, got %q", reloadThermalURI, calls[0].URL)
	}
	if inm := calls[0].CustomHeaders["If-None-Match"]; inm != etag {
		t.Errorf("expected If-None-Match %q, got %q", etag, inm)
	}
}

// TestRefreshModified verifies Refresh returns the updated resource when the
// service reports a change (200 with a fresh body and ETag).
func TestRefreshModified(t *testing.T) {
	obj := &Entity{ODataID: reloadThermalURI, Name: "Thermal", ODataEtag: `"v1"`}

	const updated = `{"@odata.id":"` + reloadThermalURI + `","Id":"Thermal","Name":"Thermal Updated"}`
	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {okCall(updated, `"v2"`)},
		},
	}
	obj.SetClient(testClient)

	got, err := Refresh(obj)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.Name != "Thermal Updated" {
		t.Errorf("expected updated Name, got %q", got.Name)
	}
	if got.GetETag() != `"v2"` {
		t.Errorf("expected refreshed ETag \"v2\", got %q", got.GetETag())
	}
	if got.GetClient() == nil {
		t.Error("refreshed object should have its client set")
	}
}

// TestRefreshNoETag verifies that without an ETag Refresh falls back to an
// unconditional reload (no If-None-Match header).
func TestRefreshNoETag(t *testing.T) {
	obj := &Entity{ODataID: reloadThermalURI, Name: "Thermal"}

	const body = `{"@odata.id":"` + reloadThermalURI + `","Id":"Thermal","Name":"Thermal"}`
	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {okCall(body, "")},
		},
	}
	obj.SetClient(testClient)

	if _, err := Refresh(obj); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	calls := testClient.CapturedCalls()
	if len(calls) != 1 {
		t.Fatalf("expected 1 call, got %d", len(calls))
	}
	if _, ok := calls[0].CustomHeaders["If-None-Match"]; ok {
		t.Errorf("expected no If-None-Match header, got %q", calls[0].CustomHeaders["If-None-Match"])
	}
}

// TestReloadWithHeaders verifies Reload forwards caller-supplied headers verbatim.
func TestReloadWithHeaders(t *testing.T) {
	obj := &Entity{ODataID: reloadThermalURI, Name: "Thermal", ODataEtag: `"v1"`}

	testClient := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {notModifiedCall(`"v1"`)},
		},
	}
	obj.SetClient(testClient)

	headers := map[string]string{"If-None-Match": "unquoted-v1"}
	_, err := Reload[Entity](obj, headers)
	if !errors.Is(err, ErrNotModified) {
		t.Fatalf("expected ErrNotModified, got %v", err)
	}

	calls := testClient.CapturedCalls()
	if inm := calls[0].CustomHeaders["If-None-Match"]; inm != "unquoted-v1" {
		t.Errorf("expected caller header forwarded verbatim, got %q", inm)
	}
}

// TestReloadNoClient and TestReloadNoODataID cover the guard clauses.
func TestReloadGuards(t *testing.T) {
	noClient := &Entity{ODataID: reloadThermalURI}
	if _, err := Reload[Entity](noClient, nil); err == nil {
		t.Error("expected error when no client is set")
	}

	noURI := &Entity{}
	noURI.SetClient(&TestClient{})
	if _, err := Reload[Entity](noURI, nil); err == nil {
		t.Error("expected error when object has no @odata.id")
	}
}
