//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"github.com/stmcginnis/gofish/common"
)

// DataProtectionLineOfService describes a replica that protects data from loss.
// The requirements must be met collectively by the communication path and the
// replica.
type DataProtectionLineOfService struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// IsIsolated is True shall indicate that the replica is in a separate
	// fault domain from its source. The default value of this property is
	// false.
	IsIsolated bool
	// MinLifetime shall be an ISO 8601 duration that specifies
	// the minimum required lifetime of the replica. Note: The maximum number
	// of replicas can be determined using this value together with the
	// replicaSchedule.
	MinLifetime string
	// RecoveryGeographicObjective specifies the geographic scope of the failure
	// domain.
	RecoveryGeographicObjective FailureDomainScope
	// RecoveryPointObjectiveTime shall be an ISO 8601 duration that specifies
	// the maximum time over which source data may be lost on failure. In the
	// case that IsIsolated = false, failure of the domain is not a
	// consideration.
	RecoveryPointObjectiveTime string
	// RecoveryTimeObjective shall be an enumeration that
	// indicates the maximum time required to access an alternate replica. In
	// the case that IsIsolated = false, failure of the domain is not a
	// consideration.
	RecoveryTimeObjective RecoveryAccessScope
	// ReplicaAccessLocation is used if the data access location of the
	// replica is required to be at a specific location.   Note 1: The
	// location value may be granular.  Note 2: A value may be required for
	// some regulatory compliance.
	ReplicaAccessLocation common.Location
	// ReplicaClassOfService shall reference the class of
	// service that defines the required service levels of the replica.
	ReplicaClassOfService ClassOfService
	// ReplicaType is the type of replica.
	ReplicaType ReplicaType
	// Schedule if a replica is made periodically, the value shall define
	// the schedule.
	Schedule common.Schedule
}

// GetDataProtectionLineOfService will get a DataProtectionLineOfService instance from the service.
func GetDataProtectionLineOfService(c common.Client, uri string) (*DataProtectionLineOfService, error) {
	var dataProtectionLineOfService DataProtectionLineOfService
	return &dataProtectionLineOfService, dataProtectionLineOfService.Get(c, uri, &dataProtectionLineOfService)
}

// ListReferencedDataProtectionLineOfServices gets the collection of DataProtectionLineOfService from
// a provided reference.
func ListReferencedDataProtectionLineOfServices(c common.Client, link string) ([]*DataProtectionLineOfService, error) { //nolint:dupl
	var result []*DataProtectionLineOfService
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *DataProtectionLineOfService
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		dataprotectionlineofservice, err := GetDataProtectionLineOfService(c, link)
		ch <- GetResult{Item: dataprotectionlineofservice, Link: link, Error: err}
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

// ReplicaRequest is a request for a replica.
type ReplicaRequest struct {
	// ReplicaName shall be the names of the replica.
	ReplicaName string
	// ReplicaSource shall reference a resource to be
	// replicated.
	ReplicaSource common.Link
}
