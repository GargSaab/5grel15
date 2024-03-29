/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// RangingSlPosAuth - Indicates whether the UE is authorized to use related services. 
type RangingSlPosAuth struct {

	RgSlPosTargetAuth UeAuth `json:"rgSlPosTargetAuth,omitempty"`

	RgSlPosSlRefAuth UeAuth `json:"rgSlPosSlRefAuth,omitempty"`

	RgSlPosLocAuth UeAuth `json:"rgSlPosLocAuth,omitempty"`

	RgSlPosClientAuth UeAuth `json:"rgSlPosClientAuth,omitempty"`

	RgSlPosServerAuth UeAuth `json:"rgSlPosServerAuth,omitempty"`
}

// AssertRangingSlPosAuthRequired checks if the required fields are not zero-ed
func AssertRangingSlPosAuthRequired(obj RangingSlPosAuth) error {
	if err := AssertUeAuthRequired(obj.RgSlPosTargetAuth); err != nil {
		return err
	}
	if err := AssertUeAuthRequired(obj.RgSlPosSlRefAuth); err != nil {
		return err
	}
	if err := AssertUeAuthRequired(obj.RgSlPosLocAuth); err != nil {
		return err
	}
	if err := AssertUeAuthRequired(obj.RgSlPosClientAuth); err != nil {
		return err
	}
	if err := AssertUeAuthRequired(obj.RgSlPosServerAuth); err != nil {
		return err
	}
	return nil
}

// AssertRangingSlPosAuthConstraints checks if the values respects the defined constraints
func AssertRangingSlPosAuthConstraints(obj RangingSlPosAuth) error {
	return nil
}
