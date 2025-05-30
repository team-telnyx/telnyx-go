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

// checks if the GetSubRequestByPortingOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetSubRequestByPortingOrder{}

// GetSubRequestByPortingOrder struct for GetSubRequestByPortingOrder
type GetSubRequestByPortingOrder struct {
	// Identifies the Sub Request associated with the Porting Order
	SubRequestId *string `json:"sub_request_id,omitempty"`
	// Identifies the Port Request associated with the Porting Order
	PortRequestId *string `json:"port_request_id,omitempty"`
}

// NewGetSubRequestByPortingOrder instantiates a new GetSubRequestByPortingOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetSubRequestByPortingOrder() *GetSubRequestByPortingOrder {
	this := GetSubRequestByPortingOrder{}
	return &this
}

// NewGetSubRequestByPortingOrderWithDefaults instantiates a new GetSubRequestByPortingOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetSubRequestByPortingOrderWithDefaults() *GetSubRequestByPortingOrder {
	this := GetSubRequestByPortingOrder{}
	return &this
}

// GetSubRequestId returns the SubRequestId field value if set, zero value otherwise.
func (o *GetSubRequestByPortingOrder) GetSubRequestId() string {
	if o == nil || IsNil(o.SubRequestId) {
		var ret string
		return ret
	}
	return *o.SubRequestId
}

// GetSubRequestIdOk returns a tuple with the SubRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubRequestByPortingOrder) GetSubRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.SubRequestId) {
		return nil, false
	}
	return o.SubRequestId, true
}

// HasSubRequestId returns a boolean if a field has been set.
func (o *GetSubRequestByPortingOrder) HasSubRequestId() bool {
	if o != nil && !IsNil(o.SubRequestId) {
		return true
	}

	return false
}

// SetSubRequestId gets a reference to the given string and assigns it to the SubRequestId field.
func (o *GetSubRequestByPortingOrder) SetSubRequestId(v string) {
	o.SubRequestId = &v
}

// GetPortRequestId returns the PortRequestId field value if set, zero value otherwise.
func (o *GetSubRequestByPortingOrder) GetPortRequestId() string {
	if o == nil || IsNil(o.PortRequestId) {
		var ret string
		return ret
	}
	return *o.PortRequestId
}

// GetPortRequestIdOk returns a tuple with the PortRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetSubRequestByPortingOrder) GetPortRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortRequestId) {
		return nil, false
	}
	return o.PortRequestId, true
}

// HasPortRequestId returns a boolean if a field has been set.
func (o *GetSubRequestByPortingOrder) HasPortRequestId() bool {
	if o != nil && !IsNil(o.PortRequestId) {
		return true
	}

	return false
}

// SetPortRequestId gets a reference to the given string and assigns it to the PortRequestId field.
func (o *GetSubRequestByPortingOrder) SetPortRequestId(v string) {
	o.PortRequestId = &v
}

func (o GetSubRequestByPortingOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetSubRequestByPortingOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubRequestId) {
		toSerialize["sub_request_id"] = o.SubRequestId
	}
	if !IsNil(o.PortRequestId) {
		toSerialize["port_request_id"] = o.PortRequestId
	}
	return toSerialize, nil
}

type NullableGetSubRequestByPortingOrder struct {
	value *GetSubRequestByPortingOrder
	isSet bool
}

func (v NullableGetSubRequestByPortingOrder) Get() *GetSubRequestByPortingOrder {
	return v.value
}

func (v *NullableGetSubRequestByPortingOrder) Set(val *GetSubRequestByPortingOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableGetSubRequestByPortingOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableGetSubRequestByPortingOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetSubRequestByPortingOrder(val *GetSubRequestByPortingOrder) *NullableGetSubRequestByPortingOrder {
	return &NullableGetSubRequestByPortingOrder{value: val, isSet: true}
}

func (v NullableGetSubRequestByPortingOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetSubRequestByPortingOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


