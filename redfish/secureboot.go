//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// ResetKeysType is method for resetting keys.
type ResetKeysType string

const (
	// ResetAllKeysToDefaultResetKeysType Reset the content of all UEFI
	// Secure Boot key databases (PK, KEK, DB, DBX) to their default values.
	ResetAllKeysToDefaultResetKeysType ResetKeysType = "ResetAllKeysToDefault"
	// DeleteAllKeysResetKeysType Delete the content of all UEFI Secure Boot
	// key databases (PK, KEK, DB, DBX). This puts the system in Setup Mode.
	DeleteAllKeysResetKeysType ResetKeysType = "DeleteAllKeys"
	// DeletePKResetKeysType Delete the content of the PK UEFI Secure Boot
	// database. This puts the system in Setup Mode.
	DeletePKResetKeysType ResetKeysType = "DeletePK"
)

// SecureBootCurrentBootType is the type of secure boot.
type SecureBootCurrentBootType string

const (

	// EnabledSecureBootCurrentBootType Secure Boot is currently enabled.
	EnabledSecureBootCurrentBootType SecureBootCurrentBootType = "Enabled"
	// DisabledSecureBootCurrentBootType Secure Boot is currently disabled.
	DisabledSecureBootCurrentBootType SecureBootCurrentBootType = "Disabled"
)

// SecureBootModeType is the boot mode.
type SecureBootModeType string

const (

	// SetupModeSecureBootModeType Secure Boot is currently in Setup Mode.
	SetupModeSecureBootModeType SecureBootModeType = "SetupMode"
	// UserModeSecureBootModeType Secure Boot is currently in User Mode.
	UserModeSecureBootModeType SecureBootModeType = "UserMode"
	// AuditModeSecureBootModeType Secure Boot is currently in Audit Mode.
	AuditModeSecureBootModeType SecureBootModeType = "AuditMode"
	// DeployedModeSecureBootModeType Secure Boot is currently in Deployed
	// Mode.
	DeployedModeSecureBootModeType SecureBootModeType = "DeployedMode"
)

// SecureBoot is used to represent a UEFI Secure Boot resource.
type SecureBoot struct {
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
	// SecureBootCurrentBoot shall indicate the UEFI Secure Boot state during
	// the current boot cycle.
	SecureBootCurrentBoot SecureBootCurrentBootType
	// SecureBootEnable set to true enables UEFI Secure Boot, and setting it to
	// false disables it. This property can be enabled only in UEFI boot mode.
	SecureBootEnable bool
	// SecureBootMode shall contain the current Secure Boot mode, as defined in
	// the UEFI Specification.
	SecureBootMode SecureBootModeType
}

// GetSecureBoot will get a SecureBoot instance from the service.
func GetSecureBoot(c common.Client, uri string) (*SecureBoot, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var secureboot SecureBoot
	err = json.NewDecoder(resp.Body).Decode(&secureboot)
	if err != nil {
		return nil, err
	}

	secureboot.SetClient(c)
	return &secureboot, nil
}

// ListReferencedSecureBoots gets the collection of SecureBoot from
// a provided reference.
func ListReferencedSecureBoots(c common.Client, link string) ([]*SecureBoot, error) {
	var result []*SecureBoot
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, securebootLink := range links.ItemLinks {
		secureboot, err := GetSecureBoot(c, securebootLink)
		if err != nil {
			return result, err
		}
		result = append(result, secureboot)
	}

	return result, nil
}
