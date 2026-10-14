//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a PCIeDevice object from the raw JSON.
func (p *PCIeDevice) UnmarshalJSON(b []byte) error {
	type temp PCIeDevice
	type pActions struct {
		CXLReferenceDCMemory ActionTarget `json:"#PCIeDevice.CXLReferenceDCMemory"`
	}
	type pLinks struct {
		Chassis            Links `json:"Chassis"`
		ConnectedPCIePorts Links `json:"ConnectedPCIePorts"`
		PCIeFunctions      Links `json:"PCIeFunctions"`
		FabricAdapters     Links `json:"FabricAdapters"`
		Processors         Links `json:"Processors"`
		Switch             Link  `json:"Switch"`
	}
	var tmp struct {
		temp
		Actions            pActions
		Links              pLinks
		Assembly           ResolvedLink[Assembly]           `json:"Assembly"`
		CXLLogicalDevices  Link                             `json:"CXLLogicalDevices"`
		Certificates       Link                             `json:"Certificates"`
		EnvironmentMetrics ResolvedLink[EnvironmentMetrics] `json:"EnvironmentMetrics"`
		PCIeFunctions      Link                             `json:"PCIeFunctions"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*p = PCIeDevice(tmp.temp)

	// Extract the links to other entities for later
	p.cXLReferenceDCMemoryTarget = tmp.Actions.CXLReferenceDCMemory.Target
	p.chassis = tmp.Links.Chassis.ToStrings()
	p.connectedPCIePorts = tmp.Links.ConnectedPCIePorts.ToStrings()
	p.pCIeFunctionsLink = tmp.Links.PCIeFunctions.ToStrings()
	p.fabricAdapters = tmp.Links.FabricAdapters.ToStrings()
	p.processors = tmp.Links.Processors.ToStrings()
	p.switch_ = tmp.Links.Switch.String()
	p.assembly = tmp.Assembly

	p.cXLLogicalDevices = tmp.CXLLogicalDevices.String()
	p.certificates = tmp.Certificates.String()
	p.environmentMetrics = tmp.EnvironmentMetrics

	p.pCIeFunctions = tmp.PCIeFunctions.String()

	// This is a read/write object, so we need to save the raw object data for later
	p.RawData = b

	return nil
}

// PCIeFunctions gets the PCIeFunctions collection.
func (p *PCIeDevice) PCIeFunctions() ([]*PCIeFunction, error) {
	if len(p.pCIeFunctionsLink) > 0 {
		return GetObjects[PCIeFunction](p.client, p.pCIeFunctionsLink)
	}

	if p.pCIeFunctions == "" {
		return nil, nil
	}
	return GetCollectionObjects[PCIeFunction](p.client, p.pCIeFunctions)
}

func (p *PCIeInterface) UnmarshalJSON(b []byte) error {
	// If it's an empty array, ignore zero value.
	if string(b) == "[]" {
		return nil
	}

	type temp PCIeInterface
	var t struct {
		temp

		// Workaround for vendors returning LanesInUse/MaxLanes as strings.
		LanesInUse any
		MaxLanes   any
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*p = PCIeInterface(t.temp)
	p.LanesInUse = toInt(t.LanesInUse)
	p.MaxLanes = toInt(t.MaxLanes)
	return nil
}
