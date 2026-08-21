//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"encoding/xml"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stmcginnis/gofish/schemas"
)

const updateListxml = "<?xml version=\"1.0\"?>\n<CIM xmlns:fo=\"http://www.w3.org/1999/XSL/Format\" CIMVERSION=\"2.0\" DTDVERSION=\"2.0\">\n  <MESSAGE ID=\"4711\" PROTOCOLVERSION=\"1.0\">\n    <SIMPLEREQ>\n      <VALUE.NAMEDINSTANCE>\n        <INSTANCENAME CLASSNAME=\"DCIM_RepoUpdateSWID\"><PROPERTY NAME=\"Criticality\" TYPE=\"string\"><VALUE>3</VALUE></PROPERTY><PROPERTY NAME=\"DisplayName\" TYPE=\"string\"><VALUE>Mellanox ConnectX-6 Single Port VPI HDR100 QSFP Adapter - B8:CE:F6:71:A3:D0</VALUE></PROPERTY><PROPERTY NAME=\"BaseLocation\" TYPE=\"string\"><VALUE/></PROPERTY><PROPERTY NAME=\"PackagePath\" TYPE=\"string\"><VALUE>FOLDER09614044M/2/Network_Firmware_3VFTR_WN64_20.36.10.10.EXE</VALUE></PROPERTY><PROPERTY NAME=\"PackageName\" TYPE=\"string\"><VALUE>Network_Firmware_3VFTR_WN64_20.36.10.10.EXE</VALUE></PROPERTY><PROPERTY NAME=\"PackageVersion\" TYPE=\"string\"><VALUE>20.36.10.10</VALUE></PROPERTY><PROPERTY NAME=\"RebootType\" TYPE=\"string\"><VALUE>HOST</VALUE></PROPERTY><PROPERTY NAME=\"JobID\" TYPE=\"string\"><VALUE/></PROPERTY>\n    <PROPERTY NAME=\"Target\" TYPE=\"string\"><VALUE>DCIM:INSTALLED#781__InfiniBand.Slot.6-1</VALUE></PROPERTY><PROPERTY NAME=\"ComponentID\" TYPE=\"string\"><VALUE>105399</VALUE></PROPERTY><PROPERTY NAME=\"ComponentType\" TYPE=\"string\"><VALUE>FRMW</VALUE></PROPERTY><PROPERTY.ARRAY NAME=\"ComponentInfoValue\" TYPE=\"string\"><VALUE.ARRAY><VALUE>15B3:101B:15B3:0018</VALUE></VALUE.ARRAY></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInfoName\" TYPE=\"string\"><VALUE.ARRAY><VALUE>VendorID:DeviceID:SubVendorID:SubDeviceID</VALUE></VALUE.ARRAY></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInfoTarget\" TYPE=\"string\"><VALUE.ARRAY><VALUE>DCIM:INSTALLED#781__InfiniBand.Slot.6-1</VALUE></VALUE.ARRAY></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInstalledVersion\" TYPE=\"string\"><VALUE.ARRAY><VALUE>20.35.10.12</VALUE></VALUE.ARRAY></PROPERTY.ARRAY>\n    \n  </INSTANCENAME>\n      </VALUE.NAMEDINSTANCE>\n      <VALUE.NAMEDINSTANCE>\n        <INSTANCENAME CLASSNAME=\"DCIM_RepoUpdateSWID\"><PROPERTY NAME=\"Criticality\" TYPE=\"string\"><VALUE>2</VALUE></PROPERTY><PROPERTY NAME=\"DisplayName\" TYPE=\"string\"><VALUE>BIOS</VALUE></PROPERTY><PROPERTY NAME=\"BaseLocation\" TYPE=\"string\"><VALUE/></PROPERTY><PROPERTY NAME=\"PackagePath\" TYPE=\"string\"><VALUE>FOLDER09680475M/5/BIOS_YNJ0T_WN64_1.10.2_02.EXE</VALUE></PROPERTY><PROPERTY NAME=\"PackageName\" TYPE=\"string\"><VALUE>BIOS_YNJ0T_WN64_1.10.2_02.EXE</VALUE></PROPERTY><PROPERTY NAME=\"PackageVersion\" TYPE=\"string\"><VALUE>1.10.2</VALUE></PROPERTY><PROPERTY NAME=\"RebootType\" TYPE=\"string\"><VALUE>HOST</VALUE></PROPERTY><PROPERTY NAME=\"JobID\" TYPE=\"string\"><VALUE/></PROPERTY>\n    <PROPERTY NAME=\"Target\" TYPE=\"string\"><VALUE>DCIM:INSTALLED#741__BIOS.Setup.1-1</VALUE></PROPERTY><PROPERTY NAME=\"ComponentID\" TYPE=\"string\"><VALUE>159</VALUE></PROPERTY><PROPERTY NAME=\"ComponentType\" TYPE=\"string\"><VALUE>BIOS</VALUE></PROPERTY><PROPERTY.ARRAY NAME=\"ComponentInfoValue\" TYPE=\"string\"><VALUE.ARRAY/></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInfoName\" TYPE=\"string\"><VALUE.ARRAY/></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInfoTarget\" TYPE=\"string\"><VALUE.ARRAY><VALUE>DCIM:INSTALLED#741__BIOS.Setup.1-1</VALUE></VALUE.ARRAY></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInstalledVersion\" TYPE=\"string\"><VALUE.ARRAY><VALUE>1.9.2</VALUE></VALUE.ARRAY></PROPERTY.ARRAY>\n    \n  </INSTANCENAME>\n      </VALUE.NAMEDINSTANCE>\n      <VALUE.NAMEDINSTANCE>\n        <INSTANCENAME CLASSNAME=\"DCIM_RepoUpdateSWID\"><PROPERTY NAME=\"Criticality\" TYPE=\"string\"><VALUE>1</VALUE></PROPERTY><PROPERTY NAME=\"DisplayName\" TYPE=\"string\"><VALUE>Intel(R) Ethernet 10G 2P X710 OCP - 68:05:CA:D5:BF:62</VALUE></PROPERTY><PROPERTY NAME=\"BaseLocation\" TYPE=\"string\"><VALUE/></PROPERTY><PROPERTY NAME=\"PackagePath\" TYPE=\"string\"><VALUE>FOLDER09683859M/1/Network_Firmware_9NPPG_WN64_22.0.9_A00.EXE</VALUE></PROPERTY><PROPERTY NAME=\"PackageName\" TYPE=\"string\"><VALUE>Network_Firmware_9NPPG_WN64_22.0.9_A00.EXE</VALUE></PROPERTY><PROPERTY NAME=\"PackageVersion\" TYPE=\"string\"><VALUE>22.0.9</VALUE></PROPERTY><PROPERTY NAME=\"RebootType\" TYPE=\"string\"><VALUE>HOST</VALUE></PROPERTY><PROPERTY NAME=\"JobID\" TYPE=\"string\"><VALUE/></PROPERTY>\n    <PROPERTY NAME=\"Target\" TYPE=\"string\"><VALUE>DCIM:INSTALLED#701__NIC.Integrated.1-1-1</VALUE></PROPERTY><PROPERTY NAME=\"ComponentID\" TYPE=\"string\"><VALUE>108122</VALUE></PROPERTY><PROPERTY NAME=\"ComponentType\" TYPE=\"string\"><VALUE>FRMW</VALUE></PROPERTY><PROPERTY.ARRAY NAME=\"ComponentInfoValue\" TYPE=\"string\"><VALUE.ARRAY><VALUE>8086:1572:8086:0013</VALUE><VALUE>8086:1572:8086:0000</VALUE></VALUE.ARRAY></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInfoName\" TYPE=\"string\"><VALUE.ARRAY><VALUE>VendorID:DeviceID:SubVendorID:SubDeviceID</VALUE><VALUE>VendorID:DeviceID:SubVendorID:SubDeviceID</VALUE></VALUE.ARRAY></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInfoTarget\" TYPE=\"string\"><VALUE.ARRAY><VALUE>DCIM:INSTALLED#701__NIC.Integrated.1-1-1</VALUE><VALUE>DCIM:INSTALLED#701__NIC.Integrated.1-2-1</VALUE></VALUE.ARRAY></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInstalledVersion\" TYPE=\"string\"><VALUE.ARRAY><VALUE>21.5.9</VALUE><VALUE>21.5.9</VALUE></VALUE.ARRAY></PROPERTY.ARRAY>\n    \n  </INSTANCENAME>\n      </VALUE.NAMEDINSTANCE>\n      <VALUE.NAMEDINSTANCE>\n        <INSTANCENAME CLASSNAME=\"DCIM_RepoUpdateSWID\"><PROPERTY NAME=\"Criticality\" TYPE=\"string\"><VALUE>3</VALUE></PROPERTY><PROPERTY NAME=\"DisplayName\" TYPE=\"string\"><VALUE>Dell 64 Bit uEFI Diagnostics, version 4301, 4301A79, 4301.80</VALUE></PROPERTY><PROPERTY NAME=\"BaseLocation\" TYPE=\"string\"><VALUE/></PROPERTY><PROPERTY NAME=\"PackagePath\" TYPE=\"string\"><VALUE>FOLDER09610634M/1/Diagnostics_Application_WJXG4_WN64_4301A89_4301.90.EXE</VALUE></PROPERTY><PROPERTY NAME=\"PackageName\" TYPE=\"string\"><VALUE>Diagnostics_Application_WJXG4_WN64_4301A89_4301.90.EXE</VALUE></PROPERTY><PROPERTY NAME=\"PackageVersion\" TYPE=\"string\"><VALUE>4301A89</VALUE></PROPERTY><PROPERTY NAME=\"RebootType\" TYPE=\"string\"><VALUE>NONE</VALUE></PROPERTY><PROPERTY NAME=\"JobID\" TYPE=\"string\"><VALUE/></PROPERTY>\n    <PROPERTY NAME=\"Target\" TYPE=\"string\"><VALUE>DCIM:INSTALLED#802__Diagnostics.Embedded.1:LC.Embedded.1</VALUE></PROPERTY><PROPERTY NAME=\"ComponentID\" TYPE=\"string\"><VALUE>25806</VALUE></PROPERTY><PROPERTY NAME=\"ComponentType\" TYPE=\"string\"><VALUE>APAC</VALUE></PROPERTY><PROPERTY.ARRAY NAME=\"ComponentInfoValue\" TYPE=\"string\"><VALUE.ARRAY/></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInfoName\" TYPE=\"string\"><VALUE.ARRAY/></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInfoTarget\" TYPE=\"string\"><VALUE.ARRAY><VALUE>DCIM:INSTALLED#802__Diagnostics.Embedded.1:LC.Embedded.1</VALUE></VALUE.ARRAY></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInstalledVersion\" TYPE=\"string\"><VALUE.ARRAY><VALUE>4301A79</VALUE></VALUE.ARRAY></PROPERTY.ARRAY>\n    \n  </INSTANCENAME>\n      </VALUE.NAMEDINSTANCE>\n      <VALUE.NAMEDINSTANCE>\n        <INSTANCENAME CLASSNAME=\"DCIM_RepoUpdateSWID\"><PROPERTY NAME=\"Criticality\" TYPE=\"string\"><VALUE>1</VALUE></PROPERTY><PROPERTY NAME=\"DisplayName\" TYPE=\"string\"><VALUE>Integrated Dell Remote Access Controller</VALUE></PROPERTY><PROPERTY NAME=\"BaseLocation\" TYPE=\"string\"><VALUE/></PROPERTY><PROPERTY NAME=\"PackagePath\" TYPE=\"string\"><VALUE>FOLDER10032207M/1/iDRAC-with-Lifecycle-Controller_Firmware_Y0CWW_WN64_6.10.80.00_A00.EXE</VALUE></PROPERTY><PROPERTY NAME=\"PackageName\" TYPE=\"string\"><VALUE>iDRAC-with-Lifecycle-Controller_Firmware_Y0CWW_WN64_6.10.80.00_A00.EXE</VALUE></PROPERTY><PROPERTY NAME=\"PackageVersion\" TYPE=\"string\"><VALUE>6.10.80.00</VALUE></PROPERTY><PROPERTY NAME=\"RebootType\" TYPE=\"string\"><VALUE>IDRAC</VALUE></PROPERTY><PROPERTY NAME=\"JobID\" TYPE=\"string\"><VALUE/></PROPERTY>\n    <PROPERTY NAME=\"Target\" TYPE=\"string\"><VALUE>DCIM:INSTALLED#iDRAC.Embedded.1-1#IDRACinfo</VALUE></PROPERTY><PROPERTY NAME=\"ComponentID\" TYPE=\"string\"><VALUE>25227</VALUE></PROPERTY><PROPERTY NAME=\"ComponentType\" TYPE=\"string\"><VALUE>FRMW</VALUE></PROPERTY><PROPERTY.ARRAY NAME=\"ComponentInfoValue\" TYPE=\"string\"><VALUE.ARRAY/></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInfoName\" TYPE=\"string\"><VALUE.ARRAY/></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInfoTarget\" TYPE=\"string\"><VALUE.ARRAY><VALUE>DCIM:INSTALLED#iDRAC.Embedded.1-1#IDRACinfo</VALUE></VALUE.ARRAY></PROPERTY.ARRAY><PROPERTY.ARRAY NAME=\"ComponentInstalledVersion\" TYPE=\"string\"><VALUE.ARRAY><VALUE>6.10.30.20</VALUE></VALUE.ARRAY></PROPERTY.ARRAY>\n    \n  </INSTANCENAME>\n      </VALUE.NAMEDINSTANCE>\n    </SIMPLEREQ>\n  </MESSAGE>\n</CIM>\n"

func TestParseUpdateList(t *testing.T) {
	var list xmlUpdateList

	err := xml.Unmarshal([]byte(updateListxml), &list)
	if err != nil {
		t.Error(err)
	}

	ul := list.parseFromXML()

	if len(ul) != 5 {
		t.Errorf("Invalid update list length: %d", len(ul))
	}

	if ul[0].DisplayName != "Mellanox ConnectX-6 Single Port VPI HDR100 QSFP Adapter - B8:CE:F6:71:A3:D0" {
		t.Errorf("Invalid firmware display name: %s", ul[0].DisplayName)
	}

	if ul[3].PackageVersion != "4301A89" {
		t.Errorf("Invalid firmware package version: %s", ul[3].PackageVersion)
	}
}

func TestInstallFromRepository(t *testing.T) {
	const (
		actionURI = "/redfish/v1/Systems/System.Embedded.1/Oem/Dell/DellSoftwareInstallationService/Actions/DellSoftwareInstallationService.InstallFromRepository"
		jobURI    = "/redfish/v1/Managers/iDRAC.Embedded.1/Oem/Dell/Jobs/JID_456"
	)

	validBody := &InstallFromRepoBody{
		IPAddress: "192.0.2.1",
		ShareName: "/repository",
		ShareType: NFSIFRShareType,
	}

	t.Run("missing request body", func(t *testing.T) {
		client := &schemas.TestClient{}
		service := newSoftwareInstallationService(client, actionURI)

		_, err := service.InstallFromRepository(nil)
		if err == nil || !strings.Contains(err.Error(), "request body is required") {
			t.Fatalf("expected request body validation error, got %v", err)
		}
		if calls := client.CapturedCalls(); len(calls) != 0 {
			t.Fatalf("expected no API calls, got %d", len(calls))
		}
	})

	t.Run("missing Location header", func(t *testing.T) {
		client := &schemas.TestClient{CustomReturnForActions: map[string][]any{
			http.MethodPost: {
				&http.Response{StatusCode: http.StatusAccepted, Header: http.Header{}, Body: io.NopCloser(strings.NewReader(""))},
			},
		}}
		service := newSoftwareInstallationService(client, actionURI)

		_, err := service.InstallFromRepository(validBody)
		if err == nil || !strings.Contains(err.Error(), "missing Location header") {
			t.Fatalf("expected missing Location header error, got %v", err)
		}
		if calls := client.CapturedCalls(); len(calls) != 1 || calls[0].Action != http.MethodPost {
			t.Fatalf("expected only the action POST, got %#v", calls)
		}
	})

	t.Run("follows job Location", func(t *testing.T) {
		client := &schemas.TestClient{CustomReturnForActions: map[string][]any{
			http.MethodPost: {
				&http.Response{StatusCode: http.StatusAccepted, Header: http.Header{"Location": []string{jobURI}}, Body: io.NopCloser(strings.NewReader(""))},
			},
			http.MethodGet: {
				&http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`{"Id":"JID_456","JobState":"Running"}`))},
			},
		}}
		service := newSoftwareInstallationService(client, actionURI)

		job, err := service.InstallFromRepository(validBody)
		if err != nil {
			t.Fatalf("unexpected install error: %v", err)
		}
		if job.ID != "JID_456" || job.JobState != "Running" {
			t.Fatalf("unexpected job: %#v", job)
		}
		calls := client.CapturedCalls()
		if len(calls) != 2 || calls[0].URL != actionURI || calls[1].URL != jobURI {
			t.Fatalf("unexpected API calls: %#v", calls)
		}
	})
}

func newSoftwareInstallationService(client schemas.Client, actionURI string) *SoftwareInstallationService {
	service := &SoftwareInstallationService{}
	service.SetClient(client)
	service.Actions.InstallFromRepository.Target = actionURI
	return service
}
