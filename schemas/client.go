//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import (
	"io"
	"net/http"
)

type ClientSettings struct {
	DefaultQueryOptions []QueryGroupOption
}

// Client is a connection to a Redfish service.
type Client interface {
	GetSettings() ClientSettings
	Get(url string) (*http.Response, error)
	GetWithHeaders(url string, customHeaders map[string]string) (*http.Response, error)
	Post(url string, payload any) (*http.Response, error)
	PostWithHeaders(url string, payload any, customHeaders map[string]string) (*http.Response, error)
	PostMultipart(url string, payload map[string]io.Reader) (*http.Response, error)
	PostMultipartWithHeaders(url string, payload map[string]io.Reader, customHeaders map[string]string) (*http.Response, error)
	Patch(url string, payload any) (*http.Response, error)
	PatchWithHeaders(url string, payload any, customHeaders map[string]string) (*http.Response, error)
	Put(url string, payload any) (*http.Response, error)
	PutWithHeaders(url string, payload any, customHeaders map[string]string) (*http.Response, error)
	Delete(url string) (*http.Response, error)
	DeleteWithHeaders(url string, customHeaders map[string]string) (*http.Response, error)
}
