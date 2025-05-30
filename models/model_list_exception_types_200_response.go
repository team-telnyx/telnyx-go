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

// checks if the ListExceptionTypes200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListExceptionTypes200Response{}

// ListExceptionTypes200Response struct for ListExceptionTypes200Response
type ListExceptionTypes200Response struct {
	Data []PortingOrdersExceptionType `json:"data,omitempty"`
}

// NewListExceptionTypes200Response instantiates a new ListExceptionTypes200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListExceptionTypes200Response() *ListExceptionTypes200Response {
	this := ListExceptionTypes200Response{}
	return &this
}

// NewListExceptionTypes200ResponseWithDefaults instantiates a new ListExceptionTypes200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListExceptionTypes200ResponseWithDefaults() *ListExceptionTypes200Response {
	this := ListExceptionTypes200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListExceptionTypes200Response) GetData() []PortingOrdersExceptionType {
	if o == nil || IsNil(o.Data) {
		var ret []PortingOrdersExceptionType
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListExceptionTypes200Response) GetDataOk() ([]PortingOrdersExceptionType, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListExceptionTypes200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PortingOrdersExceptionType and assigns it to the Data field.
func (o *ListExceptionTypes200Response) SetData(v []PortingOrdersExceptionType) {
	o.Data = v
}

func (o ListExceptionTypes200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListExceptionTypes200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListExceptionTypes200Response struct {
	value *ListExceptionTypes200Response
	isSet bool
}

func (v NullableListExceptionTypes200Response) Get() *ListExceptionTypes200Response {
	return v.value
}

func (v *NullableListExceptionTypes200Response) Set(val *ListExceptionTypes200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListExceptionTypes200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListExceptionTypes200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListExceptionTypes200Response(val *ListExceptionTypes200Response) *NullableListExceptionTypes200Response {
	return &NullableListExceptionTypes200Response{value: val, isSet: true}
}

func (v NullableListExceptionTypes200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListExceptionTypes200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


