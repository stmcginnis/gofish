//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2024.1 - #Triggers.v1_4_0.Triggers

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// DirectionOfCrossingEnum is The value shall indicate the direction of crossing
// that corresponds to a trigger.
type DirectionOfCrossingEnum string

const (
	// IncreasingDirectionOfCrossingEnum is a trigger condition is met when the
	// metric value crosses the trigger value while increasing.
	IncreasingDirectionOfCrossingEnum DirectionOfCrossingEnum = "Increasing"
	// DecreasingDirectionOfCrossingEnum is a trigger is met when the metric value
	// crosses the trigger value while decreasing.
	DecreasingDirectionOfCrossingEnum DirectionOfCrossingEnum = "Decreasing"
)

// DiscreteTriggerConditionEnum is This type shall specify the condition, in
// relationship to the discrete trigger values, which constitutes a trigger.
type DiscreteTriggerConditionEnum string

const (
	// SpecifiedDiscreteTriggerConditionEnum is a discrete trigger condition is met
	// when the metric value becomes one of the values that the 'DiscreteTriggers'
	// property lists.
	SpecifiedDiscreteTriggerConditionEnum DiscreteTriggerConditionEnum = "Specified"
	// ChangedDiscreteTriggerConditionEnum is a discrete trigger condition is met
	// whenever the metric value changes.
	ChangedDiscreteTriggerConditionEnum DiscreteTriggerConditionEnum = "Changed"
)

// MetricTypeEnum is This type shall specify the type of metric for which the
// trigger is configured.
type MetricTypeEnum string

const (
	// NumericMetricTypeEnum is a numeric value trigger.
	NumericMetricTypeEnum MetricTypeEnum = "Numeric"
	// DiscreteMetricTypeEnum is a discrete value trigger.
	DiscreteMetricTypeEnum MetricTypeEnum = "Discrete"
)

type ThresholdActivation string

const (
	// IncreasingThresholdActivation This threshold is activated when the reading
	// changes from a value lower than the threshold to a value higher than the
	// threshold.
	IncreasingThresholdActivation ThresholdActivation = "Increasing"
	// DecreasingThresholdActivation This threshold is activated when the reading
	// changes from a value higher than the threshold to a value lower than the
	// threshold.
	DecreasingThresholdActivation ThresholdActivation = "Decreasing"
	// EitherThresholdActivation This threshold is activated when either the
	// 'Increasing' or 'Decreasing' conditions are met.
	EitherThresholdActivation ThresholdActivation = "Either"
	// DisabledThresholdActivation shall indicate the threshold is disabled and no
	// actions shall be taken as a result of the reading crossing the threshold
	// value.
	DisabledThresholdActivation ThresholdActivation = "Disabled"
)

// TriggerActionEnum is This type shall specify the actions to perform when a
// trigger condition is met.
type TriggerActionEnum string

const (
	// LogToLogServiceTriggerActionEnum shall log the occurrence of the condition
	// to the log that the 'LogService' property in the telemetry service resource
	// describes. The message for the created log entry shall follow the guidance
	// specified by the 'TriggerActionMessage' property.
	LogToLogServiceTriggerActionEnum TriggerActionEnum = "LogToLogService"
	// RedfishEventTriggerActionEnum shall send an event to subscribers. The
	// message key for the event shall follow the guidance specified by
	// TriggerActionMessage.
	RedfishEventTriggerActionEnum TriggerActionEnum = "RedfishEvent"
	// RedfishMetricReportTriggerActionEnum shall force the metric reports managed
	// by the metric report definitions specified by the 'MetricReportDefinitions'
	// property to be updated, regardless of the 'MetricReportDefinitionType'
	// property value. The actions specified in the 'ReportActions' property of
	// each 'MetricReportDefinition' resource shall be performed.
	RedfishMetricReportTriggerActionEnum TriggerActionEnum = "RedfishMetricReport"
)

// TriggerActionMessage is The value shall indicate the message used to complete
// the specified trigger actions.
type TriggerActionMessage string

const (
	// TelemetryTriggerActionMessage shall indicate that messages generated in
	// response to a trigger action shall utilize messages from the Telemetry
	// Message Registry. If this property is not supplied or supported, this value
	// should be used as the default for this trigger.
	TelemetryTriggerActionMessage TriggerActionMessage = "Telemetry"
	// DriveMediaLifeTriggerActionMessage shall indicate that messages generated in
	// response to a trigger action shall utilize the 'MediaLifeLeftLow' message
	// from the Storage Device Message Registry.
	DriveMediaLifeTriggerActionMessage TriggerActionMessage = "DriveMediaLife"
	// ConnectionSpeedTriggerActionMessage shall indicate that messages generated
	// in response to a trigger action shall utilize the 'ConnectionSpeedLow'
	// message from the Network Device Message Registry.
	ConnectionSpeedTriggerActionMessage TriggerActionMessage = "ConnectionSpeed"
)

// Triggers shall contain a trigger condition that applies to metrics.
type Triggers struct {
	common.Entity
	// DiscreteTriggerCondition shall contain the conditions when a discrete metric
	// triggers.
	DiscreteTriggerCondition DiscreteTriggerConditionEnum
	// DiscreteTriggers shall contain a list of values to which to compare a metric
	// reading. This property shall be present when the 'DiscreteTriggerCondition'
	// property is 'Specified'.
	DiscreteTriggers []DiscreteTrigger
	// EventTriggers shall contain an array of 'MessageId' values that specify when
	// a trigger condition is met based on an event. When the service generates an
	// event and if it contains a 'MessageId' within this array, a trigger
	// condition shall be met. The 'MetricType' property should not be present if
	// this resource is configured for event-based triggers.
	//
	// Version added: v1.1.0
	EventTriggers []string
	// HysteresisDuration shall indicate the duration the metric value no longer
	// violates the threshold before the threshold is deactivated. A duration of
	// zero seconds, or if the property is not present in the resource, shall
	// indicate the threshold is deactivated immediately once the metric value no
	// longer violates the threshold. The threshold shall not deactivate until the
	// conditions of both 'HysteresisReading' and 'HysteresisDuration' are met.
	//
	// Version added: v1.3.0
	HysteresisDuration string
	// HysteresisReading shall indicate the offset from the reading for this sensor
	// and the threshold value that deactivates the threshold. For example, a value
	// of '-2' indicates the metric reading shall fall 2 units below an upper
	// threshold value to deactivate the threshold. The value of the property shall
	// use the same units as the 'Reading' property. A value of '0', or if the
	// property is not present in the resource, shall indicate the threshold is
	// deactivated when the metric value no longer violates the threshold. The
	// threshold shall not deactivate until the conditions of both
	// 'HysteresisReading' and 'HysteresisDuration' are met.
	//
	// Version added: v1.3.0
	HysteresisReading *float64 `json:",omitempty"`
	// MetricIDs shall contain the labels for the metric definitions that contain
	// the property identifiers for this trigger. This property shall match the
	// value of the 'Id' property of the corresponding metric definitions.
	//
	// Version added: v1.2.0
	MetricIDs []string
	// MetricProperties shall contain an array of URIs with wildcards and property
	// identifiers for this trigger. Use a set of curly braces to delimit each
	// wildcard in the URI. Replace each wildcard with its corresponding entry in
	// the 'Wildcard' array property. A URI that contains wildcards shall link to a
	// resource property to which this trigger definition applies after all
	// wildcards are replaced with their corresponding entries in the Wildcard
	// property. The property identifiers portion of the URI shall follow the
	// RFC6901-defined JSON fragment notation rules.
	MetricProperties []string
	// MetricType shall contain the metric type of the trigger.
	MetricType MetricTypeEnum
	// NumericThresholds shall contain the list of thresholds to which to compare a
	// numeric metric value.
	NumericThresholds Thresholds
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TriggerActionMessage shall contain the message definition used to generate a
	// Redfish event or a log entry as requested by the values of 'TriggerActions'.
	//
	// Version added: v1.4.0
	TriggerActionMessage TriggerActionMessage
	// TriggerActions shall contain the actions that the trigger initiates.
	TriggerActions []TriggerActionEnum
	// TriggerEnabled shall indicate whether the trigger is enabled. If 'true', it
	// is enabled. If 'false', it is disabled and none of the actions listed in
	// 'TriggerActions' will occur.
	//
	// Version added: v1.4.0
	TriggerEnabled bool
	// Wildcards shall contain the wildcards and their substitution values for the
	// entries in the 'MetricProperties' array property. Each wildcard shall have a
	// corresponding entry in this array property.
	Wildcards []Wildcard
	// metricReportDefinitions are the URIs for MetricReportDefinitions.
	metricReportDefinitions []string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a Triggers object from the raw JSON.
func (t *Triggers) UnmarshalJSON(b []byte) error {
	type temp Triggers
	type tLinks struct {
		MetricReportDefinitions common.Links `json:"MetricReportDefinitions"`
	}
	var tmp struct {
		temp
		Links tLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = Triggers(tmp.temp)

	// Extract the links to other entities for later
	t.metricReportDefinitions = tmp.Links.MetricReportDefinitions.ToStrings()

	// This is a read/write object, so we need to save the raw object data for later
	t.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (t *Triggers) Update() error {
	readWriteFields := []string{
		"DiscreteTriggers",
		"EventTriggers",
		"HysteresisDuration",
		"HysteresisReading",
		"MetricIDs",
		"MetricProperties",
		"NumericThresholds",
		"Status",
		"TriggerActionMessage",
		"TriggerEnabled",
		"Wildcards",
	}

	return t.UpdateFromRawData(t, t.rawData, readWriteFields)
}

// GetTriggers will get a Triggers instance from the service.
func GetTriggers(c common.Client, uri string) (*Triggers, error) {
	return common.GetObject[Triggers](c, uri)
}

// ListReferencedTriggerss gets the collection of Triggers from
// a provided reference.
func ListReferencedTriggerss(c common.Client, link string) ([]*Triggers, error) {
	return common.GetCollectionObjects[Triggers](c, link)
}

// MetricReportDefinitions gets the MetricReportDefinitions linked resources.
func (t *Triggers) MetricReportDefinitions(client common.Client) ([]*MetricReportDefinition, error) {
	return common.GetObjects[MetricReportDefinition](client, t.metricReportDefinitions)
}

// DiscreteTrigger shall contain the characteristics of the discrete trigger.
type DiscreteTrigger struct {
	// DwellTime shall contain the amount of time that a trigger event persists
	// before the 'TriggerActions' are performed.
	DwellTime string
	// Name is the name of the resource or array element.
	Name string
	// Severity shall contain the 'Severity' property to be used in the event
	// message.
	Severity common.Health
	// Value shall contain the value discrete metric that constitutes a trigger
	// event. The 'DwellTime' shall be measured from this point in time.
	Value string
}

// Thresholds shall contain a set of thresholds for a metric.
type Thresholds struct {
	// LowerCritical shall contain the value at which the 'MetricProperties'
	// property is below the normal range and may require attention. The value of
	// the property shall use the same units as the 'MetricProperties' property.
	LowerCritical Threshold
	// LowerWarning shall contain the value at which the 'MetricProperties'
	// property is below the normal range. The value of the property shall use the
	// same units as the 'MetricProperties' property.
	LowerWarning Threshold
	// UpperCritical shall contain the value at which the 'MetricProperties'
	// property is above the normal range and may require attention. The value of
	// the property shall use the same units as the 'MetricProperties' property.
	UpperCritical Threshold
	// UpperWarning shall contain the value at which the 'MetricProperties'
	// property is above the normal range. The value of the property shall use the
	// same units as the 'MetricProperties' property.
	UpperWarning Threshold
}

// Wildcard shall contain a wildcard and its substitution values.
type Wildcard struct {
	// Name is the name of the resource or array element.
	Name string
	// Values shall contain the list of values to substitute for the wildcard. A
	// single value of '*' shall indicate that the wildcard matches any available
	// values when substituted for a URI segment. If this property is not present,
	// the value shall be assumed to be '*'.
	Values []string
}
