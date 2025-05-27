# TaskStatusResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | Pointer to **string** |  | [optional] 
**TaskName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**BackgroundTaskStatus**](BackgroundTaskStatus.md) |  | [optional] 
**CreatedAt** | Pointer to **string** |  | [optional] 
**FinishedAt** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskStatusResponseData

`func NewTaskStatusResponseData() *TaskStatusResponseData`

NewTaskStatusResponseData instantiates a new TaskStatusResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskStatusResponseDataWithDefaults

`func NewTaskStatusResponseDataWithDefaults() *TaskStatusResponseData`

NewTaskStatusResponseDataWithDefaults instantiates a new TaskStatusResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *TaskStatusResponseData) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *TaskStatusResponseData) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *TaskStatusResponseData) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *TaskStatusResponseData) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetTaskName

`func (o *TaskStatusResponseData) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *TaskStatusResponseData) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *TaskStatusResponseData) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.

### HasTaskName

`func (o *TaskStatusResponseData) HasTaskName() bool`

HasTaskName returns a boolean if a field has been set.

### GetStatus

`func (o *TaskStatusResponseData) GetStatus() BackgroundTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TaskStatusResponseData) GetStatusOk() (*BackgroundTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TaskStatusResponseData) SetStatus(v BackgroundTaskStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TaskStatusResponseData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TaskStatusResponseData) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TaskStatusResponseData) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TaskStatusResponseData) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TaskStatusResponseData) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetFinishedAt

`func (o *TaskStatusResponseData) GetFinishedAt() string`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *TaskStatusResponseData) GetFinishedAtOk() (*string, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *TaskStatusResponseData) SetFinishedAt(v string)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *TaskStatusResponseData) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


