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

// checks if the GetCustomerServiceRecord404Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetCustomerServiceRecord404Response{}

// GetCustomerServiceRecord404Response struct for GetCustomerServiceRecord404Response
type GetCustomerServiceRecord404Response struct {
	Errors []ResourceNotFoundError `json:"errors,omitempty"`
}

// NewGetCustomerServiceRecord404Response instantiates a new GetCustomerServiceRecord404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCustomerServiceRecord404Response() *GetCustomerServiceRecord404Response {
	this := GetCustomerServiceRecord404Response{}
	return &this
}

// NewGetCustomerServiceRecord404ResponseWithDefaults instantiates a new GetCustomerServiceRecord404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCustomerServiceRecord404ResponseWithDefaults() *GetCustomerServiceRecord404Response {
	this := GetCustomerServiceRecord404Response{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *GetCustomerServiceRecord404Response) GetErrors() []ResourceNotFoundError {
	if o == nil || IsNil(o.Errors) {
		var ret []ResourceNotFoundError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCustomerServiceRecord404Response) GetErrorsOk() ([]ResourceNotFoundError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *GetCustomerServiceRecord404Response) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ResourceNotFoundError and assigns it to the Errors field.
func (o *GetCustomerServiceRecord404Response) SetErrors(v []ResourceNotFoundError) {
	o.Errors = v
}

func (o GetCustomerServiceRecord404Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetCustomerServiceRecord404Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableGetCustomerServiceRecord404Response struct {
	value *GetCustomerServiceRecord404Response
	isSet bool
}

func (v NullableGetCustomerServiceRecord404Response) Get() *GetCustomerServiceRecord404Response {
	return v.value
}

func (v *NullableGetCustomerServiceRecord404Response) Set(val *GetCustomerServiceRecord404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCustomerServiceRecord404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCustomerServiceRecord404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCustomerServiceRecord404Response(val *GetCustomerServiceRecord404Response) *NullableGetCustomerServiceRecord404Response {
	return &NullableGetCustomerServiceRecord404Response{value: val, isSet: true}
}

func (v NullableGetCustomerServiceRecord404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCustomerServiceRecord404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


