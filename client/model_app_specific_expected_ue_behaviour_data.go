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

// checks if the AppSpecificExpectedUeBehaviourData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppSpecificExpectedUeBehaviourData{}

// AppSpecificExpectedUeBehaviourData struct for AppSpecificExpectedUeBehaviourData
type AppSpecificExpectedUeBehaviourData struct {
	// String providing an application identifier.
	AppId *string `json:"appId,omitempty"`
	TrafficFilters []FlowInfo `json:"trafficFilters,omitempty"`
	// indicating a time in seconds.
	ExpectedInactivityTime *int32 `json:"expectedInactivityTime,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime *time.Time `json:"validityTime,omitempty"`
	ConfidenceLevel *string `json:"confidenceLevel,omitempty"`
	AccuracyLevel *string `json:"accuracyLevel,omitempty"`
}

// NewAppSpecificExpectedUeBehaviourData instantiates a new AppSpecificExpectedUeBehaviourData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppSpecificExpectedUeBehaviourData() *AppSpecificExpectedUeBehaviourData {
	this := AppSpecificExpectedUeBehaviourData{}
	return &this
}

// NewAppSpecificExpectedUeBehaviourDataWithDefaults instantiates a new AppSpecificExpectedUeBehaviourData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppSpecificExpectedUeBehaviourDataWithDefaults() *AppSpecificExpectedUeBehaviourData {
	this := AppSpecificExpectedUeBehaviourData{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *AppSpecificExpectedUeBehaviourData) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSpecificExpectedUeBehaviourData) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *AppSpecificExpectedUeBehaviourData) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *AppSpecificExpectedUeBehaviourData) SetAppId(v string) {
	o.AppId = &v
}

// GetTrafficFilters returns the TrafficFilters field value if set, zero value otherwise.
func (o *AppSpecificExpectedUeBehaviourData) GetTrafficFilters() []FlowInfo {
	if o == nil || IsNil(o.TrafficFilters) {
		var ret []FlowInfo
		return ret
	}
	return o.TrafficFilters
}

// GetTrafficFiltersOk returns a tuple with the TrafficFilters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSpecificExpectedUeBehaviourData) GetTrafficFiltersOk() ([]FlowInfo, bool) {
	if o == nil || IsNil(o.TrafficFilters) {
		return nil, false
	}
	return o.TrafficFilters, true
}

// HasTrafficFilters returns a boolean if a field has been set.
func (o *AppSpecificExpectedUeBehaviourData) HasTrafficFilters() bool {
	if o != nil && !IsNil(o.TrafficFilters) {
		return true
	}

	return false
}

// SetTrafficFilters gets a reference to the given []FlowInfo and assigns it to the TrafficFilters field.
func (o *AppSpecificExpectedUeBehaviourData) SetTrafficFilters(v []FlowInfo) {
	o.TrafficFilters = v
}

// GetExpectedInactivityTime returns the ExpectedInactivityTime field value if set, zero value otherwise.
func (o *AppSpecificExpectedUeBehaviourData) GetExpectedInactivityTime() int32 {
	if o == nil || IsNil(o.ExpectedInactivityTime) {
		var ret int32
		return ret
	}
	return *o.ExpectedInactivityTime
}

// GetExpectedInactivityTimeOk returns a tuple with the ExpectedInactivityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSpecificExpectedUeBehaviourData) GetExpectedInactivityTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpectedInactivityTime) {
		return nil, false
	}
	return o.ExpectedInactivityTime, true
}

// HasExpectedInactivityTime returns a boolean if a field has been set.
func (o *AppSpecificExpectedUeBehaviourData) HasExpectedInactivityTime() bool {
	if o != nil && !IsNil(o.ExpectedInactivityTime) {
		return true
	}

	return false
}

// SetExpectedInactivityTime gets a reference to the given int32 and assigns it to the ExpectedInactivityTime field.
func (o *AppSpecificExpectedUeBehaviourData) SetExpectedInactivityTime(v int32) {
	o.ExpectedInactivityTime = &v
}

// GetValidityTime returns the ValidityTime field value if set, zero value otherwise.
func (o *AppSpecificExpectedUeBehaviourData) GetValidityTime() time.Time {
	if o == nil || IsNil(o.ValidityTime) {
		var ret time.Time
		return ret
	}
	return *o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSpecificExpectedUeBehaviourData) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ValidityTime) {
		return nil, false
	}
	return o.ValidityTime, true
}

// HasValidityTime returns a boolean if a field has been set.
func (o *AppSpecificExpectedUeBehaviourData) HasValidityTime() bool {
	if o != nil && !IsNil(o.ValidityTime) {
		return true
	}

	return false
}

// SetValidityTime gets a reference to the given time.Time and assigns it to the ValidityTime field.
func (o *AppSpecificExpectedUeBehaviourData) SetValidityTime(v time.Time) {
	o.ValidityTime = &v
}

// GetConfidenceLevel returns the ConfidenceLevel field value if set, zero value otherwise.
func (o *AppSpecificExpectedUeBehaviourData) GetConfidenceLevel() string {
	if o == nil || IsNil(o.ConfidenceLevel) {
		var ret string
		return ret
	}
	return *o.ConfidenceLevel
}

// GetConfidenceLevelOk returns a tuple with the ConfidenceLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSpecificExpectedUeBehaviourData) GetConfidenceLevelOk() (*string, bool) {
	if o == nil || IsNil(o.ConfidenceLevel) {
		return nil, false
	}
	return o.ConfidenceLevel, true
}

// HasConfidenceLevel returns a boolean if a field has been set.
func (o *AppSpecificExpectedUeBehaviourData) HasConfidenceLevel() bool {
	if o != nil && !IsNil(o.ConfidenceLevel) {
		return true
	}

	return false
}

// SetConfidenceLevel gets a reference to the given string and assigns it to the ConfidenceLevel field.
func (o *AppSpecificExpectedUeBehaviourData) SetConfidenceLevel(v string) {
	o.ConfidenceLevel = &v
}

// GetAccuracyLevel returns the AccuracyLevel field value if set, zero value otherwise.
func (o *AppSpecificExpectedUeBehaviourData) GetAccuracyLevel() string {
	if o == nil || IsNil(o.AccuracyLevel) {
		var ret string
		return ret
	}
	return *o.AccuracyLevel
}

// GetAccuracyLevelOk returns a tuple with the AccuracyLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppSpecificExpectedUeBehaviourData) GetAccuracyLevelOk() (*string, bool) {
	if o == nil || IsNil(o.AccuracyLevel) {
		return nil, false
	}
	return o.AccuracyLevel, true
}

// HasAccuracyLevel returns a boolean if a field has been set.
func (o *AppSpecificExpectedUeBehaviourData) HasAccuracyLevel() bool {
	if o != nil && !IsNil(o.AccuracyLevel) {
		return true
	}

	return false
}

// SetAccuracyLevel gets a reference to the given string and assigns it to the AccuracyLevel field.
func (o *AppSpecificExpectedUeBehaviourData) SetAccuracyLevel(v string) {
	o.AccuracyLevel = &v
}

func (o AppSpecificExpectedUeBehaviourData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppSpecificExpectedUeBehaviourData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.TrafficFilters) {
		toSerialize["trafficFilters"] = o.TrafficFilters
	}
	if !IsNil(o.ExpectedInactivityTime) {
		toSerialize["expectedInactivityTime"] = o.ExpectedInactivityTime
	}
	if !IsNil(o.ValidityTime) {
		toSerialize["validityTime"] = o.ValidityTime
	}
	if !IsNil(o.ConfidenceLevel) {
		toSerialize["confidenceLevel"] = o.ConfidenceLevel
	}
	if !IsNil(o.AccuracyLevel) {
		toSerialize["accuracyLevel"] = o.AccuracyLevel
	}
	return toSerialize, nil
}

type NullableAppSpecificExpectedUeBehaviourData struct {
	value *AppSpecificExpectedUeBehaviourData
	isSet bool
}

func (v NullableAppSpecificExpectedUeBehaviourData) Get() *AppSpecificExpectedUeBehaviourData {
	return v.value
}

func (v *NullableAppSpecificExpectedUeBehaviourData) Set(val *AppSpecificExpectedUeBehaviourData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppSpecificExpectedUeBehaviourData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppSpecificExpectedUeBehaviourData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppSpecificExpectedUeBehaviourData(val *AppSpecificExpectedUeBehaviourData) *NullableAppSpecificExpectedUeBehaviourData {
	return &NullableAppSpecificExpectedUeBehaviourData{value: val, isSet: true}
}

func (v NullableAppSpecificExpectedUeBehaviourData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppSpecificExpectedUeBehaviourData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


