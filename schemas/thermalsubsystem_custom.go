//
// SPDX-License-Identifier: BSD-3-Clause
//

package schemas

// Fans gets the Fans collection.
func (t *ThermalSubsystem) Fans() ([]*ThermalFan, error) {
	if t.fans == "" {
		return nil, nil
	}
	return GetCollectionObjects[ThermalFan](t.client, t.fans)
}
