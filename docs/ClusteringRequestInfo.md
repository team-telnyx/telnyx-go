# ClusteringRequestInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **string** |  | 
**Status** | [**TaskStatus**](TaskStatus.md) |  | 
**Bucket** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**FinishedAt** | **time.Time** |  | 
**MinClusterSize** | **int32** |  | 
**MinSubclusterSize** | **int32** |  | 

## Methods

### NewClusteringRequestInfo

`func NewClusteringRequestInfo(taskId string, status TaskStatus, bucket string, createdAt time.Time, finishedAt time.Time, minClusterSize int32, minSubclusterSize int32, ) *ClusteringRequestInfo`

NewClusteringRequestInfo instantiates a new ClusteringRequestInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusteringRequestInfoWithDefaults

`func NewClusteringRequestInfoWithDefaults() *ClusteringRequestInfo`

NewClusteringRequestInfoWithDefaults instantiates a new ClusteringRequestInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *ClusteringRequestInfo) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ClusteringRequestInfo) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ClusteringRequestInfo) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.


### GetStatus

`func (o *ClusteringRequestInfo) GetStatus() TaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ClusteringRequestInfo) GetStatusOk() (*TaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ClusteringRequestInfo) SetStatus(v TaskStatus)`

SetStatus sets Status field to given value.


### GetBucket

`func (o *ClusteringRequestInfo) GetBucket() string`

GetBucket returns the Bucket field if non-nil, zero value otherwise.

### GetBucketOk

`func (o *ClusteringRequestInfo) GetBucketOk() (*string, bool)`

GetBucketOk returns a tuple with the Bucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucket

`func (o *ClusteringRequestInfo) SetBucket(v string)`

SetBucket sets Bucket field to given value.


### GetCreatedAt

`func (o *ClusteringRequestInfo) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ClusteringRequestInfo) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ClusteringRequestInfo) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetFinishedAt

`func (o *ClusteringRequestInfo) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ClusteringRequestInfo) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ClusteringRequestInfo) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.


### GetMinClusterSize

`func (o *ClusteringRequestInfo) GetMinClusterSize() int32`

GetMinClusterSize returns the MinClusterSize field if non-nil, zero value otherwise.

### GetMinClusterSizeOk

`func (o *ClusteringRequestInfo) GetMinClusterSizeOk() (*int32, bool)`

GetMinClusterSizeOk returns a tuple with the MinClusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinClusterSize

`func (o *ClusteringRequestInfo) SetMinClusterSize(v int32)`

SetMinClusterSize sets MinClusterSize field to given value.


### GetMinSubclusterSize

`func (o *ClusteringRequestInfo) GetMinSubclusterSize() int32`

GetMinSubclusterSize returns the MinSubclusterSize field if non-nil, zero value otherwise.

### GetMinSubclusterSizeOk

`func (o *ClusteringRequestInfo) GetMinSubclusterSizeOk() (*int32, bool)`

GetMinSubclusterSizeOk returns a tuple with the MinSubclusterSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSubclusterSize

`func (o *ClusteringRequestInfo) SetMinSubclusterSize(v int32)`

SetMinSubclusterSize sets MinSubclusterSize field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


