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

// checks if the CreateGlobalIpHealthCheck202Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateGlobalIpHealthCheck202Response{}

// CreateGlobalIpHealthCheck202Response struct for CreateGlobalIpHealthCheck202Response
type CreateGlobalIpHealthCheck202Response struct {
	Data *GlobalIPHealthCheck `json:"data,omitempty"`
}

// NewCreateGlobalIpHealthCheck202Response instantiates a new CreateGlobalIpHealthCheck202Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateGlobalIpHealthCheck202Response() *CreateGlobalIpHealthCheck202Response {
	this := CreateGlobalIpHealthCheck202Response{}
	return &this
}

// NewCreateGlobalIpHealthCheck202ResponseWithDefaults instantiates a new CreateGlobalIpHealthCheck202Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateGlobalIpHealthCheck202ResponseWithDefaults() *CreateGlobalIpHealthCheck202Response {
	this := CreateGlobalIpHealthCheck202Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CreateGlobalIpHealthCheck202Response) GetData() GlobalIPHealthCheck {
	if o == nil || IsNil(o.Data) {
		var ret GlobalIPHealthCheck
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateGlobalIpHealthCheck202Response) GetDataOk() (*GlobalIPHealthCheck, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CreateGlobalIpHealthCheck202Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GlobalIPHealthCheck and assigns it to the Data field.
func (o *CreateGlobalIpHealthCheck202Response) SetData(v GlobalIPHealthCheck) {
	o.Data = &v
}

func (o CreateGlobalIpHealthCheck202Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateGlobalIpHealthCheck202Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCreateGlobalIpHealthCheck202Response struct {
	value *CreateGlobalIpHealthCheck202Response
	isSet bool
}

func (v NullableCreateGlobalIpHealthCheck202Response) Get() *CreateGlobalIpHealthCheck202Response {
	return v.value
}

func (v *NullableCreateGlobalIpHealthCheck202Response) Set(val *CreateGlobalIpHealthCheck202Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateGlobalIpHealthCheck202Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateGlobalIpHealthCheck202Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateGlobalIpHealthCheck202Response(val *CreateGlobalIpHealthCheck202Response) *NullableCreateGlobalIpHealthCheck202Response {
	return &NullableCreateGlobalIpHealthCheck202Response{value: val, isSet: true}
}

func (v NullableCreateGlobalIpHealthCheck202Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateGlobalIpHealthCheck202Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


