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
	"fmt"
	"gopkg.in/validator.v2"
)

// ChatCompletionRequestToolsInner - struct for ChatCompletionRequestToolsInner
type ChatCompletionRequestToolsInner struct {
	ChatCompletionToolParam *ChatCompletionToolParam
	Retrieval *Retrieval
}

// ChatCompletionToolParamAsChatCompletionRequestToolsInner is a convenience function that returns ChatCompletionToolParam wrapped in ChatCompletionRequestToolsInner
func ChatCompletionToolParamAsChatCompletionRequestToolsInner(v *ChatCompletionToolParam) ChatCompletionRequestToolsInner {
	return ChatCompletionRequestToolsInner{
		ChatCompletionToolParam: v,
	}
}

// RetrievalAsChatCompletionRequestToolsInner is a convenience function that returns Retrieval wrapped in ChatCompletionRequestToolsInner
func RetrievalAsChatCompletionRequestToolsInner(v *Retrieval) ChatCompletionRequestToolsInner {
	return ChatCompletionRequestToolsInner{
		Retrieval: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ChatCompletionRequestToolsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ChatCompletionToolParam
	err = newStrictDecoder(data).Decode(&dst.ChatCompletionToolParam)
	if err == nil {
		jsonChatCompletionToolParam, _ := json.Marshal(dst.ChatCompletionToolParam)
		if string(jsonChatCompletionToolParam) == "{}" { // empty struct
			dst.ChatCompletionToolParam = nil
		} else {
			if err = validator.Validate(dst.ChatCompletionToolParam); err != nil {
				dst.ChatCompletionToolParam = nil
			} else {
				match++
			}
		}
	} else {
		dst.ChatCompletionToolParam = nil
	}

	// try to unmarshal data into Retrieval
	err = newStrictDecoder(data).Decode(&dst.Retrieval)
	if err == nil {
		jsonRetrieval, _ := json.Marshal(dst.Retrieval)
		if string(jsonRetrieval) == "{}" { // empty struct
			dst.Retrieval = nil
		} else {
			if err = validator.Validate(dst.Retrieval); err != nil {
				dst.Retrieval = nil
			} else {
				match++
			}
		}
	} else {
		dst.Retrieval = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ChatCompletionToolParam = nil
		dst.Retrieval = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ChatCompletionRequestToolsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ChatCompletionRequestToolsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ChatCompletionRequestToolsInner) MarshalJSON() ([]byte, error) {
	if src.ChatCompletionToolParam != nil {
		return json.Marshal(&src.ChatCompletionToolParam)
	}

	if src.Retrieval != nil {
		return json.Marshal(&src.Retrieval)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ChatCompletionRequestToolsInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.ChatCompletionToolParam != nil {
		return obj.ChatCompletionToolParam
	}

	if obj.Retrieval != nil {
		return obj.Retrieval
	}

	// all schemas are nil
	return nil
}

// Get the actual instance value
func (obj ChatCompletionRequestToolsInner) GetActualInstanceValue() (interface{}) {
	if obj.ChatCompletionToolParam != nil {
		return *obj.ChatCompletionToolParam
	}

	if obj.Retrieval != nil {
		return *obj.Retrieval
	}

	// all schemas are nil
	return nil
}

type NullableChatCompletionRequestToolsInner struct {
	value *ChatCompletionRequestToolsInner
	isSet bool
}

func (v NullableChatCompletionRequestToolsInner) Get() *ChatCompletionRequestToolsInner {
	return v.value
}

func (v *NullableChatCompletionRequestToolsInner) Set(val *ChatCompletionRequestToolsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChatCompletionRequestToolsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChatCompletionRequestToolsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChatCompletionRequestToolsInner(val *ChatCompletionRequestToolsInner) *NullableChatCompletionRequestToolsInner {
	return &NullableChatCompletionRequestToolsInner{value: val, isSet: true}
}

func (v NullableChatCompletionRequestToolsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChatCompletionRequestToolsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


