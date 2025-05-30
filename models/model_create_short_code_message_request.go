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

// checks if the CreateShortCodeMessageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShortCodeMessageRequest{}

// CreateShortCodeMessageRequest struct for CreateShortCodeMessageRequest
type CreateShortCodeMessageRequest struct {
	// Phone number, in +E.164 format, used to send the message.
	From string `json:"from"`
	// Receiving address (+E.164 formatted phone number or short code).
	To string `json:"to"`
	// Message body (i.e., content) as a non-empty string.  **Required for SMS**
	Text *string `json:"text,omitempty"`
	// Subject of multimedia message
	Subject *string `json:"subject,omitempty"`
	// A list of media URLs. The total media size must be less than 1 MB.  **Required for MMS**
	MediaUrls []string `json:"media_urls,omitempty"`
	// The URL where webhooks related to this message will be sent.
	WebhookUrl *string `json:"webhook_url,omitempty"`
	// The failover URL where webhooks related to this message will be sent if sending to the primary URL fails.
	WebhookFailoverUrl *string `json:"webhook_failover_url,omitempty"`
	// If the profile this number is associated with has webhooks, use them for delivery notifications. If webhooks are also specified on the message itself, they will be attempted first, then those on the profile.
	UseProfileWebhooks *bool `json:"use_profile_webhooks,omitempty"`
	// The protocol for sending the message, either SMS or MMS.
	Type *string `json:"type,omitempty"`
	// Automatically detect if an SMS message is unusually long and exceeds a recommended limit of message parts.
	AutoDetect *bool `json:"auto_detect,omitempty"`
}

type _CreateShortCodeMessageRequest CreateShortCodeMessageRequest

// NewCreateShortCodeMessageRequest instantiates a new CreateShortCodeMessageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShortCodeMessageRequest(from string, to string) *CreateShortCodeMessageRequest {
	this := CreateShortCodeMessageRequest{}
	this.From = from
	this.To = to
	var useProfileWebhooks bool = true
	this.UseProfileWebhooks = &useProfileWebhooks
	var autoDetect bool = false
	this.AutoDetect = &autoDetect
	return &this
}

// NewCreateShortCodeMessageRequestWithDefaults instantiates a new CreateShortCodeMessageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShortCodeMessageRequestWithDefaults() *CreateShortCodeMessageRequest {
	this := CreateShortCodeMessageRequest{}
	var useProfileWebhooks bool = true
	this.UseProfileWebhooks = &useProfileWebhooks
	var autoDetect bool = false
	this.AutoDetect = &autoDetect
	return &this
}

// GetFrom returns the From field value
func (o *CreateShortCodeMessageRequest) GetFrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *CreateShortCodeMessageRequest) GetFromOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *CreateShortCodeMessageRequest) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *CreateShortCodeMessageRequest) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *CreateShortCodeMessageRequest) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *CreateShortCodeMessageRequest) SetTo(v string) {
	o.To = v
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *CreateShortCodeMessageRequest) GetText() string {
	if o == nil || IsNil(o.Text) {
		var ret string
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShortCodeMessageRequest) GetTextOk() (*string, bool) {
	if o == nil || IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *CreateShortCodeMessageRequest) HasText() bool {
	if o != nil && !IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given string and assigns it to the Text field.
func (o *CreateShortCodeMessageRequest) SetText(v string) {
	o.Text = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CreateShortCodeMessageRequest) GetSubject() string {
	if o == nil || IsNil(o.Subject) {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShortCodeMessageRequest) GetSubjectOk() (*string, bool) {
	if o == nil || IsNil(o.Subject) {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CreateShortCodeMessageRequest) HasSubject() bool {
	if o != nil && !IsNil(o.Subject) {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CreateShortCodeMessageRequest) SetSubject(v string) {
	o.Subject = &v
}

// GetMediaUrls returns the MediaUrls field value if set, zero value otherwise.
func (o *CreateShortCodeMessageRequest) GetMediaUrls() []string {
	if o == nil || IsNil(o.MediaUrls) {
		var ret []string
		return ret
	}
	return o.MediaUrls
}

// GetMediaUrlsOk returns a tuple with the MediaUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShortCodeMessageRequest) GetMediaUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.MediaUrls) {
		return nil, false
	}
	return o.MediaUrls, true
}

// HasMediaUrls returns a boolean if a field has been set.
func (o *CreateShortCodeMessageRequest) HasMediaUrls() bool {
	if o != nil && !IsNil(o.MediaUrls) {
		return true
	}

	return false
}

// SetMediaUrls gets a reference to the given []string and assigns it to the MediaUrls field.
func (o *CreateShortCodeMessageRequest) SetMediaUrls(v []string) {
	o.MediaUrls = v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *CreateShortCodeMessageRequest) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShortCodeMessageRequest) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *CreateShortCodeMessageRequest) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *CreateShortCodeMessageRequest) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetWebhookFailoverUrl returns the WebhookFailoverUrl field value if set, zero value otherwise.
func (o *CreateShortCodeMessageRequest) GetWebhookFailoverUrl() string {
	if o == nil || IsNil(o.WebhookFailoverUrl) {
		var ret string
		return ret
	}
	return *o.WebhookFailoverUrl
}

// GetWebhookFailoverUrlOk returns a tuple with the WebhookFailoverUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShortCodeMessageRequest) GetWebhookFailoverUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookFailoverUrl) {
		return nil, false
	}
	return o.WebhookFailoverUrl, true
}

// HasWebhookFailoverUrl returns a boolean if a field has been set.
func (o *CreateShortCodeMessageRequest) HasWebhookFailoverUrl() bool {
	if o != nil && !IsNil(o.WebhookFailoverUrl) {
		return true
	}

	return false
}

// SetWebhookFailoverUrl gets a reference to the given string and assigns it to the WebhookFailoverUrl field.
func (o *CreateShortCodeMessageRequest) SetWebhookFailoverUrl(v string) {
	o.WebhookFailoverUrl = &v
}

// GetUseProfileWebhooks returns the UseProfileWebhooks field value if set, zero value otherwise.
func (o *CreateShortCodeMessageRequest) GetUseProfileWebhooks() bool {
	if o == nil || IsNil(o.UseProfileWebhooks) {
		var ret bool
		return ret
	}
	return *o.UseProfileWebhooks
}

// GetUseProfileWebhooksOk returns a tuple with the UseProfileWebhooks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShortCodeMessageRequest) GetUseProfileWebhooksOk() (*bool, bool) {
	if o == nil || IsNil(o.UseProfileWebhooks) {
		return nil, false
	}
	return o.UseProfileWebhooks, true
}

// HasUseProfileWebhooks returns a boolean if a field has been set.
func (o *CreateShortCodeMessageRequest) HasUseProfileWebhooks() bool {
	if o != nil && !IsNil(o.UseProfileWebhooks) {
		return true
	}

	return false
}

// SetUseProfileWebhooks gets a reference to the given bool and assigns it to the UseProfileWebhooks field.
func (o *CreateShortCodeMessageRequest) SetUseProfileWebhooks(v bool) {
	o.UseProfileWebhooks = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CreateShortCodeMessageRequest) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShortCodeMessageRequest) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CreateShortCodeMessageRequest) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CreateShortCodeMessageRequest) SetType(v string) {
	o.Type = &v
}

// GetAutoDetect returns the AutoDetect field value if set, zero value otherwise.
func (o *CreateShortCodeMessageRequest) GetAutoDetect() bool {
	if o == nil || IsNil(o.AutoDetect) {
		var ret bool
		return ret
	}
	return *o.AutoDetect
}

// GetAutoDetectOk returns a tuple with the AutoDetect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateShortCodeMessageRequest) GetAutoDetectOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoDetect) {
		return nil, false
	}
	return o.AutoDetect, true
}

// HasAutoDetect returns a boolean if a field has been set.
func (o *CreateShortCodeMessageRequest) HasAutoDetect() bool {
	if o != nil && !IsNil(o.AutoDetect) {
		return true
	}

	return false
}

// SetAutoDetect gets a reference to the given bool and assigns it to the AutoDetect field.
func (o *CreateShortCodeMessageRequest) SetAutoDetect(v bool) {
	o.AutoDetect = &v
}

func (o CreateShortCodeMessageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShortCodeMessageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["from"] = o.From
	toSerialize["to"] = o.To
	if !IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	if !IsNil(o.Subject) {
		toSerialize["subject"] = o.Subject
	}
	if !IsNil(o.MediaUrls) {
		toSerialize["media_urls"] = o.MediaUrls
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhook_url"] = o.WebhookUrl
	}
	if !IsNil(o.WebhookFailoverUrl) {
		toSerialize["webhook_failover_url"] = o.WebhookFailoverUrl
	}
	if !IsNil(o.UseProfileWebhooks) {
		toSerialize["use_profile_webhooks"] = o.UseProfileWebhooks
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.AutoDetect) {
		toSerialize["auto_detect"] = o.AutoDetect
	}
	return toSerialize, nil
}

func (o *CreateShortCodeMessageRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"from",
		"to",
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

	varCreateShortCodeMessageRequest := _CreateShortCodeMessageRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateShortCodeMessageRequest)

	if err != nil {
		return err
	}

	*o = CreateShortCodeMessageRequest(varCreateShortCodeMessageRequest)

	return err
}

type NullableCreateShortCodeMessageRequest struct {
	value *CreateShortCodeMessageRequest
	isSet bool
}

func (v NullableCreateShortCodeMessageRequest) Get() *CreateShortCodeMessageRequest {
	return v.value
}

func (v *NullableCreateShortCodeMessageRequest) Set(val *CreateShortCodeMessageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShortCodeMessageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShortCodeMessageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShortCodeMessageRequest(val *CreateShortCodeMessageRequest) *NullableCreateShortCodeMessageRequest {
	return &NullableCreateShortCodeMessageRequest{value: val, isSet: true}
}

func (v NullableCreateShortCodeMessageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShortCodeMessageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


