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

// checks if the LocationResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocationResponseData{}

// LocationResponseData struct for LocationResponseData
type LocationResponseData struct {
	LocationId *string `json:"location_id,omitempty"`
	StaticEmergencyAddressId *string `json:"static_emergency_address_id,omitempty"`
	AcceptedAddressSuggestions *bool `json:"accepted_address_suggestions,omitempty"`
}

// NewLocationResponseData instantiates a new LocationResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocationResponseData() *LocationResponseData {
	this := LocationResponseData{}
	return &this
}

// NewLocationResponseDataWithDefaults instantiates a new LocationResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationResponseDataWithDefaults() *LocationResponseData {
	this := LocationResponseData{}
	return &this
}

// GetLocationId returns the LocationId field value if set, zero value otherwise.
func (o *LocationResponseData) GetLocationId() string {
	if o == nil || IsNil(o.LocationId) {
		var ret string
		return ret
	}
	return *o.LocationId
}

// GetLocationIdOk returns a tuple with the LocationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationResponseData) GetLocationIdOk() (*string, bool) {
	if o == nil || IsNil(o.LocationId) {
		return nil, false
	}
	return o.LocationId, true
}

// HasLocationId returns a boolean if a field has been set.
func (o *LocationResponseData) HasLocationId() bool {
	if o != nil && !IsNil(o.LocationId) {
		return true
	}

	return false
}

// SetLocationId gets a reference to the given string and assigns it to the LocationId field.
func (o *LocationResponseData) SetLocationId(v string) {
	o.LocationId = &v
}

// GetStaticEmergencyAddressId returns the StaticEmergencyAddressId field value if set, zero value otherwise.
func (o *LocationResponseData) GetStaticEmergencyAddressId() string {
	if o == nil || IsNil(o.StaticEmergencyAddressId) {
		var ret string
		return ret
	}
	return *o.StaticEmergencyAddressId
}

// GetStaticEmergencyAddressIdOk returns a tuple with the StaticEmergencyAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationResponseData) GetStaticEmergencyAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.StaticEmergencyAddressId) {
		return nil, false
	}
	return o.StaticEmergencyAddressId, true
}

// HasStaticEmergencyAddressId returns a boolean if a field has been set.
func (o *LocationResponseData) HasStaticEmergencyAddressId() bool {
	if o != nil && !IsNil(o.StaticEmergencyAddressId) {
		return true
	}

	return false
}

// SetStaticEmergencyAddressId gets a reference to the given string and assigns it to the StaticEmergencyAddressId field.
func (o *LocationResponseData) SetStaticEmergencyAddressId(v string) {
	o.StaticEmergencyAddressId = &v
}

// GetAcceptedAddressSuggestions returns the AcceptedAddressSuggestions field value if set, zero value otherwise.
func (o *LocationResponseData) GetAcceptedAddressSuggestions() bool {
	if o == nil || IsNil(o.AcceptedAddressSuggestions) {
		var ret bool
		return ret
	}
	return *o.AcceptedAddressSuggestions
}

// GetAcceptedAddressSuggestionsOk returns a tuple with the AcceptedAddressSuggestions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocationResponseData) GetAcceptedAddressSuggestionsOk() (*bool, bool) {
	if o == nil || IsNil(o.AcceptedAddressSuggestions) {
		return nil, false
	}
	return o.AcceptedAddressSuggestions, true
}

// HasAcceptedAddressSuggestions returns a boolean if a field has been set.
func (o *LocationResponseData) HasAcceptedAddressSuggestions() bool {
	if o != nil && !IsNil(o.AcceptedAddressSuggestions) {
		return true
	}

	return false
}

// SetAcceptedAddressSuggestions gets a reference to the given bool and assigns it to the AcceptedAddressSuggestions field.
func (o *LocationResponseData) SetAcceptedAddressSuggestions(v bool) {
	o.AcceptedAddressSuggestions = &v
}

func (o LocationResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocationResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LocationId) {
		toSerialize["location_id"] = o.LocationId
	}
	if !IsNil(o.StaticEmergencyAddressId) {
		toSerialize["static_emergency_address_id"] = o.StaticEmergencyAddressId
	}
	if !IsNil(o.AcceptedAddressSuggestions) {
		toSerialize["accepted_address_suggestions"] = o.AcceptedAddressSuggestions
	}
	return toSerialize, nil
}

type NullableLocationResponseData struct {
	value *LocationResponseData
	isSet bool
}

func (v NullableLocationResponseData) Get() *LocationResponseData {
	return v.value
}

func (v *NullableLocationResponseData) Set(val *LocationResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationResponseData(val *LocationResponseData) *NullableLocationResponseData {
	return &NullableLocationResponseData{value: val, isSet: true}
}

func (v NullableLocationResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


