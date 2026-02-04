//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/AutomationInstrumentation.v1_0_0.json
// 2025.2 - #AutomationInstrumentation.v1_0_0.AutomationInstrumentation

package schemas

import (
	"encoding/json"
)

// AutomationInstrumentation shall represent automation node instrumentation for
// a Redfish implementation.
type AutomationInstrumentation struct {
	Entity
	// CurrentAmps shall contain the current, in ampere units, for this automation
	// node. The value of the 'DataSourceUri' property, if present, shall reference
	// a resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Current'.
	CurrentAmps SensorCurrentExcerpt
	// NodeControl shall contain the control for this resource. The value of the
	// 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Control'.
	NodeControl ControlNodeExcerpt
	// NodeState shall specify the current state of the automation node.
	NodeState NodeState
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PID shall contain a PID-based control loop for this resource. The value of
	// the 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Control'.
	PID ControlSingleLoopExcerpt
	// Status shall contain any status or health properties of a resource.
	Status Status
	// TemperatureCelsius shall contain the temperature, in degree Celsius units,
	// for this resource. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Temperature'.
	TemperatureCelsius SensorExcerpt
	// Voltage shall contain the voltage, in volt units, for this automation node.
	// The value of the 'DataSourceUri' property, if present, shall reference a
	// resource of type 'Sensor' with the 'ReadingType' property containing the
	// value 'Voltage'.
	Voltage SensorVoltageExcerpt
}

// GetAutomationInstrumentation will get a AutomationInstrumentation instance from the service.
func GetAutomationInstrumentation(c Client, uri string) (*AutomationInstrumentation, error) {
	return GetObject[AutomationInstrumentation](c, uri)
}

// ListReferencedAutomationInstrumentations gets the collection of AutomationInstrumentation from
// a provided reference.
func ListReferencedAutomationInstrumentations(c Client, link string) ([]*AutomationInstrumentation, error) {
	return GetCollectionObjects[AutomationInstrumentation](c, link)
}
