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

// checks if the ReferRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReferRequest{}

// ReferRequest struct for ReferRequest
type ReferRequest struct {
	// The SIP URI to which the call will be referred to.
	SipAddress string `json:"sip_address"`
	// Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string.
	ClientState *string `json:"client_state,omitempty"`
	// Use this field to avoid execution of duplicate commands. Telnyx will ignore subsequent commands with the same `command_id` as one that has already been executed.
	CommandId *string `json:"command_id,omitempty"`
	// Custom headers to be added to the SIP INVITE.
	CustomHeaders []CustomSipHeader `json:"custom_headers,omitempty"`
	// SIP Authentication username used for SIP challenges.
	SipAuthUsername *string `json:"sip_auth_username,omitempty"`
	// SIP Authentication password used for SIP challenges.
	SipAuthPassword *string `json:"sip_auth_password,omitempty"`
	// SIP headers to be added to the request. Currently only User-to-User header is supported.
	SipHeaders []SipHeader `json:"sip_headers,omitempty"`
}

type _ReferRequest ReferRequest

// NewReferRequest instantiates a new ReferRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReferRequest(sipAddress string) *ReferRequest {
	this := ReferRequest{}
	this.SipAddress = sipAddress
	return &this
}

// NewReferRequestWithDefaults instantiates a new ReferRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReferRequestWithDefaults() *ReferRequest {
	this := ReferRequest{}
	return &this
}

// GetSipAddress returns the SipAddress field value
func (o *ReferRequest) GetSipAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SipAddress
}

// GetSipAddressOk returns a tuple with the SipAddress field value
// and a boolean to check if the value has been set.
func (o *ReferRequest) GetSipAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SipAddress, true
}

// SetSipAddress sets field value
func (o *ReferRequest) SetSipAddress(v string) {
	o.SipAddress = v
}

// GetClientState returns the ClientState field value if set, zero value otherwise.
func (o *ReferRequest) GetClientState() string {
	if o == nil || IsNil(o.ClientState) {
		var ret string
		return ret
	}
	return *o.ClientState
}

// GetClientStateOk returns a tuple with the ClientState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferRequest) GetClientStateOk() (*string, bool) {
	if o == nil || IsNil(o.ClientState) {
		return nil, false
	}
	return o.ClientState, true
}

// HasClientState returns a boolean if a field has been set.
func (o *ReferRequest) HasClientState() bool {
	if o != nil && !IsNil(o.ClientState) {
		return true
	}

	return false
}

// SetClientState gets a reference to the given string and assigns it to the ClientState field.
func (o *ReferRequest) SetClientState(v string) {
	o.ClientState = &v
}

// GetCommandId returns the CommandId field value if set, zero value otherwise.
func (o *ReferRequest) GetCommandId() string {
	if o == nil || IsNil(o.CommandId) {
		var ret string
		return ret
	}
	return *o.CommandId
}

// GetCommandIdOk returns a tuple with the CommandId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferRequest) GetCommandIdOk() (*string, bool) {
	if o == nil || IsNil(o.CommandId) {
		return nil, false
	}
	return o.CommandId, true
}

// HasCommandId returns a boolean if a field has been set.
func (o *ReferRequest) HasCommandId() bool {
	if o != nil && !IsNil(o.CommandId) {
		return true
	}

	return false
}

// SetCommandId gets a reference to the given string and assigns it to the CommandId field.
func (o *ReferRequest) SetCommandId(v string) {
	o.CommandId = &v
}

// GetCustomHeaders returns the CustomHeaders field value if set, zero value otherwise.
func (o *ReferRequest) GetCustomHeaders() []CustomSipHeader {
	if o == nil || IsNil(o.CustomHeaders) {
		var ret []CustomSipHeader
		return ret
	}
	return o.CustomHeaders
}

// GetCustomHeadersOk returns a tuple with the CustomHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferRequest) GetCustomHeadersOk() ([]CustomSipHeader, bool) {
	if o == nil || IsNil(o.CustomHeaders) {
		return nil, false
	}
	return o.CustomHeaders, true
}

// HasCustomHeaders returns a boolean if a field has been set.
func (o *ReferRequest) HasCustomHeaders() bool {
	if o != nil && !IsNil(o.CustomHeaders) {
		return true
	}

	return false
}

// SetCustomHeaders gets a reference to the given []CustomSipHeader and assigns it to the CustomHeaders field.
func (o *ReferRequest) SetCustomHeaders(v []CustomSipHeader) {
	o.CustomHeaders = v
}

// GetSipAuthUsername returns the SipAuthUsername field value if set, zero value otherwise.
func (o *ReferRequest) GetSipAuthUsername() string {
	if o == nil || IsNil(o.SipAuthUsername) {
		var ret string
		return ret
	}
	return *o.SipAuthUsername
}

// GetSipAuthUsernameOk returns a tuple with the SipAuthUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferRequest) GetSipAuthUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.SipAuthUsername) {
		return nil, false
	}
	return o.SipAuthUsername, true
}

// HasSipAuthUsername returns a boolean if a field has been set.
func (o *ReferRequest) HasSipAuthUsername() bool {
	if o != nil && !IsNil(o.SipAuthUsername) {
		return true
	}

	return false
}

// SetSipAuthUsername gets a reference to the given string and assigns it to the SipAuthUsername field.
func (o *ReferRequest) SetSipAuthUsername(v string) {
	o.SipAuthUsername = &v
}

// GetSipAuthPassword returns the SipAuthPassword field value if set, zero value otherwise.
func (o *ReferRequest) GetSipAuthPassword() string {
	if o == nil || IsNil(o.SipAuthPassword) {
		var ret string
		return ret
	}
	return *o.SipAuthPassword
}

// GetSipAuthPasswordOk returns a tuple with the SipAuthPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferRequest) GetSipAuthPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.SipAuthPassword) {
		return nil, false
	}
	return o.SipAuthPassword, true
}

// HasSipAuthPassword returns a boolean if a field has been set.
func (o *ReferRequest) HasSipAuthPassword() bool {
	if o != nil && !IsNil(o.SipAuthPassword) {
		return true
	}

	return false
}

// SetSipAuthPassword gets a reference to the given string and assigns it to the SipAuthPassword field.
func (o *ReferRequest) SetSipAuthPassword(v string) {
	o.SipAuthPassword = &v
}

// GetSipHeaders returns the SipHeaders field value if set, zero value otherwise.
func (o *ReferRequest) GetSipHeaders() []SipHeader {
	if o == nil || IsNil(o.SipHeaders) {
		var ret []SipHeader
		return ret
	}
	return o.SipHeaders
}

// GetSipHeadersOk returns a tuple with the SipHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReferRequest) GetSipHeadersOk() ([]SipHeader, bool) {
	if o == nil || IsNil(o.SipHeaders) {
		return nil, false
	}
	return o.SipHeaders, true
}

// HasSipHeaders returns a boolean if a field has been set.
func (o *ReferRequest) HasSipHeaders() bool {
	if o != nil && !IsNil(o.SipHeaders) {
		return true
	}

	return false
}

// SetSipHeaders gets a reference to the given []SipHeader and assigns it to the SipHeaders field.
func (o *ReferRequest) SetSipHeaders(v []SipHeader) {
	o.SipHeaders = v
}

func (o ReferRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReferRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sip_address"] = o.SipAddress
	if !IsNil(o.ClientState) {
		toSerialize["client_state"] = o.ClientState
	}
	if !IsNil(o.CommandId) {
		toSerialize["command_id"] = o.CommandId
	}
	if !IsNil(o.CustomHeaders) {
		toSerialize["custom_headers"] = o.CustomHeaders
	}
	if !IsNil(o.SipAuthUsername) {
		toSerialize["sip_auth_username"] = o.SipAuthUsername
	}
	if !IsNil(o.SipAuthPassword) {
		toSerialize["sip_auth_password"] = o.SipAuthPassword
	}
	if !IsNil(o.SipHeaders) {
		toSerialize["sip_headers"] = o.SipHeaders
	}
	return toSerialize, nil
}

func (o *ReferRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sip_address",
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

	varReferRequest := _ReferRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varReferRequest)

	if err != nil {
		return err
	}

	*o = ReferRequest(varReferRequest)

	return err
}

type NullableReferRequest struct {
	value *ReferRequest
	isSet bool
}

func (v NullableReferRequest) Get() *ReferRequest {
	return v.value
}

func (v *NullableReferRequest) Set(val *ReferRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReferRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReferRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReferRequest(val *ReferRequest) *NullableReferRequest {
	return &NullableReferRequest{value: val, isSet: true}
}

func (v NullableReferRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReferRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


