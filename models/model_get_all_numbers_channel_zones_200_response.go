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

// checks if the GetAllNumbersChannelZones200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllNumbersChannelZones200Response{}

// GetAllNumbersChannelZones200Response struct for GetAllNumbersChannelZones200Response
type GetAllNumbersChannelZones200Response struct {
	Data []GetAllNumbersChannelZones200ResponseDataInner `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
}

// NewGetAllNumbersChannelZones200Response instantiates a new GetAllNumbersChannelZones200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllNumbersChannelZones200Response() *GetAllNumbersChannelZones200Response {
	this := GetAllNumbersChannelZones200Response{}
	return &this
}

// NewGetAllNumbersChannelZones200ResponseWithDefaults instantiates a new GetAllNumbersChannelZones200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllNumbersChannelZones200ResponseWithDefaults() *GetAllNumbersChannelZones200Response {
	this := GetAllNumbersChannelZones200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAllNumbersChannelZones200Response) GetData() []GetAllNumbersChannelZones200ResponseDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GetAllNumbersChannelZones200ResponseDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllNumbersChannelZones200Response) GetDataOk() ([]GetAllNumbersChannelZones200ResponseDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAllNumbersChannelZones200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GetAllNumbersChannelZones200ResponseDataInner and assigns it to the Data field.
func (o *GetAllNumbersChannelZones200Response) SetData(v []GetAllNumbersChannelZones200ResponseDataInner) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetAllNumbersChannelZones200Response) GetMeta() PaginationMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllNumbersChannelZones200Response) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetAllNumbersChannelZones200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *GetAllNumbersChannelZones200Response) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

func (o GetAllNumbersChannelZones200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllNumbersChannelZones200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetAllNumbersChannelZones200Response struct {
	value *GetAllNumbersChannelZones200Response
	isSet bool
}

func (v NullableGetAllNumbersChannelZones200Response) Get() *GetAllNumbersChannelZones200Response {
	return v.value
}

func (v *NullableGetAllNumbersChannelZones200Response) Set(val *GetAllNumbersChannelZones200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllNumbersChannelZones200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllNumbersChannelZones200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllNumbersChannelZones200Response(val *GetAllNumbersChannelZones200Response) *NullableGetAllNumbersChannelZones200Response {
	return &NullableGetAllNumbersChannelZones200Response{value: val, isSet: true}
}

func (v NullableGetAllNumbersChannelZones200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllNumbersChannelZones200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


