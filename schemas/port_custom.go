//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a GenZ object from the raw JSON.
func (g *PortGenZ) UnmarshalJSON(b []byte) error {
	type temp PortGenZ
	var tmp struct {
		temp
		LPRT Link `json:"LPRT"`
		MPRT Link `json:"MPRT"`
		VCAT Link `json:"VCAT"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*g = PortGenZ(tmp.temp)

	// Extract the links to other entities for later
	g.lPRT = tmp.LPRT.String()
	g.mPRT = tmp.MPRT.String()
	g.vCAT = tmp.VCAT.String()

	return nil
}

// LPRT gets the LPRT collection.
func (g *PortGenZ) LPRT(client Client) ([]*RouteEntry, error) {
	if g.lPRT == "" {
		return nil, nil
	}
	return GetCollectionObjects[RouteEntry](client, g.lPRT)
}

// MPRT gets the MPRT collection.
func (g *PortGenZ) MPRT(client Client) ([]*RouteEntry, error) {
	if g.mPRT == "" {
		return nil, nil
	}
	return GetCollectionObjects[RouteEntry](client, g.mPRT)
}

// VCAT gets the VCAT collection.
func (g *PortGenZ) VCAT(client Client) ([]*VCATEntry, error) {
	if g.vCAT == "" {
		return nil, nil
	}
	return GetCollectionObjects[VCATEntry](client, g.vCAT)
}
