//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"fmt"

	"github.com/stmcginnis/gofish/common"
)

// ActionInfoDataTypes is the datatype for an ActionInfo value.
type ActionInfoDataTypes string

const (
	BooleanActionInfoDataTypes     ActionInfoDataTypes = "Boolean"
	NumberActionInfoDataTypes      ActionInfoDataTypes = "Number"
	NumberArrayActionInfoDataTypes ActionInfoDataTypes = "NumberArray"
	ObjectActionInfoDataTypes      ActionInfoDataTypes = "Object"
	ObjectArrayActionInfoDataTypes ActionInfoDataTypes = "ObjectArray"
	StringActionInfoDataTypes      ActionInfoDataTypes = "String"
	StringArrayActionInfoDataTypes ActionInfoDataTypes = "StringArray"
)

type ActionInfo struct {
	common.Resource
	// Parameters is the list of parameters included in the specified Redfish action.
	// This property shall list the parameters included in the specified Redfish
	// action for this resource.
	Parameters []ActionInfoParameter
}

type ActionInfoParameter struct {
	// AllowableNumbers are the allowable numeric values or duration values, inclusive ranges of values,
	// and incremental step values for this parameter as applied to this action target.
	AllowableNumbers []string
	// AllowablePattern shall contain a regular expression that describes the allowable values for this
	// parameter as applied to this action target.
	AllowablePattern string
	// AllowableValueDescriptions shall contain the descriptions of allowable values for this parameter.
	AllowableValueDescriptions []string
	// AllowableValues shall indicate the allowable values for this parameter as applied to this action target.
	AllowableValues []string
	// ArraySizeMaximum shall contain the maximum number of array elements that this service supports for this parameter.
	ArraySizeMaximum uint64
	// ArraySizeMinimum shall contain the minimum number of array elements required by this service for this parameter.
	ArraySizeMinimum uint64
	// DataType shall contain the JSON property type for this parameter.
	DataType ActionInfoDataTypes
	// MaximumValue integer or number property shall contain the maximum value that this service supports.
	MaximumValue string
	// MinimumValue integer or number property shall contain the minimum value that this service supports.
	MinimumValue string
	// Name shall contain the name of the parameter included in a Redfish action.
	Name string
	// ObjectDataType shall describe the entity type definition in @odata.type format for the parameter.
	ObjectDataType string
	// Required shall indicate whether the parameter is required to complete this action.
	Required bool
}

func GetActionInfo(c common.Client, uri string) (*ActionInfo, error) {
	return common.GetObject[ActionInfo](c, uri)
}

func (actionInfo *ActionInfo) GetParamValues(name string, dataType ActionInfoDataTypes) ([]string, error) {
	for idx := range actionInfo.Parameters {
		param := &actionInfo.Parameters[idx]
		if param.Name != name || (param.DataType != "" && param.DataType != dataType) {
			continue
		}

		return param.AllowableValues, nil
	}
	return nil, fmt.Errorf("failed to find supported values of type %s", name)
}
