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

// PortingOrderActivationStatus Activation status
type PortingOrderActivationStatus string

// List of PortingOrderActivationStatus
const (
	NEW PortingOrderActivationStatus = "New"
	PENDING PortingOrderActivationStatus = "Pending"
	CONFLICT PortingOrderActivationStatus = "Conflict"
	CANCEL_PENDING PortingOrderActivationStatus = "Cancel Pending"
	FAILED PortingOrderActivationStatus = "Failed"
	CONCURRED PortingOrderActivationStatus = "Concurred"
	ACTIVATE_RDY PortingOrderActivationStatus = "Activate RDY"
	DISCONNECT_PENDING PortingOrderActivationStatus = "Disconnect Pending"
	CONCURRENCE_SENT PortingOrderActivationStatus = "Concurrence Sent"
	OLD PortingOrderActivationStatus = "Old"
	SENDING PortingOrderActivationStatus = "Sending"
	ACTIVE PortingOrderActivationStatus = "Active"
	CANCELLED PortingOrderActivationStatus = "Cancelled"
)

// All allowed values of PortingOrderActivationStatus enum
var AllowedPortingOrderActivationStatusEnumValues = []PortingOrderActivationStatus{
	"New",
	"Pending",
	"Conflict",
	"Cancel Pending",
	"Failed",
	"Concurred",
	"Activate RDY",
	"Disconnect Pending",
	"Concurrence Sent",
	"Old",
	"Sending",
	"Active",
	"Cancelled",
}

func (v *PortingOrderActivationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PortingOrderActivationStatus(value)
	for _, existing := range AllowedPortingOrderActivationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PortingOrderActivationStatus", value)
}

// NewPortingOrderActivationStatusFromValue returns a pointer to a valid PortingOrderActivationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPortingOrderActivationStatusFromValue(v string) (*PortingOrderActivationStatus, error) {
	ev := PortingOrderActivationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PortingOrderActivationStatus: valid values are %v", v, AllowedPortingOrderActivationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PortingOrderActivationStatus) IsValid() bool {
	for _, existing := range AllowedPortingOrderActivationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PortingOrderActivationStatus value
func (v PortingOrderActivationStatus) Ptr() *PortingOrderActivationStatus {
	return &v
}

type NullablePortingOrderActivationStatus struct {
	value *PortingOrderActivationStatus
	isSet bool
}

func (v NullablePortingOrderActivationStatus) Get() *PortingOrderActivationStatus {
	return v.value
}

func (v *NullablePortingOrderActivationStatus) Set(val *PortingOrderActivationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePortingOrderActivationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePortingOrderActivationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortingOrderActivationStatus(val *PortingOrderActivationStatus) *NullablePortingOrderActivationStatus {
	return &NullablePortingOrderActivationStatus{value: val, isSet: true}
}

func (v NullablePortingOrderActivationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortingOrderActivationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

