//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"encoding/json"
	"errors"

	"github.com/stmcginnis/gofish/schemas"
)

// Dump represents a dump from the DumpService.
// NOTE: This is another one where the jsonschema reported by SMC appears to be
// wildly inaccurate. Use with caution.
type Dump struct {
	schemas.Entity

	AttestationFile []string
}

// GetDump will get a Dump instance from the service.
func GetDump(c schemas.Client, uri string) (*Dump, error) {
	return schemas.GetObject[Dump](c, uri)
}

// ListReferencedDumps gets the collection of Dumps from
// a provided reference.
func ListReferencedDumps(c schemas.Client, uri string) ([]*Dump, error) {
	return schemas.GetCollectionObjects[Dump](c, uri)
}

// DumpService is the dump service instance associated with the system.
type DumpService struct {
	schemas.Entity

	// Link to a DumpCollection.
	dumps string

	createDumpTarget string
	deleteAllTarget  string
	collectTarget    string
}

// UnmarshalJSON unmarshals an DumpService object from the raw JSON.
func (ds *DumpService) UnmarshalJSON(b []byte) error {
	type temp DumpService
	var t struct {
		temp
		Dumps   schemas.Link
		Actions struct {
			CreateDump schemas.ActionTarget `json:"#SmcDumpService.CreateDump"`
			DeleteAll  schemas.ActionTarget `json:"#SmcDumpService.DeleteAll"`
			Collect    schemas.ActionTarget `json:"#OemDumpService.Collect"`
		}
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*ds = DumpService(t.temp)

	ds.dumps = t.Dumps.String()

	ds.createDumpTarget = t.Actions.CreateDump.Target
	ds.deleteAllTarget = t.Actions.DeleteAll.Target
	ds.collectTarget = t.Actions.Collect.Target

	return nil
}

// GetDefaultDumpService will get the default DumpService instance from the service.
func GetDefaultDumpService(c schemas.Client) (*DumpService, error) {
	return schemas.GetObject[DumpService](c, "/redfish/v1/Oem/Supermicro/DumpService/")
}

// GetDumpService will get a DumpService instance from the service.
func GetDumpService(c schemas.Client, uri string) (*DumpService, error) {
	return schemas.GetObject[DumpService](c, uri)
}

// CreateDump creates a new dump. Allowable dumpType is usually only
// "Host Dump".
func (ds *DumpService) CreateDump(dumpType string) error {
	if ds.createDumpTarget == "" {
		return errors.New("create dump is not supported by this system")
	}

	return ds.Post(ds.createDumpTarget, map[string]any{
		"DumpType": dumpType,
	})
}

// DeleteAll deletes all dumps.
func (ds *DumpService) DeleteAll() error {
	if ds.deleteAllTarget == "" {
		return errors.New("delete all is not supported by this system")
	}

	return ds.Post(ds.deleteAllTarget, nil)
}

// Collect collects a dump.
// dumptType is usually only "HGXLogDump".
// actionType is usually one of "Create", "Delete", "Download", or "Query".
func (ds *DumpService) Collect(dumpType, actionType string) error {
	if ds.collectTarget == "" {
		return errors.New("collect is not supported by this system")
	}

	return ds.Post(ds.collectTarget, map[string]any{
		"DumpType":   dumpType,
		"ActionType": actionType,
	})
}

// Dumps will get the Dumps from the service.
func (ds *DumpService) Dumps() ([]*Dump, error) {
	return ListReferencedDumps(ds.GetClient(), ds.dumps)
}
