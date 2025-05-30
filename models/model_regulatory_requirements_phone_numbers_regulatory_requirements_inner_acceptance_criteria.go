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

// checks if the RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria{}

// RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria struct for RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria
type RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria struct {
	LocalityLimit *string `json:"locality_limit,omitempty"`
	FieldValue *string `json:"field_value,omitempty"`
	FieldType *string `json:"field_type,omitempty"`
}

// NewRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria instantiates a new RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria() *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria {
	this := RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria{}
	return &this
}

// NewRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteriaWithDefaults instantiates a new RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteriaWithDefaults() *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria {
	this := RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria{}
	return &this
}

// GetLocalityLimit returns the LocalityLimit field value if set, zero value otherwise.
func (o *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) GetLocalityLimit() string {
	if o == nil || IsNil(o.LocalityLimit) {
		var ret string
		return ret
	}
	return *o.LocalityLimit
}

// GetLocalityLimitOk returns a tuple with the LocalityLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) GetLocalityLimitOk() (*string, bool) {
	if o == nil || IsNil(o.LocalityLimit) {
		return nil, false
	}
	return o.LocalityLimit, true
}

// HasLocalityLimit returns a boolean if a field has been set.
func (o *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) HasLocalityLimit() bool {
	if o != nil && !IsNil(o.LocalityLimit) {
		return true
	}

	return false
}

// SetLocalityLimit gets a reference to the given string and assigns it to the LocalityLimit field.
func (o *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) SetLocalityLimit(v string) {
	o.LocalityLimit = &v
}

// GetFieldValue returns the FieldValue field value if set, zero value otherwise.
func (o *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) GetFieldValue() string {
	if o == nil || IsNil(o.FieldValue) {
		var ret string
		return ret
	}
	return *o.FieldValue
}

// GetFieldValueOk returns a tuple with the FieldValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) GetFieldValueOk() (*string, bool) {
	if o == nil || IsNil(o.FieldValue) {
		return nil, false
	}
	return o.FieldValue, true
}

// HasFieldValue returns a boolean if a field has been set.
func (o *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) HasFieldValue() bool {
	if o != nil && !IsNil(o.FieldValue) {
		return true
	}

	return false
}

// SetFieldValue gets a reference to the given string and assigns it to the FieldValue field.
func (o *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) SetFieldValue(v string) {
	o.FieldValue = &v
}

// GetFieldType returns the FieldType field value if set, zero value otherwise.
func (o *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) GetFieldType() string {
	if o == nil || IsNil(o.FieldType) {
		var ret string
		return ret
	}
	return *o.FieldType
}

// GetFieldTypeOk returns a tuple with the FieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) GetFieldTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FieldType) {
		return nil, false
	}
	return o.FieldType, true
}

// HasFieldType returns a boolean if a field has been set.
func (o *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) HasFieldType() bool {
	if o != nil && !IsNil(o.FieldType) {
		return true
	}

	return false
}

// SetFieldType gets a reference to the given string and assigns it to the FieldType field.
func (o *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) SetFieldType(v string) {
	o.FieldType = &v
}

func (o RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LocalityLimit) {
		toSerialize["locality_limit"] = o.LocalityLimit
	}
	if !IsNil(o.FieldValue) {
		toSerialize["field_value"] = o.FieldValue
	}
	if !IsNil(o.FieldType) {
		toSerialize["field_type"] = o.FieldType
	}
	return toSerialize, nil
}

type NullableRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria struct {
	value *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria
	isSet bool
}

func (v NullableRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) Get() *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria {
	return v.value
}

func (v *NullableRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) Set(val *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria(val *RegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) *NullableRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria {
	return &NullableRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria{value: val, isSet: true}
}

func (v NullableRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegulatoryRequirementsPhoneNumbersRegulatoryRequirementsInnerAcceptanceCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


