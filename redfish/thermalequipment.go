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
	cdus []string
	// CoolingLoopRedundancy shall contain redundancy information for the set of cooling loops attached to this
	// equipment. The values of the RedundancyGroup array shall reference resources of type CoolingLoop.
	CoolingLoopRedundancy []RedundantGroup
	// CoolingLoops shall contain a link to a resource collection of type CoolingLoopCollection that contains the set
	// of cooling loops managed by the service.
	coolingLoops []string
	// Description provides a description of this resource.
	Description string
	// HeatExchangers shall contain a link to a resource collection of type CoolingUnitCollection that contains a set
	// of heat exchanger units.
	heatExchangers []string
	// ImmersionUnits shall contain a link to a resource collection of type CoolingUnitCollection that contains a set
	// of immersion cooling units.
	immersionUnits []string
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
		CDUs           common.LinksCollection
		CoolingLoops   common.LinksCollection
		HeatExchangers common.LinksCollection
		ImmersionUnits common.LinksCollection
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thermalequipment = ThermalEquipment(t.temp)

	// Extract the links to other entities for later
	thermalequipment.cdus = t.CDUs.ToStrings()
	thermalequipment.coolingLoops = t.CoolingLoops.ToStrings()
	thermalequipment.heatExchangers = t.HeatExchangers.ToStrings()
	thermalequipment.immersionUnits = t.ImmersionUnits.ToStrings()

	return nil
}

// CDUs gets a collection of coolant distribution units.
func (thermalequipment *ThermalEquipment) CDUs() ([]*CoolingUnit, error) {
	var result []*CoolingUnit

	collectionError := common.NewCollectionError()
	for _, uri := range thermalequipment.cdus {
		item, err := GetCoolingUnit(thermalequipment.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// CoolingLoops gets a collection of cooling loops.
func (thermalequipment *ThermalEquipment) CoolingLoops() ([]*CoolingLoop, error) {
	var result []*CoolingLoop

	collectionError := common.NewCollectionError()
	for _, uri := range thermalequipment.coolingLoops {
		item, err := GetCoolingLoop(thermalequipment.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// HeatExchangers gets a collection of heat exchangers.
func (thermalequipment *ThermalEquipment) HeatExchangers() ([]*CoolingUnit, error) {
	var result []*CoolingUnit

	collectionError := common.NewCollectionError()
	for _, uri := range thermalequipment.heatExchangers {
		item, err := GetCoolingUnit(thermalequipment.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// ImmersionUnits gets a collection of immersion cooling units.
func (thermalequipment *ThermalEquipment) ImmersionUnits() ([]*CoolingUnit, error) {
	var result []*CoolingUnit

	collectionError := common.NewCollectionError()
	for _, uri := range thermalequipment.immersionUnits {
		item, err := GetCoolingUnit(thermalequipment.GetClient(), uri)
		if err != nil {
			collectionError.Failures[uri] = err
		} else {
			result = append(result, item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}

// GetThermalEquipment will get a ThermalEquipment instance from the service.
func GetThermalEquipment(c common.Client, uri string) (*ThermalEquipment, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var thermalequipment ThermalEquipment
	err = json.NewDecoder(resp.Body).Decode(&thermalequipment)
	if err != nil {
		return nil, err
	}

	thermalequipment.SetClient(c)
	return &thermalequipment, nil
}

// ListReferencedThermalEquipments gets the collection of ThermalEquipment from
// a provided reference.
func ListReferencedThermalEquipments(c common.Client, link string) ([]*ThermalEquipment, error) {
	var result []*ThermalEquipment
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *ThermalEquipment
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		thermalequipment, err := GetThermalEquipment(c, link)
		ch <- GetResult{Item: thermalequipment, Link: link, Error: err}
	}

	go func() {
		err := common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			result = append(result, r.Item)
		}
	}

	if collectionError.Empty() {
		return result, nil
	}

	return result, collectionError
}
