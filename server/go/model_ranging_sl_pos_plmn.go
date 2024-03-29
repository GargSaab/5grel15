/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// RangingSlPosPlmn - Contains the PLMN identities where the Ranging/SL Positioning services are authorised to use and the authorised Ranging/SL Positioning services on this given PLMNs.
type RangingSlPosPlmn struct {

	RangingSlPosPlmn PlmnId `json:"rangingSlPosPlmn"`

	RangingSlPosAllowed []RangingSlPosAllowed `json:"rangingSlPosAllowed,omitempty"`
}

// AssertRangingSlPosPlmnRequired checks if the required fields are not zero-ed
func AssertRangingSlPosPlmnRequired(obj RangingSlPosPlmn) error {
	elements := map[string]interface{}{
		"rangingSlPosPlmn": obj.RangingSlPosPlmn,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertPlmnIdRequired(obj.RangingSlPosPlmn); err != nil {
		return err
	}
	for _, el := range obj.RangingSlPosAllowed {
		if err := AssertRangingSlPosAllowedRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertRangingSlPosPlmnConstraints checks if the values respects the defined constraints
func AssertRangingSlPosPlmnConstraints(obj RangingSlPosPlmn) error {
	return nil
}
