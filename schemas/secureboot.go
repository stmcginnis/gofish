//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2020.1 - #SecureBoot.v1_1_2.SecureBoot

package schemas

import (
	"encoding/json"
)

type ResetKeysType string

const (
	// ResetAllKeysToDefaultResetKeysType Reset the contents of all UEFI Secure
	// Boot key databases, including the PK key database, to the default values.
	ResetAllKeysToDefaultResetKeysType ResetKeysType = "ResetAllKeysToDefault"
	// DeleteAllKeysResetKeysType Delete the contents of all UEFI Secure Boot key
	// databases, including the PK key database. This puts the system in Setup
	// Mode.
	DeleteAllKeysResetKeysType ResetKeysType = "DeleteAllKeys"
	// DeletePKResetKeysType Delete the contents of the PK UEFI Secure Boot
	// database. This puts the system in Setup Mode.
	DeletePKResetKeysType ResetKeysType = "DeletePK"
)

type SecureBootCurrentBootType string

const (
	// EnabledSecureBootCurrentBootType UEFI Secure Boot is currently enabled.
	EnabledSecureBootCurrentBootType SecureBootCurrentBootType = "Enabled"
	// DisabledSecureBootCurrentBootType UEFI Secure Boot is currently disabled.
	DisabledSecureBootCurrentBootType SecureBootCurrentBootType = "Disabled"
)

type SecureBootModeType string

const (
	// SetupModeSecureBootModeType UEFI Secure Boot is currently in Setup Mode.
	SetupModeSecureBootModeType SecureBootModeType = "SetupMode"
	// UserModeSecureBootModeType UEFI Secure Boot is currently in User Mode.
	UserModeSecureBootModeType SecureBootModeType = "UserMode"
	// AuditModeSecureBootModeType UEFI Secure Boot is currently in Audit Mode.
	AuditModeSecureBootModeType SecureBootModeType = "AuditMode"
	// DeployedModeSecureBootModeType UEFI Secure Boot is currently in Deployed
	// Mode.
	DeployedModeSecureBootModeType SecureBootModeType = "DeployedMode"
)

// SecureBoot This resource contains UEFI Secure Boot information for a Redfish
// implementation.
type SecureBoot struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// SecureBootCurrentBoot shall indicate the UEFI Secure Boot state during the
	// current boot cycle.
	SecureBootCurrentBoot SecureBootCurrentBootType
	// SecureBootDatabases shall be a link to a resource collection of type
	// 'SecureBootDatabaseCollection'.
	//
	// Version added: v1.1.0
	secureBootDatabases string
	// SecureBootEnable shall indicate whether the UEFI Secure Boot takes effect on
	// next boot. This property can be enabled in UEFI boot mode only.
	SecureBootEnable bool
	// SecureBootMode shall contain the current UEFI Secure Boot mode, as defined
	// in the UEFI Specification.
	SecureBootMode SecureBootModeType
	// resetKeysTarget is the URL to send ResetKeys requests.
	resetKeysTarget string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a SecureBoot object from the raw JSON.
func (s *SecureBoot) UnmarshalJSON(b []byte) error {
	type temp SecureBoot
	type sActions struct {
		ResetKeys ActionTarget `json:"#SecureBoot.ResetKeys"`
	}
	var tmp struct {
		temp
		Actions             sActions
		SecureBootDatabases Link `json:"SecureBootDatabases"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*s = SecureBoot(tmp.temp)

	// Extract the links to other entities for later
	s.resetKeysTarget = tmp.Actions.ResetKeys.Target
	s.secureBootDatabases = tmp.SecureBootDatabases.String()

	// This is a read/write object, so we need to save the raw object data for later
	s.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (s *SecureBoot) Update() error {
	readWriteFields := []string{
		"SecureBootEnable",
	}

	return s.UpdateFromRawData(s, s.RawData, readWriteFields)
}

// GetSecureBoot will get a SecureBoot instance from the service.
func GetSecureBoot(c Client, uri string) (*SecureBoot, error) {
	return GetObject[SecureBoot](c, uri)
}

// ListReferencedSecureBoots gets the collection of SecureBoot from
// a provided reference.
func ListReferencedSecureBoots(c Client, link string) ([]*SecureBoot, error) {
	return GetCollectionObjects[SecureBoot](c, link)
}

// This action shall reset the UEFI Secure Boot key databases. The
// 'ResetAllKeysToDefault' value shall reset all UEFI Secure Boot key databases
// to their default values. The 'DeleteAllKeys' value shall delete the contents
// of all UEFI Secure Boot key databases. The 'DeletePK' value shall delete the
// contents of the PK Secure Boot key database.
// resetKeysType - This parameter shall specify the type of reset or delete to
// perform on the UEFI Secure Boot databases.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (s *SecureBoot) ResetKeys(resetKeysType ResetKeysType) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ResetKeysType"] = resetKeysType
	resp, taskInfo, err := PostWithTask(s.client,
		s.resetKeysTarget, payload, s.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// SecureBootDatabases gets the SecureBootDatabases collection.
func (s *SecureBoot) SecureBootDatabases() ([]*SecureBootDatabase, error) {
	if s.secureBootDatabases == "" {
		return nil, nil
	}
	return GetCollectionObjects[SecureBootDatabase](s.client, s.secureBootDatabases)
}
