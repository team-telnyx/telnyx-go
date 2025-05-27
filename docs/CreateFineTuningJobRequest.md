# CreateFineTuningJobRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** | The base model that is being fine-tuned. | 
**TrainingFile** | **string** | The storage bucket or object used for training. | 
**Suffix** | Pointer to **string** | Optional suffix to append to the fine tuned model&#39;s name. | [optional] 
**Hyperparameters** | Pointer to [**CreateFineTuningJobRequestHyperparameters**](CreateFineTuningJobRequestHyperparameters.md) |  | [optional] 

## Methods

### NewCreateFineTuningJobRequest

`func NewCreateFineTuningJobRequest(model string, trainingFile string, ) *CreateFineTuningJobRequest`

NewCreateFineTuningJobRequest instantiates a new CreateFineTuningJobRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFineTuningJobRequestWithDefaults

`func NewCreateFineTuningJobRequestWithDefaults() *CreateFineTuningJobRequest`

NewCreateFineTuningJobRequestWithDefaults instantiates a new CreateFineTuningJobRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *CreateFineTuningJobRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CreateFineTuningJobRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CreateFineTuningJobRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetTrainingFile

`func (o *CreateFineTuningJobRequest) GetTrainingFile() string`

GetTrainingFile returns the TrainingFile field if non-nil, zero value otherwise.

### GetTrainingFileOk

`func (o *CreateFineTuningJobRequest) GetTrainingFileOk() (*string, bool)`

GetTrainingFileOk returns a tuple with the TrainingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingFile

`func (o *CreateFineTuningJobRequest) SetTrainingFile(v string)`

SetTrainingFile sets TrainingFile field to given value.


### GetSuffix

`func (o *CreateFineTuningJobRequest) GetSuffix() string`

GetSuffix returns the Suffix field if non-nil, zero value otherwise.

### GetSuffixOk

`func (o *CreateFineTuningJobRequest) GetSuffixOk() (*string, bool)`

GetSuffixOk returns a tuple with the Suffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuffix

`func (o *CreateFineTuningJobRequest) SetSuffix(v string)`

SetSuffix sets Suffix field to given value.

### HasSuffix

`func (o *CreateFineTuningJobRequest) HasSuffix() bool`

HasSuffix returns a boolean if a field has been set.

### GetHyperparameters

`func (o *CreateFineTuningJobRequest) GetHyperparameters() CreateFineTuningJobRequestHyperparameters`

GetHyperparameters returns the Hyperparameters field if non-nil, zero value otherwise.

### GetHyperparametersOk

`func (o *CreateFineTuningJobRequest) GetHyperparametersOk() (*CreateFineTuningJobRequestHyperparameters, bool)`

GetHyperparametersOk returns a tuple with the Hyperparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperparameters

`func (o *CreateFineTuningJobRequest) SetHyperparameters(v CreateFineTuningJobRequestHyperparameters)`

SetHyperparameters sets Hyperparameters field to given value.

### HasHyperparameters

`func (o *CreateFineTuningJobRequest) HasHyperparameters() bool`

HasHyperparameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


