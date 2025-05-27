# CallAIGatherEndedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallControlId** | Pointer to **string** | Call ID used to issue commands via Call Control API. | [optional] 
**ConnectionId** | Pointer to **string** | Telnyx connection ID used in the call. | [optional] 
**CallLegId** | Pointer to **string** | ID that is unique to the call and can be used to correlate webhook events. | [optional] 
**CallSessionId** | Pointer to **string** | ID that is unique to the call session and can be used to correlate webhook events. Call session is a group of related call legs that logically belong to the same phone call, e.g. an inbound and outbound leg of a transferred call. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 
**From** | Pointer to **string** | Number or SIP URI placing the call. | [optional] 
**To** | Pointer to **string** | Destination number or SIP URI of the call. | [optional] 
**MessageHistory** | Pointer to [**[]CallAIGatherEndedPayloadMessageHistoryInner**](CallAIGatherEndedPayloadMessageHistoryInner.md) | The history of the messages exchanged during the AI gather | [optional] 
**Result** | Pointer to **map[string]interface{}** | The result of the AI gather, its type depends of the &#x60;parameters&#x60; provided in the command | [optional] 
**Status** | Pointer to **string** | Reflects how command ended. | [optional] 

## Methods

### NewCallAIGatherEndedPayload

`func NewCallAIGatherEndedPayload() *CallAIGatherEndedPayload`

NewCallAIGatherEndedPayload instantiates a new CallAIGatherEndedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallAIGatherEndedPayloadWithDefaults

`func NewCallAIGatherEndedPayloadWithDefaults() *CallAIGatherEndedPayload`

NewCallAIGatherEndedPayloadWithDefaults instantiates a new CallAIGatherEndedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallAIGatherEndedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallAIGatherEndedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallAIGatherEndedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallAIGatherEndedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallAIGatherEndedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallAIGatherEndedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallAIGatherEndedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallAIGatherEndedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallAIGatherEndedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallAIGatherEndedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallAIGatherEndedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallAIGatherEndedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallAIGatherEndedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallAIGatherEndedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallAIGatherEndedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallAIGatherEndedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallAIGatherEndedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallAIGatherEndedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallAIGatherEndedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallAIGatherEndedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetFrom

`func (o *CallAIGatherEndedPayload) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallAIGatherEndedPayload) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallAIGatherEndedPayload) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallAIGatherEndedPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CallAIGatherEndedPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallAIGatherEndedPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallAIGatherEndedPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallAIGatherEndedPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetMessageHistory

`func (o *CallAIGatherEndedPayload) GetMessageHistory() []CallAIGatherEndedPayloadMessageHistoryInner`

GetMessageHistory returns the MessageHistory field if non-nil, zero value otherwise.

### GetMessageHistoryOk

`func (o *CallAIGatherEndedPayload) GetMessageHistoryOk() (*[]CallAIGatherEndedPayloadMessageHistoryInner, bool)`

GetMessageHistoryOk returns a tuple with the MessageHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHistory

`func (o *CallAIGatherEndedPayload) SetMessageHistory(v []CallAIGatherEndedPayloadMessageHistoryInner)`

SetMessageHistory sets MessageHistory field to given value.

### HasMessageHistory

`func (o *CallAIGatherEndedPayload) HasMessageHistory() bool`

HasMessageHistory returns a boolean if a field has been set.

### GetResult

`func (o *CallAIGatherEndedPayload) GetResult() map[string]interface{}`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *CallAIGatherEndedPayload) GetResultOk() (*map[string]interface{}, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *CallAIGatherEndedPayload) SetResult(v map[string]interface{})`

SetResult sets Result field to given value.

### HasResult

`func (o *CallAIGatherEndedPayload) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetStatus

`func (o *CallAIGatherEndedPayload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CallAIGatherEndedPayload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CallAIGatherEndedPayload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CallAIGatherEndedPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


