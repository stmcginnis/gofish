//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// UnmarshalJSON unmarshals a Certificate object from the raw JSON.
func (c *Certificate) UnmarshalJSON(b []byte) error {
	type temp Certificate
	type cActions struct {
		ForceAutomaticRenew ActionTarget `json:"#Certificate.ForceAutomaticRenew"`
		Rekey               ActionTarget `json:"#Certificate.Rekey"`
		Renew               ActionTarget `json:"#Certificate.Renew"`
	}
	type cLinks struct {
		Issuer   Link  `json:"Issuer"`
		Subjects Links `json:"Subjects"`
	}
	var tmp struct {
		temp
		Actions cActions
		Links   cLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = Certificate(tmp.temp)

	// Extract the links to other entities for later
	c.forceAutomaticRenewTarget = tmp.Actions.ForceAutomaticRenew.Target
	c.rekeyTarget = tmp.Actions.Rekey.Target
	c.renewTarget = tmp.Actions.Renew.Target
	c.issuer = tmp.Links.Issuer.String()
	c.subjects = tmp.Links.Subjects.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}
