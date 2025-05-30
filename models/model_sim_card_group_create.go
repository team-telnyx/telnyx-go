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

// checks if the SIMCardGroupCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SIMCardGroupCreate{}

// SIMCardGroupCreate struct for SIMCardGroupCreate
type SIMCardGroupCreate struct {
	// A user friendly name for the SIM card group.
	Name string `json:"name"`
	DataLimit *SIMCardGroupDataLimit `json:"data_limit,omitempty"`
}

type _SIMCardGroupCreate SIMCardGroupCreate

// NewSIMCardGroupCreate instantiates a new SIMCardGroupCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSIMCardGroupCreate(name string) *SIMCardGroupCreate {
	this := SIMCardGroupCreate{}
	this.Name = name
	return &this
}

// NewSIMCardGroupCreateWithDefaults instantiates a new SIMCardGroupCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSIMCardGroupCreateWithDefaults() *SIMCardGroupCreate {
	this := SIMCardGroupCreate{}
	return &this
}

// GetName returns the Name field value
func (o *SIMCardGroupCreate) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SIMCardGroupCreate) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SIMCardGroupCreate) SetName(v string) {
	o.Name = v
}

// GetDataLimit returns the DataLimit field value if set, zero value otherwise.
func (o *SIMCardGroupCreate) GetDataLimit() SIMCardGroupDataLimit {
	if o == nil || IsNil(o.DataLimit) {
		var ret SIMCardGroupDataLimit
		return ret
	}
	return *o.DataLimit
}

// GetDataLimitOk returns a tuple with the DataLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardGroupCreate) GetDataLimitOk() (*SIMCardGroupDataLimit, bool) {
	if o == nil || IsNil(o.DataLimit) {
		return nil, false
	}
	return o.DataLimit, true
}

// HasDataLimit returns a boolean if a field has been set.
func (o *SIMCardGroupCreate) HasDataLimit() bool {
	if o != nil && !IsNil(o.DataLimit) {
		return true
	}

	return false
}

// SetDataLimit gets a reference to the given SIMCardGroupDataLimit and assigns it to the DataLimit field.
func (o *SIMCardGroupCreate) SetDataLimit(v SIMCardGroupDataLimit) {
	o.DataLimit = &v
}

func (o SIMCardGroupCreate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SIMCardGroupCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.DataLimit) {
		toSerialize["data_limit"] = o.DataLimit
	}
	return toSerialize, nil
}

func (o *SIMCardGroupCreate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
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

	varSIMCardGroupCreate := _SIMCardGroupCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSIMCardGroupCreate)

	if err != nil {
		return err
	}

	*o = SIMCardGroupCreate(varSIMCardGroupCreate)

	return err
}

type NullableSIMCardGroupCreate struct {
	value *SIMCardGroupCreate
	isSet bool
}

func (v NullableSIMCardGroupCreate) Get() *SIMCardGroupCreate {
	return v.value
}

func (v *NullableSIMCardGroupCreate) Set(val *SIMCardGroupCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableSIMCardGroupCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableSIMCardGroupCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSIMCardGroupCreate(val *SIMCardGroupCreate) *NullableSIMCardGroupCreate {
	return &NullableSIMCardGroupCreate{value: val, isSet: true}
}

func (v NullableSIMCardGroupCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSIMCardGroupCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


