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

// checks if the CreatedUserBundlesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreatedUserBundlesResponse{}

// CreatedUserBundlesResponse struct for CreatedUserBundlesResponse
type CreatedUserBundlesResponse struct {
	Data []UserBundle `json:"data"`
}

type _CreatedUserBundlesResponse CreatedUserBundlesResponse

// NewCreatedUserBundlesResponse instantiates a new CreatedUserBundlesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreatedUserBundlesResponse(data []UserBundle) *CreatedUserBundlesResponse {
	this := CreatedUserBundlesResponse{}
	this.Data = data
	return &this
}

// NewCreatedUserBundlesResponseWithDefaults instantiates a new CreatedUserBundlesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreatedUserBundlesResponseWithDefaults() *CreatedUserBundlesResponse {
	this := CreatedUserBundlesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CreatedUserBundlesResponse) GetData() []UserBundle {
	if o == nil {
		var ret []UserBundle
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CreatedUserBundlesResponse) GetDataOk() ([]UserBundle, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CreatedUserBundlesResponse) SetData(v []UserBundle) {
	o.Data = v
}

func (o CreatedUserBundlesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreatedUserBundlesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *CreatedUserBundlesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varCreatedUserBundlesResponse := _CreatedUserBundlesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreatedUserBundlesResponse)

	if err != nil {
		return err
	}

	*o = CreatedUserBundlesResponse(varCreatedUserBundlesResponse)

	return err
}

type NullableCreatedUserBundlesResponse struct {
	value *CreatedUserBundlesResponse
	isSet bool
}

func (v NullableCreatedUserBundlesResponse) Get() *CreatedUserBundlesResponse {
	return v.value
}

func (v *NullableCreatedUserBundlesResponse) Set(val *CreatedUserBundlesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreatedUserBundlesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreatedUserBundlesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreatedUserBundlesResponse(val *CreatedUserBundlesResponse) *NullableCreatedUserBundlesResponse {
	return &NullableCreatedUserBundlesResponse{value: val, isSet: true}
}

func (v NullableCreatedUserBundlesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreatedUserBundlesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


