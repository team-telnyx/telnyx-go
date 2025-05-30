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

// BackgroundTaskStatus Status of an embeddings task.
type BackgroundTaskStatus string

// List of BackgroundTaskStatus
const (
	QUEUED BackgroundTaskStatus = "queued"
	PROCESSING BackgroundTaskStatus = "processing"
	SUCCESS BackgroundTaskStatus = "success"
	FAILURE BackgroundTaskStatus = "failure"
	PARTIAL_SUCCESS BackgroundTaskStatus = "partial_success"
)

// All allowed values of BackgroundTaskStatus enum
var AllowedBackgroundTaskStatusEnumValues = []BackgroundTaskStatus{
	"queued",
	"processing",
	"success",
	"failure",
	"partial_success",
}

func (v *BackgroundTaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BackgroundTaskStatus(value)
	for _, existing := range AllowedBackgroundTaskStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BackgroundTaskStatus", value)
}

// NewBackgroundTaskStatusFromValue returns a pointer to a valid BackgroundTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBackgroundTaskStatusFromValue(v string) (*BackgroundTaskStatus, error) {
	ev := BackgroundTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BackgroundTaskStatus: valid values are %v", v, AllowedBackgroundTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BackgroundTaskStatus) IsValid() bool {
	for _, existing := range AllowedBackgroundTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BackgroundTaskStatus value
func (v BackgroundTaskStatus) Ptr() *BackgroundTaskStatus {
	return &v
}

type NullableBackgroundTaskStatus struct {
	value *BackgroundTaskStatus
	isSet bool
}

func (v NullableBackgroundTaskStatus) Get() *BackgroundTaskStatus {
	return v.value
}

func (v *NullableBackgroundTaskStatus) Set(val *BackgroundTaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBackgroundTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBackgroundTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackgroundTaskStatus(val *BackgroundTaskStatus) *NullableBackgroundTaskStatus {
	return &NullableBackgroundTaskStatus{value: val, isSet: true}
}

func (v NullableBackgroundTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackgroundTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

