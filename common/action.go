package common

import (
	"encoding/json"
	"fmt"
)

// ActionTarget is contains the target endpoint for object Actions.
//
// Deprecated: Action should be used instead.
type ActionTarget struct {
	Target string
}

// Action contains the target and ActionInfo endpoints for object Actions.
type Action struct {
	Target string
	Info   string `json:"@Redfish.ActionInfo"`
}

type ParameterDataType string

const (
	ParameterDataTypeBoolean     ParameterDataType = "Boolean"
	ParameterDataTypeNumber      ParameterDataType = "Number"
	ParameterDataTypeNumberArray ParameterDataType = "NumberArray"
	ParameterDataTypeString      ParameterDataType = "String"
	ParameterDataTypeStringArray ParameterDataType = "StringArray"
	ParameterDataTypeObject      ParameterDataType = "Object"
	ParameterDataTypeObjectArray ParameterDataType = "ObjectArray"
)

// UnmarshalJSON unmarshals a DataType
func (pdt *ParameterDataType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t := ParameterDataType(s)

	switch t {
	case ParameterDataTypeBoolean,
		ParameterDataTypeNumber,
		ParameterDataTypeNumberArray,
		ParameterDataTypeString,
		ParameterDataTypeStringArray,
		ParameterDataTypeObject,
		ParameterDataTypeObjectArray:
		*pdt = t
		return nil
	default:
		return fmt.Errorf("Unknown DataType specified: %s", s)
	}
}

type Parameter struct {
	Name            string // required
	Required        bool
	DataType        ParameterDataType
	ObjectDataType  string
	AllowableValues []string
}

// UnmarshalJSON unmarshals an ActionInfo
func (p *Parameter) UnmarshalJSON(b []byte) error {
	var t struct {
		Name            string // required
		Required        bool
		DataType        ParameterDataType
		ObjectDataType  string
		AllowableValues []string
	}
	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	if t.Name == "" {
		return fmt.Errorf("Parameter required field Name is missing")
	}

	*p = t
	return nil
}

type ActionInfo struct {
	ResourceCollection // Name and Id fields are required
	Parameters         []Parameter
}

// UnmarshalJSON unmarshals an ActionInfo
func (ai *ActionInfo) UnmarshalJSON(b []byte) error {
	var t struct {
		ResourceCollection
		Parameters []Parameter
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	if t.Name == "" {
		return fmt.Errorf("ActionInfo required field Name is missing")
	}
	if t.ID == "" {
		return fmt.Errorf("ActionInfo required field ID is missing")
	}

	*ai = t
	return nil
}
