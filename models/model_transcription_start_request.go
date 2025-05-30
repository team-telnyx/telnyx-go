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

// checks if the TranscriptionStartRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TranscriptionStartRequest{}

// TranscriptionStartRequest struct for TranscriptionStartRequest
type TranscriptionStartRequest struct {
	// Engine to use for speech recognition. `A` - `Google`, `B` - `Telnyx`.
	TranscriptionEngine *string `json:"transcription_engine,omitempty"`
	TranscriptionEngineConfig *TranscriptionStartRequestTranscriptionEngineConfig `json:"transcription_engine_config,omitempty"`
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState *string `json:"client_state,omitempty"`
	// Indicates which leg of the call will be transcribed. Use `inbound` for the leg that requested the transcription, `outbound` for the other leg, and `both` for both legs of the call. Will default to `inbound`.
	TranscriptionTracks *string `json:"transcription_tracks,omitempty"`
	// Use this field to avoid duplicate commands. Telnyx will ignore any command with the same `command_id` for the same `call_control_id`.
	CommandId *string `json:"command_id,omitempty"`
}

// NewTranscriptionStartRequest instantiates a new TranscriptionStartRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranscriptionStartRequest() *TranscriptionStartRequest {
	this := TranscriptionStartRequest{}
	var transcriptionEngine string = "A"
	this.TranscriptionEngine = &transcriptionEngine
	var transcriptionTracks string = "inbound"
	this.TranscriptionTracks = &transcriptionTracks
	return &this
}

// NewTranscriptionStartRequestWithDefaults instantiates a new TranscriptionStartRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranscriptionStartRequestWithDefaults() *TranscriptionStartRequest {
	this := TranscriptionStartRequest{}
	var transcriptionEngine string = "A"
	this.TranscriptionEngine = &transcriptionEngine
	var transcriptionTracks string = "inbound"
	this.TranscriptionTracks = &transcriptionTracks
	return &this
}

// GetTranscriptionEngine returns the TranscriptionEngine field value if set, zero value otherwise.
func (o *TranscriptionStartRequest) GetTranscriptionEngine() string {
	if o == nil || IsNil(o.TranscriptionEngine) {
		var ret string
		return ret
	}
	return *o.TranscriptionEngine
}

// GetTranscriptionEngineOk returns a tuple with the TranscriptionEngine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptionStartRequest) GetTranscriptionEngineOk() (*string, bool) {
	if o == nil || IsNil(o.TranscriptionEngine) {
		return nil, false
	}
	return o.TranscriptionEngine, true
}

// HasTranscriptionEngine returns a boolean if a field has been set.
func (o *TranscriptionStartRequest) HasTranscriptionEngine() bool {
	if o != nil && !IsNil(o.TranscriptionEngine) {
		return true
	}

	return false
}

// SetTranscriptionEngine gets a reference to the given string and assigns it to the TranscriptionEngine field.
func (o *TranscriptionStartRequest) SetTranscriptionEngine(v string) {
	o.TranscriptionEngine = &v
}

// GetTranscriptionEngineConfig returns the TranscriptionEngineConfig field value if set, zero value otherwise.
func (o *TranscriptionStartRequest) GetTranscriptionEngineConfig() TranscriptionStartRequestTranscriptionEngineConfig {
	if o == nil || IsNil(o.TranscriptionEngineConfig) {
		var ret TranscriptionStartRequestTranscriptionEngineConfig
		return ret
	}
	return *o.TranscriptionEngineConfig
}

// GetTranscriptionEngineConfigOk returns a tuple with the TranscriptionEngineConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptionStartRequest) GetTranscriptionEngineConfigOk() (*TranscriptionStartRequestTranscriptionEngineConfig, bool) {
	if o == nil || IsNil(o.TranscriptionEngineConfig) {
		return nil, false
	}
	return o.TranscriptionEngineConfig, true
}

// HasTranscriptionEngineConfig returns a boolean if a field has been set.
func (o *TranscriptionStartRequest) HasTranscriptionEngineConfig() bool {
	if o != nil && !IsNil(o.TranscriptionEngineConfig) {
		return true
	}

	return false
}

// SetTranscriptionEngineConfig gets a reference to the given TranscriptionStartRequestTranscriptionEngineConfig and assigns it to the TranscriptionEngineConfig field.
func (o *TranscriptionStartRequest) SetTranscriptionEngineConfig(v TranscriptionStartRequestTranscriptionEngineConfig) {
	o.TranscriptionEngineConfig = &v
}

// GetClientState returns the ClientState field value if set, zero value otherwise.
func (o *TranscriptionStartRequest) GetClientState() string {
	if o == nil || IsNil(o.ClientState) {
		var ret string
		return ret
	}
	return *o.ClientState
}

// GetClientStateOk returns a tuple with the ClientState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptionStartRequest) GetClientStateOk() (*string, bool) {
	if o == nil || IsNil(o.ClientState) {
		return nil, false
	}
	return o.ClientState, true
}

// HasClientState returns a boolean if a field has been set.
func (o *TranscriptionStartRequest) HasClientState() bool {
	if o != nil && !IsNil(o.ClientState) {
		return true
	}

	return false
}

// SetClientState gets a reference to the given string and assigns it to the ClientState field.
func (o *TranscriptionStartRequest) SetClientState(v string) {
	o.ClientState = &v
}

// GetTranscriptionTracks returns the TranscriptionTracks field value if set, zero value otherwise.
func (o *TranscriptionStartRequest) GetTranscriptionTracks() string {
	if o == nil || IsNil(o.TranscriptionTracks) {
		var ret string
		return ret
	}
	return *o.TranscriptionTracks
}

// GetTranscriptionTracksOk returns a tuple with the TranscriptionTracks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptionStartRequest) GetTranscriptionTracksOk() (*string, bool) {
	if o == nil || IsNil(o.TranscriptionTracks) {
		return nil, false
	}
	return o.TranscriptionTracks, true
}

// HasTranscriptionTracks returns a boolean if a field has been set.
func (o *TranscriptionStartRequest) HasTranscriptionTracks() bool {
	if o != nil && !IsNil(o.TranscriptionTracks) {
		return true
	}

	return false
}

// SetTranscriptionTracks gets a reference to the given string and assigns it to the TranscriptionTracks field.
func (o *TranscriptionStartRequest) SetTranscriptionTracks(v string) {
	o.TranscriptionTracks = &v
}

// GetCommandId returns the CommandId field value if set, zero value otherwise.
func (o *TranscriptionStartRequest) GetCommandId() string {
	if o == nil || IsNil(o.CommandId) {
		var ret string
		return ret
	}
	return *o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranscriptionStartRequest) GetCommandIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommandId) {
		return nil, false
	}
	return o.CommandId, true
}

// HasCommandId returns a boolean if a field has been set.
func (o *TranscriptionStartRequest) HasCommandId() bool {
	if o != nil && !IsNil(o.CommandId) {
		return true
	}

	return false
}

// SetCommandId gets a reference to the given string and assigns it to the CommandId field.
func (o *TranscriptionStartRequest) SetCommandId(v string) {
	o.CommandId = &v
}

func (o TranscriptionStartRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TranscriptionStartRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TranscriptionEngine) {
		toSerialize["transcription_engine"] = o.TranscriptionEngine
	}
	if !IsNil(o.TranscriptionEngineConfig) {
		toSerialize["transcription_engine_config"] = o.TranscriptionEngineConfig
	}
	if !IsNil(o.ClientState) {
		toSerialize["client_state"] = o.ClientState
	}
	if !IsNil(o.TranscriptionTracks) {
		toSerialize["transcription_tracks"] = o.TranscriptionTracks
	}
	if !IsNil(o.CommandId) {
		toSerialize["command_id"] = o.CommandId
	}
	return toSerialize, nil
}

type NullableTranscriptionStartRequest struct {
	value *TranscriptionStartRequest
	isSet bool
}

func (v NullableTranscriptionStartRequest) Get() *TranscriptionStartRequest {
	return v.value
}

func (v *NullableTranscriptionStartRequest) Set(val *TranscriptionStartRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTranscriptionStartRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTranscriptionStartRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranscriptionStartRequest(val *TranscriptionStartRequest) *NullableTranscriptionStartRequest {
	return &NullableTranscriptionStartRequest{value: val, isSet: true}
}

func (v NullableTranscriptionStartRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranscriptionStartRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


