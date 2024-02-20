/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AerialUeIndication Indicates the Aerial service for the UE is allowed or not allowed, possible values are - AERIAL_UE_ALLOWED: Aerial service for the UE is allowed. - AERIAL_UE_NOT_ALLOWED: Aerial service for the UE is not allowed. 
type AerialUeIndication struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AerialUeIndication) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AerialUeIndication)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AerialUeIndication) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAerialUeIndication struct {
	value *AerialUeIndication
	isSet bool
}

func (v NullableAerialUeIndication) Get() *AerialUeIndication {
	return v.value
}

func (v *NullableAerialUeIndication) Set(val *AerialUeIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableAerialUeIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableAerialUeIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAerialUeIndication(val *AerialUeIndication) *NullableAerialUeIndication {
	return &NullableAerialUeIndication{value: val, isSet: true}
}

func (v NullableAerialUeIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAerialUeIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


