//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

import "encoding/json"

// PhysicalPortAssignment gets the PhysicalPortAssignment linked resource.
func (n *NetworkDeviceFunction) PhysicalPortAssignment() (*NetworkPort, error) {
	if n.physicalPortAssignment == "" {
		return nil, nil
	}
	return GetObject[NetworkPort](n.client, n.physicalPortAssignment)
}

// postUnmarshal records the physical port assignment, which services report
// either in Links or as the deprecated top-level property.
func (n *NetworkDeviceFunction) postUnmarshal(b []byte) error {
	var t struct {
		Links struct {
			PhysicalPortAssignment Link
		}
		PhysicalPortAssignment Link
	}
	if err := json.Unmarshal(b, &t); err != nil {
		return err
	}

	n.physicalPortAssignment = t.Links.PhysicalPortAssignment.String()
	if n.physicalPortAssignment == "" {
		n.physicalPortAssignment = t.PhysicalPortAssignment.String()
	}

	return nil
}
