/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type ExpectedUeBehaviourThreshold struct {

	ExpecedUeBehaviourDatasets []ExpecedUeBehaviourDataset `json:"expecedUeBehaviourDatasets,omitempty"`

	SingleNssais []Snssai `json:"singleNssais,omitempty"`

	Dnns []string `json:"dnns,omitempty"`

	ConfidenceLevel string `json:"confidenceLevel,omitempty"`

	AccuracyLevel string `json:"accuracyLevel,omitempty"`
}

// AssertExpectedUeBehaviourThresholdRequired checks if the required fields are not zero-ed
func AssertExpectedUeBehaviourThresholdRequired(obj ExpectedUeBehaviourThreshold) error {
	for _, el := range obj.ExpecedUeBehaviourDatasets {
		if err := AssertExpecedUeBehaviourDatasetRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.SingleNssais {
		if err := AssertSnssaiRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertExpectedUeBehaviourThresholdConstraints checks if the values respects the defined constraints
func AssertExpectedUeBehaviourThresholdConstraints(obj ExpectedUeBehaviourThreshold) error {
	return nil
}