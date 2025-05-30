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

// checks if the WebhookTool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookTool{}

// WebhookTool struct for WebhookTool
type WebhookTool struct {
	Type string `json:"type"`
	Webhook WebhookToolParams `json:"webhook"`
}

type _WebhookTool WebhookTool

// NewWebhookTool instantiates a new WebhookTool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookTool(type_ string, webhook WebhookToolParams) *WebhookTool {
	this := WebhookTool{}
	this.Type = type_
	this.Webhook = webhook
	return &this
}

// NewWebhookToolWithDefaults instantiates a new WebhookTool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookToolWithDefaults() *WebhookTool {
	this := WebhookTool{}
	return &this
}

// GetType returns the Type field value
func (o *WebhookTool) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *WebhookTool) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *WebhookTool) SetType(v string) {
	o.Type = v
}

// GetWebhook returns the Webhook field value
func (o *WebhookTool) GetWebhook() WebhookToolParams {
	if o == nil {
		var ret WebhookToolParams
		return ret
	}

	return o.Webhook
}

// GetWebhookOk returns a tuple with the Webhook field value
// and a boolean to check if the value has been set.
func (o *WebhookTool) GetWebhookOk() (*WebhookToolParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Webhook, true
}

// SetWebhook sets field value
func (o *WebhookTool) SetWebhook(v WebhookToolParams) {
	o.Webhook = v
}

func (o WebhookTool) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookTool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["webhook"] = o.Webhook
	return toSerialize, nil
}

func (o *WebhookTool) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"webhook",
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

	varWebhookTool := _WebhookTool{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWebhookTool)

	if err != nil {
		return err
	}

	*o = WebhookTool(varWebhookTool)

	return err
}

type NullableWebhookTool struct {
	value *WebhookTool
	isSet bool
}

func (v NullableWebhookTool) Get() *WebhookTool {
	return v.value
}

func (v *NullableWebhookTool) Set(val *WebhookTool) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookTool) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookTool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookTool(val *WebhookTool) *NullableWebhookTool {
	return &NullableWebhookTool{value: val, isSet: true}
}

func (v NullableWebhookTool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookTool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


