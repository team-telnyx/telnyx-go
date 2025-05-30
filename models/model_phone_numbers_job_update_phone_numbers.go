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

// checks if the PhoneNumbersJobUpdatePhoneNumbers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhoneNumbersJobUpdatePhoneNumbers{}

// PhoneNumbersJobUpdatePhoneNumbers struct for PhoneNumbersJobUpdatePhoneNumbers
type PhoneNumbersJobUpdatePhoneNumbers struct {
	Data *PhoneNumbersJob `json:"data,omitempty"`
}

// NewPhoneNumbersJobUpdatePhoneNumbers instantiates a new PhoneNumbersJobUpdatePhoneNumbers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneNumbersJobUpdatePhoneNumbers() *PhoneNumbersJobUpdatePhoneNumbers {
	this := PhoneNumbersJobUpdatePhoneNumbers{}
	return &this
}

// NewPhoneNumbersJobUpdatePhoneNumbersWithDefaults instantiates a new PhoneNumbersJobUpdatePhoneNumbers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneNumbersJobUpdatePhoneNumbersWithDefaults() *PhoneNumbersJobUpdatePhoneNumbers {
	this := PhoneNumbersJobUpdatePhoneNumbers{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *PhoneNumbersJobUpdatePhoneNumbers) GetData() PhoneNumbersJob {
	if o == nil || IsNil(o.Data) {
		var ret PhoneNumbersJob
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PhoneNumbersJobUpdatePhoneNumbers) GetDataOk() (*PhoneNumbersJob, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *PhoneNumbersJobUpdatePhoneNumbers) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PhoneNumbersJob and assigns it to the Data field.
func (o *PhoneNumbersJobUpdatePhoneNumbers) SetData(v PhoneNumbersJob) {
	o.Data = &v
}

func (o PhoneNumbersJobUpdatePhoneNumbers) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhoneNumbersJobUpdatePhoneNumbers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullablePhoneNumbersJobUpdatePhoneNumbers struct {
	value *PhoneNumbersJobUpdatePhoneNumbers
	isSet bool
}

func (v NullablePhoneNumbersJobUpdatePhoneNumbers) Get() *PhoneNumbersJobUpdatePhoneNumbers {
	return v.value
}

func (v *NullablePhoneNumbersJobUpdatePhoneNumbers) Set(val *PhoneNumbersJobUpdatePhoneNumbers) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneNumbersJobUpdatePhoneNumbers) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneNumbersJobUpdatePhoneNumbers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneNumbersJobUpdatePhoneNumbers(val *PhoneNumbersJobUpdatePhoneNumbers) *NullablePhoneNumbersJobUpdatePhoneNumbers {
	return &NullablePhoneNumbersJobUpdatePhoneNumbers{value: val, isSet: true}
}

func (v NullablePhoneNumbersJobUpdatePhoneNumbers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneNumbersJobUpdatePhoneNumbers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


