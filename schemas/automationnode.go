//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/AutomationNode.v1_0_0.json
// 2025.2 - #AutomationNode.v1_0_0.AutomationNode

package schemas

import (
	"encoding/json"
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
	Entity
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
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of a resource.
	Status Status
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
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a AutomationNode object from the raw JSON.
func (a *AutomationNode) UnmarshalJSON(b []byte) error {
	type temp AutomationNode
	type aActions struct {
		Reset       ActionTarget `json:"#AutomationNode.Reset"`
		SendTrigger ActionTarget `json:"#AutomationNode.SendTrigger"`
		Start       ActionTarget `json:"#AutomationNode.Start"`
		Stop        ActionTarget `json:"#AutomationNode.Stop"`
		Wait        ActionTarget `json:"#AutomationNode.Wait"`
	}
	type aLinks struct {
		AutomationNodeGroup Links `json:"AutomationNodeGroup"`
		Chassis             Links `json:"Chassis"`
		OutputControl       Link  `json:"OutputControl"`
		PidFeedbackSensor   Link  `json:"PidFeedbackSensor"`
		PositionSensor      Link  `json:"PositionSensor"`
		VelocitySensor      Link  `json:"VelocitySensor"`
	}
	var tmp struct {
		temp
		Actions         aActions
		Links           aLinks
		Instrumentation Link `json:"Instrumentation"`
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
	a.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (a *AutomationNode) Update() error {
	readWriteFields := []string{
		"MotionAxis",
		"MotionProfile",
	}

	return a.UpdateFromRawData(a, a.RawData, readWriteFields)
}

// GetAutomationNode will get a AutomationNode instance from the service.
func GetAutomationNode(c Client, uri string) (*AutomationNode, error) {
	return GetObject[AutomationNode](c, uri)
}

// ListReferencedAutomationNodes gets the collection of AutomationNode from
// a provided reference.
func ListReferencedAutomationNodes(c Client, link string) ([]*AutomationNode, error) {
	return GetCollectionObjects[AutomationNode](c, link)
}

// This action shall reset the node to its power-on state.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (a *AutomationNode) Reset() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(a.client,
		a.resetTarget, payload, a.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall pulse the trigger signal for the node.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (a *AutomationNode) SendTrigger() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(a.client,
		a.sendTriggerTarget, payload, a.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall transition the node to the 'Running' state. This action
// shall pulse the trigger signal to other connected nodes that are in the
// 'Waiting' state.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (a *AutomationNode) Start() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(a.client,
		a.startTarget, payload, a.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall transition the node to the 'Idle' state.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (a *AutomationNode) Stop() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(a.client,
		a.stopTarget, payload, a.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// This action shall transition the node to the 'Waiting' state.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (a *AutomationNode) Wait() (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	resp, taskInfo, err := PostWithTask(a.client,
		a.waitTarget, payload, a.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// AutomationNodeGroup gets the AutomationNodeGroup linked resources.
func (a *AutomationNode) AutomationNodeGroup() ([]*AutomationNode, error) {
	return GetObjects[AutomationNode](a.client, a.automationNodeGroup)
}

// Chassis gets the Chassis linked resources.
func (a *AutomationNode) Chassis() ([]*Chassis, error) {
	return GetObjects[Chassis](a.client, a.chassis)
}

// OutputControl gets the OutputControl linked resource.
func (a *AutomationNode) OutputControl() (*Control, error) {
	if a.outputControl == "" {
		return nil, nil
	}
	return GetObject[Control](a.client, a.outputControl)
}

// PidFeedbackSensor gets the PidFeedbackSensor linked resource.
func (a *AutomationNode) PidFeedbackSensor() (*Sensor, error) {
	if a.pidFeedbackSensor == "" {
		return nil, nil
	}
	return GetObject[Sensor](a.client, a.pidFeedbackSensor)
}

// PositionSensor gets the PositionSensor linked resource.
func (a *AutomationNode) PositionSensor() (*Sensor, error) {
	if a.positionSensor == "" {
		return nil, nil
	}
	return GetObject[Sensor](a.client, a.positionSensor)
}

// VelocitySensor gets the VelocitySensor linked resource.
func (a *AutomationNode) VelocitySensor() (*Sensor, error) {
	if a.velocitySensor == "" {
		return nil, nil
	}
	return GetObject[Sensor](a.client, a.velocitySensor)
}

// Instrumentation gets the Instrumentation linked resource.
func (a *AutomationNode) Instrumentation() (*AutomationInstrumentation, error) {
	if a.instrumentation == "" {
		return nil, nil
	}
	return GetObject[AutomationInstrumentation](a.client, a.instrumentation)
}
