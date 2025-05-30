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

// checks if the CreateNumberOrderRequestPhoneNumbersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNumberOrderRequestPhoneNumbersInner{}

// CreateNumberOrderRequestPhoneNumbersInner struct for CreateNumberOrderRequestPhoneNumbersInner
type CreateNumberOrderRequestPhoneNumbersInner struct {
	// e164_phone_number
	PhoneNumber string `json:"phone_number"`
	// ID of requirement group to use to satisfy number requirements
	RequirementGroupId *string `json:"requirement_group_id,omitempty"`
	// ID of bundle to associate the number to
	BundleId *string `json:"bundle_id,omitempty"`
}

type _CreateNumberOrderRequestPhoneNumbersInner CreateNumberOrderRequestPhoneNumbersInner

// NewCreateNumberOrderRequestPhoneNumbersInner instantiates a new CreateNumberOrderRequestPhoneNumbersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNumberOrderRequestPhoneNumbersInner(phoneNumber string) *CreateNumberOrderRequestPhoneNumbersInner {
	this := CreateNumberOrderRequestPhoneNumbersInner{}
	this.PhoneNumber = phoneNumber
	return &this
}

// NewCreateNumberOrderRequestPhoneNumbersInnerWithDefaults instantiates a new CreateNumberOrderRequestPhoneNumbersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNumberOrderRequestPhoneNumbersInnerWithDefaults() *CreateNumberOrderRequestPhoneNumbersInner {
	this := CreateNumberOrderRequestPhoneNumbersInner{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *CreateNumberOrderRequestPhoneNumbersInner) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *CreateNumberOrderRequestPhoneNumbersInner) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *CreateNumberOrderRequestPhoneNumbersInner) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetRequirementGroupId returns the RequirementGroupId field value if set, zero value otherwise.
func (o *CreateNumberOrderRequestPhoneNumbersInner) GetRequirementGroupId() string {
	if o == nil || IsNil(o.RequirementGroupId) {
		var ret string
		return ret
	}
	return *o.RequirementGroupId
}

// GetRequirementGroupIdOk returns a tuple with the RequirementGroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberOrderRequestPhoneNumbersInner) GetRequirementGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequirementGroupId) {
		return nil, false
	}
	return o.RequirementGroupId, true
}

// HasRequirementGroupId returns a boolean if a field has been set.
func (o *CreateNumberOrderRequestPhoneNumbersInner) HasRequirementGroupId() bool {
	if o != nil && !IsNil(o.RequirementGroupId) {
		return true
	}

	return false
}

// SetRequirementGroupId gets a reference to the given string and assigns it to the RequirementGroupId field.
func (o *CreateNumberOrderRequestPhoneNumbersInner) SetRequirementGroupId(v string) {
	o.RequirementGroupId = &v
}

// GetBundleId returns the BundleId field value if set, zero value otherwise.
func (o *CreateNumberOrderRequestPhoneNumbersInner) GetBundleId() string {
	if o == nil || IsNil(o.BundleId) {
		var ret string
		return ret
	}
	return *o.BundleId
}

// GetBundleIdOk returns a tuple with the BundleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNumberOrderRequestPhoneNumbersInner) GetBundleIdOk() (*string, bool) {
	if o == nil || IsNil(o.BundleId) {
		return nil, false
	}
	return o.BundleId, true
}

// HasBundleId returns a boolean if a field has been set.
func (o *CreateNumberOrderRequestPhoneNumbersInner) HasBundleId() bool {
	if o != nil && !IsNil(o.BundleId) {
		return true
	}

	return false
}

// SetBundleId gets a reference to the given string and assigns it to the BundleId field.
func (o *CreateNumberOrderRequestPhoneNumbersInner) SetBundleId(v string) {
	o.BundleId = &v
}

func (o CreateNumberOrderRequestPhoneNumbersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNumberOrderRequestPhoneNumbersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phone_number"] = o.PhoneNumber
	if !IsNil(o.RequirementGroupId) {
		toSerialize["requirement_group_id"] = o.RequirementGroupId
	}
	if !IsNil(o.BundleId) {
		toSerialize["bundle_id"] = o.BundleId
	}
	return toSerialize, nil
}

func (o *CreateNumberOrderRequestPhoneNumbersInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phone_number",
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

	varCreateNumberOrderRequestPhoneNumbersInner := _CreateNumberOrderRequestPhoneNumbersInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateNumberOrderRequestPhoneNumbersInner)

	if err != nil {
		return err
	}

	*o = CreateNumberOrderRequestPhoneNumbersInner(varCreateNumberOrderRequestPhoneNumbersInner)

	return err
}

type NullableCreateNumberOrderRequestPhoneNumbersInner struct {
	value *CreateNumberOrderRequestPhoneNumbersInner
	isSet bool
}

func (v NullableCreateNumberOrderRequestPhoneNumbersInner) Get() *CreateNumberOrderRequestPhoneNumbersInner {
	return v.value
}

func (v *NullableCreateNumberOrderRequestPhoneNumbersInner) Set(val *CreateNumberOrderRequestPhoneNumbersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateNumberOrderRequestPhoneNumbersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateNumberOrderRequestPhoneNumbersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateNumberOrderRequestPhoneNumbersInner(val *CreateNumberOrderRequestPhoneNumbersInner) *NullableCreateNumberOrderRequestPhoneNumbersInner {
	return &NullableCreateNumberOrderRequestPhoneNumbersInner{value: val, isSet: true}
}

func (v NullableCreateNumberOrderRequestPhoneNumbersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateNumberOrderRequestPhoneNumbersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


