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

// checks if the ValidationCodesPhoneNumbersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationCodesPhoneNumbersInner{}

// ValidationCodesPhoneNumbersInner struct for ValidationCodesPhoneNumbersInner
type ValidationCodesPhoneNumbersInner struct {
	PhoneNumber string `json:"phone_number"`
	Status string `json:"status"`
}

type _ValidationCodesPhoneNumbersInner ValidationCodesPhoneNumbersInner

// NewValidationCodesPhoneNumbersInner instantiates a new ValidationCodesPhoneNumbersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationCodesPhoneNumbersInner(phoneNumber string, status string) *ValidationCodesPhoneNumbersInner {
	this := ValidationCodesPhoneNumbersInner{}
	this.PhoneNumber = phoneNumber
	this.Status = status
	return &this
}

// NewValidationCodesPhoneNumbersInnerWithDefaults instantiates a new ValidationCodesPhoneNumbersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationCodesPhoneNumbersInnerWithDefaults() *ValidationCodesPhoneNumbersInner {
	this := ValidationCodesPhoneNumbersInner{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *ValidationCodesPhoneNumbersInner) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *ValidationCodesPhoneNumbersInner) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *ValidationCodesPhoneNumbersInner) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetStatus returns the Status field value
func (o *ValidationCodesPhoneNumbersInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ValidationCodesPhoneNumbersInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ValidationCodesPhoneNumbersInner) SetStatus(v string) {
	o.Status = v
}

func (o ValidationCodesPhoneNumbersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationCodesPhoneNumbersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phone_number"] = o.PhoneNumber
	toSerialize["status"] = o.Status
	return toSerialize, nil
}

func (o *ValidationCodesPhoneNumbersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phone_number",
		"status",
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

	varValidationCodesPhoneNumbersInner := _ValidationCodesPhoneNumbersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varValidationCodesPhoneNumbersInner)

	if err != nil {
		return err
	}

	*o = ValidationCodesPhoneNumbersInner(varValidationCodesPhoneNumbersInner)

	return err
}

type NullableValidationCodesPhoneNumbersInner struct {
	value *ValidationCodesPhoneNumbersInner
	isSet bool
}

func (v NullableValidationCodesPhoneNumbersInner) Get() *ValidationCodesPhoneNumbersInner {
	return v.value
}

func (v *NullableValidationCodesPhoneNumbersInner) Set(val *ValidationCodesPhoneNumbersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationCodesPhoneNumbersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationCodesPhoneNumbersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationCodesPhoneNumbersInner(val *ValidationCodesPhoneNumbersInner) *NullableValidationCodesPhoneNumbersInner {
	return &NullableValidationCodesPhoneNumbersInner{value: val, isSet: true}
}

func (v NullableValidationCodesPhoneNumbersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationCodesPhoneNumbersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


