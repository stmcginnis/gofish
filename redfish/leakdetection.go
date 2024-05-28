//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// LeakDetection shall represent the leak detection functionality present in a service for a Redfish
// implementation.
type LeakDetection struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// LeakDetectorGroups shall contain an array of leak detection groups.
	LeakDetectorGroups []LeakDetectorGroup
	// LeakDetectors shall contain a link to a resource collection of type LeakDetectorCollection.
	leakDetectors string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// UnmarshalJSON unmarshals a LeakDetection object from the raw JSON.
func (leakdetection *LeakDetection) UnmarshalJSON(b []byte) error {
	type temp LeakDetection
	var t struct {
		temp
		LeakDetectors common.Link
	}

	err := json.Unmarshal(b, &t)
	if err != nil {
		return err
	}

	*leakdetection = LeakDetection(t.temp)

	// Extract the links to other entities for later
	leakdetection.leakDetectors = t.LeakDetectors.String()

	return nil
}

// LeakDetectors gets the leak detectors within this subsystem.
func (leakdetection *LeakDetection) LeakDetectors() ([]*LeakDetector, error) {
	return ListReferencedLeakDetectors(leakdetection.GetClient(), leakdetection.leakDetectors)
}

// GetLeakDetection will get a LeakDetection instance from the service.
func GetLeakDetection(c common.Client, uri string) (*LeakDetection, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var leakdetection LeakDetection
	err = json.NewDecoder(resp.Body).Decode(&leakdetection)
	if err != nil {
		return nil, err
	}

	leakdetection.SetClient(c)
	return &leakdetection, nil
}

// ListReferencedLeakDetections gets the collection of LeakDetection from
// a provided reference.
func ListReferencedLeakDetections(c common.Client, link string) ([]*LeakDetection, error) { //nolint:dupl
	if link == "" {
		return nil, nil
	}

	type GetResult struct {
		Item  *LeakDetection
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		leakdetection, err := GetLeakDetection(c, link)
		ch <- GetResult{Item: leakdetection, Link: link, Error: err}
	}

	var links []string
	var err error
	go func() {
		links, err = common.CollectList(get, c, link)
		if err != nil {
			collectionError.Failures[link] = err
		}
		close(ch)
	}()

	// Save unordered results into link-to-LeakDetection helper map.
	unorderedResults := map[string]*LeakDetection{}
	for r := range ch {
		if r.Error != nil {
			collectionError.Failures[r.Link] = r.Error
		} else {
			unorderedResults[r.Link] = r.Item
		}
	}

	if !collectionError.Empty() {
		return nil, collectionError
	}
	// Build the final ordered slice based on the original order from the links list.
	results := make([]*LeakDetection, len(links))
	for i, link := range links {
		results[i] = unorderedResults[link]
	}

	return results, nil
}

// LeakDetectorGroup shall contain a group of leak detection equipment that reports a unified status.
type LeakDetectorGroup struct {
	// Detectors shall contain the states of all leak detection devices in this detector group. The value of the
	// DataSourceUri property, if present, shall reference a resource of type LeakDetector.
	Detectors []LeakDetectorArrayExcerpt
	// Detectors@odata.count
	DetectorsCount int `json:"Detectors@odata.count"`
	// GroupName shall contain the name used to describe this group of leak detectors and related equipment.
	GroupName string
	// HumidityPercent shall contain the humidity, in percent units, for this resource. The value of the DataSourceUri
	// property, if present, shall reference a resource of type Sensor with the ReadingType property containing the
	// value 'Humidity'.
	HumidityPercent SensorExcerpt
}
