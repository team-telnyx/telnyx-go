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

// checks if the DocServiceRecord type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocServiceRecord{}

// DocServiceRecord struct for DocServiceRecord
type DocServiceRecord struct {
	// Identifies the resource.
	Id *string `json:"id,omitempty"`
	// Identifies the type of the resource.
	RecordType *string `json:"record_type,omitempty"`
	// ISO 8601 formatted date-time indicating when the resource was created.
	CreatedAt *string `json:"created_at,omitempty"`
	// ISO 8601 formatted date-time indicating when the resource was updated.
	UpdatedAt *string `json:"updated_at,omitempty"`
}

// NewDocServiceRecord instantiates a new DocServiceRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocServiceRecord() *DocServiceRecord {
	this := DocServiceRecord{}
	return &this
}

// NewDocServiceRecordWithDefaults instantiates a new DocServiceRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocServiceRecordWithDefaults() *DocServiceRecord {
	this := DocServiceRecord{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DocServiceRecord) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocServiceRecord) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DocServiceRecord) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DocServiceRecord) SetId(v string) {
	o.Id = &v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *DocServiceRecord) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocServiceRecord) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *DocServiceRecord) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *DocServiceRecord) SetRecordType(v string) {
	o.RecordType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DocServiceRecord) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocServiceRecord) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DocServiceRecord) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *DocServiceRecord) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DocServiceRecord) GetUpdatedAt() string {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret string
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DocServiceRecord) GetUpdatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DocServiceRecord) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given string and assigns it to the UpdatedAt field.
func (o *DocServiceRecord) SetUpdatedAt(v string) {
	o.UpdatedAt = &v
}

func (o DocServiceRecord) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocServiceRecord) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	return toSerialize, nil
}

type NullableDocServiceRecord struct {
	value *DocServiceRecord
	isSet bool
}

func (v NullableDocServiceRecord) Get() *DocServiceRecord {
	return v.value
}

func (v *NullableDocServiceRecord) Set(val *DocServiceRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableDocServiceRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableDocServiceRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocServiceRecord(val *DocServiceRecord) *NullableDocServiceRecord {
	return &NullableDocServiceRecord{value: val, isSet: true}
}

func (v NullableDocServiceRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocServiceRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


