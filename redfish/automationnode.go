//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #AutomationNode.v1_0_0.AutomationNode

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type MotionAxisType string

const (
	// XMotionAxisType X axis.
	XMotionAxisType MotionAxisType = "X"
	// YMotionAxisType Y axis.
	YMotionAxisType MotionAxisType = "Y"
	// ZMotionAxisType Z axis.
	ZMotionAxisType MotionAxisType = "Z"
	// TwoAxisMotionAxisType 2-axis.
	TwoAxisMotionAxisType MotionAxisType = "TwoAxis"
	// ThreeAxisMotionAxisType 3-axis.
	ThreeAxisMotionAxisType MotionAxisType = "ThreeAxis"
)

type MotionProfileType string

const (
	// TrapezoidalMotionProfileType Trapezoidal profile.
	TrapezoidalMotionProfileType MotionProfileType = "Trapezoidal"
	// SCurveMotionProfileType S-curve profile.
	SCurveMotionProfileType MotionProfileType = "SCurve"
	// NoneMotionProfileType No profile specified.
	NoneMotionProfileType MotionProfileType = "None"
)

type NodeState string

const (
	// IdleNodeState The node's controller is idle.
	IdleNodeState NodeState = "Idle"
	// DoneNodeState The node's controller has reached its destination position.
	DoneNodeState NodeState = "Done"
	// WaitingNodeState The node's controller is waiting to start.
	WaitingNodeState NodeState = "Waiting"
	// ConditionStopNodeState The node's controller has stopped due to a condition
	// fault.
	ConditionStopNodeState NodeState = "ConditionStop"
	// ErrorStopNodeState The node's controller has stopped due to an error.
	ErrorStopNodeState NodeState = "ErrorStop"
	// RunningNodeState The node's controller is running.
	RunningNodeState NodeState = "Running"
)

type NodeType string

const (
	// MotionPositionNodeType is a position-based profiled motion node where
	// position, velocity, and acceleration are all controlled.
	MotionPositionNodeType NodeType = "MotionPosition"
	// MotionVelocityNodeType is a velocity-based profiled motion node where
	// position, velocity and acceleration are all controlled.
	MotionVelocityNodeType NodeType = "MotionVelocity"
	// MotionPositionGroupNodeType is a multi-axis position-based profiled motion
	// node where position, velocity and acceleration are all controlled.
	MotionPositionGroupNodeType NodeType = "MotionPositionGroup"
	// PIDNodeType is a node that attempts to match a set point using a PID control
	// algorithm.
	PIDNodeType NodeType = "PID"
	// SimpleNodeType is a simple node that incorporates no automated control
	// function.
	SimpleNodeType NodeType = "Simple"
)

// AutomationNode shall represent an automation node for a Redfish
// implementation.
type AutomationNode struct {
	common.Entity
	// Instrumentation shall contain a link to a resource of type
	// 'AutomationInstrumentation' that represents the instrumentation for this
	// automation node.
	instrumentation string
	// MotionAxis shall contain the primary axis of motion for this motion-related
	// node.
	MotionAxis MotionAxisType
	// MotionProfile shall contain the type of motion profile used for this
	// motion-based node.
	MotionProfile MotionProfileType
	// NodeState shall contain the current state of the automation node.
	NodeState NodeState
	// NodeType shall contain the type for the automation node.
	NodeType NodeType
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of a resource.
	Status common.Status
	// resetTarget is the URL to send Reset requests.
	resetTarget string
	// sendTriggerTarget is the URL to send SendTrigger requests.
	sendTriggerTarget string
	// startTarget is the URL to send Start requests.
	startTarget string
	// stopTarget is the URL to send Stop requests.
	stopTarget string
	// waitTarget is the URL to send Wait requests.
	waitTarget string
	// automationNodeGroup are the URIs for AutomationNodeGroup.
	automationNodeGroup []string
	// chassis are the URIs for Chassis.
	chassis []string
	// outputControl is the URI for OutputControl.
	outputControl string
	// pidFeedbackSensor is the URI for PidFeedbackSensor.
	pidFeedbackSensor string
	// positionSensor is the URI for PositionSensor.
	positionSensor string
	// velocitySensor is the URI for VelocitySensor.
	velocitySensor string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a AutomationNode object from the raw JSON.
func (a *AutomationNode) UnmarshalJSON(b []byte) error {
	type temp AutomationNode
	type aActions struct {
		Reset       common.ActionTarget `json:"#AutomationNode.Reset"`
		SendTrigger common.ActionTarget `json:"#AutomationNode.SendTrigger"`
		Start       common.ActionTarget `json:"#AutomationNode.Start"`
		Stop        common.ActionTarget `json:"#AutomationNode.Stop"`
		Wait        common.ActionTarget `json:"#AutomationNode.Wait"`
	}
	type aLinks struct {
		AutomationNodeGroup common.Links `json:"AutomationNodeGroup"`
		Chassis             common.Links `json:"Chassis"`
		OutputControl       common.Link  `json:"OutputControl"`
		PidFeedbackSensor   common.Link  `json:"PidFeedbackSensor"`
		PositionSensor      common.Link  `json:"PositionSensor"`
		VelocitySensor      common.Link  `json:"VelocitySensor"`
	}
	var tmp struct {
		temp
		Actions         aActions
		Links           aLinks
		Instrumentation common.Link `json:"instrumentation"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = AutomationNode(tmp.temp)

	// Extract the links to other entities for later
	a.resetTarget = tmp.Actions.Reset.Target
	a.sendTriggerTarget = tmp.Actions.SendTrigger.Target
	a.startTarget = tmp.Actions.Start.Target
	a.stopTarget = tmp.Actions.Stop.Target
	a.waitTarget = tmp.Actions.Wait.Target
	a.automationNodeGroup = tmp.Links.AutomationNodeGroup.ToStrings()
	a.chassis = tmp.Links.Chassis.ToStrings()
	a.outputControl = tmp.Links.OutputControl.String()
	a.pidFeedbackSensor = tmp.Links.PidFeedbackSensor.String()
	a.positionSensor = tmp.Links.PositionSensor.String()
	a.velocitySensor = tmp.Links.VelocitySensor.String()
	a.instrumentation = tmp.Instrumentation.String()

	// This is a read/write object, so we need to save the raw object data for later
	a.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (a *AutomationNode) Update() error {
	readWriteFields := []string{
		"MotionAxis",
		"MotionProfile",
		"Status",
	}

	return a.UpdateFromRawData(a, a.rawData, readWriteFields)
}

// GetAutomationNode will get a AutomationNode instance from the service.
func GetAutomationNode(c common.Client, uri string) (*AutomationNode, error) {
	return common.GetObject[AutomationNode](c, uri)
}

// ListReferencedAutomationNodes gets the collection of AutomationNode from
// a provided reference.
func ListReferencedAutomationNodes(c common.Client, link string) ([]*AutomationNode, error) {
	return common.GetCollectionObjects[AutomationNode](c, link)
}

// Reset shall reset the node to its power-on state.
func (a *AutomationNode) Reset() error {
	payload := make(map[string]any)
	return a.Post(a.resetTarget, payload)
}

// SendTrigger shall pulse the trigger signal for the node.
func (a *AutomationNode) SendTrigger() error {
	payload := make(map[string]any)
	return a.Post(a.sendTriggerTarget, payload)
}

// Start shall transition the node to the 'Running' state. This action
// shall pulse the trigger signal to other connected nodes that are in the
// 'Waiting' state.
func (a *AutomationNode) Start() error {
	payload := make(map[string]any)
	return a.Post(a.startTarget, payload)
}

// Stop shall transition the node to the 'Idle' state.
func (a *AutomationNode) Stop() error {
	payload := make(map[string]any)
	return a.Post(a.stopTarget, payload)
}

// Wait shall transition the node to the 'Waiting' state.
func (a *AutomationNode) Wait() error {
	payload := make(map[string]any)
	return a.Post(a.waitTarget, payload)
}

// AutomationNodeGroup gets the AutomationNodeGroup linked resources.
func (a *AutomationNode) AutomationNodeGroup(client common.Client) ([]*AutomationNode, error) {
	return common.GetObjects[AutomationNode](client, a.automationNodeGroup)
}

// Chassis gets the Chassis linked resources.
func (a *AutomationNode) Chassis(client common.Client) ([]*Chassis, error) {
	return common.GetObjects[Chassis](client, a.chassis)
}

// OutputControl gets the OutputControl linked resource.
func (a *AutomationNode) OutputControl(client common.Client) (*Control, error) {
	if a.outputControl == "" {
		return nil, nil
	}
	return common.GetObject[Control](client, a.outputControl)
}

// PidFeedbackSensor gets the PidFeedbackSensor linked resource.
func (a *AutomationNode) PidFeedbackSensor(client common.Client) (*Sensor, error) {
	if a.pidFeedbackSensor == "" {
		return nil, nil
	}
	return common.GetObject[Sensor](client, a.pidFeedbackSensor)
}

// PositionSensor gets the PositionSensor linked resource.
func (a *AutomationNode) PositionSensor(client common.Client) (*Sensor, error) {
	if a.positionSensor == "" {
		return nil, nil
	}
	return common.GetObject[Sensor](client, a.positionSensor)
}

// VelocitySensor gets the VelocitySensor linked resource.
func (a *AutomationNode) VelocitySensor(client common.Client) (*Sensor, error) {
	if a.velocitySensor == "" {
		return nil, nil
	}
	return common.GetObject[Sensor](client, a.velocitySensor)
}

// Instrumentation gets the Instrumentation linked resource.
func (a *AutomationNode) Instrumentation(client common.Client) (*AutomationInstrumentation, error) {
	if a.instrumentation == "" {
		return nil, nil
	}
	return common.GetObject[AutomationInstrumentation](client, a.instrumentation)
}
