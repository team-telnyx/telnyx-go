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

// checks if the EmbeddingSimilaritySearchResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbeddingSimilaritySearchResponse{}

// EmbeddingSimilaritySearchResponse struct for EmbeddingSimilaritySearchResponse
type EmbeddingSimilaritySearchResponse struct {
	Data []EmbeddingSimilaritySearchDocument `json:"data"`
}

type _EmbeddingSimilaritySearchResponse EmbeddingSimilaritySearchResponse

// NewEmbeddingSimilaritySearchResponse instantiates a new EmbeddingSimilaritySearchResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddingSimilaritySearchResponse(data []EmbeddingSimilaritySearchDocument) *EmbeddingSimilaritySearchResponse {
	this := EmbeddingSimilaritySearchResponse{}
	this.Data = data
	return &this
}

// NewEmbeddingSimilaritySearchResponseWithDefaults instantiates a new EmbeddingSimilaritySearchResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddingSimilaritySearchResponseWithDefaults() *EmbeddingSimilaritySearchResponse {
	this := EmbeddingSimilaritySearchResponse{}
	return &this
}

// GetData returns the Data field value
func (o *EmbeddingSimilaritySearchResponse) GetData() []EmbeddingSimilaritySearchDocument {
	if o == nil {
		var ret []EmbeddingSimilaritySearchDocument
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EmbeddingSimilaritySearchResponse) GetDataOk() ([]EmbeddingSimilaritySearchDocument, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *EmbeddingSimilaritySearchResponse) SetData(v []EmbeddingSimilaritySearchDocument) {
	o.Data = v
}

func (o EmbeddingSimilaritySearchResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbeddingSimilaritySearchResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *EmbeddingSimilaritySearchResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varEmbeddingSimilaritySearchResponse := _EmbeddingSimilaritySearchResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmbeddingSimilaritySearchResponse)

	if err != nil {
		return err
	}

	*o = EmbeddingSimilaritySearchResponse(varEmbeddingSimilaritySearchResponse)

	return err
}

type NullableEmbeddingSimilaritySearchResponse struct {
	value *EmbeddingSimilaritySearchResponse
	isSet bool
}

func (v NullableEmbeddingSimilaritySearchResponse) Get() *EmbeddingSimilaritySearchResponse {
	return v.value
}

func (v *NullableEmbeddingSimilaritySearchResponse) Set(val *EmbeddingSimilaritySearchResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddingSimilaritySearchResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddingSimilaritySearchResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddingSimilaritySearchResponse(val *EmbeddingSimilaritySearchResponse) *NullableEmbeddingSimilaritySearchResponse {
	return &NullableEmbeddingSimilaritySearchResponse{value: val, isSet: true}
}

func (v NullableEmbeddingSimilaritySearchResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddingSimilaritySearchResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


