/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// ProseSubscriptionData - Contains the ProSe Subscription Data.
type ProseSubscriptionData struct {

	ProseServiceAuth ProseServiceAuth `json:"proseServiceAuth,omitempty"`

	// String representing a bit rate; the prefixes follow the standard symbols from The International System of Units, and represent x1000 multipliers, with the exception that prefix \"K\" is used to represent the standard symbol \"k\". 
	NrUePc5Ambr string `json:"nrUePc5Ambr,omitempty"`

	ProseAllowedPlmn []ProSeAllowedPlmn `json:"proseAllowedPlmn,omitempty"`
}

// AssertProseSubscriptionDataRequired checks if the required fields are not zero-ed
func AssertProseSubscriptionDataRequired(obj ProseSubscriptionData) error {
	if err := AssertProseServiceAuthRequired(obj.ProseServiceAuth); err != nil {
		return err
	}
	for _, el := range obj.ProseAllowedPlmn {
		if err := AssertProSeAllowedPlmnRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertProseSubscriptionDataConstraints checks if the values respects the defined constraints
func AssertProseSubscriptionDataConstraints(obj ProseSubscriptionData) error {
	return nil
}
