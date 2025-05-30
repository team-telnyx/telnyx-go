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

// checks if the CreateUploadRequestResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateUploadRequestResponse{}

// CreateUploadRequestResponse struct for CreateUploadRequestResponse
type CreateUploadRequestResponse struct {
	// Describes wether or not the operation was successful
	Success *bool `json:"success,omitempty"`
	// Ticket id of the upload request
	TicketId *string `json:"ticket_id,omitempty"`
}

// NewCreateUploadRequestResponse instantiates a new CreateUploadRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateUploadRequestResponse() *CreateUploadRequestResponse {
	this := CreateUploadRequestResponse{}
	return &this
}

// NewCreateUploadRequestResponseWithDefaults instantiates a new CreateUploadRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateUploadRequestResponseWithDefaults() *CreateUploadRequestResponse {
	this := CreateUploadRequestResponse{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *CreateUploadRequestResponse) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUploadRequestResponse) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *CreateUploadRequestResponse) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *CreateUploadRequestResponse) SetSuccess(v bool) {
	o.Success = &v
}

// GetTicketId returns the TicketId field value if set, zero value otherwise.
func (o *CreateUploadRequestResponse) GetTicketId() string {
	if o == nil || IsNil(o.TicketId) {
		var ret string
		return ret
	}
	return *o.TicketId
}

// GetTicketIdOk returns a tuple with the TicketId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateUploadRequestResponse) GetTicketIdOk() (*string, bool) {
	if o == nil || IsNil(o.TicketId) {
		return nil, false
	}
	return o.TicketId, true
}

// HasTicketId returns a boolean if a field has been set.
func (o *CreateUploadRequestResponse) HasTicketId() bool {
	if o != nil && !IsNil(o.TicketId) {
		return true
	}

	return false
}

// SetTicketId gets a reference to the given string and assigns it to the TicketId field.
func (o *CreateUploadRequestResponse) SetTicketId(v string) {
	o.TicketId = &v
}

func (o CreateUploadRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateUploadRequestResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.TicketId) {
		toSerialize["ticket_id"] = o.TicketId
	}
	return toSerialize, nil
}

type NullableCreateUploadRequestResponse struct {
	value *CreateUploadRequestResponse
	isSet bool
}

func (v NullableCreateUploadRequestResponse) Get() *CreateUploadRequestResponse {
	return v.value
}

func (v *NullableCreateUploadRequestResponse) Set(val *CreateUploadRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateUploadRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateUploadRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateUploadRequestResponse(val *CreateUploadRequestResponse) *NullableCreateUploadRequestResponse {
	return &NullableCreateUploadRequestResponse{value: val, isSet: true}
}

func (v NullableCreateUploadRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateUploadRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


