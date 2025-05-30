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

// checks if the UpdatePartnerCampaignRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePartnerCampaignRequest{}

// UpdatePartnerCampaignRequest struct for UpdatePartnerCampaignRequest
type UpdatePartnerCampaignRequest struct {
	// Webhook to which campaign status updates are sent.
	WebhookURL *string `json:"webhookURL,omitempty"`
	// Webhook failover to which campaign status updates are sent.
	WebhookFailoverURL *string `json:"webhookFailoverURL,omitempty"`
}

// NewUpdatePartnerCampaignRequest instantiates a new UpdatePartnerCampaignRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePartnerCampaignRequest() *UpdatePartnerCampaignRequest {
	this := UpdatePartnerCampaignRequest{}
	return &this
}

// NewUpdatePartnerCampaignRequestWithDefaults instantiates a new UpdatePartnerCampaignRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePartnerCampaignRequestWithDefaults() *UpdatePartnerCampaignRequest {
	this := UpdatePartnerCampaignRequest{}
	return &this
}

// GetWebhookURL returns the WebhookURL field value if set, zero value otherwise.
func (o *UpdatePartnerCampaignRequest) GetWebhookURL() string {
	if o == nil || IsNil(o.WebhookURL) {
		var ret string
		return ret
	}
	return *o.WebhookURL
}

// GetWebhookURLOk returns a tuple with the WebhookURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePartnerCampaignRequest) GetWebhookURLOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookURL) {
		return nil, false
	}
	return o.WebhookURL, true
}

// HasWebhookURL returns a boolean if a field has been set.
func (o *UpdatePartnerCampaignRequest) HasWebhookURL() bool {
	if o != nil && !IsNil(o.WebhookURL) {
		return true
	}

	return false
}

// SetWebhookURL gets a reference to the given string and assigns it to the WebhookURL field.
func (o *UpdatePartnerCampaignRequest) SetWebhookURL(v string) {
	o.WebhookURL = &v
}

// GetWebhookFailoverURL returns the WebhookFailoverURL field value if set, zero value otherwise.
func (o *UpdatePartnerCampaignRequest) GetWebhookFailoverURL() string {
	if o == nil || IsNil(o.WebhookFailoverURL) {
		var ret string
		return ret
	}
	return *o.WebhookFailoverURL
}

// GetWebhookFailoverURLOk returns a tuple with the WebhookFailoverURL field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePartnerCampaignRequest) GetWebhookFailoverURLOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookFailoverURL) {
		return nil, false
	}
	return o.WebhookFailoverURL, true
}

// HasWebhookFailoverURL returns a boolean if a field has been set.
func (o *UpdatePartnerCampaignRequest) HasWebhookFailoverURL() bool {
	if o != nil && !IsNil(o.WebhookFailoverURL) {
		return true
	}

	return false
}

// SetWebhookFailoverURL gets a reference to the given string and assigns it to the WebhookFailoverURL field.
func (o *UpdatePartnerCampaignRequest) SetWebhookFailoverURL(v string) {
	o.WebhookFailoverURL = &v
}

func (o UpdatePartnerCampaignRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePartnerCampaignRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.WebhookURL) {
		toSerialize["webhookURL"] = o.WebhookURL
	}
	if !IsNil(o.WebhookFailoverURL) {
		toSerialize["webhookFailoverURL"] = o.WebhookFailoverURL
	}
	return toSerialize, nil
}

type NullableUpdatePartnerCampaignRequest struct {
	value *UpdatePartnerCampaignRequest
	isSet bool
}

func (v NullableUpdatePartnerCampaignRequest) Get() *UpdatePartnerCampaignRequest {
	return v.value
}

func (v *NullableUpdatePartnerCampaignRequest) Set(val *UpdatePartnerCampaignRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdatePartnerCampaignRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdatePartnerCampaignRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdatePartnerCampaignRequest(val *UpdatePartnerCampaignRequest) *NullableUpdatePartnerCampaignRequest {
	return &NullableUpdatePartnerCampaignRequest{value: val, isSet: true}
}

func (v NullableUpdatePartnerCampaignRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdatePartnerCampaignRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


