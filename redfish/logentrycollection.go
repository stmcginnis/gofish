//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"github.com/stmcginnis/gofish/common"
)

// LogEntry shall represent the log format for log services.
type LogEntryCollection struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Members contain log entries of this log entry collection
	Members []*LogEntry `json:"Members"`
	// MembersCount is the number of log entries of this log entry collection.
	MembersCount int `json:"Members@odata.count"`
	// MembersNextLink is the link used for pagination
	MembersNextLink string `json:"Members@odata.nextLink"`
}

func GetLogEntryCollection(c common.Client, uri string) (*LogEntryCollection, error) {
	var logEntryCollection LogEntryCollection
	return &logEntryCollection, logEntryCollection.Get(c, uri, &logEntryCollection)
}

func (logEntryCollection *LogEntryCollection) GetNext() (nextLogEntryCollection *LogEntryCollection, err error) {
	if logEntryCollection.MembersNextLink != "" {
		return GetLogEntryCollection(logEntryCollection.Client, logEntryCollection.MembersNextLink)
	}
	return
}
