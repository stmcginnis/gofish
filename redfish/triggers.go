//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
)

// DirectionOfCrossingEnum shall indicate the direction of crossing that corresponds to a trigger.
type DirectionOfCrossingEnum string

const (
	// IncreasingDirectionOfCrossingEnum A trigger condition is met when the metric value crosses the trigger value
	// while increasing.
	IncreasingDirectionOfCrossingEnum DirectionOfCrossingEnum = "Increasing"
	// DecreasingDirectionOfCrossingEnum A trigger is met when the metric value crosses the trigger value while
	// decreasing.
	DecreasingDirectionOfCrossingEnum DirectionOfCrossingEnum = "Decreasing"
)

// DiscreteTriggerConditionEnum shall specify the condition, in relationship to the discrete trigger
// values, which constitutes a trigger.
type DiscreteTriggerConditionEnum string

const (
	// SpecifiedDiscreteTriggerConditionEnum A discrete trigger condition is met when the metric value becomes one of
	// the values that the DiscreteTriggers property lists.
	SpecifiedDiscreteTriggerConditionEnum DiscreteTriggerConditionEnum = "Specified"
	// ChangedDiscreteTriggerConditionEnum A discrete trigger condition is met whenever the metric value changes.
	ChangedDiscreteTriggerConditionEnum DiscreteTriggerConditionEnum = "Changed"
)

// MetricTypeEnum shall specify the type of metric for which the trigger is configured.
type MetricTypeEnum string

const (
	// NumericMetricTypeEnum The trigger is for numeric sensor.
	NumericMetricTypeEnum MetricTypeEnum = "Numeric"
	// DiscreteMetricTypeEnum The trigger is for a discrete sensor.
	DiscreteMetricTypeEnum MetricTypeEnum = "Discrete"
)

// TriggerActionEnum shall specify the actions to perform when a trigger condition is met.
type TriggerActionEnum string

const (
	// LogToLogServiceTriggerActionEnum shall log the occurrence of the condition to the log that the LogService
	// property in the telemetry service resource describes.
	LogToLogServiceTriggerActionEnum TriggerActionEnum = "LogToLogService"
	// RedfishEventTriggerActionEnum shall send an event to subscribers.
	RedfishEventTriggerActionEnum TriggerActionEnum = "RedfishEvent"
	// RedfishMetricReportTriggerActionEnum shall force the metric reports managed by the MetricReportDefinitions
	// specified by the MetricReportDefinitions property to be updated, regardless of the MetricReportDefinitionType
	// property value. The actions specified in the ReportActions property of each MetricReportDefinition shall be
	// performed.
	RedfishMetricReportTriggerActionEnum TriggerActionEnum = "RedfishMetricReport"
)

// DiscreteTrigger shall contain the characteristics of the discrete trigger.
type DiscreteTrigger struct {
	common.Entity
	// DwellTime shall contain the amount of time that a trigger event persists before the TriggerActions are
	// performed.
	DwellTime string
	// Severity shall contain the Severity property to be used in the event message.
	Severity common.Health
	// Value shall contain the value discrete metric that constitutes a trigger event. The DwellTime shall be measured
	// from this point in time.
	Value string
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a DiscreteTrigger object from the raw JSON.
func (discretetrigger *DiscreteTrigger) UnmarshalJSON(b []byte) error {
	type temp DiscreteTrigger
	var t struct {
		temp
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*discretetrigger = DiscreteTrigger(t.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	discretetrigger.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (discretetrigger *DiscreteTrigger) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(DiscreteTrigger)
	original.UnmarshalJSON(discretetrigger.rawData)

	readWriteFields := []string{
		"DwellTime",
		"Severity",
		"Value",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(discretetrigger).Elem()

	return discretetrigger.Entity.Update(originalElement, currentElement, readWriteFields)
}

// Triggers shall contain a trigger that applies to metrics.
type Triggers struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// DiscreteTriggerCondition shall contain the conditions when a discrete metric triggers.
	DiscreteTriggerCondition DiscreteTriggerConditionEnum
	// DiscreteTriggers shall contain a list of values to which to compare a metric reading. This property shall be
	// present when the DiscreteTriggerCondition property is 'Specified'.
	DiscreteTriggers []DiscreteTrigger
	// EventTriggers shall contain an array of MessageIds that specify when a trigger condition is met based on an
	// event. When the service generates an event and if it contains a MessageId within this array, a trigger condition
	// shall be met. The MetricType property should not be present if this resource is configured for event-based
	// triggers.
	EventTriggers []string
	// HysteresisDuration shall indicate the duration the sensor value no longer violates the threshold before the
	// threshold is deactivated. A duration of zero seconds, or if the property is not present in the resource, shall
	// indicate the threshold is deactivated immediately once the sensor value no longer violates the threshold. The
	// threshold shall not deactivate until the conditions of both HysteresisReading and HysteresisDuration are met.
	HysteresisDuration string
	// HysteresisReading shall indicate the offset from the reading for this sensor and the threshold value that
	// deactivates the threshold. For example, a value of '-2' indicates the sensor reading shall fall 2 units below an
	// upper threshold value to deactivate the threshold. The value of the property shall use the same units as the
	// Reading property. A value of '0', or if the property is not present in the resource, shall indicate the
	// threshold is deactivated when the sensor value no longer violates the threshold. The threshold shall not
	// deactivate until the conditions of both HysteresisReading and HysteresisDuration are met.
	HysteresisReading float64
	// MetricIDs shall contain the labels for the metric definitions that contain the property identifiers for this
	// trigger. This property shall match the value of the Id property of the corresponding metric definitions.
	MetricIDs []string
	// MetricProperties shall contain an array of URIs with wildcards and property identifiers for this trigger. Use a
	// set of curly braces to delimit each wildcard in the URI. Replace each wildcard with its corresponding entry in
	// the Wildcard array property. A URI that contains wildcards shall link to a resource property to which the metric
	// definition applies after all wildcards are replaced with their corresponding entries in the Wildcard array
	// property. The property identifiers portion of the URI shall follow the RFC6901-defined JSON fragment notation
	// rules.
	MetricProperties []string
	// MetricType shall contain the metric type of the trigger.
	MetricType MetricTypeEnum
	// NumericThresholds shall contain the list of thresholds to which to compare a numeric metric value.
	NumericThresholds Thresholds
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// TriggerActions shall contain the actions that the trigger initiates.
	TriggerActions []TriggerActionEnum
	// Wildcards shall contain the wildcards and their substitution values for the entries in the MetricProperties
	// array property. Each wildcard shall have a corresponding entry in this array property.
	Wildcards []Wildcard
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte

	metricReportDefinitions []string
	// MetricReportDefinitionsCount is the number of MetricReportDefinitions.
	MetricReportDefinitionsCount int
}

// UnmarshalJSON unmarshals a Triggers object from the raw JSON.
func (triggers *Triggers) UnmarshalJSON(b []byte) error {
	type temp Triggers
	type Links struct {
		// MetricReportDefinitions shall contain a set of links to metric report definitions that generate new metric
		// reports when a trigger condition is met and when the TriggerActions property contains 'RedfishMetricReport'.
		MetricReportDefinitions common.Links
		// MetricReportDefinitions@odata.count
		MetricReportDefinitionsCount int `json:"MetricReportDefinitions@odata.count"`
		// Oem shall contain the OEM extensions. All values for properties contained in this object shall conform to the
		// Redfish Specification-described requirements.
		OEM json.RawMessage `json:"Oem"`
	}
	var t struct {
		temp
		Links Links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*triggers = Triggers(t.temp)

	// Extract the links to other entities for later
	triggers.metricReportDefinitions = t.Links.MetricReportDefinitions.ToStrings()
	triggers.MetricReportDefinitionsCount = t.Links.MetricReportDefinitionsCount

	// This is a read/write object, so we need to save the raw object data for later
	triggers.rawData = b

	return nil
}

// MetricReportDefinitions gets the metric report definitions that generate new metric
// reports when a trigger condition is met and when the TriggerActions property contains 'RedfishMetricReport'.
func (triggers *Triggers) MetricReportDefinitions() ([]*MetricReportDefinition, error) {
	return common.GetObjects[MetricReportDefinition](triggers.GetClient(), triggers.metricReportDefinitions)
}

// Update commits updates to this object's properties to the running system.
func (triggers *Triggers) Update() error {
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(Triggers)
	original.UnmarshalJSON(triggers.rawData)

	readWriteFields := []string{
		"EventTriggers",
		"HysteresisDuration",
		"HysteresisReading",
		"MetricIds",
		"MetricProperties",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(triggers).Elem()

	return triggers.Entity.Update(originalElement, currentElement, readWriteFields)
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
