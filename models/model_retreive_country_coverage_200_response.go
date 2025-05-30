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

// checks if the RetreiveCountryCoverage200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetreiveCountryCoverage200Response{}

// RetreiveCountryCoverage200Response struct for RetreiveCountryCoverage200Response
type RetreiveCountryCoverage200Response struct {
	Data []CountryCoverage `json:"data,omitempty"`
}

// NewRetreiveCountryCoverage200Response instantiates a new RetreiveCountryCoverage200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetreiveCountryCoverage200Response() *RetreiveCountryCoverage200Response {
	this := RetreiveCountryCoverage200Response{}
	return &this
}

// NewRetreiveCountryCoverage200ResponseWithDefaults instantiates a new RetreiveCountryCoverage200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetreiveCountryCoverage200ResponseWithDefaults() *RetreiveCountryCoverage200Response {
	this := RetreiveCountryCoverage200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *RetreiveCountryCoverage200Response) GetData() []CountryCoverage {
	if o == nil || IsNil(o.Data) {
		var ret []CountryCoverage
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetreiveCountryCoverage200Response) GetDataOk() ([]CountryCoverage, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *RetreiveCountryCoverage200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []CountryCoverage and assigns it to the Data field.
func (o *RetreiveCountryCoverage200Response) SetData(v []CountryCoverage) {
	o.Data = v
}

func (o RetreiveCountryCoverage200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetreiveCountryCoverage200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableRetreiveCountryCoverage200Response struct {
	value *RetreiveCountryCoverage200Response
	isSet bool
}

func (v NullableRetreiveCountryCoverage200Response) Get() *RetreiveCountryCoverage200Response {
	return v.value
}

func (v *NullableRetreiveCountryCoverage200Response) Set(val *RetreiveCountryCoverage200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableRetreiveCountryCoverage200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableRetreiveCountryCoverage200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetreiveCountryCoverage200Response(val *RetreiveCountryCoverage200Response) *NullableRetreiveCountryCoverage200Response {
	return &NullableRetreiveCountryCoverage200Response{value: val, isSet: true}
}

func (v NullableRetreiveCountryCoverage200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetreiveCountryCoverage200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


