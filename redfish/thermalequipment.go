//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2025.2 - #ThermalEquipment.v1_2_0.ThermalEquipment

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// ThermalEquipment shall represent the set of cooling equipment for a Redfish
// implementation.
type ThermalEquipment struct {
	common.Entity
	// CDUs shall contain a link to a resource collection of type
	// 'CoolingUnitCollection' that contains a set of coolant distribution units.
	cDUs string
	// CoolingLoopRedundancy shall contain redundancy information for the set of
	// cooling loops attached to this equipment. The values of the
	// 'RedundancyGroup' array shall reference resources of type 'CoolingLoop'.
	//
	// Version added: v1.1.0
	CoolingLoopRedundancy []common.RedundantGroup
	// CoolingLoops shall contain a link to a resource collection of type
	// 'CoolingLoopCollection' that contains the set of cooling loops managed by
	// the service.
	coolingLoops string
	// HeatExchangers shall contain a link to a resource collection of type
	// 'CoolingUnitCollection' that contains a set of heat exchanger units.
	heatExchangers string
	// ImmersionUnits shall contain a link to a resource collection of type
	// 'CoolingUnitCollection' that contains a set of immersion cooling units.
	immersionUnits string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RPUs shall contain a link to a resource collection of type
	// 'CoolingUnitCollection' that contains a set of reservoir and pumping units.
	//
	// Version added: v1.2.0
	rPUs string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a ThermalEquipment object from the raw JSON.
func (t *ThermalEquipment) UnmarshalJSON(b []byte) error {
	type temp ThermalEquipment
	var tmp struct {
		temp
		CDUs           common.Link `json:"cDUs"`
		CoolingLoops   common.Link `json:"coolingLoops"`
		HeatExchangers common.Link `json:"heatExchangers"`
		ImmersionUnits common.Link `json:"immersionUnits"`
		RPUs           common.Link `json:"rPUs"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*t = ThermalEquipment(tmp.temp)

	// Extract the links to other entities for later
	t.cDUs = tmp.CDUs.String()
	t.coolingLoops = tmp.CoolingLoops.String()
	t.heatExchangers = tmp.HeatExchangers.String()
	t.immersionUnits = tmp.ImmersionUnits.String()
	t.rPUs = tmp.RPUs.String()

	// This is a read/write object, so we need to save the raw object data for later
	t.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (t *ThermalEquipment) Update() error {
	readWriteFields := []string{
		"CoolingLoopRedundancy",
		"Status",
	}

	return t.UpdateFromRawData(t, t.rawData, readWriteFields)
}

// GetThermalEquipment will get a ThermalEquipment instance from the service.
func GetThermalEquipment(c common.Client, uri string) (*ThermalEquipment, error) {
	return common.GetObject[ThermalEquipment](c, uri)
}

// ListReferencedThermalEquipments gets the collection of ThermalEquipment from
// a provided reference.
func ListReferencedThermalEquipments(c common.Client, link string) ([]*ThermalEquipment, error) {
	return common.GetCollectionObjects[ThermalEquipment](c, link)
}

// CDUs gets the CDUs collection.
func (t *ThermalEquipment) CDUs(client common.Client) ([]*CoolingUnit, error) {
	if t.cDUs == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[CoolingUnit](client, t.cDUs)
}

// CoolingLoops gets the CoolingLoops collection.
func (t *ThermalEquipment) CoolingLoops(client common.Client) ([]*CoolingLoop, error) {
	if t.coolingLoops == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[CoolingLoop](client, t.coolingLoops)
}

// HeatExchangers gets the HeatExchangers collection.
func (t *ThermalEquipment) HeatExchangers(client common.Client) ([]*CoolingUnit, error) {
	if t.heatExchangers == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[CoolingUnit](client, t.heatExchangers)
}

// ImmersionUnits gets the ImmersionUnits collection.
func (t *ThermalEquipment) ImmersionUnits(client common.Client) ([]*CoolingUnit, error) {
	if t.immersionUnits == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[CoolingUnit](client, t.immersionUnits)
}

// RPUs gets the RPUs collection.
func (t *ThermalEquipment) RPUs(client common.Client) ([]*CoolingUnit, error) {
	if t.rPUs == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[CoolingUnit](client, t.rPUs)
}
