# QueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**Queue**](Queue.md) |  | [optional] 

## Methods

### NewQueueResponse

`func NewQueueResponse() *QueueResponse`

NewQueueResponse instantiates a new QueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueResponseWithDefaults

`func NewQueueResponseWithDefaults() *QueueResponse`

NewQueueResponseWithDefaults instantiates a new QueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *QueueResponse) GetData() Queue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *QueueResponse) GetDataOk() (*Queue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *QueueResponse) SetData(v Queue)`

SetData sets Data field to given value.

### HasData

`func (o *QueueResponse) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


