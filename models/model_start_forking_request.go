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

// checks if the StartForkingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StartForkingRequest{}

// StartForkingRequest struct for StartForkingRequest
type StartForkingRequest struct {
	// The network target, <udp:ip_address:port>, where the call's incoming RTP media packets should be forwarded.
	Rx *string `json:"rx,omitempty"`
	// Optionally specify a media type to stream. If `decrypted` selected, Telnyx will decrypt incoming SIP media before forking to the target. `rx` and `tx` are required fields if `decrypted` selected.
	StreamType *string `json:"stream_type,omitempty"`
	// The network target, <udp:ip_address:port>, where the call's outgoing RTP media packets should be forwarded.
	Tx *string `json:"tx,omitempty"`
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState *string `json:"client_state,omitempty"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.
	CommandId *string `json:"command_id,omitempty"`
}

// NewStartForkingRequest instantiates a new StartForkingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStartForkingRequest() *StartForkingRequest {
	this := StartForkingRequest{}
	var streamType string = "decrypted"
	this.StreamType = &streamType
	return &this
}

// NewStartForkingRequestWithDefaults instantiates a new StartForkingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStartForkingRequestWithDefaults() *StartForkingRequest {
	this := StartForkingRequest{}
	var streamType string = "decrypted"
	this.StreamType = &streamType
	return &this
}

// GetRx returns the Rx field value if set, zero value otherwise.
func (o *StartForkingRequest) GetRx() string {
	if o == nil || IsNil(o.Rx) {
		var ret string
		return ret
	}
	return *o.Rx
}

// GetRxOk returns a tuple with the Rx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartForkingRequest) GetRxOk() (*string, bool) {
	if o == nil || IsNil(o.Rx) {
		return nil, false
	}
	return o.Rx, true
}

// HasRx returns a boolean if a field has been set.
func (o *StartForkingRequest) HasRx() bool {
	if o != nil && !IsNil(o.Rx) {
		return true
	}

	return false
}

// SetRx gets a reference to the given string and assigns it to the Rx field.
func (o *StartForkingRequest) SetRx(v string) {
	o.Rx = &v
}

// GetStreamType returns the StreamType field value if set, zero value otherwise.
func (o *StartForkingRequest) GetStreamType() string {
	if o == nil || IsNil(o.StreamType) {
		var ret string
		return ret
	}
	return *o.StreamType
}

// GetStreamTypeOk returns a tuple with the StreamType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartForkingRequest) GetStreamTypeOk() (*string, bool) {
	if o == nil || IsNil(o.StreamType) {
		return nil, false
	}
	return o.StreamType, true
}

// HasStreamType returns a boolean if a field has been set.
func (o *StartForkingRequest) HasStreamType() bool {
	if o != nil && !IsNil(o.StreamType) {
		return true
	}

	return false
}

// SetStreamType gets a reference to the given string and assigns it to the StreamType field.
func (o *StartForkingRequest) SetStreamType(v string) {
	o.StreamType = &v
}

// GetTx returns the Tx field value if set, zero value otherwise.
func (o *StartForkingRequest) GetTx() string {
	if o == nil || IsNil(o.Tx) {
		var ret string
		return ret
	}
	return *o.Tx
}

// GetTxOk returns a tuple with the Tx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartForkingRequest) GetTxOk() (*string, bool) {
	if o == nil || IsNil(o.Tx) {
		return nil, false
	}
	return o.Tx, true
}

// HasTx returns a boolean if a field has been set.
func (o *StartForkingRequest) HasTx() bool {
	if o != nil && !IsNil(o.Tx) {
		return true
	}

	return false
}

// SetTx gets a reference to the given string and assigns it to the Tx field.
func (o *StartForkingRequest) SetTx(v string) {
	o.Tx = &v
}

// GetClientState returns the ClientState field value if set, zero value otherwise.
func (o *StartForkingRequest) GetClientState() string {
	if o == nil || IsNil(o.ClientState) {
		var ret string
		return ret
	}
	return *o.ClientState
}

// GetClientStateOk returns a tuple with the ClientState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartForkingRequest) GetClientStateOk() (*string, bool) {
	if o == nil || IsNil(o.ClientState) {
		return nil, false
	}
	return o.ClientState, true
}

// HasClientState returns a boolean if a field has been set.
func (o *StartForkingRequest) HasClientState() bool {
	if o != nil && !IsNil(o.ClientState) {
		return true
	}

	return false
}

// SetClientState gets a reference to the given string and assigns it to the ClientState field.
func (o *StartForkingRequest) SetClientState(v string) {
	o.ClientState = &v
}

// GetCommandId returns the CommandId field value if set, zero value otherwise.
func (o *StartForkingRequest) GetCommandId() string {
	if o == nil || IsNil(o.CommandId) {
		var ret string
		return ret
	}
	return *o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StartForkingRequest) GetCommandIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommandId) {
		return nil, false
	}
	return o.CommandId, true
}

// HasCommandId returns a boolean if a field has been set.
func (o *StartForkingRequest) HasCommandId() bool {
	if o != nil && !IsNil(o.CommandId) {
		return true
	}

	return false
}

// SetCommandId gets a reference to the given string and assigns it to the CommandId field.
func (o *StartForkingRequest) SetCommandId(v string) {
	o.CommandId = &v
}

func (o StartForkingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StartForkingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rx) {
		toSerialize["rx"] = o.Rx
	}
	if !IsNil(o.StreamType) {
		toSerialize["stream_type"] = o.StreamType
	}
	if !IsNil(o.Tx) {
		toSerialize["tx"] = o.Tx
	}
	if !IsNil(o.ClientState) {
		toSerialize["client_state"] = o.ClientState
	}
	if !IsNil(o.CommandId) {
		toSerialize["command_id"] = o.CommandId
	}
	return toSerialize, nil
}

type NullableStartForkingRequest struct {
	value *StartForkingRequest
	isSet bool
}

func (v NullableStartForkingRequest) Get() *StartForkingRequest {
	return v.value
}

func (v *NullableStartForkingRequest) Set(val *StartForkingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStartForkingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStartForkingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStartForkingRequest(val *StartForkingRequest) *NullableStartForkingRequest {
	return &NullableStartForkingRequest{value: val, isSet: true}
}

func (v NullableStartForkingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStartForkingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


