//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"
	"reflect"

	"github.com/stmcginnis/gofish/common"
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
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
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
	// DeltaTemperatureCelsius shall contain the change in temperature, in degree Celsius units, between the supply
	// connection and the outflow or return connection to the cooling loop. The value of the DataSourceUri property, if
	// present, shall reference a resource of type Sensor with the ReadingType property containing the value
	// 'Temperature'.
	DeltaTemperatureCelsius SensorExcerpt
	// Description provides a description of this resource.
	Description string
	// FlowLitersPerMinute shall contain the liquid flow rate, in liters per minute units, for this coolant connector.
	// The value of the DataSourceUri property, if present, shall reference a resource of type Sensor with the
	// ReadingType property containing the value 'LiquidFlowLPM'.
	FlowLitersPerMinute SensorExcerpt
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
	// Get a representation of the object's original state so we can find what
	// to update.
	original := new(CoolantConnector)
	original.UnmarshalJSON(coolantconnector.rawData)

	readWriteFields := []string{
		"CoolingLoopName",
		"CoolingManagerURI",
		"LocationIndicatorActive",
	}

	originalElement := reflect.ValueOf(original).Elem()
	currentElement := reflect.ValueOf(coolantconnector).Elem()

	return coolantconnector.Entity.Update(originalElement, currentElement, readWriteFields)
}

// GetCoolantConnector will get a CoolantConnector instance from the service.
func GetCoolantConnector(c common.Client, uri string) (*CoolantConnector, error) {
	return common.GetObject[CoolantConnector](c, uri)
}

// ListReferencedCoolantConnectors gets the collection of CoolantConnector from
// a provided reference.
func ListReferencedCoolantConnectors(c common.Client, link string) ([]*CoolantConnector, error) {
	return common.GetCollectionObjects[CoolantConnector](c, link)
}

// ConnectedChassis retrieves a collection of the Chassis at the other end of the connection.
func (coolantconnector *CoolantConnector) ConnectedChassis() ([]*Chassis, error) {
	return common.GetObjects[Chassis](coolantconnector.GetClient(), coolantconnector.connectedChassis)
}

// ConnectedCoolingLoop gets the cooling loop at the other end of the connection.
func (coolantconnector *CoolantConnector) ConnectedCoolingLoop() (*CoolingLoop, error) {
	return GetCoolingLoop(coolantconnector.GetClient(), coolantconnector.connectedCoolingLoop)
}

// ConnectedCoolingUnit gets the cooling unit at the other end of the connection.
func (coolantconnector *CoolantConnector) ConnectedCoolingUnit() (*CoolingUnit, error) {
	return GetCoolingUnit(coolantconnector.GetClient(), coolantconnector.connectedCoolingUnit)
}
