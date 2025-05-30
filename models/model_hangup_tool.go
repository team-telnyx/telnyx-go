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

// checks if the HangupTool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HangupTool{}

// HangupTool struct for HangupTool
type HangupTool struct {
	Type string `json:"type"`
	Hangup HangupToolParams `json:"hangup"`
}

type _HangupTool HangupTool

// NewHangupTool instantiates a new HangupTool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHangupTool(type_ string, hangup HangupToolParams) *HangupTool {
	this := HangupTool{}
	this.Type = type_
	this.Hangup = hangup
	return &this
}

// NewHangupToolWithDefaults instantiates a new HangupTool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHangupToolWithDefaults() *HangupTool {
	this := HangupTool{}
	return &this
}

// GetType returns the Type field value
func (o *HangupTool) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *HangupTool) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *HangupTool) SetType(v string) {
	o.Type = v
}

// GetHangup returns the Hangup field value
func (o *HangupTool) GetHangup() HangupToolParams {
	if o == nil {
		var ret HangupToolParams
		return ret
	}

	return o.Hangup
}

// GetHangupOk returns a tuple with the Hangup field value
// and a boolean to check if the value has been set.
func (o *HangupTool) GetHangupOk() (*HangupToolParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hangup, true
}

// SetHangup sets field value
func (o *HangupTool) SetHangup(v HangupToolParams) {
	o.Hangup = v
}

func (o HangupTool) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HangupTool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["hangup"] = o.Hangup
	return toSerialize, nil
}

func (o *HangupTool) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"hangup",
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

	varHangupTool := _HangupTool{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varHangupTool)

	if err != nil {
		return err
	}

	*o = HangupTool(varHangupTool)

	return err
}

type NullableHangupTool struct {
	value *HangupTool
	isSet bool
}

func (v NullableHangupTool) Get() *HangupTool {
	return v.value
}

func (v *NullableHangupTool) Set(val *HangupTool) {
	v.value = val
	v.isSet = true
}

func (v NullableHangupTool) IsSet() bool {
	return v.isSet
}

func (v *NullableHangupTool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHangupTool(val *HangupTool) *NullableHangupTool {
	return &NullableHangupTool{value: val, isSet: true}
}

func (v NullableHangupTool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHangupTool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


