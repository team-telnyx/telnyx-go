# FineTuningJob

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The name of the fine-tuned model that is being created. | 
**CreatedAt** | **int32** | The Unix timestamp (in seconds) for when the fine-tuning job was created. | 
**FinishedAt** | **NullableInt32** | The Unix timestamp (in seconds) for when the fine-tuning job was finished. The value will be null if the fine-tuning job is still running. | 
**Hyperparameters** | [**FineTuningJobHyperparameters**](FineTuningJobHyperparameters.md) |  | 
**Model** | **string** | The base model that is being fine-tuned. | 
**OrganizationId** | **string** | The organization that owns the fine-tuning job. | 
**Status** | **string** | The current status of the fine-tuning job. | 
**TrainedTokens** | **NullableInt32** | The total number of billable tokens processed by this fine-tuning job. The value will be null if the fine-tuning job is still running. | 
**TrainingFile** | **string** | The storage bucket or object used for training. | 

## Methods

### NewFineTuningJob

`func NewFineTuningJob(id string, createdAt int32, finishedAt NullableInt32, hyperparameters FineTuningJobHyperparameters, model string, organizationId string, status string, trainedTokens NullableInt32, trainingFile string, ) *FineTuningJob`

NewFineTuningJob instantiates a new FineTuningJob object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFineTuningJobWithDefaults

`func NewFineTuningJobWithDefaults() *FineTuningJob`

NewFineTuningJobWithDefaults instantiates a new FineTuningJob object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FineTuningJob) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FineTuningJob) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FineTuningJob) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *FineTuningJob) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FineTuningJob) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FineTuningJob) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetFinishedAt

`func (o *FineTuningJob) GetFinishedAt() int32`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *FineTuningJob) GetFinishedAtOk() (*int32, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *FineTuningJob) SetFinishedAt(v int32)`

SetFinishedAt sets FinishedAt field to given value.


### SetFinishedAtNil

`func (o *FineTuningJob) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *FineTuningJob) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetHyperparameters

`func (o *FineTuningJob) GetHyperparameters() FineTuningJobHyperparameters`

GetHyperparameters returns the Hyperparameters field if non-nil, zero value otherwise.

### GetHyperparametersOk

`func (o *FineTuningJob) GetHyperparametersOk() (*FineTuningJobHyperparameters, bool)`

GetHyperparametersOk returns a tuple with the Hyperparameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHyperparameters

`func (o *FineTuningJob) SetHyperparameters(v FineTuningJobHyperparameters)`

SetHyperparameters sets Hyperparameters field to given value.


### GetModel

`func (o *FineTuningJob) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *FineTuningJob) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *FineTuningJob) SetModel(v string)`

SetModel sets Model field to given value.


### GetOrganizationId

`func (o *FineTuningJob) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *FineTuningJob) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *FineTuningJob) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetStatus

`func (o *FineTuningJob) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FineTuningJob) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FineTuningJob) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTrainedTokens

`func (o *FineTuningJob) GetTrainedTokens() int32`

GetTrainedTokens returns the TrainedTokens field if non-nil, zero value otherwise.

### GetTrainedTokensOk

`func (o *FineTuningJob) GetTrainedTokensOk() (*int32, bool)`

GetTrainedTokensOk returns a tuple with the TrainedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainedTokens

`func (o *FineTuningJob) SetTrainedTokens(v int32)`

SetTrainedTokens sets TrainedTokens field to given value.


### SetTrainedTokensNil

`func (o *FineTuningJob) SetTrainedTokensNil(b bool)`

 SetTrainedTokensNil sets the value for TrainedTokens to be an explicit nil

### UnsetTrainedTokens
`func (o *FineTuningJob) UnsetTrainedTokens()`

UnsetTrainedTokens ensures that no value is present for TrainedTokens, not even an explicit nil
### GetTrainingFile

`func (o *FineTuningJob) GetTrainingFile() string`

GetTrainingFile returns the TrainingFile field if non-nil, zero value otherwise.

### GetTrainingFileOk

`func (o *FineTuningJob) GetTrainingFileOk() (*string, bool)`

GetTrainingFileOk returns a tuple with the TrainingFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrainingFile

`func (o *FineTuningJob) SetTrainingFile(v string)`

SetTrainingFile sets TrainingFile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


