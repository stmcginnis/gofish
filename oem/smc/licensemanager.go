//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"errors"

	"github.com/stmcginnis/gofish/schemas"
)

// LicenseManager is the license manager instance associated with the system.
type LicenseManager struct {
	schemas.Entity

	queryLicense string

	activateLicenseTarget string
	clearLicenseTarget    string
}

// UnmarshalJSON unmarshals a LicenseManager object from the raw JSON.
func (lm *LicenseManager) UnmarshalJSON(b []byte) error {
	type temp LicenseManager
	var t struct {
		temp
		Actions struct {
			ActivateLicense schemas.ActionTarget `json:"#LicenseManager.ActivateLicense"`
			ClearLicense    schemas.ActionTarget `json:"#LicenseManager.ClearLicense"`
		}
		QueryLicense schemas.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*lm = LicenseManager(t.temp)
	lm.queryLicense = t.QueryLicense.String()
	lm.activateLicenseTarget = t.Actions.ActivateLicense.Target
	lm.clearLicenseTarget = t.Actions.ClearLicense.Target

	return nil
}

// GetLicenseManager will get a LicenseManager instance from the service.
func GetLicenseManager(c schemas.Client, uri string) (*LicenseManager, error) {
	return schemas.GetObject[LicenseManager](c, uri)
}

// ActivateLicense performs the ActivateLicense action of the LicenseManager.
func (lm *LicenseManager) ActivateLicense() error {
	if lm.activateLicenseTarget == "" {
		return errors.New("ActivateLicense is not supported by this system")
	}

	return lm.Post(lm.activateLicenseTarget, nil)
}

// ClearLicense performs the ClearLicense action of the LicenseManager.
func (lm *LicenseManager) ClearLicense() error {
	if lm.clearLicenseTarget == "" {
		return errors.New("ClearLicense is not supported by this system")
	}

	return lm.Post(lm.clearLicenseTarget, nil)
}

// QueryLicense will get the license information from the service.
func (lm *LicenseManager) QueryLicense() (*QueryLicense, error) {
	return GetQueryLicense(lm.GetClient(), lm.queryLicense)
}

// QueryLicense contains license information.
type QueryLicense struct {
	schemas.Entity
	Licenses []string
}

// GetQueryLicense will get the QueryLicense instance from the service.
func GetQueryLicense(c schemas.Client, uri string) (*QueryLicense, error) {
	return schemas.GetObject[QueryLicense](c, uri)
}
