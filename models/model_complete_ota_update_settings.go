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

// checks if the CompleteOTAUpdateSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CompleteOTAUpdateSettings{}

// CompleteOTAUpdateSettings A JSON object representation of the operation. The information present here will relate directly to the source of the OTA request.
type CompleteOTAUpdateSettings struct {
	// A list of mobile network operators and the priority that should be applied when the SIM is connecting to the network.
	MobileNetworkOperatorsPreferences []MobileNetworkOperatorPreferencesResponse `json:"mobile_network_operators_preferences,omitempty"`
}

// NewCompleteOTAUpdateSettings instantiates a new CompleteOTAUpdateSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCompleteOTAUpdateSettings() *CompleteOTAUpdateSettings {
	this := CompleteOTAUpdateSettings{}
	return &this
}

// NewCompleteOTAUpdateSettingsWithDefaults instantiates a new CompleteOTAUpdateSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCompleteOTAUpdateSettingsWithDefaults() *CompleteOTAUpdateSettings {
	this := CompleteOTAUpdateSettings{}
	return &this
}

// GetMobileNetworkOperatorsPreferences returns the MobileNetworkOperatorsPreferences field value if set, zero value otherwise.
func (o *CompleteOTAUpdateSettings) GetMobileNetworkOperatorsPreferences() []MobileNetworkOperatorPreferencesResponse {
	if o == nil || IsNil(o.MobileNetworkOperatorsPreferences) {
		var ret []MobileNetworkOperatorPreferencesResponse
		return ret
	}
	return o.MobileNetworkOperatorsPreferences
}

// GetMobileNetworkOperatorsPreferencesOk returns a tuple with the MobileNetworkOperatorsPreferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CompleteOTAUpdateSettings) GetMobileNetworkOperatorsPreferencesOk() ([]MobileNetworkOperatorPreferencesResponse, bool) {
	if o == nil || IsNil(o.MobileNetworkOperatorsPreferences) {
		return nil, false
	}
	return o.MobileNetworkOperatorsPreferences, true
}

// HasMobileNetworkOperatorsPreferences returns a boolean if a field has been set.
func (o *CompleteOTAUpdateSettings) HasMobileNetworkOperatorsPreferences() bool {
	if o != nil && !IsNil(o.MobileNetworkOperatorsPreferences) {
		return true
	}

	return false
}

// SetMobileNetworkOperatorsPreferences gets a reference to the given []MobileNetworkOperatorPreferencesResponse and assigns it to the MobileNetworkOperatorsPreferences field.
func (o *CompleteOTAUpdateSettings) SetMobileNetworkOperatorsPreferences(v []MobileNetworkOperatorPreferencesResponse) {
	o.MobileNetworkOperatorsPreferences = v
}

func (o CompleteOTAUpdateSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CompleteOTAUpdateSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MobileNetworkOperatorsPreferences) {
		toSerialize["mobile_network_operators_preferences"] = o.MobileNetworkOperatorsPreferences
	}
	return toSerialize, nil
}

type NullableCompleteOTAUpdateSettings struct {
	value *CompleteOTAUpdateSettings
	isSet bool
}

func (v NullableCompleteOTAUpdateSettings) Get() *CompleteOTAUpdateSettings {
	return v.value
}

func (v *NullableCompleteOTAUpdateSettings) Set(val *CompleteOTAUpdateSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableCompleteOTAUpdateSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableCompleteOTAUpdateSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCompleteOTAUpdateSettings(val *CompleteOTAUpdateSettings) *NullableCompleteOTAUpdateSettings {
	return &NullableCompleteOTAUpdateSettings{value: val, isSet: true}
}

func (v NullableCompleteOTAUpdateSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCompleteOTAUpdateSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


