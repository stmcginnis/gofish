//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// SecureBootDatabaseResetKeysType is
type SecureBootDatabaseResetKeysType string

const (
	// ResetAllKeysToDefaultSecureBootDatabaseResetKeysType Reset the contents of this UEFI Secure Boot key database to the default
	// values.
	ResetAllKeysToDefaultSecureBootDatabaseResetKeysType SecureBootDatabaseResetKeysType = "ResetAllKeysToDefault"
	// DeleteAllKeysSecureBootDatabaseResetKeysType Delete the contents of this UEFI Secure Boot key database.
	DeleteAllKeysSecureBootDatabaseResetKeysType SecureBootDatabaseResetKeysType = "DeleteAllKeys"
)

// SecureBootDatabase shall be used to represent a UEFI Secure Boot database for a Redfish implementation.
type SecureBootDatabase struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// certificates shall be a link to a resource collection of type CertificateCollection.
	certificates string
	// DatabaseID shall contain the name of the UEFI Secure Boot database. This property shall contain the same value
	// as the Id property. The value shall be one of the UEFI-defined Secure Boot databases: 'PK', 'KEK' 'db', 'dbx',
	// 'dbr', 'dbt', 'PKDefault', 'KEKDefault', 'dbDefault', 'dbxDefault', 'dbrDefault', or 'dbtDefault'.
	DatabaseID string
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Signatures shall be a link to a resource collection of type SignatureCollection.
	signatures string

	resetKeysTarget string
}

// UnmarshalJSON unmarshals a SecureBootDatabase object from the raw JSON.
func (securebootdatabase *SecureBootDatabase) UnmarshalJSON(b []byte) error {
	type temp SecureBootDatabase
	var t struct {
		temp
		Actions struct {
			ResetKeys common.ActionTarget `json:"#SecureBootDatabase.ResetKeys"`
		}
		Certificates common.Link
		Signatures   common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*securebootdatabase = SecureBootDatabase(t.temp)

	// Extract the links to other entities for later
	securebootdatabase.certificates = t.Certificates.String()
	securebootdatabase.signatures = t.Signatures.String()

	securebootdatabase.resetKeysTarget = t.Actions.ResetKeys.Target

	return nil
}

// Certificates get the certificates contained in this UEFI Secure Boot database.
func (securebootdatabase *SecureBootDatabase) Certificates() ([]*Certificate, error) {
	return ListReferencedCertificates(securebootdatabase.GetClient(), securebootdatabase.certificates)
}

// Signatures get the certificates contained in this UEFI Secure Boot database.
func (securebootdatabase *SecureBootDatabase) Signatures() ([]*Signature, error) {
	return ListReferencedSignatures(securebootdatabase.GetClient(), securebootdatabase.signatures)
}

// ResetKeys will perform a reset of this UEFI Secure Boot key database. The `ResetAllKeysToDefault`
// value shall reset this UEFI Secure Boot key database to the default values. The `DeleteAllKeys`
// value shall delete the contents of this UEFI Secure Boot key database.
func (securebootdatabase *SecureBootDatabase) ResetKeys(resetType ResetKeysType) error {
	params := struct {
		ResetKeysType ResetKeysType
	}{
		ResetKeysType: resetType,
	}
	return securebootdatabase.Post(securebootdatabase.resetKeysTarget, params)
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
