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

// checks if the ConferenceStopRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConferenceStopRequest{}

// ConferenceStopRequest struct for ConferenceStopRequest
type ConferenceStopRequest struct {
	// List of call control ids identifying participants the audio file should stop be played to. If not given, the audio will be stoped to the entire conference.
	CallControlIds []string `json:"call_control_ids,omitempty"`
}

// NewConferenceStopRequest instantiates a new ConferenceStopRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConferenceStopRequest() *ConferenceStopRequest {
	this := ConferenceStopRequest{}
	return &this
}

// NewConferenceStopRequestWithDefaults instantiates a new ConferenceStopRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConferenceStopRequestWithDefaults() *ConferenceStopRequest {
	this := ConferenceStopRequest{}
	return &this
}

// GetCallControlIds returns the CallControlIds field value if set, zero value otherwise.
func (o *ConferenceStopRequest) GetCallControlIds() []string {
	if o == nil || IsNil(o.CallControlIds) {
		var ret []string
		return ret
	}
	return o.CallControlIds
}

// GetCallControlIdsOk returns a tuple with the CallControlIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceStopRequest) GetCallControlIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.CallControlIds) {
		return nil, false
	}
	return o.CallControlIds, true
}

// HasCallControlIds returns a boolean if a field has been set.
func (o *ConferenceStopRequest) HasCallControlIds() bool {
	if o != nil && !IsNil(o.CallControlIds) {
		return true
	}

	return false
}

// SetCallControlIds gets a reference to the given []string and assigns it to the CallControlIds field.
func (o *ConferenceStopRequest) SetCallControlIds(v []string) {
	o.CallControlIds = v
}

func (o ConferenceStopRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConferenceStopRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallControlIds) {
		toSerialize["call_control_ids"] = o.CallControlIds
	}
	return toSerialize, nil
}

type NullableConferenceStopRequest struct {
	value *ConferenceStopRequest
	isSet bool
}

func (v NullableConferenceStopRequest) Get() *ConferenceStopRequest {
	return v.value
}

func (v *NullableConferenceStopRequest) Set(val *ConferenceStopRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConferenceStopRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConferenceStopRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConferenceStopRequest(val *ConferenceStopRequest) *NullableConferenceStopRequest {
	return &NullableConferenceStopRequest{value: val, isSet: true}
}

func (v NullableConferenceStopRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConferenceStopRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


