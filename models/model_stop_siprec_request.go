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

// checks if the StopSiprecRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopSiprecRequest{}

// StopSiprecRequest struct for StopSiprecRequest
type StopSiprecRequest struct {
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState *string `json:"client_state,omitempty"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.
	CommandId *string `json:"command_id,omitempty"`
}

// NewStopSiprecRequest instantiates a new StopSiprecRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopSiprecRequest() *StopSiprecRequest {
	this := StopSiprecRequest{}
	return &this
}

// NewStopSiprecRequestWithDefaults instantiates a new StopSiprecRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopSiprecRequestWithDefaults() *StopSiprecRequest {
	this := StopSiprecRequest{}
	return &this
}

// GetClientState returns the ClientState field value if set, zero value otherwise.
func (o *StopSiprecRequest) GetClientState() string {
	if o == nil || IsNil(o.ClientState) {
		var ret string
		return ret
	}
	return *o.ClientState
}

// GetClientStateOk returns a tuple with the ClientState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopSiprecRequest) GetClientStateOk() (*string, bool) {
	if o == nil || IsNil(o.ClientState) {
		return nil, false
	}
	return o.ClientState, true
}

// HasClientState returns a boolean if a field has been set.
func (o *StopSiprecRequest) HasClientState() bool {
	if o != nil && !IsNil(o.ClientState) {
		return true
	}

	return false
}

// SetClientState gets a reference to the given string and assigns it to the ClientState field.
func (o *StopSiprecRequest) SetClientState(v string) {
	o.ClientState = &v
}

// GetCommandId returns the CommandId field value if set, zero value otherwise.
func (o *StopSiprecRequest) GetCommandId() string {
	if o == nil || IsNil(o.CommandId) {
		var ret string
		return ret
	}
	return *o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopSiprecRequest) GetCommandIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommandId) {
		return nil, false
	}
	return o.CommandId, true
}

// HasCommandId returns a boolean if a field has been set.
func (o *StopSiprecRequest) HasCommandId() bool {
	if o != nil && !IsNil(o.CommandId) {
		return true
	}

	return false
}

// SetCommandId gets a reference to the given string and assigns it to the CommandId field.
func (o *StopSiprecRequest) SetCommandId(v string) {
	o.CommandId = &v
}

func (o StopSiprecRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopSiprecRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientState) {
		toSerialize["client_state"] = o.ClientState
	}
	if !IsNil(o.CommandId) {
		toSerialize["command_id"] = o.CommandId
	}
	return toSerialize, nil
}

type NullableStopSiprecRequest struct {
	value *StopSiprecRequest
	isSet bool
}

func (v NullableStopSiprecRequest) Get() *StopSiprecRequest {
	return v.value
}

func (v *NullableStopSiprecRequest) Set(val *StopSiprecRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStopSiprecRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStopSiprecRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopSiprecRequest(val *StopSiprecRequest) *NullableStopSiprecRequest {
	return &NullableStopSiprecRequest{value: val, isSet: true}
}

func (v NullableStopSiprecRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopSiprecRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


