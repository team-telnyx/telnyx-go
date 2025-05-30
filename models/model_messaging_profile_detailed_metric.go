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

// checks if the MessagingProfileDetailedMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MessagingProfileDetailedMetric{}

// MessagingProfileDetailedMetric struct for MessagingProfileDetailedMetric
type MessagingProfileDetailedMetric struct {
	// The timestamp of the aggregated data.
	Timestamp *string `json:"timestamp,omitempty"`
	Metrics []MessagingProfileMessageTypeMetrics `json:"metrics,omitempty"`
}

// NewMessagingProfileDetailedMetric instantiates a new MessagingProfileDetailedMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessagingProfileDetailedMetric() *MessagingProfileDetailedMetric {
	this := MessagingProfileDetailedMetric{}
	return &this
}

// NewMessagingProfileDetailedMetricWithDefaults instantiates a new MessagingProfileDetailedMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessagingProfileDetailedMetricWithDefaults() *MessagingProfileDetailedMetric {
	this := MessagingProfileDetailedMetric{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *MessagingProfileDetailedMetric) GetTimestamp() string {
	if o == nil || IsNil(o.Timestamp) {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagingProfileDetailedMetric) GetTimestampOk() (*string, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *MessagingProfileDetailedMetric) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *MessagingProfileDetailedMetric) SetTimestamp(v string) {
	o.Timestamp = &v
}

// GetMetrics returns the Metrics field value if set, zero value otherwise.
func (o *MessagingProfileDetailedMetric) GetMetrics() []MessagingProfileMessageTypeMetrics {
	if o == nil || IsNil(o.Metrics) {
		var ret []MessagingProfileMessageTypeMetrics
		return ret
	}
	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MessagingProfileDetailedMetric) GetMetricsOk() ([]MessagingProfileMessageTypeMetrics, bool) {
	if o == nil || IsNil(o.Metrics) {
		return nil, false
	}
	return o.Metrics, true
}

// HasMetrics returns a boolean if a field has been set.
func (o *MessagingProfileDetailedMetric) HasMetrics() bool {
	if o != nil && !IsNil(o.Metrics) {
		return true
	}

	return false
}

// SetMetrics gets a reference to the given []MessagingProfileMessageTypeMetrics and assigns it to the Metrics field.
func (o *MessagingProfileDetailedMetric) SetMetrics(v []MessagingProfileMessageTypeMetrics) {
	o.Metrics = v
}

func (o MessagingProfileDetailedMetric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MessagingProfileDetailedMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Metrics) {
		toSerialize["metrics"] = o.Metrics
	}
	return toSerialize, nil
}

type NullableMessagingProfileDetailedMetric struct {
	value *MessagingProfileDetailedMetric
	isSet bool
}

func (v NullableMessagingProfileDetailedMetric) Get() *MessagingProfileDetailedMetric {
	return v.value
}

func (v *NullableMessagingProfileDetailedMetric) Set(val *MessagingProfileDetailedMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableMessagingProfileDetailedMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableMessagingProfileDetailedMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessagingProfileDetailedMetric(val *MessagingProfileDetailedMetric) *NullableMessagingProfileDetailedMetric {
	return &NullableMessagingProfileDetailedMetric{value: val, isSet: true}
}

func (v NullableMessagingProfileDetailedMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessagingProfileDetailedMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


