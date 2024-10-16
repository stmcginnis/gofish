//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
	"github.com/stmcginnis/gofish/redfish"
)

// LLDP is an instance of a LLDP object.
type LLDP struct {
	common.Entity

	// Enabled shall contain the state indicating whether to enable LLDP for a port.
	// If LLDP is disabled at the adapter level, this property shall be ignored.
	Enabled bool `json:"LLDPEnabled"`
	// LLDPReceive shall contain the LLDP data being received on this link.
	LLDPReceive redfish.LLDPReceive
	// LLDPTransmit shall contain the LLDP data being transmit on this link.
	LLDPTransmit redfish.LLDPTransmit

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a LLDP object from the raw JSON.
func (i *LLDP) UnmarshalJSON(b []byte) error {
	type temp LLDP
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*i = LLDP(t.temp)

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *LLDP) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	orig := new(LLDP)
	err := orig.UnmarshalJSON(i.RawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"Enabled",
	}

	originalElement := reflect.ValueOf(orig).Elem()
	currentElement := reflect.ValueOf(i).Elem()

	return i.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetLLDP will get a LLDP instance from the service.
func GetLLDP(c common.Client, uri string) (*LLDP, error) {
	return common.GetObject[LLDP](c, uri)
}
