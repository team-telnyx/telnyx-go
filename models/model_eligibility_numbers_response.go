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

// checks if the EligibilityNumbersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EligibilityNumbersResponse{}

// EligibilityNumbersResponse struct for EligibilityNumbersResponse
type EligibilityNumbersResponse struct {
	// List of phone numbers with their eligibility status.
	PhoneNumbers []EligibilityNumberResponse `json:"phone_numbers,omitempty"`
}

// NewEligibilityNumbersResponse instantiates a new EligibilityNumbersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEligibilityNumbersResponse() *EligibilityNumbersResponse {
	this := EligibilityNumbersResponse{}
	return &this
}

// NewEligibilityNumbersResponseWithDefaults instantiates a new EligibilityNumbersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEligibilityNumbersResponseWithDefaults() *EligibilityNumbersResponse {
	this := EligibilityNumbersResponse{}
	return &this
}

// GetPhoneNumbers returns the PhoneNumbers field value if set, zero value otherwise.
func (o *EligibilityNumbersResponse) GetPhoneNumbers() []EligibilityNumberResponse {
	if o == nil || IsNil(o.PhoneNumbers) {
		var ret []EligibilityNumberResponse
		return ret
	}
	return o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EligibilityNumbersResponse) GetPhoneNumbersOk() ([]EligibilityNumberResponse, bool) {
	if o == nil || IsNil(o.PhoneNumbers) {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// HasPhoneNumbers returns a boolean if a field has been set.
func (o *EligibilityNumbersResponse) HasPhoneNumbers() bool {
	if o != nil && !IsNil(o.PhoneNumbers) {
		return true
	}

	return false
}

// SetPhoneNumbers gets a reference to the given []EligibilityNumberResponse and assigns it to the PhoneNumbers field.
func (o *EligibilityNumbersResponse) SetPhoneNumbers(v []EligibilityNumberResponse) {
	o.PhoneNumbers = v
}

func (o EligibilityNumbersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EligibilityNumbersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PhoneNumbers) {
		toSerialize["phone_numbers"] = o.PhoneNumbers
	}
	return toSerialize, nil
}

type NullableEligibilityNumbersResponse struct {
	value *EligibilityNumbersResponse
	isSet bool
}

func (v NullableEligibilityNumbersResponse) Get() *EligibilityNumbersResponse {
	return v.value
}

func (v *NullableEligibilityNumbersResponse) Set(val *EligibilityNumbersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEligibilityNumbersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEligibilityNumbersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEligibilityNumbersResponse(val *EligibilityNumbersResponse) *NullableEligibilityNumbersResponse {
	return &NullableEligibilityNumbersResponse{value: val, isSet: true}
}

func (v NullableEligibilityNumbersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEligibilityNumbersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


