/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// UpConfidentiality - indicates whether UP confidentiality protection is required, preferred or not needed for all the traffic on the PDU Session. It shall comply with the provisions defined in table 5.4.3.5-1. 
type UpConfidentiality struct {
}

// AssertUpConfidentialityRequired checks if the required fields are not zero-ed
func AssertUpConfidentialityRequired(obj UpConfidentiality) error {
	return nil
}

// AssertUpConfidentialityConstraints checks if the values respects the defined constraints
func AssertUpConfidentialityConstraints(obj UpConfidentiality) error {
	return nil
}
