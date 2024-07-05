//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"math/big"

	"github.com/stmcginnis/gofish/common"
)

type AttributeType string

const (
	// A flag with a true or false value.
	BooleanAttributeType AttributeType = "Boolean"
	// A list of the known possible enumerated values.
	EnumerationAttributeType AttributeType = "Enumeration"
	// An integer value.
	IntegerAttributeType AttributeType = "Integer"
	// Password values that do not appear as plain text.
	// The value shall be null in responses.
	PasswordAttributeType AttributeType = "Password"
	// Free-form text in their values.
	StringAttributeType AttributeType = "String"
)

// AttributeValue represents the possible values for enumerated attribute values.
type AttributeValue struct {
	// A user-readable display string of the value for the attribute in the defined language.
	ValueDisplayName string
	// The unique value name for the attribute.
	ValueName string
}

// Attribute represents attributes and their possible values in the attribute registry.
type Attribute struct {
	// The unique name for the attribute.
	AttributeName string
	// The placeholder of the current value for the attribute.
	CurrentValue interface{}
	// The default value for the attribute.
	DefaultValue interface{}
	// The user-readable display string for the attribute
	// in the defined language.
	DisplayName string
	// The ascending order, as a number, in which this attribute appears
	// relative to other attributes.
	DisplayOrder int64
	// An indication of whether this attribute is grayed out.
	GrayOut bool
	// The help text for the attribute.
	HelpText string
	// An indication of whether this attribute is hidden in user interfaces.
	Hidden bool
	// An indication of whether this attribute is immutable.
	Immutable bool
	// An indication of whether this attribute is unique for this system
	// and should not be replicated.
	IsSystemUniqueProperty bool
	// The lower limit for an integer attribute.
	LowerBound int64
	// The maximum character length of a string attribute.
	MaxLength int64
	// The path that describes the menu hierarchy of this attribute.
	MenuPath string
	// The minimum character length of the string attribute.
	MinLength int64
	// Oem contains all the vendor specific information.
	Oem json.RawMessage
	// An indication of whether this attribute is read-only.
	ReadOnly bool
	// An indication of whether a system or device reset is required
	// for this attribute value change to take effect.
	ResetRequired bool
	// The amount to increment or decrement an integer attribute
	// each time a user requests a value change.
	ScalarIncrement int64
	// The attribute type.
	Type AttributeType
	// The UEFI device path that qualifies this attribute.
	UefiDevicePath string
	// The UEFI keyword string for this attribute.
	UefiKeywordName string
	// The UEFI namespace ID for the attribute.
	UefiNamespaceID string `json:"UefiNamespaceId"`
	// The upper limit for an integer attribute.
	UpperBound big.Int
	// 	An array of the possible values for enumerated attribute values.
	Value []AttributeValue
	// A valid regular expression, according to the Perl regular expression dialect,
	// that validates the attribute value.
	ValueExpression string
	// The warning text for the attribute.
	WarningText string
	// An indication of whether this attribute is write-only.
	WriteOnly bool
}

type DependencyType string

const (
	// A simple mapping dependency. If the condition evaluates to true,
	// the attribute or state changes to the mapped value.
	MapDependencyType DependencyType = "Map"
)

type MapToProperty string

const (
	// The dependency that affects an attribute's CurrentValue.
	CurrentValueMapToProperty MapToProperty = "CurrentValue"
	// The dependency that affects an attribute's DefaultValue.
	DefaultValueMapToProperty MapToProperty = "DefaultValue"
	// The dependency that affects an attribute's DisplayName.
	DisplayNameMapToProperty MapToProperty = "DisplayName"
	// The dependency that affects an attribute's DisplayName.
	DisplayOrderMapToProperty MapToProperty = "DisplayOrder"
	// The dependency that affects an attribute's GrayOut state.
	GrayOutMapToProperty MapToProperty = "GrayOut"
	// The dependency that affects an attribute's HelpText.
	HelpTextMapToProperty MapToProperty = "HelpText"
	// The dependency that affects an attribute's Hidden state.
	HiddenMapToProperty MapToProperty = "Hidden"
	// The dependency that affects an attribute's Immutable state.
	ImmutableMapToProperty MapToProperty = "Immutable"
	// The dependency that affects an attribute's LowerBound.
	LowerBoundMapToProperty MapToProperty = "LowerBound"
	// The dependency that affects an attribute's MaxLength.
	MaxLengthMapToProperty MapToProperty = "MaxLength"
	// The dependency that affects an attribute's MinLength.
	MinLengthMapToProperty MapToProperty = "MinLength"
	// The dependency that affects an attribute's ReadOnly state.
	ReadOnlyMapToProperty MapToProperty = "ReadOnly"
	// The dependency that affects an attribute's ScalarIncrement.
	ScalarIncrementMapToProperty MapToProperty = "ScalarIncrement"
	// The dependency that affects an attribute's UpperBound.
	UpperBoundMapToProperty MapToProperty = "UpperBound"
	// The dependency that affects an attribute's ValueExpression.
	ValueExpressionMapToProperty MapToProperty = "ValueExpression"
	// The dependency that affects an attribute's WarningText.
	WarningTextMapToProperty MapToProperty = "WarningText"
	// The dependency that affects an attribute's WriteOnly state.
	WriteOnlyMapToProperty MapToProperty = "WriteOnly"
)

type MapFromProperty string

const (
	// The dependency on an attribute's CurrentValue
	CurrentValueMapFromProperty MapFromProperty = "CurrentValue"
	// The dependency on an attribute's DefaultValue.
	DefaultValueMapFromProperty MapFromProperty = "DefaultValue"
	// The dependency on an attribute's GrayOut state.
	GrayOutMapFromProperty MapFromProperty = "GrayOut"
	// The dependency on an attribute's Hidden state.
	HiddenMapFromProperty MapFromProperty = "Hidden"
	// The dependency on an attribute's LowerBound.
	LowerBoundMapFromProperty MapFromProperty = "LowerBound"
	// The dependency on an attribute's MaxLength.
	MaxLengthMapFromProperty MapFromProperty = "MaxLength"
	// The dependency on an attribute's MinLength.
	MinLengthMapFromProperty MapFromProperty = "MinLength"
	// The dependency on an attribute's ReadOnly state.
	ReadOnlyMapFromProperty MapFromProperty = "ReadOnly"
	// The dependency on an attribute's ScalarIncrement.
	ScalarIncrementMapFromProperty MapFromProperty = "ScalarIncrement"
	// The dependency on an attribute's UpperBound.
	UpperBoundMapFromProperty MapFromProperty = "UpperBound"
	// The dependency on an attribute's WriteOnly state.
	WriteOnlyMapFromProperty MapFromProperty = "WriteOnly"
)

type MapFromCondition string

const (
	// The logical operation for 'Equal'.
	EqualCondition MapFromCondition = "EQU"
	// The logical operation for 'Greater than or Equal'.
	GreaterThanOrEqualCondition MapFromCondition = "GEQ"
	// The logical operation for 'Greater than'.
	GreaterThanCondition MapFromCondition = "GTR"
	// The logical operation for 'Less than or Equal'.
	LessThanOrEqualCondition MapFromCondition = "LEQ"
	// The logical operation for 'Less than'.
	LessThanCondition MapFromCondition = "LSS"
	// The logical operation for 'Not Equal'.
	NotEqualCondition MapFromCondition = "NEQ"
)

type MapTerms string

const (
	// The operation used for logical 'AND' of dependency terms.
	AndLogicalTerm MapTerms = "AND"
	// The operation used for logical 'OR' of dependency terms.
	OrLogicalTerm MapTerms = "OR"
)

type MapFrom struct {
	// The attribute to use to evaluate this dependency expression.
	MapFromAttribute string
	// The condition to use to evaluate this dependency expression.
	MapFromCondition MapFromCondition
	// The metadata property for the attribute that the MapFromAttribute property specifies
	// to use to evaluate this dependency expression.
	MapFromProperty MapFromProperty
	// The value to use to evaluate this dependency expression.
	MapFromValue interface{}
	// The logical term that combines two or more map-from conditions
	// in this dependency expression.
	MapTerms MapTerms
}

// The dependency expression for one or more attributes in this attribute registry.
type DependencyExpression struct {
	// An array of the map-from conditions for a mapping dependency.
	MapFrom []MapFrom
	// The AttributeName of the attribute that is affected by this dependency expression.
	MapToAttribute string
	// The metadata property for the attribute that contains
	// the map-from condition that evaluates this dependency expression.
	MapToProperty MapToProperty
	// The value that the map-to property changes to
	// if the dependency expression evaluates to true.
	MapToValue interface{}
}

// Dependency represents dependencies of attributes on this component.
type Dependency struct {
	// The dependency expression for one or more attributes in this attribute registry.
	Dependency DependencyExpression
	// The AttributeName of the attribute whose change triggers
	// the evaluation of this dependency expression.
	DependencyFor string
	// The type of the dependency structure.
	Type DependencyType
}

// Menu represents the attributes menus and their hierarchy in the attribute registry.
type Menu struct {
	// The user-readable display string of this menu in the defined language.
	DisplayName string
	// The ascending order, as a number, in which this menu appears relative to other menus.
	DisplayOrder int64
	// An indication of whether this menu is grayed out.
	GrayOut bool
	// An indication of whether this menu is hidden in user interfaces.
	Hidden bool
	// The unique name string of this menu.
	MenuName string
	// The path to the menu names that describes this menu
	// hierarchy relative to other menus.
	MenuPath string
	// Oem contains all the vendor specific information.
	Oem json.RawMessage
	// An indication of whether this menu is read-only.
	ReadOnly bool
}

// RegistryEntries shall list attributes for this component,
// along with their possible values, dependencies, and other metadata.
type RegistryEntries struct {
	// An array of attributes and their possible values in the attribute registry.
	Attributes []Attribute
	// An array of dependencies of attributes on this component.
	Dependencies []Dependency
	// An array for the attributes menus and their hierarchy in the attribute registry.
	Menus []Menu
}

type SupportedSystem struct {
	// The version of the component firmware image to which this attribute registry applies.
	FirmwareVersion string
	// The product name of the computer system to which this attribute registry applies.
	ProductName string
	// The ID of the systems to which this attribute registry applies.
	SystemID string `json:"SystemId"`
}

// AttributeRegistry shall represent an attribute registry for a Redfish implementation.
type AttributeRegistry struct {
	common.Entity

	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string

	// The RFC5646-conformant language code for the attribute registry.
	Language string
	// The organization or company that publishes this attribute registry.
	OwningEntity string
	// The list of all attributes and their metadata for this component.
	RegistryEntries RegistryEntries
	// The attribute registry version.
	RegistryVersion string
	// An array of systems that this attribute registry supports.
	SupportedSystems []SupportedSystem
}

// GetAttributeRegistry will get an AttributeRegistry instance from the Redfish service,
// e.g. BiosAttributeRegistry
func GetAttributeRegistry(c common.Client, uri string) (*AttributeRegistry, error) {
	return common.GetObject[AttributeRegistry](c, uri)
}
