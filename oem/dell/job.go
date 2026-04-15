//
// SPDX-License-Identifier: BSD-3-Clause
//

package dell

import (
	"github.com/stmcginnis/gofish/common"
)

type Job struct {
	common.Entity
	common.Message

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// This property represents completion time of Job.
	CompletionTime string
	Description    string
	// This property represents End time of Job. This is the timestamp until when the service will wait for a job to complete. If a job did not complete within this time, it will be killed and marked as failed. TIME_NA is a default value that implies EndTime is not applicable. It is optional for clients to specify this property when creating a job.
	EndTime string
	// The state of the Job.
	JobState string
	// This property represent configuration type of job. The value of this property will be one of the possible configuration type of job.
	JobType string
	// The percentage completion of the Job.
	PercentComplete int
	// This property represents start time of Job.
	StartTime string
}

// GetJob will get a Job instance from the service.
func GetJob(c common.Client, uri string) (*Job, error) {
	return common.GetObject[Job](c, uri)
}
