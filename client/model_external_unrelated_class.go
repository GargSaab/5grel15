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

// checks if the ExternalUnrelatedClass type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalUnrelatedClass{}

// ExternalUnrelatedClass struct for ExternalUnrelatedClass
type ExternalUnrelatedClass struct {
	LcsClientExternals []LcsClientExternal `json:"lcsClientExternals,omitempty"`
	AfExternals []AfExternal `json:"afExternals,omitempty"`
	LcsClientGroupExternals []LcsClientGroupExternal `json:"lcsClientGroupExternals,omitempty"`
}

// NewExternalUnrelatedClass instantiates a new ExternalUnrelatedClass object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalUnrelatedClass() *ExternalUnrelatedClass {
	this := ExternalUnrelatedClass{}
	return &this
}

// NewExternalUnrelatedClassWithDefaults instantiates a new ExternalUnrelatedClass object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalUnrelatedClassWithDefaults() *ExternalUnrelatedClass {
	this := ExternalUnrelatedClass{}
	return &this
}

// GetLcsClientExternals returns the LcsClientExternals field value if set, zero value otherwise.
func (o *ExternalUnrelatedClass) GetLcsClientExternals() []LcsClientExternal {
	if o == nil || IsNil(o.LcsClientExternals) {
		var ret []LcsClientExternal
		return ret
	}
	return o.LcsClientExternals
}

// GetLcsClientExternalsOk returns a tuple with the LcsClientExternals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalUnrelatedClass) GetLcsClientExternalsOk() ([]LcsClientExternal, bool) {
	if o == nil || IsNil(o.LcsClientExternals) {
		return nil, false
	}
	return o.LcsClientExternals, true
}

// HasLcsClientExternals returns a boolean if a field has been set.
func (o *ExternalUnrelatedClass) HasLcsClientExternals() bool {
	if o != nil && !IsNil(o.LcsClientExternals) {
		return true
	}

	return false
}

// SetLcsClientExternals gets a reference to the given []LcsClientExternal and assigns it to the LcsClientExternals field.
func (o *ExternalUnrelatedClass) SetLcsClientExternals(v []LcsClientExternal) {
	o.LcsClientExternals = v
}

// GetAfExternals returns the AfExternals field value if set, zero value otherwise.
func (o *ExternalUnrelatedClass) GetAfExternals() []AfExternal {
	if o == nil || IsNil(o.AfExternals) {
		var ret []AfExternal
		return ret
	}
	return o.AfExternals
}

// GetAfExternalsOk returns a tuple with the AfExternals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalUnrelatedClass) GetAfExternalsOk() ([]AfExternal, bool) {
	if o == nil || IsNil(o.AfExternals) {
		return nil, false
	}
	return o.AfExternals, true
}

// HasAfExternals returns a boolean if a field has been set.
func (o *ExternalUnrelatedClass) HasAfExternals() bool {
	if o != nil && !IsNil(o.AfExternals) {
		return true
	}

	return false
}

// SetAfExternals gets a reference to the given []AfExternal and assigns it to the AfExternals field.
func (o *ExternalUnrelatedClass) SetAfExternals(v []AfExternal) {
	o.AfExternals = v
}

// GetLcsClientGroupExternals returns the LcsClientGroupExternals field value if set, zero value otherwise.
func (o *ExternalUnrelatedClass) GetLcsClientGroupExternals() []LcsClientGroupExternal {
	if o == nil || IsNil(o.LcsClientGroupExternals) {
		var ret []LcsClientGroupExternal
		return ret
	}
	return o.LcsClientGroupExternals
}

// GetLcsClientGroupExternalsOk returns a tuple with the LcsClientGroupExternals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalUnrelatedClass) GetLcsClientGroupExternalsOk() ([]LcsClientGroupExternal, bool) {
	if o == nil || IsNil(o.LcsClientGroupExternals) {
		return nil, false
	}
	return o.LcsClientGroupExternals, true
}

// HasLcsClientGroupExternals returns a boolean if a field has been set.
func (o *ExternalUnrelatedClass) HasLcsClientGroupExternals() bool {
	if o != nil && !IsNil(o.LcsClientGroupExternals) {
		return true
	}

	return false
}

// SetLcsClientGroupExternals gets a reference to the given []LcsClientGroupExternal and assigns it to the LcsClientGroupExternals field.
func (o *ExternalUnrelatedClass) SetLcsClientGroupExternals(v []LcsClientGroupExternal) {
	o.LcsClientGroupExternals = v
}

func (o ExternalUnrelatedClass) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalUnrelatedClass) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LcsClientExternals) {
		toSerialize["lcsClientExternals"] = o.LcsClientExternals
	}
	if !IsNil(o.AfExternals) {
		toSerialize["afExternals"] = o.AfExternals
	}
	if !IsNil(o.LcsClientGroupExternals) {
		toSerialize["lcsClientGroupExternals"] = o.LcsClientGroupExternals
	}
	return toSerialize, nil
}

type NullableExternalUnrelatedClass struct {
	value *ExternalUnrelatedClass
	isSet bool
}

func (v NullableExternalUnrelatedClass) Get() *ExternalUnrelatedClass {
	return v.value
}

func (v *NullableExternalUnrelatedClass) Set(val *ExternalUnrelatedClass) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalUnrelatedClass) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalUnrelatedClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalUnrelatedClass(val *ExternalUnrelatedClass) *NullableExternalUnrelatedClass {
	return &NullableExternalUnrelatedClass{value: val, isSet: true}
}

func (v NullableExternalUnrelatedClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalUnrelatedClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


