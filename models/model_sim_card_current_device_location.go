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

// checks if the SIMCardCurrentDeviceLocation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SIMCardCurrentDeviceLocation{}

// SIMCardCurrentDeviceLocation Current physical location data of a given SIM card. Accuracy is given in meters.
type SIMCardCurrentDeviceLocation struct {
	Latitude *float32 `json:"latitude,omitempty"`
	Longitude *float32 `json:"longitude,omitempty"`
	Accuracy *int32 `json:"accuracy,omitempty"`
	AccuracyUnit *string `json:"accuracy_unit,omitempty"`
}

// NewSIMCardCurrentDeviceLocation instantiates a new SIMCardCurrentDeviceLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSIMCardCurrentDeviceLocation() *SIMCardCurrentDeviceLocation {
	this := SIMCardCurrentDeviceLocation{}
	var accuracyUnit string = "m"
	this.AccuracyUnit = &accuracyUnit
	return &this
}

// NewSIMCardCurrentDeviceLocationWithDefaults instantiates a new SIMCardCurrentDeviceLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSIMCardCurrentDeviceLocationWithDefaults() *SIMCardCurrentDeviceLocation {
	this := SIMCardCurrentDeviceLocation{}
	var accuracyUnit string = "m"
	this.AccuracyUnit = &accuracyUnit
	return &this
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *SIMCardCurrentDeviceLocation) GetLatitude() float32 {
	if o == nil || IsNil(o.Latitude) {
		var ret float32
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardCurrentDeviceLocation) GetLatitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Latitude) {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *SIMCardCurrentDeviceLocation) HasLatitude() bool {
	if o != nil && !IsNil(o.Latitude) {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given float32 and assigns it to the Latitude field.
func (o *SIMCardCurrentDeviceLocation) SetLatitude(v float32) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *SIMCardCurrentDeviceLocation) GetLongitude() float32 {
	if o == nil || IsNil(o.Longitude) {
		var ret float32
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardCurrentDeviceLocation) GetLongitudeOk() (*float32, bool) {
	if o == nil || IsNil(o.Longitude) {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *SIMCardCurrentDeviceLocation) HasLongitude() bool {
	if o != nil && !IsNil(o.Longitude) {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given float32 and assigns it to the Longitude field.
func (o *SIMCardCurrentDeviceLocation) SetLongitude(v float32) {
	o.Longitude = &v
}

// GetAccuracy returns the Accuracy field value if set, zero value otherwise.
func (o *SIMCardCurrentDeviceLocation) GetAccuracy() int32 {
	if o == nil || IsNil(o.Accuracy) {
		var ret int32
		return ret
	}
	return *o.Accuracy
}

// GetAccuracyOk returns a tuple with the Accuracy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardCurrentDeviceLocation) GetAccuracyOk() (*int32, bool) {
	if o == nil || IsNil(o.Accuracy) {
		return nil, false
	}
	return o.Accuracy, true
}

// HasAccuracy returns a boolean if a field has been set.
func (o *SIMCardCurrentDeviceLocation) HasAccuracy() bool {
	if o != nil && !IsNil(o.Accuracy) {
		return true
	}

	return false
}

// SetAccuracy gets a reference to the given int32 and assigns it to the Accuracy field.
func (o *SIMCardCurrentDeviceLocation) SetAccuracy(v int32) {
	o.Accuracy = &v
}

// GetAccuracyUnit returns the AccuracyUnit field value if set, zero value otherwise.
func (o *SIMCardCurrentDeviceLocation) GetAccuracyUnit() string {
	if o == nil || IsNil(o.AccuracyUnit) {
		var ret string
		return ret
	}
	return *o.AccuracyUnit
}

// GetAccuracyUnitOk returns a tuple with the AccuracyUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SIMCardCurrentDeviceLocation) GetAccuracyUnitOk() (*string, bool) {
	if o == nil || IsNil(o.AccuracyUnit) {
		return nil, false
	}
	return o.AccuracyUnit, true
}

// HasAccuracyUnit returns a boolean if a field has been set.
func (o *SIMCardCurrentDeviceLocation) HasAccuracyUnit() bool {
	if o != nil && !IsNil(o.AccuracyUnit) {
		return true
	}

	return false
}

// SetAccuracyUnit gets a reference to the given string and assigns it to the AccuracyUnit field.
func (o *SIMCardCurrentDeviceLocation) SetAccuracyUnit(v string) {
	o.AccuracyUnit = &v
}

func (o SIMCardCurrentDeviceLocation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SIMCardCurrentDeviceLocation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Latitude) {
		toSerialize["latitude"] = o.Latitude
	}
	if !IsNil(o.Longitude) {
		toSerialize["longitude"] = o.Longitude
	}
	if !IsNil(o.Accuracy) {
		toSerialize["accuracy"] = o.Accuracy
	}
	if !IsNil(o.AccuracyUnit) {
		toSerialize["accuracy_unit"] = o.AccuracyUnit
	}
	return toSerialize, nil
}

type NullableSIMCardCurrentDeviceLocation struct {
	value *SIMCardCurrentDeviceLocation
	isSet bool
}

func (v NullableSIMCardCurrentDeviceLocation) Get() *SIMCardCurrentDeviceLocation {
	return v.value
}

func (v *NullableSIMCardCurrentDeviceLocation) Set(val *SIMCardCurrentDeviceLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableSIMCardCurrentDeviceLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableSIMCardCurrentDeviceLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSIMCardCurrentDeviceLocation(val *SIMCardCurrentDeviceLocation) *NullableSIMCardCurrentDeviceLocation {
	return &NullableSIMCardCurrentDeviceLocation{value: val, isSet: true}
}

func (v NullableSIMCardCurrentDeviceLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSIMCardCurrentDeviceLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


