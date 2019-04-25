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
	"fmt"
	"net/http"

	"github.com/stmcginnis/gofish/school/common"
)

// apiClient represents a connection to a Redfish/Swordfish enabled service
// or device.
type apiClient struct {
	// Endpoint is the URL of the *fish service
	Endpoint string

	// httpClient is for direct http actions
	httpClient http.Client
}

// APIClient creates a new client connection to a Redfish service.
func APIClient(host string, port uint16, useHTTPS bool) common.Client {
	http := "http"
	if useHTTPS {
		http = "https"
	}

	endpoint := fmt.Sprintf("%s://%s:%d", http, host, port)

	client := apiClient{
		Endpoint: endpoint}
	return client
}

// Get performs a GET request against the Redfish service.
func (c apiClient) Get(url string) (*http.Response, error) {
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
	req.Close = true

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (c *apiClient) Post() {

}

func (c *apiClient) Put() {

}

func (c *apiClient) Patch() {

}

func (c *apiClient) Delete() {

}
