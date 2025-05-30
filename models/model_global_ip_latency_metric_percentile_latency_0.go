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

// checks if the GlobalIpLatencyMetricPercentileLatency0 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GlobalIpLatencyMetricPercentileLatency0{}

// GlobalIpLatencyMetricPercentileLatency0 struct for GlobalIpLatencyMetricPercentileLatency0
type GlobalIpLatencyMetricPercentileLatency0 struct {
	// The minimum latency.
	Amount *float32 `json:"amount,omitempty"`
	// The unit of the minimum latency.
	Unit *string `json:"unit,omitempty"`
}

// NewGlobalIpLatencyMetricPercentileLatency0 instantiates a new GlobalIpLatencyMetricPercentileLatency0 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGlobalIpLatencyMetricPercentileLatency0() *GlobalIpLatencyMetricPercentileLatency0 {
	this := GlobalIpLatencyMetricPercentileLatency0{}
	return &this
}

// NewGlobalIpLatencyMetricPercentileLatency0WithDefaults instantiates a new GlobalIpLatencyMetricPercentileLatency0 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGlobalIpLatencyMetricPercentileLatency0WithDefaults() *GlobalIpLatencyMetricPercentileLatency0 {
	this := GlobalIpLatencyMetricPercentileLatency0{}
	return &this
}

// GetAmount returns the Amount field value if set, zero value otherwise.
func (o *GlobalIpLatencyMetricPercentileLatency0) GetAmount() float32 {
	if o == nil || IsNil(o.Amount) {
		var ret float32
		return ret
	}
	return *o.Amount
}

// GetAmountOk returns a tuple with the Amount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIpLatencyMetricPercentileLatency0) GetAmountOk() (*float32, bool) {
	if o == nil || IsNil(o.Amount) {
		return nil, false
	}
	return o.Amount, true
}

// HasAmount returns a boolean if a field has been set.
func (o *GlobalIpLatencyMetricPercentileLatency0) HasAmount() bool {
	if o != nil && !IsNil(o.Amount) {
		return true
	}

	return false
}

// SetAmount gets a reference to the given float32 and assigns it to the Amount field.
func (o *GlobalIpLatencyMetricPercentileLatency0) SetAmount(v float32) {
	o.Amount = &v
}

// GetUnit returns the Unit field value if set, zero value otherwise.
func (o *GlobalIpLatencyMetricPercentileLatency0) GetUnit() string {
	if o == nil || IsNil(o.Unit) {
		var ret string
		return ret
	}
	return *o.Unit
}

// GetUnitOk returns a tuple with the Unit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GlobalIpLatencyMetricPercentileLatency0) GetUnitOk() (*string, bool) {
	if o == nil || IsNil(o.Unit) {
		return nil, false
	}
	return o.Unit, true
}

// HasUnit returns a boolean if a field has been set.
func (o *GlobalIpLatencyMetricPercentileLatency0) HasUnit() bool {
	if o != nil && !IsNil(o.Unit) {
		return true
	}

	return false
}

// SetUnit gets a reference to the given string and assigns it to the Unit field.
func (o *GlobalIpLatencyMetricPercentileLatency0) SetUnit(v string) {
	o.Unit = &v
}

func (o GlobalIpLatencyMetricPercentileLatency0) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GlobalIpLatencyMetricPercentileLatency0) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Amount) {
		toSerialize["amount"] = o.Amount
	}
	if !IsNil(o.Unit) {
		toSerialize["unit"] = o.Unit
	}
	return toSerialize, nil
}

type NullableGlobalIpLatencyMetricPercentileLatency0 struct {
	value *GlobalIpLatencyMetricPercentileLatency0
	isSet bool
}

func (v NullableGlobalIpLatencyMetricPercentileLatency0) Get() *GlobalIpLatencyMetricPercentileLatency0 {
	return v.value
}

func (v *NullableGlobalIpLatencyMetricPercentileLatency0) Set(val *GlobalIpLatencyMetricPercentileLatency0) {
	v.value = val
	v.isSet = true
}

func (v NullableGlobalIpLatencyMetricPercentileLatency0) IsSet() bool {
	return v.isSet
}

func (v *NullableGlobalIpLatencyMetricPercentileLatency0) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGlobalIpLatencyMetricPercentileLatency0(val *GlobalIpLatencyMetricPercentileLatency0) *NullableGlobalIpLatencyMetricPercentileLatency0 {
	return &NullableGlobalIpLatencyMetricPercentileLatency0{value: val, isSet: true}
}

func (v NullableGlobalIpLatencyMetricPercentileLatency0) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGlobalIpLatencyMetricPercentileLatency0) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


