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

// checks if the ListBulkSimCardActions200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListBulkSimCardActions200Response{}

// ListBulkSimCardActions200Response struct for ListBulkSimCardActions200Response
type ListBulkSimCardActions200Response struct {
	Data []BulkSIMCardActionDetailed `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
}

// NewListBulkSimCardActions200Response instantiates a new ListBulkSimCardActions200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListBulkSimCardActions200Response() *ListBulkSimCardActions200Response {
	this := ListBulkSimCardActions200Response{}
	return &this
}

// NewListBulkSimCardActions200ResponseWithDefaults instantiates a new ListBulkSimCardActions200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListBulkSimCardActions200ResponseWithDefaults() *ListBulkSimCardActions200Response {
	this := ListBulkSimCardActions200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListBulkSimCardActions200Response) GetData() []BulkSIMCardActionDetailed {
	if o == nil || IsNil(o.Data) {
		var ret []BulkSIMCardActionDetailed
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBulkSimCardActions200Response) GetDataOk() ([]BulkSIMCardActionDetailed, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListBulkSimCardActions200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []BulkSIMCardActionDetailed and assigns it to the Data field.
func (o *ListBulkSimCardActions200Response) SetData(v []BulkSIMCardActionDetailed) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ListBulkSimCardActions200Response) GetMeta() PaginationMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListBulkSimCardActions200Response) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ListBulkSimCardActions200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *ListBulkSimCardActions200Response) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

func (o ListBulkSimCardActions200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListBulkSimCardActions200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableListBulkSimCardActions200Response struct {
	value *ListBulkSimCardActions200Response
	isSet bool
}

func (v NullableListBulkSimCardActions200Response) Get() *ListBulkSimCardActions200Response {
	return v.value
}

func (v *NullableListBulkSimCardActions200Response) Set(val *ListBulkSimCardActions200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListBulkSimCardActions200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListBulkSimCardActions200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListBulkSimCardActions200Response(val *ListBulkSimCardActions200Response) *NullableListBulkSimCardActions200Response {
	return &NullableListBulkSimCardActions200Response{value: val, isSet: true}
}

func (v NullableListBulkSimCardActions200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListBulkSimCardActions200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


