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

// checks if the CallBridged type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallBridged{}

// CallBridged struct for CallBridged
type CallBridged struct {
	// Identifies the type of the resource.
	RecordType *string `json:"record_type,omitempty"`
	// The type of event being delivered.
	EventType *string `json:"event_type,omitempty"`
	// Identifies the type of resource.
	Id *string `json:"id,omitempty"`
	// ISO 8601 datetime of when the event occurred.
	OccurredAt *time.Time `json:"occurred_at,omitempty"`
	Payload *CallBridgedPayload `json:"payload,omitempty"`
}

// NewCallBridged instantiates a new CallBridged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallBridged() *CallBridged {
	this := CallBridged{}
	return &this
}

// NewCallBridgedWithDefaults instantiates a new CallBridged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallBridgedWithDefaults() *CallBridged {
	this := CallBridged{}
	return &this
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *CallBridged) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallBridged) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *CallBridged) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *CallBridged) SetRecordType(v string) {
	o.RecordType = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CallBridged) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallBridged) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CallBridged) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CallBridged) SetEventType(v string) {
	o.EventType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CallBridged) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallBridged) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CallBridged) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CallBridged) SetId(v string) {
	o.Id = &v
}

// GetOccurredAt returns the OccurredAt field value if set, zero value otherwise.
func (o *CallBridged) GetOccurredAt() time.Time {
	if o == nil || IsNil(o.OccurredAt) {
		var ret time.Time
		return ret
	}
	return *o.OccurredAt
}

// GetOccurredAtOk returns a tuple with the OccurredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallBridged) GetOccurredAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OccurredAt) {
		return nil, false
	}
	return o.OccurredAt, true
}

// HasOccurredAt returns a boolean if a field has been set.
func (o *CallBridged) HasOccurredAt() bool {
	if o != nil && !IsNil(o.OccurredAt) {
		return true
	}

	return false
}

// SetOccurredAt gets a reference to the given time.Time and assigns it to the OccurredAt field.
func (o *CallBridged) SetOccurredAt(v time.Time) {
	o.OccurredAt = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CallBridged) GetPayload() CallBridgedPayload {
	if o == nil || IsNil(o.Payload) {
		var ret CallBridgedPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallBridged) GetPayloadOk() (*CallBridgedPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *CallBridged) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given CallBridgedPayload and assigns it to the Payload field.
func (o *CallBridged) SetPayload(v CallBridgedPayload) {
	o.Payload = &v
}

func (o CallBridged) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallBridged) ToMap() (map[string]interface{}, error) {
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

type NullableCallBridged struct {
	value *CallBridged
	isSet bool
}

func (v NullableCallBridged) Get() *CallBridged {
	return v.value
}

func (v *NullableCallBridged) Set(val *CallBridged) {
	v.value = val
	v.isSet = true
}

func (v NullableCallBridged) IsSet() bool {
	return v.isSet
}

func (v *NullableCallBridged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallBridged(val *CallBridged) *NullableCallBridged {
	return &NullableCallBridged{value: val, isSet: true}
}

func (v NullableCallBridged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallBridged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


