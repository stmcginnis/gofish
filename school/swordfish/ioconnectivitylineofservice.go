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

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// IOConnectivityLineOfService is used to specify the characteristics of
// storage connectivity.
type IOConnectivityLineOfService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AccessProtocols shall specify the Access protocol for this service
	// option. NOTE: If multiple protocols are specified,  the corresponding
	// MaxSupportedIOPS governs the max achieved across all protocol uses. This
	// may be less than the sum of the individual max values, which may be
	// specified by individual Line of Service entries.
	AccessProtocols []common.Protocol
	// Description provides a description of this resource.
	Description string
	// MaxBytesPerSecond shall be the maximum bytes per second that a connection
	// can support.
	MaxBytesPerSecond int
	// MaxIOPS shall be the maximum IOs per second that the connection shall
	// allow for the selected access protocol.
	MaxIOPS int
}

// GetIOConnectivityLineOfService will get a IOConnectivityLineOfService instance from the service.
func GetIOConnectivityLineOfService(c common.Client, uri string) (*IOConnectivityLineOfService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ioconnectivitylineofservice IOConnectivityLineOfService
	err = json.NewDecoder(resp.Body).Decode(&ioconnectivitylineofservice)
	if err != nil {
		return nil, err
	}

	ioconnectivitylineofservice.SetClient(c)
	return &ioconnectivitylineofservice, nil
}

// ListReferencedIOConnectivityLineOfServices gets the collection of IOConnectivityLineOfService from
// a provided reference.
func ListReferencedIOConnectivityLineOfServices(c common.Client, link string) ([]*IOConnectivityLineOfService, error) {
	var result []*IOConnectivityLineOfService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, ioconnectivitylineofserviceLink := range links.ItemLinks {
		ioconnectivitylineofservice, err := GetIOConnectivityLineOfService(c, ioconnectivitylineofserviceLink)
		if err != nil {
			return result, err
		}
		result = append(result, ioconnectivitylineofservice)
	}

	return result, nil
}
