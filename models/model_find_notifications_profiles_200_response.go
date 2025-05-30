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

// checks if the FindNotificationsProfiles200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FindNotificationsProfiles200Response{}

// FindNotificationsProfiles200Response struct for FindNotificationsProfiles200Response
type FindNotificationsProfiles200Response struct {
	Data []NotificationProfile `json:"data,omitempty"`
	Meta *PaginationMeta `json:"meta,omitempty"`
}

// NewFindNotificationsProfiles200Response instantiates a new FindNotificationsProfiles200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFindNotificationsProfiles200Response() *FindNotificationsProfiles200Response {
	this := FindNotificationsProfiles200Response{}
	return &this
}

// NewFindNotificationsProfiles200ResponseWithDefaults instantiates a new FindNotificationsProfiles200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFindNotificationsProfiles200ResponseWithDefaults() *FindNotificationsProfiles200Response {
	this := FindNotificationsProfiles200Response{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *FindNotificationsProfiles200Response) GetData() []NotificationProfile {
	if o == nil || IsNil(o.Data) {
		var ret []NotificationProfile
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindNotificationsProfiles200Response) GetDataOk() ([]NotificationProfile, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *FindNotificationsProfiles200Response) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []NotificationProfile and assigns it to the Data field.
func (o *FindNotificationsProfiles200Response) SetData(v []NotificationProfile) {
	o.Data = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *FindNotificationsProfiles200Response) GetMeta() PaginationMeta {
	if o == nil || IsNil(o.Meta) {
		var ret PaginationMeta
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FindNotificationsProfiles200Response) GetMetaOk() (*PaginationMeta, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *FindNotificationsProfiles200Response) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PaginationMeta and assigns it to the Meta field.
func (o *FindNotificationsProfiles200Response) SetMeta(v PaginationMeta) {
	o.Meta = &v
}

func (o FindNotificationsProfiles200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FindNotificationsProfiles200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableFindNotificationsProfiles200Response struct {
	value *FindNotificationsProfiles200Response
	isSet bool
}

func (v NullableFindNotificationsProfiles200Response) Get() *FindNotificationsProfiles200Response {
	return v.value
}

func (v *NullableFindNotificationsProfiles200Response) Set(val *FindNotificationsProfiles200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableFindNotificationsProfiles200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableFindNotificationsProfiles200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFindNotificationsProfiles200Response(val *FindNotificationsProfiles200Response) *NullableFindNotificationsProfiles200Response {
	return &NullableFindNotificationsProfiles200Response{value: val, isSet: true}
}

func (v NullableFindNotificationsProfiles200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFindNotificationsProfiles200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


