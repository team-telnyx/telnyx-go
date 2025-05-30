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
	"bytes"
	"fmt"
)

// checks if the Queue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Queue{}

// Queue struct for Queue
type Queue struct {
	RecordType string `json:"record_type"`
	// Uniquely identifies the queue
	Id string `json:"id"`
	// Name of the queue
	Name string `json:"name"`
	// ISO 8601 formatted date of when the queue was created
	CreatedAt string `json:"created_at"`
	// ISO 8601 formatted date of when the queue was last updated
	UpdatedAt string `json:"updated_at"`
	// The number of calls currently in the queue
	CurrentSize int32 `json:"current_size"`
	// The maximum number of calls allowed in the queue
	MaxSize int32 `json:"max_size"`
	// The average time that the calls currently in the queue have spent waiting, given in seconds.
	AverageWaitTimeSecs int32 `json:"average_wait_time_secs"`
}

type _Queue Queue

// NewQueue instantiates a new Queue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueue(recordType string, id string, name string, createdAt string, updatedAt string, currentSize int32, maxSize int32, averageWaitTimeSecs int32) *Queue {
	this := Queue{}
	this.RecordType = recordType
	this.Id = id
	this.Name = name
	this.CreatedAt = createdAt
	this.UpdatedAt = updatedAt
	this.CurrentSize = currentSize
	this.MaxSize = maxSize
	this.AverageWaitTimeSecs = averageWaitTimeSecs
	return &this
}

// NewQueueWithDefaults instantiates a new Queue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueWithDefaults() *Queue {
	this := Queue{}
	return &this
}

// GetRecordType returns the RecordType field value
func (o *Queue) GetRecordType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value
// and a boolean to check if the value has been set.
func (o *Queue) GetRecordTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RecordType, true
}

// SetRecordType sets field value
func (o *Queue) SetRecordType(v string) {
	o.RecordType = v
}

// GetId returns the Id field value
func (o *Queue) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Queue) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Queue) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Queue) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Queue) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Queue) SetName(v string) {
	o.Name = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Queue) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Queue) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Queue) SetCreatedAt(v string) {
	o.CreatedAt = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *Queue) GetUpdatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *Queue) GetUpdatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *Queue) SetUpdatedAt(v string) {
	o.UpdatedAt = v
}

// GetCurrentSize returns the CurrentSize field value
func (o *Queue) GetCurrentSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.CurrentSize
}

// GetCurrentSizeOk returns a tuple with the CurrentSize field value
// and a boolean to check if the value has been set.
func (o *Queue) GetCurrentSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CurrentSize, true
}

// SetCurrentSize sets field value
func (o *Queue) SetCurrentSize(v int32) {
	o.CurrentSize = v
}

// GetMaxSize returns the MaxSize field value
func (o *Queue) GetMaxSize() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxSize
}

// GetMaxSizeOk returns a tuple with the MaxSize field value
// and a boolean to check if the value has been set.
func (o *Queue) GetMaxSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxSize, true
}

// SetMaxSize sets field value
func (o *Queue) SetMaxSize(v int32) {
	o.MaxSize = v
}

// GetAverageWaitTimeSecs returns the AverageWaitTimeSecs field value
func (o *Queue) GetAverageWaitTimeSecs() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AverageWaitTimeSecs
}

// GetAverageWaitTimeSecsOk returns a tuple with the AverageWaitTimeSecs field value
// and a boolean to check if the value has been set.
func (o *Queue) GetAverageWaitTimeSecsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AverageWaitTimeSecs, true
}

// SetAverageWaitTimeSecs sets field value
func (o *Queue) SetAverageWaitTimeSecs(v int32) {
	o.AverageWaitTimeSecs = v
}

func (o Queue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Queue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["record_type"] = o.RecordType
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	toSerialize["created_at"] = o.CreatedAt
	toSerialize["updated_at"] = o.UpdatedAt
	toSerialize["current_size"] = o.CurrentSize
	toSerialize["max_size"] = o.MaxSize
	toSerialize["average_wait_time_secs"] = o.AverageWaitTimeSecs
	return toSerialize, nil
}

func (o *Queue) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"record_type",
		"id",
		"name",
		"created_at",
		"updated_at",
		"current_size",
		"max_size",
		"average_wait_time_secs",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varQueue := _Queue{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varQueue)

	if err != nil {
		return err
	}

	*o = Queue(varQueue)

	return err
}

type NullableQueue struct {
	value *Queue
	isSet bool
}

func (v NullableQueue) Get() *Queue {
	return v.value
}

func (v *NullableQueue) Set(val *Queue) {
	v.value = val
	v.isSet = true
}

func (v NullableQueue) IsSet() bool {
	return v.isSet
}

func (v *NullableQueue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueue(val *Queue) *NullableQueue {
	return &NullableQueue{value: val, isSet: true}
}

func (v NullableQueue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


