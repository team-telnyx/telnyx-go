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

// checks if the PortoutComment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PortoutComment{}

// PortoutComment struct for PortoutComment
type PortoutComment struct {
	Id string `json:"id"`
	// Identifies the type of the resource.
	RecordType *string `json:"record_type,omitempty"`
	// Comment body
	Body string `json:"body"`
	// Identifies the associated port request
	PortoutId *string `json:"portout_id,omitempty"`
	// Identifies the user who created the comment. Will be null if created by Telnyx Admin
	UserId string `json:"user_id"`
	// Comment creation timestamp in ISO 8601 format
	CreatedAt string `json:"created_at"`
}

type _PortoutComment PortoutComment

// NewPortoutComment instantiates a new PortoutComment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPortoutComment(id string, body string, userId string, createdAt string) *PortoutComment {
	this := PortoutComment{}
	this.Id = id
	this.Body = body
	this.UserId = userId
	this.CreatedAt = createdAt
	return &this
}

// NewPortoutCommentWithDefaults instantiates a new PortoutComment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPortoutCommentWithDefaults() *PortoutComment {
	this := PortoutComment{}
	return &this
}

// GetId returns the Id field value
func (o *PortoutComment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PortoutComment) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PortoutComment) SetId(v string) {
	o.Id = v
}

// GetRecordType returns the RecordType field value if set, zero value otherwise.
func (o *PortoutComment) GetRecordType() string {
	if o == nil || IsNil(o.RecordType) {
		var ret string
		return ret
	}
	return *o.RecordType
}

// GetRecordTypeOk returns a tuple with the RecordType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortoutComment) GetRecordTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RecordType) {
		return nil, false
	}
	return o.RecordType, true
}

// HasRecordType returns a boolean if a field has been set.
func (o *PortoutComment) HasRecordType() bool {
	if o != nil && !IsNil(o.RecordType) {
		return true
	}

	return false
}

// SetRecordType gets a reference to the given string and assigns it to the RecordType field.
func (o *PortoutComment) SetRecordType(v string) {
	o.RecordType = &v
}

// GetBody returns the Body field value
func (o *PortoutComment) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *PortoutComment) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *PortoutComment) SetBody(v string) {
	o.Body = v
}

// GetPortoutId returns the PortoutId field value if set, zero value otherwise.
func (o *PortoutComment) GetPortoutId() string {
	if o == nil || IsNil(o.PortoutId) {
		var ret string
		return ret
	}
	return *o.PortoutId
}

// GetPortoutIdOk returns a tuple with the PortoutId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PortoutComment) GetPortoutIdOk() (*string, bool) {
	if o == nil || IsNil(o.PortoutId) {
		return nil, false
	}
	return o.PortoutId, true
}

// HasPortoutId returns a boolean if a field has been set.
func (o *PortoutComment) HasPortoutId() bool {
	if o != nil && !IsNil(o.PortoutId) {
		return true
	}

	return false
}

// SetPortoutId gets a reference to the given string and assigns it to the PortoutId field.
func (o *PortoutComment) SetPortoutId(v string) {
	o.PortoutId = &v
}

// GetUserId returns the UserId field value
func (o *PortoutComment) GetUserId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value
// and a boolean to check if the value has been set.
func (o *PortoutComment) GetUserIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserId, true
}

// SetUserId sets field value
func (o *PortoutComment) SetUserId(v string) {
	o.UserId = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PortoutComment) GetCreatedAt() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PortoutComment) GetCreatedAtOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PortoutComment) SetCreatedAt(v string) {
	o.CreatedAt = v
}

func (o PortoutComment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PortoutComment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.RecordType) {
		toSerialize["record_type"] = o.RecordType
	}
	toSerialize["body"] = o.Body
	if !IsNil(o.PortoutId) {
		toSerialize["portout_id"] = o.PortoutId
	}
	toSerialize["user_id"] = o.UserId
	toSerialize["created_at"] = o.CreatedAt
	return toSerialize, nil
}

func (o *PortoutComment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"body",
		"user_id",
		"created_at",
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

	varPortoutComment := _PortoutComment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPortoutComment)

	if err != nil {
		return err
	}

	*o = PortoutComment(varPortoutComment)

	return err
}

type NullablePortoutComment struct {
	value *PortoutComment
	isSet bool
}

func (v NullablePortoutComment) Get() *PortoutComment {
	return v.value
}

func (v *NullablePortoutComment) Set(val *PortoutComment) {
	v.value = val
	v.isSet = true
}

func (v NullablePortoutComment) IsSet() bool {
	return v.isSet
}

func (v *NullablePortoutComment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePortoutComment(val *PortoutComment) *NullablePortoutComment {
	return &NullablePortoutComment{value: val, isSet: true}
}

func (v NullablePortoutComment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePortoutComment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


