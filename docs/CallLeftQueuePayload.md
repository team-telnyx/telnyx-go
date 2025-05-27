# CallLeftQueuePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**ConnectionId** | Pointer to **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**Queue** | Pointer to **string** | The name of the queue | [optional] 
**QueuePosition** | Pointer to **int32** | Last position of the call in the queue. | [optional] 
**Reason** | Pointer to **string** | The reason for leaving the queue | [optional] 
**WaitTimeSecs** | Pointer to **int32** | Time call spent in the queue in seconds. | [optional] 

## Methods

### NewCallLeftQueuePayload

`func NewCallLeftQueuePayload() *CallLeftQueuePayload`

NewCallLeftQueuePayload instantiates a new CallLeftQueuePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallLeftQueuePayloadWithDefaults

`func NewCallLeftQueuePayloadWithDefaults() *CallLeftQueuePayload`

NewCallLeftQueuePayloadWithDefaults instantiates a new CallLeftQueuePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallLeftQueuePayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallLeftQueuePayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallLeftQueuePayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallLeftQueuePayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallLeftQueuePayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallLeftQueuePayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallLeftQueuePayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallLeftQueuePayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallLeftQueuePayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallLeftQueuePayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallLeftQueuePayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallLeftQueuePayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallLeftQueuePayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallLeftQueuePayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallLeftQueuePayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallLeftQueuePayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallLeftQueuePayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallLeftQueuePayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallLeftQueuePayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallLeftQueuePayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetQueue

`func (o *CallLeftQueuePayload) GetQueue() string`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *CallLeftQueuePayload) GetQueueOk() (*string, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *CallLeftQueuePayload) SetQueue(v string)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *CallLeftQueuePayload) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetQueuePosition

`func (o *CallLeftQueuePayload) GetQueuePosition() int32`

GetQueuePosition returns the QueuePosition field if non-nil, zero value otherwise.

### GetQueuePositionOk

`func (o *CallLeftQueuePayload) GetQueuePositionOk() (*int32, bool)`

GetQueuePositionOk returns a tuple with the QueuePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePosition

`func (o *CallLeftQueuePayload) SetQueuePosition(v int32)`

SetQueuePosition sets QueuePosition field to given value.

### HasQueuePosition

`func (o *CallLeftQueuePayload) HasQueuePosition() bool`

HasQueuePosition returns a boolean if a field has been set.

### GetReason

`func (o *CallLeftQueuePayload) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CallLeftQueuePayload) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CallLeftQueuePayload) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CallLeftQueuePayload) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetWaitTimeSecs

`func (o *CallLeftQueuePayload) GetWaitTimeSecs() int32`

GetWaitTimeSecs returns the WaitTimeSecs field if non-nil, zero value otherwise.

### GetWaitTimeSecsOk

`func (o *CallLeftQueuePayload) GetWaitTimeSecsOk() (*int32, bool)`

GetWaitTimeSecsOk returns a tuple with the WaitTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimeSecs

`func (o *CallLeftQueuePayload) SetWaitTimeSecs(v int32)`

SetWaitTimeSecs sets WaitTimeSecs field to given value.

### HasWaitTimeSecs

`func (o *CallLeftQueuePayload) HasWaitTimeSecs() bool`

HasWaitTimeSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


