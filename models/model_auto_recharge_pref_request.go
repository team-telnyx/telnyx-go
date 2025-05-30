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

// checks if the AutoRechargePrefRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoRechargePrefRequest{}

// AutoRechargePrefRequest struct for AutoRechargePrefRequest
type AutoRechargePrefRequest struct {
	// The threshold amount at which the account will be recharged.
	ThresholdAmount *float64 `json:"threshold_amount,omitempty"`
	// The amount to recharge the account, the actual recharge amount will be the amount necessary to reach the threshold amount plus the recharge amount.
	RechargeAmount *float64 `json:"recharge_amount,omitempty"`
	// Whether auto recharge is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	InvoiceEnabled *bool `json:"invoice_enabled,omitempty"`
	// The payment preference for auto recharge.
	Preference *string `json:"preference,omitempty"`
}

// NewAutoRechargePrefRequest instantiates a new AutoRechargePrefRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoRechargePrefRequest() *AutoRechargePrefRequest {
	this := AutoRechargePrefRequest{}
	return &this
}

// NewAutoRechargePrefRequestWithDefaults instantiates a new AutoRechargePrefRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoRechargePrefRequestWithDefaults() *AutoRechargePrefRequest {
	this := AutoRechargePrefRequest{}
	return &this
}

// GetThresholdAmount returns the ThresholdAmount field value if set, zero value otherwise.
func (o *AutoRechargePrefRequest) GetThresholdAmount() float64 {
	if o == nil || IsNil(o.ThresholdAmount) {
		var ret float64
		return ret
	}
	return *o.ThresholdAmount
}

// GetThresholdAmountOk returns a tuple with the ThresholdAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoRechargePrefRequest) GetThresholdAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.ThresholdAmount) {
		return nil, false
	}
	return o.ThresholdAmount, true
}

// HasThresholdAmount returns a boolean if a field has been set.
func (o *AutoRechargePrefRequest) HasThresholdAmount() bool {
	if o != nil && !IsNil(o.ThresholdAmount) {
		return true
	}

	return false
}

// SetThresholdAmount gets a reference to the given float64 and assigns it to the ThresholdAmount field.
func (o *AutoRechargePrefRequest) SetThresholdAmount(v float64) {
	o.ThresholdAmount = &v
}

// GetRechargeAmount returns the RechargeAmount field value if set, zero value otherwise.
func (o *AutoRechargePrefRequest) GetRechargeAmount() float64 {
	if o == nil || IsNil(o.RechargeAmount) {
		var ret float64
		return ret
	}
	return *o.RechargeAmount
}

// GetRechargeAmountOk returns a tuple with the RechargeAmount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoRechargePrefRequest) GetRechargeAmountOk() (*float64, bool) {
	if o == nil || IsNil(o.RechargeAmount) {
		return nil, false
	}
	return o.RechargeAmount, true
}

// HasRechargeAmount returns a boolean if a field has been set.
func (o *AutoRechargePrefRequest) HasRechargeAmount() bool {
	if o != nil && !IsNil(o.RechargeAmount) {
		return true
	}

	return false
}

// SetRechargeAmount gets a reference to the given float64 and assigns it to the RechargeAmount field.
func (o *AutoRechargePrefRequest) SetRechargeAmount(v float64) {
	o.RechargeAmount = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *AutoRechargePrefRequest) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoRechargePrefRequest) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *AutoRechargePrefRequest) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *AutoRechargePrefRequest) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInvoiceEnabled returns the InvoiceEnabled field value if set, zero value otherwise.
func (o *AutoRechargePrefRequest) GetInvoiceEnabled() bool {
	if o == nil || IsNil(o.InvoiceEnabled) {
		var ret bool
		return ret
	}
	return *o.InvoiceEnabled
}

// GetInvoiceEnabledOk returns a tuple with the InvoiceEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoRechargePrefRequest) GetInvoiceEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.InvoiceEnabled) {
		return nil, false
	}
	return o.InvoiceEnabled, true
}

// HasInvoiceEnabled returns a boolean if a field has been set.
func (o *AutoRechargePrefRequest) HasInvoiceEnabled() bool {
	if o != nil && !IsNil(o.InvoiceEnabled) {
		return true
	}

	return false
}

// SetInvoiceEnabled gets a reference to the given bool and assigns it to the InvoiceEnabled field.
func (o *AutoRechargePrefRequest) SetInvoiceEnabled(v bool) {
	o.InvoiceEnabled = &v
}

// GetPreference returns the Preference field value if set, zero value otherwise.
func (o *AutoRechargePrefRequest) GetPreference() string {
	if o == nil || IsNil(o.Preference) {
		var ret string
		return ret
	}
	return *o.Preference
}

// GetPreferenceOk returns a tuple with the Preference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoRechargePrefRequest) GetPreferenceOk() (*string, bool) {
	if o == nil || IsNil(o.Preference) {
		return nil, false
	}
	return o.Preference, true
}

// HasPreference returns a boolean if a field has been set.
func (o *AutoRechargePrefRequest) HasPreference() bool {
	if o != nil && !IsNil(o.Preference) {
		return true
	}

	return false
}

// SetPreference gets a reference to the given string and assigns it to the Preference field.
func (o *AutoRechargePrefRequest) SetPreference(v string) {
	o.Preference = &v
}

func (o AutoRechargePrefRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoRechargePrefRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ThresholdAmount) {
		toSerialize["threshold_amount"] = o.ThresholdAmount
	}
	if !IsNil(o.RechargeAmount) {
		toSerialize["recharge_amount"] = o.RechargeAmount
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.InvoiceEnabled) {
		toSerialize["invoice_enabled"] = o.InvoiceEnabled
	}
	if !IsNil(o.Preference) {
		toSerialize["preference"] = o.Preference
	}
	return toSerialize, nil
}

type NullableAutoRechargePrefRequest struct {
	value *AutoRechargePrefRequest
	isSet bool
}

func (v NullableAutoRechargePrefRequest) Get() *AutoRechargePrefRequest {
	return v.value
}

func (v *NullableAutoRechargePrefRequest) Set(val *AutoRechargePrefRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoRechargePrefRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoRechargePrefRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoRechargePrefRequest(val *AutoRechargePrefRequest) *NullableAutoRechargePrefRequest {
	return &NullableAutoRechargePrefRequest{value: val, isSet: true}
}

func (v NullableAutoRechargePrefRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoRechargePrefRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


