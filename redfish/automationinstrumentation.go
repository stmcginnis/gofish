//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #AutomationInstrumentation.v1_0_0.AutomationInstrumentation

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// AutomationInstrumentation shall represent automation node instrumentation for
// a Redfish implementation.
type AutomationInstrumentation struct {
	common.Entity
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
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PID shall contain a PID-based control loop for this resource. The value of
	// the 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Control'.
	PID ControlSingleLoopExcerpt
	// Status shall contain any status or health properties of a resource.
	Status common.Status
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
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a AutomationInstrumentation object from the raw JSON.
func (a *AutomationInstrumentation) UnmarshalJSON(b []byte) error {
	type temp AutomationInstrumentation
	var tmp struct {
		temp
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*a = AutomationInstrumentation(tmp.temp)

	// Extract the links to other entities for later

	// This is a read/write object, so we need to save the raw object data for later
	a.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (a *AutomationInstrumentation) Update() error {
	readWriteFields := []string{
		"CurrentAmps",
		"NodeControl",
		"PID",
		"Status",
		"TemperatureCelsius",
		"Voltage",
	}

	return a.UpdateFromRawData(a, a.rawData, readWriteFields)
}

// GetAutomationInstrumentation will get a AutomationInstrumentation instance from the service.
func GetAutomationInstrumentation(c common.Client, uri string) (*AutomationInstrumentation, error) {
	return common.GetObject[AutomationInstrumentation](c, uri)
}

// ListReferencedAutomationInstrumentations gets the collection of AutomationInstrumentation from
// a provided reference.
func ListReferencedAutomationInstrumentations(c common.Client, link string) ([]*AutomationInstrumentation, error) {
	return common.GetCollectionObjects[AutomationInstrumentation](c, link)
}
