//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// postUnmarshal records the number of PCIe functions in Links.
func (c *CXLLogicalDevice) postUnmarshal(b []byte) error {
	var t struct {
		Links struct {
			PCIeFunctionsCount int `json:"PCIeFunctions@odata.count"`
		}
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	c.PCIeFunctionsCount = t.Links.PCIeFunctionsCount

	return nil
}
