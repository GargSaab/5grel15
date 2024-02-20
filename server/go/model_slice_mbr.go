/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// SliceMbr - MBR related to slice
type SliceMbr struct {

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	Uplink string `json:"uplink"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	Downlink string `json:"downlink"`
}

// AssertSliceMbrRequired checks if the required fields are not zero-ed
func AssertSliceMbrRequired(obj SliceMbr) error {
	elements := map[string]interface{}{
		"uplink": obj.Uplink,
		"downlink": obj.Downlink,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertSliceMbrConstraints checks if the values respects the defined constraints
func AssertSliceMbrConstraints(obj SliceMbr) error {
	return nil
}
