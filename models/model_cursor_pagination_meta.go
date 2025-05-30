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

// checks if the CursorPaginationMeta type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CursorPaginationMeta{}

// CursorPaginationMeta struct for CursorPaginationMeta
type CursorPaginationMeta struct {
	Cursors *Cursor `json:"cursors,omitempty"`
	// Path to next page.
	Next *string `json:"next,omitempty"`
	// Path to previous page.
	Previous *string `json:"previous,omitempty"`
}

// NewCursorPaginationMeta instantiates a new CursorPaginationMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCursorPaginationMeta() *CursorPaginationMeta {
	this := CursorPaginationMeta{}
	return &this
}

// NewCursorPaginationMetaWithDefaults instantiates a new CursorPaginationMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCursorPaginationMetaWithDefaults() *CursorPaginationMeta {
	this := CursorPaginationMeta{}
	return &this
}

// GetCursors returns the Cursors field value if set, zero value otherwise.
func (o *CursorPaginationMeta) GetCursors() Cursor {
	if o == nil || IsNil(o.Cursors) {
		var ret Cursor
		return ret
	}
	return *o.Cursors
}

// GetCursorsOk returns a tuple with the Cursors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CursorPaginationMeta) GetCursorsOk() (*Cursor, bool) {
	if o == nil || IsNil(o.Cursors) {
		return nil, false
	}
	return o.Cursors, true
}

// HasCursors returns a boolean if a field has been set.
func (o *CursorPaginationMeta) HasCursors() bool {
	if o != nil && !IsNil(o.Cursors) {
		return true
	}

	return false
}

// SetCursors gets a reference to the given Cursor and assigns it to the Cursors field.
func (o *CursorPaginationMeta) SetCursors(v Cursor) {
	o.Cursors = &v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *CursorPaginationMeta) GetNext() string {
	if o == nil || IsNil(o.Next) {
		var ret string
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CursorPaginationMeta) GetNextOk() (*string, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *CursorPaginationMeta) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given string and assigns it to the Next field.
func (o *CursorPaginationMeta) SetNext(v string) {
	o.Next = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *CursorPaginationMeta) GetPrevious() string {
	if o == nil || IsNil(o.Previous) {
		var ret string
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CursorPaginationMeta) GetPreviousOk() (*string, bool) {
	if o == nil || IsNil(o.Previous) {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *CursorPaginationMeta) HasPrevious() bool {
	if o != nil && !IsNil(o.Previous) {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given string and assigns it to the Previous field.
func (o *CursorPaginationMeta) SetPrevious(v string) {
	o.Previous = &v
}

func (o CursorPaginationMeta) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CursorPaginationMeta) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cursors) {
		toSerialize["cursors"] = o.Cursors
	}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	if !IsNil(o.Previous) {
		toSerialize["previous"] = o.Previous
	}
	return toSerialize, nil
}

type NullableCursorPaginationMeta struct {
	value *CursorPaginationMeta
	isSet bool
}

func (v NullableCursorPaginationMeta) Get() *CursorPaginationMeta {
	return v.value
}

func (v *NullableCursorPaginationMeta) Set(val *CursorPaginationMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableCursorPaginationMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableCursorPaginationMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCursorPaginationMeta(val *CursorPaginationMeta) *NullableCursorPaginationMeta {
	return &NullableCursorPaginationMeta{value: val, isSet: true}
}

func (v NullableCursorPaginationMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCursorPaginationMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


