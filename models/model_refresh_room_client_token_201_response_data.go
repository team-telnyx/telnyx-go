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

// checks if the RefreshRoomClientToken201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshRoomClientToken201ResponseData{}

// RefreshRoomClientToken201ResponseData struct for RefreshRoomClientToken201ResponseData
type RefreshRoomClientToken201ResponseData struct {
	Token *string `json:"token,omitempty"`
	// ISO 8601 timestamp when the token expires.
	TokenExpiresAt *string `json:"token_expires_at,omitempty"`
}

// NewRefreshRoomClientToken201ResponseData instantiates a new RefreshRoomClientToken201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshRoomClientToken201ResponseData() *RefreshRoomClientToken201ResponseData {
	this := RefreshRoomClientToken201ResponseData{}
	return &this
}

// NewRefreshRoomClientToken201ResponseDataWithDefaults instantiates a new RefreshRoomClientToken201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshRoomClientToken201ResponseDataWithDefaults() *RefreshRoomClientToken201ResponseData {
	this := RefreshRoomClientToken201ResponseData{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *RefreshRoomClientToken201ResponseData) GetToken() string {
	if o == nil || IsNil(o.Token) {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshRoomClientToken201ResponseData) GetTokenOk() (*string, bool) {
	if o == nil || IsNil(o.Token) {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *RefreshRoomClientToken201ResponseData) HasToken() bool {
	if o != nil && !IsNil(o.Token) {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *RefreshRoomClientToken201ResponseData) SetToken(v string) {
	o.Token = &v
}

// GetTokenExpiresAt returns the TokenExpiresAt field value if set, zero value otherwise.
func (o *RefreshRoomClientToken201ResponseData) GetTokenExpiresAt() string {
	if o == nil || IsNil(o.TokenExpiresAt) {
		var ret string
		return ret
	}
	return *o.TokenExpiresAt
}

// GetTokenExpiresAtOk returns a tuple with the TokenExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshRoomClientToken201ResponseData) GetTokenExpiresAtOk() (*string, bool) {
	if o == nil || IsNil(o.TokenExpiresAt) {
		return nil, false
	}
	return o.TokenExpiresAt, true
}

// HasTokenExpiresAt returns a boolean if a field has been set.
func (o *RefreshRoomClientToken201ResponseData) HasTokenExpiresAt() bool {
	if o != nil && !IsNil(o.TokenExpiresAt) {
		return true
	}

	return false
}

// SetTokenExpiresAt gets a reference to the given string and assigns it to the TokenExpiresAt field.
func (o *RefreshRoomClientToken201ResponseData) SetTokenExpiresAt(v string) {
	o.TokenExpiresAt = &v
}

func (o RefreshRoomClientToken201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshRoomClientToken201ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Token) {
		toSerialize["token"] = o.Token
	}
	if !IsNil(o.TokenExpiresAt) {
		toSerialize["token_expires_at"] = o.TokenExpiresAt
	}
	return toSerialize, nil
}

type NullableRefreshRoomClientToken201ResponseData struct {
	value *RefreshRoomClientToken201ResponseData
	isSet bool
}

func (v NullableRefreshRoomClientToken201ResponseData) Get() *RefreshRoomClientToken201ResponseData {
	return v.value
}

func (v *NullableRefreshRoomClientToken201ResponseData) Set(val *RefreshRoomClientToken201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshRoomClientToken201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshRoomClientToken201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshRoomClientToken201ResponseData(val *RefreshRoomClientToken201ResponseData) *NullableRefreshRoomClientToken201ResponseData {
	return &NullableRefreshRoomClientToken201ResponseData{value: val, isSet: true}
}

func (v NullableRefreshRoomClientToken201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshRoomClientToken201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


