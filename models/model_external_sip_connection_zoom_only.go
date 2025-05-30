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
	"fmt"
)

// ExternalSipConnectionZoomOnly The service that will be consuming this connection.
type ExternalSipConnectionZoomOnly string

// List of ExternalSipConnectionZoomOnly
const (
	ZOOM ExternalSipConnectionZoomOnly = "zoom"
)

// All allowed values of ExternalSipConnectionZoomOnly enum
var AllowedExternalSipConnectionZoomOnlyEnumValues = []ExternalSipConnectionZoomOnly{
	"zoom",
}

func (v *ExternalSipConnectionZoomOnly) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExternalSipConnectionZoomOnly(value)
	for _, existing := range AllowedExternalSipConnectionZoomOnlyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExternalSipConnectionZoomOnly", value)
}

// NewExternalSipConnectionZoomOnlyFromValue returns a pointer to a valid ExternalSipConnectionZoomOnly
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExternalSipConnectionZoomOnlyFromValue(v string) (*ExternalSipConnectionZoomOnly, error) {
	ev := ExternalSipConnectionZoomOnly(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExternalSipConnectionZoomOnly: valid values are %v", v, AllowedExternalSipConnectionZoomOnlyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExternalSipConnectionZoomOnly) IsValid() bool {
	for _, existing := range AllowedExternalSipConnectionZoomOnlyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExternalSipConnectionZoomOnly value
func (v ExternalSipConnectionZoomOnly) Ptr() *ExternalSipConnectionZoomOnly {
	return &v
}

type NullableExternalSipConnectionZoomOnly struct {
	value *ExternalSipConnectionZoomOnly
	isSet bool
}

func (v NullableExternalSipConnectionZoomOnly) Get() *ExternalSipConnectionZoomOnly {
	return v.value
}

func (v *NullableExternalSipConnectionZoomOnly) Set(val *ExternalSipConnectionZoomOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalSipConnectionZoomOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalSipConnectionZoomOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalSipConnectionZoomOnly(val *ExternalSipConnectionZoomOnly) *NullableExternalSipConnectionZoomOnly {
	return &NullableExternalSipConnectionZoomOnly{value: val, isSet: true}
}

func (v NullableExternalSipConnectionZoomOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalSipConnectionZoomOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

