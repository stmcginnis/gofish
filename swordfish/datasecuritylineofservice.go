//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
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
	return common.GetObject[DataSecurityLineOfService](c, uri)
}

// ListReferencedDataSecurityLineOfServices gets the collection of DataSecurityLineOfService from
// a provided reference.
func ListReferencedDataSecurityLineOfServices(c common.Client, link string) ([]*DataSecurityLineOfService, error) {
	return common.GetCollectionObjects[DataSecurityLineOfService](c, link)
}
