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

// WebhookApiVersion Determines which webhook format will be used, Telnyx API v1 or v2.
type WebhookApiVersion string

// List of WebhookApiVersion
const (
	_1 WebhookApiVersion = "1"
	_2 WebhookApiVersion = "2"
)

// All allowed values of WebhookApiVersion enum
var AllowedWebhookApiVersionEnumValues = []WebhookApiVersion{
	"1",
	"2",
}

func (v *WebhookApiVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebhookApiVersion(value)
	for _, existing := range AllowedWebhookApiVersionEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebhookApiVersion", value)
}

// NewWebhookApiVersionFromValue returns a pointer to a valid WebhookApiVersion
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebhookApiVersionFromValue(v string) (*WebhookApiVersion, error) {
	ev := WebhookApiVersion(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebhookApiVersion: valid values are %v", v, AllowedWebhookApiVersionEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebhookApiVersion) IsValid() bool {
	for _, existing := range AllowedWebhookApiVersionEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WebhookApiVersion value
func (v WebhookApiVersion) Ptr() *WebhookApiVersion {
	return &v
}

type NullableWebhookApiVersion struct {
	value *WebhookApiVersion
	isSet bool
}

func (v NullableWebhookApiVersion) Get() *WebhookApiVersion {
	return v.value
}

func (v *NullableWebhookApiVersion) Set(val *WebhookApiVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookApiVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookApiVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookApiVersion(val *WebhookApiVersion) *NullableWebhookApiVersion {
	return &NullableWebhookApiVersion{value: val, isSet: true}
}

func (v NullableWebhookApiVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookApiVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

