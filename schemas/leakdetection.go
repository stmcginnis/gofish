//
// SPDX-License-Identifier: BSD-3-Clause
//

// http://redfish.dmtf.org/schemas/v1/LeakDetection.v1_2_0.json
// 2025.4 - #LeakDetection.v1_2_0.LeakDetection

package schemas

import (
	"encoding/json"
)

// LeakDetection shall represent the leak detection functionality present in a
// service for a Redfish implementation.
type LeakDetection struct {
	Entity
	// LeakDetectorGroups shall contain an array of leak detection groups.
	LeakDetectorGroups []LeakDetectorGroup
	// LeakDetectors shall contain a link to a resource collection of type
	// 'LeakDetectorCollection'.
	//
	// Deprecated: v1.2.0
	// This property has been deprecated in favor of the 'LeakDetectors' property
	// in the related 'Chassis' resource.
	leakDetectors string
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// OEM shall contain the OEM extensions. All values for properties that this
	// object contains shall conform to the Redfish Specification-described
	// requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status Status
}

// UnmarshalJSON unmarshals a LeakDetection object from the raw JSON.
func (l *LeakDetection) UnmarshalJSON(b []byte) error {
	type temp LeakDetection
	var tmp struct {
		temp
		LeakDetectors Link `json:"LeakDetectors"`
	}

	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*l = LeakDetection(tmp.temp)

	// Extract the links to other entities for later
	l.leakDetectors = tmp.LeakDetectors.String()

	return nil
}

// GetLeakDetection will get a LeakDetection instance from the service.
func GetLeakDetection(c Client, uri string) (*LeakDetection, error) {
	return GetObject[LeakDetection](c, uri)
}

// ListReferencedLeakDetections gets the collection of LeakDetection from
// a provided reference.
func ListReferencedLeakDetections(c Client, link string) ([]*LeakDetection, error) {
	return GetCollectionObjects[LeakDetection](c, link)
}

// LeakDetectors gets the LeakDetectors collection.
func (l *LeakDetection) LeakDetectors() ([]*LeakDetector, error) {
	if l.leakDetectors == "" {
		return nil, nil
	}
	return GetCollectionObjects[LeakDetector](l.client, l.leakDetectors)
}

// LeakDetectorGroup shall contain a group of leak detection equipment that
// reports a unified status.
type LeakDetectorGroup struct {
	// Detectors shall contain the states of all leak detection devices in this
	// detector group. The value of the 'DataSourceUri' property, if present, shall
	// reference a resource of type 'LeakDetector'.
	Detectors []LeakDetectorArrayExcerpt
	// DetectorsCount
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
	Status Status
}
