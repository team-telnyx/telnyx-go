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

// checks if the TransferTool type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferTool{}

// TransferTool struct for TransferTool
type TransferTool struct {
	Type string `json:"type"`
	Transfer TransferToolParams `json:"transfer"`
}

type _TransferTool TransferTool

// NewTransferTool instantiates a new TransferTool object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferTool(type_ string, transfer TransferToolParams) *TransferTool {
	this := TransferTool{}
	this.Type = type_
	this.Transfer = transfer
	return &this
}

// NewTransferToolWithDefaults instantiates a new TransferTool object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferToolWithDefaults() *TransferTool {
	this := TransferTool{}
	return &this
}

// GetType returns the Type field value
func (o *TransferTool) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TransferTool) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TransferTool) SetType(v string) {
	o.Type = v
}

// GetTransfer returns the Transfer field value
func (o *TransferTool) GetTransfer() TransferToolParams {
	if o == nil {
		var ret TransferToolParams
		return ret
	}

	return o.Transfer
}

// GetTransferOk returns a tuple with the Transfer field value
// and a boolean to check if the value has been set.
func (o *TransferTool) GetTransferOk() (*TransferToolParams, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transfer, true
}

// SetTransfer sets field value
func (o *TransferTool) SetTransfer(v TransferToolParams) {
	o.Transfer = v
}

func (o TransferTool) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferTool) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["transfer"] = o.Transfer
	return toSerialize, nil
}

func (o *TransferTool) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"transfer",
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

	varTransferTool := _TransferTool{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTransferTool)

	if err != nil {
		return err
	}

	*o = TransferTool(varTransferTool)

	return err
}

type NullableTransferTool struct {
	value *TransferTool
	isSet bool
}

func (v NullableTransferTool) Get() *TransferTool {
	return v.value
}

func (v *NullableTransferTool) Set(val *TransferTool) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferTool) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferTool) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferTool(val *TransferTool) *NullableTransferTool {
	return &NullableTransferTool{value: val, isSet: true}
}

func (v NullableTransferTool) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferTool) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


