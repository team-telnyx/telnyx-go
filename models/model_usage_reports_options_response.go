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

// checks if the UsageReportsOptionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageReportsOptionsResponse{}

// UsageReportsOptionsResponse An object following one of the schemas published in https://developers.telnyx.com/docs/api/v2/detail-records
type UsageReportsOptionsResponse struct {
	// Collection of product description
	Data []UsageReportsOptionsRecord `json:"data,omitempty"`
}

// NewUsageReportsOptionsResponse instantiates a new UsageReportsOptionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageReportsOptionsResponse() *UsageReportsOptionsResponse {
	this := UsageReportsOptionsResponse{}
	return &this
}

// NewUsageReportsOptionsResponseWithDefaults instantiates a new UsageReportsOptionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageReportsOptionsResponseWithDefaults() *UsageReportsOptionsResponse {
	this := UsageReportsOptionsResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *UsageReportsOptionsResponse) GetData() []UsageReportsOptionsRecord {
	if o == nil || IsNil(o.Data) {
		var ret []UsageReportsOptionsRecord
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageReportsOptionsResponse) GetDataOk() ([]UsageReportsOptionsRecord, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *UsageReportsOptionsResponse) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []UsageReportsOptionsRecord and assigns it to the Data field.
func (o *UsageReportsOptionsResponse) SetData(v []UsageReportsOptionsRecord) {
	o.Data = v
}

func (o UsageReportsOptionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageReportsOptionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableUsageReportsOptionsResponse struct {
	value *UsageReportsOptionsResponse
	isSet bool
}

func (v NullableUsageReportsOptionsResponse) Get() *UsageReportsOptionsResponse {
	return v.value
}

func (v *NullableUsageReportsOptionsResponse) Set(val *UsageReportsOptionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageReportsOptionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageReportsOptionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageReportsOptionsResponse(val *UsageReportsOptionsResponse) *NullableUsageReportsOptionsResponse {
	return &NullableUsageReportsOptionsResponse{value: val, isSet: true}
}

func (v NullableUsageReportsOptionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageReportsOptionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


