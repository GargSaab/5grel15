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

// checks if the VnGroupData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VnGroupData{}

// VnGroupData struct for VnGroupData
type VnGroupData struct {
	PduSessionTypes *PduSessionTypes `json:"pduSessionTypes,omitempty"`
	// String representing a Data Network as defined in clause 9A of 3GPP TS 23.003;  it shall contain either a DNN Network Identifier, or a full DNN with both the Network  Identifier and Operator Identifier, as specified in 3GPP TS 23.003 clause 9.1.1 and 9.1.2. It shall be coded as string in which the labels are separated by dots  (e.g. \"Label1.Label2.Label3\"). 
	Dnn *string `json:"dnn,omitempty"`
	SingleNssai *Snssai `json:"singleNssai,omitempty"`
	AppDescriptors []AppDescriptor `json:"appDescriptors,omitempty"`
	SecondaryAuth *bool `json:"secondaryAuth,omitempty"`
	DnAaaIpAddressAllocation *bool `json:"dnAaaIpAddressAllocation,omitempty"`
	DnAaaAddress NullableIpAddress `json:"dnAaaAddress,omitempty"`
	AdditionalDnAaaAddresses []IpAddress `json:"additionalDnAaaAddresses,omitempty"`
	// Fully Qualified Domain Name
	DnAaaFqdn *string `json:"dnAaaFqdn,omitempty"`
}

// NewVnGroupData instantiates a new VnGroupData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnGroupData() *VnGroupData {
	this := VnGroupData{}
	return &this
}

// NewVnGroupDataWithDefaults instantiates a new VnGroupData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnGroupDataWithDefaults() *VnGroupData {
	this := VnGroupData{}
	return &this
}

// GetPduSessionTypes returns the PduSessionTypes field value if set, zero value otherwise.
func (o *VnGroupData) GetPduSessionTypes() PduSessionTypes {
	if o == nil || IsNil(o.PduSessionTypes) {
		var ret PduSessionTypes
		return ret
	}
	return *o.PduSessionTypes
}

// GetPduSessionTypesOk returns a tuple with the PduSessionTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnGroupData) GetPduSessionTypesOk() (*PduSessionTypes, bool) {
	if o == nil || IsNil(o.PduSessionTypes) {
		return nil, false
	}
	return o.PduSessionTypes, true
}

// HasPduSessionTypes returns a boolean if a field has been set.
func (o *VnGroupData) HasPduSessionTypes() bool {
	if o != nil && !IsNil(o.PduSessionTypes) {
		return true
	}

	return false
}

// SetPduSessionTypes gets a reference to the given PduSessionTypes and assigns it to the PduSessionTypes field.
func (o *VnGroupData) SetPduSessionTypes(v PduSessionTypes) {
	o.PduSessionTypes = &v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *VnGroupData) GetDnn() string {
	if o == nil || IsNil(o.Dnn) {
		var ret string
		return ret
	}
	return *o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnGroupData) GetDnnOk() (*string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *VnGroupData) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given string and assigns it to the Dnn field.
func (o *VnGroupData) SetDnn(v string) {
	o.Dnn = &v
}

// GetSingleNssai returns the SingleNssai field value if set, zero value otherwise.
func (o *VnGroupData) GetSingleNssai() Snssai {
	if o == nil || IsNil(o.SingleNssai) {
		var ret Snssai
		return ret
	}
	return *o.SingleNssai
}

// GetSingleNssaiOk returns a tuple with the SingleNssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnGroupData) GetSingleNssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.SingleNssai) {
		return nil, false
	}
	return o.SingleNssai, true
}

// HasSingleNssai returns a boolean if a field has been set.
func (o *VnGroupData) HasSingleNssai() bool {
	if o != nil && !IsNil(o.SingleNssai) {
		return true
	}

	return false
}

// SetSingleNssai gets a reference to the given Snssai and assigns it to the SingleNssai field.
func (o *VnGroupData) SetSingleNssai(v Snssai) {
	o.SingleNssai = &v
}

// GetAppDescriptors returns the AppDescriptors field value if set, zero value otherwise.
func (o *VnGroupData) GetAppDescriptors() []AppDescriptor {
	if o == nil || IsNil(o.AppDescriptors) {
		var ret []AppDescriptor
		return ret
	}
	return o.AppDescriptors
}

// GetAppDescriptorsOk returns a tuple with the AppDescriptors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnGroupData) GetAppDescriptorsOk() ([]AppDescriptor, bool) {
	if o == nil || IsNil(o.AppDescriptors) {
		return nil, false
	}
	return o.AppDescriptors, true
}

// HasAppDescriptors returns a boolean if a field has been set.
func (o *VnGroupData) HasAppDescriptors() bool {
	if o != nil && !IsNil(o.AppDescriptors) {
		return true
	}

	return false
}

// SetAppDescriptors gets a reference to the given []AppDescriptor and assigns it to the AppDescriptors field.
func (o *VnGroupData) SetAppDescriptors(v []AppDescriptor) {
	o.AppDescriptors = v
}

// GetSecondaryAuth returns the SecondaryAuth field value if set, zero value otherwise.
func (o *VnGroupData) GetSecondaryAuth() bool {
	if o == nil || IsNil(o.SecondaryAuth) {
		var ret bool
		return ret
	}
	return *o.SecondaryAuth
}

// GetSecondaryAuthOk returns a tuple with the SecondaryAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnGroupData) GetSecondaryAuthOk() (*bool, bool) {
	if o == nil || IsNil(o.SecondaryAuth) {
		return nil, false
	}
	return o.SecondaryAuth, true
}

// HasSecondaryAuth returns a boolean if a field has been set.
func (o *VnGroupData) HasSecondaryAuth() bool {
	if o != nil && !IsNil(o.SecondaryAuth) {
		return true
	}

	return false
}

// SetSecondaryAuth gets a reference to the given bool and assigns it to the SecondaryAuth field.
func (o *VnGroupData) SetSecondaryAuth(v bool) {
	o.SecondaryAuth = &v
}

// GetDnAaaIpAddressAllocation returns the DnAaaIpAddressAllocation field value if set, zero value otherwise.
func (o *VnGroupData) GetDnAaaIpAddressAllocation() bool {
	if o == nil || IsNil(o.DnAaaIpAddressAllocation) {
		var ret bool
		return ret
	}
	return *o.DnAaaIpAddressAllocation
}

// GetDnAaaIpAddressAllocationOk returns a tuple with the DnAaaIpAddressAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnGroupData) GetDnAaaIpAddressAllocationOk() (*bool, bool) {
	if o == nil || IsNil(o.DnAaaIpAddressAllocation) {
		return nil, false
	}
	return o.DnAaaIpAddressAllocation, true
}

// HasDnAaaIpAddressAllocation returns a boolean if a field has been set.
func (o *VnGroupData) HasDnAaaIpAddressAllocation() bool {
	if o != nil && !IsNil(o.DnAaaIpAddressAllocation) {
		return true
	}

	return false
}

// SetDnAaaIpAddressAllocation gets a reference to the given bool and assigns it to the DnAaaIpAddressAllocation field.
func (o *VnGroupData) SetDnAaaIpAddressAllocation(v bool) {
	o.DnAaaIpAddressAllocation = &v
}

// GetDnAaaAddress returns the DnAaaAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnGroupData) GetDnAaaAddress() IpAddress {
	if o == nil || IsNil(o.DnAaaAddress.Get()) {
		var ret IpAddress
		return ret
	}
	return *o.DnAaaAddress.Get()
}

// GetDnAaaAddressOk returns a tuple with the DnAaaAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnGroupData) GetDnAaaAddressOk() (*IpAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnAaaAddress.Get(), o.DnAaaAddress.IsSet()
}

// HasDnAaaAddress returns a boolean if a field has been set.
func (o *VnGroupData) HasDnAaaAddress() bool {
	if o != nil && o.DnAaaAddress.IsSet() {
		return true
	}

	return false
}

// SetDnAaaAddress gets a reference to the given NullableIpAddress and assigns it to the DnAaaAddress field.
func (o *VnGroupData) SetDnAaaAddress(v IpAddress) {
	o.DnAaaAddress.Set(&v)
}
// SetDnAaaAddressNil sets the value for DnAaaAddress to be an explicit nil
func (o *VnGroupData) SetDnAaaAddressNil() {
	o.DnAaaAddress.Set(nil)
}

// UnsetDnAaaAddress ensures that no value is present for DnAaaAddress, not even an explicit nil
func (o *VnGroupData) UnsetDnAaaAddress() {
	o.DnAaaAddress.Unset()
}

// GetAdditionalDnAaaAddresses returns the AdditionalDnAaaAddresses field value if set, zero value otherwise.
func (o *VnGroupData) GetAdditionalDnAaaAddresses() []IpAddress {
	if o == nil || IsNil(o.AdditionalDnAaaAddresses) {
		var ret []IpAddress
		return ret
	}
	return o.AdditionalDnAaaAddresses
}

// GetAdditionalDnAaaAddressesOk returns a tuple with the AdditionalDnAaaAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnGroupData) GetAdditionalDnAaaAddressesOk() ([]IpAddress, bool) {
	if o == nil || IsNil(o.AdditionalDnAaaAddresses) {
		return nil, false
	}
	return o.AdditionalDnAaaAddresses, true
}

// HasAdditionalDnAaaAddresses returns a boolean if a field has been set.
func (o *VnGroupData) HasAdditionalDnAaaAddresses() bool {
	if o != nil && !IsNil(o.AdditionalDnAaaAddresses) {
		return true
	}

	return false
}

// SetAdditionalDnAaaAddresses gets a reference to the given []IpAddress and assigns it to the AdditionalDnAaaAddresses field.
func (o *VnGroupData) SetAdditionalDnAaaAddresses(v []IpAddress) {
	o.AdditionalDnAaaAddresses = v
}

// GetDnAaaFqdn returns the DnAaaFqdn field value if set, zero value otherwise.
func (o *VnGroupData) GetDnAaaFqdn() string {
	if o == nil || IsNil(o.DnAaaFqdn) {
		var ret string
		return ret
	}
	return *o.DnAaaFqdn
}

// GetDnAaaFqdnOk returns a tuple with the DnAaaFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnGroupData) GetDnAaaFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.DnAaaFqdn) {
		return nil, false
	}
	return o.DnAaaFqdn, true
}

// HasDnAaaFqdn returns a boolean if a field has been set.
func (o *VnGroupData) HasDnAaaFqdn() bool {
	if o != nil && !IsNil(o.DnAaaFqdn) {
		return true
	}

	return false
}

// SetDnAaaFqdn gets a reference to the given string and assigns it to the DnAaaFqdn field.
func (o *VnGroupData) SetDnAaaFqdn(v string) {
	o.DnAaaFqdn = &v
}

func (o VnGroupData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VnGroupData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PduSessionTypes) {
		toSerialize["pduSessionTypes"] = o.PduSessionTypes
	}
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	if !IsNil(o.SingleNssai) {
		toSerialize["singleNssai"] = o.SingleNssai
	}
	if !IsNil(o.AppDescriptors) {
		toSerialize["appDescriptors"] = o.AppDescriptors
	}
	if !IsNil(o.SecondaryAuth) {
		toSerialize["secondaryAuth"] = o.SecondaryAuth
	}
	if !IsNil(o.DnAaaIpAddressAllocation) {
		toSerialize["dnAaaIpAddressAllocation"] = o.DnAaaIpAddressAllocation
	}
	if o.DnAaaAddress.IsSet() {
		toSerialize["dnAaaAddress"] = o.DnAaaAddress.Get()
	}
	if !IsNil(o.AdditionalDnAaaAddresses) {
		toSerialize["additionalDnAaaAddresses"] = o.AdditionalDnAaaAddresses
	}
	if !IsNil(o.DnAaaFqdn) {
		toSerialize["dnAaaFqdn"] = o.DnAaaFqdn
	}
	return toSerialize, nil
}

type NullableVnGroupData struct {
	value *VnGroupData
	isSet bool
}

func (v NullableVnGroupData) Get() *VnGroupData {
	return v.value
}

func (v *NullableVnGroupData) Set(val *VnGroupData) {
	v.value = val
	v.isSet = true
}

func (v NullableVnGroupData) IsSet() bool {
	return v.isSet
}

func (v *NullableVnGroupData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnGroupData(val *VnGroupData) *NullableVnGroupData {
	return &NullableVnGroupData{value: val, isSet: true}
}

func (v NullableVnGroupData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnGroupData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


