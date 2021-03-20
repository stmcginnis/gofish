//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

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
	// resetKeysTarget is the URL to send ResetKeys requests.
	resetKeysTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a SecureBoot object from the raw JSON.
func (secureboot *SecureBoot) UnmarshalJSON(b []byte) error {
	type temp SecureBoot
	type actions struct {
		ResetKeys struct {
			Target string
		} `json:"#SecureBoot.ResetKeys"`
	}
	var t struct {
		temp
		Actions actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*secureboot = SecureBoot(t.temp)
	secureboot.resetKeysTarget = t.Actions.ResetKeys.Target

	// This is a read/write object, so we need to save the raw object data for later
	secureboot.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (secureboot *SecureBoot) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(SecureBoot)
	err := original.UnmarshalJSON(secureboot.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"SecureBootEnable",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(secureboot).Elem()

	return secureboot.Entity.Update(originalElement, currentElement, readWriteFields)
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

// ResetKeys shall perform a reset of the Secure Boot key databases. The
// ResetAllKeysToDefault value shall reset the UEFI Secure Boot key databases to
// their default values. The DeleteAllKeys value shall delete the content of the
// UEFI Secure Boot key databases. The DeletePK value shall delete the content
// of the PK Secure boot key.
func (secureboot *SecureBoot) ResetKeys(resetType ResetKeysType) error {
	type temp struct {
		ResetKeysType ResetKeysType
	}
	t := temp{ResetKeysType: resetType}

	_, err := secureboot.Client.Post(secureboot.resetKeysTarget, t)
	return err
}
