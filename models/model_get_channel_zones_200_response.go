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

// checks if the GetChannelZones200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetChannelZones200Response{}

// GetChannelZones200Response struct for GetChannelZones200Response
type GetChannelZones200Response struct {
	Data []GcbChannelZone `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
}

// NewGetChannelZones200Response instantiates a new GetChannelZones200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetChannelZones200Response() *GetChannelZones200Response {
	this := GetChannelZones200Response{}
	return &this
}

// NewGetChannelZones200ResponseWithDefaults instantiates a new GetChannelZones200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetChannelZones200ResponseWithDefaults() *GetChannelZones200Response {
	this := GetChannelZones200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetChannelZones200Response) GetData() []GcbChannelZone {
	if o == nil || IsNil(o.Data) {
		var ret []GcbChannelZone
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChannelZones200Response) GetDataOk() ([]GcbChannelZone, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetChannelZones200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GcbChannelZone and assigns it to the Data field.
func (o *GetChannelZones200Response) SetData(v []GcbChannelZone) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetChannelZones200Response) GetMeta() PaginationMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetChannelZones200Response) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetChannelZones200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *GetChannelZones200Response) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

func (o GetChannelZones200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetChannelZones200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetChannelZones200Response struct {
	value *GetChannelZones200Response
	isSet bool
}

func (v NullableGetChannelZones200Response) Get() *GetChannelZones200Response {
	return v.value
}

func (v *NullableGetChannelZones200Response) Set(val *GetChannelZones200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetChannelZones200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetChannelZones200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetChannelZones200Response(val *GetChannelZones200Response) *NullableGetChannelZones200Response {
	return &NullableGetChannelZones200Response{value: val, isSet: true}
}

func (v NullableGetChannelZones200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetChannelZones200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


