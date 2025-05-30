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

// checks if the FaxApplicationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FaxApplicationResponse{}

// FaxApplicationResponse struct for FaxApplicationResponse
type FaxApplicationResponse struct {
	Data *FaxApplication `json:"data,omitempty"`
}

// NewFaxApplicationResponse instantiates a new FaxApplicationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFaxApplicationResponse() *FaxApplicationResponse {
	this := FaxApplicationResponse{}
	return &this
}

// NewFaxApplicationResponseWithDefaults instantiates a new FaxApplicationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFaxApplicationResponseWithDefaults() *FaxApplicationResponse {
	this := FaxApplicationResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FaxApplicationResponse) GetData() FaxApplication {
	if o == nil || IsNil(o.Data) {
		var ret FaxApplication
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FaxApplicationResponse) GetDataOk() (*FaxApplication, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FaxApplicationResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given FaxApplication and assigns it to the Data field.
func (o *FaxApplicationResponse) SetData(v FaxApplication) {
	o.Data = &v
}

func (o FaxApplicationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FaxApplicationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableFaxApplicationResponse struct {
	value *FaxApplicationResponse
	isSet bool
}

func (v NullableFaxApplicationResponse) Get() *FaxApplicationResponse {
	return v.value
}

func (v *NullableFaxApplicationResponse) Set(val *FaxApplicationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFaxApplicationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFaxApplicationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFaxApplicationResponse(val *FaxApplicationResponse) *NullableFaxApplicationResponse {
	return &NullableFaxApplicationResponse{value: val, isSet: true}
}

func (v NullableFaxApplicationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFaxApplicationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


