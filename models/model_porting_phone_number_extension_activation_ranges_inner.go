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

// checks if the PortingPhoneNumberExtensionActivationRangesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortingPhoneNumberExtensionActivationRangesInner{}

// PortingPhoneNumberExtensionActivationRangesInner struct for PortingPhoneNumberExtensionActivationRangesInner
type PortingPhoneNumberExtensionActivationRangesInner struct {
	// Specifies the start of the activation range. Must be greater or equal the start of the extension range.
	StartAt *int32 `json:"start_at,omitempty"`
	// Specifies the end of the activation range. It must be no more than the end of the extension range.
	EndAt *int32 `json:"end_at,omitempty"`
}

// NewPortingPhoneNumberExtensionActivationRangesInner instantiates a new PortingPhoneNumberExtensionActivationRangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortingPhoneNumberExtensionActivationRangesInner() *PortingPhoneNumberExtensionActivationRangesInner {
	this := PortingPhoneNumberExtensionActivationRangesInner{}
	return &this
}

// NewPortingPhoneNumberExtensionActivationRangesInnerWithDefaults instantiates a new PortingPhoneNumberExtensionActivationRangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortingPhoneNumberExtensionActivationRangesInnerWithDefaults() *PortingPhoneNumberExtensionActivationRangesInner {
	this := PortingPhoneNumberExtensionActivationRangesInner{}
	return &this
}

// GetStartAt returns the StartAt field value if set, zero value otherwise.
func (o *PortingPhoneNumberExtensionActivationRangesInner) GetStartAt() int32 {
	if o == nil || IsNil(o.StartAt) {
		var ret int32
		return ret
	}
	return *o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingPhoneNumberExtensionActivationRangesInner) GetStartAtOk() (*int32, bool) {
	if o == nil || IsNil(o.StartAt) {
		return nil, false
	}
	return o.StartAt, true
}

// HasStartAt returns a boolean if a field has been set.
func (o *PortingPhoneNumberExtensionActivationRangesInner) HasStartAt() bool {
	if o != nil && !IsNil(o.StartAt) {
		return true
	}

	return false
}

// SetStartAt gets a reference to the given int32 and assigns it to the StartAt field.
func (o *PortingPhoneNumberExtensionActivationRangesInner) SetStartAt(v int32) {
	o.StartAt = &v
}

// GetEndAt returns the EndAt field value if set, zero value otherwise.
func (o *PortingPhoneNumberExtensionActivationRangesInner) GetEndAt() int32 {
	if o == nil || IsNil(o.EndAt) {
		var ret int32
		return ret
	}
	return *o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingPhoneNumberExtensionActivationRangesInner) GetEndAtOk() (*int32, bool) {
	if o == nil || IsNil(o.EndAt) {
		return nil, false
	}
	return o.EndAt, true
}

// HasEndAt returns a boolean if a field has been set.
func (o *PortingPhoneNumberExtensionActivationRangesInner) HasEndAt() bool {
	if o != nil && !IsNil(o.EndAt) {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given int32 and assigns it to the EndAt field.
func (o *PortingPhoneNumberExtensionActivationRangesInner) SetEndAt(v int32) {
	o.EndAt = &v
}

func (o PortingPhoneNumberExtensionActivationRangesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortingPhoneNumberExtensionActivationRangesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartAt) {
		toSerialize["start_at"] = o.StartAt
	}
	if !IsNil(o.EndAt) {
		toSerialize["end_at"] = o.EndAt
	}
	return toSerialize, nil
}

type NullablePortingPhoneNumberExtensionActivationRangesInner struct {
	value *PortingPhoneNumberExtensionActivationRangesInner
	isSet bool
}

func (v NullablePortingPhoneNumberExtensionActivationRangesInner) Get() *PortingPhoneNumberExtensionActivationRangesInner {
	return v.value
}

func (v *NullablePortingPhoneNumberExtensionActivationRangesInner) Set(val *PortingPhoneNumberExtensionActivationRangesInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePortingPhoneNumberExtensionActivationRangesInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePortingPhoneNumberExtensionActivationRangesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortingPhoneNumberExtensionActivationRangesInner(val *PortingPhoneNumberExtensionActivationRangesInner) *NullablePortingPhoneNumberExtensionActivationRangesInner {
	return &NullablePortingPhoneNumberExtensionActivationRangesInner{value: val, isSet: true}
}

func (v NullablePortingPhoneNumberExtensionActivationRangesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortingPhoneNumberExtensionActivationRangesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


