//
// SPDX-License-Identifier: BSD-3-Clause
//
// 1.2.4a - #ConsistencyGroup.v1_1_1.ConsistencyGroup

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type ApplicationConsistencyMethod string

const (
	// HotStandbyApplicationConsistencyMethod Supports consistency method commonly
	// orchestrated using application-specific code.
	HotStandbyApplicationConsistencyMethod ApplicationConsistencyMethod = "HotStandby"
	// VASAApplicationConsistencyMethod Supports VMware consistency requirements,
	// such as for VASA and VVOLs.
	VASAApplicationConsistencyMethod ApplicationConsistencyMethod = "VASA"
	// VDIApplicationConsistencyMethod Supports Microsoft virtual backup device
	// interface (VDI).
	VDIApplicationConsistencyMethod ApplicationConsistencyMethod = "VDI"
	// VSSApplicationConsistencyMethod Supports Microsoft VSS.
	VSSApplicationConsistencyMethod ApplicationConsistencyMethod = "VSS"
	// OtherApplicationConsistencyMethod Supports consistency method orchestrated
	// using vendor-specific code.
	OtherApplicationConsistencyMethod ApplicationConsistencyMethod = "Other"
)

// ConsistencyGroup is a collection of volumes grouped together to ensure write
// order consistency across all those volumes. A management operation on a
// consistency group, such as configuring replication properties, applies to all
// the volumes within the consistency group.
type ConsistencyGroup struct {
	common.Entity
	// ConsistencyMethod shall set the consistency method used by this group.
	ConsistencyMethod ApplicationConsistencyMethod
	// ConsistencyType shall set the consistency type used by this group.
	ConsistencyType ConsistencyType
	// IsConsistent shall be set to true when the consistency group is in a
	// consistent state.
	IsConsistent bool
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RemoteReplicaTargets shall reference the URIs to the remote target replicas
	// that are sourced by this replica. Remote indicates that the replica is
	// managed by a separate Swordfish service instance.
	//
	// Version added: v1.1.0
	RemoteReplicaTargets []string
	// ReplicaInfo shall describe the replication relationship between this storage
	// group and a corresponding source storage group.
	ReplicaInfo ReplicaInfo
	// ReplicaTargets shall reference the target replicas that are sourced by this
	// replica.
	ReplicaTargets []string
	// ReplicaTargets@odata.count
	ReplicaTargetsCount int `json:"ReplicaTargets@odata.count"`
	// Status shall contain the status of the ConsistencyGroup.
	Status common.Status
	// Volumes is an array of references to volumes managed by this storage group.
	volumes []string
	// Volumes@odata.count
	VolumesCount int `json:"Volumes@odata.count"`
	// assignReplicaTargetTarget is the URL to send AssignReplicaTarget requests.
	assignReplicaTargetTarget string
	// createReplicaTargetTarget is the URL to send CreateReplicaTarget requests.
	createReplicaTargetTarget string
	// removeReplicaRelationshipTarget is the URL to send RemoveReplicaRelationship requests.
	removeReplicaRelationshipTarget string
	// resumeReplicationTarget is the URL to send ResumeReplication requests.
	resumeReplicationTarget string
	// reverseReplicationRelationshipTarget is the URL to send ReverseReplicationRelationship requests.
	reverseReplicationRelationshipTarget string
	// splitReplicationTarget is the URL to send SplitReplication requests.
	splitReplicationTarget string
	// suspendReplicationTarget is the URL to send SuspendReplication requests.
	suspendReplicationTarget string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ConsistencyGroup object from the raw JSON.
func (c *ConsistencyGroup) UnmarshalJSON(b []byte) error {
	type temp ConsistencyGroup
	type cActions struct {
		AssignReplicaTarget            common.ActionTarget `json:"#ConsistencyGroup.AssignReplicaTarget"`
		CreateReplicaTarget            common.ActionTarget `json:"#ConsistencyGroup.CreateReplicaTarget"`
		RemoveReplicaRelationship      common.ActionTarget `json:"#ConsistencyGroup.RemoveReplicaRelationship"`
		ResumeReplication              common.ActionTarget `json:"#ConsistencyGroup.ResumeReplication"`
		ReverseReplicationRelationship common.ActionTarget `json:"#ConsistencyGroup.ReverseReplicationRelationship"`
		SplitReplication               common.ActionTarget `json:"#ConsistencyGroup.SplitReplication"`
		SuspendReplication             common.ActionTarget `json:"#ConsistencyGroup.SuspendReplication"`
	}
	var tmp struct {
		temp
		Actions cActions
		Volumes common.Links
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = ConsistencyGroup(tmp.temp)

	// Extract the links to other entities for later
	c.assignReplicaTargetTarget = tmp.Actions.AssignReplicaTarget.Target
	c.createReplicaTargetTarget = tmp.Actions.CreateReplicaTarget.Target
	c.removeReplicaRelationshipTarget = tmp.Actions.RemoveReplicaRelationship.Target
	c.resumeReplicationTarget = tmp.Actions.ResumeReplication.Target
	c.reverseReplicationRelationshipTarget = tmp.Actions.ReverseReplicationRelationship.Target
	c.splitReplicationTarget = tmp.Actions.SplitReplication.Target
	c.suspendReplicationTarget = tmp.Actions.SuspendReplication.Target

	c.volumes = tmp.Volumes.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	c.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *ConsistencyGroup) Update() error {
	readWriteFields := []string{
		"ConsistencyMethod",
		"ConsistencyType",
		"ReplicaInfo",
		"ReplicaTargets@odata.count",
		"Status",
		"Volumes",
		"Volumes@odata.count",
	}

	return c.UpdateFromRawData(c, c.rawData, readWriteFields)
}

// Volumes gets the volumes in this consistency group.
func (c *ConsistencyGroup) Volumes() ([]*Volume, error) {
	return common.GetObjects[Volume](c.GetClient(), c.volumes)
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

// AssignReplicaTarget shall be used to establish a replication relationship by
// assigning an existing consistency group to serve as a target replica for an
// existing source consistency group.
// replicaType - This parameter shall contain the type of replica relationship
// to be created.
// replicaUpdateMode - This parameter shall specify the replica update mode.
// targetConsistencyGroup - This parameter shall contain the Uri to the
// existing consistency group.
func (c *ConsistencyGroup) AssignReplicaTarget(replicaType ReplicaType, replicaUpdateMode ReplicaUpdateMode, targetConsistencyGroup string) error {
	payload := make(map[string]any)
	payload["ReplicaType"] = replicaType
	payload["ReplicaUpdateMode"] = replicaUpdateMode
	payload["TargetConsistencyGroup"] = targetConsistencyGroup
	return c.Post(c.assignReplicaTargetTarget, payload)
}

// CreateReplicaTarget shall be used to create a new consistency group resource to
// provide expanded data protection through a replica relationship with the
// specified source consistency group.
// consistencyGroupName - This parameter shall contain the Name for the target
// consistency group.
// replicaType - This parameter shall contain the type of replica relationship
// to be created.
// replicaUpdateMode - This parameter shall specify the replica update mode.
// targetStoragePool - This parameter shall contain the Uri to the existing
// StoragePool in which to create the target consistency group.
func (c *ConsistencyGroup) CreateReplicaTarget(consistencyGroupName string, replicaType ReplicaType, replicaUpdateMode ReplicaUpdateMode, targetStoragePool string) error {
	payload := make(map[string]any)
	payload["ConsistencyGroupName"] = consistencyGroupName
	payload["ReplicaType"] = replicaType
	payload["ReplicaUpdateMode"] = replicaUpdateMode
	payload["TargetStoragePool"] = targetStoragePool
	return c.Post(c.createReplicaTargetTarget, payload)
}

// RemoveReplicaRelationship shall be used to disable data synchronization between a source
// and target consistency group, remove the replication relationship, and
// optionally delete the target consistency group.
// deleteTargetConsistencyGroup - This parameter shall indicate whether or not
// to delete the target consistency group as part of the operation. If not
// specified, the system should use its default behavior.
// targetConsistencyGroup - This parameter shall contain the Uri to the
// existing target consistency group.
func (c *ConsistencyGroup) RemoveReplicaRelationship(deleteTargetConsistencyGroup bool, targetConsistencyGroup string) error {
	payload := make(map[string]any)
	payload["DeleteTargetConsistencyGroup"] = deleteTargetConsistencyGroup
	payload["TargetConsistencyGroup"] = targetConsistencyGroup
	return c.Post(c.removeReplicaRelationshipTarget, payload)
}

// ResumeReplication shall be used to resume the active data synchronization between
// a source and target consistency group, without otherwise altering the
// replication relationship.
// targetConsistencyGroup - This parameter shall contain the Uri to the
// existing target consistency group.
func (c *ConsistencyGroup) ResumeReplication(targetConsistencyGroup string) error {
	payload := make(map[string]any)
	payload["TargetConsistencyGroup"] = targetConsistencyGroup
	return c.Post(c.resumeReplicationTarget, payload)
}

// ReverseReplicationRelationship shall be used to reverse the replication relationship between a
// source and target consistency group.
// targetConsistencyGroup - This parameter shall contain the Uri to the
// existing target consistency group.
func (c *ConsistencyGroup) ReverseReplicationRelationship(targetConsistencyGroup string) error {
	payload := make(map[string]any)
	payload["TargetConsistencyGroup"] = targetConsistencyGroup
	return c.Post(c.reverseReplicationRelationshipTarget, payload)
}

// SplitReplication shall be used to split the replication relationship and suspend
// data synchronization between a source and target consistency group.
// targetConsistencyGroup - This parameter shall contain the Uri to the
// existing target consistency group.
func (c *ConsistencyGroup) SplitReplication(targetConsistencyGroup string) error {
	payload := make(map[string]any)
	payload["TargetConsistencyGroup"] = targetConsistencyGroup
	return c.Post(c.splitReplicationTarget, payload)
}

// SuspendReplication shall be used to suspend active data synchronization between a
// source and target consistency group, without otherwise altering the
// replication relationship.
// targetConsistencyGroup - This parameter shall contain the Uri to the
// existing target consistency group.
func (c *ConsistencyGroup) SuspendReplication(targetConsistencyGroup string) error {
	payload := make(map[string]any)
	payload["TargetConsistencyGroup"] = targetConsistencyGroup
	return c.Post(c.suspendReplicationTarget, payload)
}
