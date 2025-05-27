# CallEnqueuedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**Queue** | Pointer to **string** | The name of the queue | [optional] 
**CurrentPosition** | Pointer to **int32** | Current position of the call in the queue. | [optional] 
**QueueAvgWaitTimeSecs** | Pointer to **int32** | Average time call spends in the queue in seconds. | [optional] 

## Methods

### NewCallEnqueuedPayload

`func NewCallEnqueuedPayload() *CallEnqueuedPayload`

NewCallEnqueuedPayload instantiates a new CallEnqueuedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallEnqueuedPayloadWithDefaults

`func NewCallEnqueuedPayloadWithDefaults() *CallEnqueuedPayload`

NewCallEnqueuedPayloadWithDefaults instantiates a new CallEnqueuedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallEnqueuedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallEnqueuedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallEnqueuedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallEnqueuedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallEnqueuedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallEnqueuedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallEnqueuedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallEnqueuedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallEnqueuedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallEnqueuedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallEnqueuedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallEnqueuedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallEnqueuedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallEnqueuedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallEnqueuedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallEnqueuedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallEnqueuedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallEnqueuedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallEnqueuedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallEnqueuedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetQueue

`func (o *CallEnqueuedPayload) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *CallEnqueuedPayload) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *CallEnqueuedPayload) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *CallEnqueuedPayload) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetCurrentPosition

`func (o *CallEnqueuedPayload) GetCurrentPosition() int32`

GetCurrentPosition returns the CurrentPosition field if non-nil, zero value otherwise.

### GetCurrentPositionOk

`func (o *CallEnqueuedPayload) GetCurrentPositionOk() (*int32, bool)`

GetCurrentPositionOk returns a tuple with the CurrentPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentPosition

`func (o *CallEnqueuedPayload) SetCurrentPosition(v int32)`

SetCurrentPosition sets CurrentPosition field to given value.

### HasCurrentPosition

`func (o *CallEnqueuedPayload) HasCurrentPosition() bool`

HasCurrentPosition returns a boolean if a field has been set.

### GetQueueAvgWaitTimeSecs

`func (o *CallEnqueuedPayload) GetQueueAvgWaitTimeSecs() int32`

GetQueueAvgWaitTimeSecs returns the QueueAvgWaitTimeSecs field if non-nil, zero value otherwise.

### GetQueueAvgWaitTimeSecsOk

`func (o *CallEnqueuedPayload) GetQueueAvgWaitTimeSecsOk() (*int32, bool)`

GetQueueAvgWaitTimeSecsOk returns a tuple with the QueueAvgWaitTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueAvgWaitTimeSecs

`func (o *CallEnqueuedPayload) SetQueueAvgWaitTimeSecs(v int32)`

SetQueueAvgWaitTimeSecs sets QueueAvgWaitTimeSecs field to given value.

### HasQueueAvgWaitTimeSecs

`func (o *CallEnqueuedPayload) HasQueueAvgWaitTimeSecs() bool`

HasQueueAvgWaitTimeSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


