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

// checks if the PostNumbersFeatures200ResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostNumbersFeatures200ResponseDataInner{}

// PostNumbersFeatures200ResponseDataInner struct for PostNumbersFeatures200ResponseDataInner
type PostNumbersFeatures200ResponseDataInner struct {
	PhoneNumber string `json:"phone_number"`
	Features []string `json:"features"`
}

type _PostNumbersFeatures200ResponseDataInner PostNumbersFeatures200ResponseDataInner

// NewPostNumbersFeatures200ResponseDataInner instantiates a new PostNumbersFeatures200ResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostNumbersFeatures200ResponseDataInner(phoneNumber string, features []string) *PostNumbersFeatures200ResponseDataInner {
	this := PostNumbersFeatures200ResponseDataInner{}
	this.PhoneNumber = phoneNumber
	this.Features = features
	return &this
}

// NewPostNumbersFeatures200ResponseDataInnerWithDefaults instantiates a new PostNumbersFeatures200ResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostNumbersFeatures200ResponseDataInnerWithDefaults() *PostNumbersFeatures200ResponseDataInner {
	this := PostNumbersFeatures200ResponseDataInner{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *PostNumbersFeatures200ResponseDataInner) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *PostNumbersFeatures200ResponseDataInner) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *PostNumbersFeatures200ResponseDataInner) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetFeatures returns the Features field value
func (o *PostNumbersFeatures200ResponseDataInner) GetFeatures() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value
// and a boolean to check if the value has been set.
func (o *PostNumbersFeatures200ResponseDataInner) GetFeaturesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Features, true
}

// SetFeatures sets field value
func (o *PostNumbersFeatures200ResponseDataInner) SetFeatures(v []string) {
	o.Features = v
}

func (o PostNumbersFeatures200ResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostNumbersFeatures200ResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phone_number"] = o.PhoneNumber
	toSerialize["features"] = o.Features
	return toSerialize, nil
}

func (o *PostNumbersFeatures200ResponseDataInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phone_number",
		"features",
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

	varPostNumbersFeatures200ResponseDataInner := _PostNumbersFeatures200ResponseDataInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostNumbersFeatures200ResponseDataInner)

	if err != nil {
		return err
	}

	*o = PostNumbersFeatures200ResponseDataInner(varPostNumbersFeatures200ResponseDataInner)

	return err
}

type NullablePostNumbersFeatures200ResponseDataInner struct {
	value *PostNumbersFeatures200ResponseDataInner
	isSet bool
}

func (v NullablePostNumbersFeatures200ResponseDataInner) Get() *PostNumbersFeatures200ResponseDataInner {
	return v.value
}

func (v *NullablePostNumbersFeatures200ResponseDataInner) Set(val *PostNumbersFeatures200ResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePostNumbersFeatures200ResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePostNumbersFeatures200ResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostNumbersFeatures200ResponseDataInner(val *PostNumbersFeatures200ResponseDataInner) *NullablePostNumbersFeatures200ResponseDataInner {
	return &NullablePostNumbersFeatures200ResponseDataInner{value: val, isSet: true}
}

func (v NullablePostNumbersFeatures200ResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostNumbersFeatures200ResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


