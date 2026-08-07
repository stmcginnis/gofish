//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
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
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// CertificatesLink shall be a link to a resource collection of type CertificateCollection.
	CertificatesLink common.Link `json:"Certificates"`
	// DatabaseID shall contain the name of the UEFI Secure Boot database. This property shall contain the same value
	// as the Id property. The value shall be one of the UEFI-defined Secure Boot databases: 'PK', 'KEK' 'db', 'dbx',
	// 'dbr', 'dbt', 'PKDefault', 'KEKDefault', 'dbDefault', 'dbxDefault', 'dbrDefault', or 'dbtDefault'.
	DatabaseID string
	// Description provides a description of this resource.
	Description string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// SignaturesLink shall be a link to a resource collection of type SignatureCollection.
	SignaturesLink common.Link `json:"Signatures"`
	// Actions shall contain the available actions for this resource.
	Actions SecureBootDatabaseActions `json:"Actions"`
}

// SecureBootDatabaseActions shall contain the available actions for a SecureBootDatabase resource.
type SecureBootDatabaseActions struct {
	// ResetKeys shall reset the UEFI Secure Boot key database.
	ResetKeys common.ActionTarget `json:"#SecureBootDatabase.ResetKeys"`
	// Oem shall contain the available OEM-specific actions for this resource.
	Oem json.RawMessage `json:"Oem"`
}

// Certificates get the certificates contained in this UEFI Secure Boot database.
func (securebootdatabase *SecureBootDatabase) Certificates(queryOpts ...common.QueryGroupOption) ([]*Certificate, error) {
	return securebootdatabase.CertificatesWithContext(common.ContextOf(securebootdatabase.GetClient()), queryOpts...)
}

// CertificatesWithContext get the certificates contained in this UEFI Secure Boot database.
func (securebootdatabase *SecureBootDatabase) CertificatesWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Certificate, error) {
	if securebootdatabase.CertificatesLink.IsZero() {
		return nil, nil
	}
	return ListReferencedCertificatesWithContext(ctx, securebootdatabase.GetClient(), securebootdatabase.CertificatesLink.String(), queryOpts...)
}

// Signatures get the certificates contained in this UEFI Secure Boot database.
func (securebootdatabase *SecureBootDatabase) Signatures(queryOpts ...common.QueryGroupOption) ([]*Signature, error) {
	return securebootdatabase.SignaturesWithContext(common.ContextOf(securebootdatabase.GetClient()), queryOpts...)
}

// SignaturesWithContext get the certificates contained in this UEFI Secure Boot database.
func (securebootdatabase *SecureBootDatabase) SignaturesWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Signature, error) {
	if securebootdatabase.SignaturesLink.IsZero() {
		return nil, nil
	}
	return ListReferencedSignaturesWithContext(ctx, securebootdatabase.GetClient(), securebootdatabase.SignaturesLink.String(), queryOpts...)
}

// ResetKeys will perform a reset of this UEFI Secure Boot key database. The `ResetAllKeysToDefault`
// value shall reset this UEFI Secure Boot key database to the default values. The `DeleteAllKeys`
// value shall delete the contents of this UEFI Secure Boot key database.
func (securebootdatabase *SecureBootDatabase) ResetKeys(resetType ResetKeysType) error {
	return securebootdatabase.ResetKeysWithContext(common.ContextOf(securebootdatabase.GetClient()), resetType)
}

// ResetKeysWithContext will perform a reset of this UEFI Secure Boot key database. The `ResetAllKeysToDefault`
// value shall reset this UEFI Secure Boot key database to the default values. The `DeleteAllKeys`
// value shall delete the contents of this UEFI Secure Boot key database.
func (securebootdatabase *SecureBootDatabase) ResetKeysWithContext(ctx context.Context, resetType ResetKeysType) error {
	if securebootdatabase.Actions.ResetKeys.Target == "" {
		return ErrActionNotSupported
	}
	params := struct {
		ResetKeysType ResetKeysType
	}{
		ResetKeysType: resetType,
	}
	return securebootdatabase.PostWithContext(ctx, securebootdatabase.Actions.ResetKeys.Target, params)
}

// GetSecureBootDatabase will get a SecureBootDatabase instance from the service.
func GetSecureBootDatabase(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*SecureBootDatabase, error) {
	return GetSecureBootDatabaseWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetSecureBootDatabaseWithContext will get a SecureBootDatabase instance from the service.
func GetSecureBootDatabaseWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*SecureBootDatabase, error) {
	return common.GetObjectWithContext[SecureBootDatabase](ctx, c, uri, queryOpts...)
}

// ListReferencedSecureBootDatabases gets the collection of SecureBootDatabase from
// a provided reference.
func ListReferencedSecureBootDatabases(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*SecureBootDatabase, error) {
	return ListReferencedSecureBootDatabasesWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedSecureBootDatabasesWithContext gets the collection of SecureBootDatabase from
// a provided reference.
func ListReferencedSecureBootDatabasesWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*SecureBootDatabase, error) {
	return common.GetCollectionObjectsWithContext[SecureBootDatabase](ctx, c, link, queryOpts...)
}
