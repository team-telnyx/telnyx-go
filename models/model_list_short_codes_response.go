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

// checks if the ListShortCodesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListShortCodesResponse{}

// ListShortCodesResponse struct for ListShortCodesResponse
type ListShortCodesResponse struct {
	Data []ShortCode `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
}

// NewListShortCodesResponse instantiates a new ListShortCodesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListShortCodesResponse() *ListShortCodesResponse {
	this := ListShortCodesResponse{}
	return &this
}

// NewListShortCodesResponseWithDefaults instantiates a new ListShortCodesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListShortCodesResponseWithDefaults() *ListShortCodesResponse {
	this := ListShortCodesResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListShortCodesResponse) GetData() []ShortCode {
	if o == nil || IsNil(o.Data) {
		var ret []ShortCode
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListShortCodesResponse) GetDataOk() ([]ShortCode, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListShortCodesResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ShortCode and assigns it to the Data field.
func (o *ListShortCodesResponse) SetData(v []ShortCode) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListShortCodesResponse) GetMeta() PaginationMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListShortCodesResponse) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListShortCodesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *ListShortCodesResponse) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

func (o ListShortCodesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListShortCodesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableListShortCodesResponse struct {
	value *ListShortCodesResponse
	isSet bool
}

func (v NullableListShortCodesResponse) Get() *ListShortCodesResponse {
	return v.value
}

func (v *NullableListShortCodesResponse) Set(val *ListShortCodesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListShortCodesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListShortCodesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListShortCodesResponse(val *ListShortCodesResponse) *NullableListShortCodesResponse {
	return &NullableListShortCodesResponse{value: val, isSet: true}
}

func (v NullableListShortCodesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListShortCodesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


