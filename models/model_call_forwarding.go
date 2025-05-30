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

// checks if the CallForwarding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallForwarding{}

// CallForwarding The call forwarding settings for a phone number.
type CallForwarding struct {
	// Indicates if call forwarding will be enabled for this number if forwards_to and forwarding_type are filled in. Defaults to true for backwards compatibility with APIV1 use of numbers endpoints.
	CallForwardingEnabled *bool `json:"call_forwarding_enabled,omitempty"`
	// The phone number to which inbound calls to this number are forwarded. Inbound calls will not be forwarded if this field is left blank. If set, must be a +E.164-formatted phone number.
	ForwardsTo *string `json:"forwards_to,omitempty"`
	// Call forwarding type. 'forwards_to' must be set for this to have an effect.
	ForwardingType *string `json:"forwarding_type,omitempty"`
}

// NewCallForwarding instantiates a new CallForwarding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallForwarding() *CallForwarding {
	this := CallForwarding{}
	var callForwardingEnabled bool = true
	this.CallForwardingEnabled = &callForwardingEnabled
	return &this
}

// NewCallForwardingWithDefaults instantiates a new CallForwarding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallForwardingWithDefaults() *CallForwarding {
	this := CallForwarding{}
	var callForwardingEnabled bool = true
	this.CallForwardingEnabled = &callForwardingEnabled
	return &this
}

// GetCallForwardingEnabled returns the CallForwardingEnabled field value if set, zero value otherwise.
func (o *CallForwarding) GetCallForwardingEnabled() bool {
	if o == nil || IsNil(o.CallForwardingEnabled) {
		var ret bool
		return ret
	}
	return *o.CallForwardingEnabled
}

// GetCallForwardingEnabledOk returns a tuple with the CallForwardingEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallForwarding) GetCallForwardingEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.CallForwardingEnabled) {
		return nil, false
	}
	return o.CallForwardingEnabled, true
}

// HasCallForwardingEnabled returns a boolean if a field has been set.
func (o *CallForwarding) HasCallForwardingEnabled() bool {
	if o != nil && !IsNil(o.CallForwardingEnabled) {
		return true
	}

	return false
}

// SetCallForwardingEnabled gets a reference to the given bool and assigns it to the CallForwardingEnabled field.
func (o *CallForwarding) SetCallForwardingEnabled(v bool) {
	o.CallForwardingEnabled = &v
}

// GetForwardsTo returns the ForwardsTo field value if set, zero value otherwise.
func (o *CallForwarding) GetForwardsTo() string {
	if o == nil || IsNil(o.ForwardsTo) {
		var ret string
		return ret
	}
	return *o.ForwardsTo
}

// GetForwardsToOk returns a tuple with the ForwardsTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallForwarding) GetForwardsToOk() (*string, bool) {
	if o == nil || IsNil(o.ForwardsTo) {
		return nil, false
	}
	return o.ForwardsTo, true
}

// HasForwardsTo returns a boolean if a field has been set.
func (o *CallForwarding) HasForwardsTo() bool {
	if o != nil && !IsNil(o.ForwardsTo) {
		return true
	}

	return false
}

// SetForwardsTo gets a reference to the given string and assigns it to the ForwardsTo field.
func (o *CallForwarding) SetForwardsTo(v string) {
	o.ForwardsTo = &v
}

// GetForwardingType returns the ForwardingType field value if set, zero value otherwise.
func (o *CallForwarding) GetForwardingType() string {
	if o == nil || IsNil(o.ForwardingType) {
		var ret string
		return ret
	}
	return *o.ForwardingType
}

// GetForwardingTypeOk returns a tuple with the ForwardingType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallForwarding) GetForwardingTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ForwardingType) {
		return nil, false
	}
	return o.ForwardingType, true
}

// HasForwardingType returns a boolean if a field has been set.
func (o *CallForwarding) HasForwardingType() bool {
	if o != nil && !IsNil(o.ForwardingType) {
		return true
	}

	return false
}

// SetForwardingType gets a reference to the given string and assigns it to the ForwardingType field.
func (o *CallForwarding) SetForwardingType(v string) {
	o.ForwardingType = &v
}

func (o CallForwarding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallForwarding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallForwardingEnabled) {
		toSerialize["call_forwarding_enabled"] = o.CallForwardingEnabled
	}
	if !IsNil(o.ForwardsTo) {
		toSerialize["forwards_to"] = o.ForwardsTo
	}
	if !IsNil(o.ForwardingType) {
		toSerialize["forwarding_type"] = o.ForwardingType
	}
	return toSerialize, nil
}

type NullableCallForwarding struct {
	value *CallForwarding
	isSet bool
}

func (v NullableCallForwarding) Get() *CallForwarding {
	return v.value
}

func (v *NullableCallForwarding) Set(val *CallForwarding) {
	v.value = val
	v.isSet = true
}

func (v NullableCallForwarding) IsSet() bool {
	return v.isSet
}

func (v *NullableCallForwarding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallForwarding(val *CallForwarding) *NullableCallForwarding {
	return &NullableCallForwarding{value: val, isSet: true}
}

func (v NullableCallForwarding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallForwarding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


