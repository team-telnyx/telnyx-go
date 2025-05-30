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

// checks if the PortoutEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortoutEvent{}

// PortoutEvent struct for PortoutEvent
type PortoutEvent struct {
	// Uniquely identifies the event.
	Id *string `json:"id,omitempty"`
	// Identifies the event type
	EventType *string `json:"event_type,omitempty"`
	// Identifies the port-out order associated with the event.
	PortoutId *string `json:"portout_id,omitempty"`
	// Indicates the notification methods used.
	AvailableNotificationMethods []string `json:"available_notification_methods,omitempty"`
	// The status of the payload generation.
	PayloadStatus *string `json:"payload_status,omitempty"`
	Payload *PortoutEventPayload `json:"payload,omitempty"`
	// Identifies the type of the resource.
	RecordType *string `json:"record_type,omitempty"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// NewPortoutEvent instantiates a new PortoutEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortoutEvent() *PortoutEvent {
	this := PortoutEvent{}
	return &this
}

// NewPortoutEventWithDefaults instantiates a new PortoutEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortoutEventWithDefaults() *PortoutEvent {
	this := PortoutEvent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PortoutEvent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortoutEvent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PortoutEvent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PortoutEvent) SetId(v string) {
	o.Id = &v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *PortoutEvent) GetEventType() string {
	if o == nil || IsNil(o.EventType) {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortoutEvent) GetEventTypeOk() (*string, bool) {
	if o == nil || IsNil(o.EventType) {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *PortoutEvent) HasEventType() bool {
	if o != nil && !IsNil(o.EventType) {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *PortoutEvent) SetEventType(v string) {
	o.EventType = &v
}

// GetPortoutId returns the PortoutId field value if set, zero value otherwise.
func (o *PortoutEvent) GetPortoutId() string {
	if o == nil || IsNil(o.PortoutId) {
		var ret string
		return ret
	}
	return *o.PortoutId
}

// GetPortoutIdOk returns a tuple with the PortoutId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortoutEvent) GetPortoutIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortoutId) {
		return nil, false
	}
	return o.PortoutId, true
}

// HasPortoutId returns a boolean if a field has been set.
func (o *PortoutEvent) HasPortoutId() bool {
	if o != nil && !IsNil(o.PortoutId) {
		return true
	}

	return false
}

// SetPortoutId gets a reference to the given string and assigns it to the PortoutId field.
func (o *PortoutEvent) SetPortoutId(v string) {
	o.PortoutId = &v
}

// GetAvailableNotificationMethods returns the AvailableNotificationMethods field value if set, zero value otherwise.
func (o *PortoutEvent) GetAvailableNotificationMethods() []string {
	if o == nil || IsNil(o.AvailableNotificationMethods) {
		var ret []string
		return ret
	}
	return o.AvailableNotificationMethods
}

// GetAvailableNotificationMethodsOk returns a tuple with the AvailableNotificationMethods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortoutEvent) GetAvailableNotificationMethodsOk() ([]string, bool) {
	if o == nil || IsNil(o.AvailableNotificationMethods) {
		return nil, false
	}
	return o.AvailableNotificationMethods, true
}

// HasAvailableNotificationMethods returns a boolean if a field has been set.
func (o *PortoutEvent) HasAvailableNotificationMethods() bool {
	if o != nil && !IsNil(o.AvailableNotificationMethods) {
		return true
	}

	return false
}

// SetAvailableNotificationMethods gets a reference to the given []string and assigns it to the AvailableNotificationMethods field.
func (o *PortoutEvent) SetAvailableNotificationMethods(v []string) {
	o.AvailableNotificationMethods = v
}

// GetPayloadStatus returns the PayloadStatus field value if set, zero value otherwise.
func (o *PortoutEvent) GetPayloadStatus() string {
	if o == nil || IsNil(o.PayloadStatus) {
		var ret string
		return ret
	}
	return *o.PayloadStatus
}

// GetPayloadStatusOk returns a tuple with the PayloadStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortoutEvent) GetPayloadStatusOk() (*string, bool) {
	if o == nil || IsNil(o.PayloadStatus) {
		return nil, false
	}
	return o.PayloadStatus, true
}

// HasPayloadStatus returns a boolean if a field has been set.
func (o *PortoutEvent) HasPayloadStatus() bool {
	if o != nil && !IsNil(o.PayloadStatus) {
		return true
	}

	return false
}

// SetPayloadStatus gets a reference to the given string and assigns it to the PayloadStatus field.
func (o *PortoutEvent) SetPayloadStatus(v string) {
	o.PayloadStatus = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *PortoutEvent) GetPayload() PortoutEventPayload {
	if o == nil || IsNil(o.Payload) {
		var ret PortoutEventPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortoutEvent) GetPayloadOk() (*PortoutEventPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *PortoutEvent) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given PortoutEventPayload and assigns it to the Payload field.
func (o *PortoutEvent) SetPayload(v PortoutEventPayload) {
	o.Payload = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *PortoutEvent) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortoutEvent) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *PortoutEvent) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *PortoutEvent) SetRecordType(v string) {
	o.RecordType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PortoutEvent) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortoutEvent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PortoutEvent) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PortoutEvent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PortoutEvent) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortoutEvent) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PortoutEvent) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PortoutEvent) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

func (o PortoutEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortoutEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.EventType) {
		toSerialize["event_type"] = o.EventType
	}
	if !IsNil(o.PortoutId) {
		toSerialize["portout_id"] = o.PortoutId
	}
	if !IsNil(o.AvailableNotificationMethods) {
		toSerialize["available_notification_methods"] = o.AvailableNotificationMethods
	}
	if !IsNil(o.PayloadStatus) {
		toSerialize["payload_status"] = o.PayloadStatus
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullablePortoutEvent struct {
	value *PortoutEvent
	isSet bool
}

func (v NullablePortoutEvent) Get() *PortoutEvent {
	return v.value
}

func (v *NullablePortoutEvent) Set(val *PortoutEvent) {
	v.value = val
	v.isSet = true
}

func (v NullablePortoutEvent) IsSet() bool {
	return v.isSet
}

func (v *NullablePortoutEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortoutEvent(val *PortoutEvent) *NullablePortoutEvent {
	return &NullablePortoutEvent{value: val, isSet: true}
}

func (v NullablePortoutEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortoutEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


