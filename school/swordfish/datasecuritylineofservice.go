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

// DataSecurityLineOfService is used to describe data security service
// level requirements.
type DataSecurityLineOfService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// AntivirusEngineProvider shall specify an AntiVirus provider.
	AntivirusEngineProvider string
	// AntivirusScanPolicies shall specify the
	// policy for triggering an AntiVirus scan.
	AntivirusScanPolicies []AntiVirusScanTrigger
	// ChannelEncryptionStrength shall specify a key size in a symmetric
	// encryption algorithm for transport channel encryption.
	ChannelEncryptionStrength KeySize
	// DataSanitizationPolicy shall specify the data sanitization policy.
	DataSanitizationPolicy DataSanitizationPolicy
	// Description provides a description of this resource.
	Description string
	// HostAuthenticationType shall specify the
	// authentication type for hosts (servers) or initiator endpoints.
	HostAuthenticationType AuthenticationType
	// MediaEncryptionStrength shall specify a key
	// size in a symmetric encryption algorithm for media encryption.
	MediaEncryptionStrength KeySize
	// SecureChannelProtocol shall specify the
	// protocol that provide encrypted communication.
	SecureChannelProtocol SecureChannelProtocol
	// UserAuthenticationType shall specify the
	// authentication type for users (or programs).
	UserAuthenticationType AuthenticationType
}

// GetDataSecurityLineOfService will get a DataSecurityLineOfService instance from the service.
func GetDataSecurityLineOfService(c common.Client, uri string) (*DataSecurityLineOfService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var datasecuritylineofservice DataSecurityLineOfService
	err = json.NewDecoder(resp.Body).Decode(&datasecuritylineofservice)
	if err != nil {
		return nil, err
	}

	datasecuritylineofservice.SetClient(c)
	return &datasecuritylineofservice, nil
}

// ListReferencedDataSecurityLineOfServices gets the collection of DataSecurityLineOfService from
// a provided reference.
func ListReferencedDataSecurityLineOfServices(c common.Client, link string) ([]*DataSecurityLineOfService, error) {
	var result []*DataSecurityLineOfService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, datasecuritylineofserviceLink := range links.ItemLinks {
		datasecuritylineofservice, err := GetDataSecurityLineOfService(c, datasecuritylineofserviceLink)
		if err != nil {
			return result, err
		}
		result = append(result, datasecuritylineofservice)
	}

	return result, nil
}
