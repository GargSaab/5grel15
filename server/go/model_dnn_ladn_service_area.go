/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type DnnLadnServiceArea struct {

	Dnn AccessAndMobilitySubscriptionDataSubscribedDnnListInner `json:"dnn"`

	LadnServiceArea []Tai `json:"ladnServiceArea"`
}

// AssertDnnLadnServiceAreaRequired checks if the required fields are not zero-ed
func AssertDnnLadnServiceAreaRequired(obj DnnLadnServiceArea) error {
	elements := map[string]interface{}{
		"dnn": obj.Dnn,
		"ladnServiceArea": obj.LadnServiceArea,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertAccessAndMobilitySubscriptionDataSubscribedDnnListInnerRequired(obj.Dnn); err != nil {
		return err
	}
	for _, el := range obj.LadnServiceArea {
		if err := AssertTaiRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertDnnLadnServiceAreaConstraints checks if the values respects the defined constraints
func AssertDnnLadnServiceAreaConstraints(obj DnnLadnServiceArea) error {
	return nil
}