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

// checks if the CallLeftQueueEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallLeftQueueEvent{}

// CallLeftQueueEvent struct for CallLeftQueueEvent
type CallLeftQueueEvent struct {
	Data *CallLeftQueue `json:"data,omitempty"`
}

// NewCallLeftQueueEvent instantiates a new CallLeftQueueEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallLeftQueueEvent() *CallLeftQueueEvent {
	this := CallLeftQueueEvent{}
	return &this
}

// NewCallLeftQueueEventWithDefaults instantiates a new CallLeftQueueEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallLeftQueueEventWithDefaults() *CallLeftQueueEvent {
	this := CallLeftQueueEvent{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CallLeftQueueEvent) GetData() CallLeftQueue {
	if o == nil || IsNil(o.Data) {
		var ret CallLeftQueue
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallLeftQueueEvent) GetDataOk() (*CallLeftQueue, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CallLeftQueueEvent) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CallLeftQueue and assigns it to the Data field.
func (o *CallLeftQueueEvent) SetData(v CallLeftQueue) {
	o.Data = &v
}

func (o CallLeftQueueEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallLeftQueueEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCallLeftQueueEvent struct {
	value *CallLeftQueueEvent
	isSet bool
}

func (v NullableCallLeftQueueEvent) Get() *CallLeftQueueEvent {
	return v.value
}

func (v *NullableCallLeftQueueEvent) Set(val *CallLeftQueueEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableCallLeftQueueEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableCallLeftQueueEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallLeftQueueEvent(val *CallLeftQueueEvent) *NullableCallLeftQueueEvent {
	return &NullableCallLeftQueueEvent{value: val, isSet: true}
}

func (v NullableCallLeftQueueEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallLeftQueueEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


