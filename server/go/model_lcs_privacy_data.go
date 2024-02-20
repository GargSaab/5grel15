/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type LcsPrivacyData struct {

	Lpi Lpi `json:"lpi,omitempty"`

	UnrelatedClass UnrelatedClass `json:"unrelatedClass,omitempty"`

	PlmnOperatorClasses []PlmnOperatorClass `json:"plmnOperatorClasses,omitempty"`

	EvtRptExpectedArea GeographicArea `json:"evtRptExpectedArea,omitempty"`

	AreaUsageInd AreaUsageInd `json:"areaUsageInd,omitempty"`

	UpLocRepIndAf UpLocRepIndAf `json:"upLocRepIndAf,omitempty"`
}

// AssertLcsPrivacyDataRequired checks if the required fields are not zero-ed
func AssertLcsPrivacyDataRequired(obj LcsPrivacyData) error {
	if err := AssertLpiRequired(obj.Lpi); err != nil {
		return err
	}
	if err := AssertUnrelatedClassRequired(obj.UnrelatedClass); err != nil {
		return err
	}
	for _, el := range obj.PlmnOperatorClasses {
		if err := AssertPlmnOperatorClassRequired(el); err != nil {
			return err
		}
	}
	if err := AssertGeographicAreaRequired(obj.EvtRptExpectedArea); err != nil {
		return err
	}
	if err := AssertAreaUsageIndRequired(obj.AreaUsageInd); err != nil {
		return err
	}
	if err := AssertUpLocRepIndAfRequired(obj.UpLocRepIndAf); err != nil {
		return err
	}
	return nil
}

// AssertLcsPrivacyDataConstraints checks if the values respects the defined constraints
func AssertLcsPrivacyDataConstraints(obj LcsPrivacyData) error {
	return nil
}