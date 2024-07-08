//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"
	"errors"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

type ApplicationConsistencyMethod string

const (
	// HostStandbyApplicationConsistencyMethod supports consistency method commonly orchestrated using application-specific code.
	HostStandbyApplicationConsistencyMethod ApplicationConsistencyMethod = "HotStandby"
	// OtherApplicationConsistencyMethod supports consistency method orchestrated using vendor-specific code.
	OtherApplicationConsistencyMethod ApplicationConsistencyMethod = "Other"
	// VASAApplicationConsistencyMethod supports VMware consistency requirements, such as for VASA and VVOLs.
	VASAApplicationConsistencyMethod ApplicationConsistencyMethod = "VASA"
	// VDIApplicationConsistencyMethod supports Microsoft virtual backup device interface (VDI).
	VDIApplicationConsistencyMethod ApplicationConsistencyMethod = "VDI"
	// VSSApplicationConsistencyMethod supports Microsoft VSS.
	VSSApplicationConsistencyMethod ApplicationConsistencyMethod = "VSS"
)

// ConsistencyGroup A collection of volumes grouped together to ensure write order consistency across all those
// volumes. A management operation on a consistency group, such as configuring replication properties, applies to
// all the volumes within the consistency group.
type ConsistencyGroup struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// ConsistencyMethod shall set the consistency method used by this group.
	ConsistencyMethod ApplicationConsistencyMethod
	// ConsistencyType shall set the consistency type used by this group.
	ConsistencyType ConsistencyType
	// Description provides a description of this resource.
	Description string
	// IsConsistent shall be set to true when the consistency group is in a consistent state.
	IsConsistent bool
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// RemoteReplicaTargets shall reference the URIs to the remote target replicas that are sourced by this replica.
	// Remote indicates that the replica is managed by a separate Swordfish service instance.
	RemoteReplicaTargets []string
	// ReplicaInfo shall describe the replication relationship between this storage group and a corresponding source
	// storage group.
	ReplicaInfo ReplicaInfo
	// ReplicaTargets shall reference the target replicas that are sourced by this replica.
	ReplicaTargets []string
	// ReplicaTargetsCount is the number of replica targets.
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// Status shall contain the status of the ConsistencyGroup.
	Status common.Status
	// Volumes is an array of references to volumes managed by this storage group.
	volumes []string
	// VolumesCount is the number of volumes.
	VolumesCount int `json:"Volumes@odata.count"`
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	assignReplicaTargetTarget            string
	createReplicaTargetTarget            string
	removeReplicaRelationshipTarget      string
	resumeReplicationTarget              string
	reverseReplicationRelationshipTarget string
	splitReplicationTarget               string
	suspendReplicationTarget             string
}

// UnmarshalJSON unmarshals a ConsistencyGroup object from the raw JSON.
func (consistencygroup *ConsistencyGroup) UnmarshalJSON(b []byte) error {
	type temp ConsistencyGroup
	var t struct {
		temp
		Actions struct {
			AssignReplicaTarget            common.ActionTarget `json:"#ConsistencyGroup.AssignReplicaTarget"`
			CreateReplicaTarget            common.ActionTarget `json:"#ConsistencyGroup.CreateReplicaTarget"`
			RemoveReplicaRelationship      common.ActionTarget `json:"#ConsistencyGroup.RemoveReplicaRelationship"`
			ResumeReplication              common.ActionTarget `json:"#ConsistencyGroup.ResumeReplication"`
			ReverseReplicationRelationship common.ActionTarget `json:"#ConsistencyGroup.ReverseReplicationRelationship"`
			SplitReplication               common.ActionTarget `json:"#ConsistencyGroup.SplitReplication"`
			SuspendReplication             common.ActionTarget `json:"#ConsistencyGroup.SuspendReplication"`
		}
		Volumes common.Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*consistencygroup = ConsistencyGroup(t.temp)

	// Extract the links to other entities for later
	consistencygroup.assignReplicaTargetTarget = t.Actions.AssignReplicaTarget.Target
	consistencygroup.createReplicaTargetTarget = t.Actions.CreateReplicaTarget.Target
	consistencygroup.removeReplicaRelationshipTarget = t.Actions.RemoveReplicaRelationship.Target
	consistencygroup.resumeReplicationTarget = t.Actions.ResumeReplication.Target
	consistencygroup.reverseReplicationRelationshipTarget = t.Actions.ReverseReplicationRelationship.Target
	consistencygroup.splitReplicationTarget = t.Actions.SplitReplication.Target
	consistencygroup.suspendReplicationTarget = t.Actions.SuspendReplication.Target

	consistencygroup.volumes = t.Volumes.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	consistencygroup.rawData = b

	return nil
}

// AssignReplicaTarget will establish a replication relationship by assigning an existing consistency group
// to serve as a target replica for an existing source consistency group.
//
// `replicaType` is the type of replica relationship to be created (e.g., Clone, Mirror, Snap).
// `updateMode` is the replica update mode (synchronous vs asynchronous).
// `targetGroupURI` is the Uri to the existing consistency group.
func (consistencygroup *ConsistencyGroup) AssignReplicaTarget(replicaType ReplicaType, updateMode ReplicaUpdateMode, targetGroupURI string) error {
	if consistencygroup.assignReplicaTargetTarget == "" {
		return errors.New("method not supported by this service")
	}

	payload := struct {
		ReplicateType          string
		ReplicaUpdateMode      string
		TargetConsistencyGroup string
	}{
		ReplicateType:          string(replicaType),
		ReplicaUpdateMode:      string(updateMode),
		TargetConsistencyGroup: targetGroupURI,
	}

	return consistencygroup.Post(consistencygroup.assignReplicaTargetTarget, payload)
}

// CreateReplicaTarget will create a new consistency group resource to provide expanded data protection
// through a replica relationship with the specified source consistency group.
//
// `groupName` is the name for the target consistency group.
// `replicaType` is the type of replica relationship to be created (e.g., Clone, Mirror, Snap).
// `updateMode` is the replica update mode (synchronous vs asynchronous).
// `targetGroupURI` is the Uri to the existing consistency group.
func (consistencygroup *ConsistencyGroup) CreateReplicaTarget(groupName string, replicaType ReplicaType, updateMode ReplicaUpdateMode, targetGroupURI string) error {
	if consistencygroup.createReplicaTargetTarget == "" {
		return errors.New("method not supported by this service")
	}

	payload := struct {
		ConsistencyGroupName   string
		ReplicateType          string
		ReplicaUpdateMode      string
		TargetConsistencyGroup string
	}{
		ConsistencyGroupName:   groupName,
		ReplicateType:          string(replicaType),
		ReplicaUpdateMode:      string(updateMode),
		TargetConsistencyGroup: targetGroupURI,
	}

	return consistencygroup.Post(consistencygroup.createReplicaTargetTarget, payload)
}

// RemoveReplicaRelationship will disable data synchronization between a source and target consistency group,
// remove the replication relationship, and optionally delete the target consistency group.
//
// `deleteTarget` indicates whether or not to delete the target consistency group as part of the operation.
// `targetGroupURI` is the Uri to the existing consistency group.
func (consistencygroup *ConsistencyGroup) RemoveReplicaRelationship(deleteTarget bool, targetGroupURI string) error {
	if consistencygroup.removeReplicaRelationshipTarget == "" {
		return errors.New("method not supported by this service")
	}

	payload := struct {
		DeleteTargetConsistencyGroup bool
		TargetConsistencyGroup       string
	}{
		DeleteTargetConsistencyGroup: deleteTarget,
		TargetConsistencyGroup:       targetGroupURI,
	}

	return consistencygroup.Post(consistencygroup.removeReplicaRelationshipTarget, payload)
}

// ResumeReplication will resume the active data synchronization between a source and target
// consistency group, without otherwise altering the replication relationship.
//
// `targetGroupURI` is the Uri to the existing consistency group.
func (consistencygroup *ConsistencyGroup) ResumeReplication(targetGroupURI string) error {
	if consistencygroup.resumeReplicationTarget == "" {
		return errors.New("method not supported by this service")
	}

	payload := struct {
		TargetConsistencyGroup string
	}{
		TargetConsistencyGroup: targetGroupURI,
	}

	return consistencygroup.Post(consistencygroup.resumeReplicationTarget, payload)
}

// ReverseReplicationRelationship will resume the active data synchronization between a source and target
// consistency group, without otherwise altering the replication relationship.
//
// `targetGroupURI` is the Uri to the existing consistency group.
func (consistencygroup *ConsistencyGroup) ReverseReplicationRelationship(targetGroupURI string) error {
	if consistencygroup.reverseReplicationRelationshipTarget == "" {
		return errors.New("method not supported by this service")
	}

	payload := struct {
		TargetConsistencyGroup string
	}{
		TargetConsistencyGroup: targetGroupURI,
	}

	return consistencygroup.Post(consistencygroup.reverseReplicationRelationshipTarget, payload)
}

// SplitReplication will split the replication relationship and suspend data synchronization
// between a source and target consistency group.
//
// `targetGroupURI` is the Uri to the existing consistency group.
func (consistencygroup *ConsistencyGroup) SplitReplication(targetGroupURI string) error {
	if consistencygroup.splitReplicationTarget == "" {
		return errors.New("method not supported by this service")
	}

	payload := struct {
		TargetConsistencyGroup string
	}{
		TargetConsistencyGroup: targetGroupURI,
	}

	return consistencygroup.Post(consistencygroup.splitReplicationTarget, payload)
}

// SuspendReplication will suspend active data synchronization between a source and target
// consistency group, without otherwise altering the replication relationship.
//
// `targetGroupURI` is the Uri to the existing consistency group.
func (consistencygroup *ConsistencyGroup) SuspendReplication(targetGroupURI string) error {
	if consistencygroup.suspendReplicationTarget == "" {
		return errors.New("method not supported by this service")
	}

	payload := struct {
		TargetConsistencyGroup string
	}{
		TargetConsistencyGroup: targetGroupURI,
	}

	return consistencygroup.Post(consistencygroup.suspendReplicationTarget, payload)
}

// Volumes gets the volumes in this consistency group.
func (consistencygroup *ConsistencyGroup) Volumes() ([]*Volume, error) {
	return common.GetObjects[Volume](consistencygroup.GetClient(), consistencygroup.volumes)
}

// Update commits updates to this object's properties to the running system.
func (consistencygroup *ConsistencyGroup) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(ConsistencyGroup)
	original.UnmarshalJSON(consistencygroup.rawData)

	readWriteFields := []string{
		"ConsistencyMethod",
		"ConsistencyType",
		"Volumes",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(consistencygroup).Elem()

	return consistencygroup.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetConsistencyGroup will get a ConsistencyGroup instance from the service.
func GetConsistencyGroup(c common.Client, uri string) (*ConsistencyGroup, error) {
	return common.GetObject[ConsistencyGroup](c, uri)
}

// ListReferencedConsistencyGroups gets the collection of ConsistencyGroup from
// a provided reference.
func ListReferencedConsistencyGroups(c common.Client, link string) ([]*ConsistencyGroup, error) {
	return common.GetCollectionObjects[ConsistencyGroup](c, link)
}
