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

// checks if the ListLoaConfigurations200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListLoaConfigurations200Response{}

// ListLoaConfigurations200Response struct for ListLoaConfigurations200Response
type ListLoaConfigurations200Response struct {
	Data []PortingLOAConfiguration `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
}

// NewListLoaConfigurations200Response instantiates a new ListLoaConfigurations200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListLoaConfigurations200Response() *ListLoaConfigurations200Response {
	this := ListLoaConfigurations200Response{}
	return &this
}

// NewListLoaConfigurations200ResponseWithDefaults instantiates a new ListLoaConfigurations200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListLoaConfigurations200ResponseWithDefaults() *ListLoaConfigurations200Response {
	this := ListLoaConfigurations200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListLoaConfigurations200Response) GetData() []PortingLOAConfiguration {
	if o == nil || IsNil(o.Data) {
		var ret []PortingLOAConfiguration
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLoaConfigurations200Response) GetDataOk() ([]PortingLOAConfiguration, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListLoaConfigurations200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PortingLOAConfiguration and assigns it to the Data field.
func (o *ListLoaConfigurations200Response) SetData(v []PortingLOAConfiguration) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListLoaConfigurations200Response) GetMeta() PaginationMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListLoaConfigurations200Response) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListLoaConfigurations200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *ListLoaConfigurations200Response) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

func (o ListLoaConfigurations200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListLoaConfigurations200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableListLoaConfigurations200Response struct {
	value *ListLoaConfigurations200Response
	isSet bool
}

func (v NullableListLoaConfigurations200Response) Get() *ListLoaConfigurations200Response {
	return v.value
}

func (v *NullableListLoaConfigurations200Response) Set(val *ListLoaConfigurations200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListLoaConfigurations200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListLoaConfigurations200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListLoaConfigurations200Response(val *ListLoaConfigurations200Response) *NullableListLoaConfigurations200Response {
	return &NullableListLoaConfigurations200Response{value: val, isSet: true}
}

func (v NullableListLoaConfigurations200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListLoaConfigurations200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


