// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gofish

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"strings"

	"github.com/stmcginnis/gofish/school/common"
)

// ApiClient represents a connection to a Redfish/Swordfish enabled service
// or device.
type ApiClient struct {
	// Endpoint is the URL of the *fish service
	Endpoint string

	// Token is the session token to be used for all requests issued
	Token string

	// httpClient is for direct http actions
	httpClient *http.Client
}

// APIClient creates a new client connection to a Redfish service.
func APIClient(endpoint string, httpClient *http.Client) (c *ApiClient, err error) {
	if !strings.HasPrefix(endpoint, "http") {
		return c, fmt.Errorf("endpoint must starts with http or https")
	}
	client := &ApiClient{Endpoint: endpoint}
	if httpClient != nil {
		client.httpClient = httpClient
	} else {
		client.httpClient = &http.Client{}
	}
	return client, err
}

// Get performs a GET request against the Redfish service.
func (c *ApiClient) Get(url string) (*http.Response, error) {
	relativePath := url
	if relativePath == "" {
		relativePath = common.DefaultServiceRoot
	}

	endpoint := fmt.Sprintf("%s%s", c.Endpoint, relativePath)
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gofish/1.0.0")
	req.Header.Set("Accept", "application/json")
	if c.Token != "" {
		req.Header.Set("X-Auth-Token", c.Token)
	}
	req.Close = true

	resp, err := c.httpClient.Do(req)
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
func (c *ApiClient) Post(url string, payload []byte) (*http.Response, error) {
	relativePath := url
	if relativePath == "" {
		relativePath = common.DefaultServiceRoot
	}

	endpoint := fmt.Sprintf("%s%s", c.Endpoint, relativePath)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "gofish/1.0.0")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if c.Token != "" {
		req.Header.Set("X-Auth-Token", c.Token)
	}
	req.Close = true

	resp, err := c.httpClient.Do(req)
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

// Put makes a PUT call. TODO: Implement
func (c *ApiClient) Put() {

}

// Patch makes a PATCH call. TODO: Implement
func (c *ApiClient) Patch() {

}

// Delete performs a Delete request against the Redfish service.
func (c *ApiClient) Delete(url string) error {
	relativePath := url
	if relativePath == "" {
		relativePath = common.DefaultServiceRoot
	}

	endpoint := fmt.Sprintf("%s%s", c.Endpoint, relativePath)
	req, err := http.NewRequest("DELETE", endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Set("User-Agent", "gofish/1.0.0")
	req.Header.Set("Accept", "application/json")
	if c.Token != "" {
		req.Header.Set("X-Auth-Token", c.Token)
	}
	req.Close = true

	resp, err := c.httpClient.Do(req)
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
