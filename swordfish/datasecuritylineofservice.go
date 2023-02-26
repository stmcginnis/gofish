//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// DataSecurityLineOfService is used to describe data security service
// level requirements.
type DataSecurityLineOfService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
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
func ListReferencedDataSecurityLineOfServices(c common.Client, link string) ([]*DataSecurityLineOfService, error) { //nolint:dupl
	var result []*DataSecurityLineOfService
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *DataSecurityLineOfService
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		datasecuritylineofservice, err := GetDataSecurityLineOfService(c, link)
		ch <- GetResult{Item: datasecuritylineofservice, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
