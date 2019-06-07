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

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/school/common"
)

// IOStatistics is used to represent the IO statistics of the requested
// object.
type IOStatistics struct {
	// NonIORequestTime shall be an ISO 8601 conformant duration describing the
	// time that the resource is busy processing non IO requests.
	NonIORequestTime string
	// NonIORequests shall represent the total count from the time of last reset
	// or wrap of non IO requests.
	NonIORequests int
	// ReadHitIORequests shall represent the total count from the time of last
	// reset or wrap of read IO requests satisfied from memory.
	ReadHitIORequests int
	// ReadIOKiBytes shall represent the total number of kibibytes read from the
	// time of last reset or wrap.
	ReadIOKiBytes int
	// ReadIORequestTime shall be an ISO 8601 conformant duration describing the
	// time that the resource is busy processing read requests.
	ReadIORequestTime string
	// ReadIORequests shall represent the total count from the time of last
	// reset or wrap of read IO requests satisfied from either media or memory
	// (i.e. from a storage device or from a cache).
	ReadIORequests int
	// WriteHitIORequests shall represent the total count from the time of last
	// reset or wrap of write IO requests coallesced into memory.
	WriteHitIORequests int
	// WriteIOKiBytes shall represent the total number of kibibytes written from
	// the time of last reset or wrap.
	WriteIOKiBytes int
	// WriteIORequestTime shall be an ISO 8601 conformant duration describing
	// the time that the resource is busy processing write requests.
	WriteIORequestTime string
	// WriteIORequests shall represent the total count from the time of last
	// reset or wrap of write IO requests.
	WriteIORequests int
}

// GetIOStatistics will get a IOStatistics instance from the service.
func GetIOStatistics(c common.Client, uri string) (*IOStatistics, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var iostatistics IOStatistics
	err = json.NewDecoder(resp.Body).Decode(&iostatistics)
	if err != nil {
		return nil, err
	}
	return &iostatistics, nil
}

// ListReferencedIOStatisticss gets the collection of IOStatistics from
// a provided reference.
func ListReferencedIOStatisticss(c common.Client, link string) ([]*IOStatistics, error) {
	var result []*IOStatistics
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, iostatisticsLink := range links.ItemLinks {
		iostatistics, err := GetIOStatistics(c, iostatisticsLink)
		if err != nil {
			return result, err
		}
		result = append(result, iostatistics)
	}

	return result, nil
}
