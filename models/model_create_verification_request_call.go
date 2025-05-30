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
	"bytes"
	"fmt"
)

// checks if the CreateVerificationRequestCall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateVerificationRequestCall{}

// CreateVerificationRequestCall The request body when creating a verification.
type CreateVerificationRequestCall struct {
	// +E164 formatted phone number.
	PhoneNumber string `json:"phone_number"`
	// The identifier of the associated Verify profile.
	VerifyProfileId string `json:"verify_profile_id"`
	// Send a self-generated numeric code to the end-user
	CustomCode NullableString `json:"custom_code,omitempty"`
	// The number of seconds the verification code is valid for.
	TimeoutSecs *int32 `json:"timeout_secs,omitempty"`
}

type _CreateVerificationRequestCall CreateVerificationRequestCall

// NewCreateVerificationRequestCall instantiates a new CreateVerificationRequestCall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateVerificationRequestCall(phoneNumber string, verifyProfileId string) *CreateVerificationRequestCall {
	this := CreateVerificationRequestCall{}
	this.PhoneNumber = phoneNumber
	this.VerifyProfileId = verifyProfileId
	return &this
}

// NewCreateVerificationRequestCallWithDefaults instantiates a new CreateVerificationRequestCall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateVerificationRequestCallWithDefaults() *CreateVerificationRequestCall {
	this := CreateVerificationRequestCall{}
	return &this
}

// GetPhoneNumber returns the PhoneNumber field value
func (o *CreateVerificationRequestCall) GetPhoneNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PhoneNumber
}

// GetPhoneNumberOk returns a tuple with the PhoneNumber field value
// and a boolean to check if the value has been set.
func (o *CreateVerificationRequestCall) GetPhoneNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PhoneNumber, true
}

// SetPhoneNumber sets field value
func (o *CreateVerificationRequestCall) SetPhoneNumber(v string) {
	o.PhoneNumber = v
}

// GetVerifyProfileId returns the VerifyProfileId field value
func (o *CreateVerificationRequestCall) GetVerifyProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VerifyProfileId
}

// GetVerifyProfileIdOk returns a tuple with the VerifyProfileId field value
// and a boolean to check if the value has been set.
func (o *CreateVerificationRequestCall) GetVerifyProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VerifyProfileId, true
}

// SetVerifyProfileId sets field value
func (o *CreateVerificationRequestCall) SetVerifyProfileId(v string) {
	o.VerifyProfileId = v
}

// GetCustomCode returns the CustomCode field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateVerificationRequestCall) GetCustomCode() string {
	if o == nil || IsNil(o.CustomCode.Get()) {
		var ret string
		return ret
	}
	return *o.CustomCode.Get()
}

// GetCustomCodeOk returns a tuple with the CustomCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateVerificationRequestCall) GetCustomCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomCode.Get(), o.CustomCode.IsSet()
}

// HasCustomCode returns a boolean if a field has been set.
func (o *CreateVerificationRequestCall) HasCustomCode() bool {
	if o != nil && o.CustomCode.IsSet() {
		return true
	}

	return false
}

// SetCustomCode gets a reference to the given NullableString and assigns it to the CustomCode field.
func (o *CreateVerificationRequestCall) SetCustomCode(v string) {
	o.CustomCode.Set(&v)
}
// SetCustomCodeNil sets the value for CustomCode to be an explicit nil
func (o *CreateVerificationRequestCall) SetCustomCodeNil() {
	o.CustomCode.Set(nil)
}

// UnsetCustomCode ensures that no value is present for CustomCode, not even an explicit nil
func (o *CreateVerificationRequestCall) UnsetCustomCode() {
	o.CustomCode.Unset()
}

// GetTimeoutSecs returns the TimeoutSecs field value if set, zero value otherwise.
func (o *CreateVerificationRequestCall) GetTimeoutSecs() int32 {
	if o == nil || IsNil(o.TimeoutSecs) {
		var ret int32
		return ret
	}
	return *o.TimeoutSecs
}

// GetTimeoutSecsOk returns a tuple with the TimeoutSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateVerificationRequestCall) GetTimeoutSecsOk() (*int32, bool) {
	if o == nil || IsNil(o.TimeoutSecs) {
		return nil, false
	}
	return o.TimeoutSecs, true
}

// HasTimeoutSecs returns a boolean if a field has been set.
func (o *CreateVerificationRequestCall) HasTimeoutSecs() bool {
	if o != nil && !IsNil(o.TimeoutSecs) {
		return true
	}

	return false
}

// SetTimeoutSecs gets a reference to the given int32 and assigns it to the TimeoutSecs field.
func (o *CreateVerificationRequestCall) SetTimeoutSecs(v int32) {
	o.TimeoutSecs = &v
}

func (o CreateVerificationRequestCall) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateVerificationRequestCall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["phone_number"] = o.PhoneNumber
	toSerialize["verify_profile_id"] = o.VerifyProfileId
	if o.CustomCode.IsSet() {
		toSerialize["custom_code"] = o.CustomCode.Get()
	}
	if !IsNil(o.TimeoutSecs) {
		toSerialize["timeout_secs"] = o.TimeoutSecs
	}
	return toSerialize, nil
}

func (o *CreateVerificationRequestCall) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"phone_number",
		"verify_profile_id",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCreateVerificationRequestCall := _CreateVerificationRequestCall{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateVerificationRequestCall)

	if err != nil {
		return err
	}

	*o = CreateVerificationRequestCall(varCreateVerificationRequestCall)

	return err
}

type NullableCreateVerificationRequestCall struct {
	value *CreateVerificationRequestCall
	isSet bool
}

func (v NullableCreateVerificationRequestCall) Get() *CreateVerificationRequestCall {
	return v.value
}

func (v *NullableCreateVerificationRequestCall) Set(val *CreateVerificationRequestCall) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateVerificationRequestCall) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateVerificationRequestCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateVerificationRequestCall(val *CreateVerificationRequestCall) *NullableCreateVerificationRequestCall {
	return &NullableCreateVerificationRequestCall{value: val, isSet: true}
}

func (v NullableCreateVerificationRequestCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateVerificationRequestCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


