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

// checks if the RCSMessage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RCSMessage{}

// RCSMessage struct for RCSMessage
type RCSMessage struct {
	// RCS Agent ID
	AgentId string `json:"agent_id"`
	// Phone number in +E.164 format
	To string `json:"to"`
	// A valid messaging profile ID
	MessagingProfileId string `json:"messaging_profile_id"`
	// Message type - must be set to \"RCS\"
	Type *string `json:"type,omitempty"`
	// The URL where webhooks related to this message will be sent.
	WebhookUrl *string `json:"webhook_url,omitempty"`
	AgentMessage RCSAgentMessage `json:"agent_message"`
	SmsFallback *SMSFallback `json:"sms_fallback,omitempty"`
}

type _RCSMessage RCSMessage

// NewRCSMessage instantiates a new RCSMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRCSMessage(agentId string, to string, messagingProfileId string, agentMessage RCSAgentMessage) *RCSMessage {
	this := RCSMessage{}
	this.AgentId = agentId
	this.To = to
	this.MessagingProfileId = messagingProfileId
	this.AgentMessage = agentMessage
	return &this
}

// NewRCSMessageWithDefaults instantiates a new RCSMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRCSMessageWithDefaults() *RCSMessage {
	this := RCSMessage{}
	return &this
}

// GetAgentId returns the AgentId field value
func (o *RCSMessage) GetAgentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value
// and a boolean to check if the value has been set.
func (o *RCSMessage) GetAgentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentId, true
}

// SetAgentId sets field value
func (o *RCSMessage) SetAgentId(v string) {
	o.AgentId = v
}

// GetTo returns the To field value
func (o *RCSMessage) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *RCSMessage) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *RCSMessage) SetTo(v string) {
	o.To = v
}

// GetMessagingProfileId returns the MessagingProfileId field value
func (o *RCSMessage) GetMessagingProfileId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MessagingProfileId
}

// GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field value
// and a boolean to check if the value has been set.
func (o *RCSMessage) GetMessagingProfileIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessagingProfileId, true
}

// SetMessagingProfileId sets field value
func (o *RCSMessage) SetMessagingProfileId(v string) {
	o.MessagingProfileId = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RCSMessage) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RCSMessage) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RCSMessage) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *RCSMessage) SetType(v string) {
	o.Type = &v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *RCSMessage) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RCSMessage) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *RCSMessage) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *RCSMessage) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

// GetAgentMessage returns the AgentMessage field value
func (o *RCSMessage) GetAgentMessage() RCSAgentMessage {
	if o == nil {
		var ret RCSAgentMessage
		return ret
	}

	return o.AgentMessage
}

// GetAgentMessageOk returns a tuple with the AgentMessage field value
// and a boolean to check if the value has been set.
func (o *RCSMessage) GetAgentMessageOk() (*RCSAgentMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentMessage, true
}

// SetAgentMessage sets field value
func (o *RCSMessage) SetAgentMessage(v RCSAgentMessage) {
	o.AgentMessage = v
}

// GetSmsFallback returns the SmsFallback field value if set, zero value otherwise.
func (o *RCSMessage) GetSmsFallback() SMSFallback {
	if o == nil || IsNil(o.SmsFallback) {
		var ret SMSFallback
		return ret
	}
	return *o.SmsFallback
}

// GetSmsFallbackOk returns a tuple with the SmsFallback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RCSMessage) GetSmsFallbackOk() (*SMSFallback, bool) {
	if o == nil || IsNil(o.SmsFallback) {
		return nil, false
	}
	return o.SmsFallback, true
}

// HasSmsFallback returns a boolean if a field has been set.
func (o *RCSMessage) HasSmsFallback() bool {
	if o != nil && !IsNil(o.SmsFallback) {
		return true
	}

	return false
}

// SetSmsFallback gets a reference to the given SMSFallback and assigns it to the SmsFallback field.
func (o *RCSMessage) SetSmsFallback(v SMSFallback) {
	o.SmsFallback = &v
}

func (o RCSMessage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RCSMessage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agent_id"] = o.AgentId
	toSerialize["to"] = o.To
	toSerialize["messaging_profile_id"] = o.MessagingProfileId
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhook_url"] = o.WebhookUrl
	}
	toSerialize["agent_message"] = o.AgentMessage
	if !IsNil(o.SmsFallback) {
		toSerialize["sms_fallback"] = o.SmsFallback
	}
	return toSerialize, nil
}

func (o *RCSMessage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agent_id",
		"to",
		"messaging_profile_id",
		"agent_message",
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

	varRCSMessage := _RCSMessage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRCSMessage)

	if err != nil {
		return err
	}

	*o = RCSMessage(varRCSMessage)

	return err
}

type NullableRCSMessage struct {
	value *RCSMessage
	isSet bool
}

func (v NullableRCSMessage) Get() *RCSMessage {
	return v.value
}

func (v *NullableRCSMessage) Set(val *RCSMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableRCSMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableRCSMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRCSMessage(val *RCSMessage) *NullableRCSMessage {
	return &NullableRCSMessage{value: val, isSet: true}
}

func (v NullableRCSMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRCSMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


