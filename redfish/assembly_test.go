//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"strings"
	"testing"
)

var assemblyBody = strings.NewReader(
	`{
		"@odata.id": "/redfish/v1/Assembly/1",
		"@odata.context": "/redfish/v1/$metadata#Assembly.Assembly",
		"@odata.type": "#Assembly.v1_2_1.Assembly",
		"Assemblies": [
			{
				"@odata.id": "/redfish/v1/Assembly/1/AssemblyData",
				"BinaryDataURI": "/image/path",
				"Description": "Assembly Description",
				"EngineeringChangeLevel": "1.2.3",
				"MemberId": "1",
				"Model": "AssemblyOne",
				"Name": "Assembly One",
				"PartNumber": "Assembly Part",
				"PhysicalContext": "ContextInfo",
				"Producer": "Vendor One",
				"ProductionDate": "2013-09-11T17:03:55+00:00",
				"SKU": "1234",
				"SerialNumber": "12345",
				"SparePartNumber": "12345-1234",
				"Status": {
					"State": "Enabled",
					"Health": "OK",
					"HealthRollup": "OK"
				},
				"Vendor": "Vendor One",
				"Version": "1.2.3.4"
			}
		],
		"Assemblies@odata.count": 1,
		"Description": "Assembly Description",
		"Id": "1",
		"Name": "Assembly One"
	}`)

var assemblyDataBody = strings.NewReader(
	`{
		"@odata.id": "/redfish/v1/Assembly/1/AssemblyData",
		"BinaryDataURI": "/image/path",
		"Description": "Assembly Description",
		"EngineeringChangeLevel": "1.2.3",
		"MemberId": "1",
		"Model": "AssemblyOne",
		"Name": "Assembly One",
		"PartNumber": "Assembly Part",
		"PhysicalContext": "ContextInfo",
		"Producer": "Vendor One",
		"ProductionDate": "2013-09-11T17:03:55+00:00",
		"SKU": "1234",
		"SerialNumber": "12345",
		"SparePartNumber": "12345-1234",
		"Status": {
			"State": "Enabled",
			"Health": "OK",
			"HealthRollup": "OK"
		},
		"Vendor": "Vendor One",
		"Version": "1.2.3.4"
	}`)

// TestAssembly tests the parsing of Assembly objects.
func TestAssembly(t *testing.T) {
	var result Assembly
	err := json.NewDecoder(assemblyBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.ID != "1" {
		t.Errorf("Received invalid ID: %s", result.ID)
	}

	if result.Name != "Assembly One" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if len(result.Assemblies) != result.AssembliesCount {
		t.Errorf("Parsed %d assemblies but should be %d",
			len(result.Assemblies), result.AssembliesCount)
	}
}

// TestAssemblyData tests the parsing of AssemblyData objects.
func TestAssemblyData(t *testing.T) {
	var result AssemblyData
	err := json.NewDecoder(assemblyDataBody).Decode(&result)

	if err != nil {
		t.Errorf("Error decoding JSON: %s", err)
	}

	if result.MemberID != "1" {
		t.Errorf("Received invalid member ID: %s", result.MemberID)
	}

	if result.Model != "AssemblyOne" {
		t.Errorf("Received invalid model: %s", result.Model)
	}

	if result.Name != "Assembly One" {
		t.Errorf("Received invalid name: %s", result.Name)
	}

	if result.BinaryDataURI != "/image/path" {
		t.Errorf("Received invalid binary data URI: %s", result.BinaryDataURI)
	}

	if result.Status.State != "Enabled" {
		t.Errorf("Received invalid status state: %s", result.Status.State)
	}
}
