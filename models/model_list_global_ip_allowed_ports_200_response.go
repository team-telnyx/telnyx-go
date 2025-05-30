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

// checks if the ListGlobalIpAllowedPorts200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListGlobalIpAllowedPorts200Response{}

// ListGlobalIpAllowedPorts200Response struct for ListGlobalIpAllowedPorts200Response
type ListGlobalIpAllowedPorts200Response struct {
	Data []GlobalIPAllowedPort `json:"data,omitempty"`
}

// NewListGlobalIpAllowedPorts200Response instantiates a new ListGlobalIpAllowedPorts200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListGlobalIpAllowedPorts200Response() *ListGlobalIpAllowedPorts200Response {
	this := ListGlobalIpAllowedPorts200Response{}
	return &this
}

// NewListGlobalIpAllowedPorts200ResponseWithDefaults instantiates a new ListGlobalIpAllowedPorts200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListGlobalIpAllowedPorts200ResponseWithDefaults() *ListGlobalIpAllowedPorts200Response {
	this := ListGlobalIpAllowedPorts200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ListGlobalIpAllowedPorts200Response) GetData() []GlobalIPAllowedPort {
	if o == nil || IsNil(o.Data) {
		var ret []GlobalIPAllowedPort
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListGlobalIpAllowedPorts200Response) GetDataOk() ([]GlobalIPAllowedPort, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ListGlobalIpAllowedPorts200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GlobalIPAllowedPort and assigns it to the Data field.
func (o *ListGlobalIpAllowedPorts200Response) SetData(v []GlobalIPAllowedPort) {
	o.Data = v
}

func (o ListGlobalIpAllowedPorts200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListGlobalIpAllowedPorts200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableListGlobalIpAllowedPorts200Response struct {
	value *ListGlobalIpAllowedPorts200Response
	isSet bool
}

func (v NullableListGlobalIpAllowedPorts200Response) Get() *ListGlobalIpAllowedPorts200Response {
	return v.value
}

func (v *NullableListGlobalIpAllowedPorts200Response) Set(val *ListGlobalIpAllowedPorts200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableListGlobalIpAllowedPorts200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableListGlobalIpAllowedPorts200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListGlobalIpAllowedPorts200Response(val *ListGlobalIpAllowedPorts200Response) *NullableListGlobalIpAllowedPorts200Response {
	return &NullableListGlobalIpAllowedPorts200Response{value: val, isSet: true}
}

func (v NullableListGlobalIpAllowedPorts200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListGlobalIpAllowedPorts200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


