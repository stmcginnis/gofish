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

const testETag = `"v1"`

type etagHeaderCase struct {
	name        string
	header      http.Header
	wantIfMatch string
}

func etagHeaderSet(key, value string) http.Header {
	h := make(http.Header)
	h.Set(key, value)
	return h
}

func etagHeaderRaw(key, value string) http.Header {
	return http.Header{key: []string{value}}
}

// etagHeaderCases covers the ETag lookup behaviors callers rely on: Header.Set
// (canonicalizes any casing), raw map keys that skip canonicalization, a
// missing header, and an empty value.
func etagHeaderCases() []etagHeaderCase {
	return []etagHeaderCase{
		{name: "Header.Set ETag", header: etagHeaderSet("ETag", testETag), wantIfMatch: testETag},
		{name: "Header.Set etag", header: etagHeaderSet("etag", testETag), wantIfMatch: testETag},
		{name: "raw map key Etag", header: etagHeaderRaw("Etag", testETag), wantIfMatch: testETag},
		{name: "raw map key ETag", header: etagHeaderRaw("ETag", testETag), wantIfMatch: testETag},
		{name: "raw map key etag", header: etagHeaderRaw("etag", testETag), wantIfMatch: testETag},
		{name: "missing header", header: http.Header{}, wantIfMatch: ""},
		{name: "empty header value", header: etagHeaderRaw("Etag", ""), wantIfMatch: ""},
		{name: "empty header slice", header: http.Header{"Etag": []string{}}, wantIfMatch: ""},
	}
}

func getResponse(body string, header http.Header) *http.Response {
	return &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(strings.NewReader(body)),
		Header:     header,
	}
}

func assertIfMatch(t *testing.T, headers map[string]string, want string) {
	t.Helper()
	got, ok := headers["If-Match"]
	if want == "" {
		if ok {
			t.Errorf("expected If-Match to be omitted, got %q", got)
		}
		return
	}
	if got != want {
		t.Errorf("If-Match = %q, want %q", got, want)
	}
}

func TestGetObjectETagHeaderCasing(t *testing.T) {
	const body = `{"@odata.id":"/redfish/v1/Chassis/1","Id":"1","Name":"c"}`

	for _, tc := range etagHeaderCases() {
		t.Run(tc.name, func(t *testing.T) {
			client := &TestClient{
				CustomReturnForActions: map[string][]any{
					http.MethodGet: {getResponse(body, tc.header)},
				},
			}

			obj, err := GetObject[Entity](client, "/redfish/v1/Chassis/1")
			if err != nil {
				t.Fatalf("GetObject: %v", err)
			}
			if got := obj.GetETag(); got != tc.wantIfMatch {
				t.Errorf("GetETag() = %q, want %q", got, tc.wantIfMatch)
			}
		})
	}
}

func TestGetObjectPrefersODataEtagOverHeader(t *testing.T) {
	const odataEtag = "from-body"
	body := `{"@odata.id":"/redfish/v1/Chassis/1","@odata.etag":"` + odataEtag + `","Id":"1","Name":"c"}`

	cases := []struct {
		name   string
		header http.Header
	}{
		{name: "missing header", header: http.Header{}},
		{name: "Header.Set ETag", header: etagHeaderSet("ETag", testETag)},
		{name: "raw map key ETag", header: etagHeaderRaw("ETag", testETag)},
		{name: "empty header value", header: etagHeaderRaw("Etag", "")},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			client := &TestClient{
				CustomReturnForActions: map[string][]any{
					http.MethodGet: {getResponse(body, tc.header)},
				},
			}

			obj, err := GetObject[Entity](client, "/redfish/v1/Chassis/1")
			if err != nil {
				t.Fatalf("GetObject: %v", err)
			}
			if got := obj.GetETag(); got != odataEtag {
				t.Errorf("GetETag() = %q, want @odata.etag %q", got, odataEtag)
			}
		})
	}
}

func TestEntityGetETagHeaderCasing(t *testing.T) {
	const body = `{"@odata.id":"/redfish/v1/Chassis/1","Id":"1","Name":"c"}`

	for _, tc := range etagHeaderCases() {
		t.Run(tc.name, func(t *testing.T) {
			client := &TestClient{
				CustomReturnForActions: map[string][]any{
					http.MethodGet: {getResponse(body, tc.header)},
				},
			}

			var e Entity
			if err := e.Get(client, "/redfish/v1/Chassis/1", &e); err != nil {
				t.Fatalf("Entity.Get: %v", err)
			}
			if got := e.GetETag(); got != tc.wantIfMatch {
				t.Errorf("GetETag() = %q, want %q", got, tc.wantIfMatch)
			}
		})
	}
}

func TestEntityGetPrefersODataEtagOverHeader(t *testing.T) {
	const odataEtag = "from-body"
	body := `{"@odata.id":"/redfish/v1/Chassis/1","@odata.etag":"` + odataEtag + `","Id":"1","Name":"c"}`

	client := &TestClient{
		CustomReturnForActions: map[string][]any{
			http.MethodGet: {getResponse(body, etagHeaderRaw("ETag", testETag))},
		},
	}

	var e Entity
	if err := e.Get(client, "/redfish/v1/Chassis/1", &e); err != nil {
		t.Fatalf("Entity.Get: %v", err)
	}
	if got := e.GetETag(); got != odataEtag {
		t.Errorf("GetETag() = %q, want @odata.etag %q", got, odataEtag)
	}
}

func TestUpdateBiosAttributesIfMatchFromETag(t *testing.T) {
	for _, tc := range etagHeaderCases() {
		t.Run(tc.name, func(t *testing.T) {
			var bios Bios
			if err := json.NewDecoder(strings.NewReader(biosBody)).Decode(&bios); err != nil {
				t.Fatalf("decode Bios: %v", err)
			}

			client := &TestClient{
				CustomReturnForActions: map[string][]any{
					http.MethodGet: {getResponse("{}", tc.header)},
				},
			}
			bios.SetClient(client)

			if err := bios.UpdateBiosAttributes(SettingsAttributes{"AssetTag": "test"}); err != nil {
				t.Fatalf("UpdateBiosAttributes: %v", err)
			}

			calls := client.CapturedCalls()
			if len(calls) != 2 {
				t.Fatalf("expected GET then PATCH, got %d calls: %#v", len(calls), calls)
			}
			assertIfMatch(t, calls[1].CustomHeaders, tc.wantIfMatch)
		})
	}
}

func TestSetBootIfMatchFromETag(t *testing.T) {
	for _, tc := range etagHeaderCases() {
		t.Run(tc.name, func(t *testing.T) {
			var system ComputerSystem
			if err := json.NewDecoder(strings.NewReader(computerSystemBodyWithSettings)).Decode(&system); err != nil {
				t.Fatalf("decode ComputerSystem: %v", err)
			}

			client := &TestClient{
				CustomReturnForActions: map[string][]any{
					http.MethodPatch: {
						&http.Response{
							StatusCode: http.StatusBadRequest,
							Body: io.NopCloser(strings.NewReader(
								`{"error":{"code":"Base.1.0.GeneralError","message":"Boot is a read-only property"}}`)),
						},
						nil,
					},
					http.MethodGet: {getResponse("{}", tc.header)},
				},
			}
			system.SetClient(client)

			err := system.SetBoot(&Boot{
				BootSourceOverrideTarget:  PxeBootSource,
				BootSourceOverrideEnabled: OnceBootSourceOverrideEnabled,
			})
			if err != nil {
				t.Fatalf("SetBoot: %v", err)
			}

			calls := client.CapturedCalls()
			if len(calls) != 3 {
				t.Fatalf("expected PATCH, GET, PATCH, got %d calls: %#v", len(calls), calls)
			}
			assertIfMatch(t, calls[2].CustomHeaders, tc.wantIfMatch)
		})
	}
}

func TestUpdateBootAttributesIfMatchFromETag(t *testing.T) {
	for _, tc := range etagHeaderCases() {
		t.Run(tc.name, func(t *testing.T) {
			var system ComputerSystem
			if err := json.NewDecoder(strings.NewReader(computerSystemBodyWithSettings)).Decode(&system); err != nil {
				t.Fatalf("decode ComputerSystem: %v", err)
			}

			client := &TestClient{
				CustomReturnForActions: map[string][]any{
					http.MethodGet: {getResponse("{}", tc.header)},
				},
			}
			system.SetClient(client)

			err := system.UpdateBootAttributes(SettingsAttributes{"BootSourceOverrideTarget": "Pxe"})
			if err != nil {
				t.Fatalf("UpdateBootAttributes: %v", err)
			}

			calls := client.CapturedCalls()
			if len(calls) != 2 {
				t.Fatalf("expected GET then PATCH, got %d calls: %#v", len(calls), calls)
			}
			assertIfMatch(t, calls[1].CustomHeaders, tc.wantIfMatch)
		})
	}
}

func TestSecureBootDatabaseAddCertificateETagFromHeader(t *testing.T) {
	for _, tc := range etagHeaderCases() {
		t.Run(tc.name, func(t *testing.T) {
			client := &TestClient{
				CustomReturnForActions: map[string][]any{
					http.MethodPost: {testResponse(http.StatusCreated, secureBootCertificateBody, tc.header)},
				},
			}
			db := newTestSecureBootDatabase(t, client)

			certificate, err := db.AddCertificate(testPEMCertificate, PEMCertificateType, "")
			if err != nil {
				t.Fatalf("AddCertificate: %v", err)
			}
			if certificate == nil {
				t.Fatal("expected created certificate")
			}
			if got := certificate.GetETag(); got != tc.wantIfMatch {
				t.Errorf("GetETag() = %q, want %q", got, tc.wantIfMatch)
			}
		})
	}
}
