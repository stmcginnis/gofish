// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// AlarmTrips shall contain properties describing the types of alarms that have
// been raised by the memory. These alarams shall be reset when the system
// resets. Note that if they are re-discovered they can be reasserted.
type AlarmTrips struct {
	// AddressParityError shall be true if an Address Parity Error was detected
	// which could not be corrected by retry.
	AddressParityError bool
	// CorrectableECCError shall be true if the correctable error threshold
	// crossing alarm trip was detected.
	CorrectableECCError bool
	// SpareBlock shall be true if the spare block capacity crossing alarm trip
	// was detected.
	SpareBlock bool
	// Temperature shall be true if a temperature threshold alarm trip was detected.
	Temperature bool
	// UncorrectableECCError shall be true if the uncorrectable error threshold
	// alarm trip was detected.
	UncorrectableECCError bool
}

// CurrentPeriod shall describe the metrics of the memory since last time the
// ClearCurrentPeriod Action was performed or the system reset.
type CurrentPeriod struct {
	// BlocksRead shall be number of blocks read since reset.
	BlocksRead int
	// BlocksWritten shall be mumber of blocks written since reset.
	BlocksWritten int
}

// HealthData shall contain properties which describe the HealthData metrics for
// the current resource.
type HealthData struct {
	// AlarmTrips shall contain properties describe the types of alarms that
	// have been raised by the memory.
	AlarmTrips string
	// DataLossDetected shall be data loss detection status, with true
	// indicating data loss detected.
	DataLossDetected bool
	// LastShutdownSuccess shall be the status ofthe  last shutdown, with true
	// indicating success.
	LastShutdownSuccess bool
	// PerformanceDegraded shall be performance degraded mode status, with true
	// indicating perfomance degraded.
	PerformanceDegraded bool
	// PredictedMediaLifeLeftPercent is This property shall contain an indicator
	// of the percentage of life remaining in the media.
	PredictedMediaLifeLeftPercent int
	// RemainingSpareBlockPercentage shall be the remaining spare blocks in percentage.
	RemainingSpareBlockPercentage int
}

// LifeTime shall describe the metrics of the memory since manufacturing.
type LifeTime struct {
	// BlocksRead shall be number of blocks read for the lifetime of the Memory.
	BlocksRead int
	// BlocksWritten shall be number of blocks written for the lifetime of the Memory.
	BlocksWritten int
}

// MemoryMetrics is used to represent the Memory Metrics for a single Memory
// device in a Redfish implementation.
type MemoryMetrics struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// BlockSizeBytes shall be the block size in bytes of all stucture elements.
	BlockSizeBytes int
	// CurrentPeriod shall contain properties which describe the CurrentPeriod
	// metrics for the current resource.
	CurrentPeriod CurrentPeriod
	// Description provides a description of this resource.
	Description string
	// HealthData shall contain properties which describe the HealthData metrics
	// for the current resource.
	HealthData HealthData
	// LifeTime shall contain properties which describe the LifeTime metrics for
	// the current resource.
	LifeTime LifeTime
}

// GetMemoryMetrics will get a MemoryMetrics instance from the service.
func GetMemoryMetrics(c common.Client, uri string) (*MemoryMetrics, error) {
	resp, err := c.Get(uri)
	defer resp.Body.Close()

	var memorymetrics MemoryMetrics
	err = json.NewDecoder(resp.Body).Decode(&memorymetrics)
	if err != nil {
		return nil, err
	}

	memorymetrics.SetClient(c)
	return &memorymetrics, nil
}

// ListReferencedMemoryMetricss gets the collection of MemoryMetrics from
// a provided reference.
func ListReferencedMemoryMetricss(c common.Client, link string) ([]*MemoryMetrics, error) {
	var result []*MemoryMetrics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, memorymetricsLink := range links.ItemLinks {
		memorymetrics, err := GetMemoryMetrics(c, memorymetricsLink)
		if err != nil {
			return result, err
		}
		result = append(result, memorymetrics)
	}

	return result, nil
}
