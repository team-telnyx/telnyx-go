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
	"time"
)

// checks if the BucketAPIUsageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BucketAPIUsageResponse{}

// BucketAPIUsageResponse struct for BucketAPIUsageResponse
type BucketAPIUsageResponse struct {
	Categories []BucketOps `json:"categories,omitempty"`
	Total *BucketOpsTotal `json:"total,omitempty"`
	// The time the usage was recorded
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewBucketAPIUsageResponse instantiates a new BucketAPIUsageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBucketAPIUsageResponse() *BucketAPIUsageResponse {
	this := BucketAPIUsageResponse{}
	return &this
}

// NewBucketAPIUsageResponseWithDefaults instantiates a new BucketAPIUsageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBucketAPIUsageResponseWithDefaults() *BucketAPIUsageResponse {
	this := BucketAPIUsageResponse{}
	return &this
}

// GetCategories returns the Categories field value if set, zero value otherwise.
func (o *BucketAPIUsageResponse) GetCategories() []BucketOps {
	if o == nil || IsNil(o.Categories) {
		var ret []BucketOps
		return ret
	}
	return o.Categories
}

// GetCategoriesOk returns a tuple with the Categories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketAPIUsageResponse) GetCategoriesOk() ([]BucketOps, bool) {
	if o == nil || IsNil(o.Categories) {
		return nil, false
	}
	return o.Categories, true
}

// HasCategories returns a boolean if a field has been set.
func (o *BucketAPIUsageResponse) HasCategories() bool {
	if o != nil && !IsNil(o.Categories) {
		return true
	}

	return false
}

// SetCategories gets a reference to the given []BucketOps and assigns it to the Categories field.
func (o *BucketAPIUsageResponse) SetCategories(v []BucketOps) {
	o.Categories = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *BucketAPIUsageResponse) GetTotal() BucketOpsTotal {
	if o == nil || IsNil(o.Total) {
		var ret BucketOpsTotal
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketAPIUsageResponse) GetTotalOk() (*BucketOpsTotal, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *BucketAPIUsageResponse) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given BucketOpsTotal and assigns it to the Total field.
func (o *BucketAPIUsageResponse) SetTotal(v BucketOpsTotal) {
	o.Total = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *BucketAPIUsageResponse) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BucketAPIUsageResponse) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *BucketAPIUsageResponse) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *BucketAPIUsageResponse) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o BucketAPIUsageResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BucketAPIUsageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Categories) {
		toSerialize["categories"] = o.Categories
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableBucketAPIUsageResponse struct {
	value *BucketAPIUsageResponse
	isSet bool
}

func (v NullableBucketAPIUsageResponse) Get() *BucketAPIUsageResponse {
	return v.value
}

func (v *NullableBucketAPIUsageResponse) Set(val *BucketAPIUsageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBucketAPIUsageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBucketAPIUsageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBucketAPIUsageResponse(val *BucketAPIUsageResponse) *NullableBucketAPIUsageResponse {
	return &NullableBucketAPIUsageResponse{value: val, isSet: true}
}

func (v NullableBucketAPIUsageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBucketAPIUsageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


