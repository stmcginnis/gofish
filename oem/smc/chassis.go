//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/schemas"
)

// Chassis is a Supermicro OEM instance of a Chassis.
type Chassis struct {
	schemas.Chassis
	BoardSerialNumber string
	GUID              string
	BoardID           string
}

// FromChassis converts a standard Chassis object to the OEM implementation.
func FromChassis(chassis *schemas.Chassis) (*Chassis, error) {
	cs := Chassis{
		Chassis: *chassis,
	}

	var t struct {
		Oem struct {
			Supermicro struct {
				BoardSerialNumber string `json:"BoardSerialNumber"`
				GUID              string `json:"GUID"`
				BoardID           string `json:"BoardID"`
			} `json:"Supermicro"`
		} `json:"Oem"`
	}

	err := json.Unmarshal(chassis.RawData, &t)
	if err != nil {
		return nil, err
	}

	cs.BoardSerialNumber = t.Oem.Supermicro.BoardSerialNumber
	cs.GUID = t.Oem.Supermicro.GUID
	cs.BoardID = t.Oem.Supermicro.BoardID

	cs.SetClient(chassis.GetClient())
	return &cs, nil
}
