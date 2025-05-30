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
	"time"
)

// checks if the CallRecordingError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallRecordingError{}

// CallRecordingError struct for CallRecordingError
type CallRecordingError struct {
	// Identifies the type of the resource.
	RecordType *string `json:"record_type,omitempty"`
	// The type of event being delivered.
	EventType *string `json:"event_type,omitempty"`
	// Identifies the type of resource.
	Id *string `json:"id,omitempty"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt *time.Time `json:"occurred_at,omitempty"`
	Payload *CallRecordingErrorPayload `json:"payload,omitempty"`
}

// NewCallRecordingError instantiates a new CallRecordingError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallRecordingError() *CallRecordingError {
	this := CallRecordingError{}
	return &this
}

// NewCallRecordingErrorWithDefaults instantiates a new CallRecordingError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallRecordingErrorWithDefaults() *CallRecordingError {
	this := CallRecordingError{}
	return &this
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *CallRecordingError) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRecordingError) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *CallRecordingError) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *CallRecordingError) SetRecordType(v string) {
	o.RecordType = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CallRecordingError) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRecordingError) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CallRecordingError) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CallRecordingError) SetEventType(v string) {
	o.EventType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CallRecordingError) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRecordingError) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CallRecordingError) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CallRecordingError) SetId(v string) {
	o.Id = &v
}

// GetOccurredAt returns the OccurredAt field value if set, zero value otherwise.
func (o *CallRecordingError) GetOccurredAt() time.Time {
	if o == nil || IsNil(o.OccurredAt) {
		var ret time.Time
		return ret
	}
	return *o.OccurredAt
}

// GetOccurredAtOk returns a tuple with the OccurredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRecordingError) GetOccurredAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OccurredAt) {
		return nil, false
	}
	return o.OccurredAt, true
}

// HasOccurredAt returns a boolean if a field has been set.
func (o *CallRecordingError) HasOccurredAt() bool {
	if o != nil && !IsNil(o.OccurredAt) {
		return true
	}

	return false
}

// SetOccurredAt gets a reference to the given time.Time and assigns it to the OccurredAt field.
func (o *CallRecordingError) SetOccurredAt(v time.Time) {
	o.OccurredAt = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CallRecordingError) GetPayload() CallRecordingErrorPayload {
	if o == nil || IsNil(o.Payload) {
		var ret CallRecordingErrorPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRecordingError) GetPayloadOk() (*CallRecordingErrorPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *CallRecordingError) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given CallRecordingErrorPayload and assigns it to the Payload field.
func (o *CallRecordingError) SetPayload(v CallRecordingErrorPayload) {
	o.Payload = &v
}

func (o CallRecordingError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallRecordingError) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.OccurredAt) {
		toSerialize["occurred_at"] = o.OccurredAt
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	return toSerialize, nil
}

type NullableCallRecordingError struct {
	value *CallRecordingError
	isSet bool
}

func (v NullableCallRecordingError) Get() *CallRecordingError {
	return v.value
}

func (v *NullableCallRecordingError) Set(val *CallRecordingError) {
	v.value = val
	v.isSet = true
}

func (v NullableCallRecordingError) IsSet() bool {
	return v.isSet
}

func (v *NullableCallRecordingError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallRecordingError(val *CallRecordingError) *NullableCallRecordingError {
	return &NullableCallRecordingError{value: val, isSet: true}
}

func (v NullableCallRecordingError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallRecordingError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


