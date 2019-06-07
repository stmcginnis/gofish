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

// IOConnectivityLoSCapabilities describes capabilities of the system to
// support various IO Connectivity service options.
type IOConnectivityLoSCapabilities struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
	// MaxSupportedBytesPerSecond shall be the maximum bytes per second that a
	// connection can support.
	MaxSupportedBytesPerSecond int
	// MaxSupportedIOPS shall be the maximum IOPS that a connection can support.
	MaxSupportedIOPS int
	// SupportedAccessProtocols is Access protocols supported by this service
	// option. NOTE: SMB+NFS* requires that SMB and at least one of NFSv3 or
	// NFXv4 are also selected, (i.e. {'SMB', 'NFSv4', 'SMB+NFS*'}).
	SupportedAccessProtocols []common.Protocol
	// SupportedLinesOfService shall contain known and
	// supported IOConnectivityLinesOfService.
	SupportedLinesOfService []IOConnectivityLineOfService
	// SupportedLinesOfServiceCount is the number of IOConnectivityLineOfServices.
	SupportedLinesOfServiceCount int `json:"SupportedLinesOfService@odata.count"`
}

// GetIOConnectivityLoSCapabilities will get a IOConnectivityLoSCapabilities
// instance from the service.
func GetIOConnectivityLoSCapabilities(c common.Client, uri string) (*IOConnectivityLoSCapabilities, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var ioconnectivityloscapabilities IOConnectivityLoSCapabilities
	err = json.NewDecoder(resp.Body).Decode(&ioconnectivityloscapabilities)
	if err != nil {
		return nil, err
	}

	ioconnectivityloscapabilities.SetClient(c)
	return &ioconnectivityloscapabilities, nil
}

// ListReferencedIOConnectivityLoSCapabilitiess gets the collection of
// IOConnectivityLoSCapabilities from a provided reference.
func ListReferencedIOConnectivityLoSCapabilitiess(c common.Client, link string) ([]*IOConnectivityLoSCapabilities, error) {
	var result []*IOConnectivityLoSCapabilities
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, ioconnectivityloscapabilitiesLink := range links.ItemLinks {
		ioconnectivityloscapabilities, err := GetIOConnectivityLoSCapabilities(c, ioconnectivityloscapabilitiesLink)
		if err != nil {
			return result, err
		}
		result = append(result, ioconnectivityloscapabilities)
	}

	return result, nil
}
