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

// checks if the CreateDocumentRequestOneOf1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateDocumentRequestOneOf1{}

// CreateDocumentRequestOneOf1 struct for CreateDocumentRequestOneOf1
type CreateDocumentRequestOneOf1 struct {
	// The Base64 encoded contents of the file you are uploading.
	File string `json:"file"`
	// The filename of the document.
	Filename *string `json:"filename,omitempty"`
	// A customer reference string for customer look ups.
	CustomerReference *string `json:"customer_reference,omitempty"`
}

type _CreateDocumentRequestOneOf1 CreateDocumentRequestOneOf1

// NewCreateDocumentRequestOneOf1 instantiates a new CreateDocumentRequestOneOf1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDocumentRequestOneOf1(file string) *CreateDocumentRequestOneOf1 {
	this := CreateDocumentRequestOneOf1{}
	this.File = file
	return &this
}

// NewCreateDocumentRequestOneOf1WithDefaults instantiates a new CreateDocumentRequestOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDocumentRequestOneOf1WithDefaults() *CreateDocumentRequestOneOf1 {
	this := CreateDocumentRequestOneOf1{}
	return &this
}

// GetFile returns the File field value
func (o *CreateDocumentRequestOneOf1) GetFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequestOneOf1) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *CreateDocumentRequestOneOf1) SetFile(v string) {
	o.File = v
}

// GetFilename returns the Filename field value if set, zero value otherwise.
func (o *CreateDocumentRequestOneOf1) GetFilename() string {
	if o == nil || IsNil(o.Filename) {
		var ret string
		return ret
	}
	return *o.Filename
}

// GetFilenameOk returns a tuple with the Filename field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequestOneOf1) GetFilenameOk() (*string, bool) {
	if o == nil || IsNil(o.Filename) {
		return nil, false
	}
	return o.Filename, true
}

// HasFilename returns a boolean if a field has been set.
func (o *CreateDocumentRequestOneOf1) HasFilename() bool {
	if o != nil && !IsNil(o.Filename) {
		return true
	}

	return false
}

// SetFilename gets a reference to the given string and assigns it to the Filename field.
func (o *CreateDocumentRequestOneOf1) SetFilename(v string) {
	o.Filename = &v
}

// GetCustomerReference returns the CustomerReference field value if set, zero value otherwise.
func (o *CreateDocumentRequestOneOf1) GetCustomerReference() string {
	if o == nil || IsNil(o.CustomerReference) {
		var ret string
		return ret
	}
	return *o.CustomerReference
}

// GetCustomerReferenceOk returns a tuple with the CustomerReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDocumentRequestOneOf1) GetCustomerReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerReference) {
		return nil, false
	}
	return o.CustomerReference, true
}

// HasCustomerReference returns a boolean if a field has been set.
func (o *CreateDocumentRequestOneOf1) HasCustomerReference() bool {
	if o != nil && !IsNil(o.CustomerReference) {
		return true
	}

	return false
}

// SetCustomerReference gets a reference to the given string and assigns it to the CustomerReference field.
func (o *CreateDocumentRequestOneOf1) SetCustomerReference(v string) {
	o.CustomerReference = &v
}

func (o CreateDocumentRequestOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateDocumentRequestOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["file"] = o.File
	if !IsNil(o.Filename) {
		toSerialize["filename"] = o.Filename
	}
	if !IsNil(o.CustomerReference) {
		toSerialize["customer_reference"] = o.CustomerReference
	}
	return toSerialize, nil
}

func (o *CreateDocumentRequestOneOf1) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"file",
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

	varCreateDocumentRequestOneOf1 := _CreateDocumentRequestOneOf1{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateDocumentRequestOneOf1)

	if err != nil {
		return err
	}

	*o = CreateDocumentRequestOneOf1(varCreateDocumentRequestOneOf1)

	return err
}

type NullableCreateDocumentRequestOneOf1 struct {
	value *CreateDocumentRequestOneOf1
	isSet bool
}

func (v NullableCreateDocumentRequestOneOf1) Get() *CreateDocumentRequestOneOf1 {
	return v.value
}

func (v *NullableCreateDocumentRequestOneOf1) Set(val *CreateDocumentRequestOneOf1) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDocumentRequestOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDocumentRequestOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDocumentRequestOneOf1(val *CreateDocumentRequestOneOf1) *NullableCreateDocumentRequestOneOf1 {
	return &NullableCreateDocumentRequestOneOf1{value: val, isSet: true}
}

func (v NullableCreateDocumentRequestOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDocumentRequestOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


