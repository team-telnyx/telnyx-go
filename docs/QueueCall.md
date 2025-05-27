# QueueCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** |  | 
**CallSessionId** | **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call | 
**CallLegId** | **string** | ID that is unique to the call and can be used to correlate webhook events | 
**CallControlId** | **string** | Unique identifier and token for controlling the call. | 
**ConnectionId** | **string** | Call Control App ID (formerly Telnyx connection ID) used in the call. | 
**From** | **string** | Number or SIP URI placing the call. | 
**To** | **string** | Destination number or SIP URI of the call. | 
**EnqueuedAt** | **string** | ISO 8601 formatted date of when the call was put in the queue | 
**WaitTimeSecs** | **int32** | The time the call has been waiting in the queue, given in seconds | 
**QueuePosition** | **int32** | Current position of the call in the queue | 
**QueueId** | **string** | Unique identifier of the queue the call is in. | 

## Methods

### NewQueueCall

`func NewQueueCall(recordType string, callSessionId string, callLegId string, callControlId string, connectionId string, from string, to string, enqueuedAt string, waitTimeSecs int32, queuePosition int32, queueId string, ) *QueueCall`

NewQueueCall instantiates a new QueueCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueCallWithDefaults

`func NewQueueCallWithDefaults() *QueueCall`

NewQueueCallWithDefaults instantiates a new QueueCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *QueueCall) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *QueueCall) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *QueueCall) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetCallSessionId

`func (o *QueueCall) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *QueueCall) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *QueueCall) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.


### GetCallLegId

`func (o *QueueCall) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *QueueCall) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *QueueCall) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.


### GetCallControlId

`func (o *QueueCall) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *QueueCall) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *QueueCall) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.


### GetConnectionId

`func (o *QueueCall) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *QueueCall) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *QueueCall) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetFrom

`func (o *QueueCall) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *QueueCall) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *QueueCall) SetFrom(v string)`

SetFrom sets From field to given value.


### GetTo

`func (o *QueueCall) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *QueueCall) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *QueueCall) SetTo(v string)`

SetTo sets To field to given value.


### GetEnqueuedAt

`func (o *QueueCall) GetEnqueuedAt() string`

GetEnqueuedAt returns the EnqueuedAt field if non-nil, zero value otherwise.

### GetEnqueuedAtOk

`func (o *QueueCall) GetEnqueuedAtOk() (*string, bool)`

GetEnqueuedAtOk returns a tuple with the EnqueuedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnqueuedAt

`func (o *QueueCall) SetEnqueuedAt(v string)`

SetEnqueuedAt sets EnqueuedAt field to given value.


### GetWaitTimeSecs

`func (o *QueueCall) GetWaitTimeSecs() int32`

GetWaitTimeSecs returns the WaitTimeSecs field if non-nil, zero value otherwise.

### GetWaitTimeSecsOk

`func (o *QueueCall) GetWaitTimeSecsOk() (*int32, bool)`

GetWaitTimeSecsOk returns a tuple with the WaitTimeSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitTimeSecs

`func (o *QueueCall) SetWaitTimeSecs(v int32)`

SetWaitTimeSecs sets WaitTimeSecs field to given value.


### GetQueuePosition

`func (o *QueueCall) GetQueuePosition() int32`

GetQueuePosition returns the QueuePosition field if non-nil, zero value otherwise.

### GetQueuePositionOk

`func (o *QueueCall) GetQueuePositionOk() (*int32, bool)`

GetQueuePositionOk returns a tuple with the QueuePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueuePosition

`func (o *QueueCall) SetQueuePosition(v int32)`

SetQueuePosition sets QueuePosition field to given value.


### GetQueueId

`func (o *QueueCall) GetQueueId() string`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *QueueCall) GetQueueIdOk() (*string, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *QueueCall) SetQueueId(v string)`

SetQueueId sets QueueId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


