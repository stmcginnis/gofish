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

// LinesOfServiceCollection contains the lines of service for the CoS
type LinesOfServiceCollection struct {
	DataProtectionLineOfService []DataProtectionLineOfService
	DataSecurityLineOfService   DataSecurityLineOfService
	DataStorageLineOfService    DataStorageLineOfService
	IOConnectivityLineOfService IOConnectivityLineOfService
	IOPerformanceLineOfService  IOPerformanceLineOfService
}

// ClassesOfService is a Swordfish storage system instance.
type ClassesOfService struct {
	common.Entity
	ClassesOfServiceVersion string
	Description             string
	LinesOfService          LinesOfServiceCollection
}

// GetClassesOfService will get a ClassesOfService instance from the Swordfish service.
func GetClassesOfService(c common.Client, uri string) (*ClassesOfService, error) {
	resp, err := c.Get(uri)
	defer resp.Body.Close()

	var ClassesOfService ClassesOfService
	err = json.NewDecoder(resp.Body).Decode(&ClassesOfService)
	if err != nil {
		return nil, err
	}

	ClassesOfService.SetClient(c)
	return &ClassesOfService, nil
}

// ListReferencedClassesOfServices gets the collection of ClassesOfServices
func ListReferencedClassesOfServices(c common.Client, link string) ([]*ClassesOfService, error) {
	var result []*ClassesOfService
	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, ClassesOfServiceLink := range links.ItemLinks {
		ClassesOfService, err := GetClassesOfService(c, ClassesOfServiceLink)
		if err != nil {
			return result, err
		}
		result = append(result, ClassesOfService)
	}

	return result, nil
}
