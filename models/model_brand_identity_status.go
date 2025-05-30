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

// BrandIdentityStatus The verification status of an active brand
type BrandIdentityStatus string

// List of BrandIdentityStatus
const (
	VERIFIED BrandIdentityStatus = "VERIFIED"
	UNVERIFIED BrandIdentityStatus = "UNVERIFIED"
	SELF_DECLARED BrandIdentityStatus = "SELF_DECLARED"
	VETTED_VERIFIED BrandIdentityStatus = "VETTED_VERIFIED"
)

// All allowed values of BrandIdentityStatus enum
var AllowedBrandIdentityStatusEnumValues = []BrandIdentityStatus{
	"VERIFIED",
	"UNVERIFIED",
	"SELF_DECLARED",
	"VETTED_VERIFIED",
}

func (v *BrandIdentityStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BrandIdentityStatus(value)
	for _, existing := range AllowedBrandIdentityStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BrandIdentityStatus", value)
}

// NewBrandIdentityStatusFromValue returns a pointer to a valid BrandIdentityStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBrandIdentityStatusFromValue(v string) (*BrandIdentityStatus, error) {
	ev := BrandIdentityStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BrandIdentityStatus: valid values are %v", v, AllowedBrandIdentityStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BrandIdentityStatus) IsValid() bool {
	for _, existing := range AllowedBrandIdentityStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BrandIdentityStatus value
func (v BrandIdentityStatus) Ptr() *BrandIdentityStatus {
	return &v
}

type NullableBrandIdentityStatus struct {
	value *BrandIdentityStatus
	isSet bool
}

func (v NullableBrandIdentityStatus) Get() *BrandIdentityStatus {
	return v.value
}

func (v *NullableBrandIdentityStatus) Set(val *BrandIdentityStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableBrandIdentityStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableBrandIdentityStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBrandIdentityStatus(val *BrandIdentityStatus) *NullableBrandIdentityStatus {
	return &NullableBrandIdentityStatus{value: val, isSet: true}
}

func (v NullableBrandIdentityStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBrandIdentityStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

