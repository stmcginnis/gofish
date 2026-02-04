//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/ThermalEquipment.v1_2_0.json
// 2025.2 - #ThermalEquipment.v1_2_0.ThermalEquipment

package schemas

import (
	"encoding/json"
)

// ThermalEquipment shall represent the set of cooling equipment for a Redfish
// implementation.
type ThermalEquipment struct {
	Entity
	// CDUs shall contain a link to a resource collection of type
	// 'CoolingUnitCollection' that contains a set of coolant distribution units.
	cDUs string
	// CoolingLoopRedundancy shall contain redundancy information for the set of
	// cooling loops attached to this equipment. The values of the
	// 'RedundancyGroup' array shall reference resources of type 'CoolingLoop'.
	//
	// Version added: v1.1.0
	CoolingLoopRedundancy []RedundantGroup
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
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// RPUs shall contain a link to a resource collection of type
	// 'CoolingUnitCollection' that contains a set of reservoir and pumping units.
	//
	// Version added: v1.2.0
	rPUs string
	// Status shall contain any status or health properties of the resource.
	Status Status
}

// UnmarshalJSON unmarshals a ThermalEquipment object from the raw JSON.
func (t *ThermalEquipment) UnmarshalJSON(b []byte) error {
	type temp ThermalEquipment
	var tmp struct {
		temp
		CDUs           Link `json:"CDUs"`
		CoolingLoops   Link `json:"CoolingLoops"`
		HeatExchangers Link `json:"HeatExchangers"`
		ImmersionUnits Link `json:"ImmersionUnits"`
		RPUs           Link `json:"RPUs"`
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

	return nil
}

// GetThermalEquipment will get a ThermalEquipment instance from the service.
func GetThermalEquipment(c Client, uri string) (*ThermalEquipment, error) {
	return GetObject[ThermalEquipment](c, uri)
}

// ListReferencedThermalEquipments gets the collection of ThermalEquipment from
// a provided reference.
func ListReferencedThermalEquipments(c Client, link string) ([]*ThermalEquipment, error) {
	return GetCollectionObjects[ThermalEquipment](c, link)
}

// CDUs gets the CDUs collection.
func (t *ThermalEquipment) CDUs() ([]*CoolingUnit, error) {
	if t.cDUs == "" {
		return nil, nil
	}
	return GetCollectionObjects[CoolingUnit](t.client, t.cDUs)
}

// CoolingLoops gets the CoolingLoops collection.
func (t *ThermalEquipment) CoolingLoops() ([]*CoolingLoop, error) {
	if t.coolingLoops == "" {
		return nil, nil
	}
	return GetCollectionObjects[CoolingLoop](t.client, t.coolingLoops)
}

// HeatExchangers gets the HeatExchangers collection.
func (t *ThermalEquipment) HeatExchangers() ([]*CoolingUnit, error) {
	if t.heatExchangers == "" {
		return nil, nil
	}
	return GetCollectionObjects[CoolingUnit](t.client, t.heatExchangers)
}

// ImmersionUnits gets the ImmersionUnits collection.
func (t *ThermalEquipment) ImmersionUnits() ([]*CoolingUnit, error) {
	if t.immersionUnits == "" {
		return nil, nil
	}
	return GetCollectionObjects[CoolingUnit](t.client, t.immersionUnits)
}

// RPUs gets the RPUs collection.
func (t *ThermalEquipment) RPUs() ([]*CoolingUnit, error) {
	if t.rPUs == "" {
		return nil, nil
	}
	return GetCollectionObjects[CoolingUnit](t.client, t.rPUs)
}
