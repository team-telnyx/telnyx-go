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

// checks if the ListVerifiedNumbersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListVerifiedNumbersResponse{}

// ListVerifiedNumbersResponse A paginated list of Verified Numbers
type ListVerifiedNumbersResponse struct {
	Data []VerifiedNumberResponse `json:"data"`
	Meta Meta `json:"meta"`
}

type _ListVerifiedNumbersResponse ListVerifiedNumbersResponse

// NewListVerifiedNumbersResponse instantiates a new ListVerifiedNumbersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVerifiedNumbersResponse(data []VerifiedNumberResponse, meta Meta) *ListVerifiedNumbersResponse {
	this := ListVerifiedNumbersResponse{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewListVerifiedNumbersResponseWithDefaults instantiates a new ListVerifiedNumbersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVerifiedNumbersResponseWithDefaults() *ListVerifiedNumbersResponse {
	this := ListVerifiedNumbersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListVerifiedNumbersResponse) GetData() []VerifiedNumberResponse {
	if o == nil {
		var ret []VerifiedNumberResponse
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListVerifiedNumbersResponse) GetDataOk() ([]VerifiedNumberResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListVerifiedNumbersResponse) SetData(v []VerifiedNumberResponse) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *ListVerifiedNumbersResponse) GetMeta() Meta {
	if o == nil {
		var ret Meta
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *ListVerifiedNumbersResponse) GetMetaOk() (*Meta, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *ListVerifiedNumbersResponse) SetMeta(v Meta) {
	o.Meta = v
}

func (o ListVerifiedNumbersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListVerifiedNumbersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *ListVerifiedNumbersResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"meta",
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

	varListVerifiedNumbersResponse := _ListVerifiedNumbersResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varListVerifiedNumbersResponse)

	if err != nil {
		return err
	}

	*o = ListVerifiedNumbersResponse(varListVerifiedNumbersResponse)

	return err
}

type NullableListVerifiedNumbersResponse struct {
	value *ListVerifiedNumbersResponse
	isSet bool
}

func (v NullableListVerifiedNumbersResponse) Get() *ListVerifiedNumbersResponse {
	return v.value
}

func (v *NullableListVerifiedNumbersResponse) Set(val *ListVerifiedNumbersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListVerifiedNumbersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListVerifiedNumbersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVerifiedNumbersResponse(val *ListVerifiedNumbersResponse) *NullableListVerifiedNumbersResponse {
	return &NullableListVerifiedNumbersResponse{value: val, isSet: true}
}

func (v NullableListVerifiedNumbersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVerifiedNumbersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


