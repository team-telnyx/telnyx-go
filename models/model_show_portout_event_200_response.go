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

// checks if the ShowPortoutEvent200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShowPortoutEvent200Response{}

// ShowPortoutEvent200Response struct for ShowPortoutEvent200Response
type ShowPortoutEvent200Response struct {
	Data *PortoutEvent `json:"data,omitempty"`
}

// NewShowPortoutEvent200Response instantiates a new ShowPortoutEvent200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShowPortoutEvent200Response() *ShowPortoutEvent200Response {
	this := ShowPortoutEvent200Response{}
	return &this
}

// NewShowPortoutEvent200ResponseWithDefaults instantiates a new ShowPortoutEvent200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShowPortoutEvent200ResponseWithDefaults() *ShowPortoutEvent200Response {
	this := ShowPortoutEvent200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ShowPortoutEvent200Response) GetData() PortoutEvent {
	if o == nil || IsNil(o.Data) {
		var ret PortoutEvent
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShowPortoutEvent200Response) GetDataOk() (*PortoutEvent, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ShowPortoutEvent200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given PortoutEvent and assigns it to the Data field.
func (o *ShowPortoutEvent200Response) SetData(v PortoutEvent) {
	o.Data = &v
}

func (o ShowPortoutEvent200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShowPortoutEvent200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableShowPortoutEvent200Response struct {
	value *ShowPortoutEvent200Response
	isSet bool
}

func (v NullableShowPortoutEvent200Response) Get() *ShowPortoutEvent200Response {
	return v.value
}

func (v *NullableShowPortoutEvent200Response) Set(val *ShowPortoutEvent200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableShowPortoutEvent200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableShowPortoutEvent200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShowPortoutEvent200Response(val *ShowPortoutEvent200Response) *NullableShowPortoutEvent200Response {
	return &NullableShowPortoutEvent200Response{value: val, isSet: true}
}

func (v NullableShowPortoutEvent200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShowPortoutEvent200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


