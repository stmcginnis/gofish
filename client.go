//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"

	"strings"
	"time"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

const userAgent = "gofish/1.0"
const applicationJSON = "application/json"

// APIClient represents a connection to a Redfish/Swordfish enabled service
// or device.
type APIClient struct {
	// Endpoint is the URL of the *fish service
	endpoint string

	// HTTPClient is for direct http actions
	HTTPClient *http.Client

	// Service is the ServiceRoot of this Redfish instance
	Service *Service

	// Auth information saved for later to be able to log out
	Auth *redfish.AuthToken

	// dumpWriter will receive HTTP dumps if non-nil.
	dumpWriter io.Writer
}

// ClientConfig holds the settings for establishing a connection.
type ClientConfig struct {
	// Endpoint is the URL of the redfish service
	Endpoint string

	// Username is the optional user name to authenticate with.
	Username string

	// Password is the password to use for authentication.
	Password string

	// Insecure controls whether to enforce SSL certificate validity.
	Insecure bool

	// Controls TLS handshake timeout
	TLSHandshakeTimeout int

	// HTTPClient is the optional client to connect with.
	HTTPClient *http.Client

	// DumpWriter is an optional io.Writer to receive dumps of HTTP
	// requests and responses.
	DumpWriter io.Writer

	// BasicAuth tells the APIClient if basic auth should be used (true) or token based auth must be used (false)
	BasicAuth bool
}

// Connect creates a new client connection to a Redfish service.
func Connect(config ClientConfig) (c *APIClient, err error) {

	if !strings.HasPrefix(config.Endpoint, "http") {
		return c, fmt.Errorf("endpoint must starts with http or https")
	}

	client := &APIClient{
		endpoint:   config.Endpoint,
		dumpWriter: config.DumpWriter,
	}

	if config.TLSHandshakeTimeout == 0 {
		config.TLSHandshakeTimeout = 10
	}

	if config.HTTPClient == nil {
		defaultTransport := http.DefaultTransport.(*http.Transport)
		transport := &http.Transport{
			Proxy:                 defaultTransport.Proxy,
			DialContext:           defaultTransport.DialContext,
			MaxIdleConns:          defaultTransport.MaxIdleConns,
			IdleConnTimeout:       defaultTransport.IdleConnTimeout,
			ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
			TLSHandshakeTimeout:   time.Duration(config.TLSHandshakeTimeout) * time.Second,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: config.Insecure,
			},
		}
		client.HTTPClient = &http.Client{Transport: transport}
	} else {
		client.HTTPClient = config.HTTPClient
	}

	if config.Username != "" {
		// Authenticate with the service
		service, err := ServiceRoot(client)
		if err != nil {
			return nil, err
		}

		var auth *redfish.AuthToken
		if config.BasicAuth {
			auth = &redfish.AuthToken{
				Username:  config.Username,
				Password:  config.Password,
				BasicAuth: true,
			}
		} else {
			auth, err = service.CreateSession(config.Username, config.Password)
			if err != nil {
				return nil, err
			}
		}

		client.Service = service
		client.Auth = auth
	}

	return client, err
}

// ConnectDefault creates an unauthenticated connection to a Redfish service.
func ConnectDefault(endpoint string) (c *APIClient, err error) {
	if !strings.HasPrefix(endpoint, "http") {
		return c, fmt.Errorf("endpoint must starts with http or https")
	}

	client := &APIClient{endpoint: endpoint}
	client.HTTPClient = &http.Client{}

	// Fetch the service root
	service, err := ServiceRoot(client)
	if err != nil {
		return nil, err
	}
	client.Service = service

	return client, err
}

// Get performs a GET request against the Redfish service.
func (c *APIClient) Get(url string) (*http.Response, error) {
	relativePath := url
	if relativePath == "" {
		relativePath = common.DefaultServiceRoot
	}

	return c.runRequest(http.MethodGet, relativePath, nil)
}

// Post performs a Post request against the Redfish service.
func (c *APIClient) Post(url string, payload interface{}) (*http.Response, error) {
	return c.runRequest(http.MethodPost, url, payload)
}

// Put performs a Put request against the Redfish service.
func (c *APIClient) Put(url string, payload interface{}) (*http.Response, error) {
	return c.runRequest(http.MethodPut, url, payload)
}

// Patch performs a Patch request against the Redfish service.
func (c *APIClient) Patch(url string, payload interface{}) (*http.Response, error) {
	return c.runRequest(http.MethodPatch, url, payload)
}

// Delete performs a Delete request against the Redfish service.
func (c *APIClient) Delete(url string) error {
	resp, err := c.runRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}
	if resp != nil && resp.Body != nil {
		resp.Body.Close()
	}
	return nil
}

// runRequest actually performs the REST calls.
func (c *APIClient) runRequest(method string, url string, payload interface{}) (*http.Response, error) {
	if url == "" {
		return nil, fmt.Errorf("unable to execute request, no target provided")
	}

	var payloadBuffer io.ReadSeeker
	if payload != nil {
		body, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		payloadBuffer = bytes.NewReader(body)
	}

	endpoint := fmt.Sprintf("%s%s", c.endpoint, url)
	req, err := http.NewRequest(method, endpoint, payloadBuffer)
	if err != nil {
		return nil, err
	}

	// Add common headers
	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", applicationJSON)

	// Add content info if present
	if payload != nil {
		req.Header.Set("Content-Type", applicationJSON)
	}

	// Add auth info if authenticated
	if c.Auth != nil {
		if c.Auth.Token != "" {
			req.Header.Set("X-Auth-Token", c.Auth.Token)
		} else {
			if c.Auth.BasicAuth == true && c.Auth.Username != "" && c.Auth.Password != "" {
				encodedAuth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%v:%v", c.Auth.Username, c.Auth.Password)))
				req.Header.Set("Authorization", fmt.Sprintf("Basic %v", encodedAuth))
			}
		}
	}
	req.Close = true

	// Dump request if needed.
	if c.dumpWriter != nil {
		d, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return nil, err
		}

		d = append(d, '\n')
		_, err = c.dumpWriter.Write(d)
		if err != nil {
			panic(err)
		}
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	// Dump response if needed.
	if c.dumpWriter != nil {
		d, err := httputil.DumpResponse(resp, true)
		if err != nil {
			defer resp.Body.Close()
			return nil, err
		}

		d = append(d, '\n')
		_, err = c.dumpWriter.Write(d)
		if err != nil {
			panic(err)
		}
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 && resp.StatusCode != 202 && resp.StatusCode != 204 {
		payload, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, string(payload))
	}

	return resp, err
}

// Logout will delete any active session. Useful to defer logout when creating
// a new connection.
func (c *APIClient) Logout() {
	if c.Service != nil && c.Auth != nil {
		_ = c.Service.DeleteSession(c.Auth.Session)
	}
}
