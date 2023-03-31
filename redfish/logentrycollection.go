//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

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
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a LogEntry object from the raw JSON.
func (logEntryCollection *LogEntryCollection) UnmarshalJSON(b []byte) (err error) {
	type temp LogEntryCollection
	var t struct {
		temp
	}

	if err = json.Unmarshal(b, &t); err != nil {
		return
	}

	*logEntryCollection = LogEntryCollection(t.temp)
	logEntryCollection.rawData = b
	return
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
