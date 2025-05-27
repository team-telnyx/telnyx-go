# BackgroundTasksQueryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** |  | 
**TaskId** | **string** |  | 
**TaskName** | **string** |  | 
**Status** | [**BackgroundTaskStatus**](BackgroundTaskStatus.md) |  | 
**CreatedAt** | **time.Time** |  | 
**FinishedAt** | Pointer to **time.Time** |  | [optional] 
**Bucket** | Pointer to **string** |  | [optional] 

## Methods

### NewBackgroundTasksQueryResponse

`func NewBackgroundTasksQueryResponse(userId string, taskId string, taskName string, status BackgroundTaskStatus, createdAt time.Time, ) *BackgroundTasksQueryResponse`

NewBackgroundTasksQueryResponse instantiates a new BackgroundTasksQueryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBackgroundTasksQueryResponseWithDefaults

`func NewBackgroundTasksQueryResponseWithDefaults() *BackgroundTasksQueryResponse`

NewBackgroundTasksQueryResponseWithDefaults instantiates a new BackgroundTasksQueryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *BackgroundTasksQueryResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *BackgroundTasksQueryResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *BackgroundTasksQueryResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetTaskId

`func (o *BackgroundTasksQueryResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *BackgroundTasksQueryResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *BackgroundTasksQueryResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetTaskName

`func (o *BackgroundTasksQueryResponse) GetTaskName() string`

GetTaskName returns the TaskName field if non-nil, zero value otherwise.

### GetTaskNameOk

`func (o *BackgroundTasksQueryResponse) GetTaskNameOk() (*string, bool)`

GetTaskNameOk returns a tuple with the TaskName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskName

`func (o *BackgroundTasksQueryResponse) SetTaskName(v string)`

SetTaskName sets TaskName field to given value.


### GetStatus

`func (o *BackgroundTasksQueryResponse) GetStatus() BackgroundTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BackgroundTasksQueryResponse) GetStatusOk() (*BackgroundTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BackgroundTasksQueryResponse) SetStatus(v BackgroundTaskStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *BackgroundTasksQueryResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BackgroundTasksQueryResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BackgroundTasksQueryResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFinishedAt

`func (o *BackgroundTasksQueryResponse) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *BackgroundTasksQueryResponse) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *BackgroundTasksQueryResponse) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *BackgroundTasksQueryResponse) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### GetBucket

`func (o *BackgroundTasksQueryResponse) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *BackgroundTasksQueryResponse) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *BackgroundTasksQueryResponse) SetBucket(v string)`

SetBucket sets Bucket field to given value.

### HasBucket

`func (o *BackgroundTasksQueryResponse) HasBucket() bool`

HasBucket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


