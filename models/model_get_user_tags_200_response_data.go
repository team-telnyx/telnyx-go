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

// checks if the GetUserTags200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetUserTags200ResponseData{}

// GetUserTags200ResponseData struct for GetUserTags200ResponseData
type GetUserTags200ResponseData struct {
	// A list of your tags on the given resource type. NOTE: The casing of the tags returned will not necessarily match the casing of the tags when they were added to a resource. This is because the tags will have the casing of the first time they were used within the Telnyx system regardless of source.
	OutboundProfileTags []string `json:"outbound_profile_tags,omitempty"`
	// A list of your tags on the given resource type. NOTE: The casing of the tags returned will not necessarily match the casing of the tags when they were added to a resource. This is because the tags will have the casing of the first time they were used within the Telnyx system regardless of source.
	NumberTags []string `json:"number_tags,omitempty"`
}

// NewGetUserTags200ResponseData instantiates a new GetUserTags200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetUserTags200ResponseData() *GetUserTags200ResponseData {
	this := GetUserTags200ResponseData{}
	return &this
}

// NewGetUserTags200ResponseDataWithDefaults instantiates a new GetUserTags200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetUserTags200ResponseDataWithDefaults() *GetUserTags200ResponseData {
	this := GetUserTags200ResponseData{}
	return &this
}

// GetOutboundProfileTags returns the OutboundProfileTags field value if set, zero value otherwise.
func (o *GetUserTags200ResponseData) GetOutboundProfileTags() []string {
	if o == nil || IsNil(o.OutboundProfileTags) {
		var ret []string
		return ret
	}
	return o.OutboundProfileTags
}

// GetOutboundProfileTagsOk returns a tuple with the OutboundProfileTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTags200ResponseData) GetOutboundProfileTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.OutboundProfileTags) {
		return nil, false
	}
	return o.OutboundProfileTags, true
}

// HasOutboundProfileTags returns a boolean if a field has been set.
func (o *GetUserTags200ResponseData) HasOutboundProfileTags() bool {
	if o != nil && !IsNil(o.OutboundProfileTags) {
		return true
	}

	return false
}

// SetOutboundProfileTags gets a reference to the given []string and assigns it to the OutboundProfileTags field.
func (o *GetUserTags200ResponseData) SetOutboundProfileTags(v []string) {
	o.OutboundProfileTags = v
}

// GetNumberTags returns the NumberTags field value if set, zero value otherwise.
func (o *GetUserTags200ResponseData) GetNumberTags() []string {
	if o == nil || IsNil(o.NumberTags) {
		var ret []string
		return ret
	}
	return o.NumberTags
}

// GetNumberTagsOk returns a tuple with the NumberTags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetUserTags200ResponseData) GetNumberTagsOk() ([]string, bool) {
	if o == nil || IsNil(o.NumberTags) {
		return nil, false
	}
	return o.NumberTags, true
}

// HasNumberTags returns a boolean if a field has been set.
func (o *GetUserTags200ResponseData) HasNumberTags() bool {
	if o != nil && !IsNil(o.NumberTags) {
		return true
	}

	return false
}

// SetNumberTags gets a reference to the given []string and assigns it to the NumberTags field.
func (o *GetUserTags200ResponseData) SetNumberTags(v []string) {
	o.NumberTags = v
}

func (o GetUserTags200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetUserTags200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OutboundProfileTags) {
		toSerialize["outbound_profile_tags"] = o.OutboundProfileTags
	}
	if !IsNil(o.NumberTags) {
		toSerialize["number_tags"] = o.NumberTags
	}
	return toSerialize, nil
}

type NullableGetUserTags200ResponseData struct {
	value *GetUserTags200ResponseData
	isSet bool
}

func (v NullableGetUserTags200ResponseData) Get() *GetUserTags200ResponseData {
	return v.value
}

func (v *NullableGetUserTags200ResponseData) Set(val *GetUserTags200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableGetUserTags200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableGetUserTags200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetUserTags200ResponseData(val *GetUserTags200ResponseData) *NullableGetUserTags200ResponseData {
	return &NullableGetUserTags200ResponseData{value: val, isSet: true}
}

func (v NullableGetUserTags200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetUserTags200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


