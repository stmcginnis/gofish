//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// ThermalEquipment shall represent the set of cooling equipment for a Redfish implementation.
type ThermalEquipment struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// CDUs shall contain a link to a resource collection of type CoolingUnitCollection that contains a set of coolant
	// distribution units.
	cdus string
	// CoolingLoopRedundancy shall contain redundancy information for the set of cooling loops attached to this
	// equipment. The values of the RedundancyGroup array shall reference resources of type CoolingLoop.
	CoolingLoopRedundancy []RedundantGroup
	// CoolingLoops shall contain a link to a resource collection of type CoolingLoopCollection that contains the set
	// of cooling loops managed by the service.
	coolingLoops string
	// Description provides a description of this resource.
	Description string
	// HeatExchangers shall contain a link to a resource collection of type CoolingUnitCollection that contains a set
	// of heat exchanger units.
	heatExchangers string
	// ImmersionUnits shall contain a link to a resource collection of type CoolingUnitCollection that contains a set
	// of immersion cooling units.
	immersionUnits string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a ThermalEquipment object from the raw JSON.
func (thermalequipment *ThermalEquipment) UnmarshalJSON(b []byte) error {
	type temp ThermalEquipment
	var t struct {
		temp
		CDUs           common.Link
		CoolingLoops   common.Link
		HeatExchangers common.Link
		ImmersionUnits common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thermalequipment = ThermalEquipment(t.temp)

	// Extract the links to other entities for later
	thermalequipment.cdus = t.CDUs.String()
	thermalequipment.coolingLoops = t.CoolingLoops.String()
	thermalequipment.heatExchangers = t.HeatExchangers.String()
	thermalequipment.immersionUnits = t.ImmersionUnits.String()

	return nil
}

// CDUs gets a collection of coolant distribution units.
func (thermalequipment *ThermalEquipment) CDUs() ([]*CoolingUnit, error) {
	return ListReferencedCoolingUnits(thermalequipment.GetClient(), thermalequipment.cdus)
}

// CoolingLoops gets a collection of cooling loops.
func (thermalequipment *ThermalEquipment) CoolingLoops() ([]*CoolingLoop, error) {
	return ListReferencedCoolingLoops(thermalequipment.GetClient(), thermalequipment.coolingLoops)
}

// HeatExchangers gets a collection of heat exchangers.
func (thermalequipment *ThermalEquipment) HeatExchangers() ([]*CoolingUnit, error) {
	return ListReferencedCoolingUnits(thermalequipment.GetClient(), thermalequipment.heatExchangers)
}

// ImmersionUnits gets a collection of immersion cooling units.
func (thermalequipment *ThermalEquipment) ImmersionUnits() ([]*CoolingUnit, error) {
	return ListReferencedCoolingUnits(thermalequipment.GetClient(), thermalequipment.immersionUnits)
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
