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

// checks if the CreateInboundIpRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateInboundIpRequest{}

// CreateInboundIpRequest struct for CreateInboundIpRequest
type CreateInboundIpRequest struct {
	// This setting allows you to set the format with which the caller's number (ANI) is sent for inbound phone calls.
	AniNumberFormat *string `json:"ani_number_format,omitempty"`
	DnisNumberFormat *string `json:"dnis_number_format,omitempty"`
	// Defines the list of codecs that Telnyx will send for inbound calls to a specific number on your portal account, in priority order. This only works when the Connection the number is assigned to uses Media Handling mode: default. OPUS and H.264 codecs are available only when using TCP or TLS transport for SIP.
	Codecs []string `json:"codecs,omitempty"`
	// Default routing method to be used when a number is associated with the connection. Must be one of the routing method types or left blank, other values are not allowed.
	DefaultRoutingMethod *string `json:"default_routing_method,omitempty"`
	// When set, this will limit the total number of inbound calls to phone numbers associated with this connection.
	ChannelLimit *int32 `json:"channel_limit,omitempty"`
	// Generate ringback tone through 183 session progress message with early media.
	GenerateRingbackTone *bool `json:"generate_ringback_tone,omitempty"`
	// When set, inbound phone calls will receive ISUP parameters via SIP headers. (Only when available and only when using TCP or TLS transport.)
	IsupHeadersEnabled *bool `json:"isup_headers_enabled,omitempty"`
	// Enable PRACK messages as defined in RFC3262.
	PrackEnabled *bool `json:"prack_enabled,omitempty"`
	// Defaults to true.
	SipCompactHeadersEnabled *bool `json:"sip_compact_headers_enabled,omitempty"`
	// Selects which `sip_region` to receive inbound calls from. If null, the default region (US) will be used.
	SipRegion *string `json:"sip_region,omitempty"`
	// Specifies a subdomain that can be used to receive Inbound calls to a Connection, in the same way a phone number is used, from a SIP endpoint. Example: the subdomain \"example.sip.telnyx.com\" can be called from any SIP endpoint by using the SIP URI \"sip:@example.sip.telnyx.com\" where the user part can be any alphanumeric value. Please note TLS encrypted calls are not allowed for subdomain calls.
	SipSubdomain *string `json:"sip_subdomain,omitempty"`
	// This option can be enabled to receive calls from: \"Anyone\" (any SIP endpoint in the public Internet) or \"Only my connections\" (any connection assigned to the same Telnyx user).
	SipSubdomainReceiveSettings *string `json:"sip_subdomain_receive_settings,omitempty"`
	// Time(sec) before aborting if connection is not made.
	Timeout1xxSecs *int32 `json:"timeout_1xx_secs,omitempty"`
	// Time(sec) before aborting if call is unanswered (min: 1, max: 600).
	Timeout2xxSecs *int32 `json:"timeout_2xx_secs,omitempty"`
	// When enabled the SIP Connection will receive the Identity header with Shaken/Stir data in the SIP INVITE message of inbound calls, even when using UDP transport.
	ShakenStirEnabled *bool `json:"shaken_stir_enabled,omitempty"`
}

// NewCreateInboundIpRequest instantiates a new CreateInboundIpRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateInboundIpRequest() *CreateInboundIpRequest {
	this := CreateInboundIpRequest{}
	var aniNumberFormat string = "E.164-national"
	this.AniNumberFormat = &aniNumberFormat
	var dnisNumberFormat string = "e164"
	this.DnisNumberFormat = &dnisNumberFormat
	var generateRingbackTone bool = false
	this.GenerateRingbackTone = &generateRingbackTone
	var isupHeadersEnabled bool = false
	this.IsupHeadersEnabled = &isupHeadersEnabled
	var prackEnabled bool = false
	this.PrackEnabled = &prackEnabled
	var sipCompactHeadersEnabled bool = true
	this.SipCompactHeadersEnabled = &sipCompactHeadersEnabled
	var sipRegion string = "US"
	this.SipRegion = &sipRegion
	var timeout1xxSecs int32 = 3
	this.Timeout1xxSecs = &timeout1xxSecs
	var timeout2xxSecs int32 = 90
	this.Timeout2xxSecs = &timeout2xxSecs
	var shakenStirEnabled bool = false
	this.ShakenStirEnabled = &shakenStirEnabled
	return &this
}

// NewCreateInboundIpRequestWithDefaults instantiates a new CreateInboundIpRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateInboundIpRequestWithDefaults() *CreateInboundIpRequest {
	this := CreateInboundIpRequest{}
	var aniNumberFormat string = "E.164-national"
	this.AniNumberFormat = &aniNumberFormat
	var dnisNumberFormat string = "e164"
	this.DnisNumberFormat = &dnisNumberFormat
	var generateRingbackTone bool = false
	this.GenerateRingbackTone = &generateRingbackTone
	var isupHeadersEnabled bool = false
	this.IsupHeadersEnabled = &isupHeadersEnabled
	var prackEnabled bool = false
	this.PrackEnabled = &prackEnabled
	var sipCompactHeadersEnabled bool = true
	this.SipCompactHeadersEnabled = &sipCompactHeadersEnabled
	var sipRegion string = "US"
	this.SipRegion = &sipRegion
	var timeout1xxSecs int32 = 3
	this.Timeout1xxSecs = &timeout1xxSecs
	var timeout2xxSecs int32 = 90
	this.Timeout2xxSecs = &timeout2xxSecs
	var shakenStirEnabled bool = false
	this.ShakenStirEnabled = &shakenStirEnabled
	return &this
}

// GetAniNumberFormat returns the AniNumberFormat field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetAniNumberFormat() string {
	if o == nil || IsNil(o.AniNumberFormat) {
		var ret string
		return ret
	}
	return *o.AniNumberFormat
}

// GetAniNumberFormatOk returns a tuple with the AniNumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetAniNumberFormatOk() (*string, bool) {
	if o == nil || IsNil(o.AniNumberFormat) {
		return nil, false
	}
	return o.AniNumberFormat, true
}

// HasAniNumberFormat returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasAniNumberFormat() bool {
	if o != nil && !IsNil(o.AniNumberFormat) {
		return true
	}

	return false
}

// SetAniNumberFormat gets a reference to the given string and assigns it to the AniNumberFormat field.
func (o *CreateInboundIpRequest) SetAniNumberFormat(v string) {
	o.AniNumberFormat = &v
}

// GetDnisNumberFormat returns the DnisNumberFormat field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetDnisNumberFormat() string {
	if o == nil || IsNil(o.DnisNumberFormat) {
		var ret string
		return ret
	}
	return *o.DnisNumberFormat
}

// GetDnisNumberFormatOk returns a tuple with the DnisNumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetDnisNumberFormatOk() (*string, bool) {
	if o == nil || IsNil(o.DnisNumberFormat) {
		return nil, false
	}
	return o.DnisNumberFormat, true
}

// HasDnisNumberFormat returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasDnisNumberFormat() bool {
	if o != nil && !IsNil(o.DnisNumberFormat) {
		return true
	}

	return false
}

// SetDnisNumberFormat gets a reference to the given string and assigns it to the DnisNumberFormat field.
func (o *CreateInboundIpRequest) SetDnisNumberFormat(v string) {
	o.DnisNumberFormat = &v
}

// GetCodecs returns the Codecs field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetCodecs() []string {
	if o == nil || IsNil(o.Codecs) {
		var ret []string
		return ret
	}
	return o.Codecs
}

// GetCodecsOk returns a tuple with the Codecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetCodecsOk() ([]string, bool) {
	if o == nil || IsNil(o.Codecs) {
		return nil, false
	}
	return o.Codecs, true
}

// HasCodecs returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasCodecs() bool {
	if o != nil && !IsNil(o.Codecs) {
		return true
	}

	return false
}

// SetCodecs gets a reference to the given []string and assigns it to the Codecs field.
func (o *CreateInboundIpRequest) SetCodecs(v []string) {
	o.Codecs = v
}

// GetDefaultRoutingMethod returns the DefaultRoutingMethod field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetDefaultRoutingMethod() string {
	if o == nil || IsNil(o.DefaultRoutingMethod) {
		var ret string
		return ret
	}
	return *o.DefaultRoutingMethod
}

// GetDefaultRoutingMethodOk returns a tuple with the DefaultRoutingMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetDefaultRoutingMethodOk() (*string, bool) {
	if o == nil || IsNil(o.DefaultRoutingMethod) {
		return nil, false
	}
	return o.DefaultRoutingMethod, true
}

// HasDefaultRoutingMethod returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasDefaultRoutingMethod() bool {
	if o != nil && !IsNil(o.DefaultRoutingMethod) {
		return true
	}

	return false
}

// SetDefaultRoutingMethod gets a reference to the given string and assigns it to the DefaultRoutingMethod field.
func (o *CreateInboundIpRequest) SetDefaultRoutingMethod(v string) {
	o.DefaultRoutingMethod = &v
}

// GetChannelLimit returns the ChannelLimit field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetChannelLimit() int32 {
	if o == nil || IsNil(o.ChannelLimit) {
		var ret int32
		return ret
	}
	return *o.ChannelLimit
}

// GetChannelLimitOk returns a tuple with the ChannelLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetChannelLimitOk() (*int32, bool) {
	if o == nil || IsNil(o.ChannelLimit) {
		return nil, false
	}
	return o.ChannelLimit, true
}

// HasChannelLimit returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasChannelLimit() bool {
	if o != nil && !IsNil(o.ChannelLimit) {
		return true
	}

	return false
}

// SetChannelLimit gets a reference to the given int32 and assigns it to the ChannelLimit field.
func (o *CreateInboundIpRequest) SetChannelLimit(v int32) {
	o.ChannelLimit = &v
}

// GetGenerateRingbackTone returns the GenerateRingbackTone field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetGenerateRingbackTone() bool {
	if o == nil || IsNil(o.GenerateRingbackTone) {
		var ret bool
		return ret
	}
	return *o.GenerateRingbackTone
}

// GetGenerateRingbackToneOk returns a tuple with the GenerateRingbackTone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetGenerateRingbackToneOk() (*bool, bool) {
	if o == nil || IsNil(o.GenerateRingbackTone) {
		return nil, false
	}
	return o.GenerateRingbackTone, true
}

// HasGenerateRingbackTone returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasGenerateRingbackTone() bool {
	if o != nil && !IsNil(o.GenerateRingbackTone) {
		return true
	}

	return false
}

// SetGenerateRingbackTone gets a reference to the given bool and assigns it to the GenerateRingbackTone field.
func (o *CreateInboundIpRequest) SetGenerateRingbackTone(v bool) {
	o.GenerateRingbackTone = &v
}

// GetIsupHeadersEnabled returns the IsupHeadersEnabled field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetIsupHeadersEnabled() bool {
	if o == nil || IsNil(o.IsupHeadersEnabled) {
		var ret bool
		return ret
	}
	return *o.IsupHeadersEnabled
}

// GetIsupHeadersEnabledOk returns a tuple with the IsupHeadersEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetIsupHeadersEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.IsupHeadersEnabled) {
		return nil, false
	}
	return o.IsupHeadersEnabled, true
}

// HasIsupHeadersEnabled returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasIsupHeadersEnabled() bool {
	if o != nil && !IsNil(o.IsupHeadersEnabled) {
		return true
	}

	return false
}

// SetIsupHeadersEnabled gets a reference to the given bool and assigns it to the IsupHeadersEnabled field.
func (o *CreateInboundIpRequest) SetIsupHeadersEnabled(v bool) {
	o.IsupHeadersEnabled = &v
}

// GetPrackEnabled returns the PrackEnabled field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetPrackEnabled() bool {
	if o == nil || IsNil(o.PrackEnabled) {
		var ret bool
		return ret
	}
	return *o.PrackEnabled
}

// GetPrackEnabledOk returns a tuple with the PrackEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetPrackEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.PrackEnabled) {
		return nil, false
	}
	return o.PrackEnabled, true
}

// HasPrackEnabled returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasPrackEnabled() bool {
	if o != nil && !IsNil(o.PrackEnabled) {
		return true
	}

	return false
}

// SetPrackEnabled gets a reference to the given bool and assigns it to the PrackEnabled field.
func (o *CreateInboundIpRequest) SetPrackEnabled(v bool) {
	o.PrackEnabled = &v
}

// GetSipCompactHeadersEnabled returns the SipCompactHeadersEnabled field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetSipCompactHeadersEnabled() bool {
	if o == nil || IsNil(o.SipCompactHeadersEnabled) {
		var ret bool
		return ret
	}
	return *o.SipCompactHeadersEnabled
}

// GetSipCompactHeadersEnabledOk returns a tuple with the SipCompactHeadersEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetSipCompactHeadersEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.SipCompactHeadersEnabled) {
		return nil, false
	}
	return o.SipCompactHeadersEnabled, true
}

// HasSipCompactHeadersEnabled returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasSipCompactHeadersEnabled() bool {
	if o != nil && !IsNil(o.SipCompactHeadersEnabled) {
		return true
	}

	return false
}

// SetSipCompactHeadersEnabled gets a reference to the given bool and assigns it to the SipCompactHeadersEnabled field.
func (o *CreateInboundIpRequest) SetSipCompactHeadersEnabled(v bool) {
	o.SipCompactHeadersEnabled = &v
}

// GetSipRegion returns the SipRegion field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetSipRegion() string {
	if o == nil || IsNil(o.SipRegion) {
		var ret string
		return ret
	}
	return *o.SipRegion
}

// GetSipRegionOk returns a tuple with the SipRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetSipRegionOk() (*string, bool) {
	if o == nil || IsNil(o.SipRegion) {
		return nil, false
	}
	return o.SipRegion, true
}

// HasSipRegion returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasSipRegion() bool {
	if o != nil && !IsNil(o.SipRegion) {
		return true
	}

	return false
}

// SetSipRegion gets a reference to the given string and assigns it to the SipRegion field.
func (o *CreateInboundIpRequest) SetSipRegion(v string) {
	o.SipRegion = &v
}

// GetSipSubdomain returns the SipSubdomain field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetSipSubdomain() string {
	if o == nil || IsNil(o.SipSubdomain) {
		var ret string
		return ret
	}
	return *o.SipSubdomain
}

// GetSipSubdomainOk returns a tuple with the SipSubdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetSipSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.SipSubdomain) {
		return nil, false
	}
	return o.SipSubdomain, true
}

// HasSipSubdomain returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasSipSubdomain() bool {
	if o != nil && !IsNil(o.SipSubdomain) {
		return true
	}

	return false
}

// SetSipSubdomain gets a reference to the given string and assigns it to the SipSubdomain field.
func (o *CreateInboundIpRequest) SetSipSubdomain(v string) {
	o.SipSubdomain = &v
}

// GetSipSubdomainReceiveSettings returns the SipSubdomainReceiveSettings field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetSipSubdomainReceiveSettings() string {
	if o == nil || IsNil(o.SipSubdomainReceiveSettings) {
		var ret string
		return ret
	}
	return *o.SipSubdomainReceiveSettings
}

// GetSipSubdomainReceiveSettingsOk returns a tuple with the SipSubdomainReceiveSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetSipSubdomainReceiveSettingsOk() (*string, bool) {
	if o == nil || IsNil(o.SipSubdomainReceiveSettings) {
		return nil, false
	}
	return o.SipSubdomainReceiveSettings, true
}

// HasSipSubdomainReceiveSettings returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasSipSubdomainReceiveSettings() bool {
	if o != nil && !IsNil(o.SipSubdomainReceiveSettings) {
		return true
	}

	return false
}

// SetSipSubdomainReceiveSettings gets a reference to the given string and assigns it to the SipSubdomainReceiveSettings field.
func (o *CreateInboundIpRequest) SetSipSubdomainReceiveSettings(v string) {
	o.SipSubdomainReceiveSettings = &v
}

// GetTimeout1xxSecs returns the Timeout1xxSecs field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetTimeout1xxSecs() int32 {
	if o == nil || IsNil(o.Timeout1xxSecs) {
		var ret int32
		return ret
	}
	return *o.Timeout1xxSecs
}

// GetTimeout1xxSecsOk returns a tuple with the Timeout1xxSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetTimeout1xxSecsOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout1xxSecs) {
		return nil, false
	}
	return o.Timeout1xxSecs, true
}

// HasTimeout1xxSecs returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasTimeout1xxSecs() bool {
	if o != nil && !IsNil(o.Timeout1xxSecs) {
		return true
	}

	return false
}

// SetTimeout1xxSecs gets a reference to the given int32 and assigns it to the Timeout1xxSecs field.
func (o *CreateInboundIpRequest) SetTimeout1xxSecs(v int32) {
	o.Timeout1xxSecs = &v
}

// GetTimeout2xxSecs returns the Timeout2xxSecs field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetTimeout2xxSecs() int32 {
	if o == nil || IsNil(o.Timeout2xxSecs) {
		var ret int32
		return ret
	}
	return *o.Timeout2xxSecs
}

// GetTimeout2xxSecsOk returns a tuple with the Timeout2xxSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetTimeout2xxSecsOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout2xxSecs) {
		return nil, false
	}
	return o.Timeout2xxSecs, true
}

// HasTimeout2xxSecs returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasTimeout2xxSecs() bool {
	if o != nil && !IsNil(o.Timeout2xxSecs) {
		return true
	}

	return false
}

// SetTimeout2xxSecs gets a reference to the given int32 and assigns it to the Timeout2xxSecs field.
func (o *CreateInboundIpRequest) SetTimeout2xxSecs(v int32) {
	o.Timeout2xxSecs = &v
}

// GetShakenStirEnabled returns the ShakenStirEnabled field value if set, zero value otherwise.
func (o *CreateInboundIpRequest) GetShakenStirEnabled() bool {
	if o == nil || IsNil(o.ShakenStirEnabled) {
		var ret bool
		return ret
	}
	return *o.ShakenStirEnabled
}

// GetShakenStirEnabledOk returns a tuple with the ShakenStirEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateInboundIpRequest) GetShakenStirEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ShakenStirEnabled) {
		return nil, false
	}
	return o.ShakenStirEnabled, true
}

// HasShakenStirEnabled returns a boolean if a field has been set.
func (o *CreateInboundIpRequest) HasShakenStirEnabled() bool {
	if o != nil && !IsNil(o.ShakenStirEnabled) {
		return true
	}

	return false
}

// SetShakenStirEnabled gets a reference to the given bool and assigns it to the ShakenStirEnabled field.
func (o *CreateInboundIpRequest) SetShakenStirEnabled(v bool) {
	o.ShakenStirEnabled = &v
}

func (o CreateInboundIpRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateInboundIpRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AniNumberFormat) {
		toSerialize["ani_number_format"] = o.AniNumberFormat
	}
	if !IsNil(o.DnisNumberFormat) {
		toSerialize["dnis_number_format"] = o.DnisNumberFormat
	}
	if !IsNil(o.Codecs) {
		toSerialize["codecs"] = o.Codecs
	}
	if !IsNil(o.DefaultRoutingMethod) {
		toSerialize["default_routing_method"] = o.DefaultRoutingMethod
	}
	if !IsNil(o.ChannelLimit) {
		toSerialize["channel_limit"] = o.ChannelLimit
	}
	if !IsNil(o.GenerateRingbackTone) {
		toSerialize["generate_ringback_tone"] = o.GenerateRingbackTone
	}
	if !IsNil(o.IsupHeadersEnabled) {
		toSerialize["isup_headers_enabled"] = o.IsupHeadersEnabled
	}
	if !IsNil(o.PrackEnabled) {
		toSerialize["prack_enabled"] = o.PrackEnabled
	}
	if !IsNil(o.SipCompactHeadersEnabled) {
		toSerialize["sip_compact_headers_enabled"] = o.SipCompactHeadersEnabled
	}
	if !IsNil(o.SipRegion) {
		toSerialize["sip_region"] = o.SipRegion
	}
	if !IsNil(o.SipSubdomain) {
		toSerialize["sip_subdomain"] = o.SipSubdomain
	}
	if !IsNil(o.SipSubdomainReceiveSettings) {
		toSerialize["sip_subdomain_receive_settings"] = o.SipSubdomainReceiveSettings
	}
	if !IsNil(o.Timeout1xxSecs) {
		toSerialize["timeout_1xx_secs"] = o.Timeout1xxSecs
	}
	if !IsNil(o.Timeout2xxSecs) {
		toSerialize["timeout_2xx_secs"] = o.Timeout2xxSecs
	}
	if !IsNil(o.ShakenStirEnabled) {
		toSerialize["shaken_stir_enabled"] = o.ShakenStirEnabled
	}
	return toSerialize, nil
}

type NullableCreateInboundIpRequest struct {
	value *CreateInboundIpRequest
	isSet bool
}

func (v NullableCreateInboundIpRequest) Get() *CreateInboundIpRequest {
	return v.value
}

func (v *NullableCreateInboundIpRequest) Set(val *CreateInboundIpRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateInboundIpRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateInboundIpRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateInboundIpRequest(val *CreateInboundIpRequest) *NullableCreateInboundIpRequest {
	return &NullableCreateInboundIpRequest{value: val, isSet: true}
}

func (v NullableCreateInboundIpRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateInboundIpRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


