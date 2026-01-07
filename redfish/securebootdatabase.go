//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.1 - #SecureBootDatabase.v1_0_3.SecureBootDatabase

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// SecureBootDatabase shall be used to represent a UEFI Secure Boot database for
// a Redfish implementation.
type SecureBootDatabase struct {
	common.Entity
	// Certificates shall be a link to a resource collection of type
	// 'CertificateCollection'.
	certificates string
	// DatabaseId shall contain the name of the UEFI Secure Boot database. This
	// property shall contain the same value as the 'Id' property. The value shall
	// be one of the UEFI-defined Secure Boot databases: 'PK', 'KEK' 'db', 'dbx',
	// 'dbr', 'dbt', 'PKDefault', 'KEKDefault', 'dbDefault', 'dbxDefault',
	// 'dbrDefault', or 'dbtDefault'.
	DatabaseID string `json:"DatabaseId"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Signatures shall be a link to a resource collection of type
	// 'SignatureCollection'.
	signatures string
	// resetKeysTarget is the URL to send ResetKeys requests.
	resetKeysTarget string
}

// UnmarshalJSON unmarshals a SecureBootDatabase object from the raw JSON.
func (s *SecureBootDatabase) UnmarshalJSON(b []byte) error {
	type temp SecureBootDatabase
	type sActions struct {
		ResetKeys common.ActionTarget `json:"#SecureBootDatabase.ResetKeys"`
	}
	var tmp struct {
		temp
		Actions      sActions
		Certificates common.Link `json:"certificates"`
		Signatures   common.Link `json:"signatures"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SecureBootDatabase(tmp.temp)

	// Extract the links to other entities for later
	s.resetKeysTarget = tmp.Actions.ResetKeys.Target
	s.certificates = tmp.Certificates.String()
	s.signatures = tmp.Signatures.String()

	return nil
}

// GetSecureBootDatabase will get a SecureBootDatabase instance from the service.
func GetSecureBootDatabase(c common.Client, uri string) (*SecureBootDatabase, error) {
	return common.GetObject[SecureBootDatabase](c, uri)
}

// ListReferencedSecureBootDatabases gets the collection of SecureBootDatabase from
// a provided reference.
func ListReferencedSecureBootDatabases(c common.Client, link string) ([]*SecureBootDatabase, error) {
	return common.GetCollectionObjects[SecureBootDatabase](c, link)
}

// ResetKeys shall perform a reset of this UEFI Secure Boot key database. The
// 'ResetAllKeysToDefault' value shall reset this UEFI Secure Boot key database
// to the default values. The 'DeleteAllKeys' value shall delete the contents
// of this UEFI Secure Boot key database.
// resetKeysType - This parameter shall specify the type of reset or delete to
// perform on this UEFI Secure Boot database.
func (s *SecureBootDatabase) ResetKeys(resetKeysType ResetKeysType) error {
	payload := make(map[string]any)
	payload["ResetKeysType"] = resetKeysType
	return s.Post(s.resetKeysTarget, payload)
}

// Certificates gets the Certificates collection.
func (s *SecureBootDatabase) Certificates(client common.Client) ([]*Certificate, error) {
	if s.certificates == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Certificate](client, s.certificates)
}

// Signatures gets the Signatures collection.
func (s *SecureBootDatabase) Signatures(client common.Client) ([]*Signature, error) {
	if s.signatures == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[Signature](client, s.signatures)
}
