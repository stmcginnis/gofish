//
// SPDX-License-Identifier: BSD-3-Clause
//

package redfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

type LeakDetectorType string

const (
	// MoistureLeakDetectorType A moisture sensor.
	MoistureLeakDetectorType LeakDetectorType = "Moisture"
	// FloatSwitchLeakDetectorType A float switch.
	FloatSwitchLeakDetectorType LeakDetectorType = "FloatSwitch"
)

// LeakDetector shall represent a state-based or digital-value leak detector for a Redfish implementation.
type LeakDetector struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// DetectorState shall contain the state of the leak detector.
	DetectorState common.Health
	// LeakDetectorType shall contain the reading type of the leak detection sensor.
	LeakDetectorType LeakDetectorType
	// Location shall indicate the location information for this leak detector.
	Location common.Location
	// Manufacturer shall contain the name of the organization responsible for producing the leak detector. This
	// organization may be the entity from whom the leak detector is purchased, but this is not necessarily true.
	Manufacturer string
	// Model shall contain the name by which the manufacturer generally refers to the leak detector.
	Model string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// PartNumber shall contain a part number assigned by the organization that is responsible for producing or
	// manufacturing the leak detector.
	PartNumber string
	// PhysicalContext shall contain a description of the affected component or region within the equipment to which
	// this leak detector applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region within the equipment to which this
	// leak detector applies. This property generally differentiates multiple leak detectors within the same
	// PhysicalContext instance.
	PhysicalSubContext PhysicalSubContext
	// SKU shall contain the stock-keeping unit number for this leak detector.
	SKU string
	// SensingFrequency shall contain the time interval between readings of the physical leak detector.
	SensingFrequency float64
	// SerialNumber shall contain a manufacturer-allocated number that identifies the leak detector.
	SerialNumber string
	// SparePartNumber shall contain the spare part number of the leak detector.
	SparePartNumber string
	// Status shall contain any status or health properties of the resource.
	Status common.Status
}

// GetLeakDetector will get a LeakDetector instance from the service.
func GetLeakDetector(c common.Client, uri string) (*LeakDetector, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var leakdetector LeakDetector
	err = json.NewDecoder(resp.Body).Decode(&leakdetector)
	if err != nil {
		return nil, err
	}

	leakdetector.SetClient(c)
	return &leakdetector, nil
}

// ListReferencedLeakDetectors gets the collection of LeakDetector from
// a provided reference.
func ListReferencedLeakDetectors(c common.Client, link string) ([]*LeakDetector, error) {
	var result []*LeakDetector
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *LeakDetector
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		leakdetector, err := GetLeakDetector(c, link)
		ch <- GetResult{Item: leakdetector, Link: link, Error: err}
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

// LeakDetectorArrayExcerpt shall represent a state-based or digital-value leak detector for a Redfish
// implementation.
type LeakDetectorArrayExcerpt struct {
	// DataSourceURI shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy.
	DataSourceURI string
	// DetectorState shall contain the state of the leak detector.
	DetectorState common.Health
	// PhysicalContext shall contain a description of the affected component or region within the equipment to which
	// this leak detector applies.
	PhysicalContext PhysicalContext
	// PhysicalSubContext shall contain a description of the usage or sub-region within the equipment to which this
	// leak detector applies. This property generally differentiates multiple leak detectors within the same
	// PhysicalContext instance.
	PhysicalSubContext PhysicalSubContext
}

// LeakDetectorExcerpt shall represent a state-based or digital-value leak detector for a Redfish implementation.
type LeakDetectorExcerpt struct {
	// DataSourceURI shall contain a URI to the resource that provides the source of the excerpt contained within this
	// copy.
	DataSourceURI string
	// DetectorState shall contain the state of the leak detector.
	DetectorState common.Health
}
