/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type VnGroupData struct {

	PduSessionTypes PduSessionTypes `json:"pduSessionTypes,omitempty"`

	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn string `json:"dnn,omitempty"`

	SingleNssai Snssai `json:"singleNssai,omitempty"`

	AppDescriptors []AppDescriptor `json:"appDescriptors,omitempty"`

	SecondaryAuth bool `json:"secondaryAuth,omitempty"`

	DnAaaIpAddressAllocation bool `json:"dnAaaIpAddressAllocation,omitempty"`

	DnAaaAddress *IpAddress `json:"dnAaaAddress,omitempty"`

	AdditionalDnAaaAddresses []IpAddress `json:"additionalDnAaaAddresses,omitempty"`

	// Fully Qualified Domain Name
	DnAaaFqdn string `json:"dnAaaFqdn,omitempty"`
}

// AssertVnGroupDataRequired checks if the required fields are not zero-ed
func AssertVnGroupDataRequired(obj VnGroupData) error {
	if err := AssertPduSessionTypesRequired(obj.PduSessionTypes); err != nil {
		return err
	}
	if err := AssertSnssaiRequired(obj.SingleNssai); err != nil {
		return err
	}
	for _, el := range obj.AppDescriptors {
		if err := AssertAppDescriptorRequired(el); err != nil {
			return err
		}
	}
	if obj.DnAaaAddress != nil {
		if err := AssertIpAddressRequired(*obj.DnAaaAddress); err != nil {
			return err
		}
	}
	for _, el := range obj.AdditionalDnAaaAddresses {
		if err := AssertIpAddressRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertVnGroupDataConstraints checks if the values respects the defined constraints
func AssertVnGroupDataConstraints(obj VnGroupData) error {
	return nil
}
