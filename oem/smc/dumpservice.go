//
// SPDX-License-Identifier: BSD-3-Clause
//

package smc

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/coreweave/gofish/common"
)

// Dump represents a dump from the DumpService.
// NOTE: This is another one where the jsonschema reported by SMC appears to be
// wildly inaccurate. Use with caution.
type Dump struct {
	common.Entity

	AttestationFile []string
}

// GetDump will get a Dump instance from the service.
func GetDump(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Dump, error) {
	return GetDumpWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetDumpWithContext will get a Dump instance from the service.
func GetDumpWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*Dump, error) {
	return common.GetObjectWithContext[Dump](ctx, c, uri, queryOpts...)
}

// ListReferencedDumps gets the collection of Dumps from
// a provided reference.
func ListReferencedDumps(c common.Client, uri string, queryOpts ...common.QueryGroupOption) ([]*Dump, error) {
	return ListReferencedDumpsWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// ListReferencedDumpsWithContext gets the collection of Dumps from
// a provided reference.
func ListReferencedDumpsWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) ([]*Dump, error) {
	return common.GetCollectionObjectsWithContext[Dump](ctx, c, uri, queryOpts...)
}

// DumpService is the dump service instance associated with the system.
type DumpService struct {
	common.Entity

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
		Dumps   common.Link
		Actions struct {
			CreateDump common.ActionTarget `json:"#SmcDumpService.CreateDump"`
			DeleteAll  common.ActionTarget `json:"#SmcDumpService.DeleteAll"`
			Collect    common.ActionTarget `json:"#OemDumpService.Collect"`
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
func GetDefaultDumpService(c common.Client, queryOpts ...common.QueryGroupOption) (*DumpService, error) {
	return GetDefaultDumpServiceWithContext(common.ContextOf(c), c, queryOpts...)
}

// GetDefaultDumpServiceWithContext will get the default DumpService instance from the service.
func GetDefaultDumpServiceWithContext(ctx context.Context, c common.Client, queryOpts ...common.QueryGroupOption) (*DumpService, error) {
	return common.GetObjectWithContext[DumpService](ctx, c, "/redfish/v1/Oem/Supermicro/DumpService/", queryOpts...)
}

// GetDumpService will get a DumpService instance from the service.
func GetDumpService(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*DumpService, error) {
	return GetDumpServiceWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetDumpServiceWithContext will get a DumpService instance from the service.
func GetDumpServiceWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*DumpService, error) {
	return common.GetObjectWithContext[DumpService](ctx, c, uri, queryOpts...)
}

// CreateDump creates a new dump. Allowable dumpType is usually only
// "Host Dump".
func (ds *DumpService) CreateDump(dumpType string) error {
	return ds.CreateDumpWithContext(common.ContextOf(ds.GetClient()), dumpType)
}

// CreateDumpWithContext creates a new dump. Allowable dumpType is usually only
// "Host Dump".
func (ds *DumpService) CreateDumpWithContext(ctx context.Context, dumpType string) error {
	if ds.createDumpTarget == "" {
		return errors.New("create dump is not supported by this system")
	}

	return ds.PostWithContext(ctx, ds.createDumpTarget, map[string]any{
		"DumpType": dumpType,
	})
}

// DeleteAll deletes all dumps.
func (ds *DumpService) DeleteAll() error {
	return ds.DeleteAllWithContext(common.ContextOf(ds.GetClient()))
}

// DeleteAllWithContext deletes all dumps.
func (ds *DumpService) DeleteAllWithContext(ctx context.Context) error {
	if ds.deleteAllTarget == "" {
		return errors.New("delete all is not supported by this system")
	}

	return ds.PostWithContext(ctx, ds.deleteAllTarget, nil)
}

// Collect collects a dump.
// dumptType is usually only "HGXLogDump".
// actionType is usually one of "Create", "Delete", "Download", or "Query".
func (ds *DumpService) Collect(dumpType, actionType string) error {
	return ds.CollectWithContext(common.ContextOf(ds.GetClient()), dumpType, actionType)
}

// CollectWithContext collects a dump.
// dumptType is usually only "HGXLogDump".
// actionType is usually one of "Create", "Delete", "Download", or "Query".
func (ds *DumpService) CollectWithContext(ctx context.Context, dumpType, actionType string) error {
	if ds.collectTarget == "" {
		return errors.New("collect is not supported by this system")
	}

	return ds.PostWithContext(ctx, ds.collectTarget, map[string]any{
		"DumpType":   dumpType,
		"ActionType": actionType,
	})
}

// Dumps will get the Dumps from the service.
func (ds *DumpService) Dumps(queryOpts ...common.QueryGroupOption) ([]*Dump, error) {
	return ds.DumpsWithContext(common.ContextOf(ds.GetClient()), queryOpts...)
}

// DumpsWithContext will get the Dumps from the service.
func (ds *DumpService) DumpsWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Dump, error) {
	return ListReferencedDumpsWithContext(ctx, ds.GetClient(), ds.dumps, queryOpts...)
}
