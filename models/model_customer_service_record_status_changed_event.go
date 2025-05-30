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

// checks if the CustomerServiceRecordStatusChangedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerServiceRecordStatusChangedEvent{}

// CustomerServiceRecordStatusChangedEvent struct for CustomerServiceRecordStatusChangedEvent
type CustomerServiceRecordStatusChangedEvent struct {
	// Uniquely identifies the callback event.
	Id *string `json:"id,omitempty"`
	// The type of the callback event.
	EventType *string `json:"event_type,omitempty"`
	Payload *CustomerServiceRecordStatusChangedEventPayload `json:"payload,omitempty"`
	// ISO 8601 formatted date indicating when the callback event occurred.
	OccurredAt *time.Time `json:"occurred_at,omitempty"`
	// Identifies the type of the resource.
	RecordType *string `json:"record_type,omitempty"`
}

// NewCustomerServiceRecordStatusChangedEvent instantiates a new CustomerServiceRecordStatusChangedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerServiceRecordStatusChangedEvent() *CustomerServiceRecordStatusChangedEvent {
	this := CustomerServiceRecordStatusChangedEvent{}
	return &this
}

// NewCustomerServiceRecordStatusChangedEventWithDefaults instantiates a new CustomerServiceRecordStatusChangedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerServiceRecordStatusChangedEventWithDefaults() *CustomerServiceRecordStatusChangedEvent {
	this := CustomerServiceRecordStatusChangedEvent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CustomerServiceRecordStatusChangedEvent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordStatusChangedEvent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CustomerServiceRecordStatusChangedEvent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CustomerServiceRecordStatusChangedEvent) SetId(v string) {
	o.Id = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *CustomerServiceRecordStatusChangedEvent) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordStatusChangedEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *CustomerServiceRecordStatusChangedEvent) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *CustomerServiceRecordStatusChangedEvent) SetEventType(v string) {
	o.EventType = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *CustomerServiceRecordStatusChangedEvent) GetPayload() CustomerServiceRecordStatusChangedEventPayload {
	if o == nil || IsNil(o.Payload) {
		var ret CustomerServiceRecordStatusChangedEventPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordStatusChangedEvent) GetPayloadOk() (*CustomerServiceRecordStatusChangedEventPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *CustomerServiceRecordStatusChangedEvent) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given CustomerServiceRecordStatusChangedEventPayload and assigns it to the Payload field.
func (o *CustomerServiceRecordStatusChangedEvent) SetPayload(v CustomerServiceRecordStatusChangedEventPayload) {
	o.Payload = &v
}

// GetOccurredAt returns the OccurredAt field value if set, zero value otherwise.
func (o *CustomerServiceRecordStatusChangedEvent) GetOccurredAt() time.Time {
	if o == nil || IsNil(o.OccurredAt) {
		var ret time.Time
		return ret
	}
	return *o.OccurredAt
}

// GetOccurredAtOk returns a tuple with the OccurredAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordStatusChangedEvent) GetOccurredAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.OccurredAt) {
		return nil, false
	}
	return o.OccurredAt, true
}

// HasOccurredAt returns a boolean if a field has been set.
func (o *CustomerServiceRecordStatusChangedEvent) HasOccurredAt() bool {
	if o != nil && !IsNil(o.OccurredAt) {
		return true
	}

	return false
}

// SetOccurredAt gets a reference to the given time.Time and assigns it to the OccurredAt field.
func (o *CustomerServiceRecordStatusChangedEvent) SetOccurredAt(v time.Time) {
	o.OccurredAt = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *CustomerServiceRecordStatusChangedEvent) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerServiceRecordStatusChangedEvent) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *CustomerServiceRecordStatusChangedEvent) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *CustomerServiceRecordStatusChangedEvent) SetRecordType(v string) {
	o.RecordType = &v
}

func (o CustomerServiceRecordStatusChangedEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerServiceRecordStatusChangedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.OccurredAt) {
		toSerialize["occurred_at"] = o.OccurredAt
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	return toSerialize, nil
}

type NullableCustomerServiceRecordStatusChangedEvent struct {
	value *CustomerServiceRecordStatusChangedEvent
	isSet bool
}

func (v NullableCustomerServiceRecordStatusChangedEvent) Get() *CustomerServiceRecordStatusChangedEvent {
	return v.value
}

func (v *NullableCustomerServiceRecordStatusChangedEvent) Set(val *CustomerServiceRecordStatusChangedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerServiceRecordStatusChangedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerServiceRecordStatusChangedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerServiceRecordStatusChangedEvent(val *CustomerServiceRecordStatusChangedEvent) *NullableCustomerServiceRecordStatusChangedEvent {
	return &NullableCustomerServiceRecordStatusChangedEvent{value: val, isSet: true}
}

func (v NullableCustomerServiceRecordStatusChangedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerServiceRecordStatusChangedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


