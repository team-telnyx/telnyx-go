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

// checks if the ListUploadsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUploadsResponse{}

// ListUploadsResponse struct for ListUploadsResponse
type ListUploadsResponse struct {
	Data []Upload `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
}

// NewListUploadsResponse instantiates a new ListUploadsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUploadsResponse() *ListUploadsResponse {
	this := ListUploadsResponse{}
	return &this
}

// NewListUploadsResponseWithDefaults instantiates a new ListUploadsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUploadsResponseWithDefaults() *ListUploadsResponse {
	this := ListUploadsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListUploadsResponse) GetData() []Upload {
	if o == nil || IsNil(o.Data) {
		var ret []Upload
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUploadsResponse) GetDataOk() ([]Upload, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListUploadsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []Upload and assigns it to the Data field.
func (o *ListUploadsResponse) SetData(v []Upload) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListUploadsResponse) GetMeta() PaginationMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUploadsResponse) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListUploadsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *ListUploadsResponse) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

func (o ListUploadsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUploadsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableListUploadsResponse struct {
	value *ListUploadsResponse
	isSet bool
}

func (v NullableListUploadsResponse) Get() *ListUploadsResponse {
	return v.value
}

func (v *NullableListUploadsResponse) Set(val *ListUploadsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListUploadsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListUploadsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUploadsResponse(val *ListUploadsResponse) *NullableListUploadsResponse {
	return &NullableListUploadsResponse{value: val, isSet: true}
}

func (v NullableListUploadsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUploadsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


