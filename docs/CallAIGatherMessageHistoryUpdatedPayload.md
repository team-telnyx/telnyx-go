# CallAIGatherMessageHistoryUpdatedPayload

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

## Methods

### NewCallAIGatherMessageHistoryUpdatedPayload

`func NewCallAIGatherMessageHistoryUpdatedPayload() *CallAIGatherMessageHistoryUpdatedPayload`

NewCallAIGatherMessageHistoryUpdatedPayload instantiates a new CallAIGatherMessageHistoryUpdatedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallAIGatherMessageHistoryUpdatedPayloadWithDefaults

`func NewCallAIGatherMessageHistoryUpdatedPayloadWithDefaults() *CallAIGatherMessageHistoryUpdatedPayload`

NewCallAIGatherMessageHistoryUpdatedPayloadWithDefaults instantiates a new CallAIGatherMessageHistoryUpdatedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallControlId

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *CallAIGatherMessageHistoryUpdatedPayload) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.

### HasCallControlId

`func (o *CallAIGatherMessageHistoryUpdatedPayload) HasCallControlId() bool`

HasCallControlId returns a boolean if a field has been set.

### GetConnectionId

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CallAIGatherMessageHistoryUpdatedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CallAIGatherMessageHistoryUpdatedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallLegId

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *CallAIGatherMessageHistoryUpdatedPayload) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *CallAIGatherMessageHistoryUpdatedPayload) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *CallAIGatherMessageHistoryUpdatedPayload) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *CallAIGatherMessageHistoryUpdatedPayload) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetClientState

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *CallAIGatherMessageHistoryUpdatedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *CallAIGatherMessageHistoryUpdatedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.

### GetFrom

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallAIGatherMessageHistoryUpdatedPayload) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallAIGatherMessageHistoryUpdatedPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallAIGatherMessageHistoryUpdatedPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallAIGatherMessageHistoryUpdatedPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetMessageHistory

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetMessageHistory() []CallAIGatherEndedPayloadMessageHistoryInner`

GetMessageHistory returns the MessageHistory field if non-nil, zero value otherwise.

### GetMessageHistoryOk

`func (o *CallAIGatherMessageHistoryUpdatedPayload) GetMessageHistoryOk() (*[]CallAIGatherEndedPayloadMessageHistoryInner, bool)`

GetMessageHistoryOk returns a tuple with the MessageHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageHistory

`func (o *CallAIGatherMessageHistoryUpdatedPayload) SetMessageHistory(v []CallAIGatherEndedPayloadMessageHistoryInner)`

SetMessageHistory sets MessageHistory field to given value.

### HasMessageHistory

`func (o *CallAIGatherMessageHistoryUpdatedPayload) HasMessageHistory() bool`

HasMessageHistory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


