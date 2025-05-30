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

// checks if the CreateFaxApplicationRequestInbound type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFaxApplicationRequestInbound{}

// CreateFaxApplicationRequestInbound struct for CreateFaxApplicationRequestInbound
type CreateFaxApplicationRequestInbound struct {
	// When set, this will limit the number of concurrent inbound calls to phone numbers associated with this connection.
	ChannelLimit *int32 `json:"channel_limit,omitempty"`
	// Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain \"example.sip.telnyx.com\" can be called from any SIP endpoint by using the SIP URI \"sip:@example.sip.telnyx.com\" where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls.
	SipSubdomain *string `json:"sip_subdomain,omitempty"`
	// This option can be enabled to receive calls from: \"Anyone\" (any SIP endpoint in the public Internet) or \"Only my connections\" (any connection assigned to the same Telnyx user).
	SipSubdomainReceiveSettings *string `json:"sip_subdomain_receive_settings,omitempty"`
}

// NewCreateFaxApplicationRequestInbound instantiates a new CreateFaxApplicationRequestInbound object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFaxApplicationRequestInbound() *CreateFaxApplicationRequestInbound {
	this := CreateFaxApplicationRequestInbound{}
	var sipSubdomainReceiveSettings string = "from_anyone"
	this.SipSubdomainReceiveSettings = &sipSubdomainReceiveSettings
	return &this
}

// NewCreateFaxApplicationRequestInboundWithDefaults instantiates a new CreateFaxApplicationRequestInbound object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFaxApplicationRequestInboundWithDefaults() *CreateFaxApplicationRequestInbound {
	this := CreateFaxApplicationRequestInbound{}
	var sipSubdomainReceiveSettings string = "from_anyone"
	this.SipSubdomainReceiveSettings = &sipSubdomainReceiveSettings
	return &this
}

// GetChannelLimit returns the ChannelLimit field value if set, zero value otherwise.
func (o *CreateFaxApplicationRequestInbound) GetChannelLimit() int32 {
	if o == nil || IsNil(o.ChannelLimit) {
		var ret int32
		return ret
	}
	return *o.ChannelLimit
}

// GetChannelLimitOk returns a tuple with the ChannelLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFaxApplicationRequestInbound) GetChannelLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.ChannelLimit) {
		return nil, false
	}
	return o.ChannelLimit, true
}

// HasChannelLimit returns a boolean if a field has been set.
func (o *CreateFaxApplicationRequestInbound) HasChannelLimit() bool {
	if o != nil && !IsNil(o.ChannelLimit) {
		return true
	}

	return false
}

// SetChannelLimit gets a reference to the given int32 and assigns it to the ChannelLimit field.
func (o *CreateFaxApplicationRequestInbound) SetChannelLimit(v int32) {
	o.ChannelLimit = &v
}

// GetSipSubdomain returns the SipSubdomain field value if set, zero value otherwise.
func (o *CreateFaxApplicationRequestInbound) GetSipSubdomain() string {
	if o == nil || IsNil(o.SipSubdomain) {
		var ret string
		return ret
	}
	return *o.SipSubdomain
}

// GetSipSubdomainOk returns a tuple with the SipSubdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFaxApplicationRequestInbound) GetSipSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.SipSubdomain) {
		return nil, false
	}
	return o.SipSubdomain, true
}

// HasSipSubdomain returns a boolean if a field has been set.
func (o *CreateFaxApplicationRequestInbound) HasSipSubdomain() bool {
	if o != nil && !IsNil(o.SipSubdomain) {
		return true
	}

	return false
}

// SetSipSubdomain gets a reference to the given string and assigns it to the SipSubdomain field.
func (o *CreateFaxApplicationRequestInbound) SetSipSubdomain(v string) {
	o.SipSubdomain = &v
}

// GetSipSubdomainReceiveSettings returns the SipSubdomainReceiveSettings field value if set, zero value otherwise.
func (o *CreateFaxApplicationRequestInbound) GetSipSubdomainReceiveSettings() string {
	if o == nil || IsNil(o.SipSubdomainReceiveSettings) {
		var ret string
		return ret
	}
	return *o.SipSubdomainReceiveSettings
}

// GetSipSubdomainReceiveSettingsOk returns a tuple with the SipSubdomainReceiveSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFaxApplicationRequestInbound) GetSipSubdomainReceiveSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.SipSubdomainReceiveSettings) {
		return nil, false
	}
	return o.SipSubdomainReceiveSettings, true
}

// HasSipSubdomainReceiveSettings returns a boolean if a field has been set.
func (o *CreateFaxApplicationRequestInbound) HasSipSubdomainReceiveSettings() bool {
	if o != nil && !IsNil(o.SipSubdomainReceiveSettings) {
		return true
	}

	return false
}

// SetSipSubdomainReceiveSettings gets a reference to the given string and assigns it to the SipSubdomainReceiveSettings field.
func (o *CreateFaxApplicationRequestInbound) SetSipSubdomainReceiveSettings(v string) {
	o.SipSubdomainReceiveSettings = &v
}

func (o CreateFaxApplicationRequestInbound) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFaxApplicationRequestInbound) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ChannelLimit) {
		toSerialize["channel_limit"] = o.ChannelLimit
	}
	if !IsNil(o.SipSubdomain) {
		toSerialize["sip_subdomain"] = o.SipSubdomain
	}
	if !IsNil(o.SipSubdomainReceiveSettings) {
		toSerialize["sip_subdomain_receive_settings"] = o.SipSubdomainReceiveSettings
	}
	return toSerialize, nil
}

type NullableCreateFaxApplicationRequestInbound struct {
	value *CreateFaxApplicationRequestInbound
	isSet bool
}

func (v NullableCreateFaxApplicationRequestInbound) Get() *CreateFaxApplicationRequestInbound {
	return v.value
}

func (v *NullableCreateFaxApplicationRequestInbound) Set(val *CreateFaxApplicationRequestInbound) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFaxApplicationRequestInbound) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFaxApplicationRequestInbound) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFaxApplicationRequestInbound(val *CreateFaxApplicationRequestInbound) *NullableCreateFaxApplicationRequestInbound {
	return &NullableCreateFaxApplicationRequestInbound{value: val, isSet: true}
}

func (v NullableCreateFaxApplicationRequestInbound) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFaxApplicationRequestInbound) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


