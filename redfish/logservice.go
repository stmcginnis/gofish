//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// LogEntryTypes is the type of log entry.
type LogEntryTypes string

const (
	// EventLogEntryTypes contains Redfish-defined messages (events).
	EventLogEntryTypes LogEntryTypes = "Event"
	// SELLogEntryTypes contains legacy IPMI System Event Log (SEL) entries.
	SELLogEntryTypes LogEntryTypes = "SEL"
	// MultipleLogEntryTypes contains multiple Log Entry types or a
	// single entry type cannot be guaranteed by the Log Service.
	MultipleLogEntryTypes LogEntryTypes = "Multiple"
	// OEMLogEntryTypes contains entries in an OEM-defined format.
	OEMLogEntryTypes LogEntryTypes = "OEM"
)

// OverWritePolicy is the log overwriting policy.
type OverWritePolicy string

const (

	// UnknownOverWritePolicy means the policy is not known or is undefined.
	UnknownOverWritePolicy OverWritePolicy = "Unknown"
	// WrapsWhenFullOverWritePolicy means when full, new entries to the Log will
	// overwrite previous entries.
	WrapsWhenFullOverWritePolicy OverWritePolicy = "WrapsWhenFull"
	// NeverOverWritesOverWritePolicy means when full, new entries to the Log will
	// be discarded.
	NeverOverWritesOverWritePolicy OverWritePolicy = "NeverOverWrites"
)

// LogService is used to represent a log service for a Redfish
// implementation.
type LogService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// DateTime shall represent the current DateTime value that the log service
	// is using, with offset from UTC, in Redfish Timestamp format.
	DateTime string
	// DateTimeLocalOffset shall represent the offset from UTC time that the
	// current value of DataTime property contains.
	DateTimeLocalOffset string
	// Description provides a description of this resource.
	Description string
	// Entries shall reference a collection of resources of type LogEntry.
	entries string
	// LogEntryType shall represent the
	// EntryType of all LogEntry resources contained in the Entries
	// collection. If a single EntryType for all LogEntry resources cannot
	// be determined or guaranteed by the Service, the value of this property
	// shall be 'Multiple'.
	LogEntryType LogEntryTypes
	// MaxNumberOfRecords shall be the maximum numbers of LogEntry resources in
	// the Entries collection for this service.
	MaxNumberOfRecords uint64
	// OverWritePolicy shall indicate the
	// policy of the log service when the MaxNumberOfRecords has been
	// reached. Unknown indicates the log overwrite policy is unknown.
	// WrapsWhenFull indicates that the log overwrites its entries with new
	// entries when the log has reached its maximum capacity. NeverOverwrites
	// indicates that the log never overwrites its entries by the new entries
	// and ceases logging when the limit has been reached.
	OverWritePolicy OverWritePolicy
	// ServiceEnabled shall be a boolean
	// indicating whether this service is enabled.
	ServiceEnabled bool
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// clearLogTarget is the URL to send ClearLog actions to.
	clearLogTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a LogService object from the raw JSON.
func (logservice *LogService) UnmarshalJSON(b []byte) error {
	type temp LogService
	type Actions struct {
		ClearLog struct {
			Target string
		} `json:"#LogService.ClearLog"`
	}
	var t struct {
		temp
		Entries common.Link
		Actions Actions
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*logservice = LogService(t.temp)
	logservice.entries = t.Entries.String()
	logservice.clearLogTarget = t.Actions.ClearLog.Target

	// This is a read/write object, so we need to save the raw object data for later
	logservice.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (logservice *LogService) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(LogService)
	err := original.UnmarshalJSON(logservice.rawData)
	if err != nil {
		return err
	}

	readWriteFields := []string{
		"DateTime",
		"DateTimeLocalOffset",
		"ServiceEnabled",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(logservice).Elem()

	return logservice.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetLogService will get a LogService instance from the service.
func GetLogService(c common.Client, uri string) (*LogService, error) {
	var logService LogService
	return &logService, logService.Get(c, uri, &logService)
}

// ListReferencedLogServices gets the collection of LogService from a provided reference.
func ListReferencedLogServices(c common.Client, link string) ([]*LogService, error) {
	var result []*LogService
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *LogService
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		logservice, err := GetLogService(c, link)
		ch <- GetResult{Item: logservice, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// Entries gets the log entries of this service.
func (logservice *LogService) Entries() ([]*LogEntry, error) {
	return ListReferencedLogEntrys(logservice.GetClient(), logservice.entries)
}

// FilteredEntries gets the log entries of this service with filtering applied (e.g. skip, top).
func (logservice *LogService) FilteredEntries(options ...common.FilterOption) ([]*LogEntry, error) {
	var filter common.Filter
	filter.SetFilter(options...)
	return ListReferencedLogEntrys(logservice.GetClient(), fmt.Sprintf("%s%s", logservice.entries, filter))
}

// ClearLog shall delete all entries found in the Entries collection for this
// Log Service.
func (logservice *LogService) ClearLog() error {
	t := struct {
		Action string
	}{
		Action: "LogService.ClearLog",
	}

	return logservice.Post(logservice.clearLogTarget, t)
}
