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

// checks if the ListInboundChannels200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInboundChannels200ResponseData{}

// ListInboundChannels200ResponseData struct for ListInboundChannels200ResponseData
type ListInboundChannels200ResponseData struct {
	// The current number of concurrent channels set for the account
	Channels *int32 `json:"channels,omitempty"`
	// Identifies the type of the response
	RecordType *string `json:"record_type,omitempty"`
}

// NewListInboundChannels200ResponseData instantiates a new ListInboundChannels200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInboundChannels200ResponseData() *ListInboundChannels200ResponseData {
	this := ListInboundChannels200ResponseData{}
	return &this
}

// NewListInboundChannels200ResponseDataWithDefaults instantiates a new ListInboundChannels200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInboundChannels200ResponseDataWithDefaults() *ListInboundChannels200ResponseData {
	this := ListInboundChannels200ResponseData{}
	return &this
}

// GetChannels returns the Channels field value if set, zero value otherwise.
func (o *ListInboundChannels200ResponseData) GetChannels() int32 {
	if o == nil || IsNil(o.Channels) {
		var ret int32
		return ret
	}
	return *o.Channels
}

// GetChannelsOk returns a tuple with the Channels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInboundChannels200ResponseData) GetChannelsOk() (*int32, bool) {
	if o == nil || IsNil(o.Channels) {
		return nil, false
	}
	return o.Channels, true
}

// HasChannels returns a boolean if a field has been set.
func (o *ListInboundChannels200ResponseData) HasChannels() bool {
	if o != nil && !IsNil(o.Channels) {
		return true
	}

	return false
}

// SetChannels gets a reference to the given int32 and assigns it to the Channels field.
func (o *ListInboundChannels200ResponseData) SetChannels(v int32) {
	o.Channels = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *ListInboundChannels200ResponseData) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInboundChannels200ResponseData) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *ListInboundChannels200ResponseData) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *ListInboundChannels200ResponseData) SetRecordType(v string) {
	o.RecordType = &v
}

func (o ListInboundChannels200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInboundChannels200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Channels) {
		toSerialize["channels"] = o.Channels
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	return toSerialize, nil
}

type NullableListInboundChannels200ResponseData struct {
	value *ListInboundChannels200ResponseData
	isSet bool
}

func (v NullableListInboundChannels200ResponseData) Get() *ListInboundChannels200ResponseData {
	return v.value
}

func (v *NullableListInboundChannels200ResponseData) Set(val *ListInboundChannels200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableListInboundChannels200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableListInboundChannels200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInboundChannels200ResponseData(val *ListInboundChannels200ResponseData) *NullableListInboundChannels200ResponseData {
	return &NullableListInboundChannels200ResponseData{value: val, isSet: true}
}

func (v NullableListInboundChannels200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInboundChannels200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


