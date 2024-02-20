/*
 * Nudm_SDM
 *
 * Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 
 *
 * API version: 2.3.0-alpha.5
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi




// IpAddr - Contains an IP adresse.
type IpAddr struct {

	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr string `json:"ipv4Addr,omitempty"`

	Ipv6Addr Ipv6Addr `json:"ipv6Addr,omitempty"`

	Ipv6Prefix Ipv6Prefix `json:"ipv6Prefix,omitempty"`
}

// AssertIpAddrRequired checks if the required fields are not zero-ed
func AssertIpAddrRequired(obj IpAddr) error {
	if err := AssertIpv6AddrRequired(obj.Ipv6Addr); err != nil {
		return err
	}
	if err := AssertIpv6PrefixRequired(obj.Ipv6Prefix); err != nil {
		return err
	}
	return nil
}

// AssertIpAddrConstraints checks if the values respects the defined constraints
func AssertIpAddrConstraints(obj IpAddr) error {
	return nil
}
