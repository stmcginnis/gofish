//
// SPDX-License-Identifier: BSD-3-Clause
//

package swordfish

import (
	"encoding/json"

	"github.com/stmcginnis/gofish/common"
)

// FeaturesRegistry shall be used to represent a Feature registry for a Redfish implementation.
type FeaturesRegistry struct {
	common.Entity
	// ODataContext is the odata context.
	ODataContext string `json:"@odata.context"`
	// ODataEtag is the odata etag.
	ODataEtag string `json:"@odata.etag"`
	// ODataType is the odata type.
	ODataType string `json:"@odata.type"`
	// Description provides a description of this resource.
	Description string
	// Features shall represent the suffix to be used in the FeatureId and shall be unique within this message
	// registry.
	Features []SupportedFeature
	// Language shall be a string consisting of an RFC 5646 language code.
	Language string
	// Oem shall contain the OEM extensions. All values for properties that this object contains shall conform to the
	// Redfish Specification-described requirements.
	OEM json.RawMessage `json:"Oem"`
	// OwningEntity shall be a string that represents the publisher of this registry.
	OwningEntity string
	// RegistryPrefix shall be the prefix used in IDs which uniquely identifies all of the Features in this registry as
	// belonging to this registry.
	RegistryPrefix string
	// RegistryVersion shall be the version of this message registry. The format of this string shall be of the format
	// majorversion.minorversion.errata.
	RegistryVersion string
}

// GetFeaturesRegistry will get a FeaturesRegistry instance from the service.
func GetFeaturesRegistry(c common.Client, uri string) (*FeaturesRegistry, error) {
	resp, err := c.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var featuresregistry FeaturesRegistry
	err = json.NewDecoder(resp.Body).Decode(&featuresregistry)
	if err != nil {
		return nil, err
	}

	featuresregistry.SetClient(c)
	return &featuresregistry, nil
}

// ListReferencedFeaturesRegistrys gets the collection of FeaturesRegistry from
// a provided reference.
func ListReferencedFeaturesRegistrys(c common.Client, link string) ([]*FeaturesRegistry, error) {
	var result []*FeaturesRegistry
	if link == "" {
		return result, nil
	}

	type GetResult struct {
		Item  *FeaturesRegistry
		Link  string
		Error error
	}

	ch := make(chan GetResult)
	collectionError := common.NewCollectionError()
	get := func(link string) {
		featuresregistry, err := GetFeaturesRegistry(c, link)
		ch <- GetResult{Item: featuresregistry, Link: link, Error: err}
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

// FeaturesRegistryProperty shall represent the suffix to be used in the Feature and shall be unique within this
// registry.
type FeaturesRegistryProperty struct {
}

// SupportedFeature shall name a feature.
type SupportedFeature struct {
	// CorrespondingProfileDefinition shall define a profile definition that contains the named profile declaration.
	CorrespondingProfileDefinition string
	// Description provides a description of this resource.
	Description string
	// FeatureName shall be the unique name of the feature prefixed by the defining organization separated by a period
	// (e.g. 'vendor.feature').
	FeatureName string
	// Version shall uniquely identify the version of the feature, using the major.minor.errata format.
	Version string
}
