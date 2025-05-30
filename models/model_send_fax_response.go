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

// checks if the SendFaxResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SendFaxResponse{}

// SendFaxResponse struct for SendFaxResponse
type SendFaxResponse struct {
	Data *Fax `json:"data,omitempty"`
}

// NewSendFaxResponse instantiates a new SendFaxResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSendFaxResponse() *SendFaxResponse {
	this := SendFaxResponse{}
	return &this
}

// NewSendFaxResponseWithDefaults instantiates a new SendFaxResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSendFaxResponseWithDefaults() *SendFaxResponse {
	this := SendFaxResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SendFaxResponse) GetData() Fax {
	if o == nil || IsNil(o.Data) {
		var ret Fax
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SendFaxResponse) GetDataOk() (*Fax, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SendFaxResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given Fax and assigns it to the Data field.
func (o *SendFaxResponse) SetData(v Fax) {
	o.Data = &v
}

func (o SendFaxResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SendFaxResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableSendFaxResponse struct {
	value *SendFaxResponse
	isSet bool
}

func (v NullableSendFaxResponse) Get() *SendFaxResponse {
	return v.value
}

func (v *NullableSendFaxResponse) Set(val *SendFaxResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSendFaxResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSendFaxResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSendFaxResponse(val *SendFaxResponse) *NullableSendFaxResponse {
	return &NullableSendFaxResponse{value: val, isSet: true}
}

func (v NullableSendFaxResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSendFaxResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


