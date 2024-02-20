/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2023, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.5
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the MbsrOperationAllowed type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsrOperationAllowed{}

// MbsrOperationAllowed struct for MbsrOperationAllowed
type MbsrOperationAllowed struct {
	MbsrOperationAllowedInd *bool `json:"mbsrOperationAllowedInd,omitempty"`
	MbsrValidTimePeriod *ValidTimePeriod `json:"mbsrValidTimePeriod,omitempty"`
}

// NewMbsrOperationAllowed instantiates a new MbsrOperationAllowed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsrOperationAllowed() *MbsrOperationAllowed {
	this := MbsrOperationAllowed{}
	return &this
}

// NewMbsrOperationAllowedWithDefaults instantiates a new MbsrOperationAllowed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsrOperationAllowedWithDefaults() *MbsrOperationAllowed {
	this := MbsrOperationAllowed{}
	return &this
}

// GetMbsrOperationAllowedInd returns the MbsrOperationAllowedInd field value if set, zero value otherwise.
func (o *MbsrOperationAllowed) GetMbsrOperationAllowedInd() bool {
	if o == nil || IsNil(o.MbsrOperationAllowedInd) {
		var ret bool
		return ret
	}
	return *o.MbsrOperationAllowedInd
}

// GetMbsrOperationAllowedIndOk returns a tuple with the MbsrOperationAllowedInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsrOperationAllowed) GetMbsrOperationAllowedIndOk() (*bool, bool) {
	if o == nil || IsNil(o.MbsrOperationAllowedInd) {
		return nil, false
	}
	return o.MbsrOperationAllowedInd, true
}

// HasMbsrOperationAllowedInd returns a boolean if a field has been set.
func (o *MbsrOperationAllowed) HasMbsrOperationAllowedInd() bool {
	if o != nil && !IsNil(o.MbsrOperationAllowedInd) {
		return true
	}

	return false
}

// SetMbsrOperationAllowedInd gets a reference to the given bool and assigns it to the MbsrOperationAllowedInd field.
func (o *MbsrOperationAllowed) SetMbsrOperationAllowedInd(v bool) {
	o.MbsrOperationAllowedInd = &v
}

// GetMbsrValidTimePeriod returns the MbsrValidTimePeriod field value if set, zero value otherwise.
func (o *MbsrOperationAllowed) GetMbsrValidTimePeriod() ValidTimePeriod {
	if o == nil || IsNil(o.MbsrValidTimePeriod) {
		var ret ValidTimePeriod
		return ret
	}
	return *o.MbsrValidTimePeriod
}

// GetMbsrValidTimePeriodOk returns a tuple with the MbsrValidTimePeriod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsrOperationAllowed) GetMbsrValidTimePeriodOk() (*ValidTimePeriod, bool) {
	if o == nil || IsNil(o.MbsrValidTimePeriod) {
		return nil, false
	}
	return o.MbsrValidTimePeriod, true
}

// HasMbsrValidTimePeriod returns a boolean if a field has been set.
func (o *MbsrOperationAllowed) HasMbsrValidTimePeriod() bool {
	if o != nil && !IsNil(o.MbsrValidTimePeriod) {
		return true
	}

	return false
}

// SetMbsrValidTimePeriod gets a reference to the given ValidTimePeriod and assigns it to the MbsrValidTimePeriod field.
func (o *MbsrOperationAllowed) SetMbsrValidTimePeriod(v ValidTimePeriod) {
	o.MbsrValidTimePeriod = &v
}

func (o MbsrOperationAllowed) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsrOperationAllowed) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MbsrOperationAllowedInd) {
		toSerialize["mbsrOperationAllowedInd"] = o.MbsrOperationAllowedInd
	}
	if !IsNil(o.MbsrValidTimePeriod) {
		toSerialize["mbsrValidTimePeriod"] = o.MbsrValidTimePeriod
	}
	return toSerialize, nil
}

type NullableMbsrOperationAllowed struct {
	value *MbsrOperationAllowed
	isSet bool
}

func (v NullableMbsrOperationAllowed) Get() *MbsrOperationAllowed {
	return v.value
}

func (v *NullableMbsrOperationAllowed) Set(val *MbsrOperationAllowed) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsrOperationAllowed) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsrOperationAllowed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsrOperationAllowed(val *MbsrOperationAllowed) *NullableMbsrOperationAllowed {
	return &NullableMbsrOperationAllowed{value: val, isSet: true}
}

func (v NullableMbsrOperationAllowed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsrOperationAllowed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


