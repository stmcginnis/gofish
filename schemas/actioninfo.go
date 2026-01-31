//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.3 - #ActionInfo.v1_5_0.ActionInfo

package schemas

import (
	"encoding/json"
	"fmt"
)

type ParameterTypes string

const (
	// BooleanParameterTypes is a boolean.
	BooleanParameterTypes ParameterTypes = "Boolean"
	// NumberParameterTypes is a number.
	NumberParameterTypes ParameterTypes = "Number"
	// NumberArrayParameterTypes is an array of numbers.
	NumberArrayParameterTypes ParameterTypes = "NumberArray"
	// StringParameterTypes is a string.
	StringParameterTypes ParameterTypes = "String"
	// StringArrayParameterTypes is an array of strings.
	StringArrayParameterTypes ParameterTypes = "StringArray"
	// ObjectParameterTypes is an embedded JSON object.
	ObjectParameterTypes ParameterTypes = "Object"
	// ObjectArrayParameterTypes is an array of JSON objects.
	ObjectArrayParameterTypes ParameterTypes = "ObjectArray"
)

// ActionInfo shall represent the supported parameters and other information for
// a Redfish action on a target within a Redfish implementation. Supported
// parameters can differ among vendors and even among resource instances. This
// data can ensure that action requests from applications contain supported
// parameters.
type ActionInfo struct {
	Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Parameters shall list the parameters included in the specified Redfish
	// action for this resource.
	Parameters []ActionInfoParameter
}

// GetActionInfo will get a ActionInfo instance from the service.
func GetActionInfo(c Client, uri string) (*ActionInfo, error) {
	return GetObject[ActionInfo](c, uri)
}

// ListReferencedActionInfos gets the collection of ActionInfo from
// a provided reference.
func ListReferencedActionInfos(c Client, link string) ([]*ActionInfo, error) {
	return GetCollectionObjects[ActionInfo](c, link)
}

func (a *ActionInfo) GetParamValues(name string, dataType ParameterTypes) ([]string, error) {
	for idx := range a.Parameters {
		param := &a.Parameters[idx]
		if param.Name != name || (param.DataType != "" && param.DataType != dataType) {
			continue
		}

		return param.AllowableValues, nil
	}
	return nil, fmt.Errorf("failed to find supported values of type %s", name)
}

// ActionInfoParameter shall contain information about a parameter included in a Redfish
// action for this resource.
type ActionInfoParameter struct {
	// AllowableNumbers shall indicate the allowable numeric values, inclusive
	// ranges of values, and incremental step values for this parameter as applied
	// to this action target, as defined in the 'Allowable values for numbers and
	// durations' clause of the Redfish Specification. For arrays, this property
	// shall represent the allowable values for each array member. This property
	// shall only be present for numeric parameters or string parameters that
	// specify a duration.
	//
	// Version added: v1.3.0
	AllowableNumbers []string
	// AllowablePattern shall contain a regular expression that describes the
	// allowable values for this parameter as applied to this action target. For
	// arrays, this property shall represent the allowable values for each array
	// member. This property shall only be present for string parameters.
	//
	// Version added: v1.3.0
	AllowablePattern string
	// AllowableValueDescriptions shall contain the descriptions of allowable
	// values for this parameter. The descriptions shall appear in the same array
	// order as the 'AllowableValues' property. For arrays, this property shall
	// represent the descriptions of allowable values for each array member.
	//
	// Version added: v1.4.0
	AllowableValueDescriptions []string
	// AllowableValues shall indicate the allowable values for this parameter as
	// applied to this action target. For arrays, this property shall represent the
	// allowable values for each array member.
	AllowableValues []string
	// ArraySizeMaximum shall contain the maximum number of array elements that
	// this service supports for this parameter. This property shall not be present
	// for non-array parameters.
	//
	// Version added: v1.2.0
	ArraySizeMaximum *int `json:",omitempty"`
	// ArraySizeMinimum shall contain the minimum number of array elements required
	// by this service for this parameter. This property shall not be present for
	// non-array parameters.
	//
	// Version added: v1.2.0
	ArraySizeMinimum *int `json:",omitempty"`
	// DataType shall contain the JSON property type for this parameter.
	DataType ParameterTypes
	// DefaultValue shall contain the default value for this parameter if the
	// client does not provide the parameter. This property shall not be present if
	// 'Required' contains 'true'. If 'DataType' does not contain 'String', the
	// service shall convert the value to an RFC8259-defined JSON string.
	//
	// Version added: v1.5.0
	DefaultValue string
	// MaximumValue shall contain the maximum value that this service supports. For
	// arrays, this property shall represent the maximum value for each array
	// member. This property shall not be present for non-integer or number
	// parameters.
	//
	// Version added: v1.1.0
	MaximumValue *float64 `json:",omitempty"`
	// MinimumValue shall contain the minimum value that this service supports. For
	// arrays, this property shall represent the minimum value for each array
	// member. This property shall not be present for non-integer or number
	// parameters.
	//
	// Version added: v1.1.0
	MinimumValue *float64 `json:",omitempty"`
	// Name is the name of the resource or array element.
	Name string
	// NoDefaultValue shall indicate that there is no default value for this
	// parameter. For example, if username and password parameters are optional,
	// the absence of the parameters indicates no credentials are used. This
	// property shall not be present if 'Required' contains 'true' or if
	// 'DefaultValue' is present.
	//
	// Version added: v1.5.0
	NoDefaultValue bool
	// ObjectDataType shall describe the entity type definition in '@odata.type'
	// format for the parameter. This property shall be required for parameters
	// with a data type of 'Object' or 'ObjectArray', and shall not be present for
	// parameters with other data types.
	ObjectDataType string
	// Required shall indicate whether the parameter is required to complete this
	// action.
	Required bool
}
