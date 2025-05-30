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

// checks if the CountryCoverageLocal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CountryCoverageLocal{}

// CountryCoverageLocal struct for CountryCoverageLocal
type CountryCoverageLocal struct {
	Features []string `json:"features,omitempty"`
	Reservable *bool `json:"reservable,omitempty"`
	Quickship *bool `json:"quickship,omitempty"`
	InternationalSms *bool `json:"international_sms,omitempty"`
	P2p *bool `json:"p2p,omitempty"`
}

// NewCountryCoverageLocal instantiates a new CountryCoverageLocal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCountryCoverageLocal() *CountryCoverageLocal {
	this := CountryCoverageLocal{}
	return &this
}

// NewCountryCoverageLocalWithDefaults instantiates a new CountryCoverageLocal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCountryCoverageLocalWithDefaults() *CountryCoverageLocal {
	this := CountryCoverageLocal{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *CountryCoverageLocal) GetFeatures() []string {
	if o == nil || IsNil(o.Features) {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryCoverageLocal) GetFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *CountryCoverageLocal) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *CountryCoverageLocal) SetFeatures(v []string) {
	o.Features = v
}

// GetReservable returns the Reservable field value if set, zero value otherwise.
func (o *CountryCoverageLocal) GetReservable() bool {
	if o == nil || IsNil(o.Reservable) {
		var ret bool
		return ret
	}
	return *o.Reservable
}

// GetReservableOk returns a tuple with the Reservable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryCoverageLocal) GetReservableOk() (*bool, bool) {
	if o == nil || IsNil(o.Reservable) {
		return nil, false
	}
	return o.Reservable, true
}

// HasReservable returns a boolean if a field has been set.
func (o *CountryCoverageLocal) HasReservable() bool {
	if o != nil && !IsNil(o.Reservable) {
		return true
	}

	return false
}

// SetReservable gets a reference to the given bool and assigns it to the Reservable field.
func (o *CountryCoverageLocal) SetReservable(v bool) {
	o.Reservable = &v
}

// GetQuickship returns the Quickship field value if set, zero value otherwise.
func (o *CountryCoverageLocal) GetQuickship() bool {
	if o == nil || IsNil(o.Quickship) {
		var ret bool
		return ret
	}
	return *o.Quickship
}

// GetQuickshipOk returns a tuple with the Quickship field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryCoverageLocal) GetQuickshipOk() (*bool, bool) {
	if o == nil || IsNil(o.Quickship) {
		return nil, false
	}
	return o.Quickship, true
}

// HasQuickship returns a boolean if a field has been set.
func (o *CountryCoverageLocal) HasQuickship() bool {
	if o != nil && !IsNil(o.Quickship) {
		return true
	}

	return false
}

// SetQuickship gets a reference to the given bool and assigns it to the Quickship field.
func (o *CountryCoverageLocal) SetQuickship(v bool) {
	o.Quickship = &v
}

// GetInternationalSms returns the InternationalSms field value if set, zero value otherwise.
func (o *CountryCoverageLocal) GetInternationalSms() bool {
	if o == nil || IsNil(o.InternationalSms) {
		var ret bool
		return ret
	}
	return *o.InternationalSms
}

// GetInternationalSmsOk returns a tuple with the InternationalSms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryCoverageLocal) GetInternationalSmsOk() (*bool, bool) {
	if o == nil || IsNil(o.InternationalSms) {
		return nil, false
	}
	return o.InternationalSms, true
}

// HasInternationalSms returns a boolean if a field has been set.
func (o *CountryCoverageLocal) HasInternationalSms() bool {
	if o != nil && !IsNil(o.InternationalSms) {
		return true
	}

	return false
}

// SetInternationalSms gets a reference to the given bool and assigns it to the InternationalSms field.
func (o *CountryCoverageLocal) SetInternationalSms(v bool) {
	o.InternationalSms = &v
}

// GetP2p returns the P2p field value if set, zero value otherwise.
func (o *CountryCoverageLocal) GetP2p() bool {
	if o == nil || IsNil(o.P2p) {
		var ret bool
		return ret
	}
	return *o.P2p
}

// GetP2pOk returns a tuple with the P2p field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CountryCoverageLocal) GetP2pOk() (*bool, bool) {
	if o == nil || IsNil(o.P2p) {
		return nil, false
	}
	return o.P2p, true
}

// HasP2p returns a boolean if a field has been set.
func (o *CountryCoverageLocal) HasP2p() bool {
	if o != nil && !IsNil(o.P2p) {
		return true
	}

	return false
}

// SetP2p gets a reference to the given bool and assigns it to the P2p field.
func (o *CountryCoverageLocal) SetP2p(v bool) {
	o.P2p = &v
}

func (o CountryCoverageLocal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CountryCoverageLocal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.Reservable) {
		toSerialize["reservable"] = o.Reservable
	}
	if !IsNil(o.Quickship) {
		toSerialize["quickship"] = o.Quickship
	}
	if !IsNil(o.InternationalSms) {
		toSerialize["international_sms"] = o.InternationalSms
	}
	if !IsNil(o.P2p) {
		toSerialize["p2p"] = o.P2p
	}
	return toSerialize, nil
}

type NullableCountryCoverageLocal struct {
	value *CountryCoverageLocal
	isSet bool
}

func (v NullableCountryCoverageLocal) Get() *CountryCoverageLocal {
	return v.value
}

func (v *NullableCountryCoverageLocal) Set(val *CountryCoverageLocal) {
	v.value = val
	v.isSet = true
}

func (v NullableCountryCoverageLocal) IsSet() bool {
	return v.isSet
}

func (v *NullableCountryCoverageLocal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCountryCoverageLocal(val *CountryCoverageLocal) *NullableCountryCoverageLocal {
	return &NullableCountryCoverageLocal{value: val, isSet: true}
}

func (v NullableCountryCoverageLocal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCountryCoverageLocal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


