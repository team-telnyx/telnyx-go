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

// checks if the CreateAdditionalDocumentsRequestAdditionalDocumentsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateAdditionalDocumentsRequestAdditionalDocumentsInner{}

// CreateAdditionalDocumentsRequestAdditionalDocumentsInner struct for CreateAdditionalDocumentsRequestAdditionalDocumentsInner
type CreateAdditionalDocumentsRequestAdditionalDocumentsInner struct {
	// The type of document being created.
	DocumentType *string `json:"document_type,omitempty"`
	// The document identification
	DocumentId *string `json:"document_id,omitempty"`
}

// NewCreateAdditionalDocumentsRequestAdditionalDocumentsInner instantiates a new CreateAdditionalDocumentsRequestAdditionalDocumentsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAdditionalDocumentsRequestAdditionalDocumentsInner() *CreateAdditionalDocumentsRequestAdditionalDocumentsInner {
	this := CreateAdditionalDocumentsRequestAdditionalDocumentsInner{}
	return &this
}

// NewCreateAdditionalDocumentsRequestAdditionalDocumentsInnerWithDefaults instantiates a new CreateAdditionalDocumentsRequestAdditionalDocumentsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAdditionalDocumentsRequestAdditionalDocumentsInnerWithDefaults() *CreateAdditionalDocumentsRequestAdditionalDocumentsInner {
	this := CreateAdditionalDocumentsRequestAdditionalDocumentsInner{}
	return &this
}

// GetDocumentType returns the DocumentType field value if set, zero value otherwise.
func (o *CreateAdditionalDocumentsRequestAdditionalDocumentsInner) GetDocumentType() string {
	if o == nil || IsNil(o.DocumentType) {
		var ret string
		return ret
	}
	return *o.DocumentType
}

// GetDocumentTypeOk returns a tuple with the DocumentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAdditionalDocumentsRequestAdditionalDocumentsInner) GetDocumentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentType) {
		return nil, false
	}
	return o.DocumentType, true
}

// HasDocumentType returns a boolean if a field has been set.
func (o *CreateAdditionalDocumentsRequestAdditionalDocumentsInner) HasDocumentType() bool {
	if o != nil && !IsNil(o.DocumentType) {
		return true
	}

	return false
}

// SetDocumentType gets a reference to the given string and assigns it to the DocumentType field.
func (o *CreateAdditionalDocumentsRequestAdditionalDocumentsInner) SetDocumentType(v string) {
	o.DocumentType = &v
}

// GetDocumentId returns the DocumentId field value if set, zero value otherwise.
func (o *CreateAdditionalDocumentsRequestAdditionalDocumentsInner) GetDocumentId() string {
	if o == nil || IsNil(o.DocumentId) {
		var ret string
		return ret
	}
	return *o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAdditionalDocumentsRequestAdditionalDocumentsInner) GetDocumentIdOk() (*string, bool) {
	if o == nil || IsNil(o.DocumentId) {
		return nil, false
	}
	return o.DocumentId, true
}

// HasDocumentId returns a boolean if a field has been set.
func (o *CreateAdditionalDocumentsRequestAdditionalDocumentsInner) HasDocumentId() bool {
	if o != nil && !IsNil(o.DocumentId) {
		return true
	}

	return false
}

// SetDocumentId gets a reference to the given string and assigns it to the DocumentId field.
func (o *CreateAdditionalDocumentsRequestAdditionalDocumentsInner) SetDocumentId(v string) {
	o.DocumentId = &v
}

func (o CreateAdditionalDocumentsRequestAdditionalDocumentsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateAdditionalDocumentsRequestAdditionalDocumentsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DocumentType) {
		toSerialize["document_type"] = o.DocumentType
	}
	if !IsNil(o.DocumentId) {
		toSerialize["document_id"] = o.DocumentId
	}
	return toSerialize, nil
}

type NullableCreateAdditionalDocumentsRequestAdditionalDocumentsInner struct {
	value *CreateAdditionalDocumentsRequestAdditionalDocumentsInner
	isSet bool
}

func (v NullableCreateAdditionalDocumentsRequestAdditionalDocumentsInner) Get() *CreateAdditionalDocumentsRequestAdditionalDocumentsInner {
	return v.value
}

func (v *NullableCreateAdditionalDocumentsRequestAdditionalDocumentsInner) Set(val *CreateAdditionalDocumentsRequestAdditionalDocumentsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAdditionalDocumentsRequestAdditionalDocumentsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAdditionalDocumentsRequestAdditionalDocumentsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAdditionalDocumentsRequestAdditionalDocumentsInner(val *CreateAdditionalDocumentsRequestAdditionalDocumentsInner) *NullableCreateAdditionalDocumentsRequestAdditionalDocumentsInner {
	return &NullableCreateAdditionalDocumentsRequestAdditionalDocumentsInner{value: val, isSet: true}
}

func (v NullableCreateAdditionalDocumentsRequestAdditionalDocumentsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAdditionalDocumentsRequestAdditionalDocumentsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


