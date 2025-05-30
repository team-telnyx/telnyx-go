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

// checks if the UpdateVerifyProfileSMSRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateVerifyProfileSMSRequest{}

// UpdateVerifyProfileSMSRequest struct for UpdateVerifyProfileSMSRequest
type UpdateVerifyProfileSMSRequest struct {
	// The message template identifier selected from /verify_profiles/templates
	MessagingTemplateId *string `json:"messaging_template_id,omitempty"`
	// The name that identifies the application requesting 2fa in the verification message.
	AppName *string `json:"app_name,omitempty" validate:"regexp=^[A-Za-z0-9 -]{1,30}$"`
	// The alphanumeric sender ID to use when sending to destinations that require an alphanumeric sender ID.
	AlphaSender NullableString `json:"alpha_sender,omitempty" validate:"regexp=^[A-Za-z0-9 ]{1,11}$"`
	// The length of the verify code to generate.
	CodeLength *int32 `json:"code_length,omitempty"`
	// Enabled country destinations to send verification codes. The elements in the list must be valid ISO 3166-1 alpha-2 country codes. If set to `[\"*\"]`, all destinations will be allowed.
	WhitelistedDestinations []string `json:"whitelisted_destinations,omitempty"`
	// For every request that is initiated via this Verify profile, this sets the number of seconds before a verification request code expires. Once the verification request expires, the user cannot use the code to verify their identity.
	DefaultVerificationTimeoutSecs *int32 `json:"default_verification_timeout_secs,omitempty"`
}

// NewUpdateVerifyProfileSMSRequest instantiates a new UpdateVerifyProfileSMSRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateVerifyProfileSMSRequest() *UpdateVerifyProfileSMSRequest {
	this := UpdateVerifyProfileSMSRequest{}
	var alphaSender string = "Telnyx"
	this.AlphaSender = *NewNullableString(&alphaSender)
	var codeLength int32 = 5
	this.CodeLength = &codeLength
	var defaultVerificationTimeoutSecs int32 = 300
	this.DefaultVerificationTimeoutSecs = &defaultVerificationTimeoutSecs
	return &this
}

// NewUpdateVerifyProfileSMSRequestWithDefaults instantiates a new UpdateVerifyProfileSMSRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateVerifyProfileSMSRequestWithDefaults() *UpdateVerifyProfileSMSRequest {
	this := UpdateVerifyProfileSMSRequest{}
	var alphaSender string = "Telnyx"
	this.AlphaSender = *NewNullableString(&alphaSender)
	var codeLength int32 = 5
	this.CodeLength = &codeLength
	var defaultVerificationTimeoutSecs int32 = 300
	this.DefaultVerificationTimeoutSecs = &defaultVerificationTimeoutSecs
	return &this
}

// GetMessagingTemplateId returns the MessagingTemplateId field value if set, zero value otherwise.
func (o *UpdateVerifyProfileSMSRequest) GetMessagingTemplateId() string {
	if o == nil || IsNil(o.MessagingTemplateId) {
		var ret string
		return ret
	}
	return *o.MessagingTemplateId
}

// GetMessagingTemplateIdOk returns a tuple with the MessagingTemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVerifyProfileSMSRequest) GetMessagingTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.MessagingTemplateId) {
		return nil, false
	}
	return o.MessagingTemplateId, true
}

// HasMessagingTemplateId returns a boolean if a field has been set.
func (o *UpdateVerifyProfileSMSRequest) HasMessagingTemplateId() bool {
	if o != nil && !IsNil(o.MessagingTemplateId) {
		return true
	}

	return false
}

// SetMessagingTemplateId gets a reference to the given string and assigns it to the MessagingTemplateId field.
func (o *UpdateVerifyProfileSMSRequest) SetMessagingTemplateId(v string) {
	o.MessagingTemplateId = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *UpdateVerifyProfileSMSRequest) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVerifyProfileSMSRequest) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *UpdateVerifyProfileSMSRequest) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *UpdateVerifyProfileSMSRequest) SetAppName(v string) {
	o.AppName = &v
}

// GetAlphaSender returns the AlphaSender field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateVerifyProfileSMSRequest) GetAlphaSender() string {
	if o == nil || IsNil(o.AlphaSender.Get()) {
		var ret string
		return ret
	}
	return *o.AlphaSender.Get()
}

// GetAlphaSenderOk returns a tuple with the AlphaSender field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateVerifyProfileSMSRequest) GetAlphaSenderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AlphaSender.Get(), o.AlphaSender.IsSet()
}

// HasAlphaSender returns a boolean if a field has been set.
func (o *UpdateVerifyProfileSMSRequest) HasAlphaSender() bool {
	if o != nil && o.AlphaSender.IsSet() {
		return true
	}

	return false
}

// SetAlphaSender gets a reference to the given NullableString and assigns it to the AlphaSender field.
func (o *UpdateVerifyProfileSMSRequest) SetAlphaSender(v string) {
	o.AlphaSender.Set(&v)
}
// SetAlphaSenderNil sets the value for AlphaSender to be an explicit nil
func (o *UpdateVerifyProfileSMSRequest) SetAlphaSenderNil() {
	o.AlphaSender.Set(nil)
}

// UnsetAlphaSender ensures that no value is present for AlphaSender, not even an explicit nil
func (o *UpdateVerifyProfileSMSRequest) UnsetAlphaSender() {
	o.AlphaSender.Unset()
}

// GetCodeLength returns the CodeLength field value if set, zero value otherwise.
func (o *UpdateVerifyProfileSMSRequest) GetCodeLength() int32 {
	if o == nil || IsNil(o.CodeLength) {
		var ret int32
		return ret
	}
	return *o.CodeLength
}

// GetCodeLengthOk returns a tuple with the CodeLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVerifyProfileSMSRequest) GetCodeLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.CodeLength) {
		return nil, false
	}
	return o.CodeLength, true
}

// HasCodeLength returns a boolean if a field has been set.
func (o *UpdateVerifyProfileSMSRequest) HasCodeLength() bool {
	if o != nil && !IsNil(o.CodeLength) {
		return true
	}

	return false
}

// SetCodeLength gets a reference to the given int32 and assigns it to the CodeLength field.
func (o *UpdateVerifyProfileSMSRequest) SetCodeLength(v int32) {
	o.CodeLength = &v
}

// GetWhitelistedDestinations returns the WhitelistedDestinations field value if set, zero value otherwise.
func (o *UpdateVerifyProfileSMSRequest) GetWhitelistedDestinations() []string {
	if o == nil || IsNil(o.WhitelistedDestinations) {
		var ret []string
		return ret
	}
	return o.WhitelistedDestinations
}

// GetWhitelistedDestinationsOk returns a tuple with the WhitelistedDestinations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVerifyProfileSMSRequest) GetWhitelistedDestinationsOk() ([]string, bool) {
	if o == nil || IsNil(o.WhitelistedDestinations) {
		return nil, false
	}
	return o.WhitelistedDestinations, true
}

// HasWhitelistedDestinations returns a boolean if a field has been set.
func (o *UpdateVerifyProfileSMSRequest) HasWhitelistedDestinations() bool {
	if o != nil && !IsNil(o.WhitelistedDestinations) {
		return true
	}

	return false
}

// SetWhitelistedDestinations gets a reference to the given []string and assigns it to the WhitelistedDestinations field.
func (o *UpdateVerifyProfileSMSRequest) SetWhitelistedDestinations(v []string) {
	o.WhitelistedDestinations = v
}

// GetDefaultVerificationTimeoutSecs returns the DefaultVerificationTimeoutSecs field value if set, zero value otherwise.
func (o *UpdateVerifyProfileSMSRequest) GetDefaultVerificationTimeoutSecs() int32 {
	if o == nil || IsNil(o.DefaultVerificationTimeoutSecs) {
		var ret int32
		return ret
	}
	return *o.DefaultVerificationTimeoutSecs
}

// GetDefaultVerificationTimeoutSecsOk returns a tuple with the DefaultVerificationTimeoutSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateVerifyProfileSMSRequest) GetDefaultVerificationTimeoutSecsOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultVerificationTimeoutSecs) {
		return nil, false
	}
	return o.DefaultVerificationTimeoutSecs, true
}

// HasDefaultVerificationTimeoutSecs returns a boolean if a field has been set.
func (o *UpdateVerifyProfileSMSRequest) HasDefaultVerificationTimeoutSecs() bool {
	if o != nil && !IsNil(o.DefaultVerificationTimeoutSecs) {
		return true
	}

	return false
}

// SetDefaultVerificationTimeoutSecs gets a reference to the given int32 and assigns it to the DefaultVerificationTimeoutSecs field.
func (o *UpdateVerifyProfileSMSRequest) SetDefaultVerificationTimeoutSecs(v int32) {
	o.DefaultVerificationTimeoutSecs = &v
}

func (o UpdateVerifyProfileSMSRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateVerifyProfileSMSRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessagingTemplateId) {
		toSerialize["messaging_template_id"] = o.MessagingTemplateId
	}
	if !IsNil(o.AppName) {
		toSerialize["app_name"] = o.AppName
	}
	if o.AlphaSender.IsSet() {
		toSerialize["alpha_sender"] = o.AlphaSender.Get()
	}
	if !IsNil(o.CodeLength) {
		toSerialize["code_length"] = o.CodeLength
	}
	if !IsNil(o.WhitelistedDestinations) {
		toSerialize["whitelisted_destinations"] = o.WhitelistedDestinations
	}
	if !IsNil(o.DefaultVerificationTimeoutSecs) {
		toSerialize["default_verification_timeout_secs"] = o.DefaultVerificationTimeoutSecs
	}
	return toSerialize, nil
}

type NullableUpdateVerifyProfileSMSRequest struct {
	value *UpdateVerifyProfileSMSRequest
	isSet bool
}

func (v NullableUpdateVerifyProfileSMSRequest) Get() *UpdateVerifyProfileSMSRequest {
	return v.value
}

func (v *NullableUpdateVerifyProfileSMSRequest) Set(val *UpdateVerifyProfileSMSRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateVerifyProfileSMSRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateVerifyProfileSMSRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateVerifyProfileSMSRequest(val *UpdateVerifyProfileSMSRequest) *NullableUpdateVerifyProfileSMSRequest {
	return &NullableUpdateVerifyProfileSMSRequest{value: val, isSet: true}
}

func (v NullableUpdateVerifyProfileSMSRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateVerifyProfileSMSRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


