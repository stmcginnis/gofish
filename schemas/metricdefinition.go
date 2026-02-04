//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/MetricDefinition.v1_3_5.json
// 2022.1 - #MetricDefinition.v1_3_5.MetricDefinition

package schemas

import (
	"encoding/json"
)

// Calculable is The type shall describe the types of calculations that can be
// applied to the metric reading.
type Calculable string

const (
	// NonCalculatableCalculable No calculations should be performed on the metric
	// reading.
	NonCalculatableCalculable Calculable = "NonCalculatable"
	// SummableCalculable The sum of the metric reading across multiple instances
	// is meaningful.
	SummableCalculable Calculable = "Summable"
	// NonSummableCalculable The sum of the metric reading across multiple
	// instances is not meaningful.
	NonSummableCalculable Calculable = "NonSummable"
)

type CalculationAlgorithmEnum string

const (
	// AverageCalculationAlgorithmEnum shall be calculated as the average metric
	// reading over a sliding time interval. The time interval shall contain the
	// 'CalculationTimeInterval' property value.
	AverageCalculationAlgorithmEnum CalculationAlgorithmEnum = "Average"
	// MaximumCalculationAlgorithmEnum shall be calculated as the maximum metric
	// reading over a sliding time interval. The time interval shall contain the
	// 'CalculationTimeInterval' property value.
	MaximumCalculationAlgorithmEnum CalculationAlgorithmEnum = "Maximum"
	// MinimumCalculationAlgorithmEnum shall be calculated as the minimum metric
	// reading over a sliding time interval. The time interval shall contain the
	// 'CalculationTimeInterval' property value.
	MinimumCalculationAlgorithmEnum CalculationAlgorithmEnum = "Minimum"
	// OEMCalculationAlgorithmEnum shall be calculated as specified by an OEM. The
	// 'OEMCalculationAlgorithm' property shall contain the specific OEM
	// calculation algorithm.
	OEMCalculationAlgorithmEnum CalculationAlgorithmEnum = "OEM"
	// SummationCalculationAlgorithmEnum shall indicate the metric is calculated as
	// the sum of the specified metric reading over a duration. The duration shall
	// be the 'CollectionDuration' property value.
	SummationCalculationAlgorithmEnum CalculationAlgorithmEnum = "Summation" // from metricreportdefinition.go
)

type MetricDefinitionImplementationType string

const (
	// PhysicalSensorMetricDefinitionImplementationType The metric is implemented as a physical
	// sensor.
	PhysicalSensorMetricDefinitionImplementationType MetricDefinitionImplementationType = "PhysicalSensor"
	// CalculatedMetricDefinitionImplementationType The metric is implemented by applying a
	// calculation on another metric property. The calculation is specified in the
	// 'CalculationAlgorithm' property.
	CalculatedMetricDefinitionImplementationType MetricDefinitionImplementationType = "Calculated"
	// SynthesizedMetricDefinitionImplementationType The metric is implemented by applying a
	// calculation on one or more metric properties. The calculation is not
	// provided.
	SynthesizedMetricDefinitionImplementationType MetricDefinitionImplementationType = "Synthesized"
	// DigitalMeterMetricDefinitionImplementationType The metric is implemented as digital meter.
	DigitalMeterMetricDefinitionImplementationType MetricDefinitionImplementationType = "DigitalMeter"
)

// MetricDataType is This type shall describe the data type of the related
// metric values as defined by JSON data types.
type MetricDataType string

const (
	// BooleanMetricDataType The JSON boolean definition.
	BooleanMetricDataType MetricDataType = "Boolean"
	// DateTimeMetricDataType The JSON string definition with the date-time format.
	DateTimeMetricDataType MetricDataType = "DateTime"
	// DecimalMetricDataType The JSON decimal definition.
	DecimalMetricDataType MetricDataType = "Decimal"
	// IntegerMetricDataType The JSON integer definition.
	IntegerMetricDataType MetricDataType = "Integer"
	// StringMetricDataType The JSON string definition.
	StringMetricDataType MetricDataType = "String"
	// EnumerationMetricDataType The JSON string definition with a set of defined
	// enumerations.
	EnumerationMetricDataType MetricDataType = "Enumeration"
)

// MetricType is This property shall contain the type of metric.
type MetricType string

const (
	// NumericMetricType The metric is a numeric metric. The metric value is any
	// real number.
	NumericMetricType MetricType = "Numeric"
	// DiscreteMetricType shall indicate discrete states.
	DiscreteMetricType MetricType = "Discrete"
	// GaugeMetricType The metric is a gauge metric. The metric value is a real
	// number. When the metric value reaches the gauge's extrema, it stays at that
	// value, until the reading falls within the extrema.
	GaugeMetricType MetricType = "Gauge"
	// CounterMetricType The metric is a counter metric. The metric reading is a
	// non-negative integer that increases monotonically. When a counter reaches
	// its maximum, the value resets to 0 and resumes counting.
	CounterMetricType MetricType = "Counter"
	// CountdownMetricType The metric is a countdown metric. The metric reading is
	// a non-negative integer that decreases monotonically. When a counter reaches
	// its minimum, the value resets to preset value and resumes counting down.
	CountdownMetricType MetricType = "Countdown"
	// StringMetricType The metric is a non-discrete string metric. The metric
	// reading is a non-discrete string that displays some non-discrete,
	// non-numeric data.
	StringMetricType MetricType = "String"
)

// MetricDefinition shall contain the metadata information for a metric in a
// Redfish implementation.
type MetricDefinition struct {
	Entity
	// Accuracy shall contain the percent error +/- of the measured versus actual
	// values. The property is not meaningful when the 'MetricType' property is
	// 'Discrete'.
	Accuracy *float64 `json:",omitempty"`
	// Calculable shall specify whether the metric can be used in a calculation.
	Calculable Calculable
	// CalculationAlgorithm shall contain the calculation performed to obtain the
	// metric.
	CalculationAlgorithm CalculationAlgorithmEnum
	// CalculationParameters shall list the metric properties that are part of a
	// calculation that this metric definition defines. This property should be
	// present if 'MetricDefinitionImplementationType' contains 'Synthesized' or 'Calculated'.
	CalculationParameters []CalculationParamsType
	// CalculationTimeInterval shall specify the time interval over the metric
	// calculation is performed.
	CalculationTimeInterval string
	// Calibration shall contain the calibration offset added to the metric
	// reading. The value shall have the units specified in the 'Units' property.
	// The property is not meaningful when the 'MetricType' property is 'Discrete'.
	Calibration *float64 `json:",omitempty"`
	// DiscreteValues shall specify the possible values of the discrete metric.
	// This property shall have values when the 'MetricType' property is
	// 'Discrete'.
	DiscreteValues []string
	// Implementation shall specify the implementation of the metric.
	Implementation MetricDefinitionImplementationType
	// IsLinear shall indicate whether the metric values are linear versus
	// non-linear. Linear metrics can use a greater than relation to compared them.
	// An example of linear metrics include performance metrics. Examples of
	// non-linear metrics include error codes.
	IsLinear bool
	// LogicalContexts shall contain the logical contexts related to the metric.
	// This property should be present when the 'PhysicalContext' property does not
	// provide complete information and additional context information is needed.
	// For example, if the metric refers to capacity or performance.
	//
	// Version added: v1.3.0
	LogicalContexts []LogicalContext
	// MaxReadingRange shall indicate the highest possible value for a related
	// MetricValue. The value shall have the units specified in the property Units.
	// The property is not meaningful when the 'MetricType' property is 'Discrete'.
	MaxReadingRange *float64 `json:",omitempty"`
	// MetricDataType shall specify the data-type of the metric.
	MetricDataType MetricDataType
	// MetricProperties shall list the URIs with wildcards and property identifiers
	// that this metric defines. A set of curly braces shall delimit each wildcard
	// in the URI. The corresponding entry in the 'Wildcard' property shall replace
	// each wildcard. After each wildcard is replaced, it shall identify a resource
	// property to which the metric definition applies. The property identifiers
	// portion of the URI shall follow RFC6901-defined JSON pointer notation rules.
	// This property should not be present if 'MetricDefinitionImplementationType' contains
	// 'Synthesized' or 'Calculated'.
	MetricProperties []string
	// MetricType shall specify the type of metric.
	MetricType MetricType
	// MinReadingRange shall contain the lowest possible value for the metric
	// reading. The value shall have the units specified in the property Units. The
	// property is not meaningful when the 'MetricType' property is 'Discrete'.
	MinReadingRange *float64 `json:",omitempty"`
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// OEMCalculationAlgorithm shall contain the OEM-defined calculation performed
	// to obtain the metric. This property shall be present if
	// 'CalculationAlgorithm' is 'OEM'.
	//
	// Version added: v1.1.0
	OEMCalculationAlgorithm string
	// PhysicalContext shall contain the physical context of the metric.
	PhysicalContext PhysicalContext
	// Precision shall specify the number of significant digits in the metric
	// reading. The property is not meaningful when the 'MetricType' property is
	// 'Discrete'.
	Precision *int `json:",omitempty"`
	// SensingInterval shall specify the time interval between when a metric is
	// updated.
	SensingInterval string
	// TimestampAccuracy shall specify the expected + or - variability of the
	// timestamp.
	TimestampAccuracy string
	// Units shall specify the units of the metric. This property shall be
	// consistent with the case-sensitive ('C/s' column) Unified Code for Units of
	// Measure. Note: Not all units of measure are covered by UCUM.
	Units string
	// Wildcards shall contain a list of wildcards and their replacement strings,
	// which are applied to the 'MetricProperties' array property. Each wildcard
	// shall have a corresponding entry in this array property.
	Wildcards []Wildcard
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a MetricDefinition object from the raw JSON.
func (m *MetricDefinition) UnmarshalJSON(b []byte) error {
	type temp MetricDefinition
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = MetricDefinition(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	m.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (m *MetricDefinition) Update() error {
	readWriteFields := []string{
		"Calculable",
		"CalculationTimeInterval",
		"DiscreteValues",
		"IsLinear",
		"MetricDataType",
		"MetricProperties",
		"MetricType",
		"SensingInterval",
		"Units",
	}

	return m.UpdateFromRawData(m, m.RawData, readWriteFields)
}

// GetMetricDefinition will get a MetricDefinition instance from the service.
func GetMetricDefinition(c Client, uri string) (*MetricDefinition, error) {
	return GetObject[MetricDefinition](c, uri)
}

// ListReferencedMetricDefinitions gets the collection of MetricDefinition from
// a provided reference.
func ListReferencedMetricDefinitions(c Client, link string) ([]*MetricDefinition, error) {
	return GetCollectionObjects[MetricDefinition](c, link)
}

// CalculationParamsType shall contain the parameters for a metric calculation.
type CalculationParamsType struct {
	// ResultMetric shall contain a URI with wildcards and property identifiers of
	// the metric property that stores the result of the calculation. A set of
	// curly braces shall delimit each wildcard in the URI. The corresponding entry
	// in the 'Wildcard' property shall replace each wildcard. After each wildcard
	// is replaced, it shall identify a resource property to which the metric
	// definition applies. The property identifiers portion of the URI shall follow
	// RFC6901-defined JSON pointer notation rules.
	ResultMetric string
	// SourceMetric shall contain a URI with wildcards and property identifiers of
	// the metric property used as the input into the calculation. A set of curly
	// braces shall delimit each wildcard in the URI. The corresponding entry in
	// the 'Wildcard' property shall replace each wildcard. After each wildcard is
	// replaced, it shall identify a resource property to which the metric
	// definition applies. The property identifiers portion of the URI shall follow
	// RFC6901-defined JSON pointer notation rules.
	SourceMetric string
}

// Wildcard shall contain a wildcard and its substitution values.
type Wildcard struct {
	// Name is the name of the resource or array element.
	Name string
	// Values shall contain the list of values to substitute for the wildcard. The
	// value '*' shall indicate all possible values for the wildcard.
	Values []string
}
