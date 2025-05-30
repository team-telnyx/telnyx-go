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

// checks if the ForbiddenError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForbiddenError{}

// ForbiddenError struct for ForbiddenError
type ForbiddenError struct {
	Code interface{} `json:"code,omitempty"`
	Title interface{} `json:"title,omitempty"`
	Detail interface{} `json:"detail,omitempty"`
	Source *UpdateOutboundChannelsDefaultResponseErrorsInnerSource `json:"source,omitempty"`
	Meta *ForbiddenErrorAllOfMeta `json:"meta,omitempty"`
}

// NewForbiddenError instantiates a new ForbiddenError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForbiddenError() *ForbiddenError {
	this := ForbiddenError{}
	return &this
}

// NewForbiddenErrorWithDefaults instantiates a new ForbiddenError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForbiddenErrorWithDefaults() *ForbiddenError {
	this := ForbiddenError{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForbiddenError) GetCode() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForbiddenError) GetCodeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return &o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ForbiddenError) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given interface{} and assigns it to the Code field.
func (o *ForbiddenError) SetCode(v interface{}) {
	o.Code = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForbiddenError) GetTitle() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForbiddenError) GetTitleOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return &o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ForbiddenError) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given interface{} and assigns it to the Title field.
func (o *ForbiddenError) SetTitle(v interface{}) {
	o.Title = v
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ForbiddenError) GetDetail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ForbiddenError) GetDetailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return &o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ForbiddenError) HasDetail() bool {
	if o != nil && !IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given interface{} and assigns it to the Detail field.
func (o *ForbiddenError) SetDetail(v interface{}) {
	o.Detail = v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *ForbiddenError) GetSource() UpdateOutboundChannelsDefaultResponseErrorsInnerSource {
	if o == nil || IsNil(o.Source) {
		var ret UpdateOutboundChannelsDefaultResponseErrorsInnerSource
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForbiddenError) GetSourceOk() (*UpdateOutboundChannelsDefaultResponseErrorsInnerSource, bool) {
	if o == nil || IsNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *ForbiddenError) HasSource() bool {
	if o != nil && !IsNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given UpdateOutboundChannelsDefaultResponseErrorsInnerSource and assigns it to the Source field.
func (o *ForbiddenError) SetSource(v UpdateOutboundChannelsDefaultResponseErrorsInnerSource) {
	o.Source = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ForbiddenError) GetMeta() ForbiddenErrorAllOfMeta {
	if o == nil || IsNil(o.Meta) {
		var ret ForbiddenErrorAllOfMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForbiddenError) GetMetaOk() (*ForbiddenErrorAllOfMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ForbiddenError) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given ForbiddenErrorAllOfMeta and assigns it to the Meta field.
func (o *ForbiddenError) SetMeta(v ForbiddenErrorAllOfMeta) {
	o.Meta = &v
}

func (o ForbiddenError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForbiddenError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if !IsNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableForbiddenError struct {
	value *ForbiddenError
	isSet bool
}

func (v NullableForbiddenError) Get() *ForbiddenError {
	return v.value
}

func (v *NullableForbiddenError) Set(val *ForbiddenError) {
	v.value = val
	v.isSet = true
}

func (v NullableForbiddenError) IsSet() bool {
	return v.isSet
}

func (v *NullableForbiddenError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForbiddenError(val *ForbiddenError) *NullableForbiddenError {
	return &NullableForbiddenError{value: val, isSet: true}
}

func (v NullableForbiddenError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForbiddenError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


