//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// postUnmarshal records Links.Issuer, which has no getter because the Issuer
// property owns that name.
func (c *Certificate) postUnmarshal(b []byte) error {
	var t struct {
		Links struct {
			Issuer Link
		}
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	c.issuer = t.Links.Issuer.String()

	return nil
}
