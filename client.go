//
// SPDX-License-Identifier: BSD-3-Clause
//

package gofish

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"

	"strings"

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
	auth *redfish.AuthToken
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

	// HTTPClient is the optional client to connect with.
	HTTPClient *http.Client
}

// Connect creates a new client connection to a Redfish service.
func Connect(config ClientConfig) (c *APIClient, err error) {

	if !strings.HasPrefix(config.Endpoint, "http") {
		return c, fmt.Errorf("endpoint must starts with http or https")
	}

	client := &APIClient{endpoint: config.Endpoint}

	if config.HTTPClient == nil {
		defaultTransport := http.DefaultTransport.(*http.Transport)
		transport := &http.Transport{
			Proxy:                 defaultTransport.Proxy,
			DialContext:           defaultTransport.DialContext,
			MaxIdleConns:          defaultTransport.MaxIdleConns,
			IdleConnTimeout:       defaultTransport.IdleConnTimeout,
			ExpectContinueTimeout: defaultTransport.ExpectContinueTimeout,
			TLSHandshakeTimeout:   defaultTransport.TLSHandshakeTimeout,
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

		auth, err := service.CreateSession(config.Username, config.Password)
		if err != nil {
			return nil, err
		}

		client.Service = service
		client.auth = auth
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

	endpoint := fmt.Sprintf("%s%s", c.endpoint, relativePath)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Accept", applicationJSON)
	if c.auth != nil && c.auth.Token != "" {
		req.Header.Set("X-Auth-Token", c.auth.Token)
	}
	req.Close = true

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		payload, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		return nil, fmt.Errorf("%d: %s", resp.StatusCode, string(payload))
	}

	return resp, err
}

// Post performs a Post request against the Redfish service.
func (c *APIClient) Post(url string, payload []byte) (*http.Response, error) {
	relativePath := url
	if relativePath == "" {
		relativePath = common.DefaultServiceRoot
	}

	endpoint := fmt.Sprintf("%s%s", c.endpoint, relativePath)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", applicationJSON)
	req.Header.Set("Accept", applicationJSON)
	if c.auth != nil && c.auth.Token != "" {
		req.Header.Set("X-Auth-Token", c.auth.Token)
	}
	req.Close = true

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
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

// Put performs a Put request against the Redfish service.
func (c *APIClient) Put(url string, payload []byte) (*http.Response, error) {
	relativePath := url
	if relativePath == "" {
		relativePath = common.DefaultServiceRoot
	}

	endpoint := fmt.Sprintf("%s%s", c.endpoint, relativePath)
	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", applicationJSON)
	req.Header.Set("Accept", applicationJSON)
	if c.auth != nil && c.auth.Token != "" {
		req.Header.Set("X-Auth-Token", c.auth.Token)
	}
	req.Close = true

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
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

// Patch performs a Patch request against the Redfish service.
func (c *APIClient) Patch(url string, payload []byte) (*http.Response, error) {
	relativePath := url
	if relativePath == "" {
		relativePath = common.DefaultServiceRoot
	}

	endpoint := fmt.Sprintf("%s%s", c.endpoint, relativePath)
	req, err := http.NewRequest("Patch", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", userAgent)
	req.Header.Set("Content-Type", applicationJSON)
	req.Header.Set("Accept", applicationJSON)
	if c.auth != nil && c.auth.Token != "" {
		req.Header.Set("X-Auth-Token", c.auth.Token)
	}
	req.Close = true

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
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

// Delete performs a Delete request against the Redfish service.
func (c *APIClient) Delete(url string) error {
	relativePath := url
	if relativePath == "" {
		relativePath = common.DefaultServiceRoot
	}

	endpoint := fmt.Sprintf("%s%s", c.endpoint, relativePath)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "gofish/1.0.0")
	req.Header.Set("Accept", "application/json")
	if c.auth != nil && c.auth.Token != "" {
		req.Header.Set("X-Auth-Token", c.auth.Token)
	}
	req.Close = true

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 202 && resp.StatusCode != 204 {
		payload, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		return fmt.Errorf("%d: %s", resp.StatusCode, string(payload))
	}

	return err
}

// Logout will delete any active session. Useful to defer logout when creating
// a new connection.
func (c *APIClient) Logout() {
	if c.Service != nil && c.auth != nil {
		c.Service.DeleteSession(c.auth.Session)
	}
}
