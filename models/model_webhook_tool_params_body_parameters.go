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

// checks if the WebhookToolParamsBodyParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookToolParamsBodyParameters{}

// WebhookToolParamsBodyParameters The body parameters the webhook tool accepts, described as a JSON Schema object. These parameters will be passed to the webhook as the body of the request. See the [JSON Schema reference](https://json-schema.org/understanding-json-schema) for documentation about the format
type WebhookToolParamsBodyParameters struct {
	// The properties of the body parameters.
	Properties map[string]interface{} `json:"properties,omitempty"`
	// The required properties of the body parameters.
	Required []string `json:"required,omitempty"`
	Type *string `json:"type,omitempty"`
}

// NewWebhookToolParamsBodyParameters instantiates a new WebhookToolParamsBodyParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookToolParamsBodyParameters() *WebhookToolParamsBodyParameters {
	this := WebhookToolParamsBodyParameters{}
	return &this
}

// NewWebhookToolParamsBodyParametersWithDefaults instantiates a new WebhookToolParamsBodyParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookToolParamsBodyParametersWithDefaults() *WebhookToolParamsBodyParameters {
	this := WebhookToolParamsBodyParameters{}
	return &this
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *WebhookToolParamsBodyParameters) GetProperties() map[string]interface{} {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]interface{}
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookToolParamsBodyParameters) GetPropertiesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Properties) {
		return map[string]interface{}{}, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *WebhookToolParamsBodyParameters) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *WebhookToolParamsBodyParameters) SetProperties(v map[string]interface{}) {
	o.Properties = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *WebhookToolParamsBodyParameters) GetRequired() []string {
	if o == nil || IsNil(o.Required) {
		var ret []string
		return ret
	}
	return o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookToolParamsBodyParameters) GetRequiredOk() ([]string, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *WebhookToolParamsBodyParameters) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given []string and assigns it to the Required field.
func (o *WebhookToolParamsBodyParameters) SetRequired(v []string) {
	o.Required = v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WebhookToolParamsBodyParameters) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookToolParamsBodyParameters) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WebhookToolParamsBodyParameters) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *WebhookToolParamsBodyParameters) SetType(v string) {
	o.Type = &v
}

func (o WebhookToolParamsBodyParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookToolParamsBodyParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableWebhookToolParamsBodyParameters struct {
	value *WebhookToolParamsBodyParameters
	isSet bool
}

func (v NullableWebhookToolParamsBodyParameters) Get() *WebhookToolParamsBodyParameters {
	return v.value
}

func (v *NullableWebhookToolParamsBodyParameters) Set(val *WebhookToolParamsBodyParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookToolParamsBodyParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookToolParamsBodyParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookToolParamsBodyParameters(val *WebhookToolParamsBodyParameters) *NullableWebhookToolParamsBodyParameters {
	return &NullableWebhookToolParamsBodyParameters{value: val, isSet: true}
}

func (v NullableWebhookToolParamsBodyParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookToolParamsBodyParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


