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

// checks if the RetrieveVerificationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetrieveVerificationResponse{}

// RetrieveVerificationResponse struct for RetrieveVerificationResponse
type RetrieveVerificationResponse struct {
	Data Verification `json:"data"`
}

type _RetrieveVerificationResponse RetrieveVerificationResponse

// NewRetrieveVerificationResponse instantiates a new RetrieveVerificationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetrieveVerificationResponse(data Verification) *RetrieveVerificationResponse {
	this := RetrieveVerificationResponse{}
	this.Data = data
	return &this
}

// NewRetrieveVerificationResponseWithDefaults instantiates a new RetrieveVerificationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetrieveVerificationResponseWithDefaults() *RetrieveVerificationResponse {
	this := RetrieveVerificationResponse{}
	return &this
}

// GetData returns the Data field value
func (o *RetrieveVerificationResponse) GetData() Verification {
	if o == nil {
		var ret Verification
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RetrieveVerificationResponse) GetDataOk() (*Verification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *RetrieveVerificationResponse) SetData(v Verification) {
	o.Data = v
}

func (o RetrieveVerificationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetrieveVerificationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *RetrieveVerificationResponse) UnmarshalJSON(data []byte) (err error) {
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

	varRetrieveVerificationResponse := _RetrieveVerificationResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRetrieveVerificationResponse)

	if err != nil {
		return err
	}

	*o = RetrieveVerificationResponse(varRetrieveVerificationResponse)

	return err
}

type NullableRetrieveVerificationResponse struct {
	value *RetrieveVerificationResponse
	isSet bool
}

func (v NullableRetrieveVerificationResponse) Get() *RetrieveVerificationResponse {
	return v.value
}

func (v *NullableRetrieveVerificationResponse) Set(val *RetrieveVerificationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrieveVerificationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrieveVerificationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrieveVerificationResponse(val *RetrieveVerificationResponse) *NullableRetrieveVerificationResponse {
	return &NullableRetrieveVerificationResponse{value: val, isSet: true}
}

func (v NullableRetrieveVerificationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetrieveVerificationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


