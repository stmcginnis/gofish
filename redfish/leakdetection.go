//
// SPDX-License-Identifier: BSD-3-Clause
//
// 2024.3 - #LeakDetection.v1_1_0.LeakDetection

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// LeakDetection shall represent the leak detection functionality present in a
// service for a Redfish implementation.
type LeakDetection struct {
	common.Entity
	// LeakDetectorGroups shall contain an array of leak detection groups.
	LeakDetectorGroups []LeakDetectorGroup
	// LeakDetectors shall contain a link to a resource collection of type
	// 'LeakDetectorCollection'.
	leakDetectors string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Oem shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
	// rawData holds the original serialized JSON so we can compare updates.
	rawData []byte
}

// UnmarshalJSON unmarshals a LeakDetection object from the raw JSON.
func (l *LeakDetection) UnmarshalJSON(b []byte) error {
	type temp LeakDetection
	var tmp struct {
		temp
		LeakDetectors common.Link `json:"leakDetectors"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*l = LeakDetection(tmp.temp)

	// Extract the links to other entities for later
	l.leakDetectors = tmp.LeakDetectors.String()

	// This is a read/write object, so we need to save the raw object data for later
	l.rawData = b

	return nil
}

// Update commits updates to this object's properties to the running system.
func (l *LeakDetection) Update() error {
	readWriteFields := []string{
		"LeakDetectorGroups",
		"Status",
	}

	return l.UpdateFromRawData(l, l.rawData, readWriteFields)
}

// GetLeakDetection will get a LeakDetection instance from the service.
func GetLeakDetection(c common.Client, uri string) (*LeakDetection, error) {
	return common.GetObject[LeakDetection](c, uri)
}

// ListReferencedLeakDetections gets the collection of LeakDetection from
// a provided reference.
func ListReferencedLeakDetections(c common.Client, link string) ([]*LeakDetection, error) {
	return common.GetCollectionObjects[LeakDetection](c, link)
}

// LeakDetectors gets the LeakDetectors collection.
func (l *LeakDetection) LeakDetectors(client common.Client) ([]*LeakDetector, error) {
	if l.leakDetectors == "" {
		return nil, nil
	}
	return common.GetCollectionObjects[LeakDetector](client, l.leakDetectors)
}

// LeakDetectorGroup shall contain a group of leak detection equipment that
// reports a unified status.
type LeakDetectorGroup struct {
	// Detectors shall contain the states of all leak detection devices in this
	// detector group. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'LeakDetector'.
	Detectors []LeakDetectorArrayExcerpt
	// Detectors@odata.count
	DetectorsCount int `json:"Detectors@odata.count"`
	// GroupName shall contain the name used to describe this group of leak
	// detectors and related equipment.
	GroupName string
	// HumidityPercent shall contain the humidity, in percent units, for this
	// resource. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'Sensor' with the 'ReadingType' property
	// containing the value 'Humidity'.
	HumidityPercent SensorExcerpt
	// Status shall contain any status or health properties of the resource.
	//
	// Version added: v1.1.0
	Status common.Status
}
