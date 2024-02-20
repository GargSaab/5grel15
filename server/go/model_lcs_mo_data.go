/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type LcsMoData struct {

	AllowedServiceClasses []LcsMoServiceClass `json:"allowedServiceClasses"`

	MoAssistanceDataTypes LcsBroadcastAssistanceTypesData `json:"moAssistanceDataTypes,omitempty"`
}

// AssertLcsMoDataRequired checks if the required fields are not zero-ed
func AssertLcsMoDataRequired(obj LcsMoData) error {
	elements := map[string]interface{}{
		"allowedServiceClasses": obj.AllowedServiceClasses,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	for _, el := range obj.AllowedServiceClasses {
		if err := AssertLcsMoServiceClassRequired(el); err != nil {
			return err
		}
	}
	if err := AssertLcsBroadcastAssistanceTypesDataRequired(obj.MoAssistanceDataTypes); err != nil {
		return err
	}
	return nil
}

// AssertLcsMoDataConstraints checks if the values respects the defined constraints
func AssertLcsMoDataConstraints(obj LcsMoData) error {
	return nil
}
