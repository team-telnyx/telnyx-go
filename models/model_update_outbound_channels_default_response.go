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

// checks if the UpdateOutboundChannelsDefaultResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateOutboundChannelsDefaultResponse{}

// UpdateOutboundChannelsDefaultResponse struct for UpdateOutboundChannelsDefaultResponse
type UpdateOutboundChannelsDefaultResponse struct {
	Errors []UpdateOutboundChannelsDefaultResponseErrorsInner `json:"errors,omitempty"`
}

// NewUpdateOutboundChannelsDefaultResponse instantiates a new UpdateOutboundChannelsDefaultResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateOutboundChannelsDefaultResponse() *UpdateOutboundChannelsDefaultResponse {
	this := UpdateOutboundChannelsDefaultResponse{}
	return &this
}

// NewUpdateOutboundChannelsDefaultResponseWithDefaults instantiates a new UpdateOutboundChannelsDefaultResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateOutboundChannelsDefaultResponseWithDefaults() *UpdateOutboundChannelsDefaultResponse {
	this := UpdateOutboundChannelsDefaultResponse{}
	return &this
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *UpdateOutboundChannelsDefaultResponse) GetErrors() []UpdateOutboundChannelsDefaultResponseErrorsInner {
	if o == nil || IsNil(o.Errors) {
		var ret []UpdateOutboundChannelsDefaultResponseErrorsInner
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateOutboundChannelsDefaultResponse) GetErrorsOk() ([]UpdateOutboundChannelsDefaultResponseErrorsInner, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *UpdateOutboundChannelsDefaultResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []UpdateOutboundChannelsDefaultResponseErrorsInner and assigns it to the Errors field.
func (o *UpdateOutboundChannelsDefaultResponse) SetErrors(v []UpdateOutboundChannelsDefaultResponseErrorsInner) {
	o.Errors = v
}

func (o UpdateOutboundChannelsDefaultResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateOutboundChannelsDefaultResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	return toSerialize, nil
}

type NullableUpdateOutboundChannelsDefaultResponse struct {
	value *UpdateOutboundChannelsDefaultResponse
	isSet bool
}

func (v NullableUpdateOutboundChannelsDefaultResponse) Get() *UpdateOutboundChannelsDefaultResponse {
	return v.value
}

func (v *NullableUpdateOutboundChannelsDefaultResponse) Set(val *UpdateOutboundChannelsDefaultResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateOutboundChannelsDefaultResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateOutboundChannelsDefaultResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateOutboundChannelsDefaultResponse(val *UpdateOutboundChannelsDefaultResponse) *NullableUpdateOutboundChannelsDefaultResponse {
	return &NullableUpdateOutboundChannelsDefaultResponse{value: val, isSet: true}
}

func (v NullableUpdateOutboundChannelsDefaultResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateOutboundChannelsDefaultResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


