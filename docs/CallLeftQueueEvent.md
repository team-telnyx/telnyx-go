# CallLeftQueueEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CallLeftQueue**](CallLeftQueue.md) |  | [optional] 

## Methods

### NewCallLeftQueueEvent

`func NewCallLeftQueueEvent() *CallLeftQueueEvent`

NewCallLeftQueueEvent instantiates a new CallLeftQueueEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallLeftQueueEventWithDefaults

`func NewCallLeftQueueEventWithDefaults() *CallLeftQueueEvent`

NewCallLeftQueueEventWithDefaults instantiates a new CallLeftQueueEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CallLeftQueueEvent) GetData() CallLeftQueue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CallLeftQueueEvent) GetDataOk() (*CallLeftQueue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CallLeftQueueEvent) SetData(v CallLeftQueue)`

SetData sets Data field to given value.

### HasData

`func (o *CallLeftQueueEvent) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


