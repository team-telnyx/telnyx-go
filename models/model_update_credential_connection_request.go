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

// checks if the UpdateCredentialConnectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCredentialConnectionRequest{}

// UpdateCredentialConnectionRequest struct for UpdateCredentialConnectionRequest
type UpdateCredentialConnectionRequest struct {
	// Defaults to true
	Active *bool `json:"active,omitempty"`
	// The user name to be used as part of the credentials. Must be 4-32 characters long and alphanumeric values only (no spaces or special characters).
	UserName *string `json:"user_name,omitempty"`
	// The password to be used as part of the credentials. Must be 8 to 128 characters long.
	Password *string `json:"password,omitempty"`
	AnchorsiteOverride *AnchorsiteOverride `json:"anchorsite_override,omitempty"`
	// A user-assigned name to help manage the connection.
	ConnectionName *string `json:"connection_name,omitempty"`
	// This feature enables inbound SIP URI calls to your Credential Auth Connection. If enabled for all (unrestricted) then anyone who calls the SIP URI <your-username>@telnyx.com will be connected to your Connection. You can also choose to allow only calls that are originated on any Connections under your account (internal).
	SipUriCallingPreference *string `json:"sip_uri_calling_preference,omitempty"`
	// When enabled, Telnyx will generate comfort noise when you place the call on hold. If disabled, you will need to generate comfort noise or on hold music to avoid RTP timeout.
	DefaultOnHoldComfortNoiseEnabled *bool `json:"default_on_hold_comfort_noise_enabled,omitempty"`
	DtmfType *DtmfType `json:"dtmf_type,omitempty"`
	// Encode the SIP contact header sent by Telnyx to avoid issues for NAT or ALG scenarios.
	EncodeContactHeaderEnabled *bool `json:"encode_contact_header_enabled,omitempty"`
	EncryptedMedia NullableEncryptedMedia `json:"encrypted_media,omitempty"`
	// Enable on-net T38 if you prefer the sender and receiver negotiating T38 directly if both are on the Telnyx network. If this is disabled, Telnyx will be able to use T38 on just one leg of the call depending on each leg's settings.
	OnnetT38PassthroughEnabled *bool `json:"onnet_t38_passthrough_enabled,omitempty"`
	// The uuid of the push credential for Ios
	IosPushCredentialId NullableString `json:"ios_push_credential_id,omitempty"`
	// The uuid of the push credential for Android
	AndroidPushCredentialId NullableString `json:"android_push_credential_id,omitempty"`
	// The URL where webhooks related to this connection will be sent. Must include a scheme, such as 'https'.
	WebhookEventUrl *string `json:"webhook_event_url,omitempty"`
	// The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverUrl NullableString `json:"webhook_event_failover_url,omitempty"`
	// Determines which webhook format will be used, Telnyx API v1 or v2.
	WebhookApiVersion *string `json:"webhook_api_version,omitempty"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs NullableInt32 `json:"webhook_timeout_secs,omitempty"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitempty"`
	RtcpSettings *ConnectionRtcpSettings `json:"rtcp_settings,omitempty"`
	Inbound *CredentialInbound `json:"inbound,omitempty"`
	Outbound *CredentialOutbound `json:"outbound,omitempty"`
}

// NewUpdateCredentialConnectionRequest instantiates a new UpdateCredentialConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCredentialConnectionRequest() *UpdateCredentialConnectionRequest {
	this := UpdateCredentialConnectionRequest{}
	var anchorsiteOverride AnchorsiteOverride = LATENCY
	this.AnchorsiteOverride = &anchorsiteOverride
	var defaultOnHoldComfortNoiseEnabled bool = false
	this.DefaultOnHoldComfortNoiseEnabled = &defaultOnHoldComfortNoiseEnabled
	var dtmfType DtmfType = RFC_2833
	this.DtmfType = &dtmfType
	var encodeContactHeaderEnabled bool = false
	this.EncodeContactHeaderEnabled = &encodeContactHeaderEnabled
	var onnetT38PassthroughEnabled bool = false
	this.OnnetT38PassthroughEnabled = &onnetT38PassthroughEnabled
	var webhookEventFailoverUrl string = ""
	this.WebhookEventFailoverUrl = *NewNullableString(&webhookEventFailoverUrl)
	var webhookApiVersion string = "1"
	this.WebhookApiVersion = &webhookApiVersion
	return &this
}

// NewUpdateCredentialConnectionRequestWithDefaults instantiates a new UpdateCredentialConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCredentialConnectionRequestWithDefaults() *UpdateCredentialConnectionRequest {
	this := UpdateCredentialConnectionRequest{}
	var anchorsiteOverride AnchorsiteOverride = LATENCY
	this.AnchorsiteOverride = &anchorsiteOverride
	var defaultOnHoldComfortNoiseEnabled bool = false
	this.DefaultOnHoldComfortNoiseEnabled = &defaultOnHoldComfortNoiseEnabled
	var dtmfType DtmfType = RFC_2833
	this.DtmfType = &dtmfType
	var encodeContactHeaderEnabled bool = false
	this.EncodeContactHeaderEnabled = &encodeContactHeaderEnabled
	var onnetT38PassthroughEnabled bool = false
	this.OnnetT38PassthroughEnabled = &onnetT38PassthroughEnabled
	var webhookEventFailoverUrl string = ""
	this.WebhookEventFailoverUrl = *NewNullableString(&webhookEventFailoverUrl)
	var webhookApiVersion string = "1"
	this.WebhookApiVersion = &webhookApiVersion
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateCredentialConnectionRequest) SetActive(v bool) {
	o.Active = &v
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetUserName() string {
	if o == nil || IsNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetUserNameOk() (*string, bool) {
	if o == nil || IsNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasUserName() bool {
	if o != nil && !IsNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *UpdateCredentialConnectionRequest) SetUserName(v string) {
	o.UserName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateCredentialConnectionRequest) SetPassword(v string) {
	o.Password = &v
}

// GetAnchorsiteOverride returns the AnchorsiteOverride field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetAnchorsiteOverride() AnchorsiteOverride {
	if o == nil || IsNil(o.AnchorsiteOverride) {
		var ret AnchorsiteOverride
		return ret
	}
	return *o.AnchorsiteOverride
}

// GetAnchorsiteOverrideOk returns a tuple with the AnchorsiteOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetAnchorsiteOverrideOk() (*AnchorsiteOverride, bool) {
	if o == nil || IsNil(o.AnchorsiteOverride) {
		return nil, false
	}
	return o.AnchorsiteOverride, true
}

// HasAnchorsiteOverride returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasAnchorsiteOverride() bool {
	if o != nil && !IsNil(o.AnchorsiteOverride) {
		return true
	}

	return false
}

// SetAnchorsiteOverride gets a reference to the given AnchorsiteOverride and assigns it to the AnchorsiteOverride field.
func (o *UpdateCredentialConnectionRequest) SetAnchorsiteOverride(v AnchorsiteOverride) {
	o.AnchorsiteOverride = &v
}

// GetConnectionName returns the ConnectionName field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetConnectionName() string {
	if o == nil || IsNil(o.ConnectionName) {
		var ret string
		return ret
	}
	return *o.ConnectionName
}

// GetConnectionNameOk returns a tuple with the ConnectionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetConnectionNameOk() (*string, bool) {
	if o == nil || IsNil(o.ConnectionName) {
		return nil, false
	}
	return o.ConnectionName, true
}

// HasConnectionName returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasConnectionName() bool {
	if o != nil && !IsNil(o.ConnectionName) {
		return true
	}

	return false
}

// SetConnectionName gets a reference to the given string and assigns it to the ConnectionName field.
func (o *UpdateCredentialConnectionRequest) SetConnectionName(v string) {
	o.ConnectionName = &v
}

// GetSipUriCallingPreference returns the SipUriCallingPreference field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetSipUriCallingPreference() string {
	if o == nil || IsNil(o.SipUriCallingPreference) {
		var ret string
		return ret
	}
	return *o.SipUriCallingPreference
}

// GetSipUriCallingPreferenceOk returns a tuple with the SipUriCallingPreference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetSipUriCallingPreferenceOk() (*string, bool) {
	if o == nil || IsNil(o.SipUriCallingPreference) {
		return nil, false
	}
	return o.SipUriCallingPreference, true
}

// HasSipUriCallingPreference returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasSipUriCallingPreference() bool {
	if o != nil && !IsNil(o.SipUriCallingPreference) {
		return true
	}

	return false
}

// SetSipUriCallingPreference gets a reference to the given string and assigns it to the SipUriCallingPreference field.
func (o *UpdateCredentialConnectionRequest) SetSipUriCallingPreference(v string) {
	o.SipUriCallingPreference = &v
}

// GetDefaultOnHoldComfortNoiseEnabled returns the DefaultOnHoldComfortNoiseEnabled field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetDefaultOnHoldComfortNoiseEnabled() bool {
	if o == nil || IsNil(o.DefaultOnHoldComfortNoiseEnabled) {
		var ret bool
		return ret
	}
	return *o.DefaultOnHoldComfortNoiseEnabled
}

// GetDefaultOnHoldComfortNoiseEnabledOk returns a tuple with the DefaultOnHoldComfortNoiseEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetDefaultOnHoldComfortNoiseEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.DefaultOnHoldComfortNoiseEnabled) {
		return nil, false
	}
	return o.DefaultOnHoldComfortNoiseEnabled, true
}

// HasDefaultOnHoldComfortNoiseEnabled returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasDefaultOnHoldComfortNoiseEnabled() bool {
	if o != nil && !IsNil(o.DefaultOnHoldComfortNoiseEnabled) {
		return true
	}

	return false
}

// SetDefaultOnHoldComfortNoiseEnabled gets a reference to the given bool and assigns it to the DefaultOnHoldComfortNoiseEnabled field.
func (o *UpdateCredentialConnectionRequest) SetDefaultOnHoldComfortNoiseEnabled(v bool) {
	o.DefaultOnHoldComfortNoiseEnabled = &v
}

// GetDtmfType returns the DtmfType field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetDtmfType() DtmfType {
	if o == nil || IsNil(o.DtmfType) {
		var ret DtmfType
		return ret
	}
	return *o.DtmfType
}

// GetDtmfTypeOk returns a tuple with the DtmfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetDtmfTypeOk() (*DtmfType, bool) {
	if o == nil || IsNil(o.DtmfType) {
		return nil, false
	}
	return o.DtmfType, true
}

// HasDtmfType returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasDtmfType() bool {
	if o != nil && !IsNil(o.DtmfType) {
		return true
	}

	return false
}

// SetDtmfType gets a reference to the given DtmfType and assigns it to the DtmfType field.
func (o *UpdateCredentialConnectionRequest) SetDtmfType(v DtmfType) {
	o.DtmfType = &v
}

// GetEncodeContactHeaderEnabled returns the EncodeContactHeaderEnabled field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetEncodeContactHeaderEnabled() bool {
	if o == nil || IsNil(o.EncodeContactHeaderEnabled) {
		var ret bool
		return ret
	}
	return *o.EncodeContactHeaderEnabled
}

// GetEncodeContactHeaderEnabledOk returns a tuple with the EncodeContactHeaderEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetEncodeContactHeaderEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.EncodeContactHeaderEnabled) {
		return nil, false
	}
	return o.EncodeContactHeaderEnabled, true
}

// HasEncodeContactHeaderEnabled returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasEncodeContactHeaderEnabled() bool {
	if o != nil && !IsNil(o.EncodeContactHeaderEnabled) {
		return true
	}

	return false
}

// SetEncodeContactHeaderEnabled gets a reference to the given bool and assigns it to the EncodeContactHeaderEnabled field.
func (o *UpdateCredentialConnectionRequest) SetEncodeContactHeaderEnabled(v bool) {
	o.EncodeContactHeaderEnabled = &v
}

// GetEncryptedMedia returns the EncryptedMedia field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCredentialConnectionRequest) GetEncryptedMedia() EncryptedMedia {
	if o == nil || IsNil(o.EncryptedMedia.Get()) {
		var ret EncryptedMedia
		return ret
	}
	return *o.EncryptedMedia.Get()
}

// GetEncryptedMediaOk returns a tuple with the EncryptedMedia field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCredentialConnectionRequest) GetEncryptedMediaOk() (*EncryptedMedia, bool) {
	if o == nil {
		return nil, false
	}
	return o.EncryptedMedia.Get(), o.EncryptedMedia.IsSet()
}

// HasEncryptedMedia returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasEncryptedMedia() bool {
	if o != nil && o.EncryptedMedia.IsSet() {
		return true
	}

	return false
}

// SetEncryptedMedia gets a reference to the given NullableEncryptedMedia and assigns it to the EncryptedMedia field.
func (o *UpdateCredentialConnectionRequest) SetEncryptedMedia(v EncryptedMedia) {
	o.EncryptedMedia.Set(&v)
}
// SetEncryptedMediaNil sets the value for EncryptedMedia to be an explicit nil
func (o *UpdateCredentialConnectionRequest) SetEncryptedMediaNil() {
	o.EncryptedMedia.Set(nil)
}

// UnsetEncryptedMedia ensures that no value is present for EncryptedMedia, not even an explicit nil
func (o *UpdateCredentialConnectionRequest) UnsetEncryptedMedia() {
	o.EncryptedMedia.Unset()
}

// GetOnnetT38PassthroughEnabled returns the OnnetT38PassthroughEnabled field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetOnnetT38PassthroughEnabled() bool {
	if o == nil || IsNil(o.OnnetT38PassthroughEnabled) {
		var ret bool
		return ret
	}
	return *o.OnnetT38PassthroughEnabled
}

// GetOnnetT38PassthroughEnabledOk returns a tuple with the OnnetT38PassthroughEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetOnnetT38PassthroughEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.OnnetT38PassthroughEnabled) {
		return nil, false
	}
	return o.OnnetT38PassthroughEnabled, true
}

// HasOnnetT38PassthroughEnabled returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasOnnetT38PassthroughEnabled() bool {
	if o != nil && !IsNil(o.OnnetT38PassthroughEnabled) {
		return true
	}

	return false
}

// SetOnnetT38PassthroughEnabled gets a reference to the given bool and assigns it to the OnnetT38PassthroughEnabled field.
func (o *UpdateCredentialConnectionRequest) SetOnnetT38PassthroughEnabled(v bool) {
	o.OnnetT38PassthroughEnabled = &v
}

// GetIosPushCredentialId returns the IosPushCredentialId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCredentialConnectionRequest) GetIosPushCredentialId() string {
	if o == nil || IsNil(o.IosPushCredentialId.Get()) {
		var ret string
		return ret
	}
	return *o.IosPushCredentialId.Get()
}

// GetIosPushCredentialIdOk returns a tuple with the IosPushCredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCredentialConnectionRequest) GetIosPushCredentialIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IosPushCredentialId.Get(), o.IosPushCredentialId.IsSet()
}

// HasIosPushCredentialId returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasIosPushCredentialId() bool {
	if o != nil && o.IosPushCredentialId.IsSet() {
		return true
	}

	return false
}

// SetIosPushCredentialId gets a reference to the given NullableString and assigns it to the IosPushCredentialId field.
func (o *UpdateCredentialConnectionRequest) SetIosPushCredentialId(v string) {
	o.IosPushCredentialId.Set(&v)
}
// SetIosPushCredentialIdNil sets the value for IosPushCredentialId to be an explicit nil
func (o *UpdateCredentialConnectionRequest) SetIosPushCredentialIdNil() {
	o.IosPushCredentialId.Set(nil)
}

// UnsetIosPushCredentialId ensures that no value is present for IosPushCredentialId, not even an explicit nil
func (o *UpdateCredentialConnectionRequest) UnsetIosPushCredentialId() {
	o.IosPushCredentialId.Unset()
}

// GetAndroidPushCredentialId returns the AndroidPushCredentialId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCredentialConnectionRequest) GetAndroidPushCredentialId() string {
	if o == nil || IsNil(o.AndroidPushCredentialId.Get()) {
		var ret string
		return ret
	}
	return *o.AndroidPushCredentialId.Get()
}

// GetAndroidPushCredentialIdOk returns a tuple with the AndroidPushCredentialId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCredentialConnectionRequest) GetAndroidPushCredentialIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AndroidPushCredentialId.Get(), o.AndroidPushCredentialId.IsSet()
}

// HasAndroidPushCredentialId returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasAndroidPushCredentialId() bool {
	if o != nil && o.AndroidPushCredentialId.IsSet() {
		return true
	}

	return false
}

// SetAndroidPushCredentialId gets a reference to the given NullableString and assigns it to the AndroidPushCredentialId field.
func (o *UpdateCredentialConnectionRequest) SetAndroidPushCredentialId(v string) {
	o.AndroidPushCredentialId.Set(&v)
}
// SetAndroidPushCredentialIdNil sets the value for AndroidPushCredentialId to be an explicit nil
func (o *UpdateCredentialConnectionRequest) SetAndroidPushCredentialIdNil() {
	o.AndroidPushCredentialId.Set(nil)
}

// UnsetAndroidPushCredentialId ensures that no value is present for AndroidPushCredentialId, not even an explicit nil
func (o *UpdateCredentialConnectionRequest) UnsetAndroidPushCredentialId() {
	o.AndroidPushCredentialId.Unset()
}

// GetWebhookEventUrl returns the WebhookEventUrl field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetWebhookEventUrl() string {
	if o == nil || IsNil(o.WebhookEventUrl) {
		var ret string
		return ret
	}
	return *o.WebhookEventUrl
}

// GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetWebhookEventUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookEventUrl) {
		return nil, false
	}
	return o.WebhookEventUrl, true
}

// HasWebhookEventUrl returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasWebhookEventUrl() bool {
	if o != nil && !IsNil(o.WebhookEventUrl) {
		return true
	}

	return false
}

// SetWebhookEventUrl gets a reference to the given string and assigns it to the WebhookEventUrl field.
func (o *UpdateCredentialConnectionRequest) SetWebhookEventUrl(v string) {
	o.WebhookEventUrl = &v
}

// GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCredentialConnectionRequest) GetWebhookEventFailoverUrl() string {
	if o == nil || IsNil(o.WebhookEventFailoverUrl.Get()) {
		var ret string
		return ret
	}
	return *o.WebhookEventFailoverUrl.Get()
}

// GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCredentialConnectionRequest) GetWebhookEventFailoverUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebhookEventFailoverUrl.Get(), o.WebhookEventFailoverUrl.IsSet()
}

// HasWebhookEventFailoverUrl returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasWebhookEventFailoverUrl() bool {
	if o != nil && o.WebhookEventFailoverUrl.IsSet() {
		return true
	}

	return false
}

// SetWebhookEventFailoverUrl gets a reference to the given NullableString and assigns it to the WebhookEventFailoverUrl field.
func (o *UpdateCredentialConnectionRequest) SetWebhookEventFailoverUrl(v string) {
	o.WebhookEventFailoverUrl.Set(&v)
}
// SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil
func (o *UpdateCredentialConnectionRequest) SetWebhookEventFailoverUrlNil() {
	o.WebhookEventFailoverUrl.Set(nil)
}

// UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
func (o *UpdateCredentialConnectionRequest) UnsetWebhookEventFailoverUrl() {
	o.WebhookEventFailoverUrl.Unset()
}

// GetWebhookApiVersion returns the WebhookApiVersion field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetWebhookApiVersion() string {
	if o == nil || IsNil(o.WebhookApiVersion) {
		var ret string
		return ret
	}
	return *o.WebhookApiVersion
}

// GetWebhookApiVersionOk returns a tuple with the WebhookApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetWebhookApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookApiVersion) {
		return nil, false
	}
	return o.WebhookApiVersion, true
}

// HasWebhookApiVersion returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasWebhookApiVersion() bool {
	if o != nil && !IsNil(o.WebhookApiVersion) {
		return true
	}

	return false
}

// SetWebhookApiVersion gets a reference to the given string and assigns it to the WebhookApiVersion field.
func (o *UpdateCredentialConnectionRequest) SetWebhookApiVersion(v string) {
	o.WebhookApiVersion = &v
}

// GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateCredentialConnectionRequest) GetWebhookTimeoutSecs() int32 {
	if o == nil || IsNil(o.WebhookTimeoutSecs.Get()) {
		var ret int32
		return ret
	}
	return *o.WebhookTimeoutSecs.Get()
}

// GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateCredentialConnectionRequest) GetWebhookTimeoutSecsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebhookTimeoutSecs.Get(), o.WebhookTimeoutSecs.IsSet()
}

// HasWebhookTimeoutSecs returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasWebhookTimeoutSecs() bool {
	if o != nil && o.WebhookTimeoutSecs.IsSet() {
		return true
	}

	return false
}

// SetWebhookTimeoutSecs gets a reference to the given NullableInt32 and assigns it to the WebhookTimeoutSecs field.
func (o *UpdateCredentialConnectionRequest) SetWebhookTimeoutSecs(v int32) {
	o.WebhookTimeoutSecs.Set(&v)
}
// SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil
func (o *UpdateCredentialConnectionRequest) SetWebhookTimeoutSecsNil() {
	o.WebhookTimeoutSecs.Set(nil)
}

// UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
func (o *UpdateCredentialConnectionRequest) UnsetWebhookTimeoutSecs() {
	o.WebhookTimeoutSecs.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateCredentialConnectionRequest) SetTags(v []string) {
	o.Tags = v
}

// GetRtcpSettings returns the RtcpSettings field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetRtcpSettings() ConnectionRtcpSettings {
	if o == nil || IsNil(o.RtcpSettings) {
		var ret ConnectionRtcpSettings
		return ret
	}
	return *o.RtcpSettings
}

// GetRtcpSettingsOk returns a tuple with the RtcpSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetRtcpSettingsOk() (*ConnectionRtcpSettings, bool) {
	if o == nil || IsNil(o.RtcpSettings) {
		return nil, false
	}
	return o.RtcpSettings, true
}

// HasRtcpSettings returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasRtcpSettings() bool {
	if o != nil && !IsNil(o.RtcpSettings) {
		return true
	}

	return false
}

// SetRtcpSettings gets a reference to the given ConnectionRtcpSettings and assigns it to the RtcpSettings field.
func (o *UpdateCredentialConnectionRequest) SetRtcpSettings(v ConnectionRtcpSettings) {
	o.RtcpSettings = &v
}

// GetInbound returns the Inbound field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetInbound() CredentialInbound {
	if o == nil || IsNil(o.Inbound) {
		var ret CredentialInbound
		return ret
	}
	return *o.Inbound
}

// GetInboundOk returns a tuple with the Inbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetInboundOk() (*CredentialInbound, bool) {
	if o == nil || IsNil(o.Inbound) {
		return nil, false
	}
	return o.Inbound, true
}

// HasInbound returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasInbound() bool {
	if o != nil && !IsNil(o.Inbound) {
		return true
	}

	return false
}

// SetInbound gets a reference to the given CredentialInbound and assigns it to the Inbound field.
func (o *UpdateCredentialConnectionRequest) SetInbound(v CredentialInbound) {
	o.Inbound = &v
}

// GetOutbound returns the Outbound field value if set, zero value otherwise.
func (o *UpdateCredentialConnectionRequest) GetOutbound() CredentialOutbound {
	if o == nil || IsNil(o.Outbound) {
		var ret CredentialOutbound
		return ret
	}
	return *o.Outbound
}

// GetOutboundOk returns a tuple with the Outbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCredentialConnectionRequest) GetOutboundOk() (*CredentialOutbound, bool) {
	if o == nil || IsNil(o.Outbound) {
		return nil, false
	}
	return o.Outbound, true
}

// HasOutbound returns a boolean if a field has been set.
func (o *UpdateCredentialConnectionRequest) HasOutbound() bool {
	if o != nil && !IsNil(o.Outbound) {
		return true
	}

	return false
}

// SetOutbound gets a reference to the given CredentialOutbound and assigns it to the Outbound field.
func (o *UpdateCredentialConnectionRequest) SetOutbound(v CredentialOutbound) {
	o.Outbound = &v
}

func (o UpdateCredentialConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCredentialConnectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.UserName) {
		toSerialize["user_name"] = o.UserName
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.AnchorsiteOverride) {
		toSerialize["anchorsite_override"] = o.AnchorsiteOverride
	}
	if !IsNil(o.ConnectionName) {
		toSerialize["connection_name"] = o.ConnectionName
	}
	if !IsNil(o.SipUriCallingPreference) {
		toSerialize["sip_uri_calling_preference"] = o.SipUriCallingPreference
	}
	if !IsNil(o.DefaultOnHoldComfortNoiseEnabled) {
		toSerialize["default_on_hold_comfort_noise_enabled"] = o.DefaultOnHoldComfortNoiseEnabled
	}
	if !IsNil(o.DtmfType) {
		toSerialize["dtmf_type"] = o.DtmfType
	}
	if !IsNil(o.EncodeContactHeaderEnabled) {
		toSerialize["encode_contact_header_enabled"] = o.EncodeContactHeaderEnabled
	}
	if o.EncryptedMedia.IsSet() {
		toSerialize["encrypted_media"] = o.EncryptedMedia.Get()
	}
	if !IsNil(o.OnnetT38PassthroughEnabled) {
		toSerialize["onnet_t38_passthrough_enabled"] = o.OnnetT38PassthroughEnabled
	}
	if o.IosPushCredentialId.IsSet() {
		toSerialize["ios_push_credential_id"] = o.IosPushCredentialId.Get()
	}
	if o.AndroidPushCredentialId.IsSet() {
		toSerialize["android_push_credential_id"] = o.AndroidPushCredentialId.Get()
	}
	if !IsNil(o.WebhookEventUrl) {
		toSerialize["webhook_event_url"] = o.WebhookEventUrl
	}
	if o.WebhookEventFailoverUrl.IsSet() {
		toSerialize["webhook_event_failover_url"] = o.WebhookEventFailoverUrl.Get()
	}
	if !IsNil(o.WebhookApiVersion) {
		toSerialize["webhook_api_version"] = o.WebhookApiVersion
	}
	if o.WebhookTimeoutSecs.IsSet() {
		toSerialize["webhook_timeout_secs"] = o.WebhookTimeoutSecs.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.RtcpSettings) {
		toSerialize["rtcp_settings"] = o.RtcpSettings
	}
	if !IsNil(o.Inbound) {
		toSerialize["inbound"] = o.Inbound
	}
	if !IsNil(o.Outbound) {
		toSerialize["outbound"] = o.Outbound
	}
	return toSerialize, nil
}

type NullableUpdateCredentialConnectionRequest struct {
	value *UpdateCredentialConnectionRequest
	isSet bool
}

func (v NullableUpdateCredentialConnectionRequest) Get() *UpdateCredentialConnectionRequest {
	return v.value
}

func (v *NullableUpdateCredentialConnectionRequest) Set(val *UpdateCredentialConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateCredentialConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateCredentialConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateCredentialConnectionRequest(val *UpdateCredentialConnectionRequest) *NullableUpdateCredentialConnectionRequest {
	return &NullableUpdateCredentialConnectionRequest{value: val, isSet: true}
}

func (v NullableUpdateCredentialConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateCredentialConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


