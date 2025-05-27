# AssignmentTaskStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** |  | 
**Status** | [**TaskStatus**](TaskStatus.md) | The status of the task associated with assigning a messaging profile to a campaign. | 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAssignmentTaskStatusResponse

`func NewAssignmentTaskStatusResponse(taskId string, status TaskStatus, ) *AssignmentTaskStatusResponse`

NewAssignmentTaskStatusResponse instantiates a new AssignmentTaskStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignmentTaskStatusResponseWithDefaults

`func NewAssignmentTaskStatusResponseWithDefaults() *AssignmentTaskStatusResponse`

NewAssignmentTaskStatusResponseWithDefaults instantiates a new AssignmentTaskStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *AssignmentTaskStatusResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *AssignmentTaskStatusResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *AssignmentTaskStatusResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetStatus

`func (o *AssignmentTaskStatusResponse) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AssignmentTaskStatusResponse) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AssignmentTaskStatusResponse) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.


### GetCreatedAt

`func (o *AssignmentTaskStatusResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AssignmentTaskStatusResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AssignmentTaskStatusResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AssignmentTaskStatusResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *AssignmentTaskStatusResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *AssignmentTaskStatusResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *AssignmentTaskStatusResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *AssignmentTaskStatusResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


