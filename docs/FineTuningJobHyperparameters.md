# FineTuningJobHyperparameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NEpochs** | **int32** | The number of epochs to train the model for. An epoch refers to one full cycle through the training dataset. | [default to 3]

## Methods

### NewFineTuningJobHyperparameters

`func NewFineTuningJobHyperparameters(nEpochs int32, ) *FineTuningJobHyperparameters`

NewFineTuningJobHyperparameters instantiates a new FineTuningJobHyperparameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFineTuningJobHyperparametersWithDefaults

`func NewFineTuningJobHyperparametersWithDefaults() *FineTuningJobHyperparameters`

NewFineTuningJobHyperparametersWithDefaults instantiates a new FineTuningJobHyperparameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNEpochs

`func (o *FineTuningJobHyperparameters) GetNEpochs() int32`

GetNEpochs returns the NEpochs field if non-nil, zero value otherwise.

### GetNEpochsOk

`func (o *FineTuningJobHyperparameters) GetNEpochsOk() (*int32, bool)`

GetNEpochsOk returns a tuple with the NEpochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNEpochs

`func (o *FineTuningJobHyperparameters) SetNEpochs(v int32)`

SetNEpochs sets NEpochs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


