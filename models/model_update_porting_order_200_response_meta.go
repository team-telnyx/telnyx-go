/*
Telnyx API

SIP trunking, SMS, MMS, Call Control and Telephony Data Services.

API version: 2.0.0
Contact: support@telnyx.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package telnyx

import (
	"encoding/json"
)

// checks if the UpdatePortingOrder200ResponseMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePortingOrder200ResponseMeta{}

// UpdatePortingOrder200ResponseMeta struct for UpdatePortingOrder200ResponseMeta
type UpdatePortingOrder200ResponseMeta struct {
	// Link to list all phone numbers
	PhoneNumbersUrl *string `json:"phone_numbers_url,omitempty"`
}

// NewUpdatePortingOrder200ResponseMeta instantiates a new UpdatePortingOrder200ResponseMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePortingOrder200ResponseMeta() *UpdatePortingOrder200ResponseMeta {
	this := UpdatePortingOrder200ResponseMeta{}
	return &this
}

// NewUpdatePortingOrder200ResponseMetaWithDefaults instantiates a new UpdatePortingOrder200ResponseMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePortingOrder200ResponseMetaWithDefaults() *UpdatePortingOrder200ResponseMeta {
	this := UpdatePortingOrder200ResponseMeta{}
	return &this
}

// GetPhoneNumbersUrl returns the PhoneNumbersUrl field value if set, zero value otherwise.
func (o *UpdatePortingOrder200ResponseMeta) GetPhoneNumbersUrl() string {
	if o == nil || IsNil(o.PhoneNumbersUrl) {
		var ret string
		return ret
	}
	return *o.PhoneNumbersUrl
}

// GetPhoneNumbersUrlOk returns a tuple with the PhoneNumbersUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePortingOrder200ResponseMeta) GetPhoneNumbersUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PhoneNumbersUrl) {
		return nil, false
	}
	return o.PhoneNumbersUrl, true
}

// HasPhoneNumbersUrl returns a boolean if a field has been set.
func (o *UpdatePortingOrder200ResponseMeta) HasPhoneNumbersUrl() bool {
	if o != nil && !IsNil(o.PhoneNumbersUrl) {
		return true
	}

	return false
}

// SetPhoneNumbersUrl gets a reference to the given string and assigns it to the PhoneNumbersUrl field.
func (o *UpdatePortingOrder200ResponseMeta) SetPhoneNumbersUrl(v string) {
	o.PhoneNumbersUrl = &v
}

func (o UpdatePortingOrder200ResponseMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePortingOrder200ResponseMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PhoneNumbersUrl) {
		toSerialize["phone_numbers_url"] = o.PhoneNumbersUrl
	}
	return toSerialize, nil
}

type NullableUpdatePortingOrder200ResponseMeta struct {
	value *UpdatePortingOrder200ResponseMeta
	isSet bool
}

func (v NullableUpdatePortingOrder200ResponseMeta) Get() *UpdatePortingOrder200ResponseMeta {
	return v.value
}

func (v *NullableUpdatePortingOrder200ResponseMeta) Set(val *UpdatePortingOrder200ResponseMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePortingOrder200ResponseMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePortingOrder200ResponseMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePortingOrder200ResponseMeta(val *UpdatePortingOrder200ResponseMeta) *NullableUpdatePortingOrder200ResponseMeta {
	return &NullableUpdatePortingOrder200ResponseMeta{value: val, isSet: true}
}

func (v NullableUpdatePortingOrder200ResponseMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePortingOrder200ResponseMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


