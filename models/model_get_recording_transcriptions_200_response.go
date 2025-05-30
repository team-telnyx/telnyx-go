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

// checks if the GetRecordingTranscriptions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRecordingTranscriptions200Response{}

// GetRecordingTranscriptions200Response struct for GetRecordingTranscriptions200Response
type GetRecordingTranscriptions200Response struct {
	Data []RecordingTranscription `json:"data,omitempty"`
	Meta *CursorPaginationMeta `json:"meta,omitempty"`
}

// NewGetRecordingTranscriptions200Response instantiates a new GetRecordingTranscriptions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRecordingTranscriptions200Response() *GetRecordingTranscriptions200Response {
	this := GetRecordingTranscriptions200Response{}
	return &this
}

// NewGetRecordingTranscriptions200ResponseWithDefaults instantiates a new GetRecordingTranscriptions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRecordingTranscriptions200ResponseWithDefaults() *GetRecordingTranscriptions200Response {
	this := GetRecordingTranscriptions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GetRecordingTranscriptions200Response) GetData() []RecordingTranscription {
	if o == nil || IsNil(o.Data) {
		var ret []RecordingTranscription
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordingTranscriptions200Response) GetDataOk() ([]RecordingTranscription, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GetRecordingTranscriptions200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []RecordingTranscription and assigns it to the Data field.
func (o *GetRecordingTranscriptions200Response) SetData(v []RecordingTranscription) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GetRecordingTranscriptions200Response) GetMeta() CursorPaginationMeta {
	if o == nil || IsNil(o.Meta) {
		var ret CursorPaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRecordingTranscriptions200Response) GetMetaOk() (*CursorPaginationMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GetRecordingTranscriptions200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given CursorPaginationMeta and assigns it to the Meta field.
func (o *GetRecordingTranscriptions200Response) SetMeta(v CursorPaginationMeta) {
	o.Meta = &v
}

func (o GetRecordingTranscriptions200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRecordingTranscriptions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableGetRecordingTranscriptions200Response struct {
	value *GetRecordingTranscriptions200Response
	isSet bool
}

func (v NullableGetRecordingTranscriptions200Response) Get() *GetRecordingTranscriptions200Response {
	return v.value
}

func (v *NullableGetRecordingTranscriptions200Response) Set(val *GetRecordingTranscriptions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRecordingTranscriptions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRecordingTranscriptions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRecordingTranscriptions200Response(val *GetRecordingTranscriptions200Response) *NullableGetRecordingTranscriptions200Response {
	return &NullableGetRecordingTranscriptions200Response{value: val, isSet: true}
}

func (v NullableGetRecordingTranscriptions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRecordingTranscriptions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


