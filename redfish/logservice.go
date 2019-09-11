//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

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
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
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
	logservice.entries = string(t.Entries)
	logservice.clearLogTarget = t.Actions.ClearLog.Target

	return nil
}

// GetLogService will get a LogService instance from the service.
func GetLogService(c common.Client, uri string) (*LogService, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var logservice LogService
	err = json.NewDecoder(resp.Body).Decode(&logservice)
	if err != nil {
		return nil, err
	}

	logservice.SetClient(c)
	return &logservice, nil
}

// ListReferencedLogServices gets the collection of LogService from a provided reference.
func ListReferencedLogServices(c common.Client, link string) ([]*LogService, error) {
	var result []*LogService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, logserviceLink := range links.ItemLinks {
		logservice, err := GetLogService(c, logserviceLink)
		if err != nil {
			return result, err
		}
		result = append(result, logservice)
	}

	return result, nil
}

// Entries gets the log entries of this service.
func (logservice *LogService) Entries() ([]*LogEntry, error) {
	return ListReferencedLogEntrys(logservice.Client, logservice.entries)
}

// ClearLog shall delete all entries found in the Entries collection for this
// Log Service.
func (logservice *LogService) ClearLog() error {
	var payload []byte
	_, err := logservice.Client.Post(logservice.clearLogTarget, payload)
	return err
}
