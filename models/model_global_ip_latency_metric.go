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

// checks if the GlobalIpLatencyMetric type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalIpLatencyMetric{}

// GlobalIpLatencyMetric struct for GlobalIpLatencyMetric
type GlobalIpLatencyMetric struct {
	// The timestamp of the metric.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	GlobalIp *GlobalIpAssignmentHealthMetricGlobalIp `json:"global_ip,omitempty"`
	ProberLocation *GlobalIpLatencyMetricProberLocation `json:"prober_location,omitempty"`
	MeanLatency *GlobalIpLatencyMetricMeanLatency `json:"mean_latency,omitempty"`
	PercentileLatency *GlobalIpLatencyMetricPercentileLatency `json:"percentile_latency,omitempty"`
}

// NewGlobalIpLatencyMetric instantiates a new GlobalIpLatencyMetric object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalIpLatencyMetric() *GlobalIpLatencyMetric {
	this := GlobalIpLatencyMetric{}
	return &this
}

// NewGlobalIpLatencyMetricWithDefaults instantiates a new GlobalIpLatencyMetric object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalIpLatencyMetricWithDefaults() *GlobalIpLatencyMetric {
	this := GlobalIpLatencyMetric{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GlobalIpLatencyMetric) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIpLatencyMetric) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GlobalIpLatencyMetric) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *GlobalIpLatencyMetric) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetGlobalIp returns the GlobalIp field value if set, zero value otherwise.
func (o *GlobalIpLatencyMetric) GetGlobalIp() GlobalIpAssignmentHealthMetricGlobalIp {
	if o == nil || IsNil(o.GlobalIp) {
		var ret GlobalIpAssignmentHealthMetricGlobalIp
		return ret
	}
	return *o.GlobalIp
}

// GetGlobalIpOk returns a tuple with the GlobalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIpLatencyMetric) GetGlobalIpOk() (*GlobalIpAssignmentHealthMetricGlobalIp, bool) {
	if o == nil || IsNil(o.GlobalIp) {
		return nil, false
	}
	return o.GlobalIp, true
}

// HasGlobalIp returns a boolean if a field has been set.
func (o *GlobalIpLatencyMetric) HasGlobalIp() bool {
	if o != nil && !IsNil(o.GlobalIp) {
		return true
	}

	return false
}

// SetGlobalIp gets a reference to the given GlobalIpAssignmentHealthMetricGlobalIp and assigns it to the GlobalIp field.
func (o *GlobalIpLatencyMetric) SetGlobalIp(v GlobalIpAssignmentHealthMetricGlobalIp) {
	o.GlobalIp = &v
}

// GetProberLocation returns the ProberLocation field value if set, zero value otherwise.
func (o *GlobalIpLatencyMetric) GetProberLocation() GlobalIpLatencyMetricProberLocation {
	if o == nil || IsNil(o.ProberLocation) {
		var ret GlobalIpLatencyMetricProberLocation
		return ret
	}
	return *o.ProberLocation
}

// GetProberLocationOk returns a tuple with the ProberLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIpLatencyMetric) GetProberLocationOk() (*GlobalIpLatencyMetricProberLocation, bool) {
	if o == nil || IsNil(o.ProberLocation) {
		return nil, false
	}
	return o.ProberLocation, true
}

// HasProberLocation returns a boolean if a field has been set.
func (o *GlobalIpLatencyMetric) HasProberLocation() bool {
	if o != nil && !IsNil(o.ProberLocation) {
		return true
	}

	return false
}

// SetProberLocation gets a reference to the given GlobalIpLatencyMetricProberLocation and assigns it to the ProberLocation field.
func (o *GlobalIpLatencyMetric) SetProberLocation(v GlobalIpLatencyMetricProberLocation) {
	o.ProberLocation = &v
}

// GetMeanLatency returns the MeanLatency field value if set, zero value otherwise.
func (o *GlobalIpLatencyMetric) GetMeanLatency() GlobalIpLatencyMetricMeanLatency {
	if o == nil || IsNil(o.MeanLatency) {
		var ret GlobalIpLatencyMetricMeanLatency
		return ret
	}
	return *o.MeanLatency
}

// GetMeanLatencyOk returns a tuple with the MeanLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIpLatencyMetric) GetMeanLatencyOk() (*GlobalIpLatencyMetricMeanLatency, bool) {
	if o == nil || IsNil(o.MeanLatency) {
		return nil, false
	}
	return o.MeanLatency, true
}

// HasMeanLatency returns a boolean if a field has been set.
func (o *GlobalIpLatencyMetric) HasMeanLatency() bool {
	if o != nil && !IsNil(o.MeanLatency) {
		return true
	}

	return false
}

// SetMeanLatency gets a reference to the given GlobalIpLatencyMetricMeanLatency and assigns it to the MeanLatency field.
func (o *GlobalIpLatencyMetric) SetMeanLatency(v GlobalIpLatencyMetricMeanLatency) {
	o.MeanLatency = &v
}

// GetPercentileLatency returns the PercentileLatency field value if set, zero value otherwise.
func (o *GlobalIpLatencyMetric) GetPercentileLatency() GlobalIpLatencyMetricPercentileLatency {
	if o == nil || IsNil(o.PercentileLatency) {
		var ret GlobalIpLatencyMetricPercentileLatency
		return ret
	}
	return *o.PercentileLatency
}

// GetPercentileLatencyOk returns a tuple with the PercentileLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIpLatencyMetric) GetPercentileLatencyOk() (*GlobalIpLatencyMetricPercentileLatency, bool) {
	if o == nil || IsNil(o.PercentileLatency) {
		return nil, false
	}
	return o.PercentileLatency, true
}

// HasPercentileLatency returns a boolean if a field has been set.
func (o *GlobalIpLatencyMetric) HasPercentileLatency() bool {
	if o != nil && !IsNil(o.PercentileLatency) {
		return true
	}

	return false
}

// SetPercentileLatency gets a reference to the given GlobalIpLatencyMetricPercentileLatency and assigns it to the PercentileLatency field.
func (o *GlobalIpLatencyMetric) SetPercentileLatency(v GlobalIpLatencyMetricPercentileLatency) {
	o.PercentileLatency = &v
}

func (o GlobalIpLatencyMetric) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalIpLatencyMetric) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.GlobalIp) {
		toSerialize["global_ip"] = o.GlobalIp
	}
	if !IsNil(o.ProberLocation) {
		toSerialize["prober_location"] = o.ProberLocation
	}
	if !IsNil(o.MeanLatency) {
		toSerialize["mean_latency"] = o.MeanLatency
	}
	if !IsNil(o.PercentileLatency) {
		toSerialize["percentile_latency"] = o.PercentileLatency
	}
	return toSerialize, nil
}

type NullableGlobalIpLatencyMetric struct {
	value *GlobalIpLatencyMetric
	isSet bool
}

func (v NullableGlobalIpLatencyMetric) Get() *GlobalIpLatencyMetric {
	return v.value
}

func (v *NullableGlobalIpLatencyMetric) Set(val *GlobalIpLatencyMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalIpLatencyMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalIpLatencyMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalIpLatencyMetric(val *GlobalIpLatencyMetric) *NullableGlobalIpLatencyMetric {
	return &NullableGlobalIpLatencyMetric{value: val, isSet: true}
}

func (v NullableGlobalIpLatencyMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalIpLatencyMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


