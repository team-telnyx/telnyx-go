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

// checks if the ListPortoutReports200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListPortoutReports200Response{}

// ListPortoutReports200Response struct for ListPortoutReports200Response
type ListPortoutReports200Response struct {
	Data []PortoutReport `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
}

// NewListPortoutReports200Response instantiates a new ListPortoutReports200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPortoutReports200Response() *ListPortoutReports200Response {
	this := ListPortoutReports200Response{}
	return &this
}

// NewListPortoutReports200ResponseWithDefaults instantiates a new ListPortoutReports200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPortoutReports200ResponseWithDefaults() *ListPortoutReports200Response {
	this := ListPortoutReports200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListPortoutReports200Response) GetData() []PortoutReport {
	if o == nil || IsNil(o.Data) {
		var ret []PortoutReport
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPortoutReports200Response) GetDataOk() ([]PortoutReport, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListPortoutReports200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []PortoutReport and assigns it to the Data field.
func (o *ListPortoutReports200Response) SetData(v []PortoutReport) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListPortoutReports200Response) GetMeta() PaginationMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListPortoutReports200Response) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListPortoutReports200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *ListPortoutReports200Response) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

func (o ListPortoutReports200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListPortoutReports200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableListPortoutReports200Response struct {
	value *ListPortoutReports200Response
	isSet bool
}

func (v NullableListPortoutReports200Response) Get() *ListPortoutReports200Response {
	return v.value
}

func (v *NullableListPortoutReports200Response) Set(val *ListPortoutReports200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListPortoutReports200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListPortoutReports200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPortoutReports200Response(val *ListPortoutReports200Response) *NullableListPortoutReports200Response {
	return &NullableListPortoutReports200Response{value: val, isSet: true}
}

func (v NullableListPortoutReports200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPortoutReports200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


