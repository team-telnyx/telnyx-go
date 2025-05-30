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

// checks if the UpdateExternalConnectionRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateExternalConnectionRequest{}

// UpdateExternalConnectionRequest struct for UpdateExternalConnectionRequest
type UpdateExternalConnectionRequest struct {
	// Specifies whether the connection can be used.
	Active *bool `json:"active,omitempty"`
	// The URL where webhooks related to this connection will be sent. Must include a scheme, such as 'https'.
	WebhookEventUrl *string `json:"webhook_event_url,omitempty"`
	// The failover URL where webhooks related to this connection will be sent if sending to the primary URL fails. Must include a scheme, such as 'https'.
	WebhookEventFailoverUrl NullableString `json:"webhook_event_failover_url,omitempty"`
	// Tags associated with the connection.
	Tags []string `json:"tags,omitempty"`
	// Specifies how many seconds to wait before timing out a webhook.
	WebhookTimeoutSecs NullableInt32 `json:"webhook_timeout_secs,omitempty"`
	Inbound *CreateExternalConnectionRequestInbound `json:"inbound,omitempty"`
	Outbound *CreateExternalConnectionRequestOutbound `json:"outbound,omitempty"`
}

// NewUpdateExternalConnectionRequest instantiates a new UpdateExternalConnectionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateExternalConnectionRequest() *UpdateExternalConnectionRequest {
	this := UpdateExternalConnectionRequest{}
	var active bool = true
	this.Active = &active
	var webhookEventFailoverUrl string = ""
	this.WebhookEventFailoverUrl = *NewNullableString(&webhookEventFailoverUrl)
	return &this
}

// NewUpdateExternalConnectionRequestWithDefaults instantiates a new UpdateExternalConnectionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateExternalConnectionRequestWithDefaults() *UpdateExternalConnectionRequest {
	this := UpdateExternalConnectionRequest{}
	var active bool = true
	this.Active = &active
	var webhookEventFailoverUrl string = ""
	this.WebhookEventFailoverUrl = *NewNullableString(&webhookEventFailoverUrl)
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateExternalConnectionRequest) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateExternalConnectionRequest) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *UpdateExternalConnectionRequest) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateExternalConnectionRequest) SetActive(v bool) {
	o.Active = &v
}

// GetWebhookEventUrl returns the WebhookEventUrl field value if set, zero value otherwise.
func (o *UpdateExternalConnectionRequest) GetWebhookEventUrl() string {
	if o == nil || IsNil(o.WebhookEventUrl) {
		var ret string
		return ret
	}
	return *o.WebhookEventUrl
}

// GetWebhookEventUrlOk returns a tuple with the WebhookEventUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateExternalConnectionRequest) GetWebhookEventUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookEventUrl) {
		return nil, false
	}
	return o.WebhookEventUrl, true
}

// HasWebhookEventUrl returns a boolean if a field has been set.
func (o *UpdateExternalConnectionRequest) HasWebhookEventUrl() bool {
	if o != nil && !IsNil(o.WebhookEventUrl) {
		return true
	}

	return false
}

// SetWebhookEventUrl gets a reference to the given string and assigns it to the WebhookEventUrl field.
func (o *UpdateExternalConnectionRequest) SetWebhookEventUrl(v string) {
	o.WebhookEventUrl = &v
}

// GetWebhookEventFailoverUrl returns the WebhookEventFailoverUrl field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateExternalConnectionRequest) GetWebhookEventFailoverUrl() string {
	if o == nil || IsNil(o.WebhookEventFailoverUrl.Get()) {
		var ret string
		return ret
	}
	return *o.WebhookEventFailoverUrl.Get()
}

// GetWebhookEventFailoverUrlOk returns a tuple with the WebhookEventFailoverUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateExternalConnectionRequest) GetWebhookEventFailoverUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebhookEventFailoverUrl.Get(), o.WebhookEventFailoverUrl.IsSet()
}

// HasWebhookEventFailoverUrl returns a boolean if a field has been set.
func (o *UpdateExternalConnectionRequest) HasWebhookEventFailoverUrl() bool {
	if o != nil && o.WebhookEventFailoverUrl.IsSet() {
		return true
	}

	return false
}

// SetWebhookEventFailoverUrl gets a reference to the given NullableString and assigns it to the WebhookEventFailoverUrl field.
func (o *UpdateExternalConnectionRequest) SetWebhookEventFailoverUrl(v string) {
	o.WebhookEventFailoverUrl.Set(&v)
}
// SetWebhookEventFailoverUrlNil sets the value for WebhookEventFailoverUrl to be an explicit nil
func (o *UpdateExternalConnectionRequest) SetWebhookEventFailoverUrlNil() {
	o.WebhookEventFailoverUrl.Set(nil)
}

// UnsetWebhookEventFailoverUrl ensures that no value is present for WebhookEventFailoverUrl, not even an explicit nil
func (o *UpdateExternalConnectionRequest) UnsetWebhookEventFailoverUrl() {
	o.WebhookEventFailoverUrl.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *UpdateExternalConnectionRequest) GetTags() []string {
	if o == nil || IsNil(o.Tags) {
		var ret []string
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateExternalConnectionRequest) GetTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *UpdateExternalConnectionRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *UpdateExternalConnectionRequest) SetTags(v []string) {
	o.Tags = v
}

// GetWebhookTimeoutSecs returns the WebhookTimeoutSecs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateExternalConnectionRequest) GetWebhookTimeoutSecs() int32 {
	if o == nil || IsNil(o.WebhookTimeoutSecs.Get()) {
		var ret int32
		return ret
	}
	return *o.WebhookTimeoutSecs.Get()
}

// GetWebhookTimeoutSecsOk returns a tuple with the WebhookTimeoutSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateExternalConnectionRequest) GetWebhookTimeoutSecsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.WebhookTimeoutSecs.Get(), o.WebhookTimeoutSecs.IsSet()
}

// HasWebhookTimeoutSecs returns a boolean if a field has been set.
func (o *UpdateExternalConnectionRequest) HasWebhookTimeoutSecs() bool {
	if o != nil && o.WebhookTimeoutSecs.IsSet() {
		return true
	}

	return false
}

// SetWebhookTimeoutSecs gets a reference to the given NullableInt32 and assigns it to the WebhookTimeoutSecs field.
func (o *UpdateExternalConnectionRequest) SetWebhookTimeoutSecs(v int32) {
	o.WebhookTimeoutSecs.Set(&v)
}
// SetWebhookTimeoutSecsNil sets the value for WebhookTimeoutSecs to be an explicit nil
func (o *UpdateExternalConnectionRequest) SetWebhookTimeoutSecsNil() {
	o.WebhookTimeoutSecs.Set(nil)
}

// UnsetWebhookTimeoutSecs ensures that no value is present for WebhookTimeoutSecs, not even an explicit nil
func (o *UpdateExternalConnectionRequest) UnsetWebhookTimeoutSecs() {
	o.WebhookTimeoutSecs.Unset()
}

// GetInbound returns the Inbound field value if set, zero value otherwise.
func (o *UpdateExternalConnectionRequest) GetInbound() CreateExternalConnectionRequestInbound {
	if o == nil || IsNil(o.Inbound) {
		var ret CreateExternalConnectionRequestInbound
		return ret
	}
	return *o.Inbound
}

// GetInboundOk returns a tuple with the Inbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateExternalConnectionRequest) GetInboundOk() (*CreateExternalConnectionRequestInbound, bool) {
	if o == nil || IsNil(o.Inbound) {
		return nil, false
	}
	return o.Inbound, true
}

// HasInbound returns a boolean if a field has been set.
func (o *UpdateExternalConnectionRequest) HasInbound() bool {
	if o != nil && !IsNil(o.Inbound) {
		return true
	}

	return false
}

// SetInbound gets a reference to the given CreateExternalConnectionRequestInbound and assigns it to the Inbound field.
func (o *UpdateExternalConnectionRequest) SetInbound(v CreateExternalConnectionRequestInbound) {
	o.Inbound = &v
}

// GetOutbound returns the Outbound field value if set, zero value otherwise.
func (o *UpdateExternalConnectionRequest) GetOutbound() CreateExternalConnectionRequestOutbound {
	if o == nil || IsNil(o.Outbound) {
		var ret CreateExternalConnectionRequestOutbound
		return ret
	}
	return *o.Outbound
}

// GetOutboundOk returns a tuple with the Outbound field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateExternalConnectionRequest) GetOutboundOk() (*CreateExternalConnectionRequestOutbound, bool) {
	if o == nil || IsNil(o.Outbound) {
		return nil, false
	}
	return o.Outbound, true
}

// HasOutbound returns a boolean if a field has been set.
func (o *UpdateExternalConnectionRequest) HasOutbound() bool {
	if o != nil && !IsNil(o.Outbound) {
		return true
	}

	return false
}

// SetOutbound gets a reference to the given CreateExternalConnectionRequestOutbound and assigns it to the Outbound field.
func (o *UpdateExternalConnectionRequest) SetOutbound(v CreateExternalConnectionRequestOutbound) {
	o.Outbound = &v
}

func (o UpdateExternalConnectionRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateExternalConnectionRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.WebhookEventUrl) {
		toSerialize["webhook_event_url"] = o.WebhookEventUrl
	}
	if o.WebhookEventFailoverUrl.IsSet() {
		toSerialize["webhook_event_failover_url"] = o.WebhookEventFailoverUrl.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if o.WebhookTimeoutSecs.IsSet() {
		toSerialize["webhook_timeout_secs"] = o.WebhookTimeoutSecs.Get()
	}
	if !IsNil(o.Inbound) {
		toSerialize["inbound"] = o.Inbound
	}
	if !IsNil(o.Outbound) {
		toSerialize["outbound"] = o.Outbound
	}
	return toSerialize, nil
}

type NullableUpdateExternalConnectionRequest struct {
	value *UpdateExternalConnectionRequest
	isSet bool
}

func (v NullableUpdateExternalConnectionRequest) Get() *UpdateExternalConnectionRequest {
	return v.value
}

func (v *NullableUpdateExternalConnectionRequest) Set(val *UpdateExternalConnectionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateExternalConnectionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateExternalConnectionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateExternalConnectionRequest(val *UpdateExternalConnectionRequest) *NullableUpdateExternalConnectionRequest {
	return &NullableUpdateExternalConnectionRequest{value: val, isSet: true}
}

func (v NullableUpdateExternalConnectionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateExternalConnectionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


