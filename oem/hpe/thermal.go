//
// SPDX-License-Identifier: BSD-3-Clause
//

package hpe

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/redfish"
)

type Fan struct {
	redfish.ThermalFan
	Oem FanOem
}

type FanOem struct {
	Hpe struct {
		OdataContext string `json:"@odata.context"`
		OdataType    string `json:"@odata.type"`
		Location     string `json:"Location"`
		Redundant    bool   `json:"Redundant"`
		HotPluggable bool   `json:"HotPluggable"`
	} `json:"Hpe"`
}

type Thermal struct {
	redfish.Thermal
	Fans []Fan
	Oem  ThermalOem
}

type ThermalOem struct {
	Hpe struct {
		OdataContext         string `json:"@odata.context"`
		OdataType            string `json:"@odata.type"`
		ThermalConfiguration string `json:"ThermalConfiguration"`
		FanPercentMinimum    int    `json:"FanPercentMinimum"`
	} `json:"Hpe"`
}

func FromThermal(thermal *redfish.Thermal) (Thermal, error) {
	oem := ThermalOem{}

	_ = json.Unmarshal(thermal.Oem, &oem)

	fans := make([]Fan, 0, len(thermal.Fans))

	for i := range thermal.Fans {
		fan, err := FromFan(&thermal.Fans[i])
		if err != nil {
			return Thermal{}, err
		}

		fans = append(fans, fan)
	}

	return Thermal{
		Thermal: *thermal,
		Fans:    fans,
		Oem:     oem,
	}, nil
}

func FromFan(fan *redfish.ThermalFan) (Fan, error) {
	oem := FanOem{}

	_ = json.Unmarshal(fan.OEM, &oem)

	return Fan{
		ThermalFan: *fan,
		Oem:        oem,
	}, nil
}
