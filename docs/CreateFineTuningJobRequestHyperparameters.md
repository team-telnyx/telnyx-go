# CreateFineTuningJobRequestHyperparameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NEpochs** | Pointer to **int32** | The number of epochs to train the model for. An epoch refers to one full cycle through the training dataset. &#39;auto&#39; decides the optimal number of epochs based on the size of the dataset. If setting the number manually, we support any number between 1 and 50 epochs. | [optional] [default to 3]

## Methods

### NewCreateFineTuningJobRequestHyperparameters

`func NewCreateFineTuningJobRequestHyperparameters() *CreateFineTuningJobRequestHyperparameters`

NewCreateFineTuningJobRequestHyperparameters instantiates a new CreateFineTuningJobRequestHyperparameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFineTuningJobRequestHyperparametersWithDefaults

`func NewCreateFineTuningJobRequestHyperparametersWithDefaults() *CreateFineTuningJobRequestHyperparameters`

NewCreateFineTuningJobRequestHyperparametersWithDefaults instantiates a new CreateFineTuningJobRequestHyperparameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNEpochs

`func (o *CreateFineTuningJobRequestHyperparameters) GetNEpochs() int32`

GetNEpochs returns the NEpochs field if non-nil, zero value otherwise.

### GetNEpochsOk

`func (o *CreateFineTuningJobRequestHyperparameters) GetNEpochsOk() (*int32, bool)`

GetNEpochsOk returns a tuple with the NEpochs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNEpochs

`func (o *CreateFineTuningJobRequestHyperparameters) SetNEpochs(v int32)`

SetNEpochs sets NEpochs field to given value.

### HasNEpochs

`func (o *CreateFineTuningJobRequestHyperparameters) HasNEpochs() bool`

HasNEpochs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


