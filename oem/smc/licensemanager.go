//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/coreweave/gofish/common"
)

// LicenseManager is the license manager instance associated with the system.
type LicenseManager struct {
	common.Entity

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
			ActivateLicense common.ActionTarget `json:"#LicenseManager.ActivateLicense"`
			ClearLicense    common.ActionTarget `json:"#LicenseManager.ClearLicense"`
		}
		QueryLicense common.Link
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
func GetLicenseManager(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*LicenseManager, error) {
	return GetLicenseManagerWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetLicenseManagerWithContext will get a LicenseManager instance from the service.
func GetLicenseManagerWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*LicenseManager, error) {
	return common.GetObjectWithContext[LicenseManager](ctx, c, uri, queryOpts...)
}

// ActivateLicense performs the ActivateLicense action of the LicenseManager.
func (lm *LicenseManager) ActivateLicense() error {
	return lm.ActivateLicenseWithContext(common.ContextOf(lm.GetClient()))
}

// ActivateLicenseWithContext performs the ActivateLicense action of the LicenseManager.
func (lm *LicenseManager) ActivateLicenseWithContext(ctx context.Context) error {
	if lm.activateLicenseTarget == "" {
		return errors.New("ActivateLicense is not supported by this system")
	}

	return lm.PostWithContext(ctx, lm.activateLicenseTarget, nil)
}

// ClearLicense performs the ClearLicense action of the LicenseManager.
func (lm *LicenseManager) ClearLicense() error {
	return lm.ClearLicenseWithContext(common.ContextOf(lm.GetClient()))
}

// ClearLicenseWithContext performs the ClearLicense action of the LicenseManager.
func (lm *LicenseManager) ClearLicenseWithContext(ctx context.Context) error {
	if lm.clearLicenseTarget == "" {
		return errors.New("ClearLicense is not supported by this system")
	}

	return lm.PostWithContext(ctx, lm.clearLicenseTarget, nil)
}

// QueryLicense will get the license information from the service.
func (lm *LicenseManager) QueryLicense(queryOpts ...common.QueryGroupOption) (*QueryLicense, error) {
	return lm.QueryLicenseWithContext(common.ContextOf(lm.GetClient()), queryOpts...)
}

// QueryLicenseWithContext will get the license information from the service.
func (lm *LicenseManager) QueryLicenseWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*QueryLicense, error) {
	return GetQueryLicenseWithContext(ctx, lm.GetClient(), lm.queryLicense, queryOpts...)
}

// QueryLicense contains license information.
type QueryLicense struct {
	common.Entity
	Licenses []string
}

// GetQueryLicense will get the QueryLicense instance from the service.
func GetQueryLicense(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*QueryLicense, error) {
	return GetQueryLicenseWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetQueryLicenseWithContext will get the QueryLicense instance from the service.
func GetQueryLicenseWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*QueryLicense, error) {
	return common.GetObjectWithContext[QueryLicense](ctx, c, uri, queryOpts...)
}
