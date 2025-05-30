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

// checks if the TransferToolParamsCustomHeadersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TransferToolParamsCustomHeadersInner{}

// TransferToolParamsCustomHeadersInner struct for TransferToolParamsCustomHeadersInner
type TransferToolParamsCustomHeadersInner struct {
	Name *string `json:"name,omitempty"`
	// The value of the header. Note that we support mustache templating for the value. For example you can use `{{#integration_secret}}test-secret{{/integration_secret}}` to pass the value of the integration secret.
	Value *string `json:"value,omitempty"`
}

// NewTransferToolParamsCustomHeadersInner instantiates a new TransferToolParamsCustomHeadersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferToolParamsCustomHeadersInner() *TransferToolParamsCustomHeadersInner {
	this := TransferToolParamsCustomHeadersInner{}
	return &this
}

// NewTransferToolParamsCustomHeadersInnerWithDefaults instantiates a new TransferToolParamsCustomHeadersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferToolParamsCustomHeadersInnerWithDefaults() *TransferToolParamsCustomHeadersInner {
	this := TransferToolParamsCustomHeadersInner{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TransferToolParamsCustomHeadersInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferToolParamsCustomHeadersInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TransferToolParamsCustomHeadersInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TransferToolParamsCustomHeadersInner) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *TransferToolParamsCustomHeadersInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransferToolParamsCustomHeadersInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *TransferToolParamsCustomHeadersInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *TransferToolParamsCustomHeadersInner) SetValue(v string) {
	o.Value = &v
}

func (o TransferToolParamsCustomHeadersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TransferToolParamsCustomHeadersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableTransferToolParamsCustomHeadersInner struct {
	value *TransferToolParamsCustomHeadersInner
	isSet bool
}

func (v NullableTransferToolParamsCustomHeadersInner) Get() *TransferToolParamsCustomHeadersInner {
	return v.value
}

func (v *NullableTransferToolParamsCustomHeadersInner) Set(val *TransferToolParamsCustomHeadersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferToolParamsCustomHeadersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferToolParamsCustomHeadersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferToolParamsCustomHeadersInner(val *TransferToolParamsCustomHeadersInner) *NullableTransferToolParamsCustomHeadersInner {
	return &NullableTransferToolParamsCustomHeadersInner{value: val, isSet: true}
}

func (v NullableTransferToolParamsCustomHeadersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferToolParamsCustomHeadersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


