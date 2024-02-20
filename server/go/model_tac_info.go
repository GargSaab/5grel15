/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// TacInfo - contains tracking area information (tracking area codes).
type TacInfo struct {

	TacList []string `json:"tacList"`
}

// AssertTacInfoRequired checks if the required fields are not zero-ed
func AssertTacInfoRequired(obj TacInfo) error {
	elements := map[string]interface{}{
		"tacList": obj.TacList,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertTacInfoConstraints checks if the values respects the defined constraints
func AssertTacInfoConstraints(obj TacInfo) error {
	return nil
}