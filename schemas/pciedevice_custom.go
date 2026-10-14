//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

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

// postUnmarshal records Links.PCIeFunctions, used when the service has no
// PCIeFunctions collection.
func (p *PCIeDevice) postUnmarshal(b []byte) error {
	var t struct {
		Links struct {
			PCIeFunctions Links
		}
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	p.pCIeFunctionsLink = t.Links.PCIeFunctions.ToStrings()

	return nil
}
