/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




type PlmnRestriction struct {

	RatRestrictions []RatType `json:"ratRestrictions,omitempty"`

	ForbiddenAreas []Area `json:"forbiddenAreas,omitempty"`

	ServiceAreaRestriction ServiceAreaRestriction `json:"serviceAreaRestriction,omitempty"`

	CoreNetworkTypeRestrictions []CoreNetworkType `json:"coreNetworkTypeRestrictions,omitempty"`

	AccessTypeRestrictions []AccessType `json:"accessTypeRestrictions,omitempty"`

	PrimaryRatRestrictions []RatType `json:"primaryRatRestrictions,omitempty"`

	SecondaryRatRestrictions []RatType `json:"secondaryRatRestrictions,omitempty"`
}

// AssertPlmnRestrictionRequired checks if the required fields are not zero-ed
func AssertPlmnRestrictionRequired(obj PlmnRestriction) error {
	for _, el := range obj.RatRestrictions {
		if err := AssertRatTypeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.ForbiddenAreas {
		if err := AssertAreaRequired(el); err != nil {
			return err
		}
	}
	if err := AssertServiceAreaRestrictionRequired(obj.ServiceAreaRestriction); err != nil {
		return err
	}
	for _, el := range obj.CoreNetworkTypeRestrictions {
		if err := AssertCoreNetworkTypeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.PrimaryRatRestrictions {
		if err := AssertRatTypeRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.SecondaryRatRestrictions {
		if err := AssertRatTypeRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertPlmnRestrictionConstraints checks if the values respects the defined constraints
func AssertPlmnRestrictionConstraints(obj PlmnRestriction) error {
	return nil
}
