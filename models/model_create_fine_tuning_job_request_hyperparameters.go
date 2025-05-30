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

// checks if the CreateFineTuningJobRequestHyperparameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateFineTuningJobRequestHyperparameters{}

// CreateFineTuningJobRequestHyperparameters The hyperparameters used for the fine-tuning job.
type CreateFineTuningJobRequestHyperparameters struct {
	// The number of epochs to train the model for. An epoch refers to one full cycle through the training dataset. 'auto' decides the optimal number of epochs based on the size of the dataset. If setting the number manually, we support any number between 1 and 50 epochs.
	NEpochs *int32 `json:"n_epochs,omitempty"`
}

// NewCreateFineTuningJobRequestHyperparameters instantiates a new CreateFineTuningJobRequestHyperparameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFineTuningJobRequestHyperparameters() *CreateFineTuningJobRequestHyperparameters {
	this := CreateFineTuningJobRequestHyperparameters{}
	var nEpochs int32 = 3
	this.NEpochs = &nEpochs
	return &this
}

// NewCreateFineTuningJobRequestHyperparametersWithDefaults instantiates a new CreateFineTuningJobRequestHyperparameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFineTuningJobRequestHyperparametersWithDefaults() *CreateFineTuningJobRequestHyperparameters {
	this := CreateFineTuningJobRequestHyperparameters{}
	var nEpochs int32 = 3
	this.NEpochs = &nEpochs
	return &this
}

// GetNEpochs returns the NEpochs field value if set, zero value otherwise.
func (o *CreateFineTuningJobRequestHyperparameters) GetNEpochs() int32 {
	if o == nil || IsNil(o.NEpochs) {
		var ret int32
		return ret
	}
	return *o.NEpochs
}

// GetNEpochsOk returns a tuple with the NEpochs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateFineTuningJobRequestHyperparameters) GetNEpochsOk() (*int32, bool) {
	if o == nil || IsNil(o.NEpochs) {
		return nil, false
	}
	return o.NEpochs, true
}

// HasNEpochs returns a boolean if a field has been set.
func (o *CreateFineTuningJobRequestHyperparameters) HasNEpochs() bool {
	if o != nil && !IsNil(o.NEpochs) {
		return true
	}

	return false
}

// SetNEpochs gets a reference to the given int32 and assigns it to the NEpochs field.
func (o *CreateFineTuningJobRequestHyperparameters) SetNEpochs(v int32) {
	o.NEpochs = &v
}

func (o CreateFineTuningJobRequestHyperparameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateFineTuningJobRequestHyperparameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NEpochs) {
		toSerialize["n_epochs"] = o.NEpochs
	}
	return toSerialize, nil
}

type NullableCreateFineTuningJobRequestHyperparameters struct {
	value *CreateFineTuningJobRequestHyperparameters
	isSet bool
}

func (v NullableCreateFineTuningJobRequestHyperparameters) Get() *CreateFineTuningJobRequestHyperparameters {
	return v.value
}

func (v *NullableCreateFineTuningJobRequestHyperparameters) Set(val *CreateFineTuningJobRequestHyperparameters) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFineTuningJobRequestHyperparameters) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFineTuningJobRequestHyperparameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFineTuningJobRequestHyperparameters(val *CreateFineTuningJobRequestHyperparameters) *NullableCreateFineTuningJobRequestHyperparameters {
	return &NullableCreateFineTuningJobRequestHyperparameters{value: val, isSet: true}
}

func (v NullableCreateFineTuningJobRequestHyperparameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFineTuningJobRequestHyperparameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


