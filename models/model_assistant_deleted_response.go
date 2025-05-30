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

// checks if the AssistantDeletedResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssistantDeletedResponse{}

// AssistantDeletedResponse Aligns with the OpenAI API: https://platform.openai.com/docs/api-reference/assistants/deleteAssistant
type AssistantDeletedResponse struct {
	Id string `json:"id"`
	Object string `json:"object"`
	Deleted bool `json:"deleted"`
}

type _AssistantDeletedResponse AssistantDeletedResponse

// NewAssistantDeletedResponse instantiates a new AssistantDeletedResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssistantDeletedResponse(id string, object string, deleted bool) *AssistantDeletedResponse {
	this := AssistantDeletedResponse{}
	this.Id = id
	this.Object = object
	this.Deleted = deleted
	return &this
}

// NewAssistantDeletedResponseWithDefaults instantiates a new AssistantDeletedResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssistantDeletedResponseWithDefaults() *AssistantDeletedResponse {
	this := AssistantDeletedResponse{}
	return &this
}

// GetId returns the Id field value
func (o *AssistantDeletedResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AssistantDeletedResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AssistantDeletedResponse) SetId(v string) {
	o.Id = v
}

// GetObject returns the Object field value
func (o *AssistantDeletedResponse) GetObject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Object
}

// GetObjectOk returns a tuple with the Object field value
// and a boolean to check if the value has been set.
func (o *AssistantDeletedResponse) GetObjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Object, true
}

// SetObject sets field value
func (o *AssistantDeletedResponse) SetObject(v string) {
	o.Object = v
}

// GetDeleted returns the Deleted field value
func (o *AssistantDeletedResponse) GetDeleted() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value
// and a boolean to check if the value has been set.
func (o *AssistantDeletedResponse) GetDeletedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Deleted, true
}

// SetDeleted sets field value
func (o *AssistantDeletedResponse) SetDeleted(v bool) {
	o.Deleted = v
}

func (o AssistantDeletedResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssistantDeletedResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["object"] = o.Object
	toSerialize["deleted"] = o.Deleted
	return toSerialize, nil
}

func (o *AssistantDeletedResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"object",
		"deleted",
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

	varAssistantDeletedResponse := _AssistantDeletedResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssistantDeletedResponse)

	if err != nil {
		return err
	}

	*o = AssistantDeletedResponse(varAssistantDeletedResponse)

	return err
}

type NullableAssistantDeletedResponse struct {
	value *AssistantDeletedResponse
	isSet bool
}

func (v NullableAssistantDeletedResponse) Get() *AssistantDeletedResponse {
	return v.value
}

func (v *NullableAssistantDeletedResponse) Set(val *AssistantDeletedResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAssistantDeletedResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAssistantDeletedResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssistantDeletedResponse(val *AssistantDeletedResponse) *NullableAssistantDeletedResponse {
	return &NullableAssistantDeletedResponse{value: val, isSet: true}
}

func (v NullableAssistantDeletedResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssistantDeletedResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


