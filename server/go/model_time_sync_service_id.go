/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi


import (
	"errors"
)



// TimeSyncServiceId - Time Synchronization Service ID
type TimeSyncServiceId struct {

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	SNssai Snssai `json:"sNssai,omitempty"`

	Reference string `json:"reference"`

	TempVals []TemporalValidity `json:"tempVals,omitempty"`

	CoverageArea []Tai `json:"coverageArea,omitempty"`

	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	UuTimeSyncErrBdgt int32 `json:"uuTimeSyncErrBdgt,omitempty"`
}

// AssertTimeSyncServiceIdRequired checks if the required fields are not zero-ed
func AssertTimeSyncServiceIdRequired(obj TimeSyncServiceId) error {
	elements := map[string]interface{}{
		"reference": obj.Reference,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertSnssaiRequired(obj.SNssai); err != nil {
		return err
	}
	for _, el := range obj.TempVals {
		if err := AssertTemporalValidityRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.CoverageArea {
		if err := AssertTaiRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertTimeSyncServiceIdConstraints checks if the values respects the defined constraints
func AssertTimeSyncServiceIdConstraints(obj TimeSyncServiceId) error {
	if obj.UuTimeSyncErrBdgt < 0 {
		return &ParsingError{Err: errors.New(errMsgMinValueConstraint)}
	}
	return nil
}