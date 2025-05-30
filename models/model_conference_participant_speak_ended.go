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

// checks if the ConferenceParticipantSpeakEnded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConferenceParticipantSpeakEnded{}

// ConferenceParticipantSpeakEnded struct for ConferenceParticipantSpeakEnded
type ConferenceParticipantSpeakEnded struct {
	// Identifies the type of the resource.
	RecordType *string `json:"record_type,omitempty"`
	// The type of event being delivered.
	EventType *string `json:"event_type,omitempty"`
	// Identifies the type of resource.
	Id *string `json:"id,omitempty"`
	Payload *ConferenceParticipantSpeakEndedPayload `json:"payload,omitempty"`
}

// NewConferenceParticipantSpeakEnded instantiates a new ConferenceParticipantSpeakEnded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConferenceParticipantSpeakEnded() *ConferenceParticipantSpeakEnded {
	this := ConferenceParticipantSpeakEnded{}
	return &this
}

// NewConferenceParticipantSpeakEndedWithDefaults instantiates a new ConferenceParticipantSpeakEnded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConferenceParticipantSpeakEndedWithDefaults() *ConferenceParticipantSpeakEnded {
	this := ConferenceParticipantSpeakEnded{}
	return &this
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *ConferenceParticipantSpeakEnded) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceParticipantSpeakEnded) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *ConferenceParticipantSpeakEnded) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *ConferenceParticipantSpeakEnded) SetRecordType(v string) {
	o.RecordType = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *ConferenceParticipantSpeakEnded) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceParticipantSpeakEnded) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *ConferenceParticipantSpeakEnded) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *ConferenceParticipantSpeakEnded) SetEventType(v string) {
	o.EventType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ConferenceParticipantSpeakEnded) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceParticipantSpeakEnded) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ConferenceParticipantSpeakEnded) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ConferenceParticipantSpeakEnded) SetId(v string) {
	o.Id = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *ConferenceParticipantSpeakEnded) GetPayload() ConferenceParticipantSpeakEndedPayload {
	if o == nil || IsNil(o.Payload) {
		var ret ConferenceParticipantSpeakEndedPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceParticipantSpeakEnded) GetPayloadOk() (*ConferenceParticipantSpeakEndedPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *ConferenceParticipantSpeakEnded) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given ConferenceParticipantSpeakEndedPayload and assigns it to the Payload field.
func (o *ConferenceParticipantSpeakEnded) SetPayload(v ConferenceParticipantSpeakEndedPayload) {
	o.Payload = &v
}

func (o ConferenceParticipantSpeakEnded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConferenceParticipantSpeakEnded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableConferenceParticipantSpeakEnded struct {
	value *ConferenceParticipantSpeakEnded
	isSet bool
}

func (v NullableConferenceParticipantSpeakEnded) Get() *ConferenceParticipantSpeakEnded {
	return v.value
}

func (v *NullableConferenceParticipantSpeakEnded) Set(val *ConferenceParticipantSpeakEnded) {
	v.value = val
	v.isSet = true
}

func (v NullableConferenceParticipantSpeakEnded) IsSet() bool {
	return v.isSet
}

func (v *NullableConferenceParticipantSpeakEnded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConferenceParticipantSpeakEnded(val *ConferenceParticipantSpeakEnded) *NullableConferenceParticipantSpeakEnded {
	return &NullableConferenceParticipantSpeakEnded{value: val, isSet: true}
}

func (v NullableConferenceParticipantSpeakEnded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConferenceParticipantSpeakEnded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


