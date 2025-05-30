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

// checks if the SIMCardActionStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SIMCardActionStatus{}

// SIMCardActionStatus struct for SIMCardActionStatus
type SIMCardActionStatus struct {
	// The current status of the SIM card action.
	Value *string `json:"value,omitempty"`
	// It describes why the SIM card action is in the current status. This will be <code>null</code> for self-explanatory statuses, such as <code>in-progress</code> and <code>completed</code> but will include further information on statuses like <code>interrupted</code> and <code>failed</code>.
	Reason *string `json:"reason,omitempty"`
}

// NewSIMCardActionStatus instantiates a new SIMCardActionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSIMCardActionStatus() *SIMCardActionStatus {
	this := SIMCardActionStatus{}
	return &this
}

// NewSIMCardActionStatusWithDefaults instantiates a new SIMCardActionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSIMCardActionStatusWithDefaults() *SIMCardActionStatus {
	this := SIMCardActionStatus{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SIMCardActionStatus) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardActionStatus) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SIMCardActionStatus) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SIMCardActionStatus) SetValue(v string) {
	o.Value = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SIMCardActionStatus) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardActionStatus) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SIMCardActionStatus) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SIMCardActionStatus) SetReason(v string) {
	o.Reason = &v
}

func (o SIMCardActionStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SIMCardActionStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableSIMCardActionStatus struct {
	value *SIMCardActionStatus
	isSet bool
}

func (v NullableSIMCardActionStatus) Get() *SIMCardActionStatus {
	return v.value
}

func (v *NullableSIMCardActionStatus) Set(val *SIMCardActionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSIMCardActionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSIMCardActionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSIMCardActionStatus(val *SIMCardActionStatus) *NullableSIMCardActionStatus {
	return &NullableSIMCardActionStatus{value: val, isSet: true}
}

func (v NullableSIMCardActionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSIMCardActionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


