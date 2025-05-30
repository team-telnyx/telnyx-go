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

// checks if the ChatCompletionToolParam type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChatCompletionToolParam{}

// ChatCompletionToolParam struct for ChatCompletionToolParam
type ChatCompletionToolParam struct {
	Type string `json:"type"`
	Function FunctionDefinition `json:"function"`
}

type _ChatCompletionToolParam ChatCompletionToolParam

// NewChatCompletionToolParam instantiates a new ChatCompletionToolParam object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChatCompletionToolParam(type_ string, function FunctionDefinition) *ChatCompletionToolParam {
	this := ChatCompletionToolParam{}
	this.Type = type_
	this.Function = function
	return &this
}

// NewChatCompletionToolParamWithDefaults instantiates a new ChatCompletionToolParam object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChatCompletionToolParamWithDefaults() *ChatCompletionToolParam {
	this := ChatCompletionToolParam{}
	return &this
}

// GetType returns the Type field value
func (o *ChatCompletionToolParam) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionToolParam) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ChatCompletionToolParam) SetType(v string) {
	o.Type = v
}

// GetFunction returns the Function field value
func (o *ChatCompletionToolParam) GetFunction() FunctionDefinition {
	if o == nil {
		var ret FunctionDefinition
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *ChatCompletionToolParam) GetFunctionOk() (*FunctionDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *ChatCompletionToolParam) SetFunction(v FunctionDefinition) {
	o.Function = v
}

func (o ChatCompletionToolParam) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChatCompletionToolParam) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["function"] = o.Function
	return toSerialize, nil
}

func (o *ChatCompletionToolParam) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"function",
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

	varChatCompletionToolParam := _ChatCompletionToolParam{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varChatCompletionToolParam)

	if err != nil {
		return err
	}

	*o = ChatCompletionToolParam(varChatCompletionToolParam)

	return err
}

type NullableChatCompletionToolParam struct {
	value *ChatCompletionToolParam
	isSet bool
}

func (v NullableChatCompletionToolParam) Get() *ChatCompletionToolParam {
	return v.value
}

func (v *NullableChatCompletionToolParam) Set(val *ChatCompletionToolParam) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionToolParam) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionToolParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionToolParam(val *ChatCompletionToolParam) *NullableChatCompletionToolParam {
	return &NullableChatCompletionToolParam{value: val, isSet: true}
}

func (v NullableChatCompletionToolParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionToolParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


