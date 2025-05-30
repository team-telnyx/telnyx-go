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

// checks if the CreateTexmlApplicationRequestOutbound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateTexmlApplicationRequestOutbound{}

// CreateTexmlApplicationRequestOutbound struct for CreateTexmlApplicationRequestOutbound
type CreateTexmlApplicationRequestOutbound struct {
	// When set, this will limit the total number of outbound calls to phone numbers associated with this connection.
	ChannelLimit *int32 `json:"channel_limit,omitempty"`
	// Identifies the associated outbound voice profile.
	OutboundVoiceProfileId *string `json:"outbound_voice_profile_id,omitempty"`
}

// NewCreateTexmlApplicationRequestOutbound instantiates a new CreateTexmlApplicationRequestOutbound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTexmlApplicationRequestOutbound() *CreateTexmlApplicationRequestOutbound {
	this := CreateTexmlApplicationRequestOutbound{}
	return &this
}

// NewCreateTexmlApplicationRequestOutboundWithDefaults instantiates a new CreateTexmlApplicationRequestOutbound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTexmlApplicationRequestOutboundWithDefaults() *CreateTexmlApplicationRequestOutbound {
	this := CreateTexmlApplicationRequestOutbound{}
	return &this
}

// GetChannelLimit returns the ChannelLimit field value if set, zero value otherwise.
func (o *CreateTexmlApplicationRequestOutbound) GetChannelLimit() int32 {
	if o == nil || IsNil(o.ChannelLimit) {
		var ret int32
		return ret
	}
	return *o.ChannelLimit
}

// GetChannelLimitOk returns a tuple with the ChannelLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTexmlApplicationRequestOutbound) GetChannelLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.ChannelLimit) {
		return nil, false
	}
	return o.ChannelLimit, true
}

// HasChannelLimit returns a boolean if a field has been set.
func (o *CreateTexmlApplicationRequestOutbound) HasChannelLimit() bool {
	if o != nil && !IsNil(o.ChannelLimit) {
		return true
	}

	return false
}

// SetChannelLimit gets a reference to the given int32 and assigns it to the ChannelLimit field.
func (o *CreateTexmlApplicationRequestOutbound) SetChannelLimit(v int32) {
	o.ChannelLimit = &v
}

// GetOutboundVoiceProfileId returns the OutboundVoiceProfileId field value if set, zero value otherwise.
func (o *CreateTexmlApplicationRequestOutbound) GetOutboundVoiceProfileId() string {
	if o == nil || IsNil(o.OutboundVoiceProfileId) {
		var ret string
		return ret
	}
	return *o.OutboundVoiceProfileId
}

// GetOutboundVoiceProfileIdOk returns a tuple with the OutboundVoiceProfileId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTexmlApplicationRequestOutbound) GetOutboundVoiceProfileIdOk() (*string, bool) {
	if o == nil || IsNil(o.OutboundVoiceProfileId) {
		return nil, false
	}
	return o.OutboundVoiceProfileId, true
}

// HasOutboundVoiceProfileId returns a boolean if a field has been set.
func (o *CreateTexmlApplicationRequestOutbound) HasOutboundVoiceProfileId() bool {
	if o != nil && !IsNil(o.OutboundVoiceProfileId) {
		return true
	}

	return false
}

// SetOutboundVoiceProfileId gets a reference to the given string and assigns it to the OutboundVoiceProfileId field.
func (o *CreateTexmlApplicationRequestOutbound) SetOutboundVoiceProfileId(v string) {
	o.OutboundVoiceProfileId = &v
}

func (o CreateTexmlApplicationRequestOutbound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateTexmlApplicationRequestOutbound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelLimit) {
		toSerialize["channel_limit"] = o.ChannelLimit
	}
	if !IsNil(o.OutboundVoiceProfileId) {
		toSerialize["outbound_voice_profile_id"] = o.OutboundVoiceProfileId
	}
	return toSerialize, nil
}

type NullableCreateTexmlApplicationRequestOutbound struct {
	value *CreateTexmlApplicationRequestOutbound
	isSet bool
}

func (v NullableCreateTexmlApplicationRequestOutbound) Get() *CreateTexmlApplicationRequestOutbound {
	return v.value
}

func (v *NullableCreateTexmlApplicationRequestOutbound) Set(val *CreateTexmlApplicationRequestOutbound) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTexmlApplicationRequestOutbound) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTexmlApplicationRequestOutbound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTexmlApplicationRequestOutbound(val *CreateTexmlApplicationRequestOutbound) *NullableCreateTexmlApplicationRequestOutbound {
	return &NullableCreateTexmlApplicationRequestOutbound{value: val, isSet: true}
}

func (v NullableCreateTexmlApplicationRequestOutbound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTexmlApplicationRequestOutbound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


