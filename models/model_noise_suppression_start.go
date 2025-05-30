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

// checks if the NoiseSuppressionStart type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NoiseSuppressionStart{}

// NoiseSuppressionStart struct for NoiseSuppressionStart
type NoiseSuppressionStart struct {
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState *string `json:"client_state,omitempty"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.
	CommandId *string `json:"command_id,omitempty"`
	Direction *NoiseSuppressionDirection `json:"direction,omitempty"`
}

// NewNoiseSuppressionStart instantiates a new NoiseSuppressionStart object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNoiseSuppressionStart() *NoiseSuppressionStart {
	this := NoiseSuppressionStart{}
	var direction NoiseSuppressionDirection = INBOUND
	this.Direction = &direction
	return &this
}

// NewNoiseSuppressionStartWithDefaults instantiates a new NoiseSuppressionStart object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNoiseSuppressionStartWithDefaults() *NoiseSuppressionStart {
	this := NoiseSuppressionStart{}
	var direction NoiseSuppressionDirection = INBOUND
	this.Direction = &direction
	return &this
}

// GetClientState returns the ClientState field value if set, zero value otherwise.
func (o *NoiseSuppressionStart) GetClientState() string {
	if o == nil || IsNil(o.ClientState) {
		var ret string
		return ret
	}
	return *o.ClientState
}

// GetClientStateOk returns a tuple with the ClientState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoiseSuppressionStart) GetClientStateOk() (*string, bool) {
	if o == nil || IsNil(o.ClientState) {
		return nil, false
	}
	return o.ClientState, true
}

// HasClientState returns a boolean if a field has been set.
func (o *NoiseSuppressionStart) HasClientState() bool {
	if o != nil && !IsNil(o.ClientState) {
		return true
	}

	return false
}

// SetClientState gets a reference to the given string and assigns it to the ClientState field.
func (o *NoiseSuppressionStart) SetClientState(v string) {
	o.ClientState = &v
}

// GetCommandId returns the CommandId field value if set, zero value otherwise.
func (o *NoiseSuppressionStart) GetCommandId() string {
	if o == nil || IsNil(o.CommandId) {
		var ret string
		return ret
	}
	return *o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoiseSuppressionStart) GetCommandIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommandId) {
		return nil, false
	}
	return o.CommandId, true
}

// HasCommandId returns a boolean if a field has been set.
func (o *NoiseSuppressionStart) HasCommandId() bool {
	if o != nil && !IsNil(o.CommandId) {
		return true
	}

	return false
}

// SetCommandId gets a reference to the given string and assigns it to the CommandId field.
func (o *NoiseSuppressionStart) SetCommandId(v string) {
	o.CommandId = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *NoiseSuppressionStart) GetDirection() NoiseSuppressionDirection {
	if o == nil || IsNil(o.Direction) {
		var ret NoiseSuppressionDirection
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NoiseSuppressionStart) GetDirectionOk() (*NoiseSuppressionDirection, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *NoiseSuppressionStart) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given NoiseSuppressionDirection and assigns it to the Direction field.
func (o *NoiseSuppressionStart) SetDirection(v NoiseSuppressionDirection) {
	o.Direction = &v
}

func (o NoiseSuppressionStart) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NoiseSuppressionStart) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ClientState) {
		toSerialize["client_state"] = o.ClientState
	}
	if !IsNil(o.CommandId) {
		toSerialize["command_id"] = o.CommandId
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	return toSerialize, nil
}

type NullableNoiseSuppressionStart struct {
	value *NoiseSuppressionStart
	isSet bool
}

func (v NullableNoiseSuppressionStart) Get() *NoiseSuppressionStart {
	return v.value
}

func (v *NullableNoiseSuppressionStart) Set(val *NoiseSuppressionStart) {
	v.value = val
	v.isSet = true
}

func (v NullableNoiseSuppressionStart) IsSet() bool {
	return v.isSet
}

func (v *NullableNoiseSuppressionStart) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoiseSuppressionStart(val *NoiseSuppressionStart) *NullableNoiseSuppressionStart {
	return &NullableNoiseSuppressionStart{value: val, isSet: true}
}

func (v NullableNoiseSuppressionStart) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoiseSuppressionStart) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


