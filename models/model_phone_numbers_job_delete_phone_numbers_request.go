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
	"bytes"
	"fmt"
)

// checks if the PhoneNumbersJobDeletePhoneNumbersRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PhoneNumbersJobDeletePhoneNumbersRequest{}

// PhoneNumbersJobDeletePhoneNumbersRequest struct for PhoneNumbersJobDeletePhoneNumbersRequest
type PhoneNumbersJobDeletePhoneNumbersRequest struct {
	PhoneNumbers []string `json:"phone_numbers"`
}

type _PhoneNumbersJobDeletePhoneNumbersRequest PhoneNumbersJobDeletePhoneNumbersRequest

// NewPhoneNumbersJobDeletePhoneNumbersRequest instantiates a new PhoneNumbersJobDeletePhoneNumbersRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPhoneNumbersJobDeletePhoneNumbersRequest(phoneNumbers []string) *PhoneNumbersJobDeletePhoneNumbersRequest {
	this := PhoneNumbersJobDeletePhoneNumbersRequest{}
	this.PhoneNumbers = phoneNumbers
	return &this
}

// NewPhoneNumbersJobDeletePhoneNumbersRequestWithDefaults instantiates a new PhoneNumbersJobDeletePhoneNumbersRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPhoneNumbersJobDeletePhoneNumbersRequestWithDefaults() *PhoneNumbersJobDeletePhoneNumbersRequest {
	this := PhoneNumbersJobDeletePhoneNumbersRequest{}
	return &this
}

// GetPhoneNumbers returns the PhoneNumbers field value
func (o *PhoneNumbersJobDeletePhoneNumbersRequest) GetPhoneNumbers() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value
// and a boolean to check if the value has been set.
func (o *PhoneNumbersJobDeletePhoneNumbersRequest) GetPhoneNumbersOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// SetPhoneNumbers sets field value
func (o *PhoneNumbersJobDeletePhoneNumbersRequest) SetPhoneNumbers(v []string) {
	o.PhoneNumbers = v
}

func (o PhoneNumbersJobDeletePhoneNumbersRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PhoneNumbersJobDeletePhoneNumbersRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phone_numbers"] = o.PhoneNumbers
	return toSerialize, nil
}

func (o *PhoneNumbersJobDeletePhoneNumbersRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phone_numbers",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPhoneNumbersJobDeletePhoneNumbersRequest := _PhoneNumbersJobDeletePhoneNumbersRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPhoneNumbersJobDeletePhoneNumbersRequest)

	if err != nil {
		return err
	}

	*o = PhoneNumbersJobDeletePhoneNumbersRequest(varPhoneNumbersJobDeletePhoneNumbersRequest)

	return err
}

type NullablePhoneNumbersJobDeletePhoneNumbersRequest struct {
	value *PhoneNumbersJobDeletePhoneNumbersRequest
	isSet bool
}

func (v NullablePhoneNumbersJobDeletePhoneNumbersRequest) Get() *PhoneNumbersJobDeletePhoneNumbersRequest {
	return v.value
}

func (v *NullablePhoneNumbersJobDeletePhoneNumbersRequest) Set(val *PhoneNumbersJobDeletePhoneNumbersRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePhoneNumbersJobDeletePhoneNumbersRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePhoneNumbersJobDeletePhoneNumbersRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhoneNumbersJobDeletePhoneNumbersRequest(val *PhoneNumbersJobDeletePhoneNumbersRequest) *NullablePhoneNumbersJobDeletePhoneNumbersRequest {
	return &NullablePhoneNumbersJobDeletePhoneNumbersRequest{value: val, isSet: true}
}

func (v NullablePhoneNumbersJobDeletePhoneNumbersRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhoneNumbersJobDeletePhoneNumbersRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


