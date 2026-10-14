//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a FabricAdapterGenZ object from the raw JSON.
func (g *FabricAdapterGenZ) UnmarshalJSON(b []byte) error {
	type temp FabricAdapterGenZ
	var tmp struct {
		temp
		MSDT          Link `json:"MSDT"`
		RequestorVCAT Link `json:"RequestorVCAT"`
		ResponderVCAT Link `json:"ResponderVCAT"`
		SSDT          Link `json:"SSDT"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*g = FabricAdapterGenZ(tmp.temp)

	// Extract the links to other entities for later
	g.mSDT = tmp.MSDT.String()
	g.requestorVCAT = tmp.RequestorVCAT.String()
	g.responderVCAT = tmp.ResponderVCAT.String()
	g.sSDT = tmp.SSDT.String()

	return nil
}

// MSDT gets the MSDT collection.
func (g *FabricAdapterGenZ) MSDT(client Client) ([]*RouteEntry, error) {
	if g.mSDT == "" {
		return nil, nil
	}
	return GetCollectionObjects[RouteEntry](client, g.mSDT)
}

// RequestorVCAT gets the RequestorVCAT collection.
func (g *FabricAdapterGenZ) RequestorVCAT(client Client) ([]*VCATEntry, error) {
	if g.requestorVCAT == "" {
		return nil, nil
	}
	return GetCollectionObjects[VCATEntry](client, g.requestorVCAT)
}

// ResponderVCAT gets the ResponderVCAT collection.
func (g *FabricAdapterGenZ) ResponderVCAT(client Client) ([]*VCATEntry, error) {
	if g.responderVCAT == "" {
		return nil, nil
	}
	return GetCollectionObjects[VCATEntry](client, g.responderVCAT)
}

// SSDT gets the SSDT collection.
func (g *FabricAdapterGenZ) SSDT(client Client) ([]*RouteEntry, error) {
	if g.sSDT == "" {
		return nil, nil
	}
	return GetCollectionObjects[RouteEntry](client, g.sSDT)
}
