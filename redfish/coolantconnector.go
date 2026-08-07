//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"context"
	"encoding/json"

	"github.com/coreweave/gofish/common"
)

type CoolantConnectorType string

const (
	// PairCoolantConnectorType is a connection pair.
	PairCoolantConnectorType CoolantConnectorType = "Pair"
	// SupplyCoolantConnectorType is a supply or intake connection.
	SupplyCoolantConnectorType CoolantConnectorType = "Supply"
	// ReturnCoolantConnectorType is a return or outflow connection.
	ReturnCoolantConnectorType CoolantConnectorType = "Return"
	// InlineCoolantConnectorType An inline connection or measurement point.
	InlineCoolantConnectorType CoolantConnectorType = "Inline"
	// ClosedCoolantConnectorType is a closed or self-contained loop.
	ClosedCoolantConnectorType CoolantConnectorType = "Closed"
)

// CoolantConnector shall represent a coolant connector for a Redfish implementation.
type CoolantConnector struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Coolant shall contain details regarding the coolant contained or used by this unit.
	Coolant Coolant
	// CoolantConnectorType shall contain the type of coolant connector.
	CoolantConnectorType CoolantConnectorType
	// CoolingLoopName shall contain the name of the cooling loop attached to this interface. If the 'CoolingLoop' link
	// property is present, this property shall contain the value of the 'Id' property in the resource referenced by
	// that link.
	CoolingLoopName string
	// CoolingManagerURI shall contain a URI to the application or device that provides administration or management of
	// the cooling loop associated with this interface.
	CoolingManagerURI string
	// DeltaPressurekPa shall contain the pressure, in kilopascal units, for the difference in pressure between the
	// supply and outflow or return connection to the cooling loop. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value
	// 'PressurekPa'.
	DeltaPressurekPa SensorExcerpt
	// DeltaPressureControlkPa contain the control for the desired pressure difference, in kilopascal units, for this
	// coolant connector. This control shall only be present for the secondary coolant connector.
	DeltaPressureControlkPa ControlSingleExcerpt
	// DeltaTemperatureCelsius shall contain the change in temperature, in degree Celsius units, between the supply
	// connection and the outflow or return connection to the cooling loop. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value
	// 'Temperature'.
	DeltaTemperatureCelsius SensorExcerpt
	// DeltaTemperatureControlCelsius contain the control for the desired temperature difference, in degree Celsius
	// for this coolant connector. This control shall only be present for the secondary coolant connector.
	DeltaTemperatureControlCelsius ControlSingleExcerpt
	// Description provides a description of this resource.
	Description string
	// FlowLitersPerMinute shall contain the liquid flow rate, in liters per minute units, for this coolant connector.
	// The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'LiquidFlowLPM'.
	FlowLitersPerMinute SensorExcerpt
	// FlowControlLiterPerMinute shall contain the control for the liquid flow rate, in liters per minute units,
	// for this coolant connector. This control shall only be present for the secondary coolant connector.
	FlowControlLitersPerMinute ControlSingleExcerpt
	// HeatRemovedkW shall contain the amount of heat removed, in kilowatt units, by the coolant flow through this
	// connector. The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with
	// the ReadingType property containing the value 'Heat'.
	HeatRemovedkW SensorExcerpt
	// LocationIndicatorActive shall contain the state of the indicator used to physically identify or locate this
	// resource.
	LocationIndicatorActive bool
	// RatedFlowLitersPerMinute shall contain the rated liquid flow, in liters per minute units, for this loop
	// interface.
	RatedFlowLitersPerMinute float64
	// RatedFlowPressurekPa shall contain the pressure, in kilopascal units, that the rated liquid flow is valid for
	// this connector.
	RatedFlowPressurekPa float64
	// RatedPressurekPa shall contain the rated maximum pressure, in kilopascal units, for this connector.
	RatedPressurekPa float64
	// ReturnPressurekPa shall contain the pressure, in kilopascal units, for the outflow or return connection to the
	// cooling loop. The value of the DataSourceUri property, if present, shall reference a resource of type Sensor
	// with the ReadingType property containing the value 'PressurekPa'.
	ReturnPressurekPa SensorExcerpt
	// ReturnTemperatureCelsius shall contain the temperature, in degree Celsius units, for the outflow or return
	// connection to the cooling loop. The value of the DataSourceUri property, if present, shall reference a resource
	// of type Sensor with the ReadingType property containing the value 'Temperature'.
	ReturnTemperatureCelsius SensorExcerpt
	// ReturnTemperatureControlCelsius contain the control for the desired return temperature, in degree Celsius units
	// for this coolant connector.
	ReturnTemperatureControlCelsius ControlSingleExcerpt
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// SupplyPressurekPa shall contain the pressure, in kilopascal units, for the intake or supply connection to the
	// cooling loop. The value of the DataSourceUri property, if present, shall reference a resource of type Sensor
	// with the ReadingType property containing the value 'PressurekPa'.
	SupplyPressurekPa SensorExcerpt
	// SupplyTemperatureCelsius shall contain the temperature, in degree Celsius units, for the intake or supply
	// connection to the cooling loop. The value of the DataSourceUri property, if present, shall reference a resource
	// of type Sensor with the ReadingType property containing the value 'Temperature'.
	SupplyTemperatureCelsius SensorExcerpt
	// SupplyTemperatureControlCelsius contain the control for the desired supply temperature in degree Celsius units
	// of this coolant connector.
	SupplyTemperatureControlCelsius ControlSingleExcerpt
	// ValvePositionControlPercent contain the control for the desired valve position (% open) of this connector.
	ValvePositionControlPercent ControlSingleExcerpt
	// ValvePositionPercent the valve position (percent open) of this connector.
	ValvePositionPercent SensorExcerpt
	// rawData holds the original serialized JSON so we can compare updates.
	rawData          []byte
	connectedChassis []string
	// ConnectedChassisCount is the number of connected chassis at the other end of the connection.
	ConnectedChassisCount int
	connectedCoolingLoop  string
	connectedCoolingUnit  string
}

// UnmarshalJSON unmarshals a CoolantConnector object from the raw JSON.
func (coolantconnector *CoolantConnector) UnmarshalJSON(b []byte) error {
	type temp CoolantConnector
	type links struct {
		// ConnectedChassis shall contain an array of links to resources of type Chassis that represent the chassis at the
		// other end of the connection.
		ConnectedChassis common.Links
		// ConnectedChassis@odata.count
		ConnectedChassisCount int `json:"ConnectedChassis@odata.count"`
		// ConnectedCoolingLoop shall contain a link to a resource of type CoolingLoop that represents the cooling loop at
		// the other end of the connection.
		ConnectedCoolingLoop common.Link
		// ConnectedCoolingUnit shall contain a link to a resource of type CoolingUnit that represents the cooling unit at
		// the other end of the connection.
		ConnectedCoolingUnit common.Link
	}
	var t struct {
		temp
		Links links
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*coolantconnector = CoolantConnector(t.temp)

	// Extract the links to other entities for later
	coolantconnector.connectedChassis = t.Links.ConnectedChassis.ToStrings()
	coolantconnector.ConnectedChassisCount = t.Links.ConnectedChassisCount
	coolantconnector.connectedCoolingLoop = t.Links.ConnectedCoolingLoop.String()
	coolantconnector.connectedCoolingUnit = t.Links.ConnectedCoolingUnit.String()

	// This is a read/write object, so we need to save the raw object data for later
	coolantconnector.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (coolantconnector *CoolantConnector) Update() error {
	return coolantconnector.UpdateWithContext(common.ContextOf(coolantconnector.GetClient()))
}

// UpdateWithContext commits updates to this object's properties to the running system.
func (coolantconnector *CoolantConnector) UpdateWithContext(ctx context.Context) error {
	readWriteFields := []string{
		"CoolingLoopName",
		"CoolingManagerURI",
		"DeltaPressureControlkPa",
		"DeltaTemperatureControlCelsius",
		"FlowControlLitersPerMinute",
		"LocationIndicatorActive",
		"ReturnTemperatureControlCelsius",
		"SupplyTemperatureControlCelsius",
		"ValvePositionControlPercent",
	}

	return coolantconnector.UpdateFromRawDataWithContext(ctx, coolantconnector, coolantconnector.rawData, readWriteFields)
}

// GetCoolantConnector will get a CoolantConnector instance from the service.
func GetCoolantConnector(c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*CoolantConnector, error) {
	return GetCoolantConnectorWithContext(common.ContextOf(c), c, uri, queryOpts...)
}

// GetCoolantConnectorWithContext will get a CoolantConnector instance from the service.
func GetCoolantConnectorWithContext(ctx context.Context, c common.Client, uri string, queryOpts ...common.QueryGroupOption) (*CoolantConnector, error) {
	return common.GetObjectWithContext[CoolantConnector](ctx, c, uri, queryOpts...)
}

// ListReferencedCoolantConnectors gets the collection of CoolantConnector from
// a provided reference.
func ListReferencedCoolantConnectors(c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*CoolantConnector, error) {
	return ListReferencedCoolantConnectorsWithContext(common.ContextOf(c), c, link, queryOpts...)
}

// ListReferencedCoolantConnectorsWithContext gets the collection of CoolantConnector from
// a provided reference.
func ListReferencedCoolantConnectorsWithContext(ctx context.Context, c common.Client, link string, queryOpts ...common.QueryGroupOption) ([]*CoolantConnector, error) {
	return common.GetCollectionObjectsWithContext[CoolantConnector](ctx, c, link, queryOpts...)
}

// ConnectedChassis retrieves a collection of the Chassis at the other end of the connection.
func (coolantconnector *CoolantConnector) ConnectedChassis(queryOpts ...common.QueryGroupOption) ([]*Chassis, error) {
	return coolantconnector.ConnectedChassisWithContext(common.ContextOf(coolantconnector.GetClient()), queryOpts...)
}

// ConnectedChassisWithContext retrieves a collection of the Chassis at the other end of the connection.
func (coolantconnector *CoolantConnector) ConnectedChassisWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) ([]*Chassis, error) {
	return common.GetObjectsWithContext[Chassis](ctx, coolantconnector.GetClient(), coolantconnector.connectedChassis, queryOpts...)
}

// ConnectedCoolingLoop gets the cooling loop at the other end of the connection.
func (coolantconnector *CoolantConnector) ConnectedCoolingLoop(queryOpts ...common.QueryGroupOption) (*CoolingLoop, error) {
	return coolantconnector.ConnectedCoolingLoopWithContext(common.ContextOf(coolantconnector.GetClient()), queryOpts...)
}

// ConnectedCoolingLoopWithContext gets the cooling loop at the other end of the connection.
func (coolantconnector *CoolantConnector) ConnectedCoolingLoopWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*CoolingLoop, error) {
	return GetCoolingLoopWithContext(ctx, coolantconnector.GetClient(), coolantconnector.connectedCoolingLoop, queryOpts...)
}

// ConnectedCoolingUnit gets the cooling unit at the other end of the connection.
func (coolantconnector *CoolantConnector) ConnectedCoolingUnit(queryOpts ...common.QueryGroupOption) (*CoolingUnit, error) {
	return coolantconnector.ConnectedCoolingUnitWithContext(common.ContextOf(coolantconnector.GetClient()), queryOpts...)
}

// ConnectedCoolingUnitWithContext gets the cooling unit at the other end of the connection.
func (coolantconnector *CoolantConnector) ConnectedCoolingUnitWithContext(ctx context.Context, queryOpts ...common.QueryGroupOption) (*CoolingUnit, error) {
	return GetCoolingUnitWithContext(ctx, coolantconnector.GetClient(), coolantconnector.connectedCoolingUnit, queryOpts...)
}
