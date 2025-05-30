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

// checks if the UnusedUserBundlesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnusedUserBundlesResponse{}

// UnusedUserBundlesResponse struct for UnusedUserBundlesResponse
type UnusedUserBundlesResponse struct {
	Data []UnusedUserBundlesSchema `json:"data"`
}

type _UnusedUserBundlesResponse UnusedUserBundlesResponse

// NewUnusedUserBundlesResponse instantiates a new UnusedUserBundlesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnusedUserBundlesResponse(data []UnusedUserBundlesSchema) *UnusedUserBundlesResponse {
	this := UnusedUserBundlesResponse{}
	this.Data = data
	return &this
}

// NewUnusedUserBundlesResponseWithDefaults instantiates a new UnusedUserBundlesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnusedUserBundlesResponseWithDefaults() *UnusedUserBundlesResponse {
	this := UnusedUserBundlesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UnusedUserBundlesResponse) GetData() []UnusedUserBundlesSchema {
	if o == nil {
		var ret []UnusedUserBundlesSchema
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UnusedUserBundlesResponse) GetDataOk() ([]UnusedUserBundlesSchema, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *UnusedUserBundlesResponse) SetData(v []UnusedUserBundlesSchema) {
	o.Data = v
}

func (o UnusedUserBundlesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnusedUserBundlesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *UnusedUserBundlesResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varUnusedUserBundlesResponse := _UnusedUserBundlesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUnusedUserBundlesResponse)

	if err != nil {
		return err
	}

	*o = UnusedUserBundlesResponse(varUnusedUserBundlesResponse)

	return err
}

type NullableUnusedUserBundlesResponse struct {
	value *UnusedUserBundlesResponse
	isSet bool
}

func (v NullableUnusedUserBundlesResponse) Get() *UnusedUserBundlesResponse {
	return v.value
}

func (v *NullableUnusedUserBundlesResponse) Set(val *UnusedUserBundlesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUnusedUserBundlesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUnusedUserBundlesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnusedUserBundlesResponse(val *UnusedUserBundlesResponse) *NullableUnusedUserBundlesResponse {
	return &NullableUnusedUserBundlesResponse{value: val, isSet: true}
}

func (v NullableUnusedUserBundlesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnusedUserBundlesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


