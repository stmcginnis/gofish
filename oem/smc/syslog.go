//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/schemas"
)

type SyslogServer struct {
	ServerIP string `json:"ServerIP"`
	Port     *int   `json:"Port"`
}

// Syslog is an instance of a Syslog object.
type Syslog struct {
	schemas.Entity

	Enabled bool   `json:"EnableSyslog"`
	Server  string `json:"SyslogServer"`
	Port    int    `json:"SyslogPortNumber"`

	// Servers is populated for both formats. On firmware that reports a
	// singular server it holds the values from the first entry.
	Servers []SyslogServer `json:"-"`

	// serverCollection records whether the BMC reported SyslogServer as a
	// collection, which determines the shape of the update payload.
	serverCollection bool

	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a Syslog object from the raw JSON.
func (i *Syslog) UnmarshalJSON(b []byte) error {
	type temp Syslog
	var t struct {
		temp
		SyslogServer json.RawMessage
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*i = Syslog(t.temp)

	// Starting with BMC firmware versions Gen 13 1.10 and Gen 14 1.08 the
	// SyslogServer changed from a singular string value to an array of
	// SyslogServer entries.
	if len(t.SyslogServer) != 0 {
		var servers []SyslogServer
		if err := json.Unmarshal(t.SyslogServer, &servers); err == nil {
			i.Servers = servers
			i.serverCollection = true

			// Capture first server for backwards compatibility
			if len(servers) > 0 {
				i.Server = servers[0].ServerIP
				i.Port = 0
				if servers[0].Port != nil {
					i.Port = *servers[0].Port
				}
			}
		} else {
			// Not a collection, fall back to the singular string value.
			if err := json.Unmarshal(t.SyslogServer, &i.Server); err != nil {
				return err
			}

			server := SyslogServer{ServerIP: i.Server}
			if i.Port != 0 {
				server.Port = &i.Port
			}
			i.Servers = []SyslogServer{server}
		}
	}

	// This is a read/write object, so we need to save the raw object data for later
	i.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (i *Syslog) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	orig := new(Syslog)
	err := orig.UnmarshalJSON(i.RawData)
	if err != nil {
		return err
	}

	if i.serverCollection {
		return i.updateServers(orig)
	}

	readWriteFields := []string{
		"Enabled",
		"EnableSyslog",
		"Port",
		"SyslogPortNumber",
		"Server",
		"SyslogServer",
	}

	originalElement := reflect.ValueOf(orig).Elem()
	currentElement := reflect.ValueOf(i).Elem()

	return i.Entity.Update(originalElement, currentElement, readWriteFields)
}

// updateServers sends updates to firmware that reports SyslogServer as a
// collection. Server and Port mirror the first collection entry, so a change
// to either is folded into that entry before the collection is sent.
func (i *Syslog) updateServers(orig *Syslog) error {
	servers := make([]SyslogServer, len(i.Servers))
	copy(servers, i.Servers)

	if len(servers) > 0 {
		if i.Server != orig.Server {
			servers[0].ServerIP = i.Server
		}
		if i.Port != orig.Port {
			port := i.Port
			servers[0].Port = &port
		}
	}

	payload := map[string]any{}
	if i.Enabled != orig.Enabled {
		payload["EnableSyslog"] = i.Enabled
	}
	if !reflect.DeepEqual(servers, orig.Servers) {
		payload["SyslogServer"] = servers
	}

	if len(payload) == 0 {
		return nil
	}

	return i.Patch(i.ODataID, payload)
}

// GetSyslog will get a Syslog instance from the service.
func GetSyslog(c schemas.Client, uri string) (*Syslog, error) {
	return schemas.GetObject[Syslog](c, uri)
}
