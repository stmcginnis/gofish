// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
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

// ClassOfService is This resource shall define a service option composed
// of one or more line of service entities. ITIL defines a service
// option as a choice of utility or warranty for a service.
type ClassOfService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataID is the odata identifier.
	ODataID string `json:"@odata.id"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ClassOfServiceVersion is The version describing the creation or last
	// modification of this service option specification. The string
	// representing the version shall be in the form: M + '.' + N + '.' + U
	// Where: M - The major version (in numeric form). N - The minor version
	// (in numeric form). U - The update (e.g. errata or patch in numeric
	// form).
	ClassOfServiceVersion string
	// DataProtectionLinesOfService shall be a set of data protection service
	// options. Within a class of service, one data protection service option
	// shall be present for each replication session.
	DataProtectionLinesOfService []DataProtectionLineOfService
	// DataProtectionLinesOfServiceCount is the number of DataProtectionLineOfService.
	DataProtectionLinesOfServiceCount int `json:"DataProtectionLinesOfService@odata.count"`
	// DataSecurityLinesOfService shall be a set of data security service options.
	DataSecurityLinesOfService []DataSecurityLineOfService
	// DataSecurityLinesOfServiceCount is number of DataSercurityLineOfService.
	DataSecurityLinesOfServiceCount int `json:"DataSecurityLinesOfService@odata.count"`
	// DataStorageLinesOfService shall be a set of data protection service options.
	DataStorageLinesOfService []DataStorageLineOfService
	// DataStorageLinesOfServiceCount is the number of DataStorageLinesOfService.
	DataStorageLinesOfServiceCount int `json:"DataStorageLinesOfService@odata.count"`
	// Description provides a description of this resource.
	Description string
	// IOConnectivityLinesOfService shall be a set of IO connectivity service
	// options. Within a class of service, at most one IO connectivity service
	// option may be present for a value of AccessProtocol.
	IOConnectivityLinesOfService []IOConnectivityLineOfService
	// IOConnectivityLinesOfServiceCount is
	IOConnectivityLinesOfServiceCount int `json:"IOConnectivityLinesOfService@odata.count"`
	// IOPerformanceLinesOfService shall be a set of IO
	// performance service options.
	IOPerformanceLinesOfService []IOPerformanceLineOfService
	// IOPerformanceLinesOfServiceCount is
	IOPerformanceLinesOfServiceCount int `json:"IOPerformanceLinesOfService@odata.count"`
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
}

// GetClassOfService will get a ClassOfService instance from the service.
func GetClassOfService(c common.Client, uri string) (*ClassOfService, error) {
	resp, err := c.Get(uri)
	defer resp.Body.Close()

	var classofservice ClassOfService
	err = json.NewDecoder(resp.Body).Decode(&classofservice)
	if err != nil {
		return nil, err
	}

	classofservice.SetClient(c)
	return &classofservice, nil
}

// ListReferencedClassOfServices gets the collection of ClassOfService from
// a provided reference.
func ListReferencedClassOfServices(c common.Client, link string) ([]*ClassOfService, error) {
	var result []*ClassOfService
	if link == "" {
		return result, nil
	}

	links, err := common.GetCollection(c, link)
	if err != nil {
		return result, err
	}

	for _, classofserviceLink := range links.ItemLinks {
		classofservice, err := GetClassOfService(c, classofserviceLink)
		if err != nil {
			return result, err
		}
		result = append(result, classofservice)
	}

	return result, nil
}
