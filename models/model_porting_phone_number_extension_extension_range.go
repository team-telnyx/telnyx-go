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

// checks if the PortingPhoneNumberExtensionExtensionRange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortingPhoneNumberExtensionExtensionRange{}

// PortingPhoneNumberExtensionExtensionRange Specifies the extension range for this porting phone number extension.
type PortingPhoneNumberExtensionExtensionRange struct {
	// Specifies the start of the extension range for this porting phone number extension.
	StartAt *int32 `json:"start_at,omitempty"`
	// Specifies the end of the extension range for this porting phone number extension.
	EndAt *int32 `json:"end_at,omitempty"`
}

// NewPortingPhoneNumberExtensionExtensionRange instantiates a new PortingPhoneNumberExtensionExtensionRange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortingPhoneNumberExtensionExtensionRange() *PortingPhoneNumberExtensionExtensionRange {
	this := PortingPhoneNumberExtensionExtensionRange{}
	return &this
}

// NewPortingPhoneNumberExtensionExtensionRangeWithDefaults instantiates a new PortingPhoneNumberExtensionExtensionRange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortingPhoneNumberExtensionExtensionRangeWithDefaults() *PortingPhoneNumberExtensionExtensionRange {
	this := PortingPhoneNumberExtensionExtensionRange{}
	return &this
}

// GetStartAt returns the StartAt field value if set, zero value otherwise.
func (o *PortingPhoneNumberExtensionExtensionRange) GetStartAt() int32 {
	if o == nil || IsNil(o.StartAt) {
		var ret int32
		return ret
	}
	return *o.StartAt
}

// GetStartAtOk returns a tuple with the StartAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingPhoneNumberExtensionExtensionRange) GetStartAtOk() (*int32, bool) {
	if o == nil || IsNil(o.StartAt) {
		return nil, false
	}
	return o.StartAt, true
}

// HasStartAt returns a boolean if a field has been set.
func (o *PortingPhoneNumberExtensionExtensionRange) HasStartAt() bool {
	if o != nil && !IsNil(o.StartAt) {
		return true
	}

	return false
}

// SetStartAt gets a reference to the given int32 and assigns it to the StartAt field.
func (o *PortingPhoneNumberExtensionExtensionRange) SetStartAt(v int32) {
	o.StartAt = &v
}

// GetEndAt returns the EndAt field value if set, zero value otherwise.
func (o *PortingPhoneNumberExtensionExtensionRange) GetEndAt() int32 {
	if o == nil || IsNil(o.EndAt) {
		var ret int32
		return ret
	}
	return *o.EndAt
}

// GetEndAtOk returns a tuple with the EndAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortingPhoneNumberExtensionExtensionRange) GetEndAtOk() (*int32, bool) {
	if o == nil || IsNil(o.EndAt) {
		return nil, false
	}
	return o.EndAt, true
}

// HasEndAt returns a boolean if a field has been set.
func (o *PortingPhoneNumberExtensionExtensionRange) HasEndAt() bool {
	if o != nil && !IsNil(o.EndAt) {
		return true
	}

	return false
}

// SetEndAt gets a reference to the given int32 and assigns it to the EndAt field.
func (o *PortingPhoneNumberExtensionExtensionRange) SetEndAt(v int32) {
	o.EndAt = &v
}

func (o PortingPhoneNumberExtensionExtensionRange) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortingPhoneNumberExtensionExtensionRange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartAt) {
		toSerialize["start_at"] = o.StartAt
	}
	if !IsNil(o.EndAt) {
		toSerialize["end_at"] = o.EndAt
	}
	return toSerialize, nil
}

type NullablePortingPhoneNumberExtensionExtensionRange struct {
	value *PortingPhoneNumberExtensionExtensionRange
	isSet bool
}

func (v NullablePortingPhoneNumberExtensionExtensionRange) Get() *PortingPhoneNumberExtensionExtensionRange {
	return v.value
}

func (v *NullablePortingPhoneNumberExtensionExtensionRange) Set(val *PortingPhoneNumberExtensionExtensionRange) {
	v.value = val
	v.isSet = true
}

func (v NullablePortingPhoneNumberExtensionExtensionRange) IsSet() bool {
	return v.isSet
}

func (v *NullablePortingPhoneNumberExtensionExtensionRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortingPhoneNumberExtensionExtensionRange(val *PortingPhoneNumberExtensionExtensionRange) *NullablePortingPhoneNumberExtensionExtensionRange {
	return &NullablePortingPhoneNumberExtensionExtensionRange{value: val, isSet: true}
}

func (v NullablePortingPhoneNumberExtensionExtensionRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortingPhoneNumberExtensionExtensionRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


