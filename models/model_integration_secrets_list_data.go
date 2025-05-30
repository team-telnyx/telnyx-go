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

// checks if the IntegrationSecretsListData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IntegrationSecretsListData{}

// IntegrationSecretsListData struct for IntegrationSecretsListData
type IntegrationSecretsListData struct {
	Data []IntegrationSecret `json:"data"`
	Meta Metadata `json:"meta"`
}

type _IntegrationSecretsListData IntegrationSecretsListData

// NewIntegrationSecretsListData instantiates a new IntegrationSecretsListData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIntegrationSecretsListData(data []IntegrationSecret, meta Metadata) *IntegrationSecretsListData {
	this := IntegrationSecretsListData{}
	this.Data = data
	this.Meta = meta
	return &this
}

// NewIntegrationSecretsListDataWithDefaults instantiates a new IntegrationSecretsListData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIntegrationSecretsListDataWithDefaults() *IntegrationSecretsListData {
	this := IntegrationSecretsListData{}
	return &this
}

// GetData returns the Data field value
func (o *IntegrationSecretsListData) GetData() []IntegrationSecret {
	if o == nil {
		var ret []IntegrationSecret
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *IntegrationSecretsListData) GetDataOk() ([]IntegrationSecret, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *IntegrationSecretsListData) SetData(v []IntegrationSecret) {
	o.Data = v
}

// GetMeta returns the Meta field value
func (o *IntegrationSecretsListData) GetMeta() Metadata {
	if o == nil {
		var ret Metadata
		return ret
	}

	return o.Meta
}

// GetMetaOk returns a tuple with the Meta field value
// and a boolean to check if the value has been set.
func (o *IntegrationSecretsListData) GetMetaOk() (*Metadata, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Meta, true
}

// SetMeta sets field value
func (o *IntegrationSecretsListData) SetMeta(v Metadata) {
	o.Meta = v
}

func (o IntegrationSecretsListData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IntegrationSecretsListData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["meta"] = o.Meta
	return toSerialize, nil
}

func (o *IntegrationSecretsListData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"meta",
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

	varIntegrationSecretsListData := _IntegrationSecretsListData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varIntegrationSecretsListData)

	if err != nil {
		return err
	}

	*o = IntegrationSecretsListData(varIntegrationSecretsListData)

	return err
}

type NullableIntegrationSecretsListData struct {
	value *IntegrationSecretsListData
	isSet bool
}

func (v NullableIntegrationSecretsListData) Get() *IntegrationSecretsListData {
	return v.value
}

func (v *NullableIntegrationSecretsListData) Set(val *IntegrationSecretsListData) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrationSecretsListData) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrationSecretsListData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrationSecretsListData(val *IntegrationSecretsListData) *NullableIntegrationSecretsListData {
	return &NullableIntegrationSecretsListData{value: val, isSet: true}
}

func (v NullableIntegrationSecretsListData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrationSecretsListData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


