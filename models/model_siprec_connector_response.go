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

// checks if the SiprecConnectorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiprecConnectorResponse{}

// SiprecConnectorResponse struct for SiprecConnectorResponse
type SiprecConnectorResponse struct {
	Data SIPRECConnector `json:"data"`
}

type _SiprecConnectorResponse SiprecConnectorResponse

// NewSiprecConnectorResponse instantiates a new SiprecConnectorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiprecConnectorResponse(data SIPRECConnector) *SiprecConnectorResponse {
	this := SiprecConnectorResponse{}
	this.Data = data
	return &this
}

// NewSiprecConnectorResponseWithDefaults instantiates a new SiprecConnectorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiprecConnectorResponseWithDefaults() *SiprecConnectorResponse {
	this := SiprecConnectorResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SiprecConnectorResponse) GetData() SIPRECConnector {
	if o == nil {
		var ret SIPRECConnector
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SiprecConnectorResponse) GetDataOk() (*SIPRECConnector, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SiprecConnectorResponse) SetData(v SIPRECConnector) {
	o.Data = v
}

func (o SiprecConnectorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiprecConnectorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SiprecConnectorResponse) UnmarshalJSON(data []byte) (err error) {
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

	varSiprecConnectorResponse := _SiprecConnectorResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSiprecConnectorResponse)

	if err != nil {
		return err
	}

	*o = SiprecConnectorResponse(varSiprecConnectorResponse)

	return err
}

type NullableSiprecConnectorResponse struct {
	value *SiprecConnectorResponse
	isSet bool
}

func (v NullableSiprecConnectorResponse) Get() *SiprecConnectorResponse {
	return v.value
}

func (v *NullableSiprecConnectorResponse) Set(val *SiprecConnectorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSiprecConnectorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSiprecConnectorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiprecConnectorResponse(val *SiprecConnectorResponse) *NullableSiprecConnectorResponse {
	return &NullableSiprecConnectorResponse{value: val, isSet: true}
}

func (v NullableSiprecConnectorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiprecConnectorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


