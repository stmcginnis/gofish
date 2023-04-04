//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// ClassOfService shall define a service option composed
// of one or more line of service entities. ITIL defines a service
// option as a choice of utility or warranty for a service.
type ClassOfService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ClassOfServiceVersion is the version describing the creation or last
	// modification of this service option specification. The string
	// representing the version shall be in the form: M + '.' + N + '.' + U
	// Where: M - The major version (in numeric form). N - The minor version
	// (in numeric form). U - The update (e.g. errata or patch in numeric
	// form).
	ClassOfServiceVersion string
	// DataProtectionLinesOfService shall be a set of data protection service
	// options. Within a class of service, one data protection service option
	// shall be present for each replication session.
	dataProtectionLinesOfService []string
	// DataProtectionLinesOfServiceCount is the number of DataProtectionLineOfService.
	DataProtectionLinesOfServiceCount int `json:"DataProtectionLinesOfService@odata.count"`
	// DataSecurityLinesOfService shall be a set of data security service options.
	dataSecurityLinesOfService []string
	// DataSecurityLinesOfServiceCount is number of DataSecurityLineOfService.
	DataSecurityLinesOfServiceCount int `json:"DataSecurityLinesOfService@odata.count"`
	// DataStorageLinesOfService shall be a set of data protection service options.
	dataStorageLinesOfService []string
	// DataStorageLinesOfServiceCount is the number of DataStorageLinesOfService.
	DataStorageLinesOfServiceCount int `json:"DataStorageLinesOfService@odata.count"`
	// Description provides a description of this resource.
	Description string
	// IOConnectivityLinesOfService shall be a set of IO connectivity service
	// options. Within a class of service, at most one IO connectivity service
	// option may be present for a value of AccessProtocol.
	ioConnectivityLinesOfService []string
	// IOConnectivityLinesOfServiceCount is the number of IOConnectivityLinesOfService.
	IOConnectivityLinesOfServiceCount int `json:"IOConnectivityLinesOfService@odata.count"`
	// IOPerformanceLinesOfService shall be a set of IO performance service options.
	ioPerformanceLinesOfService []string
	// IOPerformanceLinesOfServiceCount is the number of IOPerformanceLinesOfService.
	IOPerformanceLinesOfServiceCount int `json:"IOPerformanceLinesOfService@odata.count"`
	// Identifier shall be unique within the managed ecosystem.
	Identifier common.Identifier
}

// UnmarshalJSON unmarshals a ClassOfService object from the raw JSON.
func (classofservice *ClassOfService) UnmarshalJSON(b []byte) error {
	type temp ClassOfService
	var t struct {
		temp
		DataProtectionLinesOfService common.Links
		DataSecurityLinesOfService   common.Links
		DataStorageLinesOfService    common.Links
		IOConnectivityLinesOfService common.Links
		IOPerformanceLinesOfService  common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	// Extract the links to other entities for later
	*classofservice = ClassOfService(t.temp)
	classofservice.dataProtectionLinesOfService = t.DataProtectionLinesOfService.ToStrings()
	classofservice.dataSecurityLinesOfService = t.DataSecurityLinesOfService.ToStrings()
	classofservice.dataStorageLinesOfService = t.DataStorageLinesOfService.ToStrings()
	classofservice.ioConnectivityLinesOfService = t.IOConnectivityLinesOfService.ToStrings()
	classofservice.ioPerformanceLinesOfService = t.IOPerformanceLinesOfService.ToStrings()

	return nil
}

// GetClassOfService will get a ClassOfService instance from the service.
func GetClassOfService(c common.Client, uri string) (*ClassOfService, error) {
	var classOfService ClassOfService
	return &classOfService, classOfService.Get(c, uri, &classOfService)
}

// ListReferencedClassOfServices gets the collection of ClassOfService from
// a provided reference.
func ListReferencedClassOfServices(c common.Client, link string) ([]*ClassOfService, error) { //nolint:dupl
	var result []*ClassOfService
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *ClassOfService
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		classofservice, err := GetClassOfService(c, link)
		ch <- GetResult{Item: classofservice, Link: link, Error: err}
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

// DataProtectionLinesOfServices gets the DataProtectionLinesOfService that are
// part of this ClassOfService.
func (classofservice *ClassOfService) DataProtectionLinesOfServices() ([]*DataProtectionLineOfService, error) {
	var result []*DataProtectionLineOfService

	collectionError := common.NewCollectionError()
	for _, dpLosLink := range classofservice.dataProtectionLinesOfService {
		dpLos, err := GetDataProtectionLineOfService(classofservice.GetClient(), dpLosLink)
		if err != nil {
			collectionError.Failures[dpLosLink] = err
		} else {
			result = append(result, dpLos)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// DataSecurityLinesOfServices gets the DataSecurityLinesOfService that are
// part of this ClassOfService.
func (classofservice *ClassOfService) DataSecurityLinesOfServices() ([]*DataSecurityLineOfService, error) {
	var result []*DataSecurityLineOfService

	collectionError := common.NewCollectionError()
	for _, dsLosLink := range classofservice.dataSecurityLinesOfService {
		dsLos, err := GetDataSecurityLineOfService(classofservice.GetClient(), dsLosLink)
		if err != nil {
			collectionError.Failures[dsLosLink] = err
		} else {
			result = append(result, dsLos)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// DataStorageLinesOfServices gets the DataStorageLinesOfService that are
// part of this ClassOfService.
func (classofservice *ClassOfService) DataStorageLinesOfServices() ([]*DataStorageLineOfService, error) {
	var result []*DataStorageLineOfService

	collectionError := common.NewCollectionError()
	for _, dsLosLink := range classofservice.dataStorageLinesOfService {
		dsLos, err := GetDataStorageLineOfService(classofservice.GetClient(), dsLosLink)
		if err != nil {
			collectionError.Failures[dsLosLink] = err
		} else {
			result = append(result, dsLos)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// IOConnectivityLinesOfServices gets the IOConnectivityLinesOfService that are
// part of this ClassOfService.
func (classofservice *ClassOfService) IOConnectivityLinesOfServices() ([]*IOConnectivityLineOfService, error) {
	var result []*IOConnectivityLineOfService

	collectionError := common.NewCollectionError()
	for _, ioLosLink := range classofservice.dataSecurityLinesOfService {
		ioLos, err := GetIOConnectivityLineOfService(classofservice.GetClient(), ioLosLink)
		if err != nil {
			collectionError.Failures[ioLosLink] = err
		} else {
			result = append(result, ioLos)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// IOPerformanceLinesOfServices gets the IOPerformanceLinesOfService that are
// part of this ClassOfService.
func (classofservice *ClassOfService) IOPerformanceLinesOfServices() ([]*IOPerformanceLineOfService, error) {
	var result []*IOPerformanceLineOfService

	collectionError := common.NewCollectionError()
	for _, ioLosLink := range classofservice.dataSecurityLinesOfService {
		ioLos, err := GetIOPerformanceLineOfService(classofservice.GetClient(), ioLosLink)
		if err != nil {
			collectionError.Failures[ioLosLink] = err
		} else {
			result = append(result, ioLos)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
