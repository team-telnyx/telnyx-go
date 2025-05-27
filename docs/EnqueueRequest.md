# EnqueueRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueName** | **string** | The name of the queue the call should be put in. If a queue with a given name doesn&#39;t exist yet it will be created. | 
**ClientState** | Pointer to **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | [optional] 
**CommandId** | Pointer to **string** | Use this field to avoid duplicate commands. Telnyx will ignore any command with the same &#x60;command_id&#x60; for the same &#x60;call_control_id&#x60;. | [optional] 
**MaxWaitTimeSecs** | Pointer to **int32** | The number of seconds after which the call will be removed from the queue. | [optional] 
**MaxSize** | Pointer to **int32** | The maximum number of calls allowed in the queue at a given time. Can&#39;t be modified for an existing queue. | [optional] [default to 100]

## Methods

### NewEnqueueRequest

`func NewEnqueueRequest(queueName string, ) *EnqueueRequest`

NewEnqueueRequest instantiates a new EnqueueRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnqueueRequestWithDefaults

`func NewEnqueueRequestWithDefaults() *EnqueueRequest`

NewEnqueueRequestWithDefaults instantiates a new EnqueueRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueName

`func (o *EnqueueRequest) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *EnqueueRequest) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *EnqueueRequest) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.


### GetClientState

`func (o *EnqueueRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *EnqueueRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *EnqueueRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *EnqueueRequest) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetCommandId

`func (o *EnqueueRequest) GetCommandId() string`

GetCommandId returns the CommandId field if non-nil, zero value otherwise.

### GetCommandIdOk

`func (o *EnqueueRequest) GetCommandIdOk() (*string, bool)`

GetCommandIdOk returns a tuple with the CommandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandId

`func (o *EnqueueRequest) SetCommandId(v string)`

SetCommandId sets CommandId field to given value.

### HasCommandId

`func (o *EnqueueRequest) HasCommandId() bool`

HasCommandId returns a boolean if a field has been set.

### GetMaxWaitTimeSecs

`func (o *EnqueueRequest) GetMaxWaitTimeSecs() int32`

GetMaxWaitTimeSecs returns the MaxWaitTimeSecs field if non-nil, zero value otherwise.

### GetMaxWaitTimeSecsOk

`func (o *EnqueueRequest) GetMaxWaitTimeSecsOk() (*int32, bool)`

GetMaxWaitTimeSecsOk returns a tuple with the MaxWaitTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxWaitTimeSecs

`func (o *EnqueueRequest) SetMaxWaitTimeSecs(v int32)`

SetMaxWaitTimeSecs sets MaxWaitTimeSecs field to given value.

### HasMaxWaitTimeSecs

`func (o *EnqueueRequest) HasMaxWaitTimeSecs() bool`

HasMaxWaitTimeSecs returns a boolean if a field has been set.

### GetMaxSize

`func (o *EnqueueRequest) GetMaxSize() int32`

GetMaxSize returns the MaxSize field if non-nil, zero value otherwise.

### GetMaxSizeOk

`func (o *EnqueueRequest) GetMaxSizeOk() (*int32, bool)`

GetMaxSizeOk returns a tuple with the MaxSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSize

`func (o *EnqueueRequest) SetMaxSize(v int32)`

SetMaxSize sets MaxSize field to given value.

### HasMaxSize

`func (o *EnqueueRequest) HasMaxSize() bool`

HasMaxSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


