//
// SPDX-License-Identifier: BSD-3-Clause
//

package supermicro

import (
	"encoding/json"
	"strings"
	"testing"
)

var supermicroIKVMBodyHTML = `{
    "@odata.type": "#IKVM.v1_0_1.IKVM",
    "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/IKVM",
    "Id": "IKVM",
    "Name": "IKVM",
    "Current interface": "HTML 5",
    "URI": "/redfish/ohkei8Quei5odeem.IKVM"
}`

var supermicroIKVMBodyJAVA = `{
    "@odata.type": "#IKVM.v1_0_1.IKVM",
    "@odata.id": "/redfish/v1/Managers/1/Oem/Supermicro/IKVM",
    "Id": "IKVM",
    "Name": "IKVM",
    "Current interface": "JAVA plug-in"
}`

func TestSupermicroIKVM(t *testing.T) {
	var ikvm *IKVM
	err := json.NewDecoder(strings.NewReader(supermicroIKVMBodyJAVA)).Decode(&ikvm)

	if err != nil {
		t.Errorf("Error decoding IKVM-JAVA JSON: %v", err)
	}

	if ikvm.CurrentInterface != JavaIKVMInterface {
		t.Errorf("Expected JAVA KVM interface, got: %s", ikvm.CurrentInterface)
	}

	if ikvm.URI != "" {
		t.Errorf("JAVA KVM interface should have an empty URI, got %s", ikvm.URI)
	}

	err = json.NewDecoder(strings.NewReader(supermicroIKVMBodyHTML)).Decode(&ikvm)
	if err != nil {
		t.Errorf("Error decoding IKVM-HTML JSON: %v", err)
	}

	if ikvm.CurrentInterface != HTMLIKVMInterface {
		t.Errorf("Expected HTML KVM interface, got %s", ikvm.CurrentInterface)
	}

	if ikvm.URI != "/redfish/ohkei8Quei5odeem.IKVM" {
		t.Errorf("Expected HTML KVM interface URI \"/redfish/ohkei8Quei5odeem.IKVM\", got %s", ikvm.URI)
	}
}
