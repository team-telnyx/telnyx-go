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

// checks if the InlineObject1Meta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InlineObject1Meta{}

// InlineObject1Meta struct for InlineObject1Meta
type InlineObject1Meta struct {
	// Link to list porting order
	PortingOrderUrl *string `json:"porting_order_url,omitempty"`
}

// NewInlineObject1Meta instantiates a new InlineObject1Meta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject1Meta() *InlineObject1Meta {
	this := InlineObject1Meta{}
	return &this
}

// NewInlineObject1MetaWithDefaults instantiates a new InlineObject1Meta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject1MetaWithDefaults() *InlineObject1Meta {
	this := InlineObject1Meta{}
	return &this
}

// GetPortingOrderUrl returns the PortingOrderUrl field value if set, zero value otherwise.
func (o *InlineObject1Meta) GetPortingOrderUrl() string {
	if o == nil || IsNil(o.PortingOrderUrl) {
		var ret string
		return ret
	}
	return *o.PortingOrderUrl
}

// GetPortingOrderUrlOk returns a tuple with the PortingOrderUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject1Meta) GetPortingOrderUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PortingOrderUrl) {
		return nil, false
	}
	return o.PortingOrderUrl, true
}

// HasPortingOrderUrl returns a boolean if a field has been set.
func (o *InlineObject1Meta) HasPortingOrderUrl() bool {
	if o != nil && !IsNil(o.PortingOrderUrl) {
		return true
	}

	return false
}

// SetPortingOrderUrl gets a reference to the given string and assigns it to the PortingOrderUrl field.
func (o *InlineObject1Meta) SetPortingOrderUrl(v string) {
	o.PortingOrderUrl = &v
}

func (o InlineObject1Meta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InlineObject1Meta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PortingOrderUrl) {
		toSerialize["porting_order_url"] = o.PortingOrderUrl
	}
	return toSerialize, nil
}

type NullableInlineObject1Meta struct {
	value *InlineObject1Meta
	isSet bool
}

func (v NullableInlineObject1Meta) Get() *InlineObject1Meta {
	return v.value
}

func (v *NullableInlineObject1Meta) Set(val *InlineObject1Meta) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject1Meta) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject1Meta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject1Meta(val *InlineObject1Meta) *NullableInlineObject1Meta {
	return &NullableInlineObject1Meta{value: val, isSet: true}
}

func (v NullableInlineObject1Meta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject1Meta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


