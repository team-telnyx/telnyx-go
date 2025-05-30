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

// checks if the GetAllTelephonyCredentialResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAllTelephonyCredentialResponse{}

// GetAllTelephonyCredentialResponse struct for GetAllTelephonyCredentialResponse
type GetAllTelephonyCredentialResponse struct {
	Data []TelephonyCredential `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
}

// NewGetAllTelephonyCredentialResponse instantiates a new GetAllTelephonyCredentialResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAllTelephonyCredentialResponse() *GetAllTelephonyCredentialResponse {
	this := GetAllTelephonyCredentialResponse{}
	return &this
}

// NewGetAllTelephonyCredentialResponseWithDefaults instantiates a new GetAllTelephonyCredentialResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAllTelephonyCredentialResponseWithDefaults() *GetAllTelephonyCredentialResponse {
	this := GetAllTelephonyCredentialResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetAllTelephonyCredentialResponse) GetData() []TelephonyCredential {
	if o == nil || IsNil(o.Data) {
		var ret []TelephonyCredential
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllTelephonyCredentialResponse) GetDataOk() ([]TelephonyCredential, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetAllTelephonyCredentialResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []TelephonyCredential and assigns it to the Data field.
func (o *GetAllTelephonyCredentialResponse) SetData(v []TelephonyCredential) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetAllTelephonyCredentialResponse) GetMeta() PaginationMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAllTelephonyCredentialResponse) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetAllTelephonyCredentialResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *GetAllTelephonyCredentialResponse) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

func (o GetAllTelephonyCredentialResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAllTelephonyCredentialResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetAllTelephonyCredentialResponse struct {
	value *GetAllTelephonyCredentialResponse
	isSet bool
}

func (v NullableGetAllTelephonyCredentialResponse) Get() *GetAllTelephonyCredentialResponse {
	return v.value
}

func (v *NullableGetAllTelephonyCredentialResponse) Set(val *GetAllTelephonyCredentialResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAllTelephonyCredentialResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAllTelephonyCredentialResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAllTelephonyCredentialResponse(val *GetAllTelephonyCredentialResponse) *NullableGetAllTelephonyCredentialResponse {
	return &NullableGetAllTelephonyCredentialResponse{value: val, isSet: true}
}

func (v NullableGetAllTelephonyCredentialResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAllTelephonyCredentialResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


