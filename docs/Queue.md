# Queue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** |  | 
**Id** | **string** | Uniquely identifies the queue | 
**Name** | **string** | Name of the queue | 
**CreatedAt** | **string** | ISO 8601 formatted date of when the queue was created | 
**UpdatedAt** | **string** | ISO 8601 formatted date of when the queue was last updated | 
**CurrentSize** | **int32** | The number of calls currently in the queue | 
**MaxSize** | **int32** | The maximum number of calls allowed in the queue | 
**AverageWaitTimeSecs** | **int32** | The average time that the calls currently in the queue have spent waiting, given in seconds. | 

## Methods

### NewQueue

`func NewQueue(recordType string, id string, name string, createdAt string, updatedAt string, currentSize int32, maxSize int32, averageWaitTimeSecs int32, ) *Queue`

NewQueue instantiates a new Queue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueWithDefaults

`func NewQueueWithDefaults() *Queue`

NewQueueWithDefaults instantiates a new Queue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *Queue) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *Queue) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *Queue) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetId

`func (o *Queue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Queue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Queue) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Queue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Queue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Queue) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *Queue) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Queue) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Queue) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Queue) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Queue) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Queue) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetCurrentSize

`func (o *Queue) GetCurrentSize() int32`

GetCurrentSize returns the CurrentSize field if non-nil, zero value otherwise.

### GetCurrentSizeOk

`func (o *Queue) GetCurrentSizeOk() (*int32, bool)`

GetCurrentSizeOk returns a tuple with the CurrentSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSize

`func (o *Queue) SetCurrentSize(v int32)`

SetCurrentSize sets CurrentSize field to given value.


### GetMaxSize

`func (o *Queue) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *Queue) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *Queue) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.


### GetAverageWaitTimeSecs

`func (o *Queue) GetAverageWaitTimeSecs() int32`

GetAverageWaitTimeSecs returns the AverageWaitTimeSecs field if non-nil, zero value otherwise.

### GetAverageWaitTimeSecsOk

`func (o *Queue) GetAverageWaitTimeSecsOk() (*int32, bool)`

GetAverageWaitTimeSecsOk returns a tuple with the AverageWaitTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageWaitTimeSecs

`func (o *Queue) SetAverageWaitTimeSecs(v int32)`

SetAverageWaitTimeSecs sets AverageWaitTimeSecs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


