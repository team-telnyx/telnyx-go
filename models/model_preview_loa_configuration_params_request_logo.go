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

// checks if the PreviewLoaConfigurationParamsRequestLogo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviewLoaConfigurationParamsRequestLogo{}

// PreviewLoaConfigurationParamsRequestLogo The logo of the LOA configuration
type PreviewLoaConfigurationParamsRequestLogo struct {
	// The document identification
	DocumentId string `json:"document_id"`
}

type _PreviewLoaConfigurationParamsRequestLogo PreviewLoaConfigurationParamsRequestLogo

// NewPreviewLoaConfigurationParamsRequestLogo instantiates a new PreviewLoaConfigurationParamsRequestLogo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviewLoaConfigurationParamsRequestLogo(documentId string) *PreviewLoaConfigurationParamsRequestLogo {
	this := PreviewLoaConfigurationParamsRequestLogo{}
	this.DocumentId = documentId
	return &this
}

// NewPreviewLoaConfigurationParamsRequestLogoWithDefaults instantiates a new PreviewLoaConfigurationParamsRequestLogo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviewLoaConfigurationParamsRequestLogoWithDefaults() *PreviewLoaConfigurationParamsRequestLogo {
	this := PreviewLoaConfigurationParamsRequestLogo{}
	return &this
}

// GetDocumentId returns the DocumentId field value
func (o *PreviewLoaConfigurationParamsRequestLogo) GetDocumentId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentId
}

// GetDocumentIdOk returns a tuple with the DocumentId field value
// and a boolean to check if the value has been set.
func (o *PreviewLoaConfigurationParamsRequestLogo) GetDocumentIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentId, true
}

// SetDocumentId sets field value
func (o *PreviewLoaConfigurationParamsRequestLogo) SetDocumentId(v string) {
	o.DocumentId = v
}

func (o PreviewLoaConfigurationParamsRequestLogo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreviewLoaConfigurationParamsRequestLogo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["document_id"] = o.DocumentId
	return toSerialize, nil
}

func (o *PreviewLoaConfigurationParamsRequestLogo) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"document_id",
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

	varPreviewLoaConfigurationParamsRequestLogo := _PreviewLoaConfigurationParamsRequestLogo{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPreviewLoaConfigurationParamsRequestLogo)

	if err != nil {
		return err
	}

	*o = PreviewLoaConfigurationParamsRequestLogo(varPreviewLoaConfigurationParamsRequestLogo)

	return err
}

type NullablePreviewLoaConfigurationParamsRequestLogo struct {
	value *PreviewLoaConfigurationParamsRequestLogo
	isSet bool
}

func (v NullablePreviewLoaConfigurationParamsRequestLogo) Get() *PreviewLoaConfigurationParamsRequestLogo {
	return v.value
}

func (v *NullablePreviewLoaConfigurationParamsRequestLogo) Set(val *PreviewLoaConfigurationParamsRequestLogo) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviewLoaConfigurationParamsRequestLogo) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviewLoaConfigurationParamsRequestLogo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviewLoaConfigurationParamsRequestLogo(val *PreviewLoaConfigurationParamsRequestLogo) *NullablePreviewLoaConfigurationParamsRequestLogo {
	return &NullablePreviewLoaConfigurationParamsRequestLogo{value: val, isSet: true}
}

func (v NullablePreviewLoaConfigurationParamsRequestLogo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviewLoaConfigurationParamsRequestLogo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


