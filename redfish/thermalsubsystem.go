//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// ThermalSubsystem shall represent a thermal subsystem for a Redfish implementation.
type ThermalSubsystem struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// CoolantConnectorRedundancy shall contain redundancy information for the set of coolant connectors attached to
	// this equipment. The values of the RedundancyGroup array shall reference resources of type CoolantConnector.
	CoolantConnectorRedundancy []RedundantGroup
	// CoolantConnectors shall contain a link to a resource collection of type CoolantConnectorCollection that contains
	// the coolant connectors for this equipment.
	coolantConnectors []string
	// Description provides a description of this resource.
	Description string
	// FanRedundancy shall contain redundancy information for the groups of fans in this subsystem.
	FanRedundancy []RedundantGroup
	// Fans shall contain a link to a resource collection of type FanCollection.
	fans []string
	// Heaters shall contain a link to a resource collection of type HeaterCollection.
	heaters []string
	// LeakDetection shall contain a link to a resource collection of type LeakDetection.
	leakDetection []string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Pumps shall contain a link to a resource collection of type PumpCollection that contains details for the pumps
	// included in this equipment.
	pumps []string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// ThermalMetrics shall contain a link to a resource collection of type ThermalMetrics.
	thermalMetrics []string
}

// UnmarshalJSON unmarshals a ThermalSubsystem object from the raw JSON.
func (thermalsubsystem *ThermalSubsystem) UnmarshalJSON(b []byte) error {
	type temp ThermalSubsystem
	var t struct {
		temp
		CoolantConnectors common.LinksCollection
		Fans              common.LinksCollection
		Heaters           common.LinksCollection
		LeakDetection     common.LinksCollection
		Pumps             common.LinksCollection
		ThermalMetrics    common.LinksCollection
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*thermalsubsystem = ThermalSubsystem(t.temp)

	// Extract the links to other entities for later
	thermalsubsystem.coolantConnectors = t.CoolantConnectors.ToStrings()
	thermalsubsystem.fans = t.Fans.ToStrings()
	thermalsubsystem.heaters = t.Heaters.ToStrings()
	thermalsubsystem.leakDetection = t.LeakDetection.ToStrings()
	thermalsubsystem.pumps = t.Pumps.ToStrings()
	thermalsubsystem.thermalMetrics = t.ThermalMetrics.ToStrings()

	return nil
}

// CoolantConnectors gets the coolant connectors for this equipment.
func (thermalsubsystem *ThermalSubsystem) CoolantConnectors() ([]*CoolantConnector, error) {
	var result []*CoolantConnector

	collectionError := common.NewCollectionError()
	for _, uri := range thermalsubsystem.coolantConnectors {
		item, err := GetCoolantConnector(thermalsubsystem.GetClient(), uri)
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

// Fans gets the fans for this equipment.
func (thermalsubsystem *ThermalSubsystem) Fans() ([]*Fan, error) {
	var result []*Fan

	collectionError := common.NewCollectionError()
	for _, uri := range thermalsubsystem.fans {
		item, err := GetFan(thermalsubsystem.GetClient(), uri)
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

// Heaters gets the heaters within this subsystem.
func (thermalsubsystem *ThermalSubsystem) Heaters() ([]*Heater, error) {
	var result []*Heater

	collectionError := common.NewCollectionError()
	for _, uri := range thermalsubsystem.heaters {
		item, err := GetHeater(thermalsubsystem.GetClient(), uri)
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

// LeakDetection gets the leak detection system within this chassis.
func (thermalsubsystem *ThermalSubsystem) LeakDetection() ([]*LeakDetection, error) {
	var result []*LeakDetection

	collectionError := common.NewCollectionError()
	for _, uri := range thermalsubsystem.leakDetection {
		item, err := GetLeakDetection(thermalsubsystem.GetClient(), uri)
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

// Pumps gets the pumps for this equipment.
func (thermalsubsystem *ThermalSubsystem) Pumps() ([]*Pump, error) {
	var result []*Pump

	collectionError := common.NewCollectionError()
	for _, uri := range thermalsubsystem.pumps {
		item, err := GetPump(thermalsubsystem.GetClient(), uri)
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

// ThermalMetrics gets the summary of thermal metrics for this subsystem.
func (thermalsubsystem *ThermalSubsystem) ThermalMetrics() ([]*ThermalMetrics, error) {
	var result []*ThermalMetrics

	collectionError := common.NewCollectionError()
	for _, uri := range thermalsubsystem.thermalMetrics {
		item, err := GetThermalMetrics(thermalsubsystem.GetClient(), uri)
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

// GetThermalSubsystem will get a ThermalSubsystem instance from the service.
func GetThermalSubsystem(c common.Client, uri string) (*ThermalSubsystem, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var thermalsubsystem ThermalSubsystem
	err = json.NewDecoder(resp.Body).Decode(&thermalsubsystem)
	if err != nil {
		return nil, err
	}

	thermalsubsystem.SetClient(c)
	return &thermalsubsystem, nil
}

// ListReferencedThermalSubsystems gets the collection of ThermalSubsystem from
// a provided reference.
func ListReferencedThermalSubsystems(c common.Client, link string) ([]*ThermalSubsystem, error) {
	var result []*ThermalSubsystem
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *ThermalSubsystem
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		thermalsubsystem, err := GetThermalSubsystem(c, link)
		ch <- GetResult{Item: thermalsubsystem, Link: link, Error: err}
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
