//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"strings"
	"testing"
)

var licenseManagerBody = `{
  "@odata.type": "#LicenseManager.v1_0_0.LicenseManager",
  "@odata.id": "/redfish/v1/Managers/1/LicenseManager",
  "Id": "License Manager",
  "Name": "License Manager",
  "QueryLicense": {
    "@odata.id": "/redfish/v1/Managers/1/LicenseManager/QueryLicense"
  },
  "Actions": {
    "Oem": {},
    "#LicenseManager.ActivateLicense": {
      "target": "/redfish/v1/Managers/1/LicenseManager/Actions/LicenseManager.ActivateLicense"
    },
    "#LicenseManager.ClearLicense": {
      "target": "/redfish/v1/Managers/1/LicenseManager/Actions/LicenseManager.ClearLicense"
    }
  },
  "@odata.etag": "\"1e5832264809b8196fcffe3926e503d2\""
}`

var queryLicenseBody = `{
  "@odata.type": "#QueryLicense.v1_0_0.QueryLicense",
  "@odata.id": "/redfish/v1/Managers/1/LicenseManager/QueryLicense",
  "Id": "1",
  "Name": "LicenseManager",
  "Licenses": [
    "{\"ProductKey\":{\"Node\":{\"LicenseID\":\"1\",\"LicenseName\":\"SFT-OOB-LIC\",\"CreateDate\":\"20240918\"},\"Signature\":\"IDCGR8VNy9Uy2ZIVW0iZs19aS3R6oZjelbSHTDHL6LrgF8Pq/kRzBlRQ4jNc36gXmK7Bl75rHAD2dPyjzdimOZDH/N5iR1XuQFI72/FtJhpnaKyN+l+I4QuArSEZodM9IaBQcttz4QQZyu3Oa5qKGx68PySwLG2prNid62Ts1E2Ni7KpjxtxETzczsDedNoLVi9R2g/2FXhTJiY3XIGANcgJxD/oaGWjM9ebHtsQgY1bXaw8bT6MAJv1pD7iwxuMi+RBrWopvOH7qVylD/vTE9FG6RpwOU79RMmvEcAjXOz32ezHyZECpk5NiRQhHXt7/y2uTh1Bclcl8zdnmHiSy==\"}}",
    "{\"ProductKey\":{\"Node\":{\"LicenseID\":\"2\",\"LicenseName\":\"SFT-DCMS-SINGLE\",\"CreateDate\":\"20240918\"},\"Signature\":\"oI13iOXFGszhTdfgvd95DWaErCFNMVhMden7uZ/p/vSFno3rXasyeh3eSUEH3j6BsIs/vxdv0QqxR4BfZEhkPh9+3wdqTaQksvulTyiLV5SQ8Pw5iBtz/5KjAQVKuNNsbva0ZH78an5/X7ZARtIxvXHk+Fbbb6UGx4WYQ4qoz4RfQ6wO0aYa7FV1+fMHLhHgxZS85zOmfY5oEwvkoQgnK2qtlrCmhBNe66KyBv+OhukZm1gagW0EoXq43oPplwOXJDMXfzldY9rML7u6MhaY8KQk4Pss3HakO8pmGVpAROs4mpxCPLobHfCiomaJjqv8/bNofzE5E2jCTM8VfcwGV==\"}}"
  ],
  "@odata.etag": "\"bfb9e657900835da4081bf525cf454d0\""
}`

// TestLicenseManager tests the parsing of LicenseManager objects.
func TestLicenseManager(t *testing.T) {
	var result LicenseManager
	err := json.NewDecoder(strings.NewReader(licenseManagerBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "License Manager" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.queryLicense != "/redfish/v1/Managers/1/LicenseManager/QueryLicense" {
		t.Errorf("Invalid query license link: %s", result.queryLicense)
	}

	if result.activateLicenseTarget != "/redfish/v1/Managers/1/LicenseManager/Actions/LicenseManager.ActivateLicense" {
		t.Errorf("Invalid activate license target link: %s", result.activateLicenseTarget)
	}

	if result.clearLicenseTarget != "/redfish/v1/Managers/1/LicenseManager/Actions/LicenseManager.ClearLicense" {
		t.Errorf("Invalid clear license target link: %s", result.clearLicenseTarget)
	}
}

// TestQueryLicense tests the parsing of QueryLicense objects.
func TestQueryLicense(t *testing.T) {
	var result QueryLicense
	err := json.NewDecoder(strings.NewReader(queryLicenseBody)).Decode(&result)
	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if len(result.Licenses) != 2 {
		t.Errorf("Expected 2 licenses, got %d", len(result.Licenses))
	}

	if !strings.Contains(result.Licenses[0], "SFT-OOB-LIC") {
		t.Errorf("Expected license to contain a license string, got %s", result.Licenses[0])
	}
}
