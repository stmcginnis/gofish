//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/CoolantConnector.v1_3_0.json
// 2025.3 - #CoolantConnector.v1_3_0.CoolantConnector

package schemas

import (
	"encoding/json"
)

type CoolantConnectorType string

const (
	// PairCoolantConnectorType is a connection pair.
	PairCoolantConnectorType CoolantConnectorType = "Pair"
	// SupplyCoolantConnectorType is a supply or intake connection.
	SupplyCoolantConnectorType CoolantConnectorType = "Supply"
	// ReturnCoolantConnectorType is a return or outflow connection.
	ReturnCoolantConnectorType CoolantConnectorType = "Return"
	// InlineCoolantConnectorType is an inline connection or measurement point.
	InlineCoolantConnectorType CoolantConnectorType = "Inline"
	// ClosedCoolantConnectorType is a closed or self-contained loop.
	ClosedCoolantConnectorType CoolantConnectorType = "Closed"
)

type ValveState string

const (
	// OpenValveState Open the valve for this connector to its normal operating
	// position.
	OpenValveState ValveState = "Open"
	// ClosedValveState Close the valve for this connector.
	ClosedValveState ValveState = "Closed"
)

type ValveStateReason string

const (
	// NormalValveStateReason shall indicate a normal operation of the valve, or a
	// return to a normal operating state. The 'State' property within 'Status'
	// shall reflect the chosen value of 'ValveState', where a value of 'Open'
	// shall indicate an 'Enabled' state, and a 'Closed' value shall indicate a
	// 'Disabled' state.
	NormalValveStateReason ValveStateReason = "Normal"
	// NotInUseValveStateReason shall indicate the valve is not in use or is not
	// connected, and therefore should remain closed. The 'State' property within
	// 'Status' shall indicate 'Absent'.
	NotInUseValveStateReason ValveStateReason = "NotInUse"
	// LeakDetectedValveStateReason shall indicate a leak was detected by an
	// external source. The 'Health' of the resource may be affected by this change
	// in state. The 'Health', 'State', and 'Condition' properties within 'Status'
	// should reflect the reaction taken by the service in response to a detected
	// leak.
	LeakDetectedValveStateReason ValveStateReason = "LeakDetected"
)

// CoolantConnector shall represent a coolant connector for a Redfish
// implementation.
type CoolantConnector struct {
	Entity
	// Coolant shall contain details regarding the coolant contained or used by
	// this unit.
	Coolant Coolant
	// CoolantConnectorType shall contain the type of coolant connector.
	CoolantConnectorType CoolantConnectorType
	// CoolingLoopName shall contain the name of the cooling loop attached to this
	// interface. If the 'CoolingLoop' link property is present, this property
	// shall contain the value of the 'Id' property in the resource referenced by
	// that link.
	CoolingLoopName string
	// CoolingManagerURI shall contain a URI to the application or device that
	// provides administration or management of the cooling loop associated with
	// this interface.
	CoolingManagerURI string
	// DeltaPressureControlkPa shall contain the desired differential pressure, in
	// kilopascal units, of this coolant connector. The value of the
	// 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Control' with the 'ControlType' property containing the value
	// 'PressurekPa'. This property shall only be present for secondary coolant
	// connectors. Services may automatically change other controls if a client
	// attempts to enable this control to prevent conflicts.
	//
	// Version added: v1.1.0
	DeltaPressureControlkPa ControlSingleLoopExcerpt
	// DeltaPressurekPa shall contain the pressure, in kilopascal units, for the
	// difference in pressure between the supply and outflow or return connection
	// to the cooling loop. The value of the 'DataSourceUri' property, if present,
	// shall reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'PressurekPa'.
	DeltaPressurekPa SensorExcerpt
	// DeltaTemperatureCelsius shall contain the change in temperature, in degree
	// Celsius units, between the supply connection and the outflow or return
	// connection to the cooling loop. The value of the 'DataSourceUri' property,
	// if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'Temperature'.
	DeltaTemperatureCelsius SensorExcerpt
	// DeltaTemperatureControlCelsius shall contain the desired differential
	// temperature, in degree Celsius units, of this coolant connector. The value
	// of the 'DataSourceUri' property, if present, shall reference a resource of
	// type 'Control' with the 'ControlType' property containing the value
	// 'Temperature'. This property shall only be present for secondary coolant
	// connectors. Services may automatically change other controls if a client
	// attempts to enable this control to prevent conflicts.
	//
	// Version added: v1.1.0
	DeltaTemperatureControlCelsius ControlSingleLoopExcerpt
	// FlowControlLitersPerMinute shall contain the desired liquid flow rate, in
	// liters per minute units, of this coolant connector. The value of the
	// 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Control' with the 'ControlType' property containing the value
	// 'LiquidFlowLPM'. This property shall only be present for secondary coolant
	// connectors. Services may automatically change other controls if a client
	// attempts to enable this control to prevent conflicts.
	//
	// Version added: v1.1.0
	FlowControlLitersPerMinute ControlSingleLoopExcerpt
	// FlowLitersPerMinute shall contain the liquid flow rate, in liters per minute
	// units, for this coolant connector. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'LiquidFlowLPM'.
	FlowLitersPerMinute SensorExcerpt
	// HeatRemovedkW shall contain the amount of heat removed, in kilowatt units,
	// by the coolant flow through this connector. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Sensor' with the
	// 'ReadingType' property containing the value 'Heat'.
	HeatRemovedkW SensorExcerpt
	// LocationIndicatorActive shall contain the state of the indicator used to
	// physically identify or locate this resource.
	LocationIndicatorActive bool
	// Manufacturer shall contain the name of the organization responsible for
	// producing the component. This organization may be the entity from whom the
	// component is purchased, but this is not necessarily true. This property is
	// generally used only for replaceable or user-configurable components.
	//
	// Version added: v1.3.0
	Manufacturer string
	// Model shall contain the name by which the manufacturer generally refers to
	// the component. This property is generally used only for replaceable or
	// user-configurable components.
	//
	// Version added: v1.3.0
	Model string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain a part number assigned by the organization that is
	// responsible for producing or manufacturing the component. This property is
	// generally used only for replaceable or user-configurable components.
	//
	// Version added: v1.3.0
	PartNumber string
	// RatedFlowLitersPerMinute shall contain the rated liquid flow, in liters per
	// minute units, for this loop interface.
	RatedFlowLitersPerMinute *float64 `json:",omitempty"`
	// RatedFlowPressurekPa shall contain the pressure, in kilopascal units, that
	// the rated liquid flow is valid for this connector.
	RatedFlowPressurekPa *float64 `json:",omitempty"`
	// RatedPressurekPa shall contain the rated maximum pressure, in kilopascal
	// units, for this connector.
	RatedPressurekPa *float64 `json:",omitempty"`
	// ReturnPressurekPa shall contain the pressure, in kilopascal units, for the
	// outflow or return connection to the cooling loop. The value of the
	// 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value 'PressurekPa'.
	ReturnPressurekPa SensorExcerpt
	// ReturnTemperatureCelsius shall contain the temperature, in degree Celsius
	// units, for the outflow or return connection to the cooling loop. The value
	// of the 'DataSourceUri' property, if present, shall reference a resource of
	// type 'Sensor' with the 'ReadingType' property containing the value
	// 'Temperature'.
	ReturnTemperatureCelsius SensorExcerpt
	// ReturnTemperatureControlCelsius shall contain the desired return
	// temperature, in degree Celsius units, of this coolant connector. The value
	// of the 'DataSourceUri' property, if present, shall reference a resource of
	// type 'Control' with the 'ControlType' property containing the value
	// 'Temperature'. This property shall only be present for secondary coolant
	// connectors. Services may automatically change other controls if a client
	// attempts to enable this control to prevent conflicts.
	//
	// Version added: v1.1.0
	ReturnTemperatureControlCelsius ControlSingleLoopExcerpt
	// SKU shall contain the stock-keeping unit number for this component. This
	// property is generally used only for replaceable or user-configurable
	// components.
	//
	// Version added: v1.3.0
	SKU string
	// SerialNumber shall contain a manufacturer-allocated number that identifies
	// the component. This property is generally used only for replaceable or
	// user-configurable components.
	//
	// Version added: v1.3.0
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the component. This
	// property is generally used only for replaceable or user-configurable
	// components.
	//
	// Version added: v1.3.0
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status Status
	// SupplyPressurekPa shall contain the pressure, in kilopascal units, for the
	// intake or supply connection to the cooling loop. The value of the
	// 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value 'PressurekPa'.
	SupplyPressurekPa SensorExcerpt
	// SupplyTemperatureCelsius shall contain the temperature, in degree Celsius
	// units, for the intake or supply connection to the cooling loop. The value of
	// the 'DataSourceUri' property, if present, shall reference a resource of type
	// 'Sensor' with the 'ReadingType' property containing the value 'Temperature'.
	SupplyTemperatureCelsius SensorExcerpt
	// SupplyTemperatureControlCelsius shall contain the desired supply
	// temperature, in degree Celsius units, of this coolant connector. The value
	// of the 'DataSourceUri' property, if present, shall reference a resource of
	// type 'Control' with the 'ControlType' property containing the value
	// 'Temperature'. This property shall only be present for secondary coolant
	// connectors. Services may automatically change other controls if a client
	// attempts to enable this control to prevent conflicts.
	//
	// Version added: v1.1.0
	SupplyTemperatureControlCelsius ControlSingleLoopExcerpt
	// ValvePositionControlPercent shall contain the desired valve position, in
	// percent units, of this coolant connector. The value of the 'DataSourceUri'
	// property, if present, shall reference a resource of type 'Control' with the
	// 'ControlType' property containing the value 'Valve'. Services may
	// automatically change other controls if a client attempts to enable this
	// control to prevent conflicts.
	//
	// Version added: v1.2.0
	ValvePositionControlPercent ControlSingleLoopExcerpt
	// ValvePositionPercent shall contain the valve position, in percent units, of
	// this connector. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Valve'.
	//
	// Version added: v1.2.0
	ValvePositionPercent SensorExcerpt
	// valveControlTarget is the URL to send ValveControl requests.
	valveControlTarget string
	// connectedChassis are the URIs for ConnectedChassis.
	connectedChassis []string
	// connectedCoolingLoop is the URI for ConnectedCoolingLoop.
	connectedCoolingLoop string
	// connectedCoolingUnit is the URI for ConnectedCoolingUnit.
	connectedCoolingUnit string
	// RawData holds the original serialized JSON so we can compare updates.
	RawData []byte
}

// UnmarshalJSON unmarshals a CoolantConnector object from the raw JSON.
func (c *CoolantConnector) UnmarshalJSON(b []byte) error {
	type temp CoolantConnector
	type cActions struct {
		ValveControl ActionTarget `json:"#CoolantConnector.ValveControl"`
	}
	type cLinks struct {
		ConnectedChassis     Links `json:"ConnectedChassis"`
		ConnectedCoolingLoop Link  `json:"ConnectedCoolingLoop"`
		ConnectedCoolingUnit Link  `json:"ConnectedCoolingUnit"`
	}
	var tmp struct {
		temp
		Actions cActions
		Links   cLinks
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*c = CoolantConnector(tmp.temp)

	// Extract the links to other entities for later
	c.valveControlTarget = tmp.Actions.ValveControl.Target
	c.connectedChassis = tmp.Links.ConnectedChassis.ToStrings()
	c.connectedCoolingLoop = tmp.Links.ConnectedCoolingLoop.String()
	c.connectedCoolingUnit = tmp.Links.ConnectedCoolingUnit.String()

	// This is a read/write object, so we need to save the raw object data for later
	c.RawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (c *CoolantConnector) Update() error {
	readWriteFields := []string{
		"CoolingLoopName",
		"CoolingManagerURI",
		"LocationIndicatorActive",
	}

	return c.UpdateFromRawData(c, c.RawData, readWriteFields)
}

// GetCoolantConnector will get a CoolantConnector instance from the service.
func GetCoolantConnector(c Client, uri string) (*CoolantConnector, error) {
	return GetObject[CoolantConnector](c, uri)
}

// ListReferencedCoolantConnectors gets the collection of CoolantConnector from
// a provided reference.
func ListReferencedCoolantConnectors(c Client, link string) ([]*CoolantConnector, error) {
	return GetCollectionObjects[CoolantConnector](c, link)
}

// This action shall set the operating state of the coolant connector
// represented by the resource.
// valveState - This parameter shall contain the desired valve state for the
// coolant connector. If this parameter is not provided, the service shall not
// change the valve state for this connector. Upon successful completion, the
// value of the 'State' property within 'Status' shall reflect this value.
// valveStateReason - This parameter shall contain the reason for desired state
// for the coolant connector. Upon successful completion, the value of the
// 'Health' property within 'Status' shall reflect the result of the
// 'ValveState' if the reason is applicable. If this parameter is not provided,
// the value shall be 'Normal'.
//
// If TaskMonitorInfo is not nil it can be used to monitor async tasks.
func (c *CoolantConnector) ValveControl(valveState ValveState, valveStateReason ValveStateReason) (*TaskMonitorInfo, error) {
	payload := make(map[string]any)
	payload["ValveState"] = valveState
	payload["ValveStateReason"] = valveStateReason
	resp, taskInfo, err := PostWithTask(c.client,
		c.valveControlTarget, payload, c.Headers(), false)
	defer DeferredCleanupHTTPResponse(resp)
	return taskInfo, err
}

// ConnectedChassis gets the ConnectedChassis linked resources.
func (c *CoolantConnector) ConnectedChassis() ([]*Chassis, error) {
	return GetObjects[Chassis](c.client, c.connectedChassis)
}

// ConnectedCoolingLoop gets the ConnectedCoolingLoop linked resource.
func (c *CoolantConnector) ConnectedCoolingLoop() (*CoolingLoop, error) {
	if c.connectedCoolingLoop == "" {
		return nil, nil
	}
	return GetObject[CoolingLoop](c.client, c.connectedCoolingLoop)
}

// ConnectedCoolingUnit gets the ConnectedCoolingUnit linked resource.
func (c *CoolantConnector) ConnectedCoolingUnit() (*CoolingUnit, error) {
	if c.connectedCoolingUnit == "" {
		return nil, nil
	}
	return GetObject[CoolingUnit](c.client, c.connectedCoolingUnit)
}
