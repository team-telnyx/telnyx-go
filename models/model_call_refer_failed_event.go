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

// checks if the CallReferFailedEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallReferFailedEvent{}

// CallReferFailedEvent struct for CallReferFailedEvent
type CallReferFailedEvent struct {
	Data *CallReferFailed `json:"data,omitempty"`
}

// NewCallReferFailedEvent instantiates a new CallReferFailedEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallReferFailedEvent() *CallReferFailedEvent {
	this := CallReferFailedEvent{}
	return &this
}

// NewCallReferFailedEventWithDefaults instantiates a new CallReferFailedEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallReferFailedEventWithDefaults() *CallReferFailedEvent {
	this := CallReferFailedEvent{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CallReferFailedEvent) GetData() CallReferFailed {
	if o == nil || IsNil(o.Data) {
		var ret CallReferFailed
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallReferFailedEvent) GetDataOk() (*CallReferFailed, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CallReferFailedEvent) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CallReferFailed and assigns it to the Data field.
func (o *CallReferFailedEvent) SetData(v CallReferFailed) {
	o.Data = &v
}

func (o CallReferFailedEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallReferFailedEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCallReferFailedEvent struct {
	value *CallReferFailedEvent
	isSet bool
}

func (v NullableCallReferFailedEvent) Get() *CallReferFailedEvent {
	return v.value
}

func (v *NullableCallReferFailedEvent) Set(val *CallReferFailedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableCallReferFailedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableCallReferFailedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallReferFailedEvent(val *CallReferFailedEvent) *NullableCallReferFailedEvent {
	return &NullableCallReferFailedEvent{value: val, isSet: true}
}

func (v NullableCallReferFailedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallReferFailedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


