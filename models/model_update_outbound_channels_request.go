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
	"bytes"
	"fmt"
)

// checks if the UpdateOutboundChannelsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOutboundChannelsRequest{}

// UpdateOutboundChannelsRequest struct for UpdateOutboundChannelsRequest
type UpdateOutboundChannelsRequest struct {
	// The new number of concurrent channels for the account
	Channels int32 `json:"channels"`
}

type _UpdateOutboundChannelsRequest UpdateOutboundChannelsRequest

// NewUpdateOutboundChannelsRequest instantiates a new UpdateOutboundChannelsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOutboundChannelsRequest(channels int32) *UpdateOutboundChannelsRequest {
	this := UpdateOutboundChannelsRequest{}
	this.Channels = channels
	return &this
}

// NewUpdateOutboundChannelsRequestWithDefaults instantiates a new UpdateOutboundChannelsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOutboundChannelsRequestWithDefaults() *UpdateOutboundChannelsRequest {
	this := UpdateOutboundChannelsRequest{}
	return &this
}

// GetChannels returns the Channels field value
func (o *UpdateOutboundChannelsRequest) GetChannels() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value
// and a boolean to check if the value has been set.
func (o *UpdateOutboundChannelsRequest) GetChannelsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Channels, true
}

// SetChannels sets field value
func (o *UpdateOutboundChannelsRequest) SetChannels(v int32) {
	o.Channels = v
}

func (o UpdateOutboundChannelsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOutboundChannelsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["channels"] = o.Channels
	return toSerialize, nil
}

func (o *UpdateOutboundChannelsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"channels",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateOutboundChannelsRequest := _UpdateOutboundChannelsRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateOutboundChannelsRequest)

	if err != nil {
		return err
	}

	*o = UpdateOutboundChannelsRequest(varUpdateOutboundChannelsRequest)

	return err
}

type NullableUpdateOutboundChannelsRequest struct {
	value *UpdateOutboundChannelsRequest
	isSet bool
}

func (v NullableUpdateOutboundChannelsRequest) Get() *UpdateOutboundChannelsRequest {
	return v.value
}

func (v *NullableUpdateOutboundChannelsRequest) Set(val *UpdateOutboundChannelsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOutboundChannelsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOutboundChannelsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOutboundChannelsRequest(val *UpdateOutboundChannelsRequest) *NullableUpdateOutboundChannelsRequest {
	return &NullableUpdateOutboundChannelsRequest{value: val, isSet: true}
}

func (v NullableUpdateOutboundChannelsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOutboundChannelsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


